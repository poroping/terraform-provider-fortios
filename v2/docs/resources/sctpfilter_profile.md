---
subcategory: "FortiGate Sctp-Filter"
layout: "fortios"
page_title: "FortiOS: fortios_sctpfilter_profile"
description: |-
  Configure SCTP filter profiles.
---

## fortios_sctpfilter_profile
Configure SCTP filter profiles.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comment` - Comment.
* `name` - Profile name.
* `ppid_filters` - PPID filters list. The structure of `ppid_filters` block is documented below.

The `ppid_filters` block contains:

* `action` - Action taken when PPID is matched. Valid values: `pass` `reset` `replace` .
* `comment` - Comment.
* `id` - ID.
* `ppid` - Payload protocol identifier.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_sctpfilter_profile can be imported using:
```sh
terraform import fortios_sctpfilter_profile.labelname {{mkey}}
```
