---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_sslserver"
description: |-
  Get information on a fortios Configure SSL servers.
---

# Data Source: fortios_firewall_sslserver
Use this data source to get information on a fortios Configure SSL servers.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Server name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `add_header_x_forwarded_proto` - Enable/disable adding an X-Forwarded-Proto header to forwarded requests.
* `ip` - IPv4 address of the SSL server.
* `mapped_port` - Mapped server service port (1 - 65535, default = 80).
* `name` - Server name.
* `port` - Server service port (1 - 65535, default = 443).
* `ssl_algorithm` - Relative strength of encryption algorithms accepted in negotiation.
* `ssl_cert` - Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
* `ssl_client_renegotiation` - Allow or block client renegotiation by server.
* `ssl_dh_bits` - Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048).
* `ssl_max_version` - Highest SSL/TLS version to negotiate.
* `ssl_min_version` - Lowest SSL/TLS version to negotiate.
* `ssl_mode` - SSL/TLS mode for encryption and decryption of traffic.
* `ssl_send_empty_frags` - Enable/disable sending empty fragments to avoid attack on CBC IV.
* `url_rewrite` - Enable/disable rewriting the URL.
