---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_setting"
description: |-
  Configure user authentication setting.
---

## fortios_user_setting
Configure user authentication setting.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `auth_blackout_time` - Time in seconds an IP address is denied access after failing to authenticate five times within one minute.
* `auth_ca_cert` - HTTPS CA certificate for policy authentication. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `auth_cert` - HTTPS server certificate for policy authentication. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `auth_http_basic` - Enable/disable use of HTTP basic authentication for identity-based firewall policies. Valid values: `enable` `disable` .
* `auth_invalid_max` - Maximum number of failed authentication attempts before the user is blocked.
* `auth_lockout_duration` - Lockout period in seconds after too many login failures.
* `auth_lockout_threshold` - Maximum number of failed login attempts before login lockout is triggered.
* `auth_on_demand` - Always/implicitly trigger firewall authentication on demand. Valid values: `always` `implicitly` .
* `auth_portal_timeout` - Time in minutes before captive portal user have to re-authenticate (1 - 30 min, default 3 min).
* `auth_secure_http` - Enable/disable redirecting HTTP user authentication to more secure HTTPS. Valid values: `enable` `disable` .
* `auth_src_mac` - Enable/disable source MAC for user identity. Valid values: `enable` `disable` .
* `auth_ssl_allow_renegotiation` - Allow/forbid SSL re-negotiation for HTTPS authentication. Valid values: `enable` `disable` .
* `auth_ssl_max_proto_version` - Maximum supported protocol version for SSL/TLS connections (default is no limit). Valid values: `sslv3` `tlsv1` `tlsv1-1` `tlsv1-2` `tlsv1-3` .
* `auth_ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default` `SSLv3` `TLSv1` `TLSv1-1` `TLSv1-2` .
* `auth_ssl_sigalgs` - Set signature algorithms related to HTTPS authentication (affects TLS version <= 1.2 only, default is to enable all). Valid values: `no-rsa-pss` `all` .
* `auth_timeout` - Time in minutes before the firewall user authentication timeout requires the user to re-authenticate.
* `auth_timeout_type` - Control if authenticated users have to login again after a hard timeout, after an idle timeout, or after a session timeout. Valid values: `idle-timeout` `hard-timeout` `new-session` .
* `auth_type` - Supported firewall policy authentication protocols/methods. Valid values: `http` `https` `ftp` `telnet` .
* `per_policy_disclaimer` - Enable/disable per policy disclaimer. Valid values: `enable` `disable` .
* `radius_ses_timeout_act` - Set the RADIUS session timeout to a hard timeout or to ignore RADIUS server session timeouts. Valid values: `hard-timeout` `ignore-timeout` .
* `auth_ports` - Set up non-standard ports for authentication with HTTP, HTTPS, FTP, and TELNET. The structure of `auth_ports` block is documented below.

The `auth_ports` block contains:

* `id` - ID.
* `port` - Non-standard port for firewall user authentication.
* `type` - Service type. Valid values: `http` `https` `ftp` `telnet` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_user_setting can be imported using:
```sh
terraform import fortios_user_setting.labelname {{mkey}}
```
