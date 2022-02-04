---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemautoupdate_tunneling"
description: |-
  Configure web proxy tunneling for the FDN.
---

## fortios_systemautoupdate_tunneling
Configure web proxy tunneling for the FDN.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `address` - Web proxy IP address or FQDN.
* `password` - Web proxy password.
* `port` - Web proxy port.
* `status` - Enable/disable web proxy tunneling. Valid values: `enable` `disable` .
* `username` - Web proxy username.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_systemautoupdate_tunneling can be imported using:
```sh
terraform import fortios_systemautoupdate_tunneling.labelname {{mkey}}
```
