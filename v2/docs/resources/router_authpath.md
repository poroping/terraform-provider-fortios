---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_authpath"
description: |-
  Configure authentication based routing.
---

## fortios_router_authpath
Configure authentication based routing.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `device` - Outgoing interface. This attribute must reference one of the following datasources: `system.interface.name` .
* `gateway` - Gateway IP address.
* `name` - Name of the entry.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_authpath can be imported using:
```sh
terraform import fortios_router_authpath.labelname {{mkey}}
```
