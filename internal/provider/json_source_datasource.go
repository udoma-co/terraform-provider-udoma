package provider

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bluekeyes/go-gitdiff/gitdiff"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/jsonpatch"
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/tf"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource = &jsonSource{}
)

func NewJsonSource() datasource.DataSource {
	return &jsonSource{}
}

// jsonSource defines the data source implementation.
type jsonSource struct{}

// jsonSourceModel describes the data source data model.
type jsonSourceModel struct {
	Source  types.String       `tfsdk:"source"`
	Patches []string           `tfsdk:"patches"`
	Content tf.JsonObjectValue `tfsdk:"content"`
}

func (r *jsonSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_json_source"
}

func (r *jsonSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "Resource represents a JSON source file that may contain references to other files",

		Attributes: map[string]schema.Attribute{
			"source": schema.StringAttribute{
				Required:    true,
				Description: "The file location to read the JSON from",
			},
			"patches": schema.ListAttribute{
				Optional: true,
				Description: "An optional list of patches to apply to the original JSON content. These patches can " +
					"be either JSON patches or unified diff patches. If a unified diff patch is provided, it will be " +
					"applied to the content of the source file before any JSON patches are applied.",
				ElementType: types.StringType,
			},
			"content": schema.StringAttribute{
				CustomType:  tf.JsonObjectType{},
				Computed:    true,
				Description: "The content read from the file, with other files referenced by the source attribute embedded",
			},
		},
	}
}

// Read does not need to perform any operations as the state in ReadResourceResponse is already populated.
func (r *jsonSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	// Retrieve values from config
	var config jsonSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	jsonPatches := []jsonpatch.Patch{}
	diffPatches := []*gitdiff.File{}

	for _, patch := range config.Patches {

		patchContents, err := readMergedFile(patch, diffPatches)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Reading Patch File",
				"Could not read file, unexpected error: "+err.Error(),
			)
			return
		}

		// patch can be either a JSON patch or a unified diff patch. We can
		// distinguish them by checking if the first character is a '[' for
		// JSON patch or a '-' for unified diff patch.
		if len(patchContents) == 0 {
			resp.Diagnostics.AddError(
				"Error Reading Patch File",
				"Patch file is empty",
			)
			return
		}
		if patchContents[0] == '[' {
			// JSON patch

			patch, err := jsonpatch.DecodePatch(patchContents)
			if err != nil {
				resp.Diagnostics.AddError(
					"Error Decoding Json Patch",
					"Could not decode json patch, unexpected error: "+err.Error(),
				)
				return
			}

			jsonPatches = append(jsonPatches, patch)

		} else if patchContents[0] == '-' {
			// Unified diff patch

			files, _, err := gitdiff.Parse(bytes.NewReader(patchContents))
			if err != nil {
				resp.Diagnostics.AddError(
					"Error Parsing Diff Patch",
					"Could not parse diff patch, unexpected error: "+err.Error(),
				)
				return
			}

			resp.Diagnostics.AddWarning(
				"Diff Patch Warning",
				"Running diff for file "+files[0].OldName,
			)

			diffPatches = append(diffPatches, files...)

		} else {
			resp.Diagnostics.AddError(
				"Error Reading Patch File",
				"Patch file does not start with a valid character, expected '[' for JSON patch or '-' for unified diff patch",
			)
			return
		}

	}

	fileContents, err := readMergedFile(config.Source.ValueString(), diffPatches)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Json Source",
			"Could not read source file, unexpected error: "+err.Error(),
		)
		return
	}

	for _, patch := range jsonPatches {
		fileContents, err = patch.Apply(fileContents)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Applying Json Patch",
				"Could not apply json patch, unexpected error: "+err.Error(),
			)
			return
		}
	}

	for _, patch := range diffPatches {

		if !strings.HasSuffix(config.Source.String(), patch.OldName) {
			// skip patches that do not match the source file
			continue
		}

		var output bytes.Buffer
		if err := gitdiff.Apply(&output, bytes.NewReader(fileContents), patch); err != nil {
			resp.Diagnostics.AddError(
				"Error Applying Diff Patch",
				"Could not apply diff patch, unexpected error: "+err.Error(),
			)
			return
		}

		fileContents = output.Bytes()
	}

	if len(config.Patches) > 0 {
		// normalize the content to remove whitespace drift
		var data any
		if err := json.Unmarshal(fileContents, &data); err != nil {
			resp.Diagnostics.AddError(
				"Error Normalizing Json Content",
				"Could not normalize content, unexpected error: "+err.Error(),
			)
			return
		}
		fileContents, err = json.Marshal(data)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Normalizing Json Content",
				"Could not normalize content, unexpected error: "+err.Error(),
			)
			return
		}
	}

	config.Content = tf.NewJsonObjectValue(string(fileContents))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// readMergedFile reads the file with the given name and embeds any nested
