---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallipmacbinding_table"
description: |-
  Get information on a fortios Configure IP to MAC address pairs in the IP/MAC binding table.
---

# Data Source: fortios_firewallipmacbinding_table
Use this data source to get information on a fortios Configure IP to MAC address pairs in the IP/MAC binding table.


## Example Usage

```hcl

```

## Argument Reference

* `seq_num` - (Required) Entry number.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `ip` - IPv4 address portion of the pair (format: xxx.xxx.xxx.xxx).
* `mac` - MAC address portion of the pair (format = xx:xx:xx:xx:xx:xx in hexadecimal).
* `name` - Name of the pair (optional, default = no name).
* `seq_num` - Entry number.
* `status` - Enable/disable this IP-mac binding pair.
