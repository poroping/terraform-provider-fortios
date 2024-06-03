---
subcategory: "FortiGate Alertemail"
layout: "fortios"
page_title: "FortiOS: fortios_alertemail_setting"
description: |-
  Get information on a fortios Configure alert email settings.
---

# Data Source: fortios_alertemail_setting
Use this data source to get information on a fortios Configure alert email settings.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `fds_license_expiring_days` - Number of days to send alert email prior to FortiGuard license expiration (1 - 100 days, default = 15).
* `fds_license_expiring_warning` - Enable/disable FortiGuard license expiration warnings in alert email.
* `fds_update_logs` - Enable/disable FortiGuard update logs in alert email.
* `fips_cc_errors` - Enable/disable FIPS and Common Criteria error logs in alert email.
* `fsso_disconnect_logs` - Enable/disable logging of FSSO collector agent disconnect.
* `ha_logs` - Enable/disable HA logs in alert email.
* `ips_logs` - Enable/disable IPS logs in alert email.
* `ipsec_errors_logs` - Enable/disable IPsec error logs in alert email.
* `ppp_errors_logs` - Enable/disable PPP error logs in alert email.
* `admin_login_logs` - Enable/disable administrator login/logout logs in alert email.
* `alert_interval` - Alert alert interval in minutes.
* `amc_interface_bypass_mode` - Enable/disable Fortinet Advanced Mezzanine Card (AMC) interface bypass mode logs in alert email.
* `antivirus_logs` - Enable/disable antivirus logs in alert email.
* `configuration_changes_logs` - Enable/disable configuration change logs in alert email.
* `critical_interval` - Critical alert interval in minutes.
* `debug_interval` - Debug alert interval in minutes.
* `email_interval` - Interval between sending alert emails (1 - 99999 min, default = 5).
* `emergency_interval` - Emergency alert interval in minutes.
* `error_interval` - Error alert interval in minutes.
* `filter_mode` - How to filter log messages that are sent to alert emails.
* `firewall_authentication_failure_logs` - Enable/disable firewall authentication failure logs in alert email.
* `fortiguard_log_quota_warning` - Enable/disable FortiCloud log quota warnings in alert email.
* `information_interval` - Information alert interval in minutes.
* `local_disk_usage` - Disk usage percentage at which to send alert email (1 - 99 percent, default = 75).
* `log_disk_usage_warning` - Enable/disable disk usage warnings in alert email.
* `mailto1` - Email address to send alert email to (usually a system administrator) (max. 63 characters).
* `mailto2` - Optional second email address to send alert email to (max. 63 characters).
* `mailto3` - Optional third email address to send alert email to (max. 63 characters).
* `notification_interval` - Notification alert interval in minutes.
* `severity` - Lowest severity level to log.
* `ssh_logs` - Enable/disable SSH logs in alert email.
* `sslvpn_authentication_errors_logs` - Enable/disable SSL-VPN authentication error logs in alert email.
* `username` - Name that appears in the From: field of alert emails (max. 63 characters).
* `violation_traffic_logs` - Enable/disable violation traffic logs in alert email.
* `warning_interval` - Warning alert interval in minutes.
* `webfilter_logs` - Enable/disable web filter logs in alert email.
