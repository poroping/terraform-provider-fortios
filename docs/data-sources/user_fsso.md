---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_fsso"
description: |-
  Get information on a fortios Configure Fortinet Single Sign On (FSSO) agents.
---

# Data Source: fortios_user_fsso
Use this data source to get information on a fortios Configure Fortinet Single Sign On (FSSO) agents.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `group_poll_interval` - Interval in minutes within to fetch groups from FSSO server, or unset to disable.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `ldap_poll` - Enable/disable automatic fetching of groups from LDAP server.
* `ldap_poll_filter` - Filter used to fetch groups.
* `ldap_poll_interval` - Interval in minutes within to fetch groups from LDAP server.
* `ldap_server` - LDAP server to get group information.
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
* `source_ip` - Source IP for communications to FSSO agent.
* `source_ip6` - IPv6 source for communications to FSSO agent.
* `ssl` - Enable/disable use of SSL.
* `ssl_server_host_ip_check` - Enable/disable server host/IP verification.
* `ssl_trusted_cert` - Trusted server certificate or CA certificate.
* `type` - Server type.
* `user_info_server` - LDAP server to get user information.
