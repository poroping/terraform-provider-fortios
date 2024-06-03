---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_federatedupgrade"
description: |-
  Get information on a fortios Coordinate federated upgrades within the Security Fabric.
---

# Data Source: fortios_system_federatedupgrade
Use this data source to get information on a fortios Coordinate federated upgrades within the Security Fabric.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `failure_device` - Serial number of the node to include.
* `failure_reason` - Reason for upgrade failure.
* `ha_reboot_controller` - Serial number of the FortiGate unit that will control the reboot process for the federated upgrade of the HA cluster.
* `next_path_index` - The index of the next image to upgrade to.
* `status` - Current status of the upgrade.
* `upgrade_id` - Unique identifier for this upgrade.
* `node_list` - Nodes which will be included in the upgrade.The structure of `node_list` block is documented below.

The `node_list` block contains:

* `coordinating_fortigate` - Serial number of the FortiGate unit that controls this device.
* `device_type` - Fortinet device type.
* `maximum_minutes` - Maximum number of minutes to allow for immediate upgrade preparation.
* `serial` - Serial number of the node to include.
* `setup_time` - Upgrade preparation start time in UTC (hh:mm yyyy/mm/dd UTC).
* `time` - Scheduled upgrade execution time in UTC (hh:mm yyyy/mm/dd UTC).
* `timing` - Run immediately or at a scheduled time.
* `upgrade_path` - Fortinet OS image versions to upgrade through in major-minor-patch format, such as 7-0-4.
