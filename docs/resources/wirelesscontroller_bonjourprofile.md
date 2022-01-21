---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_bonjourprofile"
description: |-
  Configure Bonjour profiles. Bonjour is Apple's zero configuration networking protocol. Bonjour profiles allow APs and FortiAPs to connnect to networks using Bonjour.
---

## fortios_wirelesscontroller_bonjourprofile
Configure Bonjour profiles. Bonjour is Apple's zero configuration networking protocol. Bonjour profiles allow APs and FortiAPs to connnect to networks using Bonjour.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comment` - Comment.
* `name` - Bonjour profile name.
* `policy_list` - Bonjour policy list. The structure of `policy_list` block is documented below.

The `policy_list` block contains:

* `description` - Description.
* `from_vlan` - VLAN ID from which the Bonjour service is advertised (0 - 4094, default = 0).
* `policy_id` - Policy ID.
* `services` - Bonjour services for the VLAN connecting to the Bonjour network. Valid values: `all` `airplay` `afp` `bit-torrent` `ftp` `ichat` `itunes` `printers` `samba` `scanners` `ssh` `chromecast` .
* `to_vlan` - VLAN ID to which the Bonjour service is made available (0 - 4094, default = all).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontroller_bonjourprofile can be imported using:
```sh
terraform import fortios_wirelesscontroller_bonjourprofile.labelname {{mkey}}
```
