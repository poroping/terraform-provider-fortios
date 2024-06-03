---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_global"
description: |-
  Configure Web proxy global settings.
---

## fortios_webproxy_global
Configure Web proxy global settings.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `fast_policy_match` - Enable/disable fast matching algorithm for explicit and transparent proxy policy. Valid values: `enable` `disable` .
* `forward_proxy_auth` - Enable/disable forwarding proxy authentication headers. Valid values: `enable` `disable` .
* `forward_server_affinity_timeout` - Period of time before the source IP's traffic is no longer assigned to the forwarding server (6 - 60 min, default = 30).
* `ldap_user_cache` - Enable/disable LDAP user cache for explicit and transparent proxy user. Valid values: `enable` `disable` .
* `learn_client_ip` - Enable/disable learning the client's IP address from headers. Valid values: `enable` `disable` .
* `learn_client_ip_from_header` - Learn client IP address from the specified headers. Valid values: `true-client-ip` `x-real-ip` `x-forwarded-for` .
* `log_policy_pending` - Enable/disable logging sessions that are pending on policy matching. Valid values: `enable` `disable` .
* `max_message_length` - Maximum length of HTTP message, not including body (16 - 256 Kbytes, default = 32).
* `max_request_length` - Maximum length of HTTP request line (2 - 64 Kbytes, default = 8).
* `max_waf_body_cache_length` - Maximum length of HTTP messages processed by Web Application Firewall (WAF) (10 - 1024 Kbytes, default = 32).
* `policy_category_deep_inspect` - Enable/disable deep inspection for application level category policy matching. Valid values: `enable` `disable` .
* `proxy_fqdn` - Fully Qualified Domain Name (FQDN) that clients connect to (default = default.fqdn) to connect to the explicit web proxy.
* `src_affinity_exempt_addr` - IPv4 source addresses to exempt proxy affinity.
* `src_affinity_exempt_addr6` - IPv6 source addresses to exempt proxy affinity.
* `ssl_ca_cert` - SSL CA certificate for SSL interception. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `ssl_cert` - SSL certificate for SSL interception. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `strict_web_check` - Enable/disable strict web checking to block web sites that send incorrect headers that don't conform to HTTP 1.1. Valid values: `enable` `disable` .
* `tunnel_non_http` - Enable/disable allowing non-HTTP traffic. Allowed non-HTTP traffic is tunneled. Valid values: `enable` `disable` .
* `unknown_http_version` - Action to take when an unknown version of HTTP is encountered: reject, allow (tunnel), or proceed with best-effort. Valid values: `reject` `tunnel` `best-effort` .
* `webproxy_profile` - Name of the web proxy profile to apply when explicit proxy traffic is allowed by default and traffic is accepted that does not match an explicit proxy policy. This attribute must reference one of the following datasources: `web-proxy.profile.name` .
* `learn_client_ip_srcaddr` - Source address name (srcaddr or srcaddr6 must be set). The structure of `learn_client_ip_srcaddr` block is documented below.

The `learn_client_ip_srcaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `learn_client_ip_srcaddr6` - IPv6 Source address name (srcaddr or srcaddr6 must be set). The structure of `learn_client_ip_srcaddr6` block is documented below.

The `learn_client_ip_srcaddr6` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_webproxy_global can be imported using:
```sh
terraform import fortios_webproxy_global.labelname {{mkey}}
```
