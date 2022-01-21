---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logfortiguard_overridesetting"
description: |-
  Get information on a fortios Override global FortiCloud logging settings for this VDOM.
---

# Data Source: fortios_logfortiguard_overridesetting
Use this data source to get information on a fortios Override global FortiCloud logging settings for this VDOM.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `access_config` - Enable/disable FortiCloud access to configuration and data.
* `max_log_rate` - FortiCloud maximum log rate in MBps (0 = unlimited).
* `override` - Overriding FortiCloud settings for this VDOM or use global settings.
* `priority` - Set log transmission priority.
* `status` - Enable/disable logging to FortiCloud.
* `upload_day` - Day of week to roll logs.
* `upload_interval` - Frequency of uploading log files to FortiCloud.
* `upload_option` - Configure how log messages are sent to FortiCloud.
* `upload_time` - Time of day to roll logs (hh:mm).
