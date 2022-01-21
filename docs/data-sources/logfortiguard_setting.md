---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logfortiguard_setting"
description: |-
  Get information on a fortios Configure logging to FortiCloud.
---

# Data Source: fortios_logfortiguard_setting
Use this data source to get information on a fortios Configure logging to FortiCloud.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `access_config` - Enable/disable FortiCloud access to configuration and data.
* `conn_timeout` - FortiGate Cloud connection timeout in seconds.
* `enc_algorithm` - Configure the level of SSL protection for secure communication with FortiCloud.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `max_log_rate` - FortiCloud maximum log rate in MBps (0 = unlimited).
* `priority` - Set log transmission priority.
* `source_ip` - Source IP address used to connect FortiCloud.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `status` - Enable/disable logging to FortiCloud.
* `upload_day` - Day of week to roll logs.
* `upload_interval` - Frequency of uploading log files to FortiCloud.
* `upload_option` - Configure how log messages are sent to FortiCloud.
* `upload_time` - Time of day to roll logs (hh:mm).
