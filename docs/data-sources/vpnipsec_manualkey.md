---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnipsec_manualkey"
description: |-
  Get information on a fortios Configure IPsec manual keys.
---

# Data Source: fortios_vpnipsec_manualkey
Use this data source to get information on a fortios Configure IPsec manual keys.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) IPsec tunnel name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `authentication` - Authentication algorithm. Must be the same for both ends of the tunnel.
* `authkey` - Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
* `enckey` - Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
* `encryption` - Encryption algorithm. Must be the same for both ends of the tunnel.
* `interface` - Name of the physical, aggregate, or VLAN interface.
* `local_gw` - Local gateway.
* `localspi` - Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
* `name` - IPsec tunnel name.
* `npu_offload` - Enable/disable NPU offloading.
* `remote_gw` - Peer gateway.
* `remotespi` - Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
