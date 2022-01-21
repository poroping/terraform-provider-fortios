---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallipmacbinding_setting"
description: |-
  Get information on a fortios Configure IP to MAC binding settings.
---

# Data Source: fortios_firewallipmacbinding_setting
Use this data source to get information on a fortios Configure IP to MAC binding settings.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `bindthroughfw` - Enable/disable use of IP/MAC binding to filter packets that would normally go through the firewall.
* `bindtofw` - Enable/disable use of IP/MAC binding to filter packets that would normally go to the firewall.
* `undefinedhost` - Select action to take on packets with IP/MAC addresses not in the binding list (default = block).
