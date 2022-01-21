---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_peergrp"
description: |-
  Configure peer groups.
---

## fortios_user_peergrp
Configure peer groups.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `name` - Peer group name.
* `member` - Peer group members. The structure of `member` block is documented below.

The `member` block contains:

* `name` - Peer group member name. This attribute must reference one of the following datasources: `user.peer.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_user_peergrp can be imported using:
```sh
terraform import fortios_user_peergrp.labelname {{mkey}}
```
