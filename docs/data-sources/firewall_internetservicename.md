---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservicename"
description: |-
  Get information on a fortios Define internet service names.
---

# Data Source: fortios_firewall_internetservicename
Use this data source to get information on a fortios Define internet service names.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Internet Service name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `city_id` - City ID.
* `country_id` - Country or Area ID.
* `internet_service_id` - Internet Service ID.
* `name` - Internet Service name.
* `region_id` - Region ID.
* `type` - Internet Service name type.
