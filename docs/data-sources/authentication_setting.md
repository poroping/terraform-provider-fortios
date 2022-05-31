---
subcategory: "FortiGate Authentication"
layout: "fortios"
page_title: "FortiOS: fortios_authentication_setting"
description: |-
  Get information on a fortios Configure authentication setting.
---

# Data Source: fortios_authentication_setting
Use this data source to get information on a fortios Configure authentication setting.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `active_auth_scheme` - Active authentication method (scheme name).
* `auth_https` - Enable/disable redirecting HTTP user authentication to HTTPS.
* `captive_portal` - Captive portal host name.
* `captive_portal_ip` - Captive portal IP address.
* `captive_portal_ip6` - Captive portal IPv6 address.
* `captive_portal_port` - Captive portal port number (1 - 65535, default = 7830).
* `captive_portal_ssl_port` - Captive portal SSL port number (1 - 65535, default = 7831).
* `captive_portal_type` - Captive portal type.
* `captive_portal6` - IPv6 captive portal host name.
* `cert_auth` - Enable/disable redirecting certificate authentication to HTTPS portal.
* `cert_captive_portal` - Certificate captive portal host name.
* `cert_captive_portal_ip` - Certificate captive portal IP address.
* `cert_captive_portal_port` - Certificate captive portal port number (1 - 65535, default = 7832).
* `cookie_max_age` - Persistent web portal cookie maximum age in minutes (30 - 10080 (1 week), default = 480 (8 hours)).
* `cookie_refresh_div` - Refresh rate divider of persistent web portal cookie (default = 2). Refresh value = cookie-max-age/cookie-refresh-div.
* `ip_auth_cookie` - Enable/disable persistent cookie on IP based web portal authentication (default = disable).
* `persistent_cookie` - Enable/disable persistent cookie on web portal authentication (default = enable).
* `sso_auth_scheme` - Single-Sign-On authentication method (scheme name).
* `update_time` - Time of the last update.
* `dev_range` - Address range for the IP based device query.The structure of `dev_range` block is documented below.

The `dev_range` block contains:

* `name` - Address name.
* `user_cert_ca` - CA certificate used for client certificate verification.The structure of `user_cert_ca` block is documented below.

The `user_cert_ca` block contains:

* `name` - CA certificate list.
