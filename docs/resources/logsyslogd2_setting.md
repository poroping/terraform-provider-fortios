---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logsyslogd2_setting"
description: |-
  Global settings for remote syslog server.
---

## fortios_logsyslogd2_setting
Global settings for remote syslog server.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `certificate` - Certificate used to communicate with Syslog server. This attribute must reference one of the following datasources: `certificate.local.name` .
* `enc_algorithm` - Enable/disable reliable syslogging with TLS encryption. Valid values: `high-medium` `high` `low` `disable` .
* `facility` - Remote syslog facility. Valid values: `kernel` `user` `mail` `daemon` `auth` `syslog` `lpr` `news` `uucp` `cron` `authpriv` `ftp` `ntp` `audit` `alert` `clock` `local0` `local1` `local2` `local3` `local4` `local5` `local6` `local7` .
* `format` - Log format. Valid values: `default` `csv` `cef` `rfc5424` .
* `interface` - Specify outgoing interface to reach server. This attribute must reference one of the following datasources: `system.interface.name` .
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `max_log_rate` - Syslog maximum log rate in MBps (0 = unlimited).
* `mode` - Remote syslog logging over UDP/Reliable TCP. Valid values: `udp` `legacy-reliable` `reliable` .
* `port` - Server listen port.
* `priority` - Set log transmission priority. Valid values: `default` `low` .
* `server` - Address of remote syslog server.
* `source_ip` - Source IP address of syslog.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default` `SSLv3` `TLSv1` `TLSv1-1` `TLSv1-2` .
* `status` - Enable/disable remote syslog logging. Valid values: `enable` `disable` .
* `custom_field_name` - Custom field name for CEF format logging. The structure of `custom_field_name` block is documented below.

The `custom_field_name` block contains:

* `custom` - Field custom name [A-Za-z0-9_].
* `id` - Entry ID.
* `name` - Field name [A-Za-z0-9_].

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_logsyslogd2_setting can be imported using:
```sh
terraform import fortios_logsyslogd2_setting.labelname {{mkey}}
```
