---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_setting"
description: |-
  Get information on a fortios Configure user authentication setting.
---

# Data Source: fortios_user_setting
Use this data source to get information on a fortios Configure user authentication setting.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `auth_blackout_time` - Time in seconds an IP address is denied access after failing to authenticate five times within one minute.
* `auth_ca_cert` - HTTPS CA certificate for policy authentication.
* `auth_cert` - HTTPS server certificate for policy authentication.
* `auth_http_basic` - Enable/disable use of HTTP basic authentication for identity-based firewall policies.
* `auth_invalid_max` - Maximum number of failed authentication attempts before the user is blocked.
* `auth_lockout_duration` - Lockout period in seconds after too many login failures.
* `auth_lockout_threshold` - Maximum number of failed login attempts before login lockout is triggered.
* `auth_on_demand` - Always/implicitly trigger firewall authentication on demand.
* `auth_portal_timeout` - Time in minutes before captive portal user have to re-authenticate (1 - 30 min, default 3 min).
* `auth_secure_http` - Enable/disable redirecting HTTP user authentication to more secure HTTPS.
* `auth_src_mac` - Enable/disable source MAC for user identity.
* `auth_ssl_allow_renegotiation` - Allow/forbid SSL re-negotiation for HTTPS authentication.
* `auth_ssl_max_proto_version` - Maximum supported protocol version for SSL/TLS connections (default is no limit).
* `auth_ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `auth_ssl_sigalgs` - Set signature algorithms related to HTTPS authentication (affects TLS version <= 1.2 only, default is to enable all).
* `auth_timeout` - Time in minutes before the firewall user authentication timeout requires the user to re-authenticate.
* `auth_timeout_type` - Control if authenticated users have to login again after a hard timeout, after an idle timeout, or after a session timeout.
* `auth_type` - Supported firewall policy authentication protocols/methods.
* `per_policy_disclaimer` - Enable/disable per policy disclaimer.
* `radius_ses_timeout_act` - Set the RADIUS session timeout to a hard timeout or to ignore RADIUS server session timeouts.
* `auth_ports` - Set up non-standard ports for authentication with HTTP, HTTPS, FTP, and TELNET.The structure of `auth_ports` block is documented below.

The `auth_ports` block contains:

* `id` - ID.
* `port` - Non-standard port for firewall user authentication.
* `type` - Service type.
