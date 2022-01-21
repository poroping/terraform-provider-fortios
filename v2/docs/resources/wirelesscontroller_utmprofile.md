---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_utmprofile"
description: |-
  Configure UTM (Unified Threat Management) profile.
---

## fortios_wirelesscontroller_utmprofile
Configure UTM (Unified Threat Management) profile.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `antivirus_profile` - AntiVirus profile name. This attribute must reference one of the following datasources: `antivirus.profile.name` .
* `application_list` - Application control list name. This attribute must reference one of the following datasources: `application.list.name` .
* `comment` - Comment.
* `ips_sensor` - IPS sensor name. This attribute must reference one of the following datasources: `ips.sensor.name` .
* `name` - UTM profile name.
* `scan_botnet_connections` - Block or monitor connections to Botnet servers or disable Botnet scanning. Valid values: `disable` `monitor` `block` .
* `utm_log` - Enable/disable UTM logging. Valid values: `enable` `disable` .
* `webfilter_profile` - WebFilter profile name. This attribute must reference one of the following datasources: `webfilter.profile.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontroller_utmprofile can be imported using:
```sh
terraform import fortios_wirelesscontroller_utmprofile.labelname {{mkey}}
```
