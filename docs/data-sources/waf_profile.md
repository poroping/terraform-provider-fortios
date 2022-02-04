---
subcategory: "FortiGate Waf"
layout: "fortios"
page_title: "FortiOS: fortios_waf_profile"
description: |-
  Get information on a fortios Configure Web application firewall configuration.
---

# Data Source: fortios_waf_profile
Use this data source to get information on a fortios Configure Web application firewall configuration.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) WAF Profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `extended_log` - Enable/disable extended logging.
* `external` - Disable/Enable external HTTP Inspection.
* `name` - WAF Profile name.
* `address_list` - Address block and allow lists.The structure of `address_list` block is documented below.

The `address_list` block contains:

* `blocked_log` - Enable/disable logging on blocked addresses.
* `severity` - Severity.
* `status` - Status.
* `blocked_address` - Blocked address.The structure of `blocked_address` block is documented below.

The `blocked_address` block contains:

* `name` - Address name.
* `trusted_address` - Trusted address.The structure of `trusted_address` block is documented below.

The `trusted_address` block contains:

* `name` - Address name.
* `constraint` - WAF HTTP protocol restrictions.The structure of `constraint` block is documented below.

The `constraint` block contains:

* `content_length` - HTTP content length in request.The structure of `content_length` block is documented below.

The `content_length` block contains:

* `action` - Action.
* `length` - Length of HTTP content in bytes (0 to 2147483647).
* `log` - Enable/disable logging.
* `severity` - Severity.
* `status` - Enable/disable the constraint.
* `exception` - HTTP constraint exception.The structure of `exception` block is documented below.

The `exception` block contains:

* `address` - Host address.
* `content_length` - HTTP content length in request.
* `header_length` - HTTP header length in request.
* `hostname` - Enable/disable hostname check.
* `id` - Exception ID.
* `line_length` - HTTP line length in request.
* `malformed` - Enable/disable malformed HTTP request check.
* `max_cookie` - Maximum number of cookies in HTTP request.
* `max_header_line` - Maximum number of HTTP header line.
* `max_range_segment` - Maximum number of range segments in HTTP range line.
* `max_url_param` - Maximum number of parameters in URL.
* `method` - Enable/disable HTTP method check.
* `param_length` - Maximum length of parameter in URL, HTTP POST request or HTTP body.
* `pattern` - URL pattern.
* `regex` - Enable/disable regular expression based pattern match.
* `url_param_length` - Maximum length of parameter in URL.
* `version` - Enable/disable HTTP version check.
* `header_length` - HTTP header length in request.The structure of `header_length` block is documented below.

The `header_length` block contains:

* `action` - Action.
* `length` - Length of HTTP header in bytes (0 to 2147483647).
* `log` - Enable/disable logging.
* `severity` - Severity.
* `status` - Enable/disable the constraint.
* `hostname` - Enable/disable hostname check.The structure of `hostname` block is documented below.

The `hostname` block contains:

* `action` - Action.
* `log` - Enable/disable logging.
* `severity` - Severity.
* `status` - Enable/disable the constraint.
* `line_length` - HTTP line length in request.The structure of `line_length` block is documented below.

The `line_length` block contains:

* `action` - Action.
* `length` - Length of HTTP line in bytes (0 to 2147483647).
* `log` - Enable/disable logging.
* `severity` - Severity.
* `status` - Enable/disable the constraint.
* `malformed` - Enable/disable malformed HTTP request check.The structure of `malformed` block is documented below.

The `malformed` block contains:

* `action` - Action.
* `log` - Enable/disable logging.
* `severity` - Severity.
* `status` - Enable/disable the constraint.
* `max_cookie` - Maximum number of cookies in HTTP request.The structure of `max_cookie` block is documented below.

The `max_cookie` block contains:

* `action` - Action.
* `log` - Enable/disable logging.
* `max_cookie` - Maximum number of cookies in HTTP request (0 to 2147483647).
* `severity` - Severity.
* `status` - Enable/disable the constraint.
* `max_header_line` - Maximum number of HTTP header line.The structure of `max_header_line` block is documented below.

The `max_header_line` block contains:

