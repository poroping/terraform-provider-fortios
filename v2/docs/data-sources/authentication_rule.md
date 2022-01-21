---
subcategory: "FortiGate Authentication"
layout: "fortios"
page_title: "FortiOS: fortios_authentication_rule"
description: |-
  Get information on a fortios Configure Authentication Rules.
---

# Data Source: fortios_authentication_rule
Use this data source to get information on a fortios Configure Authentication Rules.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Authentication rule name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `active_auth_method` - Select an active authentication method.
* `comments` - Comment.
* `ip_based` - Enable/disable IP-based authentication. When enabled, previously authenticated users from the same IP address will be exempted.
* `name` - Authentication rule name.
* `protocol` - Authentication is required for the selected protocol (default = HTTP).
* `sso_auth_method` - Select a single-sign on (SSO) authentication method.
* `status` - Enable/disable this authentication rule.
* `transaction_based` - Enable/disable transaction based authentication (default = disable).
* `web_auth_cookie` - Enable/disable Web authentication cookies (default = disable).
* `web_portal` - Enable/disable web portal for proxy transparent policy (default = enable).
* `dstaddr` - Select an IPv4 destination address from available options. Required for web proxy authentication.The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name.
* `dstaddr6` - Select an IPv6 destination address from available options. Required for web proxy authentication.The structure of `dstaddr6` block is documented below.

The `dstaddr6` block contains:

* `name` - Address name.
* `srcaddr` - Authentication is required for the selected IPv4 source address.The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name.
* `srcaddr6` - Authentication is required for the selected IPv6 source address.The structure of `srcaddr6` block is documented below.

The `srcaddr6` block contains:

* `name` - Address name.
* `srcintf` - Incoming (ingress) interface.The structure of `srcintf` block is documented below.

The `srcintf` block contains:

* `name` - Interface name.
