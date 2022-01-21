---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetserviceextension"
description: |-
  Get information on a fortios Configure Internet Services Extension.
---

# Data Source: fortios_firewall_internetserviceextension
Use this data source to get information on a fortios Configure Internet Services Extension.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Internet Service ID in the Internet Service database.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `id` - Internet Service ID in the Internet Service database.
* `disable_entry` - Disable entries in the Internet Service database.The structure of `disable_entry` block is documented below.

The `disable_entry` block contains:

* `id` - Disable entry ID.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `ip_range` - IP ranges in the disable entry.The structure of `ip_range` block is documented below.

The `ip_range` block contains:

* `end_ip` - End IP address.
* `id` - Disable entry range ID.
* `start_ip` - Start IP address.
* `port_range` - Port ranges in the disable entry.The structure of `port_range` block is documented below.

The `port_range` block contains:

* `end_port` - Ending TCP/UDP/SCTP destination port (1 to 65535).
* `id` - Custom entry port range ID.
* `start_port` - Starting TCP/UDP/SCTP destination port (1 to 65535).
* `entry` - Entries added to the Internet Service extension database.The structure of `entry` block is documented below.

The `entry` block contains:

* `id` - Entry ID(1-255).
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `dst` - Destination address or address group name.The structure of `dst` block is documented below.

The `dst` block contains:

* `name` - Select the destination address or address group object from available options.
* `port_range` - Port ranges in the custom entry.The structure of `port_range` block is documented below.

The `port_range` block contains:

* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (1 to 65535).
* `id` - Custom entry port range ID.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (1 to 65535).
