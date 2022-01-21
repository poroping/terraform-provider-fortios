---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_forwardserver"
description: |-
  Configure forward-server addresses.
---

## fortios_webproxy_forwardserver
Configure forward-server addresses.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `addr_type` - Address type of the forwarding proxy server: IP or FQDN. Valid values: `ip` `fqdn` .
* `comment` - Comment.
* `fqdn` - Forward server Fully Qualified Domain Name (FQDN).
* `healthcheck` - Enable/disable forward server health checking. Attempts to connect through the remote forwarding server to a destination to verify that the forwarding server is operating normally. Valid values: `disable` `enable` .
* `ip` - Forward proxy server IP address.
* `monitor` - URL for forward server health check monitoring (default = http://www.google.com).
* `name` - Server name.
* `password` - HTTP authentication password.
* `port` - Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).
* `server_down_option` - Action to take when the forward server is found to be down: block sessions until the server is back up or pass sessions to their destination. Valid values: `block` `pass` .
* `username` - HTTP authentication user name.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_webproxy_forwardserver can be imported using:
```sh
terraform import fortios_webproxy_forwardserver.labelname {{mkey}}
```
