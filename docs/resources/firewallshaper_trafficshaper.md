---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallshaper_trafficshaper"
description: |-
  Configure shared traffic shaper.
---

## fortios_firewallshaper_trafficshaper
Configure shared traffic shaper.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `bandwidth_unit` - Unit of measurement for guaranteed and maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps` `mbps` `gbps` .
* `diffserv` - Enable/disable changing the DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable` `disable` .
* `diffservcode` - DiffServ setting to be applied to traffic accepted by this shaper.
* `dscp_marking_method` - Select DSCP marking method. Valid values: `multi-stage` `static` .
* `exceed_bandwidth` - Exceed bandwidth used for DSCP multi-stage marking. Units depend on the bandwidth-unit setting.
* `exceed_class_id` - Class ID for traffic in [guaranteed-bandwidth, maximum-bandwidth]. This attribute must reference one of the following datasources: `firewall.traffic-class.class-id` .
* `exceed_dscp` - DSCP mark for traffic in [guaranteed-bandwidth, exceed-bandwidth].
* `guaranteed_bandwidth` - Amount of bandwidth guaranteed for this shaper (0 - 16776000). Units depend on the bandwidth-unit setting.
* `maximum_bandwidth` - Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.
* `maximum_dscp` - DSCP mark for traffic in [exceed-bandwidth, maximum-bandwidth].
* `name` - Traffic shaper name.
* `overhead` - Per-packet size overhead used in rate computations.
* `per_policy` - Enable/disable applying a separate shaper for each policy. For example, if enabled the guaranteed bandwidth is applied separately for each policy. Valid values: `disable` `enable` .
* `priority` - Higher priority traffic is more likely to be forwarded without delays and without compromising the guaranteed bandwidth. Valid values: `low` `medium` `high` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewallshaper_trafficshaper can be imported using:
```sh
terraform import fortios_firewallshaper_trafficshaper.labelname {{mkey}}
```
