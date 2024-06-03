---
subcategory: "FortiGate Webfilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_fortiguard"
description: |-
  Get information on a fortios Configure FortiGuard Web Filter service.
---

# Data Source: fortios_webfilter_fortiguard
Use this data source to get information on a fortios Configure FortiGuard Web Filter service.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `cache_mem_percent` - Maximum percentage of available memory allocated to caching (1 - 15).
* `cache_mode` - Cache entry expiration mode.
* `cache_prefix_match` - Enable/disable prefix matching in the cache.
* `close_ports` - Close ports used for HTTP/HTTPS override authentication and disable user overrides.
* `embed_image` - Enable/disable embedding images into replacement messages (default = enable).
* `ovrd_auth_https` - Enable/disable use of HTTPS for override authentication.
* `ovrd_auth_port_http` - Port to use for FortiGuard Web Filter HTTP override authentication.
* `ovrd_auth_port_https` - Port to use for FortiGuard Web Filter HTTPS override authentication in proxy mode.
* `ovrd_auth_port_https_flow` - Port to use for FortiGuard Web Filter HTTPS override authentication in flow mode.
* `ovrd_auth_port_warning` - Port to use for FortiGuard Web Filter Warning override authentication.
* `request_packet_size_limit` - Limit size of URL request packets sent to FortiGuard server (0 for default).
* `warn_auth_https` - Enable/disable use of HTTPS for warning and authentication.
