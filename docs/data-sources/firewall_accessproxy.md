---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_accessproxy"
description: |-
  Get information on a fortios Configure IPv4 access proxy.
---

# Data Source: fortios_firewall_accessproxy
Use this data source to get information on a fortios Configure IPv4 access proxy.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Access Proxy name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `client_cert` - Enable/disable to request client certificate.
* `decrypted_traffic_mirror` - Decrypted traffic mirror.
* `empty_cert_action` - Action of an empty client certificate.
* `ldb_method` - Method used to distribute sessions to SSL real servers.
* `log_blocked_traffic` - Enable/disable logging of blocked traffic.
* `name` - Access Proxy name.
* `server_pubkey_auth` - Enable/disable SSH real server public key authentication.
* `vip` - Virtual IP name.
* `api_gateway` - Set IPv4 API Gateway.The structure of `api_gateway` block is documented below.

The `api_gateway` block contains:

* `http_cookie_age` - Time in minutes that client web browsers should keep a cookie. Default is 60 minutes. 0 = no time limit.
* `http_cookie_domain` - Domain that HTTP cookie persistence should apply to.
* `http_cookie_domain_from_host` - Enable/disable use of HTTP cookie domain from host field in HTTP.
* `http_cookie_generation` - Generation of HTTP cookie to be accepted. Changing invalidates all existing cookies.
* `http_cookie_path` - Limit HTTP cookie persistence to the specified path.
* `http_cookie_share` - Control sharing of cookies across API Gateway. same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing.
* `https_cookie_secure` - Enable/disable verification that inserted HTTPS cookies are secure.
* `id` - API Gateway ID.
* `ldb_method` - Method used to distribute sessions to real servers.
* `persistence` - Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session.
* `saml_redirect` - Enable/disable SAML redirection after successful authentication.
* `saml_server` - SAML service provider configuration for VIP authentication.
* `service` - Service.
* `ssl_algorithm` - Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength.
* `ssl_dh_bits` - Number of bits to use in the Diffie-Hellman exchange for RSA encryption of SSL sessions.
* `ssl_max_version` - Highest SSL/TLS version acceptable from a server.
* `ssl_min_version` - Lowest SSL/TLS version acceptable from a server.
* `url_map` - URL pattern to match.
* `url_map_type` - Type of url-map.
* `virtual_host` - Virtual host.
* `realservers` - Select the real servers that this Access Proxy will distribute traffic to.The structure of `realservers` block is documented below.

The `realservers` block contains:

* `addr_type` - Type of address.
* `address` - Address or address group of the real server.
* `health_check` - Enable to check the responsiveness of the real server before forwarding traffic.
* `health_check_proto` - Protocol of the health check monitor to use when polling to determine server's connectivity status.
* `holddown_interval` - Enable/disable holddown timer. Server will be considered active and reachable once the holddown period has expired (30 seconds).
* `http_host` - HTTP server domain name in HTTP header.
* `id` - Real server ID.
* `ip` - IP address of the real server.
* `mappedport` - Port for communicating with the real server.
* `port` - Port for communicating with the real server.
* `ssh_client_cert` - Set access-proxy SSH client certificate profile.
* `ssh_host_key_validation` - Enable/disable SSH real server host key validation.
* `status` - Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent.
* `type` - TCP forwarding server type.
* `weight` - Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.
* `ssh_host_key` - One or more server host key.The structure of `ssh_host_key` block is documented below.

The `ssh_host_key` block contains:

* `name` - Server host key name.
* `ssl_cipher_suites` - SSL/TLS cipher suites to offer to a server, ordered by priority.The structure of `ssl_cipher_suites` block is documented below.

The `ssl_cipher_suites` block contains:

* `cipher` - Cipher suite name.
* `priority` - SSL/TLS cipher suites priority.
* `versions` - SSL/TLS versions that the cipher suite can be used with.
* `api_gateway6` - Set IPv6 API Gateway.The structure of `api_gateway6` block is documented below.

The `api_gateway6` block contains:

