---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemautoupdate_pushupdate"
description: |-
  Get information on a fortios Configure push updates.
---

# Data Source: fortios_systemautoupdate_pushupdate
Use this data source to get information on a fortios Configure push updates.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `address` - IPv4 or IPv6 address used by FortiGuard servers to send push updates to this FortiGate.
* `override` - Enable/disable push update override server.
* `port` - Push update override port. (Do not overlap with other service ports)
* `status` - Enable/disable push updates.
