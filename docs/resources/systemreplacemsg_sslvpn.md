---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemreplacemsg_sslvpn"
description: |-
  Replacement messages.
---

## fortios_systemreplacemsg_sslvpn
Replacement messages.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `msg_type` to be defined.

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none` `text` `html` .
* `header` - Header flag. Valid values: `none` `http` `8bit` .
* `msg_type` - Message type.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_systemreplacemsg_sslvpn can be imported using:
```sh
terraform import fortios_systemreplacemsg_sslvpn.labelname {{mkey}}
```
