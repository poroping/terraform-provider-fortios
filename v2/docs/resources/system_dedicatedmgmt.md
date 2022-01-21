---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_dedicatedmgmt"
description: |-
  Configure dedicated management.
---

## fortios_system_dedicatedmgmt
Configure dedicated management.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `default_gateway` - Default gateway for dedicated management interface.
* `dhcp_end_ip` - DHCP end IP for dedicated management.
* `dhcp_netmask` - DHCP netmask.
* `dhcp_server` - Enable/disable DHCP server on management interface. Valid values: `enable` `disable` .
* `dhcp_start_ip` - DHCP start IP for dedicated management.
* `interface` - Dedicated management interface. This attribute must reference one of the following datasources: `system.interface.name` .
* `status` - Enable/disable dedicated management. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_dedicatedmgmt can be imported using:
```sh
terraform import fortios_system_dedicatedmgmt.labelname {{mkey}}
```
