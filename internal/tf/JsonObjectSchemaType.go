package tf

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Ensure the implementation satisfies the expected interfaces
var _ basetypes.StringTypable = JsonObjectType{}

type JsonObjectType struct {
	basetypes.StringType
	// ... potentially other fields ...
}

func (t JsonObjectType) Equal(o attr.Type) bool {
	other, ok := o.(JsonObjectType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

func (t JsonObjectType) String() string {
	return "CustomStringType"
}

func (t JsonObjectType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	// CustomStringValue defined in the value type section
	value := JsonObjectValue{
		StringValue: in,
	}

	return value, nil
}

func (t JsonObjectType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.StringType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	stringValue, ok := attrValue.(basetypes.StringValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	stringValuable, diags := t.ValueFromString(ctx, stringValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting StringValue to StringValuable: %v", diags)
	}

	return stringValuable, nil
}

func (t JsonObjectType) ValueType(ctx context.Context) attr.Value {
	// CustomStringValue defined in the value type section
	return JsonObjectValue{}
}

// CustomStringType defined in the schema type section
func (t JsonObjectType) Validate(ctx context.Context, value tftypes.Value, valuePath path.Path) diag.Diagnostics {
	if value.IsNull() || !value.IsKnown() {
		return nil
	}

	var diags diag.Diagnostics
	var valueString string

	if err := value.As(&valueString); err != nil {
		diags.AddAttributeError(
			valuePath,
			"Invalid Terraform Value",
			"An unexpected error occurred while attempting to convert a Terraform value to a string. "+
				"This generally is an issue with the provider schema implementation. "+
				"Please contact the provider developers.\n\n"+
				"Path: "+valuePath.String()+"\n"+
				"Error: "+err.Error(),
		)

		return diags
	}

	var target any
	if err := json.Unmarshal([]byte(valueString), &target); err != nil {
		diags.AddAttributeError(
			valuePath,
			"Invalid JSON String Value",
			"An unexpected error occurred while converting a string value that was expected to be JSON. "+
				"Please ensure the value is valid JSON.\n\n"+
				"Path: "+valuePath.String()+"\n"+
				"Given Value: "+valueString+"\n"+
				"Error: "+err.Error(),
		)

		return diags
	}

	return diags
}
