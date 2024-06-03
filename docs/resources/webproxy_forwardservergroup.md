---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_forwardservergroup"
description: |-
  Configure a forward server group consisting or multiple forward servers. Supports failover and load balancing.
---

## fortios_webproxy_forwardservergroup
Configure a forward server group consisting or multiple forward servers. Supports failover and load balancing.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `affinity` - Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global). Valid values: `enable` `disable` .
* `group_down_option` - Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination. Valid values: `block` `pass` .
* `ldb_method` - Load balance method: weighted or least-session. Valid values: `weighted` `least-session` `active-passive` .
* `name` - Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
* `server_list` - Add web forward servers to a list to form a server group. Optionally assign weights to each server. The structure of `server_list` block is documented below.

The `server_list` block contains:

* `name` - Forward server name. This attribute must reference one of the following datasources: `web-proxy.forward-server.name` .
* `weight` - Optionally assign a weight of the forwarding server for weighted load balancing (1 - 100, default = 10).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_webproxy_forwardservergroup can be imported using:
```sh
terraform import fortios_webproxy_forwardservergroup.labelname {{mkey}}
```
