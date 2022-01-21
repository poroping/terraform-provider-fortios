---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_adgrp"
description: |-
  Get information on a fortios Configure FSSO groups.
---

# Data Source: fortios_user_adgrp
Use this data source to get information on a fortios Configure FSSO groups.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `connector_source` - FSSO connector source.
* `id` - Group ID.
* `name` - Name.
* `server_name` - FSSO agent name.
