---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "udoma_attachment Resource - terraform-provider-udoma"
subcategory: ""
description: |-
  Resource represents an attachment
---

# udoma_attachment (Resource)

Resource represents an attachment



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `source` (String) The file location to read the JSON from

### Read-Only

- `created` (Number) The date and time the attachment was created
- `file_name` (String) The name of the attachment file
- `file_sha256` (String) The SHA256 sum of the file
- `file_size` (Number) Size of the attachment
- `file_type` (String) Metadata about the attachment
- `id` (String) The unique identifier for the attachment
- `url` (String) The URL to download the attachment
