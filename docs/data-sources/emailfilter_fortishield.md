---
subcategory: "FortiGate Emailfilter"
layout: "fortios"
page_title: "FortiOS: fortios_emailfilter_fortishield"
description: |-
  Get information on a fortios Configure FortiGuard - AntiSpam.
---

# Data Source: fortios_emailfilter_fortishield
Use this data source to get information on a fortios Configure FortiGuard - AntiSpam.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `spam_submit_force` - Enable/disable force insertion of a new mime entity for the submission text.
* `spam_submit_srv` - Hostname of the spam submission server.
* `spam_submit_txt2htm` - Enable/disable conversion of text email to HTML email.
