---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollerautoconfig_default"
description: |-
  Policies which are applied automatically to all ISL/ICL/FortiLink interfaces.
---

## fortios_switchcontrollerautoconfig_default
Policies which are applied automatically to all ISL/ICL/FortiLink interfaces.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `fgt_policy` - Default FortiLink auto-config policy. This attribute must reference one of the following datasources: `switch-controller.auto-config.policy.name` .
* `icl_policy` - Default ICL auto-config policy. This attribute must reference one of the following datasources: `switch-controller.auto-config.policy.name` .
* `isl_policy` - Default ISL auto-config policy. This attribute must reference one of the following datasources: `switch-controller.auto-config.policy.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_switchcontrollerautoconfig_default can be imported using:
```sh
terraform import fortios_switchcontrollerautoconfig_default.labelname {{mkey}}
```
