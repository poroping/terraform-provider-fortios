---
subcategory: "FortiGate Dnsfilter"
layout: "fortios"
page_title: "FortiOS: fortios_dnsfilter_domainfilter"
description: |-
  Get information on a fortios Configure DNS domain filters.
---

# Data Source: fortios_dnsfilter_domainfilter
Use this data source to get information on a fortios Configure DNS domain filters.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Optional comments.
* `id` - ID.
* `name` - Name of table.
* `entries` - DNS domain filter entries.The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - Action to take for domain filter matches.
* `domain` - Domain entries to be filtered.
* `id` - Id.
* `status` - Enable/disable this domain filter.
* `type` - DNS domain filter type.
