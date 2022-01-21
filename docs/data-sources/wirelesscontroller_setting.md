---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_setting"
description: |-
  Get information on a fortios VDOM wireless controller configuration.
---

# Data Source: fortios_wirelesscontroller_setting
Use this data source to get information on a fortios VDOM wireless controller configuration.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `account_id` - FortiCloud customer account ID.
* `country` - Country or region in which the FortiGate is located. The country determines the 802.11 bands and channels that are available.
* `darrp_optimize` - Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).
* `device_holdoff` - Lower limit of creation time of device for identification in minutes (0 - 60, default = 5).
* `device_idle` - Upper limit of idle time of device for identification in minutes (0 - 14400, default = 1440).
* `device_weight` - Upper limit of confidence of device for identification (0 - 255, default = 1, 0 = disable).
* `duplicate_ssid` - Enable/disable allowing Virtual Access Points (VAPs) to use the same SSID name in the same VDOM.
* `fake_ssid_action` - Actions taken for detected fake SSID.
* `fapc_compatibility` - Enable/disable FAP-C series compatibility.
* `firmware_provision_on_authorization` - Enable/disable automatic provisioning of latest firmware on authorization.
* `phishing_ssid_detect` - Enable/disable phishing SSID detection.
* `wfa_compatibility` - Enable/disable WFA compatibility.
* `darrp_optimize_schedules` - Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space.The structure of `darrp_optimize_schedules` block is documented below.

The `darrp_optimize_schedules` block contains:

* `name` - Schedule name.
* `offending_ssid` - Configure offending SSID.The structure of `offending_ssid` block is documented below.

The `offending_ssid` block contains:

* `action` - Actions taken for detected offending SSID.
* `id` - ID.
* `ssid_pattern` - Define offending SSID pattern (case insensitive), eg: word, word*, *word, wo*rd.
