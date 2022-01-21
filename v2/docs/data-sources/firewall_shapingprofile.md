---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_shapingprofile"
description: |-
  Get information on a fortios Configure shaping profiles.
---

# Data Source: fortios_firewall_shapingprofile
Use this data source to get information on a fortios Configure shaping profiles.


## Example Usage

```hcl

```

## Argument Reference

* `profile_name` - (Required) Shaping profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `default_class_id` - Default class ID to handle unclassified packets (including all local traffic).
* `profile_name` - Shaping profile name.
* `type` - Select shaping profile type: policing / queuing.
* `shaping_entries` - Define shaping entries of this shaping profile.The structure of `shaping_entries` block is documented below.

The `shaping_entries` block contains:

* `burst_in_msec` - Number of bytes that can be burst at maximum-bandwidth speed. Formula: burst = maximum-bandwidth*burst-in-msec.
* `cburst_in_msec` - Number of bytes that can be burst as fast as the interface can transmit. Formula: cburst = maximum-bandwidth*cburst-in-msec.
* `class_id` - Class ID.
* `guaranteed_bandwidth_percentage` - Guaranteed bandwith in percentage.
* `id` - ID number.
* `limit` - Hard limit on the real queue size in packets.
* `max` - Average queue size in packets at which RED drop probability is maximal.
* `maximum_bandwidth_percentage` - Maximum bandwith in percentage.
* `min` - Average queue size in packets at which RED drop becomes a possibility.
* `priority` - Priority.
* `red_probability` - Maximum probability (in percentage) for RED marking.
