---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ipsurlfilterdns6"
description: |-
  Get information on a fortios Configure IPS URL filter IPv6 DNS servers.
---

# Data Source: fortios_system_ipsurlfilterdns6
Use this data source to get information on a fortios Configure IPS URL filter IPv6 DNS servers.


## Example Usage

```hcl

```

## Argument Reference

* `address6` - (Required) IPv6 address of DNS server.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `address6` - IPv6 address of DNS server.
* `status` - Enable/disable this server for IPv6 DNS queries.
