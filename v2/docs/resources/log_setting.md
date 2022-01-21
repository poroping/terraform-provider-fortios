---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_log_setting"
description: |-
  Configure general log settings.
---

## fortios_log_setting
Configure general log settings.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `anonymization_hash` - User name anonymization hash salt.
* `brief_traffic_format` - Enable/disable brief format traffic logging. Valid values: `enable` `disable` .
* `daemon_log` - Enable/disable daemon logging. Valid values: `enable` `disable` .
* `expolicy_implicit_log` - Enable/disable explicit proxy firewall implicit policy logging. Valid values: `enable` `disable` .
* `faz_override` - Enable/disable override FortiAnalyzer settings. Valid values: `enable` `disable` .
* `fortiview_weekly_data` - Enable/disable FortiView weekly data. Valid values: `enable` `disable` .
* `fwpolicy_implicit_log` - Enable/disable implicit firewall policy logging. Valid values: `enable` `disable` .
* `fwpolicy6_implicit_log` - Enable/disable implicit firewall policy6 logging. Valid values: `enable` `disable` .
* `local_in_allow` - Enable/disable local-in-allow logging. Valid values: `enable` `disable` .
* `local_in_deny_broadcast` - Enable/disable local-in-deny-broadcast logging. Valid values: `enable` `disable` .
* `local_in_deny_unicast` - Enable/disable local-in-deny-unicast logging. Valid values: `enable` `disable` .
* `local_out` - Enable/disable local-out logging. Valid values: `enable` `disable` .
* `log_invalid_packet` - Enable/disable invalid packet traffic logging. Valid values: `enable` `disable` .
* `log_policy_comment` - Enable/disable inserting policy comments into traffic logs. Valid values: `enable` `disable` .
* `log_policy_name` - Enable/disable inserting policy name into traffic logs. Valid values: `enable` `disable` .
* `log_user_in_upper` - Enable/disable logs with user-in-upper. Valid values: `enable` `disable` .
* `neighbor_event` - Enable/disable neighbor event logging. Valid values: `enable` `disable` .
* `resolve_ip` - Enable/disable adding resolved domain names to traffic logs if possible. Valid values: `enable` `disable` .
* `resolve_port` - Enable/disable adding resolved service names to traffic logs. Valid values: `enable` `disable` .
* `syslog_override` - Enable/disable override Syslog settings. Valid values: `enable` `disable` .
* `user_anonymize` - Enable/disable anonymizing user names in log messages. Valid values: `enable` `disable` .
* `custom_log_fields` - Custom fields to append to all log messages. The structure of `custom_log_fields` block is documented below.

The `custom_log_fields` block contains:

* `field_id` - Custom log field. This attribute must reference one of the following datasources: `log.custom-field.id` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_log_setting can be imported using:
```sh
terraform import fortios_log_setting.labelname {{mkey}}
```
