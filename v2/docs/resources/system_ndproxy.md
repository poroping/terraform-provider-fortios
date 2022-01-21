---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ndproxy"
description: |-
  Configure IPv6 neighbor discovery proxy (RFC4389).
---

## fortios_system_ndproxy
Configure IPv6 neighbor discovery proxy (RFC4389).

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `status` - Enable/disable neighbor discovery proxy. Valid values: `enable` `disable` .
* `member` - Interfaces using the neighbor discovery proxy. The structure of `member` block is documented below.

The `member` block contains:

* `interface_name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_ndproxy can be imported using:
```sh
terraform import fortios_system_ndproxy.labelname {{mkey}}
```
