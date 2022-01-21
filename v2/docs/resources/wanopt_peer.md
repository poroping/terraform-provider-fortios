---
subcategory: "FortiGate Wanopt"
layout: "fortios"
page_title: "FortiOS: fortios_wanopt_peer"
description: |-
  Configure WAN optimization peers.
---

## fortios_wanopt_peer
Configure WAN optimization peers.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `peer_host_id` to be defined.

* `ip` - Peer IP address.
* `peer_host_id` - Peer host ID.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wanopt_peer can be imported using:
```sh
terraform import fortios_wanopt_peer.labelname {{mkey}}
```
