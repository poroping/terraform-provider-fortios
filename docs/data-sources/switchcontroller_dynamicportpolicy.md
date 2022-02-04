---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_dynamicportpolicy"
description: |-
  Get information on a fortios Configure Dynamic port policy to be applied on the managed FortiSwitch ports through DPP device.
---

# Data Source: fortios_switchcontroller_dynamicportpolicy
Use this data source to get information on a fortios Configure Dynamic port policy to be applied on the managed FortiSwitch ports through DPP device.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Dynamic port policy name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `description` - Description for the Dynamic port policy.
* `fortilink` - FortiLink interface for which this Dynamic port policy belongs to.
* `name` - Dynamic port policy name.
* `policy` - Port policies with matching criteria and actions.The structure of `policy` block is documented below.

The `policy` block contains:

* `802_1x` - 802.1x security policy to be applied when using this policy.
* `bounce_port_link` - Enable/disable bouncing (administratively bring the link down, up) of a switch port where this policy is applied. Helps to clear and reassign VLAN from lldp-profile.
* `category` - Category of Dynamic port policy.
* `description` - Description for the policy.
* `family` - Match policy based on family.
* `host` - Match policy based on host.
* `hw_vendor` - Match policy based on hardware vendor.
* `lldp_profile` - LLDP profile to be applied when using this policy.
* `mac` - Match policy based on MAC address.
* `name` - Policy name.
* `qos_policy` - QoS policy to be applied when using this policy.
* `status` - Enable/disable policy.
* `type` - Match policy based on type.
* `vlan_policy` - VLAN policy to be applied when using this policy.
* `interface_tags` - Match policy based on the FortiSwitch interface object tags.The structure of `interface_tags` block is documented below.

The `interface_tags` block contains:

* `tag_name` - FortiSwitch port tag name.
