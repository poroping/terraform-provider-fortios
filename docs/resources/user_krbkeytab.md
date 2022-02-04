---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_krbkeytab"
description: |-
  Configure Kerberos keytab entries.
---

## fortios_user_krbkeytab
Configure Kerberos keytab entries.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `keytab` - Base64 coded keytab file containing a pre-shared key.
* `name` - Kerberos keytab entry name.
* `pac_data` - Enable/disable parsing PAC data in the ticket. Valid values: `enable` `disable` .
* `principal` - Kerberos service principal. For example, HTTP/myfgt.example.com@example.com.
* `ldap_server` - LDAP server name(s). The structure of `ldap_server` block is documented below.

The `ldap_server` block contains:

* `name` - LDAP server name. This attribute must reference one of the following datasources: `user.ldap.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_user_krbkeytab can be imported using:
```sh
terraform import fortios_user_krbkeytab.labelname {{mkey}}
```
