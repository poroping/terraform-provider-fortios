---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnsslweb_portal"
description: |-
  Get information on a fortios Portal.
---

# Data Source: fortios_vpnsslweb_portal
Use this data source to get information on a fortios Portal.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Portal name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `allow_user_access` - Allow user access to SSL-VPN applications.
* `auto_connect` - Enable/disable automatic connect by client when system is up.
* `clipboard` - Enable to support RDP/VPC clipboard functionality.
* `custom_lang` - Change the web portal display language. Overrides config system global set language. You can use config system custom-language and execute system custom-language to add custom language files.
* `customize_forticlient_download_url` - Enable support of customized download URL for FortiClient.
* `display_bookmark` - Enable to display the web portal bookmark widget.
* `display_connection_tools` - Enable to display the web portal connection tools widget.
* `display_history` - Enable to display the web portal user login history widget.
* `display_status` - Enable to display the web portal status widget.
* `dns_server1` - IPv4 DNS server 1.
* `dns_server2` - IPv4 DNS server 2.
* `dns_suffix` - DNS suffix.
* `exclusive_routing` - Enable/disable all traffic go through tunnel only.
* `forticlient_download` - Enable/disable download option for FortiClient.
* `forticlient_download_method` - FortiClient download method.
* `heading` - Web portal heading message.
* `hide_sso_credential` - Enable to prevent SSO credential being sent to client.
* `host_check` - Type of host checking performed on endpoints.
* `host_check_interval` - Periodic host check interval. Value of 0 means disabled and host checking only happens when the endpoint connects.
* `ip_mode` - Method by which users of this SSL-VPN tunnel obtain IP addresses.
* `ipv6_dns_server1` - IPv6 DNS server 1.
* `ipv6_dns_server2` - IPv6 DNS server 2.
* `ipv6_exclusive_routing` - Enable/disable all IPv6 traffic go through tunnel only.
* `ipv6_service_restriction` - Enable/disable IPv6 tunnel service restriction.
* `ipv6_split_tunneling` - Enable/disable IPv6 split tunneling.
* `ipv6_split_tunneling_routing_negate` - Enable to negate IPv6 split tunneling routing address.
* `ipv6_tunnel_mode` - Enable/disable IPv6 SSL-VPN tunnel mode.
* `ipv6_wins_server1` - IPv6 WINS server 1.
* `ipv6_wins_server2` - IPv6 WINS server 2.
* `keep_alive` - Enable/disable automatic reconnect for FortiClient connections.
* `limit_user_logins` - Enable to limit each user to one SSL-VPN session at a time.
* `mac_addr_action` - Client MAC address action.
* `mac_addr_check` - Enable/disable MAC address host checking.
* `macos_forticlient_download_url` - Download URL for Mac FortiClient.
* `name` - Portal name.
* `os_check` - Enable to let the FortiGate decide action based on client OS.
* `prefer_ipv6_dns` - Prefer to query IPv6 DNS server first if enabled.
* `redir_url` - Client login redirect URL.
* `rewrite_ip_uri_ui` - Rewrite contents for URI contains IP and /ui/ (default = disable).
* `save_password` - Enable/disable FortiClient saving the user's password.
* `service_restriction` - Enable/disable tunnel service restriction.
* `skip_check_for_browser` - Enable to skip host check for browser support.
* `skip_check_for_unsupported_os` - Enable to skip host check if client OS does not support it.
* `smb_max_version` - SMB maximum client protocol version.
* `smb_min_version` - SMB minimum client protocol version.
* `smb_ntlmv1_auth` - Enable support of NTLMv1 for Samba authentication.
* `smbv1` - SMB version 1.
* `split_tunneling` - Enable/disable IPv4 split tunneling.
* `split_tunneling_routing_negate` - Enable to negate split tunneling routing address.
* `theme` - Web portal color scheme.
* `transform_backward_slashes` - Transform backward slashes to forward slashes in URLs.
* `tunnel_mode` - Enable/disable IPv4 SSL-VPN tunnel mode.
* `use_sdwan` - Use SD-WAN rules to get output interface.
* `user_bookmark` - Enable to allow web portal users to create their own bookmarks.
* `user_group_bookmark` - Enable to allow web portal users to create bookmarks for all users in the same user group.
* `web_mode` - Enable/disable SSL-VPN web mode.
* `windows_forticlient_download_url` - Download URL for Windows FortiClient.
* `wins_server1` - IPv4 WINS server 1.
* `wins_server2` - IPv4 WINS server 1.
* `bookmark_group` - Portal bookmark group.The structure of `bookmark_group` block is documented below.

The `bookmark_group` block contains:

* `name` - Bookmark group name.
* `bookmarks` - Bookmark table.The structure of `bookmarks` block is documented below.

The `bookmarks` block contains:

* `additional_params` - Additional parameters.
* `apptype` - Application type.
* `color_depth` - Color depth per pixel.
* `description` - Description.
* `domain` - Login domain.
* `folder` - Network shared file folder parameter.
* `height` - Screen height (range from 480 - 65535, default = 768).
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
* `width` - Screen width (range from 640 - 65535, default = 1024).
* `form_data` - Form data.The structure of `form_data` block is documented below.

The `form_data` block contains:

* `name` - Name.
* `value` - Value.
* `host_check_policy` - One or more policies to require the endpoint to have specific security software.The structure of `host_check_policy` block is documented below.

The `host_check_policy` block contains:

* `name` - Host check software list name.
* `ip_pools` - IPv4 firewall source address objects reserved for SSL-VPN tunnel mode clients.The structure of `ip_pools` block is documented below.

The `ip_pools` block contains:

* `name` - Address name.
* `ipv6_pools` - IPv6 firewall source address objects reserved for SSL-VPN tunnel mode clients.The structure of `ipv6_pools` block is documented below.

The `ipv6_pools` block contains:

* `name` - Address name.
* `ipv6_split_tunneling_routing_address` - IPv6 SSL-VPN tunnel mode firewall address objects that override firewall policy destination addresses to control split-tunneling access.The structure of `ipv6_split_tunneling_routing_address` block is documented below.

The `ipv6_split_tunneling_routing_address` block contains:

* `name` - Address name.
* `mac_addr_check_rule` - Client MAC address check rule.The structure of `mac_addr_check_rule` block is documented below.

The `mac_addr_check_rule` block contains:

* `mac_addr_mask` - Client MAC address mask.
* `name` - Client MAC address check rule name.
* `mac_addr_list` - Client MAC address list.The structure of `mac_addr_list` block is documented below.

The `mac_addr_list` block contains:

* `addr` - Client MAC address.
* `os_check_list` - SSL-VPN OS checks.The structure of `os_check_list` block is documented below.

The `os_check_list` block contains:

* `action` - OS check options.
* `latest_patch_level` - Latest OS patch level.
* `name` - Name.
* `tolerance` - OS patch level tolerance.
* `split_dns` - Split DNS for SSL-VPN.The structure of `split_dns` block is documented below.

The `split_dns` block contains:

* `dns_server1` - DNS server 1.
* `dns_server2` - DNS server 2.
* `domains` - Split DNS domains used for SSL-VPN clients separated by comma.
* `id` - ID.
* `ipv6_dns_server1` - IPv6 DNS server 1.
* `ipv6_dns_server2` - IPv6 DNS server 2.
* `split_tunneling_routing_address` - IPv4 SSL-VPN tunnel mode firewall address objects that override firewall policy destination addresses to control split-tunneling access.The structure of `split_tunneling_routing_address` block is documented below.

The `split_tunneling_routing_address` block contains:

* `name` - Address name.
