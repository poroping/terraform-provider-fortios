---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_storage"
description: |-
  Get information on a fortios Configure logical storage.
---

# Data Source: fortios_system_storage
Use this data source to get information on a fortios Configure logical storage.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Storage name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `device` - Partition device.
* `media_status` - The physical status of current media.
* `name` - Storage name.
* `order` - Set storage order.
* `partition` - Label of underlying partition.
* `size` - Partition size.
* `status` - Enable/disable storage.
* `usage` - Use hard disk for logging or WAN Optimization (default = log).
* `wanopt_mode` - WAN Optimization mode (default = mix).
