---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_profile"
description: |-
  Configure web proxy profiles.
---

## fortios_webproxy_profile
Configure web proxy profiles.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `header_client_ip` - Action to take on the HTTP client-IP header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass` `add` `remove` .
* `header_front_end_https` - Action to take on the HTTP front-end-HTTPS header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass` `add` `remove` .
* `header_via_request` - Action to take on the HTTP via header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass` `add` `remove` .
* `header_via_response` - Action to take on the HTTP via header in forwarded responses: forwards (pass), adds, or removes the HTTP header. Valid values: `pass` `add` `remove` .
* `header_x_authenticated_groups` - Action to take on the HTTP x-authenticated-groups header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass` `add` `remove` .
* `header_x_authenticated_user` - Action to take on the HTTP x-authenticated-user header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass` `add` `remove` .
* `header_x_forwarded_client_cert` - Action to take on the HTTP x-forwarded-client-cert header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass` `add` `remove` .
* `header_x_forwarded_for` - Action to take on the HTTP x-forwarded-for header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass` `add` `remove` .
* `log_header_change` - Enable/disable logging HTTP header changes. Valid values: `enable` `disable` .
* `name` - Profile name.
* `strip_encoding` - Enable/disable stripping unsupported encoding from the request header. Valid values: `enable` `disable` .
* `headers` - Configure HTTP forwarded requests headers. The structure of `headers` block is documented below.

The `headers` block contains:

* `action` - Action when the HTTP header is forwarded. Valid values: `add-to-request` `add-to-response` `remove-from-request` `remove-from-response` .
* `add_option` - Configure options to append content to existing HTTP header or add new HTTP header. Valid values: `append` `new-on-not-found` `new` .
* `base64_encoding` - Enable/disable use of base64 encoding of HTTP content. Valid values: `disable` `enable` .
* `content` - HTTP header content.
* `id` - HTTP forwarded header id.
* `name` - HTTP forwarded header name.
* `protocol` - Configure protocol(s) to take add-option action on (HTTP, HTTPS, or both). Valid values: `https` `http` .
* `dstaddr` - Destination address and address group names. The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `dstaddr6` - Destination address and address group names (IPv6). The structure of `dstaddr6` block is documented below.

The `dstaddr6` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_webproxy_profile can be imported using:
```sh
terraform import fortios_webproxy_profile.labelname {{mkey}}
```
