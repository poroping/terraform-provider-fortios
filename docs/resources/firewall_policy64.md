---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_policy64"
description: |-
  Configure IPv6 to IPv4 policies.
---

## fortios_firewall_policy64
Configure IPv6 to IPv4 policies.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `policyid` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `action` - Policy action. Valid values: `accept` `deny` .
* `comments` - Comment.
* `dstintf` - Destination interface name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` .
* `fixedport` - Enable/disable policy fixed port. Valid values: `enable` `disable` .
* `ippool` - Enable/disable policy64 IP pool. Valid values: `enable` `disable` .
* `logtraffic` - Enable/disable policy log traffic. Valid values: `enable` `disable` .
* `logtraffic_start` - Record logs when a session starts and ends. Valid values: `enable` `disable` .
* `name` - Policy name.
* `per_ip_shaper` - Per-IP traffic shaper. This attribute must reference one of the following datasources: `firewall.shaper.per-ip-shaper.name` .
* `permit_any_host` - Enable/disable permit any host in. Valid values: `enable` `disable` .
* `policyid` - Policy ID (0 - 4294967294).
* `schedule` - Schedule name. This attribute must reference one of the following datasources: `firewall.schedule.onetime.name` `firewall.schedule.recurring.name` `firewall.schedule.group.name` .
* `srcintf` - Source interface name. This attribute must reference one of the following datasources: `system.zone.name` `system.interface.name` .
* `status` - Enable/disable policy status. Valid values: `enable` `disable` .
* `tcp_mss_receiver` - TCP MSS value of receiver.
* `tcp_mss_sender` - TCP MSS value of sender.
* `traffic_shaper` - Traffic shaper. This attribute must reference one of the following datasources: `firewall.shaper.traffic-shaper.name` .
* `traffic_shaper_reverse` - Reverse traffic shaper. This attribute must reference one of the following datasources: `firewall.shaper.traffic-shaper.name` .
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dstaddr` - Destination address name. The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` `firewall.vip64.name` `firewall.vipgrp64.name` .
* `poolname` - Policy IP pool names. The structure of `poolname` block is documented below.

The `poolname` block contains:

* `name` - IP pool name. This attribute must reference one of the following datasources: `firewall.ippool.name` .
* `service` - Service name. The structure of `service` block is documented below.

The `service` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.service.custom.name` `firewall.service.group.name` .
* `srcaddr` - Source address name. The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_policy64 can be imported using:
```sh
terraform import fortios_firewall_policy64.labelname {{mkey}}
```