* `http_cookie_age` - Time in minutes that client web browsers should keep a cookie. Default is 60 minutes. 0 = no time limit.
* `http_cookie_domain` - Domain that HTTP cookie persistence should apply to.
* `http_cookie_domain_from_host` - Enable/disable use of HTTP cookie domain from host field in HTTP.
* `http_cookie_generation` - Generation of HTTP cookie to be accepted. Changing invalidates all existing cookies.
* `http_cookie_path` - Limit HTTP cookie persistence to the specified path.
* `http_cookie_share` - Control sharing of cookies across API Gateway. same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing.
* `https_cookie_secure` - Enable/disable verification that inserted HTTPS cookies are secure.
* `id` - API Gateway ID.
* `ldb_method` - Method used to distribute sessions to real servers.
* `persistence` - Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session.
* `saml_redirect` - Enable/disable SAML redirection after successful authentication.
* `saml_server` - SAML service provider configuration for VIP authentication.
* `service` - Service.
* `ssl_algorithm` - Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength.
* `ssl_dh_bits` - Number of bits to use in the Diffie-Hellman exchange for RSA encryption of SSL sessions.
* `ssl_max_version` - Highest SSL/TLS version acceptable from a server.
* `ssl_min_version` - Lowest SSL/TLS version acceptable from a server.
* `url_map` - URL pattern to match.
* `url_map_type` - Type of url-map.
* `virtual_host` - Virtual host.
* `realservers` - Select the real servers that this Access Proxy will distribute traffic to.The structure of `realservers` block is documented below.

The `realservers` block contains:

* `addr_type` - Type of address.
* `address` - Address or address group of the real server.
* `health_check` - Enable to check the responsiveness of the real server before forwarding traffic.
* `health_check_proto` - Protocol of the health check monitor to use when polling to determine server's connectivity status.
* `holddown_interval` - Enable/disable holddown timer. Server will be considered active and reachable once the holddown period has expired (30 seconds).
* `http_host` - HTTP server domain name in HTTP header.
* `id` - Real server ID.
* `ip` - IPv6 address of the real server.
* `mappedport` - Port for communicating with the real server.
* `port` - Port for communicating with the real server.
* `ssh_client_cert` - Set access-proxy SSH client certificate profile.
* `ssh_host_key_validation` - Enable/disable SSH real server host key validation.
* `status` - Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent.
* `type` - TCP forwarding server type.
* `weight` - Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.
* `ssh_host_key` - One or more server host key.The structure of `ssh_host_key` block is documented below.

The `ssh_host_key` block contains:

* `name` - Server host key name.
* `ssl_cipher_suites` - SSL/TLS cipher suites to offer to a server, ordered by priority.The structure of `ssl_cipher_suites` block is documented below.

The `ssl_cipher_suites` block contains:

* `cipher` - Cipher suite name.
* `priority` - SSL/TLS cipher suites priority.
* `versions` - SSL/TLS versions that the cipher suite can be used with.
* `realservers` - Select the SSL real servers that this Access Proxy will distribute traffic to.The structure of `realservers` block is documented below.

The `realservers` block contains:

* `id` - Real server ID.
* `ip` - IP address of the real server.
* `port` - Port for communicating with the real server.
* `status` - Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent.
* `weight` - Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.
* `server_pubkey_auth_settings` - Server SSH public key authentication settings.The structure of `server_pubkey_auth_settings` block is documented below.

The `server_pubkey_auth_settings` block contains:

* `auth_ca` - Name of the SSH server public key authentication CA.
* `permit_agent_forwarding` - Enable/disable appending permit-agent-forwarding certificate extension.
* `permit_port_forwarding` - Enable/disable appending permit-port-forwarding certificate extension.
* `permit_pty` - Enable/disable appending permit-pty certificate extension.
* `permit_user_rc` - Enable/disable appending permit-user-rc certificate extension.
* `permit_x11_forwarding` - Enable/disable appending permit-x11-forwarding certificate extension.
* `source_address` - Enable/disable appending source-address certificate critical option. This option ensure certificate only accepted from FortiGate source address.
* `cert_extension` - Configure certificate extension for user certificate.The structure of `cert_extension` block is documented below.

The `cert_extension` block contains:

* `critical` - Critical option.
* `data` - Name of certificate extension.
* `name` - Name of certificate extension.
* `type` - Type of certificate extension.
