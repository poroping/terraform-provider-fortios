---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_snmp"
description: |-
  Get information on a fortios Configure SNMP.
---

# Data Source: fortios_wirelesscontroller_snmp
Use this data source to get information on a fortios Configure SNMP.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `contact_info` - Contact Information.
* `engine_id` - AC SNMP engineId string (maximum 24 characters).
* `trap_high_cpu_threshold` - CPU usage when trap is sent.
* `trap_high_mem_threshold` - Memory usage when trap is sent.
* `community` - SNMP Community Configuration.The structure of `community` block is documented below.

The `community` block contains:

* `id` - Community ID.
* `name` - Community name.
* `query_v1_status` - Enable/disable SNMP v1 queries.
* `query_v2c_status` - Enable/disable SNMP v2c queries.
* `status` - Enable/disable this SNMP community.
* `trap_v1_status` - Enable/disable SNMP v1 traps.
* `trap_v2c_status` - Enable/disable SNMP v2c traps.
* `hosts` - Configure IPv4 SNMP managers (hosts).The structure of `hosts` block is documented below.

The `hosts` block contains:

* `id` - Host entry ID.
* `ip` - IPv4 address of the SNMP manager (host).
* `user` - SNMP User Configuration.The structure of `user` block is documented below.

The `user` block contains:

* `auth_proto` - Authentication protocol.
* `auth_pwd` - Password for authentication protocol.
* `name` - SNMP User Name
* `notify_hosts` - Configure SNMP User Notify Hosts.
* `priv_proto` - Privacy (encryption) protocol.
* `priv_pwd` - Password for privacy (encryption) protocol.
* `queries` - Enable/disable SNMP queries for this user.
* `security_level` - Security level for message authentication and encryption.
* `status` - SNMP User Enable
* `trap_status` - Enable/disable traps for this SNMP user.
