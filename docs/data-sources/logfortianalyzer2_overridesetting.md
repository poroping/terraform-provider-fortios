---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logfortianalyzer2_overridesetting"
description: |-
  Get information on a fortios Override FortiAnalyzer settings.
---

# Data Source: fortios_logfortianalyzer2_overridesetting
Use this data source to get information on a fortios Override FortiAnalyzer settings.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `access_config` - Enable/disable FortiAnalyzer access to configuration and data.
* `alt_server` - Alternate FortiAnalyzer.
* `certificate` - Certificate used to communicate with FortiAnalyzer.
* `certificate_verification` - Enable/disable identity verification of FortiAnalyzer by use of certificate.
* `conn_timeout` - FortiAnalyzer connection time-out in seconds (for status and log buffer).
* `enc_algorithm` - Configure the level of SSL protection for secure communication with FortiAnalyzer.
* `fallback_to_primary` - Enable/disable this FortiGate unit to fallback to the primary FortiAnalyzer when it is available.
* `hmac_algorithm` - OFTP login hash algorithm.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `ips_archive` - Enable/disable IPS packet archive logging.
* `max_log_rate` - FortiAnalyzer maximum log rate in MBps (0 = unlimited).
* `monitor_failure_retry_period` - Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
* `monitor_keepalive_period` - Time between OFTP keepalives in seconds (for status and log buffer).
* `preshared_key` - Preshared-key used for auto-authorization on FortiAnalyzer.
* `priority` - Set log transmission priority.
* `reliable` - Enable/disable reliable logging to FortiAnalyzer.
* `server` - The remote FortiAnalyzer.
* `server_cert_ca` - Mandatory CA on FortiGate in certificate chain of server.
* `source_ip` - Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `status` - Enable/disable logging to FortiAnalyzer.
* `upload_day` - Day of week (month) to upload logs.
* `upload_interval` - Frequency to upload log files to FortiAnalyzer.
* `upload_option` - Enable/disable logging to hard disk and then uploading to FortiAnalyzer.
* `upload_time` - Time to upload logs (hh:mm).
* `use_management_vdom` - Enable/disable use of management VDOM IP address as source IP for logs sent to FortiAnalyzer.
* `serial` - Serial numbers of the FortiAnalyzer.The structure of `serial` block is documented below.

The `serial` block contains:

* `name` - Serial Number.
