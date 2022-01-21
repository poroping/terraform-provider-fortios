---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_widsprofile"
description: |-
  Configure wireless intrusion detection system (WIDS) profiles.
---

## fortios_wirelesscontroller_widsprofile
Configure wireless intrusion detection system (WIDS) profiles.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `ap_auto_suppress` - Enable/disable on-wire rogue AP auto-suppression (default = disable). Valid values: `enable` `disable` .
* `ap_bgscan_duration` - Listen time on scanning a channel (10 - 1000 msec, default = 30).
* `ap_bgscan_idle` - Wait time for channel inactivity before scanning this channel (0 - 1000 msec, default = 20).
* `ap_bgscan_intv` - Period between successive channel scans (1 - 600 sec, default = 3).
* `ap_bgscan_period` - Period between background scans (10 - 3600 sec, default = 600).
* `ap_bgscan_report_intv` - Period between background scan reports (15 - 600 sec, default = 30).
* `ap_fgscan_report_intv` - Period between foreground scan reports (15 - 600 sec, default = 15).
* `ap_scan` - Enable/disable rogue AP detection. Valid values: `disable` `enable` .
* `ap_scan_passive` - Enable/disable passive scanning. Enable means do not send probe request on any channels (default = disable). Valid values: `enable` `disable` .
* `ap_scan_threshold` - Minimum signal level/threshold in dBm required for the AP to report detected rogue AP (-95 to -20, default = -90).
* `asleap_attack` - Enable/disable asleap attack detection (default = disable). Valid values: `enable` `disable` .
* `assoc_flood_thresh` - The threshold value for association frame flooding.
* `assoc_flood_time` - Number of seconds after which a station is considered not connected.
* `assoc_frame_flood` - Enable/disable association frame flooding detection (default = disable). Valid values: `enable` `disable` .
* `auth_flood_thresh` - The threshold value for authentication frame flooding.
* `auth_flood_time` - Number of seconds after which a station is considered not connected.
* `auth_frame_flood` - Enable/disable authentication frame flooding detection (default = disable). Valid values: `enable` `disable` .
* `comment` - Comment.
* `deauth_broadcast` - Enable/disable broadcasting de-authentication detection (default = disable). Valid values: `enable` `disable` .
* `deauth_unknown_src_thresh` - Threshold value per second to deauth unknown src for DoS attack (0: no limit).
* `eapol_fail_flood` - Enable/disable EAPOL-Failure flooding (to AP) detection (default = disable). Valid values: `enable` `disable` .
* `eapol_fail_intv` - The detection interval for EAPOL-Failure flooding (1 - 3600 sec).
* `eapol_fail_thresh` - The threshold value for EAPOL-Failure flooding in specified interval.
* `eapol_logoff_flood` - Enable/disable EAPOL-Logoff flooding (to AP) detection (default = disable). Valid values: `enable` `disable` .
* `eapol_logoff_intv` - The detection interval for EAPOL-Logoff flooding (1 - 3600 sec).
* `eapol_logoff_thresh` - The threshold value for EAPOL-Logoff flooding in specified interval.
* `eapol_pre_fail_flood` - Enable/disable premature EAPOL-Failure flooding (to STA) detection (default = disable). Valid values: `enable` `disable` .
* `eapol_pre_fail_intv` - The detection interval for premature EAPOL-Failure flooding (1 - 3600 sec).
* `eapol_pre_fail_thresh` - The threshold value for premature EAPOL-Failure flooding in specified interval.
* `eapol_pre_succ_flood` - Enable/disable premature EAPOL-Success flooding (to STA) detection (default = disable). Valid values: `enable` `disable` .
* `eapol_pre_succ_intv` - The detection interval for premature EAPOL-Success flooding (1 - 3600 sec).
* `eapol_pre_succ_thresh` - The threshold value for premature EAPOL-Success flooding in specified interval.
* `eapol_start_flood` - Enable/disable EAPOL-Start flooding (to AP) detection (default = disable). Valid values: `enable` `disable` .
* `eapol_start_intv` - The detection interval for EAPOL-Start flooding (1 - 3600 sec).
* `eapol_start_thresh` - The threshold value for EAPOL-Start flooding in specified interval.
* `eapol_succ_flood` - Enable/disable EAPOL-Success flooding (to AP) detection (default = disable). Valid values: `enable` `disable` .
* `eapol_succ_intv` - The detection interval for EAPOL-Success flooding (1 - 3600 sec).
* `eapol_succ_thresh` - The threshold value for EAPOL-Success flooding in specified interval.
* `invalid_mac_oui` - Enable/disable invalid MAC OUI detection. Valid values: `enable` `disable` .
* `long_duration_attack` - Enable/disable long duration attack detection based on user configured threshold (default = disable). Valid values: `enable` `disable` .
* `long_duration_thresh` - Threshold value for long duration attack detection (1000 - 32767 usec, default = 8200).
* `name` - WIDS profile name.
* `null_ssid_probe_resp` - Enable/disable null SSID probe response detection (default = disable). Valid values: `enable` `disable` .
* `sensor_mode` - Scan nearby WiFi stations (default = disable). Valid values: `disable` `foreign` `both` .
* `spoofed_deauth` - Enable/disable spoofed de-authentication attack detection (default = disable). Valid values: `enable` `disable` .
* `weak_wep_iv` - Enable/disable weak WEP IV (Initialization Vector) detection (default = disable). Valid values: `enable` `disable` .
* `wireless_bridge` - Enable/disable wireless bridge detection (default = disable). Valid values: `enable` `disable` .
* `ap_bgscan_disable_schedules` - Firewall schedules for turning off FortiAP radio background scan. Background scan will be disabled when at least one of the schedules is valid. Separate multiple schedule names with a space. The structure of `ap_bgscan_disable_schedules` block is documented below.

The `ap_bgscan_disable_schedules` block contains:

* `name` - Schedule name. This attribute must reference one of the following datasources: `firewall.schedule.group.name` `firewall.schedule.recurring.name` `firewall.schedule.onetime.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontroller_widsprofile can be imported using:
```sh
terraform import fortios_wirelesscontroller_widsprofile.labelname {{mkey}}
```
