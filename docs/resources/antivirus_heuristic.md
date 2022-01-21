---
subcategory: "FortiGate Antivirus"
layout: "fortios"
page_title: "FortiOS: fortios_antivirus_heuristic"
description: |-
  Configure global heuristic options.
---

## fortios_antivirus_heuristic
Configure global heuristic options.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `mode` - Enable/disable heuristics and determine how the system behaves if heuristics detects a problem. Valid values: `pass` `block` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_antivirus_heuristic can be imported using:
```sh
terraform import fortios_antivirus_heuristic.labelname {{mkey}}
```
