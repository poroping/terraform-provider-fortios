---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_interfacepolicy6"
description: |-
  Configure IPv6 interface policies.
---

## fortios_firewall_interfacepolicy6
Configure IPv6 interface policies.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `policyid` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `application_list` - Application list name. This attribute must reference one of the following datasources: `application.list.name` .
* `application_list_status` - Enable/disable application control. Valid values: `enable` `disable` .
* `av_profile` - Antivirus profile. This attribute must reference one of the following datasources: `antivirus.profile.name` .
* `av_profile_status` - Enable/disable antivirus. Valid values: `enable` `disable` .
* `comments` - Comments.
* `dlp_sensor` - DLP sensor name. This attribute must reference one of the following datasources: `dlp.sensor.name` .
* `dlp_sensor_status` - Enable/disable DLP. Valid values: `enable` `disable` .
* `dsri` - Enable/disable DSRI. Valid values: `enable` `disable` .
* `emailfilter_profile` - Email filter profile. This attribute must reference one of the following datasources: `emailfilter.profile.name` .
* `emailfilter_profile_status` - Enable/disable email filter. Valid values: `enable` `disable` .
* `interface` - Monitored interface name from available interfaces. This attribute must reference one of the following datasources: `system.zone.name` `system.interface.name` .
* `ips_sensor` - IPS sensor name. This attribute must reference one of the following datasources: `ips.sensor.name` .
* `ips_sensor_status` - Enable/disable IPS. Valid values: `enable` `disable` .
* `logtraffic` - Logging type to be used in this policy (Options: all | utm | disable, Default: utm). Valid values: `all` `utm` `disable` .
* `policyid` - Policy ID (0 - 4294967295).
* `status` - Enable/disable this policy. Valid values: `enable` `disable` .
* `webfilter_profile` - Web filter profile. This attribute must reference one of the following datasources: `webfilter.profile.name` .
* `webfilter_profile_status` - Enable/disable web filtering. Valid values: `enable` `disable` .
* `dstaddr6` - IPv6 address object to limit traffic monitoring to network traffic sent to the specified address or range. The structure of `dstaddr6` block is documented below.

The `dstaddr6` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `service6` - Service name. The structure of `service6` block is documented below.

The `service6` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.service.custom.name` `firewall.service.group.name` .
* `srcaddr6` - IPv6 address object to limit traffic monitoring to network traffic sent from the specified address or range. The structure of `srcaddr6` block is documented below.

The `srcaddr6` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_interfacepolicy6 can be imported using:
```sh
terraform import fortios_firewall_interfacepolicy6.labelname {{mkey}}
```
