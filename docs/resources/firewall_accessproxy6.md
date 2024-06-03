---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_accessproxy6"
description: |-
  Configure IPv6 access proxy.
---

## fortios_firewall_accessproxy6
Configure IPv6 access proxy.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `add_vhost_domain_to_dnsdb` - Enable/disable adding vhost/domain to dnsdb for ztna dox tunnel. Valid values: `enable` `disable` .
* `auth_portal` - Enable/disable authentication portal. Valid values: `disable` `enable` .
* `auth_virtual_host` - Virtual host for authentication portal. This attribute must reference one of the following datasources: `firewall.access-proxy-virtual-host.name` .
* `client_cert` - Enable/disable to request client certificate. Valid values: `disable` `enable` .
* `decrypted_traffic_mirror` - Decrypted traffic mirror. This attribute must reference one of the following datasources: `firewall.decrypted-traffic-mirror.name` .
* `empty_cert_action` - Action of an empty client certificate. Valid values: `accept` `block` `accept-unmanageable` .
* `http_supported_max_version` - Maximum supported HTTP versions. default = HTTP2 Valid values: `http1` `http2` .
* `log_blocked_traffic` - Enable/disable logging of blocked traffic. Valid values: `enable` `disable` .
* `name` - Access Proxy name.
* `svr_pool_multiplex` - Enable/disable server pool multiplexing. Share connected server in HTTP, HTTPS, and web-portal api-gateway. Valid values: `enable` `disable` .
* `svr_pool_server_max_request` - Maximum number of requests that servers in server pool handle before disconnecting (default = unlimited).
* `svr_pool_ttl` - Time-to-live in the server pool for idle connections to servers.
* `user_agent_detect` - Enable/disable to detect device type by HTTP user-agent if no client certificate provided. Valid values: `disable` `enable` .
* `vip` - Virtual IP name. This attribute must reference one of the following datasources: `firewall.vip6.name` .
* `api_gateway` - Set IPv4 API Gateway. The structure of `api_gateway` block is documented below.

The `api_gateway` block contains:

* `http_cookie_age` - Time in minutes that client web browsers should keep a cookie. Default is 60 minutes. 0 = no time limit.
* `http_cookie_domain` - Domain that HTTP cookie persistence should apply to.
* `http_cookie_domain_from_host` - Enable/disable use of HTTP cookie domain from host field in HTTP. Valid values: `disable` `enable` .
* `http_cookie_generation` - Generation of HTTP cookie to be accepted. Changing invalidates all existing cookies.
* `http_cookie_path` - Limit HTTP cookie persistence to the specified path.
* `http_cookie_share` - Control sharing of cookies across API Gateway. Use of same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing. Valid values: `disable` `same-ip` .
* `https_cookie_secure` - Enable/disable verification that inserted HTTPS cookies are secure. Valid values: `disable` `enable` .
* `id` - API Gateway ID.
* `ldb_method` - Method used to distribute sessions to real servers. Valid values: `static` `round-robin` `weighted` `first-alive` `http-host` .
* `persistence` - Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session. Valid values: `none` `http-cookie` .
* `saml_redirect` - Enable/disable SAML redirection after successful authentication. Valid values: `disable` `enable` .
* `saml_server` - SAML service provider configuration for VIP authentication. This attribute must reference one of the following datasources: `user.saml.name` .
* `service` - Service. Valid values: `http` `https` `tcp-forwarding` `samlsp` `web-portal` `saas` .
* `ssl_algorithm` - Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength. Valid values: `high` `medium` `low` .
* `ssl_dh_bits` - Number of bits to use in the Diffie-Hellman exchange for RSA encryption of SSL sessions. Valid values: `768` `1024` `1536` `2048` `3072` `4096` .
* `ssl_max_version` - Highest SSL/TLS version acceptable from a server. Valid values: `tls-1.0` `tls-1.1` `tls-1.2` `tls-1.3` .
* `ssl_min_version` - Lowest SSL/TLS version acceptable from a server. Valid values: `tls-1.0` `tls-1.1` `tls-1.2` `tls-1.3` .
* `ssl_renegotiation` - Enable/disable secure renegotiation to comply with RFC 5746. Valid values: `enable` `disable` .
* `ssl_vpn_web_portal` - SSL-VPN web portal. This attribute must reference one of the following datasources: `vpn.ssl.web.portal.name` .
* `url_map` - URL pattern to match.
* `url_map_type` - Type of url-map. Valid values: `sub-string` `wildcard` `regex` .
* `virtual_host` - Virtual host. This attribute must reference one of the following datasources: `firewall.access-proxy-virtual-host.name` .
* `application` - SaaS application controlled by this Access Proxy. The structure of `application` block is documented below.

