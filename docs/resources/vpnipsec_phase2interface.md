---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnipsec_phase2interface"
description: |-
  Configure VPN autokey tunnel.
---

## fortios_vpnipsec_phase2interface
Configure VPN autokey tunnel.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `add_route` - Enable/disable automatic route addition. Valid values: `phase1` `enable` `disable` .
* `auto_discovery_forwarder` - Enable/disable forwarding short-cut messages. Valid values: `phase1` `enable` `disable` .
* `auto_discovery_sender` - Enable/disable sending short-cut messages. Valid values: `phase1` `enable` `disable` .
* `auto_negotiate` - Enable/disable IPsec SA auto-negotiation. Valid values: `enable` `disable` .
* `comments` - Comment.
* `dhcp_ipsec` - Enable/disable DHCP-IPsec. Valid values: `enable` `disable` .
* `dhgrp` - Phase2 DH group. Valid values: `1` `2` `5` `14` `15` `16` `17` `18` `19` `20` `21` `27` `28` `29` `30` `31` `32` .
* `diffserv` - Enable/disable applying DSCP value to the IPsec tunnel outer IP header. Valid values: `enable` `disable` .
* `diffservcode` - DSCP value to be applied to the IPsec tunnel outer IP header.
* `dst_addr_type` - Remote proxy ID type. Valid values: `subnet` `range` `ip` `name` `subnet6` `range6` `ip6` `name6` .
* `dst_end_ip` - Remote proxy ID IPv4 end.
* `dst_end_ip6` - Remote proxy ID IPv6 end.
* `dst_name` - Remote proxy ID name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `dst_name6` - Remote proxy ID name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `dst_port` - Quick mode destination port (1 - 65535 or 0 for all).
* `dst_start_ip` - Remote proxy ID IPv4 start.
* `dst_start_ip6` - Remote proxy ID IPv6 start.
* `dst_subnet` - Remote proxy ID IPv4 subnet.
* `dst_subnet6` - Remote proxy ID IPv6 subnet.
* `encapsulation` - ESP encapsulation mode. Valid values: `tunnel-mode` `transport-mode` .
* `initiator_ts_narrow` - Enable/disable traffic selector narrowing for IKEv2 initiator. Valid values: `enable` `disable` .
* `ipv4_df` - Enable/disable setting and resetting of IPv4 'Don't Fragment' bit. Valid values: `enable` `disable` .
* `keepalive` - Enable/disable keep alive. Valid values: `enable` `disable` .
* `keylife_type` - Keylife type. Valid values: `seconds` `kbs` `both` .
* `keylifekbs` - Phase2 key life in number of kilobytes of traffic (5120 - 4294967295).
* `keylifeseconds` - Phase2 key life in time in seconds (120 - 172800).
* `l2tp` - Enable/disable L2TP over IPsec. Valid values: `enable` `disable` .
* `name` - IPsec tunnel name.
* `pfs` - Enable/disable PFS feature. Valid values: `enable` `disable` .
* `phase1name` - Phase 1 determines the options required for phase 2. This attribute must reference one of the following datasources: `vpn.ipsec.phase1-interface.name` .
* `proposal` - Phase2 proposal. Valid values: `null-md5` `null-sha1` `null-sha256` `null-sha384` `null-sha512` `des-null` `des-md5` `des-sha1` `des-sha256` `des-sha384` `des-sha512` `3des-null` `3des-md5` `3des-sha1` `3des-sha256` `3des-sha384` `3des-sha512` `aes128-null` `aes128-md5` `aes128-sha1` `aes128-sha256` `aes128-sha384` `aes128-sha512` `aes128gcm` `aes192-null` `aes192-md5` `aes192-sha1` `aes192-sha256` `aes192-sha384` `aes192-sha512` `aes256-null` `aes256-md5` `aes256-sha1` `aes256-sha256` `aes256-sha384` `aes256-sha512` `aes256gcm` `chacha20poly1305` `aria128-null` `aria128-md5` `aria128-sha1` `aria128-sha256` `aria128-sha384` `aria128-sha512` `aria192-null` `aria192-md5` `aria192-sha1` `aria192-sha256` `aria192-sha384` `aria192-sha512` `aria256-null` `aria256-md5` `aria256-sha1` `aria256-sha256` `aria256-sha384` `aria256-sha512` `seed-null` `seed-md5` `seed-sha1` `seed-sha256` `seed-sha384` `seed-sha512` .
* `protocol` - Quick mode protocol selector (1 - 255 or 0 for all).
* `replay` - Enable/disable replay detection. Valid values: `enable` `disable` .
* `route_overlap` - Action for overlapping routes. Valid values: `use-old` `use-new` `allow` .
* `single_source` - Enable/disable single source IP restriction. Valid values: `enable` `disable` .
* `src_addr_type` - Local proxy ID type. Valid values: `subnet` `range` `ip` `name` `subnet6` `range6` `ip6` `name6` .
* `src_end_ip` - Local proxy ID end.
* `src_end_ip6` - Local proxy ID IPv6 end.
* `src_name` - Local proxy ID name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `src_name6` - Local proxy ID name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `src_port` - Quick mode source port (1 - 65535 or 0 for all).
* `src_start_ip` - Local proxy ID start.
* `src_start_ip6` - Local proxy ID IPv6 start.
* `src_subnet` - Local proxy ID subnet.
* `src_subnet6` - Local proxy ID IPv6 subnet.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_vpnipsec_phase2interface can be imported using:
```sh
terraform import fortios_vpnipsec_phase2interface.labelname {{mkey}}
```
