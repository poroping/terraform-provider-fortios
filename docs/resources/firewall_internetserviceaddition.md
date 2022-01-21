---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetserviceaddition"
description: |-
  Configure Internet Services Addition.
---

## fortios_firewall_internetserviceaddition
Configure Internet Services Addition.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comment` - Comment.
* `id` - Internet Service ID in the Internet Service database. This attribute must reference one of the following datasources: `firewall.internet-service.id` .
* `entry` - Entries added to the Internet Service addition database. The structure of `entry` block is documented below.

The `entry` block contains:

* `id` - Entry ID(1-255).
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `port_range` - Port ranges in the custom entry. The structure of `port_range` block is documented below.

The `port_range` block contains:

* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (1 to 65535).
* `id` - Custom entry port range ID.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (1 to 65535).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_internetserviceaddition can be imported using:
```sh
terraform import fortios_firewall_internetserviceaddition.labelname {{mkey}}
```
