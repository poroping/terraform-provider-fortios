---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_geoipoverride"
description: |-
  Configure geographical location mapping for IP address(es) to override mappings from FortiGuard.
---

## fortios_system_geoipoverride
Configure geographical location mapping for IP address(es) to override mappings from FortiGuard.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `country_id` - Two character Country ID code.
* `description` - Description.
* `name` - Location name.
* `ip_range` - Table of IP ranges assigned to country. The structure of `ip_range` block is documented below.

The `ip_range` block contains:

* `end_ip` - Ending IP address, inclusive, of the address range (format: xxx.xxx.xxx.xxx).
* `id` - ID of individual entry in the IP range table.
* `start_ip` - Starting IP address, inclusive, of the address range (format: xxx.xxx.xxx.xxx).
* `ip6_range` - Table of IPv6 ranges assigned to country. The structure of `ip6_range` block is documented below.

The `ip6_range` block contains:

* `end_ip` - Ending IP address, inclusive, of the address range (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `id` - ID of individual entry in the IPv6 range table.
* `start_ip` - Starting IP address, inclusive, of the address range (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_geoipoverride can be imported using:
```sh
terraform import fortios_system_geoipoverride.labelname {{mkey}}
```
