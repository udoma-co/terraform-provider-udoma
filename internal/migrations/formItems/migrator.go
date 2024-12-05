package v1

import (
	"encoding/json"
	"fmt"

	v1 "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
)

// Migrate converts the previous version CustomForm to the latest version. The
// difference is the structure of the FormInputs that now have a different type
// of Items. The other fields remain the same and are simply copied over via
// JSON marshal/unmarshal.
func (o *CustomForm) Migrate() (*v1.CustomForm, error) {
	ret := &v1.CustomForm{
		Layout:      []v1.FormItem{},
		Groups:      []v1.FormGroup{},
		Inputs:      []v1.FormInput{},
		Validations: []v1.FormValidation{},
	}

	if err := TypeChangeViaJson(o.Layout, &ret.Layout); err != nil {
		return nil, fmt.Errorf("could not convert layout: %w", err)
	}

	if err := TypeChangeViaJson(o.Groups, &ret.Groups); err != nil {
		return nil, fmt.Errorf("could not convert groups: %w", err)
	}

	if err := TypeChangeViaJson(o.Validations, &ret.Validations); err != nil {
		return nil, fmt.Errorf("could not convert validations: %w", err)
	}

	for _, input := range o.Inputs {
		ret.Inputs = append(ret.Inputs, input.migrate())
	}

	return ret, nil
}

func (o *FormInput) migrate() v1.FormInput {
	ret := v1.FormInput{
		Id:               o.Id,
		Target:           o.Target,
		Type:             v1.FormInputType(o.Type),
		Attributes:       o.Attributes,
		Label:            o.Label,
		ViewLabel:        o.ViewLabel,
		Placeholder:      o.Placeholder,
		DefaultValue:     o.DefaultValue,
		Required:         o.Required,
		Ephemeral:        o.Ephemeral,
		PropagateChanges: o.PropagateChanges,
	}

	if o.Items != nil {
		ret.Items = make([]v1.InputItem, len(o.Items))
		for i, item := range o.Items {
			ret.Items[i] = v1.InputItem{
				Id: item,
				Label: map[string]string{
					"de": item,
				},
			}
		}
	}

	return ret
}

// TypeChangeViaJson converts the provided input to the provided
// output via json. This is currently used when convering public
// to regular API types, or vice versa.
func TypeChangeViaJson(input any, output any) error {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return fmt.Errorf("could not marshal input: %w", err)
	}

	if err = json.Unmarshal(jsonBytes, output); err != nil {
		return fmt.Errorf("could not unmarshal input: %w", err)
	}

	return nil
}
