---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_arrpprofile"
description: |-
  Get information on a fortios Configure WiFi Automatic Radio Resource Provisioning (ARRP) profiles.
---

# Data Source: fortios_wirelesscontroller_arrpprofile
Use this data source to get information on a fortios Configure WiFi Automatic Radio Resource Provisioning (ARRP) profiles.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) WiFi ARRP profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `include_dfs_channel` - Enable/disable use of DFS channel in DARRP channel selection phase 1 (default = disable).
* `include_weather_channel` - Enable/disable use of weather channel in DARRP channel selection phase 1 (default = disable).
* `monitor_period` - Period in seconds to measure average transmit retries and receive errors (default = 300).
* `name` - WiFi ARRP profile name.
* `selection_period` - Period in seconds to measure average channel load, noise floor, spectral RSSI (default = 3600).
* `threshold_ap` - Threshold to reject channel in DARRP channel selection phase 1 due to surrounding APs (0 - 500, default = 250).
* `threshold_channel_load` - Threshold in percentage to reject channel in DARRP channel selection phase 1 due to channel load (0 - 100, default = 60).
* `threshold_noise_floor` - Threshold in dBm to reject channel in DARRP channel selection phase 1 due to noise floor (-95 to -20, default = -85).
* `threshold_rx_errors` - Threshold in percentage for receive errors to trigger channel reselection in DARRP monitor stage (0 - 100, default = 50).
* `threshold_spectral_rssi` - Threshold in dBm to reject channel in DARRP channel selection phase 1 due to spectral RSSI (-95 to -20, default = -65).
* `threshold_tx_retries` - Threshold in percentage for transmit retries to trigger channel reselection in DARRP monitor stage (0 - 1000, default = 300).
* `weight_channel_load` - Weight in DARRP channel score calculation for channel load (0 - 2000, default = 20).
* `weight_dfs_channel` - Weight in DARRP channel score calculation for DFS channel (0 - 2000, default = 500).
* `weight_managed_ap` - Weight in DARRP channel score calculation for managed APs (0 - 2000, default = 50).
* `weight_noise_floor` - Weight in DARRP channel score calculation for noise floor (0 - 2000, default = 40).
* `weight_rogue_ap` - Weight in DARRP channel score calculation for rogue APs (0 - 2000, default = 10).
* `weight_spectral_rssi` - Weight in DARRP channel score calculation for spectral RSSI (0 - 2000, default = 40).
* `weight_weather_channel` - Weight in DARRP channel score calculation for weather channel (0 - 2000, default = 1000).
