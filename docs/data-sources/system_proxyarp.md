---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_proxyarp"
description: |-
  Get information on a fortios Configure proxy-ARP.
---

# Data Source: fortios_system_proxyarp
Use this data source to get information on a fortios Configure proxy-ARP.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Unique integer ID of the entry.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `end_ip` - End IP of IP range to be proxied.
* `id` - Unique integer ID of the entry.
* `interface` - Interface acting proxy-ARP.
* `ip` - IP address or start IP to be proxied.
