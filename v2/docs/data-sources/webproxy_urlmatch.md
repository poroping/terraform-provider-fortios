---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_urlmatch"
description: |-
  Get information on a fortios Exempt URLs from web proxy forwarding and caching.
---

# Data Source: fortios_webproxy_urlmatch
Use this data source to get information on a fortios Exempt URLs from web proxy forwarding and caching.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Configure a name for the URL to be exempted.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `cache_exemption` - Enable/disable exempting this URL pattern from caching.
* `comment` - Comment.
* `forward_server` - Forward server name.
* `name` - Configure a name for the URL to be exempted.
* `status` - Enable/disable exempting the URLs matching the URL pattern from web proxy forwarding and caching.
* `url_pattern` - URL pattern to be exempted from web proxy forwarding and caching.
