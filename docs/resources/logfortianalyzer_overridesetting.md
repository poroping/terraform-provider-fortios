---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logfortianalyzer_overridesetting"
description: |-
  Override FortiAnalyzer settings.
---

## fortios_logfortianalyzer_overridesetting
Override FortiAnalyzer settings.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `access_config` - Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable` `disable` .
* `certificate` - Certificate used to communicate with FortiAnalyzer. This attribute must reference one of the following datasources: `certificate.local.name` .
* `certificate_verification` - Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable` `disable` .
* `conn_timeout` - FortiAnalyzer connection time-out in seconds (for status and log buffer).
* `enc_algorithm` - Configure the level of SSL protection for secure communication with FortiAnalyzer. Valid values: `high-medium` `high` `low` .
* `hmac_algorithm` - OFTP login hash algorithm. Valid values: `sha256` `sha1` .
* `interface` - Specify outgoing interface to reach server. This attribute must reference one of the following datasources: `system.interface.name` .
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `ips_archive` - Enable/disable IPS packet archive logging. Valid values: `enable` `disable` .
* `max_log_rate` - FortiAnalyzer maximum log rate in MBps (0 = unlimited).
* `monitor_failure_retry_period` - Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
* `monitor_keepalive_period` - Time between OFTP keepalives in seconds (for status and log buffer).
* `preshared_key` - Preshared-key used for auto-authorization on FortiAnalyzer.
* `priority` - Set log transmission priority. Valid values: `default` `low` .
* `reliable` - Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable` `disable` .
* `server` - The remote FortiAnalyzer.
* `source_ip` - Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default` `SSLv3` `TLSv1` `TLSv1-1` `TLSv1-2` `TLSv1-3` .
* `status` - Enable/disable logging to FortiAnalyzer. Valid values: `enable` `disable` .
* `upload_day` - Day of week (month) to upload logs.
* `upload_interval` - Frequency to upload log files to FortiAnalyzer. Valid values: `daily` `weekly` `monthly` .
* `upload_option` - Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload` `realtime` `1-minute` `5-minute` .
* `upload_time` - Time to upload logs (hh:mm).
* `use_management_vdom` - Enable/disable use of management VDOM IP address as source IP for logs sent to FortiAnalyzer. Valid values: `enable` `disable` .
* `serial` - Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.

The `serial` block contains:

* `name` - Serial Number.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_logfortianalyzer_overridesetting can be imported using:
```sh
terraform import fortios_logfortianalyzer_overridesetting.labelname {{mkey}}
```
