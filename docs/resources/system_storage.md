---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_storage"
description: |-
  Configure logical storage.
---

## fortios_system_storage
Configure logical storage.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `device` - Partition device.
* `media_status` - The physical status of current media. Valid values: `enable` `disable` `fail` .
* `name` - Storage name.
* `order` - Set storage order.
* `partition` - Label of underlying partition.
* `size` - Partition size.
* `status` - Enable/disable storage. Valid values: `enable` `disable` .
* `usage` - Use hard disk for logging and WAN Optimization. Valid values: `log` .
* `wanopt_mode` - WAN Optimization mode (default = mix). Valid values: `mix` `wanopt` `webcache` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_storage can be imported using:
```sh
terraform import fortios_system_storage.labelname {{mkey}}
```
