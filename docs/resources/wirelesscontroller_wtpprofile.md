---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_wtpprofile"
description: |-
  Configure WTP profiles or FortiAP profiles that define radio settings for manageable FortiAP platforms.
---

## fortios_wirelesscontroller_wtpprofile
Configure WTP profiles or FortiAP profiles that define radio settings for manageable FortiAP platforms.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `allowaccess` - Control management access to the managed WTP, FortiAP, or AP. Separate entries with a space. Valid values: `https` `ssh` `snmp` .
* `ap_country` - Country in which this WTP, FortiAP, or AP will operate (default = NA, automatically use the country configured for the current VDOM). Valid values: `--` `AF` `AL` `DZ` `AS` `AO` `AR` `AM` `AU` `AT` `AZ` `BS` `BH` `BD` `BB` `BY` `BE` `BZ` `BJ` `BM` `BT` `BO` `BA` `BW` `BR` `BN` `BG` `BF` `KH` `CM` `KY` `CF` `TD` `CL` `CN` `CX` `CO` `CG` `CD` `CR` `HR` `CY` `CZ` `DK` `DM` `DO` `EC` `EG` `SV` `ET` `EE` `GF` `PF` `FO` `FJ` `FI` `FR` `GE` `DE` `GH` `GI` `GR` `GL` `GD` `GP` `GU` `GT` `GY` `HT` `HN` `HK` `HU` `IS` `IN` `ID` `IQ` `IE` `IM` `IL` `IT` `CI` `JM` `JO` `KZ` `KE` `KR` `KW` `LA` `LV` `LB` `LS` `LY` `LI` `LT` `LU` `MO` `MK` `MG` `MW` `MY` `MV` `ML` `MT` `MH` `MQ` `MR` `MU` `YT` `MX` `FM` `MD` `MC` `MA` `MZ` `MM` `NA` `NP` `NL` `AN` `AW` `NZ` `NI` `NE` `NO` `MP` `OM` `PK` `PW` `PA` `PG` `PY` `PE` `PH` `PL` `PT` `PR` `QA` `RE` `RO` `RU` `RW` `BL` `KN` `LC` `MF` `PM` `VC` `SA` `SN` `RS` `ME` `SL` `SG` `SK` `SI` `ZA` `ES` `LK` `SE` `SR` `CH` `TW` `TZ` `TH` `TG` `TT` `TN` `TR` `TM` `AE` `TC` `UG` `UA` `GB` `US` `PS` `UY` `UZ` `VU` `VE` `VN` `VI` `WF` `YE` `ZM` `ZW` `JP` `CA` .
* `ap_handoff` - Enable/disable AP handoff of clients to other APs (default = disable). Valid values: `enable` `disable` .
* `apcfg_profile` - AP local configuration profile name. This attribute must reference one of the following datasources: `wireless-controller.apcfg-profile.name` .
* `ble_profile` - Bluetooth Low Energy profile name. This attribute must reference one of the following datasources: `wireless-controller.ble-profile.name` .
* `comment` - Comment.
* `console_login` - Enable/disable FortiAP console login access (default = enable). Valid values: `enable` `disable` .
* `control_message_offload` - Enable/disable CAPWAP control message data channel offload. Valid values: `ebp-frame` `aeroscout-tag` `ap-list` `sta-list` `sta-cap-list` `stats` `aeroscout-mu` `sta-health` `spectral-analysis` .
* `dtls_in_kernel` - Enable/disable data channel DTLS in kernel. Valid values: `enable` `disable` .
* `dtls_policy` - WTP data channel DTLS policy (default = clear-text). Valid values: `clear-text` `dtls-enabled` `ipsec-vpn` .
* `energy_efficient_ethernet` - Enable/disable use of energy efficient Ethernet on WTP. Valid values: `enable` `disable` .
* `ext_info_enable` - Enable/disable station/VAP/radio extension information. Valid values: `enable` `disable` .
* `frequency_handoff` - Enable/disable frequency handoff of clients to other channels (default = disable). Valid values: `enable` `disable` .
* `handoff_roaming` - Enable/disable client load balancing during roaming to avoid roaming delay (default = disable). Valid values: `enable` `disable` .
* `handoff_rssi` - Minimum received signal strength indicator (RSSI) value for handoff (20 - 30, default = 25).
* `handoff_sta_thresh` - Threshold value for AP handoff.
* `indoor_outdoor_deployment` - Set to allow indoor/outdoor-only channels under regulatory rules (default = platform-determined). Valid values: `platform-determined` `outdoor` `indoor` .
* `ip_fragment_preventing` - Method(s) by which IP fragmentation is prevented for control and data packets through CAPWAP tunnel (default = tcp-mss-adjust). Valid values: `tcp-mss-adjust` `icmp-unreachable` .
* `led_state` - Enable/disable use of LEDs on WTP (default = disable). Valid values: `enable` `disable` .
* `lldp` - Enable/disable Link Layer Discovery Protocol (LLDP) for the WTP, FortiAP, or AP (default = enable). Valid values: `enable` `disable` .
* `login_passwd` - Set the managed WTP, FortiAP, or AP's administrator password.
* `login_passwd_change` - Change or reset the administrator password of a managed WTP, FortiAP or AP (yes, default, or no, default = no). Valid values: `yes` `default` `no` .
* `max_clients` - Maximum number of stations (STAs) supported by the WTP (default = 0, meaning no client limitation).
* `name` - WTP (or FortiAP or AP) profile name.
* `poe_mode` - Set the WTP, FortiAP, or AP's PoE mode. Valid values: `auto` `8023af` `8023at` `power-adapter` `full` `high` `low` .
* `split_tunneling_acl_local_ap_subnet` - Enable/disable automatically adding local subnetwork of FortiAP to split-tunneling ACL (default = disable). Valid values: `enable` `disable` .
* `split_tunneling_acl_path` - Split tunneling ACL path is local/tunnel. Valid values: `tunnel` `local` .
* `syslog_profile` - System log server configuration profile name. This attribute must reference one of the following datasources: `wireless-controller.syslog-profile.name` .
* `tun_mtu_downlink` - The MTU of downlink CAPWAP tunnel (576 - 1500 bytes or 0; 0 means the local MTU of FortiAP; default = 0).
* `tun_mtu_uplink` - The maximum transmission unit (MTU) of uplink CAPWAP tunnel (576 - 1500 bytes or 0; 0 means the local MTU of FortiAP; default = 0).
* `wan_port_auth` - Set WAN port authentication mode (default = none). Valid values: `none` `802.1x` .
* `wan_port_auth_methods` - WAN port 802.1x supplicant EAP methods (default = all). Valid values: `all` `EAP-FAST` `EAP-TLS` `EAP-PEAP` .
* `wan_port_auth_password` - Set WAN port 802.1x supplicant password.
* `wan_port_auth_usrname` - Set WAN port 802.1x supplicant user name.
* `wan_port_mode` - Enable/disable using a WAN port as a LAN port. Valid values: `wan-lan` `wan-only` .
* `deny_mac_list` - List of MAC addresses that are denied access to this WTP, FortiAP, or AP. The structure of `deny_mac_list` block is documented below.

