---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_h2qposuprovidernai"
description: |-
  Get information on a fortios Configure online sign up (OSU) provider NAI list.
---

# Data Source: fortios_wirelesscontrollerhotspot20_h2qposuprovidernai
Use this data source to get information on a fortios Configure online sign up (OSU) provider NAI list.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) OSU provider NAI ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - OSU provider NAI ID.
* `nai_list` - OSU NAI list.The structure of `nai_list` block is documented below.

The `nai_list` block contains:

* `name` - OSU NAI ID.
* `osu_nai` - OSU NAI.
