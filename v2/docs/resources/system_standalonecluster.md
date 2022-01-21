---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_standalonecluster"
description: |-
  Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes.
---

## fortios_system_standalonecluster
Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `encryption` - Enable/disable encryption when synchronizing sessions. Valid values: `enable` `disable` .
* `group_member_id` - Cluster member ID (0 - 15).
* `layer2_connection` - Indicate whether layer 2 connections are present among FGSP members. Valid values: `available` `unavailable` .
* `psksecret` - Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
* `session_sync_dev` - Offload session-sync process to kernel and sync sessions using connected interface(s) directly. This attribute must reference one of the following datasources: `system.interface.name` .
* `standalone_group_id` - Cluster group ID (0 - 255). Must be the same for all members.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_standalonecluster can be imported using:
```sh
terraform import fortios_system_standalonecluster.labelname {{mkey}}
```
