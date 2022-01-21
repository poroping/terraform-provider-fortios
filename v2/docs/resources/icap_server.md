---
subcategory: "FortiGate Icap"
layout: "fortios"
page_title: "FortiOS: fortios_icap_server"
description: |-
  Configure ICAP servers.
---

## fortios_icap_server
Configure ICAP servers.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `ip_address` - IPv4 address of the ICAP server.
* `ip_version` - IP version. Valid values: `4` `6` .
* `ip6_address` - IPv6 address of the ICAP server.
* `max_connections` - Maximum number of concurrent connections to ICAP server. Must not be less than wad-worker-count.
* `name` - Server name.
* `port` - ICAP server port.
* `secure` - Enable/disable secure connection to ICAP server. Valid values: `enable` `disable` .
* `ssl_cert` - CA certificate name. This attribute must reference one of the following datasources: `vpn.certificate.ca.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_icap_server can be imported using:
```sh
terraform import fortios_icap_server.labelname {{mkey}}
```
