---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnssl_client"
description: |-
  Client.
---

## fortios_vpnssl_client
Client.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `certificate` - Certificate to offer to SSL-VPN server if it requests one. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `comment` - Comment.
* `distance` - Distance for routes added by SSL-VPN (1 - 255).
* `interface` - SSL interface to send/receive traffic over. This attribute must reference one of the following datasources: `system.interface.name` .
* `name` - SSL-VPN tunnel name.
* `peer` - Authenticate peer's certificate with the peer/peergrp. This attribute must reference one of the following datasources: `user.peer.name` `user.peergrp.name` .
* `port` - SSL-VPN server port.
* `priority` - Priority for routes added by SSL-VPN (1 - 65535).
* `psk` - Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
* `realm` - Realm name configured on SSL-VPN server.
* `server` - IPv4, IPv6 or DNS address of the SSL-VPN server.
* `source_ip` - IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
* `status` - Enable/disable this SSL-VPN client configuration. Valid values: `enable` `disable` .
* `user` - Username to offer to the peer to authenticate the client.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_vpnssl_client can be imported using:
```sh
terraform import fortios_vpnssl_client.labelname {{mkey}}
```
