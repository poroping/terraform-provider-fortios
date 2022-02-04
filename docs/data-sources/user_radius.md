---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_radius"
description: |-
  Get information on a fortios Configure RADIUS server entries.
---

# Data Source: fortios_user_radius
Use this data source to get information on a fortios Configure RADIUS server entries.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) RADIUS server entry name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `acct_all_servers` - Enable/disable sending of accounting messages to all configured servers (default = disable).
* `acct_interim_interval` - Time in seconds between each accounting interim update message.
* `all_usergroup` - Enable/disable automatically including this RADIUS server in all user groups.
* `auth_type` - Authentication methods/protocols permitted for this RADIUS server.
* `group_override_attr_type` - RADIUS attribute type to override user group information.
* `h3c_compatibility` - Enable/disable compatibility with the H3C, a mechanism that performs security checking for authentication.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `name` - RADIUS server entry name.
* `nas_ip` - IP address used to communicate with the RADIUS server and used as NAS-IP-Address and Called-Station-ID attributes.
* `password_encoding` - Password encoding.
* `password_renewal` - Enable/disable password renewal.
* `radius_coa` - Enable to allow a mechanism to change the attributes of an authentication, authorization, and accounting session after it is authenticated.
* `radius_port` - RADIUS service port number.
* `rsso` - Enable/disable RADIUS based single sign on feature.
* `rsso_context_timeout` - Time in seconds before the logged out user is removed from the "user context list" of logged on users.
* `rsso_endpoint_attribute` - RADIUS attributes used to extract the user end point identifier from the RADIUS Start record.
* `rsso_endpoint_block_attribute` - RADIUS attributes used to block a user.
* `rsso_ep_one_ip_only` - Enable/disable the replacement of old IP addresses with new ones for the same endpoint on RADIUS accounting Start messages.
* `rsso_flush_ip_session` - Enable/disable flushing user IP sessions on RADIUS accounting Stop messages.
* `rsso_log_flags` - Events to log.
* `rsso_log_period` - Time interval in seconds that group event log messages will be generated for dynamic profile events.
* `rsso_radius_response` - Enable/disable sending RADIUS response packets after receiving Start and Stop records.
* `rsso_radius_server_port` - UDP port to listen on for RADIUS Start and Stop records.
* `rsso_secret` - RADIUS secret used by the RADIUS accounting server.
* `rsso_validate_request_secret` - Enable/disable validating the RADIUS request shared secret in the Start or End record.
* `secondary_secret` - Secret key to access the secondary server.
* `secondary_server` - Secondary RADIUS CN domain name or IP address.
* `secret` - Pre-shared secret key used to access the primary RADIUS server.
* `server` - Primary RADIUS server CN domain name or IP address.
* `source_ip` - Source IP address for communications to the RADIUS server.
* `sso_attribute` - RADIUS attribute that contains the profile group name to be extracted from the RADIUS Start record.
* `sso_attribute_key` - Key prefix for SSO group value in the SSO attribute.
* `sso_attribute_value_override` - Enable/disable override old attribute value with new value for the same endpoint.
* `switch_controller_acct_fast_framedip_detect` - Switch controller accounting message Framed-IP detection from DHCP snooping (seconds, default=2).
* `switch_controller_service_type` - RADIUS service type.
* `tertiary_secret` - Secret key to access the tertiary server.
* `tertiary_server` - Tertiary RADIUS CN domain name or IP address.
* `timeout` - Time in seconds between re-sending authentication requests.
* `use_management_vdom` - Enable/disable using management VDOM to send requests.
* `username_case_sensitive` - Enable/disable case sensitive user names.
* `accounting_server` - Additional accounting servers.The structure of `accounting_server` block is documented below.

The `accounting_server` block contains:

* `id` - ID (0 - 4294967295).
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `port` - RADIUS accounting port number.
* `secret` - Secret key.
* `server` - Server CN domain name or IP address.
* `source_ip` - Source IP address for communications to the RADIUS server.
* `status` - Status.
* `class` - Class attribute name(s).The structure of `class` block is documented below.

The `class` block contains:

* `name` - Class name.
