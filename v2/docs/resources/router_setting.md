---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_setting"
description: |-
  Configure router settings.
---

## fortios_router_setting
Configure router settings.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `hostname` - Hostname for this virtual domain router.
* `show_filter` - Prefix-list as filter for showing routes. This attribute must reference one of the following datasources: `router.prefix-list.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_setting can be imported using:
```sh
terraform import fortios_router_setting.labelname {{mkey}}
```
