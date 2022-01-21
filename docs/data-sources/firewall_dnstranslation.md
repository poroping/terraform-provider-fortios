---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_dnstranslation"
description: |-
  Get information on a fortios Configure DNS translation.
---

# Data Source: fortios_firewall_dnstranslation
Use this data source to get information on a fortios Configure DNS translation.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `dst` - IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
* `id` - ID.
* `netmask` - If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
* `src` - IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
