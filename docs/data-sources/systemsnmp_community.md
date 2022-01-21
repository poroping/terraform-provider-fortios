---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemsnmp_community"
description: |-
  Get information on a fortios SNMP community configuration.
---

# Data Source: fortios_systemsnmp_community
Use this data source to get information on a fortios SNMP community configuration.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Community ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `events` - SNMP trap events.
* `id` - Community ID.
* `name` - Community name.
* `query_v1_port` - SNMP v1 query port (default = 161).
* `query_v1_status` - Enable/disable SNMP v1 queries.
* `query_v2c_port` - SNMP v2c query port (default = 161).
* `query_v2c_status` - Enable/disable SNMP v2c queries.
* `status` - Enable/disable this SNMP community.
* `trap_v1_lport` - SNMP v1 trap local port (default = 162).
* `trap_v1_rport` - SNMP v1 trap remote port (default = 162).
* `trap_v1_status` - Enable/disable SNMP v1 traps.
* `trap_v2c_lport` - SNMP v2c trap local port (default = 162).
* `trap_v2c_rport` - SNMP v2c trap remote port (default = 162).
* `trap_v2c_status` - Enable/disable SNMP v2c traps.
* `hosts` - Configure IPv4 SNMP managers (hosts).The structure of `hosts` block is documented below.

The `hosts` block contains:

* `ha_direct` - Enable/disable direct management of HA cluster members.
* `host_type` - Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both. No traps will be sent when IP type is subnet.
* `id` - Host entry ID.
* `ip` - IPv4 address of the SNMP manager (host).
* `source_ip` - Source IPv4 address for SNMP traps.
* `hosts6` - Configure IPv6 SNMP managers.The structure of `hosts6` block is documented below.

The `hosts6` block contains:

* `ha_direct` - Enable/disable direct management of HA cluster members.
* `host_type` - Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both.
* `id` - Host6 entry ID.
* `ipv6` - SNMP manager IPv6 address prefix.
* `source_ipv6` - Source IPv6 address for SNMP traps.
