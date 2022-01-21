---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_snmpcommunity"
description: |-
  Configure FortiSwitch SNMP v1/v2c communities globally.
---

## fortios_switchcontroller_snmpcommunity
Configure FortiSwitch SNMP v1/v2c communities globally.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `events` - SNMP notifications (traps) to send. Valid values: `cpu-high` `mem-low` `log-full` `intf-ip` `ent-conf-change` .
* `id` - SNMP community ID.
* `name` - SNMP community name.
* `query_v1_port` - SNMP v1 query port (default = 161).
* `query_v1_status` - Enable/disable SNMP v1 queries. Valid values: `disable` `enable` .
* `query_v2c_port` - SNMP v2c query port (default = 161).
* `query_v2c_status` - Enable/disable SNMP v2c queries. Valid values: `disable` `enable` .
* `status` - Enable/disable this SNMP community. Valid values: `disable` `enable` .
* `trap_v1_lport` - SNMP v2c trap local port (default = 162).
* `trap_v1_rport` - SNMP v2c trap remote port (default = 162).
* `trap_v1_status` - Enable/disable SNMP v1 traps. Valid values: `disable` `enable` .
* `trap_v2c_lport` - SNMP v2c trap local port (default = 162).
* `trap_v2c_rport` - SNMP v2c trap remote port (default = 162).
* `trap_v2c_status` - Enable/disable SNMP v2c traps. Valid values: `disable` `enable` .
* `hosts` - Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.

The `hosts` block contains:

* `id` - Host entry ID.
* `ip` - IPv4 address of the SNMP manager (host).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_switchcontroller_snmpcommunity can be imported using:
```sh
terraform import fortios_switchcontroller_snmpcommunity.labelname {{mkey}}
```
