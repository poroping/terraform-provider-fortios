---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_federatedupgrade"
description: |-
  Coordinate federated upgrades within the Security Fabric.
---

## fortios_system_federatedupgrade
Coordinate federated upgrades within the Security Fabric.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `failure_device` - Serial number of the node to include.
* `failure_reason` - Reason for upgrade failure. Valid values: `none` `internal` `timeout` `device-type-unsupported` `download-failed` `device-missing` `version-unavailable` `staging-failed` `reboot-failed` `device-not-reconnected` `node-not-ready` `no-final-confirmation` `no-confirmation-query` `config-error-log-nonempty` `node-failed` .
* `ha_reboot_controller` - Serial number of the FortiGate unit that will control the reboot process for the federated upgrade of the HA cluster.
* `next_path_index` - The index of the next image to upgrade to.
* `status` - Current status of the upgrade. Valid values: `disabled` `initialized` `downloading` `device-disconnected` `ready` `coordinating` `staging` `final-check` `upgrade-devices` `cancelled` `confirmed` `done` `failed` .
* `upgrade_id` - Unique identifier for this upgrade.
* `node_list` - Nodes which will be included in the upgrade. The structure of `node_list` block is documented below.

The `node_list` block contains:

* `coordinating_fortigate` - Serial number of the FortiGate unit that controls this device.
* `device_type` - Fortinet device type. Valid values: `fortigate` `fortiswitch` `fortiap` `fortiextender` .
* `maximum_minutes` - Maximum number of minutes to allow for immediate upgrade preparation.
* `serial` - Serial number of the node to include.
* `setup_time` - Upgrade preparation start time in UTC (hh:mm yyyy/mm/dd UTC).
* `time` - Scheduled upgrade execution time in UTC (hh:mm yyyy/mm/dd UTC).
* `timing` - Run immediately or at a scheduled time. Valid values: `immediate` `scheduled` .
* `upgrade_path` - Fortinet OS image versions to upgrade through in major-minor-patch format, such as 7-0-4.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_federatedupgrade can be imported using:
```sh
terraform import fortios_system_federatedupgrade.labelname {{mkey}}
```
