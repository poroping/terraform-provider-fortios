---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_h2qptermsandconditions"
description: |-
  Configure terms and conditions.
---

## fortios_wirelesscontrollerhotspot20_h2qptermsandconditions
Configure terms and conditions.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `filename` - Filename.
* `name` - Terms and Conditions ID.
* `timestamp` - Timestamp.
* `url` - URL.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontrollerhotspot20_h2qptermsandconditions can be imported using:
```sh
terraform import fortios_wirelesscontrollerhotspot20_h2qptermsandconditions.labelname {{mkey}}
```
