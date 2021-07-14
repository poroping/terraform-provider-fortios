---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_access_proxy_virtual_host"
description: |-
  Configure Access Proxy virtual hosts.
---

## fortios_firewall_access_proxy_virtual_host
Configure Access Proxy virtual hosts.

## Example Usage

WIP

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `host` - The host name.
* `host_type` - Type of host pattern. Valid values: `sub-string` `wildcard` .
* `name` - Virtual host name.
* `ssl_certificate` - SSL certificate for this host.This attribute must reference one of the following datasources: `vpn.certificate.local.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Address can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_address.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
