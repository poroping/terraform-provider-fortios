---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_vip46"
description: |-
  Configure IPv4 to IPv6 virtual IPs.
---

## fortios_firewall_vip46
Configure IPv4 to IPv6 virtual IPs.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `arp_reply` - Enable ARP reply. Valid values: `disable` `enable` .
* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `extip` - Start-external-IP [-end-external-IP].
* `extport` - External service port.
* `id` - Custom defined id.
* `ldb_method` - Load balance method. Valid values: `static` `round-robin` `weighted` `least-session` `least-rtt` `first-alive` .
* `mappedip` - Start-mapped-IPv6-address [-end mapped-IPv6-address].
* `mappedport` - Mapped service port.
* `name` - VIP46 name.
* `portforward` - Enable port forwarding. Valid values: `disable` `enable` .
* `protocol` - Mapped port protocol. Valid values: `tcp` `udp` .
* `server_type` - Server type. Valid values: `http` `tcp` `udp` `ip` .
* `type` - VIP type: static NAT or server load balance. Valid values: `static-nat` `server-load-balance` .
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `monitor` - Health monitors. The structure of `monitor` block is documented below.

The `monitor` block contains:

* `name` - Health monitor name. This attribute must reference one of the following datasources: `firewall.ldb-monitor.name` .
* `realservers` - Real servers. The structure of `realservers` block is documented below.

The `realservers` block contains:

* `client_ip` - Restrict server to a client IP in this range.
* `healthcheck` - Per server health check. Valid values: `disable` `enable` `vip` .
* `holddown_interval` - Hold down interval.
* `id` - Real server ID.
* `ip` - Mapped server IPv6.
* `max_connections` - Maximum number of connections allowed to server.
* `port` - Mapped server port.
* `status` - Server administrative status. Valid values: `active` `standby` `disable` .
* `weight` - weight
* `monitor` - Health monitors. The structure of `monitor` block is documented below.

The `monitor` block contains:

* `name` - Health monitor name. This attribute must reference one of the following datasources: `firewall.ldb-monitor.name` .
* `src_filter` - Source IP filter (x.x.x.x/x). The structure of `src_filter` block is documented below.

The `src_filter` block contains:

* `range` - Src-filter range.
* `srcintf_filter` - Interfaces to which the VIP46 applies. Separate the names with spaces. The structure of `srcintf_filter` block is documented below.

The `srcintf_filter` block contains:

* `interface_name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_vip46 can be imported using:
```sh
terraform import fortios_firewall_vip46.labelname {{mkey}}
```
