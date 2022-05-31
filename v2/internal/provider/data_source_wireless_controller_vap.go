// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Virtual Access Points (VAPs).

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWirelessControllerVap() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Virtual Access Points (VAPs).",

		ReadContext: dataSourceWirelessControllerVapRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"access_control_list": {
				Type:        schema.TypeString,
				Description: "Profile name for access-control-list.",
				Computed:    true,
			},
			"acct_interim_interval": {
				Type:        schema.TypeInt,
				Description: "WiFi RADIUS accounting interim interval (60 - 86400 sec, default = 0).",
				Computed:    true,
			},
			"additional_akms": {
				Type:        schema.TypeString,
				Description: "Additional AKMs.",
				Computed:    true,
			},
			"address_group": {
				Type:        schema.TypeString,
				Description: "Firewall Address Group Name.",
				Computed:    true,
			},
			"address_group_policy": {
				Type:        schema.TypeString,
				Description: "Configure MAC address filtering policy for MAC addresses that are in the address-group.",
				Computed:    true,
			},
			"antivirus_profile": {
				Type:        schema.TypeString,
				Description: "AntiVirus profile name.",
				Computed:    true,
			},
			"application_detection_engine": {
				Type:        schema.TypeString,
				Description: "Enable/disable application detection engine (default = disable).",
				Computed:    true,
			},
			"application_list": {
				Type:        schema.TypeString,
				Description: "Application control list name.",
				Computed:    true,
			},
			"application_report_intv": {
				Type:        schema.TypeInt,
				Description: "Application report interval (30 - 864000 sec, default = 120).",
				Computed:    true,
			},
			"atf_weight": {
				Type:        schema.TypeInt,
				Description: "Airtime weight in percentage (default = 20).",
				Computed:    true,
			},
			"auth": {
				Type:        schema.TypeString,
				Description: "Authentication protocol.",
				Computed:    true,
			},
			"auth_cert": {
				Type:        schema.TypeString,
				Description: "HTTPS server certificate.",
				Computed:    true,
			},
			"auth_portal_addr": {
				Type:        schema.TypeString,
				Description: "Address of captive portal.",
				Computed:    true,
			},
			"beacon_advertising": {
				Type:        schema.TypeString,
				Description: "Fortinet beacon advertising IE data   (default = empty).",
				Computed:    true,
			},
			"broadcast_ssid": {
				Type:        schema.TypeString,
				Description: "Enable/disable broadcasting the SSID (default = enable).",
				Computed:    true,
			},
			"broadcast_suppression": {
				Type:        schema.TypeString,
				Description: "Optional suppression of broadcast messages. For example, you can keep DHCP messages, ARP broadcasts, and so on off of the wireless network.",
				Computed:    true,
			},
			"bss_color_partial": {
				Type:        schema.TypeString,
				Description: "Enable/disable 802.11ax partial BSS color (default = enable).",
				Computed:    true,
			},
			"bstm_disassociation_imminent": {
				Type:        schema.TypeString,
				Description: "Enable/disable forcing of disassociation after the BSTM request timer has been reached (default = enable).",
				Computed:    true,
			},
			"bstm_load_balancing_disassoc_timer": {
				Type:        schema.TypeInt,
				Description: "Time interval for client to voluntarily leave AP before forcing a disassociation due to AP load-balancing (0 to 30, default = 10).",
				Computed:    true,
			},
			"bstm_rssi_disassoc_timer": {
				Type:        schema.TypeInt,
				Description: "Time interval for client to voluntarily leave AP before forcing a disassociation due to low RSSI (0 to 2000, default = 200).",
				Computed:    true,
			},
			"captive_portal_ac_name": {
				Type:        schema.TypeString,
				Description: "Local-bridging captive portal ac-name.",
				Computed:    true,
			},
			"captive_portal_auth_timeout": {
				Type:        schema.TypeInt,
				Description: "Hard timeout - AP will always clear the session after timeout regardless of traffic (0 - 864000 sec, default = 0).",
				Computed:    true,
			},
			"captive_portal_macauth_radius_secret": {
				Type:        schema.TypeString,
				Description: "Secret key to access the macauth RADIUS server.",
				Computed:    true,
				Sensitive:   true,
			},
			"captive_portal_macauth_radius_server": {
				Type:        schema.TypeString,
				Description: "Captive portal external RADIUS server domain name or IP address.",
				Computed:    true,
			},
			"captive_portal_radius_secret": {
				Type:        schema.TypeString,
				Description: "Secret key to access the RADIUS server.",
				Computed:    true,
				Sensitive:   true,
			},
			"captive_portal_radius_server": {
				Type:        schema.TypeString,
				Description: "Captive portal RADIUS server domain name or IP address.",
				Computed:    true,
			},
			"captive_portal_session_timeout_interval": {
				Type:        schema.TypeInt,
				Description: "Session timeout interval (0 - 864000 sec, default = 0).",
				Computed:    true,
			},
			"dhcp_address_enforcement": {
				Type:        schema.TypeString,
				Description: "Enable/disable DHCP address enforcement (default = disable).",
				Computed:    true,
			},
			"dhcp_lease_time": {
				Type:        schema.TypeInt,
				Description: "DHCP lease time in seconds for NAT IP address.",
				Computed:    true,
			},
			"dhcp_option43_insertion": {
				Type:        schema.TypeString,
				Description: "Enable/disable insertion of DHCP option 43 (default = enable).",
				Computed:    true,
			},
			"dhcp_option82_circuit_id_insertion": {
				Type:        schema.TypeString,
				Description: "Enable/disable DHCP option 82 circuit-id insert (default = disable).",
				Computed:    true,
			},
			"dhcp_option82_insertion": {
				Type:        schema.TypeString,
				Description: "Enable/disable DHCP option 82 insert (default = disable).",
				Computed:    true,
			},
			"dhcp_option82_remote_id_insertion": {
				Type:        schema.TypeString,
				Description: "Enable/disable DHCP option 82 remote-id insert (default = disable).",
				Computed:    true,
			},
			"dynamic_vlan": {
				Type:        schema.TypeString,
				Description: "Enable/disable dynamic VLAN assignment.",
				Computed:    true,
			},
			"eap_reauth": {
				Type:        schema.TypeString,
				Description: "Enable/disable EAP re-authentication for WPA-Enterprise security.",
				Computed:    true,
			},
			"eap_reauth_intv": {
				Type:        schema.TypeInt,
				Description: "EAP re-authentication interval (1800 - 864000 sec, default = 86400).",
				Computed:    true,
			},
			"eapol_key_retries": {
				Type:        schema.TypeString,
				Description: "Enable/disable retransmission of EAPOL-Key frames (message 3/4 and group message 1/2) (default = enable).",
				Computed:    true,
			},
			"encrypt": {
				Type:        schema.TypeString,
				Description: "Encryption protocol to use (only available when security is set to a WPA type).",
				Computed:    true,
			},
			"external_fast_roaming": {
				Type:        schema.TypeString,
				Description: "Enable/disable fast roaming or pre-authentication with external APs not managed by the FortiGate (default = disable).",
				Computed:    true,
			},
			"external_logout": {
				Type:        schema.TypeString,
				Description: "URL of external authentication logout server.",
				Computed:    true,
			},
			"external_web": {
				Type:        schema.TypeString,
				Description: "URL of external authentication web server.",
				Computed:    true,
			},
			"external_web_format": {
				Type:        schema.TypeString,
				Description: "URL query parameter detection (default = auto-detect).",
				Computed:    true,
			},
			"fast_bss_transition": {
				Type:        schema.TypeString,
				Description: "Enable/disable 802.11r Fast BSS Transition (FT) (default = disable).",
				Computed:    true,
			},
			"fast_roaming": {
				Type:        schema.TypeString,
				Description: "Enable/disable fast-roaming, or pre-authentication, where supported by clients (default = disable).",
				Computed:    true,
			},
			"ft_mobility_domain": {
				Type:        schema.TypeInt,
				Description: "Mobility domain identifier in FT (1 - 65535, default = 1000).",
				Computed:    true,
			},
			"ft_over_ds": {
				Type:        schema.TypeString,
				Description: "Enable/disable FT over the Distribution System (DS).",
				Computed:    true,
			},
			"ft_r0_key_lifetime": {
				Type:        schema.TypeInt,
				Description: "Lifetime of the PMK-R0 key in FT, 1-65535 minutes.",
				Computed:    true,
			},
			"gas_comeback_delay": {
				Type:        schema.TypeInt,
				Description: "GAS comeback delay (0 or 100 - 10000 milliseconds, default = 500).",
				Computed:    true,
			},
			"gas_fragmentation_limit": {
				Type:        schema.TypeInt,
				Description: "GAS fragmentation limit (512 - 4096, default = 1024).",
				Computed:    true,
			},
			"gtk_rekey": {
				Type:        schema.TypeString,
				Description: "Enable/disable GTK rekey for WPA security.",
				Computed:    true,
			},
			"gtk_rekey_intv": {
				Type:        schema.TypeInt,
				Description: "GTK rekey interval (1800 - 864000 sec, default = 86400).",
				Computed:    true,
			},
			"high_efficiency": {
				Type:        schema.TypeString,
				Description: "Enable/disable 802.11ax high efficiency (default = enable).",
				Computed:    true,
			},
			"hotspot20_profile": {
				Type:        schema.TypeString,
				Description: "Hotspot 2.0 profile name.",
				Computed:    true,
			},
			"igmp_snooping": {
				Type:        schema.TypeString,
				Description: "Enable/disable IGMP snooping.",
				Computed:    true,
			},
			"intra_vap_privacy": {
				Type:        schema.TypeString,
				Description: "Enable/disable blocking communication between clients on the same SSID (called intra-SSID privacy) (default = disable).",
				Computed:    true,
			},
			"ip": {
				Type:        schema.TypeString,
				Description: "IP address and subnet mask for the local standalone NAT subnet.",
				Computed:    true,
			},
			"ips_sensor": {
				Type:        schema.TypeString,
				Description: "IPS sensor name.",
				Computed:    true,
			},
			"ipv6_rules": {
				Type:        schema.TypeString,
				Description: "Optional rules of IPv6 packets. For example, you can keep RA, RS and so on off of the wireless network.",
				Computed:    true,
			},
			"key": {
				Type:        schema.TypeString,
				Description: "WEP Key.",
				Computed:    true,
				Sensitive:   true,
			},
			"keyindex": {
				Type:        schema.TypeInt,
				Description: "WEP key index (1 - 4).",
				Computed:    true,
			},
			"l3_roaming": {
				Type:        schema.TypeString,
				Description: "Enable/disable layer 3 roaming (default = disable).",
				Computed:    true,
			},
			"ldpc": {
				Type:        schema.TypeString,
				Description: "VAP low-density parity-check (LDPC) coding configuration.",
				Computed:    true,
			},
			"local_authentication": {
				Type:        schema.TypeString,
				Description: "Enable/disable AP local authentication.",
				Computed:    true,
			},
			"local_bridging": {
				Type:        schema.TypeString,
				Description: "Enable/disable bridging of wireless and Ethernet interfaces on the FortiAP (default = disable).",
				Computed:    true,
			},
			"local_lan": {
				Type:        schema.TypeString,
				Description: "Allow/deny traffic destined for a Class A, B, or C private IP address (default = allow).",
				Computed:    true,
			},
			"local_standalone": {
				Type:        schema.TypeString,
				Description: "Enable/disable AP local standalone (default = disable).",
				Computed:    true,
			},
			"local_standalone_dns": {
				Type:        schema.TypeString,
				Description: "Enable/disable AP local standalone DNS.",
				Computed:    true,
			},
			"local_standalone_dns_ip": {
				Type:        schema.TypeString,
				Description: "IPv4 addresses for the local standalone DNS.",
				Computed:    true,
			},
			"local_standalone_nat": {
				Type:        schema.TypeString,
				Description: "Enable/disable AP local standalone NAT mode.",
				Computed:    true,
			},
			"mac_auth_bypass": {
				Type:        schema.TypeString,
				Description: "Enable/disable MAC authentication bypass.",
				Computed:    true,
			},
			"mac_called_station_delimiter": {
				Type:        schema.TypeString,
				Description: "MAC called station delimiter (default = hyphen).",
				Computed:    true,
			},
			"mac_calling_station_delimiter": {
				Type:        schema.TypeString,
				Description: "MAC calling station delimiter (default = hyphen).",
				Computed:    true,
			},
			"mac_case": {
				Type:        schema.TypeString,
				Description: "MAC case (default = uppercase).",
				Computed:    true,
			},
			"mac_filter": {
				Type:        schema.TypeString,
				Description: "Enable/disable MAC filtering to block wireless clients by mac address.",
				Computed:    true,
			},
			"mac_filter_list": {
				Type:        schema.TypeList,
				Description: "Create a list of MAC addresses for MAC address filtering.",
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
							Description: "MAC address.",
							Computed:    true,
						},
						"mac_filter_policy": {
							Type:        schema.TypeString,
							Description: "Deny or allow the client with this MAC address.",
							Computed:    true,
						},
					},
				},
			},
			"mac_filter_policy_other": {
				Type:        schema.TypeString,
				Description: "Allow or block clients with MAC addresses that are not in the filter list.",
				Computed:    true,
			},
			"mac_password_delimiter": {
				Type:        schema.TypeString,
				Description: "MAC authentication password delimiter (default = hyphen).",
				Computed:    true,
			},
			"mac_username_delimiter": {
				Type:        schema.TypeString,
				Description: "MAC authentication username delimiter (default = hyphen).",
				Computed:    true,
			},
			"max_clients": {
				Type:        schema.TypeInt,
				Description: "Maximum number of clients that can connect simultaneously to the VAP (default = 0, meaning no limitation).",
				Computed:    true,
			},
			"max_clients_ap": {
				Type:        schema.TypeInt,
				Description: "Maximum number of clients that can connect simultaneously to the VAP per AP radio (default = 0, meaning no limitation).",
				Computed:    true,
			},
			"mbo": {
				Type:        schema.TypeString,
				Description: "Enable/disable Multiband Operation (default = disable).",
				Computed:    true,
			},
			"mbo_cell_data_conn_pref": {
				Type:        schema.TypeString,
				Description: "MBO cell data connection preference (0, 1, or 255, default = 1).",
				Computed:    true,
			},
			"me_disable_thresh": {
				Type:        schema.TypeInt,
				Description: "Disable multicast enhancement when this many clients are receiving multicast traffic.",
				Computed:    true,
			},
			"mesh_backhaul": {
				Type:        schema.TypeString,
				Description: "Enable/disable using this VAP as a WiFi mesh backhaul (default = disable). This entry is only available when security is set to a WPA type or open.",
				Computed:    true,
			},
			"mpsk": {
				Type:        schema.TypeString,
				Description: "Enable/disable multiple PSK authentication.",
				Computed:    true,
			},
			"mpsk_concurrent_clients": {
				Type:        schema.TypeInt,
				Description: "Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).",
				Computed:    true,
			},
			"mpsk_key": {
				Type:        schema.TypeList,
				Description: "List of multiple PSK entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comment": {
							Type:        schema.TypeString,
							Description: "Comment.",
							Computed:    true,
						},
						"concurrent_clients": {
							Type:        schema.TypeString,
							Description: "Number of clients that can connect using this pre-shared key.",
							Computed:    true,
						},
						"key_name": {
							Type:        schema.TypeString,
							Description: "Pre-shared key name.",
							Computed:    true,
						},
						"mpsk_schedules": {
							Type:        schema.TypeList,
							Description: "Firewall schedule for MPSK passphrase. The passphrase will be effective only when at least one schedule is valid.",
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
						"passphrase": {
							Type:        schema.TypeString,
							Description: "WPA Pre-shared key.",
							Computed:    true,
							Sensitive:   true,
						},
					},
				},
			},
			"mpsk_profile": {
				Type:        schema.TypeString,
				Description: "MPSK profile name.",
				Computed:    true,
			},
			"mu_mimo": {
				Type:        schema.TypeString,
				Description: "Enable/disable Multi-user MIMO (default = enable).",
				Computed:    true,
			},
			"multicast_enhance": {
				Type:        schema.TypeString,
				Description: "Enable/disable converting multicast to unicast to improve performance (default = disable).",
				Computed:    true,
			},
			"multicast_rate": {
				Type:        schema.TypeString,
				Description: "Multicast rate (0, 6000, 12000, or 24000 kbps, default = 0).",
				Computed:    true,
			},
			"nac": {
				Type:        schema.TypeString,
				Description: "Enable/disable network access control.",
				Computed:    true,
			},
			"nac_profile": {
				Type:        schema.TypeString,
				Description: "NAC profile name.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Virtual AP name.",
				Required:    true,
			},
			"neighbor_report_dual_band": {
				Type:        schema.TypeString,
				Description: "Enable/disable dual-band neighbor report (default = disable).",
				Computed:    true,
			},
			"okc": {
				Type:        schema.TypeString,
				Description: "Enable/disable Opportunistic Key Caching (OKC) (default = enable).",
				Computed:    true,
			},
			"osen": {
				Type:        schema.TypeString,
				Description: "Enable/disable OSEN as part of key management (default = disable).",
				Computed:    true,
			},
			"owe_groups": {
				Type:        schema.TypeString,
				Description: "OWE-Groups.",
				Computed:    true,
			},
			"owe_transition": {
				Type:        schema.TypeString,
				Description: "Enable/disable OWE transition mode support.",
				Computed:    true,
			},
			"owe_transition_ssid": {
				Type:        schema.TypeString,
				Description: "OWE transition mode peer SSID.",
				Computed:    true,
			},
			"passphrase": {
				Type:        schema.TypeString,
				Description: "WPA pre-shared key (PSK) to be used to authenticate WiFi users.",
				Computed:    true,
				Sensitive:   true,
			},
			"pmf": {
				Type:        schema.TypeString,
				Description: "Protected Management Frames (PMF) support (default = disable).",
				Computed:    true,
			},
			"pmf_assoc_comeback_timeout": {
				Type:        schema.TypeInt,
				Description: "Protected Management Frames (PMF) comeback maximum timeout (1-20 sec).",
				Computed:    true,
			},
			"pmf_sa_query_retry_timeout": {
				Type:        schema.TypeInt,
				Description: "Protected Management Frames (PMF) SA query retry timeout interval (1 - 5 100s of msec).",
				Computed:    true,
			},
			"port_macauth": {
				Type:        schema.TypeString,
				Description: "Enable/disable LAN port MAC authentication (default = disable).",
				Computed:    true,
			},
			"port_macauth_reauth_timeout": {
				Type:        schema.TypeInt,
				Description: "LAN port MAC authentication re-authentication timeout value (default = 7200 sec).",
				Computed:    true,
			},
			"port_macauth_timeout": {
				Type:        schema.TypeInt,
				Description: "LAN port MAC authentication idle timeout value (default = 600 sec).",
				Computed:    true,
			},
			"portal_message_override_group": {
				Type:        schema.TypeString,
				Description: "Replacement message group for this VAP (only available when security is set to a captive portal type).",
				Computed:    true,
			},
			"portal_message_overrides": {
				Type:        schema.TypeList,
				Description: "Individual message overrides.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_disclaimer_page": {
							Type:        schema.TypeString,
							Description: "Override auth-disclaimer-page message with message from portal-message-overrides group.",
							Computed:    true,
						},
						"auth_login_failed_page": {
							Type:        schema.TypeString,
							Description: "Override auth-login-failed-page message with message from portal-message-overrides group.",
							Computed:    true,
						},
						"auth_login_page": {
							Type:        schema.TypeString,
							Description: "Override auth-login-page message with message from portal-message-overrides group.",
							Computed:    true,
						},
						"auth_reject_page": {
							Type:        schema.TypeString,
							Description: "Override auth-reject-page message with message from portal-message-overrides group.",
							Computed:    true,
						},
					},
				},
			},
			"portal_type": {
				Type:        schema.TypeString,
				Description: "Captive portal functionality. Configure how the captive portal authenticates users and whether it includes a disclaimer.",
				Computed:    true,
			},
			"primary_wag_profile": {
				Type:        schema.TypeString,
				Description: "Primary wireless access gateway profile name.",
				Computed:    true,
			},
			"probe_resp_suppression": {
				Type:        schema.TypeString,
				Description: "Enable/disable probe response suppression (to ignore weak signals) (default = disable).",
				Computed:    true,
			},
			"probe_resp_threshold": {
				Type:        schema.TypeString,
				Description: "Minimum signal level/threshold in dBm required for the AP response to probe requests (-95 to -20, default = -80).",
				Computed:    true,
			},
			"ptk_rekey": {
				Type:        schema.TypeString,
				Description: "Enable/disable PTK rekey for WPA-Enterprise security.",
				Computed:    true,
			},
			"ptk_rekey_intv": {
				Type:        schema.TypeInt,
				Description: "PTK rekey interval (1800 - 864000 sec, default = 86400).",
				Computed:    true,
			},
			"qos_profile": {
				Type:        schema.TypeString,
				Description: "Quality of service profile name.",
				Computed:    true,
			},
			"quarantine": {
				Type:        schema.TypeString,
				Description: "Enable/disable station quarantine (default = enable).",
				Computed:    true,
			},
			"radio_2g_threshold": {
				Type:        schema.TypeString,
				Description: "Minimum signal level/threshold in dBm required for the AP response to receive a packet in 2.4G band (-95 to -20, default = -79).",
				Computed:    true,
			},
			"radio_5g_threshold": {
				Type:        schema.TypeString,
				Description: "Minimum signal level/threshold in dBm required for the AP response to receive a packet in 5G band(-95 to -20, default = -76).",
				Computed:    true,
			},
			"radio_sensitivity": {
				Type:        schema.TypeString,
				Description: "Enable/disable software radio sensitivity (to ignore weak signals) (default = disable).",
				Computed:    true,
			},
			"radius_mac_auth": {
				Type:        schema.TypeString,
				Description: "Enable/disable RADIUS-based MAC authentication of clients (default = disable).",
				Computed:    true,
			},
			"radius_mac_auth_server": {
				Type:        schema.TypeString,
				Description: "RADIUS-based MAC authentication server.",
				Computed:    true,
			},
			"radius_mac_auth_usergroups": {
				Type:        schema.TypeList,
				Description: "Selective user groups that are permitted for RADIUS mac authentication.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "User group name.",
							Computed:    true,
						},
					},
				},
			},
			"radius_mac_mpsk_auth": {
				Type:        schema.TypeString,
				Description: "Enable/disable RADIUS-based MAC authentication of clients for MPSK authentication (default = disable).",
				Computed:    true,
			},
			"radius_mac_mpsk_timeout": {
				Type:        schema.TypeInt,
				Description: "RADIUS MAC MPSK cache timeout interval (1800 - 864000, default = 86400).",
				Computed:    true,
			},
			"radius_server": {
				Type:        schema.TypeString,
				Description: "RADIUS server to be used to authenticate WiFi users.",
				Computed:    true,
			},
			"rates_11a": {
				Type:        schema.TypeString,
				Description: "Allowed data rates for 802.11a.",
				Computed:    true,
			},
			"rates_11ac_ss12": {
				Type:        schema.TypeString,
				Description: "Allowed data rates for 802.11ac with 1 or 2 spatial streams.",
				Computed:    true,
			},
			"rates_11ac_ss34": {
				Type:        schema.TypeString,
				Description: "Allowed data rates for 802.11ac with 3 or 4 spatial streams.",
				Computed:    true,
			},
			"rates_11ax_ss12": {
				Type:        schema.TypeString,
				Description: "Allowed data rates for 802.11ax with 1 or 2 spatial streams.",
				Computed:    true,
			},
			"rates_11ax_ss34": {
				Type:        schema.TypeString,
				Description: "Allowed data rates for 802.11ax with 3 or 4 spatial streams.",
				Computed:    true,
			},
			"rates_11bg": {
				Type:        schema.TypeString,
				Description: "Allowed data rates for 802.11b/g.",
				Computed:    true,
			},
			"rates_11n_ss12": {
				Type:        schema.TypeString,
				Description: "Allowed data rates for 802.11n with 1 or 2 spatial streams.",
				Computed:    true,
			},
			"rates_11n_ss34": {
				Type:        schema.TypeString,
				Description: "Allowed data rates for 802.11n with 3 or 4 spatial streams.",
				Computed:    true,
			},
			"sae_groups": {
				Type:        schema.TypeString,
				Description: "SAE-Groups.",
				Computed:    true,
			},
			"sae_password": {
				Type:        schema.TypeString,
				Description: "WPA3 SAE password to be used to authenticate WiFi users.",
				Computed:    true,
				Sensitive:   true,
			},
			"scan_botnet_connections": {
				Type:        schema.TypeString,
				Description: "Block or monitor connections to Botnet servers or disable Botnet scanning.",
				Computed:    true,
			},
			"schedule": {
				Type:        schema.TypeList,
				Description: "Firewall schedules for enabling this VAP on the FortiAP. This VAP will be enabled when at least one of the schedules is valid. Separate multiple schedule names with a space.",
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
			"secondary_wag_profile": {
				Type:        schema.TypeString,
				Description: "Secondary wireless access gateway profile name.",
				Computed:    true,
			},
			"security": {
				Type:        schema.TypeString,
				Description: "Security mode for the wireless interface (default = wpa2-only-personal).",
				Computed:    true,
			},
			"security_exempt_list": {
				Type:        schema.TypeString,
				Description: "Optional security exempt list for captive portal authentication.",
				Computed:    true,
			},
			"security_redirect_url": {
				Type:        schema.TypeString,
				Description: "Optional URL for redirecting users after they pass captive portal authentication.",
				Computed:    true,
			},
			"selected_usergroups": {
				Type:        schema.TypeList,
				Description: "Selective user groups that are permitted to authenticate.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "User group name.",
							Computed:    true,
						},
					},
				},
			},
			"split_tunneling": {
				Type:        schema.TypeString,
				Description: "Enable/disable split tunneling (default = disable).",
				Computed:    true,
			},
			"ssid": {
				Type:        schema.TypeString,
				Description: "IEEE 802.11 service set identifier (SSID) for the wireless interface. Users who wish to use the wireless network must configure their computers to access this SSID name.",
				Computed:    true,
			},
			"sticky_client_remove": {
				Type:        schema.TypeString,
				Description: "Enable/disable sticky client remove to maintain good signal level clients in SSID (default = disable).",
				Computed:    true,
			},
			"sticky_client_threshold_2g": {
				Type:        schema.TypeString,
				Description: "Minimum signal level/threshold in dBm required for the 2G client to be serviced by the AP (-95 to -20, default = -79).",
				Computed:    true,
			},
			"sticky_client_threshold_5g": {
				Type:        schema.TypeString,
				Description: "Minimum signal level/threshold in dBm required for the 5G client to be serviced by the AP (-95 to -20, default = -76).",
				Computed:    true,
			},
			"target_wake_time": {
				Type:        schema.TypeString,
				Description: "Enable/disable 802.11ax target wake time (default = enable).",
				Computed:    true,
			},
			"tkip_counter_measure": {
				Type:        schema.TypeString,
				Description: "Enable/disable TKIP counter measure.",
				Computed:    true,
			},
			"tunnel_echo_interval": {
				Type:        schema.TypeInt,
				Description: "The time interval to send echo to both primary and secondary tunnel peers (1 - 65535 sec, default = 300).",
				Computed:    true,
			},
			"tunnel_fallback_interval": {
				Type:        schema.TypeInt,
				Description: "The time interval for secondary tunnel to fall back to primary tunnel (0 - 65535 sec, default = 7200).",
				Computed:    true,
			},
			"usergroup": {
				Type:        schema.TypeList,
				Description: "Firewall user group to be used to authenticate WiFi users.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "User group name.",
							Computed:    true,
						},
					},
				},
			},
			"utm_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable UTM logging.",
				Computed:    true,
			},
			"utm_profile": {
				Type:        schema.TypeString,
				Description: "UTM profile name.",
				Computed:    true,
			},
			"utm_status": {
				Type:        schema.TypeString,
				Description: "Enable to add one or more security profiles (AV, IPS, etc.) to the VAP.",
				Computed:    true,
			},
			"vlan_auto": {
				Type:        schema.TypeString,
				Description: "Enable/disable automatic management of SSID VLAN interface.",
				Computed:    true,
			},
			"vlan_name": {
				Type:        schema.TypeList,
				Description: "Table for mapping VLAN name to VLAN ID.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "VLAN name.",
							Computed:    true,
						},
						"vlan_id": {
							Type:        schema.TypeInt,
							Description: "VLAN ID.",
							Computed:    true,
						},
					},
				},
			},
			"vlan_pool": {
				Type:        schema.TypeList,
				Description: "VLAN pool.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"wtp_group": {
							Type:        schema.TypeString,
							Description: "WTP group name.",
							Computed:    true,
						},
					},
				},
			},
			"vlan_pooling": {
				Type:        schema.TypeString,
				Description: "Enable/disable VLAN pooling, to allow grouping of multiple wireless controller VLANs into VLAN pools (default = disable). When set to wtp-group, VLAN pooling occurs with VLAN assignment by wtp-group.",
				Computed:    true,
			},
			"vlanid": {
				Type:        schema.TypeInt,
				Description: "Optional VLAN ID.",
				Computed:    true,
			},
			"voice_enterprise": {
				Type:        schema.TypeString,
				Description: "Enable/disable 802.11k and 802.11v assisted Voice-Enterprise roaming (default = disable).",
				Computed:    true,
			},
			"webfilter_profile": {
				Type:        schema.TypeString,
				Description: "WebFilter profile name.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWirelessControllerVapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("name")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadWirelessControllerVap(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerVap dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerVap(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
