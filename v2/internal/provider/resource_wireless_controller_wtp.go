// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Wireless Termination Points (WTPs), that is, FortiAPs or APs to be managed by FortiGate.

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

func resourceWirelessControllerWtp() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Wireless Termination Points (WTPs), that is, FortiAPs or APs to be managed by FortiGate.",

		CreateContext: resourceWirelessControllerWtpCreate,
		ReadContext:   resourceWirelessControllerWtpRead,
		UpdateContext: resourceWirelessControllerWtpUpdate,
		DeleteContext: resourceWirelessControllerWtpDelete,

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
			"admin": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"discovered", "disable", "enable"}, false),

				Description: "Configure how the FortiGate operating as a wireless controller discovers and manages this WTP, AP or FortiAP.",
				Optional:    true,
				Computed:    true,
			},
			"allowaccess": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Control management access to the managed WTP, FortiAP, or AP. Separate entries with a space.",
				Optional:         true,
				Computed:         true,
			},
			"apcfg_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "AP local configuration profile name.",
				Optional:    true,
				Computed:    true,
			},
			"bonjour_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Bonjour profile name.",
				Optional:    true,
				Computed:    true,
			},
			"coordinate_latitude": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),

				Description: "WTP latitude coordinate.",
				Optional:    true,
				Computed:    true,
			},
			"coordinate_longitude": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),

				Description: "WTP longitude coordinate.",
				Optional:    true,
				Computed:    true,
			},
			"firmware_provision": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Firmware version to provision to this FortiAP on bootup (major.minor.build, i.e. 6.2.1234).",
				Optional:    true,
				Computed:    true,
			},
			"firmware_provision_latest": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "once"}, false),

				Description: "Enable/disable one-time automatic provisioning of the latest firmware version.",
				Optional:    true,
				Computed:    true,
			},
			"image_download": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable WTP image download.",
				Optional:    true,
				Computed:    true,
			},
			"index": {
				Type: schema.TypeInt,

				Description: "Index (0 - 4294967295).",
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
			"led_state": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to allow the FortiAPs LEDs to light. Disable to keep the LEDs off. You may want to keep the LEDs off so they are not distracting in low light areas etc.",
				Optional:    true,
				Computed:    true,
			},
			"location": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Field for describing the physical location of the WTP, AP or FortiAP.",
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
			"mesh_bridge_enable": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"default", "enable", "disable"}, false),

				Description: "Enable/disable mesh Ethernet bridge when WTP is configured as a mesh branch/leaf AP.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "WTP, AP or FortiAP configuration name.",
				Optional:    true,
				Computed:    true,
			},
			"override_allowaccess": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to override the WTP profile management access configuration.",
				Optional:    true,
				Computed:    true,
			},
			"override_ip_fragment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable overriding the WTP profile IP fragment prevention setting.",
				Optional:    true,
				Computed:    true,
			},
			"override_lan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to override the WTP profile LAN port setting.",
				Optional:    true,
				Computed:    true,
			},
			"override_led_state": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to override the profile LED state setting for this FortiAP. You must enable this option to use the led-state command to turn off the FortiAP's LEDs.",
				Optional:    true,
				Computed:    true,
			},
			"override_login_passwd_change": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to override the WTP profile login-password (administrator password) setting.",
				Optional:    true,
				Computed:    true,
			},
			"override_split_tunnel": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable overriding the WTP profile split tunneling setting.",
				Optional:    true,
				Computed:    true,
			},
			"override_wan_port_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable overriding the wan-port-mode in the WTP profile.",
				Optional:    true,
				Computed:    true,
			},
			"radio_1": {
				Type:        schema.TypeList,
				Description: "Configuration options for radio 1.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
							ValidateFunc: validation.StringInSlice([]string{"802.11a", "802.11b", "802.11g", "802.11n", "802.11n-5G", "802.11ac", "802.11ax-5G", "802.11ax", "802.11ac-2G", "802.11n,g-only", "802.11g-only", "802.11n-only", "802.11n-5G-only", "802.11ac,n-only", "802.11ac-only", "802.11ax,ac-only", "802.11ax,ac,n-only", "802.11ax-5G-only", "802.11ax,n-only", "802.11ax,n,g-only", "802.11ax-only"}, false),

							Description: "WiFi band that Radio 1 operates on.",
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
						"drma_manual_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ap", "monitor", "ncf", "ncf-peek"}, false),

							Description: "Radio mode to be used for DRMA manual mode (default = ncf).",
							Optional:    true,
							Computed:    true,
						},
						"override_analysis": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to override the WTP profile spectrum analysis configuration.",
							Optional:    true,
							Computed:    true,
						},
						"override_band": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to override the WTP profile band setting.",
							Optional:    true,
							Computed:    true,
						},
						"override_channel": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to override WTP profile channel settings.",
							Optional:    true,
							Computed:    true,
						},
						"override_txpower": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to override the WTP profile power level configuration.",
							Optional:    true,
							Computed:    true,
						},
						"override_vaps": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to override WTP profile Virtual Access Point (VAP) settings.",
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
						"spectrum_analysis": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "scan-only", "disable"}, false),

							Description: "Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.",
							Optional:    true,
							Computed:    true,
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
					},
				},
			},
			"radio_2": {
				Type:        schema.TypeList,
				Description: "Configuration options for radio 2.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
							ValidateFunc: validation.StringInSlice([]string{"802.11a", "802.11b", "802.11g", "802.11n", "802.11n-5G", "802.11ac", "802.11ax-5G", "802.11ax", "802.11ac-2G", "802.11n,g-only", "802.11g-only", "802.11n-only", "802.11n-5G-only", "802.11ac,n-only", "802.11ac-only", "802.11ax,ac-only", "802.11ax,ac,n-only", "802.11ax-5G-only", "802.11ax,n-only", "802.11ax,n,g-only", "802.11ax-only"}, false),

							Description: "WiFi band that Radio 2 operates on.",
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
						"drma_manual_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ap", "monitor", "ncf", "ncf-peek"}, false),

							Description: "Radio mode to be used for DRMA manual mode (default = ncf).",
							Optional:    true,
							Computed:    true,
						},
						"override_analysis": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to override the WTP profile spectrum analysis configuration.",
							Optional:    true,
							Computed:    true,
						},
						"override_band": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to override the WTP profile band setting.",
							Optional:    true,
							Computed:    true,
						},
						"override_channel": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to override WTP profile channel settings.",
							Optional:    true,
							Computed:    true,
						},
						"override_txpower": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to override the WTP profile power level configuration.",
							Optional:    true,
							Computed:    true,
						},
						"override_vaps": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to override WTP profile Virtual Access Point (VAP) settings.",
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
						"spectrum_analysis": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "scan-only", "disable"}, false),

							Description: "Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.",
							Optional:    true,
							Computed:    true,
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
					},
				},
			},
			"radio_3": {
				Type:        schema.TypeList,
				Description: "Configuration options for radio 3.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
							ValidateFunc: validation.StringInSlice([]string{"802.11a", "802.11b", "802.11g", "802.11n", "802.11n-5G", "802.11ac", "802.11ax-5G", "802.11ax", "802.11ac-2G", "802.11n,g-only", "802.11g-only", "802.11n-only", "802.11n-5G-only", "802.11ac,n-only", "802.11ac-only", "802.11ax,ac-only", "802.11ax,ac,n-only", "802.11ax-5G-only", "802.11ax,n-only", "802.11ax,n,g-only", "802.11ax-only"}, false),

							Description: "WiFi band that Radio 3 operates on.",
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
						"drma_manual_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ap", "monitor", "ncf", "ncf-peek"}, false),

							Description: "Radio mode to be used for DRMA manual mode (default = ncf).",
							Optional:    true,
							Computed:    true,
						},
						"override_analysis": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to override the WTP profile spectrum analysis configuration.",
							Optional:    true,
							Computed:    true,
						},
						"override_band": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to override the WTP profile band setting.",
							Optional:    true,
							Computed:    true,
						},
						"override_channel": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to override WTP profile channel settings.",
							Optional:    true,
							Computed:    true,
						},
						"override_txpower": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to override the WTP profile power level configuration.",
							Optional:    true,
							Computed:    true,
						},
						"override_vaps": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to override WTP profile Virtual Access Point (VAP) settings.",
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
						"spectrum_analysis": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "scan-only", "disable"}, false),

							Description: "Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.",
							Optional:    true,
							Computed:    true,
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
					},
				},
			},
			"radio_4": {
				Type:        schema.TypeList,
				Description: "Configuration options for radio 4.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
							ValidateFunc: validation.StringInSlice([]string{"802.11a", "802.11b", "802.11g", "802.11n", "802.11n-5G", "802.11ac", "802.11ax-5G", "802.11ax", "802.11ac-2G", "802.11n,g-only", "802.11g-only", "802.11n-only", "802.11n-5G-only", "802.11ac,n-only", "802.11ac-only", "802.11ax,ac-only", "802.11ax,ac,n-only", "802.11ax-5G-only", "802.11ax,n-only", "802.11ax,n,g-only", "802.11ax-only"}, false),

							Description: "WiFi band that Radio 4 operates on.",
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
						"drma_manual_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ap", "monitor", "ncf", "ncf-peek"}, false),

							Description: "Radio mode to be used for DRMA manual mode (default = ncf).",
							Optional:    true,
							Computed:    true,
						},
						"override_analysis": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to override the WTP profile spectrum analysis configuration.",
							Optional:    true,
							Computed:    true,
						},
						"override_band": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to override the WTP profile band setting.",
							Optional:    true,
							Computed:    true,
						},
						"override_channel": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to override WTP profile channel settings.",
							Optional:    true,
							Computed:    true,
						},
						"override_txpower": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to override the WTP profile power level configuration.",
							Optional:    true,
							Computed:    true,
						},
						"override_vaps": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to override WTP profile Virtual Access Point (VAP) settings.",
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
						"spectrum_analysis": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "scan-only", "disable"}, false),

							Description: "Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.",
							Optional:    true,
							Computed:    true,
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
					},
				},
			},
			"region": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Region name WTP is associated with.",
				Optional:    true,
				Computed:    true,
			},
			"region_x": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Relative horizontal region coordinate (between 0 and 1).",
				Optional:    true,
				Computed:    true,
			},
			"region_y": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Relative vertical region coordinate (between 0 and 1).",
				Optional:    true,
				Computed:    true,
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
			"uuid": {
				Type: schema.TypeString,

				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Optional:    true,
				Computed:    true,
			},
			"wan_port_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"wan-lan", "wan-only"}, false),

				Description: "Enable/disable using the FortiAP WAN port as a LAN port.",
				Optional:    true,
				Computed:    true,
			},
			"wtp_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "WTP ID.",
				ForceNew:    true,
				Required:    true,
			},
			"wtp_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"normal", "remote"}, false),

				Description: "WTP, AP, or FortiAP operating mode; normal (by default) or remote. A tunnel mode SSID can be assigned to an AP in normal mode but not remote mode, while a local-bridge mode SSID can be assigned to an AP in either normal mode or remote mode.",
				Optional:    true,
				Computed:    true,
			},
			"wtp_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "WTP profile name to apply to this WTP, AP or FortiAP.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWirelessControllerWtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "wtp_id"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating WirelessControllerWtp resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerWtp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerWtp(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerWtp")
	}

	return resourceWirelessControllerWtpRead(ctx, d, meta)
}

func resourceWirelessControllerWtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerWtp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerWtp(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerWtp resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerWtp")
	}

	return resourceWirelessControllerWtpRead(ctx, d, meta)
}

func resourceWirelessControllerWtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerWtp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerWtp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerWtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerWtp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerWtp resource: %v", err)
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

	diags := refreshObjectWirelessControllerWtp(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWirelessControllerWtpLan(d *schema.ResourceData, v *models.WirelessControllerWtpLan, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.WirelessControllerWtpLan{*v}
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

func flattenWirelessControllerWtpRadio1(d *schema.ResourceData, v *models.WirelessControllerWtpRadio1, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.WirelessControllerWtpRadio1{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
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

			if tmp := cfg.Channel; tmp != nil {
				v["channel"] = flattenWirelessControllerWtpRadio1Channel(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "channel"), sort)
			}

			if tmp := cfg.DrmaManualMode; tmp != nil {
				v["drma_manual_mode"] = *tmp
			}

			if tmp := cfg.OverrideAnalysis; tmp != nil {
				v["override_analysis"] = *tmp
			}

			if tmp := cfg.OverrideBand; tmp != nil {
				v["override_band"] = *tmp
			}

			if tmp := cfg.OverrideChannel; tmp != nil {
				v["override_channel"] = *tmp
			}

			if tmp := cfg.OverrideTxpower; tmp != nil {
				v["override_txpower"] = *tmp
			}

			if tmp := cfg.OverrideVaps; tmp != nil {
				v["override_vaps"] = *tmp
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

			if tmp := cfg.SpectrumAnalysis; tmp != nil {
				v["spectrum_analysis"] = *tmp
			}

			if tmp := cfg.VapAll; tmp != nil {
				v["vap_all"] = *tmp
			}

			if tmp := cfg.Vaps; tmp != nil {
				v["vaps"] = flattenWirelessControllerWtpRadio1Vaps(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "vaps"), sort)
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWirelessControllerWtpRadio1Channel(d *schema.ResourceData, v *[]models.WirelessControllerWtpRadio1Channel, prefix string, sort bool) interface{} {
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

func flattenWirelessControllerWtpRadio1Vaps(d *schema.ResourceData, v *[]models.WirelessControllerWtpRadio1Vaps, prefix string, sort bool) interface{} {
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

func flattenWirelessControllerWtpRadio2(d *schema.ResourceData, v *models.WirelessControllerWtpRadio2, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.WirelessControllerWtpRadio2{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
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

			if tmp := cfg.Channel; tmp != nil {
				v["channel"] = flattenWirelessControllerWtpRadio2Channel(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "channel"), sort)
			}

			if tmp := cfg.DrmaManualMode; tmp != nil {
				v["drma_manual_mode"] = *tmp
			}

			if tmp := cfg.OverrideAnalysis; tmp != nil {
				v["override_analysis"] = *tmp
			}

			if tmp := cfg.OverrideBand; tmp != nil {
				v["override_band"] = *tmp
			}

			if tmp := cfg.OverrideChannel; tmp != nil {
				v["override_channel"] = *tmp
			}

			if tmp := cfg.OverrideTxpower; tmp != nil {
				v["override_txpower"] = *tmp
			}

			if tmp := cfg.OverrideVaps; tmp != nil {
				v["override_vaps"] = *tmp
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

			if tmp := cfg.SpectrumAnalysis; tmp != nil {
				v["spectrum_analysis"] = *tmp
			}

			if tmp := cfg.VapAll; tmp != nil {
				v["vap_all"] = *tmp
			}

			if tmp := cfg.Vaps; tmp != nil {
				v["vaps"] = flattenWirelessControllerWtpRadio2Vaps(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "vaps"), sort)
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWirelessControllerWtpRadio2Channel(d *schema.ResourceData, v *[]models.WirelessControllerWtpRadio2Channel, prefix string, sort bool) interface{} {
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

func flattenWirelessControllerWtpRadio2Vaps(d *schema.ResourceData, v *[]models.WirelessControllerWtpRadio2Vaps, prefix string, sort bool) interface{} {
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

func flattenWirelessControllerWtpRadio3(d *schema.ResourceData, v *models.WirelessControllerWtpRadio3, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.WirelessControllerWtpRadio3{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
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

			if tmp := cfg.Channel; tmp != nil {
				v["channel"] = flattenWirelessControllerWtpRadio3Channel(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "channel"), sort)
			}

			if tmp := cfg.DrmaManualMode; tmp != nil {
				v["drma_manual_mode"] = *tmp
			}

			if tmp := cfg.OverrideAnalysis; tmp != nil {
				v["override_analysis"] = *tmp
			}

			if tmp := cfg.OverrideBand; tmp != nil {
				v["override_band"] = *tmp
			}

			if tmp := cfg.OverrideChannel; tmp != nil {
				v["override_channel"] = *tmp
			}

			if tmp := cfg.OverrideTxpower; tmp != nil {
				v["override_txpower"] = *tmp
			}

			if tmp := cfg.OverrideVaps; tmp != nil {
				v["override_vaps"] = *tmp
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

			if tmp := cfg.SpectrumAnalysis; tmp != nil {
				v["spectrum_analysis"] = *tmp
			}

			if tmp := cfg.VapAll; tmp != nil {
				v["vap_all"] = *tmp
			}

			if tmp := cfg.Vaps; tmp != nil {
				v["vaps"] = flattenWirelessControllerWtpRadio3Vaps(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "vaps"), sort)
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWirelessControllerWtpRadio3Channel(d *schema.ResourceData, v *[]models.WirelessControllerWtpRadio3Channel, prefix string, sort bool) interface{} {
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

func flattenWirelessControllerWtpRadio3Vaps(d *schema.ResourceData, v *[]models.WirelessControllerWtpRadio3Vaps, prefix string, sort bool) interface{} {
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

func flattenWirelessControllerWtpRadio4(d *schema.ResourceData, v *models.WirelessControllerWtpRadio4, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.WirelessControllerWtpRadio4{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
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

			if tmp := cfg.Channel; tmp != nil {
				v["channel"] = flattenWirelessControllerWtpRadio4Channel(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "channel"), sort)
			}

			if tmp := cfg.DrmaManualMode; tmp != nil {
				v["drma_manual_mode"] = *tmp
			}

			if tmp := cfg.OverrideAnalysis; tmp != nil {
				v["override_analysis"] = *tmp
			}

			if tmp := cfg.OverrideBand; tmp != nil {
				v["override_band"] = *tmp
			}

			if tmp := cfg.OverrideChannel; tmp != nil {
				v["override_channel"] = *tmp
			}

			if tmp := cfg.OverrideTxpower; tmp != nil {
				v["override_txpower"] = *tmp
			}

			if tmp := cfg.OverrideVaps; tmp != nil {
				v["override_vaps"] = *tmp
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

			if tmp := cfg.SpectrumAnalysis; tmp != nil {
				v["spectrum_analysis"] = *tmp
			}

			if tmp := cfg.VapAll; tmp != nil {
				v["vap_all"] = *tmp
			}

			if tmp := cfg.Vaps; tmp != nil {
				v["vaps"] = flattenWirelessControllerWtpRadio4Vaps(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "vaps"), sort)
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWirelessControllerWtpRadio4Channel(d *schema.ResourceData, v *[]models.WirelessControllerWtpRadio4Channel, prefix string, sort bool) interface{} {
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

func flattenWirelessControllerWtpRadio4Vaps(d *schema.ResourceData, v *[]models.WirelessControllerWtpRadio4Vaps, prefix string, sort bool) interface{} {
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

func flattenWirelessControllerWtpSplitTunnelingAcl(d *schema.ResourceData, v *[]models.WirelessControllerWtpSplitTunnelingAcl, prefix string, sort bool) interface{} {
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

func refreshObjectWirelessControllerWtp(d *schema.ResourceData, o *models.WirelessControllerWtp, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Admin != nil {
		v := *o.Admin

		if err = d.Set("admin", v); err != nil {
			return diag.Errorf("error reading admin: %v", err)
		}
	}

	if o.Allowaccess != nil {
		v := *o.Allowaccess

		if err = d.Set("allowaccess", v); err != nil {
			return diag.Errorf("error reading allowaccess: %v", err)
		}
	}

	if o.ApcfgProfile != nil {
		v := *o.ApcfgProfile

		if err = d.Set("apcfg_profile", v); err != nil {
			return diag.Errorf("error reading apcfg_profile: %v", err)
		}
	}

	if o.BonjourProfile != nil {
		v := *o.BonjourProfile

		if err = d.Set("bonjour_profile", v); err != nil {
			return diag.Errorf("error reading bonjour_profile: %v", err)
		}
	}

	if o.CoordinateLatitude != nil {
		v := *o.CoordinateLatitude

		if err = d.Set("coordinate_latitude", v); err != nil {
			return diag.Errorf("error reading coordinate_latitude: %v", err)
		}
	}

	if o.CoordinateLongitude != nil {
		v := *o.CoordinateLongitude

		if err = d.Set("coordinate_longitude", v); err != nil {
			return diag.Errorf("error reading coordinate_longitude: %v", err)
		}
	}

	if o.FirmwareProvision != nil {
		v := *o.FirmwareProvision

		if err = d.Set("firmware_provision", v); err != nil {
			return diag.Errorf("error reading firmware_provision: %v", err)
		}
	}

	if o.FirmwareProvisionLatest != nil {
		v := *o.FirmwareProvisionLatest

		if err = d.Set("firmware_provision_latest", v); err != nil {
			return diag.Errorf("error reading firmware_provision_latest: %v", err)
		}
	}

	if o.ImageDownload != nil {
		v := *o.ImageDownload

		if err = d.Set("image_download", v); err != nil {
			return diag.Errorf("error reading image_download: %v", err)
		}
	}

	if o.Index != nil {
		v := *o.Index

		if err = d.Set("index", v); err != nil {
			return diag.Errorf("error reading index: %v", err)
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
			if err = d.Set("lan", flattenWirelessControllerWtpLan(d, o.Lan, "lan", sort)); err != nil {
				return diag.Errorf("error reading lan: %v", err)
			}
		}
	}

	if o.LedState != nil {
		v := *o.LedState

		if err = d.Set("led_state", v); err != nil {
			return diag.Errorf("error reading led_state: %v", err)
		}
	}

	if o.Location != nil {
		v := *o.Location

		if err = d.Set("location", v); err != nil {
			return diag.Errorf("error reading location: %v", err)
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

	if o.MeshBridgeEnable != nil {
		v := *o.MeshBridgeEnable

		if err = d.Set("mesh_bridge_enable", v); err != nil {
			return diag.Errorf("error reading mesh_bridge_enable: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.OverrideAllowaccess != nil {
		v := *o.OverrideAllowaccess

		if err = d.Set("override_allowaccess", v); err != nil {
			return diag.Errorf("error reading override_allowaccess: %v", err)
		}
	}

	if o.OverrideIpFragment != nil {
		v := *o.OverrideIpFragment

		if err = d.Set("override_ip_fragment", v); err != nil {
			return diag.Errorf("error reading override_ip_fragment: %v", err)
		}
	}

	if o.OverrideLan != nil {
		v := *o.OverrideLan

		if err = d.Set("override_lan", v); err != nil {
			return diag.Errorf("error reading override_lan: %v", err)
		}
	}

	if o.OverrideLedState != nil {
		v := *o.OverrideLedState

		if err = d.Set("override_led_state", v); err != nil {
			return diag.Errorf("error reading override_led_state: %v", err)
		}
	}

	if o.OverrideLoginPasswdChange != nil {
		v := *o.OverrideLoginPasswdChange

		if err = d.Set("override_login_passwd_change", v); err != nil {
			return diag.Errorf("error reading override_login_passwd_change: %v", err)
		}
	}

	if o.OverrideSplitTunnel != nil {
		v := *o.OverrideSplitTunnel

		if err = d.Set("override_split_tunnel", v); err != nil {
			return diag.Errorf("error reading override_split_tunnel: %v", err)
		}
	}

	if o.OverrideWanPortMode != nil {
		v := *o.OverrideWanPortMode

		if err = d.Set("override_wan_port_mode", v); err != nil {
			return diag.Errorf("error reading override_wan_port_mode: %v", err)
		}
	}

	if _, ok := d.GetOk("radio_1"); ok {
		if o.Radio1 != nil {
			if err = d.Set("radio_1", flattenWirelessControllerWtpRadio1(d, o.Radio1, "radio_1", sort)); err != nil {
				return diag.Errorf("error reading radio_1: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("radio_2"); ok {
		if o.Radio2 != nil {
			if err = d.Set("radio_2", flattenWirelessControllerWtpRadio2(d, o.Radio2, "radio_2", sort)); err != nil {
				return diag.Errorf("error reading radio_2: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("radio_3"); ok {
		if o.Radio3 != nil {
			if err = d.Set("radio_3", flattenWirelessControllerWtpRadio3(d, o.Radio3, "radio_3", sort)); err != nil {
				return diag.Errorf("error reading radio_3: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("radio_4"); ok {
		if o.Radio4 != nil {
			if err = d.Set("radio_4", flattenWirelessControllerWtpRadio4(d, o.Radio4, "radio_4", sort)); err != nil {
				return diag.Errorf("error reading radio_4: %v", err)
			}
		}
	}

	if o.Region != nil {
		v := *o.Region

		if err = d.Set("region", v); err != nil {
			return diag.Errorf("error reading region: %v", err)
		}
	}

	if o.RegionX != nil {
		v := *o.RegionX

		if err = d.Set("region_x", v); err != nil {
			return diag.Errorf("error reading region_x: %v", err)
		}
	}

	if o.RegionY != nil {
		v := *o.RegionY

		if err = d.Set("region_y", v); err != nil {
			return diag.Errorf("error reading region_y: %v", err)
		}
	}

	if o.SplitTunnelingAcl != nil {
		if err = d.Set("split_tunneling_acl", flattenWirelessControllerWtpSplitTunnelingAcl(d, o.SplitTunnelingAcl, "split_tunneling_acl", sort)); err != nil {
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

	if o.Uuid != nil {
		v := *o.Uuid

		if err = d.Set("uuid", v); err != nil {
			return diag.Errorf("error reading uuid: %v", err)
		}
	}

	if o.WanPortMode != nil {
		v := *o.WanPortMode

		if err = d.Set("wan_port_mode", v); err != nil {
			return diag.Errorf("error reading wan_port_mode: %v", err)
		}
	}

	if o.WtpId != nil {
		v := *o.WtpId

		if err = d.Set("wtp_id", v); err != nil {
			return diag.Errorf("error reading wtp_id: %v", err)
		}
	}

	if o.WtpMode != nil {
		v := *o.WtpMode

		if err = d.Set("wtp_mode", v); err != nil {
			return diag.Errorf("error reading wtp_mode: %v", err)
		}
	}

	if o.WtpProfile != nil {
		v := *o.WtpProfile

		if err = d.Set("wtp_profile", v); err != nil {
			return diag.Errorf("error reading wtp_profile: %v", err)
		}
	}

	return nil
}

func expandWirelessControllerWtpLan(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.WirelessControllerWtpLan, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpLan

	for i := range l {
		tmp := models.WirelessControllerWtpLan{}
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

func expandWirelessControllerWtpRadio1(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.WirelessControllerWtpRadio1, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpRadio1

	for i := range l {
		tmp := models.WirelessControllerWtpRadio1{}
		var pre_append string

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

		pre_append = fmt.Sprintf("%s.%d.channel", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWirelessControllerWtpRadio1Channel(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerWtpRadio1Channel
			// 	}
			tmp.Channel = v2

		}

		pre_append = fmt.Sprintf("%s.%d.drma_manual_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DrmaManualMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_analysis", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideAnalysis = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_band", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideBand = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_channel", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideChannel = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_txpower", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideTxpower = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_vaps", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideVaps = &v2
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

		pre_append = fmt.Sprintf("%s.%d.spectrum_analysis", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SpectrumAnalysis = &v2
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
			v2, _ := expandWirelessControllerWtpRadio1Vaps(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerWtpRadio1Vaps
			// 	}
			tmp.Vaps = v2

		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandWirelessControllerWtpRadio1Channel(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerWtpRadio1Channel, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpRadio1Channel

	for i := range l {
		tmp := models.WirelessControllerWtpRadio1Channel{}
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

func expandWirelessControllerWtpRadio1Vaps(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerWtpRadio1Vaps, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpRadio1Vaps

	for i := range l {
		tmp := models.WirelessControllerWtpRadio1Vaps{}
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

func expandWirelessControllerWtpRadio2(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.WirelessControllerWtpRadio2, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpRadio2

	for i := range l {
		tmp := models.WirelessControllerWtpRadio2{}
		var pre_append string

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

		pre_append = fmt.Sprintf("%s.%d.channel", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWirelessControllerWtpRadio2Channel(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerWtpRadio2Channel
			// 	}
			tmp.Channel = v2

		}

		pre_append = fmt.Sprintf("%s.%d.drma_manual_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DrmaManualMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_analysis", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideAnalysis = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_band", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideBand = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_channel", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideChannel = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_txpower", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideTxpower = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_vaps", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideVaps = &v2
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

		pre_append = fmt.Sprintf("%s.%d.spectrum_analysis", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SpectrumAnalysis = &v2
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
			v2, _ := expandWirelessControllerWtpRadio2Vaps(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerWtpRadio2Vaps
			// 	}
			tmp.Vaps = v2

		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandWirelessControllerWtpRadio2Channel(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerWtpRadio2Channel, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpRadio2Channel

	for i := range l {
		tmp := models.WirelessControllerWtpRadio2Channel{}
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

func expandWirelessControllerWtpRadio2Vaps(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerWtpRadio2Vaps, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpRadio2Vaps

	for i := range l {
		tmp := models.WirelessControllerWtpRadio2Vaps{}
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

func expandWirelessControllerWtpRadio3(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.WirelessControllerWtpRadio3, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpRadio3

	for i := range l {
		tmp := models.WirelessControllerWtpRadio3{}
		var pre_append string

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

		pre_append = fmt.Sprintf("%s.%d.channel", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWirelessControllerWtpRadio3Channel(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerWtpRadio3Channel
			// 	}
			tmp.Channel = v2

		}

		pre_append = fmt.Sprintf("%s.%d.drma_manual_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DrmaManualMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_analysis", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideAnalysis = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_band", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideBand = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_channel", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideChannel = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_txpower", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideTxpower = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_vaps", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideVaps = &v2
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

		pre_append = fmt.Sprintf("%s.%d.spectrum_analysis", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SpectrumAnalysis = &v2
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
			v2, _ := expandWirelessControllerWtpRadio3Vaps(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerWtpRadio3Vaps
			// 	}
			tmp.Vaps = v2

		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandWirelessControllerWtpRadio3Channel(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerWtpRadio3Channel, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpRadio3Channel

	for i := range l {
		tmp := models.WirelessControllerWtpRadio3Channel{}
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

func expandWirelessControllerWtpRadio3Vaps(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerWtpRadio3Vaps, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpRadio3Vaps

	for i := range l {
		tmp := models.WirelessControllerWtpRadio3Vaps{}
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

func expandWirelessControllerWtpRadio4(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.WirelessControllerWtpRadio4, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpRadio4

	for i := range l {
		tmp := models.WirelessControllerWtpRadio4{}
		var pre_append string

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

		pre_append = fmt.Sprintf("%s.%d.channel", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWirelessControllerWtpRadio4Channel(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerWtpRadio4Channel
			// 	}
			tmp.Channel = v2

		}

		pre_append = fmt.Sprintf("%s.%d.drma_manual_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DrmaManualMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_analysis", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideAnalysis = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_band", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideBand = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_channel", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideChannel = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_txpower", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideTxpower = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_vaps", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideVaps = &v2
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

		pre_append = fmt.Sprintf("%s.%d.spectrum_analysis", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SpectrumAnalysis = &v2
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
			v2, _ := expandWirelessControllerWtpRadio4Vaps(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerWtpRadio4Vaps
			// 	}
			tmp.Vaps = v2

		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandWirelessControllerWtpRadio4Channel(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerWtpRadio4Channel, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpRadio4Channel

	for i := range l {
		tmp := models.WirelessControllerWtpRadio4Channel{}
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

func expandWirelessControllerWtpRadio4Vaps(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerWtpRadio4Vaps, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpRadio4Vaps

	for i := range l {
		tmp := models.WirelessControllerWtpRadio4Vaps{}
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

func expandWirelessControllerWtpSplitTunnelingAcl(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerWtpSplitTunnelingAcl, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWtpSplitTunnelingAcl

	for i := range l {
		tmp := models.WirelessControllerWtpSplitTunnelingAcl{}
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

func getObjectWirelessControllerWtp(d *schema.ResourceData, sv string) (*models.WirelessControllerWtp, diag.Diagnostics) {
	obj := models.WirelessControllerWtp{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("admin"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin", sv)
				diags = append(diags, e)
			}
			obj.Admin = &v2
		}
	}
	if v1, ok := d.GetOk("allowaccess"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("allowaccess", sv)
				diags = append(diags, e)
			}
			obj.Allowaccess = &v2
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
	if v1, ok := d.GetOk("bonjour_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bonjour_profile", sv)
				diags = append(diags, e)
			}
			obj.BonjourProfile = &v2
		}
	}
	if v1, ok := d.GetOk("coordinate_latitude"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("coordinate_latitude", sv)
				diags = append(diags, e)
			}
			obj.CoordinateLatitude = &v2
		}
	}
	if v1, ok := d.GetOk("coordinate_longitude"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("coordinate_longitude", sv)
				diags = append(diags, e)
			}
			obj.CoordinateLongitude = &v2
		}
	}
	if v1, ok := d.GetOk("firmware_provision"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("firmware_provision", sv)
				diags = append(diags, e)
			}
			obj.FirmwareProvision = &v2
		}
	}
	if v1, ok := d.GetOk("firmware_provision_latest"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("firmware_provision_latest", sv)
				diags = append(diags, e)
			}
			obj.FirmwareProvisionLatest = &v2
		}
	}
	if v1, ok := d.GetOk("image_download"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("image_download", sv)
				diags = append(diags, e)
			}
			obj.ImageDownload = &v2
		}
	}
	if v1, ok := d.GetOk("index"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("index", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Index = &tmp
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
		t, err := expandWirelessControllerWtpLan(d, v, "lan", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Lan = t
		}
	} else if d.HasChange("lan") {
		old, new := d.GetChange("lan")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Lan = &models.WirelessControllerWtpLan{}
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
	if v1, ok := d.GetOk("location"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("location", sv)
				diags = append(diags, e)
			}
			obj.Location = &v2
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
	if v1, ok := d.GetOk("mesh_bridge_enable"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mesh_bridge_enable", sv)
				diags = append(diags, e)
			}
			obj.MeshBridgeEnable = &v2
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
	if v1, ok := d.GetOk("override_allowaccess"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("override_allowaccess", sv)
				diags = append(diags, e)
			}
			obj.OverrideAllowaccess = &v2
		}
	}
	if v1, ok := d.GetOk("override_ip_fragment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("override_ip_fragment", sv)
				diags = append(diags, e)
			}
			obj.OverrideIpFragment = &v2
		}
	}
	if v1, ok := d.GetOk("override_lan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("override_lan", sv)
				diags = append(diags, e)
			}
			obj.OverrideLan = &v2
		}
	}
	if v1, ok := d.GetOk("override_led_state"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("override_led_state", sv)
				diags = append(diags, e)
			}
			obj.OverrideLedState = &v2
		}
	}
	if v1, ok := d.GetOk("override_login_passwd_change"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("override_login_passwd_change", sv)
				diags = append(diags, e)
			}
			obj.OverrideLoginPasswdChange = &v2
		}
	}
	if v1, ok := d.GetOk("override_split_tunnel"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("override_split_tunnel", sv)
				diags = append(diags, e)
			}
			obj.OverrideSplitTunnel = &v2
		}
	}
	if v1, ok := d.GetOk("override_wan_port_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("override_wan_port_mode", sv)
				diags = append(diags, e)
			}
			obj.OverrideWanPortMode = &v2
		}
	}
	if v, ok := d.GetOk("radio_1"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("radio_1", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerWtpRadio1(d, v, "radio_1", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Radio1 = t
		}
	} else if d.HasChange("radio_1") {
		old, new := d.GetChange("radio_1")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Radio1 = &models.WirelessControllerWtpRadio1{}
		}
	}
	if v, ok := d.GetOk("radio_2"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("radio_2", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerWtpRadio2(d, v, "radio_2", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Radio2 = t
		}
	} else if d.HasChange("radio_2") {
		old, new := d.GetChange("radio_2")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Radio2 = &models.WirelessControllerWtpRadio2{}
		}
	}
	if v, ok := d.GetOk("radio_3"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("radio_3", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerWtpRadio3(d, v, "radio_3", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Radio3 = t
		}
	} else if d.HasChange("radio_3") {
		old, new := d.GetChange("radio_3")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Radio3 = &models.WirelessControllerWtpRadio3{}
		}
	}
	if v, ok := d.GetOk("radio_4"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("radio_4", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerWtpRadio4(d, v, "radio_4", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Radio4 = t
		}
	} else if d.HasChange("radio_4") {
		old, new := d.GetChange("radio_4")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Radio4 = &models.WirelessControllerWtpRadio4{}
		}
	}
	if v1, ok := d.GetOk("region"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("region", sv)
				diags = append(diags, e)
			}
			obj.Region = &v2
		}
	}
	if v1, ok := d.GetOk("region_x"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("region_x", sv)
				diags = append(diags, e)
			}
			obj.RegionX = &v2
		}
	}
	if v1, ok := d.GetOk("region_y"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("region_y", sv)
				diags = append(diags, e)
			}
			obj.RegionY = &v2
		}
	}
	if v, ok := d.GetOk("split_tunneling_acl"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("split_tunneling_acl", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerWtpSplitTunnelingAcl(d, v, "split_tunneling_acl", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SplitTunnelingAcl = t
		}
	} else if d.HasChange("split_tunneling_acl") {
		old, new := d.GetChange("split_tunneling_acl")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SplitTunnelingAcl = &[]models.WirelessControllerWtpSplitTunnelingAcl{}
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
	if v1, ok := d.GetOk("uuid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("uuid", sv)
				diags = append(diags, e)
			}
			obj.Uuid = &v2
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
	if v1, ok := d.GetOk("wtp_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wtp_id", sv)
				diags = append(diags, e)
			}
			obj.WtpId = &v2
		}
	}
	if v1, ok := d.GetOk("wtp_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("wtp_mode", sv)
				diags = append(diags, e)
			}
			obj.WtpMode = &v2
		}
	}
	if v1, ok := d.GetOk("wtp_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wtp_profile", sv)
				diags = append(diags, e)
			}
			obj.WtpProfile = &v2
		}
	}
	return &obj, diags
}