The `application` block contains:

* `name` - SaaS application name.
* `realservers` - Select the real servers that this Access Proxy will distribute traffic to. The structure of `realservers` block is documented below.

The `realservers` block contains:

* `addr_type` - Type of address. Valid values: `ip` `fqdn` .
* `address` - Address or address group of the real server. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `domain` - Wildcard domain name of the real server.
* `health_check` - Enable to check the responsiveness of the real server before forwarding traffic. Valid values: `disable` `enable` .
* `health_check_proto` - Protocol of the health check monitor to use when polling to determine server's connectivity status. Valid values: `ping` `http` `tcp-connect` .
* `holddown_interval` - Enable/disable holddown timer. Server will be considered active and reachable once the holddown period has expired (30 seconds). Valid values: `enable` `disable` .
* `http_host` - HTTP server domain name in HTTP header.
* `id` - Real server ID.
* `ip` - IP address of the real server.
* `mappedport` - Port for communicating with the real server.
* `port` - Port for communicating with the real server.
* `ssh_client_cert` - Set access-proxy SSH client certificate profile. This attribute must reference one of the following datasources: `firewall.access-proxy-ssh-client-cert.name` .
* `ssh_host_key_validation` - Enable/disable SSH real server host key validation. Valid values: `disable` `enable` .
* `status` - Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent. Valid values: `active` `standby` `disable` .
* `translate_host` - Enable/disable translation of hostname/IP from virtual server to real server. Valid values: `enable` `disable` .
* `type` - TCP forwarding server type. Valid values: `tcp-forwarding` `ssh` .
* `weight` - Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.
* `ssh_host_key` - One or more server host key. The structure of `ssh_host_key` block is documented below.

The `ssh_host_key` block contains:

* `name` - Server host key name. This attribute must reference one of the following datasources: `firewall.ssh.host-key.name` .
* `ssl_cipher_suites` - SSL/TLS cipher suites to offer to a server, ordered by priority. The structure of `ssl_cipher_suites` block is documented below.

The `ssl_cipher_suites` block contains:

