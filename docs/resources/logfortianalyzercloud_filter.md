---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logfortianalyzercloud_filter"
description: |-
  Filters for FortiAnalyzer Cloud.
---

## fortios_logfortianalyzercloud_filter
Filters for FortiAnalyzer Cloud.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `anomaly` - Enable/disable anomaly logging. Valid values: `enable` `disable` .
* `dlp_archive` - Enable/disable DLP archive logging. Valid values: `enable` `disable` .
* `filter` - FortiAnalyzer Cloud log filter.
* `filter_type` - Include/exclude logs that match the filter. Valid values: `include` `exclude` .
* `forward_traffic` - Enable/disable forward traffic logging. Valid values: `enable` `disable` .
* `gtp` - Enable/disable GTP messages logging. Valid values: `enable` `disable` .
* `local_traffic` - Enable/disable local in or out traffic logging. Valid values: `enable` `disable` .
* `multicast_traffic` - Enable/disable multicast traffic logging. Valid values: `enable` `disable` .
* `severity` - Lowest severity level to log. Valid values: `emergency` `alert` `critical` `error` `warning` `notification` `information` `debug` .
* `sniffer_traffic` - Enable/disable sniffer traffic logging. Valid values: `enable` `disable` .
* `voip` - Enable/disable VoIP logging. Valid values: `enable` `disable` .
* `ztna_traffic` - Enable/disable ztna traffic logging. Valid values: `enable` `disable` .
* `free_style` - Free style filters. The structure of `free_style` block is documented below.

The `free_style` block contains:

* `category` - Log category. Valid values: `traffic` `event` `virus` `webfilter` `attack` `spam` `anomaly` `voip` `dlp` `app-ctrl` `waf` `gtp` `dns` `ssh` `ssl` `file-filter` `icap` .
* `filter` - Free style filter string.
* `filter_type` - Include/exclude logs that match the filter. Valid values: `include` `exclude` .
* `id` - Entry ID.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_logfortianalyzercloud_filter can be imported using:
```sh
terraform import fortios_logfortianalyzercloud_filter.labelname {{mkey}}
```
