---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "udoma_property_handover_template Resource - terraform-provider-udoma"
subcategory: ""
description: |-
  
---

# udoma_property_handover_template (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `inputs` (Attributes) A custom form to collect data with (see [below for nested schema](#nestedatt--inputs))
- `name` (String) The name of the property handover template.

### Optional

- `description` (String) The description of the property handover template.
- `version` (Number) The version of the property handover template.

### Read-Only

- `id` (String) The ID of the property handover template.

<a id="nestedatt--inputs"></a>
### Nested Schema for `inputs`

Required:

- `inputs` (Attributes List) The inputs that will be displayed to the user (see [below for nested schema](#nestedatt--inputs--inputs))
- `layout` (Attributes List) The layout of the form, which groups and inputs will be displayed (see [below for nested schema](#nestedatt--inputs--layout))

Optional:

- `groups` (Attributes List) The groups of inputs that will be displayed to the user (see [below for nested schema](#nestedatt--inputs--groups))
- `validation` (Attributes List) The validations that will be performed on the data provided by the user (see [below for nested schema](#nestedatt--inputs--validation))

<a id="nestedatt--inputs--inputs"></a>
### Nested Schema for `inputs.inputs`

Required:

- `id` (String) The ID of the input field, used to identify it within the form. If no target value is set, this will also be how the value is stored and later accessed in the collected data.
- `type` (String) The type of data that will be retrieved by that input. This also controls how the input will be displayed to the user.

Optional:

- `attributes` (Map of String) Optional set of settings for the input field, such as min/max values. Those will vary depending on the type.
- `default_value` (String) Optional default value for the input field (as a JSON string)
- `display_condition` (Attributes) Optional condition that must be met for the input to be displayed. If the condition is not met, the input will be hidden. This overrides the required attribute. (see [below for nested schema](#nestedatt--inputs--inputs--display_condition))
- `ephemeral` (Boolean) If true, the value of the input will not be persisted. This is useful for checkboxes for accepting terms and conditions, etc.
- `items` (Attributes List) Only used when the type is select or multi select. This is a list of values that the user can choose from. (see [below for nested schema](#nestedatt--inputs--inputs--items))
- `label` (Map of String) The label for the input field as a language dictionary (e.g. 'What is your name?')
- `placeholder` (Map of String) An optional placeholder to use for the input field. If not set, a default one will be used based on the label of the input.
- `propagate_changes` (Boolean) If true, changes to the input will be propagated to event listeners for the custom form.
- `readonly` (Boolean) If true, the user will not be able to change the value of this input
- `required` (Boolean) If true, the user will be required to provide a valid value for this input
- `target` (String) Optional attribute name to use when exporting the value of this input. If not set, the ID will be used.
- `tooltip` (Map of String) Optional tooltip to be displayed next to the label (hover or click on icon). This is useful for providing additional information to the user.
- `view_label` (Map of String) Optional label to be displayed when viewing the input field (e.g. 'Name'). This is useful when the label is not suitable to be displayed in a view.

<a id="nestedatt--inputs--inputs--display_condition"></a>
### Nested Schema for `inputs.inputs.display_condition`

Required:

- `operand` (String) the operand of the condition
- `source` (String) the source field, which will be checked
- `value` (String) the value of the condition, as serialized JSON


<a id="nestedatt--inputs--inputs--items"></a>
### Nested Schema for `inputs.inputs.items`

Required:

- `id` (String) The technical value that will be stored in the collected data
- `label` (Map of String) The label to be displayed to the user as a language dictionary



<a id="nestedatt--inputs--layout"></a>
### Nested Schema for `inputs.layout`

Required:

- `ref_id` (String) The ID of the entity that will be referenced
- `ref_type` (String) The type of the entity that will be referenced


<a id="nestedatt--inputs--groups"></a>
### Nested Schema for `inputs.groups`

Required:

- `id` (String) The ID of the group
- `items` (Attributes List) the IDs of the inputs that will be displayed in the group (see [below for nested schema](#nestedatt--inputs--groups--items))
- `type` (String) The type of the group

Optional:

- `bottom_divider` (Boolean) if true, a divider will be displayed below the group
- `icon` (String) the icon to display (only used for pages)
- `info` (Map of String) The info of the group
- `label` (Map of String) The label of the group
- `min_size` (Number) the minimum number of items that must be submitted in the group (only used for repeat groups)
- `nested_display` (Boolean) If true, the nested group will be displayed in a nested UI (only works on mobile and tablet). This is useful for more complex group that require more space to be displayed. The group will be displayed in a separate screen, and the user will be able to navigate back and forth between the group and the main form.
- `subtitle` (List of String) Optional subtitle for repeat groups. This will be computed dynamically based on the data provided by the user and will be displayed as a subtitle in the group accordion. The values are interpreted as JSON paths based on the data for the group.
- `target` (String) the attribute name to use when exporting the result of this group (only used for repeat groups)
- `top_divider` (Boolean) if true, a divider will be displayed above the group

<a id="nestedatt--inputs--groups--items"></a>
### Nested Schema for `inputs.groups.items`

Required:

- `ref_id` (String) The ID of the entity that will be referenced
- `ref_type` (String) The type of the entity that will be referenced



<a id="nestedatt--inputs--validation"></a>
### Nested Schema for `inputs.validation`

Required:

- `expression` (String) a JS expression that will be evaluated to determine if the validation passes or fails
- `id` (String) the ID of the validation, used to identify it
- `message` (Map of String) a map of values, where the key and values are strings
- `target` (String) the index of the input should be highlighted if the validation fails (nesting is supported via dot notation)