* `cipher` - Cipher suite name. Valid values: `TLS-AES-128-GCM-SHA256` `TLS-AES-256-GCM-SHA384` `TLS-CHACHA20-POLY1305-SHA256` `TLS-ECDHE-RSA-WITH-CHACHA20-POLY1305-SHA256` `TLS-ECDHE-ECDSA-WITH-CHACHA20-POLY1305-SHA256` `TLS-DHE-RSA-WITH-CHACHA20-POLY1305-SHA256` `TLS-DHE-RSA-WITH-AES-128-CBC-SHA` `TLS-DHE-RSA-WITH-AES-256-CBC-SHA` `TLS-DHE-RSA-WITH-AES-128-CBC-SHA256` `TLS-DHE-RSA-WITH-AES-128-GCM-SHA256` `TLS-DHE-RSA-WITH-AES-256-CBC-SHA256` `TLS-DHE-RSA-WITH-AES-256-GCM-SHA384` `TLS-DHE-DSS-WITH-AES-128-CBC-SHA` `TLS-DHE-DSS-WITH-AES-256-CBC-SHA` `TLS-DHE-DSS-WITH-AES-128-CBC-SHA256` `TLS-DHE-DSS-WITH-AES-128-GCM-SHA256` `TLS-DHE-DSS-WITH-AES-256-CBC-SHA256` `TLS-DHE-DSS-WITH-AES-256-GCM-SHA384` `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA` `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA256` `TLS-ECDHE-RSA-WITH-AES-128-GCM-SHA256` `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA` `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA384` `TLS-ECDHE-RSA-WITH-AES-256-GCM-SHA384` `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA` `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA256` `TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256` `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA` `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA384` `TLS-ECDHE-ECDSA-WITH-AES-256-GCM-SHA384` `TLS-RSA-WITH-AES-128-CBC-SHA` `TLS-RSA-WITH-AES-256-CBC-SHA` `TLS-RSA-WITH-AES-128-CBC-SHA256` `TLS-RSA-WITH-AES-128-GCM-SHA256` `TLS-RSA-WITH-AES-256-CBC-SHA256` `TLS-RSA-WITH-AES-256-GCM-SHA384` `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA` `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA` `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA256` `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA256` `TLS-DHE-RSA-WITH-3DES-EDE-CBC-SHA` `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA` `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA` `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA` `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA` `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA256` `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA256` `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA256` `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA256` `TLS-DHE-RSA-WITH-SEED-CBC-SHA` `TLS-DHE-DSS-WITH-SEED-CBC-SHA` `TLS-DHE-RSA-WITH-ARIA-128-CBC-SHA256` `TLS-DHE-RSA-WITH-ARIA-256-CBC-SHA384` `TLS-DHE-DSS-WITH-ARIA-128-CBC-SHA256` `TLS-DHE-DSS-WITH-ARIA-256-CBC-SHA384` `TLS-RSA-WITH-SEED-CBC-SHA` `TLS-RSA-WITH-ARIA-128-CBC-SHA256` `TLS-RSA-WITH-ARIA-256-CBC-SHA384` `TLS-ECDHE-RSA-WITH-ARIA-128-CBC-SHA256` `TLS-ECDHE-RSA-WITH-ARIA-256-CBC-SHA384` `TLS-ECDHE-ECDSA-WITH-ARIA-128-CBC-SHA256` `TLS-ECDHE-ECDSA-WITH-ARIA-256-CBC-SHA384` `TLS-ECDHE-RSA-WITH-RC4-128-SHA` `TLS-ECDHE-RSA-WITH-3DES-EDE-CBC-SHA` `TLS-DHE-DSS-WITH-3DES-EDE-CBC-SHA` `TLS-RSA-WITH-3DES-EDE-CBC-SHA` `TLS-RSA-WITH-RC4-128-MD5` `TLS-RSA-WITH-RC4-128-SHA` `TLS-DHE-RSA-WITH-DES-CBC-SHA` `TLS-DHE-DSS-WITH-DES-CBC-SHA` `TLS-RSA-WITH-DES-CBC-SHA` .
* `priority` - SSL/TLS cipher suites priority.
* `versions` - SSL/TLS versions that the cipher suite can be used with. Valid values: `tls-1.0` `tls-1.1` `tls-1.2` `tls-1.3` .
* `api_gateway6` - Set IPv6 API Gateway. The structure of `api_gateway6` block is documented below.

The `api_gateway6` block contains:

