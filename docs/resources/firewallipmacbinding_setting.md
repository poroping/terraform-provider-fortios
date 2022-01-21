---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallipmacbinding_setting"
description: |-
  Configure IP to MAC binding settings.
---

## fortios_firewallipmacbinding_setting
Configure IP to MAC binding settings.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `bindthroughfw` - Enable/disable use of IP/MAC binding to filter packets that would normally go through the firewall. Valid values: `enable` `disable` .
* `bindtofw` - Enable/disable use of IP/MAC binding to filter packets that would normally go to the firewall. Valid values: `enable` `disable` .
* `undefinedhost` - Select action to take on packets with IP/MAC addresses not in the binding list (default = block). Valid values: `allow` `block` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewallipmacbinding_setting can be imported using:
```sh
terraform import fortios_firewallipmacbinding_setting.labelname {{mkey}}
```
