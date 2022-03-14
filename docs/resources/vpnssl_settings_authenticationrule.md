---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnssl_settings_authentication_rule"
description: |-
  Authentication rule for SSL-VPN.
---

## fortios_vpnssl_settings_authenticationrule
Authentication rule for SSL-VPN.

~> This resource is configuring a child table of the parent resource: `fortios_vpnssl_settings`. If this resource is used the parent resource must NOT modify this child table or state inconsistencies will occur.


## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `auth` - SSL-VPN authentication method restriction. Valid values: `any` `local` `radius` `tacacs+` `ldap` `peer` .
* `cipher` - SSL-VPN cipher strength. Valid values: `any` `high` `medium` .
* `client_cert` - Enable/disable SSL-VPN client certificate restrictive. Valid values: `enable` `disable` .
* `id` - ID (0 - 4294967295).
* `portal` - SSL-VPN portal. This attribute must reference one of the following datasources: `vpn.ssl.web.portal.name` .
* `realm` - SSL-VPN realm. This attribute must reference one of the following datasources: `vpn.ssl.web.realm.url-path` .
* `source_address_negate` - Enable/disable negated source address match. Valid values: `enable` `disable` .
* `source_address6_negate` - Enable/disable negated source IPv6 address match. Valid values: `enable` `disable` .
* `user_peer` - Name of user peer. This attribute must reference one of the following datasources: `user.peer.name` .
* `groups` - User groups. The structure of `groups` block is documented below.

The `groups` block contains:

* `name` - Group name. This attribute must reference one of the following datasources: `user.group.name` .
* `source_address` - Source address of incoming traffic. The structure of `source_address` block is documented below.

The `source_address` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` `system.external-resource.name` .
* `source_address6` - IPv6 source address of incoming traffic. The structure of `source_address6` block is documented below.

The `source_address6` block contains:

* `name` - IPv6 address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` `system.external-resource.name` .
* `source_interface` - SSL-VPN source interface of incoming traffic. The structure of `source_interface` block is documented below.

The `source_interface` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` .
* `users` - User name. The structure of `users` block is documented below.

The `users` block contains:

* `name` - User name. This attribute must reference one of the following datasources: `user.local.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_vpnssl_authenticationrule can be imported using:
```sh
terraform import fortios_vpnssl_authenticationrule.labelname {{mkey}}
```
