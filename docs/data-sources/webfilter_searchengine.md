---
subcategory: "FortiGate Webfilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_searchengine"
description: |-
  Get information on a fortios Configure web filter search engines.
---

# Data Source: fortios_webfilter_searchengine
Use this data source to get information on a fortios Configure web filter search engines.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Search engine name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `charset` - Search engine charset.
* `hostname` - Hostname (regular expression).
* `name` - Search engine name.
* `query` - Code used to prefix a query (must end with an equals character).
* `safesearch` - Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.
* `safesearch_str` - Safe search parameter used in the URL in URL mode. In translate mode, it provides either the regex to translate the URL or the special case to translate the URL.
* `url` - URL (regular expression).
