---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_bleprofile"
description: |-
  Configure Bluetooth Low Energy profile.
---

## fortios_wirelesscontroller_bleprofile
Configure Bluetooth Low Energy profile.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `advertising` - Advertising type. Valid values: `ibeacon` `eddystone-uid` `eddystone-url` .
* `beacon_interval` - Beacon interval (default = 100 msec).
* `ble_scanning` - Enable/disable Bluetooth Low Energy (BLE) scanning. Valid values: `enable` `disable` .
* `comment` - Comment.
* `eddystone_instance` - Eddystone instance ID.
* `eddystone_namespace` - Eddystone namespace ID.
* `eddystone_url` - Eddystone URL.
* `ibeacon_uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `major_id` - Major ID.
* `minor_id` - Minor ID.
* `name` - Bluetooth Low Energy profile name.
* `txpower` - Transmit power level (default = 0). Valid values: `0` `1` `2` `3` `4` `5` `6` `7` `8` `9` `10` `11` `12` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontroller_bleprofile can be imported using:
```sh
terraform import fortios_wirelesscontroller_bleprofile.labelname {{mkey}}
```
