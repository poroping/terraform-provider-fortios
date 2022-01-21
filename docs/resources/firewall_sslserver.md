---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_sslserver"
description: |-
  Configure SSL servers.
---

## fortios_firewall_sslserver
Configure SSL servers.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `add_header_x_forwarded_proto` - Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable` `disable` .
* `ip` - IPv4 address of the SSL server.
* `mapped_port` - Mapped server service port (1 - 65535, default = 80).
* `name` - Server name.
* `port` - Server service port (1 - 65535, default = 443).
* `ssl_algorithm` - Relative strength of encryption algorithms accepted in negotiation. Valid values: `high` `medium` `low` .
* `ssl_cert` - Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL"). This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `ssl_client_renegotiation` - Allow or block client renegotiation by server. Valid values: `allow` `deny` `secure` .
* `ssl_dh_bits` - Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768` `1024` `1536` `2048` .
* `ssl_max_version` - Highest SSL/TLS version to negotiate. Valid values: `tls-1.0` `tls-1.1` `tls-1.2` `tls-1.3` .
* `ssl_min_version` - Lowest SSL/TLS version to negotiate. Valid values: `tls-1.0` `tls-1.1` `tls-1.2` `tls-1.3` .
* `ssl_mode` - SSL/TLS mode for encryption and decryption of traffic. Valid values: `half` `full` .
* `ssl_send_empty_frags` - Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable` `disable` .
* `url_rewrite` - Enable/disable rewriting the URL. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_sslserver can be imported using:
```sh
terraform import fortios_firewall_sslserver.labelname {{mkey}}
```
