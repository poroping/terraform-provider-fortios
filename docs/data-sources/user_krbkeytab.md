---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_krbkeytab"
description: |-
  Get information on a fortios Configure Kerberos keytab entries.
---

# Data Source: fortios_user_krbkeytab
Use this data source to get information on a fortios Configure Kerberos keytab entries.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Kerberos keytab entry name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `keytab` - base64 coded keytab file containing a pre-shared key.
* `name` - Kerberos keytab entry name.
* `pac_data` - Enable/disable parsing PAC data in the ticket.
* `principal` - Kerberos service principal, e.g. HTTP/fgt.example.com@EXAMPLE.COM.
* `ldap_server` - LDAP server name(s).The structure of `ldap_server` block is documented below.

The `ldap_server` block contains:

* `name` - LDAP server name.
