---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnsslweb_portal"
description: |-
  Portal.
---

## fortios_vpnsslweb_portal
Portal.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `allow_user_access` - Allow user access to SSL-VPN applications. Valid values: `web` `ftp` `smb` `sftp` `telnet` `ssh` `vnc` `rdp` `ping` .
* `auto_connect` - Enable/disable automatic connect by client when system is up. Valid values: `enable` `disable` .
* `clipboard` - Enable to support RDP/VPC clipboard functionality. Valid values: `enable` `disable` .
* `custom_lang` - Change the web portal display language. Overrides config system global set language. You can use config system custom-language and execute system custom-language to add custom language files. This attribute must reference one of the following datasources: `system.custom-language.name` .
* `customize_forticlient_download_url` - Enable support of customized download URL for FortiClient. Valid values: `enable` `disable` .
* `display_bookmark` - Enable to display the web portal bookmark widget. Valid values: `enable` `disable` .
* `display_connection_tools` - Enable to display the web portal connection tools widget. Valid values: `enable` `disable` .
* `display_history` - Enable to display the web portal user login history widget. Valid values: `enable` `disable` .
* `display_status` - Enable to display the web portal status widget. Valid values: `enable` `disable` .
* `dns_server1` - IPv4 DNS server 1.
* `dns_server2` - IPv4 DNS server 2.
* `dns_suffix` - DNS suffix.
* `exclusive_routing` - Enable/disable all traffic go through tunnel only. Valid values: `enable` `disable` .
* `forticlient_download` - Enable/disable download option for FortiClient. Valid values: `enable` `disable` .
* `forticlient_download_method` - FortiClient download method. Valid values: `direct` `ssl-vpn` .
* `heading` - Web portal heading message.
* `hide_sso_credential` - Enable to prevent SSO credential being sent to client. Valid values: `enable` `disable` .
* `host_check` - Type of host checking performed on endpoints. Valid values: `none` `av` `fw` `av-fw` `custom` .
* `host_check_interval` - Periodic host check interval. Value of 0 means disabled and host checking only happens when the endpoint connects.
* `ip_mode` - Method by which users of this SSL-VPN tunnel obtain IP addresses. Valid values: `range` `user-group` .
* `ipv6_dns_server1` - IPv6 DNS server 1.
* `ipv6_dns_server2` - IPv6 DNS server 2.
* `ipv6_exclusive_routing` - Enable/disable all IPv6 traffic go through tunnel only. Valid values: `enable` `disable` .
* `ipv6_service_restriction` - Enable/disable IPv6 tunnel service restriction. Valid values: `enable` `disable` .
* `ipv6_split_tunneling` - Enable/disable IPv6 split tunneling. Valid values: `enable` `disable` .
* `ipv6_split_tunneling_routing_negate` - Enable to negate IPv6 split tunneling routing address. Valid values: `enable` `disable` .
* `ipv6_tunnel_mode` - Enable/disable IPv6 SSL-VPN tunnel mode. Valid values: `enable` `disable` .
* `ipv6_wins_server1` - IPv6 WINS server 1.
* `ipv6_wins_server2` - IPv6 WINS server 2.
* `keep_alive` - Enable/disable automatic reconnect for FortiClient connections. Valid values: `enable` `disable` .
* `limit_user_logins` - Enable to limit each user to one SSL-VPN session at a time. Valid values: `enable` `disable` .
* `mac_addr_action` - Client MAC address action. Valid values: `allow` `deny` .
* `mac_addr_check` - Enable/disable MAC address host checking. Valid values: `enable` `disable` .
* `macos_forticlient_download_url` - Download URL for Mac FortiClient.
* `name` - Portal name.
* `os_check` - Enable to let the FortiGate decide action based on client OS. Valid values: `enable` `disable` .
* `prefer_ipv6_dns` - prefer to query IPv6 dns first if enabled. Valid values: `enable` `disable` .
* `redir_url` - Client login redirect URL.
* `rewrite_ip_uri_ui` - Rewrite contents for URI contains IP and "/ui/". (default = disable) Valid values: `enable` `disable` .
* `save_password` - Enable/disable FortiClient saving the user's password. Valid values: `enable` `disable` .
* `service_restriction` - Enable/disable tunnel service restriction. Valid values: `enable` `disable` .
* `skip_check_for_browser` - Enable to skip host check for browser support. Valid values: `enable` `disable` .
* `skip_check_for_unsupported_os` - Enable to skip host check if client OS does not support it. Valid values: `enable` `disable` .
* `smb_max_version` - SMB maximum client protocol version. Valid values: `smbv1` `smbv2` `smbv3` .
* `smb_min_version` - SMB minimum client protocol version. Valid values: `smbv1` `smbv2` `smbv3` .
* `smb_ntlmv1_auth` - Enable support of NTLMv1 for Samba authentication. Valid values: `enable` `disable` .
* `smbv1` - smbv1 Valid values: `enable` `disable` .
* `split_tunneling` - Enable/disable IPv4 split tunneling. Valid values: `enable` `disable` .
* `split_tunneling_routing_negate` - Enable to negate split tunneling routing address. Valid values: `enable` `disable` .
* `theme` - Web portal color scheme. Valid values: `jade` `neutrino` `mariner` `graphite` `melongene` `dark-matter` `onyx` `eclipse` .
* `transform_backward_slashes` - Transform backward slashes to forward slashes in URLs. Valid values: `enable` `disable` .
* `tunnel_mode` - Enable/disable IPv4 SSL-VPN tunnel mode. Valid values: `enable` `disable` .
* `use_sdwan` - Use SD-WAN rules to get output interface. Valid values: `enable` `disable` .
* `user_bookmark` - Enable to allow web portal users to create their own bookmarks. Valid values: `enable` `disable` .
* `user_group_bookmark` - Enable to allow web portal users to create bookmarks for all users in the same user group. Valid values: `enable` `disable` .
* `web_mode` - Enable/disable SSL-VPN web mode. Valid values: `enable` `disable` .
* `windows_forticlient_download_url` - Download URL for Windows FortiClient.
* `wins_server1` - IPv4 WINS server 1.
* `wins_server2` - IPv4 WINS server 1.
* `bookmark_group` - Portal bookmark group. The structure of `bookmark_group` block is documented below.

