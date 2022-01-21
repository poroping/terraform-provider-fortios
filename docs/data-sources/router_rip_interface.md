---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_rip_interface"
description: |-
  Get information on a fortios RIP interface configuration.
---

# Data Source: fortios_router_rip_interface
Use this data source to get information on a fortios RIP interface configuration.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Interface name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `auth_keychain` - Authentication key-chain name.
* `auth_mode` - Authentication mode.
* `auth_string` - Authentication string/password.
* `flags` - flags
* `name` - Interface name.
* `receive_version` - Receive version.
* `send_version` - Send version.
* `send_version2_broadcast` - Enable/disable broadcast version 1 compatible packets.
* `split_horizon` - Enable/disable split horizon.
* `split_horizon_status` - Enable/disable split horizon.
