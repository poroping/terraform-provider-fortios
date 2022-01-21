---
subcategory: "FortiGate Webfilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_ipsurlfiltersetting6"
description: |-
  Get information on a fortios Configure IPS URL filter settings for IPv6.
---

# Data Source: fortios_webfilter_ipsurlfiltersetting6
Use this data source to get information on a fortios Configure IPS URL filter settings for IPv6.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `device` - Interface for this route.
* `distance` - Administrative distance (1 - 255) for this route.
* `gateway6` - Gateway IPv6 address for this route.
* `geo_filter` - Filter based on geographical location. Route will NOT be installed if the resolved IPv6 address belongs to the country in the filter.
