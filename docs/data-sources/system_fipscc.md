---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fipscc"
description: |-
  Get information on a fortios Configure FIPS-CC mode.
---

# Data Source: fortios_system_fipscc
Use this data source to get information on a fortios Configure FIPS-CC mode.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `entropy_token` - Enable/disable/dynamic entropy token.
* `key_generation_self_test` - Enable/disable self tests after key generation.
* `self_test_period` - Self test period.
* `status` - Enable/disable ciphers for FIPS mode of operation.
