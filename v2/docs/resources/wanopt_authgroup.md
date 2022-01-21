---
subcategory: "FortiGate Wanopt"
layout: "fortios"
page_title: "FortiOS: fortios_wanopt_authgroup"
description: |-
  Configure WAN optimization authentication groups.
---

## fortios_wanopt_authgroup
Configure WAN optimization authentication groups.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `auth_method` - Select certificate or pre-shared key authentication for this authentication group. Valid values: `cert` `psk` .
* `cert` - Name of certificate to identify this peer. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `name` - Auth-group name.
* `peer` - If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command. This attribute must reference one of the following datasources: `wanopt.peer.peer-host-id` .
* `peer_accept` - Determine if this auth group accepts, any peer, a list of defined peers, or just one peer. Valid values: `any` `defined` `one` .
* `psk` - Pre-shared key used by the peers in this authentication group.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wanopt_authgroup can be imported using:
```sh
terraform import fortios_wanopt_authgroup.labelname {{mkey}}
```
