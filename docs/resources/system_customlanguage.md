---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_customlanguage"
description: |-
  Configure custom languages.
---

## fortios_system_customlanguage
Configure custom languages.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `comments` - Comment.
* `filename` - Custom language file path.
* `name` - Name.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_customlanguage can be imported using:
```sh
terraform import fortios_system_customlanguage.labelname {{mkey}}
```
