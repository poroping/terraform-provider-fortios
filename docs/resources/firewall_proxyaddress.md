---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_proxyaddress"
description: |-
  Configure web proxy address.
---

## fortios_firewall_proxyaddress
Configure web proxy address.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `case_sensitivity` - Enable to make the pattern case sensitive. Valid values: `disable` `enable` .
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
* `comment` - Optional comments.
* `header` - HTTP header name as a regular expression.
* `header_name` - Name of HTTP header.
* `host` - Address object for the host. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` `firewall.proxy-address.name` .
* `host_regex` - Host name as a regular expression.
* `method` - HTTP request methods to be used. Valid values: `get` `post` `put` `head` `connect` `trace` `options` `delete` .
* `name` - Address name.
* `path` - URL path as a regular expression.
* `query` - Match the query part of the URL as a regular expression.
* `referrer` - Enable/disable use of referrer field in the HTTP header to match the address. Valid values: `enable` `disable` .
* `type` - Proxy address type. Valid values: `host-regex` `url` `category` `method` `ua` `header` `src-advanced` `dst-advanced` .
* `ua` - Names of browsers to be used as user agent. Valid values: `chrome` `ms` `firefox` `safari` `other` .
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable visibility of the object in the GUI. Valid values: `enable` `disable` .
* `category` - FortiGuard category ID. The structure of `category` block is documented below.

The `category` block contains:

* `id` - FortiGuard category ID.
* `header_group` - HTTP header group. The structure of `header_group` block is documented below.

The `header_group` block contains:

* `case_sensitivity` - Case sensitivity in pattern. Valid values: `disable` `enable` .
* `header` - HTTP header regular expression.
* `header_name` - HTTP header.
* `id` - ID.
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.

The `tagging` block contains:

* `category` - Tag category. This attribute must reference one of the following datasources: `system.object-tagging.category` .
* `name` - Tagging entry name.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name. This attribute must reference one of the following datasources: `system.object-tagging.tags.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_proxyaddress can be imported using:
```sh
terraform import fortios_firewall_proxyaddress.labelname {{mkey}}
```
