---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnipsec_phase2interface"
description: |-
  Get information on a fortios Configure VPN autokey tunnel.
---

# Data Source: fortios_vpnipsec_phase2interface
Use this data source to get information on a fortios Configure VPN autokey tunnel.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) IPsec tunnel name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `add_route` - Enable/disable automatic route addition.
* `auto_discovery_forwarder` - Enable/disable forwarding short-cut messages.
* `auto_discovery_sender` - Enable/disable sending short-cut messages.
* `auto_negotiate` - Enable/disable IPsec SA auto-negotiation.
* `comments` - Comment.
* `dhcp_ipsec` - Enable/disable DHCP-IPsec.
* `dhgrp` - Phase2 DH group.
* `diffserv` - Enable/disable applying DSCP value to the IPsec tunnel outer IP header.
* `diffservcode` - DSCP value to be applied to the IPsec tunnel outer IP header.
* `dst_addr_type` - Remote proxy ID type.
* `dst_end_ip` - Remote proxy ID IPv4 end.
* `dst_end_ip6` - Remote proxy ID IPv6 end.
* `dst_name` - Remote proxy ID name.
* `dst_name6` - Remote proxy ID name.
* `dst_port` - Quick mode destination port (1 - 65535 or 0 for all).
* `dst_start_ip` - Remote proxy ID IPv4 start.
* `dst_start_ip6` - Remote proxy ID IPv6 start.
* `dst_subnet` - Remote proxy ID IPv4 subnet.
* `dst_subnet6` - Remote proxy ID IPv6 subnet.
* `encapsulation` - ESP encapsulation mode.
* `inbound_dscp_copy` - Enable/disable copying of the DSCP in the ESP header to the inner IP header.
* `initiator_ts_narrow` - Enable/disable traffic selector narrowing for IKEv2 initiator.
* `ipv4_df` - Enable/disable setting and resetting of IPv4 'Don't Fragment' bit.
* `keepalive` - Enable/disable keep alive.
* `keylife_type` - Keylife type.
* `keylifekbs` - Phase2 key life in number of kilobytes of traffic (5120 - 4294967295).
* `keylifeseconds` - Phase2 key life in time in seconds (120 - 172800).
* `l2tp` - Enable/disable L2TP over IPsec.
* `name` - IPsec tunnel name.
* `pfs` - Enable/disable PFS feature.
* `phase1name` - Phase 1 determines the options required for phase 2.
* `proposal` - Phase2 proposal.
* `protocol` - Quick mode protocol selector (1 - 255 or 0 for all).
* `replay` - Enable/disable replay detection.
* `route_overlap` - Action for overlapping routes.
* `single_source` - Enable/disable single source IP restriction.
* `src_addr_type` - Local proxy ID type.
* `src_end_ip` - Local proxy ID end.
* `src_end_ip6` - Local proxy ID IPv6 end.
* `src_name` - Local proxy ID name.
* `src_name6` - Local proxy ID name.
* `src_port` - Quick mode source port (1 - 65535 or 0 for all).
* `src_start_ip` - Local proxy ID start.
* `src_start_ip6` - Local proxy ID IPv6 start.
* `src_subnet` - Local proxy ID subnet.
* `src_subnet6` - Local proxy ID IPv6 subnet.
