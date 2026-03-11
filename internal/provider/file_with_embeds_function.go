// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	fileembed "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/file_embed"
)

var _ function.Function = &FileWithEmbedsFunction{}

type FileWithEmbedsFunction struct{}

func NewFileWithEmbedsFunction() function.Function {
	return &FileWithEmbedsFunction{}
}

func (f *FileWithEmbedsFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "file_with_embeds"
}

func (f *FileWithEmbedsFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Load a file with embedded content and return its content as a string",
		Description: "Given a file location as a string, return its content as a string, with any embedded files processed.",

		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "source",
				Description: "The source file location as a string.",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *FileWithEmbedsFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {

	var fileName string
	resp.Error = req.Arguments.Get(ctx, &fileName)
	if resp.Error != nil {
		return
	}

	// load file content with any embedded files processed
	data, err := fileembed.ReadFileWithEmbeds(fileName, nil)
	if err != nil {
		resp.Error = function.NewArgumentFuncError(0, fmt.Sprintf("Error reading file with embeds: %v", err))
		return
	}

	resp.Error = resp.Result.Set(ctx, types.StringValue(string(data)))
}
