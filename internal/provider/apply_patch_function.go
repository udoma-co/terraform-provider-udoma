// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/bluekeyes/go-gitdiff/gitdiff"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ function.Function = &ApplyPatchFunction{}

type ApplyPatchFunction struct{}

func NewApplyPatchFunction() function.Function {
	return &ApplyPatchFunction{}
}

func (f *ApplyPatchFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "apply_patch"
}

func (f *ApplyPatchFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Apply a git diff to a file and return the result",
		Description: "Given a file as a string, apply a git diff to it and return the result as a string.",

		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "source",
				Description: "The source file to apply the diff to, as a string.",
			},
			function.StringParameter{
				Name:        "patch",
				Description: "The patch file (git diff structure) to apply to the source, as a string.",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *ApplyPatchFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {

	var source, patch string
	resp.Error = req.Arguments.Get(ctx, &source, &patch)
	if resp.Error != nil {
		return
	}

	if patch == "" {
		// no diff provided, return the source as is
		resp.Result.Set(ctx, types.StringValue(source))
		return
	}

	diffPatches, _, err := gitdiff.Parse(bytes.NewReader([]byte(patch)))
	if err != nil {
		resp.Error = function.NewArgumentFuncError(1, fmt.Sprintf("Error parsing patch file: %v", err.Error()))
		return
	}

	sourceData := []byte(source)
	// apply any diff patches that match the file name
	for _, patch := range diffPatches {

		var output bytes.Buffer
		if err := gitdiff.Apply(&output, bytes.NewReader(sourceData), patch); err != nil {
			resp.Error = function.NewFuncError(fmt.Sprintf("Could not apply diff patch to file: %v", err.Error()))
			return
		}
		sourceData = output.Bytes()
	}

	res := strings.TrimSpace(string(sourceData))
	resp.Error = resp.Result.Set(ctx, &res)
}
