---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "udoma_report_definition Resource - terraform-provider-udoma"
subcategory: ""
description: |-
  Resource represents a defintion of a report
---

# udoma_report_definition (Resource)

Resource represents a defintion of a report



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the report definition, shown in the admin page
- `result_schema` (Attributes) The schema of the result, which will be used to display it in the UI (see [below for nested schema](#nestedatt--result_schema))
- `script` (String) JS script that is executed to generate the report

### Optional

- `description` (String) The description of the report definition
- `parameters` (Attributes) Optional input definition (see [below for nested schema](#nestedatt--parameters))

### Read-Only

- `created_at` (Number) The date and time the report definition was created
- `id` (String) The unique identifier for the report definition
- `last_updated` (String)
- `updated_at` (Number) The date and time the report definition was last modified

<a id="nestedatt--result_schema"></a>
### Nested Schema for `result_schema`

Required:

- `attributes` (Attributes List) The attributes of the result (see [below for nested schema](#nestedatt--result_schema--attributes))
- `result_type` (String) The type of the result (object, table).

Optional:

- `table_row_id_attribute` (String) The attribute that will be used as the row identifier in the table

<a id="nestedatt--result_schema--attributes"></a>
### Nested Schema for `result_schema.attributes`

Required:

- `id` (String) The unique identifier for the attribute

Optional:

- `is_reference` (Boolean) Whether the attribute is a reference to another entity
- `label` (Map of String) The label of the attribute, which will be displayed in the UI
- `sequence` (Number) The sequence of the attribute, which will be used to order the attributes in the UI



<a id="nestedatt--parameters"></a>
### Nested Schema for `parameters`

Required:

- `inputs` (Attributes List) The inputs that will be displayed to the user (see [below for nested schema](#nestedatt--parameters--inputs))
- `layout` (Attributes List) The layout of the form, which groups and inputs will be displayed (see [below for nested schema](#nestedatt--parameters--layout))

Optional:

- `groups` (Attributes List) The groups of inputs that will be displayed to the user (see [below for nested schema](#nestedatt--parameters--groups))
- `validation` (Attributes List) The validations that will be performed on the data provided by the user (see [below for nested schema](#nestedatt--parameters--validation))

<a id="nestedatt--parameters--inputs"></a>
### Nested Schema for `parameters.inputs`

Required:

- `id` (String) the ID of the input field, used to identify it and later access the data
- `type` (String) The type of the input

Optional:

- `attributes` (Map of String) a map of values, where the key and values are strings
- `default_value` (String) optional default value for the input field (as a JSON string)
- `ephemeral` (Boolean) if true, the value of the input will not be persisted
- `items` (Attributes List) Only used when the type is select or multi select. This is a list of values that the user can choose from. (see [below for nested schema](#nestedatt--parameters--inputs--items))
- `label` (Map of String) a map of values, where the key and values are strings
- `placeholder` (Map of String) a map of values, where the key and values are strings
- `propagate_changes` (Boolean) if true, changes to the input will be propagated to event listeners for the custom form
- `readonly` (Boolean) if true, the user will not be able to change the value of this input
- `required` (Boolean) if true, the user will be required to provide a value
- `target` (String) the attribute name to use when exporting the result of this input
- `view_label` (Map of String) a map of values, where the key and values are strings

<a id="nestedatt--parameters--inputs--items"></a>
### Nested Schema for `parameters.inputs.items`

Required:

- `id` (String) The technical value that will be stored in the collected data
- `label` (Map of String) The label to be displayed to the user as a language dictionary



<a id="nestedatt--parameters--layout"></a>
### Nested Schema for `parameters.layout`

Required:

- `ref_id` (String) The ID of the entity that will be referenced
- `ref_type` (String) The type of the entity that will be referenced


<a id="nestedatt--parameters--groups"></a>
### Nested Schema for `parameters.groups`

Required:

- `id` (String) The ID of the group
- `items` (Attributes List) the IDs of the inputs that will be displayed in the group (see [below for nested schema](#nestedatt--parameters--groups--items))
- `type` (String) The type of the group

Optional:

- `bottom_divider` (Boolean) if true, a divider will be displayed below the group
- `info` (Map of String) The info of the group
- `label` (Map of String) The label of the group
- `min_size` (Number) the minimum number of items that must be submitted in the group (only used for repeat groups)
- `target` (String) the attribute name to use when exporting the result of this group (only used for repeat groups)
- `top_divider` (Boolean) if true, a divider will be displayed above the group

<a id="nestedatt--parameters--groups--items"></a>
### Nested Schema for `parameters.groups.items`

Required:

- `ref_id` (String) The ID of the entity that will be referenced
- `ref_type` (String) The type of the entity that will be referenced



<a id="nestedatt--parameters--validation"></a>
### Nested Schema for `parameters.validation`

Required:

- `expression` (String) a JS expression that will be evaluated to determine if the validation passes or fails
- `id` (String) the ID of the validation, used to identify it
- `message` (Map of String) a map of values, where the key and values are strings
- `target` (String) the index of the input should be highlighted if the validation fails (nesting is supported via dot notation)