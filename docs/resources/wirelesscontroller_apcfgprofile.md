---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_apcfgprofile"
description: |-
  Configure AP local configuration profiles.
---

## fortios_wirelesscontroller_apcfgprofile
Configure AP local configuration profiles.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `ac_ip` - IP address of the validation controller that AP must be able to join after applying AP local configuration.
* `ac_port` - Port of the validation controller that AP must be able to join after applying AP local configuration (1024 - 49150, default = 5246).
* `ac_timer` - Maximum waiting time for the AP to join the validation controller after applying AP local configuration (3 - 30 min, default = 10).
* `ac_type` - Validation controller type (default = default). Valid values: `default` `specify` `apcfg` .
* `ap_family` - FortiAP family type (default = fap). Valid values: `fap` `fap-u` `fap-c` .
* `comment` - Comment.
* `name` - AP local configuration profile name.
* `command_list` - AP local configuration command list. The structure of `command_list` block is documented below.

The `command_list` block contains:

* `id` - Command ID.
* `name` - AP local configuration command name.
* `passwd_value` - AP local configuration command password value.
* `type` - The command type (default = non-password). Valid values: `non-password` `password` .
* `value` - AP local configuration command value.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontroller_apcfgprofile can be imported using:
```sh
terraform import fortios_wirelesscontroller_apcfgprofile.labelname {{mkey}}
```
