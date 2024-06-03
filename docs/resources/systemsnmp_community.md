---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemsnmp_community"
description: |-
  SNMP community configuration.
---

## fortios_systemsnmp_community
SNMP community configuration.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `events` - SNMP trap events. Valid values: `cpu-high` `mem-low` `log-full` `intf-ip` `vpn-tun-up` `vpn-tun-down` `ha-switch` `ha-hb-failure` `ips-signature` `ips-anomaly` `av-virus` `av-oversize` `av-pattern` `av-fragmented` `fm-if-change` `fm-conf-change` `bgp-established` `bgp-backward-transition` `ha-member-up` `ha-member-down` `ent-conf-change` `av-conserve` `av-bypass` `av-oversize-passed` `av-oversize-blocked` `ips-pkg-update` `ips-fail-open` `temperature-high` `voltage-alert` `power-supply-failure` `faz-disconnect` `fan-failure` `wc-ap-up` `wc-ap-down` `fswctl-session-up` `fswctl-session-down` `load-balance-real-server-down` `device-new` `per-cpu-high` `dhcp` `pool-usage` `ospf-nbr-state-change` `ospf-virtnbr-state-change` .
* `id` - Community ID.
* `mib_view` - SNMP access control MIB view. This attribute must reference one of the following datasources: `system.snmp.mib-view.name` .
* `name` - Community name.
* `query_v1_port` - SNMP v1 query port (default = 161).
* `query_v1_status` - Enable/disable SNMP v1 queries. Valid values: `enable` `disable` .
* `query_v2c_port` - SNMP v2c query port (default = 161).
* `query_v2c_status` - Enable/disable SNMP v2c queries. Valid values: `enable` `disable` .
* `status` - Enable/disable this SNMP community. Valid values: `enable` `disable` .
* `trap_v1_lport` - SNMP v1 trap local port (default = 162).
* `trap_v1_rport` - SNMP v1 trap remote port (default = 162).
* `trap_v1_status` - Enable/disable SNMP v1 traps. Valid values: `enable` `disable` .
* `trap_v2c_lport` - SNMP v2c trap local port (default = 162).
* `trap_v2c_rport` - SNMP v2c trap remote port (default = 162).
* `trap_v2c_status` - Enable/disable SNMP v2c traps. Valid values: `enable` `disable` .
* `hosts` - Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.

The `hosts` block contains:

* `ha_direct` - Enable/disable direct management of HA cluster members. Valid values: `enable` `disable` .
* `host_type` - Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both. No traps will be sent when IP type is subnet. Valid values: `any` `query` `trap` .
* `id` - Host entry ID.
* `ip` - IPv4 address of the SNMP manager (host).
* `source_ip` - Source IPv4 address for SNMP traps.
* `hosts6` - Configure IPv6 SNMP managers. The structure of `hosts6` block is documented below.

The `hosts6` block contains:

* `ha_direct` - Enable/disable direct management of HA cluster members. Valid values: `enable` `disable` .
* `host_type` - Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both. Valid values: `any` `query` `trap` .
* `id` - Host6 entry ID.
* `ipv6` - SNMP manager IPv6 address prefix.
* `source_ipv6` - Source IPv6 address for SNMP traps.
* `vdoms` - SNMP access control VDOMs. The structure of `vdoms` block is documented below.

The `vdoms` block contains:

* `name` - VDOM name. This attribute must reference one of the following datasources: `system.vdom.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_systemsnmp_community can be imported using:
```sh
terraform import fortios_systemsnmp_community.labelname {{mkey}}
```
