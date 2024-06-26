---
subcategory: "FortiGate Wanopt"
layout: "fortios"
page_title: "FortiOS: fortios_wanopt_webcache"
description: |-
  Configure global Web cache settings.
---

## fortios_wanopt_webcache
Configure global Web cache settings.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `always_revalidate` - Enable/disable revalidation of requested cached objects, which have content on the server, before serving it to the client. Valid values: `enable` `disable` .
* `cache_by_default` - Enable/disable caching content that lacks explicit caching policies from the server. Valid values: `enable` `disable` .
* `cache_cookie` - Enable/disable caching cookies. Since cookies contain information for or about individual users, they not usually cached. Valid values: `enable` `disable` .
* `cache_expired` - Enable/disable caching type-1 objects that are already expired on arrival. Valid values: `enable` `disable` .
* `default_ttl` - Default object expiry time (default = 1440 min (1 day); maximum = 5256000 min (10 years)). This only applies to those objects that do not have an expiry time set by the web server.
* `external` - Enable/disable external Web caching. Valid values: `enable` `disable` .
* `fresh_factor` - Frequency that the server is checked to see if any objects have expired (1 - 100, default = 100). The higher the fresh factor, the less often the checks occur.
* `host_validate` - Enable/disable validating "Host:" with original server IP. Valid values: `enable` `disable` .
* `ignore_conditional` - Enable/disable controlling the behavior of cache-control HTTP 1.1 header values. Valid values: `enable` `disable` .
* `ignore_ie_reload` - Enable/disable ignoring the PNC-interpretation of Internet Explorer's Accept: / header. Valid values: `enable` `disable` .
* `ignore_ims` - Enable/disable ignoring the if-modified-since (IMS) header. Valid values: `enable` `disable` .
* `ignore_pnc` - Enable/disable ignoring the pragma no-cache (PNC) header. Valid values: `enable` `disable` .
* `max_object_size` - Maximum cacheable object size in kB (1 - 2147483 kb (2GB). All objects that exceed this are delivered to the client but not stored in the web cache.
* `max_ttl` - Maximum time an object can stay in the web cache without checking to see if it has expired on the server (default = 7200 min (5 days); maximum = 5256000 min (10 years)).
* `min_ttl` - Minimum time an object can stay in the web cache without checking to see if it has expired on the server (default = 5 min; maximum = 5256000 (10 years)).
* `neg_resp_time` - Time in minutes to cache negative responses or errors (0 - 4294967295, default = 0  which means negative responses are not cached).
* `reval_pnc` - Enable/disable revalidation of pragma-no-cache (PNC) to address bandwidth concerns. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wanopt_webcache can be imported using:
```sh
terraform import fortios_wanopt_webcache.labelname {{mkey}}
```
