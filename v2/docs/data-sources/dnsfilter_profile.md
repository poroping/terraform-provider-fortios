---
subcategory: "FortiGate Dnsfilter"
layout: "fortios"
page_title: "FortiOS: fortios_dnsfilter_profile"
description: |-
  Get information on a fortios Configure DNS domain filter profile.
---

# Data Source: fortios_dnsfilter_profile
Use this data source to get information on a fortios Configure DNS domain filter profile.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `block_action` - Action to take for blocked domains.
* `block_botnet` - Enable/disable blocking botnet C&C DNS lookups.
* `comment` - Comment.
* `log_all_domain` - Enable/disable logging of all domains visited (detailed DNS logging).
* `name` - Profile name.
* `redirect_portal` - IPv4 address of the SDNS redirect portal.
* `redirect_portal6` - IPv6 address of the SDNS redirect portal.
* `safe_search` - Enable/disable Google, Bing, YouTube, Qwant, DuckDuckGo safe search.
* `sdns_domain_log` - Enable/disable domain filtering and botnet domain logging.
* `sdns_ftgd_err_log` - Enable/disable FortiGuard SDNS rating error logging.
* `youtube_restrict` - Set safe search for YouTube restriction level.
* `dns_translation` - DNS translation settings.The structure of `dns_translation` block is documented below.

The `dns_translation` block contains:

* `addr_type` - DNS translation type (IPv4 or IPv6).
* `dst` - IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
* `dst6` - IPv6 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src6.
* `id` - ID.
* `netmask` - If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
* `prefix` - If src6 and dst6 are subnets rather than single IP addresses, enter the prefix for both src6 and dst6 (1 - 128, default = 128).
* `src` - IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
* `src6` - IPv6 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst6.
* `status` - Enable/disable this DNS translation entry.
* `domain_filter` - Domain filter settings.The structure of `domain_filter` block is documented below.

The `domain_filter` block contains:

* `domain_filter_table` - DNS domain filter table ID.
* `external_ip_blocklist` - One or more external IP block lists.The structure of `external_ip_blocklist` block is documented below.

The `external_ip_blocklist` block contains:

* `name` - External domain block list name.
* `ftgd_dns` - FortiGuard DNS Filter settings.The structure of `ftgd_dns` block is documented below.

The `ftgd_dns` block contains:

* `options` - FortiGuard DNS filter options.
* `filters` - FortiGuard DNS domain filters.The structure of `filters` block is documented below.

The `filters` block contains:

* `action` - Action to take for DNS requests matching the category.
* `category` - Category number.
* `id` - ID number.
* `log` - Enable/disable DNS filter logging for this DNS profile.
