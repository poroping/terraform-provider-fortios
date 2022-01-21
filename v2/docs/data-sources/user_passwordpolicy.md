---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_passwordpolicy"
description: |-
  Get information on a fortios Configure user password policy.
---

# Data Source: fortios_user_passwordpolicy
Use this data source to get information on a fortios Configure user password policy.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Password policy name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `expire_days` - Time in days before the user's password expires.
* `expired_password_renewal` - Enable/disable renewal of a password that already is expired.
* `name` - Password policy name.
* `warn_days` - Time in days before a password expiration warning message is displayed to the user upon login.
