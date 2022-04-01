---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_ldap"
description: |-
  Configure LDAP server entries.
---

## fortios_user_ldap
Configure LDAP server entries.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `account_key_filter` - Account key filter, using the UPN as the search filter.
* `account_key_processing` - Account key processing operation, either keep or strip domain string of UPN in the token. Valid values: `same` `strip` .
* `antiphish` - Enable/disable AntiPhishing credential backend. Valid values: `enable` `disable` .
* `ca_cert` - CA certificate name. This attribute must reference one of the following datasources: `vpn.certificate.ca.name` .
* `client_cert` - Client certificate name. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `client_cert_auth` - Enable/disable using client certificate for TLS authentication. Valid values: `enable` `disable` .
* `cnid` - Common name identifier for the LDAP server. The common name identifier for most LDAP servers is "cn".
* `dn` - Distinguished name used to look up entries on the LDAP server.
* `group_filter` - Filter used for group matching.
* `group_member_check` - Group member checking methods. Valid values: `user-attr` `group-object` `posix-group-object` .
* `group_object_filter` - Filter used for group searching.
* `group_search_base` - Search base used for group searching.
* `interface` - Specify outgoing interface to reach server. This attribute must reference one of the following datasources: `system.interface.name` .
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `member_attr` - Name of attribute from which to get group membership.
* `name` - LDAP server entry name.
* `obtain_user_info` - Enable/disable obtaining of user information. Valid values: `enable` `disable` .
* `password` - Password for initial binding.
* `password_attr` - Name of attribute to get password hash.
* `password_expiry_warning` - Enable/disable password expiry warnings. Valid values: `enable` `disable` .
* `password_renewal` - Enable/disable online password renewal. Valid values: `enable` `disable` .
* `port` - Port to be used for communication with the LDAP server (default = 389).
* `search_type` - Search type. Valid values: `recursive` .
* `secondary_server` - Secondary LDAP server CN domain name or IP.
* `secure` - Port to be used for authentication. Valid values: `disable` `starttls` `ldaps` .
* `server` - LDAP server CN domain name or IP.
* `server_identity_check` - Enable/disable LDAP server identity check (verify server domain name/IP address against the server certificate). Valid values: `enable` `disable` .
* `source_ip` - FortiGate IP address to be used for communication with the LDAP server.
* `source_port` - Source port to be used for communication with the LDAP server.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default` `SSLv3` `TLSv1` `TLSv1-1` `TLSv1-2` .
* `tertiary_server` - Tertiary LDAP server CN domain name or IP.
* `two_factor` - Enable/disable two-factor authentication. Valid values: `disable` `fortitoken-cloud` .
* `two_factor_authentication` - Authentication method by FortiToken Cloud. Valid values: `fortitoken` `email` `sms` .
* `two_factor_notification` - Notification method for user activation by FortiToken Cloud. Valid values: `email` `sms` .
* `type` - Authentication type for LDAP searches. Valid values: `simple` `anonymous` `regular` .
* `user_info_exchange_server` - MS Exchange server from which to fetch user information. This attribute must reference one of the following datasources: `user.exchange.name` .
* `username` - Username (full DN) for initial binding.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_user_ldap can be imported using:
```sh
terraform import fortios_user_ldap.labelname {{mkey}}
```
