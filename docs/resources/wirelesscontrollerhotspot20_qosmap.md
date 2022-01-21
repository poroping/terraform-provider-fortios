---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_qosmap"
description: |-
  Configure QoS map set.
---

## fortios_wirelesscontrollerhotspot20_qosmap
Configure QoS map set.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `name` - QOS-MAP name.
* `dscp_except` - Differentiated Services Code Point (DSCP) exceptions. The structure of `dscp_except` block is documented below.

The `dscp_except` block contains:

* `dscp` - DSCP value.
* `index` - DSCP exception index.
* `up` - User priority.
* `dscp_range` - Differentiated Services Code Point (DSCP) ranges. The structure of `dscp_range` block is documented below.

The `dscp_range` block contains:

* `high` - DSCP high value.
* `index` - DSCP range index.
* `low` - DSCP low value.
* `up` - User priority.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontrollerhotspot20_qosmap can be imported using:
```sh
terraform import fortios_wirelesscontrollerhotspot20_qosmap.labelname {{mkey}}
```