* `action` - Action.
* `log` - Enable/disable logging.
* `max_header_line` - Maximum number HTTP header lines (0 to 2147483647).
* `severity` - Severity.
* `status` - Enable/disable the constraint.
* `max_range_segment` - Maximum number of range segments in HTTP range line.The structure of `max_range_segment` block is documented below.

The `max_range_segment` block contains:

* `action` - Action.
* `log` - Enable/disable logging.
* `max_range_segment` - Maximum number of range segments in HTTP range line (0 to 2147483647).
* `severity` - Severity.
* `status` - Enable/disable the constraint.
* `max_url_param` - Maximum number of parameters in URL.The structure of `max_url_param` block is documented below.

The `max_url_param` block contains:

* `action` - Action.
* `log` - Enable/disable logging.
* `max_url_param` - Maximum number of parameters in URL (0 to 2147483647).
* `severity` - Severity.
* `status` - Enable/disable the constraint.
* `method` - Enable/disable HTTP method check.The structure of `method` block is documented below.

The `method` block contains:

* `action` - Action.
* `log` - Enable/disable logging.
* `severity` - Severity.
* `status` - Enable/disable the constraint.
* `param_length` - Maximum length of parameter in URL, HTTP POST request or HTTP body.The structure of `param_length` block is documented below.

The `param_length` block contains:

* `action` - Action.
* `length` - Maximum length of parameter in URL, HTTP POST request or HTTP body in bytes (0 to 2147483647).
* `log` - Enable/disable logging.
* `severity` - Severity.
* `status` - Enable/disable the constraint.
* `url_param_length` - Maximum length of parameter in URL.The structure of `url_param_length` block is documented below.

The `url_param_length` block contains:

* `action` - Action.
* `length` - Maximum length of URL parameter in bytes (0 to 2147483647).
* `log` - Enable/disable logging.
* `severity` - Severity.
* `status` - Enable/disable the constraint.
* `version` - Enable/disable HTTP version check.The structure of `version` block is documented below.

The `version` block contains:

* `action` - Action.
* `log` - Enable/disable logging.
* `severity` - Severity.
* `status` - Enable/disable the constraint.
* `method` - Method restriction.The structure of `method` block is documented below.

The `method` block contains:

* `default_allowed_methods` - Methods.
* `log` - Enable/disable logging.
* `severity` - Severity.
* `status` - Status.
* `method_policy` - HTTP method policy.The structure of `method_policy` block is documented below.

The `method_policy` block contains:

* `address` - Host address.
* `allowed_methods` - Allowed Methods.
* `id` - HTTP method policy ID.
* `pattern` - URL pattern.
* `regex` - Enable/disable regular expression based pattern match.
* `signature` - WAF signatures.The structure of `signature` block is documented below.

The `signature` block contains:

* `credit_card_detection_threshold` - The minimum number of Credit cards to detect violation.
* `custom_signature` - Custom signature.The structure of `custom_signature` block is documented below.

The `custom_signature` block contains:

* `action` - Action.
* `case_sensitivity` - Case sensitivity in pattern.
* `direction` - Traffic direction.
* `log` - Enable/disable logging.
* `name` - Signature name.
* `pattern` - Match pattern.
* `severity` - Severity.
* `status` - Status.
* `target` - Match HTTP target.
* `disabled_signature` - Disabled signatures.The structure of `disabled_signature` block is documented below.

The `disabled_signature` block contains:

* `id` - Signature ID.
* `disabled_sub_class` - Disabled signature subclasses.The structure of `disabled_sub_class` block is documented below.

The `disabled_sub_class` block contains:

* `id` - Signature subclass ID.
* `main_class` - Main signature class.The structure of `main_class` block is documented below.

The `main_class` block contains:

* `action` - Action.
* `id` - Main signature class ID.
* `log` - Enable/disable logging.
* `severity` - Severity.
* `status` - Status.
* `url_access` - URL access list.The structure of `url_access` block is documented below.

The `url_access` block contains:

* `action` - Action.
* `address` - Host address.
* `id` - URL access ID.
* `log` - Enable/disable logging.
* `severity` - Severity.
* `access_pattern` - URL access pattern.The structure of `access_pattern` block is documented below.

The `access_pattern` block contains:

* `id` - URL access pattern ID.
* `negate` - Enable/disable match negation.
* `pattern` - URL pattern.
* `regex` - Enable/disable regular expression based pattern match.
* `srcaddr` - Source address.
