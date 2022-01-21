---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logwebtrends_setting"
description: |-
  Settings for WebTrends.
---

## fortios_logwebtrends_setting
Settings for WebTrends.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `server` - Address of the remote WebTrends server.
* `status` - Enable/disable logging to WebTrends. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_logwebtrends_setting can be imported using:
```sh
terraform import fortios_logwebtrends_setting.labelname {{mkey}}
```
