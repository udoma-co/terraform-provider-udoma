package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"

	"github.com/bluekeyes/go-gitdiff/gitdiff"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	fileembed "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/file_embed"
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

		patchContents, err := fileembed.ReadFileWithEmbeds(patch, diffPatches)
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

			diffPatches = append(diffPatches, files...)

		} else {
			resp.Diagnostics.AddError(
				"Error Reading Patch File",
				"Patch file does not start with a valid character, expected '[' for JSON patch or '-' for unified diff patch",
			)
			return
		}

	}

	fileContents, err := fileembed.ReadFileWithEmbeds(config.Source.ValueString(), diffPatches)
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
