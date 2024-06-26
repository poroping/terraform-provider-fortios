---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf_network"
description: |-
  OSPF network configuration.
---

## fortios_router_ospf_network
OSPF network configuration.

~> This resource is configuring a child table of the parent resource: `fortios_router_ospf`. If this resource is used the parent resource must NOT modify this child table or state inconsistencies will occur.


## Example Usage

```hcl
resource "fortios_router_ospf" "example" {
  area {
    id = "0.0.0.0"
  }

  lifecycle { # hack to deal with API bug https://github.com/poroping/terraform-provider-fortios/issues/7
    ignore_changes = [
      "redistribute",
      "network",
      "neighbor",
      "ospf_interface",
      "summary_address",
      "distribute_list"
    ]
  }

}

resource "fortios_router_ospf_network" "example" {
  prefix = "10.0.0.0/24"
  area   = fortios_router_ospf.example.area[0].id
}
```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.

* `area` - Attach the network to area.
* `comments` - Comment.
* `id` - Network entry ID.
* `prefix` - Prefix.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_network can be imported using:
```sh
terraform import fortios_router_network.labelname {{mkey}}
```