The `bookmark_group` block contains:

* `name` - Bookmark group name.
* `bookmarks` - Bookmark table. The structure of `bookmarks` block is documented below.

The `bookmarks` block contains:

* `additional_params` - Additional parameters.
* `apptype` - Application type. Valid values: `ftp` `rdp` `sftp` `smb` `ssh` `telnet` `vnc` `web` .
* `color_depth` - Color depth per pixel. Valid values: `32` `16` `8` .
* `description` - Description.
* `domain` - Login domain.
* `folder` - Network shared file folder parameter.
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
* `form_data` - Form data. The structure of `form_data` block is documented below.

The `form_data` block contains:

* `name` - Name.
* `value` - Value.
* `host_check_policy` - One or more policies to require the endpoint to have specific security software. The structure of `host_check_policy` block is documented below.

The `host_check_policy` block contains:

* `name` - Host check software list name. This attribute must reference one of the following datasources: `vpn.ssl.web.host-check-software.name` .
* `ip_pools` - IPv4 firewall source address objects reserved for SSL-VPN tunnel mode clients. The structure of `ip_pools` block is documented below.

The `ip_pools` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `ipv6_pools` - IPv6 firewall source address objects reserved for SSL-VPN tunnel mode clients. The structure of `ipv6_pools` block is documented below.

The `ipv6_pools` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `ipv6_split_tunneling_routing_address` - IPv6 SSL-VPN tunnel mode firewall address objects that override firewall policy destination addresses to control split-tunneling access. The structure of `ipv6_split_tunneling_routing_address` block is documented below.

The `ipv6_split_tunneling_routing_address` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `mac_addr_check_rule` - Client MAC address check rule. The structure of `mac_addr_check_rule` block is documented below.

The `mac_addr_check_rule` block contains:

* `mac_addr_mask` - Client MAC address mask.
* `name` - Client MAC address check rule name.
* `mac_addr_list` - Client MAC address list. The structure of `mac_addr_list` block is documented below.

The `mac_addr_list` block contains:

* `addr` - Client MAC address.
* `os_check_list` - SSL-VPN OS checks. The structure of `os_check_list` block is documented below.

The `os_check_list` block contains:

* `action` - OS check options. Valid values: `deny` `allow` `check-up-to-date` .
* `latest_patch_level` - Latest OS patch level.
* `name` - Name.
* `tolerance` - OS patch level tolerance.
* `split_dns` - Split DNS for SSL-VPN. The structure of `split_dns` block is documented below.

The `split_dns` block contains:

* `dns_server1` - DNS server 1.
* `dns_server2` - DNS server 2.
* `domains` - Split DNS domains used for SSL-VPN clients separated by comma(,).
* `id` - ID.
* `ipv6_dns_server1` - IPv6 DNS server 1.
* `ipv6_dns_server2` - IPv6 DNS server 2.
* `split_tunneling_routing_address` - IPv4 SSL-VPN tunnel mode firewall address objects that override firewall policy destination addresses to control split-tunneling access. The structure of `split_tunneling_routing_address` block is documented below.

The `split_tunneling_routing_address` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_vpnsslweb_portal can be imported using:
```sh
terraform import fortios_vpnsslweb_portal.labelname {{mkey}}
```
