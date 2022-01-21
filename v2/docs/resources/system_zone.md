---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_zone"
description: |-
  Configure zones to group two or more interfaces. When a zone is created you can configure policies for the zone instead of individual interfaces in the zone.
---

## fortios_system_zone
Configure zones to group two or more interfaces. When a zone is created you can configure policies for the zone instead of individual interfaces in the zone.

## Example Usage

```hcl
resource "fortios_system_interface" "example" {
  allow_append = true
  vdomparam    = "root"

  name = "TESTACCINT2"
  type = "aggregate"
  vdom = "root"
}

resource "fortios_system_zone" "example" {
  vdomparam = fortios_system_interface.example.vdom

  name = "TESTZONE"
  interface {
    interface_name = fortios_system_interface.example.name
  }
}
```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `description` - Description.
* `intrazone` - Allow or deny traffic routing between different interfaces in the same zone (default = deny). Valid values: `allow` `deny` .
* `name` - Zone name.
* `interface` - Add interfaces to this zone. Interfaces must not be assigned to another zone or have firewall policies defined. The structure of `interface` block is documented below.

The `interface` block contains:

* `interface_name` - Select interfaces to add to the zone. This attribute must reference one of the following datasources: `system.interface.name` .
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.

The `tagging` block contains:

* `category` - Tag category. This attribute must reference one of the following datasources: `system.object-tagging.category` .
* `name` - Tagging entry name.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name. This attribute must reference one of the following datasources: `system.object-tagging.tags.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_zone can be imported using:
```sh
terraform import fortios_system_zone.labelname {{mkey}}
```
