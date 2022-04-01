---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemsnmp_user"
description: |-
  Get information on a fortios SNMP user configuration.
---

# Data Source: fortios_systemsnmp_user
Use this data source to get information on a fortios SNMP user configuration.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) SNMP user name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `auth_proto` - Authentication protocol.
* `auth_pwd` - Password for authentication protocol.
* `events` - SNMP notifications (traps) to send.
* `ha_direct` - Enable/disable direct management of HA cluster members.
* `mib_view` - SNMP access control MIB view.
* `name` - SNMP user name.
* `notify_hosts` - SNMP managers to send notifications (traps) to.
* `notify_hosts6` - IPv6 SNMP managers to send notifications (traps) to.
* `priv_proto` - Privacy (encryption) protocol.
* `priv_pwd` - Password for privacy (encryption) protocol.
* `queries` - Enable/disable SNMP queries for this user.
* `query_port` - SNMPv3 query port (default = 161).
* `security_level` - Security level for message authentication and encryption.
* `source_ip` - Source IP for SNMP trap.
* `source_ipv6` - Source IPv6 for SNMP trap.
* `status` - Enable/disable this SNMP user.
* `trap_lport` - SNMPv3 local trap port (default = 162).
* `trap_rport` - SNMPv3 trap remote port (default = 162).
* `trap_status` - Enable/disable traps for this SNMP user.
* `vdoms` - SNMP access control VDOMs.The structure of `vdoms` block is documented below.

The `vdoms` block contains:

* `name` - VDOM name
