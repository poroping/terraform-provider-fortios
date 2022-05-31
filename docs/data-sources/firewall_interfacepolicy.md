---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_interfacepolicy"
description: |-
  Get information on a fortios Configure IPv4 interface policies.
---

# Data Source: fortios_firewall_interfacepolicy
Use this data source to get information on a fortios Configure IPv4 interface policies.


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
* `webfilter_profile` - Web filter profile.
* `webfilter_profile_status` - Enable/disable web filtering.
* `dstaddr` - Address object to limit traffic monitoring to network traffic sent to the specified address or range.The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name.
* `service` - Service object from available options.The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service name.
* `srcaddr` - Address object to limit traffic monitoring to network traffic sent from the specified address or range.The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name.
