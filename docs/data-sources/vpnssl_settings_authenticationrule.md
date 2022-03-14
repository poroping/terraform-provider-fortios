---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnssl_settings_authentication_rule"
description: |-
  Get information on a fortios Authentication rule for SSL-VPN.
---

# Data Source: fortios_vpnssl_settings_authenticationrule
Use this data source to get information on a fortios Authentication rule for SSL-VPN.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) ID (0 - 4294967295).
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `auth` - SSL-VPN authentication method restriction.
* `cipher` - SSL-VPN cipher strength.
* `client_cert` - Enable/disable SSL-VPN client certificate restrictive.
* `id` - ID (0 - 4294967295).
* `portal` - SSL-VPN portal.
* `realm` - SSL-VPN realm.
* `source_address_negate` - Enable/disable negated source address match.
* `source_address6_negate` - Enable/disable negated source IPv6 address match.
* `user_peer` - Name of user peer.
* `groups` - User groups.The structure of `groups` block is documented below.

The `groups` block contains:

* `name` - Group name.
* `source_address` - Source address of incoming traffic.The structure of `source_address` block is documented below.

The `source_address` block contains:

* `name` - Address name.
* `source_address6` - IPv6 source address of incoming traffic.The structure of `source_address6` block is documented below.

The `source_address6` block contains:

* `name` - IPv6 address name.
* `source_interface` - SSL-VPN source interface of incoming traffic.The structure of `source_interface` block is documented below.

The `source_interface` block contains:

* `name` - Interface name.
* `users` - User name.The structure of `users` block is documented below.

The `users` block contains:

* `name` - User name.
