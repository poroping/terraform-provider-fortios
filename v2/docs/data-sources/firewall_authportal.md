---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_authportal"
description: |-
  Get information on a fortios Configure firewall authentication portals.
---

# Data Source: fortios_firewall_authportal
Use this data source to get information on a fortios Configure firewall authentication portals.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `identity_based_route` - Name of the identity-based route that applies to this portal.
* `portal_addr` - Address (or FQDN) of the authentication portal.
* `portal_addr6` - IPv6 address (or FQDN) of authentication portal.
* `groups` - Firewall user groups permitted to authenticate through this portal. Separate group names with spaces.The structure of `groups` block is documented below.

The `groups` block contains:

* `name` - Group name.
