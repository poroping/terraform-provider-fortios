---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallipmacbinding_table"
description: |-
  Configure IP to MAC address pairs in the IP/MAC binding table.
---

## fortios_firewallipmacbinding_table
Configure IP to MAC address pairs in the IP/MAC binding table.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `seq_num` to be defined.

* `ip` - IPv4 address portion of the pair (format: xxx.xxx.xxx.xxx).
* `mac` - MAC address portion of the pair (format: xx:xx:xx:xx:xx:xx in hexidecimal).
* `name` - Name of the pair (optional, default = no name).
* `seq_num` - Entry number.
* `status` - Enable/disable this IP-mac binding pair. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewallipmacbinding_table can be imported using:
```sh
terraform import fortios_firewallipmacbinding_table.labelname {{mkey}}
```
