---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_localinpolicy"
description: |-
  Configure user defined IPv4 local-in policies.
---

## fortios_firewall_localinpolicy
Configure user defined IPv4 local-in policies.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `policyid` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `action` - Action performed on traffic matching the policy (default = deny). Valid values: `accept` `deny` .
* `comments` - Comment.
* `dstaddr_negate` - When enabled dstaddr specifies what the destination address must NOT be. Valid values: `enable` `disable` .
* `ha_mgmt_intf_only` - Enable/disable dedicating the HA management interface only for local-in policy. Valid values: `enable` `disable` .
* `intf` - Incoming interface name from available options. This attribute must reference one of the following datasources: `system.zone.name` `system.interface.name` .
* `policyid` - User defined local in policy ID.
* `schedule` - Schedule object from available options. This attribute must reference one of the following datasources: `firewall.schedule.onetime.name` `firewall.schedule.recurring.name` `firewall.schedule.group.name` .
* `service_negate` - When enabled service specifies what the service must NOT be. Valid values: `enable` `disable` .
* `srcaddr_negate` - When enabled srcaddr specifies what the source address must NOT be. Valid values: `enable` `disable` .
* `status` - Enable/disable this local-in policy. Valid values: `enable` `disable` .
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dstaddr` - Destination address object from available options. The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `service` - Service object from available options. The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service name. This attribute must reference one of the following datasources: `firewall.service.custom.name` `firewall.service.group.name` .
* `srcaddr` - Source address object from available options. The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_localinpolicy can be imported using:
```sh
terraform import fortios_firewall_localinpolicy.labelname {{mkey}}
```
