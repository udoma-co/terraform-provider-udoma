---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "udoma_document Resource - terraform-provider-udoma"
subcategory: ""
description: |-
  Resource represents a document
---

# udoma_document (Resource)

Resource represents a document



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `attachment` (String) The attachment to be used in the document
- `name` (String) The name of the document
- `ref_type` (String) The type of the reference

### Optional

- `path` (String) The path of the document in the storage
- `public` (Boolean) The public status of the document
- `ref_id` (String) The identifier of the reference

### Read-Only

- `created_at` (Number) The time the document was created
- `id` (String) The unique identifier for the document
- `updated_at` (Number) The time the document was last updated
