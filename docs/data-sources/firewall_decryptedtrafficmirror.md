---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_decryptedtrafficmirror"
description: |-
  Get information on a fortios Configure decrypted traffic mirror.
---

# Data Source: fortios_firewall_decryptedtrafficmirror
Use this data source to get information on a fortios Configure decrypted traffic mirror.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `dstmac` - Set destination MAC address for mirrored traffic.
* `name` - Name.
* `traffic_source` - Source of decrypted traffic to be mirrored.
* `traffic_type` - Types of decrypted traffic to be mirrored.
* `interface` - Decrypted traffic mirror interface.The structure of `interface` block is documented below.

The `interface` block contains:

* `name` - Decrypted traffic mirror interface.
