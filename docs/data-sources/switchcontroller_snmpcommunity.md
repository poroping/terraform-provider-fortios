---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_snmpcommunity"
description: |-
  Get information on a fortios Configure FortiSwitch SNMP v1/v2c communities globally.
---

# Data Source: fortios_switchcontroller_snmpcommunity
Use this data source to get information on a fortios Configure FortiSwitch SNMP v1/v2c communities globally.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) SNMP community ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `events` - SNMP notifications (traps) to send.
* `id` - SNMP community ID.
* `name` - SNMP community name.
* `query_v1_port` - SNMP v1 query port (default = 161).
* `query_v1_status` - Enable/disable SNMP v1 queries.
* `query_v2c_port` - SNMP v2c query port (default = 161).
* `query_v2c_status` - Enable/disable SNMP v2c queries.
* `status` - Enable/disable this SNMP community.
* `trap_v1_lport` - SNMP v2c trap local port (default = 162).
* `trap_v1_rport` - SNMP v2c trap remote port (default = 162).
* `trap_v1_status` - Enable/disable SNMP v1 traps.
* `trap_v2c_lport` - SNMP v2c trap local port (default = 162).
* `trap_v2c_rport` - SNMP v2c trap remote port (default = 162).
* `trap_v2c_status` - Enable/disable SNMP v2c traps.
* `hosts` - Configure IPv4 SNMP managers (hosts).The structure of `hosts` block is documented below.

The `hosts` block contains:

* `id` - Host entry ID.
* `ip` - IPv4 address of the SNMP manager (host).
