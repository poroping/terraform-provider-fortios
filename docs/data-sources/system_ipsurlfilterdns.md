---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ipsurlfilterdns"
description: |-
  Get information on a fortios Configure IPS URL filter DNS servers.
---

# Data Source: fortios_system_ipsurlfilterdns
Use this data source to get information on a fortios Configure IPS URL filter DNS servers.


## Example Usage

```hcl

```

## Argument Reference

* `address` - (Required) DNS server IP address.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `address` - DNS server IP address.
* `ipv6_capability` - Enable/disable this server for IPv6 queries.
* `status` - Enable/disable using this DNS server for IPS URL filter DNS queries.
