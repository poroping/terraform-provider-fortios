---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_vap"
description: |-
  Configure Virtual Access Points (VAPs).
---

## fortios_wirelesscontroller_vap
Configure Virtual Access Points (VAPs).

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `access_control_list` - Profile name for access-control-list. This attribute must reference one of the following datasources: `wireless-controller.access-control-list.name` .
* `acct_interim_interval` - WiFi RADIUS accounting interim interval (60 - 86400 sec, default = 0).
* `additional_akms` - Additional AKMs. Valid values: `akm6` .
* `address_group` - Address group ID. This attribute must reference one of the following datasources: `wireless-controller.addrgrp.id` .
* `antivirus_profile` - AntiVirus profile name. This attribute must reference one of the following datasources: `antivirus.profile.name` .
* `application_list` - Application control list name. This attribute must reference one of the following datasources: `application.list.name` .
* `atf_weight` - Airtime weight in percentage (default = 20).
* `auth` - Authentication protocol. Valid values: `psk` `radius` `usergroup` .
* `auth_cert` - HTTPS server certificate. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `auth_portal_addr` - Address of captive portal.
* `beacon_advertising` - Fortinet beacon advertising IE data   (default = empty). Valid values: `name` `model` `serial-number` .
* `broadcast_ssid` - Enable/disable broadcasting the SSID (default = enable). Valid values: `enable` `disable` .
* `broadcast_suppression` - Optional suppression of broadcast messages. For example, you can keep DHCP messages, ARP broadcasts, and so on off of the wireless network. Valid values: `dhcp-up` `dhcp-down` `dhcp-starvation` `dhcp-ucast` `arp-known` `arp-unknown` `arp-reply` `arp-poison` `arp-proxy` `netbios-ns` `netbios-ds` `ipv6` `all-other-mc` `all-other-bc` .
* `bss_color_partial` - Enable/disable 802.11ax partial BSS color (default = enable). Valid values: `enable` `disable` .
* `bstm_disassociation_imminent` - Enable/disable forcing of disassociation after the BSTM request timer has been reached (default = enable). Valid values: `enable` `disable` .
* `bstm_load_balancing_disassoc_timer` - Time interval for client to voluntarily leave AP before forcing a disassociation due to AP load-balancing (0 to 30, default = 10).
* `bstm_rssi_disassoc_timer` - Time interval for client to voluntarily leave AP before forcing a disassociation due to low RSSI (0 to 2000, default = 200).
* `captive_portal_ac_name` - Local-bridging captive portal ac-name.
* `captive_portal_auth_timeout` - Hard timeout - AP will always clear the session after timeout regardless of traffic (0 - 864000 sec, default = 0).
* `captive_portal_macauth_radius_secret` - Secret key to access the macauth RADIUS server.
* `captive_portal_macauth_radius_server` - Captive portal external RADIUS server domain name or IP address.
* `captive_portal_radius_secret` - Secret key to access the RADIUS server.
* `captive_portal_radius_server` - Captive portal RADIUS server domain name or IP address.
* `captive_portal_session_timeout_interval` - Session timeout interval (0 - 864000 sec, default = 0).
* `dhcp_address_enforcement` - Enable/disable DHCP address enforcement (default = disable). Valid values: `enable` `disable` .
* `dhcp_lease_time` - DHCP lease time in seconds for NAT IP address.
* `dhcp_option43_insertion` - Enable/disable insertion of DHCP option 43 (default = enable). Valid values: `enable` `disable` .
* `dhcp_option82_circuit_id_insertion` - Enable/disable DHCP option 82 circuit-id insert (default = disable). Valid values: `style-1` `style-2` `style-3` `disable` .
* `dhcp_option82_insertion` - Enable/disable DHCP option 82 insert (default = disable). Valid values: `enable` `disable` .
* `dhcp_option82_remote_id_insertion` - Enable/disable DHCP option 82 remote-id insert (default = disable). Valid values: `style-1` `disable` .
* `dynamic_vlan` - Enable/disable dynamic VLAN assignment. Valid values: `enable` `disable` .
* `eap_reauth` - Enable/disable EAP re-authentication for WPA-Enterprise security. Valid values: `enable` `disable` .
* `eap_reauth_intv` - EAP re-authentication interval (1800 - 864000 sec, default = 86400).
* `eapol_key_retries` - Enable/disable retransmission of EAPOL-Key frames (message 3/4 and group message 1/2) (default = enable). Valid values: `disable` `enable` .
* `encrypt` - Encryption protocol to use (only available when security is set to a WPA type). Valid values: `TKIP` `AES` `TKIP-AES` .
* `external_fast_roaming` - Enable/disable fast roaming or pre-authentication with external APs not managed by the FortiGate (default = disable). Valid values: `enable` `disable` .
* `external_logout` - URL of external authentication logout server.
* `external_web` - URL of external authentication web server.
* `external_web_format` - URL query parameter detection (default = auto-detect). Valid values: `auto-detect` `no-query-string` `partial-query-string` .
* `fast_bss_transition` - Enable/disable 802.11r Fast BSS Transition (FT) (default = disable). Valid values: `disable` `enable` .
* `fast_roaming` - Enable/disable fast-roaming, or pre-authentication, where supported by clients (default = disable). Valid values: `enable` `disable` .
* `ft_mobility_domain` - Mobility domain identifier in FT (1 - 65535, default = 1000).
* `ft_over_ds` - Enable/disable FT over the Distribution System (DS). Valid values: `disable` `enable` .
* `ft_r0_key_lifetime` - Lifetime of the PMK-R0 key in FT, 1-65535 minutes.
* `gas_comeback_delay` - GAS comeback delay (0 or 100 - 10000 milliseconds, default = 500).
* `gas_fragmentation_limit` - GAS fragmentation limit (512 - 4096, default = 1024).
* `gtk_rekey` - Enable/disable GTK rekey for WPA security. Valid values: `enable` `disable` .
* `gtk_rekey_intv` - GTK rekey interval (1800 - 864000 sec, default = 86400).
* `high_efficiency` - Enable/disable 802.11ax high efficiency (default = enable). Valid values: `enable` `disable` .
* `hotspot20_profile` - Hotspot 2.0 profile name. This attribute must reference one of the following datasources: `wireless-controller.hotspot20.hs-profile.name` .
* `igmp_snooping` - Enable/disable IGMP snooping. Valid values: `enable` `disable` .
* `intra_vap_privacy` - Enable/disable blocking communication between clients on the same SSID (called intra-SSID privacy) (default = disable). Valid values: `enable` `disable` .
* `ip` - IP address and subnet mask for the local standalone NAT subnet.
* `ips_sensor` - IPS sensor name. This attribute must reference one of the following datasources: `ips.sensor.name` .
* `ipv6_rules` - Optional rules of IPv6 packets. For example, you can keep RA, RS and so on off of the wireless network. Valid values: `drop-icmp6ra` `drop-icmp6rs` `drop-llmnr6` `drop-icmp6mld2` `drop-dhcp6s` `drop-dhcp6c` `ndp-proxy` `drop-ns-dad` `drop-ns-nondad` .
* `key` - WEP Key.
* `keyindex` - WEP key index (1 - 4).
* `ldpc` - VAP low-density parity-check (LDPC) coding configuration. Valid values: `disable` `rx` `tx` `rxtx` .
* `local_authentication` - Enable/disable AP local authentication. Valid values: `enable` `disable` .
* `local_bridging` - Enable/disable bridging of wireless and Ethernet interfaces on the FortiAP (default = disable). Valid values: `enable` `disable` .
* `local_lan` - Allow/deny traffic destined for a Class A, B, or C private IP address (default = allow). Valid values: `allow` `deny` .
* `local_standalone` - Enable/disable AP local standalone (default = disable). Valid values: `enable` `disable` .
* `local_standalone_dns` - Enable/disable AP local standalone DNS. Valid values: `enable` `disable` .
* `local_standalone_dns_ip` - IPv4 addresses for the local standalone DNS.
* `local_standalone_nat` - Enable/disable AP local standalone NAT mode. Valid values: `enable` `disable` .
* `mac_auth_bypass` - Enable/disable MAC authentication bypass. Valid values: `enable` `disable` .
* `mac_called_station_delimiter` - MAC called station delimiter (default = hyphen). Valid values: `hyphen` `single-hyphen` `colon` `none` .
* `mac_calling_station_delimiter` - MAC calling station delimiter (default = hyphen). Valid values: `hyphen` `single-hyphen` `colon` `none` .
* `mac_case` - MAC case (default = uppercase). Valid values: `uppercase` `lowercase` .
* `mac_filter` - Enable/disable MAC filtering to block wireless clients by mac address. Valid values: `enable` `disable` .
* `mac_filter_policy_other` - Allow or block clients with MAC addresses that are not in the filter list. Valid values: `allow` `deny` .
* `mac_password_delimiter` - MAC authentication password delimiter (default = hyphen). Valid values: `hyphen` `single-hyphen` `colon` `none` .
* `mac_username_delimiter` - MAC authentication username delimiter (default = hyphen). Valid values: `hyphen` `single-hyphen` `colon` `none` .
* `max_clients` - Maximum number of clients that can connect simultaneously to the VAP (default = 0, meaning no limitation).
* `max_clients_ap` - Maximum number of clients that can connect simultaneously to the VAP per AP radio (default = 0, meaning no limitation).
* `mbo` - Enable/disable Multiband Operation (default = disable). Valid values: `disable` `enable` .
* `mbo_cell_data_conn_pref` - MBO cell data connection preference (0, 1, or 255, default = 1). Valid values: `excluded` `prefer-not` `prefer-use` .
* `me_disable_thresh` - Disable multicast enhancement when this many clients are receiving multicast traffic.
* `mesh_backhaul` - Enable/disable using this VAP as a WiFi mesh backhaul (default = disable). This entry is only available when security is set to a WPA type or open. Valid values: `enable` `disable` .
* `mpsk` - Enable/disable multiple PSK authentication. Valid values: `enable` `disable` .
* `mpsk_concurrent_clients` - Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).
* `mpsk_profile` - MPSK profile name. This attribute must reference one of the following datasources: `wireless-controller.mpsk-profile.name` .
* `mu_mimo` - Enable/disable Multi-user MIMO (default = enable). Valid values: `enable` `disable` .
* `multicast_enhance` - Enable/disable converting multicast to unicast to improve performance (default = disable). Valid values: `enable` `disable` .
* `multicast_rate` - Multicast rate (0, 6000, 12000, or 24000 kbps, default = 0). Valid values: `0` `6000` `12000` `24000` .
* `nac` - Enable/disable network access control. Valid values: `enable` `disable` .
* `nac_profile` - NAC profile name. This attribute must reference one of the following datasources: `wireless-controller.nac-profile.name` .
* `name` - Virtual AP name.
* `neighbor_report_dual_band` - Enable/disable dual-band neighbor report (default = disable). Valid values: `disable` `enable` .
* `okc` - Enable/disable Opportunistic Key Caching (OKC) (default = enable). Valid values: `disable` `enable` .
* `osen` - Enable/disable OSEN as part of key management (default = disable). Valid values: `enable` `disable` .
* `owe_groups` - OWE-Groups. Valid values: `19` `20` `21` .
* `owe_transition` - Enable/disable OWE transition mode support. Valid values: `disable` `enable` .
* `owe_transition_ssid` - OWE transition mode peer SSID.
* `passphrase` - WPA pre-shared key (PSK) to be used to authenticate WiFi users.
* `pmf` - Protected Management Frames (PMF) support (default = disable). Valid values: `disable` `enable` `optional` .
* `pmf_assoc_comeback_timeout` - Protected Management Frames (PMF) comeback maximum timeout (1-20 sec).
* `pmf_sa_query_retry_timeout` - Protected Management Frames (PMF) SA query retry timeout interval (1 - 5 100s of msec).
* `port_macauth` - Enable/disable LAN port MAC authentication (default = disable). Valid values: `disable` `radius` `address-group` .
* `port_macauth_reauth_timeout` - LAN port MAC authentication re-authentication timeout value (default = 7200 sec).
* `port_macauth_timeout` - LAN port MAC authentication idle timeout value (default = 600 sec).
* `portal_message_override_group` - Replacement message group for this VAP (only available when security is set to a captive portal type). This attribute must reference one of the following datasources: `system.replacemsg-group.name` .
* `portal_type` - Captive portal functionality. Configure how the captive portal authenticates users and whether it includes a disclaimer. Valid values: `auth` `auth+disclaimer` `disclaimer` `email-collect` `cmcc` `cmcc-macauth` `auth-mac` `external-auth` `external-macauth` .
* `primary_wag_profile` - Primary wireless access gateway profile name. This attribute must reference one of the following datasources: `wireless-controller.wag-profile.name` .
* `probe_resp_suppression` - Enable/disable probe response suppression (to ignore weak signals) (default = disable). Valid values: `enable` `disable` .
* `probe_resp_threshold` - Minimum signal level/threshold in dBm required for the AP response to probe requests (-95 to -20, default = -80).
* `ptk_rekey` - Enable/disable PTK rekey for WPA-Enterprise security. Valid values: `enable` `disable` .
* `ptk_rekey_intv` - PTK rekey interval (1800 - 864000 sec, default = 86400).
* `qos_profile` - Quality of service profile name. This attribute must reference one of the following datasources: `wireless-controller.qos-profile.name` .
* `quarantine` - Enable/disable station quarantine (default = enable). Valid values: `enable` `disable` .
* `radio_2g_threshold` - Minimum signal level/threshold in dBm required for the AP response to receive a packet in 2.4G band (-95 to -20, default = -79).
* `radio_5g_threshold` - Minimum signal level/threshold in dBm required for the AP response to receive a packet in 5G band(-95 to -20, default = -76).
* `radio_sensitivity` - Enable/disable software radio sensitivity (to ignore weak signals) (default = disable). Valid values: `enable` `disable` .
* `radius_mac_auth` - Enable/disable RADIUS-based MAC authentication of clients (default = disable). Valid values: `enable` `disable` .
* `radius_mac_auth_server` - RADIUS-based MAC authentication server. This attribute must reference one of the following datasources: `user.radius.name` .
* `radius_mac_mpsk_auth` - Enable/disable RADIUS-based MAC authentication of clients for MPSK authentication (default = disable). Valid values: `enable` `disable` .
* `radius_mac_mpsk_timeout` - RADIUS MAC MPSK cache timeout interval (1800 - 864000, default = 86400).
* `radius_server` - RADIUS server to be used to authenticate WiFi users. This attribute must reference one of the following datasources: `user.radius.name` .
* `rates_11a` - Allowed data rates for 802.11a. Valid values: `1` `1-basic` `2` `2-basic` `5.5` `5.5-basic` `11` `11-basic` `6` `6-basic` `9` `9-basic` `12` `12-basic` `18` `18-basic` `24` `24-basic` `36` `36-basic` `48` `48-basic` `54` `54-basic` .
* `rates_11ac_ss12` - Allowed data rates for 802.11ac with 1 or 2 spatial streams. Valid values: `mcs0/1` `mcs1/1` `mcs2/1` `mcs3/1` `mcs4/1` `mcs5/1` `mcs6/1` `mcs7/1` `mcs8/1` `mcs9/1` `mcs10/1` `mcs11/1` `mcs0/2` `mcs1/2` `mcs2/2` `mcs3/2` `mcs4/2` `mcs5/2` `mcs6/2` `mcs7/2` `mcs8/2` `mcs9/2` `mcs10/2` `mcs11/2` .
* `rates_11ac_ss34` - Allowed data rates for 802.11ac with 3 or 4 spatial streams. Valid values: `mcs0/3` `mcs1/3` `mcs2/3` `mcs3/3` `mcs4/3` `mcs5/3` `mcs6/3` `mcs7/3` `mcs8/3` `mcs9/3` `mcs10/3` `mcs11/3` `mcs0/4` `mcs1/4` `mcs2/4` `mcs3/4` `mcs4/4` `mcs5/4` `mcs6/4` `mcs7/4` `mcs8/4` `mcs9/4` `mcs10/4` `mcs11/4` .
* `rates_11ax_ss12` - Allowed data rates for 802.11ax with 1 or 2 spatial streams. Valid values: `mcs0/1` `mcs1/1` `mcs2/1` `mcs3/1` `mcs4/1` `mcs5/1` `mcs6/1` `mcs7/1` `mcs8/1` `mcs9/1` `mcs10/1` `mcs11/1` `mcs0/2` `mcs1/2` `mcs2/2` `mcs3/2` `mcs4/2` `mcs5/2` `mcs6/2` `mcs7/2` `mcs8/2` `mcs9/2` `mcs10/2` `mcs11/2` .
* `rates_11ax_ss34` - Allowed data rates for 802.11ax with 3 or 4 spatial streams. Valid values: `mcs0/3` `mcs1/3` `mcs2/3` `mcs3/3` `mcs4/3` `mcs5/3` `mcs6/3` `mcs7/3` `mcs8/3` `mcs9/3` `mcs10/3` `mcs11/3` `mcs0/4` `mcs1/4` `mcs2/4` `mcs3/4` `mcs4/4` `mcs5/4` `mcs6/4` `mcs7/4` `mcs8/4` `mcs9/4` `mcs10/4` `mcs11/4` .
* `rates_11bg` - Allowed data rates for 802.11b/g. Valid values: `1` `1-basic` `2` `2-basic` `5.5` `5.5-basic` `11` `11-basic` `6` `6-basic` `9` `9-basic` `12` `12-basic` `18` `18-basic` `24` `24-basic` `36` `36-basic` `48` `48-basic` `54` `54-basic` .
* `rates_11n_ss12` - Allowed data rates for 802.11n with 1 or 2 spatial streams. Valid values: `mcs0/1` `mcs1/1` `mcs2/1` `mcs3/1` `mcs4/1` `mcs5/1` `mcs6/1` `mcs7/1` `mcs8/2` `mcs9/2` `mcs10/2` `mcs11/2` `mcs12/2` `mcs13/2` `mcs14/2` `mcs15/2` .
* `rates_11n_ss34` - Allowed data rates for 802.11n with 3 or 4 spatial streams. Valid values: `mcs16/3` `mcs17/3` `mcs18/3` `mcs19/3` `mcs20/3` `mcs21/3` `mcs22/3` `mcs23/3` `mcs24/4` `mcs25/4` `mcs26/4` `mcs27/4` `mcs28/4` `mcs29/4` `mcs30/4` `mcs31/4` .
* `sae_groups` - SAE-Groups. Valid values: `19` `20` `21` .
* `sae_password` - WPA3 SAE password to be used to authenticate WiFi users.
* `scan_botnet_connections` - Block or monitor connections to Botnet servers or disable Botnet scanning. Valid values: `disable` `monitor` `block` .
* `secondary_wag_profile` - Secondary wireless access gateway profile name. This attribute must reference one of the following datasources: `wireless-controller.wag-profile.name` .
* `security` - Security mode for the wireless interface (default = wpa2-only-personal). Valid values: `open` `captive-portal` `wep64` `wep128` `wpa-personal` `wpa-personal+captive-portal` `wpa-enterprise` `wpa-only-personal` `wpa-only-personal+captive-portal` `wpa-only-enterprise` `wpa2-only-personal` `wpa2-only-personal+captive-portal` `wpa2-only-enterprise` `wpa3-enterprise` `wpa3-only-enterprise` `wpa3-enterprise-transition` `wpa3-sae` `wpa3-sae-transition` `owe` `osen` .
* `security_exempt_list` - Optional security exempt list for captive portal authentication. This attribute must reference one of the following datasources: `user.security-exempt-list.name` .
* `security_redirect_url` - Optional URL for redirecting users after they pass captive portal authentication.
* `split_tunneling` - Enable/disable split tunneling (default = disable). Valid values: `enable` `disable` .
* `ssid` - IEEE 802.11 service set identifier (SSID) for the wireless interface. Users who wish to use the wireless network must configure their computers to access this SSID name.
* `sticky_client_remove` - Enable/disable sticky client remove to maintain good signal level clients in SSID (default = disable). Valid values: `enable` `disable` .
* `sticky_client_threshold_2g` - Minimum signal level/threshold in dBm required for the 2G client to be serviced by the AP (-95 to -20, default = -79).
* `sticky_client_threshold_5g` - Minimum signal level/threshold in dBm required for the 5G client to be serviced by the AP (-95 to -20, default = -76).
* `target_wake_time` - Enable/disable 802.11ax target wake time (default = enable). Valid values: `enable` `disable` .
* `tkip_counter_measure` - Enable/disable TKIP counter measure. Valid values: `enable` `disable` .
* `tunnel_echo_interval` - The time interval to send echo to both primary and secondary tunnel peers (1 - 65535 sec, default = 300).
* `tunnel_fallback_interval` - The time interval for secondary tunnel to fall back to primary tunnel (0 - 65535 sec, default = 7200).
* `utm_log` - Enable/disable UTM logging. Valid values: `enable` `disable` .
* `utm_profile` - UTM profile name. This attribute must reference one of the following datasources: `wireless-controller.utm-profile.name` .
* `utm_status` - Enable to add one or more security profiles (AV, IPS, etc.) to the VAP. Valid values: `enable` `disable` .
* `vlan_auto` - Enable/disable automatic management of SSID VLAN interface. Valid values: `enable` `disable` .
* `vlan_pooling` - Enable/disable VLAN pooling, to allow grouping of multiple wireless controller VLANs into VLAN pools (default = disable). When set to wtp-group, VLAN pooling occurs with VLAN assignment by wtp-group. Valid values: `wtp-group` `round-robin` `hash` `disable` .
* `vlanid` - Optional VLAN ID.
* `voice_enterprise` - Enable/disable 802.11k and 802.11v assisted Voice-Enterprise roaming (default = disable). Valid values: `disable` `enable` .
* `webfilter_profile` - WebFilter profile name. This attribute must reference one of the following datasources: `webfilter.profile.name` .
* `mac_filter_list` - Create a list of MAC addresses for MAC address filtering. The structure of `mac_filter_list` block is documented below.

