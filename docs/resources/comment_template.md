---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "udoma_comment_template Resource - terraform-provider-udoma"
subcategory: ""
description: |-
  Comment template resource
---

# udoma_comment_template (Resource)

Comment template resource



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `display_name` (String) The display name of the comment template
- `template` (String) The template that will be used to generate the comment

### Optional

- `access_list` (List of String) The list of users that have access to the comment template
- `is_deny_list` (Boolean) If set true the access list will act as a deny list instead of an allow list
- `script` (String) The script that will be executed to generate the comment

### Read-Only

- `created_at` (Number) The date and time the template was created
- `id` (String) The unique identifier of the comment template
- `updated_at` (Number) The date and time the template was last updated
