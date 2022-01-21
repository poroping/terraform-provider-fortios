---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_accessproxyvirtualhost"
description: |-
  Get information on a fortios Configure Access Proxy virtual hosts.
---

# Data Source: fortios_firewall_accessproxyvirtualhost
Use this data source to get information on a fortios Configure Access Proxy virtual hosts.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Virtual host name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `host` - The host name.
* `host_type` - Type of host pattern.
* `name` - Virtual host name.
* `ssl_certificate` - SSL certificate for this host.
