---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnipsec_forticlient"
description: |-
  Configure FortiClient policy realm.
---

## fortios_vpnipsec_forticlient
Configure FortiClient policy realm.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `realm` to be defined.

* `phase2name` - Phase 2 tunnel name that you defined in the FortiClient dialup configuration. This attribute must reference one of the following datasources: `vpn.ipsec.phase2.name` `vpn.ipsec.phase2-interface.name` .
* `realm` - FortiClient realm name.
* `status` - Enable/disable this FortiClient configuration. Valid values: `enable` `disable` .
* `usergroupname` - User group name for FortiClient users. This attribute must reference one of the following datasources: `user.group.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_vpnipsec_forticlient can be imported using:
```sh
terraform import fortios_vpnipsec_forticlient.labelname {{mkey}}
```
