---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_interfacepolicy6"
description: |-
  Get information on a fortios Configure IPv6 interface policies.
---

# Data Source: fortios_firewall_interfacepolicy6
Use this data source to get information on a fortios Configure IPv6 interface policies.


## Example Usage

```hcl

```

## Argument Reference

* `policyid` - (Required) Policy ID (0 - 4294967295).
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `application_list` - Application list name.
* `application_list_status` - Enable/disable application control.
* `av_profile` - Antivirus profile.
* `av_profile_status` - Enable/disable antivirus.
* `comments` - Comments.
* `dlp_profile` - DLP profile name.
* `dlp_profile_status` - Enable/disable DLP.
* `dlp_sensor` - DLP sensor name.
* `dlp_sensor_status` - Enable/disable DLP.
* `dsri` - Enable/disable DSRI.
* `emailfilter_profile` - Email filter profile.
* `emailfilter_profile_status` - Enable/disable email filter.
* `interface` - Monitored interface name from available interfaces.
* `ips_sensor` - IPS sensor name.
* `ips_sensor_status` - Enable/disable IPS.
* `logtraffic` - Logging type to be used in this policy (Options: all | utm | disable, Default: utm).
* `policyid` - Policy ID (0 - 4294967295).
* `status` - Enable/disable this policy.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `webfilter_profile` - Web filter profile.
* `webfilter_profile_status` - Enable/disable web filtering.
* `dstaddr6` - IPv6 address object to limit traffic monitoring to network traffic sent to the specified address or range.The structure of `dstaddr6` block is documented below.

The `dstaddr6` block contains:

* `name` - Address name.
* `service6` - Service name.The structure of `service6` block is documented below.

The `service6` block contains:

* `name` - Service name.
* `srcaddr6` - IPv6 address object to limit traffic monitoring to network traffic sent from the specified address or range.The structure of `srcaddr6` block is documented below.

The `srcaddr6` block contains:

* `name` - Address name.
