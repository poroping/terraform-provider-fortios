---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_forwardserver"
description: |-
  Get information on a fortios Configure forward-server addresses.
---

# Data Source: fortios_webproxy_forwardserver
Use this data source to get information on a fortios Configure forward-server addresses.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Server name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `addr_type` - Address type of the forwarding proxy server: IP or FQDN.
* `comment` - Comment.
* `fqdn` - Forward server Fully Qualified Domain Name (FQDN).
* `healthcheck` - Enable/disable forward server health checking. Attempts to connect through the remote forwarding server to a destination to verify that the forwarding server is operating normally.
* `ip` - Forward proxy server IP address.
* `monitor` - URL for forward server health check monitoring (default = http://www.google.com).
* `name` - Server name.
* `password` - HTTP authentication password.
* `port` - Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).
* `server_down_option` - Action to take when the forward server is found to be down: block sessions until the server is back up or pass sessions to their destination.
* `username` - HTTP authentication user name.
