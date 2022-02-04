---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollerinitialconfig_template"
description: |-
  Get information on a fortios Configure template for auto-generated VLANs.
---

# Data Source: fortios_switchcontrollerinitialconfig_template
Use this data source to get information on a fortios Configure template for auto-generated VLANs.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Initial config template name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `allowaccess` - Permitted types of management access to this interface.
* `auto_ip` - Automatically allocate interface address and subnet block.
* `dhcp_server` - Enable/disable a DHCP server on this interface.
* `ip` - Interface IPv4 address and subnet mask.
* `name` - Initial config template name.
* `vlanid` - Unique VLAN ID.
