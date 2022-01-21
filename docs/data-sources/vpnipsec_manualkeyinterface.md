---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnipsec_manualkeyinterface"
description: |-
  Get information on a fortios Configure IPsec manual keys.
---

# Data Source: fortios_vpnipsec_manualkeyinterface
Use this data source to get information on a fortios Configure IPsec manual keys.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) IPsec tunnel name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `addr_type` - IP version to use for IP packets.
* `auth_alg` - Authentication algorithm. Must be the same for both ends of the tunnel.
* `auth_key` - Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
* `enc_alg` - Encryption algorithm. Must be the same for both ends of the tunnel.
* `enc_key` - Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
* `interface` - Name of the physical, aggregate, or VLAN interface.
* `ip_version` - IP version to use for VPN interface.
* `local_gw` - IPv4 address of the local gateway's external interface.
* `local_gw6` - Local IPv6 address of VPN gateway.
* `local_spi` - Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
* `name` - IPsec tunnel name.
* `npu_offload` - Enable/disable offloading IPsec VPN manual key sessions to NPUs.
* `remote_gw` - IPv4 address of the remote gateway's external interface.
* `remote_gw6` - Remote IPv6 address of VPN gateway.
* `remote_spi` - Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
