---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_forwardservergroup"
description: |-
  Get information on a fortios Configure a forward server group consisting or multiple forward servers. Supports failover and load balancing.
---

# Data Source: fortios_webproxy_forwardservergroup
Use this data source to get information on a fortios Configure a forward server group consisting or multiple forward servers. Supports failover and load balancing.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `affinity` - Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global).
* `group_down_option` - Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination.
* `ldb_method` - Load balance method: weighted or least-session.
* `name` - Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
* `server_list` - Add web forward servers to a list to form a server group. Optionally assign weights to each server.The structure of `server_list` block is documented below.

The `server_list` block contains:

* `name` - Forward server name.
* `weight` - Optionally assign a weight of the forwarding server for weighted load balancing (1 - 100, default = 10)
