---
subcategory: "FortiGate Wanopt"
layout: "fortios"
page_title: "FortiOS: fortios_wanopt_authgroup"
description: |-
  Get information on a fortios Configure WAN optimization authentication groups.
---

# Data Source: fortios_wanopt_authgroup
Use this data source to get information on a fortios Configure WAN optimization authentication groups.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Auth-group name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `auth_method` - Select certificate or pre-shared key authentication for this authentication group.
* `cert` - Name of certificate to identify this peer.
* `name` - Auth-group name.
* `peer` - If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command.
* `peer_accept` - Determine if this auth group accepts, any peer, a list of defined peers, or just one peer.
* `psk` - Pre-shared key used by the peers in this authentication group.
