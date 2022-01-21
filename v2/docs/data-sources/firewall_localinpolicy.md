---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_localinpolicy"
description: |-
  Get information on a fortios Configure user defined IPv4 local-in policies.
---

# Data Source: fortios_firewall_localinpolicy
Use this data source to get information on a fortios Configure user defined IPv4 local-in policies.


## Example Usage

```hcl

```

## Argument Reference

* `policyid` - (Required) User defined local in policy ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `action` - Action performed on traffic matching the policy (default = deny).
* `comments` - Comment.
* `dstaddr_negate` - When enabled dstaddr specifies what the destination address must NOT be.
* `ha_mgmt_intf_only` - Enable/disable dedicating the HA management interface only for local-in policy.
* `intf` - Incoming interface name from available options.
* `policyid` - User defined local in policy ID.
* `schedule` - Schedule object from available options.
* `service_negate` - When enabled service specifies what the service must NOT be.
* `srcaddr_negate` - When enabled srcaddr specifies what the source address must NOT be.
* `status` - Enable/disable this local-in policy.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dstaddr` - Destination address object from available options.The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name.
* `service` - Service object from available options.The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service name.
* `srcaddr` - Source address object from available options.The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name.
