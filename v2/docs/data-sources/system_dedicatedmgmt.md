---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_dedicatedmgmt"
description: |-
  Get information on a fortios Configure dedicated management.
---

# Data Source: fortios_system_dedicatedmgmt
Use this data source to get information on a fortios Configure dedicated management.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `default_gateway` - Default gateway for dedicated management interface.
* `dhcp_end_ip` - DHCP end IP for dedicated management.
* `dhcp_netmask` - DHCP netmask.
* `dhcp_server` - Enable/disable DHCP server on management interface.
* `dhcp_start_ip` - DHCP start IP for dedicated management.
* `interface` - Dedicated management interface.
* `status` - Enable/disable dedicated management.
