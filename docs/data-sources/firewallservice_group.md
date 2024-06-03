---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallservice_group"
description: |-
  Get information on a fortios Configure service groups.
---

# Data Source: fortios_firewallservice_group
Use this data source to get information on a fortios Configure service groups.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Service group name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `fabric_object` - Security Fabric global object setting.
* `name` - Service group name.
* `proxy` - Enable/disable web proxy service group.
* `member` - Service objects contained within the group.The structure of `member` block is documented below.

The `member` block contains:

* `name` - Service or service group name.
