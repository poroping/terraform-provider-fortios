---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_passwordpolicyguestadmin"
description: |-
  Get information on a fortios Configure the password policy for guest administrators.
---

# Data Source: fortios_system_passwordpolicyguestadmin
Use this data source to get information on a fortios Configure the password policy for guest administrators.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `apply_to` - Guest administrator to which this password policy applies.
* `change_4_characters` - Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled).
* `expire_day` - Number of days after which passwords expire (1 - 999 days, default = 90).
* `expire_status` - Enable/disable password expiration.
* `min_change_characters` - Minimum number of unique characters in new password which do not exist in old password (0 - 128, default = 0. This attribute overrides reuse-password if both are enabled).
* `min_lower_case_letter` - Minimum number of lowercase characters in password (0 - 128, default = 0).
* `min_non_alphanumeric` - Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
* `min_number` - Minimum number of numeric characters in password (0 - 128, default = 0).
* `min_upper_case_letter` - Minimum number of uppercase characters in password (0 - 128, default = 0).
* `minimum_length` - Minimum password length (8 - 128, default = 8).
* `reuse_password` - Enable/disable reuse of password. If both reuse-password and min-change-characters are enabled, min-change-characters overrides.
* `status` - Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys.