The `mac_filter_list` block contains:

* `id` - ID.
* `mac` - MAC address.
* `mac_filter_policy` - Deny or allow the client with this MAC address. Valid values: `allow` `deny` .
* `mpsk_key` - List of multiple PSK entries. The structure of `mpsk_key` block is documented below.

The `mpsk_key` block contains:

* `comment` - Comment.
* `concurrent_clients` - Number of clients that can connect using this pre-shared key.
* `key_name` - Pre-shared key name.
* `passphrase` - WPA Pre-shared key.
* `mpsk_schedules` - Firewall schedule for MPSK passphrase. The passphrase will be effective only when at least one schedule is valid. The structure of `mpsk_schedules` block is documented below.

The `mpsk_schedules` block contains:

* `name` - Schedule name. This attribute must reference one of the following datasources: `firewall.schedule.group.name` `firewall.schedule.recurring.name` `firewall.schedule.onetime.name` .
* `portal_message_overrides` - Individual message overrides. The structure of `portal_message_overrides` block is documented below.

The `portal_message_overrides` block contains:

* `auth_disclaimer_page` - Override auth-disclaimer-page message with message from portal-message-overrides group.
* `auth_login_failed_page` - Override auth-login-failed-page message with message from portal-message-overrides group.
* `auth_login_page` - Override auth-login-page message with message from portal-message-overrides group.
* `auth_reject_page` - Override auth-reject-page message with message from portal-message-overrides group.
* `radius_mac_auth_usergroups` - Selective user groups that are permitted for RADIUS mac authentication. The structure of `radius_mac_auth_usergroups` block is documented below.

