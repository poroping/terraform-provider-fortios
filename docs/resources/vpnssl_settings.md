---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnssl_settings"
description: |-
  Configure SSL-VPN.
---

## fortios_vpnssl_settings
Configure SSL-VPN.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `algorithm` - Force the SSL-VPN security level. High allows only high. Medium allows medium and high. Low allows any. Valid values: `high` `medium` `default` `low` .
* `auth_session_check_source_ip` - Enable/disable checking of source IP for authentication session. Valid values: `enable` `disable` .
* `auth_timeout` - SSL-VPN authentication timeout (1 - 259200 sec (3 days), 0 for no timeout).
* `auto_tunnel_static_route` - Enable/disable to auto-create static routes for the SSL-VPN tunnel IP addresses. Valid values: `enable` `disable` .
* `banned_cipher` - Select one or more cipher technologies that cannot be used in SSL-VPN negotiations. Only applies to TLS 1.2 and below. Valid values: `RSA` `DHE` `ECDHE` `DSS` `ECDSA` `AES` `AESGCM` `CAMELLIA` `3DES` `SHA1` `SHA256` `SHA384` `STATIC` `CHACHA20` `ARIA` `AESCCM` .
* `browser_language_detection` - Enable/disable overriding the configured system language based on the preferred language of the browser. Valid values: `enable` `disable` .
* `check_referer` - Enable/disable verification of referer field in HTTP request header. Valid values: `enable` `disable` .
* `ciphersuite` - Select one or more TLS 1.3 ciphersuites to enable. Does not affect ciphers in TLS 1.2 and below. At least one must be enabled. To disable all, set ssl-max-proto-ver to tls1-2 or below. Valid values: `TLS-AES-128-GCM-SHA256` `TLS-AES-256-GCM-SHA384` `TLS-CHACHA20-POLY1305-SHA256` `TLS-AES-128-CCM-SHA256` `TLS-AES-128-CCM-8-SHA256` .
* `client_sigalgs` - Set signature algorithms related to client authentication. Affects TLS version <= 1.2 only. Valid values: `no-rsa-pss` `all` .
* `default_portal` - Default SSL-VPN portal. This attribute must reference one of the following datasources: `vpn.ssl.web.portal.name` .
* `deflate_compression_level` - Compression level (0~9).
* `deflate_min_data_size` - Minimum amount of data that triggers compression (200 - 65535 bytes).
* `dns_server1` - DNS server 1.
* `dns_server2` - DNS server 2.
* `dns_suffix` - DNS suffix used for SSL-VPN clients.
* `dtls_hello_timeout` - SSLVPN maximum DTLS hello timeout (10 - 60 sec, default = 10).
* `dtls_max_proto_ver` - DTLS maximum protocol version. Valid values: `dtls1-0` `dtls1-2` .
* `dtls_min_proto_ver` - DTLS minimum protocol version. Valid values: `dtls1-0` `dtls1-2` .
* `dtls_tunnel` - Enable/disable DTLS to prevent eavesdropping, tampering, or message forgery. Valid values: `enable` `disable` .
* `dual_stack_mode` - Tunnel mode: enable parallel IPv4 and IPv6 tunnel. Web mode: support IPv4 and IPv6 bookmarks in the portal. Valid values: `enable` `disable` .
* `encode_2f_sequence` - Encode \2F sequence to forward slash in URLs. Valid values: `enable` `disable` .
* `encrypt_and_store_password` - Encrypt and store user passwords for SSL-VPN web sessions. Valid values: `enable` `disable` .
* `force_two_factor_auth` - Enable/disable only PKI users with two-factor authentication for SSL-VPNs. Valid values: `enable` `disable` .
* `header_x_forwarded_for` - Forward the same, add, or remove HTTP header. Valid values: `pass` `add` `remove` .
* `hsts_include_subdomains` - Add HSTS includeSubDomains response header. Valid values: `enable` `disable` .
* `http_compression` - Enable/disable to allow HTTP compression over SSL-VPN tunnels. Valid values: `enable` `disable` .
* `http_only_cookie` - Enable/disable SSL-VPN support for HttpOnly cookies. Valid values: `enable` `disable` .
* `http_request_body_timeout` - SSL-VPN session is disconnected if an HTTP request body is not received within this time (1 - 60 sec, default = 20).
* `http_request_header_timeout` - SSL-VPN session is disconnected if an HTTP request header is not received within this time (1 - 60 sec, default = 20).
* `https_redirect` - Enable/disable redirect of port 80 to SSL-VPN port. Valid values: `enable` `disable` .
* `idle_timeout` - SSL-VPN disconnects if idle for specified time in seconds.
* `ipv6_dns_server1` - IPv6 DNS server 1.
* `ipv6_dns_server2` - IPv6 DNS server 2.
* `ipv6_wins_server1` - IPv6 WINS server 1.
* `ipv6_wins_server2` - IPv6 WINS server 2.
* `login_attempt_limit` - SSL-VPN maximum login attempt times before block (0 - 10, default = 2, 0 = no limit).
* `login_block_time` - Time for which a user is blocked from logging in after too many failed login attempts (0 - 86400 sec, default = 60).
* `login_timeout` - SSLVPN maximum login timeout (10 - 180 sec, default = 30).
* `port` - SSL-VPN access port (1 - 65535).
* `port_precedence` - Enable/disable, Enable means that if SSL-VPN connections are allowed on an interface admin GUI connections are blocked on that interface. Valid values: `enable` `disable` .
* `reqclientcert` - Enable/disable to require client certificates for all SSL-VPN users. Valid values: `enable` `disable` .
* `route_source_interface` - Enable to allow SSL-VPN sessions to bypass routing and bind to the incoming interface. Valid values: `enable` `disable` .
* `saml_redirect_port` - SAML local redirect port in the machine running FortiClient (0 - 65535). 0 is to disable redirection on FGT side.
* `server_hostname` - Server hostname for HTTPS. When set, will be used for SSL VPN web proxy host header for any redirection.
* `servercert` - Name of the server certificate to be used for SSL-VPNs. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `source_address_negate` - Enable/disable negated source address match. Valid values: `enable` `disable` .
* `source_address6_negate` - Enable/disable negated source IPv6 address match. Valid values: `enable` `disable` .
* `ssl_client_renegotiation` - Enable/disable to allow client renegotiation by the server if the tunnel goes down. Valid values: `disable` `enable` .
* `ssl_insert_empty_fragment` - Enable/disable insertion of empty fragment. Valid values: `enable` `disable` .
* `ssl_max_proto_ver` - SSL maximum protocol version. Valid values: `tls1-0` `tls1-1` `tls1-2` `tls1-3` .
* `ssl_min_proto_ver` - SSL minimum protocol version. Valid values: `tls1-0` `tls1-1` `tls1-2` `tls1-3` .
* `status` - Enable/disable SSL-VPN. Valid values: `enable` `disable` .
* `tlsv1_0` - tlsv1-0 Valid values: `enable` `disable` .
* `tlsv1_1` - tlsv1-1 Valid values: `enable` `disable` .
* `tlsv1_2` - tlsv1-2 Valid values: `enable` `disable` .
* `tlsv1_3` - tlsv1-3 Valid values: `enable` `disable` .
* `transform_backward_slashes` - Transform backward slashes to forward slashes in URLs. Valid values: `enable` `disable` .
* `tunnel_addr_assigned_method` - Method used for assigning address for tunnel. Valid values: `first-available` `round-robin` .
* `tunnel_connect_without_reauth` - Enable/disable tunnel connection without re-authorization if previous connection dropped. Valid values: `enable` `disable` .
* `tunnel_user_session_timeout` - Time out value to clean up user session after tunnel connection is dropped (1 - 255 sec, default=30).
* `unsafe_legacy_renegotiation` - Enable/disable unsafe legacy re-negotiation. Valid values: `enable` `disable` .
* `url_obscuration` - Enable/disable to obscure the host name of the URL of the web browser display. Valid values: `enable` `disable` .
* `user_peer` - Name of user peer. This attribute must reference one of the following datasources: `user.peer.name` .
* `web_mode_snat` - Enable/disable use of IP pools defined in firewall policy while using web-mode. Valid values: `enable` `disable` .
* `wins_server1` - WINS server 1.
* `wins_server2` - WINS server 2.
* `x_content_type_options` - Add HTTP X-Content-Type-Options header. Valid values: `enable` `disable` .
* `ztna_trusted_client` - Enable/disable verification of device certificate for SSLVPN ZTNA session. Valid values: `enable` `disable` .
* `authentication_rule` - Authentication rule for SSL-VPN. The structure of `authentication_rule` block is documented below.

