---
subcategory: "FortiGate Authentication"
layout: "fortios"
page_title: "FortiOS: fortios_authentication_rule"
description: |-
  Configure Authentication Rules.
---

## fortios_authentication_rule
Configure Authentication Rules.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `active_auth_method` - Select an active authentication method. This attribute must reference one of the following datasources: `authentication.scheme.name` .
* `comments` - Comment.
* `ip_based` - Enable/disable IP-based authentication. When enabled, previously authenticated users from the same IP address will be exempted. Valid values: `enable` `disable` .
* `name` - Authentication rule name.
* `protocol` - Authentication is required for the selected protocol (default = HTTP). Valid values: `http` `ftp` `socks` `ssh` .
* `sso_auth_method` - Select a single-sign on (SSO) authentication method. This attribute must reference one of the following datasources: `authentication.scheme.name` .
* `status` - Enable/disable this authentication rule. Valid values: `enable` `disable` .
* `transaction_based` - Enable/disable transaction based authentication (default = disable). Valid values: `enable` `disable` .
* `web_auth_cookie` - Enable/disable Web authentication cookies (default = disable). Valid values: `enable` `disable` .
* `web_portal` - Enable/disable web portal for proxy transparent policy (default = enable). Valid values: `enable` `disable` .
* `dstaddr` - Select an IPv4 destination address from available options. Required for web proxy authentication. The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` `firewall.proxy-address.name` `firewall.proxy-addrgrp.name` `system.external-resource.name` .
* `dstaddr6` - Select an IPv6 destination address from available options. Required for web proxy authentication. The structure of `dstaddr6` block is documented below.

The `dstaddr6` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `srcaddr` - Authentication is required for the selected IPv4 source address. The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` `firewall.proxy-address.name` `firewall.proxy-addrgrp.name` `system.external-resource.name` .
* `srcaddr6` - Authentication is required for the selected IPv6 source address. The structure of `srcaddr6` block is documented below.

The `srcaddr6` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `srcintf` - Incoming (ingress) interface. The structure of `srcintf` block is documented below.

The `srcintf` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_authentication_rule can be imported using:
```sh
terraform import fortios_authentication_rule.labelname {{mkey}}
```
