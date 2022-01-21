---
subcategory: "FortiGate Application"
layout: "fortios"
page_title: "FortiOS: fortios_application_custom"
description: |-
  Configure custom application signatures.
---

## fortios_application_custom
Configure custom application signatures.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `tag` to be defined.

* `behavior` - Custom application signature behavior.
* `category` - Custom application category ID (use ? to view available options).
* `comment` - Comment.
* `id` - Custom application category ID (use ? to view available options).
* `protocol` - Custom application signature protocol.
* `signature` - The text that makes up the actual custom application signature.
* `tag` - Signature tag.
* `technology` - Custom application signature technology.
* `vendor` - Custom application signature vendor.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_application_custom can be imported using:
```sh
terraform import fortios_application_custom.labelname {{mkey}}
```