The `deny_mac_list` block contains:

* `id` - ID.
* `mac` - A WiFi device with this MAC address is denied access to this WTP, FortiAP or AP.
* `esl_ses_dongle` - ESL SES-imagotag dongle configuration. The structure of `esl_ses_dongle` block is documented below.

The `esl_ses_dongle` block contains:

* `apc_addr_type` - ESL SES-imagotag APC address type (default = fqdn). Valid values: `fqdn` `ip` .
* `apc_fqdn` - FQDN of ESL SES-imagotag Access Point Controller (APC).
* `apc_ip` - IP address of ESL SES-imagotag Access Point Controller (APC).
* `apc_port` - Port of ESL SES-imagotag Access Point Controller (APC).
* `coex_level` - ESL SES-imagotag dongle coexistence level (default = none). Valid values: `none` .
* `compliance_level` - Compliance levels for the ESL solution integration (default = compliance-level-2). Valid values: `compliance-level-2` .
* `esl_channel` - ESL SES-imagotag dongle channel (default = 127). Valid values: `-1` `0` `1` `2` `3` `4` `5` `6` `7` `8` `9` `10` `127` .
* `output_power` - ESL SES-imagotag dongle output power (default = A). Valid values: `a` `b` `c` `d` `e` `f` `g` `h` .
* `scd_enable` - Enable/disable ESL SES-imagotag Serial Communication Daemon (SCD) (default = disable). Valid values: `enable` `disable` .
* `tls_cert_verification` - Enable/disable TLS certificate verification (default = enable). Valid values: `enable` `disable` .
* `tls_fqdn_verification` - Enable/disable TLS certificate verification (default = disable). Valid values: `enable` `disable` .
* `lan` - WTP LAN port mapping. The structure of `lan` block is documented below.

The `lan` block contains:

* `port_esl_mode` - ESL port mode. Valid values: `offline` `nat-to-wan` `bridge-to-wan` `bridge-to-ssid` .
* `port_esl_ssid` - Bridge ESL port to SSID. This attribute must reference one of the following datasources: `system.interface.name` .
* `port_mode` - LAN port mode. Valid values: `offline` `nat-to-wan` `bridge-to-wan` `bridge-to-ssid` .
* `port_ssid` - Bridge LAN port to SSID. This attribute must reference one of the following datasources: `system.interface.name` .
* `port1_mode` - LAN port 1 mode. Valid values: `offline` `nat-to-wan` `bridge-to-wan` `bridge-to-ssid` .
* `port1_ssid` - Bridge LAN port 1 to SSID. This attribute must reference one of the following datasources: `system.interface.name` .
* `port2_mode` - LAN port 2 mode. Valid values: `offline` `nat-to-wan` `bridge-to-wan` `bridge-to-ssid` .
* `port2_ssid` - Bridge LAN port 2 to SSID. This attribute must reference one of the following datasources: `system.interface.name` .
* `port3_mode` - LAN port 3 mode. Valid values: `offline` `nat-to-wan` `bridge-to-wan` `bridge-to-ssid` .
* `port3_ssid` - Bridge LAN port 3 to SSID. This attribute must reference one of the following datasources: `system.interface.name` .
* `port4_mode` - LAN port 4 mode. Valid values: `offline` `nat-to-wan` `bridge-to-wan` `bridge-to-ssid` .
* `port4_ssid` - Bridge LAN port 4 to SSID. This attribute must reference one of the following datasources: `system.interface.name` .
* `port5_mode` - LAN port 5 mode. Valid values: `offline` `nat-to-wan` `bridge-to-wan` `bridge-to-ssid` .
* `port5_ssid` - Bridge LAN port 5 to SSID. This attribute must reference one of the following datasources: `system.interface.name` .
* `port6_mode` - LAN port 6 mode. Valid values: `offline` `nat-to-wan` `bridge-to-wan` `bridge-to-ssid` .
* `port6_ssid` - Bridge LAN port 6 to SSID. This attribute must reference one of the following datasources: `system.interface.name` .
* `port7_mode` - LAN port 7 mode. Valid values: `offline` `nat-to-wan` `bridge-to-wan` `bridge-to-ssid` .
* `port7_ssid` - Bridge LAN port 7 to SSID. This attribute must reference one of the following datasources: `system.interface.name` .
* `port8_mode` - LAN port 8 mode. Valid values: `offline` `nat-to-wan` `bridge-to-wan` `bridge-to-ssid` .
* `port8_ssid` - Bridge LAN port 8 to SSID. This attribute must reference one of the following datasources: `system.interface.name` .
* `lbs` - Set various location based service (LBS) options. The structure of `lbs` block is documented below.

