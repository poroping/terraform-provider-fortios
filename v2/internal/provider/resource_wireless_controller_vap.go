// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Virtual Access Points (VAPs).

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

func resourceWirelessControllerVap() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Virtual Access Points (VAPs).",

		CreateContext: resourceWirelessControllerVapCreate,
		ReadContext:   resourceWirelessControllerVapRead,
		UpdateContext: resourceWirelessControllerVapUpdate,
		DeleteContext: resourceWirelessControllerVapDelete,

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
			"access_control_list": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Profile name for access-control-list.",
				Optional:    true,
				Computed:    true,
			},
			"acct_interim_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 86400),

				Description: "WiFi RADIUS accounting interim interval (60 - 86400 sec, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"additional_akms": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Additional AKMs.",
				Optional:         true,
				Computed:         true,
			},
			"address_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Address group ID.",
				Optional:    true,
				Computed:    true,
			},
			"antivirus_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "AntiVirus profile name.",
				Optional:    true,
				Computed:    true,
			},
			"application_list": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Application control list name.",
				Optional:    true,
				Computed:    true,
			},
			"atf_weight": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),

				Description: "Airtime weight in percentage (default = 20).",
				Optional:    true,
				Computed:    true,
			},
			"auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"psk", "radius", "usergroup"}, false),

				Description: "Authentication protocol.",
				Optional:    true,
				Computed:    true,
			},
			"auth_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "HTTPS server certificate.",
				Optional:    true,
				Computed:    true,
			},
			"auth_portal_addr": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Address of captive portal.",
				Optional:    true,
				Computed:    true,
			},
			"beacon_advertising": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Fortinet beacon advertising IE data   (default = empty).",
				Optional:         true,
				Computed:         true,
			},
			"broadcast_ssid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable broadcasting the SSID (default = enable).",
				Optional:    true,
				Computed:    true,
			},
			"broadcast_suppression": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Optional suppression of broadcast messages. For example, you can keep DHCP messages, ARP broadcasts, and so on off of the wireless network.",
				Optional:         true,
				Computed:         true,
			},
			"bss_color_partial": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable 802.11ax partial BSS color (default = enable).",
				Optional:    true,
				Computed:    true,
			},
			"bstm_disassociation_imminent": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable forcing of disassociation after the BSTM request timer has been reached (default = enable).",
				Optional:    true,
				Computed:    true,
			},
			"bstm_load_balancing_disassoc_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),

				Description: "Time interval for client to voluntarily leave AP before forcing a disassociation due to AP load-balancing (0 to 30, default = 10).",
				Optional:    true,
				Computed:    true,
			},
			"bstm_rssi_disassoc_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 2000),

				Description: "Time interval for client to voluntarily leave AP before forcing a disassociation due to low RSSI (0 to 2000, default = 200).",
				Optional:    true,
				Computed:    true,
			},
			"captive_portal_ac_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Local-bridging captive portal ac-name.",
				Optional:    true,
				Computed:    true,
			},
			"captive_portal_auth_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 864000),

				Description: "Hard timeout - AP will always clear the session after timeout regardless of traffic (0 - 864000 sec, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"captive_portal_macauth_radius_secret": {
				Type: schema.TypeString,

				Description: "Secret key to access the macauth RADIUS server.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"captive_portal_macauth_radius_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Captive portal external RADIUS server domain name or IP address.",
				Optional:    true,
				Computed:    true,
			},
			"captive_portal_radius_secret": {
				Type: schema.TypeString,

				Description: "Secret key to access the RADIUS server.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"captive_portal_radius_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Captive portal RADIUS server domain name or IP address.",
				Optional:    true,
				Computed:    true,
			},
			"captive_portal_session_timeout_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 864000),

				Description: "Session timeout interval (0 - 864000 sec, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_address_enforcement": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable DHCP address enforcement (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_lease_time": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 8640000),

				Description: "DHCP lease time in seconds for NAT IP address.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_option43_insertion": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable insertion of DHCP option 43 (default = enable).",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_option82_circuit_id_insertion": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"style-1", "style-2", "style-3", "disable"}, false),

				Description: "Enable/disable DHCP option 82 circuit-id insert (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_option82_insertion": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable DHCP option 82 insert (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_option82_remote_id_insertion": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"style-1", "disable"}, false),

				Description: "Enable/disable DHCP option 82 remote-id insert (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"dynamic_vlan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable dynamic VLAN assignment.",
				Optional:    true,
				Computed:    true,
			},
			"eap_reauth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable EAP re-authentication for WPA-Enterprise security.",
				Optional:    true,
				Computed:    true,
			},
			"eap_reauth_intv": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1800, 864000),

				Description: "EAP re-authentication interval (1800 - 864000 sec, default = 86400).",
				Optional:    true,
				Computed:    true,
			},
			"eapol_key_retries": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable retransmission of EAPOL-Key frames (message 3/4 and group message 1/2) (default = enable).",
				Optional:    true,
				Computed:    true,
			},
			"encrypt": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"TKIP", "AES", "TKIP-AES"}, false),

				Description: "Encryption protocol to use (only available when security is set to a WPA type).",
				Optional:    true,
				Computed:    true,
			},
			"external_fast_roaming": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable fast roaming or pre-authentication with external APs not managed by the FortiGate (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"external_logout": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "URL of external authentication logout server.",
				Optional:    true,
				Computed:    true,
			},
			"external_web": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),

				Description: "URL of external authentication web server.",
				Optional:    true,
				Computed:    true,
			},
			"external_web_format": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto-detect", "no-query-string", "partial-query-string"}, false),

				Description: "URL query parameter detection (default = auto-detect).",
				Optional:    true,
				Computed:    true,
			},
			"fast_bss_transition": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable 802.11r Fast BSS Transition (FT) (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"fast_roaming": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable fast-roaming, or pre-authentication, where supported by clients (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"ft_mobility_domain": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Mobility domain identifier in FT (1 - 65535, default = 1000).",
				Optional:    true,
				Computed:    true,
			},
			"ft_over_ds": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable FT over the Distribution System (DS).",
				Optional:    true,
				Computed:    true,
			},
			"ft_r0_key_lifetime": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Lifetime of the PMK-R0 key in FT, 1-65535 minutes.",
				Optional:    true,
				Computed:    true,
			},
			"gas_comeback_delay": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(100, 10000),

				Description: "GAS comeback delay (0 or 100 - 10000 milliseconds, default = 500).",
				Optional:    true,
				Computed:    true,
			},
			"gas_fragmentation_limit": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(512, 4096),

				Description: "GAS fragmentation limit (512 - 4096, default = 1024).",
				Optional:    true,
				Computed:    true,
			},
			"gtk_rekey": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable GTK rekey for WPA security.",
				Optional:    true,
				Computed:    true,
			},
			"gtk_rekey_intv": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1800, 864000),

				Description: "GTK rekey interval (1800 - 864000 sec, default = 86400).",
				Optional:    true,
				Computed:    true,
			},
			"high_efficiency": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable 802.11ax high efficiency (default = enable).",
				Optional:    true,
				Computed:    true,
			},
			"hotspot20_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Hotspot 2.0 profile name.",
				Optional:    true,
				Computed:    true,
			},
			"igmp_snooping": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IGMP snooping.",
				Optional:    true,
				Computed:    true,
			},
			"intra_vap_privacy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable blocking communication between clients on the same SSID (called intra-SSID privacy) (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"ip": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4ClassnetHost,

				Description: "IP address and subnet mask for the local standalone NAT subnet.",
				Optional:    true,
				Computed:    true,
			},
			"ips_sensor": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "IPS sensor name.",
				Optional:    true,
				Computed:    true,
			},
			"ipv6_rules": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Optional rules of IPv6 packets. For example, you can keep RA, RS and so on off of the wireless network.",
				Optional:         true,
				Computed:         true,
			},
			"key": {
				Type: schema.TypeString,

				Description: "WEP Key.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"keyindex": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 4),

				Description: "WEP key index (1 - 4).",
				Optional:    true,
				Computed:    true,
			},
			"ldpc": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "rx", "tx", "rxtx"}, false),

				Description: "VAP low-density parity-check (LDPC) coding configuration.",
				Optional:    true,
				Computed:    true,
			},
			"local_authentication": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable AP local authentication.",
				Optional:    true,
				Computed:    true,
			},
			"local_bridging": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable bridging of wireless and Ethernet interfaces on the FortiAP (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"local_lan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"allow", "deny"}, false),

				Description: "Allow/deny traffic destined for a Class A, B, or C private IP address (default = allow).",
				Optional:    true,
				Computed:    true,
			},
			"local_standalone": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable AP local standalone (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"local_standalone_dns": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable AP local standalone DNS.",
				Optional:    true,
				Computed:    true,
			},
			"local_standalone_dns_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IPv4 addresses for the local standalone DNS.",
				Optional:    true,
				Computed:    true,
			},
			"local_standalone_nat": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable AP local standalone NAT mode.",
				Optional:    true,
				Computed:    true,
			},
			"mac_auth_bypass": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable MAC authentication bypass.",
				Optional:    true,
				Computed:    true,
			},
			"mac_called_station_delimiter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"hyphen", "single-hyphen", "colon", "none"}, false),

				Description: "MAC called station delimiter (default = hyphen).",
				Optional:    true,
				Computed:    true,
			},
			"mac_calling_station_delimiter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"hyphen", "single-hyphen", "colon", "none"}, false),

				Description: "MAC calling station delimiter (default = hyphen).",
				Optional:    true,
				Computed:    true,
			},
			"mac_case": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"uppercase", "lowercase"}, false),

				Description: "MAC case (default = uppercase).",
				Optional:    true,
				Computed:    true,
			},
			"mac_filter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable MAC filtering to block wireless clients by mac address.",
				Optional:    true,
				Computed:    true,
			},
			"mac_filter_list": {
				Type:        schema.TypeList,
				Description: "Create a list of MAC addresses for MAC address filtering.",
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

							Description: "MAC address.",
							Optional:    true,
							Computed:    true,
						},
						"mac_filter_policy": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "deny"}, false),

							Description: "Deny or allow the client with this MAC address.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"mac_filter_policy_other": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"allow", "deny"}, false),

				Description: "Allow or block clients with MAC addresses that are not in the filter list.",
				Optional:    true,
				Computed:    true,
			},
			"mac_password_delimiter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"hyphen", "single-hyphen", "colon", "none"}, false),

				Description: "MAC authentication password delimiter (default = hyphen).",
				Optional:    true,
				Computed:    true,
			},
			"mac_username_delimiter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"hyphen", "single-hyphen", "colon", "none"}, false),

				Description: "MAC authentication username delimiter (default = hyphen).",
				Optional:    true,
				Computed:    true,
			},
			"max_clients": {
				Type: schema.TypeInt,

				Description: "Maximum number of clients that can connect simultaneously to the VAP (default = 0, meaning no limitation).",
				Optional:    true,
				Computed:    true,
			},
			"max_clients_ap": {
				Type: schema.TypeInt,

				Description: "Maximum number of clients that can connect simultaneously to the VAP per AP radio (default = 0, meaning no limitation).",
				Optional:    true,
				Computed:    true,
			},
			"mbo": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable Multiband Operation (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"mbo_cell_data_conn_pref": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"excluded", "prefer-not", "prefer-use"}, false),

				Description: "MBO cell data connection preference (0, 1, or 255, default = 1).",
				Optional:    true,
				Computed:    true,
			},
			"me_disable_thresh": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 256),

				Description: "Disable multicast enhancement when this many clients are receiving multicast traffic.",
				Optional:    true,
				Computed:    true,
			},
			"mesh_backhaul": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable using this VAP as a WiFi mesh backhaul (default = disable). This entry is only available when security is set to a WPA type or open.",
				Optional:    true,
				Computed:    true,
			},
			"mpsk": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable multiple PSK authentication.",
				Optional:    true,
				Computed:    true,
			},
			"mpsk_concurrent_clients": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).",
				Optional:    true,
				Computed:    true,
			},
			"mpsk_key": {
				Type:        schema.TypeList,
				Description: "List of multiple PSK entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comment": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Comment.",
							Optional:    true,
							Computed:    true,
						},
						"concurrent_clients": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Number of clients that can connect using this pre-shared key.",
							Optional:    true,
							Computed:    true,
						},
						"key_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Pre-shared key name.",
							Optional:    true,
							Computed:    true,
						},
						"mpsk_schedules": {
							Type:        schema.TypeList,
							Description: "Firewall schedule for MPSK passphrase. The passphrase will be effective only when at least one schedule is valid.",
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
						"passphrase": {
							Type: schema.TypeString,

							Description: "WPA Pre-shared key.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
					},
				},
			},
			"mpsk_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "MPSK profile name.",
				Optional:    true,
				Computed:    true,
			},
			"mu_mimo": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Multi-user MIMO (default = enable).",
				Optional:    true,
				Computed:    true,
			},
			"multicast_enhance": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable converting multicast to unicast to improve performance (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"multicast_rate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"0", "6000", "12000", "24000"}, false),

				Description: "Multicast rate (0, 6000, 12000, or 24000 kbps, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"nac": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable network access control.",
				Optional:    true,
				Computed:    true,
			},
			"nac_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "NAC profile name.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Virtual AP name.",
				ForceNew:    true,
				Required:    true,
			},
			"neighbor_report_dual_band": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable dual-band neighbor report (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"okc": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable Opportunistic Key Caching (OKC) (default = enable).",
				Optional:    true,
				Computed:    true,
			},
			"osen": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable OSEN as part of key management (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"owe_groups": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "OWE-Groups.",
				Optional:         true,
				Computed:         true,
			},
			"owe_transition": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable OWE transition mode support.",
				Optional:    true,
				Computed:    true,
			},
			"owe_transition_ssid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),

				Description: "OWE transition mode peer SSID.",
				Optional:    true,
				Computed:    true,
			},
			"passphrase": {
				Type: schema.TypeString,

				Description: "WPA pre-shared key (PSK) to be used to authenticate WiFi users.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"pmf": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable", "optional"}, false),

				Description: "Protected Management Frames (PMF) support (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"pmf_assoc_comeback_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 20),

				Description: "Protected Management Frames (PMF) comeback maximum timeout (1-20 sec).",
				Optional:    true,
				Computed:    true,
			},
			"pmf_sa_query_retry_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 5),

				Description: "Protected Management Frames (PMF) SA query retry timeout interval (1 - 5 100s of msec).",
				Optional:    true,
				Computed:    true,
			},
			"port_macauth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "radius", "address-group"}, false),

				Description: "Enable/disable LAN port MAC authentication (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"port_macauth_reauth_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(120, 65535),

				Description: "LAN port MAC authentication re-authentication timeout value (default = 7200 sec).",
				Optional:    true,
				Computed:    true,
			},
			"port_macauth_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 65535),

				Description: "LAN port MAC authentication idle timeout value (default = 600 sec).",
				Optional:    true,
				Computed:    true,
			},
			"portal_message_override_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Replacement message group for this VAP (only available when security is set to a captive portal type).",
				Optional:    true,
				Computed:    true,
			},
			"portal_message_overrides": {
				Type:        schema.TypeList,
				Description: "Individual message overrides.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_disclaimer_page": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Override auth-disclaimer-page message with message from portal-message-overrides group.",
							Optional:    true,
							Computed:    true,
						},
						"auth_login_failed_page": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Override auth-login-failed-page message with message from portal-message-overrides group.",
							Optional:    true,
							Computed:    true,
						},
						"auth_login_page": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Override auth-login-page message with message from portal-message-overrides group.",
							Optional:    true,
							Computed:    true,
						},
						"auth_reject_page": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Override auth-reject-page message with message from portal-message-overrides group.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"portal_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auth", "auth+disclaimer", "disclaimer", "email-collect", "cmcc", "cmcc-macauth", "auth-mac", "external-auth", "external-macauth"}, false),

				Description: "Captive portal functionality. Configure how the captive portal authenticates users and whether it includes a disclaimer.",
				Optional:    true,
				Computed:    true,
			},
			"primary_wag_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Primary wireless access gateway profile name.",
				Optional:    true,
				Computed:    true,
			},
			"probe_resp_suppression": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable probe response suppression (to ignore weak signals) (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"probe_resp_threshold": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),

				Description: "Minimum signal level/threshold in dBm required for the AP response to probe requests (-95 to -20, default = -80).",
				Optional:    true,
				Computed:    true,
			},
			"ptk_rekey": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable PTK rekey for WPA-Enterprise security.",
				Optional:    true,
				Computed:    true,
			},
			"ptk_rekey_intv": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1800, 864000),

				Description: "PTK rekey interval (1800 - 864000 sec, default = 86400).",
				Optional:    true,
				Computed:    true,
			},
			"qos_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Quality of service profile name.",
				Optional:    true,
				Computed:    true,
			},
			"quarantine": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable station quarantine (default = enable).",
				Optional:    true,
				Computed:    true,
			},
			"radio_2g_threshold": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),

				Description: "Minimum signal level/threshold in dBm required for the AP response to receive a packet in 2.4G band (-95 to -20, default = -79).",
				Optional:    true,
				Computed:    true,
			},
			"radio_5g_threshold": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),

				Description: "Minimum signal level/threshold in dBm required for the AP response to receive a packet in 5G band(-95 to -20, default = -76).",
				Optional:    true,
				Computed:    true,
			},
			"radio_sensitivity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable software radio sensitivity (to ignore weak signals) (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"radius_mac_auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable RADIUS-based MAC authentication of clients (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"radius_mac_auth_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "RADIUS-based MAC authentication server.",
				Optional:    true,
				Computed:    true,
			},
			"radius_mac_auth_usergroups": {
				Type:        schema.TypeList,
				Description: "Selective user groups that are permitted for RADIUS mac authentication.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "User group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"radius_mac_mpsk_auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable RADIUS-based MAC authentication of clients for MPSK authentication (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"radius_mac_mpsk_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1800, 864000),

				Description: "RADIUS MAC MPSK cache timeout interval (1800 - 864000, default = 86400).",
				Optional:    true,
				Computed:    true,
			},
			"radius_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "RADIUS server to be used to authenticate WiFi users.",
				Optional:    true,
				Computed:    true,
			},
			"rates_11a": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Allowed data rates for 802.11a.",
				Optional:         true,
				Computed:         true,
			},
			"rates_11ac_ss12": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Allowed data rates for 802.11ac with 1 or 2 spatial streams.",
				Optional:         true,
				Computed:         true,
			},
			"rates_11ac_ss34": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Allowed data rates for 802.11ac with 3 or 4 spatial streams.",
				Optional:         true,
				Computed:         true,
			},
			"rates_11ax_ss12": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Allowed data rates for 802.11ax with 1 or 2 spatial streams.",
				Optional:         true,
				Computed:         true,
			},
			"rates_11ax_ss34": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Allowed data rates for 802.11ax with 3 or 4 spatial streams.",
				Optional:         true,
				Computed:         true,
			},
			"rates_11bg": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Allowed data rates for 802.11b/g.",
				Optional:         true,
				Computed:         true,
			},
			"rates_11n_ss12": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Allowed data rates for 802.11n with 1 or 2 spatial streams.",
				Optional:         true,
				Computed:         true,
			},
			"rates_11n_ss34": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Allowed data rates for 802.11n with 3 or 4 spatial streams.",
				Optional:         true,
				Computed:         true,
			},
			"sae_groups": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "SAE-Groups.",
				Optional:         true,
				Computed:         true,
			},
			"sae_password": {
				Type: schema.TypeString,

				Description: "WPA3 SAE password to be used to authenticate WiFi users.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"scan_botnet_connections": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "monitor", "block"}, false),

				Description: "Block or monitor connections to Botnet servers or disable Botnet scanning.",
				Optional:    true,
				Computed:    true,
			},
			"schedule": {
				Type:        schema.TypeList,
				Description: "Firewall schedules for enabling this VAP on the FortiAP. This VAP will be enabled when at least one of the schedules is valid. Separate multiple schedule names with a space.",
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
			"secondary_wag_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Secondary wireless access gateway profile name.",
				Optional:    true,
				Computed:    true,
			},
			"security": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"open", "captive-portal", "wep64", "wep128", "wpa-personal", "wpa-personal+captive-portal", "wpa-enterprise", "wpa-only-personal", "wpa-only-personal+captive-portal", "wpa-only-enterprise", "wpa2-only-personal", "wpa2-only-personal+captive-portal", "wpa2-only-enterprise", "wpa3-enterprise", "wpa3-only-enterprise", "wpa3-enterprise-transition", "wpa3-sae", "wpa3-sae-transition", "owe", "osen"}, false),

				Description: "Security mode for the wireless interface (default = wpa2-only-personal).",
				Optional:    true,
				Computed:    true,
			},
			"security_exempt_list": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Optional security exempt list for captive portal authentication.",
				Optional:    true,
				Computed:    true,
			},
			"security_redirect_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),

				Description: "Optional URL for redirecting users after they pass captive portal authentication.",
				Optional:    true,
				Computed:    true,
			},
			"selected_usergroups": {
				Type:        schema.TypeList,
				Description: "Selective user groups that are permitted to authenticate.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "User group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"split_tunneling": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable split tunneling (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"ssid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),

				Description: "IEEE 802.11 service set identifier (SSID) for the wireless interface. Users who wish to use the wireless network must configure their computers to access this SSID name.",
				Optional:    true,
				Computed:    true,
			},
			"sticky_client_remove": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable sticky client remove to maintain good signal level clients in SSID (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"sticky_client_threshold_2g": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),

				Description: "Minimum signal level/threshold in dBm required for the 2G client to be serviced by the AP (-95 to -20, default = -79).",
				Optional:    true,
				Computed:    true,
			},
			"sticky_client_threshold_5g": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),

				Description: "Minimum signal level/threshold in dBm required for the 5G client to be serviced by the AP (-95 to -20, default = -76).",
				Optional:    true,
				Computed:    true,
			},
			"target_wake_time": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable 802.11ax target wake time (default = enable).",
				Optional:    true,
				Computed:    true,
			},
			"tkip_counter_measure": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable TKIP counter measure.",
				Optional:    true,
				Computed:    true,
			},
			"tunnel_echo_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "The time interval to send echo to both primary and secondary tunnel peers (1 - 65535 sec, default = 300).",
				Optional:    true,
				Computed:    true,
			},
			"tunnel_fallback_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "The time interval for secondary tunnel to fall back to primary tunnel (0 - 65535 sec, default = 7200).",
				Optional:    true,
				Computed:    true,
			},
			"usergroup": {
				Type:        schema.TypeList,
				Description: "Firewall user group to be used to authenticate WiFi users.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "User group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"utm_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable UTM logging.",
				Optional:    true,
				Computed:    true,
			},
			"utm_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "UTM profile name.",
				Optional:    true,
				Computed:    true,
			},
			"utm_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to add one or more security profiles (AV, IPS, etc.) to the VAP.",
				Optional:    true,
				Computed:    true,
			},
			"vlan_auto": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable automatic management of SSID VLAN interface.",
				Optional:    true,
				Computed:    true,
			},
			"vlan_name": {
				Type:        schema.TypeList,
				Description: "Table for mapping VLAN name to VLAN ID.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "VLAN name.",
							Optional:    true,
							Computed:    true,
						},
						"vlan_id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4094),

							Description: "VLAN ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"vlan_pool": {
				Type:        schema.TypeList,
				Description: "VLAN pool.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4094),

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"wtp_group": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "WTP group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"vlan_pooling": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"wtp-group", "round-robin", "hash", "disable"}, false),

				Description: "Enable/disable VLAN pooling, to allow grouping of multiple wireless controller VLANs into VLAN pools (default = disable). When set to wtp-group, VLAN pooling occurs with VLAN assignment by wtp-group.",
				Optional:    true,
				Computed:    true,
			},
			"vlanid": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4094),

				Description: "Optional VLAN ID.",
				Optional:    true,
				Computed:    true,
			},
			"voice_enterprise": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable 802.11k and 802.11v assisted Voice-Enterprise roaming (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"webfilter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "WebFilter profile name.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWirelessControllerVapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WirelessControllerVap resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerVap(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerVap(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerVap")
	}

	return resourceWirelessControllerVapRead(ctx, d, meta)
}

func resourceWirelessControllerVapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerVap(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerVap(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerVap resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerVap")
	}

	return resourceWirelessControllerVapRead(ctx, d, meta)
}

func resourceWirelessControllerVapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerVap(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerVap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerVapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerVap(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerVap resource: %v", err)
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

	diags := refreshObjectWirelessControllerVap(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWirelessControllerVapMacFilterList(d *schema.ResourceData, v *[]models.WirelessControllerVapMacFilterList, prefix string, sort bool) interface{} {
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

			if tmp := cfg.MacFilterPolicy; tmp != nil {
				v["mac_filter_policy"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenWirelessControllerVapMpskKey(d *schema.ResourceData, v *[]models.WirelessControllerVapMpskKey, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Comment; tmp != nil {
				v["comment"] = *tmp
			}

			if tmp := cfg.ConcurrentClients; tmp != nil {
				v["concurrent_clients"] = *tmp
			}

			if tmp := cfg.KeyName; tmp != nil {
				v["key_name"] = *tmp
			}

			if tmp := cfg.MpskSchedules; tmp != nil {
				v["mpsk_schedules"] = flattenWirelessControllerVapMpskKeyMpskSchedules(d, tmp, prefix+"mpsk_schedules", sort)
			}

			if tmp := cfg.Passphrase; tmp != nil {
				v["passphrase"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "key_name")
	}

	return flat
}

func flattenWirelessControllerVapMpskKeyMpskSchedules(d *schema.ResourceData, v *[]models.WirelessControllerVapMpskKeyMpskSchedules, prefix string, sort bool) interface{} {
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

func flattenWirelessControllerVapPortalMessageOverrides(d *schema.ResourceData, v *[]models.WirelessControllerVapPortalMessageOverrides, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AuthDisclaimerPage; tmp != nil {
				v["auth_disclaimer_page"] = *tmp
			}

			if tmp := cfg.AuthLoginFailedPage; tmp != nil {
				v["auth_login_failed_page"] = *tmp
			}

			if tmp := cfg.AuthLoginPage; tmp != nil {
				v["auth_login_page"] = *tmp
			}

			if tmp := cfg.AuthRejectPage; tmp != nil {
				v["auth_reject_page"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWirelessControllerVapRadiusMacAuthUsergroups(d *schema.ResourceData, v *[]models.WirelessControllerVapRadiusMacAuthUsergroups, prefix string, sort bool) interface{} {
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

func flattenWirelessControllerVapSchedule(d *schema.ResourceData, v *[]models.WirelessControllerVapSchedule, prefix string, sort bool) interface{} {
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

func flattenWirelessControllerVapSelectedUsergroups(d *schema.ResourceData, v *[]models.WirelessControllerVapSelectedUsergroups, prefix string, sort bool) interface{} {
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

func flattenWirelessControllerVapUsergroup(d *schema.ResourceData, v *[]models.WirelessControllerVapUsergroup, prefix string, sort bool) interface{} {
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

func flattenWirelessControllerVapVlanName(d *schema.ResourceData, v *[]models.WirelessControllerVapVlanName, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.VlanId; tmp != nil {
				v["vlan_id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenWirelessControllerVapVlanPool(d *schema.ResourceData, v *[]models.WirelessControllerVapVlanPool, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.WtpGroup; tmp != nil {
				v["wtp_group"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectWirelessControllerVap(d *schema.ResourceData, o *models.WirelessControllerVap, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AccessControlList != nil {
		v := *o.AccessControlList

		if err = d.Set("access_control_list", v); err != nil {
			return diag.Errorf("error reading access_control_list: %v", err)
		}
	}

	if o.AcctInterimInterval != nil {
		v := *o.AcctInterimInterval

		if err = d.Set("acct_interim_interval", v); err != nil {
			return diag.Errorf("error reading acct_interim_interval: %v", err)
		}
	}

	if o.AdditionalAkms != nil {
		v := *o.AdditionalAkms

		if err = d.Set("additional_akms", v); err != nil {
			return diag.Errorf("error reading additional_akms: %v", err)
		}
	}

	if o.AddressGroup != nil {
		v := *o.AddressGroup

		if err = d.Set("address_group", v); err != nil {
			return diag.Errorf("error reading address_group: %v", err)
		}
	}

	if o.AntivirusProfile != nil {
		v := *o.AntivirusProfile

		if err = d.Set("antivirus_profile", v); err != nil {
			return diag.Errorf("error reading antivirus_profile: %v", err)
		}
	}

	if o.ApplicationList != nil {
		v := *o.ApplicationList

		if err = d.Set("application_list", v); err != nil {
			return diag.Errorf("error reading application_list: %v", err)
		}
	}

	if o.AtfWeight != nil {
		v := *o.AtfWeight

		if err = d.Set("atf_weight", v); err != nil {
			return diag.Errorf("error reading atf_weight: %v", err)
		}
	}

	if o.Auth != nil {
		v := *o.Auth

		if err = d.Set("auth", v); err != nil {
			return diag.Errorf("error reading auth: %v", err)
		}
	}

	if o.AuthCert != nil {
		v := *o.AuthCert

		if err = d.Set("auth_cert", v); err != nil {
			return diag.Errorf("error reading auth_cert: %v", err)
		}
	}

	if o.AuthPortalAddr != nil {
		v := *o.AuthPortalAddr

		if err = d.Set("auth_portal_addr", v); err != nil {
			return diag.Errorf("error reading auth_portal_addr: %v", err)
		}
	}

	if o.BeaconAdvertising != nil {
		v := *o.BeaconAdvertising

		if err = d.Set("beacon_advertising", v); err != nil {
			return diag.Errorf("error reading beacon_advertising: %v", err)
		}
	}

	if o.BroadcastSsid != nil {
		v := *o.BroadcastSsid

		if err = d.Set("broadcast_ssid", v); err != nil {
			return diag.Errorf("error reading broadcast_ssid: %v", err)
		}
	}

	if o.BroadcastSuppression != nil {
		v := *o.BroadcastSuppression

		if err = d.Set("broadcast_suppression", v); err != nil {
			return diag.Errorf("error reading broadcast_suppression: %v", err)
		}
	}

	if o.BssColorPartial != nil {
		v := *o.BssColorPartial

		if err = d.Set("bss_color_partial", v); err != nil {
			return diag.Errorf("error reading bss_color_partial: %v", err)
		}
	}

	if o.BstmDisassociationImminent != nil {
		v := *o.BstmDisassociationImminent

		if err = d.Set("bstm_disassociation_imminent", v); err != nil {
			return diag.Errorf("error reading bstm_disassociation_imminent: %v", err)
		}
	}

	if o.BstmLoadBalancingDisassocTimer != nil {
		v := *o.BstmLoadBalancingDisassocTimer

		if err = d.Set("bstm_load_balancing_disassoc_timer", v); err != nil {
			return diag.Errorf("error reading bstm_load_balancing_disassoc_timer: %v", err)
		}
	}

	if o.BstmRssiDisassocTimer != nil {
		v := *o.BstmRssiDisassocTimer

		if err = d.Set("bstm_rssi_disassoc_timer", v); err != nil {
			return diag.Errorf("error reading bstm_rssi_disassoc_timer: %v", err)
		}
	}

	if o.CaptivePortalAcName != nil {
		v := *o.CaptivePortalAcName

		if err = d.Set("captive_portal_ac_name", v); err != nil {
			return diag.Errorf("error reading captive_portal_ac_name: %v", err)
		}
	}

	if o.CaptivePortalAuthTimeout != nil {
		v := *o.CaptivePortalAuthTimeout

		if err = d.Set("captive_portal_auth_timeout", v); err != nil {
			return diag.Errorf("error reading captive_portal_auth_timeout: %v", err)
		}
	}

	if o.CaptivePortalMacauthRadiusSecret != nil {
		v := *o.CaptivePortalMacauthRadiusSecret

		if v == "ENC XXXX" {
		} else if err = d.Set("captive_portal_macauth_radius_secret", v); err != nil {
			return diag.Errorf("error reading captive_portal_macauth_radius_secret: %v", err)
		}
	}

	if o.CaptivePortalMacauthRadiusServer != nil {
		v := *o.CaptivePortalMacauthRadiusServer

		if err = d.Set("captive_portal_macauth_radius_server", v); err != nil {
			return diag.Errorf("error reading captive_portal_macauth_radius_server: %v", err)
		}
	}

	if o.CaptivePortalRadiusSecret != nil {
		v := *o.CaptivePortalRadiusSecret

		if v == "ENC XXXX" {
		} else if err = d.Set("captive_portal_radius_secret", v); err != nil {
			return diag.Errorf("error reading captive_portal_radius_secret: %v", err)
		}
	}

	if o.CaptivePortalRadiusServer != nil {
		v := *o.CaptivePortalRadiusServer

		if err = d.Set("captive_portal_radius_server", v); err != nil {
			return diag.Errorf("error reading captive_portal_radius_server: %v", err)
		}
	}

	if o.CaptivePortalSessionTimeoutInterval != nil {
		v := *o.CaptivePortalSessionTimeoutInterval

		if err = d.Set("captive_portal_session_timeout_interval", v); err != nil {
			return diag.Errorf("error reading captive_portal_session_timeout_interval: %v", err)
		}
	}

	if o.DhcpAddressEnforcement != nil {
		v := *o.DhcpAddressEnforcement

		if err = d.Set("dhcp_address_enforcement", v); err != nil {
			return diag.Errorf("error reading dhcp_address_enforcement: %v", err)
		}
	}

	if o.DhcpLeaseTime != nil {
		v := *o.DhcpLeaseTime

		if err = d.Set("dhcp_lease_time", v); err != nil {
			return diag.Errorf("error reading dhcp_lease_time: %v", err)
		}
	}

	if o.DhcpOption43Insertion != nil {
		v := *o.DhcpOption43Insertion

		if err = d.Set("dhcp_option43_insertion", v); err != nil {
			return diag.Errorf("error reading dhcp_option43_insertion: %v", err)
		}
	}

	if o.DhcpOption82CircuitIdInsertion != nil {
		v := *o.DhcpOption82CircuitIdInsertion

		if err = d.Set("dhcp_option82_circuit_id_insertion", v); err != nil {
			return diag.Errorf("error reading dhcp_option82_circuit_id_insertion: %v", err)
		}
	}

	if o.DhcpOption82Insertion != nil {
		v := *o.DhcpOption82Insertion

		if err = d.Set("dhcp_option82_insertion", v); err != nil {
			return diag.Errorf("error reading dhcp_option82_insertion: %v", err)
		}
	}

	if o.DhcpOption82RemoteIdInsertion != nil {
		v := *o.DhcpOption82RemoteIdInsertion

		if err = d.Set("dhcp_option82_remote_id_insertion", v); err != nil {
			return diag.Errorf("error reading dhcp_option82_remote_id_insertion: %v", err)
		}
	}

	if o.DynamicVlan != nil {
		v := *o.DynamicVlan

		if err = d.Set("dynamic_vlan", v); err != nil {
			return diag.Errorf("error reading dynamic_vlan: %v", err)
		}
	}

	if o.EapReauth != nil {
		v := *o.EapReauth

		if err = d.Set("eap_reauth", v); err != nil {
			return diag.Errorf("error reading eap_reauth: %v", err)
		}
	}

	if o.EapReauthIntv != nil {
		v := *o.EapReauthIntv

		if err = d.Set("eap_reauth_intv", v); err != nil {
			return diag.Errorf("error reading eap_reauth_intv: %v", err)
		}
	}

	if o.EapolKeyRetries != nil {
		v := *o.EapolKeyRetries

		if err = d.Set("eapol_key_retries", v); err != nil {
			return diag.Errorf("error reading eapol_key_retries: %v", err)
		}
	}

	if o.Encrypt != nil {
		v := *o.Encrypt

		if err = d.Set("encrypt", v); err != nil {
			return diag.Errorf("error reading encrypt: %v", err)
		}
	}

	if o.ExternalFastRoaming != nil {
		v := *o.ExternalFastRoaming

		if err = d.Set("external_fast_roaming", v); err != nil {
			return diag.Errorf("error reading external_fast_roaming: %v", err)
		}
	}

	if o.ExternalLogout != nil {
		v := *o.ExternalLogout

		if err = d.Set("external_logout", v); err != nil {
			return diag.Errorf("error reading external_logout: %v", err)
		}
	}

	if o.ExternalWeb != nil {
		v := *o.ExternalWeb

		if err = d.Set("external_web", v); err != nil {
			return diag.Errorf("error reading external_web: %v", err)
		}
	}

	if o.ExternalWebFormat != nil {
		v := *o.ExternalWebFormat

		if err = d.Set("external_web_format", v); err != nil {
			return diag.Errorf("error reading external_web_format: %v", err)
		}
	}

	if o.FastBssTransition != nil {
		v := *o.FastBssTransition

		if err = d.Set("fast_bss_transition", v); err != nil {
			return diag.Errorf("error reading fast_bss_transition: %v", err)
		}
	}

	if o.FastRoaming != nil {
		v := *o.FastRoaming

		if err = d.Set("fast_roaming", v); err != nil {
			return diag.Errorf("error reading fast_roaming: %v", err)
		}
	}

	if o.FtMobilityDomain != nil {
		v := *o.FtMobilityDomain

		if err = d.Set("ft_mobility_domain", v); err != nil {
			return diag.Errorf("error reading ft_mobility_domain: %v", err)
		}
	}

	if o.FtOverDs != nil {
		v := *o.FtOverDs

		if err = d.Set("ft_over_ds", v); err != nil {
			return diag.Errorf("error reading ft_over_ds: %v", err)
		}
	}

	if o.FtR0KeyLifetime != nil {
		v := *o.FtR0KeyLifetime

		if err = d.Set("ft_r0_key_lifetime", v); err != nil {
			return diag.Errorf("error reading ft_r0_key_lifetime: %v", err)
		}
	}

	if o.GasComebackDelay != nil {
		v := *o.GasComebackDelay

		if err = d.Set("gas_comeback_delay", v); err != nil {
			return diag.Errorf("error reading gas_comeback_delay: %v", err)
		}
	}

	if o.GasFragmentationLimit != nil {
		v := *o.GasFragmentationLimit

		if err = d.Set("gas_fragmentation_limit", v); err != nil {
			return diag.Errorf("error reading gas_fragmentation_limit: %v", err)
		}
	}

	if o.GtkRekey != nil {
		v := *o.GtkRekey

		if err = d.Set("gtk_rekey", v); err != nil {
			return diag.Errorf("error reading gtk_rekey: %v", err)
		}
	}

	if o.GtkRekeyIntv != nil {
		v := *o.GtkRekeyIntv

		if err = d.Set("gtk_rekey_intv", v); err != nil {
			return diag.Errorf("error reading gtk_rekey_intv: %v", err)
		}
	}

	if o.HighEfficiency != nil {
		v := *o.HighEfficiency

		if err = d.Set("high_efficiency", v); err != nil {
			return diag.Errorf("error reading high_efficiency: %v", err)
		}
	}

	if o.Hotspot20Profile != nil {
		v := *o.Hotspot20Profile

		if err = d.Set("hotspot20_profile", v); err != nil {
			return diag.Errorf("error reading hotspot20_profile: %v", err)
		}
	}

	if o.IgmpSnooping != nil {
		v := *o.IgmpSnooping

		if err = d.Set("igmp_snooping", v); err != nil {
			return diag.Errorf("error reading igmp_snooping: %v", err)
		}
	}

	if o.IntraVapPrivacy != nil {
		v := *o.IntraVapPrivacy

		if err = d.Set("intra_vap_privacy", v); err != nil {
			return diag.Errorf("error reading intra_vap_privacy: %v", err)
		}
	}

	if o.Ip != nil {
		v := *o.Ip
		if current, ok := d.GetOk("ip"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("ip", v); err != nil {
			return diag.Errorf("error reading ip: %v", err)
		}
	}

	if o.IpsSensor != nil {
		v := *o.IpsSensor

		if err = d.Set("ips_sensor", v); err != nil {
			return diag.Errorf("error reading ips_sensor: %v", err)
		}
	}

	if o.Ipv6Rules != nil {
		v := *o.Ipv6Rules

		if err = d.Set("ipv6_rules", v); err != nil {
			return diag.Errorf("error reading ipv6_rules: %v", err)
		}
	}

	if o.Key != nil {
		v := *o.Key

		if v == "ENC XXXX" {
		} else if err = d.Set("key", v); err != nil {
			return diag.Errorf("error reading key: %v", err)
		}
	}

	if o.Keyindex != nil {
		v := *o.Keyindex

		if err = d.Set("keyindex", v); err != nil {
			return diag.Errorf("error reading keyindex: %v", err)
		}
	}

	if o.Ldpc != nil {
		v := *o.Ldpc

		if err = d.Set("ldpc", v); err != nil {
			return diag.Errorf("error reading ldpc: %v", err)
		}
	}

	if o.LocalAuthentication != nil {
		v := *o.LocalAuthentication

		if err = d.Set("local_authentication", v); err != nil {
			return diag.Errorf("error reading local_authentication: %v", err)
		}
	}

	if o.LocalBridging != nil {
		v := *o.LocalBridging

		if err = d.Set("local_bridging", v); err != nil {
			return diag.Errorf("error reading local_bridging: %v", err)
		}
	}

	if o.LocalLan != nil {
		v := *o.LocalLan

		if err = d.Set("local_lan", v); err != nil {
			return diag.Errorf("error reading local_lan: %v", err)
		}
	}

	if o.LocalStandalone != nil {
		v := *o.LocalStandalone

		if err = d.Set("local_standalone", v); err != nil {
			return diag.Errorf("error reading local_standalone: %v", err)
		}
	}

	if o.LocalStandaloneDns != nil {
		v := *o.LocalStandaloneDns

		if err = d.Set("local_standalone_dns", v); err != nil {
			return diag.Errorf("error reading local_standalone_dns: %v", err)
		}
	}

	if o.LocalStandaloneDnsIp != nil {
		v := *o.LocalStandaloneDnsIp

		if err = d.Set("local_standalone_dns_ip", v); err != nil {
			return diag.Errorf("error reading local_standalone_dns_ip: %v", err)
		}
	}

	if o.LocalStandaloneNat != nil {
		v := *o.LocalStandaloneNat

		if err = d.Set("local_standalone_nat", v); err != nil {
			return diag.Errorf("error reading local_standalone_nat: %v", err)
		}
	}

	if o.MacAuthBypass != nil {
		v := *o.MacAuthBypass

		if err = d.Set("mac_auth_bypass", v); err != nil {
			return diag.Errorf("error reading mac_auth_bypass: %v", err)
		}
	}

	if o.MacCalledStationDelimiter != nil {
		v := *o.MacCalledStationDelimiter

		if err = d.Set("mac_called_station_delimiter", v); err != nil {
			return diag.Errorf("error reading mac_called_station_delimiter: %v", err)
		}
	}

	if o.MacCallingStationDelimiter != nil {
		v := *o.MacCallingStationDelimiter

		if err = d.Set("mac_calling_station_delimiter", v); err != nil {
			return diag.Errorf("error reading mac_calling_station_delimiter: %v", err)
		}
	}

	if o.MacCase != nil {
		v := *o.MacCase

		if err = d.Set("mac_case", v); err != nil {
			return diag.Errorf("error reading mac_case: %v", err)
		}
	}

	if o.MacFilter != nil {
		v := *o.MacFilter

		if err = d.Set("mac_filter", v); err != nil {
			return diag.Errorf("error reading mac_filter: %v", err)
		}
	}

	if o.MacFilterList != nil {
		if err = d.Set("mac_filter_list", flattenWirelessControllerVapMacFilterList(d, o.MacFilterList, "mac_filter_list", sort)); err != nil {
			return diag.Errorf("error reading mac_filter_list: %v", err)
		}
	}

	if o.MacFilterPolicyOther != nil {
		v := *o.MacFilterPolicyOther

		if err = d.Set("mac_filter_policy_other", v); err != nil {
			return diag.Errorf("error reading mac_filter_policy_other: %v", err)
		}
	}

	if o.MacPasswordDelimiter != nil {
		v := *o.MacPasswordDelimiter

		if err = d.Set("mac_password_delimiter", v); err != nil {
			return diag.Errorf("error reading mac_password_delimiter: %v", err)
		}
	}

	if o.MacUsernameDelimiter != nil {
		v := *o.MacUsernameDelimiter

		if err = d.Set("mac_username_delimiter", v); err != nil {
			return diag.Errorf("error reading mac_username_delimiter: %v", err)
		}
	}

	if o.MaxClients != nil {
		v := *o.MaxClients

		if err = d.Set("max_clients", v); err != nil {
			return diag.Errorf("error reading max_clients: %v", err)
		}
	}

	if o.MaxClientsAp != nil {
		v := *o.MaxClientsAp

		if err = d.Set("max_clients_ap", v); err != nil {
			return diag.Errorf("error reading max_clients_ap: %v", err)
		}
	}

	if o.Mbo != nil {
		v := *o.Mbo

		if err = d.Set("mbo", v); err != nil {
			return diag.Errorf("error reading mbo: %v", err)
		}
	}

	if o.MboCellDataConnPref != nil {
		v := *o.MboCellDataConnPref

		if err = d.Set("mbo_cell_data_conn_pref", v); err != nil {
			return diag.Errorf("error reading mbo_cell_data_conn_pref: %v", err)
		}
	}

	if o.MeDisableThresh != nil {
		v := *o.MeDisableThresh

		if err = d.Set("me_disable_thresh", v); err != nil {
			return diag.Errorf("error reading me_disable_thresh: %v", err)
		}
	}

	if o.MeshBackhaul != nil {
		v := *o.MeshBackhaul

		if err = d.Set("mesh_backhaul", v); err != nil {
			return diag.Errorf("error reading mesh_backhaul: %v", err)
		}
	}

	if o.Mpsk != nil {
		v := *o.Mpsk

		if err = d.Set("mpsk", v); err != nil {
			return diag.Errorf("error reading mpsk: %v", err)
		}
	}

	if o.MpskConcurrentClients != nil {
		v := *o.MpskConcurrentClients

		if err = d.Set("mpsk_concurrent_clients", v); err != nil {
			return diag.Errorf("error reading mpsk_concurrent_clients: %v", err)
		}
	}

	if o.MpskKey != nil {
		if err = d.Set("mpsk_key", flattenWirelessControllerVapMpskKey(d, o.MpskKey, "mpsk_key", sort)); err != nil {
			return diag.Errorf("error reading mpsk_key: %v", err)
		}
	}

	if o.MpskProfile != nil {
		v := *o.MpskProfile

		if err = d.Set("mpsk_profile", v); err != nil {
			return diag.Errorf("error reading mpsk_profile: %v", err)
		}
	}

	if o.MuMimo != nil {
		v := *o.MuMimo

		if err = d.Set("mu_mimo", v); err != nil {
			return diag.Errorf("error reading mu_mimo: %v", err)
		}
	}

	if o.MulticastEnhance != nil {
		v := *o.MulticastEnhance

		if err = d.Set("multicast_enhance", v); err != nil {
			return diag.Errorf("error reading multicast_enhance: %v", err)
		}
	}

	if o.MulticastRate != nil {
		v := *o.MulticastRate

		if err = d.Set("multicast_rate", v); err != nil {
			return diag.Errorf("error reading multicast_rate: %v", err)
		}
	}

	if o.Nac != nil {
		v := *o.Nac

		if err = d.Set("nac", v); err != nil {
			return diag.Errorf("error reading nac: %v", err)
		}
	}

	if o.NacProfile != nil {
		v := *o.NacProfile

		if err = d.Set("nac_profile", v); err != nil {
			return diag.Errorf("error reading nac_profile: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.NeighborReportDualBand != nil {
		v := *o.NeighborReportDualBand

		if err = d.Set("neighbor_report_dual_band", v); err != nil {
			return diag.Errorf("error reading neighbor_report_dual_band: %v", err)
		}
	}

	if o.Okc != nil {
		v := *o.Okc

		if err = d.Set("okc", v); err != nil {
			return diag.Errorf("error reading okc: %v", err)
		}
	}

	if o.Osen != nil {
		v := *o.Osen

		if err = d.Set("osen", v); err != nil {
			return diag.Errorf("error reading osen: %v", err)
		}
	}

	if o.OweGroups != nil {
		v := *o.OweGroups

		if err = d.Set("owe_groups", v); err != nil {
			return diag.Errorf("error reading owe_groups: %v", err)
		}
	}

	if o.OweTransition != nil {
		v := *o.OweTransition

		if err = d.Set("owe_transition", v); err != nil {
			return diag.Errorf("error reading owe_transition: %v", err)
		}
	}

	if o.OweTransitionSsid != nil {
		v := *o.OweTransitionSsid

		if err = d.Set("owe_transition_ssid", v); err != nil {
			return diag.Errorf("error reading owe_transition_ssid: %v", err)
		}
	}

	if o.Passphrase != nil {
		v := *o.Passphrase

		if v == "ENC XXXX" {
		} else if err = d.Set("passphrase", v); err != nil {
			return diag.Errorf("error reading passphrase: %v", err)
		}
	}

	if o.Pmf != nil {
		v := *o.Pmf

		if err = d.Set("pmf", v); err != nil {
			return diag.Errorf("error reading pmf: %v", err)
		}
	}

	if o.PmfAssocComebackTimeout != nil {
		v := *o.PmfAssocComebackTimeout

		if err = d.Set("pmf_assoc_comeback_timeout", v); err != nil {
			return diag.Errorf("error reading pmf_assoc_comeback_timeout: %v", err)
		}
	}

	if o.PmfSaQueryRetryTimeout != nil {
		v := *o.PmfSaQueryRetryTimeout

		if err = d.Set("pmf_sa_query_retry_timeout", v); err != nil {
			return diag.Errorf("error reading pmf_sa_query_retry_timeout: %v", err)
		}
	}

	if o.PortMacauth != nil {
		v := *o.PortMacauth

		if err = d.Set("port_macauth", v); err != nil {
			return diag.Errorf("error reading port_macauth: %v", err)
		}
	}

	if o.PortMacauthReauthTimeout != nil {
		v := *o.PortMacauthReauthTimeout

		if err = d.Set("port_macauth_reauth_timeout", v); err != nil {
			return diag.Errorf("error reading port_macauth_reauth_timeout: %v", err)
		}
	}

	if o.PortMacauthTimeout != nil {
		v := *o.PortMacauthTimeout

		if err = d.Set("port_macauth_timeout", v); err != nil {
			return diag.Errorf("error reading port_macauth_timeout: %v", err)
		}
	}

	if o.PortalMessageOverrideGroup != nil {
		v := *o.PortalMessageOverrideGroup

		if err = d.Set("portal_message_override_group", v); err != nil {
			return diag.Errorf("error reading portal_message_override_group: %v", err)
		}
	}

	if o.PortalMessageOverrides != nil {
		if err = d.Set("portal_message_overrides", flattenWirelessControllerVapPortalMessageOverrides(d, o.PortalMessageOverrides, "portal_message_overrides", sort)); err != nil {
			return diag.Errorf("error reading portal_message_overrides: %v", err)
		}
	}

	if o.PortalType != nil {
		v := *o.PortalType

		if err = d.Set("portal_type", v); err != nil {
			return diag.Errorf("error reading portal_type: %v", err)
		}
	}

	if o.PrimaryWagProfile != nil {
		v := *o.PrimaryWagProfile

		if err = d.Set("primary_wag_profile", v); err != nil {
			return diag.Errorf("error reading primary_wag_profile: %v", err)
		}
	}

	if o.ProbeRespSuppression != nil {
		v := *o.ProbeRespSuppression

		if err = d.Set("probe_resp_suppression", v); err != nil {
			return diag.Errorf("error reading probe_resp_suppression: %v", err)
		}
	}

	if o.ProbeRespThreshold != nil {
		v := *o.ProbeRespThreshold

		if err = d.Set("probe_resp_threshold", v); err != nil {
			return diag.Errorf("error reading probe_resp_threshold: %v", err)
		}
	}

	if o.PtkRekey != nil {
		v := *o.PtkRekey

		if err = d.Set("ptk_rekey", v); err != nil {
			return diag.Errorf("error reading ptk_rekey: %v", err)
		}
	}

	if o.PtkRekeyIntv != nil {
		v := *o.PtkRekeyIntv

		if err = d.Set("ptk_rekey_intv", v); err != nil {
			return diag.Errorf("error reading ptk_rekey_intv: %v", err)
		}
	}

	if o.QosProfile != nil {
		v := *o.QosProfile

		if err = d.Set("qos_profile", v); err != nil {
			return diag.Errorf("error reading qos_profile: %v", err)
		}
	}

	if o.Quarantine != nil {
		v := *o.Quarantine

		if err = d.Set("quarantine", v); err != nil {
			return diag.Errorf("error reading quarantine: %v", err)
		}
	}

	if o.Radio2gThreshold != nil {
		v := *o.Radio2gThreshold

		if err = d.Set("radio_2g_threshold", v); err != nil {
			return diag.Errorf("error reading radio_2g_threshold: %v", err)
		}
	}

	if o.Radio5gThreshold != nil {
		v := *o.Radio5gThreshold

		if err = d.Set("radio_5g_threshold", v); err != nil {
			return diag.Errorf("error reading radio_5g_threshold: %v", err)
		}
	}

	if o.RadioSensitivity != nil {
		v := *o.RadioSensitivity

		if err = d.Set("radio_sensitivity", v); err != nil {
			return diag.Errorf("error reading radio_sensitivity: %v", err)
		}
	}

	if o.RadiusMacAuth != nil {
		v := *o.RadiusMacAuth

		if err = d.Set("radius_mac_auth", v); err != nil {
			return diag.Errorf("error reading radius_mac_auth: %v", err)
		}
	}

	if o.RadiusMacAuthServer != nil {
		v := *o.RadiusMacAuthServer

		if err = d.Set("radius_mac_auth_server", v); err != nil {
			return diag.Errorf("error reading radius_mac_auth_server: %v", err)
		}
	}

	if o.RadiusMacAuthUsergroups != nil {
		if err = d.Set("radius_mac_auth_usergroups", flattenWirelessControllerVapRadiusMacAuthUsergroups(d, o.RadiusMacAuthUsergroups, "radius_mac_auth_usergroups", sort)); err != nil {
			return diag.Errorf("error reading radius_mac_auth_usergroups: %v", err)
		}
	}

	if o.RadiusMacMpskAuth != nil {
		v := *o.RadiusMacMpskAuth

		if err = d.Set("radius_mac_mpsk_auth", v); err != nil {
			return diag.Errorf("error reading radius_mac_mpsk_auth: %v", err)
		}
	}

	if o.RadiusMacMpskTimeout != nil {
		v := *o.RadiusMacMpskTimeout

		if err = d.Set("radius_mac_mpsk_timeout", v); err != nil {
			return diag.Errorf("error reading radius_mac_mpsk_timeout: %v", err)
		}
	}

	if o.RadiusServer != nil {
		v := *o.RadiusServer

		if err = d.Set("radius_server", v); err != nil {
			return diag.Errorf("error reading radius_server: %v", err)
		}
	}

	if o.Rates11a != nil {
		v := *o.Rates11a

		if err = d.Set("rates_11a", v); err != nil {
			return diag.Errorf("error reading rates_11a: %v", err)
		}
	}

	if o.Rates11acSs12 != nil {
		v := *o.Rates11acSs12

		if err = d.Set("rates_11ac_ss12", v); err != nil {
			return diag.Errorf("error reading rates_11ac_ss12: %v", err)
		}
	}

	if o.Rates11acSs34 != nil {
		v := *o.Rates11acSs34

		if err = d.Set("rates_11ac_ss34", v); err != nil {
			return diag.Errorf("error reading rates_11ac_ss34: %v", err)
		}
	}

	if o.Rates11axSs12 != nil {
		v := *o.Rates11axSs12

		if err = d.Set("rates_11ax_ss12", v); err != nil {
			return diag.Errorf("error reading rates_11ax_ss12: %v", err)
		}
	}

	if o.Rates11axSs34 != nil {
		v := *o.Rates11axSs34

		if err = d.Set("rates_11ax_ss34", v); err != nil {
			return diag.Errorf("error reading rates_11ax_ss34: %v", err)
		}
	}

	if o.Rates11bg != nil {
		v := *o.Rates11bg

		if err = d.Set("rates_11bg", v); err != nil {
			return diag.Errorf("error reading rates_11bg: %v", err)
		}
	}

	if o.Rates11nSs12 != nil {
		v := *o.Rates11nSs12

		if err = d.Set("rates_11n_ss12", v); err != nil {
			return diag.Errorf("error reading rates_11n_ss12: %v", err)
		}
	}

	if o.Rates11nSs34 != nil {
		v := *o.Rates11nSs34

		if err = d.Set("rates_11n_ss34", v); err != nil {
			return diag.Errorf("error reading rates_11n_ss34: %v", err)
		}
	}

	if o.SaeGroups != nil {
		v := *o.SaeGroups

		if err = d.Set("sae_groups", v); err != nil {
			return diag.Errorf("error reading sae_groups: %v", err)
		}
	}

	if o.SaePassword != nil {
		v := *o.SaePassword

		if v == "ENC XXXX" {
		} else if err = d.Set("sae_password", v); err != nil {
			return diag.Errorf("error reading sae_password: %v", err)
		}
	}

	if o.ScanBotnetConnections != nil {
		v := *o.ScanBotnetConnections

		if err = d.Set("scan_botnet_connections", v); err != nil {
			return diag.Errorf("error reading scan_botnet_connections: %v", err)
		}
	}

	if o.Schedule != nil {
		if err = d.Set("schedule", flattenWirelessControllerVapSchedule(d, o.Schedule, "schedule", sort)); err != nil {
			return diag.Errorf("error reading schedule: %v", err)
		}
	}

	if o.SecondaryWagProfile != nil {
		v := *o.SecondaryWagProfile

		if err = d.Set("secondary_wag_profile", v); err != nil {
			return diag.Errorf("error reading secondary_wag_profile: %v", err)
		}
	}

	if o.Security != nil {
		v := *o.Security

		if err = d.Set("security", v); err != nil {
			return diag.Errorf("error reading security: %v", err)
		}
	}

	if o.SecurityExemptList != nil {
		v := *o.SecurityExemptList

		if err = d.Set("security_exempt_list", v); err != nil {
			return diag.Errorf("error reading security_exempt_list: %v", err)
		}
	}

	if o.SecurityRedirectUrl != nil {
		v := *o.SecurityRedirectUrl

		if err = d.Set("security_redirect_url", v); err != nil {
			return diag.Errorf("error reading security_redirect_url: %v", err)
		}
	}

	if o.SelectedUsergroups != nil {
		if err = d.Set("selected_usergroups", flattenWirelessControllerVapSelectedUsergroups(d, o.SelectedUsergroups, "selected_usergroups", sort)); err != nil {
			return diag.Errorf("error reading selected_usergroups: %v", err)
		}
	}

	if o.SplitTunneling != nil {
		v := *o.SplitTunneling

		if err = d.Set("split_tunneling", v); err != nil {
			return diag.Errorf("error reading split_tunneling: %v", err)
		}
	}

	if o.Ssid != nil {
		v := *o.Ssid

		if err = d.Set("ssid", v); err != nil {
			return diag.Errorf("error reading ssid: %v", err)
		}
	}

	if o.StickyClientRemove != nil {
		v := *o.StickyClientRemove

		if err = d.Set("sticky_client_remove", v); err != nil {
			return diag.Errorf("error reading sticky_client_remove: %v", err)
		}
	}

	if o.StickyClientThreshold2g != nil {
		v := *o.StickyClientThreshold2g

		if err = d.Set("sticky_client_threshold_2g", v); err != nil {
			return diag.Errorf("error reading sticky_client_threshold_2g: %v", err)
		}
	}

	if o.StickyClientThreshold5g != nil {
		v := *o.StickyClientThreshold5g

		if err = d.Set("sticky_client_threshold_5g", v); err != nil {
			return diag.Errorf("error reading sticky_client_threshold_5g: %v", err)
		}
	}

	if o.TargetWakeTime != nil {
		v := *o.TargetWakeTime

		if err = d.Set("target_wake_time", v); err != nil {
			return diag.Errorf("error reading target_wake_time: %v", err)
		}
	}

	if o.TkipCounterMeasure != nil {
		v := *o.TkipCounterMeasure

		if err = d.Set("tkip_counter_measure", v); err != nil {
			return diag.Errorf("error reading tkip_counter_measure: %v", err)
		}
	}

	if o.TunnelEchoInterval != nil {
		v := *o.TunnelEchoInterval

		if err = d.Set("tunnel_echo_interval", v); err != nil {
			return diag.Errorf("error reading tunnel_echo_interval: %v", err)
		}
	}

	if o.TunnelFallbackInterval != nil {
		v := *o.TunnelFallbackInterval

		if err = d.Set("tunnel_fallback_interval", v); err != nil {
			return diag.Errorf("error reading tunnel_fallback_interval: %v", err)
		}
	}

	if o.Usergroup != nil {
		if err = d.Set("usergroup", flattenWirelessControllerVapUsergroup(d, o.Usergroup, "usergroup", sort)); err != nil {
			return diag.Errorf("error reading usergroup: %v", err)
		}
	}

	if o.UtmLog != nil {
		v := *o.UtmLog

		if err = d.Set("utm_log", v); err != nil {
			return diag.Errorf("error reading utm_log: %v", err)
		}
	}

	if o.UtmProfile != nil {
		v := *o.UtmProfile

		if err = d.Set("utm_profile", v); err != nil {
			return diag.Errorf("error reading utm_profile: %v", err)
		}
	}

	if o.UtmStatus != nil {
		v := *o.UtmStatus

		if err = d.Set("utm_status", v); err != nil {
			return diag.Errorf("error reading utm_status: %v", err)
		}
	}

	if o.VlanAuto != nil {
		v := *o.VlanAuto

		if err = d.Set("vlan_auto", v); err != nil {
			return diag.Errorf("error reading vlan_auto: %v", err)
		}
	}

	if o.VlanName != nil {
		if err = d.Set("vlan_name", flattenWirelessControllerVapVlanName(d, o.VlanName, "vlan_name", sort)); err != nil {
			return diag.Errorf("error reading vlan_name: %v", err)
		}
	}

	if o.VlanPool != nil {
		if err = d.Set("vlan_pool", flattenWirelessControllerVapVlanPool(d, o.VlanPool, "vlan_pool", sort)); err != nil {
			return diag.Errorf("error reading vlan_pool: %v", err)
		}
	}

	if o.VlanPooling != nil {
		v := *o.VlanPooling

		if err = d.Set("vlan_pooling", v); err != nil {
			return diag.Errorf("error reading vlan_pooling: %v", err)
		}
	}

	if o.Vlanid != nil {
		v := *o.Vlanid

		if err = d.Set("vlanid", v); err != nil {
			return diag.Errorf("error reading vlanid: %v", err)
		}
	}

	if o.VoiceEnterprise != nil {
		v := *o.VoiceEnterprise

		if err = d.Set("voice_enterprise", v); err != nil {
			return diag.Errorf("error reading voice_enterprise: %v", err)
		}
	}

	if o.WebfilterProfile != nil {
		v := *o.WebfilterProfile

		if err = d.Set("webfilter_profile", v); err != nil {
			return diag.Errorf("error reading webfilter_profile: %v", err)
		}
	}

	return nil
}

func expandWirelessControllerVapMacFilterList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerVapMacFilterList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerVapMacFilterList

	for i := range l {
		tmp := models.WirelessControllerVapMacFilterList{}
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

		pre_append = fmt.Sprintf("%s.%d.mac_filter_policy", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MacFilterPolicy = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerVapMpskKey(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerVapMpskKey, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerVapMpskKey

	for i := range l {
		tmp := models.WirelessControllerVapMpskKey{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.comment", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Comment = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.concurrent_clients", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ConcurrentClients = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.key_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeyName = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mpsk_schedules", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWirelessControllerVapMpskKeyMpskSchedules(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerVapMpskKeyMpskSchedules
			// 	}
			tmp.MpskSchedules = v2

		}

		pre_append = fmt.Sprintf("%s.%d.passphrase", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Passphrase = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerVapMpskKeyMpskSchedules(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerVapMpskKeyMpskSchedules, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerVapMpskKeyMpskSchedules

	for i := range l {
		tmp := models.WirelessControllerVapMpskKeyMpskSchedules{}
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

func expandWirelessControllerVapPortalMessageOverrides(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerVapPortalMessageOverrides, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerVapPortalMessageOverrides

	for i := range l {
		tmp := models.WirelessControllerVapPortalMessageOverrides{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.auth_disclaimer_page", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthDisclaimerPage = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auth_login_failed_page", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthLoginFailedPage = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auth_login_page", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthLoginPage = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auth_reject_page", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthRejectPage = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerVapRadiusMacAuthUsergroups(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerVapRadiusMacAuthUsergroups, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerVapRadiusMacAuthUsergroups

	for i := range l {
		tmp := models.WirelessControllerVapRadiusMacAuthUsergroups{}
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

func expandWirelessControllerVapSchedule(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerVapSchedule, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerVapSchedule

	for i := range l {
		tmp := models.WirelessControllerVapSchedule{}
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

func expandWirelessControllerVapSelectedUsergroups(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerVapSelectedUsergroups, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerVapSelectedUsergroups

	for i := range l {
		tmp := models.WirelessControllerVapSelectedUsergroups{}
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

func expandWirelessControllerVapUsergroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerVapUsergroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerVapUsergroup

	for i := range l {
		tmp := models.WirelessControllerVapUsergroup{}
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

func expandWirelessControllerVapVlanName(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerVapVlanName, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerVapVlanName

	for i := range l {
		tmp := models.WirelessControllerVapVlanName{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.VlanId = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerVapVlanPool(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerVapVlanPool, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerVapVlanPool

	for i := range l {
		tmp := models.WirelessControllerVapVlanPool{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.wtp_group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.WtpGroup = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWirelessControllerVap(d *schema.ResourceData, sv string) (*models.WirelessControllerVap, diag.Diagnostics) {
	obj := models.WirelessControllerVap{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("access_control_list"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("access_control_list", sv)
				diags = append(diags, e)
			}
			obj.AccessControlList = &v2
		}
	}
	if v1, ok := d.GetOk("acct_interim_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("acct_interim_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AcctInterimInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("additional_akms"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("additional_akms", sv)
				diags = append(diags, e)
			}
			obj.AdditionalAkms = &v2
		}
	}
	if v1, ok := d.GetOk("address_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("address_group", sv)
				diags = append(diags, e)
			}
			obj.AddressGroup = &v2
		}
	}
	if v1, ok := d.GetOk("antivirus_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("antivirus_profile", sv)
				diags = append(diags, e)
			}
			obj.AntivirusProfile = &v2
		}
	}
	if v1, ok := d.GetOk("application_list"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("application_list", sv)
				diags = append(diags, e)
			}
			obj.ApplicationList = &v2
		}
	}
	if v1, ok := d.GetOk("atf_weight"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("atf_weight", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AtfWeight = &tmp
		}
	}
	if v1, ok := d.GetOk("auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth", sv)
				diags = append(diags, e)
			}
			obj.Auth = &v2
		}
	}
	if v1, ok := d.GetOk("auth_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.4", "") {
				e := utils.AttributeVersionWarning("auth_cert", sv)
				diags = append(diags, e)
			}
			obj.AuthCert = &v2
		}
	}
	if v1, ok := d.GetOk("auth_portal_addr"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.4", "") {
				e := utils.AttributeVersionWarning("auth_portal_addr", sv)
				diags = append(diags, e)
			}
			obj.AuthPortalAddr = &v2
		}
	}
	if v1, ok := d.GetOk("beacon_advertising"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("beacon_advertising", sv)
				diags = append(diags, e)
			}
			obj.BeaconAdvertising = &v2
		}
	}
	if v1, ok := d.GetOk("broadcast_ssid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("broadcast_ssid", sv)
				diags = append(diags, e)
			}
			obj.BroadcastSsid = &v2
		}
	}
	if v1, ok := d.GetOk("broadcast_suppression"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("broadcast_suppression", sv)
				diags = append(diags, e)
			}
			obj.BroadcastSuppression = &v2
		}
	}
	if v1, ok := d.GetOk("bss_color_partial"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("bss_color_partial", sv)
				diags = append(diags, e)
			}
			obj.BssColorPartial = &v2
		}
	}
	if v1, ok := d.GetOk("bstm_disassociation_imminent"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("bstm_disassociation_imminent", sv)
				diags = append(diags, e)
			}
			obj.BstmDisassociationImminent = &v2
		}
	}
	if v1, ok := d.GetOk("bstm_load_balancing_disassoc_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("bstm_load_balancing_disassoc_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.BstmLoadBalancingDisassocTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("bstm_rssi_disassoc_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("bstm_rssi_disassoc_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.BstmRssiDisassocTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("captive_portal_ac_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("captive_portal_ac_name", sv)
				diags = append(diags, e)
			}
			obj.CaptivePortalAcName = &v2
		}
	}
	if v1, ok := d.GetOk("captive_portal_auth_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("captive_portal_auth_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CaptivePortalAuthTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("captive_portal_macauth_radius_secret"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("captive_portal_macauth_radius_secret", sv)
				diags = append(diags, e)
			}
			obj.CaptivePortalMacauthRadiusSecret = &v2
		}
	}
	if v1, ok := d.GetOk("captive_portal_macauth_radius_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("captive_portal_macauth_radius_server", sv)
				diags = append(diags, e)
			}
			obj.CaptivePortalMacauthRadiusServer = &v2
		}
	}
	if v1, ok := d.GetOk("captive_portal_radius_secret"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("captive_portal_radius_secret", sv)
				diags = append(diags, e)
			}
			obj.CaptivePortalRadiusSecret = &v2
		}
	}
	if v1, ok := d.GetOk("captive_portal_radius_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("captive_portal_radius_server", sv)
				diags = append(diags, e)
			}
			obj.CaptivePortalRadiusServer = &v2
		}
	}
	if v1, ok := d.GetOk("captive_portal_session_timeout_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("captive_portal_session_timeout_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CaptivePortalSessionTimeoutInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("dhcp_address_enforcement"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("dhcp_address_enforcement", sv)
				diags = append(diags, e)
			}
			obj.DhcpAddressEnforcement = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_lease_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp_lease_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DhcpLeaseTime = &tmp
		}
	}
	if v1, ok := d.GetOk("dhcp_option43_insertion"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("dhcp_option43_insertion", sv)
				diags = append(diags, e)
			}
			obj.DhcpOption43Insertion = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_option82_circuit_id_insertion"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp_option82_circuit_id_insertion", sv)
				diags = append(diags, e)
			}
			obj.DhcpOption82CircuitIdInsertion = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_option82_insertion"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp_option82_insertion", sv)
				diags = append(diags, e)
			}
			obj.DhcpOption82Insertion = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_option82_remote_id_insertion"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp_option82_remote_id_insertion", sv)
				diags = append(diags, e)
			}
			obj.DhcpOption82RemoteIdInsertion = &v2
		}
	}
	if v1, ok := d.GetOk("dynamic_vlan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dynamic_vlan", sv)
				diags = append(diags, e)
			}
			obj.DynamicVlan = &v2
		}
	}
	if v1, ok := d.GetOk("eap_reauth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eap_reauth", sv)
				diags = append(diags, e)
			}
			obj.EapReauth = &v2
		}
	}
	if v1, ok := d.GetOk("eap_reauth_intv"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eap_reauth_intv", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EapReauthIntv = &tmp
		}
	}
	if v1, ok := d.GetOk("eapol_key_retries"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eapol_key_retries", sv)
				diags = append(diags, e)
			}
			obj.EapolKeyRetries = &v2
		}
	}
	if v1, ok := d.GetOk("encrypt"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("encrypt", sv)
				diags = append(diags, e)
			}
			obj.Encrypt = &v2
		}
	}
	if v1, ok := d.GetOk("external_fast_roaming"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("external_fast_roaming", sv)
				diags = append(diags, e)
			}
			obj.ExternalFastRoaming = &v2
		}
	}
	if v1, ok := d.GetOk("external_logout"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("external_logout", sv)
				diags = append(diags, e)
			}
			obj.ExternalLogout = &v2
		}
	}
	if v1, ok := d.GetOk("external_web"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("external_web", sv)
				diags = append(diags, e)
			}
			obj.ExternalWeb = &v2
		}
	}
	if v1, ok := d.GetOk("external_web_format"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("external_web_format", sv)
				diags = append(diags, e)
			}
			obj.ExternalWebFormat = &v2
		}
	}
	if v1, ok := d.GetOk("fast_bss_transition"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fast_bss_transition", sv)
				diags = append(diags, e)
			}
			obj.FastBssTransition = &v2
		}
	}
	if v1, ok := d.GetOk("fast_roaming"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fast_roaming", sv)
				diags = append(diags, e)
			}
			obj.FastRoaming = &v2
		}
	}
	if v1, ok := d.GetOk("ft_mobility_domain"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ft_mobility_domain", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FtMobilityDomain = &tmp
		}
	}
	if v1, ok := d.GetOk("ft_over_ds"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ft_over_ds", sv)
				diags = append(diags, e)
			}
			obj.FtOverDs = &v2
		}
	}
	if v1, ok := d.GetOk("ft_r0_key_lifetime"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ft_r0_key_lifetime", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FtR0KeyLifetime = &tmp
		}
	}
	if v1, ok := d.GetOk("gas_comeback_delay"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("gas_comeback_delay", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.GasComebackDelay = &tmp
		}
	}
	if v1, ok := d.GetOk("gas_fragmentation_limit"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("gas_fragmentation_limit", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.GasFragmentationLimit = &tmp
		}
	}
	if v1, ok := d.GetOk("gtk_rekey"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gtk_rekey", sv)
				diags = append(diags, e)
			}
			obj.GtkRekey = &v2
		}
	}
	if v1, ok := d.GetOk("gtk_rekey_intv"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gtk_rekey_intv", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.GtkRekeyIntv = &tmp
		}
	}
	if v1, ok := d.GetOk("high_efficiency"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("high_efficiency", sv)
				diags = append(diags, e)
			}
			obj.HighEfficiency = &v2
		}
	}
	if v1, ok := d.GetOk("hotspot20_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hotspot20_profile", sv)
				diags = append(diags, e)
			}
			obj.Hotspot20Profile = &v2
		}
	}
	if v1, ok := d.GetOk("igmp_snooping"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("igmp_snooping", sv)
				diags = append(diags, e)
			}
			obj.IgmpSnooping = &v2
		}
	}
	if v1, ok := d.GetOk("intra_vap_privacy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("intra_vap_privacy", sv)
				diags = append(diags, e)
			}
			obj.IntraVapPrivacy = &v2
		}
	}
	if v1, ok := d.GetOk("ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip", sv)
				diags = append(diags, e)
			}
			obj.Ip = &v2
		}
	}
	if v1, ok := d.GetOk("ips_sensor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("ips_sensor", sv)
				diags = append(diags, e)
			}
			obj.IpsSensor = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6_rules"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("ipv6_rules", sv)
				diags = append(diags, e)
			}
			obj.Ipv6Rules = &v2
		}
	}
	if v1, ok := d.GetOk("key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("key", sv)
				diags = append(diags, e)
			}
			obj.Key = &v2
		}
	}
	if v1, ok := d.GetOk("keyindex"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("keyindex", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Keyindex = &tmp
		}
	}
	if v1, ok := d.GetOk("ldpc"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ldpc", sv)
				diags = append(diags, e)
			}
			obj.Ldpc = &v2
		}
	}
	if v1, ok := d.GetOk("local_authentication"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_authentication", sv)
				diags = append(diags, e)
			}
			obj.LocalAuthentication = &v2
		}
	}
	if v1, ok := d.GetOk("local_bridging"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_bridging", sv)
				diags = append(diags, e)
			}
			obj.LocalBridging = &v2
		}
	}
	if v1, ok := d.GetOk("local_lan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_lan", sv)
				diags = append(diags, e)
			}
			obj.LocalLan = &v2
		}
	}
	if v1, ok := d.GetOk("local_standalone"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_standalone", sv)
				diags = append(diags, e)
			}
			obj.LocalStandalone = &v2
		}
	}
	if v1, ok := d.GetOk("local_standalone_dns"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("local_standalone_dns", sv)
				diags = append(diags, e)
			}
			obj.LocalStandaloneDns = &v2
		}
	}
	if v1, ok := d.GetOk("local_standalone_dns_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("local_standalone_dns_ip", sv)
				diags = append(diags, e)
			}
			obj.LocalStandaloneDnsIp = &v2
		}
	}
	if v1, ok := d.GetOk("local_standalone_nat"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_standalone_nat", sv)
				diags = append(diags, e)
			}
			obj.LocalStandaloneNat = &v2
		}
	}
	if v1, ok := d.GetOk("mac_auth_bypass"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mac_auth_bypass", sv)
				diags = append(diags, e)
			}
			obj.MacAuthBypass = &v2
		}
	}
	if v1, ok := d.GetOk("mac_called_station_delimiter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("mac_called_station_delimiter", sv)
				diags = append(diags, e)
			}
			obj.MacCalledStationDelimiter = &v2
		}
	}
	if v1, ok := d.GetOk("mac_calling_station_delimiter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("mac_calling_station_delimiter", sv)
				diags = append(diags, e)
			}
			obj.MacCallingStationDelimiter = &v2
		}
	}
	if v1, ok := d.GetOk("mac_case"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("mac_case", sv)
				diags = append(diags, e)
			}
			obj.MacCase = &v2
		}
	}
	if v1, ok := d.GetOk("mac_filter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mac_filter", sv)
				diags = append(diags, e)
			}
			obj.MacFilter = &v2
		}
	}
	if v, ok := d.GetOk("mac_filter_list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("mac_filter_list", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerVapMacFilterList(d, v, "mac_filter_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.MacFilterList = t
		}
	} else if d.HasChange("mac_filter_list") {
		old, new := d.GetChange("mac_filter_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.MacFilterList = &[]models.WirelessControllerVapMacFilterList{}
		}
	}
	if v1, ok := d.GetOk("mac_filter_policy_other"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mac_filter_policy_other", sv)
				diags = append(diags, e)
			}
			obj.MacFilterPolicyOther = &v2
		}
	}
	if v1, ok := d.GetOk("mac_password_delimiter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("mac_password_delimiter", sv)
				diags = append(diags, e)
			}
			obj.MacPasswordDelimiter = &v2
		}
	}
	if v1, ok := d.GetOk("mac_username_delimiter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("mac_username_delimiter", sv)
				diags = append(diags, e)
			}
			obj.MacUsernameDelimiter = &v2
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
	if v1, ok := d.GetOk("max_clients_ap"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_clients_ap", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxClientsAp = &tmp
		}
	}
	if v1, ok := d.GetOk("mbo"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("mbo", sv)
				diags = append(diags, e)
			}
			obj.Mbo = &v2
		}
	}
	if v1, ok := d.GetOk("mbo_cell_data_conn_pref"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("mbo_cell_data_conn_pref", sv)
				diags = append(diags, e)
			}
			obj.MboCellDataConnPref = &v2
		}
	}
	if v1, ok := d.GetOk("me_disable_thresh"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("me_disable_thresh", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MeDisableThresh = &tmp
		}
	}
	if v1, ok := d.GetOk("mesh_backhaul"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mesh_backhaul", sv)
				diags = append(diags, e)
			}
			obj.MeshBackhaul = &v2
		}
	}
	if v1, ok := d.GetOk("mpsk"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("mpsk", sv)
				diags = append(diags, e)
			}
			obj.Mpsk = &v2
		}
	}
	if v1, ok := d.GetOk("mpsk_concurrent_clients"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("mpsk_concurrent_clients", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MpskConcurrentClients = &tmp
		}
	}
	if v, ok := d.GetOk("mpsk_key"); ok {
		if !utils.CheckVer(sv, "", "v6.4.2") {
			e := utils.AttributeVersionWarning("mpsk_key", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerVapMpskKey(d, v, "mpsk_key", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.MpskKey = t
		}
	} else if d.HasChange("mpsk_key") {
		old, new := d.GetChange("mpsk_key")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.MpskKey = &[]models.WirelessControllerVapMpskKey{}
		}
	}
	if v1, ok := d.GetOk("mpsk_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("mpsk_profile", sv)
				diags = append(diags, e)
			}
			obj.MpskProfile = &v2
		}
	}
	if v1, ok := d.GetOk("mu_mimo"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mu_mimo", sv)
				diags = append(diags, e)
			}
			obj.MuMimo = &v2
		}
	}
	if v1, ok := d.GetOk("multicast_enhance"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("multicast_enhance", sv)
				diags = append(diags, e)
			}
			obj.MulticastEnhance = &v2
		}
	}
	if v1, ok := d.GetOk("multicast_rate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("multicast_rate", sv)
				diags = append(diags, e)
			}
			obj.MulticastRate = &v2
		}
	}
	if v1, ok := d.GetOk("nac"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("nac", sv)
				diags = append(diags, e)
			}
			obj.Nac = &v2
		}
	}
	if v1, ok := d.GetOk("nac_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("nac_profile", sv)
				diags = append(diags, e)
			}
			obj.NacProfile = &v2
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
	if v1, ok := d.GetOk("neighbor_report_dual_band"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("neighbor_report_dual_band", sv)
				diags = append(diags, e)
			}
			obj.NeighborReportDualBand = &v2
		}
	}
	if v1, ok := d.GetOk("okc"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("okc", sv)
				diags = append(diags, e)
			}
			obj.Okc = &v2
		}
	}
	if v1, ok := d.GetOk("osen"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("osen", sv)
				diags = append(diags, e)
			}
			obj.Osen = &v2
		}
	}
	if v1, ok := d.GetOk("owe_groups"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("owe_groups", sv)
				diags = append(diags, e)
			}
			obj.OweGroups = &v2
		}
	}
	if v1, ok := d.GetOk("owe_transition"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("owe_transition", sv)
				diags = append(diags, e)
			}
			obj.OweTransition = &v2
		}
	}
	if v1, ok := d.GetOk("owe_transition_ssid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("owe_transition_ssid", sv)
				diags = append(diags, e)
			}
			obj.OweTransitionSsid = &v2
		}
	}
	if v1, ok := d.GetOk("passphrase"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("passphrase", sv)
				diags = append(diags, e)
			}
			obj.Passphrase = &v2
		}
	}
	if v1, ok := d.GetOk("pmf"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pmf", sv)
				diags = append(diags, e)
			}
			obj.Pmf = &v2
		}
	}
	if v1, ok := d.GetOk("pmf_assoc_comeback_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pmf_assoc_comeback_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PmfAssocComebackTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("pmf_sa_query_retry_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pmf_sa_query_retry_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PmfSaQueryRetryTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("port_macauth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("port_macauth", sv)
				diags = append(diags, e)
			}
			obj.PortMacauth = &v2
		}
	}
	if v1, ok := d.GetOk("port_macauth_reauth_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("port_macauth_reauth_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PortMacauthReauthTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("port_macauth_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("port_macauth_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PortMacauthTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("portal_message_override_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("portal_message_override_group", sv)
				diags = append(diags, e)
			}
			obj.PortalMessageOverrideGroup = &v2
		}
	}
	if v, ok := d.GetOk("portal_message_overrides"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("portal_message_overrides", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerVapPortalMessageOverrides(d, v, "portal_message_overrides", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.PortalMessageOverrides = t
		}
	} else if d.HasChange("portal_message_overrides") {
		old, new := d.GetChange("portal_message_overrides")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.PortalMessageOverrides = &[]models.WirelessControllerVapPortalMessageOverrides{}
		}
	}
	if v1, ok := d.GetOk("portal_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("portal_type", sv)
				diags = append(diags, e)
			}
			obj.PortalType = &v2
		}
	}
	if v1, ok := d.GetOk("primary_wag_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("primary_wag_profile", sv)
				diags = append(diags, e)
			}
			obj.PrimaryWagProfile = &v2
		}
	}
	if v1, ok := d.GetOk("probe_resp_suppression"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("probe_resp_suppression", sv)
				diags = append(diags, e)
			}
			obj.ProbeRespSuppression = &v2
		}
	}
	if v1, ok := d.GetOk("probe_resp_threshold"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("probe_resp_threshold", sv)
				diags = append(diags, e)
			}
			obj.ProbeRespThreshold = &v2
		}
	}
	if v1, ok := d.GetOk("ptk_rekey"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ptk_rekey", sv)
				diags = append(diags, e)
			}
			obj.PtkRekey = &v2
		}
	}
	if v1, ok := d.GetOk("ptk_rekey_intv"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ptk_rekey_intv", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PtkRekeyIntv = &tmp
		}
	}
	if v1, ok := d.GetOk("qos_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("qos_profile", sv)
				diags = append(diags, e)
			}
			obj.QosProfile = &v2
		}
	}
	if v1, ok := d.GetOk("quarantine"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("quarantine", sv)
				diags = append(diags, e)
			}
			obj.Quarantine = &v2
		}
	}
	if v1, ok := d.GetOk("radio_2g_threshold"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("radio_2g_threshold", sv)
				diags = append(diags, e)
			}
			obj.Radio2gThreshold = &v2
		}
	}
	if v1, ok := d.GetOk("radio_5g_threshold"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("radio_5g_threshold", sv)
				diags = append(diags, e)
			}
			obj.Radio5gThreshold = &v2
		}
	}
	if v1, ok := d.GetOk("radio_sensitivity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("radio_sensitivity", sv)
				diags = append(diags, e)
			}
			obj.RadioSensitivity = &v2
		}
	}
	if v1, ok := d.GetOk("radius_mac_auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("radius_mac_auth", sv)
				diags = append(diags, e)
			}
			obj.RadiusMacAuth = &v2
		}
	}
	if v1, ok := d.GetOk("radius_mac_auth_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("radius_mac_auth_server", sv)
				diags = append(diags, e)
			}
			obj.RadiusMacAuthServer = &v2
		}
	}
	if v, ok := d.GetOk("radius_mac_auth_usergroups"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("radius_mac_auth_usergroups", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerVapRadiusMacAuthUsergroups(d, v, "radius_mac_auth_usergroups", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.RadiusMacAuthUsergroups = t
		}
	} else if d.HasChange("radius_mac_auth_usergroups") {
		old, new := d.GetChange("radius_mac_auth_usergroups")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.RadiusMacAuthUsergroups = &[]models.WirelessControllerVapRadiusMacAuthUsergroups{}
		}
	}
	if v1, ok := d.GetOk("radius_mac_mpsk_auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("radius_mac_mpsk_auth", sv)
				diags = append(diags, e)
			}
			obj.RadiusMacMpskAuth = &v2
		}
	}
	if v1, ok := d.GetOk("radius_mac_mpsk_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("radius_mac_mpsk_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RadiusMacMpskTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("radius_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("radius_server", sv)
				diags = append(diags, e)
			}
			obj.RadiusServer = &v2
		}
	}
	if v1, ok := d.GetOk("rates_11a"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rates_11a", sv)
				diags = append(diags, e)
			}
			obj.Rates11a = &v2
		}
	}
	if v1, ok := d.GetOk("rates_11ac_ss12"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rates_11ac_ss12", sv)
				diags = append(diags, e)
			}
			obj.Rates11acSs12 = &v2
		}
	}
	if v1, ok := d.GetOk("rates_11ac_ss34"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rates_11ac_ss34", sv)
				diags = append(diags, e)
			}
			obj.Rates11acSs34 = &v2
		}
	}
	if v1, ok := d.GetOk("rates_11ax_ss12"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("rates_11ax_ss12", sv)
				diags = append(diags, e)
			}
			obj.Rates11axSs12 = &v2
		}
	}
	if v1, ok := d.GetOk("rates_11ax_ss34"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("rates_11ax_ss34", sv)
				diags = append(diags, e)
			}
			obj.Rates11axSs34 = &v2
		}
	}
	if v1, ok := d.GetOk("rates_11bg"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rates_11bg", sv)
				diags = append(diags, e)
			}
			obj.Rates11bg = &v2
		}
	}
	if v1, ok := d.GetOk("rates_11n_ss12"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rates_11n_ss12", sv)
				diags = append(diags, e)
			}
			obj.Rates11nSs12 = &v2
		}
	}
	if v1, ok := d.GetOk("rates_11n_ss34"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rates_11n_ss34", sv)
				diags = append(diags, e)
			}
			obj.Rates11nSs34 = &v2
		}
	}
	if v1, ok := d.GetOk("sae_groups"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sae_groups", sv)
				diags = append(diags, e)
			}
			obj.SaeGroups = &v2
		}
	}
	if v1, ok := d.GetOk("sae_password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sae_password", sv)
				diags = append(diags, e)
			}
			obj.SaePassword = &v2
		}
	}
	if v1, ok := d.GetOk("scan_botnet_connections"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("scan_botnet_connections", sv)
				diags = append(diags, e)
			}
			obj.ScanBotnetConnections = &v2
		}
	}
	if v, ok := d.GetOk("schedule"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("schedule", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerVapSchedule(d, v, "schedule", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Schedule = t
		}
	} else if d.HasChange("schedule") {
		old, new := d.GetChange("schedule")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Schedule = &[]models.WirelessControllerVapSchedule{}
		}
	}
	if v1, ok := d.GetOk("secondary_wag_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("secondary_wag_profile", sv)
				diags = append(diags, e)
			}
			obj.SecondaryWagProfile = &v2
		}
	}
	if v1, ok := d.GetOk("security"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("security", sv)
				diags = append(diags, e)
			}
			obj.Security = &v2
		}
	}
	if v1, ok := d.GetOk("security_exempt_list"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("security_exempt_list", sv)
				diags = append(diags, e)
			}
			obj.SecurityExemptList = &v2
		}
	}
	if v1, ok := d.GetOk("security_redirect_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("security_redirect_url", sv)
				diags = append(diags, e)
			}
			obj.SecurityRedirectUrl = &v2
		}
	}
	if v, ok := d.GetOk("selected_usergroups"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("selected_usergroups", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerVapSelectedUsergroups(d, v, "selected_usergroups", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SelectedUsergroups = t
		}
	} else if d.HasChange("selected_usergroups") {
		old, new := d.GetChange("selected_usergroups")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SelectedUsergroups = &[]models.WirelessControllerVapSelectedUsergroups{}
		}
	}
	if v1, ok := d.GetOk("split_tunneling"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("split_tunneling", sv)
				diags = append(diags, e)
			}
			obj.SplitTunneling = &v2
		}
	}
	if v1, ok := d.GetOk("ssid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssid", sv)
				diags = append(diags, e)
			}
			obj.Ssid = &v2
		}
	}
	if v1, ok := d.GetOk("sticky_client_remove"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("sticky_client_remove", sv)
				diags = append(diags, e)
			}
			obj.StickyClientRemove = &v2
		}
	}
	if v1, ok := d.GetOk("sticky_client_threshold_2g"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("sticky_client_threshold_2g", sv)
				diags = append(diags, e)
			}
			obj.StickyClientThreshold2g = &v2
		}
	}
	if v1, ok := d.GetOk("sticky_client_threshold_5g"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("sticky_client_threshold_5g", sv)
				diags = append(diags, e)
			}
			obj.StickyClientThreshold5g = &v2
		}
	}
	if v1, ok := d.GetOk("target_wake_time"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("target_wake_time", sv)
				diags = append(diags, e)
			}
			obj.TargetWakeTime = &v2
		}
	}
	if v1, ok := d.GetOk("tkip_counter_measure"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tkip_counter_measure", sv)
				diags = append(diags, e)
			}
			obj.TkipCounterMeasure = &v2
		}
	}
	if v1, ok := d.GetOk("tunnel_echo_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tunnel_echo_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TunnelEchoInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("tunnel_fallback_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tunnel_fallback_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TunnelFallbackInterval = &tmp
		}
	}
	if v, ok := d.GetOk("usergroup"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("usergroup", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerVapUsergroup(d, v, "usergroup", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Usergroup = t
		}
	} else if d.HasChange("usergroup") {
		old, new := d.GetChange("usergroup")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Usergroup = &[]models.WirelessControllerVapUsergroup{}
		}
	}
	if v1, ok := d.GetOk("utm_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("utm_log", sv)
				diags = append(diags, e)
			}
			obj.UtmLog = &v2
		}
	}
	if v1, ok := d.GetOk("utm_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("utm_profile", sv)
				diags = append(diags, e)
			}
			obj.UtmProfile = &v2
		}
	}
	if v1, ok := d.GetOk("utm_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("utm_status", sv)
				diags = append(diags, e)
			}
			obj.UtmStatus = &v2
		}
	}
	if v1, ok := d.GetOk("vlan_auto"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vlan_auto", sv)
				diags = append(diags, e)
			}
			obj.VlanAuto = &v2
		}
	}
	if v, ok := d.GetOk("vlan_name"); ok {
		if !utils.CheckVer(sv, "v7.0.4", "") {
			e := utils.AttributeVersionWarning("vlan_name", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerVapVlanName(d, v, "vlan_name", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.VlanName = t
		}
	} else if d.HasChange("vlan_name") {
		old, new := d.GetChange("vlan_name")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.VlanName = &[]models.WirelessControllerVapVlanName{}
		}
	}
	if v, ok := d.GetOk("vlan_pool"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("vlan_pool", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerVapVlanPool(d, v, "vlan_pool", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.VlanPool = t
		}
	} else if d.HasChange("vlan_pool") {
		old, new := d.GetChange("vlan_pool")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.VlanPool = &[]models.WirelessControllerVapVlanPool{}
		}
	}
	if v1, ok := d.GetOk("vlan_pooling"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vlan_pooling", sv)
				diags = append(diags, e)
			}
			obj.VlanPooling = &v2
		}
	}
	if v1, ok := d.GetOk("vlanid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vlanid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Vlanid = &tmp
		}
	}
	if v1, ok := d.GetOk("voice_enterprise"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("voice_enterprise", sv)
				diags = append(diags, e)
			}
			obj.VoiceEnterprise = &v2
		}
	}
	if v1, ok := d.GetOk("webfilter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("webfilter_profile", sv)
				diags = append(diags, e)
			}
			obj.WebfilterProfile = &v2
		}
	}
	return &obj, diags
}
