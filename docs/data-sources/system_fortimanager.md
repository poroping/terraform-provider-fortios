---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fortimanager"
description: |-
  Get information on a fortios Configure FortiManager.
---

# Data Source: fortios_system_fortimanager
Use this data source to get information on a fortios Configure FortiManager.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `central_management` - Enable/disable FortiManager central management.
* `central_mgmt_auto_backup` - Enable/disable central management auto backup.
* `central_mgmt_schedule_config_restore` - Enable/disable central management schedule config restore.
* `central_mgmt_schedule_script_restore` - Enable/disable central management schedule script restore.
* `ip` - IP address.
* `ipsec` - Enable/disable FortiManager IPsec tunnel.
* `vdom` - Virtual domain name.
