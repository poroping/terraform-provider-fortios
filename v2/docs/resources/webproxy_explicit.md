---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_explicit"
description: |-
  Configure explicit Web proxy settings.
---

## fortios_webproxy_explicit
Configure explicit Web proxy settings.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `ftp_incoming_port` - Accept incoming FTP-over-HTTP requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
* `ftp_over_http` - Enable to proxy FTP-over-HTTP sessions sent from a web browser. Valid values: `enable` `disable` .
* `http_incoming_port` - Accept incoming HTTP requests on one or more ports (0 - 65535, default = 8080).
* `https_incoming_port` - Accept incoming HTTPS requests on one or more ports (0 - 65535, default = 0, use the same as HTTP).
* `https_replacement_message` - Enable/disable sending the client a replacement message for HTTPS requests. Valid values: `enable` `disable` .
* `incoming_ip` - Restrict the explicit HTTP proxy to only accept sessions from this IP address. An interface must have this IP address.
* `incoming_ip6` - Restrict the explicit web proxy to only accept sessions from this IPv6 address. An interface must have this IPv6 address.
* `ipv6_status` - Enable/disable allowing an IPv6 web proxy destination in policies and all IPv6 related entries in this command. Valid values: `enable` `disable` .
* `message_upon_server_error` - Enable/disable displaying a replacement message when a server error is detected. Valid values: `enable` `disable` .
* `outgoing_ip` - Outgoing HTTP requests will have this IP address as their source address. An interface must have this IP address.
* `outgoing_ip6` - Outgoing HTTP requests will leave this IPv6. Multiple interfaces can be specified. Interfaces must have these IPv6 addresses.
* `pac_file_data` - PAC file contents enclosed in quotes (maximum of 256K bytes).
* `pac_file_name` - Pac file name.
* `pac_file_server_port` - Port number that PAC traffic from client web browsers uses to connect to the explicit web proxy (0 - 65535, default = 0; use the same as HTTP).
* `pac_file_server_status` - Enable/disable Proxy Auto-Configuration (PAC) for users of this explicit proxy profile. Valid values: `enable` `disable` .
* `pac_file_url` - PAC file access URL.
* `pref_dns_result` - Prefer resolving addresses using the configured IPv4 or IPv6 DNS server (default = ipv4). Valid values: `ipv4` `ipv6` .
* `realm` - Authentication realm used to identify the explicit web proxy (maximum of 63 characters).
* `sec_default_action` - Accept or deny explicit web proxy sessions when no web proxy firewall policy exists. Valid values: `accept` `deny` .
* `socks` - Enable/disable the SOCKS proxy. Valid values: `enable` `disable` .
* `socks_incoming_port` - Accept incoming SOCKS proxy requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
* `ssl_algorithm` - Relative strength of encryption algorithms accepted in HTTPS deep scan: high, medium, or low. Valid values: `high` `medium` `low` .
* `status` - Enable/disable the explicit Web proxy for HTTP and HTTPS session. Valid values: `enable` `disable` .
* `strict_guest` - Enable/disable strict guest user checking by the explicit web proxy. Valid values: `enable` `disable` .
* `trace_auth_no_rsp` - Enable/disable logging timed-out authentication requests. Valid values: `enable` `disable` .
* `unknown_http_version` - How to handle HTTP sessions that do not comply with HTTP 0.9, 1.0, or 1.1. Valid values: `reject` `best-effort` .
* `pac_policy` - PAC policies. The structure of `pac_policy` block is documented below.

The `pac_policy` block contains:

* `comments` - Optional comments.
* `pac_file_data` - PAC file contents enclosed in quotes (maximum of 256K bytes).
* `pac_file_name` - Pac file name.
* `policyid` - Policy ID.
* `status` - Enable/disable policy. Valid values: `enable` `disable` .
* `dstaddr` - Destination address objects. The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `srcaddr` - Source address objects. The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` `firewall.proxy-address.name` `firewall.proxy-addrgrp.name` .
* `srcaddr6` - Source address6 objects. The structure of `srcaddr6` block is documented below.

The `srcaddr6` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_webproxy_explicit can be imported using:
```sh
terraform import fortios_webproxy_explicit.labelname {{mkey}}
```
