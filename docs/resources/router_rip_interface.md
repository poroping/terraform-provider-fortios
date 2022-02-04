---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_rip_interface"
description: |-
  RIP interface configuration.
---

## fortios_router_rip_interface
RIP interface configuration.

~> This resource is configuring a child table of the parent resource: `fortios_router_rip`. If this resource is used the parent resource must NOT modify this child table or state inconsistencies will occur.


## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `auth_keychain` - Authentication key-chain name. This attribute must reference one of the following datasources: `router.key-chain.name` .
* `auth_mode` - Authentication mode. Valid values: `none` `text` `md5` .
* `auth_string` - Authentication string/password.
* `flags` - Flags.
* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `receive_version` - Receive version. Valid values: `1` `2` .
* `send_version` - Send version. Valid values: `1` `2` .
* `send_version2_broadcast` - Enable/disable broadcast version 1 compatible packets. Valid values: `disable` `enable` .
* `split_horizon` - Enable/disable split horizon. Valid values: `poisoned` `regular` .
* `split_horizon_status` - Enable/disable split horizon. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_interface can be imported using:
```sh
terraform import fortios_router_interface.labelname {{mkey}}
```
