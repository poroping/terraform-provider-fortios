---
subcategory: "FortiGate Webfilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_urlfilter"
description: |-
  Get information on a fortios Configure URL filter lists.
---

# Data Source: fortios_webfilter_urlfilter
Use this data source to get information on a fortios Configure URL filter lists.


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
* `ip_addr_block` - Enable/disable blocking URLs when the hostname appears as an IP address.
* `ip4_mapped_ip6` - Enable/disable matching of IPv4 mapped IPv6 URLs.
* `name` - Name of URL filter list.
* `one_arm_ips_urlfilter` - Enable/disable DNS resolver for one-arm IPS URL filter operation.
* `entries` - URL filter entries.The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - Action to take for URL filter matches.
* `antiphish_action` - Action to take for AntiPhishing matches.
* `dns_address_family` - Resolve IPv4 address, IPv6 address, or both from DNS server.
* `exempt` - If action is set to exempt, select the security profile operations that exempt URLs skip. Separate multiple options with a space.
* `id` - Id.
* `referrer_host` - Referrer host name.
* `status` - Enable/disable this URL filter.
* `type` - Filter type (simple, regex, or wildcard).
* `url` - URL to be filtered.
* `web_proxy_profile` - Web proxy profile.
