---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_wtpprofile"
description: |-
  Get information on a fortios Configure WTP profiles or FortiAP profiles that define radio settings for manageable FortiAP platforms.
---

# Data Source: fortios_wirelesscontroller_wtpprofile
Use this data source to get information on a fortios Configure WTP profiles or FortiAP profiles that define radio settings for manageable FortiAP platforms.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) WTP (or FortiAP or AP) profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `allowaccess` - Control management access to the managed WTP, FortiAP, or AP. Separate entries with a space.
* `ap_country` - Country in which this WTP, FortiAP or AP will operate (default = NA, automatically use the country configured for the current VDOM).
* `ap_handoff` - Enable/disable AP handoff of clients to other APs (default = disable).
* `apcfg_profile` - AP local configuration profile name.
* `ble_profile` - Bluetooth Low Energy profile name.
* `comment` - Comment.
* `console_login` - Enable/disable FAP console login access (default = enable).
* `control_message_offload` - Enable/disable CAPWAP control message data channel offload.
* `dtls_in_kernel` - Enable/disable data channel DTLS in kernel.
* `dtls_policy` - WTP data channel DTLS policy (default = clear-text).
* `energy_efficient_ethernet` - Enable/disable use of energy efficient Ethernet on WTP.
* `ext_info_enable` - Enable/disable station/VAP/radio extension information.
* `frequency_handoff` - Enable/disable frequency handoff of clients to other channels (default = disable).
* `handoff_roaming` - Enable/disable client load balancing during roaming to avoid roaming delay (default = disable).
* `handoff_rssi` - Minimum received signal strength indicator (RSSI) value for handoff (20 - 30, default = 25).
* `handoff_sta_thresh` - Threshold value for AP handoff.
* `indoor_outdoor_deployment` - Set to allow indoor/outdoor-only channels under regulatory rules (default = platform-determined).
* `ip_fragment_preventing` - Method(s) by which IP fragmentation is prevented for control and data packets through CAPWAP tunnel (default = tcp-mss-adjust).
* `led_state` - Enable/disable use of LEDs on WTP (default = disable).
* `lldp` - Enable/disable Link Layer Discovery Protocol (LLDP) for the WTP, FortiAP, or AP (default = enable).
* `login_passwd` - Set the managed WTP, FortiAP, or AP's administrator password.
* `login_passwd_change` - Change or reset the administrator password of a managed WTP, FortiAP or AP (yes, default, or no, default = no).
* `max_clients` - Maximum number of stations (STAs) supported by the WTP (default = 0, meaning no client limitation).
* `name` - WTP (or FortiAP or AP) profile name.
* `poe_mode` - Set the WTP, FortiAP, or AP's PoE mode.
* `split_tunneling_acl_local_ap_subnet` - Enable/disable automatically adding local subnetwork of FortiAP to split-tunneling ACL (default = disable).
* `split_tunneling_acl_path` - Split tunneling ACL path is local/tunnel.
* `syslog_profile` - System log server configuration profile name.
* `tun_mtu_downlink` - The MTU of downlink CAPWAP tunnel (576 - 1500 bytes or 0; 0 means the local MTU of FortiAP; default = 0).
* `tun_mtu_uplink` - The maximum transmission unit (MTU) of uplink CAPWAP tunnel (576 - 1500 bytes or 0; 0 means the local MTU of FortiAP; default = 0).
* `wan_port_auth` - Set WAN port authentication mode (default = none).
* `wan_port_auth_methods` - WAN port 802.1x supplicant EAP methods (default = all).
* `wan_port_auth_password` - Set WAN port 802.1x supplicant password.
* `wan_port_auth_usrname` - Set WAN port 802.1x supplicant user name.
* `wan_port_mode` - Enable/disable using a WAN port as a LAN port.
* `deny_mac_list` - List of MAC addresses that are denied access to this WTP, FortiAP, or AP.The structure of `deny_mac_list` block is documented below.

The `deny_mac_list` block contains:

* `id` - ID.
* `mac` - A WiFi device with this MAC address is denied access to this WTP, FortiAP or AP.
* `esl_ses_dongle` - ESL SES-imagotag dongle configuration.The structure of `esl_ses_dongle` block is documented below.

The `esl_ses_dongle` block contains:

