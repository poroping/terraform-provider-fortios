---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_ldap"
description: |-
  Get information on a fortios Configure LDAP server entries.
---

# Data Source: fortios_user_ldap
Use this data source to get information on a fortios Configure LDAP server entries.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) LDAP server entry name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `account_key_filter` - Account key filter, using the UPN as the search filter.
* `account_key_processing` - Account key processing operation, either keep or strip domain string of UPN in the token.
* `antiphish` - Enable/disable AntiPhishing credential backend.
* `ca_cert` - CA certificate name.
* `client_cert` - Client certificate name.
* `client_cert_auth` - Enable/disable using client certificate for TLS authentication.
* `cnid` - Common name identifier for the LDAP server. The common name identifier for most LDAP servers is "cn".
* `dn` - Distinguished name used to look up entries on the LDAP server.
* `group_filter` - Filter used for group matching.
* `group_member_check` - Group member checking methods.
* `group_object_filter` - Filter used for group searching.
* `group_search_base` - Search base used for group searching.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `member_attr` - Name of attribute from which to get group membership.
* `name` - LDAP server entry name.
* `obtain_user_info` - Enable/disable obtaining of user information.
* `password` - Password for initial binding.
* `password_attr` - Name of attribute to get password hash.
* `password_expiry_warning` - Enable/disable password expiry warnings.
* `password_renewal` - Enable/disable online password renewal.
* `port` - Port to be used for communication with the LDAP server (default = 389).
* `search_type` - Search type.
* `secondary_server` - Secondary LDAP server CN domain name or IP.
* `secure` - Port to be used for authentication.
* `server` - LDAP server CN domain name or IP.
* `server_identity_check` - Enable/disable LDAP server identity check (verify server domain name/IP address against the server certificate).
* `source_ip` - FortiGate IP address to be used for communication with the LDAP server.
* `source_port` - Source port to be used for communication with the LDAP server.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `tertiary_server` - Tertiary LDAP server CN domain name or IP.
* `two_factor` - Enable/disable two-factor authentication.
* `two_factor_authentication` - Authentication method by FortiToken Cloud.
* `two_factor_notification` - Notification method for user activation by FortiToken Cloud.
* `type` - Authentication type for LDAP searches.
* `user_info_exchange_server` - MS Exchange server from which to fetch user information.
* `username` - Username (full DN) for initial binding.
