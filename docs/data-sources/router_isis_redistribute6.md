---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_isis_redistribute6"
description: |-
  Get information on a fortios IS-IS IPv6 redistribution for routing protocols.
---

# Data Source: fortios_router_isis_redistribute6
Use this data source to get information on a fortios IS-IS IPv6 redistribution for routing protocols.


## Example Usage

```hcl

```

## Argument Reference

* `protocol` - (Required) Protocol name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `level` - Level.
* `metric` - Metric.
* `metric_type` - Metric type.
* `protocol` - Protocol name.
* `routemap` - Route map name.
* `status` - Enable/disable redistribution.
