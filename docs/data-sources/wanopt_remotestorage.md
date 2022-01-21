---
subcategory: "FortiGate Wanopt"
layout: "fortios"
page_title: "FortiOS: fortios_wanopt_remotestorage"
description: |-
  Get information on a fortios Configure a remote cache device as Web cache storage.
---

# Data Source: fortios_wanopt_remotestorage
Use this data source to get information on a fortios Configure a remote cache device as Web cache storage.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `local_cache_id` - ID that this device uses to connect to the remote device.
* `remote_cache_id` - ID of the remote device to which the device connects.
* `remote_cache_ip` - IP address of the remote device to which the device connects.
* `status` - Enable/disable using remote device as Web cache storage.
