---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_qosmap"
description: |-
  Get information on a fortios Configure QoS map set.
---

# Data Source: fortios_wirelesscontrollerhotspot20_qosmap
Use this data source to get information on a fortios Configure QoS map set.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) QOS-MAP name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - QOS-MAP name.
* `dscp_except` - Differentiated Services Code Point (DSCP) exceptions.The structure of `dscp_except` block is documented below.

The `dscp_except` block contains:

* `dscp` - DSCP value.
* `index` - DSCP exception index.
* `up` - User priority.
* `dscp_range` - Differentiated Services Code Point (DSCP) ranges.The structure of `dscp_range` block is documented below.

The `dscp_range` block contains:

* `high` - DSCP high value.
* `index` - DSCP range index.
* `low` - DSCP low value.
* `up` - User priority.