* `apc_addr_type` - ESL SES-imagotag APC address type (default = fqdn).
* `apc_fqdn` - FQDN of ESL SES-imagotag Access Point Controller (APC).
* `apc_ip` - IP address of ESL SES-imagotag Access Point Controller (APC).
* `apc_port` - Port of ESL SES-imagotag Access Point Controller (APC).
* `coex_level` - ESL SES-imagotag dongle coexistence level (default = none).
* `compliance_level` - Compliance levels for the ESL solution integration (default = compliance-level-2).
* `esl_channel` - ESL SES-imagotag dongle channel (default = 127).
* `output_power` - ESL SES-imagotag dongle output power (default = A).
* `scd_enable` - Enable/disable ESL SES-imagotag Serial Communication Daemon (SCD) (default = disable).
* `tls_cert_verification` - Enable/disable TLS Certificate verification. (default = enable).
* `tls_fqdn_verification` - Enable/disable TLS Certificate verification. (default = disable).
* `lan` - WTP LAN port mapping.The structure of `lan` block is documented below.

The `lan` block contains:

* `port_esl_mode` - ESL port mode.
* `port_esl_ssid` - Bridge ESL port to SSID.
* `port_mode` - LAN port mode.
* `port_ssid` - Bridge LAN port to SSID.
* `port1_mode` - LAN port 1 mode.
* `port1_ssid` - Bridge LAN port 1 to SSID.
* `port2_mode` - LAN port 2 mode.
* `port2_ssid` - Bridge LAN port 2 to SSID.
* `port3_mode` - LAN port 3 mode.
* `port3_ssid` - Bridge LAN port 3 to SSID.
* `port4_mode` - LAN port 4 mode.
* `port4_ssid` - Bridge LAN port 4 to SSID.
* `port5_mode` - LAN port 5 mode.
* `port5_ssid` - Bridge LAN port 5 to SSID.
* `port6_mode` - LAN port 6 mode.
* `port6_ssid` - Bridge LAN port 6 to SSID.
* `port7_mode` - LAN port 7 mode.
* `port7_ssid` - Bridge LAN port 7 to SSID.
* `port8_mode` - LAN port 8 mode.
* `port8_ssid` - Bridge LAN port 8 to SSID.
* `lbs` - Set various location based service (LBS) options.The structure of `lbs` block is documented below.

The `lbs` block contains:

* `aeroscout` - Enable/disable AeroScout Real Time Location Service (RTLS) support (default = disable).
* `aeroscout_ap_mac` - Use BSSID or board MAC address as AP MAC address in AeroScout AP messages (default = bssid).
* `aeroscout_mmu_report` - Enable/disable compounded AeroScout tag and MU report (default = enable).
* `aeroscout_mu` - Enable/disable AeroScout Mobile Unit (MU) support (default = disable).
* `aeroscout_mu_factor` - AeroScout MU mode dilution factor (default = 20).
* `aeroscout_mu_timeout` - AeroScout MU mode timeout (0 - 65535 sec, default = 5).
* `aeroscout_server_ip` - IP address of AeroScout server.
* `aeroscout_server_port` - AeroScout server UDP listening port.
* `ekahau_blink_mode` - Enable/disable Ekahau blink mode (now known as AiRISTA Flow) to track and locate WiFi tags (default = disable).
* `ekahau_tag` - WiFi frame MAC address or WiFi Tag.
* `erc_server_ip` - IP address of Ekahau RTLS Controller (ERC).
* `erc_server_port` - Ekahau RTLS Controller (ERC) UDP listening port.
* `fortipresence` - Enable/disable FortiPresence to monitor the location and activity of WiFi clients even if they don't connect to this WiFi network (default = disable).
* `fortipresence_ble` - Enable/disable FortiPresence finding and reporting BLE devices.
* `fortipresence_frequency` - FortiPresence report transmit frequency (5 - 65535 sec, default = 30).
* `fortipresence_port` - UDP listening port of FortiPresence server (default = 3000).
* `fortipresence_project` - FortiPresence project name (max. 16 characters, default = fortipresence).
* `fortipresence_rogue` - Enable/disable FortiPresence finding and reporting rogue APs.
* `fortipresence_secret` - FortiPresence secret password (max. 16 characters).
* `fortipresence_server` - IP address of FortiPresence server.
* `fortipresence_server_addr_type` - FortiPresence server address type (default = ipv4).
* `fortipresence_server_fqdn` - FQDN of FortiPresence server.
* `fortipresence_unassoc` - Enable/disable FortiPresence finding and reporting unassociated stations.
* `station_locate` - Enable/disable client station locating services for all clients, whether associated or not (default = disable).
* `led_schedules` - Recurring firewall schedules for illuminating LEDs on the FortiAP. If led-state is enabled, LEDs will be visible when at least one of the schedules is valid. Separate multiple schedule names with a space.The structure of `led_schedules` block is documented below.

