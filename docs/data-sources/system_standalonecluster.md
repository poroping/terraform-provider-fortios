---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_standalonecluster"
description: |-
  Get information on a fortios Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes.
---

# Data Source: fortios_system_standalonecluster
Use this data source to get information on a fortios Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `encryption` - Enable/disable encryption when synchronizing sessions.
* `group_member_id` - Cluster member ID (0 - 15).
* `layer2_connection` - Indicate whether layer 2 connections are present among FGSP members.
* `psksecret` - Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
* `session_sync_dev` - Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
* `standalone_group_id` - Cluster group ID (0 - 255). Must be the same for all members.
