---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_vip6"
description: |-
  Get information on a fortios Configure virtual IP for IPv6.
---

# Data Source: fortios_firewall_vip6
Use this data source to get information on a fortios Configure virtual IP for IPv6.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Virtual ip6 name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `add_nat64_route` - Enable/disable adding NAT64 route.
* `arp_reply` - Enable to respond to ARP requests for this virtual IP address. Enabled by default.
* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `embedded_ipv4_address` - Enable/disable use of the lower 32 bits of the external IPv6 address as mapped IPv4 address.
* `extip` - IPv6 address or address range on the external interface that you want to map to an address or address range on the destination network.
* `extport` - Incoming port number range that you want to map to a port number range on the destination network.
* `http_cookie_age` - Time in minutes that client web browsers should keep a cookie. Default is 60 minutes. 0 = no time limit.
* `http_cookie_domain` - Domain that HTTP cookie persistence should apply to.
* `http_cookie_domain_from_host` - Enable/disable use of HTTP cookie domain from host field in HTTP.
* `http_cookie_generation` - Generation of HTTP cookie to be accepted. Changing invalidates all existing cookies.
* `http_cookie_path` - Limit HTTP cookie persistence to the specified path.
* `http_cookie_share` - Control sharing of cookies across virtual servers. Use of same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing.
* `http_ip_header` - For HTTP multiplexing, enable to add the original client IP address in the XForwarded-For HTTP header.
* `http_ip_header_name` - For HTTP multiplexing, enter a custom HTTPS header name. The original client IP address is added to this header. If empty, X-Forwarded-For is used.
* `http_multiplex` - Enable/disable HTTP multiplexing.
* `http_redirect` - Enable/disable redirection of HTTP to HTTPS.
* `https_cookie_secure` - Enable/disable verification that inserted HTTPS cookies are secure.
* `id` - Custom defined ID.
* `ipv4_mappedip` - Range of mapped IP addresses. Specify the start IP address followed by a space and the end IP address.
* `ipv4_mappedport` - IPv4 port number range on the destination network to which the external port number range is mapped.
* `ldb_method` - Method used to distribute sessions to real servers.
* `mappedip` - Mapped IPv6 address range in the format startIP-endIP.
* `mappedport` - Port number range on the destination network to which the external port number range is mapped.
* `max_embryonic_connections` - Maximum number of incomplete connections.
* `name` - Virtual ip6 name.
* `nat_source_vip` - Enable to perform SNAT on traffic from mappedip to the extip for all egress interfaces.
* `nat64` - Enable/disable DNAT64.
* `nat66` - Enable/disable DNAT66.
* `ndp_reply` - Enable/disable this FortiGate unit's ability to respond to NDP requests for this virtual IP address (default = enable).
* `outlook_web_access` - Enable to add the Front-End-Https header for Microsoft Outlook Web Access.
* `persistence` - Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session.
* `portforward` - Enable port forwarding.
* `protocol` - Protocol to use when forwarding packets.
* `server_type` - Protocol to be load balanced by the virtual server (also called the server load balance virtual IP).
* `ssl_accept_ffdhe_groups` - Enable/disable FFDHE cipher suite for SSL key exchange.
* `ssl_algorithm` - Permitted encryption algorithms for SSL sessions according to encryption strength.
* `ssl_certificate` - The name of the certificate to use for SSL handshake.
* `ssl_client_fallback` - Enable/disable support for preventing Downgrade Attacks on client connections (RFC 7507).
* `ssl_client_rekey_count` - Maximum length of data in MB before triggering a client rekey (0 = disable).
* `ssl_client_renegotiation` - Allow, deny, or require secure renegotiation of client sessions to comply with RFC 5746.
* `ssl_client_session_state_max` - Maximum number of client to FortiGate SSL session states to keep.
* `ssl_client_session_state_timeout` - Number of minutes to keep client to FortiGate SSL session state.
* `ssl_client_session_state_type` - How to expire SSL sessions for the segment of the SSL connection between the client and the FortiGate.
* `ssl_dh_bits` - Number of bits to use in the Diffie-Hellman exchange for RSA encryption of SSL sessions.
* `ssl_hpkp` - Enable/disable including HPKP header in response.
* `ssl_hpkp_age` - Number of minutes the web browser should keep HPKP.
* `ssl_hpkp_backup` - Certificate to generate backup HPKP pin from.
* `ssl_hpkp_include_subdomains` - Indicate that HPKP header applies to all subdomains.
* `ssl_hpkp_primary` - Certificate to generate primary HPKP pin from.
* `ssl_hpkp_report_uri` - URL to report HPKP violations to.
* `ssl_hsts` - Enable/disable including HSTS header in response.
* `ssl_hsts_age` - Number of seconds the client should honor the HSTS setting.
* `ssl_hsts_include_subdomains` - Indicate that HSTS header applies to all subdomains.
* `ssl_http_location_conversion` - Enable to replace HTTP with HTTPS in the reply's Location HTTP header field.
* `ssl_http_match_host` - Enable/disable HTTP host matching for location conversion.
* `ssl_max_version` - Highest SSL/TLS version acceptable from a client.
* `ssl_min_version` - Lowest SSL/TLS version acceptable from a client.
* `ssl_mode` - Apply SSL offloading between the client and the FortiGate (half) or from the client to the FortiGate and from the FortiGate to the server (full).
* `ssl_pfs` - Select the cipher suites that can be used for SSL perfect forward secrecy (PFS). Applies to both client and server sessions.
* `ssl_send_empty_frags` - Enable/disable sending empty fragments to avoid CBC IV attacks (SSL 3.0 & TLS 1.0 only). May need to be disabled for compatibility with older systems.
* `ssl_server_algorithm` - Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength.
* `ssl_server_max_version` - Highest SSL/TLS version acceptable from a server. Use the client setting by default.
* `ssl_server_min_version` - Lowest SSL/TLS version acceptable from a server. Use the client setting by default.
* `ssl_server_renegotiation` - Enable/disable secure renegotiation to comply with RFC 5746.
* `ssl_server_session_state_max` - Maximum number of FortiGate to Server SSL session states to keep.
* `ssl_server_session_state_timeout` - Number of minutes to keep FortiGate to Server SSL session state.
* `ssl_server_session_state_type` - How to expire SSL sessions for the segment of the SSL connection between the server and the FortiGate.
* `type` - Configure a static NAT server load balance VIP or access proxy.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `weblogic_server` - Enable to add an HTTP header to indicate SSL offloading for a WebLogic server.
* `websphere_server` - Enable to add an HTTP header to indicate SSL offloading for a WebSphere server.
* `monitor` - Name of the health check monitor to use when polling to determine a virtual server's connectivity status.The structure of `monitor` block is documented below.