// files, such as JSON, PNG, JPG, and JS files, into the content. For JS
// files, it applies any diff patches that were provided that match the
// file name.
func readMergedFile(fileName string, diffPatches []*gitdiff.File) ([]byte, error) {

	dirName := filepath.Dir(fileName)

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	data, err = embedFiles(data, dirName, diffPatches)
	if err != nil {
		return nil, fmt.Errorf("could not embed nested files: %w", err)
	}

	return data, nil
}

// embedFiles finds occurrences of "src:.*.json", "src:.*.png", "src:.*.jpg",
// and "src:.*.js" in the data and replaces them with the content of the
// respective files. For JSON files, it embeds the content as a string, for
// PNG and JPG  files, it embeds the content as base64 encoded strings, and
// for JS files, it embeds the content as a single line string with escaped
// characters, after applying all patches that match the file name. JSON files
// can also be embedded as a string, rather than a nested JSON structure by
// adding "|string" at the end of the file name. It also recursively embeds
// any nested files found within the JSON content. It returns the modified
// data or an error if any file could not be read.
func embedFiles(data []byte, dirName string, diffPatches []*gitdiff.File) ([]byte, error) {

	// find occurrences of "src:.*.json" and replace with the content of the file
	pattern := regexp.MustCompile(`"src:.*\.json(\|string)?"`)
	for {
		finding := pattern.FindIndex(data)
		if finding == nil {
			break
		}
		start := finding[0]
		end := finding[1]

		fileName := string(data[start+5 : end-1])

		embedAsStr := false
		if strings.HasSuffix(fileName, "|string") {
			embedAsStr = true
			fileName = fileName[:len(fileName)-7]
		}

		filePath := filepath.Join(dirName, fileName)
		fileData, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("could not read file: %w", err)
		}

		fileData, err = embedFiles(fileData, dirName, diffPatches)
		if err != nil {
			return nil, fmt.Errorf("could not embed nested files: %w", err)
		}

		if embedAsStr {
			fileData = embedString(fileData, true)
		}

		data = append(data[:start], append(fileData, data[end:]...)...)
	}

	// find occurrences of "src:.*.png" and replace with
	// the base64 encoded content of the file
	pattern = regexp.MustCompile(`"src:.*\.(png|jpg)"`)
	for {
		finding := pattern.FindIndex(data)
		if finding == nil {
			break
		}
		start := finding[0]
		end := finding[1]
		fileName := string(data[start+5 : end-1])
		filePath := filepath.Join(dirName, fileName)
		fileData, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("could not read file: %w", err)
		}

		base64Data := base64.StdEncoding.EncodeToString(fileData)

		data = append(data[:start+1], append([]byte(base64Data), data[end-1:]...)...)
	}

	// do the same for JS files, but make sure content is escaped and put in one line
	pattern = regexp.MustCompile(`"src:.*\.js"`)
	for {
		finding := pattern.FindIndex(data)
		if finding == nil {
			break
		}
		start := finding[0]
		end := finding[1]
		fileName := string(data[start+5 : end-1])
		filePath := filepath.Join(dirName, fileName)
		fileData, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("could not read file: %w", err)
		}

		// apply any diff patches that match the file name
		for _, patch := range diffPatches {
			if !strings.HasSuffix(fileName, patch.OldName) {
				// skip patches that do not match the file name
				continue
			}

			var output bytes.Buffer
			if err := gitdiff.Apply(&output, bytes.NewReader(fileData), patch); err != nil {
				return nil, fmt.Errorf("could not apply diff patch to file %s: %w", fileName, err)
			}
			fileData = output.Bytes()
		}

		sanitized := embedString(fileData, false)

		data = append(data[:start], append(sanitized, data[end:]...)...)
	}

	return data, nil
}

// embedString takes a byte slice, replaces newlines with \n, escapes quotes,
// and optionally replaces tabs with \t. It returns the modified byte slice
func embedString(in []byte, replaceTabs bool) []byte {

	ret := string(in)
	ret = regexp.MustCompile(`\r?\n`).ReplaceAllString(ret, `\n`)
	ret = regexp.MustCompile(`"`).ReplaceAllString(ret, `\"`)
	if replaceTabs {
		ret = regexp.MustCompile(`\\t`).ReplaceAllString(ret, `\\t`)
	}
	ret = `"` + ret + `"`

	return []byte(ret)
}
