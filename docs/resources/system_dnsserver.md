---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_dnsserver"
description: |-
  Configure DNS servers.
---

## fortios_system_dnsserver
Configure DNS servers.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `dnsfilter_profile` - DNS filter profile. This attribute must reference one of the following datasources: `dnsfilter.profile.name` .
* `doh` - DNS over HTTPS. Valid values: `enable` `disable` .
* `mode` - DNS server mode. Valid values: `recursive` `non-recursive` `forward-only` .
* `name` - DNS server name. This attribute must reference one of the following datasources: `system.interface.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_dnsserver can be imported using:
```sh
terraform import fortios_system_dnsserver.labelname {{mkey}}
```