The `monitor` block contains:

* `name` - Health monitor name.
* `realservers` - Select the real servers that this server load balancing VIP will distribute traffic to.The structure of `realservers` block is documented below.

The `realservers` block contains:

* `client_ip` - Only clients in this IP range can connect to this real server.
* `healthcheck` - Enable to check the responsiveness of the real server before forwarding traffic.
* `holddown_interval` - Time in seconds that the health check monitor continues to monitor an unresponsive server that should be active.
* `http_host` - HTTP server domain name in HTTP header.
* `id` - Real server ID.
* `ip` - IP address of the real server.
* `max_connections` - Max number of active connections that can directed to the real server. When reached, sessions are sent to other real servers.
* `port` - Port for communicating with the real server. Required if port forwarding is enabled.
* `status` - Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent.
* `translate_host` - Enable/disable translation of hostname/IP from virtual server to real server.
* `weight` - Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.
* `monitor` - Name of the health check monitor to use when polling to determine a virtual server's connectivity status.The structure of `monitor` block is documented below.

The `monitor` block contains:

* `name` - Health monitor name.
* `src_filter` - Source IP6 filter (x:x:x:x:x:x:x:x/x). Separate addresses with spaces.The structure of `src_filter` block is documented below.

The `src_filter` block contains:

* `range` - Source-filter range.
* `ssl_cipher_suites` - SSL/TLS cipher suites acceptable from a client, ordered by priority.The structure of `ssl_cipher_suites` block is documented below.

The `ssl_cipher_suites` block contains:

* `cipher` - Cipher suite name.
* `priority` - SSL/TLS cipher suites priority.
* `versions` - SSL/TLS versions that the cipher suite can be used with.
* `ssl_server_cipher_suites` - SSL/TLS cipher suites to offer to a server, ordered by priority.The structure of `ssl_server_cipher_suites` block is documented below.

The `ssl_server_cipher_suites` block contains:

* `cipher` - Cipher suite name.
* `priority` - SSL/TLS cipher suites priority.
* `versions` - SSL/TLS versions that the cipher suite can be used with.
