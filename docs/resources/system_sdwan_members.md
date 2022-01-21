---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sdwan_members"
description: |-
  FortiGate interfaces added to the SD-WAN.
---

## fortios_system_sdwan_members
FortiGate interfaces added to the SD-WAN.

~> This resource is configuring a child table of the parent resource: `fortios_system_sdwan`. If this resource is used the parent resource must NOT modify this child table or state inconsistencies will occur.


## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `seq_num` to be defined.

* `comment` - Comments.
* `cost` - Cost of this interface for services in SLA mode (0 - 4294967295, default = 0).
* `gateway` - The default gateway for this interface. Usually the default gateway of the Internet service provider that this interface is connected to.
* `gateway6` - IPv6 gateway.
* `ingress_spillover_threshold` - Ingress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.
* `interface` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `priority` - Priority of the interface for IPv4 (0 - 65535, default = 0). Used for SD-WAN rules or priority rules.
* `priority6` - Priority of the interface for IPv6 (1 - 65535, default = 1024). Used for SD-WAN rules or priority rules.
* `seq_num` - Sequence number(1-512).
* `source` - Source IP address used in the health-check packet to the server.
* `source6` - Source IPv6 address used in the health-check packet to the server.
* `spillover_threshold` - Egress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.
* `status` - Enable/disable this interface in the SD-WAN. Valid values: `disable` `enable` .
* `volume_ratio` - Measured volume ratio (this value / sum of all values = percentage of link volume, 1 - 255).
* `weight` - Weight of this interface for weighted load balancing. (1 - 255) More traffic is directed to interfaces with higher weights.
* `zone` - Zone name. This attribute must reference one of the following datasources: `system.sdwan.zone.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_members can be imported using:
```sh
terraform import fortios_system_members.labelname {{mkey}}
```
