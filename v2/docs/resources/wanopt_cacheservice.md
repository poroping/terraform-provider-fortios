---
subcategory: "FortiGate Wanopt"
layout: "fortios"
page_title: "FortiOS: fortios_wanopt_cacheservice"
description: |-
  Designate cache-service for wan-optimization and webcache.
---

## fortios_wanopt_cacheservice
Designate cache-service for wan-optimization and webcache.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `acceptable_connections` - Set strategy when accepting cache collaboration connection. Valid values: `any` `peers` .
* `collaboration` - Enable/disable cache-collaboration between cache-service clusters. Valid values: `enable` `disable` .
* `device_id` - Set identifier for this cache device.
* `prefer_scenario` - Set the preferred cache behavior towards the balance between latency and hit-ratio. Valid values: `balance` `prefer-speed` `prefer-cache` .
* `dst_peer` - Modify cache-service destination peer list. The structure of `dst_peer` block is documented below.

The `dst_peer` block contains:

* `auth_type` - Set authentication type for this peer.
* `device_id` - Device ID of this peer.
* `encode_type` - Set encode type for this peer.
* `ip` - Set cluster IP address of this peer.
* `priority` - Set priority for this peer.
* `src_peer` - Modify cache-service source peer list. The structure of `src_peer` block is documented below.

The `src_peer` block contains:

* `auth_type` - Set authentication type for this peer.
* `device_id` - Device ID of this peer.
* `encode_type` - Set encode type for this peer.
* `ip` - Set cluster IP address of this peer.
* `priority` - Set priority for this peer.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wanopt_cacheservice can be imported using:
```sh
terraform import fortios_wanopt_cacheservice.labelname {{mkey}}
```
