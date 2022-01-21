---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_bleprofile"
description: |-
  Get information on a fortios Configure Bluetooth Low Energy profile.
---

# Data Source: fortios_wirelesscontroller_bleprofile
Use this data source to get information on a fortios Configure Bluetooth Low Energy profile.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Bluetooth Low Energy profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `advertising` - Advertising type.
* `beacon_interval` - Beacon interval (default = 100 msec).
* `ble_scanning` - Enable/disable Bluetooth Low Energy (BLE) scanning.
* `comment` - Comment.
* `eddystone_instance` - Eddystone instance ID.
* `eddystone_namespace` - Eddystone namespace ID.
* `eddystone_url` - Eddystone URL.
* `ibeacon_uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `major_id` - Major ID.
* `minor_id` - Minor ID.
* `name` - Bluetooth Low Energy profile name.
* `txpower` - Transmit power level (default = 0).
