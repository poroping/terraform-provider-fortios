---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ipsurlfilterdns6"
description: |-
  Configure IPS URL filter IPv6 DNS servers.
---

## fortios_system_ipsurlfilterdns6
Configure IPS URL filter IPv6 DNS servers.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `address6` to be defined.

* `address6` - IPv6 address of DNS server.
* `status` - Enable/disable this server for IPv6 DNS queries. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_ipsurlfilterdns6 can be imported using:
```sh
terraform import fortios_system_ipsurlfilterdns6.labelname {{mkey}}
```
