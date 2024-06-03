---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_authportal"
description: |-
  Configure firewall authentication portals.
---

## fortios_firewall_authportal
Configure firewall authentication portals.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `identity_based_route` - Name of the identity-based route that applies to this portal. This attribute must reference one of the following datasources: `firewall.identity-based-route.name` .
* `portal_addr` - Address (or FQDN) of the authentication portal.
* `portal_addr6` - IPv6 address (or FQDN) of authentication portal.
* `proxy_auth` - Enable/disable authentication by proxy daemon (default = disable). Valid values: `enable` `disable` .
* `groups` - Firewall user groups permitted to authenticate through this portal. Separate group names with spaces. The structure of `groups` block is documented below.

The `groups` block contains:

* `name` - Group name. This attribute must reference one of the following datasources: `user.group.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_authportal can be imported using:
```sh
terraform import fortios_firewall_authportal.labelname {{mkey}}
```
