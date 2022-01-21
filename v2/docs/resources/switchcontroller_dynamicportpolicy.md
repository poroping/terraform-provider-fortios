---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_dynamicportpolicy"
description: |-
  Configure Dynamic port policy to be applied on the managed FortiSwitch ports through DPP device.
---

## fortios_switchcontroller_dynamicportpolicy
Configure Dynamic port policy to be applied on the managed FortiSwitch ports through DPP device.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `description` - Description for the Dynamic port policy.
* `fortilink` - FortiLink interface for which this Dynamic port policy belongs to. This attribute must reference one of the following datasources: `system.interface.name` .
* `name` - Dynamic port policy name.
* `policy` - Port policies with matching criteria and actions. The structure of `policy` block is documented below.

The `policy` block contains:

* `802_1x` - 802.1x security policy to be applied when using this policy. This attribute must reference one of the following datasources: `switch-controller.security-policy.802-1X.name` `switch-controller.security-policy.captive-portal.name` .
* `bounce_port_link` - Enable/disable bouncing (administratively bring the link down, up) of a switch port where this policy is applied. Helps to clear and reassign VLAN from lldp-profile. Valid values: `disable` `enable` .
* `category` - Category of Dynamic port policy. Valid values: `device` `interface-tag` .
* `description` - Description for the policy.
* `family` - Policy matching family.
* `host` - Policy matching host.
* `lldp_profile` - LLDP profile to be applied when using this policy. This attribute must reference one of the following datasources: `switch-controller.lldp-profile.name` .
* `mac` - Policy matching MAC address.
* `name` - Policy name.
* `qos_policy` - QoS policy to be applied when using this policy. This attribute must reference one of the following datasources: `switch-controller.qos.qos-policy.name` .
* `status` - Enable/disable policy. Valid values: `enable` `disable` .
* `type` - Policy matching type.
* `vlan_policy` - VLAN policy to be applied when using this policy. This attribute must reference one of the following datasources: `switch-controller.vlan-policy.name` .
* `interface_tags` - Policy matching the FortiSwitch interface object tags. The structure of `interface_tags` block is documented below.

The `interface_tags` block contains:

* `tag_name` - FortiSwitch port tag name. This attribute must reference one of the following datasources: `switch-controller.switch-interface-tag.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_switchcontroller_dynamicportpolicy can be imported using:
```sh
terraform import fortios_switchcontroller_dynamicportpolicy.labelname {{mkey}}
```
