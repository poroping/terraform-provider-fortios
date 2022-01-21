---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logmemory_filter"
description: |-
  Get information on a fortios Filters for memory buffer.
---

# Data Source: fortios_logmemory_filter
Use this data source to get information on a fortios Filters for memory buffer.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `admin` - Enable/disable admin login/logout logging.
* `anomaly` - Enable/disable anomaly logging.
* `auth` - Enable/disable firewall authentication logging.
* `cpu_memory_usage` - Enable/disable CPU & memory usage logging every 5 minutes.
* `dhcp` - Enable/disable DHCP service messages logging.
* `event` - Enable/disable event logging.
* `filter` - Memory log filter.
* `filter_type` - Include/exclude logs that match the filter.
* `forward_traffic` - Enable/disable forward traffic logging.
* `gtp` - Enable/disable GTP messages logging.
* `ha` - Enable/disable HA logging.
* `ipsec` - Enable/disable IPsec negotiation messages logging.
* `ldb_monitor` - Enable/disable VIP real server health monitoring logging.
* `local_traffic` - Enable/disable local in or out traffic logging.
* `multicast_traffic` - Enable/disable multicast traffic logging.
* `pattern` - Enable/disable pattern update logging.
* `ppp` - Enable/disable L2TP/PPTP/PPPoE logging.
* `radius` - Enable/disable RADIUS messages logging.
* `severity` - Log every message above and including this severity level.
* `sniffer_traffic` - Enable/disable sniffer traffic logging.
* `sslvpn_log_adm` - Enable/disable SSL administrator login logging.
* `sslvpn_log_auth` - Enable/disable SSL user authentication logging.
* `sslvpn_log_session` - Enable/disable SSL session logging.
* `system` - Enable/disable system activity logging.
* `vip_ssl` - Enable/disable VIP SSL logging.
* `voip` - Enable/disable VoIP logging.
* `wan_opt` - Enable/disable WAN optimization event logging.
* `wireless_activity` - Enable/disable wireless activity event logging.
* `free_style` - Free Style FiltersThe structure of `free_style` block is documented below.

The `free_style` block contains:

* `category` - Log category.
* `filter` - Free style filter string.
* `filter_type` - Include/exclude logs that match the filter.
* `id` - Entry ID.
