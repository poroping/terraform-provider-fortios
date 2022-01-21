---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_urlmatch"
description: |-
  Exempt URLs from web proxy forwarding and caching.
---

## fortios_webproxy_urlmatch
Exempt URLs from web proxy forwarding and caching.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `cache_exemption` - Enable/disable exempting this URL pattern from caching. Valid values: `enable` `disable` .
* `comment` - Comment.
* `forward_server` - Forward server name. This attribute must reference one of the following datasources: `web-proxy.forward-server.name` `web-proxy.forward-server-group.name` .
* `name` - Configure a name for the URL to be exempted.
* `status` - Enable/disable exempting the URLs matching the URL pattern from web proxy forwarding and caching. Valid values: `enable` `disable` .
* `url_pattern` - URL pattern to be exempted from web proxy forwarding and caching.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_webproxy_urlmatch can be imported using:
```sh
terraform import fortios_webproxy_urlmatch.labelname {{mkey}}
```
