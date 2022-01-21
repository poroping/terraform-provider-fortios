---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_utmprofile"
description: |-
  Get information on a fortios Configure UTM (Unified Threat Management) profile.
---

# Data Source: fortios_wirelesscontroller_utmprofile
Use this data source to get information on a fortios Configure UTM (Unified Threat Management) profile.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) UTM profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `antivirus_profile` - AntiVirus profile name.
* `application_list` - Application control list name.
* `comment` - Comment.
* `ips_sensor` - IPS sensor name.
* `name` - UTM profile name.
* `scan_botnet_connections` - Block or monitor connections to Botnet servers or disable Botnet scanning.
* `utm_log` - Enable/disable UTM logging.
* `webfilter_profile` - WebFilter profile name.
