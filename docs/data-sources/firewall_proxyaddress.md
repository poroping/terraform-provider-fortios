---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_proxyaddress"
description: |-
  Get information on a fortios Configure web proxy address.
---

# Data Source: fortios_firewall_proxyaddress
Use this data source to get information on a fortios Configure web proxy address.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Address name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `case_sensitivity` - Enable to make the pattern case sensitive.
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
* `comment` - Optional comments.
* `header` - HTTP header name as a regular expression.
* `header_name` - Name of HTTP header.
* `host` - Address object for the host.
* `host_regex` - Host name as a regular expression.
* `method` - HTTP request methods to be used.
* `name` - Address name.
* `path` - URL path as a regular expression.
* `query` - Match the query part of the URL as a regular expression.
* `referrer` - Enable/disable use of referrer field in the HTTP header to match the address.
* `type` - Proxy address type.
* `ua` - Names of browsers to be used as user agent.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable visibility of the object in the GUI.
* `category` - FortiGuard category ID.The structure of `category` block is documented below.

The `category` block contains:

* `id` - FortiGuard category ID.
* `header_group` - HTTP header group.The structure of `header_group` block is documented below.

The `header_group` block contains:

* `case_sensitivity` - Case sensitivity in pattern.
* `header` - HTTP header regular expression.
* `header_name` - HTTP header.
* `id` - ID.
* `tagging` - Config object tagging.The structure of `tagging` block is documented below.

The `tagging` block contains:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.
