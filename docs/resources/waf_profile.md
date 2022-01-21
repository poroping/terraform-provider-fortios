---
subcategory: "FortiGate Waf"
layout: "fortios"
page_title: "FortiOS: fortios_waf_profile"
description: |-
  Configure Web application firewall configuration.
---

## fortios_waf_profile
Configure Web application firewall configuration.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comment` - Comment.
* `extended_log` - Enable/disable extended logging. Valid values: `enable` `disable` .
* `external` - Disable/Enable external HTTP Inspection. Valid values: `disable` `enable` .
* `name` - WAF Profile name.
* `address_list` - Address block and allow lists. The structure of `address_list` block is documented below.

The `address_list` block contains:

* `blocked_log` - Enable/disable logging on blocked addresses. Valid values: `enable` `disable` .
* `severity` - Severity. Valid values: `high` `medium` `low` .
* `status` - Status. Valid values: `enable` `disable` .
* `blocked_address` - Blocked address. The structure of `blocked_address` block is documented below.

The `blocked_address` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `trusted_address` - Trusted address. The structure of `trusted_address` block is documented below.

The `trusted_address` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `constraint` - WAF HTTP protocol restrictions. The structure of `constraint` block is documented below.

The `constraint` block contains:

* `content_length` - HTTP content length in request. The structure of `content_length` block is documented below.

The `content_length` block contains:

* `action` - Action. Valid values: `allow` `block` .
* `length` - Length of HTTP content in bytes (0 to 2147483647).
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `severity` - Severity. Valid values: `high` `medium` `low` .
* `status` - Enable/disable the constraint. Valid values: `enable` `disable` .
* `exception` - HTTP constraint exception. The structure of `exception` block is documented below.

The `exception` block contains:

* `address` - Host address. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `content_length` - HTTP content length in request. Valid values: `enable` `disable` .
* `header_length` - HTTP header length in request. Valid values: `enable` `disable` .
* `hostname` - Enable/disable hostname check. Valid values: `enable` `disable` .
* `id` - Exception ID.
* `line_length` - HTTP line length in request. Valid values: `enable` `disable` .
* `malformed` - Enable/disable malformed HTTP request check. Valid values: `enable` `disable` .
* `max_cookie` - Maximum number of cookies in HTTP request. Valid values: `enable` `disable` .
* `max_header_line` - Maximum number of HTTP header line. Valid values: `enable` `disable` .
* `max_range_segment` - Maximum number of range segments in HTTP range line. Valid values: `enable` `disable` .
* `max_url_param` - Maximum number of parameters in URL. Valid values: `enable` `disable` .
* `method` - Enable/disable HTTP method check. Valid values: `enable` `disable` .
* `param_length` - Maximum length of parameter in URL, HTTP POST request or HTTP body. Valid values: `enable` `disable` .
* `pattern` - URL pattern.
* `regex` - Enable/disable regular expression based pattern match. Valid values: `enable` `disable` .
* `url_param_length` - Maximum length of parameter in URL. Valid values: `enable` `disable` .
* `version` - Enable/disable HTTP version check. Valid values: `enable` `disable` .
* `header_length` - HTTP header length in request. The structure of `header_length` block is documented below.

The `header_length` block contains:

* `action` - Action. Valid values: `allow` `block` .
* `length` - Length of HTTP header in bytes (0 to 2147483647).
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `severity` - Severity. Valid values: `high` `medium` `low` .
* `status` - Enable/disable the constraint. Valid values: `enable` `disable` .
* `hostname` - Enable/disable hostname check. The structure of `hostname` block is documented below.

The `hostname` block contains:

* `action` - Action. Valid values: `allow` `block` .
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `severity` - Severity. Valid values: `high` `medium` `low` .
* `status` - Enable/disable the constraint. Valid values: `enable` `disable` .
* `line_length` - HTTP line length in request. The structure of `line_length` block is documented below.

The `line_length` block contains:

* `action` - Action. Valid values: `allow` `block` .
* `length` - Length of HTTP line in bytes (0 to 2147483647).
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `severity` - Severity. Valid values: `high` `medium` `low` .
* `status` - Enable/disable the constraint. Valid values: `enable` `disable` .
* `malformed` - Enable/disable malformed HTTP request check. The structure of `malformed` block is documented below.

The `malformed` block contains:

* `action` - Action. Valid values: `allow` `block` .
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `severity` - Severity. Valid values: `high` `medium` `low` .
* `status` - Enable/disable the constraint. Valid values: `enable` `disable` .
* `max_cookie` - Maximum number of cookies in HTTP request. The structure of `max_cookie` block is documented below.

The `max_cookie` block contains:

* `action` - Action. Valid values: `allow` `block` .
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `max_cookie` - Maximum number of cookies in HTTP request (0 to 2147483647).
* `severity` - Severity. Valid values: `high` `medium` `low` .
* `status` - Enable/disable the constraint. Valid values: `enable` `disable` .
* `max_header_line` - Maximum number of HTTP header line. The structure of `max_header_line` block is documented below.

The `max_header_line` block contains:

* `action` - Action. Valid values: `allow` `block` .
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `max_header_line` - Maximum number HTTP header lines (0 to 2147483647).
* `severity` - Severity. Valid values: `high` `medium` `low` .
* `status` - Enable/disable the constraint. Valid values: `enable` `disable` .
* `max_range_segment` - Maximum number of range segments in HTTP range line. The structure of `max_range_segment` block is documented below.

The `max_range_segment` block contains:

* `action` - Action. Valid values: `allow` `block` .
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `max_range_segment` - Maximum number of range segments in HTTP range line (0 to 2147483647).
* `severity` - Severity. Valid values: `high` `medium` `low` .
* `status` - Enable/disable the constraint. Valid values: `enable` `disable` .
* `max_url_param` - Maximum number of parameters in URL. The structure of `max_url_param` block is documented below.

The `max_url_param` block contains:

* `action` - Action. Valid values: `allow` `block` .
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `max_url_param` - Maximum number of parameters in URL (0 to 2147483647).
* `severity` - Severity. Valid values: `high` `medium` `low` .
* `status` - Enable/disable the constraint. Valid values: `enable` `disable` .
* `method` - Enable/disable HTTP method check. The structure of `method` block is documented below.

The `method` block contains:

* `action` - Action. Valid values: `allow` `block` .
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `severity` - Severity. Valid values: `high` `medium` `low` .
* `status` - Enable/disable the constraint. Valid values: `enable` `disable` .
* `param_length` - Maximum length of parameter in URL, HTTP POST request or HTTP body. The structure of `param_length` block is documented below.

The `param_length` block contains:

* `action` - Action. Valid values: `allow` `block` .
* `length` - Maximum length of parameter in URL, HTTP POST request or HTTP body in bytes (0 to 2147483647).
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `severity` - Severity. Valid values: `high` `medium` `low` .
* `status` - Enable/disable the constraint. Valid values: `enable` `disable` .
* `url_param_length` - Maximum length of parameter in URL. The structure of `url_param_length` block is documented below.

The `url_param_length` block contains:

* `action` - Action. Valid values: `allow` `block` .
* `length` - Maximum length of URL parameter in bytes (0 to 2147483647).
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `severity` - Severity. Valid values: `high` `medium` `low` .
* `status` - Enable/disable the constraint. Valid values: `enable` `disable` .
* `version` - Enable/disable HTTP version check. The structure of `version` block is documented below.

The `version` block contains:

* `action` - Action. Valid values: `allow` `block` .
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `severity` - Severity. Valid values: `high` `medium` `low` .
* `status` - Enable/disable the constraint. Valid values: `enable` `disable` .
* `method` - Method restriction. The structure of `method` block is documented below.

The `method` block contains:

* `default_allowed_methods` - Methods. Valid values: `get` `post` `put` `head` `connect` `trace` `options` `delete` `others` .
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `severity` - Severity. Valid values: `high` `medium` `low` .
* `status` - Status. Valid values: `enable` `disable` .
* `method_policy` - HTTP method policy. The structure of `method_policy` block is documented below.

The `method_policy` block contains:

* `address` - Host address. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `allowed_methods` - Allowed Methods. Valid values: `get` `post` `put` `head` `connect` `trace` `options` `delete` `others` .
* `id` - HTTP method policy ID.
* `pattern` - URL pattern.
* `regex` - Enable/disable regular expression based pattern match. Valid values: `enable` `disable` .
* `signature` - WAF signatures. The structure of `signature` block is documented below.

The `signature` block contains:

* `credit_card_detection_threshold` - The minimum number of Credit cards to detect violation.
* `custom_signature` - Custom signature. The structure of `custom_signature` block is documented below.

The `custom_signature` block contains:

* `action` - Action. Valid values: `allow` `block` `erase` .
* `case_sensitivity` - Case sensitivity in pattern. Valid values: `disable` `enable` .
* `direction` - Traffic direction. Valid values: `request` `response` .
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `name` - Signature name.
* `pattern` - Match pattern.
* `severity` - Severity. Valid values: `high` `medium` `low` .
* `status` - Status. Valid values: `enable` `disable` .
* `target` - Match HTTP target. Valid values: `arg` `arg-name` `req-body` `req-cookie` `req-cookie-name` `req-filename` `req-header` `req-header-name` `req-raw-uri` `req-uri` `resp-body` `resp-hdr` `resp-status` .
* `disabled_signature` - Disabled signatures The structure of `disabled_signature` block is documented below.

The `disabled_signature` block contains:

* `id` - Signature ID. This attribute must reference one of the following datasources: `waf.signature.id` .
* `disabled_sub_class` - Disabled signature subclasses. The structure of `disabled_sub_class` block is documented below.

The `disabled_sub_class` block contains:

* `id` - Signature subclass ID. This attribute must reference one of the following datasources: `waf.sub-class.id` .
* `main_class` - Main signature class. The structure of `main_class` block is documented below.

The `main_class` block contains:

* `action` - Action. Valid values: `allow` `block` `erase` .
* `id` - Main signature class ID. This attribute must reference one of the following datasources: `waf.main-class.id` .
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `severity` - Severity. Valid values: `high` `medium` `low` .
* `status` - Status. Valid values: `enable` `disable` .
* `url_access` - URL access list The structure of `url_access` block is documented below.

The `url_access` block contains:

* `action` - Action. Valid values: `bypass` `permit` `block` .
* `address` - Host address. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `id` - URL access ID.
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `severity` - Severity. Valid values: `high` `medium` `low` .
* `access_pattern` - URL access pattern. The structure of `access_pattern` block is documented below.

The `access_pattern` block contains:

* `id` - URL access pattern ID.
* `negate` - Enable/disable match negation. Valid values: `enable` `disable` .
* `pattern` - URL pattern.
* `regex` - Enable/disable regular expression based pattern match. Valid values: `enable` `disable` .
* `srcaddr` - Source address. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_waf_profile can be imported using:
```sh
terraform import fortios_waf_profile.labelname {{mkey}}
```
