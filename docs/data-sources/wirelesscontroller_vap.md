---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_vap"
description: |-
  Get information on a fortios Configure Virtual Access Points (VAPs).
---

# Data Source: fortios_wirelesscontroller_vap
Use this data source to get information on a fortios Configure Virtual Access Points (VAPs).


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Virtual AP name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `access_control_list` - Profile name for access-control-list.
* `acct_interim_interval` - WiFi RADIUS accounting interim interval (60 - 86400 sec, default = 0).
* `additional_akms` - Additional AKMs.
* `address_group` - Firewall Address Group Name.
* `address_group_policy` - Configure MAC address filtering policy for MAC addresses that are in the address-group.
* `antivirus_profile` - AntiVirus profile name.
* `application_detection_engine` - Enable/disable application detection engine (default = disable).
* `application_dscp_marking` - Enable/disable application attribute based DSCP marking (default = disable).
* `application_list` - Application control list name.
* `application_report_intv` - Application report interval (30 - 864000 sec, default = 120).
* `atf_weight` - Airtime weight in percentage (default = 20).
* `auth` - Authentication protocol.
* `auth_cert` - HTTPS server certificate.
* `auth_portal_addr` - Address of captive portal.
* `beacon_advertising` - Fortinet beacon advertising IE data   (default = empty).
* `broadcast_ssid` - Enable/disable broadcasting the SSID (default = enable).
* `broadcast_suppression` - Optional suppression of broadcast messages. For example, you can keep DHCP messages, ARP broadcasts, and so on off of the wireless network.
* `bss_color_partial` - Enable/disable 802.11ax partial BSS color (default = enable).
* `bstm_disassociation_imminent` - Enable/disable forcing of disassociation after the BSTM request timer has been reached (default = enable).
* `bstm_load_balancing_disassoc_timer` - Time interval for client to voluntarily leave AP before forcing a disassociation due to AP load-balancing (0 to 30, default = 10).
* `bstm_rssi_disassoc_timer` - Time interval for client to voluntarily leave AP before forcing a disassociation due to low RSSI (0 to 2000, default = 200).
* `captive_portal_ac_name` - Local-bridging captive portal ac-name.
* `captive_portal_auth_timeout` - Hard timeout - AP will always clear the session after timeout regardless of traffic (0 - 864000 sec, default = 0).
* `captive_portal_fw_accounting` - Enable/disable RADIUS accounting for captive portal firewall authentication session.
* `captive_portal_macauth_radius_secret` - Secret key to access the macauth RADIUS server.
* `captive_portal_macauth_radius_server` - Captive portal external RADIUS server domain name or IP address.
* `captive_portal_radius_secret` - Secret key to access the RADIUS server.
* `captive_portal_radius_server` - Captive portal RADIUS server domain name or IP address.
* `captive_portal_session_timeout_interval` - Session timeout interval (0 - 864000 sec, default = 0).
* `dhcp_address_enforcement` - Enable/disable DHCP address enforcement (default = disable).
* `dhcp_lease_time` - DHCP lease time in seconds for NAT IP address.
* `dhcp_option43_insertion` - Enable/disable insertion of DHCP option 43 (default = enable).
* `dhcp_option82_circuit_id_insertion` - Enable/disable DHCP option 82 circuit-id insert (default = disable).
* `dhcp_option82_insertion` - Enable/disable DHCP option 82 insert (default = disable).
* `dhcp_option82_remote_id_insertion` - Enable/disable DHCP option 82 remote-id insert (default = disable).
* `dynamic_vlan` - Enable/disable dynamic VLAN assignment.
* `eap_reauth` - Enable/disable EAP re-authentication for WPA-Enterprise security.
* `eap_reauth_intv` - EAP re-authentication interval (1800 - 864000 sec, default = 86400).
* `eapol_key_retries` - Enable/disable retransmission of EAPOL-Key frames (message 3/4 and group message 1/2) (default = enable).
* `encrypt` - Encryption protocol to use (only available when security is set to a WPA type).
* `external_fast_roaming` - Enable/disable fast roaming or pre-authentication with external APs not managed by the FortiGate (default = disable).
* `external_logout` - URL of external authentication logout server.
* `external_web` - URL of external authentication web server.
* `external_web_format` - URL query parameter detection (default = auto-detect).
* `fast_bss_transition` - Enable/disable 802.11r Fast BSS Transition (FT) (default = disable).
* `fast_roaming` - Enable/disable fast-roaming, or pre-authentication, where supported by clients (default = disable).
* `ft_mobility_domain` - Mobility domain identifier in FT (1 - 65535, default = 1000).
* `ft_over_ds` - Enable/disable FT over the Distribution System (DS).
* `ft_r0_key_lifetime` - Lifetime of the PMK-R0 key in FT, 1-65535 minutes.
* `gas_comeback_delay` - GAS comeback delay (0 or 100 - 10000 milliseconds, default = 500).
* `gas_fragmentation_limit` - GAS fragmentation limit (512 - 4096, default = 1024).
* `gtk_rekey` - Enable/disable GTK rekey for WPA security.
* `gtk_rekey_intv` - GTK rekey interval (1800 - 864000 sec, default = 86400).
* `high_efficiency` - Enable/disable 802.11ax high efficiency (default = enable).
* `hotspot20_profile` - Hotspot 2.0 profile name.
* `igmp_snooping` - Enable/disable IGMP snooping.
* `intra_vap_privacy` - Enable/disable blocking communication between clients on the same SSID (called intra-SSID privacy) (default = disable).
* `ip` - IP address and subnet mask for the local standalone NAT subnet.
* `ips_sensor` - IPS sensor name.
* `ipv6_rules` - Optional rules of IPv6 packets. For example, you can keep RA, RS and so on off of the wireless network.
* `key` - WEP Key.
* `keyindex` - WEP key index (1 - 4).
* `l3_roaming` - Enable/disable layer 3 roaming (default = disable).
* `l3_roaming_mode` - Select the way that layer 3 roaming traffic is passed (default = direct).
* `ldpc` - VAP low-density parity-check (LDPC) coding configuration.
* `local_authentication` - Enable/disable AP local authentication.
* `local_bridging` - Enable/disable bridging of wireless and Ethernet interfaces on the FortiAP (default = disable).
* `local_lan` - Allow/deny traffic destined for a Class A, B, or C private IP address (default = allow).
* `local_standalone` - Enable/disable AP local standalone (default = disable).
* `local_standalone_dns` - Enable/disable AP local standalone DNS.
* `local_standalone_dns_ip` - IPv4 addresses for the local standalone DNS.
* `local_standalone_nat` - Enable/disable AP local standalone NAT mode.
* `mac_auth_bypass` - Enable/disable MAC authentication bypass.
* `mac_called_station_delimiter` - MAC called station delimiter (default = hyphen).
* `mac_calling_station_delimiter` - MAC calling station delimiter (default = hyphen).
* `mac_case` - MAC case (default = uppercase).
* `mac_filter` - Enable/disable MAC filtering to block wireless clients by mac address.
* `mac_filter_policy_other` - Allow or block clients with MAC addresses that are not in the filter list.
* `mac_password_delimiter` - MAC authentication password delimiter (default = hyphen).
* `mac_username_delimiter` - MAC authentication username delimiter (default = hyphen).
* `max_clients` - Maximum number of clients that can connect simultaneously to the VAP (default = 0, meaning no limitation).
* `max_clients_ap` - Maximum number of clients that can connect simultaneously to the VAP per AP radio (default = 0, meaning no limitation).
* `mbo` - Enable/disable Multiband Operation (default = disable).
* `mbo_cell_data_conn_pref` - MBO cell data connection preference (0, 1, or 255, default = 1).
* `me_disable_thresh` - Disable multicast enhancement when this many clients are receiving multicast traffic.
* `mesh_backhaul` - Enable/disable using this VAP as a WiFi mesh backhaul (default = disable). This entry is only available when security is set to a WPA type or open.
* `mpsk` - Enable/disable multiple PSK authentication.
* `mpsk_concurrent_clients` - Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).
* `mpsk_profile` - MPSK profile name.
* `mu_mimo` - Enable/disable Multi-user MIMO (default = enable).
* `multicast_enhance` - Enable/disable converting multicast to unicast to improve performance (default = disable).
* `multicast_rate` - Multicast rate (0, 6000, 12000, or 24000 kbps, default = 0).
* `nac` - Enable/disable network access control.
* `nac_profile` - NAC profile name.
* `name` - Virtual AP name.
* `neighbor_report_dual_band` - Enable/disable dual-band neighbor report (default = disable).
* `okc` - Enable/disable Opportunistic Key Caching (OKC) (default = enable).
* `osen` - Enable/disable OSEN as part of key management (default = disable).
* `owe_groups` - OWE-Groups.
* `owe_transition` - Enable/disable OWE transition mode support.
* `owe_transition_ssid` - OWE transition mode peer SSID.
* `passphrase` - WPA pre-shared key (PSK) to be used to authenticate WiFi users.
* `pmf` - Protected Management Frames (PMF) support (default = disable).
* `pmf_assoc_comeback_timeout` - Protected Management Frames (PMF) comeback maximum timeout (1-20 sec).
* `pmf_sa_query_retry_timeout` - Protected Management Frames (PMF) SA query retry timeout interval (1 - 5 100s of msec).
* `port_macauth` - Enable/disable LAN port MAC authentication (default = disable).
* `port_macauth_reauth_timeout` - LAN port MAC authentication re-authentication timeout value (default = 7200 sec).
* `port_macauth_timeout` - LAN port MAC authentication idle timeout value (default = 600 sec).
* `portal_message_override_group` - Replacement message group for this VAP (only available when security is set to a captive portal type).
* `portal_type` - Captive portal functionality. Configure how the captive portal authenticates users and whether it includes a disclaimer.
* `primary_wag_profile` - Primary wireless access gateway profile name.
* `probe_resp_suppression` - Enable/disable probe response suppression (to ignore weak signals) (default = disable).
* `probe_resp_threshold` - Minimum signal level/threshold in dBm required for the AP response to probe requests (-95 to -20, default = -80).
* `ptk_rekey` - Enable/disable PTK rekey for WPA-Enterprise security.
* `ptk_rekey_intv` - PTK rekey interval (1800 - 864000 sec, default = 86400).
* `qos_profile` - Quality of service profile name.
* `quarantine` - Enable/disable station quarantine (default = enable).
* `radio_2g_threshold` - Minimum signal level/threshold in dBm required for the AP response to receive a packet in 2.4G band (-95 to -20, default = -79).
* `radio_5g_threshold` - Minimum signal level/threshold in dBm required for the AP response to receive a packet in 5G band(-95 to -20, default = -76).
* `radio_sensitivity` - Enable/disable software radio sensitivity (to ignore weak signals) (default = disable).
* `radius_mac_auth` - Enable/disable RADIUS-based MAC authentication of clients (default = disable).
* `radius_mac_auth_block_interval` - Don't send RADIUS MAC auth request again if the client has been rejected within specific interval (0 or 30 - 864000 seconds, default = 0, 0 to disable blocking).
* `radius_mac_auth_server` - RADIUS-based MAC authentication server.
* `radius_mac_mpsk_auth` - Enable/disable RADIUS-based MAC authentication of clients for MPSK authentication (default = disable).
* `radius_mac_mpsk_timeout` - RADIUS MAC MPSK cache timeout interval (0 or 300 - 864000, default = 86400, 0 to disable caching).
* `radius_server` - RADIUS server to be used to authenticate WiFi users.
* `rates_11a` - Allowed data rates for 802.11a.
* `rates_11ac_mcs_map` - Comma separated list of max supported VHT MCS for spatial streams 1 through 8.
* `rates_11ac_ss12` - Allowed data rates for 802.11ac with 1 or 2 spatial streams.
* `rates_11ac_ss34` - Allowed data rates for 802.11ac with 3 or 4 spatial streams.
* `rates_11ax_mcs_map` - Comma separated list of max supported HE MCS for spatial streams 1 through 8.
* `rates_11ax_ss12` - Allowed data rates for 802.11ax with 1 or 2 spatial streams.
* `rates_11ax_ss34` - Allowed data rates for 802.11ax with 3 or 4 spatial streams.
* `rates_11bg` - Allowed data rates for 802.11b/g.
* `rates_11n_ss12` - Allowed data rates for 802.11n with 1 or 2 spatial streams.
* `rates_11n_ss34` - Allowed data rates for 802.11n with 3 or 4 spatial streams.
* `sae_groups` - SAE-Groups.
* `sae_h2e_only` - Use hash-to-element-only mechanism for PWE derivation (default = disable).
* `sae_password` - WPA3 SAE password to be used to authenticate WiFi users.
* `sae_pk` - Enable/disable WPA3 SAE-PK (default = disable).
* `sae_private_key` - Private key used for WPA3 SAE-PK authentication.
* `scan_botnet_connections` - Block or monitor connections to Botnet servers or disable Botnet scanning.
* `secondary_wag_profile` - Secondary wireless access gateway profile name.
* `security` - Security mode for the wireless interface (default = wpa2-only-personal).
* `security_exempt_list` - Optional security exempt list for captive portal authentication.
* `security_redirect_url` - Optional URL for redirecting users after they pass captive portal authentication.
* `split_tunneling` - Enable/disable split tunneling (default = disable).
* `ssid` - IEEE 802.11 service set identifier (SSID) for the wireless interface. Users who wish to use the wireless network must configure their computers to access this SSID name.
* `sticky_client_remove` - Enable/disable sticky client remove to maintain good signal level clients in SSID (default = disable).
* `sticky_client_threshold_2g` - Minimum signal level/threshold in dBm required for the 2G client to be serviced by the AP (-95 to -20, default = -79).
* `sticky_client_threshold_5g` - Minimum signal level/threshold in dBm required for the 5G client to be serviced by the AP (-95 to -20, default = -76).
* `sticky_client_threshold_6g` - Minimum signal level/threshold in dBm required for the 6G client to be serviced by the AP (-95 to -20, default = -76).
* `target_wake_time` - Enable/disable 802.11ax target wake time (default = enable).
* `tkip_counter_measure` - Enable/disable TKIP counter measure.
* `tunnel_echo_interval` - The time interval to send echo to both primary and secondary tunnel peers (1 - 65535 sec, default = 300).
* `tunnel_fallback_interval` - The time interval for secondary tunnel to fall back to primary tunnel (0 - 65535 sec, default = 7200).
* `utm_log` - Enable/disable UTM logging.
* `utm_profile` - UTM profile name.
* `utm_status` - Enable to add one or more security profiles (AV, IPS, etc.) to the VAP.
* `vlan_auto` - Enable/disable automatic management of SSID VLAN interface.
* `vlan_pooling` - Enable/disable VLAN pooling, to allow grouping of multiple wireless controller VLANs into VLAN pools (default = disable). When set to wtp-group, VLAN pooling occurs with VLAN assignment by wtp-group.
* `vlanid` - Optional VLAN ID.
* `voice_enterprise` - Enable/disable 802.11k and 802.11v assisted Voice-Enterprise roaming (default = disable).
* `webfilter_profile` - WebFilter profile name.
* `mac_filter_list` - Create a list of MAC addresses for MAC address filtering.The structure of `mac_filter_list` block is documented below.

