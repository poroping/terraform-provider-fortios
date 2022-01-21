---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_dnsserver"
description: |-
  Get information on a fortios Configure DNS servers.
---

# Data Source: fortios_system_dnsserver
Use this data source to get information on a fortios Configure DNS servers.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) DNS server name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `dnsfilter_profile` - DNS filter profile.
* `doh` - DNS over HTTPS.
* `mode` - DNS server mode.
* `name` - DNS server name.
