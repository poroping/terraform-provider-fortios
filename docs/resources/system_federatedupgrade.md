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
* `failure_reason` - Reason for upgrade failure. Valid values: `none` `internal` `timeout` `device-type-unsupported` `download-failed` `device-missing` `version-unavailable` `staging-failed` `reboot-failed` `device-not-reconnected` `node-not-ready` `no-final-confirmation` `no-confirmation-query` .
* `next_path_index` - The index of the next image to upgrade to.
* `status` - Current status of the upgrade. Valid values: `disabled` `initialized` `downloading` `device-disconnected` `ready` `staging` `final-check` `upgrade-devices` `cancelled` `confirmed` `done` `failed` .
* `upgrade_id` - Unique identifier for this upgrade.
* `node_list` - Nodes which will be included in the upgrade. The structure of `node_list` block is documented below.

The `node_list` block contains:

* `coordinating_fortigate` - Serial number of the FortiGate unit that controls this device.
* `device_type` - What type of device this node represents. Valid values: `fortigate` `fortiswitch` `fortiap` .
* `serial` - Serial number of the node to include.
* `setup_time` - When the upgrade was configured. Format hh:mm yyyy/mm/dd UTC.
* `time` - Scheduled time for the upgrade. Format hh:mm yyyy/mm/dd UTC.
* `timing` - Whether the upgrade should be run immediately, or at a scheduled time. Valid values: `immediate` `scheduled` .
* `upgrade_path` - Image IDs to upgrade through.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_federatedupgrade can be imported using:
```sh
terraform import fortios_system_federatedupgrade.labelname {{mkey}}
```
