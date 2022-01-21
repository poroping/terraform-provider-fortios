---
subcategory: "FortiGate Certificate"
layout: "fortios"
page_title: "FortiOS: fortios_certificate_remote"
description: |-
  Remote certificate as a PEM file.
---

## fortios_certificate_remote
Remote certificate as a PEM file.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `name` - Name.
* `range` - Either the global or VDOM IP address range for the remote certificate. Valid values: `global` `vdom` .
* `remote` - Remote certificate.
* `source` - Remote certificate source type. Valid values: `factory` `user` `bundle` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_certificate_remote can be imported using:
```sh
terraform import fortios_certificate_remote.labelname {{mkey}}
```
