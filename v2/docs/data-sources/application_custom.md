---
subcategory: "FortiGate Application"
layout: "fortios"
page_title: "FortiOS: fortios_application_custom"
description: |-
  Get information on a fortios Configure custom application signatures.
---

# Data Source: fortios_application_custom
Use this data source to get information on a fortios Configure custom application signatures.


## Example Usage

```hcl

```

## Argument Reference

* `tag` - (Required) Signature tag.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `behavior` - Custom application signature behavior.
* `category` - Custom application category ID (use ? to view available options).
* `comment` - Comment.
* `id` - Custom application category ID (use ? to view available options).
* `protocol` - Custom application signature protocol.
* `signature` - The text that makes up the actual custom application signature.
* `tag` - Signature tag.
* `technology` - Custom application signature technology.
* `vendor` - Custom application signature vendor.
