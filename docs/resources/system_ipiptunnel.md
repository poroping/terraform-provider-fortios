---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ipiptunnel"
description: |-
  Configure IP in IP Tunneling.
---

## fortios_system_ipiptunnel
Configure IP in IP Tunneling.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `auto_asic_offload` - Enable/disable tunnel ASIC offloading. Valid values: `enable` `disable` .
* `interface` - Interface name that is associated with the incoming traffic from available options. This attribute must reference one of the following datasources: `system.interface.name` .
* `local_gw` - IPv4 address for the local gateway.
* `name` - IPIP Tunnel name.
* `remote_gw` - IPv4 address for the remote gateway.
* `use_sdwan` - Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable` `enable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_ipiptunnel can be imported using:
```sh
terraform import fortios_system_ipiptunnel.labelname {{mkey}}
```
