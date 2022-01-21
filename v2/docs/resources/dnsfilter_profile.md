---
subcategory: "FortiGate Dnsfilter"
layout: "fortios"
page_title: "FortiOS: fortios_dnsfilter_profile"
description: |-
  Configure DNS domain filter profile.
---

## fortios_dnsfilter_profile
Configure DNS domain filter profile.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `block_action` - Action to take for blocked domains. Valid values: `block` `redirect` `block-sevrfail` .
* `block_botnet` - Enable/disable blocking botnet C&C DNS lookups. Valid values: `disable` `enable` .
* `comment` - Comment.
* `log_all_domain` - Enable/disable logging of all domains visited (detailed DNS logging). Valid values: `enable` `disable` .
* `name` - Profile name.
* `redirect_portal` - IPv4 address of the SDNS redirect portal.
* `redirect_portal6` - IPv6 address of the SDNS redirect portal.
* `safe_search` - Enable/disable Google, Bing, YouTube, Qwant, DuckDuckGo safe search. Valid values: `disable` `enable` .
* `sdns_domain_log` - Enable/disable domain filtering and botnet domain logging. Valid values: `enable` `disable` .
* `sdns_ftgd_err_log` - Enable/disable FortiGuard SDNS rating error logging. Valid values: `enable` `disable` .
* `youtube_restrict` - Set safe search for YouTube restriction level. Valid values: `strict` `moderate` .
* `dns_translation` - DNS translation settings. The structure of `dns_translation` block is documented below.

The `dns_translation` block contains:

* `addr_type` - DNS translation type (IPv4 or IPv6). Valid values: `ipv4` `ipv6` .
* `dst` - IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
* `dst6` - IPv6 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src6.
* `id` - ID.
* `netmask` - If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
* `prefix` - If src6 and dst6 are subnets rather than single IP addresses, enter the prefix for both src6 and dst6 (1 - 128, default = 128).
* `src` - IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
* `src6` - IPv6 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst6.
* `status` - Enable/disable this DNS translation entry. Valid values: `enable` `disable` .
* `domain_filter` - Domain filter settings. The structure of `domain_filter` block is documented below.

The `domain_filter` block contains:

* `domain_filter_table` - DNS domain filter table ID. This attribute must reference one of the following datasources: `dnsfilter.domain-filter.id` .
* `external_ip_blocklist` - One or more external IP block lists. The structure of `external_ip_blocklist` block is documented below.

The `external_ip_blocklist` block contains:

* `name` - External domain block list name. This attribute must reference one of the following datasources: `system.external-resource.name` .
* `ftgd_dns` - FortiGuard DNS Filter settings. The structure of `ftgd_dns` block is documented below.

The `ftgd_dns` block contains:

* `options` - FortiGuard DNS filter options. Valid values: `error-allow` `ftgd-disable` .
* `filters` - FortiGuard DNS domain filters. The structure of `filters` block is documented below.

The `filters` block contains:

* `action` - Action to take for DNS requests matching the category. Valid values: `block` `monitor` .
* `category` - Category number.
* `id` - ID number.
* `log` - Enable/disable DNS filter logging for this DNS profile. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_dnsfilter_profile can be imported using:
```sh
terraform import fortios_dnsfilter_profile.labelname {{mkey}}
```
