---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnipsec_forticlient"
description: |-
  Get information on a fortios Configure FortiClient policy realm.
---

# Data Source: fortios_vpnipsec_forticlient
Use this data source to get information on a fortios Configure FortiClient policy realm.


## Example Usage

```hcl

```

## Argument Reference

* `realm` - (Required) FortiClient realm name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `phase2name` - Phase 2 tunnel name that you defined in the FortiClient dialup configuration.
* `realm` - FortiClient realm name.
* `status` - Enable/disable this FortiClient configuration.
* `usergroupname` - User group name for FortiClient users.
