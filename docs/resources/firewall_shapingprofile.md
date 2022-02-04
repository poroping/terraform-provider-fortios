---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_shapingprofile"
description: |-
  Configure shaping profiles.
---

## fortios_firewall_shapingprofile
Configure shaping profiles.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `profile_name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comment` - Comment.
* `default_class_id` - Default class ID to handle unclassified packets (including all local traffic). This attribute must reference one of the following datasources: `firewall.traffic-class.class-id` .
* `profile_name` - Shaping profile name.
* `type` - Select shaping profile type: policing / queuing. Valid values: `policing` `queuing` .
* `shaping_entries` - Define shaping entries of this shaping profile. The structure of `shaping_entries` block is documented below.

The `shaping_entries` block contains:

* `burst_in_msec` - Number of bytes that can be burst at maximum-bandwidth speed. Formula: burst = maximum-bandwidth*burst-in-msec.
* `cburst_in_msec` - Number of bytes that can be burst as fast as the interface can transmit. Formula: cburst = maximum-bandwidth*cburst-in-msec.
* `class_id` - Class ID. This attribute must reference one of the following datasources: `firewall.traffic-class.class-id` .
* `guaranteed_bandwidth_percentage` - Guaranteed bandwidth in percentage.
* `id` - ID number.
* `limit` - Hard limit on the real queue size in packets.
* `max` - Average queue size in packets at which RED drop probability is maximal.
* `maximum_bandwidth_percentage` - Maximum bandwidth in percentage.
* `min` - Average queue size in packets at which RED drop becomes a possibility.
* `priority` - Priority. Valid values: `top` `critical` `high` `medium` `low` .
* `red_probability` - Maximum probability (in percentage) for RED marking.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_shapingprofile can be imported using:
```sh
terraform import fortios_firewall_shapingprofile.labelname {{mkey}}
```
