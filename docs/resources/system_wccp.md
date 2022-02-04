---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_wccp"
description: |-
  Configure WCCP.
---

## fortios_system_wccp
Configure WCCP.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `service_id` to be defined.

* `assignment_bucket_format` - Assignment bucket format for the WCCP cache engine. Valid values: `wccp-v2` `cisco-implementation` .
* `assignment_dstaddr_mask` - Assignment destination address mask.
* `assignment_method` - Hash key assignment preference. Valid values: `HASH` `MASK` `any` .
* `assignment_srcaddr_mask` - Assignment source address mask.
* `assignment_weight` - Assignment of hash weight/ratio for the WCCP cache engine.
* `authentication` - Enable/disable MD5 authentication. Valid values: `enable` `disable` .
* `cache_engine_method` - Method used to forward traffic to the routers or to return to the cache engine. Valid values: `GRE` `L2` .
* `cache_id` - IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
* `forward_method` - Method used to forward traffic to the cache servers. Valid values: `GRE` `L2` `any` .
* `group_address` - IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
* `password` - Password for MD5 authentication.
* `ports` - Service ports.
* `ports_defined` - Match method. Valid values: `source` `destination` .
* `primary_hash` - Hash method. Valid values: `src-ip` `dst-ip` `src-port` `dst-port` .
* `priority` - Service priority.
* `protocol` - Service protocol.
* `return_method` - Method used to decline a redirected packet and return it to the FortiGate unit. Valid values: `GRE` `L2` `any` .
* `router_id` - IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
* `router_list` - IP addresses of one or more WCCP routers.
* `server_list` - IP addresses and netmasks for up to four cache servers.
* `server_type` - Cache server type. Valid values: `forward` `proxy` .
* `service_id` - Service ID.
* `service_type` - WCCP service type used by the cache server for logical interception and redirection of traffic. Valid values: `auto` `standard` `dynamic` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_wccp can be imported using:
```sh
terraform import fortios_system_wccp.labelname {{mkey}}
```
