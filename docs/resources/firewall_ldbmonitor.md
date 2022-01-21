---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_ldbmonitor"
description: |-
  Configure server load balancing health monitors.
---

## fortios_firewall_ldbmonitor
Configure server load balancing health monitors.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `dns_match_ip` - Response IP expected from DNS server.
* `dns_protocol` - Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp` `tcp` .
* `dns_request_domain` - Fully qualified domain name to resolve for the DNS probe.
* `http_get` - URL used to send a GET request to check the health of an HTTP server.
* `http_match` - String to match the value expected in response to an HTTP-GET request.
* `http_max_redirects` - The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
* `interval` - Time between health checks (5 - 65635 sec, default = 10).
* `name` - Monitor name.
* `port` - Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (0 - 65635, default = 0).
* `retry` - Number health check attempts before the server is considered down (1 - 255, default = 3).
* `src_ip` - Source IP for ldb-monitor.
* `timeout` - Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
* `type` - Select the Monitor type used by the health check monitor to check the health of the server (PING | TCP | HTTP | HTTPS | DNS). Valid values: `ping` `tcp` `http` `https` `dns` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_ldbmonitor can be imported using:
```sh
terraform import fortios_firewall_ldbmonitor.labelname {{mkey}}
```
