---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnipsec_concentrator"
description: |-
  Get information on a fortios Concentrator configuration.
---

# Data Source: fortios_vpnipsec_concentrator
Use this data source to get information on a fortios Concentrator configuration.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Concentrator ID. (1-65535)
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `id` - Concentrator ID. (1-65535)
* `name` - Concentrator name.
* `src_check` - Enable to check source address of phase 2 selector. Disable to check only the destination selector.
* `member` - Names of up to 3 VPN tunnels to add to the concentrator.The structure of `member` block is documented below.

The `member` block contains:

* `name` - Member name.
