---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ipam"
description: |-
  Get information on a fortios Configure IP address management services.
---

# Data Source: fortios_system_ipam
Use this data source to get information on a fortios Configure IP address management services.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `pool_subnet` - Configure IPAM pool subnet, Class A - Class B subnet.
* `server_type` - Configure the type of IPAM server to use.
* `status` - Enable/disable IP address management services.
