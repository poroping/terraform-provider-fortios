---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_wccp"
description: |-
  Get information on a fortios Configure WCCP.
---

# Data Source: fortios_system_wccp
Use this data source to get information on a fortios Configure WCCP.


## Example Usage

```hcl

```

## Argument Reference

* `service_id` - (Required) Service ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `assignment_bucket_format` - Assignment bucket format for the WCCP cache engine.
* `assignment_dstaddr_mask` - Assignment destination address mask.
* `assignment_method` - Hash key assignment preference.
* `assignment_srcaddr_mask` - Assignment source address mask.
* `assignment_weight` - Assignment of hash weight/ratio for the WCCP cache engine.
* `authentication` - Enable/disable MD5 authentication.
* `cache_engine_method` - Method used to forward traffic to the routers or to return to the cache engine.
* `cache_id` - IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
* `forward_method` - Method used to forward traffic to the cache servers.
* `group_address` - IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
* `password` - Password for MD5 authentication.
* `ports` - Service ports.
* `ports_defined` - Match method.
* `primary_hash` - Hash method.
* `priority` - Service priority.
* `protocol` - Service protocol.
* `return_method` - Method used to decline a redirected packet and return it to the FortiGate unit.
* `router_id` - IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
* `router_list` - IP addresses of one or more WCCP routers.
* `server_list` - IP addresses and netmasks for up to four cache servers.
* `server_type` - Cache server type.
* `service_id` - Service ID.
* `service_type` - WCCP service type used by the cache server for logical interception and redirection of traffic.
