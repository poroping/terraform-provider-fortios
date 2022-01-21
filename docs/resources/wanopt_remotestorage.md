---
subcategory: "FortiGate Wanopt"
layout: "fortios"
page_title: "FortiOS: fortios_wanopt_remotestorage"
description: |-
  Configure a remote cache device as Web cache storage.
---

## fortios_wanopt_remotestorage
Configure a remote cache device as Web cache storage.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `local_cache_id` - ID that this device uses to connect to the remote device.
* `remote_cache_id` - ID of the remote device to which the device connects.
* `remote_cache_ip` - IP address of the remote device to which the device connects.
* `status` - Enable/disable using remote device as Web cache storage. Valid values: `disable` `enable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wanopt_remotestorage can be imported using:
```sh
terraform import fortios_wanopt_remotestorage.labelname {{mkey}}
```
