---
subcategory: "FortiGate Icap"
layout: "fortios"
page_title: "FortiOS: fortios_icap_profile"
description: |-
  Configure ICAP profiles.
---

## fortios_icap_profile
Configure ICAP profiles.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `chunk_encap` - Enable/disable chunked encapsulation (default = disable). Valid values: `disable` `enable` .
* `extension_feature` - Enable/disable ICAP extension features. Valid values: `scan-progress` .
* `icap_block_log` - Enable/disable UTM log when infection found (default = disable). Valid values: `disable` `enable` .
* `methods` - The allowed HTTP methods that will be sent to ICAP server for further processing. Valid values: `delete` `get` `head` `options` `post` `put` `trace` `other` .
* `name` - ICAP profile name.
* `preview` - Enable/disable preview of data to ICAP server. Valid values: `disable` `enable` .
* `preview_data_length` - Preview data length to be sent to ICAP server.
* `replacemsg_group` - Replacement message group. This attribute must reference one of the following datasources: `system.replacemsg-group.name` .
* `request` - Enable/disable whether an HTTP request is passed to an ICAP server. Valid values: `disable` `enable` .
* `request_failure` - Action to take if the ICAP server cannot be contacted when processing an HTTP request. Valid values: `error` `bypass` .
* `request_path` - Path component of the ICAP URI that identifies the HTTP request processing service.
* `request_server` - ICAP server to use for an HTTP request. This attribute must reference one of the following datasources: `icap.server.name` .
* `respmod_default_action` - Default action to ICAP response modification (respmod) processing. Valid values: `forward` `bypass` .
* `response` - Enable/disable whether an HTTP response is passed to an ICAP server. Valid values: `disable` `enable` .
* `response_failure` - Action to take if the ICAP server cannot be contacted when processing an HTTP response. Valid values: `error` `bypass` .
* `response_path` - Path component of the ICAP URI that identifies the HTTP response processing service.
* `response_req_hdr` - Enable/disable addition of req-hdr for ICAP response modification (respmod) processing. Valid values: `disable` `enable` .
* `response_server` - ICAP server to use for an HTTP response. This attribute must reference one of the following datasources: `icap.server.name` .
* `scan_progress_interval` - Scan progress interval value.
* `streaming_content_bypass` - Enable/disable bypassing of ICAP server for streaming content. Valid values: `disable` `enable` .
* `icap_headers` - Configure ICAP forwarded request headers. The structure of `icap_headers` block is documented below.

The `icap_headers` block contains:

* `base64_encoding` - Enable/disable use of base64 encoding of HTTP content. Valid values: `disable` `enable` .
* `content` - HTTP header content.
* `id` - HTTP forwarded header ID.
* `name` - HTTP forwarded header name.
* `respmod_forward_rules` - ICAP response mode forward rules. The structure of `respmod_forward_rules` block is documented below.

The `respmod_forward_rules` block contains:

* `action` - Action to be taken for ICAP server. Valid values: `forward` `bypass` .
* `host` - Address object for the host. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` `firewall.proxy-address.name` .
* `name` - Address name.
* `header_group` - HTTP header group. The structure of `header_group` block is documented below.

The `header_group` block contains:

* `case_sensitivity` - Enable/disable case sensitivity when matching header. Valid values: `disable` `enable` .
* `header` - HTTP header regular expression.
* `header_name` - HTTP header.
* `id` - ID.
* `http_resp_status_code` - HTTP response status code. The structure of `http_resp_status_code` block is documented below.

The `http_resp_status_code` block contains:

* `code` - HTTP response status code.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_icap_profile can be imported using:
```sh
terraform import fortios_icap_profile.labelname {{mkey}}
```
