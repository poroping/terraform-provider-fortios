---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_wisp"
description: |-
  Configure Websense Integrated Services Protocol (WISP) servers.
---

## fortios_webproxy_wisp
Configure Websense Integrated Services Protocol (WISP) servers.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `comment` - Comment.
* `max_connections` - Maximum number of web proxy WISP connections (4 - 4096, default = 64).
* `name` - Server name.
* `outgoing_ip` - WISP outgoing IP address.
* `server_ip` - WISP server IP address.
* `server_port` - WISP server port (1 - 65535, default = 15868).
* `timeout` - Period of time before WISP requests time out (1 - 15 sec, default = 5).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_webproxy_wisp can be imported using:
```sh
terraform import fortios_webproxy_wisp.labelname {{mkey}}
```
