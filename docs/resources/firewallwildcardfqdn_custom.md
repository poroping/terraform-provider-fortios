---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallwildcardfqdn_custom"
description: |-
  Config global/VDOM Wildcard FQDN address.
---

## fortios_firewallwildcardfqdn_custom
Config global/VDOM Wildcard FQDN address.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `color` - GUI icon color.
* `comment` - Comment.
* `name` - Address name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable address visibility. Valid values: `enable` `disable` .
* `wildcard_fqdn` - Wildcard FQDN.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewallwildcardfqdn_custom can be imported using:
```sh
terraform import fortios_firewallwildcardfqdn_custom.labelname {{mkey}}
```
