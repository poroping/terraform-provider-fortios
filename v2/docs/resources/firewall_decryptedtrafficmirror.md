---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_decryptedtrafficmirror"
description: |-
  Configure decrypted traffic mirror.
---

## fortios_firewall_decryptedtrafficmirror
Configure decrypted traffic mirror.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `dstmac` - Set destination MAC address for mirrored traffic.
* `name` - Name.
* `traffic_source` - Source of decrypted traffic to be mirrored. Valid values: `client` `server` `both` .
* `traffic_type` - Types of decrypted traffic to be mirrored. Valid values: `ssl` `ssh` .
* `interface` - Decrypted traffic mirror interface The structure of `interface` block is documented below.

The `interface` block contains:

* `name` - Decrypted traffic mirror interface. This attribute must reference one of the following datasources: `system.interface.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_decryptedtrafficmirror can be imported using:
```sh
terraform import fortios_firewall_decryptedtrafficmirror.labelname {{mkey}}
```
