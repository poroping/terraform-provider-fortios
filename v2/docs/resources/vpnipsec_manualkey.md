---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnipsec_manualkey"
description: |-
  Configure IPsec manual keys.
---

## fortios_vpnipsec_manualkey
Configure IPsec manual keys.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `authentication` - Authentication algorithm. Must be the same for both ends of the tunnel. Valid values: `null` `md5` `sha1` `sha256` `sha384` `sha512` .
* `authkey` - Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
* `enckey` - Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
* `encryption` - Encryption algorithm. Must be the same for both ends of the tunnel. Valid values: `null` `des` `3des` `aes128` `aes192` `aes256` `aria128` `aria192` `aria256` `seed` .
* `interface` - Name of the physical, aggregate, or VLAN interface. This attribute must reference one of the following datasources: `system.interface.name` .
* `local_gw` - Local gateway.
* `localspi` - Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
* `name` - IPsec tunnel name.
* `npu_offload` - Enable/disable NPU offloading. Valid values: `enable` `disable` .
* `remote_gw` - Peer gateway.
* `remotespi` - Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_vpnipsec_manualkey can be imported using:
```sh
terraform import fortios_vpnipsec_manualkey.labelname {{mkey}}
```
