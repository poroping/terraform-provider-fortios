---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnsslweb_userbookmark"
description: |-
  Get information on a fortios Configure SSL-VPN user bookmark.
---

# Data Source: fortios_vpnsslweb_userbookmark
Use this data source to get information on a fortios Configure SSL-VPN user bookmark.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) User and group name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `custom_lang` - Personal language.
* `name` - User and group name.
* `bookmarks` - Bookmark table.The structure of `bookmarks` block is documented below.

The `bookmarks` block contains:

* `additional_params` - Additional parameters.
* `apptype` - Application type.
* `color_depth` - Color depth per pixel.
* `description` - Description.
* `domain` - Login domain.
* `folder` - Network shared file folder parameter.
* `height` - Screen height (range from 0 - 65535, default = 0).
* `host` - Host name/IP parameter.
* `keyboard_layout` - Keyboard layout.
* `listening_port` - Listening port (0 - 65535).
* `load_balancing_info` - The load balancing information or cookie which should be provided to the connection broker.
* `logon_password` - Logon password.
* `logon_user` - Logon user.
* `name` - Bookmark name.
* `port` - Remote port.
* `preconnection_blob` - An arbitrary string which identifies the RDP source.
* `preconnection_id` - The numeric ID of the RDP source (0-4294967295).
* `remote_port` - Remote port (0 - 65535).
* `restricted_admin` - Enable/disable restricted admin mode for RDP.
* `security` - Security mode for RDP connection.
* `send_preconnection_id` - Enable/disable sending of preconnection ID.
* `server_layout` - Server side keyboard layout.
* `show_status_window` - Enable/disable showing of status window.
* `sso` - Single Sign-On.
* `sso_credential` - Single sign-on credentials.
* `sso_credential_sent_once` - Single sign-on credentials are only sent once to remote server.
* `sso_password` - SSO password.
* `sso_username` - SSO user name.
* `url` - URL parameter.
* `vnc_keyboard_layout` - Keyboard layout.
* `width` - Screen width (range from 0 - 65535, default = 0).
* `form_data` - Form data.The structure of `form_data` block is documented below.

The `form_data` block contains:

* `name` - Name.
* `value` - Value.
