---
subcategory: "FortiGate Wanopt"
layout: "fortios"
page_title: "FortiOS: fortios_wanopt_peer"
description: |-
  Get information on a fortios Configure WAN optimization peers.
---

# Data Source: fortios_wanopt_peer
Use this data source to get information on a fortios Configure WAN optimization peers.


## Example Usage

```hcl

```

## Argument Reference

* `peer_host_id` - (Required) Peer host ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `ip` - Peer IP address.
* `peer_host_id` - Peer host ID.