The `lbs` block contains:

* `aeroscout` - Enable/disable AeroScout Real Time Location Service (RTLS) support (default = disable). Valid values: `enable` `disable` .
* `aeroscout_ap_mac` - Use BSSID or board MAC address as AP MAC address in AeroScout AP messages (default = bssid). Valid values: `bssid` `board-mac` .
* `aeroscout_mmu_report` - Enable/disable compounded AeroScout tag and MU report (default = enable). Valid values: `enable` `disable` .
* `aeroscout_mu` - Enable/disable AeroScout Mobile Unit (MU) support (default = disable). Valid values: `enable` `disable` .
* `aeroscout_mu_factor` - AeroScout MU mode dilution factor (default = 20).
* `aeroscout_mu_timeout` - AeroScout MU mode timeout (0 - 65535 sec, default = 5).
* `aeroscout_server_ip` - IP address of AeroScout server.
* `aeroscout_server_port` - AeroScout server UDP listening port.
* `ekahau_blink_mode` - Enable/disable Ekahau blink mode (now known as AiRISTA Flow) to track and locate WiFi tags (default = disable). Valid values: `enable` `disable` .
* `ekahau_tag` - WiFi frame MAC address or WiFi Tag.
* `erc_server_ip` - IP address of Ekahau RTLS Controller (ERC).
* `erc_server_port` - Ekahau RTLS Controller (ERC) UDP listening port.
* `fortipresence` - Enable/disable FortiPresence to monitor the location and activity of WiFi clients even if they don't connect to this WiFi network (default = disable). Valid values: `foreign` `both` `disable` .
* `fortipresence_ble` - Enable/disable FortiPresence finding and reporting BLE devices. Valid values: `enable` `disable` .
* `fortipresence_frequency` - FortiPresence report transmit frequency (5 - 65535 sec, default = 30).
* `fortipresence_port` - UDP listening port of FortiPresence server (default = 3000).
* `fortipresence_project` - FortiPresence project name (max. 16 characters, default = fortipresence).
* `fortipresence_rogue` - Enable/disable FortiPresence finding and reporting rogue APs. Valid values: `enable` `disable` .
* `fortipresence_secret` - FortiPresence secret password (max. 16 characters).
* `fortipresence_server` - IP address of FortiPresence server.
* `fortipresence_server_addr_type` - FortiPresence server address type (default = ipv4). Valid values: `ipv4` `fqdn` .
* `fortipresence_server_fqdn` - FQDN of FortiPresence server.
* `fortipresence_unassoc` - Enable/disable FortiPresence finding and reporting unassociated stations. Valid values: `enable` `disable` .
* `station_locate` - Enable/disable client station locating services for all clients, whether associated or not (default = disable). Valid values: `enable` `disable` .
* `led_schedules` - Recurring firewall schedules for illuminating LEDs on the FortiAP. If led-state is enabled, LEDs will be visible when at least one of the schedules is valid. Separate multiple schedule names with a space. The structure of `led_schedules` block is documented below.

The `led_schedules` block contains:

* `name` - Schedule name. This attribute must reference one of the following datasources: `firewall.schedule.group.name` `firewall.schedule.recurring.name` `firewall.schedule.onetime.name` .
* `platform` - WTP, FortiAP, or AP platform. The structure of `platform` block is documented below.

The `platform` block contains:

* `ddscan` - Enable/disable use of one radio for dedicated dual-band scanning to detect RF characterization and wireless threat management. Valid values: `enable` `disable` .
* `mode` - Configure operation mode of 5G radios (default = single-5G). Valid values: `single-5G` `dual-5G` .
* `type` - WTP, FortiAP or AP platform type. There are built-in WTP profiles for all supported FortiAP models. You can select a built-in profile and customize it or create a new profile. Valid values: `AP-11N` `220B` `210B` `222B` `112B` `320B` `11C` `14C` `223B` `28C` `320C` `221C` `25D` `222C` `224D` `214B` `21D` `24D` `112D` `223C` `321C` `C220C` `C225C` `C23JD` `C24JE` `S321C` `S322C` `S323C` `S311C` `S313C` `S321CR` `S322CR` `S323CR` `S421E` `S422E` `S423E` `421E` `423E` `221E` `222E` `223E` `224E` `231E` `S221E` `S223E` `321E` `431F` `432F` `433F` `231F` `234F` `23JF` `831F` `U421E` `U422EV` `U423E` `U221EV` `U223EV` `U24JEV` `U321EV` `U323EV` `U431F` `U433F` `U231F` `U234F` `U432F` .
* `radio_1` - Configuration options for radio 1. The structure of `radio_1` block is documented below.

The `radio_1` block contains:

