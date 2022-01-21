---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomradiusserver"
description: |-
  Get information on a fortios Configure a RADIUS server to use as a RADIUS Single Sign On (RSSO) server for this VDOM.
---

# Data Source: fortios_system_vdomradiusserver
Use this data source to get information on a fortios Configure a RADIUS server to use as a RADIUS Single Sign On (RSSO) server for this VDOM.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name of the VDOM that you are adding the RADIUS server to.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - Name of the VDOM that you are adding the RADIUS server to.
* `radius_server_vdom` - Use this option to select another VDOM containing a VDOM RSSO RADIUS server to use for the current VDOM.
* `status` - Enable/disable the RSSO RADIUS server for this VDOM.
