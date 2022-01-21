---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf_neighbor"
description: |-
  OSPF neighbor configuration are used when OSPF runs on non-broadcast media
---

## fortios_router_ospf_neighbor
OSPF neighbor configuration are used when OSPF runs on non-broadcast media

~> This resource is configuring a child table of the parent resource: `fortios_router_ospf`. If this resource is used the parent resource must NOT modify this child table or state inconsistencies will occur.


## Example Usage

```hcl
resource "fortios_system_interface" "example" {
  vdom = "BUTT"
  type = "loopback"
  name = "foobar123"
  ip   = "5.5.5.0/31"
}

resource "fortios_router_ospf_ospf_interface" "example" {
  name         = fortios_system_interface.example.name
  interface    = fortios_system_interface.example.name
  network_type = "non-broadcast"
}

resource "fortios_router_ospf_neighbor" "example" {
  ip = "5.5.5.1"

  depends_on = [
    fortios_router_ospf_ospf_interface.example
  ]
}
```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.

* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `id` - Neighbor entry ID.
* `ip` - Interface IP address of the neighbor.
* `poll_interval` - Poll interval time in seconds.
* `priority` - Priority.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_neighbor can be imported using:
```sh
terraform import fortios_router_neighbor.labelname {{mkey}}
```
