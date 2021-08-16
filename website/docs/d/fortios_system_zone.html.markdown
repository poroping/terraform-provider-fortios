---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_zone"
description: |-
  Get information on a fortios Configure zones to group two or more interfaces. When a zone is created you can configure policies for the zone instead of individual interfaces in the zone.
---

# Data Source: fortios_system_zone
Use this data source to get information on a fortios Configure zones to group two or more interfaces. When a zone is created you can configure policies for the zone instead of individual interfaces in the zone.

## Example Usage

```hcl
data "fortios_system_zone" sample1 {
    name = "outside"
}

output output1 {
  value = data.fortios_system_zone.sample1
}
```

## Argument Reference

* `name` - (Required) Zone name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `description` - Description.
* `intrazone` - Allow or deny traffic routing between different interfaces in the same zone (default = deny).
* `name` - Zone name.
* `interface` - Add interfaces to this zone. Interfaces must not be assigned to another zone or have firewall policies defined.The structure of `interface` block is documented below.

The `interface` block contains:

* `interface_name` - Select interfaces to add to the zone.
* `tagging` - Config object tagging.The structure of `tagging` block is documented below.

The `tagging` block contains:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.
