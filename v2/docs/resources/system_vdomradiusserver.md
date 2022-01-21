---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomradiusserver"
description: |-
  Configure a RADIUS server to use as a RADIUS Single Sign On (RSSO) server for this VDOM.
---

## fortios_system_vdomradiusserver
Configure a RADIUS server to use as a RADIUS Single Sign On (RSSO) server for this VDOM.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `name` - Name of the VDOM that you are adding the RADIUS server to. This attribute must reference one of the following datasources: `system.vdom.name` .
* `radius_server_vdom` - Use this option to select another VDOM containing a VDOM RSSO RADIUS server to use for the current VDOM. This attribute must reference one of the following datasources: `system.vdom.name` .
* `status` - Enable/disable the RSSO RADIUS server for this VDOM. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_vdomradiusserver can be imported using:
```sh
terraform import fortios_system_vdomradiusserver.labelname {{mkey}}
```
