---
subcategory: "FortiGate Emailfilter"
layout: "fortios"
page_title: "FortiOS: fortios_emailfilter_fortishield"
description: |-
  Configure FortiGuard - AntiSpam.
---

## fortios_emailfilter_fortishield
Configure FortiGuard - AntiSpam.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `spam_submit_force` - Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable` `disable` .
* `spam_submit_srv` - Hostname of the spam submission server.
* `spam_submit_txt2htm` - Enable/disable conversion of text email to HTML email. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_emailfilter_fortishield can be imported using:
```sh
terraform import fortios_emailfilter_fortishield.labelname {{mkey}}
```
