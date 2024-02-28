package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.String = SHA256Modifier{}

type SHA256Modifier struct {
	fileSource string
}

// Description implements planmodifier.String.
func (SHA256Modifier) Description(context.Context) string {
	return "This is a SHA256 modifier, that would keep the SHA256 element updated based on file path."
}

// MarkdownDescription implements planmodifier.String.
func (SHA256Modifier) MarkdownDescription(context.Context) string {
	return "This is a SHA256 modifier, that would keep the SHA256 element updated based on file path."
}

// PlanModifyString implements planmodifier.String.
func (m SHA256Modifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	var source types.String
	req.Plan.GetAttribute(ctx, path.Root(m.fileSource), &source)

	f, err := os.Open(source.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to open file", err.Error())
	}

	sha, err := getFileSHA256(f)
	if err != nil {
		resp.Diagnostics.AddError("Failed to checksum the file", err.Error())
	}

	if req.PlanValue.ValueString() != sha {
		resp.PlanValue = types.StringValue(sha)
		resp.RequiresReplace = true
	}
}
