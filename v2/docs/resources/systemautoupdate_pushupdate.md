---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemautoupdate_pushupdate"
description: |-
  Configure push updates.
---

## fortios_systemautoupdate_pushupdate
Configure push updates.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `address` - IPv4 or IPv6 address used by FortiGuard servers to send push updates to this FortiGate.
* `override` - Enable/disable push update override server. Valid values: `enable` `disable` .
* `port` - Push update override port. (Do not overlap with other service ports)
* `status` - Enable/disable push updates. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_systemautoupdate_pushupdate can be imported using:
```sh
terraform import fortios_systemautoupdate_pushupdate.labelname {{mkey}}
```
