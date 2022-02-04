---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnssl_client"
description: |-
  Get information on a fortios Client.
---

# Data Source: fortios_vpnssl_client
Use this data source to get information on a fortios Client.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) SSL-VPN tunnel name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `certificate` - Certificate to offer to SSL-VPN server if it requests one.
* `comment` - Comment.
* `distance` - Distance for routes added by SSL-VPN (1 - 255).
* `interface` - SSL interface to send/receive traffic over.
* `name` - SSL-VPN tunnel name.
* `peer` - Authenticate peer's certificate with the peer/peergrp.
* `port` - SSL-VPN server port.
* `priority` - Priority for routes added by SSL-VPN (1 - 65535).
* `psk` - Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
* `realm` - Realm name configured on SSL-VPN server.
* `server` - IPv4, IPv6 or DNS address of the SSL-VPN server.
* `source_ip` - IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
* `status` - Enable/disable this SSL-VPN client configuration.
* `user` - Username to offer to the peer to authenticate the client.
