---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ssoadmin"
description: |-
  Get information on a fortios Configure SSO admin users.
---

# Data Source: fortios_system_ssoadmin
Use this data source to get information on a fortios Configure SSO admin users.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) SSO admin name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `accprofile` - SSO admin user access profile.
* `name` - SSO admin name.
* `vdom` - Virtual domain(s) that the administrator can access.The structure of `vdom` block is documented below.

The `vdom` block contains:

* `name` - Virtual domain name.
