---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_mpskprofile"
description: |-
  Get information on a fortios Configure MPSK profile.
---

# Data Source: fortios_wirelesscontroller_mpskprofile
Use this data source to get information on a fortios Configure MPSK profile.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) MPSK profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `mpsk_concurrent_clients` - Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).
* `name` - MPSK profile name.
* `mpsk_group` - List of multiple PSK groups.The structure of `mpsk_group` block is documented below.

The `mpsk_group` block contains:

* `name` - MPSK group name.
* `vlan_id` - Optional VLAN ID.
* `vlan_type` - MPSK group VLAN options.
* `mpsk_key` - List of multiple PSK entries.The structure of `mpsk_key` block is documented below.

The `mpsk_key` block contains:

* `comment` - Comment.
* `concurrent_client_limit_type` - MPSK client limit type options.
* `concurrent_clients` - Number of clients that can connect using this pre-shared key (1 - 65535, default is 256).
* `mac` - MAC address.
* `name` - Pre-shared key name.
* `passphrase` - WPA Pre-shared key.
* `mpsk_schedules` - Firewall schedule for MPSK passphrase. The passphrase will be effective only when at least one schedule is valid.The structure of `mpsk_schedules` block is documented below.

The `mpsk_schedules` block contains:

* `name` - Schedule name.