* `airtime_fairness` - Enable/disable airtime fairness (default = disable). Valid values: `enable` `disable` .
* `amsdu` - Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable). Valid values: `enable` `disable` .
* `ap_handoff` - Enable/disable AP handoff of clients to other APs (default = disable). Valid values: `enable` `disable` .
* `ap_sniffer_addr` - MAC address to monitor.
* `ap_sniffer_bufsize` - Sniffer buffer size (1 - 32 MB, default = 16).
* `ap_sniffer_chan` - Channel on which to operate the sniffer (default = 6).
* `ap_sniffer_ctl` - Enable/disable sniffer on WiFi control frame (default = enable). Valid values: `enable` `disable` .
* `ap_sniffer_data` - Enable/disable sniffer on WiFi data frame (default = enable). Valid values: `enable` `disable` .
* `ap_sniffer_mgmt_beacon` - Enable/disable sniffer on WiFi management Beacon frames (default = enable). Valid values: `enable` `disable` .
* `ap_sniffer_mgmt_other` - Enable/disable sniffer on WiFi management other frames  (default = enable). Valid values: `enable` `disable` .
* `ap_sniffer_mgmt_probe` - Enable/disable sniffer on WiFi management probe frames (default = enable). Valid values: `enable` `disable` .
* `arrp_profile` - Distributed Automatic Radio Resource Provisioning (DARRP) profile name to assign to the radio. This attribute must reference one of the following datasources: `wireless-controller.arrp-profile.name` .
* `auto_power_high` - The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable). Valid values: `enable` `disable` .
* `auto_power_low` - The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_target` - Target of automatic transmit power adjustment in dBm (-95 to -20, default = -70).
* `band` - WiFi band that Radio 1 operates on. Valid values: `802.11a` `802.11b` `802.11g` `802.11n` `802.11n-5G` `802.11ac` `802.11ax-5G` `802.11ax` `802.11ac-2G` `802.11n,g-only` `802.11g-only` `802.11n-only` `802.11n-5G-only` `802.11ac,n-only` `802.11ac-only` `802.11ax,ac-only` `802.11ax,ac,n-only` `802.11ax-5G-only` `802.11ax,n-only` `802.11ax,n,g-only` `802.11ax-only` .
* `band_5g_type` - WiFi 5G band type. Valid values: `5g-full` `5g-high` `5g-low` .
* `bandwidth_admission_control` - Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it. Valid values: `enable` `disable` .
* `bandwidth_capacity` - Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).
* `beacon_interval` - Beacon interval. The time between beacon frames in milliseconds. Actual range of beacon interval depends on the AP platform type (default = 100).
* `bss_color` - BSS color value for this 11ax radio (0 - 63, disable = 0, default = 0).
* `bss_color_mode` - BSS color mode for this 11ax radio (default = auto). Valid values: `auto` `static` .
* `call_admission_control` - Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them. Valid values: `enable` `disable` .
* `call_capacity` - Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).
* `channel_bonding` - Channel bandwidth: 160,80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence. Valid values: `160MHz` `80MHz` `40MHz` `20MHz` .
* `channel_utilization` - Enable/disable measuring channel utilization. Valid values: `enable` `disable` .
* `coexistence` - Enable/disable allowing both HT20 and HT40 on the same radio (default = enable). Valid values: `enable` `disable` .
* `darrp` - Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable). Valid values: `enable` `disable` .
* `drma` - Enable/disable dynamic radio mode assignment (DRMA) (default = disable). Valid values: `disable` `enable` .
* `drma_sensitivity` - Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low). Valid values: `low` `medium` `high` .
* `dtim` - Delivery Traffic Indication Map (DTIM) period (1 - 255, default = 1). Set higher to save battery life of WiFi client in power-save mode.
* `frag_threshold` - Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).
* `frequency_handoff` - Enable/disable frequency handoff of clients to other channels (default = disable). Valid values: `enable` `disable` .
* `iperf_protocol` - Iperf test protocol (default = "UDP"). Valid values: `udp` `tcp` .
* `iperf_server_port` - Iperf service port number.
* `max_clients` - Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.
* `max_distance` - Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).
* `mode` - Mode of radio 1. Radio 1 can be disabled, configured as an access point, a rogue AP monitor, a sniffer, or a station. Valid values: `disabled` `ap` `monitor` `sniffer` `sam` .
* `power_level` - Radio EIRP power level as a percentage of the maximum EIRP power (0 - 100, default = 100).
* `power_mode` - Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities. Valid values: `dBm` `percentage` .
* `power_value` - Radio EIRP power in dBm (1 - 33, default = 27).
* `powersave_optimize` - Enable client power-saving features such as TIM, AC VO, and OBSS etc. Valid values: `tim` `ac-vo` `no-obss-scan` `no-11b-rate` `client-rate-follow` .
* `protection_mode` - Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable). Valid values: `rtscts` `ctsonly` `disable` .
* `rts_threshold` - Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).
* `sam_bssid` - BSSID for WiFi network.
* `sam_captive_portal` - Enable/disable Captive Portal Authentication (default = disable). Valid values: `enable` `disable` .
* `sam_cwp_failure_string` - Failure identification on the page after an incorrect login.
* `sam_cwp_match_string` - Identification string from the captive portal login form.
* `sam_cwp_password` - Password for captive portal authentication.
* `sam_cwp_success_string` - Success identification on the page after a successful login.
* `sam_cwp_test_url` - Website the client is trying to access.
* `sam_cwp_username` - Username for captive portal authentication.
* `sam_password` - Passphrase for WiFi network connection.
* `sam_report_intv` - SAM report interval (sec), 0 for a one-time report.
* `sam_security_type` - Select WiFi network security type (default = "wpa-personal"). Valid values: `open` `wpa-personal` `wpa-enterprise` .
* `sam_server` - SAM test server IP address or domain name.
* `sam_server_fqdn` - SAM test server domain name.
* `sam_server_ip` - SAM test server IP address.
* `sam_server_type` - Select SAM server type (default = "IP"). Valid values: `ip` `fqdn` .
* `sam_ssid` - SSID for WiFi network.
* `sam_test` - Select SAM test type (default = "PING"). Valid values: `ping` `iperf` .
* `sam_username` - Username for WiFi network connection.
* `short_guard_interval` - Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns. Valid values: `enable` `disable` .
* `spectrum_analysis` - Enable/disable spectrum analysis to find interference that would negatively impact wireless performance. Valid values: `enable` `scan-only` `disable` .
* `transmit_optimize` - Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default. Valid values: `disable` `power-save` `aggr-limit` `retry-limit` `send-bar` .
* `vap_all` - Configure method for assigning SSIDs to this FortiAP (default = automatically assign tunnel SSIDs). Valid values: `tunnel` `bridge` `manual` .
* `wids_profile` - Wireless Intrusion Detection System (WIDS) profile name to assign to the radio. This attribute must reference one of the following datasources: `wireless-controller.wids-profile.name` .
* `zero_wait_dfs` - Enable/disable zero wait DFS on radio (default = enable). Valid values: `enable` `disable` .
* `channel` - Selected list of wireless radio channels. The structure of `channel` block is documented below.

The `channel` block contains:

* `chan` - Channel number.
* `vaps` - Manually selected list of Virtual Access Points (VAPs). The structure of `vaps` block is documented below.

The `vaps` block contains:

* `name` - Virtual Access Point (VAP) name. This attribute must reference one of the following datasources: `wireless-controller.vap-group.name` `system.interface.name` .
* `radio_2` - Configuration options for radio 2. The structure of `radio_2` block is documented below.

The `radio_2` block contains:

* `airtime_fairness` - Enable/disable airtime fairness (default = disable). Valid values: `enable` `disable` .
* `amsdu` - Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable). Valid values: `enable` `disable` .
* `ap_handoff` - Enable/disable AP handoff of clients to other APs (default = disable). Valid values: `enable` `disable` .
* `ap_sniffer_addr` - MAC address to monitor.
* `ap_sniffer_bufsize` - Sniffer buffer size (1 - 32 MB, default = 16).
* `ap_sniffer_chan` - Channel on which to operate the sniffer (default = 6).
* `ap_sniffer_ctl` - Enable/disable sniffer on WiFi control frame (default = enable). Valid values: `enable` `disable` .
* `ap_sniffer_data` - Enable/disable sniffer on WiFi data frame (default = enable). Valid values: `enable` `disable` .
* `ap_sniffer_mgmt_beacon` - Enable/disable sniffer on WiFi management Beacon frames (default = enable). Valid values: `enable` `disable` .
* `ap_sniffer_mgmt_other` - Enable/disable sniffer on WiFi management other frames  (default = enable). Valid values: `enable` `disable` .
* `ap_sniffer_mgmt_probe` - Enable/disable sniffer on WiFi management probe frames (default = enable). Valid values: `enable` `disable` .
* `arrp_profile` - Distributed Automatic Radio Resource Provisioning (DARRP) profile name to assign to the radio. This attribute must reference one of the following datasources: `wireless-controller.arrp-profile.name` .
* `auto_power_high` - The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable). Valid values: `enable` `disable` .
* `auto_power_low` - The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_target` - Target of automatic transmit power adjustment in dBm (-95 to -20, default = -70).
* `band` - WiFi band that Radio 2 operates on. Valid values: `802.11a` `802.11b` `802.11g` `802.11n` `802.11n-5G` `802.11ac` `802.11ax-5G` `802.11ax` `802.11ac-2G` `802.11n,g-only` `802.11g-only` `802.11n-only` `802.11n-5G-only` `802.11ac,n-only` `802.11ac-only` `802.11ax,ac-only` `802.11ax,ac,n-only` `802.11ax-5G-only` `802.11ax,n-only` `802.11ax,n,g-only` `802.11ax-only` .
* `band_5g_type` - WiFi 5G band type. Valid values: `5g-full` `5g-high` `5g-low` .
* `bandwidth_admission_control` - Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it. Valid values: `enable` `disable` .
* `bandwidth_capacity` - Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).
* `beacon_interval` - Beacon interval. The time between beacon frames in milliseconds. Actual range of beacon interval depends on the AP platform type (default = 100).
* `bss_color` - BSS color value for this 11ax radio (0 - 63, disable = 0, default = 0).
* `bss_color_mode` - BSS color mode for this 11ax radio (default = auto). Valid values: `auto` `static` .
* `call_admission_control` - Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them. Valid values: `enable` `disable` .
* `call_capacity` - Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).
* `channel_bonding` - Channel bandwidth: 160,80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence. Valid values: `160MHz` `80MHz` `40MHz` `20MHz` .
* `channel_utilization` - Enable/disable measuring channel utilization. Valid values: `enable` `disable` .
* `coexistence` - Enable/disable allowing both HT20 and HT40 on the same radio (default = enable). Valid values: `enable` `disable` .
* `darrp` - Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable). Valid values: `enable` `disable` .
* `drma` - Enable/disable dynamic radio mode assignment (DRMA) (default = disable). Valid values: `disable` `enable` .
* `drma_sensitivity` - Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low). Valid values: `low` `medium` `high` .
* `dtim` - Delivery Traffic Indication Map (DTIM) period (1 - 255, default = 1). Set higher to save battery life of WiFi client in power-save mode.
* `frag_threshold` - Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).
* `frequency_handoff` - Enable/disable frequency handoff of clients to other channels (default = disable). Valid values: `enable` `disable` .
* `iperf_protocol` - Iperf test protocol (default = "UDP"). Valid values: `udp` `tcp` .
* `iperf_server_port` - Iperf service port number.
* `max_clients` - Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.
* `max_distance` - Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).
* `mode` - Mode of radio 2. Radio 2 can be disabled, configured as an access point, a rogue AP monitor, a sniffer, or a station. Valid values: `disabled` `ap` `monitor` `sniffer` `sam` .
* `power_level` - Radio EIRP power level as a percentage of the maximum EIRP power (0 - 100, default = 100).
* `power_mode` - Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities. Valid values: `dBm` `percentage` .
* `power_value` - Radio EIRP power in dBm (1 - 33, default = 27).
* `powersave_optimize` - Enable client power-saving features such as TIM, AC VO, and OBSS etc. Valid values: `tim` `ac-vo` `no-obss-scan` `no-11b-rate` `client-rate-follow` .
* `protection_mode` - Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable). Valid values: `rtscts` `ctsonly` `disable` .
* `rts_threshold` - Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).
* `sam_bssid` - BSSID for WiFi network.
* `sam_captive_portal` - Enable/disable Captive Portal Authentication (default = disable). Valid values: `enable` `disable` .
* `sam_cwp_failure_string` - Failure identification on the page after an incorrect login.
* `sam_cwp_match_string` - Identification string from the captive portal login form.
* `sam_cwp_password` - Password for captive portal authentication.
* `sam_cwp_success_string` - Success identification on the page after a successful login.
* `sam_cwp_test_url` - Website the client is trying to access.
* `sam_cwp_username` - Username for captive portal authentication.
* `sam_password` - Passphrase for WiFi network connection.
* `sam_report_intv` - SAM report interval (sec), 0 for a one-time report.
* `sam_security_type` - Select WiFi network security type (default = "wpa-personal"). Valid values: `open` `wpa-personal` `wpa-enterprise` .
* `sam_server` - SAM test server IP address or domain name.
* `sam_server_fqdn` - SAM test server domain name.
* `sam_server_ip` - SAM test server IP address.
* `sam_server_type` - Select SAM server type (default = "IP"). Valid values: `ip` `fqdn` .
* `sam_ssid` - SSID for WiFi network.
* `sam_test` - Select SAM test type (default = "PING"). Valid values: `ping` `iperf` .
* `sam_username` - Username for WiFi network connection.
* `short_guard_interval` - Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns. Valid values: `enable` `disable` .
* `spectrum_analysis` - Enable/disable spectrum analysis to find interference that would negatively impact wireless performance. Valid values: `enable` `scan-only` `disable` .
* `transmit_optimize` - Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default. Valid values: `disable` `power-save` `aggr-limit` `retry-limit` `send-bar` .
* `vap_all` - Configure method for assigning SSIDs to this FortiAP (default = automatically assign tunnel SSIDs). Valid values: `tunnel` `bridge` `manual` .
* `wids_profile` - Wireless Intrusion Detection System (WIDS) profile name to assign to the radio. This attribute must reference one of the following datasources: `wireless-controller.wids-profile.name` .
* `zero_wait_dfs` - Enable/disable zero wait DFS on radio (default = enable). Valid values: `enable` `disable` .
* `channel` - Selected list of wireless radio channels. The structure of `channel` block is documented below.

