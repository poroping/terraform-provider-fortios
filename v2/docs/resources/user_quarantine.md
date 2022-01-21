---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_quarantine"
description: |-
  Configure quarantine support.
---

## fortios_user_quarantine
Configure quarantine support.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `firewall_groups` - Firewall address group which includes all quarantine MAC address. This attribute must reference one of the following datasources: `firewall.addrgrp.name` .
* `quarantine` - Enable/disable quarantine. Valid values: `enable` `disable` .
* `traffic_policy` - Traffic policy for quarantined MACs. This attribute must reference one of the following datasources: `switch-controller.traffic-policy.name` .
* `targets` - Quarantine entry to hold multiple MACs. The structure of `targets` block is documented below.

The `targets` block contains:

* `description` - Description for the quarantine entry.
* `entry` - Quarantine entry name.
* `macs` - Quarantine MACs. The structure of `macs` block is documented below.

The `macs` block contains:

* `description` - Description for the quarantine MAC.
* `drop` - Enable/Disable dropping of quarantined device traffic Valid values: `disable` `enable` .
* `mac` - Quarantine MAC.
* `parent` - Parent entry name.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_user_quarantine can be imported using:
```sh
terraform import fortios_user_quarantine.labelname {{mkey}}
```
