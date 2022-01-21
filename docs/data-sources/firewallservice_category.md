---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallservice_category"
description: |-
  Get information on a fortios Configure service categories.
---

# Data Source: fortios_firewallservice_category
Use this data source to get information on a fortios Configure service categories.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Service category name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `fabric_object` - Security Fabric global object setting.
* `name` - Service category name.
