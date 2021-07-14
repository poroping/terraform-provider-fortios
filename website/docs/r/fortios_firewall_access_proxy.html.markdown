---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_access_proxy"
description: |-
  Configure Access Proxy.
---

## fortios_firewall_access_proxy
Configure Access Proxy.

## Example Usage

WIP

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - String `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `client_cert` - Enable/disable to request client certificate. Valid values: `disable` `enable` .
* `empty_cert_action` - Action of an empty client certificate. Valid values: `accept` `block` .
* `ldb_method` - Method used to distribute sessions to SSL real servers. Valid values: `static` `round-robin` `weighted` `least-session` `least-rtt` `first-alive` .
* `name` - Access Proxy name.
* `server_pubkey_auth` - Enable/disable SSH real server public key authentication. Valid values: `disable` `enable` .
* `vip` - Virtual IP name.This attribute must reference one of the following datasources: `firewall.vip.name` .
* `api_gateway` - Set API Gateway.The structure of `api_gateway` block is documented below.

The `api_gateway` block contains:

* `http_cookie_age` - Time in minutes that client web browsers should keep a cookie. Default is 60 minutes. 0 = no time limit.
* `http_cookie_domain` - Domain that HTTP cookie persistence should apply to.
* `http_cookie_domain_from_host` - Enable/disable use of HTTP cookie domain from host field in HTTP. Valid values: `disable` `enable` .
* `http_cookie_generation` - Generation of HTTP cookie to be accepted. Changing invalidates all existing cookies.
* `http_cookie_path` - Limit HTTP cookie persistence to the specified path.
* `http_cookie_share` - Control sharing of cookies across API Gateway. same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing. Valid values: `disable` `same-ip` .
* `https_cookie_secure` - Enable/disable verification that inserted HTTPS cookies are secure. Valid values: `disable` `enable` .
* `id` - API Gateway ID.
* `ldb_method` - Method used to distribute sessions to real servers. Valid values: `static` `round-robin` `weighted` `least-session` `least-rtt` `first-alive` `http-host` .
* `persistence` - Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session. Valid values: `none` `http-cookie` .
* `saml_server` - SAML service provider configuration for VIP authentication.This attribute must reference one of the following datasources: `user.saml.name` .
* `service` - Service. Valid values: `http` `https` `tcp-forwarding` `samlsp` .
* `ssl_algorithm` - Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength. Valid values: `high` `medium` `low` `custom` .
* `ssl_dh_bits` - Number of bits to use in the Diffie-Hellman exchange for RSA encryption of SSL sessions. Valid values: `768` `1024` `1536` `2048` `3072` `4096` .
* `ssl_max_version` - Highest SSL/TLS version acceptable from a server. Valid values: `tls-1.0` `tls-1.1` `tls-1.2` `tls-1.3` .
* `ssl_min_version` - Lowest SSL/TLS version acceptable from a server. Valid values: `tls-1.0` `tls-1.1` `tls-1.2` `tls-1.3` .
* `url_map` - URL pattern to match.
* `url_map_type` - Type of url-map. Valid values: `sub-string` `wildcard` `regex` .
* `virtual_host` - Virtual host.This attribute must reference one of the following datasources: `firewall.access-proxy-virtual-host.name` .
* `realservers` - Select the real servers that this Access Proxy will distribute traffic to.The structure of `realservers` block is documented below.

The `realservers` block contains:

* `address` - Address or address group of the real server.This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `health_check` - Enable to check the responsiveness of the real server before forwarding traffic. Valid values: `disable` `enable` .
* `health_check_proto` - Protocol of the health check monitor to use when polling to determine server's connectivity status. Valid values: `ping` `http` `tcp-connect` .
* `http_host` - HTTP server domain name in HTTP header.
* `id` - Real server ID.
* `ip` - IP address of the real server.
* `mappedport` - Port for communicating with the real server.
* `port` - Port for communicating with the real server.
* `status` - Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent. Valid values: `active` `standby` `disable` .
* `weight` - Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.
* `ssl_cipher_suites` - SSL/TLS cipher suites to offer to a server, ordered by priority.The structure of `ssl_cipher_suites` block is documented below.

The `ssl_cipher_suites` block contains:

* `cipher` - Cipher suite name. Valid values: `TLS-AES-128-GCM-SHA256` `TLS-AES-256-GCM-SHA384` `TLS-CHACHA20-POLY1305-SHA256` `TLS-ECDHE-RSA-WITH-CHACHA20-POLY1305-SHA256` `TLS-ECDHE-ECDSA-WITH-CHACHA20-POLY1305-SHA256` `TLS-DHE-RSA-WITH-CHACHA20-POLY1305-SHA256` `TLS-DHE-RSA-WITH-AES-128-CBC-SHA` `TLS-DHE-RSA-WITH-AES-256-CBC-SHA` `TLS-DHE-RSA-WITH-AES-128-CBC-SHA256` `TLS-DHE-RSA-WITH-AES-128-GCM-SHA256` `TLS-DHE-RSA-WITH-AES-256-CBC-SHA256` `TLS-DHE-RSA-WITH-AES-256-GCM-SHA384` `TLS-DHE-DSS-WITH-AES-128-CBC-SHA` `TLS-DHE-DSS-WITH-AES-256-CBC-SHA` `TLS-DHE-DSS-WITH-AES-128-CBC-SHA256` `TLS-DHE-DSS-WITH-AES-128-GCM-SHA256` `TLS-DHE-DSS-WITH-AES-256-CBC-SHA256` `TLS-DHE-DSS-WITH-AES-256-GCM-SHA384` `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA` `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA256` `TLS-ECDHE-RSA-WITH-AES-128-GCM-SHA256` `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA` `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA384` `TLS-ECDHE-RSA-WITH-AES-256-GCM-SHA384` `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA` `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA256` `TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256` `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA384` `TLS-ECDHE-ECDSA-WITH-AES-256-GCM-SHA384` `TLS-RSA-WITH-AES-128-CBC-SHA` `TLS-RSA-WITH-AES-256-CBC-SHA` `TLS-RSA-WITH-AES-128-CBC-SHA256` `TLS-RSA-WITH-AES-128-GCM-SHA256` `TLS-RSA-WITH-AES-256-CBC-SHA256` `TLS-RSA-WITH-AES-256-GCM-SHA384` `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA` `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA` `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA256` `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA256` `TLS-DHE-RSA-WITH-3DES-EDE-CBC-SHA` `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA` `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA` `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA` `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA` `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA256` `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA256` `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA256` `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA256` `TLS-DHE-RSA-WITH-SEED-CBC-SHA` `TLS-DHE-DSS-WITH-SEED-CBC-SHA` `TLS-DHE-RSA-WITH-ARIA-128-CBC-SHA256` `TLS-DHE-RSA-WITH-ARIA-256-CBC-SHA384` `TLS-DHE-DSS-WITH-ARIA-128-CBC-SHA256` `TLS-DHE-DSS-WITH-ARIA-256-CBC-SHA384` `TLS-RSA-WITH-SEED-CBC-SHA` `TLS-RSA-WITH-ARIA-128-CBC-SHA256` `TLS-RSA-WITH-ARIA-256-CBC-SHA384` `TLS-ECDHE-RSA-WITH-ARIA-128-CBC-SHA256` `TLS-ECDHE-RSA-WITH-ARIA-256-CBC-SHA384` `TLS-ECDHE-ECDSA-WITH-ARIA-128-CBC-SHA256` `TLS-ECDHE-ECDSA-WITH-ARIA-256-CBC-SHA384` `TLS-ECDHE-RSA-WITH-RC4-128-SHA` `TLS-ECDHE-RSA-WITH-3DES-EDE-CBC-SHA` `TLS-DHE-DSS-WITH-3DES-EDE-CBC-SHA` `TLS-RSA-WITH-3DES-EDE-CBC-SHA` `TLS-RSA-WITH-RC4-128-MD5` `TLS-RSA-WITH-RC4-128-SHA` `TLS-DHE-RSA-WITH-DES-CBC-SHA` `TLS-DHE-DSS-WITH-DES-CBC-SHA` `TLS-RSA-WITH-DES-CBC-SHA` .
* `priority` - SSL/TLS cipher suites priority.
* `versions` - SSL/TLS versions that the cipher suite can be used with. Valid values: `tls-1.0` `tls-1.1` `tls-1.2` `tls-1.3` .
* `realservers` - Select the SSL real servers that this Access Proxy will distribute traffic to.The structure of `realservers` block is documented below.

The `realservers` block contains:

* `id` - Real server ID.
* `ip` - IP address of the real server.
* `port` - Port for communicating with the real server.
* `status` - Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent. Valid values: `active` `standby` `disable` .
* `weight` - Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.
* `server_pubkey_auth_settings` - Server SSH public key authentication settings.The structure of `server_pubkey_auth_settings` block is documented below.

The `server_pubkey_auth_settings` block contains:

* `auth_ca` - Name of the SSH server public key authentication CA.This attribute must reference one of the following datasources: `firewall.ssh.local-ca.name` .
* `permit_agent_forwarding` - Enable/disable appending permit-agent-forwarding certificate extension. Valid values: `enable` `disable` .
* `permit_port_forwarding` - Enable/disable appending permit-port-forwarding certificate extension. Valid values: `enable` `disable` .
* `permit_pty` - Enable/disable appending permit-pty certificate extension. Valid values: `enable` `disable` .
* `permit_user_rc` - Enable/disable appending permit-user-rc certificate extension. Valid values: `enable` `disable` .
* `permit_x11_forwarding` - Enable/disable appending permit-x11-forwarding certificate extension. Valid values: `enable` `disable` .
* `source_address` - Enable/disable appending source-address certificate critical option. This option ensure certificate only accepted from FortiGate source address. Valid values: `enable` `disable` .
* `cert_extension` - Configure certificate extension for user certificate.The structure of `cert_extension` block is documented below.

The `cert_extension` block contains:

* `critical` - Critical option. Valid values: `no` `yes` .
* `data` - Name of certificate extension.
* `name` - Name of certificate extension.
* `type` - Type of certificate extension. Valid values: `fixed` `user` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Address can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_address.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
