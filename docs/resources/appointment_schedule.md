---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "udoma_appointment_schedule Resource - terraform-provider-udoma"
subcategory: ""
description: |-
  Resource represents an appointment schedule
---

# udoma_appointment_schedule (Resource)

Resource represents an appointment schedule



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `color` (String) The color to use when displaying the schedule
- `description` (Map of String) Description of the appointment schedule in all different languages
- `gap_duration` (Number) The length of the gap in-between slots
- `name` (String) The name of the appointment schedule
- `slot_duration` (Number) The length of a single appointment
- `template_ref` (String) ID of the template to be used in the schedule

### Read-Only

- `id` (String) The unique identifier for the appointment schedule

