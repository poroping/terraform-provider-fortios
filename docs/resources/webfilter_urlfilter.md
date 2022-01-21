---
subcategory: "FortiGate Webfilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_urlfilter"
description: |-
  Configure URL filter lists.
---

## fortios_webfilter_urlfilter
Configure URL filter lists.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comment` - Optional comments.
* `id` - ID.
* `ip_addr_block` - Enable/disable blocking URLs when the hostname appears as an IP address. Valid values: `enable` `disable` .
* `name` - Name of URL filter list.
* `one_arm_ips_urlfilter` - Enable/disable DNS resolver for one-arm IPS URL filter operation. Valid values: `enable` `disable` .
* `entries` - URL filter entries. The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - Action to take for URL filter matches. Valid values: `exempt` `block` `allow` `monitor` .
* `antiphish_action` - Action to take for AntiPhishing matches. Valid values: `block` `log` .
* `dns_address_family` - Resolve IPv4 address, IPv6 address, or both from DNS server. Valid values: `ipv4` `ipv6` `both` .
* `exempt` - If action is set to exempt, select the security profile operations that exempt URLs skip. Separate multiple options with a space. Valid values: `av` `web-content` `activex-java-cookie` `dlp` `fortiguard` `range-block` `pass` `antiphish` `all` .
* `id` - Id.
* `referrer_host` - Referrer host name.
* `status` - Enable/disable this URL filter. Valid values: `enable` `disable` .
* `type` - Filter type (simple, regex, or wildcard). Valid values: `simple` `regex` `wildcard` .
* `url` - URL to be filtered.
* `web_proxy_profile` - Web proxy profile. This attribute must reference one of the following datasources: `web-proxy.profile.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_webfilter_urlfilter can be imported using:
```sh
terraform import fortios_webfilter_urlfilter.labelname {{mkey}}
```
