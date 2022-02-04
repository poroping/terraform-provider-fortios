---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_fortitoken"
description: |-
  Configure FortiToken.
---

## fortios_user_fortitoken
Configure FortiToken.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `serial_number` to be defined.

* `activation_code` - Mobile token user activation-code.
* `activation_expire` - Mobile token user activation-code expire time.
* `comments` - Comment.
* `license` - Mobile token license.
* `os_ver` - Device Mobile Version.
* `reg_id` - Device Reg ID.
* `serial_number` - Serial number.
* `status` - Status. Valid values: `active` `lock` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_user_fortitoken can be imported using:
```sh
terraform import fortios_user_fortitoken.labelname {{mkey}}
```
