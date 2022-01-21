---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_lognulldevice_setting"
description: |-
  Get information on a fortios Settings for null device logging.
---

# Data Source: fortios_lognulldevice_setting
Use this data source to get information on a fortios Settings for null device logging.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `status` - Enable/disable statistics collection for when no external logging destination, such as FortiAnalyzer, is present (data is not saved).
