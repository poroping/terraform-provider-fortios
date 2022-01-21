---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnipsec_manualkeyinterface"
description: |-
  Configure IPsec manual keys.
---

## fortios_vpnipsec_manualkeyinterface
Configure IPsec manual keys.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `addr_type` - IP version to use for IP packets. Valid values: `4` `6` .
* `auth_alg` - Authentication algorithm. Must be the same for both ends of the tunnel. Valid values: `null` `md5` `sha1` `sha256` `sha384` `sha512` .
* `auth_key` - Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
* `enc_alg` - Encryption algorithm. Must be the same for both ends of the tunnel. Valid values: `null` `des` `3des` `aes128` `aes192` `aes256` `aria128` `aria192` `aria256` `seed` .
* `enc_key` - Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
* `interface` - Name of the physical, aggregate, or VLAN interface. This attribute must reference one of the following datasources: `system.interface.name` .
* `ip_version` - IP version to use for VPN interface. Valid values: `4` `6` .
* `local_gw` - IPv4 address of the local gateway's external interface.
* `local_gw6` - Local IPv6 address of VPN gateway.
* `local_spi` - Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
* `name` - IPsec tunnel name.
* `npu_offload` - Enable/disable offloading IPsec VPN manual key sessions to NPUs. Valid values: `enable` `disable` .
* `remote_gw` - IPv4 address of the remote gateway's external interface.
* `remote_gw6` - Remote IPv6 address of VPN gateway.
* `remote_spi` - Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_vpnipsec_manualkeyinterface can be imported using:
```sh
terraform import fortios_vpnipsec_manualkeyinterface.labelname {{mkey}}
```
