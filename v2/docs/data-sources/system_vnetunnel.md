---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vnetunnel"
description: |-
  Get information on a fortios Configure virtual network enabler tunnel.
---

# Data Source: fortios_system_vnetunnel
Use this data source to get information on a fortios Configure virtual network enabler tunnel.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `bmr_hostname` - BMR hostname.
* `br` - Border relay IPv6 address.
* `interface` - Interface name.
* `ipv4_address` - Tunnel IPv4 address and netmask.
* `mode` - VNE tunnel mode.
* `ssl_certificate` - Name of local certificate for SSL connections.
* `status` - Enable/disable VNE tunnel.
* `update_url` - URL of provisioning server.
