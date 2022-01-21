---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomnetflow"
description: |-
  Configure NetFlow per VDOM.
---

## fortios_system_vdomnetflow
Configure NetFlow per VDOM.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `collector_ip` - NetFlow collector IP address.
* `collector_port` - NetFlow collector port number.
* `interface` - Specify outgoing interface to reach server. This attribute must reference one of the following datasources: `system.interface.name` .
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `source_ip` - Source IP address for communication with the NetFlow agent.
* `vdom_netflow` - Enable/disable NetFlow per VDOM. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_vdomnetflow can be imported using:
```sh
terraform import fortios_system_vdomnetflow.labelname {{mkey}}
```
