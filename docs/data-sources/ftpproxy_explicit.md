---
subcategory: "FortiGate Ftp-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_ftpproxy_explicit"
description: |-
  Get information on a fortios Configure explicit FTP proxy settings.
---

# Data Source: fortios_ftpproxy_explicit
Use this data source to get information on a fortios Configure explicit FTP proxy settings.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `incoming_ip` - Accept incoming FTP requests from this IP address. An interface must have this IP address.
* `incoming_port` - Accept incoming FTP requests on one or more ports.
* `outgoing_ip` - Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
* `sec_default_action` - Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists.
* `ssl` - Enable/disable the explicit FTPS proxy.
* `ssl_algorithm` - Relative strength of encryption algorithms accepted in negotiation.
* `ssl_cert` - Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
* `ssl_dh_bits` - Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048).
* `status` - Enable/disable the explicit FTP proxy.
