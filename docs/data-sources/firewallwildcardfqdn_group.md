---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallwildcardfqdn_group"
description: |-
  Get information on a fortios Config global Wildcard FQDN address groups.
---

# Data Source: fortios_firewallwildcardfqdn_group
Use this data source to get information on a fortios Config global Wildcard FQDN address groups.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Address group name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `color` - GUI icon color.
* `comment` - Comment.
* `name` - Address group name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable address visibility.
* `member` - Address group members.The structure of `member` block is documented below.

The `member` block contains:

* `name` - Address name.
