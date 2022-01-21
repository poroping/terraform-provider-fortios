---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpn_ocvpn"
description: |-
  Get information on a fortios Configure Overlay Controller VPN settings.
---

# Data Source: fortios_vpn_ocvpn
Use this data source to get information on a fortios Configure Overlay Controller VPN settings.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `auto_discovery` - Enable/disable auto-discovery shortcuts.
* `auto_discovery_shortcut_mode` - Control deletion of child short-cut tunnels when the parent tunnel goes down.
* `eap` - Enable/disable EAP client authentication.
* `eap_users` - EAP authentication user group.
* `ip_allocation_block` - Class B subnet reserved for private IP address assignment.
* `multipath` - Enable/disable multipath redundancy.
* `nat` - Enable/disable NAT support.
* `poll_interval` - Overlay Controller VPN polling interval.
* `role` - Set device role.
* `sdwan` - Enable/disable adding OCVPN tunnels to SD-WAN.
* `sdwan_zone` - Set SD-WAN zone.
* `status` - Enable/disable Overlay Controller cloud assisted VPN.
* `forticlient_access` - Configure FortiClient settings.The structure of `forticlient_access` block is documented below.

The `forticlient_access` block contains:

* `psksecret` - Pre-shared secret for FortiClient PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).
* `status` - Enable/disable FortiClient to access OCVPN networks.
* `auth_groups` - FortiClient user authentication groups.The structure of `auth_groups` block is documented below.

The `auth_groups` block contains:

* `auth_group` - Authentication user group for FortiClient access.
* `name` - Group name.
* `overlays` - OCVPN overlays to allow access to.The structure of `overlays` block is documented below.

The `overlays` block contains:

* `overlay_name` - Overlay name.
* `overlays` - Network overlays to register with Overlay Controller VPN service.The structure of `overlays` block is documented below.

The `overlays` block contains:

* `assign_ip` - Enable/disable mode-cfg address assignment.
* `id` - ID.
* `inter_overlay` - Allow or deny traffic from other overlays.
* `ipv4_end_ip` - End of IPv4 range.
* `ipv4_start_ip` - Start of IPv4 range.
* `name` - Overlay name.
* `overlay_name` - Overlay name.
* `subnets` - Internal subnets to register with OCVPN service.The structure of `subnets` block is documented below.

The `subnets` block contains:

* `id` - ID.
* `interface` - LAN interface.
* `subnet` - IPv4 address and subnet mask.
* `type` - Subnet type.
* `wan_interface` - FortiGate WAN interfaces to use with OCVPN.The structure of `wan_interface` block is documented below.

The `wan_interface` block contains:

* `name` - Interface name.
