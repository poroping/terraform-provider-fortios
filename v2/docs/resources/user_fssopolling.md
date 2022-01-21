---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_fssopolling"
description: |-
  Configure FSSO active directory servers for polling mode.
---

## fortios_user_fssopolling
Configure FSSO active directory servers for polling mode.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `default_domain` - Default domain managed by this Active Directory server.
* `id` - Active Directory server ID.
* `ldap_server` - LDAP server name used in LDAP connection strings. This attribute must reference one of the following datasources: `user.ldap.name` .
* `logon_history` - Number of hours of logon history to keep, 0 means keep all history.
* `password` - Password required to log into this Active Directory server
* `polling_frequency` - Polling frequency (every 1 to 30 seconds).
* `port` - Port to communicate with this Active Directory server.
* `server` - Host name or IP address of the Active Directory server.
* `smb_ntlmv1_auth` - Enable/disable support of NTLMv1 for Samba authentication. Valid values: `enable` `disable` .
* `smbv1` - Enable/disable support of SMBv1 for Samba. Valid values: `enable` `disable` .
* `status` - Enable/disable polling for the status of this Active Directory server. Valid values: `enable` `disable` .
* `user` - User name required to log into this Active Directory server.
* `adgrp` - LDAP Group Info. The structure of `adgrp` block is documented below.

The `adgrp` block contains:

* `name` - Name.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_user_fssopolling can be imported using:
```sh
terraform import fortios_user_fssopolling.labelname {{mkey}}
```
