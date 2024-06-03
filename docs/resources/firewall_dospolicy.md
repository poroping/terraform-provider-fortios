---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_dospolicy"
description: |-
  Configure IPv4 DoS policies.
---

## fortios_firewall_dospolicy
Configure IPv4 DoS policies.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `policyid` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comments` - Comment.
* `interface` - Incoming interface name from available interfaces. This attribute must reference one of the following datasources: `system.zone.name` `system.interface.name` .
* `name` - Policy name.
* `policyid` - Policy ID.
* `status` - Enable/disable this policy. Valid values: `enable` `disable` .
* `anomaly` - Anomaly name. The structure of `anomaly` block is documented below.

The `anomaly` block contains:

* `action` - Action taken when the threshold is reached. Valid values: `pass` `block` `proxy` .
* `log` - Enable/disable anomaly logging. Valid values: `enable` `disable` .
* `name` - Anomaly name.
* `quarantine` - Quarantine method. Valid values: `none` `attacker` .
* `quarantine_expiry` - Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.
* `quarantine_log` - Enable/disable quarantine logging. Valid values: `disable` `enable` .
* `status` - Enable/disable this anomaly. Valid values: `disable` `enable` .
* `synproxy_tcp_mss` - Determine TCP maximum segment size (MSS) value for packets replied by syn proxy module. Valid values: `0` `256` `512` `1024` `1300` `1360` `1460` `1500` .
* `synproxy_tcp_sack` - enable/disable TCP selective acknowledage (SACK) for packets replied by syn proxy module. Valid values: `enable` `disable` .
* `synproxy_tcp_timestamp` - enable/disable TCP timestamp option for packets replied by syn proxy module. Valid values: `enable` `disable` .
* `synproxy_tcp_window` - Determine TCP Window size for packets replied by syn proxy module. Valid values: `4096` `8192` `16384` `32768` .
* `synproxy_tcp_windowscale` - Determine TCP window scale option value for packets replied by syn proxy module. Valid values: `0` `1` `2` `3` `4` `5` `6` `7` `8` `9` `10` `11` `12` `13` `14` .
* `synproxy_tos` - Determine TCP differentiated services code point value (type of service). Valid values: `0` `10` `12` `14` `18` `20` `22` `26` `28` `30` `34` `36` `38` `40` `46` `255` .
* `synproxy_ttl` - Determine Time to live (TTL) value for packets replied by syn proxy module. Valid values: `32` `64` `128` `255` .
* `threshold` - Anomaly threshold. Number of detected instances (packets per second or concurrent session number) that triggers the anomaly action.
* `thresholddefault` - Number of detected instances (packets per second or concurrent session number) which triggers action (1 - 2147483647, default = 1000). Note that each anomaly has a different threshold value assigned to it.
* `dstaddr` - Destination address name from available addresses. The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `service` - Service object from available options. The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service name. This attribute must reference one of the following datasources: `firewall.service.custom.name` `firewall.service.group.name` .
* `srcaddr` - Source address name from available addresses. The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_dospolicy can be imported using:
```sh
terraform import fortios_firewall_dospolicy.labelname {{mkey}}
```
