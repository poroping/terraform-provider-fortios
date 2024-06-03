---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fipscc"
description: |-
  Configure FIPS-CC mode.
---

## fortios_system_fipscc
Configure FIPS-CC mode.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `entropy_token` - Enable/disable/dynamic entropy token. Valid values: `disable` .
* `key_generation_self_test` - Enable/disable self tests after key generation. Valid values: `enable` `disable` .
* `self_test_period` - Self test period.
* `status` - Enable/disable ciphers for FIPS mode of operation. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_fipscc can be imported using:
```sh
terraform import fortios_system_fipscc.labelname {{mkey}}
```
