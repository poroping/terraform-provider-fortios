---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_dns64"
description: |-
  Get information on a fortios Configure DNS64.
---

# Data Source: fortios_system_dns64
Use this data source to get information on a fortios Configure DNS64.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `always_synthesize_aaaa_record` - Enable/disable AAAA record synthesis (default = enable).
* `dns64_prefix` - DNS64 prefix must be ::/96 (default = 64:ff9b::/96).
* `status` - Enable/disable DNS64 (default = disable).
