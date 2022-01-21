---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_radius"
description: |-
  Configure RADIUS server entries.
---

## fortios_user_radius
Configure RADIUS server entries.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `acct_all_servers` - Enable/disable sending of accounting messages to all configured servers (default = disable). Valid values: `enable` `disable` .
* `acct_interim_interval` - Time in seconds between each accounting interim update message.
* `all_usergroup` - Enable/disable automatically including this RADIUS server in all user groups. Valid values: `disable` `enable` .
* `auth_type` - Authentication methods/protocols permitted for this RADIUS server. Valid values: `auto` `ms_chap_v2` `ms_chap` `chap` `pap` .
* `group_override_attr_type` - RADIUS attribute type to override user group information. Valid values: `filter-Id` `class` .
* `h3c_compatibility` - Enable/disable compatibility with the H3C, a mechanism that performs security checking for authentication. Valid values: `enable` `disable` .
* `interface` - Specify outgoing interface to reach server. This attribute must reference one of the following datasources: `system.interface.name` .
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `name` - RADIUS server entry name.
* `nas_ip` - IP address used to communicate with the RADIUS server and used as NAS-IP-Address and Called-Station-ID attributes.
* `password_encoding` - Password encoding. Valid values: `auto` `ISO-8859-1` .
* `password_renewal` - Enable/disable password renewal. Valid values: `enable` `disable` .
* `radius_coa` - Enable to allow a mechanism to change the attributes of an authentication, authorization, and accounting session after it is authenticated. Valid values: `enable` `disable` .
* `radius_port` - RADIUS service port number.
* `rsso` - Enable/disable RADIUS based single sign on feature. Valid values: `enable` `disable` .
* `rsso_context_timeout` - Time in seconds before the logged out user is removed from the "user context list" of logged on users.
* `rsso_endpoint_attribute` - RADIUS attributes used to extract the user end point identifer from the RADIUS Start record. Valid values: `User-Name` `NAS-IP-Address` `Framed-IP-Address` `Framed-IP-Netmask` `Filter-Id` `Login-IP-Host` `Reply-Message` `Callback-Number` `Callback-Id` `Framed-Route` `Framed-IPX-Network` `Class` `Called-Station-Id` `Calling-Station-Id` `NAS-Identifier` `Proxy-State` `Login-LAT-Service` `Login-LAT-Node` `Login-LAT-Group` `Framed-AppleTalk-Zone` `Acct-Session-Id` `Acct-Multi-Session-Id` .
* `rsso_endpoint_block_attribute` - RADIUS attributes used to block a user. Valid values: `User-Name` `NAS-IP-Address` `Framed-IP-Address` `Framed-IP-Netmask` `Filter-Id` `Login-IP-Host` `Reply-Message` `Callback-Number` `Callback-Id` `Framed-Route` `Framed-IPX-Network` `Class` `Called-Station-Id` `Calling-Station-Id` `NAS-Identifier` `Proxy-State` `Login-LAT-Service` `Login-LAT-Node` `Login-LAT-Group` `Framed-AppleTalk-Zone` `Acct-Session-Id` `Acct-Multi-Session-Id` .
* `rsso_ep_one_ip_only` - Enable/disable the replacement of old IP addresses with new ones for the same endpoint on RADIUS accounting Start messages. Valid values: `enable` `disable` .
* `rsso_flush_ip_session` - Enable/disable flushing user IP sessions on RADIUS accounting Stop messages. Valid values: `enable` `disable` .
* `rsso_log_flags` - Events to log. Valid values: `protocol-error` `profile-missing` `accounting-stop-missed` `accounting-event` `endpoint-block` `radiusd-other` `none` .
* `rsso_log_period` - Time interval in seconds that group event log messages will be generated for dynamic profile events.
* `rsso_radius_response` - Enable/disable sending RADIUS response packets after receiving Start and Stop records. Valid values: `enable` `disable` .
* `rsso_radius_server_port` - UDP port to listen on for RADIUS Start and Stop records.
* `rsso_secret` - RADIUS secret used by the RADIUS accounting server.
* `rsso_validate_request_secret` - Enable/disable validating the RADIUS request shared secret in the Start or End record. Valid values: `enable` `disable` .
* `secondary_secret` - Secret key to access the secondary server.
* `secondary_server` - {<name_str|ip_str>} secondary RADIUS CN domain name or IP.
* `secret` - Pre-shared secret key used to access the primary RADIUS server.
* `server` - Primary RADIUS server CN domain name or IP address.
* `source_ip` - Source IP address for communications to the RADIUS server.
* `sso_attribute` - RADIUS attribute that contains the profile group name to be extracted from the RADIUS Start record. Valid values: `User-Name` `NAS-IP-Address` `Framed-IP-Address` `Framed-IP-Netmask` `Filter-Id` `Login-IP-Host` `Reply-Message` `Callback-Number` `Callback-Id` `Framed-Route` `Framed-IPX-Network` `Class` `Called-Station-Id` `Calling-Station-Id` `NAS-Identifier` `Proxy-State` `Login-LAT-Service` `Login-LAT-Node` `Login-LAT-Group` `Framed-AppleTalk-Zone` `Acct-Session-Id` `Acct-Multi-Session-Id` .
* `sso_attribute_key` - Key prefix for SSO group value in the SSO attribute.
* `sso_attribute_value_override` - Enable/disable override old attribute value with new value for the same endpoint. Valid values: `enable` `disable` .
* `switch_controller_acct_fast_framedip_detect` - Switch controller accounting message Framed-IP detection from DHCP snooping (seconds, default=2).
* `switch_controller_service_type` - RADIUS service type. Valid values: `login` `framed` `callback-login` `callback-framed` `outbound` `administrative` `nas-prompt` `authenticate-only` `callback-nas-prompt` `call-check` `callback-administrative` .
* `tertiary_secret` - Secret key to access the tertiary server.
* `tertiary_server` - {<name_str|ip_str>} tertiary RADIUS CN domain name or IP.
* `timeout` - Time in seconds between re-sending authentication requests.
* `use_management_vdom` - Enable/disable using management VDOM to send requests. Valid values: `enable` `disable` .
* `username_case_sensitive` - Enable/disable case sensitive user names. Valid values: `enable` `disable` .
* `accounting_server` - Additional accounting servers. The structure of `accounting_server` block is documented below.

The `accounting_server` block contains:

* `id` - ID (0 - 4294967295).
* `interface` - Specify outgoing interface to reach server. This attribute must reference one of the following datasources: `system.interface.name` .
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `port` - RADIUS accounting port number.
* `secret` - Secret key.
* `server` - {<name_str|ip_str>} Server CN domain name or IP.
* `source_ip` - Source IP address for communications to the RADIUS server.
* `status` - Status. Valid values: `enable` `disable` .
* `class` - Class attribute name(s). The structure of `class` block is documented below.

The `class` block contains:

* `name` - Class name.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_user_radius can be imported using:
```sh
terraform import fortios_user_radius.labelname {{mkey}}
```
