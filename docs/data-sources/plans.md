---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gitops_plans Data Source - gitops"
subcategory: ""
description: |-
  Manages a Gitops instance order
---

# gitops_plans (Data Source)

Manages a Gitops instance order



<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `plans` (Attributes List) (see [below for nested schema](#nestedatt--plans))

<a id="nestedatt--plans"></a>
### Nested Schema for `plans`

Read-Only:

- `bits_account` (Number) Account-ID of the Gitops resource instance
- `instance_name` (String) Name of the Gitops resource instance
- `orderer_id` (String) Name of the Gitops resource orderer
- `replica_count` (Number) Replica count of the Gitops resource instance
- `service_id` (Number) Service-ID of the Gitops resource instance
- `some_value` (String) Some custom value of the Gitops resource instance
- `version` (String) Version of the Gitops resource instance
