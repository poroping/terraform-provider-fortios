---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_adgrp"
description: |-
  Configure FSSO groups.
---

## fortios_user_adgrp
Configure FSSO groups.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `connector_source` - FSSO connector source.
* `id` - Group ID.
* `name` - Name.
* `server_name` - FSSO agent name.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_user_adgrp can be imported using:
```sh
terraform import fortios_user_adgrp.labelname {{mkey}}
```
