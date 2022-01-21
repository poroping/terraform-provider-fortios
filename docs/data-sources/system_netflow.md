---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_netflow"
description: |-
  Get information on a fortios Configure NetFlow.
---

# Data Source: fortios_system_netflow
Use this data source to get information on a fortios Configure NetFlow.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `active_flow_timeout` - Timeout to report active flows (60 - 3600 sec, default = 1800).
* `collector_ip` - Collector IP.
* `collector_port` - NetFlow collector port number.
* `inactive_flow_timeout` - Timeout for periodic report of finished flows (10 - 600 sec, default = 15).
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `source_ip` - Source IP address for communication with the NetFlow agent.
* `template_tx_counter` - Counter of flowset records before resending a template flowset record.
* `template_tx_timeout` - Timeout for periodic template flowset transmission (60 - 86400 sec, default = 1800).
