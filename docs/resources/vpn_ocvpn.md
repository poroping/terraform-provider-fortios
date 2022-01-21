---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpn_ocvpn"
description: |-
  Configure Overlay Controller VPN settings.
---

## fortios_vpn_ocvpn
Configure Overlay Controller VPN settings.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `auto_discovery` - Enable/disable auto-discovery shortcuts. Valid values: `enable` `disable` .
* `auto_discovery_shortcut_mode` - Control deletion of child short-cut tunnels when the parent tunnel goes down. Valid values: `independent` `dependent` .
* `eap` - Enable/disable EAP client authentication. Valid values: `enable` `disable` .
* `eap_users` - EAP authentication user group. This attribute must reference one of the following datasources: `user.group.name` .
* `ip_allocation_block` - Class B subnet reserved for private IP address assignment.
* `multipath` - Enable/disable multipath redundancy. Valid values: `enable` `disable` .
* `nat` - Enable/disable NAT support. Valid values: `enable` `disable` .
* `poll_interval` - Overlay Controller VPN polling interval.
* `role` - Set device role. Valid values: `spoke` `primary-hub` `secondary-hub` .
* `sdwan` - Enable/disable adding OCVPN tunnels to SD-WAN. Valid values: `enable` `disable` .
* `sdwan_zone` - Set SD-WAN zone. This attribute must reference one of the following datasources: `system.sdwan.zone.name` .
* `status` - Enable/disable Overlay Controller cloud assisted VPN. Valid values: `enable` `disable` .
* `forticlient_access` - Configure FortiClient settings. The structure of `forticlient_access` block is documented below.

The `forticlient_access` block contains:

* `psksecret` - Pre-shared secret for FortiClient PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).
* `status` - Enable/disable FortiClient to access OCVPN networks. Valid values: `enable` `disable` .
* `auth_groups` - FortiClient user authentication groups. The structure of `auth_groups` block is documented below.

The `auth_groups` block contains:

* `auth_group` - Authentication user group for FortiClient access. This attribute must reference one of the following datasources: `user.group.name` .
* `name` - Group name.
* `overlays` - OCVPN overlays to allow access to. The structure of `overlays` block is documented below.

The `overlays` block contains:

* `overlay_name` - Overlay name. This attribute must reference one of the following datasources: `vpn.ocvpn.overlays.overlay-name` .
* `overlays` - Network overlays to register with Overlay Controller VPN service. The structure of `overlays` block is documented below.

The `overlays` block contains:

* `assign_ip` - Enable/disable mode-cfg address assignment. Valid values: `enable` `disable` .
* `id` - ID.
* `inter_overlay` - Allow or deny traffic from other overlays. Valid values: `allow` `deny` .
* `ipv4_end_ip` - End of IPv4 range.
* `ipv4_start_ip` - Start of IPv4 range.
* `name` - Overlay name.
* `overlay_name` - Overlay name.
* `subnets` - Internal subnets to register with OCVPN service. The structure of `subnets` block is documented below.

The `subnets` block contains:

* `id` - ID.
* `interface` - LAN interface. This attribute must reference one of the following datasources: `system.interface.name` .
* `subnet` - IPv4 address and subnet mask.
* `type` - Subnet type. Valid values: `subnet` `interface` .
* `wan_interface` - FortiGate WAN interfaces to use with OCVPN. The structure of `wan_interface` block is documented below.

The `wan_interface` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_vpn_ocvpn can be imported using:
```sh
terraform import fortios_vpn_ocvpn.labelname {{mkey}}
```
