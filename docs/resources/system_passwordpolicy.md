---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_passwordpolicy"
description: |-
  Configure password policy for locally defined administrator passwords and IPsec VPN pre-shared keys.
---

## fortios_system_passwordpolicy
Configure password policy for locally defined administrator passwords and IPsec VPN pre-shared keys.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `apply_to` - Apply password policy to administrator passwords or IPsec pre-shared keys or both. Separate entries with a space. Valid values: `admin-password` `ipsec-preshared-key` .
* `change_4_characters` - Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled). Valid values: `enable` `disable` .
* `expire_day` - Number of days after which passwords expire (1 - 999 days, default = 90).
* `expire_status` - Enable/disable password expiration. Valid values: `enable` `disable` .
* `min_change_characters` - Minimum number of unique characters in new password which do not exist in old password (0 - 128, default = 0. This attribute overrides reuse-password if both are enabled).
* `min_lower_case_letter` - Minimum number of lowercase characters in password (0 - 128, default = 0).
* `min_non_alphanumeric` - Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
* `min_number` - Minimum number of numeric characters in password (0 - 128, default = 0).
* `min_upper_case_letter` - Minimum number of uppercase characters in password (0 - 128, default = 0).
* `minimum_length` - Minimum password length (8 - 128, default = 8).
* `reuse_password` - Enable/disable reuse of password. If both reuse-password and min-change-characters are enabled, min-change-characters overrides. Valid values: `enable` `disable` .
* `status` - Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_passwordpolicy can be imported using:
```sh
terraform import fortios_system_passwordpolicy.labelname {{mkey}}
```
