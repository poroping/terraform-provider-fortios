---
subcategory: "FortiGate Icap"
layout: "fortios"
page_title: "FortiOS: fortios_icap_profile"
description: |-
  Get information on a fortios Configure ICAP profiles.
---

# Data Source: fortios_icap_profile
Use this data source to get information on a fortios Configure ICAP profiles.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) ICAP profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `chunk_encap` - Enable/disable chunked encapsulation (default = disable).
* `extension_feature` - Enable/disable ICAP extension features.
* `icap_block_log` - Enable/disable UTM log when infection found (default = disable).
* `methods` - The allowed HTTP methods that will be sent to ICAP server for further processing.
* `name` - ICAP profile name.
* `preview` - Enable/disable preview of data to ICAP server.
* `preview_data_length` - Preview data length to be sent to ICAP server.
* `replacemsg_group` - Replacement message group.
* `request` - Enable/disable whether an HTTP request is passed to an ICAP server.
* `request_failure` - Action to take if the ICAP server cannot be contacted when processing an HTTP request.
* `request_path` - Path component of the ICAP URI that identifies the HTTP request processing service.
* `request_server` - ICAP server to use for an HTTP request.
* `respmod_default_action` - Default action to ICAP response modification (respmod) processing.
* `response` - Enable/disable whether an HTTP response is passed to an ICAP server.
* `response_failure` - Action to take if the ICAP server cannot be contacted when processing an HTTP response.
* `response_path` - Path component of the ICAP URI that identifies the HTTP response processing service.
* `response_req_hdr` - Enable/disable addition of req-hdr for ICAP response modification (respmod) processing.
* `response_server` - ICAP server to use for an HTTP response.
* `scan_progress_interval` - Scan progress interval value.
* `streaming_content_bypass` - Enable/disable bypassing of ICAP server for streaming content.
* `icap_headers` - Configure ICAP forwarded request headers.The structure of `icap_headers` block is documented below.

The `icap_headers` block contains:

* `base64_encoding` - Enable/disable use of base64 encoding of HTTP content.
* `content` - HTTP header content.
* `id` - HTTP forwarded header ID.
* `name` - HTTP forwarded header name.
* `respmod_forward_rules` - ICAP response mode forward rules.The structure of `respmod_forward_rules` block is documented below.

The `respmod_forward_rules` block contains:

* `action` - Action to be taken for ICAP server.
* `host` - Address object for the host.
* `name` - Address name.
* `header_group` - HTTP header group.The structure of `header_group` block is documented below.

The `header_group` block contains:

* `case_sensitivity` - Enable/disable case sensitivity when matching header.
* `header` - HTTP header regular expression.
* `header_name` - HTTP header.
* `id` - ID.
* `http_resp_status_code` - HTTP response status code.The structure of `http_resp_status_code` block is documented below.

The `http_resp_status_code` block contains:

* `code` - HTTP response status code.
