---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_wisp"
description: |-
  Get information on a fortios Configure Wireless Internet service provider (WISP) servers.
---

# Data Source: fortios_webproxy_wisp
Use this data source to get information on a fortios Configure Wireless Internet service provider (WISP) servers.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Server name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `max_connections` - Maximum number of web proxy WISP connections (4 - 4096, default = 64).
* `name` - Server name.
* `outgoing_ip` - WISP outgoing IP address.
* `server_ip` - WISP server IP address.
* `server_port` - WISP server port (1 - 65535, default = 15868).
* `timeout` - Period of time before WISP requests time out (1 - 15 sec, default = 5).
