---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_peergrp"
description: |-
  Get information on a fortios Configure peer groups.
---

# Data Source: fortios_user_peergrp
Use this data source to get information on a fortios Configure peer groups.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Peer group name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - Peer group name.
* `member` - Peer group members.The structure of `member` block is documented below.

The `member` block contains:

* `name` - Peer group member name.
