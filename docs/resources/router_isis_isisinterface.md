---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_isis_isisinterface"
description: |-
  IS-IS interface configuration.
---

## fortios_router_isis_isisinterface
IS-IS interface configuration.

~> This resource is configuring a child table of the parent resource: `fortios_router_isis`. If this resource is used the parent resource must NOT modify this child table or state inconsistencies will occur.


## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `auth_keychain_l1` - Authentication key-chain for level 1 PDUs. This attribute must reference one of the following datasources: `router.key-chain.name` .
* `auth_keychain_l2` - Authentication key-chain for level 2 PDUs. This attribute must reference one of the following datasources: `router.key-chain.name` .
* `auth_mode_l1` - Level 1 authentication mode. Valid values: `md5` `password` .
* `auth_mode_l2` - Level 2 authentication mode. Valid values: `md5` `password` .
* `auth_password_l1` - Authentication password for level 1 PDUs.
* `auth_password_l2` - Authentication password for level 2 PDUs.
* `auth_send_only_l1` - Enable/disable authentication send-only for level 1 PDUs. Valid values: `enable` `disable` .
* `auth_send_only_l2` - Enable/disable authentication send-only for level 2 PDUs. Valid values: `enable` `disable` .
* `circuit_type` - IS-IS interface's circuit type Valid values: `level-1-2` `level-1` `level-2` .
* `csnp_interval_l1` - Level 1 CSNP interval.
* `csnp_interval_l2` - Level 2 CSNP interval.
* `hello_interval_l1` - Level 1 hello interval.
* `hello_interval_l2` - Level 2 hello interval.
* `hello_multiplier_l1` - Level 1 multiplier for Hello holding time.
* `hello_multiplier_l2` - Level 2 multiplier for Hello holding time.
* `hello_padding` - Enable/disable padding to IS-IS hello packets. Valid values: `enable` `disable` .
* `lsp_interval` - LSP transmission interval (milliseconds).
* `lsp_retransmit_interval` - LSP retransmission interval (sec).
* `mesh_group` - Enable/disable IS-IS mesh group. Valid values: `enable` `disable` .
* `mesh_group_id` - Mesh group ID <0-4294967295>, 0: mesh-group blocked.
* `metric_l1` - Level 1 metric for interface.
* `metric_l2` - Level 2 metric for interface.
* `name` - IS-IS interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `network_type` - IS-IS interface's network type Valid values: `broadcast` `point-to-point` `loopback` .
* `priority_l1` - Level 1 priority.
* `priority_l2` - Level 2 priority.
* `status` - Enable/disable interface for IS-IS. Valid values: `enable` `disable` .
* `status6` - Enable/disable IPv6 interface for IS-IS. Valid values: `enable` `disable` .
* `wide_metric_l1` - Level 1 wide metric for interface.
* `wide_metric_l2` - Level 2 wide metric for interface.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_isisinterface can be imported using:
```sh
terraform import fortios_router_isisinterface.labelname {{mkey}}
```
