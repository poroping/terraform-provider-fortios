---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetserviceappend"
description: |-
  Get information on a fortios Configure additional port mappings for Internet Services.
---

# Data Source: fortios_firewall_internetserviceappend
Use this data source to get information on a fortios Configure additional port mappings for Internet Services.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `append_port` - Appending TCP/UDP/SCTP destination port (1 to 65535).
* `match_port` - Matching TCP/UDP/SCTP destination port (0 to 65535, 0 means any port).
