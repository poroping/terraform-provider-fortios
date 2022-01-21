---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetserviceextension"
description: |-
  Configure Internet Services Extension.
---

## fortios_firewall_internetserviceextension
Configure Internet Services Extension.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comment` - Comment.
* `id` - Internet Service ID in the Internet Service database. This attribute must reference one of the following datasources: `firewall.internet-service.id` .
* `disable_entry` - Disable entries in the Internet Service database. The structure of `disable_entry` block is documented below.

The `disable_entry` block contains:

* `id` - Disable entry ID.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `ip_range` - IP ranges in the disable entry. The structure of `ip_range` block is documented below.

The `ip_range` block contains:

* `end_ip` - End IP address.
* `id` - Disable entry range ID.
* `start_ip` - Start IP address.
* `port_range` - Port ranges in the disable entry. The structure of `port_range` block is documented below.

The `port_range` block contains:

* `end_port` - Ending TCP/UDP/SCTP destination port (1 to 65535).
* `id` - Custom entry port range ID.
* `start_port` - Starting TCP/UDP/SCTP destination port (1 to 65535).
* `entry` - Entries added to the Internet Service extension database. The structure of `entry` block is documented below.

The `entry` block contains:

* `id` - Entry ID(1-255).
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `dst` - Destination address or address group name. The structure of `dst` block is documented below.

The `dst` block contains:

* `name` - Select the destination address or address group object from available options. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
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

fortios_firewall_internetserviceextension can be imported using:
```sh
terraform import fortios_firewall_internetserviceextension.labelname {{mkey}}
```