The `led_schedules` block contains:

* `name` - Schedule name.
* `platform` - WTP, FortiAP, or AP platform.The structure of `platform` block is documented below.

The `platform` block contains:

* `ddscan` - Enable/disable use of one radio for dedicated dual-band scanning to detect RF characterization and wireless threat management.
* `mode` - Configure operation mode of 5G radios (default = single-5G).
* `type` - WTP, FortiAP or AP platform type. There are built-in WTP profiles for all supported FortiAP models. You can select a built-in profile and customize it or create a new profile.
* `radio_1` - Configuration options for radio 1.The structure of `radio_1` block is documented below.

The `radio_1` block contains:

* `airtime_fairness` - Enable/disable airtime fairness (default = disable).
* `amsdu` - Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable).
* `ap_handoff` - Enable/disable AP handoff of clients to other APs (default = disable).
* `ap_sniffer_addr` - MAC address to monitor.
* `ap_sniffer_bufsize` - Sniffer buffer size (1 - 32 MB, default = 16).
* `ap_sniffer_chan` - Channel on which to operate the sniffer (default = 6).
* `ap_sniffer_ctl` - Enable/disable sniffer on WiFi control frame (default = enable).
* `ap_sniffer_data` - Enable/disable sniffer on WiFi data frame (default = enable).
* `ap_sniffer_mgmt_beacon` - Enable/disable sniffer on WiFi management Beacon frames (default = enable).
* `ap_sniffer_mgmt_other` - Enable/disable sniffer on WiFi management other frames  (default = enable).
* `ap_sniffer_mgmt_probe` - Enable/disable sniffer on WiFi management probe frames (default = enable).
* `auto_power_high` - The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable).
* `auto_power_low` - The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_target` - The target of automatic transmit power adjustment in dBm. (-95 to -20, default = -70).
* `band` - WiFi band that Radio 1 operates on.
* `band_5g_type` - WiFi 5G band type.
* `bandwidth_admission_control` - Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it.
* `bandwidth_capacity` - Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).
* `beacon_interval` - Beacon interval. The time between beacon frames in msec (the actual range of beacon interval depends on the AP platform type, default = 100).
* `bss_color` - BSS color value for this 11ax radio (0 - 63, 0 means disable. default = 0).
* `bss_color_mode` - BSS color mode for this 11ax radio (default = auto).
* `call_admission_control` - Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them.
* `call_capacity` - Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).
* `channel_bonding` - Channel bandwidth: 160,80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence.
* `channel_utilization` - Enable/disable measuring channel utilization.
* `coexistence` - Enable/disable allowing both HT20 and HT40 on the same radio (default = enable).
* `darrp` - Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable).
* `drma` - Enable/disable dynamic radio mode assignment (DRMA) (default = disable).
* `drma_sensitivity` - Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low).
* `dtim` - Delivery Traffic Indication Map (DTIM) period (1 - 255, default = 1). Set higher to save battery life of WiFi client in power-save mode.
* `frag_threshold` - Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).
* `frequency_handoff` - Enable/disable frequency handoff of clients to other channels (default = disable).
* `iperf_protocol` - Iperf test protocol (default = "UDP").
* `iperf_server_port` - Iperf service port number.
* `max_clients` - Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.
* `max_distance` - Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).
* `mode` - Mode of radio 1. Radio 1 can be disabled, configured as an access point, a rogue AP monitor, a sniffer, or a station.
* `power_level` - Radio EIRP power level as a percentage of the maximum EIRP power (0 - 100, default = 100).
* `power_mode` - Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities.
* `power_value` - Radio EIRP power in dBm (1 - 33, default = 27).
* `powersave_optimize` - Enable client power-saving features such as TIM, AC VO, and OBSS etc.
* `protection_mode` - Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable).
* `rts_threshold` - Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).
* `sam_bssid` - BSSID for WiFi network.
* `sam_captive_portal` - Enable/disable Captive Portal Authentication (default = disable).
* `sam_cwp_failure_string` - Failure identification on the page after an incorrect login.
* `sam_cwp_match_string` - Identification string from the captive portal login form.
* `sam_cwp_password` - Password for captive portal authentication.
* `sam_cwp_success_string` - Success identification on the page after a successful login.
* `sam_cwp_test_url` - Website the client is trying to access.
* `sam_cwp_username` - Username for captive portal authentication.
* `sam_password` - Passphrase for WiFi network connection.
* `sam_report_intv` - SAM report interval (sec), 0 for a one-time report.
* `sam_security_type` - Select WiFi network security type (default = "wpa-personal").
* `sam_server` - SAM test server IP address or domain name.
* `sam_server_fqdn` - SAM test server domain name.
* `sam_server_ip` - SAM test server IP address.
* `sam_server_type` - Select SAM server type (default = "IP").
* `sam_ssid` - SSID for WiFi network.
* `sam_test` - Select SAM test type (default = "PING").
* `sam_username` - Username for WiFi network connection.
* `short_guard_interval` - Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns.
* `spectrum_analysis` - Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.
* `transmit_optimize` - Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default.
* `vap_all` - Configure method for assigning SSIDs to this FortiAP (default = automatically assign tunnel SSIDs).
* `wids_profile` - Wireless Intrusion Detection System (WIDS) profile name to assign to the radio.
* `zero_wait_dfs` - Enable/disable zero wait DFS on radio (default = enable).
* `channel` - Selected list of wireless radio channels.The structure of `channel` block is documented below.

The `channel` block contains:

* `chan` - Channel number.
* `vaps` - Manually selected list of Virtual Access Points (VAPs).The structure of `vaps` block is documented below.

The `vaps` block contains:

* `name` - Virtual Access Point (VAP) name.
* `radio_2` - Configuration options for radio 2.The structure of `radio_2` block is documented below.

The `radio_2` block contains:

* `airtime_fairness` - Enable/disable airtime fairness (default = disable).
* `amsdu` - Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable).
* `ap_handoff` - Enable/disable AP handoff of clients to other APs (default = disable).
* `ap_sniffer_addr` - MAC address to monitor.
* `ap_sniffer_bufsize` - Sniffer buffer size (1 - 32 MB, default = 16).
* `ap_sniffer_chan` - Channel on which to operate the sniffer (default = 6).
* `ap_sniffer_ctl` - Enable/disable sniffer on WiFi control frame (default = enable).
* `ap_sniffer_data` - Enable/disable sniffer on WiFi data frame (default = enable).
* `ap_sniffer_mgmt_beacon` - Enable/disable sniffer on WiFi management Beacon frames (default = enable).
* `ap_sniffer_mgmt_other` - Enable/disable sniffer on WiFi management other frames  (default = enable).
* `ap_sniffer_mgmt_probe` - Enable/disable sniffer on WiFi management probe frames (default = enable).
* `auto_power_high` - The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable).
* `auto_power_low` - The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_target` - The target of automatic transmit power adjustment in dBm. (-95 to -20, default = -70).
* `band` - WiFi band that Radio 2 operates on.
* `band_5g_type` - WiFi 5G band type.
* `bandwidth_admission_control` - Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it.
* `bandwidth_capacity` - Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).
* `beacon_interval` - Beacon interval. The time between beacon frames in msec (the actual range of beacon interval depends on the AP platform type, default = 100).
* `bss_color` - BSS color value for this 11ax radio (0 - 63, 0 means disable. default = 0).
* `bss_color_mode` - BSS color mode for this 11ax radio (default = auto).
* `call_admission_control` - Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them.
* `call_capacity` - Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).
* `channel_bonding` - Channel bandwidth: 160,80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence.
* `channel_utilization` - Enable/disable measuring channel utilization.
* `coexistence` - Enable/disable allowing both HT20 and HT40 on the same radio (default = enable).
* `darrp` - Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable).
* `drma` - Enable/disable dynamic radio mode assignment (DRMA) (default = disable).
* `drma_sensitivity` - Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low).
* `dtim` - Delivery Traffic Indication Map (DTIM) period (1 - 255, default = 1). Set higher to save battery life of WiFi client in power-save mode.
* `frag_threshold` - Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).
* `frequency_handoff` - Enable/disable frequency handoff of clients to other channels (default = disable).
* `iperf_protocol` - Iperf test protocol (default = "UDP").
* `iperf_server_port` - Iperf service port number.
* `max_clients` - Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.
* `max_distance` - Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).
* `mode` - Mode of radio 2. Radio 2 can be disabled, configured as an access point, a rogue AP monitor, a sniffer, or a station.
* `power_level` - Radio EIRP power level as a percentage of the maximum EIRP power (0 - 100, default = 100).
* `power_mode` - Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities.
* `power_value` - Radio EIRP power in dBm (1 - 33, default = 27).
* `powersave_optimize` - Enable client power-saving features such as TIM, AC VO, and OBSS etc.
* `protection_mode` - Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable).
* `rts_threshold` - Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).
* `sam_bssid` - BSSID for WiFi network.
* `sam_captive_portal` - Enable/disable Captive Portal Authentication (default = disable).
* `sam_cwp_failure_string` - Failure identification on the page after an incorrect login.
* `sam_cwp_match_string` - Identification string from the captive portal login form.
* `sam_cwp_password` - Password for captive portal authentication.
* `sam_cwp_success_string` - Success identification on the page after a successful login.
* `sam_cwp_test_url` - Website the client is trying to access.
* `sam_cwp_username` - Username for captive portal authentication.
* `sam_password` - Passphrase for WiFi network connection.
* `sam_report_intv` - SAM report interval (sec), 0 for a one-time report.
* `sam_security_type` - Select WiFi network security type (default = "wpa-personal").
* `sam_server` - SAM test server IP address or domain name.
* `sam_server_fqdn` - SAM test server domain name.
* `sam_server_ip` - SAM test server IP address.
* `sam_server_type` - Select SAM server type (default = "IP").
* `sam_ssid` - SSID for WiFi network.
* `sam_test` - Select SAM test type (default = "PING").
* `sam_username` - Username for WiFi network connection.
* `short_guard_interval` - Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns.
* `spectrum_analysis` - Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.
* `transmit_optimize` - Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default.
* `vap_all` - Configure method for assigning SSIDs to this FortiAP (default = automatically assign tunnel SSIDs).
* `wids_profile` - Wireless Intrusion Detection System (WIDS) profile name to assign to the radio.
* `zero_wait_dfs` - Enable/disable zero wait DFS on radio (default = enable).
* `channel` - Selected list of wireless radio channels.The structure of `channel` block is documented below.

