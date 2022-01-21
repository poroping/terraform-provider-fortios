---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_isis_isisinterface"
description: |-
  Get information on a fortios IS-IS interface configuration.
---

# Data Source: fortios_router_isis_isisinterface
Use this data source to get information on a fortios IS-IS interface configuration.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) IS-IS interface name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `auth_keychain_l1` - Authentication key-chain for level 1 PDUs.
* `auth_keychain_l2` - Authentication key-chain for level 2 PDUs.
* `auth_mode_l1` - Level 1 authentication mode.
* `auth_mode_l2` - Level 2 authentication mode.
* `auth_password_l1` - Authentication password for level 1 PDUs.
* `auth_password_l2` - Authentication password for level 2 PDUs.
* `auth_send_only_l1` - Enable/disable authentication send-only for level 1 PDUs.
* `auth_send_only_l2` - Enable/disable authentication send-only for level 2 PDUs.
* `circuit_type` - IS-IS interface's circuit type
* `csnp_interval_l1` - Level 1 CSNP interval.
* `csnp_interval_l2` - Level 2 CSNP interval.
* `hello_interval_l1` - Level 1 hello interval.
* `hello_interval_l2` - Level 2 hello interval.
* `hello_multiplier_l1` - Level 1 multiplier for Hello holding time.
* `hello_multiplier_l2` - Level 2 multiplier for Hello holding time.
* `hello_padding` - Enable/disable padding to IS-IS hello packets.
* `lsp_interval` - LSP transmission interval (milliseconds).
* `lsp_retransmit_interval` - LSP retransmission interval (sec).
* `mesh_group` - Enable/disable IS-IS mesh group.
* `mesh_group_id` - Mesh group ID <0-4294967295>, 0: mesh-group blocked.
* `metric_l1` - Level 1 metric for interface.
* `metric_l2` - Level 2 metric for interface.
* `name` - IS-IS interface name.
* `network_type` - IS-IS interface's network type
* `priority_l1` - Level 1 priority.
* `priority_l2` - Level 2 priority.
* `status` - Enable/disable interface for IS-IS.
* `status6` - Enable/disable IPv6 interface for IS-IS.
* `wide_metric_l1` - Level 1 wide metric for interface.
* `wide_metric_l2` - Level 2 wide metric for interface.
