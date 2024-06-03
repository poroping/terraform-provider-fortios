---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logsyslogd_setting"
description: |-
  Get information on a fortios Global settings for remote syslog server.
---

# Data Source: fortios_logsyslogd_setting
Use this data source to get information on a fortios Global settings for remote syslog server.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `certificate` - Certificate used to communicate with Syslog server.
* `enc_algorithm` - Enable/disable reliable syslogging with TLS encryption.
* `facility` - Remote syslog facility.
* `format` - Log format.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `max_log_rate` - Syslog maximum log rate in MBps (0 = unlimited).
* `mode` - Remote syslog logging over UDP/Reliable TCP.
* `port` - Server listen port.
* `priority` - Set log transmission priority.
* `server` - Address of remote syslog server.
* `source_ip` - Source IP address of syslog.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `status` - Enable/disable remote syslog logging.
* `custom_field_name` - Custom field name for CEF format logging.The structure of `custom_field_name` block is documented below.

The `custom_field_name` block contains:

* `custom` - Field custom name [A-Za-z0-9_].
* `id` - Entry ID.
* `name` - Field name [A-Za-z0-9_].
