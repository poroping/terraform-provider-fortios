---
subcategory: "FortiGate Alertemail"
layout: "fortios"
page_title: "FortiOS: fortios_alertemail_setting"
description: |-
  Configure alert email settings.
---

## fortios_alertemail_setting
Configure alert email settings.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `fds_license_expiring_days` - Number of days to send alert email prior to FortiGuard license expiration (1 - 100 days, default = 100).
* `fds_license_expiring_warning` - Enable/disable FortiGuard license expiration warnings in alert email. Valid values: `enable` `disable` .
* `fds_update_logs` - Enable/disable FortiGuard update logs in alert email. Valid values: `enable` `disable` .
* `fips_cc_errors` - Enable/disable FIPS and Common Criteria error logs in alert email. Valid values: `enable` `disable` .
* `fsso_disconnect_logs` - Enable/disable logging of FSSO collector agent disconnect. Valid values: `enable` `disable` .
* `ha_logs` - Enable/disable HA logs in alert email. Valid values: `enable` `disable` .
* `ips_logs` - Enable/disable IPS logs in alert email. Valid values: `enable` `disable` .
* `ipsec_errors_logs` - Enable/disable IPsec error logs in alert email. Valid values: `enable` `disable` .
* `ppp_errors_logs` - Enable/disable PPP error logs in alert email. Valid values: `enable` `disable` .
* `admin_login_logs` - Enable/disable administrator login/logout logs in alert email. Valid values: `enable` `disable` .
* `alert_interval` - Alert alert interval in minutes.
* `amc_interface_bypass_mode` - Enable/disable Fortinet Advanced Mezzanine Card (AMC) interface bypass mode logs in alert email. Valid values: `enable` `disable` .
* `antivirus_logs` - Enable/disable antivirus logs in alert email. Valid values: `enable` `disable` .
* `configuration_changes_logs` - Enable/disable configuration change logs in alert email. Valid values: `enable` `disable` .
* `critical_interval` - Critical alert interval in minutes.
* `debug_interval` - Debug alert interval in minutes.
* `email_interval` - Interval between sending alert emails (1 - 99999 min, default = 5).
* `emergency_interval` - Emergency alert interval in minutes.
* `error_interval` - Error alert interval in minutes.
* `filter_mode` - How to filter log messages that are sent to alert emails. Valid values: `category` `threshold` .
* `firewall_authentication_failure_logs` - Enable/disable firewall authentication failure logs in alert email. Valid values: `enable` `disable` .
* `fortiguard_log_quota_warning` - Enable/disable FortiCloud log quota warnings in alert email. Valid values: `enable` `disable` .
* `information_interval` - Information alert interval in minutes.
* `local_disk_usage` - Disk usage percentage at which to send alert email (1 - 99 percent, default = 75).
* `log_disk_usage_warning` - Enable/disable disk usage warnings in alert email. Valid values: `enable` `disable` .
* `mailto1` - Email address to send alert email to (usually a system administrator) (max. 63 characters).
* `mailto2` - Optional second email address to send alert email to (max. 63 characters).
* `mailto3` - Optional third email address to send alert email to (max. 63 characters).
* `notification_interval` - Notification alert interval in minutes.
* `severity` - Lowest severity level to log. Valid values: `emergency` `alert` `critical` `error` `warning` `notification` `information` `debug` .
* `ssh_logs` - Enable/disable SSH logs in alert email. Valid values: `enable` `disable` .
* `sslvpn_authentication_errors_logs` - Enable/disable SSL-VPN authentication error logs in alert email. Valid values: `enable` `disable` .
* `username` - Name that appears in the From: field of alert emails (max. 63 characters).
* `violation_traffic_logs` - Enable/disable violation traffic logs in alert email. Valid values: `enable` `disable` .
* `warning_interval` - Warning alert interval in minutes.
* `webfilter_logs` - Enable/disable web filter logs in alert email. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_alertemail_setting can be imported using:
```sh
terraform import fortios_alertemail_setting.labelname {{mkey}}
```
