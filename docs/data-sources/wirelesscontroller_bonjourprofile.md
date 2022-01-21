---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_bonjourprofile"
description: |-
  Get information on a fortios Configure Bonjour profiles. Bonjour is Apple's zero configuration networking protocol. Bonjour profiles allow APs and FortiAPs to connnect to networks using Bonjour.
---

# Data Source: fortios_wirelesscontroller_bonjourprofile
Use this data source to get information on a fortios Configure Bonjour profiles. Bonjour is Apple's zero configuration networking protocol. Bonjour profiles allow APs and FortiAPs to connnect to networks using Bonjour.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Bonjour profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `name` - Bonjour profile name.
* `policy_list` - Bonjour policy list.The structure of `policy_list` block is documented below.

The `policy_list` block contains:

* `description` - Description.
* `from_vlan` - VLAN ID from which the Bonjour service is advertised (0 - 4094, default = 0).
* `policy_id` - Policy ID.
* `services` - Bonjour services for the VLAN connecting to the Bonjour network.
* `to_vlan` - VLAN ID to which the Bonjour service is made available (0 - 4094, default = all).
