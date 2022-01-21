---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_apcfgprofile"
description: |-
  Get information on a fortios Configure AP local configuration profiles.
---

# Data Source: fortios_wirelesscontroller_apcfgprofile
Use this data source to get information on a fortios Configure AP local configuration profiles.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) AP local configuration profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `ac_ip` - IP address of the validation controller that AP must be able to join after applying AP local configuration.
* `ac_port` - Port of the validation controller that AP must be able to join after applying AP local configuration (1024 - 49150, default = 5246).
* `ac_timer` - Maximum waiting time for the AP to join the validation controller after applying AP local configuration (3 - 30 min, default = 10).
* `ac_type` - Validation controller type (default = default).
* `ap_family` - FortiAP family type (default = fap).
* `comment` - Comment.
* `name` - AP local configuration profile name.
* `command_list` - AP local configuration command list.The structure of `command_list` block is documented below.

The `command_list` block contains:

* `id` - Command ID.
* `name` - AP local configuration command name.
* `passwd_value` - AP local configuration command password value.
* `type` - The command type (default = non-password).
* `value` - AP local configuration command value.
