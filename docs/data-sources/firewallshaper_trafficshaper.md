---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallshaper_trafficshaper"
description: |-
  Get information on a fortios Configure shared traffic shaper.
---

# Data Source: fortios_firewallshaper_trafficshaper
Use this data source to get information on a fortios Configure shared traffic shaper.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Traffic shaper name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `bandwidth_unit` - Unit of measurement for guaranteed and maximum bandwidth for this shaper (Kbps, Mbps or Gbps).
* `diffserv` - Enable/disable changing the DiffServ setting applied to traffic accepted by this shaper.
* `diffservcode` - DiffServ setting to be applied to traffic accepted by this shaper.
* `dscp_marking_method` - Select DSCP marking method.
* `exceed_bandwidth` - Exceed bandwidth used for DSCP multi-stage marking. Units depend on the bandwidth-unit setting.
* `exceed_class_id` - Class ID for traffic in guaranteed-bandwidth and maximum-bandwidth.
* `exceed_dscp` - DSCP mark for traffic in guaranteed-bandwidth and exceed-bandwidth.
* `guaranteed_bandwidth` - Amount of bandwidth guaranteed for this shaper (0 - 16776000). Units depend on the bandwidth-unit setting.
* `maximum_bandwidth` - Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.
* `maximum_dscp` - DSCP mark for traffic in exceed-bandwidth and maximum-bandwidth.
* `name` - Traffic shaper name.
* `overhead` - Per-packet size overhead used in rate computations.
* `per_policy` - Enable/disable applying a separate shaper for each policy. For example, if enabled the guaranteed bandwidth is applied separately for each policy.
* `priority` - Higher priority traffic is more likely to be forwarded without delays and without compromising the guaranteed bandwidth.
