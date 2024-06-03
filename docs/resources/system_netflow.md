---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_netflow"
description: |-
  Configure NetFlow.
---

## fortios_system_netflow
Configure NetFlow.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `active_flow_timeout` - Timeout to report active flows (60 - 3600 sec, default = 1800).
* `collector_ip` - Collector IP.
* `collector_port` - NetFlow collector port number.
* `inactive_flow_timeout` - Timeout for periodic report of finished flows (10 - 600 sec, default = 15).
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `source_ip` - Source IP address for communication with the NetFlow agent.
* `template_tx_counter` - Counter of flowset records before resending a template flowset record.
* `template_tx_timeout` - Timeout for periodic template flowset transmission (60 - 86400 sec, default = 1800).
* `collectors` - Netflow collectors. The structure of `collectors` block is documented below.

The `collectors` block contains:

* `collector_ip` - Collector IP.
* `collector_port` - NetFlow collector port number.
* `id` - ID.
* `interface` - Specify outgoing interface to reach server. This attribute must reference one of the following datasources: `system.interface.name` .
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `source_ip` - Source IP address for communication with the NetFlow agent.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_netflow can be imported using:
```sh
terraform import fortios_system_netflow.labelname {{mkey}}
```
