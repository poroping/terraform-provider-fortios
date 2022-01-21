---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_fssopolling"
description: |-
  Get information on a fortios Configure FSSO active directory servers for polling mode.
---

# Data Source: fortios_user_fssopolling
Use this data source to get information on a fortios Configure FSSO active directory servers for polling mode.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Active Directory server ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `default_domain` - Default domain managed by this Active Directory server.
* `id` - Active Directory server ID.
* `ldap_server` - LDAP server name used in LDAP connection strings.
* `logon_history` - Number of hours of logon history to keep, 0 means keep all history.
* `password` - Password required to log into this Active Directory server
* `polling_frequency` - Polling frequency (every 1 to 30 seconds).
* `port` - Port to communicate with this Active Directory server.
* `server` - Host name or IP address of the Active Directory server.
* `smb_ntlmv1_auth` - Enable/disable support of NTLMv1 for Samba authentication.
* `smbv1` - Enable/disable support of SMBv1 for Samba.
* `status` - Enable/disable polling for the status of this Active Directory server.
* `user` - User name required to log into this Active Directory server.
* `adgrp` - LDAP Group Info.The structure of `adgrp` block is documented below.

The `adgrp` block contains:

* `name` - Name.