The `authentication_rule` block contains:

* `auth` - SSL-VPN authentication method restriction. Valid values: `any` `local` `radius` `tacacs+` `ldap` `peer` .
* `cipher` - SSL-VPN cipher strength. Valid values: `any` `high` `medium` .
* `client_cert` - Enable/disable SSL-VPN client certificate restrictive. Valid values: `enable` `disable` .
* `id` - ID (0 - 4294967295).
* `portal` - SSL-VPN portal. This attribute must reference one of the following datasources: `vpn.ssl.web.portal.name` .
* `realm` - SSL-VPN realm. This attribute must reference one of the following datasources: `vpn.ssl.web.realm.url-path` .
* `source_address_negate` - Enable/disable negated source address match. Valid values: `enable` `disable` .
* `source_address6_negate` - Enable/disable negated source IPv6 address match. Valid values: `enable` `disable` .
* `user_peer` - Name of user peer. This attribute must reference one of the following datasources: `user.peer.name` .
* `groups` - User groups. The structure of `groups` block is documented below.

The `groups` block contains:

* `name` - Group name. This attribute must reference one of the following datasources: `user.group.name` .
* `source_address` - Source address of incoming traffic. The structure of `source_address` block is documented below.

The `source_address` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` `system.external-resource.name` .
* `source_address6` - IPv6 source address of incoming traffic. The structure of `source_address6` block is documented below.

The `source_address6` block contains:

* `name` - IPv6 address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` `system.external-resource.name` .
* `source_interface` - SSL-VPN source interface of incoming traffic. The structure of `source_interface` block is documented below.

The `source_interface` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` .
* `users` - User name. The structure of `users` block is documented below.

The `users` block contains:

* `name` - User name. This attribute must reference one of the following datasources: `user.local.name` .
* `source_address` - Source address of incoming traffic. The structure of `source_address` block is documented below.

The `source_address` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` `system.external-resource.name` .
* `source_address6` - IPv6 source address of incoming traffic. The structure of `source_address6` block is documented below.

The `source_address6` block contains:

* `name` - IPv6 address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` `system.external-resource.name` .
* `source_interface` - SSL-VPN source interface of incoming traffic. The structure of `source_interface` block is documented below.

The `source_interface` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` .
* `tunnel_ip_pools` - Names of the IPv4 IP Pool firewall objects that define the IP addresses reserved for remote clients. The structure of `tunnel_ip_pools` block is documented below.

The `tunnel_ip_pools` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `tunnel_ipv6_pools` - Names of the IPv6 IP Pool firewall objects that define the IP addresses reserved for remote clients. The structure of `tunnel_ipv6_pools` block is documented below.

The `tunnel_ipv6_pools` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_vpnssl_settings can be imported using:
```sh
terraform import fortios_vpnssl_settings.labelname {{mkey}}
```