The `radius_mac_auth_usergroups` block contains:

* `name` - User group name.
* `schedule` - Firewall schedules for enabling this VAP on the FortiAP. This VAP will be enabled when at least one of the schedules is valid. Separate multiple schedule names with a space. The structure of `schedule` block is documented below.

The `schedule` block contains:

* `name` - Schedule name. This attribute must reference one of the following datasources: `firewall.schedule.group.name` `firewall.schedule.recurring.name` `firewall.schedule.onetime.name` .
* `selected_usergroups` - Selective user groups that are permitted to authenticate. The structure of `selected_usergroups` block is documented below.

The `selected_usergroups` block contains:

* `name` - User group name. This attribute must reference one of the following datasources: `user.group.name` .
* `usergroup` - Firewall user group to be used to authenticate WiFi users. The structure of `usergroup` block is documented below.

The `usergroup` block contains:

* `name` - User group name. This attribute must reference one of the following datasources: `user.group.name` .
* `vlan_name` - Table for mapping VLAN name to VLAN ID. The structure of `vlan_name` block is documented below.

The `vlan_name` block contains:

* `name` - VLAN name.
* `vlan_id` - VLAN ID.
* `vlan_pool` - VLAN pool. The structure of `vlan_pool` block is documented below.

The `vlan_pool` block contains:

* `id` - ID.
* `wtp_group` - WTP group name. This attribute must reference one of the following datasources: `wireless-controller.wtp-group.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontroller_vap can be imported using:
```sh
terraform import fortios_wirelesscontroller_vap.labelname {{mkey}}
```
