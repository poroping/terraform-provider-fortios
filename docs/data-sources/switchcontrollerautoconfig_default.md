---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollerautoconfig_default"
description: |-
  Get information on a fortios Policies which are applied automatically to all ISL/ICL/FortiLink interfaces.
---

# Data Source: fortios_switchcontrollerautoconfig_default
Use this data source to get information on a fortios Policies which are applied automatically to all ISL/ICL/FortiLink interfaces.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `fgt_policy` - Default FortiLink auto-config policy.
* `icl_policy` - Default ICL auto-config policy.
* `isl_policy` - Default ISL auto-config policy.
