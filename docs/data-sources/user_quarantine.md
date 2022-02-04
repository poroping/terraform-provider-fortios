---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_quarantine"
description: |-
  Get information on a fortios Configure quarantine support.
---

# Data Source: fortios_user_quarantine
Use this data source to get information on a fortios Configure quarantine support.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `firewall_groups` - Firewall address group which includes all quarantine MAC address.
* `quarantine` - Enable/disable quarantine.
* `traffic_policy` - Traffic policy for quarantined MACs.
* `targets` - Quarantine entry to hold multiple MACs.The structure of `targets` block is documented below.

The `targets` block contains:

* `description` - Description for the quarantine entry.
* `entry` - Quarantine entry name.
* `macs` - Quarantine MACs.The structure of `macs` block is documented below.

The `macs` block contains:

* `description` - Description for the quarantine MAC.
* `drop` - Enable/disable dropping of quarantined device traffic.
* `mac` - Quarantine MAC.
* `parent` - Parent entry name.
