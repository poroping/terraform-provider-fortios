---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_console"
description: |-
  Configure console.
---

## fortios_system_console
Configure console.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `baudrate` - Console baud rate. Valid values: `9600` `19200` `38400` `57600` `115200` .
* `fortiexplorer` - Enable/disable access for FortiExplorer. Valid values: `enable` `disable` .
* `login` - Enable/disable serial console and FortiExplorer. Valid values: `enable` `disable` .
* `mode` - Console mode. Valid values: `batch` `line` .
* `output` - Console output mode. Valid values: `standard` `more` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_console can be imported using:
```sh
terraform import fortios_system_console.labelname {{mkey}}
```
