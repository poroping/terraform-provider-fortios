---
subcategory: "FortiGate Wanopt"
layout: "fortios"
page_title: "FortiOS: fortios_wanopt_webcache"
description: |-
  Get information on a fortios Configure global Web cache settings.
---

# Data Source: fortios_wanopt_webcache
Use this data source to get information on a fortios Configure global Web cache settings.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `always_revalidate` - Enable/disable revalidation of requested cached objects, which have content on the server, before serving it to the client.
* `cache_by_default` - Enable/disable caching content that lacks explicit caching policies from the server.
* `cache_cookie` - Enable/disable caching cookies. Since cookies contain information for or about individual users, they not usually cached.
* `cache_expired` - Enable/disable caching type-1 objects that are already expired on arrival.
* `default_ttl` - Default object expiry time (default = 1440 min (1 day); maximum = 5256000 min (10 years)). This only applies to those objects that do not have an expiry time set by the web server.
* `external` - Enable/disable external Web caching.
* `fresh_factor` - Frequency that the server is checked to see if any objects have expired (1 - 100, default = 100). The higher the fresh factor, the less often the checks occur.
* `host_validate` - Enable/disable validating "Host:" with original server IP.
* `ignore_conditional` - Enable/disable controlling the behavior of cache-control HTTP 1.1 header values.
* `ignore_ie_reload` - Enable/disable ignoring the PNC-interpretation of Internet Explorer's Accept: / header.
* `ignore_ims` - Enable/disable ignoring the if-modified-since (IMS) header.
* `ignore_pnc` - Enable/disable ignoring the pragma no-cache (PNC) header.
* `max_object_size` - Maximum cacheable object size in kB (1 - 2147483 kb (2GB). All objects that exceed this are delivered to the client but not stored in the web cache.
* `max_ttl` - Maximum time an object can stay in the web cache without checking to see if it has expired on the server (default = 7200 min (5 days); maximum = 5256000 min (10 years)).
* `min_ttl` - Minimum time an object can stay in the web cache without checking to see if it has expired on the server (default = 5 min; maximum = 5256000 (10 years)).
* `neg_resp_time` - Time in minutes to cache negative responses or errors (0 - 4294967295, default = 0  which means negative responses are not cached).
* `reval_pnc` - Enable/disable revalidation of pragma-no-cache (PNC) to address bandwidth concerns.
