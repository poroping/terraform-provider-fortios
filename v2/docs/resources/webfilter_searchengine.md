---
subcategory: "FortiGate Webfilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_searchengine"
description: |-
  Configure web filter search engines.
---

## fortios_webfilter_searchengine
Configure web filter search engines.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `charset` - Search engine charset. Valid values: `utf-8` `gb2312` .
* `hostname` - Hostname (regular expression).
* `name` - Search engine name.
* `query` - Code used to prefix a query (must end with an equals character).
* `safesearch` - Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header. Valid values: `disable` `url` `header` `translate` `yt-pattern` `yt-scan` `yt-video` `yt-channel` .
* `safesearch_str` - Safe search parameter used in the URL.
* `url` - URL (regular expression).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_webfilter_searchengine can be imported using:
```sh
terraform import fortios_webfilter_searchengine.labelname {{mkey}}
```
