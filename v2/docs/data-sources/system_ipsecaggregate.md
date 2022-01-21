---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ipsecaggregate"
description: |-
  Get information on a fortios Configure an aggregate of IPsec tunnels.
---

# Data Source: fortios_system_ipsecaggregate
Use this data source to get information on a fortios Configure an aggregate of IPsec tunnels.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) IPsec aggregate name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `algorithm` - Frame distribution algorithm.
* `name` - IPsec aggregate name.
* `member` - Member tunnels of the aggregate.The structure of `member` block is documented below.

The `member` block contains:

* `tunnel_name` - Tunnel name.
