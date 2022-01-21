---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_mpskprofile"
description: |-
  Configure MPSK profile.
---

## fortios_wirelesscontroller_mpskprofile
Configure MPSK profile.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `mpsk_concurrent_clients` - Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).
* `name` - MPSK profile name.
* `mpsk_group` - List of multiple PSK groups. The structure of `mpsk_group` block is documented below.

The `mpsk_group` block contains:

* `name` - MPSK group name.
* `vlan_id` - Optional VLAN ID.
* `vlan_type` - MPSK group VLAN options. Valid values: `no-vlan` `fixed-vlan` .
* `mpsk_key` - List of multiple PSK entries. The structure of `mpsk_key` block is documented below.

The `mpsk_key` block contains:

* `comment` - Comment.
* `concurrent_client_limit_type` - MPSK client limit type options. Valid values: `default` `unlimited` `specified` .
* `concurrent_clients` - Number of clients that can connect using this pre-shared key (1 - 65535, default is 256).
* `mac` - MAC address.
* `name` - Pre-shared key name.
* `passphrase` - WPA Pre-shared key.
* `mpsk_schedules` - Firewall schedule for MPSK passphrase. The passphrase will be effective only when at least one schedule is valid. The structure of `mpsk_schedules` block is documented below.

The `mpsk_schedules` block contains:

* `name` - Schedule name. This attribute must reference one of the following datasources: `firewall.schedule.group.name` `firewall.schedule.recurring.name` `firewall.schedule.onetime.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontroller_mpskprofile can be imported using:
```sh
terraform import fortios_wirelesscontroller_mpskprofile.labelname {{mkey}}
```
