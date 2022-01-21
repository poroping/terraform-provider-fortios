---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_policy64"
description: |-
  Get information on a fortios Configure IPv6 to IPv4 policies.
---

# Data Source: fortios_firewall_policy64
Use this data source to get information on a fortios Configure IPv6 to IPv4 policies.


## Example Usage

```hcl

```

## Argument Reference

* `policyid` - (Required) Policy ID (0 - 4294967294).
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `action` - Policy action.
* `comments` - Comment.
* `dstintf` - Destination interface name.
* `fixedport` - Enable/disable policy fixed port.
* `ippool` - Enable/disable policy64 IP pool.
* `logtraffic` - Enable/disable policy log traffic.
* `logtraffic_start` - Record logs when a session starts and ends.
* `name` - Policy name.
* `per_ip_shaper` - Per-IP traffic shaper.
* `permit_any_host` - Enable/disable permit any host in.
* `policyid` - Policy ID (0 - 4294967294).
* `schedule` - Schedule name.
* `srcintf` - Source interface name.
* `status` - Enable/disable policy status.
* `tcp_mss_receiver` - TCP MSS value of receiver.
* `tcp_mss_sender` - TCP MSS value of sender.
* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dstaddr` - Destination address name.The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name.
* `poolname` - Policy IP pool names.The structure of `poolname` block is documented below.

The `poolname` block contains:

* `name` - IP pool name.
* `service` - Service name.The structure of `service` block is documented below.

The `service` block contains:

* `name` - Address name.
* `srcaddr` - Source address name.The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name.
