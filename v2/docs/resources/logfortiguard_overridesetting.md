---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logfortiguard_overridesetting"
description: |-
  Override global FortiCloud logging settings for this VDOM.
---

## fortios_logfortiguard_overridesetting
Override global FortiCloud logging settings for this VDOM.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `access_config` - Enable/disable FortiCloud access to configuration and data. Valid values: `enable` `disable` .
* `max_log_rate` - FortiCloud maximum log rate in MBps (0 = unlimited).
* `override` - Overriding FortiCloud settings for this VDOM or use global settings. Valid values: `enable` `disable` .
* `priority` - Set log transmission priority. Valid values: `default` `low` .
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

fortios_logfortiguard_overridesetting can be imported using:
```sh
terraform import fortios_logfortiguard_overridesetting.labelname {{mkey}}
```
