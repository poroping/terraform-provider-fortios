---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_domaincontroller"
description: |-
  Configure domain controller entries.
---

## fortios_user_domaincontroller
Configure domain controller entries.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `ad_mode` - Set Active Directory mode. Valid values: `none` `ds` `lds` .
* `adlds_dn` - AD LDS distinguished name.
* `adlds_ip_address` - AD LDS IPv4 address.
* `adlds_ip6` - AD LDS IPv6 address.
* `adlds_port` - Port number of AD LDS service (default = 389).
* `change_detection` - Enable/disable detection of a configuration change in the Active Directory server. Valid values: `enable` `disable` .
* `change_detection_period` - Minutes to detect a configuration change in the Active Directory server (5 - 10080 minutes (7 days), default = 60).
* `dns_srv_lookup` - Enable/disable DNS service lookup. Valid values: `enable` `disable` .
* `domain_name` - Domain DNS name.
* `hostname` - Hostname of the server to connect to.
* `interface` - Specify outgoing interface to reach server. This attribute must reference one of the following datasources: `system.interface.name` .
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `ip_address` - Domain controller IPv4 address.
* `ip6` - Domain controller IPv6 address.
* `name` - Domain controller entry name.
* `password` - Password for specified username.
* `port` - Port to be used for communication with the domain controller (default = 445).
* `replication_port` - Port to be used for communication with the domain controller for replication service. Port number 0 indicates automatic discovery.
* `source_ip_address` - FortiGate IPv4 address to be used for communication with the domain controller.
* `source_ip6` - FortiGate IPv6 address to be used for communication with the domain controller.
* `source_port` - Source port to be used for communication with the domain controller.
* `username` - User name to sign in with. Must have proper permissions for service.
* `extra_server` - Extra servers. The structure of `extra_server` block is documented below.

The `extra_server` block contains:

* `id` - Server ID.
* `ip_address` - Domain controller IP address.
* `port` - Port to be used for communication with the domain controller (default = 445).
* `source_ip_address` - FortiGate IPv4 address to be used for communication with the domain controller.
* `source_port` - Source port to be used for communication with the domain controller.
* `ldap_server` - LDAP server name(s). The structure of `ldap_server` block is documented below.

The `ldap_server` block contains:

* `name` - LDAP server name. This attribute must reference one of the following datasources: `user.ldap.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_user_domaincontroller can be imported using:
```sh
terraform import fortios_user_domaincontroller.labelname {{mkey}}
```
