---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf_network"
description: |-
  Get information on a fortios OSPF network configuration.
---

# Data Source: fortios_router_ospf_network
Use this data source to get information on a fortios OSPF network configuration.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Network entry ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `area` - Attach the network to area.
* `comments` - Comment.
* `id` - Network entry ID.
* `prefix` - Prefix.
