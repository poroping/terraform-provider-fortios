---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_accessproxyvirtualhost"
description: |-
  Configure Access Proxy virtual hosts.
---

## fortios_firewall_accessproxyvirtualhost
Configure Access Proxy virtual hosts.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `host` - The host name.
* `host_type` - Type of host pattern. Valid values: `sub-string` `wildcard` .
* `name` - Virtual host name.
* `replacemsg_group` - Access-proxy-virtual-host replacement message override group. This attribute must reference one of the following datasources: `system.replacemsg-group.name` .
* `ssl_certificate` - SSL certificate for this host. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_accessproxyvirtualhost can be imported using:
```sh
terraform import fortios_firewall_accessproxyvirtualhost.labelname {{mkey}}
```
