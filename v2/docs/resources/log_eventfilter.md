---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_log_eventfilter"
description: |-
  Configure log event filters.
---

## fortios_log_eventfilter
Configure log event filters.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `cifs` - Enable/disable CIFS logging. Valid values: `enable` `disable` .
* `connector` - Enable/disable SDN connector logging. Valid values: `enable` `disable` .
* `endpoint` - Enable/disable endpoint event logging. Valid values: `enable` `disable` .
* `event` - Enable/disable event logging. Valid values: `enable` `disable` .
* `fortiextender` - Enable/disable FortiExtender logging. Valid values: `enable` `disable` .
* `ha` - Enable/disable ha event logging. Valid values: `enable` `disable` .
* `router` - Enable/disable router event logging. Valid values: `enable` `disable` .
* `sdwan` - Enable/disable SD-WAN logging. Valid values: `enable` `disable` .
* `security_rating` - Enable/disable Security Rating result logging. Valid values: `enable` `disable` .
* `switch_controller` - Enable/disable Switch-Controller logging. Valid values: `enable` `disable` .
* `system` - Enable/disable system event logging. Valid values: `enable` `disable` .
* `user` - Enable/disable user authentication event logging. Valid values: `enable` `disable` .
* `vpn` - Enable/disable VPN event logging. Valid values: `enable` `disable` .
* `wan_opt` - Enable/disable WAN optimization event logging. Valid values: `enable` `disable` .
* `wireless_activity` - Enable/disable wireless event logging. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_log_eventfilter can be imported using:
```sh
terraform import fortios_log_eventfilter.labelname {{mkey}}
```
