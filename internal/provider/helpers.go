package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	api "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
)

// omittableBooleanValue returns a new value for a boolean field that can be
// omitted. That is, if the new value is nil and the old value is false, the
// returned value is false. Otherwise, the returned value is the new value.
// This is useful for fields that are optional and default to false, i.e.
// during planing they will have a value of false, but after apply the API
// might omit the field if it is false.
func omittableBooleanValue(newValue *bool, oldValue basetypes.BoolValue) basetypes.BoolValue {
	if newValue == nil && !oldValue.IsNull() && !oldValue.ValueBool() {
		// omitted false value
		return types.BoolValue(false)
	}
	return types.BoolPointerValue(newValue)
}

// omittableStringValue returns a new value for a string field that can be
// omitted. That is, if the new value is nil and the old value is "", the
// returned value is "". Otherwise, the returned value is the new value.
// This is useful for fields that are optional and default to "", i.e.
// during planing they will have a value of "", but after apply the API
// might omit the field if it is "".
func omittableStringValue(newValue *string, oldValue basetypes.StringValue) basetypes.StringValue {
	if newValue == nil && !oldValue.IsNull() && oldValue.ValueString() == "" {
		// omitted empty value
		return types.StringValue("")
	}
	return types.StringPointerValue(newValue)
}

func idp32(in *int32) int32 {
	if in == nil {
		return 0
	}
	return *in
}

func getApiErrorMessage(err error) string {

	if err == nil {
		return ""
	}

	apiErr, ok := err.(*api.GenericOpenAPIError)
	if ok {
		return fmt.Sprintf("%s: %s", err.Error(), string(apiErr.Body()))
	}

	return err.Error()
}

func stringMapToValueMap(in map[string]string) map[string]attr.Value {
	ret := map[string]attr.Value{}
	for k, v := range in {
		ret[k] = types.StringValue(v)
	}
	return ret
}

func modelMapToStringMap(in basetypes.MapValue) map[string]string {

	if in.IsNull() || in.IsUnknown() {
		return nil
	}
	ret := make(map[string]string, len(in.Elements()))
	for k, v := range in.Elements() {
		ret[k] = (v.(types.String)).ValueString()
	}

	return ret
}

func modelListToStringSlice(in basetypes.ListValue) []string {

	if in.IsNull() || in.IsUnknown() {
		return nil
	}
	ret := make([]string, len(in.Elements()))
	for i := range in.Elements() {
		ret[i] = (in.Elements()[i].(types.String)).ValueString()
	}

	return ret
}

type EnumType interface {
	api.CaseStatusEnum | api.CaseActionEnum |
		api.UserTypeEnum | api.CaseFeedbackModeEnum
}

func modelListToEnumSlice[T EnumType](in basetypes.ListValue) []T {

	if in.IsNull() || in.IsUnknown() {
		return nil
	}
	ret := make([]T, len(in.Elements()))
	for i := range in.Elements() {
		strV := (in.Elements()[i].(types.String)).ValueString()
		ret[i] = T(strV)
	}

	return ret
}

func modelListToInt32Slice(in basetypes.ListValue) []int32 {

	if in.IsNull() || in.IsUnknown() {
		return nil
	}
	ret := make([]int32, len(in.Elements()))
	for i := range in.Elements() {
		i64 := (in.Elements()[i].(types.Int64)).ValueInt64()
		ret[i] = int32(i64)
	}

	return ret
}

func stringSliceToValueList(in []string) []attr.Value {
	ret := make([]attr.Value, len(in))
	for i := range in {
		ret[i] = types.StringValue(in[i])
	}
	return ret
}

func enumSliceToValueList[T EnumType](in []T) []attr.Value {
	ret := make([]attr.Value, len(in))
	for i := range in {
		ret[i] = types.StringValue(string(in[i]))
	}
	return ret
}

func int32SliceToValueList(in []int32) []attr.Value {
	ret := make([]attr.Value, len(in))
	for i := range in {
		ret[i] = types.Int64Value(int64(in[i]))
	}
	return ret
}

func i64ToI32Ptr(val int64) *int32 {
	res := int32(val)
	return &res
}

// func isEmptyBool(s *bool) bool {
// 	return s == nil || !*s
// }

// func isEmptyEnum[T EnumType](e *T) bool {
// 	if e == nil {
// 		return true
// 	}
// 	strV := string(*e)
// 	return strV == ""
// }
