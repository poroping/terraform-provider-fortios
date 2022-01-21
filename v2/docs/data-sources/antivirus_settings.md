---
subcategory: "FortiGate Antivirus"
layout: "fortios"
page_title: "FortiOS: fortios_antivirus_settings"
description: |-
  Get information on a fortios Configure AntiVirus settings.
---

# Data Source: fortios_antivirus_settings
Use this data source to get information on a fortios Configure AntiVirus settings.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `cache_clean_result` - Enable/disable cache of clean scan results (default = enable).
* `cache_infected_result` - Enable/disable cache of infected scan results (default = enable).
* `default_db` - Select the AV database to be used for AV scanning.
* `grayware` - Enable/disable grayware detection when an AntiVirus profile is applied to traffic.
* `machine_learning_detection` - Use machine learning based malware detection.
* `override_timeout` - Override the large file scan timeout value in seconds (30 - 3600). Zero is the default value and is used to disable this command. When disabled, the daemon adjusts the large file scan timeout based on the file size.
* `use_extreme_db` - Enable/disable the use of Extreme AVDB.
