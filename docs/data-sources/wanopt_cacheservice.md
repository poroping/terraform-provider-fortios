---
subcategory: "FortiGate Wanopt"
layout: "fortios"
page_title: "FortiOS: fortios_wanopt_cacheservice"
description: |-
  Get information on a fortios Designate cache-service for wan-optimization and webcache.
---

# Data Source: fortios_wanopt_cacheservice
Use this data source to get information on a fortios Designate cache-service for wan-optimization and webcache.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `acceptable_connections` - Set strategy when accepting cache collaboration connection.
* `collaboration` - Enable/disable cache-collaboration between cache-service clusters.
* `device_id` - Set identifier for this cache device.
* `prefer_scenario` - Set the preferred cache behavior towards the balance between latency and hit-ratio.
* `dst_peer` - Modify cache-service destination peer list.The structure of `dst_peer` block is documented below.

The `dst_peer` block contains:

* `auth_type` - Set authentication type for this peer.
* `device_id` - Device ID of this peer.
* `encode_type` - Set encode type for this peer.
* `ip` - Set cluster IP address of this peer.
* `priority` - Set priority for this peer.
* `src_peer` - Modify cache-service source peer list.The structure of `src_peer` block is documented below.

The `src_peer` block contains:

* `auth_type` - Set authentication type for this peer.
* `device_id` - Device ID of this peer.
* `encode_type` - Set encode type for this peer.
* `ip` - Set cluster IP address of this peer.
* `priority` - Set priority for this peer.