The `channel` block contains:

* `chan` - Channel number.
* `vaps` - Manually selected list of Virtual Access Points (VAPs).The structure of `vaps` block is documented below.

The `vaps` block contains:

* `name` - Virtual Access Point (VAP) name.
* `radio_3` - Configuration options for radio 3.The structure of `radio_3` block is documented below.

The `radio_3` block contains:

* `airtime_fairness` - Enable/disable airtime fairness (default = disable).
* `amsdu` - Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable).
* `ap_handoff` - Enable/disable AP handoff of clients to other APs (default = disable).
* `ap_sniffer_addr` - MAC address to monitor.
* `ap_sniffer_bufsize` - Sniffer buffer size (1 - 32 MB, default = 16).
* `ap_sniffer_chan` - Channel on which to operate the sniffer (default = 6).
* `ap_sniffer_ctl` - Enable/disable sniffer on WiFi control frame (default = enable).
* `ap_sniffer_data` - Enable/disable sniffer on WiFi data frame (default = enable).
* `ap_sniffer_mgmt_beacon` - Enable/disable sniffer on WiFi management Beacon frames (default = enable).
* `ap_sniffer_mgmt_other` - Enable/disable sniffer on WiFi management other frames  (default = enable).
* `ap_sniffer_mgmt_probe` - Enable/disable sniffer on WiFi management probe frames (default = enable).
* `auto_power_high` - The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable).
* `auto_power_low` - The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_target` - The target of automatic transmit power adjustment in dBm. (-95 to -20, default = -70).
* `band` - WiFi band that Radio 3 operates on.
* `band_5g_type` - WiFi 5G band type.
* `bandwidth_admission_control` - Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it.
* `bandwidth_capacity` - Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).
* `beacon_interval` - Beacon interval. The time between beacon frames in msec (the actual range of beacon interval depends on the AP platform type, default = 100).
* `bss_color` - BSS color value for this 11ax radio (0 - 63, 0 means disable. default = 0).
* `bss_color_mode` - BSS color mode for this 11ax radio (default = auto).
* `call_admission_control` - Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them.
* `call_capacity` - Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).
* `channel_bonding` - Channel bandwidth: 160,80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence.
* `channel_utilization` - Enable/disable measuring channel utilization.
* `coexistence` - Enable/disable allowing both HT20 and HT40 on the same radio (default = enable).
* `darrp` - Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable).
* `drma` - Enable/disable dynamic radio mode assignment (DRMA) (default = disable).
* `drma_sensitivity` - Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low).
* `dtim` - Delivery Traffic Indication Map (DTIM) period (1 - 255, default = 1). Set higher to save battery life of WiFi client in power-save mode.
* `frag_threshold` - Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).
* `frequency_handoff` - Enable/disable frequency handoff of clients to other channels (default = disable).
* `iperf_protocol` - Iperf test protocol (default = "UDP").
* `iperf_server_port` - Iperf service port number.
* `max_clients` - Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.
* `max_distance` - Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).
* `mode` - Mode of radio 3. Radio 3 can be disabled, configured as an access point, a rogue AP monitor, a sniffer, or a station.
* `power_level` - Radio EIRP power level as a percentage of the maximum EIRP power (0 - 100, default = 100).
* `power_mode` - Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities.
* `power_value` - Radio EIRP power in dBm (1 - 33, default = 27).
* `powersave_optimize` - Enable client power-saving features such as TIM, AC VO, and OBSS etc.
* `protection_mode` - Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable).
* `rts_threshold` - Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).
* `sam_bssid` - BSSID for WiFi network.
* `sam_captive_portal` - Enable/disable Captive Portal Authentication (default = disable).
* `sam_cwp_failure_string` - Failure identification on the page after an incorrect login.
* `sam_cwp_match_string` - Identification string from the captive portal login form.
* `sam_cwp_password` - Password for captive portal authentication.
* `sam_cwp_success_string` - Success identification on the page after a successful login.
* `sam_cwp_test_url` - Website the client is trying to access.
* `sam_cwp_username` - Username for captive portal authentication.
* `sam_password` - Passphrase for WiFi network connection.
* `sam_report_intv` - SAM report interval (sec), 0 for a one-time report.
* `sam_security_type` - Select WiFi network security type (default = "wpa-personal").
* `sam_server` - SAM test server IP address or domain name.
* `sam_server_fqdn` - SAM test server domain name.
* `sam_server_ip` - SAM test server IP address.
* `sam_server_type` - Select SAM server type (default = "IP").
* `sam_ssid` - SSID for WiFi network.
* `sam_test` - Select SAM test type (default = "PING").
* `sam_username` - Username for WiFi network connection.
* `short_guard_interval` - Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns.
* `spectrum_analysis` - Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.
* `transmit_optimize` - Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default.
* `vap_all` - Configure method for assigning SSIDs to this FortiAP (default = automatically assign tunnel SSIDs).
* `wids_profile` - Wireless Intrusion Detection System (WIDS) profile name to assign to the radio.
* `zero_wait_dfs` - Enable/disable zero wait DFS on radio (default = enable).
* `channel` - Selected list of wireless radio channels.The structure of `channel` block is documented below.

The `channel` block contains:

* `chan` - Channel number.
* `vaps` - Manually selected list of Virtual Access Points (VAPs).The structure of `vaps` block is documented below.

The `vaps` block contains:

* `name` - Virtual Access Point (VAP) name.
* `radio_4` - Configuration options for radio 4.The structure of `radio_4` block is documented below.

The `radio_4` block contains:

* `airtime_fairness` - Enable/disable airtime fairness (default = disable).
* `amsdu` - Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable).
* `ap_handoff` - Enable/disable AP handoff of clients to other APs (default = disable).
* `ap_sniffer_addr` - MAC address to monitor.
* `ap_sniffer_bufsize` - Sniffer buffer size (1 - 32 MB, default = 16).
* `ap_sniffer_chan` - Channel on which to operate the sniffer (default = 6).
* `ap_sniffer_ctl` - Enable/disable sniffer on WiFi control frame (default = enable).
* `ap_sniffer_data` - Enable/disable sniffer on WiFi data frame (default = enable).
* `ap_sniffer_mgmt_beacon` - Enable/disable sniffer on WiFi management Beacon frames (default = enable).
* `ap_sniffer_mgmt_other` - Enable/disable sniffer on WiFi management other frames  (default = enable).
* `ap_sniffer_mgmt_probe` - Enable/disable sniffer on WiFi management probe frames (default = enable).
* `auto_power_high` - The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable).
* `auto_power_low` - The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_target` - The target of automatic transmit power adjustment in dBm. (-95 to -20, default = -70).
* `band` - WiFi band that Radio 3 operates on.
* `band_5g_type` - WiFi 5G band type.
* `bandwidth_admission_control` - Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it.
* `bandwidth_capacity` - Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).
* `beacon_interval` - Beacon interval. The time between beacon frames in msec (the actual range of beacon interval depends on the AP platform type, default = 100).
* `bss_color` - BSS color value for this 11ax radio (0 - 63, 0 means disable. default = 0).
* `bss_color_mode` - BSS color mode for this 11ax radio (default = auto).
* `call_admission_control` - Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them.
* `call_capacity` - Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).
* `channel_bonding` - Channel bandwidth: 160,80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence.
* `channel_utilization` - Enable/disable measuring channel utilization.
* `coexistence` - Enable/disable allowing both HT20 and HT40 on the same radio (default = enable).
* `darrp` - Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable).
* `drma` - Enable/disable dynamic radio mode assignment (DRMA) (default = disable).
* `drma_sensitivity` - Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low).
* `dtim` - Delivery Traffic Indication Map (DTIM) period (1 - 255, default = 1). Set higher to save battery life of WiFi client in power-save mode.
* `frag_threshold` - Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).
* `frequency_handoff` - Enable/disable frequency handoff of clients to other channels (default = disable).
* `iperf_protocol` - Iperf test protocol (default = "UDP").
* `iperf_server_port` - Iperf service port number.
* `max_clients` - Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.
* `max_distance` - Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).
* `mode` - Mode of radio 3. Radio 3 can be disabled, configured as an access point, a rogue AP monitor, a sniffer, or a station.
* `power_level` - Radio EIRP power level as a percentage of the maximum EIRP power (0 - 100, default = 100).
* `power_mode` - Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities.
* `power_value` - Radio EIRP power in dBm (1 - 33, default = 27).
* `powersave_optimize` - Enable client power-saving features such as TIM, AC VO, and OBSS etc.
* `protection_mode` - Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable).
* `rts_threshold` - Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).
* `sam_bssid` - BSSID for WiFi network.
* `sam_captive_portal` - Enable/disable Captive Portal Authentication (default = disable).
* `sam_cwp_failure_string` - Failure identification on the page after an incorrect login.
* `sam_cwp_match_string` - Identification string from the captive portal login form.
* `sam_cwp_password` - Password for captive portal authentication.
* `sam_cwp_success_string` - Success identification on the page after a successful login.
* `sam_cwp_test_url` - Website the client is trying to access.
* `sam_cwp_username` - Username for captive portal authentication.
* `sam_password` - Passphrase for WiFi network connection.
* `sam_report_intv` - SAM report interval (sec), 0 for a one-time report.
* `sam_security_type` - Select WiFi network security type (default = "wpa-personal").
* `sam_server` - SAM test server IP address or domain name.
* `sam_server_fqdn` - SAM test server domain name.
* `sam_server_ip` - SAM test server IP address.
* `sam_server_type` - Select SAM server type (default = "IP").
* `sam_ssid` - SSID for WiFi network.
* `sam_test` - Select SAM test type (default = "PING").
* `sam_username` - Username for WiFi network connection.
* `short_guard_interval` - Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns.
* `spectrum_analysis` - Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.
* `transmit_optimize` - Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default.
* `vap_all` - Configure method for assigning SSIDs to this FortiAP (default = automatically assign tunnel SSIDs).
* `wids_profile` - Wireless Intrusion Detection System (WIDS) profile name to assign to the radio.
* `zero_wait_dfs` - Enable/disable zero wait DFS on radio (default = enable).
* `channel` - Selected list of wireless radio channels.The structure of `channel` block is documented below.

The `channel` block contains:

* `chan` - Channel number.
* `vaps` - Manually selected list of Virtual Access Points (VAPs).The structure of `vaps` block is documented below.

The `vaps` block contains:

* `name` - Virtual Access Point (VAP) name.
* `split_tunneling_acl` - Split tunneling ACL filter list.The structure of `split_tunneling_acl` block is documented below.

The `split_tunneling_acl` block contains:

* `dest_ip` - Destination IP and mask for the split-tunneling subnet.
* `id` - ID.
