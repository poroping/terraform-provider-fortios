---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fortimanager"
description: |-
  Configure FortiManager.
---

## fortios_system_fortimanager
Configure FortiManager.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `central_management` - Enable/disable FortiManager central management. Valid values: `enable` `disable` .
* `central_mgmt_auto_backup` - Enable/disable central management auto backup. Valid values: `enable` `disable` .
* `central_mgmt_schedule_config_restore` - Enable/disable central management schedule config restore. Valid values: `enable` `disable` .
* `central_mgmt_schedule_script_restore` - Enable/disable central management schedule script restore. Valid values: `enable` `disable` .
* `ip` - IP address.
* `ipsec` - Enable/disable FortiManager IPsec tunnel. Valid values: `enable` `disable` .
* `vdom` - Virtual domain name. This attribute must reference one of the following datasources: `system.vdom.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_fortimanager can be imported using:
```sh
terraform import fortios_system_fortimanager.labelname {{mkey}}
```
