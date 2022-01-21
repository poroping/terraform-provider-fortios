---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logsyslogd_filter"
description: |-
  Get information on a fortios Filters for remote system server.
---

# Data Source: fortios_logsyslogd_filter
Use this data source to get information on a fortios Filters for remote system server.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `anomaly` - Enable/disable anomaly logging.
* `filter` - Syslog filter.
* `filter_type` - Include/exclude logs that match the filter.
* `forward_traffic` - Enable/disable forward traffic logging.
* `gtp` - Enable/disable GTP messages logging.
* `local_traffic` - Enable/disable local in or out traffic logging.
* `multicast_traffic` - Enable/disable multicast traffic logging.
* `severity` - Lowest severity level to log.
* `sniffer_traffic` - Enable/disable sniffer traffic logging.
* `voip` - Enable/disable VoIP logging.
* `free_style` - Free Style FiltersThe structure of `free_style` block is documented below.

The `free_style` block contains:

* `category` - Log category.
* `filter` - Free style filter string.
* `filter_type` - Include/exclude logs that match the filter.
* `id` - Entry ID.
