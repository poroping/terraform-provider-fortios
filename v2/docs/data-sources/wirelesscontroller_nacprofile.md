---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_nacprofile"
description: |-
  Get information on a fortios Configure WiFi network access control (NAC) profiles.
---

# Data Source: fortios_wirelesscontroller_nacprofile
Use this data source to get information on a fortios Configure WiFi network access control (NAC) profiles.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `name` - Name.
* `onboarding_vlan` - VLAN interface name.
