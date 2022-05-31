---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemsnmp_user"
description: |-
  SNMP user configuration.
---

## fortios_systemsnmp_user
SNMP user configuration.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `auth_proto` - Authentication protocol. Valid values: `md5` `sha` `sha224` `sha256` `sha384` `sha512` .
* `auth_pwd` - Password for authentication protocol.
* `events` - SNMP notifications (traps) to send. Valid values: `cpu-high` `mem-low` `log-full` `intf-ip` `vpn-tun-up` `vpn-tun-down` `ha-switch` `ha-hb-failure` `ips-signature` `ips-anomaly` `av-virus` `av-oversize` `av-pattern` `av-fragmented` `fm-if-change` `fm-conf-change` `bgp-established` `bgp-backward-transition` `ha-member-up` `ha-member-down` `ent-conf-change` `av-conserve` `av-bypass` `av-oversize-passed` `av-oversize-blocked` `ips-pkg-update` `ips-fail-open` `faz-disconnect` `wc-ap-up` `wc-ap-down` `fswctl-session-up` `fswctl-session-down` `load-balance-real-server-down` `device-new` `per-cpu-high` `dhcp` `ospf-nbr-state-change` `ospf-virtnbr-state-change` .
* `ha_direct` - Enable/disable direct management of HA cluster members. Valid values: `enable` `disable` .
* `mib_view` - SNMP access control MIB view. This attribute must reference one of the following datasources: `system.snmp.mib-view.name` .
* `name` - SNMP user name.
* `notify_hosts` - SNMP managers to send notifications (traps) to.
* `notify_hosts6` - IPv6 SNMP managers to send notifications (traps) to.
* `priv_proto` - Privacy (encryption) protocol. Valid values: `aes` `des` `aes256` `aes256cisco` .
* `priv_pwd` - Password for privacy (encryption) protocol.
* `queries` - Enable/disable SNMP queries for this user. Valid values: `enable` `disable` .
* `query_port` - SNMPv3 query port (default = 161).
* `security_level` - Security level for message authentication and encryption. Valid values: `no-auth-no-priv` `auth-no-priv` `auth-priv` .
* `source_ip` - Source IP for SNMP trap.
* `source_ipv6` - Source IPv6 for SNMP trap.
* `status` - Enable/disable this SNMP user. Valid values: `enable` `disable` .
* `trap_lport` - SNMPv3 local trap port (default = 162).
* `trap_rport` - SNMPv3 trap remote port (default = 162).
* `trap_status` - Enable/disable traps for this SNMP user. Valid values: `enable` `disable` .
* `vdoms` - SNMP access control VDOMs. The structure of `vdoms` block is documented below.

The `vdoms` block contains:

* `name` - VDOM name This attribute must reference one of the following datasources: `system.vdom.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_systemsnmp_user can be imported using:
```sh
terraform import fortios_systemsnmp_user.labelname {{mkey}}
```
