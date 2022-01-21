---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_exchange"
description: |-
  Get information on a fortios Configure MS Exchange server entries.
---

# Data Source: fortios_user_exchange
Use this data source to get information on a fortios Configure MS Exchange server entries.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) MS Exchange server entry name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `auth_level` - Authentication security level used for the RPC protocol layer.
* `auth_type` - Authentication security type used for the RPC protocol layer.
* `auto_discover_kdc` - Enable/disable automatic discovery of KDC IP addresses.
* `connect_protocol` - Connection protocol used to connect to MS Exchange service.
* `domain_name` - MS Exchange server fully qualified domain name.
* `http_auth_type` - Authentication security type used for the HTTP transport.
* `ip` - Server IPv4 address.
* `name` - MS Exchange server entry name.
* `password` - Password for the specified username.
* `server_name` - MS Exchange server hostname.
* `ssl_min_proto_version` - Minimum SSL/TLS protocol version for HTTPS transport (default is to follow system global setting).
* `username` - User name used to sign in to the server. Must have proper permissions for service.
* `kdc_ip` - KDC IPv4 addresses for Kerberos authentication.The structure of `kdc_ip` block is documented below.

The `kdc_ip` block contains:

* `ipv4` - KDC IPv4 addresses for Kerberos authentication.
