---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_policy46"
description: |-
  Get information on a fortios Configure IPv4 to IPv6 policies.
---

# Data Source: fortios_firewall_policy46
Use this data source to get information on a fortios Configure IPv4 to IPv6 policies.


## Example Usage

```hcl

```

## Argument Reference

* `policyid` - (Required) Policy ID (0 - 4294967294).
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `action` - Accept or deny traffic matching the policy.
* `comments` - Comment.
* `dstintf` - Destination interface name.
* `fixedport` - Enable/disable fixed port for this policy.
* `ippool` - Enable/disable use of IP Pools for source NAT.
* `logtraffic` - Enable/disable traffic logging for this policy.
* `logtraffic_start` - Record logs when a session starts and ends.
* `name` - Policy name.
* `per_ip_shaper` - Per IP traffic shaper.
* `permit_any_host` - Enable/disable allowing any host.
* `policyid` - Policy ID (0 - 4294967294).
* `schedule` - Schedule name.
* `srcintf` - Source interface name.
* `status` - Enable/disable this policy.
* `tcp_mss_receiver` - TCP Maximum Segment Size value of receiver (0 - 65535, default = 0)
* `tcp_mss_sender` - TCP Maximum Segment Size value of sender (0 - 65535, default = 0).
* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dstaddr` - Destination address objects.The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name.
* `poolname` - IP Pool names.The structure of `poolname` block is documented below.

The `poolname` block contains:

* `name` - IP pool name.
* `service` - Service name.The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service name.
* `srcaddr` - Source address objects.The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name.
