---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_log_customfield"
description: |-
  Get information on a fortios Configure custom log fields.
---

# Data Source: fortios_log_customfield
Use this data source to get information on a fortios Configure custom log fields.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Field ID string.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `id` - Field ID string.
* `name` - Field name (max: 15 characters).
* `value` - Field value (max: 15 characters).
