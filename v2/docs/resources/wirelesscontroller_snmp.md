---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_snmp"
description: |-
  Configure SNMP.
---

## fortios_wirelesscontroller_snmp
Configure SNMP.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `contact_info` - Contact Information.
* `engine_id` - AC SNMP engineId string (maximum 24 characters).
* `trap_high_cpu_threshold` - CPU usage when trap is sent.
* `trap_high_mem_threshold` - Memory usage when trap is sent.
* `community` - SNMP Community Configuration. The structure of `community` block is documented below.

The `community` block contains:

* `id` - Community ID.
* `name` - Community name.
* `query_v1_status` - Enable/disable SNMP v1 queries. Valid values: `enable` `disable` .
* `query_v2c_status` - Enable/disable SNMP v2c queries. Valid values: `enable` `disable` .
* `status` - Enable/disable this SNMP community. Valid values: `enable` `disable` .
* `trap_v1_status` - Enable/disable SNMP v1 traps. Valid values: `enable` `disable` .
* `trap_v2c_status` - Enable/disable SNMP v2c traps. Valid values: `enable` `disable` .
* `hosts` - Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.

The `hosts` block contains:

* `id` - Host entry ID.
* `ip` - IPv4 address of the SNMP manager (host).
* `user` - SNMP User Configuration. The structure of `user` block is documented below.

The `user` block contains:

* `auth_proto` - Authentication protocol. Valid values: `md5` `sha` .
* `auth_pwd` - Password for authentication protocol.
* `name` - SNMP User Name
* `notify_hosts` - Configure SNMP User Notify Hosts.
* `priv_proto` - Privacy (encryption) protocol. Valid values: `aes` `des` `aes256` `aes256cisco` .
* `priv_pwd` - Password for privacy (encryption) protocol.
* `queries` - Enable/disable SNMP queries for this user. Valid values: `enable` `disable` .
* `security_level` - Security level for message authentication and encryption. Valid values: `no-auth-no-priv` `auth-no-priv` `auth-priv` .
* `status` - SNMP User Enable Valid values: `enable` `disable` .
* `trap_status` - Enable/disable traps for this SNMP user. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontroller_snmp can be imported using:
```sh
terraform import fortios_wirelesscontroller_snmp.labelname {{mkey}}
```
