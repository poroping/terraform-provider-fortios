---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_setting"
description: |-
  VDOM wireless controller configuration.
---

## fortios_wirelesscontroller_setting
VDOM wireless controller configuration.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `account_id` - FortiCloud customer account ID.
* `country` - Country or region in which the FortiGate is located. The country determines the 802.11 bands and channels that are available. Valid values: `--` `AF` `AL` `DZ` `AS` `AO` `AR` `AM` `AU` `AT` `AZ` `BS` `BH` `BD` `BB` `BY` `BE` `BZ` `BJ` `BM` `BT` `BO` `BA` `BW` `BR` `BN` `BG` `BF` `KH` `CM` `KY` `CF` `TD` `CL` `CN` `CX` `CO` `CG` `CD` `CR` `HR` `CY` `CZ` `DK` `DM` `DO` `EC` `EG` `SV` `ET` `EE` `GF` `PF` `FO` `FJ` `FI` `FR` `GE` `DE` `GH` `GI` `GR` `GL` `GD` `GP` `GU` `GT` `GY` `HT` `HN` `HK` `HU` `IS` `IN` `ID` `IQ` `IE` `IM` `IL` `IT` `CI` `JM` `JO` `KZ` `KE` `KR` `KW` `LA` `LV` `LB` `LS` `LY` `LI` `LT` `LU` `MO` `MK` `MG` `MW` `MY` `MV` `ML` `MT` `MH` `MQ` `MR` `MU` `YT` `MX` `FM` `MD` `MC` `MA` `MZ` `MM` `NA` `NP` `NL` `AN` `AW` `NZ` `NI` `NE` `NO` `MP` `OM` `PK` `PW` `PA` `PG` `PY` `PE` `PH` `PL` `PT` `PR` `QA` `RE` `RO` `RU` `RW` `BL` `KN` `LC` `MF` `PM` `VC` `SA` `SN` `RS` `ME` `SL` `SG` `SK` `SI` `ZA` `ES` `LK` `SE` `SR` `CH` `TW` `TZ` `TH` `TG` `TT` `TN` `TR` `TM` `AE` `TC` `UG` `UA` `GB` `US` `PS` `UY` `UZ` `VU` `VE` `VN` `VI` `WF` `YE` `ZM` `ZW` `JP` `CA` .
* `darrp_optimize` - Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).
* `device_holdoff` - Lower limit of creation time of device for identification in minutes (0 - 60, default = 5).
* `device_idle` - Upper limit of idle time of device for identification in minutes (0 - 14400, default = 1440).
* `device_weight` - Upper limit of confidence of device for identification (0 - 255, default = 1, 0 = disable).
* `duplicate_ssid` - Enable/disable allowing Virtual Access Points (VAPs) to use the same SSID name in the same VDOM. Valid values: `enable` `disable` .
* `fake_ssid_action` - Actions taken for detected fake SSID. Valid values: `log` `suppress` .
* `fapc_compatibility` - Enable/disable FAP-C series compatibility. Valid values: `enable` `disable` .
* `firmware_provision_on_authorization` - Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable` `disable` .
* `phishing_ssid_detect` - Enable/disable phishing SSID detection. Valid values: `enable` `disable` .
* `wfa_compatibility` - Enable/disable WFA compatibility. Valid values: `enable` `disable` .
* `darrp_optimize_schedules` - Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space. The structure of `darrp_optimize_schedules` block is documented below.

The `darrp_optimize_schedules` block contains:

* `name` - Schedule name. This attribute must reference one of the following datasources: `firewall.schedule.group.name` `firewall.schedule.recurring.name` `firewall.schedule.onetime.name` .
* `offending_ssid` - Configure offending SSID. The structure of `offending_ssid` block is documented below.

The `offending_ssid` block contains:

* `action` - Actions taken for detected offending SSID. Valid values: `log` `suppress` .
* `id` - ID.
* `ssid_pattern` - Define offending SSID pattern (case insensitive), eg: word, word*, *word, wo*rd.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontroller_setting can be imported using:
```sh
terraform import fortios_wirelesscontroller_setting.labelname {{mkey}}
```
