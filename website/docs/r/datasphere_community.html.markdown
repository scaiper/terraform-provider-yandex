---
layout: "yandex"
page_title: "Yandex: yandex_datasphere_community"
sidebar_current: "docs-yandex-datasphere-community"
description: |-
Allows management of a Yandex.Cloud Datasphere Community.
---

# yandex\_datasphere\_community

Allows management of Yandex Cloud Datasphere Communities

## Example Usage

```hcl
resource "yandex_datasphere_community" "my-community" {
  name = "example-datasphere-community"
  description = "Description of community"
  billing_account_id = "example-organization-id"
  labels = {
    "foo": "bar"
  }
  organization_id = "example-organization-id"
}
```

## Argument Reference

The following arguments are supported:

* `organization_id` - (Required) Organization ID where community would be created
* `name` - (Required) Name of the Datasphere Community.
* `description` -  (Optional) Datasphere Community description.
* `labels` - (Optional) A set of key/value label pairs to assign to the Datasphere Community.
* `billing_account_id` - (Optional) Billing account ID to associated with community

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - Datasphere Community unique identifier
* `created_at` - Creation timestamp of the Yandex Datasphere Community
* `created_by` - Creator account ID of the Yandex Datasphere Community


## Timeouts

This resource provides the following configuration options for timeouts:

- `create` - Default is 1 minute.
- `update` - Default is 1 minute.
- `delete` - Default is 1 minute.

## Import

A Datasphere Community can be imported using the `id` of the resource, e.g.:

```
$ terraform import yandex_datasphere_community.default community_id
```
