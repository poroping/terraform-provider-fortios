---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_log_setting"
description: |-
  Get information on a fortios Configure general log settings.
---

# Data Source: fortios_log_setting
Use this data source to get information on a fortios Configure general log settings.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `anonymization_hash` - User name anonymization hash salt.
* `brief_traffic_format` - Enable/disable brief format traffic logging.
* `daemon_log` - Enable/disable daemon logging.
* `expolicy_implicit_log` - Enable/disable explicit proxy firewall implicit policy logging.
* `faz_override` - Enable/disable override FortiAnalyzer settings.
* `fortiview_weekly_data` - Enable/disable FortiView weekly data.
* `fwpolicy_implicit_log` - Enable/disable implicit firewall policy logging.
* `fwpolicy6_implicit_log` - Enable/disable implicit firewall policy6 logging.
* `local_in_allow` - Enable/disable local-in-allow logging.
* `local_in_deny_broadcast` - Enable/disable local-in-deny-broadcast logging.
* `local_in_deny_unicast` - Enable/disable local-in-deny-unicast logging.
* `local_out` - Enable/disable local-out logging.
* `local_out_ioc_detection` - Enable/disable local-out traffic IoC detection. Requires local-out to be enabled.
* `log_invalid_packet` - Enable/disable invalid packet traffic logging.
* `log_policy_comment` - Enable/disable inserting policy comments into traffic logs.
* `log_policy_name` - Enable/disable inserting policy name into traffic logs.
* `log_user_in_upper` - Enable/disable logs with user-in-upper.
* `neighbor_event` - Enable/disable neighbor event logging.
* `resolve_ip` - Enable/disable adding resolved domain names to traffic logs if possible.
* `resolve_port` - Enable/disable adding resolved service names to traffic logs.
* `rest_api_get` - Enable/disable REST API GET request logging.
* `rest_api_set` - Enable/disable REST API POST/PUT/DELETE request logging.
* `syslog_override` - Enable/disable override Syslog settings.
* `user_anonymize` - Enable/disable anonymizing user names in log messages.
* `custom_log_fields` - Custom fields to append to all log messages.The structure of `custom_log_fields` block is documented below.

The `custom_log_fields` block contains:

* `field_id` - Custom log field.
