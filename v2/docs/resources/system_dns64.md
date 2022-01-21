---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_dns64"
description: |-
  Configure DNS64.
---

## fortios_system_dns64
Configure DNS64.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `always_synthesize_aaaa_record` - Enable/disable AAAA record synthesis (default = enable). Valid values: `enable` `disable` .
* `dns64_prefix` - DNS64 prefix must be ::/96 (default = 64:ff9b::/96).
* `status` - Enable/disable DNS64 (default = disable). Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_dns64 can be imported using:
```sh
terraform import fortios_system_dns64.labelname {{mkey}}
```
