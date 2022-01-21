---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_profile"
description: |-
  Get information on a fortios Configure web proxy profiles.
---

# Data Source: fortios_webproxy_profile
Use this data source to get information on a fortios Configure web proxy profiles.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `header_client_ip` - Action to take on the HTTP client-IP header in forwarded requests: forwards (pass), adds, or removes the HTTP header.
* `header_front_end_https` - Action to take on the HTTP front-end-HTTPS header in forwarded requests: forwards (pass), adds, or removes the HTTP header.
* `header_via_request` - Action to take on the HTTP via header in forwarded requests: forwards (pass), adds, or removes the HTTP header.
* `header_via_response` - Action to take on the HTTP via header in forwarded responses: forwards (pass), adds, or removes the HTTP header.
* `header_x_authenticated_groups` - Action to take on the HTTP x-authenticated-groups header in forwarded requests: forwards (pass), adds, or removes the HTTP header.
* `header_x_authenticated_user` - Action to take on the HTTP x-authenticated-user header in forwarded requests: forwards (pass), adds, or removes the HTTP header.
* `header_x_forwarded_client_cert` - Action to take on the HTTP x-forwarded-client-cert header in forwarded requests: forwards (pass), adds, or removes the HTTP header.
* `header_x_forwarded_for` - Action to take on the HTTP x-forwarded-for header in forwarded requests: forwards (pass), adds, or removes the HTTP header.
* `log_header_change` - Enable/disable logging HTTP header changes.
* `name` - Profile name.
* `strip_encoding` - Enable/disable stripping unsupported encoding from the request header.
* `headers` - Configure HTTP forwarded requests headers.The structure of `headers` block is documented below.

The `headers` block contains:

* `action` - Action when the HTTP header is forwarded.
* `add_option` - Configure options to append content to existing HTTP header or add new HTTP header.
* `base64_encoding` - Enable/disable use of base64 encoding of HTTP content.
* `content` - HTTP header content.
* `id` - HTTP forwarded header id.
* `name` - HTTP forwarded header name.
* `protocol` - Configure protocol(s) to take add-option action on (HTTP, HTTPS, or both).
* `dstaddr` - Destination address and address group names.The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name.
* `dstaddr6` - Destination address and address group names (IPv6).The structure of `dstaddr6` block is documented below.

The `dstaddr6` block contains:

* `name` - Address name.
