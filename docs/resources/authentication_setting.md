---
subcategory: "FortiGate Authentication"
layout: "fortios"
page_title: "FortiOS: fortios_authentication_setting"
description: |-
  Configure authentication setting.
---

## fortios_authentication_setting
Configure authentication setting.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `active_auth_scheme` - Active authentication method (scheme name). This attribute must reference one of the following datasources: `authentication.scheme.name` .
* `auth_https` - Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable` `disable` .
* `captive_portal` - Captive portal host name. This attribute must reference one of the following datasources: `firewall.address.name` .
* `captive_portal_ip` - Captive portal IP address.
* `captive_portal_ip6` - Captive portal IPv6 address.
* `captive_portal_port` - Captive portal port number (1 - 65535, default = 7830).
* `captive_portal_ssl_port` - Captive portal SSL port number (1 - 65535, default = 7831).
* `captive_portal_type` - Captive portal type. Valid values: `fqdn` `ip` .
* `captive_portal6` - IPv6 captive portal host name. This attribute must reference one of the following datasources: `firewall.address6.name` .
* `cert_auth` - Enable/disable redirecting certificate authentication to HTTPS portal. Valid values: `enable` `disable` .
* `cert_captive_portal` - Certificate captive portal host name. This attribute must reference one of the following datasources: `firewall.address.name` .
* `cert_captive_portal_ip` - Certificate captive portal IP address.
* `cert_captive_portal_port` - Certificate captive portal port number (1 - 65535, default = 7832).
* `sso_auth_scheme` - Single-Sign-On authentication method (scheme name). This attribute must reference one of the following datasources: `authentication.scheme.name` .
* `dev_range` - Address range for the IP based device query. The structure of `dev_range` block is documented below.

The `dev_range` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `user_cert_ca` - CA certificate used for client certificate verification. The structure of `user_cert_ca` block is documented below.

The `user_cert_ca` block contains:

* `name` - CA certificate list. This attribute must reference one of the following datasources: `vpn.certificate.ca.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_authentication_setting can be imported using:
```sh
terraform import fortios_authentication_setting.labelname {{mkey}}
```
