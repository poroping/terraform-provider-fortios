---
subcategory: "FortiGate Ips"
layout: "fortios"
page_title: "FortiOS: fortios_ips_viewmap"
description: |-
  Configure IPS view-map.
---

## fortios_ips_viewmap
Configure IPS view-map.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.

* `id` - View ID.
* `id_policy_id` - ID-based policy ID.
* `policy_id` - Policy ID.
* `vdom_id` - VDOM ID.
* `which` - Policy. Valid values: `firewall` `interface` `interface6` `sniffer` `sniffer6` `explicit` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_ips_viewmap can be imported using:
```sh
terraform import fortios_ips_viewmap.labelname {{mkey}}
```
