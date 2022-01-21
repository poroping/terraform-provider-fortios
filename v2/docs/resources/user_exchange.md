---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_exchange"
description: |-
  Configure MS Exchange server entries.
---

## fortios_user_exchange
Configure MS Exchange server entries.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `auth_level` - Authentication security level used for the RPC protocol layer. Valid values: `connect` `call` `packet` `integrity` `privacy` .
* `auth_type` - Authentication security type used for the RPC protocol layer. Valid values: `spnego` `ntlm` `kerberos` .
* `auto_discover_kdc` - Enable/disable automatic discovery of KDC IP addresses. Valid values: `enable` `disable` .
* `connect_protocol` - Connection protocol used to connect to MS Exchange service. Valid values: `rpc-over-tcp` `rpc-over-http` `rpc-over-https` .
* `domain_name` - MS Exchange server fully qualified domain name.
* `http_auth_type` - Authentication security type used for the HTTP transport. Valid values: `basic` `ntlm` .
* `ip` - Server IPv4 address.
* `name` - MS Exchange server entry name.
* `password` - Password for the specified username.
* `server_name` - MS Exchange server hostname.
* `ssl_min_proto_version` - Minimum SSL/TLS protocol version for HTTPS transport (default is to follow system global setting). Valid values: `default` `SSLv3` `TLSv1` `TLSv1-1` `TLSv1-2` .
* `username` - User name used to sign in to the server. Must have proper permissions for service.
* `kdc_ip` - KDC IPv4 addresses for Kerberos authentication. The structure of `kdc_ip` block is documented below.

The `kdc_ip` block contains:

* `ipv4` - KDC IPv4 addresses for Kerberos authentication.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_user_exchange can be imported using:
```sh
terraform import fortios_user_exchange.labelname {{mkey}}
```
