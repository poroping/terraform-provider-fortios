---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_autoinstall"
description: |-
  Configure USB auto installation.
---

## fortios_system_autoinstall
Configure USB auto installation.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `auto_install_config` - Enable/disable auto install the config in USB disk. Valid values: `enable` `disable` .
* `auto_install_image` - Enable/disable auto install the image in USB disk. Valid values: `enable` `disable` .
* `default_config_file` - Default config file name in USB disk.
* `default_image_file` - Default image file name in USB disk.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_autoinstall can be imported using:
```sh
terraform import fortios_system_autoinstall.labelname {{mkey}}
```
