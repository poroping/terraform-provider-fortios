---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnipsec_concentrator"
description: |-
  Concentrator configuration.
---

## fortios_vpnipsec_concentrator
Concentrator configuration.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `id` - Concentrator ID (1 - 65535).
* `name` - Concentrator name.
* `src_check` - Enable to check source address of phase 2 selector. Disable to check only the destination selector. Valid values: `disable` `enable` .
* `member` - Names of up to 3 VPN tunnels to add to the concentrator. The structure of `member` block is documented below.

The `member` block contains:

* `name` - Member name. This attribute must reference one of the following datasources: `vpn.ipsec.manualkey.name` `vpn.ipsec.phase1.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_vpnipsec_concentrator can be imported using:
```sh
terraform import fortios_vpnipsec_concentrator.labelname {{mkey}}
```
