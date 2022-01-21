---
subcategory: "FortiGate Dnsfilter"
layout: "fortios"
page_title: "FortiOS: fortios_dnsfilter_domainfilter"
description: |-
  Configure DNS domain filters.
---

## fortios_dnsfilter_domainfilter
Configure DNS domain filters.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comment` - Optional comments.
* `id` - ID.
* `name` - Name of table.
* `entries` - DNS domain filter entries. The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - Action to take for domain filter matches. Valid values: `block` `allow` `monitor` .
* `domain` - Domain entries to be filtered.
* `id` - Id.
* `status` - Enable/disable this domain filter. Valid values: `enable` `disable` .
* `type` - DNS domain filter type. Valid values: `simple` `regex` `wildcard` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_dnsfilter_domainfilter can be imported using:
```sh
terraform import fortios_dnsfilter_domainfilter.labelname {{mkey}}
```
