package provider

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

	fileContents, err := readMergedFile(config.Source.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Json Source",
			"Could not read file, unexpected error: "+err.Error(),
		)
		return
	}

	config.Content = tf.NewJsonObjectValue(string(fileContents))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func readMergedFile(fileName string) ([]byte, error) {

	dirName := filepath.Dir(fileName)

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	data, err = embedFiles(data, dirName)
	if err != nil {
		return nil, fmt.Errorf("could not embed nested files %w", err)
	}

	return data, nil
}

func embedFiles(data []byte, dirName string) ([]byte, error) {

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
			return nil, fmt.Errorf("could not read file %w", err)
		}

		fileData, err = embedFiles(fileData, dirName)
		if err != nil {
			return nil, fmt.Errorf("could not embed nested files %w", err)
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
			return nil, fmt.Errorf("could not read file %w", err)
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
			return nil, fmt.Errorf("could not read file %w", err)
		}

		sanitized := embedString(fileData, false)

		data = append(data[:start], append(sanitized, data[end:]...)...)
	}

	return data, nil
}

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