The `channel` block contains:

* `chan` - Channel number.
* `vaps` - Manually selected list of Virtual Access Points (VAPs). The structure of `vaps` block is documented below.

The `vaps` block contains:

* `name` - Virtual Access Point (VAP) name. This attribute must reference one of the following datasources: `wireless-controller.vap-group.name` `system.interface.name` .
* `radio_3` - Configuration options for radio 3. The structure of `radio_3` block is documented below.

The `radio_3` block contains:

* `airtime_fairness` - Enable/disable airtime fairness (default = disable). Valid values: `enable` `disable` .
* `amsdu` - Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable). Valid values: `enable` `disable` .
* `ap_handoff` - Enable/disable AP handoff of clients to other APs (default = disable). Valid values: `enable` `disable` .
* `ap_sniffer_addr` - MAC address to monitor.
* `ap_sniffer_bufsize` - Sniffer buffer size (1 - 32 MB, default = 16).
* `ap_sniffer_chan` - Channel on which to operate the sniffer (default = 6).
* `ap_sniffer_ctl` - Enable/disable sniffer on WiFi control frame (default = enable). Valid values: `enable` `disable` .
* `ap_sniffer_data` - Enable/disable sniffer on WiFi data frame (default = enable). Valid values: `enable` `disable` .
* `ap_sniffer_mgmt_beacon` - Enable/disable sniffer on WiFi management Beacon frames (default = enable). Valid values: `enable` `disable` .
* `ap_sniffer_mgmt_other` - Enable/disable sniffer on WiFi management other frames  (default = enable). Valid values: `enable` `disable` .
* `ap_sniffer_mgmt_probe` - Enable/disable sniffer on WiFi management probe frames (default = enable). Valid values: `enable` `disable` .
* `arrp_profile` - Distributed Automatic Radio Resource Provisioning (DARRP) profile name to assign to the radio. This attribute must reference one of the following datasources: `wireless-controller.arrp-profile.name` .
* `auto_power_high` - The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable). Valid values: `enable` `disable` .
* `auto_power_low` - The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_target` - Target of automatic transmit power adjustment in dBm (-95 to -20, default = -70).
* `band` - WiFi band that Radio 3 operates on. Valid values: `802.11a` `802.11b` `802.11g` `802.11n` `802.11n-5G` `802.11ac` `802.11ax-5G` `802.11ax` `802.11ac-2G` `802.11n,g-only` `802.11g-only` `802.11n-only` `802.11n-5G-only` `802.11ac,n-only` `802.11ac-only` `802.11ax,ac-only` `802.11ax,ac,n-only` `802.11ax-5G-only` `802.11ax,n-only` `802.11ax,n,g-only` `802.11ax-only` .
* `band_5g_type` - WiFi 5G band type. Valid values: `5g-full` `5g-high` `5g-low` .
* `bandwidth_admission_control` - Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it. Valid values: `enable` `disable` .
* `bandwidth_capacity` - Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).
* `beacon_interval` - Beacon interval. The time between beacon frames in milliseconds. Actual range of beacon interval depends on the AP platform type (default = 100).
* `bss_color` - BSS color value for this 11ax radio (0 - 63, disable = 0, default = 0).
* `bss_color_mode` - BSS color mode for this 11ax radio (default = auto). Valid values: `auto` `static` .
* `call_admission_control` - Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them. Valid values: `enable` `disable` .
* `call_capacity` - Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).
* `channel_bonding` - Channel bandwidth: 160,80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence. Valid values: `160MHz` `80MHz` `40MHz` `20MHz` .
* `channel_utilization` - Enable/disable measuring channel utilization. Valid values: `enable` `disable` .
* `coexistence` - Enable/disable allowing both HT20 and HT40 on the same radio (default = enable). Valid values: `enable` `disable` .
* `darrp` - Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable). Valid values: `enable` `disable` .
* `drma` - Enable/disable dynamic radio mode assignment (DRMA) (default = disable). Valid values: `disable` `enable` .
* `drma_sensitivity` - Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low). Valid values: `low` `medium` `high` .
* `dtim` - Delivery Traffic Indication Map (DTIM) period (1 - 255, default = 1). Set higher to save battery life of WiFi client in power-save mode.
* `frag_threshold` - Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).
* `frequency_handoff` - Enable/disable frequency handoff of clients to other channels (default = disable). Valid values: `enable` `disable` .
* `iperf_protocol` - Iperf test protocol (default = "UDP"). Valid values: `udp` `tcp` .
* `iperf_server_port` - Iperf service port number.
* `max_clients` - Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.
* `max_distance` - Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).
* `mode` - Mode of radio 3. Radio 3 can be disabled, configured as an access point, a rogue AP monitor, a sniffer, or a station. Valid values: `disabled` `ap` `monitor` `sniffer` `sam` .
* `power_level` - Radio EIRP power level as a percentage of the maximum EIRP power (0 - 100, default = 100).
* `power_mode` - Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities. Valid values: `dBm` `percentage` .
* `power_value` - Radio EIRP power in dBm (1 - 33, default = 27).
* `powersave_optimize` - Enable client power-saving features such as TIM, AC VO, and OBSS etc. Valid values: `tim` `ac-vo` `no-obss-scan` `no-11b-rate` `client-rate-follow` .
* `protection_mode` - Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable). Valid values: `rtscts` `ctsonly` `disable` .
* `rts_threshold` - Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).
* `sam_bssid` - BSSID for WiFi network.
* `sam_captive_portal` - Enable/disable Captive Portal Authentication (default = disable). Valid values: `enable` `disable` .
* `sam_cwp_failure_string` - Failure identification on the page after an incorrect login.
* `sam_cwp_match_string` - Identification string from the captive portal login form.
* `sam_cwp_password` - Password for captive portal authentication.
* `sam_cwp_success_string` - Success identification on the page after a successful login.
* `sam_cwp_test_url` - Website the client is trying to access.
* `sam_cwp_username` - Username for captive portal authentication.
* `sam_password` - Passphrase for WiFi network connection.
* `sam_report_intv` - SAM report interval (sec), 0 for a one-time report.
* `sam_security_type` - Select WiFi network security type (default = "wpa-personal"). Valid values: `open` `wpa-personal` `wpa-enterprise` .
* `sam_server` - SAM test server IP address or domain name.
* `sam_server_fqdn` - SAM test server domain name.
* `sam_server_ip` - SAM test server IP address.
* `sam_server_type` - Select SAM server type (default = "IP"). Valid values: `ip` `fqdn` .
* `sam_ssid` - SSID for WiFi network.
* `sam_test` - Select SAM test type (default = "PING"). Valid values: `ping` `iperf` .
* `sam_username` - Username for WiFi network connection.
* `short_guard_interval` - Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns. Valid values: `enable` `disable` .
* `spectrum_analysis` - Enable/disable spectrum analysis to find interference that would negatively impact wireless performance. Valid values: `enable` `scan-only` `disable` .
* `transmit_optimize` - Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default. Valid values: `disable` `power-save` `aggr-limit` `retry-limit` `send-bar` .
* `vap_all` - Configure method for assigning SSIDs to this FortiAP (default = automatically assign tunnel SSIDs). Valid values: `tunnel` `bridge` `manual` .
* `wids_profile` - Wireless Intrusion Detection System (WIDS) profile name to assign to the radio. This attribute must reference one of the following datasources: `wireless-controller.wids-profile.name` .
* `zero_wait_dfs` - Enable/disable zero wait DFS on radio (default = enable). Valid values: `enable` `disable` .
* `channel` - Selected list of wireless radio channels. The structure of `channel` block is documented below.

The `channel` block contains:

* `chan` - Channel number.
* `vaps` - Manually selected list of Virtual Access Points (VAPs). The structure of `vaps` block is documented below.

The `vaps` block contains:

* `name` - Virtual Access Point (VAP) name. This attribute must reference one of the following datasources: `wireless-controller.vap-group.name` `system.interface.name` .
* `radio_4` - Configuration options for radio 4. The structure of `radio_4` block is documented below.

The `radio_4` block contains:

* `airtime_fairness` - Enable/disable airtime fairness (default = disable). Valid values: `enable` `disable` .
* `amsdu` - Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable). Valid values: `enable` `disable` .
* `ap_handoff` - Enable/disable AP handoff of clients to other APs (default = disable). Valid values: `enable` `disable` .
* `ap_sniffer_addr` - MAC address to monitor.
* `ap_sniffer_bufsize` - Sniffer buffer size (1 - 32 MB, default = 16).
* `ap_sniffer_chan` - Channel on which to operate the sniffer (default = 6).
* `ap_sniffer_ctl` - Enable/disable sniffer on WiFi control frame (default = enable). Valid values: `enable` `disable` .
* `ap_sniffer_data` - Enable/disable sniffer on WiFi data frame (default = enable). Valid values: `enable` `disable` .
* `ap_sniffer_mgmt_beacon` - Enable/disable sniffer on WiFi management Beacon frames (default = enable). Valid values: `enable` `disable` .
* `ap_sniffer_mgmt_other` - Enable/disable sniffer on WiFi management other frames  (default = enable). Valid values: `enable` `disable` .
* `ap_sniffer_mgmt_probe` - Enable/disable sniffer on WiFi management probe frames (default = enable). Valid values: `enable` `disable` .
* `arrp_profile` - Distributed Automatic Radio Resource Provisioning (DARRP) profile name to assign to the radio. This attribute must reference one of the following datasources: `wireless-controller.arrp-profile.name` .
* `auto_power_high` - The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable). Valid values: `enable` `disable` .
* `auto_power_low` - The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_target` - Target of automatic transmit power adjustment in dBm (-95 to -20, default = -70).
* `band` - WiFi band that Radio 3 operates on. Valid values: `802.11a` `802.11b` `802.11g` `802.11n` `802.11n-5G` `802.11ac` `802.11ax-5G` `802.11ax` `802.11ac-2G` `802.11n,g-only` `802.11g-only` `802.11n-only` `802.11n-5G-only` `802.11ac,n-only` `802.11ac-only` `802.11ax,ac-only` `802.11ax,ac,n-only` `802.11ax-5G-only` `802.11ax,n-only` `802.11ax,n,g-only` `802.11ax-only` .
* `band_5g_type` - WiFi 5G band type. Valid values: `5g-full` `5g-high` `5g-low` .
* `bandwidth_admission_control` - Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it. Valid values: `enable` `disable` .
* `bandwidth_capacity` - Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).
* `beacon_interval` - Beacon interval. The time between beacon frames in milliseconds. Actual range of beacon interval depends on the AP platform type (default = 100).
* `bss_color` - BSS color value for this 11ax radio (0 - 63, disable = 0, default = 0).
* `bss_color_mode` - BSS color mode for this 11ax radio (default = auto). Valid values: `auto` `static` .
* `call_admission_control` - Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them. Valid values: `enable` `disable` .
* `call_capacity` - Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).
* `channel_bonding` - Channel bandwidth: 160,80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence. Valid values: `160MHz` `80MHz` `40MHz` `20MHz` .
* `channel_utilization` - Enable/disable measuring channel utilization. Valid values: `enable` `disable` .
* `coexistence` - Enable/disable allowing both HT20 and HT40 on the same radio (default = enable). Valid values: `enable` `disable` .
* `darrp` - Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable). Valid values: `enable` `disable` .
* `drma` - Enable/disable dynamic radio mode assignment (DRMA) (default = disable). Valid values: `disable` `enable` .
* `drma_sensitivity` - Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low). Valid values: `low` `medium` `high` .
* `dtim` - Delivery Traffic Indication Map (DTIM) period (1 - 255, default = 1). Set higher to save battery life of WiFi client in power-save mode.
* `frag_threshold` - Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).
* `frequency_handoff` - Enable/disable frequency handoff of clients to other channels (default = disable). Valid values: `enable` `disable` .
* `iperf_protocol` - Iperf test protocol (default = "UDP"). Valid values: `udp` `tcp` .
* `iperf_server_port` - Iperf service port number.
* `max_clients` - Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.
* `max_distance` - Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).
* `mode` - Mode of radio 3. Radio 3 can be disabled, configured as an access point, a rogue AP monitor, a sniffer, or a station. Valid values: `disabled` `ap` `monitor` `sniffer` `sam` .
* `power_level` - Radio EIRP power level as a percentage of the maximum EIRP power (0 - 100, default = 100).
* `power_mode` - Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities. Valid values: `dBm` `percentage` .
* `power_value` - Radio EIRP power in dBm (1 - 33, default = 27).
* `powersave_optimize` - Enable client power-saving features such as TIM, AC VO, and OBSS etc. Valid values: `tim` `ac-vo` `no-obss-scan` `no-11b-rate` `client-rate-follow` .
* `protection_mode` - Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable). Valid values: `rtscts` `ctsonly` `disable` .
* `rts_threshold` - Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).
* `sam_bssid` - BSSID for WiFi network.
* `sam_captive_portal` - Enable/disable Captive Portal Authentication (default = disable). Valid values: `enable` `disable` .
* `sam_cwp_failure_string` - Failure identification on the page after an incorrect login.
* `sam_cwp_match_string` - Identification string from the captive portal login form.
* `sam_cwp_password` - Password for captive portal authentication.
* `sam_cwp_success_string` - Success identification on the page after a successful login.
* `sam_cwp_test_url` - Website the client is trying to access.
* `sam_cwp_username` - Username for captive portal authentication.
* `sam_password` - Passphrase for WiFi network connection.
* `sam_report_intv` - SAM report interval (sec), 0 for a one-time report.
* `sam_security_type` - Select WiFi network security type (default = "wpa-personal"). Valid values: `open` `wpa-personal` `wpa-enterprise` .
* `sam_server` - SAM test server IP address or domain name.
* `sam_server_fqdn` - SAM test server domain name.
* `sam_server_ip` - SAM test server IP address.
* `sam_server_type` - Select SAM server type (default = "IP"). Valid values: `ip` `fqdn` .
* `sam_ssid` - SSID for WiFi network.
* `sam_test` - Select SAM test type (default = "PING"). Valid values: `ping` `iperf` .
* `sam_username` - Username for WiFi network connection.
* `short_guard_interval` - Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns. Valid values: `enable` `disable` .
* `spectrum_analysis` - Enable/disable spectrum analysis to find interference that would negatively impact wireless performance. Valid values: `enable` `scan-only` `disable` .
* `transmit_optimize` - Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default. Valid values: `disable` `power-save` `aggr-limit` `retry-limit` `send-bar` .
* `vap_all` - Configure method for assigning SSIDs to this FortiAP (default = automatically assign tunnel SSIDs). Valid values: `tunnel` `bridge` `manual` .
* `wids_profile` - Wireless Intrusion Detection System (WIDS) profile name to assign to the radio. This attribute must reference one of the following datasources: `wireless-controller.wids-profile.name` .
* `zero_wait_dfs` - Enable/disable zero wait DFS on radio (default = enable). Valid values: `enable` `disable` .
* `channel` - Selected list of wireless radio channels. The structure of `channel` block is documented below.

The `channel` block contains:

* `chan` - Channel number.
* `vaps` - Manually selected list of Virtual Access Points (VAPs). The structure of `vaps` block is documented below.

The `vaps` block contains:

* `name` - Virtual Access Point (VAP) name. This attribute must reference one of the following datasources: `wireless-controller.vap-group.name` `system.interface.name` .
* `split_tunneling_acl` - Split tunneling ACL filter list. The structure of `split_tunneling_acl` block is documented below.

The `split_tunneling_acl` block contains:

* `dest_ip` - Destination IP and mask for the split-tunneling subnet.
* `id` - ID.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontroller_wtpprofile can be imported using:
```sh
terraform import fortios_wirelesscontroller_wtpprofile.labelname {{mkey}}
```
