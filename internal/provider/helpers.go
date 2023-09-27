package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	api "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
)

// func isResourceTimeOut(err error) bool {
// 	if err == nil {
// 		return false
// 	}

// 	timeOut, ok := err.(*resource.TimeoutError)
// 	if !ok {
// 		return false
// 	}

// 	return timeOut.LastError == nil
// }

// func sp(in string) *string {
// 	return &in
// }

func sdp(in *string) string {
	if in == nil {
		return ""
	}
	return *in
}

// func bp(in bool) *bool {
// 	return &in
// }

func bdp(in *bool) bool {
	if in == nil {
		return false
	}
	return *in
}

func idp(in *int64) int64 {
	if in == nil {
		return 0
	}
	return *in
}

// func anyToMapStringString(v any) map[string]string {
// 	t := v.(map[string]any)
// 	nm := map[string]string{}
// 	for k, v := range t {
// 		nm[k] = v.(string)
// 	}
// 	return nm
// }

func getApiErrorMessage(err error) string {

	if err == nil {
		return ""
	}

	apiErr, ok := err.(*api.GenericOpenAPIError)
	if ok {
		return string(apiErr.Body())
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

func modelMapToStringMap(in basetypes.MapValue) *map[string]string {

	if in.IsNull() || in.IsUnknown() {
		return nil
	}
	ret := make(map[string]string, len(in.Elements()))
	for k, v := range in.Elements() {
		ret[k] = (v.(types.String)).ValueString()
	}

	return &ret
}
