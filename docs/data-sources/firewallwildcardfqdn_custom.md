---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallwildcardfqdn_custom"
description: |-
  Get information on a fortios Config global/VDOM Wildcard FQDN address.
---

# Data Source: fortios_firewallwildcardfqdn_custom
Use this data source to get information on a fortios Config global/VDOM Wildcard FQDN address.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Address name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `color` - GUI icon color.
* `comment` - Comment.
* `name` - Address name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable address visibility.
* `wildcard_fqdn` - Wildcard FQDN.