The `mac_filter_list` block contains:

* `id` - ID.
* `mac` - MAC address.
* `mac_filter_policy` - Deny or allow the client with this MAC address.
* `mpsk_key` - List of multiple PSK entries.The structure of `mpsk_key` block is documented below.

The `mpsk_key` block contains:

* `comment` - Comment.
* `concurrent_clients` - Number of clients that can connect using this pre-shared key.
* `key_name` - Pre-shared key name.
* `passphrase` - WPA Pre-shared key.
* `mpsk_schedules` - Firewall schedule for MPSK passphrase. The passphrase will be effective only when at least one schedule is valid.The structure of `mpsk_schedules` block is documented below.

The `mpsk_schedules` block contains:

* `name` - Schedule name.
* `portal_message_overrides` - Individual message overrides.The structure of `portal_message_overrides` block is documented below.

The `portal_message_overrides` block contains:

* `auth_disclaimer_page` - Override auth-disclaimer-page message with message from portal-message-overrides group.
* `auth_login_failed_page` - Override auth-login-failed-page message with message from portal-message-overrides group.
* `auth_login_page` - Override auth-login-page message with message from portal-message-overrides group.
* `auth_reject_page` - Override auth-reject-page message with message from portal-message-overrides group.
* `radius_mac_auth_usergroups` - Selective user groups that are permitted for RADIUS mac authentication.The structure of `radius_mac_auth_usergroups` block is documented below.

The `radius_mac_auth_usergroups` block contains:

* `name` - User group name.
* `schedule` - Firewall schedules for enabling this VAP on the FortiAP. This VAP will be enabled when at least one of the schedules is valid. Separate multiple schedule names with a space.The structure of `schedule` block is documented below.

The `schedule` block contains:

* `name` - Schedule name.
* `selected_usergroups` - Selective user groups that are permitted to authenticate.The structure of `selected_usergroups` block is documented below.

The `selected_usergroups` block contains:

* `name` - User group name.
* `usergroup` - Firewall user group to be used to authenticate WiFi users.The structure of `usergroup` block is documented below.

The `usergroup` block contains:

* `name` - User group name.
* `vlan_name` - Table for mapping VLAN name to VLAN ID.The structure of `vlan_name` block is documented below.

The `vlan_name` block contains:

* `name` - VLAN name.
* `vlan_id` - VLAN ID.
* `vlan_pool` - VLAN pool.The structure of `vlan_pool` block is documented below.

The `vlan_pool` block contains:

* `id` - ID.
* `wtp_group` - WTP group name.
