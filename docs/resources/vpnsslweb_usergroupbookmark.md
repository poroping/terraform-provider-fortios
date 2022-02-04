---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnsslweb_usergroupbookmark"
description: |-
  Configure SSL-VPN user group bookmark.
---

## fortios_vpnsslweb_usergroupbookmark
Configure SSL-VPN user group bookmark.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `name` - Group name. This attribute must reference one of the following datasources: `user.group.name` .
* `bookmarks` - Bookmark table. The structure of `bookmarks` block is documented below.

The `bookmarks` block contains:

* `additional_params` - Additional parameters.
* `apptype` - Application type. Valid values: `ftp` `rdp` `sftp` `smb` `ssh` `telnet` `vnc` `web` .
* `color_depth` - Color depth per pixel. Valid values: `32` `16` `8` .
* `description` - Description.
* `domain` - Login domain.
* `folder` - Network shared file folder parameter.
* `height` - Screen height (range from 480 - 65535, default = 768).
* `host` - Host name/IP parameter.
* `keyboard_layout` - Keyboard layout. Valid values: `ar-101` `ar-102` `ar-102-azerty` `can-mul` `cz` `cz-qwerty` `cz-pr` `da` `nl` `de` `de-ch` `de-ibm` `en-uk` `en-uk-ext` `en-us` `en-us-dvorak` `es` `es-var` `fi` `fi-sami` `fr` `fr-ca` `fr-ch` `fr-be` `hr` `hu` `hu-101` `it` `it-142` `ja` `ko` `lt` `lt-ibm` `lt-std` `lav-std` `lav-leg` `mk` `mk-std` `no` `no-sami` `pol-214` `pol-pr` `pt` `pt-br` `pt-br-abnt2` `ru` `ru-mne` `ru-t` `sl` `sv` `sv-sami` `tuk` `tur-f` `tur-q` `zh-sym-sg-us` `zh-sym-us` `zh-tr-hk` `zh-tr-mo` `zh-tr-us` .
* `listening_port` - Listening port (0 - 65535).
* `load_balancing_info` - The load balancing information or cookie which should be provided to the connection broker.
* `logon_password` - Logon password.
* `logon_user` - Logon user.
* `name` - Bookmark name.
* `port` - Remote port.
* `preconnection_blob` - An arbitrary string which identifies the RDP source.
* `preconnection_id` - The numeric ID of the RDP source (0-4294967295).
* `remote_port` - Remote port (0 - 65535).
* `restricted_admin` - Enable/disable restricted admin mode for RDP. Valid values: `enable` `disable` .
* `security` - Security mode for RDP connection. Valid values: `rdp` `nla` `tls` `any` .
* `send_preconnection_id` - Enable/disable sending of preconnection ID. Valid values: `enable` `disable` .
* `server_layout` - Server side keyboard layout. Valid values: `de-de-qwertz` `en-gb-qwerty` `en-us-qwerty` `es-es-qwerty` `fr-ca-qwerty` `fr-fr-azerty` `fr-ch-qwertz` `it-it-qwerty` `ja-jp-qwerty` `pt-br-qwerty` `sv-se-qwerty` `tr-tr-qwerty` `failsafe` .
* `show_status_window` - Enable/disable showing of status window. Valid values: `enable` `disable` .
* `sso` - Single Sign-On. Valid values: `disable` `static` `auto` .
* `sso_credential` - Single sign-on credentials. Valid values: `sslvpn-login` `alternative` .
* `sso_credential_sent_once` - Single sign-on credentials are only sent once to remote server. Valid values: `enable` `disable` .
* `sso_password` - SSO password.
* `sso_username` - SSO user name.
* `url` - URL parameter.
* `width` - Screen width (range from 640 - 65535, default = 1024).
* `form_data` - Form data. The structure of `form_data` block is documented below.

The `form_data` block contains:

* `name` - Name.
* `value` - Value.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_vpnsslweb_usergroupbookmark can be imported using:
```sh
terraform import fortios_vpnsslweb_usergroupbookmark.labelname {{mkey}}
```
