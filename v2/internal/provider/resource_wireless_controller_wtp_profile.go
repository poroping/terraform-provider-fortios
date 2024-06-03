// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure WTP profiles or FortiAP profiles that define radio settings for manageable FortiAP platforms.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceWirelessControllerWtpProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure WTP profiles or FortiAP profiles that define radio settings for manageable FortiAP platforms.",

		CreateContext: resourceWirelessControllerWtpProfileCreate,
		ReadContext:   resourceWirelessControllerWtpProfileRead,
		UpdateContext: resourceWirelessControllerWtpProfileUpdate,
		DeleteContext: resourceWirelessControllerWtpProfileDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"allow_append": {
				Type:        schema.TypeBool,
				Description: "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"allowaccess": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Control management access to the managed WTP, FortiAP, or AP. Separate entries with a space.",
				Optional:         true,
				Computed:         true,
			},
			"ap_country": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"--", "AF", "AL", "DZ", "AS", "AO", "AR", "AM", "AU", "AT", "AZ", "BS", "BH", "BD", "BB", "BY", "BE", "BZ", "BJ", "BM", "BT", "BO", "BA", "BW", "BR", "BN", "BG", "BF", "KH", "CM", "KY", "CF", "TD", "CL", "CN", "CX", "CO", "CG", "CD", "CR", "HR", "CY", "CZ", "DK", "DJ", "DM", "DO", "EC", "EG", "SV", "ET", "EE", "GF", "PF", "FO", "FJ", "FI", "FR", "GA", "GE", "GM", "DE", "GH", "GI", "GR", "GL", "GD", "GP", "GU", "GT", "GY", "HT", "HN", "HK", "HU", "IS", "IN", "ID", "IQ", "IE", "IM", "IL", "IT", "CI", "JM", "JO", "KZ", "KE", "KR", "KW", "LA", "LV", "LB", "LS", "LR", "LY", "LI", "LT", "LU", "MO", "MK", "MG", "MW", "MY", "MV", "ML", "MT", "MH", "MQ", "MR", "MU", "YT", "MX", "FM", "MD", "MC", "MN", "MA", "MZ", "MM", "NA", "NP", "NL", "AN", "AW", "NZ", "NI", "NE", "NG", "NO", "MP", "OM", "PK", "PW", "PA", "PG", "PY", "PE", "PH", "PL", "PT", "PR", "QA", "RE", "RO", "RU", "RW", "BL", "KN", "LC", "MF", "PM", "VC", "SA", "SN", "RS", "ME", "SL", "SG", "SK", "SI", "SO", "ZA", "ES", "LK", "SR", "SZ", "SE", "CH", "TW", "TZ", "TH", "TG", "TT", "TN", "TR", "TM", "AE", "TC", "UG", "UA", "GB", "US", "PS", "UY", "UZ", "VU", "VE", "VN", "VI", "WF", "YE", "ZM", "ZW", "JP", "CA"}, false),

				Description: "Country in which this WTP, FortiAP, or AP will operate (default = NA, automatically use the country configured for the current VDOM).",
				Optional:    true,
				Computed:    true,
			},
			"ap_handoff": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable AP handoff of clients to other APs (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"apcfg_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "AP local configuration profile name.",
				Optional:    true,
				Computed:    true,
			},
			"ble_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Bluetooth Low Energy profile name.",
				Optional:    true,
				Computed:    true,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"console_login": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiAP console login access (default = enable).",
				Optional:    true,
				Computed:    true,
			},
			"control_message_offload": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Enable/disable CAPWAP control message data channel offload.",
				Optional:         true,
				Computed:         true,
			},
			"deny_mac_list": {
				Type:        schema.TypeList,
				Description: "List of MAC addresses that are denied access to this WTP, FortiAP, or AP.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"mac": {
							Type: schema.TypeString,

							Description: "A WiFi device with this MAC address is denied access to this WTP, FortiAP or AP.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dtls_in_kernel": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable data channel DTLS in kernel.",
				Optional:    true,
				Computed:    true,
			},
			"dtls_policy": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "WTP data channel DTLS policy (default = clear-text).",
				Optional:         true,
				Computed:         true,
			},
			"energy_efficient_ethernet": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of energy efficient Ethernet on WTP.",
				Optional:    true,
				Computed:    true,
			},
			"esl_ses_dongle": {
				Type:        schema.TypeList,
				Description: "ESL SES-imagotag dongle configuration.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"apc_addr_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"fqdn", "ip"}, false),

							Description: "ESL SES-imagotag APC address type (default = fqdn).",
							Optional:    true,
							Computed:    true,
						},
						"apc_fqdn": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "FQDN of ESL SES-imagotag Access Point Controller (APC).",
							Optional:    true,
							Computed:    true,
						},
						"apc_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "IP address of ESL SES-imagotag Access Point Controller (APC).",
							Optional:    true,
							Computed:    true,
						},
						"apc_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Port of ESL SES-imagotag Access Point Controller (APC).",
							Optional:    true,
							Computed:    true,
						},
						"coex_level": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none"}, false),

							Description: "ESL SES-imagotag dongle coexistence level (default = none).",
							Optional:    true,
							Computed:    true,
						},
						"compliance_level": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"compliance-level-2"}, false),

							Description: "Compliance levels for the ESL solution integration (default = compliance-level-2).",
							Optional:    true,
							Computed:    true,
						},
						"esl_channel": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"-1", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "127"}, false),

							Description: "ESL SES-imagotag dongle channel (default = 127).",
							Optional:    true,
							Computed:    true,
						},
						"output_power": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"a", "b", "c", "d", "e", "f", "g", "h"}, false),

							Description: "ESL SES-imagotag dongle output power (default = A).",
							Optional:    true,
							Computed:    true,
						},
						"scd_enable": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable ESL SES-imagotag Serial Communication Daemon (SCD) (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"tls_cert_verification": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable TLS certificate verification (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"tls_fqdn_verification": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable TLS certificate verification (default = disable).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ext_info_enable": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable station/VAP/radio extension information.",
				Optional:    true,
				Computed:    true,
			},
			"frequency_handoff": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable frequency handoff of clients to other channels (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"handoff_roaming": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable client load balancing during roaming to avoid roaming delay (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"handoff_rssi": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(20, 30),

				Description: "Minimum received signal strength indicator (RSSI) value for handoff (20 - 30, default = 25).",
				Optional:    true,
				Computed:    true,
			},
			"handoff_sta_thresh": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 60),

				Description: "Threshold value for AP handoff.",
				Optional:    true,
				Computed:    true,
			},
			"indoor_outdoor_deployment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"platform-determined", "outdoor", "indoor"}, false),

				Description: "Set to allow indoor/outdoor-only channels under regulatory rules (default = platform-determined).",
				Optional:    true,
				Computed:    true,
			},
			"ip_fragment_preventing": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Method(s) by which IP fragmentation is prevented for control and data packets through CAPWAP tunnel (default = tcp-mss-adjust).",
				Optional:         true,
				Computed:         true,
			},
			"lan": {
				Type:        schema.TypeList,
				Description: "WTP LAN port mapping.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_esl_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"offline", "nat-to-wan", "bridge-to-wan", "bridge-to-ssid"}, false),

							Description: "ESL port mode.",
							Optional:    true,
							Computed:    true,
						},
						"port_esl_ssid": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Bridge ESL port to SSID.",
							Optional:    true,
							Computed:    true,
						},
						"port_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"offline", "nat-to-wan", "bridge-to-wan", "bridge-to-ssid"}, false),

							Description: "LAN port mode.",
							Optional:    true,
							Computed:    true,
						},
						"port_ssid": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Bridge LAN port to SSID.",
							Optional:    true,
							Computed:    true,
						},
						"port1_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"offline", "nat-to-wan", "bridge-to-wan", "bridge-to-ssid"}, false),

							Description: "LAN port 1 mode.",
							Optional:    true,
							Computed:    true,
						},
						"port1_ssid": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Bridge LAN port 1 to SSID.",
							Optional:    true,
							Computed:    true,
						},
						"port2_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"offline", "nat-to-wan", "bridge-to-wan", "bridge-to-ssid"}, false),

							Description: "LAN port 2 mode.",
							Optional:    true,
							Computed:    true,
						},
						"port2_ssid": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Bridge LAN port 2 to SSID.",
							Optional:    true,
							Computed:    true,
						},
						"port3_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"offline", "nat-to-wan", "bridge-to-wan", "bridge-to-ssid"}, false),

							Description: "LAN port 3 mode.",
							Optional:    true,
							Computed:    true,
						},
						"port3_ssid": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Bridge LAN port 3 to SSID.",
							Optional:    true,
							Computed:    true,
						},
						"port4_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"offline", "nat-to-wan", "bridge-to-wan", "bridge-to-ssid"}, false),

							Description: "LAN port 4 mode.",
							Optional:    true,
							Computed:    true,
						},
						"port4_ssid": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Bridge LAN port 4 to SSID.",
							Optional:    true,
							Computed:    true,
						},
						"port5_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"offline", "nat-to-wan", "bridge-to-wan", "bridge-to-ssid"}, false),

							Description: "LAN port 5 mode.",
							Optional:    true,
							Computed:    true,
						},
						"port5_ssid": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Bridge LAN port 5 to SSID.",
							Optional:    true,
							Computed:    true,
						},
						"port6_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"offline", "nat-to-wan", "bridge-to-wan", "bridge-to-ssid"}, false),

							Description: "LAN port 6 mode.",
							Optional:    true,
							Computed:    true,
						},
						"port6_ssid": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Bridge LAN port 6 to SSID.",
							Optional:    true,
							Computed:    true,
						},
						"port7_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"offline", "nat-to-wan", "bridge-to-wan", "bridge-to-ssid"}, false),

							Description: "LAN port 7 mode.",
							Optional:    true,
							Computed:    true,
						},
						"port7_ssid": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Bridge LAN port 7 to SSID.",
							Optional:    true,
							Computed:    true,
						},
						"port8_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"offline", "nat-to-wan", "bridge-to-wan", "bridge-to-ssid"}, false),

							Description: "LAN port 8 mode.",
							Optional:    true,
							Computed:    true,
						},
						"port8_ssid": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Bridge LAN port 8 to SSID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"lbs": {
				Type:        schema.TypeList,
				Description: "Set various location based service (LBS) options.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aeroscout": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable AeroScout Real Time Location Service (RTLS) support (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"aeroscout_ap_mac": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bssid", "board-mac"}, false),

							Description: "Use BSSID or board MAC address as AP MAC address in AeroScout AP messages (default = bssid).",
							Optional:    true,
							Computed:    true,
						},
						"aeroscout_mmu_report": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable compounded AeroScout tag and MU report (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"aeroscout_mu": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable AeroScout Mobile Unit (MU) support (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"aeroscout_mu_factor": {
							Type: schema.TypeInt,

							Description: "AeroScout MU mode dilution factor (default = 20).",
							Optional:    true,
							Computed:    true,
						},
						"aeroscout_mu_timeout": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "AeroScout MU mode timeout (0 - 65535 sec, default = 5).",
							Optional:    true,
							Computed:    true,
						},
						"aeroscout_server_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "IP address of AeroScout server.",
							Optional:    true,
							Computed:    true,
						},
						"aeroscout_server_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1024, 65535),

							Description: "AeroScout server UDP listening port.",
							Optional:    true,
							Computed:    true,
						},
						"ekahau_blink_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable Ekahau blink mode (now known as AiRISTA Flow) to track and locate WiFi tags (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"ekahau_tag": {
							Type: schema.TypeString,

							Description: "WiFi frame MAC address or WiFi Tag.",
							Optional:    true,
							Computed:    true,
						},
						"erc_server_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "IP address of Ekahau RTLS Controller (ERC).",
							Optional:    true,
							Computed:    true,
						},
						"erc_server_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1024, 65535),

							Description: "Ekahau RTLS Controller (ERC) UDP listening port.",
							Optional:    true,
							Computed:    true,
						},
						"fortipresence": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"foreign", "both", "disable"}, false),

							Description: "Enable/disable FortiPresence to monitor the location and activity of WiFi clients even if they don't connect to this WiFi network (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"fortipresence_ble": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable FortiPresence finding and reporting BLE devices.",
							Optional:    true,
							Computed:    true,
						},
						"fortipresence_frequency": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 65535),

							Description: "FortiPresence report transmit frequency (5 - 65535 sec, default = 30).",
							Optional:    true,
							Computed:    true,
						},
						"fortipresence_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(300, 65535),

							Description: "UDP listening port of FortiPresence server (default = 3000).",
							Optional:    true,
							Computed:    true,
						},
						"fortipresence_project": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 16),

							Description: "FortiPresence project name (max. 16 characters, default = fortipresence).",
							Optional:    true,
							Computed:    true,
						},
						"fortipresence_rogue": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable FortiPresence finding and reporting rogue APs.",
							Optional:    true,
							Computed:    true,
						},
						"fortipresence_secret": {
							Type: schema.TypeString,

							Description: "FortiPresence secret password (max. 16 characters).",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"fortipresence_server": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "IP address of FortiPresence server.",
							Optional:    true,
							Computed:    true,
						},
						"fortipresence_server_addr_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ipv4", "fqdn"}, false),

							Description: "FortiPresence server address type (default = ipv4).",
							Optional:    true,
							Computed:    true,
						},
						"fortipresence_server_fqdn": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "FQDN of FortiPresence server.",
							Optional:    true,
							Computed:    true,
						},
						"fortipresence_unassoc": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable FortiPresence finding and reporting unassociated stations.",
							Optional:    true,
							Computed:    true,
						},
						"station_locate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable client station locating services for all clients, whether associated or not (default = disable).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"led_schedules": {
				Type:        schema.TypeList,
				Description: "Recurring firewall schedules for illuminating LEDs on the FortiAP. If led-state is enabled, LEDs will be visible when at least one of the schedules is valid. Separate multiple schedule names with a space.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Schedule name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"led_state": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of LEDs on WTP (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"lldp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Link Layer Discovery Protocol (LLDP) for the WTP, FortiAP, or AP (default = enable).",
				Optional:    true,
				Computed:    true,
			},
			"login_passwd": {
				Type: schema.TypeString,

				Description: "Set the managed WTP, FortiAP, or AP's administrator password.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"login_passwd_change": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"yes", "default", "no"}, false),

				Description: "Change or reset the administrator password of a managed WTP, FortiAP or AP (yes, default, or no, default = no).",
				Optional:    true,
				Computed:    true,
			},
			"max_clients": {
				Type: schema.TypeInt,

				Description: "Maximum number of stations (STAs) supported by the WTP (default = 0, meaning no client limitation).",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "WTP (or FortiAP or AP) profile name.",
				ForceNew:    true,
				Required:    true,
			},
			"platform": {
				Type:        schema.TypeList,
				Description: "WTP, FortiAP, or AP platform.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ddscan": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable use of one radio for dedicated full-band scanning to detect RF characterization and wireless threat management.",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"single-5G", "dual-5G"}, false),

							Description: "Configure operation mode of 5G radios (default = single-5G).",
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"AP-11N", "220B", "210B", "222B", "112B", "320B", "11C", "14C", "223B", "28C", "320C", "221C", "25D", "222C", "224D", "214B", "21D", "24D", "112D", "223C", "321C", "C220C", "C225C", "C23JD", "C24JE", "S321C", "S322C", "S323C", "S311C", "S313C", "S321CR", "S322CR", "S323CR", "S421E", "S422E", "S423E", "421E", "423E", "221E", "222E", "223E", "224E", "231E", "S221E", "S223E", "321E", "431F", "431FL", "432F", "432FR", "433F", "433FL", "231F", "231FL", "234F", "23JF", "831F", "231G", "233G", "234G", "431G", "432G", "433G", "U421E", "U422EV", "U423E", "U221EV", "U223EV", "U24JEV", "U321EV", "U323EV", "U431F", "U433F", "U231F", "U234F", "U432F", "U231G", "U441G"}, false),

							Description: "WTP, FortiAP or AP platform type. There are built-in WTP profiles for all supported FortiAP models. You can select a built-in profile and customize it or create a new profile.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"poe_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "8023af", "8023at", "power-adapter", "full", "high", "low"}, false),

				Description: "Set the WTP, FortiAP, or AP's PoE mode.",
				Optional:    true,
				Computed:    true,
			},
			"radio_1": {
				Type:        schema.TypeList,
				Description: "Configuration options for radio 1.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"80211d": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable 802.11d countryie(default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"airtime_fairness": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable airtime fairness (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"amsdu": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_handoff": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable AP handoff of clients to other APs (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_addr": {
							Type: schema.TypeString,

							Description: "MAC address to monitor.",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_bufsize": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 32),

							Description: "Sniffer buffer size (1 - 32 MB, default = 16).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_chan": {
							Type: schema.TypeInt,

							Description: "Channel on which to operate the sniffer (default = 6).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_ctl": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sniffer on WiFi control frame (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_data": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sniffer on WiFi data frame (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_mgmt_beacon": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sniffer on WiFi management Beacon frames (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_mgmt_other": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sniffer on WiFi management other frames  (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_mgmt_probe": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sniffer on WiFi management probe frames (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"arrp_profile": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Distributed Automatic Radio Resource Provisioning (DARRP) profile name to assign to the radio.",
							Optional:    true,
							Computed:    true,
						},
						"auto_power_high": {
							Type: schema.TypeInt,

							Description: "The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).",
							Optional:    true,
							Computed:    true,
						},
						"auto_power_level": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"auto_power_low": {
							Type: schema.TypeInt,

							Description: "The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).",
							Optional:    true,
							Computed:    true,
						},
						"auto_power_target": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 7),

							Description: "Target of automatic transmit power adjustment in dBm (-95 to -20, default = -70).",
							Optional:    true,
							Computed:    true,
						},
						"band": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"802.11a", "802.11b", "802.11g", "802.11n", "802.11n-5G", "802.11ac", "802.11ax-5G", "802.11ax", "802.11ac-2G", "802.11ax-6G", "802.11n,g-only", "802.11g-only", "802.11n-only", "802.11n-5G-only", "802.11ac,n-only", "802.11ac-only", "802.11ax,ac-only", "802.11ax,ac,n-only", "802.11ax-5G-only", "802.11ax,n-only", "802.11ax,n,g-only", "802.11ax-only"}, false),

							Description: "WiFi band that Radio 1 operates on.",
							Optional:    true,
							Computed:    true,
						},
						"band_5g_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"5g-full", "5g-high", "5g-low"}, false),

							Description: "WiFi 5G band type.",
							Optional:    true,
							Computed:    true,
						},
						"bandwidth_admission_control": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it.",
							Optional:    true,
							Computed:    true,
						},
						"bandwidth_capacity": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 600000),

							Description: "Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).",
							Optional:    true,
							Computed:    true,
						},
						"beacon_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Beacon interval. The time between beacon frames in milliseconds. Actual range of beacon interval depends on the AP platform type (default = 100).",
							Optional:    true,
							Computed:    true,
						},
						"bss_color": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),

							Description: "BSS color value for this 11ax radio (0 - 63, disable = 0, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"bss_color_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"auto", "static"}, false),

							Description: "BSS color mode for this 11ax radio (default = auto).",
							Optional:    true,
							Computed:    true,
						},
						"call_admission_control": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them.",
							Optional:    true,
							Computed:    true,
						},
						"call_capacity": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 60),

							Description: "Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).",
							Optional:    true,
							Computed:    true,
						},
						"channel": {
							Type:        schema.TypeList,
							Description: "Selected list of wireless radio channels.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"chan": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 3),

										Description: "Channel number.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"channel_bonding": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"160MHz", "80MHz", "40MHz", "20MHz"}, false),

							Description: "Channel bandwidth: 160,80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence.",
							Optional:    true,
							Computed:    true,
						},
						"channel_utilization": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable measuring channel utilization.",
							Optional:    true,
							Computed:    true,
						},
						"coexistence": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable allowing both HT20 and HT40 on the same radio (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"darrp": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"drma": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable dynamic radio mode assignment (DRMA) (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"drma_sensitivity": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"low", "medium", "high"}, false),

							Description: "Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low).",
							Optional:    true,
							Computed:    true,
						},
						"dtim": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),

							Description: "Delivery Traffic Indication Map (DTIM) period (1 - 255, default = 1). Set higher to save battery life of WiFi client in power-save mode.",
							Optional:    true,
							Computed:    true,
						},
						"frag_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(800, 2346),

							Description: "Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).",
							Optional:    true,
							Computed:    true,
						},
						"frequency_handoff": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable frequency handoff of clients to other channels (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"iperf_protocol": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"udp", "tcp"}, false),

							Description: "Iperf test protocol (default = \"UDP\").",
							Optional:    true,
							Computed:    true,
						},
						"iperf_server_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Iperf service port number.",
							Optional:    true,
							Computed:    true,
						},
						"max_clients": {
							Type: schema.TypeInt,

							Description: "Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.",
							Optional:    true,
							Computed:    true,
						},
						"max_distance": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 54000),

							Description: "Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disabled", "ap", "monitor", "sniffer", "sam"}, false),

							Description: "Mode of radio 1. Radio 1 can be disabled, configured as an access point, a rogue AP monitor, a sniffer, or a station.",
							Optional:    true,
							Computed:    true,
						},
						"optional_antenna": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "FANT-04ABGN-0606-O-N", "FANT-04ABGN-1414-P-N", "FANT-04ABGN-8065-P-N", "FANT-04ABGN-0606-O-R", "FANT-04ABGN-0606-P-R", "FANT-10ACAX-1213-D-N", "FANT-08ABGN-1213-D-R"}, false),

							Description: "Optional antenna used on FAP (default = none).",
							Optional:    true,
							Computed:    true,
						},
						"power_level": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),

							Description: "Radio EIRP power level as a percentage of the maximum EIRP power (0 - 100, default = 100).",
							Optional:    true,
							Computed:    true,
						},
						"power_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"dBm", "percentage"}, false),

							Description: "Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities.",
							Optional:    true,
							Computed:    true,
						},
						"power_value": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 33),

							Description: "Radio EIRP power in dBm (1 - 33, default = 27).",
							Optional:    true,
							Computed:    true,
						},
						"powersave_optimize": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Enable client power-saving features such as TIM, AC VO, and OBSS etc.",
							Optional:         true,
							Computed:         true,
						},
						"protection_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"rtscts", "ctsonly", "disable"}, false),

							Description: "Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable).",
							Optional:    true,
							Computed:    true,
						},
						"rts_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(256, 2346),

							Description: "Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).",
							Optional:    true,
							Computed:    true,
						},
						"sam_bssid": {
							Type: schema.TypeString,

							Description: "BSSID for WiFi network.",
							Optional:    true,
							Computed:    true,
						},
						"sam_captive_portal": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable Captive Portal Authentication (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"sam_cwp_failure_string": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),

							Description: "Failure identification on the page after an incorrect login.",
							Optional:    true,
							Computed:    true,
						},
						"sam_cwp_match_string": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),

							Description: "Identification string from the captive portal login form.",
							Optional:    true,
							Computed:    true,
						},
						"sam_cwp_password": {
							Type: schema.TypeString,

							Description: "Password for captive portal authentication.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"sam_cwp_success_string": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),

							Description: "Success identification on the page after a successful login.",
							Optional:    true,
							Computed:    true,
						},
						"sam_cwp_test_url": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Website the client is trying to access.",
							Optional:    true,
							Computed:    true,
						},
						"sam_cwp_username": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Username for captive portal authentication.",
							Optional:    true,
							Computed:    true,
						},
						"sam_password": {
							Type: schema.TypeString,

							Description: "Passphrase for WiFi network connection.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"sam_report_intv": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(60, 864000),

							Description: "SAM report interval (sec), 0 for a one-time report.",
							Optional:    true,
							Computed:    true,
						},
						"sam_security_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"open", "wpa-personal", "wpa-enterprise"}, false),

							Description: "Select WiFi network security type (default = \"wpa-personal\").",
							Optional:    true,
							Computed:    true,
						},
						"sam_server": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "SAM test server IP address or domain name.",
							Optional:    true,
							Computed:    true,
						},
						"sam_server_fqdn": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "SAM test server domain name.",
							Optional:    true,
							Computed:    true,
						},
						"sam_server_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "SAM test server IP address.",
							Optional:    true,
							Computed:    true,
						},
						"sam_server_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ip", "fqdn"}, false),

							Description: "Select SAM server type (default = \"IP\").",
							Optional:    true,
							Computed:    true,
						},
						"sam_ssid": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32),

							Description: "SSID for WiFi network.",
							Optional:    true,
							Computed:    true,
						},
						"sam_test": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ping", "iperf"}, false),

							Description: "Select SAM test type (default = \"PING\").",
							Optional:    true,
							Computed:    true,
						},
						"sam_username": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Username for WiFi network connection.",
							Optional:    true,
							Computed:    true,
						},
						"short_guard_interval": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns.",
							Optional:    true,
							Computed:    true,
						},
						"spectrum_analysis": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "scan-only", "disable"}, false),

							Description: "Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.",
							Optional:    true,
							Computed:    true,
						},
						"transmit_optimize": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default.",
							Optional:         true,
							Computed:         true,
						},
						"vap_all": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"tunnel", "bridge", "manual"}, false),

							Description: "Configure method for assigning SSIDs to this FortiAP (default = automatically assign tunnel SSIDs).",
							Optional:    true,
							Computed:    true,
						},
						"vaps": {
							Type:        schema.TypeList,
							Description: "Manually selected list of Virtual Access Points (VAPs).",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Virtual Access Point (VAP) name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"wids_profile": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Wireless Intrusion Detection System (WIDS) profile name to assign to the radio.",
							Optional:    true,
							Computed:    true,
						},
						"zero_wait_dfs": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable zero wait DFS on radio (default = enable).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"radio_2": {
				Type:        schema.TypeList,
				Description: "Configuration options for radio 2.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"80211d": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable 802.11d countryie(default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"airtime_fairness": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable airtime fairness (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"amsdu": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_handoff": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable AP handoff of clients to other APs (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_addr": {
							Type: schema.TypeString,

							Description: "MAC address to monitor.",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_bufsize": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 32),

							Description: "Sniffer buffer size (1 - 32 MB, default = 16).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_chan": {
							Type: schema.TypeInt,

							Description: "Channel on which to operate the sniffer (default = 6).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_ctl": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sniffer on WiFi control frame (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_data": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sniffer on WiFi data frame (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_mgmt_beacon": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sniffer on WiFi management Beacon frames (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_mgmt_other": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sniffer on WiFi management other frames  (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_mgmt_probe": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sniffer on WiFi management probe frames (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"arrp_profile": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Distributed Automatic Radio Resource Provisioning (DARRP) profile name to assign to the radio.",
							Optional:    true,
							Computed:    true,
						},
						"auto_power_high": {
							Type: schema.TypeInt,

							Description: "The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).",
							Optional:    true,
							Computed:    true,
						},
						"auto_power_level": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"auto_power_low": {
							Type: schema.TypeInt,

							Description: "The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).",
							Optional:    true,
							Computed:    true,
						},
						"auto_power_target": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 7),

							Description: "Target of automatic transmit power adjustment in dBm (-95 to -20, default = -70).",
							Optional:    true,
							Computed:    true,
						},
						"band": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"802.11a", "802.11b", "802.11g", "802.11n", "802.11n-5G", "802.11ac", "802.11ax-5G", "802.11ax", "802.11ac-2G", "802.11ax-6G", "802.11n,g-only", "802.11g-only", "802.11n-only", "802.11n-5G-only", "802.11ac,n-only", "802.11ac-only", "802.11ax,ac-only", "802.11ax,ac,n-only", "802.11ax-5G-only", "802.11ax,n-only", "802.11ax,n,g-only", "802.11ax-only"}, false),

							Description: "WiFi band that Radio 2 operates on.",
							Optional:    true,
							Computed:    true,
						},
						"band_5g_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"5g-full", "5g-high", "5g-low"}, false),

							Description: "WiFi 5G band type.",
							Optional:    true,
							Computed:    true,
						},
						"bandwidth_admission_control": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it.",
							Optional:    true,
							Computed:    true,
						},
						"bandwidth_capacity": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 600000),

							Description: "Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).",
							Optional:    true,
							Computed:    true,
						},
						"beacon_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Beacon interval. The time between beacon frames in milliseconds. Actual range of beacon interval depends on the AP platform type (default = 100).",
							Optional:    true,
							Computed:    true,
						},
						"bss_color": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),

							Description: "BSS color value for this 11ax radio (0 - 63, disable = 0, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"bss_color_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"auto", "static"}, false),

							Description: "BSS color mode for this 11ax radio (default = auto).",
							Optional:    true,
							Computed:    true,
						},
						"call_admission_control": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them.",
							Optional:    true,
							Computed:    true,
						},
						"call_capacity": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 60),

							Description: "Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).",
							Optional:    true,
							Computed:    true,
						},
						"channel": {
							Type:        schema.TypeList,
							Description: "Selected list of wireless radio channels.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"chan": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 3),

										Description: "Channel number.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"channel_bonding": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"160MHz", "80MHz", "40MHz", "20MHz"}, false),

							Description: "Channel bandwidth: 160,80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence.",
							Optional:    true,
							Computed:    true,
						},
						"channel_utilization": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable measuring channel utilization.",
							Optional:    true,
							Computed:    true,
						},
						"coexistence": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable allowing both HT20 and HT40 on the same radio (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"darrp": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"drma": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable dynamic radio mode assignment (DRMA) (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"drma_sensitivity": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"low", "medium", "high"}, false),

							Description: "Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low).",
							Optional:    true,
							Computed:    true,
						},
						"dtim": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),

							Description: "Delivery Traffic Indication Map (DTIM) period (1 - 255, default = 1). Set higher to save battery life of WiFi client in power-save mode.",
							Optional:    true,
							Computed:    true,
						},
						"frag_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(800, 2346),

							Description: "Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).",
							Optional:    true,
							Computed:    true,
						},
						"frequency_handoff": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable frequency handoff of clients to other channels (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"iperf_protocol": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"udp", "tcp"}, false),

							Description: "Iperf test protocol (default = \"UDP\").",
							Optional:    true,
							Computed:    true,
						},
						"iperf_server_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Iperf service port number.",
							Optional:    true,
							Computed:    true,
						},
						"max_clients": {
							Type: schema.TypeInt,

							Description: "Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.",
							Optional:    true,
							Computed:    true,
						},
						"max_distance": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 54000),

							Description: "Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disabled", "ap", "monitor", "sniffer", "sam"}, false),

							Description: "Mode of radio 2. Radio 2 can be disabled, configured as an access point, a rogue AP monitor, a sniffer, or a station.",
							Optional:    true,
							Computed:    true,
						},
						"optional_antenna": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "FANT-04ABGN-0606-O-N", "FANT-04ABGN-1414-P-N", "FANT-04ABGN-8065-P-N", "FANT-04ABGN-0606-O-R", "FANT-04ABGN-0606-P-R", "FANT-10ACAX-1213-D-N", "FANT-08ABGN-1213-D-R"}, false),

							Description: "Optional antenna used on FAP (default = none).",
							Optional:    true,
							Computed:    true,
						},
						"power_level": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),

							Description: "Radio EIRP power level as a percentage of the maximum EIRP power (0 - 100, default = 100).",
							Optional:    true,
							Computed:    true,
						},
						"power_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"dBm", "percentage"}, false),

							Description: "Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities.",
							Optional:    true,
							Computed:    true,
						},
						"power_value": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 33),

							Description: "Radio EIRP power in dBm (1 - 33, default = 27).",
							Optional:    true,
							Computed:    true,
						},
						"powersave_optimize": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Enable client power-saving features such as TIM, AC VO, and OBSS etc.",
							Optional:         true,
							Computed:         true,
						},
						"protection_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"rtscts", "ctsonly", "disable"}, false),

							Description: "Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable).",
							Optional:    true,
							Computed:    true,
						},
						"rts_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(256, 2346),

							Description: "Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).",
							Optional:    true,
							Computed:    true,
						},
						"sam_bssid": {
							Type: schema.TypeString,

							Description: "BSSID for WiFi network.",
							Optional:    true,
							Computed:    true,
						},
						"sam_captive_portal": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable Captive Portal Authentication (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"sam_cwp_failure_string": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),

							Description: "Failure identification on the page after an incorrect login.",
							Optional:    true,
							Computed:    true,
						},
						"sam_cwp_match_string": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),

							Description: "Identification string from the captive portal login form.",
							Optional:    true,
							Computed:    true,
						},
						"sam_cwp_password": {
							Type: schema.TypeString,

							Description: "Password for captive portal authentication.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"sam_cwp_success_string": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),

							Description: "Success identification on the page after a successful login.",
							Optional:    true,
							Computed:    true,
						},
						"sam_cwp_test_url": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Website the client is trying to access.",
							Optional:    true,
							Computed:    true,
						},
						"sam_cwp_username": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Username for captive portal authentication.",
							Optional:    true,
							Computed:    true,
						},
						"sam_password": {
							Type: schema.TypeString,

							Description: "Passphrase for WiFi network connection.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"sam_report_intv": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(60, 864000),

							Description: "SAM report interval (sec), 0 for a one-time report.",
							Optional:    true,
							Computed:    true,
						},
						"sam_security_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"open", "wpa-personal", "wpa-enterprise"}, false),

							Description: "Select WiFi network security type (default = \"wpa-personal\").",
							Optional:    true,
							Computed:    true,
						},
						"sam_server": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "SAM test server IP address or domain name.",
							Optional:    true,
							Computed:    true,
						},
						"sam_server_fqdn": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "SAM test server domain name.",
							Optional:    true,
							Computed:    true,
						},
						"sam_server_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "SAM test server IP address.",
							Optional:    true,
							Computed:    true,
						},
						"sam_server_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ip", "fqdn"}, false),

							Description: "Select SAM server type (default = \"IP\").",
							Optional:    true,
							Computed:    true,
						},
						"sam_ssid": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32),

							Description: "SSID for WiFi network.",
							Optional:    true,
							Computed:    true,
						},
						"sam_test": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ping", "iperf"}, false),

							Description: "Select SAM test type (default = \"PING\").",
							Optional:    true,
							Computed:    true,
						},
						"sam_username": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Username for WiFi network connection.",
							Optional:    true,
							Computed:    true,
						},
						"short_guard_interval": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns.",
							Optional:    true,
							Computed:    true,
						},
						"spectrum_analysis": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "scan-only", "disable"}, false),

							Description: "Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.",
							Optional:    true,
							Computed:    true,
						},
						"transmit_optimize": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default.",
							Optional:         true,
							Computed:         true,
						},
						"vap_all": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"tunnel", "bridge", "manual"}, false),

							Description: "Configure method for assigning SSIDs to this FortiAP (default = automatically assign tunnel SSIDs).",
							Optional:    true,
							Computed:    true,
						},
						"vaps": {
							Type:        schema.TypeList,
							Description: "Manually selected list of Virtual Access Points (VAPs).",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Virtual Access Point (VAP) name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"wids_profile": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Wireless Intrusion Detection System (WIDS) profile name to assign to the radio.",
							Optional:    true,
							Computed:    true,
						},
						"zero_wait_dfs": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable zero wait DFS on radio (default = enable).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"radio_3": {
				Type:        schema.TypeList,
				Description: "Configuration options for radio 3.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"80211d": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable 802.11d countryie(default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"airtime_fairness": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable airtime fairness (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"amsdu": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_handoff": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable AP handoff of clients to other APs (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_addr": {
							Type: schema.TypeString,

							Description: "MAC address to monitor.",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_bufsize": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 32),

							Description: "Sniffer buffer size (1 - 32 MB, default = 16).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_chan": {
							Type: schema.TypeInt,

							Description: "Channel on which to operate the sniffer (default = 6).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_ctl": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sniffer on WiFi control frame (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_data": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sniffer on WiFi data frame (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_mgmt_beacon": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sniffer on WiFi management Beacon frames (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_mgmt_other": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sniffer on WiFi management other frames  (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_mgmt_probe": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sniffer on WiFi management probe frames (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"arrp_profile": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Distributed Automatic Radio Resource Provisioning (DARRP) profile name to assign to the radio.",
							Optional:    true,
							Computed:    true,
						},
						"auto_power_high": {
							Type: schema.TypeInt,

							Description: "The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).",
							Optional:    true,
							Computed:    true,
						},
						"auto_power_level": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"auto_power_low": {
							Type: schema.TypeInt,

							Description: "The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).",
							Optional:    true,
							Computed:    true,
						},
						"auto_power_target": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 7),

							Description: "Target of automatic transmit power adjustment in dBm (-95 to -20, default = -70).",
							Optional:    true,
							Computed:    true,
						},
						"band": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"802.11a", "802.11b", "802.11g", "802.11n", "802.11n-5G", "802.11ac", "802.11ax-5G", "802.11ax", "802.11ac-2G", "802.11ax-6G", "802.11n,g-only", "802.11g-only", "802.11n-only", "802.11n-5G-only", "802.11ac,n-only", "802.11ac-only", "802.11ax,ac-only", "802.11ax,ac,n-only", "802.11ax-5G-only", "802.11ax,n-only", "802.11ax,n,g-only", "802.11ax-only"}, false),

							Description: "WiFi band that Radio 3 operates on.",
							Optional:    true,
							Computed:    true,
						},
						"band_5g_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"5g-full", "5g-high", "5g-low"}, false),

							Description: "WiFi 5G band type.",
							Optional:    true,
							Computed:    true,
						},
						"bandwidth_admission_control": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it.",
							Optional:    true,
							Computed:    true,
						},
						"bandwidth_capacity": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 600000),

							Description: "Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).",
							Optional:    true,
							Computed:    true,
						},
						"beacon_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Beacon interval. The time between beacon frames in milliseconds. Actual range of beacon interval depends on the AP platform type (default = 100).",
							Optional:    true,
							Computed:    true,
						},
						"bss_color": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),

							Description: "BSS color value for this 11ax radio (0 - 63, disable = 0, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"bss_color_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"auto", "static"}, false),

							Description: "BSS color mode for this 11ax radio (default = auto).",
							Optional:    true,
							Computed:    true,
						},
						"call_admission_control": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them.",
							Optional:    true,
							Computed:    true,
						},
						"call_capacity": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 60),

							Description: "Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).",
							Optional:    true,
							Computed:    true,
						},
						"channel": {
							Type:        schema.TypeList,
							Description: "Selected list of wireless radio channels.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"chan": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 3),

										Description: "Channel number.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"channel_bonding": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"160MHz", "80MHz", "40MHz", "20MHz"}, false),

							Description: "Channel bandwidth: 160,80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence.",
							Optional:    true,
							Computed:    true,
						},
						"channel_utilization": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable measuring channel utilization.",
							Optional:    true,
							Computed:    true,
						},
						"coexistence": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable allowing both HT20 and HT40 on the same radio (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"darrp": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"drma": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable dynamic radio mode assignment (DRMA) (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"drma_sensitivity": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"low", "medium", "high"}, false),

							Description: "Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low).",
							Optional:    true,
							Computed:    true,
						},
						"dtim": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),

							Description: "Delivery Traffic Indication Map (DTIM) period (1 - 255, default = 1). Set higher to save battery life of WiFi client in power-save mode.",
							Optional:    true,
							Computed:    true,
						},
						"frag_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(800, 2346),

							Description: "Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).",
							Optional:    true,
							Computed:    true,
						},
						"frequency_handoff": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable frequency handoff of clients to other channels (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"iperf_protocol": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"udp", "tcp"}, false),

							Description: "Iperf test protocol (default = \"UDP\").",
							Optional:    true,
							Computed:    true,
						},
						"iperf_server_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Iperf service port number.",
							Optional:    true,
							Computed:    true,
						},
						"max_clients": {
							Type: schema.TypeInt,

							Description: "Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.",
							Optional:    true,
							Computed:    true,
						},
						"max_distance": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 54000),

							Description: "Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disabled", "ap", "monitor", "sniffer", "sam"}, false),

							Description: "Mode of radio 3. Radio 3 can be disabled, configured as an access point, a rogue AP monitor, a sniffer, or a station.",
							Optional:    true,
							Computed:    true,
						},
						"optional_antenna": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "FANT-04ABGN-0606-O-N", "FANT-04ABGN-1414-P-N", "FANT-04ABGN-8065-P-N", "FANT-04ABGN-0606-O-R", "FANT-04ABGN-0606-P-R", "FANT-10ACAX-1213-D-N", "FANT-08ABGN-1213-D-R"}, false),

							Description: "Optional antenna used on FAP (default = none).",
							Optional:    true,
							Computed:    true,
						},
						"power_level": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),

							Description: "Radio EIRP power level as a percentage of the maximum EIRP power (0 - 100, default = 100).",
							Optional:    true,
							Computed:    true,
						},
						"power_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"dBm", "percentage"}, false),

							Description: "Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities.",
							Optional:    true,
							Computed:    true,
						},
						"power_value": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 33),

							Description: "Radio EIRP power in dBm (1 - 33, default = 27).",
							Optional:    true,
							Computed:    true,
						},
						"powersave_optimize": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Enable client power-saving features such as TIM, AC VO, and OBSS etc.",
							Optional:         true,
							Computed:         true,
						},
						"protection_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"rtscts", "ctsonly", "disable"}, false),

							Description: "Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable).",
							Optional:    true,
							Computed:    true,
						},
						"rts_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(256, 2346),

							Description: "Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).",
							Optional:    true,
							Computed:    true,
						},
						"sam_bssid": {
							Type: schema.TypeString,

							Description: "BSSID for WiFi network.",
							Optional:    true,
							Computed:    true,
						},
						"sam_captive_portal": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable Captive Portal Authentication (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"sam_cwp_failure_string": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),

							Description: "Failure identification on the page after an incorrect login.",
							Optional:    true,
							Computed:    true,
						},
						"sam_cwp_match_string": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),

							Description: "Identification string from the captive portal login form.",
							Optional:    true,
							Computed:    true,
						},
						"sam_cwp_password": {
							Type: schema.TypeString,

							Description: "Password for captive portal authentication.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"sam_cwp_success_string": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),

							Description: "Success identification on the page after a successful login.",
							Optional:    true,
							Computed:    true,
						},
						"sam_cwp_test_url": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Website the client is trying to access.",
							Optional:    true,
							Computed:    true,
						},
						"sam_cwp_username": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Username for captive portal authentication.",
							Optional:    true,
							Computed:    true,
						},
						"sam_password": {
							Type: schema.TypeString,

							Description: "Passphrase for WiFi network connection.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"sam_report_intv": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(60, 864000),

							Description: "SAM report interval (sec), 0 for a one-time report.",
							Optional:    true,
							Computed:    true,
						},
						"sam_security_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"open", "wpa-personal", "wpa-enterprise"}, false),

							Description: "Select WiFi network security type (default = \"wpa-personal\").",
							Optional:    true,
							Computed:    true,
						},
						"sam_server": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "SAM test server IP address or domain name.",
							Optional:    true,
							Computed:    true,
						},
						"sam_server_fqdn": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "SAM test server domain name.",
							Optional:    true,
							Computed:    true,
						},
						"sam_server_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "SAM test server IP address.",
							Optional:    true,
							Computed:    true,
						},
						"sam_server_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ip", "fqdn"}, false),

							Description: "Select SAM server type (default = \"IP\").",
							Optional:    true,
							Computed:    true,
						},
						"sam_ssid": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32),

							Description: "SSID for WiFi network.",
							Optional:    true,
							Computed:    true,
						},
						"sam_test": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ping", "iperf"}, false),

							Description: "Select SAM test type (default = \"PING\").",
							Optional:    true,
							Computed:    true,
						},
						"sam_username": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Username for WiFi network connection.",
							Optional:    true,
							Computed:    true,
						},
						"short_guard_interval": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns.",
							Optional:    true,
							Computed:    true,
						},
						"spectrum_analysis": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "scan-only", "disable"}, false),

							Description: "Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.",
							Optional:    true,
							Computed:    true,
						},
						"transmit_optimize": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default.",
							Optional:         true,
							Computed:         true,
						},
						"vap_all": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"tunnel", "bridge", "manual"}, false),

							Description: "Configure method for assigning SSIDs to this FortiAP (default = automatically assign tunnel SSIDs).",
							Optional:    true,
							Computed:    true,
						},
						"vaps": {
							Type:        schema.TypeList,
							Description: "Manually selected list of Virtual Access Points (VAPs).",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Virtual Access Point (VAP) name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"wids_profile": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Wireless Intrusion Detection System (WIDS) profile name to assign to the radio.",
							Optional:    true,
							Computed:    true,
						},
						"zero_wait_dfs": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable zero wait DFS on radio (default = enable).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"radio_4": {
				Type:        schema.TypeList,
				Description: "Configuration options for radio 4.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"80211d": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable 802.11d countryie(default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"airtime_fairness": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable airtime fairness (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"amsdu": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_handoff": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable AP handoff of clients to other APs (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_addr": {
							Type: schema.TypeString,

							Description: "MAC address to monitor.",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_bufsize": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 32),

							Description: "Sniffer buffer size (1 - 32 MB, default = 16).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_chan": {
							Type: schema.TypeInt,

							Description: "Channel on which to operate the sniffer (default = 6).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_ctl": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sniffer on WiFi control frame (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_data": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sniffer on WiFi data frame (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_mgmt_beacon": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sniffer on WiFi management Beacon frames (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_mgmt_other": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sniffer on WiFi management other frames  (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"ap_sniffer_mgmt_probe": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sniffer on WiFi management probe frames (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"arrp_profile": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Distributed Automatic Radio Resource Provisioning (DARRP) profile name to assign to the radio.",
							Optional:    true,
							Computed:    true,
						},
						"auto_power_high": {
							Type: schema.TypeInt,

							Description: "The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).",
							Optional:    true,
							Computed:    true,
						},
						"auto_power_level": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"auto_power_low": {
							Type: schema.TypeInt,

							Description: "The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).",
							Optional:    true,
							Computed:    true,
						},
						"auto_power_target": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 7),

							Description: "Target of automatic transmit power adjustment in dBm (-95 to -20, default = -70).",
							Optional:    true,
							Computed:    true,
						},
						"band": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"802.11a", "802.11b", "802.11g", "802.11n", "802.11n-5G", "802.11ac", "802.11ax-5G", "802.11ax", "802.11ac-2G", "802.11ax-6G", "802.11n,g-only", "802.11g-only", "802.11n-only", "802.11n-5G-only", "802.11ac,n-only", "802.11ac-only", "802.11ax,ac-only", "802.11ax,ac,n-only", "802.11ax-5G-only", "802.11ax,n-only", "802.11ax,n,g-only", "802.11ax-only"}, false),

							Description: "WiFi band that Radio 3 operates on.",
							Optional:    true,
							Computed:    true,
						},
						"band_5g_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"5g-full", "5g-high", "5g-low"}, false),

							Description: "WiFi 5G band type.",
							Optional:    true,
							Computed:    true,
						},
						"bandwidth_admission_control": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it.",
							Optional:    true,
							Computed:    true,
						},
						"bandwidth_capacity": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 600000),

							Description: "Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).",
							Optional:    true,
							Computed:    true,
						},
						"beacon_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Beacon interval. The time between beacon frames in milliseconds. Actual range of beacon interval depends on the AP platform type (default = 100).",
							Optional:    true,
							Computed:    true,
						},
						"bss_color": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),

							Description: "BSS color value for this 11ax radio (0 - 63, disable = 0, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"bss_color_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"auto", "static"}, false),

							Description: "BSS color mode for this 11ax radio (default = auto).",
							Optional:    true,
							Computed:    true,
						},
						"call_admission_control": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them.",
							Optional:    true,
							Computed:    true,
						},
						"call_capacity": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 60),

							Description: "Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).",
							Optional:    true,
							Computed:    true,
						},
						"channel": {
							Type:        schema.TypeList,
							Description: "Selected list of wireless radio channels.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"chan": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 3),

										Description: "Channel number.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"channel_bonding": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"160MHz", "80MHz", "40MHz", "20MHz"}, false),

							Description: "Channel bandwidth: 160,80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence.",
							Optional:    true,
							Computed:    true,
						},
						"channel_utilization": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable measuring channel utilization.",
							Optional:    true,
							Computed:    true,
						},
						"coexistence": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable allowing both HT20 and HT40 on the same radio (default = enable).",
							Optional:    true,
							Computed:    true,
						},
						"darrp": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"drma": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable dynamic radio mode assignment (DRMA) (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"drma_sensitivity": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"low", "medium", "high"}, false),

							Description: "Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low).",
							Optional:    true,
							Computed:    true,
						},
						"dtim": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),

							Description: "Delivery Traffic Indication Map (DTIM) period (1 - 255, default = 1). Set higher to save battery life of WiFi client in power-save mode.",
							Optional:    true,
							Computed:    true,
						},
						"frag_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(800, 2346),

							Description: "Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).",
							Optional:    true,
							Computed:    true,
						},
						"frequency_handoff": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable frequency handoff of clients to other channels (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"iperf_protocol": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"udp", "tcp"}, false),

							Description: "Iperf test protocol (default = \"UDP\").",
							Optional:    true,
							Computed:    true,
						},
						"iperf_server_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Iperf service port number.",
							Optional:    true,
							Computed:    true,
						},
						"max_clients": {
							Type: schema.TypeInt,

							Description: "Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.",
							Optional:    true,
							Computed:    true,
						},
						"max_distance": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 54000),

							Description: "Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disabled", "ap", "monitor", "sniffer", "sam"}, false),

							Description: "Mode of radio 3. Radio 3 can be disabled, configured as an access point, a rogue AP monitor, a sniffer, or a station.",
							Optional:    true,
							Computed:    true,
						},
						"optional_antenna": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "FANT-04ABGN-0606-O-N", "FANT-04ABGN-1414-P-N", "FANT-04ABGN-8065-P-N", "FANT-04ABGN-0606-O-R", "FANT-04ABGN-0606-P-R", "FANT-10ACAX-1213-D-N", "FANT-08ABGN-1213-D-R"}, false),

							Description: "Optional antenna used on FAP (default = none).",
							Optional:    true,
							Computed:    true,
						},
						"power_level": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),

							Description: "Radio EIRP power level as a percentage of the maximum EIRP power (0 - 100, default = 100).",
							Optional:    true,
							Computed:    true,
						},
						"power_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"dBm", "percentage"}, false),

							Description: "Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities.",
							Optional:    true,
							Computed:    true,
						},
						"power_value": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 33),

							Description: "Radio EIRP power in dBm (1 - 33, default = 27).",
							Optional:    true,
							Computed:    true,
						},
						"powersave_optimize": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Enable client power-saving features such as TIM, AC VO, and OBSS etc.",
							Optional:         true,
							Computed:         true,
						},
						"protection_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"rtscts", "ctsonly", "disable"}, false),

							Description: "Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable).",
							Optional:    true,
							Computed:    true,
						},
						"rts_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(256, 2346),

							Description: "Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).",
							Optional:    true,
							Computed:    true,
						},
						"sam_bssid": {
							Type: schema.TypeString,

							Description: "BSSID for WiFi network.",
							Optional:    true,
							Computed:    true,
						},
						"sam_captive_portal": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable Captive Portal Authentication (default = disable).",
							Optional:    true,
							Computed:    true,
						},
						"sam_cwp_failure_string": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),

							Description: "Failure identification on the page after an incorrect login.",
							Optional:    true,
							Computed:    true,
						},
						"sam_cwp_match_string": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),

							Description: "Identification string from the captive portal login form.",
							Optional:    true,
							Computed:    true,
						},
						"sam_cwp_password": {
							Type: schema.TypeString,

							Description: "Password for captive portal authentication.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"sam_cwp_success_string": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),

							Description: "Success identification on the page after a successful login.",
							Optional:    true,
							Computed:    true,
						},
						"sam_cwp_test_url": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Website the client is trying to access.",
							Optional:    true,
							Computed:    true,
						},
						"sam_cwp_username": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Username for captive portal authentication.",
							Optional:    true,
							Computed:    true,
						},
						"sam_password": {
							Type: schema.TypeString,

							Description: "Passphrase for WiFi network connection.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"sam_report_intv": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(60, 864000),

							Description: "SAM report interval (sec), 0 for a one-time report.",
							Optional:    true,
							Computed:    true,
						},
						"sam_security_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"open", "wpa-personal", "wpa-enterprise"}, false),

							Description: "Select WiFi network security type (default = \"wpa-personal\").",
							Optional:    true,
							Computed:    true,
						},
						"sam_server": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "SAM test server IP address or domain name.",
							Optional:    true,
							Computed:    true,
						},
						"sam_server_fqdn": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "SAM test server domain name.",
							Optional:    true,
							Computed:    true,
						},
						"sam_server_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "SAM test server IP address.",
							Optional:    true,
							Computed:    true,
						},
						"sam_server_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ip", "fqdn"}, false),

							Description: "Select SAM server type (default = \"IP\").",
							Optional:    true,
							Computed:    true,
						},
						"sam_ssid": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32),

							Description: "SSID for WiFi network.",
							Optional:    true,
							Computed:    true,
						},
						"sam_test": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ping", "iperf"}, false),

							Description: "Select SAM test type (default = \"PING\").",
							Optional:    true,
							Computed:    true,
						},
						"sam_username": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Username for WiFi network connection.",
							Optional:    true,
							Computed:    true,
						},
						"short_guard_interval": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns.",
							Optional:    true,
							Computed:    true,
						},
						"spectrum_analysis": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "scan-only", "disable"}, false),

							Description: "Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.",
							Optional:    true,
							Computed:    true,
						},
						"transmit_optimize": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default.",
							Optional:         true,
							Computed:         true,
						},
						"vap_all": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"tunnel", "bridge", "manual"}, false),

							Description: "Configure method for assigning SSIDs to this FortiAP (default = automatically assign tunnel SSIDs).",
							Optional:    true,
							Computed:    true,
						},
						"vaps": {
							Type:        schema.TypeList,
							Description: "Manually selected list of Virtual Access Points (VAPs).",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Virtual Access Point (VAP) name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"wids_profile": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Wireless Intrusion Detection System (WIDS) profile name to assign to the radio.",
							Optional:    true,
							Computed:    true,
						},
						"zero_wait_dfs": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable zero wait DFS on radio (default = enable).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"split_tunneling_acl": {
				Type:        schema.TypeList,
				Description: "Split tunneling ACL filter list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dest_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validators.FortiValidateIPv4Classnet,

							Description: "Destination IP and mask for the split-tunneling subnet.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"split_tunneling_acl_local_ap_subnet": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable automatically adding local subnetwork of FortiAP to split-tunneling ACL (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"split_tunneling_acl_path": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"tunnel", "local"}, false),

				Description: "Split tunneling ACL path is local/tunnel.",
				Optional:    true,
				Computed:    true,
			},
			"syslog_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "System log server configuration profile name.",
				Optional:    true,
				Computed:    true,
			},
			"tun_mtu_downlink": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(576, 1500),

				Description: "The MTU of downlink CAPWAP tunnel (576 - 1500 bytes or 0; 0 means the local MTU of FortiAP; default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"tun_mtu_uplink": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(576, 1500),

				Description: "The maximum transmission unit (MTU) of uplink CAPWAP tunnel (576 - 1500 bytes or 0; 0 means the local MTU of FortiAP; default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"wan_port_auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "802.1x"}, false),

				Description: "Set WAN port authentication mode (default = none).",
				Optional:    true,
				Computed:    true,
			},
			"wan_port_auth_methods": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"all", "EAP-FAST", "EAP-TLS", "EAP-PEAP"}, false),

				Description: "WAN port 802.1x supplicant EAP methods (default = all).",
				Optional:    true,
				Computed:    true,
			},
			"wan_port_auth_password": {
				Type: schema.TypeString,

				Description: "Set WAN port 802.1x supplicant password.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"wan_port_auth_usrname": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Set WAN port 802.1x supplicant user name.",
				Optional:    true,
				Computed:    true,
			},
			"wan_port_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"wan-lan", "wan-only"}, false),

				Description: "Enable/disable using a WAN port as a LAN port.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWirelessControllerWtpProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	var diags diag.Diagnostics
	var err error
	// c.Retries = 1
	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	allow_append := false
	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}
	urlparams.AllowAppend = &allow_append

	mkey := ""
	key := "name"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating WirelessControllerWtpProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerWtpProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerWtpProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerWtpProfile")
	}

	return resourceWirelessControllerWtpProfileRead(ctx, d, meta)
}

func resourceWirelessControllerWtpProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerWtpProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerWtpProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerWtpProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerWtpProfile")
	}

	return resourceWirelessControllerWtpProfileRead(ctx, d, meta)
}

func resourceWirelessControllerWtpProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerWtpProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerWtpProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerWtpProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	ptp := true
	urlparams.PlainTextPassword = &ptp

	o, err := c.Cmdb.ReadWirelessControllerWtpProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerWtpProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
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

func flattenWirelessControllerWtpProfileDenyMacList(d *schema.ResourceData, v *[]models.WirelessControllerWtpProfileDenyMacList, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Mac; tmp != nil {
				v["mac"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenWirelessControllerWtpProfileEslSesDongle(d *schema.ResourceData, v *models.WirelessControllerWtpProfileEslSesDongle, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.WirelessControllerWtpProfileEslSesDongle{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ApcAddrType; tmp != nil {
				v["apc_addr_type"] = *tmp
			}

			if tmp := cfg.ApcFqdn; tmp != nil {
				v["apc_fqdn"] = *tmp
			}

			if tmp := cfg.ApcIp; tmp != nil {
				v["apc_ip"] = *tmp
			}

			if tmp := cfg.ApcPort; tmp != nil {
				v["apc_port"] = *tmp
			}

			if tmp := cfg.CoexLevel; tmp != nil {
				v["coex_level"] = *tmp
			}

			if tmp := cfg.ComplianceLevel; tmp != nil {
				v["compliance_level"] = *tmp
			}

			if tmp := cfg.EslChannel; tmp != nil {
				v["esl_channel"] = *tmp
			}

			if tmp := cfg.OutputPower; tmp != nil {
				v["output_power"] = *tmp
			}

			if tmp := cfg.ScdEnable; tmp != nil {
				v["scd_enable"] = *tmp
			}

			if tmp := cfg.TlsCertVerification; tmp != nil {
				v["tls_cert_verification"] = *tmp
			}

			if tmp := cfg.TlsFqdnVerification; tmp != nil {
				v["tls_fqdn_verification"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWirelessControllerWtpProfileLan(d *schema.ResourceData, v *models.WirelessControllerWtpProfileLan, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.WirelessControllerWtpProfileLan{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.PortEslMode; tmp != nil {
				v["port_esl_mode"] = *tmp
			}

			if tmp := cfg.PortEslSsid; tmp != nil {
				v["port_esl_ssid"] = *tmp
			}

			if tmp := cfg.PortMode; tmp != nil {
				v["port_mode"] = *tmp
			}

			if tmp := cfg.PortSsid; tmp != nil {
				v["port_ssid"] = *tmp
			}

			if tmp := cfg.Port1Mode; tmp != nil {
				v["port1_mode"] = *tmp
			}

			if tmp := cfg.Port1Ssid; tmp != nil {
				v["port1_ssid"] = *tmp
			}

			if tmp := cfg.Port2Mode; tmp != nil {
				v["port2_mode"] = *tmp
			}

			if tmp := cfg.Port2Ssid; tmp != nil {
				v["port2_ssid"] = *tmp
			}

			if tmp := cfg.Port3Mode; tmp != nil {
				v["port3_mode"] = *tmp
			}

			if tmp := cfg.Port3Ssid; tmp != nil {
				v["port3_ssid"] = *tmp
			}

			if tmp := cfg.Port4Mode; tmp != nil {
				v["port4_mode"] = *tmp
			}

			if tmp := cfg.Port4Ssid; tmp != nil {
				v["port4_ssid"] = *tmp
			}

			if tmp := cfg.Port5Mode; tmp != nil {
				v["port5_mode"] = *tmp
			}

			if tmp := cfg.Port5Ssid; tmp != nil {
				v["port5_ssid"] = *tmp
			}

			if tmp := cfg.Port6Mode; tmp != nil {
				v["port6_mode"] = *tmp
			}

			if tmp := cfg.Port6Ssid; tmp != nil {
				v["port6_ssid"] = *tmp
			}

			if tmp := cfg.Port7Mode; tmp != nil {
				v["port7_mode"] = *tmp
			}

			if tmp := cfg.Port7Ssid; tmp != nil {
				v["port7_ssid"] = *tmp
			}

			if tmp := cfg.Port8Mode; tmp != nil {
				v["port8_mode"] = *tmp
			}

			if tmp := cfg.Port8Ssid; tmp != nil {
				v["port8_ssid"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWirelessControllerWtpProfileLbs(d *schema.ResourceData, v *models.WirelessControllerWtpProfileLbs, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.WirelessControllerWtpProfileLbs{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Aeroscout; tmp != nil {
				v["aeroscout"] = *tmp
			}

			if tmp := cfg.AeroscoutApMac; tmp != nil {
				v["aeroscout_ap_mac"] = *tmp
			}

			if tmp := cfg.AeroscoutMmuReport; tmp != nil {
				v["aeroscout_mmu_report"] = *tmp
			}

			if tmp := cfg.AeroscoutMu; tmp != nil {
				v["aeroscout_mu"] = *tmp
			}

			if tmp := cfg.AeroscoutMuFactor; tmp != nil {
				v["aeroscout_mu_factor"] = *tmp
			}

			if tmp := cfg.AeroscoutMuTimeout; tmp != nil {
				v["aeroscout_mu_timeout"] = *tmp
			}

			if tmp := cfg.AeroscoutServerIp; tmp != nil {
				v["aeroscout_server_ip"] = *tmp
			}

			if tmp := cfg.AeroscoutServerPort; tmp != nil {
				v["aeroscout_server_port"] = *tmp
			}

			if tmp := cfg.EkahauBlinkMode; tmp != nil {
				v["ekahau_blink_mode"] = *tmp
			}

			if tmp := cfg.EkahauTag; tmp != nil {
				v["ekahau_tag"] = *tmp
			}

			if tmp := cfg.ErcServerIp; tmp != nil {
				v["erc_server_ip"] = *tmp
			}

			if tmp := cfg.ErcServerPort; tmp != nil {
				v["erc_server_port"] = *tmp
			}

			if tmp := cfg.Fortipresence; tmp != nil {
				v["fortipresence"] = *tmp
			}

			if tmp := cfg.FortipresenceBle; tmp != nil {
				v["fortipresence_ble"] = *tmp
			}

			if tmp := cfg.FortipresenceFrequency; tmp != nil {
				v["fortipresence_frequency"] = *tmp
			}

			if tmp := cfg.FortipresencePort; tmp != nil {
				v["fortipresence_port"] = *tmp
			}

			if tmp := cfg.FortipresenceProject; tmp != nil {
				v["fortipresence_project"] = *tmp
			}

			if tmp := cfg.FortipresenceRogue; tmp != nil {
				v["fortipresence_rogue"] = *tmp
			}

			if tmp := cfg.FortipresenceSecret; tmp != nil {
				v["fortipresence_secret"] = *tmp
			}

			if tmp := cfg.FortipresenceServer; tmp != nil {
				v["fortipresence_server"] = *tmp
			}

			if tmp := cfg.FortipresenceServerAddrType; tmp != nil {
				v["fortipresence_server_addr_type"] = *tmp
			}

			if tmp := cfg.FortipresenceServerFqdn; tmp != nil {
				v["fortipresence_server_fqdn"] = *tmp
			}

			if tmp := cfg.FortipresenceUnassoc; tmp != nil {
				v["fortipresence_unassoc"] = *tmp
			}

			if tmp := cfg.StationLocate; tmp != nil {
				v["station_locate"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWirelessControllerWtpProfileLedSchedules(d *schema.ResourceData, v *[]models.WirelessControllerWtpProfileLedSchedules, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenWirelessControllerWtpProfilePlatform(d *schema.ResourceData, v *models.WirelessControllerWtpProfilePlatform, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.WirelessControllerWtpProfilePlatform{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Ddscan; tmp != nil {
				v["ddscan"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWirelessControllerWtpProfileRadio1(d *schema.ResourceData, v *models.WirelessControllerWtpProfileRadio1, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.WirelessControllerWtpProfileRadio1{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.The80211d; tmp != nil {
				v["80211d"] = *tmp
			}

			if tmp := cfg.AirtimeFairness; tmp != nil {
				v["airtime_fairness"] = *tmp
			}

			if tmp := cfg.Amsdu; tmp != nil {
				v["amsdu"] = *tmp
			}

			if tmp := cfg.ApHandoff; tmp != nil {
				v["ap_handoff"] = *tmp
			}

			if tmp := cfg.ApSnifferAddr; tmp != nil {
				v["ap_sniffer_addr"] = *tmp
			}

			if tmp := cfg.ApSnifferBufsize; tmp != nil {
				v["ap_sniffer_bufsize"] = *tmp
			}

			if tmp := cfg.ApSnifferChan; tmp != nil {
				v["ap_sniffer_chan"] = *tmp
			}

			if tmp := cfg.ApSnifferCtl; tmp != nil {
				v["ap_sniffer_ctl"] = *tmp
			}

			if tmp := cfg.ApSnifferData; tmp != nil {
				v["ap_sniffer_data"] = *tmp
			}

			if tmp := cfg.ApSnifferMgmtBeacon; tmp != nil {
				v["ap_sniffer_mgmt_beacon"] = *tmp
			}

			if tmp := cfg.ApSnifferMgmtOther; tmp != nil {
				v["ap_sniffer_mgmt_other"] = *tmp
			}

			if tmp := cfg.ApSnifferMgmtProbe; tmp != nil {
				v["ap_sniffer_mgmt_probe"] = *tmp
			}

			if tmp := cfg.ArrpProfile; tmp != nil {
				v["arrp_profile"] = *tmp
			}

			if tmp := cfg.AutoPowerHigh; tmp != nil {
				v["auto_power_high"] = *tmp
			}

			if tmp := cfg.AutoPowerLevel; tmp != nil {
				v["auto_power_level"] = *tmp
			}

			if tmp := cfg.AutoPowerLow; tmp != nil {
				v["auto_power_low"] = *tmp
			}

			if tmp := cfg.AutoPowerTarget; tmp != nil {
				v["auto_power_target"] = *tmp
			}

			if tmp := cfg.Band; tmp != nil {
				v["band"] = *tmp
			}

			if tmp := cfg.Band5gType; tmp != nil {
				v["band_5g_type"] = *tmp
			}

			if tmp := cfg.BandwidthAdmissionControl; tmp != nil {
				v["bandwidth_admission_control"] = *tmp
			}

			if tmp := cfg.BandwidthCapacity; tmp != nil {
				v["bandwidth_capacity"] = *tmp
			}

			if tmp := cfg.BeaconInterval; tmp != nil {
				v["beacon_interval"] = *tmp
			}

			if tmp := cfg.BssColor; tmp != nil {
				v["bss_color"] = *tmp
			}

			if tmp := cfg.BssColorMode; tmp != nil {
				v["bss_color_mode"] = *tmp
			}

			if tmp := cfg.CallAdmissionControl; tmp != nil {
				v["call_admission_control"] = *tmp
			}

			if tmp := cfg.CallCapacity; tmp != nil {
				v["call_capacity"] = *tmp
			}

			if tmp := cfg.Channel; tmp != nil {
				v["channel"] = flattenWirelessControllerWtpProfileRadio1Channel(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "channel"), sort)
			}

			if tmp := cfg.ChannelBonding; tmp != nil {
				v["channel_bonding"] = *tmp
			}

			if tmp := cfg.ChannelUtilization; tmp != nil {
				v["channel_utilization"] = *tmp
			}

			if tmp := cfg.Coexistence; tmp != nil {
				v["coexistence"] = *tmp
			}

			if tmp := cfg.Darrp; tmp != nil {
				v["darrp"] = *tmp
			}

			if tmp := cfg.Drma; tmp != nil {
				v["drma"] = *tmp
			}

			if tmp := cfg.DrmaSensitivity; tmp != nil {
				v["drma_sensitivity"] = *tmp
			}

			if tmp := cfg.Dtim; tmp != nil {
				v["dtim"] = *tmp
			}

			if tmp := cfg.FragThreshold; tmp != nil {
				v["frag_threshold"] = *tmp
			}

			if tmp := cfg.FrequencyHandoff; tmp != nil {
				v["frequency_handoff"] = *tmp
			}

			if tmp := cfg.IperfProtocol; tmp != nil {
				v["iperf_protocol"] = *tmp
			}

			if tmp := cfg.IperfServerPort; tmp != nil {
				v["iperf_server_port"] = *tmp
			}

			if tmp := cfg.MaxClients; tmp != nil {
				v["max_clients"] = *tmp
			}

			if tmp := cfg.MaxDistance; tmp != nil {
				v["max_distance"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			if tmp := cfg.OptionalAntenna; tmp != nil {
				v["optional_antenna"] = *tmp
			}

			if tmp := cfg.PowerLevel; tmp != nil {
				v["power_level"] = *tmp
			}

			if tmp := cfg.PowerMode; tmp != nil {
				v["power_mode"] = *tmp
			}

			if tmp := cfg.PowerValue; tmp != nil {
				v["power_value"] = *tmp
			}

			if tmp := cfg.PowersaveOptimize; tmp != nil {
				v["powersave_optimize"] = *tmp
			}

			if tmp := cfg.ProtectionMode; tmp != nil {
				v["protection_mode"] = *tmp
			}

			if tmp := cfg.RtsThreshold; tmp != nil {
				v["rts_threshold"] = *tmp
			}

			if tmp := cfg.SamBssid; tmp != nil {
				v["sam_bssid"] = *tmp
			}

			if tmp := cfg.SamCaptivePortal; tmp != nil {
				v["sam_captive_portal"] = *tmp
			}

			if tmp := cfg.SamCwpFailureString; tmp != nil {
				v["sam_cwp_failure_string"] = *tmp
			}

			if tmp := cfg.SamCwpMatchString; tmp != nil {
				v["sam_cwp_match_string"] = *tmp
			}

			if tmp := cfg.SamCwpPassword; tmp != nil {
				v["sam_cwp_password"] = *tmp
			}

			if tmp := cfg.SamCwpSuccessString; tmp != nil {
				v["sam_cwp_success_string"] = *tmp
			}

			if tmp := cfg.SamCwpTestUrl; tmp != nil {
				v["sam_cwp_test_url"] = *tmp
			}

			if tmp := cfg.SamCwpUsername; tmp != nil {
				v["sam_cwp_username"] = *tmp
			}

			if tmp := cfg.SamPassword; tmp != nil {
				v["sam_password"] = *tmp
			}

			if tmp := cfg.SamReportIntv; tmp != nil {
				v["sam_report_intv"] = *tmp
			}

			if tmp := cfg.SamSecurityType; tmp != nil {
				v["sam_security_type"] = *tmp
			}

			if tmp := cfg.SamServer; tmp != nil {
				v["sam_server"] = *tmp
			}

			if tmp := cfg.SamServerFqdn; tmp != nil {
				v["sam_server_fqdn"] = *tmp
			}

			if tmp := cfg.SamServerIp; tmp != nil {
				v["sam_server_ip"] = *tmp
			}

			if tmp := cfg.SamServerType; tmp != nil {
				v["sam_server_type"] = *tmp
			}

			if tmp := cfg.SamSsid; tmp != nil {
				v["sam_ssid"] = *tmp
			}

			if tmp := cfg.SamTest; tmp != nil {
				v["sam_test"] = *tmp
			}

			if tmp := cfg.SamUsername; tmp != nil {
				v["sam_username"] = *tmp
			}

			if tmp := cfg.ShortGuardInterval; tmp != nil {
				v["short_guard_interval"] = *tmp
			}

			if tmp := cfg.SpectrumAnalysis; tmp != nil {
				v["spectrum_analysis"] = *tmp
			}

			if tmp := cfg.TransmitOptimize; tmp != nil {
				v["transmit_optimize"] = *tmp
			}

			if tmp := cfg.VapAll; tmp != nil {
				v["vap_all"] = *tmp
			}

			if tmp := cfg.Vaps; tmp != nil {
				v["vaps"] = flattenWirelessControllerWtpProfileRadio1Vaps(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "vaps"), sort)
			}

			if tmp := cfg.WidsProfile; tmp != nil {
				v["wids_profile"] = *tmp
			}

			if tmp := cfg.ZeroWaitDfs; tmp != nil {
				v["zero_wait_dfs"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWirelessControllerWtpProfileRadio1Channel(d *schema.ResourceData, v *[]models.WirelessControllerWtpProfileRadio1Channel, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Chan; tmp != nil {
				v["chan"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "chan")
	}

	return flat
}

func flattenWirelessControllerWtpProfileRadio1Vaps(d *schema.ResourceData, v *[]models.WirelessControllerWtpProfileRadio1Vaps, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenWirelessControllerWtpProfileRadio2(d *schema.ResourceData, v *models.WirelessControllerWtpProfileRadio2, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.WirelessControllerWtpProfileRadio2{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.The80211d; tmp != nil {
				v["80211d"] = *tmp
			}

			if tmp := cfg.AirtimeFairness; tmp != nil {
				v["airtime_fairness"] = *tmp
			}

			if tmp := cfg.Amsdu; tmp != nil {
				v["amsdu"] = *tmp
			}

			if tmp := cfg.ApHandoff; tmp != nil {
				v["ap_handoff"] = *tmp
			}

			if tmp := cfg.ApSnifferAddr; tmp != nil {
				v["ap_sniffer_addr"] = *tmp
			}

			if tmp := cfg.ApSnifferBufsize; tmp != nil {
				v["ap_sniffer_bufsize"] = *tmp
			}

			if tmp := cfg.ApSnifferChan; tmp != nil {
				v["ap_sniffer_chan"] = *tmp
			}

			if tmp := cfg.ApSnifferCtl; tmp != nil {
				v["ap_sniffer_ctl"] = *tmp
			}

			if tmp := cfg.ApSnifferData; tmp != nil {
				v["ap_sniffer_data"] = *tmp
			}

			if tmp := cfg.ApSnifferMgmtBeacon; tmp != nil {
				v["ap_sniffer_mgmt_beacon"] = *tmp
			}

			if tmp := cfg.ApSnifferMgmtOther; tmp != nil {
				v["ap_sniffer_mgmt_other"] = *tmp
			}

			if tmp := cfg.ApSnifferMgmtProbe; tmp != nil {
				v["ap_sniffer_mgmt_probe"] = *tmp
			}

			if tmp := cfg.ArrpProfile; tmp != nil {
				v["arrp_profile"] = *tmp
			}

			if tmp := cfg.AutoPowerHigh; tmp != nil {
				v["auto_power_high"] = *tmp
			}

			if tmp := cfg.AutoPowerLevel; tmp != nil {
				v["auto_power_level"] = *tmp
			}

			if tmp := cfg.AutoPowerLow; tmp != nil {
				v["auto_power_low"] = *tmp
			}

			if tmp := cfg.AutoPowerTarget; tmp != nil {
				v["auto_power_target"] = *tmp
			}

			if tmp := cfg.Band; tmp != nil {
				v["band"] = *tmp
			}

			if tmp := cfg.Band5gType; tmp != nil {
				v["band_5g_type"] = *tmp
			}

			if tmp := cfg.BandwidthAdmissionControl; tmp != nil {
				v["bandwidth_admission_control"] = *tmp
			}

			if tmp := cfg.BandwidthCapacity; tmp != nil {
				v["bandwidth_capacity"] = *tmp
			}

			if tmp := cfg.BeaconInterval; tmp != nil {
				v["beacon_interval"] = *tmp
			}

			if tmp := cfg.BssColor; tmp != nil {
				v["bss_color"] = *tmp
			}

			if tmp := cfg.BssColorMode; tmp != nil {
				v["bss_color_mode"] = *tmp
			}

			if tmp := cfg.CallAdmissionControl; tmp != nil {
				v["call_admission_control"] = *tmp
			}

			if tmp := cfg.CallCapacity; tmp != nil {
				v["call_capacity"] = *tmp
			}

			if tmp := cfg.Channel; tmp != nil {
				v["channel"] = flattenWirelessControllerWtpProfileRadio2Channel(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "channel"), sort)
			}

			if tmp := cfg.ChannelBonding; tmp != nil {
				v["channel_bonding"] = *tmp
			}

			if tmp := cfg.ChannelUtilization; tmp != nil {
				v["channel_utilization"] = *tmp
			}

			if tmp := cfg.Coexistence; tmp != nil {
				v["coexistence"] = *tmp
			}

			if tmp := cfg.Darrp; tmp != nil {
				v["darrp"] = *tmp
			}

			if tmp := cfg.Drma; tmp != nil {
				v["drma"] = *tmp
			}

			if tmp := cfg.DrmaSensitivity; tmp != nil {
				v["drma_sensitivity"] = *tmp
			}

			if tmp := cfg.Dtim; tmp != nil {
				v["dtim"] = *tmp
			}

			if tmp := cfg.FragThreshold; tmp != nil {
				v["frag_threshold"] = *tmp
			}

			if tmp := cfg.FrequencyHandoff; tmp != nil {
				v["frequency_handoff"] = *tmp
			}

			if tmp := cfg.IperfProtocol; tmp != nil {
				v["iperf_protocol"] = *tmp
			}

			if tmp := cfg.IperfServerPort; tmp != nil {
				v["iperf_server_port"] = *tmp
			}

			if tmp := cfg.MaxClients; tmp != nil {
				v["max_clients"] = *tmp
			}

			if tmp := cfg.MaxDistance; tmp != nil {
				v["max_distance"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			if tmp := cfg.OptionalAntenna; tmp != nil {
				v["optional_antenna"] = *tmp
			}

			if tmp := cfg.PowerLevel; tmp != nil {
				v["power_level"] = *tmp
			}

			if tmp := cfg.PowerMode; tmp != nil {
				v["power_mode"] = *tmp
			}

			if tmp := cfg.PowerValue; tmp != nil {
				v["power_value"] = *tmp
			}

			if tmp := cfg.PowersaveOptimize; tmp != nil {
				v["powersave_optimize"] = *tmp
			}

			if tmp := cfg.ProtectionMode; tmp != nil {
				v["protection_mode"] = *tmp
			}

			if tmp := cfg.RtsThreshold; tmp != nil {
				v["rts_threshold"] = *tmp
			}

			if tmp := cfg.SamBssid; tmp != nil {
				v["sam_bssid"] = *tmp
			}

			if tmp := cfg.SamCaptivePortal; tmp != nil {
				v["sam_captive_portal"] = *tmp
			}

			if tmp := cfg.SamCwpFailureString; tmp != nil {
				v["sam_cwp_failure_string"] = *tmp
			}

			if tmp := cfg.SamCwpMatchString; tmp != nil {
				v["sam_cwp_match_string"] = *tmp
			}

			if tmp := cfg.SamCwpPassword; tmp != nil {
				v["sam_cwp_password"] = *tmp
			}

			if tmp := cfg.SamCwpSuccessString; tmp != nil {
				v["sam_cwp_success_string"] = *tmp
			}

			if tmp := cfg.SamCwpTestUrl; tmp != nil {
				v["sam_cwp_test_url"] = *tmp
			}

			if tmp := cfg.SamCwpUsername; tmp != nil {
				v["sam_cwp_username"] = *tmp
			}

			if tmp := cfg.SamPassword; tmp != nil {
				v["sam_password"] = *tmp
			}

			if tmp := cfg.SamReportIntv; tmp != nil {
				v["sam_report_intv"] = *tmp
			}

			if tmp := cfg.SamSecurityType; tmp != nil {
				v["sam_security_type"] = *tmp
			}

			if tmp := cfg.SamServer; tmp != nil {
				v["sam_server"] = *tmp
			}

			if tmp := cfg.SamServerFqdn; tmp != nil {
				v["sam_server_fqdn"] = *tmp
			}

			if tmp := cfg.SamServerIp; tmp != nil {
				v["sam_server_ip"] = *tmp
			}

			if tmp := cfg.SamServerType; tmp != nil {
				v["sam_server_type"] = *tmp
			}

			if tmp := cfg.SamSsid; tmp != nil {
				v["sam_ssid"] = *tmp
			}

			if tmp := cfg.SamTest; tmp != nil {
				v["sam_test"] = *tmp
			}

			if tmp := cfg.SamUsername; tmp != nil {
				v["sam_username"] = *tmp
			}

			if tmp := cfg.ShortGuardInterval; tmp != nil {
				v["short_guard_interval"] = *tmp
			}

			if tmp := cfg.SpectrumAnalysis; tmp != nil {
				v["spectrum_analysis"] = *tmp
			}

			if tmp := cfg.TransmitOptimize; tmp != nil {
				v["transmit_optimize"] = *tmp
			}

			if tmp := cfg.VapAll; tmp != nil {
				v["vap_all"] = *tmp
			}

			if tmp := cfg.Vaps; tmp != nil {
				v["vaps"] = flattenWirelessControllerWtpProfileRadio2Vaps(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "vaps"), sort)
			}

			if tmp := cfg.WidsProfile; tmp != nil {
				v["wids_profile"] = *tmp
			}

			if tmp := cfg.ZeroWaitDfs; tmp != nil {
				v["zero_wait_dfs"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWirelessControllerWtpProfileRadio2Channel(d *schema.ResourceData, v *[]models.WirelessControllerWtpProfileRadio2Channel, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Chan; tmp != nil {
				v["chan"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "chan")
	}

	return flat
}

func flattenWirelessControllerWtpProfileRadio2Vaps(d *schema.ResourceData, v *[]models.WirelessControllerWtpProfileRadio2Vaps, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenWirelessControllerWtpProfileRadio3(d *schema.ResourceData, v *models.WirelessControllerWtpProfileRadio3, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.WirelessControllerWtpProfileRadio3{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.The80211d; tmp != nil {
				v["80211d"] = *tmp
			}

			if tmp := cfg.AirtimeFairness; tmp != nil {
				v["airtime_fairness"] = *tmp
			}

			if tmp := cfg.Amsdu; tmp != nil {
				v["amsdu"] = *tmp
			}

			if tmp := cfg.ApHandoff; tmp != nil {
				v["ap_handoff"] = *tmp
			}

			if tmp := cfg.ApSnifferAddr; tmp != nil {
				v["ap_sniffer_addr"] = *tmp
			}

			if tmp := cfg.ApSnifferBufsize; tmp != nil {
				v["ap_sniffer_bufsize"] = *tmp
			}

			if tmp := cfg.ApSnifferChan; tmp != nil {
				v["ap_sniffer_chan"] = *tmp
			}

			if tmp := cfg.ApSnifferCtl; tmp != nil {
				v["ap_sniffer_ctl"] = *tmp
			}

			if tmp := cfg.ApSnifferData; tmp != nil {
				v["ap_sniffer_data"] = *tmp
			}

			if tmp := cfg.ApSnifferMgmtBeacon; tmp != nil {
				v["ap_sniffer_mgmt_beacon"] = *tmp
			}

			if tmp := cfg.ApSnifferMgmtOther; tmp != nil {
				v["ap_sniffer_mgmt_other"] = *tmp
			}

			if tmp := cfg.ApSnifferMgmtProbe; tmp != nil {
				v["ap_sniffer_mgmt_probe"] = *tmp
			}

			if tmp := cfg.ArrpProfile; tmp != nil {
				v["arrp_profile"] = *tmp
			}

			if tmp := cfg.AutoPowerHigh; tmp != nil {
				v["auto_power_high"] = *tmp
			}

			if tmp := cfg.AutoPowerLevel; tmp != nil {
				v["auto_power_level"] = *tmp
			}

			if tmp := cfg.AutoPowerLow; tmp != nil {
				v["auto_power_low"] = *tmp
			}

			if tmp := cfg.AutoPowerTarget; tmp != nil {
				v["auto_power_target"] = *tmp
			}

			if tmp := cfg.Band; tmp != nil {
				v["band"] = *tmp
			}

			if tmp := cfg.Band5gType; tmp != nil {
				v["band_5g_type"] = *tmp
			}

			if tmp := cfg.BandwidthAdmissionControl; tmp != nil {
				v["bandwidth_admission_control"] = *tmp
			}

			if tmp := cfg.BandwidthCapacity; tmp != nil {
				v["bandwidth_capacity"] = *tmp
			}

			if tmp := cfg.BeaconInterval; tmp != nil {
				v["beacon_interval"] = *tmp
			}

			if tmp := cfg.BssColor; tmp != nil {
				v["bss_color"] = *tmp
			}

			if tmp := cfg.BssColorMode; tmp != nil {
				v["bss_color_mode"] = *tmp
			}

			if tmp := cfg.CallAdmissionControl; tmp != nil {
				v["call_admission_control"] = *tmp
			}

			if tmp := cfg.CallCapacity; tmp != nil {
				v["call_capacity"] = *tmp
			}

			if tmp := cfg.Channel; tmp != nil {
				v["channel"] = flattenWirelessControllerWtpProfileRadio3Channel(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "channel"), sort)
			}

			if tmp := cfg.ChannelBonding; tmp != nil {
				v["channel_bonding"] = *tmp
			}

			if tmp := cfg.ChannelUtilization; tmp != nil {
				v["channel_utilization"] = *tmp
			}

			if tmp := cfg.Coexistence; tmp != nil {
				v["coexistence"] = *tmp
			}

			if tmp := cfg.Darrp; tmp != nil {
				v["darrp"] = *tmp
			}

			if tmp := cfg.Drma; tmp != nil {
				v["drma"] = *tmp
			}

			if tmp := cfg.DrmaSensitivity; tmp != nil {
				v["drma_sensitivity"] = *tmp
			}

			if tmp := cfg.Dtim; tmp != nil {
				v["dtim"] = *tmp
			}

			if tmp := cfg.FragThreshold; tmp != nil {
				v["frag_threshold"] = *tmp
			}

			if tmp := cfg.FrequencyHandoff; tmp != nil {
				v["frequency_handoff"] = *tmp
			}

			if tmp := cfg.IperfProtocol; tmp != nil {
				v["iperf_protocol"] = *tmp
			}

			if tmp := cfg.IperfServerPort; tmp != nil {
				v["iperf_server_port"] = *tmp
			}

			if tmp := cfg.MaxClients; tmp != nil {
				v["max_clients"] = *tmp
			}

			if tmp := cfg.MaxDistance; tmp != nil {
				v["max_distance"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			if tmp := cfg.OptionalAntenna; tmp != nil {
				v["optional_antenna"] = *tmp
			}

			if tmp := cfg.PowerLevel; tmp != nil {
				v["power_level"] = *tmp
			}

			if tmp := cfg.PowerMode; tmp != nil {
				v["power_mode"] = *tmp
			}

			if tmp := cfg.PowerValue; tmp != nil {
				v["power_value"] = *tmp
			}

			if tmp := cfg.PowersaveOptimize; tmp != nil {
				v["powersave_optimize"] = *tmp
			}

			if tmp := cfg.ProtectionMode; tmp != nil {
				v["protection_mode"] = *tmp
			}

			if tmp := cfg.RtsThreshold; tmp != nil {
				v["rts_threshold"] = *tmp
			}

			if tmp := cfg.SamBssid; tmp != nil {
				v["sam_bssid"] = *tmp
			}

			if tmp := cfg.SamCaptivePortal; tmp != nil {
				v["sam_captive_portal"] = *tmp
			}

			if tmp := cfg.SamCwpFailureString; tmp != nil {
				v["sam_cwp_failure_string"] = *tmp
			}

			if tmp := cfg.SamCwpMatchString; tmp != nil {
				v["sam_cwp_match_string"] = *tmp
			}

			if tmp := cfg.SamCwpPassword; tmp != nil {
				v["sam_cwp_password"] = *tmp
			}

			if tmp := cfg.SamCwpSuccessString; tmp != nil {
				v["sam_cwp_success_string"] = *tmp
			}

			if tmp := cfg.SamCwpTestUrl; tmp != nil {
				v["sam_cwp_test_url"] = *tmp
			}

			if tmp := cfg.SamCwpUsername; tmp != nil {
				v["sam_cwp_username"] = *tmp
			}

			if tmp := cfg.SamPassword; tmp != nil {
				v["sam_password"] = *tmp
			}

			if tmp := cfg.SamReportIntv; tmp != nil {
				v["sam_report_intv"] = *tmp
			}

			if tmp := cfg.SamSecurityType; tmp != nil {
				v["sam_security_type"] = *tmp
			}

			if tmp := cfg.SamServer; tmp != nil {
				v["sam_server"] = *tmp
			}

			if tmp := cfg.SamServerFqdn; tmp != nil {
				v["sam_server_fqdn"] = *tmp
			}

			if tmp := cfg.SamServerIp; tmp != nil {
				v["sam_server_ip"] = *tmp
			}

			if tmp := cfg.SamServerType; tmp != nil {
				v["sam_server_type"] = *tmp
			}

			if tmp := cfg.SamSsid; tmp != nil {
				v["sam_ssid"] = *tmp
			}

			if tmp := cfg.SamTest; tmp != nil {
				v["sam_test"] = *tmp
			}

			if tmp := cfg.SamUsername; tmp != nil {
				v["sam_username"] = *tmp
			}

			if tmp := cfg.ShortGuardInterval; tmp != nil {
				v["short_guard_interval"] = *tmp
			}

			if tmp := cfg.SpectrumAnalysis; tmp != nil {
				v["spectrum_analysis"] = *tmp
			}

			if tmp := cfg.TransmitOptimize; tmp != nil {
				v["transmit_optimize"] = *tmp
			}

			if tmp := cfg.VapAll; tmp != nil {
				v["vap_all"] = *tmp
			}

			if tmp := cfg.Vaps; tmp != nil {
				v["vaps"] = flattenWirelessControllerWtpProfileRadio3Vaps(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "vaps"), sort)
			}

			if tmp := cfg.WidsProfile; tmp != nil {
				v["wids_profile"] = *tmp
			}

			if tmp := cfg.ZeroWaitDfs; tmp != nil {
				v["zero_wait_dfs"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWirelessControllerWtpProfileRadio3Channel(d *schema.ResourceData, v *[]models.WirelessControllerWtpProfileRadio3Channel, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Chan; tmp != nil {
				v["chan"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "chan")
	}

	return flat
}

func flattenWirelessControllerWtpProfileRadio3Vaps(d *schema.ResourceData, v *[]models.WirelessControllerWtpProfileRadio3Vaps, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenWirelessControllerWtpProfileRadio4(d *schema.ResourceData, v *models.WirelessControllerWtpProfileRadio4, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.WirelessControllerWtpProfileRadio4{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.The80211d; tmp != nil {
				v["80211d"] = *tmp
			}

			if tmp := cfg.AirtimeFairness; tmp != nil {
				v["airtime_fairness"] = *tmp
			}

			if tmp := cfg.Amsdu; tmp != nil {
				v["amsdu"] = *tmp
			}

			if tmp := cfg.ApHandoff; tmp != nil {
				v["ap_handoff"] = *tmp
			}

			if tmp := cfg.ApSnifferAddr; tmp != nil {
				v["ap_sniffer_addr"] = *tmp
			}

			if tmp := cfg.ApSnifferBufsize; tmp != nil {
				v["ap_sniffer_bufsize"] = *tmp
			}

			if tmp := cfg.ApSnifferChan; tmp != nil {
				v["ap_sniffer_chan"] = *tmp
			}

			if tmp := cfg.ApSnifferCtl; tmp != nil {
				v["ap_sniffer_ctl"] = *tmp
			}

			if tmp := cfg.ApSnifferData; tmp != nil {
				v["ap_sniffer_data"] = *tmp
			}

			if tmp := cfg.ApSnifferMgmtBeacon; tmp != nil {
				v["ap_sniffer_mgmt_beacon"] = *tmp
			}

			if tmp := cfg.ApSnifferMgmtOther; tmp != nil {
				v["ap_sniffer_mgmt_other"] = *tmp
			}

			if tmp := cfg.ApSnifferMgmtProbe; tmp != nil {
				v["ap_sniffer_mgmt_probe"] = *tmp
			}

			if tmp := cfg.ArrpProfile; tmp != nil {
				v["arrp_profile"] = *tmp
			}

			if tmp := cfg.AutoPowerHigh; tmp != nil {
				v["auto_power_high"] = *tmp
			}

			if tmp := cfg.AutoPowerLevel; tmp != nil {
				v["auto_power_level"] = *tmp
			}

			if tmp := cfg.AutoPowerLow; tmp != nil {
				v["auto_power_low"] = *tmp
			}

			if tmp := cfg.AutoPowerTarget; tmp != nil {
				v["auto_power_target"] = *tmp
			}

			if tmp := cfg.Band; tmp != nil {
				v["band"] = *tmp
			}

			if tmp := cfg.Band5gType; tmp != nil {
				v["band_5g_type"] = *tmp
			}

			if tmp := cfg.BandwidthAdmissionControl; tmp != nil {
				v["bandwidth_admission_control"] = *tmp
			}

			if tmp := cfg.BandwidthCapacity; tmp != nil {
				v["bandwidth_capacity"] = *tmp
			}

			if tmp := cfg.BeaconInterval; tmp != nil {
				v["beacon_interval"] = *tmp
			}

			if tmp := cfg.BssColor; tmp != nil {
				v["bss_color"] = *tmp
			}

			if tmp := cfg.BssColorMode; tmp != nil {
				v["bss_color_mode"] = *tmp
			}

			if tmp := cfg.CallAdmissionControl; tmp != nil {
				v["call_admission_control"] = *tmp
			}

			if tmp := cfg.CallCapacity; tmp != nil {
				v["call_capacity"] = *tmp
			}

			if tmp := cfg.Channel; tmp != nil {
				v["channel"] = flattenWirelessControllerWtpProfileRadio4Channel(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "channel"), sort)
			}

			if tmp := cfg.ChannelBonding; tmp != nil {
				v["channel_bonding"] = *tmp
			}

			if tmp := cfg.ChannelUtilization; tmp != nil {
				v["channel_utilization"] = *tmp
			}

			if tmp := cfg.Coexistence; tmp != nil {
				v["coexistence"] = *tmp
			}

			if tmp := cfg.Darrp; tmp != nil {
				v["darrp"] = *tmp
			}

			if tmp := cfg.Drma; tmp != nil {
				v["drma"] = *tmp
			}

			if tmp := cfg.DrmaSensitivity; tmp != nil {
				v["drma_sensitivity"] = *tmp
			}

			if tmp := cfg.Dtim; tmp != nil {
				v["dtim"] = *tmp
			}

			if tmp := cfg.FragThreshold; tmp != nil {
				v["frag_threshold"] = *tmp
			}

			if tmp := cfg.FrequencyHandoff; tmp != nil {
				v["frequency_handoff"] = *tmp
			}

			if tmp := cfg.IperfProtocol; tmp != nil {
				v["iperf_protocol"] = *tmp
			}

			if tmp := cfg.IperfServerPort; tmp != nil {
				v["iperf_server_port"] = *tmp
			}

			if tmp := cfg.MaxClients; tmp != nil {
				v["max_clients"] = *tmp
			}

			if tmp := cfg.MaxDistance; tmp != nil {
				v["max_distance"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			if tmp := cfg.OptionalAntenna; tmp != nil {
				v["optional_antenna"] = *tmp
			}

			if tmp := cfg.PowerLevel; tmp != nil {
				v["power_level"] = *tmp
			}

			if tmp := cfg.PowerMode; tmp != nil {
				v["power_mode"] = *tmp
			}

			if tmp := cfg.PowerValue; tmp != nil {
				v["power_value"] = *tmp
			}

			if tmp := cfg.PowersaveOptimize; tmp != nil {
				v["powersave_optimize"] = *tmp
			}

			if tmp := cfg.ProtectionMode; tmp != nil {
				v["protection_mode"] = *tmp
			}

			if tmp := cfg.RtsThreshold; tmp != nil {
				v["rts_threshold"] = *tmp
			}

			if tmp := cfg.SamBssid; tmp != nil {
				v["sam_bssid"] = *tmp
			}

			if tmp := cfg.SamCaptivePortal; tmp != nil {
				v["sam_captive_portal"] = *tmp
			}

			if tmp := cfg.SamCwpFailureString; tmp != nil {
				v["sam_cwp_failure_string"] = *tmp
			}

			if tmp := cfg.SamCwpMatchString; tmp != nil {
				v["sam_cwp_match_string"] = *tmp
			}

			if tmp := cfg.SamCwpPassword; tmp != nil {
				v["sam_cwp_password"] = *tmp
			}

			if tmp := cfg.SamCwpSuccessString; tmp != nil {
				v["sam_cwp_success_string"] = *tmp
			}

			if tmp := cfg.SamCwpTestUrl; tmp != nil {
				v["sam_cwp_test_url"] = *tmp
			}

			if tmp := cfg.SamCwpUsername; tmp != nil {
				v["sam_cwp_username"] = *tmp
			}

			if tmp := cfg.SamPassword; tmp != nil {
				v["sam_password"] = *tmp
			}

			if tmp := cfg.SamReportIntv; tmp != nil {
				v["sam_report_intv"] = *tmp
			}

			if tmp := cfg.SamSecurityType; tmp != nil {
				v["sam_security_type"] = *tmp
			}

			if tmp := cfg.SamServer; tmp != nil {
				v["sam_server"] = *tmp
			}

			if tmp := cfg.SamServerFqdn; tmp != nil {
				v["sam_server_fqdn"] = *tmp
			}

			if tmp := cfg.SamServerIp; tmp != nil {
				v["sam_server_ip"] = *tmp
			}

			if tmp := cfg.SamServerType; tmp != nil {
				v["sam_server_type"] = *tmp
			}

			if tmp := cfg.SamSsid; tmp != nil {
				v["sam_ssid"] = *tmp
			}

			if tmp := cfg.SamTest; tmp != nil {
				v["sam_test"] = *tmp
			}

			if tmp := cfg.SamUsername; tmp != nil {
				v["sam_username"] = *tmp
			}

			if tmp := cfg.ShortGuardInterval; tmp != nil {
				v["short_guard_interval"] = *tmp
			}

			if tmp := cfg.SpectrumAnalysis; tmp != nil {
				v["spectrum_analysis"] = *tmp
			}

			if tmp := cfg.TransmitOptimize; tmp != nil {
				v["transmit_optimize"] = *tmp
			}

			if tmp := cfg.VapAll; tmp != nil {
				v["vap_all"] = *tmp
			}

			if tmp := cfg.Vaps; tmp != nil {
				v["vaps"] = flattenWirelessControllerWtpProfileRadio4Vaps(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "vaps"), sort)
			}

			if tmp := cfg.WidsProfile; tmp != nil {
				v["wids_profile"] = *tmp
			}

			if tmp := cfg.ZeroWaitDfs; tmp != nil {
				v["zero_wait_dfs"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWirelessControllerWtpProfileRadio4Channel(d *schema.ResourceData, v *[]models.WirelessControllerWtpProfileRadio4Channel, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Chan; tmp != nil {
				v["chan"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "chan")
	}

	return flat
}

func flattenWirelessControllerWtpProfileRadio4Vaps(d *schema.ResourceData, v *[]models.WirelessControllerWtpProfileRadio4Vaps, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenWirelessControllerWtpProfileSplitTunnelingAcl(d *schema.ResourceData, v *[]models.WirelessControllerWtpProfileSplitTunnelingAcl, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.DestIp; tmp != nil {
				if tmp = utils.Ipv4Read(d, fmt.Sprintf("%s.%d.dest_ip", prefix, i), *tmp); tmp != nil {
					v["dest_ip"] = *tmp
				}
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectWirelessControllerWtpProfile(d *schema.ResourceData, o *models.WirelessControllerWtpProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Allowaccess != nil {
		v := *o.Allowaccess

		if err = d.Set("allowaccess", v); err != nil {
			return diag.Errorf("error reading allowaccess: %v", err)
		}
	}

	if o.ApCountry != nil {
		v := *o.ApCountry

		if err = d.Set("ap_country", v); err != nil {
			return diag.Errorf("error reading ap_country: %v", err)
		}
	}

	if o.ApHandoff != nil {
		v := *o.ApHandoff

		if err = d.Set("ap_handoff", v); err != nil {
			return diag.Errorf("error reading ap_handoff: %v", err)
		}
	}

	if o.ApcfgProfile != nil {
		v := *o.ApcfgProfile

		if err = d.Set("apcfg_profile", v); err != nil {
			return diag.Errorf("error reading apcfg_profile: %v", err)
		}
	}

	if o.BleProfile != nil {
		v := *o.BleProfile

		if err = d.Set("ble_profile", v); err != nil {
			return diag.Errorf("error reading ble_profile: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.ConsoleLogin != nil {
		v := *o.ConsoleLogin

		if err = d.Set("console_login", v); err != nil {
			return diag.Errorf("error reading console_login: %v", err)
		}
	}

	if o.ControlMessageOffload != nil {
		v := *o.ControlMessageOffload

		if err = d.Set("control_message_offload", v); err != nil {
			return diag.Errorf("error reading control_message_offload: %v", err)
		}
	}

	if o.DenyMacList != nil {
		if err = d.Set("deny_mac_list", flattenWirelessControllerWtpProfileDenyMacList(d, o.DenyMacList, "deny_mac_list", sort)); err != nil {
			return diag.Errorf("error reading deny_mac_list: %v", err)
		}
	}

	if o.DtlsInKernel != nil {
		v := *o.DtlsInKernel

		if err = d.Set("dtls_in_kernel", v); err != nil {
			return diag.Errorf("error reading dtls_in_kernel: %v", err)
		}
	}

	if o.DtlsPolicy != nil {
		v := *o.DtlsPolicy

		if err = d.Set("dtls_policy", v); err != nil {
			return diag.Errorf("error reading dtls_policy: %v", err)
		}
	}

	if o.EnergyEfficientEthernet != nil {
		v := *o.EnergyEfficientEthernet

		if err = d.Set("energy_efficient_ethernet", v); err != nil {
			return diag.Errorf("error reading energy_efficient_ethernet: %v", err)
		}
	}

	if _, ok := d.GetOk("esl_ses_dongle"); ok {
		if o.EslSesDongle != nil {
			if err = d.Set("esl_ses_dongle", flattenWirelessControllerWtpProfileEslSesDongle(d, o.EslSesDongle, "esl_ses_dongle", sort)); err != nil {
				return diag.Errorf("error reading esl_ses_dongle: %v", err)
			}
		}
	}

	if o.ExtInfoEnable != nil {
		v := *o.ExtInfoEnable

		if err = d.Set("ext_info_enable", v); err != nil {
			return diag.Errorf("error reading ext_info_enable: %v", err)
		}
	}

	if o.FrequencyHandoff != nil {
		v := *o.FrequencyHandoff

		if err = d.Set("frequency_handoff", v); err != nil {
			return diag.Errorf("error reading frequency_handoff: %v", err)
		}
	}

	if o.HandoffRoaming != nil {
		v := *o.HandoffRoaming

		if err = d.Set("handoff_roaming", v); err != nil {
			return diag.Errorf("error reading handoff_roaming: %v", err)
		}
	}

	if o.HandoffRssi != nil {
		v := *o.HandoffRssi

		if err = d.Set("handoff_rssi", v); err != nil {
			return diag.Errorf("error reading handoff_rssi: %v", err)
		}
	}

	if o.HandoffStaThresh != nil {
		v := *o.HandoffStaThresh

		if err = d.Set("handoff_sta_thresh", v); err != nil {
			return diag.Errorf("error reading handoff_sta_thresh: %v", err)
		}
	}

	if o.IndoorOutdoorDeployment != nil {
		v := *o.IndoorOutdoorDeployment

		if err = d.Set("indoor_outdoor_deployment", v); err != nil {
			return diag.Errorf("error reading indoor_outdoor_deployment: %v", err)
		}
	}

	if o.IpFragmentPreventing != nil {
		v := *o.IpFragmentPreventing

		if err = d.Set("ip_fragment_preventing", v); err != nil {
			return diag.Errorf("error reading ip_fragment_preventing: %v", err)
		}
	}

	if _, ok := d.GetOk("lan"); ok {
		if o.Lan != nil {
			if err = d.Set("lan", flattenWirelessControllerWtpProfileLan(d, o.Lan, "lan", sort)); err != nil {
				return diag.Errorf("error reading lan: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("lbs"); ok {
		if o.Lbs != nil {
			if err = d.Set("lbs", flattenWirelessControllerWtpProfileLbs(d, o.Lbs, "lbs", sort)); err != nil {
				return diag.Errorf("error reading lbs: %v", err)
			}
		}
	}

	if o.LedSchedules != nil {
		if err = d.Set("led_schedules", flattenWirelessControllerWtpProfileLedSchedules(d, o.LedSchedules, "led_schedules", sort)); err != nil {
			return diag.Errorf("error reading led_schedules: %v", err)
		}
	}

	if o.LedState != nil {
		v := *o.LedState

		if err = d.Set("led_state", v); err != nil {
			return diag.Errorf("error reading led_state: %v", err)
		}
	}

	if o.Lldp != nil {
		v := *o.Lldp

		if err = d.Set("lldp", v); err != nil {
			return diag.Errorf("error reading lldp: %v", err)
		}
	}

	if o.LoginPasswd != nil {
		v := *o.LoginPasswd

		if v == "ENC XXXX" {
		} else if err = d.Set("login_passwd", v); err != nil {
			return diag.Errorf("error reading login_passwd: %v", err)
		}
	}

	if o.LoginPasswdChange != nil {
		v := *o.LoginPasswdChange

		if err = d.Set("login_passwd_change", v); err != nil {
			return diag.Errorf("error reading login_passwd_change: %v", err)
		}
	}

	if o.MaxClients != nil {
		v := *o.MaxClients

		if err = d.Set("max_clients", v); err != nil {
			return diag.Errorf("error reading max_clients: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if _, ok := d.GetOk("platform"); ok {
		if o.Platform != nil {
			if err = d.Set("platform", flattenWirelessControllerWtpProfilePlatform(d, o.Platform, "platform", sort)); err != nil {
				return diag.Errorf("error reading platform: %v", err)
			}
		}
	}

	if o.PoeMode != nil {
		v := *o.PoeMode

		if err = d.Set("poe_mode", v); err != nil {
			return diag.Errorf("error reading poe_mode: %v", err)
		}
	}

	if _, ok := d.GetOk("radio_1"); ok {
		if o.Radio1 != nil {
			if err = d.Set("radio_1", flattenWirelessControllerWtpProfileRadio1(d, o.Radio1, "radio_1", sort)); err != nil {
				return diag.Errorf("error reading radio_1: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("radio_2"); ok {
		if o.Radio2 != nil {
			if err = d.Set("radio_2", flattenWirelessControllerWtpProfileRadio2(d, o.Radio2, "radio_2", sort)); err != nil {
				return diag.Errorf("error reading radio_2: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("radio_3"); ok {
		if o.Radio3 != nil {
			if err = d.Set("radio_3", flattenWirelessControllerWtpProfileRadio3(d, o.Radio3, "radio_3", sort)); err != nil {
				return diag.Errorf("error reading radio_3: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("radio_4"); ok {
		if o.Radio4 != nil {
			if err = d.Set("radio_4", flattenWirelessControllerWtpProfileRadio4(d, o.Radio4, "radio_4", sort)); err != nil {
				return diag.Errorf("error reading radio_4: %v", err)
			}
		}
	}

	if o.SplitTunnelingAcl != nil {
		if err = d.Set("split_tunneling_acl", flattenWirelessControllerWtpProfileSplitTunnelingAcl(d, o.SplitTunnelingAcl, "split_tunneling_acl", sort)); err != nil {
			return diag.Errorf("error reading split_tunneling_acl: %v", err)
		}
	}

	if o.SplitTunnelingAclLocalApSubnet != nil {
		v := *o.SplitTunnelingAclLocalApSubnet

		if err = d.Set("split_tunneling_acl_local_ap_subnet", v); err != nil {
			return diag.Errorf("error reading split_tunneling_acl_local_ap_subnet: %v", err)
		}
	}

	if o.SplitTunnelingAclPath != nil {
		v := *o.SplitTunnelingAclPath

		if err = d.Set("split_tunneling_acl_path", v); err != nil {
			return diag.Errorf("error reading split_tunneling_acl_path: %v", err)
		}
	}

	if o.SyslogProfile != nil {
		v := *o.SyslogProfile

		if err = d.Set("syslog_profile", v); err != nil {
			return diag.Errorf("error reading syslog_profile: %v", err)
		}
	}

	if o.TunMtuDownlink != nil {
		v := *o.TunMtuDownlink

		if err = d.Set("tun_mtu_downlink", v); err != nil {
			return diag.Errorf("error reading tun_mtu_downlink: %v", err)
		}
	}

	if o.TunMtuUplink != nil {
		v := *o.TunMtuUplink

		if err = d.Set("tun_mtu_uplink", v); err != nil {
			return diag.Errorf("error reading tun_mtu_uplink: %v", err)
		}
	}

	if o.WanPortAuth != nil {
		v := *o.WanPortAuth

		if err = d.Set("wan_port_auth", v); err != nil {
			return diag.Errorf("error reading wan_port_auth: %v", err)
		}
	}

	if o.WanPortAuthMethods != nil {
		v := *o.WanPortAuthMethods

		if err = d.Set("wan_port_auth_methods", v); err != nil {
			return diag.Errorf("error reading wan_port_auth_methods: %v", err)
		}
	}

	if o.WanPortAuthPassword != nil {
		v := *o.WanPortAuthPassword

		if v == "ENC XXXX" {
		} else if err = d.Set("wan_port_auth_password", v); err != nil {
			return diag.Errorf("error reading wan_port_auth_password: %v", err)
		}
	}

	if o.WanPortAuthUsrname != nil {
		v := *o.WanPortAuthUsrname

		if err = d.Set("wan_port_auth_usrname", v); err != nil {
			return diag.Errorf("error reading wan_port_auth_usrname: %v", err)
		}
	}

	if o.WanPortMode != nil {
		v := *o.WanPortMode

		if err = d.Set("wan_port_mode", v); err != nil {
			return diag.Errorf("error reading wan_port_mode: %v", err)
		}
	}

	return nil
}

func expandWirelessControllerWtpProfileDenyMacList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerWtpProfileDenyMacList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpProfileDenyMacList

	for i := range l {
		tmp := models.WirelessControllerWtpProfileDenyMacList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mac", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mac = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerWtpProfileEslSesDongle(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.WirelessControllerWtpProfileEslSesDongle, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpProfileEslSesDongle

	for i := range l {
		tmp := models.WirelessControllerWtpProfileEslSesDongle{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.apc_addr_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApcAddrType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.apc_fqdn", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApcFqdn = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.apc_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApcIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.apc_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ApcPort = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.coex_level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CoexLevel = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.compliance_level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ComplianceLevel = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.esl_channel", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EslChannel = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.output_power", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OutputPower = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.scd_enable", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ScdEnable = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tls_cert_verification", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TlsCertVerification = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tls_fqdn_verification", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TlsFqdnVerification = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandWirelessControllerWtpProfileLan(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.WirelessControllerWtpProfileLan, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpProfileLan

	for i := range l {
		tmp := models.WirelessControllerWtpProfileLan{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.port_esl_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PortEslMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port_esl_ssid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PortEslSsid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PortMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port_ssid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PortSsid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port1_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Port1Mode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port1_ssid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Port1Ssid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port2_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Port2Mode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port2_ssid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Port2Ssid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port3_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Port3Mode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port3_ssid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Port3Ssid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port4_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Port4Mode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port4_ssid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Port4Ssid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port5_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Port5Mode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port5_ssid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Port5Ssid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port6_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Port6Mode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port6_ssid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Port6Ssid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port7_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Port7Mode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port7_ssid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Port7Ssid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port8_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Port8Mode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port8_ssid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Port8Ssid = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandWirelessControllerWtpProfileLbs(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.WirelessControllerWtpProfileLbs, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpProfileLbs

	for i := range l {
		tmp := models.WirelessControllerWtpProfileLbs{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.aeroscout", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Aeroscout = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.aeroscout_ap_mac", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AeroscoutApMac = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.aeroscout_mmu_report", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AeroscoutMmuReport = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.aeroscout_mu", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AeroscoutMu = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.aeroscout_mu_factor", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AeroscoutMuFactor = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.aeroscout_mu_timeout", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AeroscoutMuTimeout = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.aeroscout_server_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AeroscoutServerIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.aeroscout_server_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AeroscoutServerPort = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ekahau_blink_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EkahauBlinkMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ekahau_tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EkahauTag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.erc_server_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ErcServerIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.erc_server_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ErcServerPort = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortipresence", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fortipresence = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortipresence_ble", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FortipresenceBle = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortipresence_frequency", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.FortipresenceFrequency = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortipresence_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.FortipresencePort = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortipresence_project", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FortipresenceProject = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortipresence_rogue", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FortipresenceRogue = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortipresence_secret", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FortipresenceSecret = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortipresence_server", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FortipresenceServer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortipresence_server_addr_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FortipresenceServerAddrType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortipresence_server_fqdn", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FortipresenceServerFqdn = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortipresence_unassoc", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FortipresenceUnassoc = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.station_locate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StationLocate = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandWirelessControllerWtpProfileLedSchedules(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerWtpProfileLedSchedules, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpProfileLedSchedules

	for i := range l {
		tmp := models.WirelessControllerWtpProfileLedSchedules{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerWtpProfilePlatform(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.WirelessControllerWtpProfilePlatform, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpProfilePlatform

	for i := range l {
		tmp := models.WirelessControllerWtpProfilePlatform{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.ddscan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ddscan = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandWirelessControllerWtpProfileRadio1(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.WirelessControllerWtpProfileRadio1, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpProfileRadio1

	for i := range l {
		tmp := models.WirelessControllerWtpProfileRadio1{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.80211d", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.The80211d = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.airtime_fairness", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AirtimeFairness = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.amsdu", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Amsdu = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_handoff", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApHandoff = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_addr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferAddr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_bufsize", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ApSnifferBufsize = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_chan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ApSnifferChan = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_ctl", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferCtl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_data", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferData = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_mgmt_beacon", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferMgmtBeacon = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_mgmt_other", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferMgmtOther = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_mgmt_probe", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferMgmtProbe = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.arrp_profile", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArrpProfile = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auto_power_high", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AutoPowerHigh = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auto_power_level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AutoPowerLevel = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auto_power_low", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AutoPowerLow = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auto_power_target", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AutoPowerTarget = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.band", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Band = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.band_5g_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Band5gType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bandwidth_admission_control", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BandwidthAdmissionControl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bandwidth_capacity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.BandwidthCapacity = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.beacon_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.BeaconInterval = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bss_color", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.BssColor = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bss_color_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BssColorMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.call_admission_control", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CallAdmissionControl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.call_capacity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.CallCapacity = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.channel", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWirelessControllerWtpProfileRadio1Channel(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerWtpProfileRadio1Channel
			// 	}
			tmp.Channel = v2

		}

		pre_append = fmt.Sprintf("%s.%d.channel_bonding", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ChannelBonding = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.channel_utilization", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ChannelUtilization = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.coexistence", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Coexistence = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.darrp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Darrp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.drma", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Drma = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.drma_sensitivity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DrmaSensitivity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dtim", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Dtim = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.frag_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.FragThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.frequency_handoff", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FrequencyHandoff = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.iperf_protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IperfProtocol = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.iperf_server_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.IperfServerPort = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_clients", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaxClients = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_distance", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaxDistance = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.optional_antenna", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OptionalAntenna = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.power_level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.PowerLevel = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.power_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PowerMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.power_value", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.PowerValue = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.powersave_optimize", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PowersaveOptimize = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.protection_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ProtectionMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rts_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.RtsThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_bssid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamBssid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_captive_portal", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCaptivePortal = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_failure_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpFailureString = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_match_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpMatchString = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_password", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpPassword = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_success_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpSuccessString = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_test_url", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpTestUrl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_username", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpUsername = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_password", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamPassword = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_report_intv", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SamReportIntv = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_security_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamSecurityType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_server", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamServer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_server_fqdn", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamServerFqdn = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_server_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamServerIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_server_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamServerType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_ssid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamSsid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_test", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamTest = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_username", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamUsername = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.short_guard_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ShortGuardInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.spectrum_analysis", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SpectrumAnalysis = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.transmit_optimize", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TransmitOptimize = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vap_all", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VapAll = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vaps", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWirelessControllerWtpProfileRadio1Vaps(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerWtpProfileRadio1Vaps
			// 	}
			tmp.Vaps = v2

		}

		pre_append = fmt.Sprintf("%s.%d.wids_profile", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.WidsProfile = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.zero_wait_dfs", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ZeroWaitDfs = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandWirelessControllerWtpProfileRadio1Channel(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerWtpProfileRadio1Channel, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpProfileRadio1Channel

	for i := range l {
		tmp := models.WirelessControllerWtpProfileRadio1Channel{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.chan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Chan = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerWtpProfileRadio1Vaps(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerWtpProfileRadio1Vaps, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpProfileRadio1Vaps

	for i := range l {
		tmp := models.WirelessControllerWtpProfileRadio1Vaps{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerWtpProfileRadio2(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.WirelessControllerWtpProfileRadio2, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpProfileRadio2

	for i := range l {
		tmp := models.WirelessControllerWtpProfileRadio2{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.80211d", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.The80211d = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.airtime_fairness", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AirtimeFairness = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.amsdu", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Amsdu = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_handoff", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApHandoff = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_addr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferAddr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_bufsize", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ApSnifferBufsize = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_chan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ApSnifferChan = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_ctl", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferCtl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_data", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferData = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_mgmt_beacon", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferMgmtBeacon = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_mgmt_other", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferMgmtOther = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_mgmt_probe", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferMgmtProbe = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.arrp_profile", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArrpProfile = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auto_power_high", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AutoPowerHigh = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auto_power_level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AutoPowerLevel = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auto_power_low", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AutoPowerLow = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auto_power_target", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AutoPowerTarget = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.band", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Band = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.band_5g_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Band5gType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bandwidth_admission_control", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BandwidthAdmissionControl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bandwidth_capacity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.BandwidthCapacity = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.beacon_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.BeaconInterval = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bss_color", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.BssColor = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bss_color_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BssColorMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.call_admission_control", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CallAdmissionControl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.call_capacity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.CallCapacity = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.channel", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWirelessControllerWtpProfileRadio2Channel(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerWtpProfileRadio2Channel
			// 	}
			tmp.Channel = v2

		}

		pre_append = fmt.Sprintf("%s.%d.channel_bonding", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ChannelBonding = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.channel_utilization", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ChannelUtilization = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.coexistence", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Coexistence = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.darrp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Darrp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.drma", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Drma = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.drma_sensitivity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DrmaSensitivity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dtim", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Dtim = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.frag_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.FragThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.frequency_handoff", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FrequencyHandoff = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.iperf_protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IperfProtocol = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.iperf_server_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.IperfServerPort = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_clients", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaxClients = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_distance", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaxDistance = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.optional_antenna", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OptionalAntenna = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.power_level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.PowerLevel = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.power_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PowerMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.power_value", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.PowerValue = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.powersave_optimize", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PowersaveOptimize = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.protection_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ProtectionMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rts_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.RtsThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_bssid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamBssid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_captive_portal", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCaptivePortal = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_failure_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpFailureString = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_match_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpMatchString = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_password", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpPassword = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_success_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpSuccessString = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_test_url", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpTestUrl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_username", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpUsername = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_password", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamPassword = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_report_intv", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SamReportIntv = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_security_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamSecurityType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_server", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamServer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_server_fqdn", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamServerFqdn = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_server_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamServerIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_server_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamServerType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_ssid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamSsid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_test", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamTest = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_username", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamUsername = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.short_guard_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ShortGuardInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.spectrum_analysis", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SpectrumAnalysis = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.transmit_optimize", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TransmitOptimize = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vap_all", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VapAll = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vaps", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWirelessControllerWtpProfileRadio2Vaps(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerWtpProfileRadio2Vaps
			// 	}
			tmp.Vaps = v2

		}

		pre_append = fmt.Sprintf("%s.%d.wids_profile", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.WidsProfile = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.zero_wait_dfs", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ZeroWaitDfs = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandWirelessControllerWtpProfileRadio2Channel(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerWtpProfileRadio2Channel, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpProfileRadio2Channel

	for i := range l {
		tmp := models.WirelessControllerWtpProfileRadio2Channel{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.chan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Chan = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerWtpProfileRadio2Vaps(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerWtpProfileRadio2Vaps, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpProfileRadio2Vaps

	for i := range l {
		tmp := models.WirelessControllerWtpProfileRadio2Vaps{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerWtpProfileRadio3(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.WirelessControllerWtpProfileRadio3, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpProfileRadio3

	for i := range l {
		tmp := models.WirelessControllerWtpProfileRadio3{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.80211d", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.The80211d = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.airtime_fairness", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AirtimeFairness = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.amsdu", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Amsdu = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_handoff", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApHandoff = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_addr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferAddr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_bufsize", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ApSnifferBufsize = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_chan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ApSnifferChan = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_ctl", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferCtl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_data", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferData = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_mgmt_beacon", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferMgmtBeacon = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_mgmt_other", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferMgmtOther = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_mgmt_probe", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferMgmtProbe = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.arrp_profile", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArrpProfile = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auto_power_high", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AutoPowerHigh = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auto_power_level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AutoPowerLevel = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auto_power_low", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AutoPowerLow = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auto_power_target", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AutoPowerTarget = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.band", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Band = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.band_5g_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Band5gType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bandwidth_admission_control", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BandwidthAdmissionControl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bandwidth_capacity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.BandwidthCapacity = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.beacon_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.BeaconInterval = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bss_color", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.BssColor = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bss_color_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BssColorMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.call_admission_control", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CallAdmissionControl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.call_capacity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.CallCapacity = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.channel", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWirelessControllerWtpProfileRadio3Channel(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerWtpProfileRadio3Channel
			// 	}
			tmp.Channel = v2

		}

		pre_append = fmt.Sprintf("%s.%d.channel_bonding", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ChannelBonding = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.channel_utilization", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ChannelUtilization = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.coexistence", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Coexistence = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.darrp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Darrp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.drma", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Drma = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.drma_sensitivity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DrmaSensitivity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dtim", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Dtim = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.frag_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.FragThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.frequency_handoff", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FrequencyHandoff = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.iperf_protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IperfProtocol = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.iperf_server_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.IperfServerPort = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_clients", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaxClients = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_distance", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaxDistance = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.optional_antenna", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OptionalAntenna = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.power_level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.PowerLevel = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.power_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PowerMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.power_value", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.PowerValue = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.powersave_optimize", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PowersaveOptimize = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.protection_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ProtectionMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rts_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.RtsThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_bssid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamBssid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_captive_portal", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCaptivePortal = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_failure_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpFailureString = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_match_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpMatchString = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_password", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpPassword = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_success_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpSuccessString = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_test_url", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpTestUrl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_username", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpUsername = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_password", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamPassword = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_report_intv", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SamReportIntv = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_security_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamSecurityType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_server", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamServer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_server_fqdn", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamServerFqdn = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_server_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamServerIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_server_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamServerType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_ssid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamSsid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_test", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamTest = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_username", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamUsername = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.short_guard_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ShortGuardInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.spectrum_analysis", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SpectrumAnalysis = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.transmit_optimize", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TransmitOptimize = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vap_all", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VapAll = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vaps", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWirelessControllerWtpProfileRadio3Vaps(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerWtpProfileRadio3Vaps
			// 	}
			tmp.Vaps = v2

		}

		pre_append = fmt.Sprintf("%s.%d.wids_profile", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.WidsProfile = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.zero_wait_dfs", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ZeroWaitDfs = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandWirelessControllerWtpProfileRadio3Channel(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerWtpProfileRadio3Channel, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpProfileRadio3Channel

	for i := range l {
		tmp := models.WirelessControllerWtpProfileRadio3Channel{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.chan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Chan = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerWtpProfileRadio3Vaps(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerWtpProfileRadio3Vaps, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpProfileRadio3Vaps

	for i := range l {
		tmp := models.WirelessControllerWtpProfileRadio3Vaps{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerWtpProfileRadio4(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.WirelessControllerWtpProfileRadio4, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpProfileRadio4

	for i := range l {
		tmp := models.WirelessControllerWtpProfileRadio4{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.80211d", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.The80211d = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.airtime_fairness", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AirtimeFairness = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.amsdu", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Amsdu = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_handoff", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApHandoff = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_addr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferAddr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_bufsize", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ApSnifferBufsize = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_chan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ApSnifferChan = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_ctl", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferCtl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_data", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferData = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_mgmt_beacon", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferMgmtBeacon = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_mgmt_other", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferMgmtOther = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ap_sniffer_mgmt_probe", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApSnifferMgmtProbe = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.arrp_profile", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArrpProfile = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auto_power_high", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AutoPowerHigh = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auto_power_level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AutoPowerLevel = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auto_power_low", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AutoPowerLow = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auto_power_target", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AutoPowerTarget = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.band", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Band = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.band_5g_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Band5gType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bandwidth_admission_control", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BandwidthAdmissionControl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bandwidth_capacity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.BandwidthCapacity = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.beacon_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.BeaconInterval = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bss_color", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.BssColor = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bss_color_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BssColorMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.call_admission_control", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CallAdmissionControl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.call_capacity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.CallCapacity = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.channel", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWirelessControllerWtpProfileRadio4Channel(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerWtpProfileRadio4Channel
			// 	}
			tmp.Channel = v2

		}

		pre_append = fmt.Sprintf("%s.%d.channel_bonding", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ChannelBonding = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.channel_utilization", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ChannelUtilization = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.coexistence", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Coexistence = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.darrp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Darrp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.drma", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Drma = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.drma_sensitivity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DrmaSensitivity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dtim", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Dtim = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.frag_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.FragThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.frequency_handoff", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FrequencyHandoff = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.iperf_protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IperfProtocol = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.iperf_server_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.IperfServerPort = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_clients", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaxClients = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_distance", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaxDistance = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.optional_antenna", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OptionalAntenna = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.power_level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.PowerLevel = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.power_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PowerMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.power_value", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.PowerValue = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.powersave_optimize", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PowersaveOptimize = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.protection_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ProtectionMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rts_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.RtsThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_bssid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamBssid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_captive_portal", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCaptivePortal = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_failure_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpFailureString = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_match_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpMatchString = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_password", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpPassword = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_success_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpSuccessString = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_test_url", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpTestUrl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_cwp_username", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamCwpUsername = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_password", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamPassword = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_report_intv", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SamReportIntv = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_security_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamSecurityType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_server", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamServer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_server_fqdn", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamServerFqdn = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_server_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamServerIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_server_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamServerType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_ssid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamSsid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_test", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamTest = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sam_username", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamUsername = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.short_guard_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ShortGuardInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.spectrum_analysis", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SpectrumAnalysis = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.transmit_optimize", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TransmitOptimize = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vap_all", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VapAll = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vaps", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWirelessControllerWtpProfileRadio4Vaps(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerWtpProfileRadio4Vaps
			// 	}
			tmp.Vaps = v2

		}

		pre_append = fmt.Sprintf("%s.%d.wids_profile", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.WidsProfile = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.zero_wait_dfs", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ZeroWaitDfs = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandWirelessControllerWtpProfileRadio4Channel(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerWtpProfileRadio4Channel, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpProfileRadio4Channel

	for i := range l {
		tmp := models.WirelessControllerWtpProfileRadio4Channel{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.chan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Chan = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerWtpProfileRadio4Vaps(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerWtpProfileRadio4Vaps, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpProfileRadio4Vaps

	for i := range l {
		tmp := models.WirelessControllerWtpProfileRadio4Vaps{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerWtpProfileSplitTunnelingAcl(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerWtpProfileSplitTunnelingAcl, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpProfileSplitTunnelingAcl

	for i := range l {
		tmp := models.WirelessControllerWtpProfileSplitTunnelingAcl{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dest_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DestIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWirelessControllerWtpProfile(d *schema.ResourceData, sv string) (*models.WirelessControllerWtpProfile, diag.Diagnostics) {
	obj := models.WirelessControllerWtpProfile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("allowaccess"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("allowaccess", sv)
				diags = append(diags, e)
			}
			obj.Allowaccess = &v2
		}
	}
	if v1, ok := d.GetOk("ap_country"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ap_country", sv)
				diags = append(diags, e)
			}
			obj.ApCountry = &v2
		}
	}
	if v1, ok := d.GetOk("ap_handoff"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("ap_handoff", sv)
				diags = append(diags, e)
			}
			obj.ApHandoff = &v2
		}
	}
	if v1, ok := d.GetOk("apcfg_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("apcfg_profile", sv)
				diags = append(diags, e)
			}
			obj.ApcfgProfile = &v2
		}
	}
	if v1, ok := d.GetOk("ble_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ble_profile", sv)
				diags = append(diags, e)
			}
			obj.BleProfile = &v2
		}
	}
	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
		}
	}
	if v1, ok := d.GetOk("console_login"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("console_login", sv)
				diags = append(diags, e)
			}
			obj.ConsoleLogin = &v2
		}
	}
	if v1, ok := d.GetOk("control_message_offload"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("control_message_offload", sv)
				diags = append(diags, e)
			}
			obj.ControlMessageOffload = &v2
		}
	}
	if v, ok := d.GetOk("deny_mac_list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("deny_mac_list", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerWtpProfileDenyMacList(d, v, "deny_mac_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DenyMacList = t
		}
	} else if d.HasChange("deny_mac_list") {
		old, new := d.GetChange("deny_mac_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DenyMacList = &[]models.WirelessControllerWtpProfileDenyMacList{}
		}
	}
	if v1, ok := d.GetOk("dtls_in_kernel"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dtls_in_kernel", sv)
				diags = append(diags, e)
			}
			obj.DtlsInKernel = &v2
		}
	}
	if v1, ok := d.GetOk("dtls_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dtls_policy", sv)
				diags = append(diags, e)
			}
			obj.DtlsPolicy = &v2
		}
	}
	if v1, ok := d.GetOk("energy_efficient_ethernet"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("energy_efficient_ethernet", sv)
				diags = append(diags, e)
			}
			obj.EnergyEfficientEthernet = &v2
		}
	}
	if v, ok := d.GetOk("esl_ses_dongle"); ok {
		if !utils.CheckVer(sv, "v7.0.1", "") {
			e := utils.AttributeVersionWarning("esl_ses_dongle", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerWtpProfileEslSesDongle(d, v, "esl_ses_dongle", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.EslSesDongle = t
		}
	} else if d.HasChange("esl_ses_dongle") {
		old, new := d.GetChange("esl_ses_dongle")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.EslSesDongle = &models.WirelessControllerWtpProfileEslSesDongle{}
		}
	}
	if v1, ok := d.GetOk("ext_info_enable"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ext_info_enable", sv)
				diags = append(diags, e)
			}
			obj.ExtInfoEnable = &v2
		}
	}
	if v1, ok := d.GetOk("frequency_handoff"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("frequency_handoff", sv)
				diags = append(diags, e)
			}
			obj.FrequencyHandoff = &v2
		}
	}
	if v1, ok := d.GetOk("handoff_roaming"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("handoff_roaming", sv)
				diags = append(diags, e)
			}
			obj.HandoffRoaming = &v2
		}
	}
	if v1, ok := d.GetOk("handoff_rssi"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("handoff_rssi", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HandoffRssi = &tmp
		}
	}
	if v1, ok := d.GetOk("handoff_sta_thresh"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("handoff_sta_thresh", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HandoffStaThresh = &tmp
		}
	}
	if v1, ok := d.GetOk("indoor_outdoor_deployment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("indoor_outdoor_deployment", sv)
				diags = append(diags, e)
			}
			obj.IndoorOutdoorDeployment = &v2
		}
	}
	if v1, ok := d.GetOk("ip_fragment_preventing"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip_fragment_preventing", sv)
				diags = append(diags, e)
			}
			obj.IpFragmentPreventing = &v2
		}
	}
	if v, ok := d.GetOk("lan"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("lan", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerWtpProfileLan(d, v, "lan", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Lan = t
		}
	} else if d.HasChange("lan") {
		old, new := d.GetChange("lan")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Lan = &models.WirelessControllerWtpProfileLan{}
		}
	}
	if v, ok := d.GetOk("lbs"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("lbs", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerWtpProfileLbs(d, v, "lbs", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Lbs = t
		}
	} else if d.HasChange("lbs") {
		old, new := d.GetChange("lbs")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Lbs = &models.WirelessControllerWtpProfileLbs{}
		}
	}
	if v, ok := d.GetOk("led_schedules"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("led_schedules", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerWtpProfileLedSchedules(d, v, "led_schedules", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.LedSchedules = t
		}
	} else if d.HasChange("led_schedules") {
		old, new := d.GetChange("led_schedules")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.LedSchedules = &[]models.WirelessControllerWtpProfileLedSchedules{}
		}
	}
	if v1, ok := d.GetOk("led_state"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("led_state", sv)
				diags = append(diags, e)
			}
			obj.LedState = &v2
		}
	}
	if v1, ok := d.GetOk("lldp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lldp", sv)
				diags = append(diags, e)
			}
			obj.Lldp = &v2
		}
	}
	if v1, ok := d.GetOk("login_passwd"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("login_passwd", sv)
				diags = append(diags, e)
			}
			obj.LoginPasswd = &v2
		}
	}
	if v1, ok := d.GetOk("login_passwd_change"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("login_passwd_change", sv)
				diags = append(diags, e)
			}
			obj.LoginPasswdChange = &v2
		}
	}
	if v1, ok := d.GetOk("max_clients"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_clients", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxClients = &tmp
		}
	}
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v, ok := d.GetOk("platform"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("platform", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerWtpProfilePlatform(d, v, "platform", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Platform = t
		}
	} else if d.HasChange("platform") {
		old, new := d.GetChange("platform")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Platform = &models.WirelessControllerWtpProfilePlatform{}
		}
	}
	if v1, ok := d.GetOk("poe_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("poe_mode", sv)
				diags = append(diags, e)
			}
			obj.PoeMode = &v2
		}
	}
	if v, ok := d.GetOk("radio_1"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("radio_1", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerWtpProfileRadio1(d, v, "radio_1", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Radio1 = t
		}
	} else if d.HasChange("radio_1") {
		old, new := d.GetChange("radio_1")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Radio1 = &models.WirelessControllerWtpProfileRadio1{}
		}
	}
	if v, ok := d.GetOk("radio_2"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("radio_2", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerWtpProfileRadio2(d, v, "radio_2", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Radio2 = t
		}
	} else if d.HasChange("radio_2") {
		old, new := d.GetChange("radio_2")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Radio2 = &models.WirelessControllerWtpProfileRadio2{}
		}
	}
	if v, ok := d.GetOk("radio_3"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("radio_3", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerWtpProfileRadio3(d, v, "radio_3", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Radio3 = t
		}
	} else if d.HasChange("radio_3") {
		old, new := d.GetChange("radio_3")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Radio3 = &models.WirelessControllerWtpProfileRadio3{}
		}
	}
	if v, ok := d.GetOk("radio_4"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("radio_4", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerWtpProfileRadio4(d, v, "radio_4", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Radio4 = t
		}
	} else if d.HasChange("radio_4") {
		old, new := d.GetChange("radio_4")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Radio4 = &models.WirelessControllerWtpProfileRadio4{}
		}
	}
	if v, ok := d.GetOk("split_tunneling_acl"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("split_tunneling_acl", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerWtpProfileSplitTunnelingAcl(d, v, "split_tunneling_acl", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SplitTunnelingAcl = t
		}
	} else if d.HasChange("split_tunneling_acl") {
		old, new := d.GetChange("split_tunneling_acl")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SplitTunnelingAcl = &[]models.WirelessControllerWtpProfileSplitTunnelingAcl{}
		}
	}
	if v1, ok := d.GetOk("split_tunneling_acl_local_ap_subnet"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("split_tunneling_acl_local_ap_subnet", sv)
				diags = append(diags, e)
			}
			obj.SplitTunnelingAclLocalApSubnet = &v2
		}
	}
	if v1, ok := d.GetOk("split_tunneling_acl_path"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("split_tunneling_acl_path", sv)
				diags = append(diags, e)
			}
			obj.SplitTunnelingAclPath = &v2
		}
	}
	if v1, ok := d.GetOk("syslog_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("syslog_profile", sv)
				diags = append(diags, e)
			}
			obj.SyslogProfile = &v2
		}
	}
	if v1, ok := d.GetOk("tun_mtu_downlink"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tun_mtu_downlink", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TunMtuDownlink = &tmp
		}
	}
	if v1, ok := d.GetOk("tun_mtu_uplink"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tun_mtu_uplink", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TunMtuUplink = &tmp
		}
	}
	if v1, ok := d.GetOk("wan_port_auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("wan_port_auth", sv)
				diags = append(diags, e)
			}
			obj.WanPortAuth = &v2
		}
	}
	if v1, ok := d.GetOk("wan_port_auth_methods"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("wan_port_auth_methods", sv)
				diags = append(diags, e)
			}
			obj.WanPortAuthMethods = &v2
		}
	}
	if v1, ok := d.GetOk("wan_port_auth_password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("wan_port_auth_password", sv)
				diags = append(diags, e)
			}
			obj.WanPortAuthPassword = &v2
		}
	}
	if v1, ok := d.GetOk("wan_port_auth_usrname"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("wan_port_auth_usrname", sv)
				diags = append(diags, e)
			}
			obj.WanPortAuthUsrname = &v2
		}
	}
	if v1, ok := d.GetOk("wan_port_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wan_port_mode", sv)
				diags = append(diags, e)
			}
			obj.WanPortMode = &v2
		}
	}
	return &obj, diags
}
