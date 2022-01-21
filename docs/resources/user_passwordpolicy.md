---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_passwordpolicy"
description: |-
  Configure user password policy.
---

## fortios_user_passwordpolicy
Configure user password policy.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `expire_days` - Time in days before the user's password expires.
* `expired_password_renewal` - Enable/disable renewal of a password that already is expired. Valid values: `enable` `disable` .
* `name` - Password policy name.
* `warn_days` - Time in days before a password expiration warning message is displayed to the user upon login.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_user_passwordpolicy can be imported using:
```sh
terraform import fortios_user_passwordpolicy.labelname {{mkey}}
```
