---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logmemory_filter"
description: |-
  Filters for memory buffer.
---

## fortios_logmemory_filter
Filters for memory buffer.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `admin` - Enable/disable admin login/logout logging. Valid values: `enable` `disable` .
* `anomaly` - Enable/disable anomaly logging. Valid values: `enable` `disable` .
* `auth` - Enable/disable firewall authentication logging. Valid values: `enable` `disable` .
* `cpu_memory_usage` - Enable/disable CPU & memory usage logging every 5 minutes. Valid values: `enable` `disable` .
* `dhcp` - Enable/disable DHCP service messages logging. Valid values: `enable` `disable` .
* `event` - Enable/disable event logging. Valid values: `enable` `disable` .
* `filter` - Memory log filter.
* `filter_type` - Include/exclude logs that match the filter. Valid values: `include` `exclude` .
* `forward_traffic` - Enable/disable forward traffic logging. Valid values: `enable` `disable` .
* `gtp` - Enable/disable GTP messages logging. Valid values: `enable` `disable` .
* `ha` - Enable/disable HA logging. Valid values: `enable` `disable` .
* `ipsec` - Enable/disable IPsec negotiation messages logging. Valid values: `enable` `disable` .
* `ldb_monitor` - Enable/disable VIP real server health monitoring logging. Valid values: `enable` `disable` .
* `local_traffic` - Enable/disable local in or out traffic logging. Valid values: `enable` `disable` .
* `multicast_traffic` - Enable/disable multicast traffic logging. Valid values: `enable` `disable` .
* `pattern` - Enable/disable pattern update logging. Valid values: `enable` `disable` .
* `ppp` - Enable/disable L2TP/PPTP/PPPoE logging. Valid values: `enable` `disable` .
* `radius` - Enable/disable RADIUS messages logging. Valid values: `enable` `disable` .
* `severity` - Log every message above and including this severity level. Valid values: `emergency` `alert` `critical` `error` `warning` `notification` `information` `debug` .
* `sniffer_traffic` - Enable/disable sniffer traffic logging. Valid values: `enable` `disable` .
* `sslvpn_log_adm` - Enable/disable SSL administrator login logging. Valid values: `enable` `disable` .
* `sslvpn_log_auth` - Enable/disable SSL user authentication logging. Valid values: `enable` `disable` .
* `sslvpn_log_session` - Enable/disable SSL session logging. Valid values: `enable` `disable` .
* `system` - Enable/disable system activity logging. Valid values: `enable` `disable` .
* `vip_ssl` - Enable/disable VIP SSL logging. Valid values: `enable` `disable` .
* `voip` - Enable/disable VoIP logging. Valid values: `enable` `disable` .
* `wan_opt` - Enable/disable WAN optimization event logging. Valid values: `enable` `disable` .
* `wireless_activity` - Enable/disable wireless activity event logging. Valid values: `enable` `disable` .
* `free_style` - Free Style Filters The structure of `free_style` block is documented below.

The `free_style` block contains:

* `category` - Log category. Valid values: `traffic` `event` `virus` `webfilter` `attack` `spam` `anomaly` `voip` `dlp` `app-ctrl` `waf` `gtp` `dns` `ssh` `ssl` `file-filter` `icap` `ztna` .
* `filter` - Free style filter string.
* `filter_type` - Include/exclude logs that match the filter. Valid values: `include` `exclude` .
* `id` - Entry ID.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_logmemory_filter can be imported using:
```sh
terraform import fortios_logmemory_filter.labelname {{mkey}}
```
