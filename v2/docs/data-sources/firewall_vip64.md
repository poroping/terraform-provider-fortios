---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_vip64"
description: |-
  Get information on a fortios Configure IPv6 to IPv4 virtual IPs.
---

# Data Source: fortios_firewall_vip64
Use this data source to get information on a fortios Configure IPv6 to IPv4 virtual IPs.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) VIP64 name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `arp_reply` - Enable ARP reply.
* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `extip` - Start-external-IPv6-address [-end-external-IPv6-address].
* `extport` - External service port.
* `id` - Custom defined id.
* `ldb_method` - Load balance method.
* `mappedip` - Start-mapped-IP [-end-mapped-IP].
* `mappedport` - Mapped service port.
* `name` - VIP64 name.
* `portforward` - Enable port forwarding.
* `protocol` - Mapped port protocol.
* `server_type` - Server type.
* `type` - VIP type: static NAT or server load balance.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `monitor` - Health monitors.The structure of `monitor` block is documented below.

The `monitor` block contains:

* `name` - Health monitor name.
* `realservers` - Real servers.The structure of `realservers` block is documented below.

The `realservers` block contains:

* `client_ip` - Restrict server to a client IP in this range.
* `healthcheck` - Per server health check.
* `holddown_interval` - Hold down interval.
* `id` - Real server ID.
* `ip` - Mapped server IP.
* `max_connections` - Maximum number of connections allowed to server.
* `port` - Mapped server port.
* `status` - Server administrative status.
* `weight` - weight
* `monitor` - Health monitors.The structure of `monitor` block is documented below.

The `monitor` block contains:

* `name` - Health monitor name.
* `src_filter` - Source IP6 filter (x:x:x:x:x:x:x:x/x).The structure of `src_filter` block is documented below.

The `src_filter` block contains:

* `range` - Src-filter range.
