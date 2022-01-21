---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemreplacemsg_fortiguardwf"
description: |-
  Get information on a fortios Replacement messages.
---

# Data Source: fortios_systemreplacemsg_fortiguardwf
Use this data source to get information on a fortios Replacement messages.


## Example Usage

```hcl

```

## Argument Reference

* `msg_type` - (Required) Message type.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `buffer` - Message string.
* `format` - Format flag.
* `header` - Header flag.
* `msg_type` - Message type.
