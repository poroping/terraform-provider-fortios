---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetserviceipblvendor"
description: |-
  Get information on a fortios IP blocklist vendor.
---

# Data Source: fortios_firewall_internetserviceipblvendor
Use this data source to get information on a fortios IP blocklist vendor.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) IP blocklist vendor ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `id` - IP blocklist vendor ID.
* `name` - IP blocklist vendor name.
