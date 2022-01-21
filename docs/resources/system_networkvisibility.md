---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_networkvisibility"
description: |-
  Configure network visibility settings.
---

## fortios_system_networkvisibility
Configure network visibility settings.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `destination_hostname_visibility` - Enable/disable logging of destination hostname visibility. Valid values: `disable` `enable` .
* `destination_location` - Enable/disable logging of destination geographical location visibility. Valid values: `disable` `enable` .
* `destination_visibility` - Enable/disable logging of destination visibility. Valid values: `disable` `enable` .
* `hostname_limit` - Limit of the number of hostname table entries (0 - 50000).
* `hostname_ttl` - TTL of hostname table entries (60 - 86400).
* `source_location` - Enable/disable logging of source geographical location visibility. Valid values: `disable` `enable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_networkvisibility can be imported using:
```sh
terraform import fortios_system_networkvisibility.labelname {{mkey}}
```
