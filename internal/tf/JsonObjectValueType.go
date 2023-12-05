package tf

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure the implementation satisfies the expected interfaces
var _ basetypes.StringValuable = JsonObjectValue{}

type JsonObjectValue struct {
	basetypes.StringValue
}

// func (v JsonObjectValue) Equal(o attr.Value) bool {
// 	other, ok := o.(JsonObjectValue)

// 	if !ok {
// 		return false
// 	}

// 	result, err := v.jsonEqual(&other)
// 	if err != nil {
// 		return false
// 	}

// 	return result
// }

func (v JsonObjectValue) Type(ctx context.Context) attr.Type {
	// JsonObjectType defined in the schema type section
	return JsonObjectType{}
}

// JsonObjectValue defined in the value type section
// Ensure the implementation satisfies the expected interfaces
var _ basetypes.StringValuableWithSemanticEquals = JsonObjectValue{}

func (v JsonObjectValue) StringSemanticEquals(ctx context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	// The framework should always pass the correct value type, but always check
	newValue, ok := newValuable.(JsonObjectValue)

	if !ok {
		diags.AddError(
			"Semantic Equality Check Error",
			"An unexpected value type was received while performing semantic equality checks. "+
				"Please report this to the provider developers.\n\n"+
				"Expected Value Type: "+fmt.Sprintf("%T", v)+"\n"+
				"Got Value Type: "+fmt.Sprintf("%T", newValuable),
		)

		return false, diags
	}

	if v.IsNull() && newValuable.IsNull() {
		return true, nil
	}

	if v.IsUnknown() && newValuable.IsUnknown() {
		return true, nil
	}

	if (!v.IsNull() && newValuable.IsNull()) ||
		(v.IsNull() && !newValuable.IsNull()) {
		return false, nil
	}

	if (!v.IsUnknown() && newValuable.IsUnknown()) ||
		(v.IsUnknown() && !newValuable.IsUnknown()) {
		return false, nil
	}

	result, err := v.jsonEqual(&newValue)

	if err != nil {
		diags.AddError(
			"Semantic Equality Check Error",
			"An unexpected error occurred while performing semantic equality checks. "+
				"Please report this to the provider developers.\n\n"+
				"Error: "+err.Error(),
		)

		return false, diags
	}

	return result, diags

}

func (v *JsonObjectValue) jsonEqual(otherValue *JsonObjectValue) (bool, error) {

	if otherValue == nil {
		return false, nil
	}

	oldValue := strings.Trim(v.ValueString(), `"`)
	newValue := strings.Trim(otherValue.ValueString(), `"`)

	if oldValue == newValue {
		return true, nil
	}

	var oldForm, newForm any

	if err := json.Unmarshal([]byte(oldValue), &oldForm); err != nil {
		return false, fmt.Errorf("failed to unmarshal old value: %w", err)
	}

	if err := json.Unmarshal([]byte(newValue), &newForm); err != nil {
		return false, fmt.Errorf("failed to unmarshal new value: %w", err)
	}

	if diff := cmp.Diff(&oldForm, &newForm); diff != "" {
		return false, nil
	}

	return true, nil
}

// NewJsonObjectNull creates a JsonObjectValue with a null value. Determine whether the value is null via IsNull method.
func NewJsonObjectNull() JsonObjectValue {
	return JsonObjectValue{
		StringValue: basetypes.NewStringNull(),
	}
}

// NewJsonObjectUnknown creates a JsonObjectValue with an unknown value. Determine whether the value is unknown via IsUnknown method.
func NewJsonObjectUnknown() JsonObjectValue {
	return JsonObjectValue{
		StringValue: basetypes.NewStringUnknown(),
	}
}

// NewJsonObjectValue creates a JsonObjectValue with a known value. Access the value via ValueString method.
func NewJsonObjectValue(value string) JsonObjectValue {
	return JsonObjectValue{
		StringValue: basetypes.NewStringValue(value),
	}
}

// NewJsonObjectPointerValue creates a JsonObjectValue with a null value if nil or a known value. Access the value via ValueStringPointer method.
func NewJsonObjectPointerValue(value *string) JsonObjectValue {
	return JsonObjectValue{
		StringValue: basetypes.NewStringPointerValue(value),
	}
}

// Unmarshal calls (encoding/json).Unmarshal with the StringValue and `target` input. A null or unknown value will produce an error diagnostic.
// See encoding/json docs for more on usage: https://pkg.go.dev/encoding/json#Unmarshal
func (v JsonObjectValue) Unmarshal(target any) diag.Diagnostics {
	var diags diag.Diagnostics

	if v.IsNull() {
		diags.Append(diag.NewErrorDiagnostic("JSON Object Unmarshal Error", "json string value is null"))
		return diags
	}

	if v.IsUnknown() {
		diags.Append(diag.NewErrorDiagnostic("JSON Object Unmarshal Error", "json string value is unknown"))
		return diags
	}

	if v.ValueString() == "" || v.ValueString() == "{}" || v.ValueString() == "[]" {
		return diags
	}

	err := json.Unmarshal([]byte(v.ValueString()), target)
	if err != nil {
		diags.Append(diag.NewErrorDiagnostic("JSON Object Unmarshal Error", err.Error()))
	}

	return diags
}
