---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logfortiguard_setting"
description: |-
  Configure logging to FortiCloud.
---

## fortios_logfortiguard_setting
Configure logging to FortiCloud.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `access_config` - Enable/disable FortiCloud access to configuration and data. Valid values: `enable` `disable` .
* `conn_timeout` - FortiGate Cloud connection timeout in seconds.
* `enc_algorithm` - Configure the level of SSL protection for secure communication with FortiCloud. Valid values: `high-medium` `high` `low` .
* `interface` - Specify outgoing interface to reach server. This attribute must reference one of the following datasources: `system.interface.name` .
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `max_log_rate` - FortiCloud maximum log rate in MBps (0 = unlimited).
* `priority` - Set log transmission priority. Valid values: `default` `low` .
* `source_ip` - Source IP address used to connect FortiCloud.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default` `SSLv3` `TLSv1` `TLSv1-1` `TLSv1-2` .
* `status` - Enable/disable logging to FortiCloud. Valid values: `enable` `disable` .
* `upload_day` - Day of week to roll logs.
* `upload_interval` - Frequency of uploading log files to FortiCloud. Valid values: `daily` `weekly` `monthly` .
* `upload_option` - Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload` `realtime` `1-minute` `5-minute` .
* `upload_time` - Time of day to roll logs (hh:mm).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_logfortiguard_setting can be imported using:
```sh
terraform import fortios_logfortiguard_setting.labelname {{mkey}}
```
