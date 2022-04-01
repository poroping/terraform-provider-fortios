---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_fsso"
description: |-
  Configure Fortinet Single Sign On (FSSO) agents.
---

## fortios_user_fsso
Configure Fortinet Single Sign On (FSSO) agents.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `group_poll_interval` - Interval in minutes within to fetch groups from FSSO server, or unset to disable.
* `interface` - Specify outgoing interface to reach server. This attribute must reference one of the following datasources: `system.interface.name` .
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `ldap_poll` - Enable/disable automatic fetching of groups from LDAP server. Valid values: `enable` `disable` .
* `ldap_poll_filter` - Filter used to fetch groups.
* `ldap_poll_interval` - Interval in minutes within to fetch groups from LDAP server.
* `ldap_server` - LDAP server to get group information. This attribute must reference one of the following datasources: `user.ldap.name` .
* `logon_timeout` - Interval in minutes to keep logons after FSSO server down.
* `name` - Name.
* `password` - Password of the first FSSO collector agent.
* `password2` - Password of the second FSSO collector agent.
* `password3` - Password of the third FSSO collector agent.
* `password4` - Password of the fourth FSSO collector agent.
* `password5` - Password of the fifth FSSO collector agent.
* `port` - Port of the first FSSO collector agent.
* `port2` - Port of the second FSSO collector agent.
* `port3` - Port of the third FSSO collector agent.
* `port4` - Port of the fourth FSSO collector agent.
* `port5` - Port of the fifth FSSO collector agent.
* `server` - Domain name or IP address of the first FSSO collector agent.
* `server2` - Domain name or IP address of the second FSSO collector agent.
* `server3` - Domain name or IP address of the third FSSO collector agent.
* `server4` - Domain name or IP address of the fourth FSSO collector agent.
* `server5` - Domain name or IP address of the fifth FSSO collector agent.
* `sni` - Server Name Indication.
* `source_ip` - Source IP for communications to FSSO agent.
* `source_ip6` - IPv6 source for communications to FSSO agent.
* `ssl` - Enable/disable use of SSL. Valid values: `enable` `disable` .
* `ssl_server_host_ip_check` - Enable/disable server host/IP verification. Valid values: `enable` `disable` .
* `ssl_trusted_cert` - Trusted server certificate or CA certificate. This attribute must reference one of the following datasources: `vpn.certificate.remote.name` `vpn.certificate.ca.name` .
* `type` - Server type. Valid values: `default` `fortinac` .
* `user_info_server` - LDAP server to get user information. This attribute must reference one of the following datasources: `user.ldap.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_user_fsso can be imported using:
```sh
terraform import fortios_user_fsso.labelname {{mkey}}
```
