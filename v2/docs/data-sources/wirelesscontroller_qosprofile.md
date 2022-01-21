---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_qosprofile"
description: |-
  Get information on a fortios Configure WiFi quality of service (QoS) profiles.
---

# Data Source: fortios_wirelesscontroller_qosprofile
Use this data source to get information on a fortios Configure WiFi quality of service (QoS) profiles.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) WiFi QoS profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `bandwidth_admission_control` - Enable/disable WMM bandwidth admission control.
* `bandwidth_capacity` - Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).
* `burst` - Enable/disable client rate burst.
* `call_admission_control` - Enable/disable WMM call admission control.
* `call_capacity` - Maximum number of Voice over WLAN (VoWLAN) phones allowed (0 - 60, default = 10).
* `comment` - Comment.
* `downlink` - Maximum downlink bandwidth for Virtual Access Points (VAPs) (0 - 2097152 Kbps, default = 0, 0 means no limit).
* `downlink_sta` - Maximum downlink bandwidth for clients (0 - 2097152 Kbps, default = 0, 0 means no limit).
* `dscp_wmm_mapping` - Enable/disable Differentiated Services Code Point (DSCP) mapping.
* `name` - WiFi QoS profile name.
* `uplink` - Maximum uplink bandwidth for Virtual Access Points (VAPs) (0 - 2097152 Kbps, default = 0, 0 means no limit).
* `uplink_sta` - Maximum uplink bandwidth for clients (0 - 2097152 Kbps, default = 0, 0 means no limit).
* `wmm` - Enable/disable WiFi multi-media (WMM) control.
* `wmm_be_dscp` - DSCP marking for best effort access (default = 0).
* `wmm_bk_dscp` - DSCP marking for background access (default = 8).
* `wmm_dscp_marking` - Enable/disable WMM Differentiated Services Code Point (DSCP) marking.
* `wmm_uapsd` - Enable/disable WMM Unscheduled Automatic Power Save Delivery (U-APSD) power save mode.
* `wmm_vi_dscp` - DSCP marking for video access (default = 32).
* `wmm_vo_dscp` - DSCP marking for voice access (default = 48).
* `dscp_wmm_be` - DSCP mapping for best effort access (default = 0 24).The structure of `dscp_wmm_be` block is documented below.

The `dscp_wmm_be` block contains:

* `id` - DSCP WMM mapping numbers (0 - 63).
* `dscp_wmm_bk` - DSCP mapping for background access (default = 8 16).The structure of `dscp_wmm_bk` block is documented below.

The `dscp_wmm_bk` block contains:

* `id` - DSCP WMM mapping numbers (0 - 63).
* `dscp_wmm_vi` - DSCP mapping for video access (default = 32 40).The structure of `dscp_wmm_vi` block is documented below.

The `dscp_wmm_vi` block contains:

* `id` - DSCP WMM mapping numbers (0 - 63).
* `dscp_wmm_vo` - DSCP mapping for voice access (default = 48 56).The structure of `dscp_wmm_vo` block is documented below.

The `dscp_wmm_vo` block contains:

* `id` - DSCP WMM mapping numbers (0 - 63).
