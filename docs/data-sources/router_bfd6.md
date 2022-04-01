---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bfd6"
description: |-
  Get information on a fortios Configure IPv6 BFD.
---

# Data Source: fortios_router_bfd6
Use this data source to get information on a fortios Configure IPv6 BFD.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `multihop_template` - BFD IPv6 multi-hop template table.The structure of `multihop_template` block is documented below.

The `multihop_template` block contains:

* `auth_mode` - Authentication mode.
* `bfd_desired_min_tx` - BFD desired minimal transmit interval (milliseconds).
* `bfd_detect_mult` - BFD detection multiplier.
* `bfd_required_min_rx` - BFD required minimal receive interval (milliseconds).
* `dst` - Destination prefix.
* `id` - ID.
* `md5_key` - MD5 key of key ID 1.
* `src` - Source prefix.
* `neighbor` - Configure neighbor of IPv6 BFD.The structure of `neighbor` block is documented below.

The `neighbor` block contains:

* `interface` - Interface to the BFD neighbor.
* `ip6_address` - IPv6 address of the BFD neighbor.
