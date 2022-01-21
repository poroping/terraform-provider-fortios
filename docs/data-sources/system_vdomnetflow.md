---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomnetflow"
description: |-
  Get information on a fortios Configure NetFlow per VDOM.
---

# Data Source: fortios_system_vdomnetflow
Use this data source to get information on a fortios Configure NetFlow per VDOM.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `collector_ip` - NetFlow collector IP address.
* `collector_port` - NetFlow collector port number.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `source_ip` - Source IP address for communication with the NetFlow agent.
* `vdom_netflow` - Enable/disable NetFlow per VDOM.
