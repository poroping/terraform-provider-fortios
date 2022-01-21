---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logfortianalyzercloud_overridesetting"
description: |-
  Override FortiAnalyzer Cloud settings.
---

## fortios_logfortianalyzercloud_overridesetting
Override FortiAnalyzer Cloud settings.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `status` - Enable/disable logging to FortiAnalyzer. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_logfortianalyzercloud_overridesetting can be imported using:
```sh
terraform import fortios_logfortianalyzercloud_overridesetting.labelname {{mkey}}
```
