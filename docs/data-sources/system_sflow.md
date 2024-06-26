---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sflow"
description: |-
  Get information on a fortios Configure sFlow.
---

# Data Source: fortios_system_sflow
Use this data source to get information on a fortios Configure sFlow.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `collector_ip` - IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
* `collector_port` - UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `source_ip` - Source IP address for sFlow agent.
