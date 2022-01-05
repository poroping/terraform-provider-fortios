// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure WTP profiles or FortiAP profiles that define radio settings for manageable FortiAP platforms.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceWirelessControllerWtpProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure WTP profiles or FortiAP profiles that define radio settings for manageable FortiAP platforms.",

		ReadContext: dataSourceWirelessControllerWtpProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"allowaccess": {
				Type:        schema.TypeString,
				Description: "Control management access to the managed WTP, FortiAP, or AP. Separate entries with a space.",
				Computed:    true,
			},
			"ap_country": {
				Type:        schema.TypeString,
				Description: "Country in which this WTP, FortiAP or AP will operate (default = NA, automatically use the country configured for the current VDOM).",
				Computed:    true,
			},
			"ap_handoff": {
				Type:        schema.TypeString,
				Description: "Enable/disable AP handoff of clients to other APs (default = disable).",
				Computed:    true,
			},
			"apcfg_profile": {
				Type:        schema.TypeString,
				Description: "AP local configuration profile name.",
				Computed:    true,
			},
			"ble_profile": {
				Type:        schema.TypeString,
				Description: "Bluetooth Low Energy profile name.",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"console_login": {
				Type:        schema.TypeString,
				Description: "Enable/disable FAP console login access (default = enable).",
				Computed:    true,
			},
			"control_message_offload": {
				Type:        schema.TypeString,
				Description: "Enable/disable CAPWAP control message data channel offload.",
				Computed:    true,
			},
			"deny_mac_list": {
				Type:        schema.TypeList,
				Description: "List of MAC addresses that are denied access to this WTP, FortiAP, or AP.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"mac": {
							Type:        schema.TypeString,
							Description: "A WiFi device with this MAC address is denied access to this WTP, FortiAP or AP.",
							Computed:    true,
						},
					},
				},
			},
			"dtls_in_kernel": {
				Type:        schema.TypeString,
				Description: "Enable/disable data channel DTLS in kernel.",
				Computed:    true,
			},
			"dtls_policy": {
				Type:        schema.TypeString,
				Description: "WTP data channel DTLS policy (default = clear-text).",
				Computed:    true,
			},
			"energy_efficient_ethernet": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of energy efficient Ethernet on WTP.",
				Computed:    true,
			},
			"esl_ses_dongle": {
				Type:        schema.TypeList,
				Description: "ESL SES-imagotag dongle configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"apc_addr_type": {
							Type:        schema.TypeString,
							Description: "ESL SES-imagotag APC address type (default = fqdn).",
							Computed:    true,
						},
						"apc_fqdn": {
							Type:        schema.TypeString,
							Description: "FQDN of ESL SES-imagotag Access Point Controller (APC).",
							Computed:    true,
						},
						"apc_ip": {
							Type:        schema.TypeString,
							Description: "IP address of ESL SES-imagotag Access Point Controller (APC).",
							Computed:    true,
						},
						"apc_port": {
							Type:        schema.TypeInt,
							Description: "Port of ESL SES-imagotag Access Point Controller (APC).",
							Computed:    true,
						},
						"coex_level": {
							Type:        schema.TypeString,
							Description: "ESL SES-imagotag dongle coexistence level (default = none).",
							Computed:    true,
						},
						"compliance_level": {
							Type:        schema.TypeString,
							Description: "Compliance levels for the ESL solution integration (default = compliance-level-2).",
							Computed:    true,
						},
						"esl_channel": {
							Type:        schema.TypeString,
							Description: "ESL SES-imagotag dongle channel (default = 127).",
							Computed:    true,
						},
						"output_power": {
							Type:        schema.TypeString,
							Description: "ESL SES-imagotag dongle output power (default = A).",
							Computed:    true,
						},
						"scd_enable": {
							Type:        schema.TypeString,
							Description: "Enable/disable ESL SES-imagotag Serial Communication Daemon (SCD) (default = disable).",
							Computed:    true,
						},
						"tls_cert_verification": {
							Type:        schema.TypeString,
							Description: "Enable/disable TLS Certificate verification. (default = enable).",
							Computed:    true,
						},
						"tls_fqdn_verification": {
							Type:        schema.TypeString,
							Description: "Enable/disable TLS Certificate verification. (default = disable).",
							Computed:    true,
						},
					},
				},
			},
			"ext_info_enable": {
				Type:        schema.TypeString,
				Description: "Enable/disable station/VAP/radio extension information.",
				Computed:    true,
			},
			"frequency_handoff": {
				Type:        schema.TypeString,
				Description: "Enable/disable frequency handoff of clients to other channels (default = disable).",
				Computed:    true,
			},
			"handoff_roaming": {
				Type:        schema.TypeString,
				Description: "Enable/disable client load balancing during roaming to avoid roaming delay (default = disable).",
				Computed:    true,
			},
			"handoff_rssi": {
				Type:        schema.TypeInt,
				Description: "Minimum received signal strength indicator (RSSI) value for handoff (20 - 30, default = 25).",
				Computed:    true,
			},
			"handoff_sta_thresh": {
				Type:        schema.TypeInt,
				Description: "Threshold value for AP handoff.",
				Computed:    true,
			},
			"indoor_outdoor_deployment": {
				Type:        schema.TypeString,
				Description: "Set to allow indoor/outdoor-only channels under regulatory rules (default = platform-determined).",
				Computed:    true,
			},
			"ip_fragment_preventing": {
				Type:        schema.TypeString,
				Description: "Method(s) by which IP fragmentation is prevented for control and data packets through CAPWAP tunnel (default = tcp-mss-adjust).",
				Computed:    true,
			},
			"lan": {
				Type:        schema.TypeList,
				Description: "WTP LAN port mapping.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_esl_mode": {
							Type:        schema.TypeString,
							Description: "ESL port mode.",
							Computed:    true,
						},
						"port_esl_ssid": {
							Type:        schema.TypeString,
							Description: "Bridge ESL port to SSID.",
							Computed:    true,
						},
						"port_mode": {
							Type:        schema.TypeString,
							Description: "LAN port mode.",
							Computed:    true,
						},
						"port_ssid": {
							Type:        schema.TypeString,
							Description: "Bridge LAN port to SSID.",
							Computed:    true,
						},
						"port1_mode": {
							Type:        schema.TypeString,
							Description: "LAN port 1 mode.",
							Computed:    true,
						},
						"port1_ssid": {
							Type:        schema.TypeString,
							Description: "Bridge LAN port 1 to SSID.",
							Computed:    true,
						},
						"port2_mode": {
							Type:        schema.TypeString,
							Description: "LAN port 2 mode.",
							Computed:    true,
						},
						"port2_ssid": {
							Type:        schema.TypeString,
							Description: "Bridge LAN port 2 to SSID.",
							Computed:    true,
						},
						"port3_mode": {
							Type:        schema.TypeString,
							Description: "LAN port 3 mode.",
							Computed:    true,
						},
						"port3_ssid": {
							Type:        schema.TypeString,
							Description: "Bridge LAN port 3 to SSID.",
							Computed:    true,
						},
						"port4_mode": {
							Type:        schema.TypeString,
							Description: "LAN port 4 mode.",
							Computed:    true,
						},
						"port4_ssid": {
							Type:        schema.TypeString,
							Description: "Bridge LAN port 4 to SSID.",
							Computed:    true,
						},
						"port5_mode": {
							Type:        schema.TypeString,
							Description: "LAN port 5 mode.",
							Computed:    true,
						},
						"port5_ssid": {
							Type:        schema.TypeString,
							Description: "Bridge LAN port 5 to SSID.",
							Computed:    true,
						},
						"port6_mode": {
							Type:        schema.TypeString,
							Description: "LAN port 6 mode.",
							Computed:    true,
						},
						"port6_ssid": {
							Type:        schema.TypeString,
							Description: "Bridge LAN port 6 to SSID.",
							Computed:    true,
						},
						"port7_mode": {
							Type:        schema.TypeString,
							Description: "LAN port 7 mode.",
							Computed:    true,
						},
						"port7_ssid": {
							Type:        schema.TypeString,
							Description: "Bridge LAN port 7 to SSID.",
							Computed:    true,
						},
						"port8_mode": {
							Type:        schema.TypeString,
							Description: "LAN port 8 mode.",
							Computed:    true,
						},
						"port8_ssid": {
							Type:        schema.TypeString,
							Description: "Bridge LAN port 8 to SSID.",
							Computed:    true,
						},
					},
				},
			},
			"lbs": {
				Type:        schema.TypeList,
				Description: "Set various location based service (LBS) options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aeroscout": {
							Type:        schema.TypeString,
							Description: "Enable/disable AeroScout Real Time Location Service (RTLS) support (default = disable).",
							Computed:    true,
						},
						"aeroscout_ap_mac": {
							Type:        schema.TypeString,
							Description: "Use BSSID or board MAC address as AP MAC address in AeroScout AP messages (default = bssid).",
							Computed:    true,
						},
						"aeroscout_mmu_report": {
							Type:        schema.TypeString,
							Description: "Enable/disable compounded AeroScout tag and MU report (default = enable).",
							Computed:    true,
						},
						"aeroscout_mu": {
							Type:        schema.TypeString,
							Description: "Enable/disable AeroScout Mobile Unit (MU) support (default = disable).",
							Computed:    true,
						},
						"aeroscout_mu_factor": {
							Type:        schema.TypeInt,
							Description: "AeroScout MU mode dilution factor (default = 20).",
							Computed:    true,
						},
						"aeroscout_mu_timeout": {
							Type:        schema.TypeInt,
							Description: "AeroScout MU mode timeout (0 - 65535 sec, default = 5).",
							Computed:    true,
						},
						"aeroscout_server_ip": {
							Type:        schema.TypeString,
							Description: "IP address of AeroScout server.",
							Computed:    true,
						},
						"aeroscout_server_port": {
							Type:        schema.TypeInt,
							Description: "AeroScout server UDP listening port.",
							Computed:    true,
						},
						"ekahau_blink_mode": {
							Type:        schema.TypeString,
							Description: "Enable/disable Ekahau blink mode (now known as AiRISTA Flow) to track and locate WiFi tags (default = disable).",
							Computed:    true,
						},
						"ekahau_tag": {
							Type:        schema.TypeString,
							Description: "WiFi frame MAC address or WiFi Tag.",
							Computed:    true,
						},
						"erc_server_ip": {
							Type:        schema.TypeString,
							Description: "IP address of Ekahau RTLS Controller (ERC).",
							Computed:    true,
						},
						"erc_server_port": {
							Type:        schema.TypeInt,
							Description: "Ekahau RTLS Controller (ERC) UDP listening port.",
							Computed:    true,
						},
						"fortipresence": {
							Type:        schema.TypeString,
							Description: "Enable/disable FortiPresence to monitor the location and activity of WiFi clients even if they don't connect to this WiFi network (default = disable).",
							Computed:    true,
						},
						"fortipresence_ble": {
							Type:        schema.TypeString,
							Description: "Enable/disable FortiPresence finding and reporting BLE devices.",
							Computed:    true,
						},
						"fortipresence_frequency": {
							Type:        schema.TypeInt,
							Description: "FortiPresence report transmit frequency (5 - 65535 sec, default = 30).",
							Computed:    true,
						},
						"fortipresence_port": {
							Type:        schema.TypeInt,
							Description: "UDP listening port of FortiPresence server (default = 3000).",
							Computed:    true,
						},
						"fortipresence_project": {
							Type:        schema.TypeString,
							Description: "FortiPresence project name (max. 16 characters, default = fortipresence).",
							Computed:    true,
						},
						"fortipresence_rogue": {
							Type:        schema.TypeString,
							Description: "Enable/disable FortiPresence finding and reporting rogue APs.",
							Computed:    true,
						},
						"fortipresence_secret": {
							Type:        schema.TypeString,
							Description: "FortiPresence secret password (max. 16 characters).",
							Computed:    true,
							Sensitive:   true,
						},
						"fortipresence_server": {
							Type:        schema.TypeString,
							Description: "IP address of FortiPresence server.",
							Computed:    true,
						},
						"fortipresence_server_addr_type": {
							Type:        schema.TypeString,
							Description: "FortiPresence server address type (default = ipv4).",
							Computed:    true,
						},
						"fortipresence_server_fqdn": {
							Type:        schema.TypeString,
							Description: "FQDN of FortiPresence server.",
							Computed:    true,
						},
						"fortipresence_unassoc": {
							Type:        schema.TypeString,
							Description: "Enable/disable FortiPresence finding and reporting unassociated stations.",
							Computed:    true,
						},
						"station_locate": {
							Type:        schema.TypeString,
							Description: "Enable/disable client station locating services for all clients, whether associated or not (default = disable).",
							Computed:    true,
						},
					},
				},
			},
			"led_schedules": {
				Type:        schema.TypeList,
				Description: "Recurring firewall schedules for illuminating LEDs on the FortiAP. If led-state is enabled, LEDs will be visible when at least one of the schedules is valid. Separate multiple schedule names with a space.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Schedule name.",
							Computed:    true,
						},
					},
				},
			},
			"led_state": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of LEDs on WTP (default = disable).",
				Computed:    true,
			},
			"lldp": {
				Type:        schema.TypeString,
				Description: "Enable/disable Link Layer Discovery Protocol (LLDP) for the WTP, FortiAP, or AP (default = enable).",
				Computed:    true,
			},
			"login_passwd": {
				Type:        schema.TypeString,
				Description: "Set the managed WTP, FortiAP, or AP's administrator password.",
				Computed:    true,
				Sensitive:   true,
			},
			"login_passwd_change": {
				Type:        schema.TypeString,
				Description: "Change or reset the administrator password of a managed WTP, FortiAP or AP (yes, default, or no, default = no).",
				Computed:    true,
			},
			"max_clients": {
				Type:        schema.TypeInt,
				Description: "Maximum number of stations (STAs) supported by the WTP (default = 0, meaning no client limitation).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "WTP (or FortiAP or AP) profile name.",
				Required:    true,
			},
			"platform": {
				Type:        schema.TypeList,
				Description: "WTP, FortiAP, or AP platform.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ddscan": {
							Type:        schema.TypeString,
							Description: "Enable/disable use of one radio for dedicated dual-band scanning to detect RF characterization and wireless threat management.",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Configure operation mode of 5G radios (default = single-5G).",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "WTP, FortiAP or AP platform type. There are built-in WTP profiles for all supported FortiAP models. You can select a built-in profile and customize it or create a new profile.",
							Computed:    true,
						},
					},
				},
			},
			"poe_mode": {
				Type:        schema.TypeString,
				Description: "Set the WTP, FortiAP, or AP's PoE mode.",
				Computed:    true,
			},
			"radio_1": {
				Type:        schema.TypeList,
				Description: "Configuration options for radio 1.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"airtime_fairness": {
							Type:        schema.TypeString,
							Description: "Enable/disable airtime fairness (default = disable).",
							Computed:    true,
						},
						"amsdu": {
							Type:        schema.TypeString,
							Description: "Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable).",
							Computed:    true,
						},
						"ap_handoff": {
							Type:        schema.TypeString,
							Description: "Enable/disable AP handoff of clients to other APs (default = disable).",
							Computed:    true,
						},
						"ap_sniffer_addr": {
							Type:        schema.TypeString,
							Description: "MAC address to monitor.",
							Computed:    true,
						},
						"ap_sniffer_bufsize": {
							Type:        schema.TypeInt,
							Description: "Sniffer buffer size (1 - 32 MB, default = 16).",
							Computed:    true,
						},
						"ap_sniffer_chan": {
							Type:        schema.TypeInt,
							Description: "Channel on which to operate the sniffer (default = 6).",
							Computed:    true,
						},
						"ap_sniffer_ctl": {
							Type:        schema.TypeString,
							Description: "Enable/disable sniffer on WiFi control frame (default = enable).",
							Computed:    true,
						},
						"ap_sniffer_data": {
							Type:        schema.TypeString,
							Description: "Enable/disable sniffer on WiFi data frame (default = enable).",
							Computed:    true,
						},
						"ap_sniffer_mgmt_beacon": {
							Type:        schema.TypeString,
							Description: "Enable/disable sniffer on WiFi management Beacon frames (default = enable).",
							Computed:    true,
						},
						"ap_sniffer_mgmt_other": {
							Type:        schema.TypeString,
							Description: "Enable/disable sniffer on WiFi management other frames  (default = enable).",
							Computed:    true,
						},
						"ap_sniffer_mgmt_probe": {
							Type:        schema.TypeString,
							Description: "Enable/disable sniffer on WiFi management probe frames (default = enable).",
							Computed:    true,
						},
						"auto_power_high": {
							Type:        schema.TypeInt,
							Description: "The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).",
							Computed:    true,
						},
						"auto_power_level": {
							Type:        schema.TypeString,
							Description: "Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable).",
							Computed:    true,
						},
						"auto_power_low": {
							Type:        schema.TypeInt,
							Description: "The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).",
							Computed:    true,
						},
						"auto_power_target": {
							Type:        schema.TypeString,
							Description: "The target of automatic transmit power adjustment in dBm. (-95 to -20, default = -70).",
							Computed:    true,
						},
						"band": {
							Type:        schema.TypeString,
							Description: "WiFi band that Radio 1 operates on.",
							Computed:    true,
						},
						"band_5g_type": {
							Type:        schema.TypeString,
							Description: "WiFi 5G band type.",
							Computed:    true,
						},
						"bandwidth_admission_control": {
							Type:        schema.TypeString,
							Description: "Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it.",
							Computed:    true,
						},
						"bandwidth_capacity": {
							Type:        schema.TypeInt,
							Description: "Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).",
							Computed:    true,
						},
						"beacon_interval": {
							Type:        schema.TypeInt,
							Description: "Beacon interval. The time between beacon frames in msec (the actual range of beacon interval depends on the AP platform type, default = 100).",
							Computed:    true,
						},
						"bss_color": {
							Type:        schema.TypeInt,
							Description: "BSS color value for this 11ax radio (0 - 63, 0 means disable. default = 0).",
							Computed:    true,
						},
						"bss_color_mode": {
							Type:        schema.TypeString,
							Description: "BSS color mode for this 11ax radio (default = auto).",
							Computed:    true,
						},
						"call_admission_control": {
							Type:        schema.TypeString,
							Description: "Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them.",
							Computed:    true,
						},
						"call_capacity": {
							Type:        schema.TypeInt,
							Description: "Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).",
							Computed:    true,
						},
						"channel": {
							Type:        schema.TypeList,
							Description: "Selected list of wireless radio channels.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"chan": {
										Type:        schema.TypeString,
										Description: "Channel number.",
										Computed:    true,
									},
								},
							},
						},
						"channel_bonding": {
							Type:        schema.TypeString,
							Description: "Channel bandwidth: 160,80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence.",
							Computed:    true,
						},
						"channel_utilization": {
							Type:        schema.TypeString,
							Description: "Enable/disable measuring channel utilization.",
							Computed:    true,
						},
						"coexistence": {
							Type:        schema.TypeString,
							Description: "Enable/disable allowing both HT20 and HT40 on the same radio (default = enable).",
							Computed:    true,
						},
						"darrp": {
							Type:        schema.TypeString,
							Description: "Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable).",
							Computed:    true,
						},
						"drma": {
							Type:        schema.TypeString,
							Description: "Enable/disable dynamic radio mode assignment (DRMA) (default = disable).",
							Computed:    true,
						},
						"drma_sensitivity": {
							Type:        schema.TypeString,
							Description: "Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low).",
							Computed:    true,
						},
						"dtim": {
							Type:        schema.TypeInt,
							Description: "Delivery Traffic Indication Map (DTIM) period (1 - 255, default = 1). Set higher to save battery life of WiFi client in power-save mode.",
							Computed:    true,
						},
						"frag_threshold": {
							Type:        schema.TypeInt,
							Description: "Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).",
							Computed:    true,
						},
						"frequency_handoff": {
							Type:        schema.TypeString,
							Description: "Enable/disable frequency handoff of clients to other channels (default = disable).",
							Computed:    true,
						},
						"iperf_protocol": {
							Type:        schema.TypeString,
							Description: "Iperf test protocol (default = \"UDP\").",
							Computed:    true,
						},
						"iperf_server_port": {
							Type:        schema.TypeInt,
							Description: "Iperf service port number.",
							Computed:    true,
						},
						"max_clients": {
							Type:        schema.TypeInt,
							Description: "Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.",
							Computed:    true,
						},
						"max_distance": {
							Type:        schema.TypeInt,
							Description: "Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Mode of radio 1. Radio 1 can be disabled, configured as an access point, a rogue AP monitor, a sniffer, or a station.",
							Computed:    true,
						},
						"power_level": {
							Type:        schema.TypeInt,
							Description: "Radio EIRP power level as a percentage of the maximum EIRP power (0 - 100, default = 100).",
							Computed:    true,
						},
						"power_mode": {
							Type:        schema.TypeString,
							Description: "Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities.",
							Computed:    true,
						},
						"power_value": {
							Type:        schema.TypeInt,
							Description: "Radio EIRP power in dBm (1 - 33, default = 27).",
							Computed:    true,
						},
						"powersave_optimize": {
							Type:        schema.TypeString,
							Description: "Enable client power-saving features such as TIM, AC VO, and OBSS etc.",
							Computed:    true,
						},
						"protection_mode": {
							Type:        schema.TypeString,
							Description: "Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable).",
							Computed:    true,
						},
						"rts_threshold": {
							Type:        schema.TypeInt,
							Description: "Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).",
							Computed:    true,
						},
						"sam_bssid": {
							Type:        schema.TypeString,
							Description: "BSSID for WiFi network.",
							Computed:    true,
						},
						"sam_captive_portal": {
							Type:        schema.TypeString,
							Description: "Enable/disable Captive Portal Authentication (default = disable).",
							Computed:    true,
						},
						"sam_cwp_failure_string": {
							Type:        schema.TypeString,
							Description: "Failure identification on the page after an incorrect login.",
							Computed:    true,
						},
						"sam_cwp_match_string": {
							Type:        schema.TypeString,
							Description: "Identification string from the captive portal login form.",
							Computed:    true,
						},
						"sam_cwp_password": {
							Type:        schema.TypeString,
							Description: "Password for captive portal authentication.",
							Computed:    true,
							Sensitive:   true,
						},
						"sam_cwp_success_string": {
							Type:        schema.TypeString,
							Description: "Success identification on the page after a successful login.",
							Computed:    true,
						},
						"sam_cwp_test_url": {
							Type:        schema.TypeString,
							Description: "Website the client is trying to access.",
							Computed:    true,
						},
						"sam_cwp_username": {
							Type:        schema.TypeString,
							Description: "Username for captive portal authentication.",
							Computed:    true,
						},
						"sam_password": {
							Type:        schema.TypeString,
							Description: "Passphrase for WiFi network connection.",
							Computed:    true,
							Sensitive:   true,
						},
						"sam_report_intv": {
							Type:        schema.TypeInt,
							Description: "SAM report interval (sec), 0 for a one-time report.",
							Computed:    true,
						},
						"sam_security_type": {
							Type:        schema.TypeString,
							Description: "Select WiFi network security type (default = \"wpa-personal\").",
							Computed:    true,
						},
						"sam_server": {
							Type:        schema.TypeString,
							Description: "SAM test server IP address or domain name.",
							Computed:    true,
						},
						"sam_server_fqdn": {
							Type:        schema.TypeString,
							Description: "SAM test server domain name.",
							Computed:    true,
						},
						"sam_server_ip": {
							Type:        schema.TypeString,
							Description: "SAM test server IP address.",
							Computed:    true,
						},
						"sam_server_type": {
							Type:        schema.TypeString,
							Description: "Select SAM server type (default = \"IP\").",
							Computed:    true,
						},
						"sam_ssid": {
							Type:        schema.TypeString,
							Description: "SSID for WiFi network.",
							Computed:    true,
						},
						"sam_test": {
							Type:        schema.TypeString,
							Description: "Select SAM test type (default = \"PING\").",
							Computed:    true,
						},
						"sam_username": {
							Type:        schema.TypeString,
							Description: "Username for WiFi network connection.",
							Computed:    true,
						},
						"short_guard_interval": {
							Type:        schema.TypeString,
							Description: "Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns.",
							Computed:    true,
						},
						"spectrum_analysis": {
							Type:        schema.TypeString,
							Description: "Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.",
							Computed:    true,
						},
						"transmit_optimize": {
							Type:        schema.TypeString,
							Description: "Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default.",
							Computed:    true,
						},
						"vap_all": {
							Type:        schema.TypeString,
							Description: "Configure method for assigning SSIDs to this FortiAP (default = automatically assign tunnel SSIDs).",
							Computed:    true,
						},
						"vaps": {
							Type:        schema.TypeList,
							Description: "Manually selected list of Virtual Access Points (VAPs).",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Virtual Access Point (VAP) name.",
										Computed:    true,
									},
								},
							},
						},
						"wids_profile": {
							Type:        schema.TypeString,
							Description: "Wireless Intrusion Detection System (WIDS) profile name to assign to the radio.",
							Computed:    true,
						},
						"zero_wait_dfs": {
							Type:        schema.TypeString,
							Description: "Enable/disable zero wait DFS on radio (default = enable).",
							Computed:    true,
						},
					},
				},
			},
			"radio_2": {
				Type:        schema.TypeList,
				Description: "Configuration options for radio 2.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"airtime_fairness": {
							Type:        schema.TypeString,
							Description: "Enable/disable airtime fairness (default = disable).",
							Computed:    true,
						},
						"amsdu": {
							Type:        schema.TypeString,
							Description: "Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable).",
							Computed:    true,
						},
						"ap_handoff": {
							Type:        schema.TypeString,
							Description: "Enable/disable AP handoff of clients to other APs (default = disable).",
							Computed:    true,
						},
						"ap_sniffer_addr": {
							Type:        schema.TypeString,
							Description: "MAC address to monitor.",
							Computed:    true,
						},
						"ap_sniffer_bufsize": {
							Type:        schema.TypeInt,
							Description: "Sniffer buffer size (1 - 32 MB, default = 16).",
							Computed:    true,
						},
						"ap_sniffer_chan": {
							Type:        schema.TypeInt,
							Description: "Channel on which to operate the sniffer (default = 6).",
							Computed:    true,
						},
						"ap_sniffer_ctl": {
							Type:        schema.TypeString,
							Description: "Enable/disable sniffer on WiFi control frame (default = enable).",
							Computed:    true,
						},
						"ap_sniffer_data": {
							Type:        schema.TypeString,
							Description: "Enable/disable sniffer on WiFi data frame (default = enable).",
							Computed:    true,
						},
						"ap_sniffer_mgmt_beacon": {
							Type:        schema.TypeString,
							Description: "Enable/disable sniffer on WiFi management Beacon frames (default = enable).",
							Computed:    true,
						},
						"ap_sniffer_mgmt_other": {
							Type:        schema.TypeString,
							Description: "Enable/disable sniffer on WiFi management other frames  (default = enable).",
							Computed:    true,
						},
						"ap_sniffer_mgmt_probe": {
							Type:        schema.TypeString,
							Description: "Enable/disable sniffer on WiFi management probe frames (default = enable).",
							Computed:    true,
						},
						"auto_power_high": {
							Type:        schema.TypeInt,
							Description: "The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).",
							Computed:    true,
						},
						"auto_power_level": {
							Type:        schema.TypeString,
							Description: "Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable).",
							Computed:    true,
						},
						"auto_power_low": {
							Type:        schema.TypeInt,
							Description: "The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).",
							Computed:    true,
						},
						"auto_power_target": {
							Type:        schema.TypeString,
							Description: "The target of automatic transmit power adjustment in dBm. (-95 to -20, default = -70).",
							Computed:    true,
						},
						"band": {
							Type:        schema.TypeString,
							Description: "WiFi band that Radio 2 operates on.",
							Computed:    true,
						},
						"band_5g_type": {
							Type:        schema.TypeString,
							Description: "WiFi 5G band type.",
							Computed:    true,
						},
						"bandwidth_admission_control": {
							Type:        schema.TypeString,
							Description: "Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it.",
							Computed:    true,
						},
						"bandwidth_capacity": {
							Type:        schema.TypeInt,
							Description: "Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).",
							Computed:    true,
						},
						"beacon_interval": {
							Type:        schema.TypeInt,
							Description: "Beacon interval. The time between beacon frames in msec (the actual range of beacon interval depends on the AP platform type, default = 100).",
							Computed:    true,
						},
						"bss_color": {
							Type:        schema.TypeInt,
							Description: "BSS color value for this 11ax radio (0 - 63, 0 means disable. default = 0).",
							Computed:    true,
						},
						"bss_color_mode": {
							Type:        schema.TypeString,
							Description: "BSS color mode for this 11ax radio (default = auto).",
							Computed:    true,
						},
						"call_admission_control": {
							Type:        schema.TypeString,
							Description: "Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them.",
							Computed:    true,
						},
						"call_capacity": {
							Type:        schema.TypeInt,
							Description: "Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).",
							Computed:    true,
						},
						"channel": {
							Type:        schema.TypeList,
							Description: "Selected list of wireless radio channels.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"chan": {
										Type:        schema.TypeString,
										Description: "Channel number.",
										Computed:    true,
									},
								},
							},
						},
						"channel_bonding": {
							Type:        schema.TypeString,
							Description: "Channel bandwidth: 160,80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence.",
							Computed:    true,
						},
						"channel_utilization": {
							Type:        schema.TypeString,
							Description: "Enable/disable measuring channel utilization.",
							Computed:    true,
						},
						"coexistence": {
							Type:        schema.TypeString,
							Description: "Enable/disable allowing both HT20 and HT40 on the same radio (default = enable).",
							Computed:    true,
						},
						"darrp": {
							Type:        schema.TypeString,
							Description: "Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable).",
							Computed:    true,
						},
						"drma": {
							Type:        schema.TypeString,
							Description: "Enable/disable dynamic radio mode assignment (DRMA) (default = disable).",
							Computed:    true,
						},
						"drma_sensitivity": {
							Type:        schema.TypeString,
							Description: "Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low).",
							Computed:    true,
						},
						"dtim": {
							Type:        schema.TypeInt,
							Description: "Delivery Traffic Indication Map (DTIM) period (1 - 255, default = 1). Set higher to save battery life of WiFi client in power-save mode.",
							Computed:    true,
						},
						"frag_threshold": {
							Type:        schema.TypeInt,
							Description: "Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).",
							Computed:    true,
						},
						"frequency_handoff": {
							Type:        schema.TypeString,
							Description: "Enable/disable frequency handoff of clients to other channels (default = disable).",
							Computed:    true,
						},
						"iperf_protocol": {
							Type:        schema.TypeString,
							Description: "Iperf test protocol (default = \"UDP\").",
							Computed:    true,
						},
						"iperf_server_port": {
							Type:        schema.TypeInt,
							Description: "Iperf service port number.",
							Computed:    true,
						},
						"max_clients": {
							Type:        schema.TypeInt,
							Description: "Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.",
							Computed:    true,
						},
						"max_distance": {
							Type:        schema.TypeInt,
							Description: "Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Mode of radio 2. Radio 2 can be disabled, configured as an access point, a rogue AP monitor, a sniffer, or a station.",
							Computed:    true,
						},
						"power_level": {
							Type:        schema.TypeInt,
							Description: "Radio EIRP power level as a percentage of the maximum EIRP power (0 - 100, default = 100).",
							Computed:    true,
						},
						"power_mode": {
							Type:        schema.TypeString,
							Description: "Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities.",
							Computed:    true,
						},
						"power_value": {
							Type:        schema.TypeInt,
							Description: "Radio EIRP power in dBm (1 - 33, default = 27).",
							Computed:    true,
						},
						"powersave_optimize": {
							Type:        schema.TypeString,
							Description: "Enable client power-saving features such as TIM, AC VO, and OBSS etc.",
							Computed:    true,
						},
						"protection_mode": {
							Type:        schema.TypeString,
							Description: "Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable).",
							Computed:    true,
						},
						"rts_threshold": {
							Type:        schema.TypeInt,
							Description: "Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).",
							Computed:    true,
						},
						"sam_bssid": {
							Type:        schema.TypeString,
							Description: "BSSID for WiFi network.",
							Computed:    true,
						},
						"sam_captive_portal": {
							Type:        schema.TypeString,
							Description: "Enable/disable Captive Portal Authentication (default = disable).",
							Computed:    true,
						},
						"sam_cwp_failure_string": {
							Type:        schema.TypeString,
							Description: "Failure identification on the page after an incorrect login.",
							Computed:    true,
						},
						"sam_cwp_match_string": {
							Type:        schema.TypeString,
							Description: "Identification string from the captive portal login form.",
							Computed:    true,
						},
						"sam_cwp_password": {
							Type:        schema.TypeString,
							Description: "Password for captive portal authentication.",
							Computed:    true,
							Sensitive:   true,
						},
						"sam_cwp_success_string": {
							Type:        schema.TypeString,
							Description: "Success identification on the page after a successful login.",
							Computed:    true,
						},
						"sam_cwp_test_url": {
							Type:        schema.TypeString,
							Description: "Website the client is trying to access.",
							Computed:    true,
						},
						"sam_cwp_username": {
							Type:        schema.TypeString,
							Description: "Username for captive portal authentication.",
							Computed:    true,
						},
						"sam_password": {
							Type:        schema.TypeString,
							Description: "Passphrase for WiFi network connection.",
							Computed:    true,
							Sensitive:   true,
						},
						"sam_report_intv": {
							Type:        schema.TypeInt,
							Description: "SAM report interval (sec), 0 for a one-time report.",
							Computed:    true,
						},
						"sam_security_type": {
							Type:        schema.TypeString,
							Description: "Select WiFi network security type (default = \"wpa-personal\").",
							Computed:    true,
						},
						"sam_server": {
							Type:        schema.TypeString,
							Description: "SAM test server IP address or domain name.",
							Computed:    true,
						},
						"sam_server_fqdn": {
							Type:        schema.TypeString,
							Description: "SAM test server domain name.",
							Computed:    true,
						},
						"sam_server_ip": {
							Type:        schema.TypeString,
							Description: "SAM test server IP address.",
							Computed:    true,
						},
						"sam_server_type": {
							Type:        schema.TypeString,
							Description: "Select SAM server type (default = \"IP\").",
							Computed:    true,
						},
						"sam_ssid": {
							Type:        schema.TypeString,
							Description: "SSID for WiFi network.",
							Computed:    true,
						},
						"sam_test": {
							Type:        schema.TypeString,
							Description: "Select SAM test type (default = \"PING\").",
							Computed:    true,
						},
						"sam_username": {
							Type:        schema.TypeString,
							Description: "Username for WiFi network connection.",
							Computed:    true,
						},
						"short_guard_interval": {
							Type:        schema.TypeString,
							Description: "Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns.",
							Computed:    true,
						},
						"spectrum_analysis": {
							Type:        schema.TypeString,
							Description: "Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.",
							Computed:    true,
						},
						"transmit_optimize": {
							Type:        schema.TypeString,
							Description: "Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default.",
							Computed:    true,
						},
						"vap_all": {
							Type:        schema.TypeString,
							Description: "Configure method for assigning SSIDs to this FortiAP (default = automatically assign tunnel SSIDs).",
							Computed:    true,
						},
						"vaps": {
							Type:        schema.TypeList,
							Description: "Manually selected list of Virtual Access Points (VAPs).",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Virtual Access Point (VAP) name.",
										Computed:    true,
									},
								},
							},
						},
						"wids_profile": {
							Type:        schema.TypeString,
							Description: "Wireless Intrusion Detection System (WIDS) profile name to assign to the radio.",
							Computed:    true,
						},
						"zero_wait_dfs": {
							Type:        schema.TypeString,
							Description: "Enable/disable zero wait DFS on radio (default = enable).",
							Computed:    true,
						},
					},
				},
			},
			"radio_3": {
				Type:        schema.TypeList,
				Description: "Configuration options for radio 3.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"airtime_fairness": {
							Type:        schema.TypeString,
							Description: "Enable/disable airtime fairness (default = disable).",
							Computed:    true,
						},
						"amsdu": {
							Type:        schema.TypeString,
							Description: "Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable).",
							Computed:    true,
						},
						"ap_handoff": {
							Type:        schema.TypeString,
							Description: "Enable/disable AP handoff of clients to other APs (default = disable).",
							Computed:    true,
						},
						"ap_sniffer_addr": {
							Type:        schema.TypeString,
							Description: "MAC address to monitor.",
							Computed:    true,
						},
						"ap_sniffer_bufsize": {
							Type:        schema.TypeInt,
							Description: "Sniffer buffer size (1 - 32 MB, default = 16).",
							Computed:    true,
						},
						"ap_sniffer_chan": {
							Type:        schema.TypeInt,
							Description: "Channel on which to operate the sniffer (default = 6).",
							Computed:    true,
						},
						"ap_sniffer_ctl": {
							Type:        schema.TypeString,
							Description: "Enable/disable sniffer on WiFi control frame (default = enable).",
							Computed:    true,
						},
						"ap_sniffer_data": {
							Type:        schema.TypeString,
							Description: "Enable/disable sniffer on WiFi data frame (default = enable).",
							Computed:    true,
						},
						"ap_sniffer_mgmt_beacon": {
							Type:        schema.TypeString,
							Description: "Enable/disable sniffer on WiFi management Beacon frames (default = enable).",
							Computed:    true,
						},
						"ap_sniffer_mgmt_other": {
							Type:        schema.TypeString,
							Description: "Enable/disable sniffer on WiFi management other frames  (default = enable).",
							Computed:    true,
						},
						"ap_sniffer_mgmt_probe": {
							Type:        schema.TypeString,
							Description: "Enable/disable sniffer on WiFi management probe frames (default = enable).",
							Computed:    true,
						},
						"auto_power_high": {
							Type:        schema.TypeInt,
							Description: "The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).",
							Computed:    true,
						},
						"auto_power_level": {
							Type:        schema.TypeString,
							Description: "Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable).",
							Computed:    true,
						},
						"auto_power_low": {
							Type:        schema.TypeInt,
							Description: "The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).",
							Computed:    true,
						},
						"auto_power_target": {
							Type:        schema.TypeString,
							Description: "The target of automatic transmit power adjustment in dBm. (-95 to -20, default = -70).",
							Computed:    true,
						},
						"band": {
							Type:        schema.TypeString,
							Description: "WiFi band that Radio 3 operates on.",
							Computed:    true,
						},
						"band_5g_type": {
							Type:        schema.TypeString,
							Description: "WiFi 5G band type.",
							Computed:    true,
						},
						"bandwidth_admission_control": {
							Type:        schema.TypeString,
							Description: "Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it.",
							Computed:    true,
						},
						"bandwidth_capacity": {
							Type:        schema.TypeInt,
							Description: "Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).",
							Computed:    true,
						},
						"beacon_interval": {
							Type:        schema.TypeInt,
							Description: "Beacon interval. The time between beacon frames in msec (the actual range of beacon interval depends on the AP platform type, default = 100).",
							Computed:    true,
						},
						"bss_color": {
							Type:        schema.TypeInt,
							Description: "BSS color value for this 11ax radio (0 - 63, 0 means disable. default = 0).",
							Computed:    true,
						},
						"bss_color_mode": {
							Type:        schema.TypeString,
							Description: "BSS color mode for this 11ax radio (default = auto).",
							Computed:    true,
						},
						"call_admission_control": {
							Type:        schema.TypeString,
							Description: "Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them.",
							Computed:    true,
						},
						"call_capacity": {
							Type:        schema.TypeInt,
							Description: "Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).",
							Computed:    true,
						},
						"channel": {
							Type:        schema.TypeList,
							Description: "Selected list of wireless radio channels.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"chan": {
										Type:        schema.TypeString,
										Description: "Channel number.",
										Computed:    true,
									},
								},
							},
						},
						"channel_bonding": {
							Type:        schema.TypeString,
							Description: "Channel bandwidth: 160,80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence.",
							Computed:    true,
						},
						"channel_utilization": {
							Type:        schema.TypeString,
							Description: "Enable/disable measuring channel utilization.",
							Computed:    true,
						},
						"coexistence": {
							Type:        schema.TypeString,
							Description: "Enable/disable allowing both HT20 and HT40 on the same radio (default = enable).",
							Computed:    true,
						},
						"darrp": {
							Type:        schema.TypeString,
							Description: "Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable).",
							Computed:    true,
						},
						"drma": {
							Type:        schema.TypeString,
							Description: "Enable/disable dynamic radio mode assignment (DRMA) (default = disable).",
							Computed:    true,
						},
						"drma_sensitivity": {
							Type:        schema.TypeString,
							Description: "Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low).",
							Computed:    true,
						},
						"dtim": {
							Type:        schema.TypeInt,
							Description: "Delivery Traffic Indication Map (DTIM) period (1 - 255, default = 1). Set higher to save battery life of WiFi client in power-save mode.",
							Computed:    true,
						},
						"frag_threshold": {
							Type:        schema.TypeInt,
							Description: "Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).",
							Computed:    true,
						},
						"frequency_handoff": {
							Type:        schema.TypeString,
							Description: "Enable/disable frequency handoff of clients to other channels (default = disable).",
							Computed:    true,
						},
						"iperf_protocol": {
							Type:        schema.TypeString,
							Description: "Iperf test protocol (default = \"UDP\").",
							Computed:    true,
						},
						"iperf_server_port": {
							Type:        schema.TypeInt,
							Description: "Iperf service port number.",
							Computed:    true,
						},
						"max_clients": {
							Type:        schema.TypeInt,
							Description: "Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.",
							Computed:    true,
						},
						"max_distance": {
							Type:        schema.TypeInt,
							Description: "Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Mode of radio 3. Radio 3 can be disabled, configured as an access point, a rogue AP monitor, a sniffer, or a station.",
							Computed:    true,
						},
						"power_level": {
							Type:        schema.TypeInt,
							Description: "Radio EIRP power level as a percentage of the maximum EIRP power (0 - 100, default = 100).",
							Computed:    true,
						},
						"power_mode": {
							Type:        schema.TypeString,
							Description: "Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities.",
							Computed:    true,
						},
						"power_value": {
							Type:        schema.TypeInt,
							Description: "Radio EIRP power in dBm (1 - 33, default = 27).",
							Computed:    true,
						},
						"powersave_optimize": {
							Type:        schema.TypeString,
							Description: "Enable client power-saving features such as TIM, AC VO, and OBSS etc.",
							Computed:    true,
						},
						"protection_mode": {
							Type:        schema.TypeString,
							Description: "Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable).",
							Computed:    true,
						},
						"rts_threshold": {
							Type:        schema.TypeInt,
							Description: "Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).",
							Computed:    true,
						},
						"sam_bssid": {
							Type:        schema.TypeString,
							Description: "BSSID for WiFi network.",
							Computed:    true,
						},
						"sam_captive_portal": {
							Type:        schema.TypeString,
							Description: "Enable/disable Captive Portal Authentication (default = disable).",
							Computed:    true,
						},
						"sam_cwp_failure_string": {
							Type:        schema.TypeString,
							Description: "Failure identification on the page after an incorrect login.",
							Computed:    true,
						},
						"sam_cwp_match_string": {
							Type:        schema.TypeString,
							Description: "Identification string from the captive portal login form.",
							Computed:    true,
						},
						"sam_cwp_password": {
							Type:        schema.TypeString,
							Description: "Password for captive portal authentication.",
							Computed:    true,
							Sensitive:   true,
						},
						"sam_cwp_success_string": {
							Type:        schema.TypeString,
							Description: "Success identification on the page after a successful login.",
							Computed:    true,
						},
						"sam_cwp_test_url": {
							Type:        schema.TypeString,
							Description: "Website the client is trying to access.",
							Computed:    true,
						},
						"sam_cwp_username": {
							Type:        schema.TypeString,
							Description: "Username for captive portal authentication.",
							Computed:    true,
						},
						"sam_password": {
							Type:        schema.TypeString,
							Description: "Passphrase for WiFi network connection.",
							Computed:    true,
							Sensitive:   true,
						},
						"sam_report_intv": {
							Type:        schema.TypeInt,
							Description: "SAM report interval (sec), 0 for a one-time report.",
							Computed:    true,
						},
						"sam_security_type": {
							Type:        schema.TypeString,
							Description: "Select WiFi network security type (default = \"wpa-personal\").",
							Computed:    true,
						},
						"sam_server": {
							Type:        schema.TypeString,
							Description: "SAM test server IP address or domain name.",
							Computed:    true,
						},
						"sam_server_fqdn": {
							Type:        schema.TypeString,
							Description: "SAM test server domain name.",
							Computed:    true,
						},
						"sam_server_ip": {
							Type:        schema.TypeString,
							Description: "SAM test server IP address.",
							Computed:    true,
						},
						"sam_server_type": {
							Type:        schema.TypeString,
							Description: "Select SAM server type (default = \"IP\").",
							Computed:    true,
						},
						"sam_ssid": {
							Type:        schema.TypeString,
							Description: "SSID for WiFi network.",
							Computed:    true,
						},
						"sam_test": {
							Type:        schema.TypeString,
							Description: "Select SAM test type (default = \"PING\").",
							Computed:    true,
						},
						"sam_username": {
							Type:        schema.TypeString,
							Description: "Username for WiFi network connection.",
							Computed:    true,
						},
						"short_guard_interval": {
							Type:        schema.TypeString,
							Description: "Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns.",
							Computed:    true,
						},
						"spectrum_analysis": {
							Type:        schema.TypeString,
							Description: "Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.",
							Computed:    true,
						},
						"transmit_optimize": {
							Type:        schema.TypeString,
							Description: "Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default.",
							Computed:    true,
						},
						"vap_all": {
							Type:        schema.TypeString,
							Description: "Configure method for assigning SSIDs to this FortiAP (default = automatically assign tunnel SSIDs).",
							Computed:    true,
						},
						"vaps": {
							Type:        schema.TypeList,
							Description: "Manually selected list of Virtual Access Points (VAPs).",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Virtual Access Point (VAP) name.",
										Computed:    true,
									},
								},
							},
						},
						"wids_profile": {
							Type:        schema.TypeString,
							Description: "Wireless Intrusion Detection System (WIDS) profile name to assign to the radio.",
							Computed:    true,
						},
						"zero_wait_dfs": {
							Type:        schema.TypeString,
							Description: "Enable/disable zero wait DFS on radio (default = enable).",
							Computed:    true,
						},
					},
				},
			},
			"radio_4": {
				Type:        schema.TypeList,
				Description: "Configuration options for radio 4.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"airtime_fairness": {
							Type:        schema.TypeString,
							Description: "Enable/disable airtime fairness (default = disable).",
							Computed:    true,
						},
						"amsdu": {
							Type:        schema.TypeString,
							Description: "Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable).",
							Computed:    true,
						},
						"ap_handoff": {
							Type:        schema.TypeString,
							Description: "Enable/disable AP handoff of clients to other APs (default = disable).",
							Computed:    true,
						},
						"ap_sniffer_addr": {
							Type:        schema.TypeString,
							Description: "MAC address to monitor.",
							Computed:    true,
						},
						"ap_sniffer_bufsize": {
							Type:        schema.TypeInt,
							Description: "Sniffer buffer size (1 - 32 MB, default = 16).",
							Computed:    true,
						},
						"ap_sniffer_chan": {
							Type:        schema.TypeInt,
							Description: "Channel on which to operate the sniffer (default = 6).",
							Computed:    true,
						},
						"ap_sniffer_ctl": {
							Type:        schema.TypeString,
							Description: "Enable/disable sniffer on WiFi control frame (default = enable).",
							Computed:    true,
						},
						"ap_sniffer_data": {
							Type:        schema.TypeString,
							Description: "Enable/disable sniffer on WiFi data frame (default = enable).",
							Computed:    true,
						},
						"ap_sniffer_mgmt_beacon": {
							Type:        schema.TypeString,
							Description: "Enable/disable sniffer on WiFi management Beacon frames (default = enable).",
							Computed:    true,
						},
						"ap_sniffer_mgmt_other": {
							Type:        schema.TypeString,
							Description: "Enable/disable sniffer on WiFi management other frames  (default = enable).",
							Computed:    true,
						},
						"ap_sniffer_mgmt_probe": {
							Type:        schema.TypeString,
							Description: "Enable/disable sniffer on WiFi management probe frames (default = enable).",
							Computed:    true,
						},
						"auto_power_high": {
							Type:        schema.TypeInt,
							Description: "The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).",
							Computed:    true,
						},
						"auto_power_level": {
							Type:        schema.TypeString,
							Description: "Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable).",
							Computed:    true,
						},
						"auto_power_low": {
							Type:        schema.TypeInt,
							Description: "The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).",
							Computed:    true,
						},
						"auto_power_target": {
							Type:        schema.TypeString,
							Description: "The target of automatic transmit power adjustment in dBm. (-95 to -20, default = -70).",
							Computed:    true,
						},
						"band": {
							Type:        schema.TypeString,
							Description: "WiFi band that Radio 3 operates on.",
							Computed:    true,
						},
						"band_5g_type": {
							Type:        schema.TypeString,
							Description: "WiFi 5G band type.",
							Computed:    true,
						},
						"bandwidth_admission_control": {
							Type:        schema.TypeString,
							Description: "Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it.",
							Computed:    true,
						},
						"bandwidth_capacity": {
							Type:        schema.TypeInt,
							Description: "Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).",
							Computed:    true,
						},
						"beacon_interval": {
							Type:        schema.TypeInt,
							Description: "Beacon interval. The time between beacon frames in msec (the actual range of beacon interval depends on the AP platform type, default = 100).",
							Computed:    true,
						},
						"bss_color": {
							Type:        schema.TypeInt,
							Description: "BSS color value for this 11ax radio (0 - 63, 0 means disable. default = 0).",
							Computed:    true,
						},
						"bss_color_mode": {
							Type:        schema.TypeString,
							Description: "BSS color mode for this 11ax radio (default = auto).",
							Computed:    true,
						},
						"call_admission_control": {
							Type:        schema.TypeString,
							Description: "Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them.",
							Computed:    true,
						},
						"call_capacity": {
							Type:        schema.TypeInt,
							Description: "Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).",
							Computed:    true,
						},
						"channel": {
							Type:        schema.TypeList,
							Description: "Selected list of wireless radio channels.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"chan": {
										Type:        schema.TypeString,
										Description: "Channel number.",
										Computed:    true,
									},
								},
							},
						},
						"channel_bonding": {
							Type:        schema.TypeString,
							Description: "Channel bandwidth: 160,80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence.",
							Computed:    true,
						},
						"channel_utilization": {
							Type:        schema.TypeString,
							Description: "Enable/disable measuring channel utilization.",
							Computed:    true,
						},
						"coexistence": {
							Type:        schema.TypeString,
							Description: "Enable/disable allowing both HT20 and HT40 on the same radio (default = enable).",
							Computed:    true,
						},
						"darrp": {
							Type:        schema.TypeString,
							Description: "Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable).",
							Computed:    true,
						},
						"drma": {
							Type:        schema.TypeString,
							Description: "Enable/disable dynamic radio mode assignment (DRMA) (default = disable).",
							Computed:    true,
						},
						"drma_sensitivity": {
							Type:        schema.TypeString,
							Description: "Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low).",
							Computed:    true,
						},
						"dtim": {
							Type:        schema.TypeInt,
							Description: "Delivery Traffic Indication Map (DTIM) period (1 - 255, default = 1). Set higher to save battery life of WiFi client in power-save mode.",
							Computed:    true,
						},
						"frag_threshold": {
							Type:        schema.TypeInt,
							Description: "Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).",
							Computed:    true,
						},
						"frequency_handoff": {
							Type:        schema.TypeString,
							Description: "Enable/disable frequency handoff of clients to other channels (default = disable).",
							Computed:    true,
						},
						"iperf_protocol": {
							Type:        schema.TypeString,
							Description: "Iperf test protocol (default = \"UDP\").",
							Computed:    true,
						},
						"iperf_server_port": {
							Type:        schema.TypeInt,
							Description: "Iperf service port number.",
							Computed:    true,
						},
						"max_clients": {
							Type:        schema.TypeInt,
							Description: "Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.",
							Computed:    true,
						},
						"max_distance": {
							Type:        schema.TypeInt,
							Description: "Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Mode of radio 3. Radio 3 can be disabled, configured as an access point, a rogue AP monitor, a sniffer, or a station.",
							Computed:    true,
						},
						"power_level": {
							Type:        schema.TypeInt,
							Description: "Radio EIRP power level as a percentage of the maximum EIRP power (0 - 100, default = 100).",
							Computed:    true,
						},
						"power_mode": {
							Type:        schema.TypeString,
							Description: "Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities.",
							Computed:    true,
						},
						"power_value": {
							Type:        schema.TypeInt,
							Description: "Radio EIRP power in dBm (1 - 33, default = 27).",
							Computed:    true,
						},
						"powersave_optimize": {
							Type:        schema.TypeString,
							Description: "Enable client power-saving features such as TIM, AC VO, and OBSS etc.",
							Computed:    true,
						},
						"protection_mode": {
							Type:        schema.TypeString,
							Description: "Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable).",
							Computed:    true,
						},
						"rts_threshold": {
							Type:        schema.TypeInt,
							Description: "Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).",
							Computed:    true,
						},
						"sam_bssid": {
							Type:        schema.TypeString,
							Description: "BSSID for WiFi network.",
							Computed:    true,
						},
						"sam_captive_portal": {
							Type:        schema.TypeString,
							Description: "Enable/disable Captive Portal Authentication (default = disable).",
							Computed:    true,
						},
						"sam_cwp_failure_string": {
							Type:        schema.TypeString,
							Description: "Failure identification on the page after an incorrect login.",
							Computed:    true,
						},
						"sam_cwp_match_string": {
							Type:        schema.TypeString,
							Description: "Identification string from the captive portal login form.",
							Computed:    true,
						},
						"sam_cwp_password": {
							Type:        schema.TypeString,
							Description: "Password for captive portal authentication.",
							Computed:    true,
							Sensitive:   true,
						},
						"sam_cwp_success_string": {
							Type:        schema.TypeString,
							Description: "Success identification on the page after a successful login.",
							Computed:    true,
						},
						"sam_cwp_test_url": {
							Type:        schema.TypeString,
							Description: "Website the client is trying to access.",
							Computed:    true,
						},
						"sam_cwp_username": {
							Type:        schema.TypeString,
							Description: "Username for captive portal authentication.",
							Computed:    true,
						},
						"sam_password": {
							Type:        schema.TypeString,
							Description: "Passphrase for WiFi network connection.",
							Computed:    true,
							Sensitive:   true,
						},
						"sam_report_intv": {
							Type:        schema.TypeInt,
							Description: "SAM report interval (sec), 0 for a one-time report.",
							Computed:    true,
						},
						"sam_security_type": {
							Type:        schema.TypeString,
							Description: "Select WiFi network security type (default = \"wpa-personal\").",
							Computed:    true,
						},
						"sam_server": {
							Type:        schema.TypeString,
							Description: "SAM test server IP address or domain name.",
							Computed:    true,
						},
						"sam_server_fqdn": {
							Type:        schema.TypeString,
							Description: "SAM test server domain name.",
							Computed:    true,
						},
						"sam_server_ip": {
							Type:        schema.TypeString,
							Description: "SAM test server IP address.",
							Computed:    true,
						},
						"sam_server_type": {
							Type:        schema.TypeString,
							Description: "Select SAM server type (default = \"IP\").",
							Computed:    true,
						},
						"sam_ssid": {
							Type:        schema.TypeString,
							Description: "SSID for WiFi network.",
							Computed:    true,
						},
						"sam_test": {
							Type:        schema.TypeString,
							Description: "Select SAM test type (default = \"PING\").",
							Computed:    true,
						},
						"sam_username": {
							Type:        schema.TypeString,
							Description: "Username for WiFi network connection.",
							Computed:    true,
						},
						"short_guard_interval": {
							Type:        schema.TypeString,
							Description: "Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns.",
							Computed:    true,
						},
						"spectrum_analysis": {
							Type:        schema.TypeString,
							Description: "Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.",
							Computed:    true,
						},
						"transmit_optimize": {
							Type:        schema.TypeString,
							Description: "Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default.",
							Computed:    true,
						},
						"vap_all": {
							Type:        schema.TypeString,
							Description: "Configure method for assigning SSIDs to this FortiAP (default = automatically assign tunnel SSIDs).",
							Computed:    true,
						},
						"vaps": {
							Type:        schema.TypeList,
							Description: "Manually selected list of Virtual Access Points (VAPs).",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Virtual Access Point (VAP) name.",
										Computed:    true,
									},
								},
							},
						},
						"wids_profile": {
							Type:        schema.TypeString,
							Description: "Wireless Intrusion Detection System (WIDS) profile name to assign to the radio.",
							Computed:    true,
						},
						"zero_wait_dfs": {
							Type:        schema.TypeString,
							Description: "Enable/disable zero wait DFS on radio (default = enable).",
							Computed:    true,
						},
					},
				},
			},
			"split_tunneling_acl": {
				Type:        schema.TypeList,
				Description: "Split tunneling ACL filter list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dest_ip": {
							Type:        schema.TypeString,
							Description: "Destination IP and mask for the split-tunneling subnet.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
					},
				},
			},
			"split_tunneling_acl_local_ap_subnet": {
				Type:        schema.TypeString,
				Description: "Enable/disable automatically adding local subnetwork of FortiAP to split-tunneling ACL (default = disable).",
				Computed:    true,
			},
			"split_tunneling_acl_path": {
				Type:        schema.TypeString,
				Description: "Split tunneling ACL path is local/tunnel.",
				Computed:    true,
			},
			"syslog_profile": {
				Type:        schema.TypeString,
				Description: "System log server configuration profile name.",
				Computed:    true,
			},
			"tun_mtu_downlink": {
				Type:        schema.TypeInt,
				Description: "The MTU of downlink CAPWAP tunnel (576 - 1500 bytes or 0; 0 means the local MTU of FortiAP; default = 0).",
				Computed:    true,
			},
			"tun_mtu_uplink": {
				Type:        schema.TypeInt,
				Description: "The maximum transmission unit (MTU) of uplink CAPWAP tunnel (576 - 1500 bytes or 0; 0 means the local MTU of FortiAP; default = 0).",
				Computed:    true,
			},
			"wan_port_auth": {
				Type:        schema.TypeString,
				Description: "Set WAN port authentication mode (default = none).",
				Computed:    true,
			},
			"wan_port_auth_methods": {
				Type:        schema.TypeString,
				Description: "WAN port 802.1x supplicant EAP methods (default = all).",
				Computed:    true,
			},
			"wan_port_auth_password": {
				Type:        schema.TypeString,
				Description: "Set WAN port 802.1x supplicant password.",
				Computed:    true,
				Sensitive:   true,
			},
			"wan_port_auth_usrname": {
				Type:        schema.TypeString,
				Description: "Set WAN port 802.1x supplicant user name.",
				Computed:    true,
			},
			"wan_port_mode": {
				Type:        schema.TypeString,
				Description: "Enable/disable using a WAN port as a LAN port.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWirelessControllerWtpProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

	c := meta.(*apiClient).Client
	// c.Retries = 1

	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	o, err := c.Cmdb.ReadWirelessControllerWtpProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerWtpProfile dataSource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] dataSource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	sort := false
	if v, ok := d.GetOk("dynamic_sort_subtable"); ok {
		if b, ok := v.(bool); ok {
			sort = b
		}
	}

	diags := refreshObjectWirelessControllerWtpProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
