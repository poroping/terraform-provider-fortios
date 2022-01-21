---
subcategory: "FortiGate Icap"
layout: "fortios"
page_title: "FortiOS: fortios_icap_server"
description: |-
  Get information on a fortios Configure ICAP servers.
---

# Data Source: fortios_icap_server
Use this data source to get information on a fortios Configure ICAP servers.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Server name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `ip_address` - IPv4 address of the ICAP server.
* `ip_version` - IP version.
* `ip6_address` - IPv6 address of the ICAP server.
* `max_connections` - Maximum number of concurrent connections to ICAP server. Must not be less than wad-worker-count.
* `name` - Server name.
* `port` - ICAP server port.
* `secure` - Enable/disable secure connection to ICAP server.
* `ssl_cert` - CA certificate name.
