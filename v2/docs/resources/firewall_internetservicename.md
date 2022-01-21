---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservicename"
description: |-
  Define internet service names.
---

## fortios_firewall_internetservicename
Define internet service names.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `city_id` - City ID. This attribute must reference one of the following datasources: `firewall.city.id` .
* `country_id` - Country or Area ID. This attribute must reference one of the following datasources: `firewall.country.id` .
* `internet_service_id` - Internet Service ID. This attribute must reference one of the following datasources: `firewall.internet-service.id` .
* `name` - Internet Service name.
* `region_id` - Region ID. This attribute must reference one of the following datasources: `firewall.region.id` .
* `type` - Internet Service name type. Valid values: `default` `location` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_internetservicename can be imported using:
```sh
terraform import fortios_firewall_internetservicename.labelname {{mkey}}
```