* `http_cookie_age` - Time in minutes that client web browsers should keep a cookie. Default is 60 minutes. 0 = no time limit.
* `http_cookie_domain` - Domain that HTTP cookie persistence should apply to.
* `http_cookie_domain_from_host` - Enable/disable use of HTTP cookie domain from host field in HTTP. Valid values: `disable` `enable` .
* `http_cookie_generation` - Generation of HTTP cookie to be accepted. Changing invalidates all existing cookies.
* `http_cookie_path` - Limit HTTP cookie persistence to the specified path.
* `http_cookie_share` - Control sharing of cookies across API Gateway. Use of same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing. Valid values: `disable` `same-ip` .
* `https_cookie_secure` - Enable/disable verification that inserted HTTPS cookies are secure. Valid values: `disable` `enable` .
* `id` - API Gateway ID.
* `ldb_method` - Method used to distribute sessions to real servers. Valid values: `static` `round-robin` `weighted` `first-alive` `http-host` .
* `persistence` - Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session. Valid values: `none` `http-cookie` .
* `saml_redirect` - Enable/disable SAML redirection after successful authentication. Valid values: `disable` `enable` .
* `saml_server` - SAML service provider configuration for VIP authentication. This attribute must reference one of the following datasources: `user.saml.name` .
* `service` - Service. Valid values: `http` `https` `tcp-forwarding` `samlsp` `web-portal` `saas` .
* `ssl_algorithm` - Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength. Valid values: `high` `medium` `low` .
* `ssl_dh_bits` - Number of bits to use in the Diffie-Hellman exchange for RSA encryption of SSL sessions. Valid values: `768` `1024` `1536` `2048` `3072` `4096` .
* `ssl_max_version` - Highest SSL/TLS version acceptable from a server. Valid values: `tls-1.0` `tls-1.1` `tls-1.2` `tls-1.3` .
* `ssl_min_version` - Lowest SSL/TLS version acceptable from a server. Valid values: `tls-1.0` `tls-1.1` `tls-1.2` `tls-1.3` .
* `ssl_renegotiation` - Enable/disable secure renegotiation to comply with RFC 5746. Valid values: `enable` `disable` .
* `ssl_vpn_web_portal` - SSL-VPN web portal. This attribute must reference one of the following datasources: `vpn.ssl.web.portal.name` .
* `url_map` - URL pattern to match.
* `url_map_type` - Type of url-map. Valid values: `sub-string` `wildcard` `regex` .
* `virtual_host` - Virtual host. This attribute must reference one of the following datasources: `firewall.access-proxy-virtual-host.name` .
* `application` - SaaS application controlled by this Access Proxy. The structure of `application` block is documented below.

The `application` block contains:

* `name` - SaaS application name.
* `realservers` - Select the real servers that this Access Proxy will distribute traffic to. The structure of `realservers` block is documented below.

The `realservers` block contains:

* `addr_type` - Type of address. Valid values: `ip` `fqdn` .
* `address` - Address or address group of the real server. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `domain` - Wildcard domain name of the real server.
* `health_check` - Enable to check the responsiveness of the real server before forwarding traffic. Valid values: `disable` `enable` .
* `health_check_proto` - Protocol of the health check monitor to use when polling to determine server's connectivity status. Valid values: `ping` `http` `tcp-connect` .
* `holddown_interval` - Enable/disable holddown timer. Server will be considered active and reachable once the holddown period has expired (30 seconds). Valid values: `enable` `disable` .
* `http_host` - HTTP server domain name in HTTP header.
* `id` - Real server ID.
* `ip` - IPv6 address of the real server.
* `mappedport` - Port for communicating with the real server.
* `port` - Port for communicating with the real server.
* `ssh_client_cert` - Set access-proxy SSH client certificate profile. This attribute must reference one of the following datasources: `firewall.access-proxy-ssh-client-cert.name` .
* `ssh_host_key_validation` - Enable/disable SSH real server host key validation. Valid values: `disable` `enable` .
* `status` - Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent. Valid values: `active` `standby` `disable` .
* `translate_host` - Enable/disable translation of hostname/IP from virtual server to real server. Valid values: `enable` `disable` .
* `type` - TCP forwarding server type. Valid values: `tcp-forwarding` `ssh` .
* `weight` - Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.
* `ssh_host_key` - One or more server host key. The structure of `ssh_host_key` block is documented below.

The `ssh_host_key` block contains:

* `name` - Server host key name. This attribute must reference one of the following datasources: `firewall.ssh.host-key.name` .
* `ssl_cipher_suites` - SSL/TLS cipher suites to offer to a server, ordered by priority. The structure of `ssl_cipher_suites` block is documented below.

The `ssl_cipher_suites` block contains:

* `cipher` - Cipher suite name. Valid values: `TLS-AES-128-GCM-SHA256` `TLS-AES-256-GCM-SHA384` `TLS-CHACHA20-POLY1305-SHA256` `TLS-ECDHE-RSA-WITH-CHACHA20-POLY1305-SHA256` `TLS-ECDHE-ECDSA-WITH-CHACHA20-POLY1305-SHA256` `TLS-DHE-RSA-WITH-CHACHA20-POLY1305-SHA256` `TLS-DHE-RSA-WITH-AES-128-CBC-SHA` `TLS-DHE-RSA-WITH-AES-256-CBC-SHA` `TLS-DHE-RSA-WITH-AES-128-CBC-SHA256` `TLS-DHE-RSA-WITH-AES-128-GCM-SHA256` `TLS-DHE-RSA-WITH-AES-256-CBC-SHA256` `TLS-DHE-RSA-WITH-AES-256-GCM-SHA384` `TLS-DHE-DSS-WITH-AES-128-CBC-SHA` `TLS-DHE-DSS-WITH-AES-256-CBC-SHA` `TLS-DHE-DSS-WITH-AES-128-CBC-SHA256` `TLS-DHE-DSS-WITH-AES-128-GCM-SHA256` `TLS-DHE-DSS-WITH-AES-256-CBC-SHA256` `TLS-DHE-DSS-WITH-AES-256-GCM-SHA384` `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA` `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA256` `TLS-ECDHE-RSA-WITH-AES-128-GCM-SHA256` `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA` `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA384` `TLS-ECDHE-RSA-WITH-AES-256-GCM-SHA384` `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA` `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA256` `TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256` `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA` `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA384` `TLS-ECDHE-ECDSA-WITH-AES-256-GCM-SHA384` `TLS-RSA-WITH-AES-128-CBC-SHA` `TLS-RSA-WITH-AES-256-CBC-SHA` `TLS-RSA-WITH-AES-128-CBC-SHA256` `TLS-RSA-WITH-AES-128-GCM-SHA256` `TLS-RSA-WITH-AES-256-CBC-SHA256` `TLS-RSA-WITH-AES-256-GCM-SHA384` `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA` `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA` `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA256` `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA256` `TLS-DHE-RSA-WITH-3DES-EDE-CBC-SHA` `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA` `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA` `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA` `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA` `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA256` `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA256` `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA256` `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA256` `TLS-DHE-RSA-WITH-SEED-CBC-SHA` `TLS-DHE-DSS-WITH-SEED-CBC-SHA` `TLS-DHE-RSA-WITH-ARIA-128-CBC-SHA256` `TLS-DHE-RSA-WITH-ARIA-256-CBC-SHA384` `TLS-DHE-DSS-WITH-ARIA-128-CBC-SHA256` `TLS-DHE-DSS-WITH-ARIA-256-CBC-SHA384` `TLS-RSA-WITH-SEED-CBC-SHA` `TLS-RSA-WITH-ARIA-128-CBC-SHA256` `TLS-RSA-WITH-ARIA-256-CBC-SHA384` `TLS-ECDHE-RSA-WITH-ARIA-128-CBC-SHA256` `TLS-ECDHE-RSA-WITH-ARIA-256-CBC-SHA384` `TLS-ECDHE-ECDSA-WITH-ARIA-128-CBC-SHA256` `TLS-ECDHE-ECDSA-WITH-ARIA-256-CBC-SHA384` `TLS-ECDHE-RSA-WITH-RC4-128-SHA` `TLS-ECDHE-RSA-WITH-3DES-EDE-CBC-SHA` `TLS-DHE-DSS-WITH-3DES-EDE-CBC-SHA` `TLS-RSA-WITH-3DES-EDE-CBC-SHA` `TLS-RSA-WITH-RC4-128-MD5` `TLS-RSA-WITH-RC4-128-SHA` `TLS-DHE-RSA-WITH-DES-CBC-SHA` `TLS-DHE-DSS-WITH-DES-CBC-SHA` `TLS-RSA-WITH-DES-CBC-SHA` .
* `priority` - SSL/TLS cipher suites priority.
* `versions` - SSL/TLS versions that the cipher suite can be used with. Valid values: `tls-1.0` `tls-1.1` `tls-1.2` `tls-1.3` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_accessproxy6 can be imported using:
```sh
terraform import fortios_firewall_accessproxy6.labelname {{mkey}}
```
