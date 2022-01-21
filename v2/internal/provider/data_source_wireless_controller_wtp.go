// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Wireless Termination Points (WTPs), that is, FortiAPs or APs to be managed by FortiGate.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWirelessControllerWtp() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Wireless Termination Points (WTPs), that is, FortiAPs or APs to be managed by FortiGate.",

		ReadContext: dataSourceWirelessControllerWtpRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"admin": {
				Type:        schema.TypeString,
				Description: "Configure how the FortiGate operating as a wireless controller discovers and manages this WTP, AP or FortiAP.",
				Computed:    true,
			},
			"allowaccess": {
				Type:        schema.TypeString,
				Description: "Control management access to the managed WTP, FortiAP, or AP. Separate entries with a space.",
				Computed:    true,
			},
			"apcfg_profile": {
				Type:        schema.TypeString,
				Description: "AP local configuration profile name.",
				Computed:    true,
			},
			"bonjour_profile": {
				Type:        schema.TypeString,
				Description: "Bonjour profile name.",
				Computed:    true,
			},
			"coordinate_latitude": {
				Type:        schema.TypeString,
				Description: "WTP latitude coordinate.",
				Computed:    true,
			},
			"coordinate_longitude": {
				Type:        schema.TypeString,
				Description: "WTP longitude coordinate.",
				Computed:    true,
			},
			"firmware_provision": {
				Type:        schema.TypeString,
				Description: "Firmware version to provision to this FortiAP on bootup (major.minor.build, i.e. 6.2.1234).",
				Computed:    true,
			},
			"firmware_provision_latest": {
				Type:        schema.TypeString,
				Description: "Enable/disable one-time automatic provisioning of the latest firmware version.",
				Computed:    true,
			},
			"image_download": {
				Type:        schema.TypeString,
				Description: "Enable/disable WTP image download.",
				Computed:    true,
			},
			"index": {
				Type:        schema.TypeInt,
				Description: "Index (0 - 4294967295).",
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
			"led_state": {
				Type:        schema.TypeString,
				Description: "Enable to allow the FortiAPs LEDs to light. Disable to keep the LEDs off. You may want to keep the LEDs off so they are not distracting in low light areas etc.",
				Computed:    true,
			},
			"location": {
				Type:        schema.TypeString,
				Description: "Field for describing the physical location of the WTP, AP or FortiAP.",
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
			"mesh_bridge_enable": {
				Type:        schema.TypeString,
				Description: "Enable/disable mesh Ethernet bridge when WTP is configured as a mesh branch/leaf AP.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "WTP, AP or FortiAP configuration name.",
				Computed:    true,
			},
			"override_allowaccess": {
				Type:        schema.TypeString,
				Description: "Enable to override the WTP profile management access configuration.",
				Computed:    true,
			},
			"override_ip_fragment": {
				Type:        schema.TypeString,
				Description: "Enable/disable overriding the WTP profile IP fragment prevention setting.",
				Computed:    true,
			},
			"override_lan": {
				Type:        schema.TypeString,
				Description: "Enable to override the WTP profile LAN port setting.",
				Computed:    true,
			},
			"override_led_state": {
				Type:        schema.TypeString,
				Description: "Enable to override the profile LED state setting for this FortiAP. You must enable this option to use the led-state command to turn off the FortiAP's LEDs.",
				Computed:    true,
			},
			"override_login_passwd_change": {
				Type:        schema.TypeString,
				Description: "Enable to override the WTP profile login-password (administrator password) setting.",
				Computed:    true,
			},
			"override_split_tunnel": {
				Type:        schema.TypeString,
				Description: "Enable/disable overriding the WTP profile split tunneling setting.",
				Computed:    true,
			},
			"override_wan_port_mode": {
				Type:        schema.TypeString,
				Description: "Enable/disable overriding the wan-port-mode in the WTP profile.",
				Computed:    true,
			},
			"radio_1": {
				Type:        schema.TypeList,
				Description: "Configuration options for radio 1.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"drma_manual_mode": {
							Type:        schema.TypeString,
							Description: "Radio mode to be used for DRMA manual mode (default = ncf).",
							Computed:    true,
						},
						"override_analysis": {
							Type:        schema.TypeString,
							Description: "Enable to override the WTP profile spectrum analysis configuration.",
							Computed:    true,
						},
						"override_band": {
							Type:        schema.TypeString,
							Description: "Enable to override the WTP profile band setting.",
							Computed:    true,
						},
						"override_channel": {
							Type:        schema.TypeString,
							Description: "Enable to override WTP profile channel settings.",
							Computed:    true,
						},
						"override_txpower": {
							Type:        schema.TypeString,
							Description: "Enable to override the WTP profile power level configuration.",
							Computed:    true,
						},
						"override_vaps": {
							Type:        schema.TypeString,
							Description: "Enable to override WTP profile Virtual Access Point (VAP) settings.",
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
						"spectrum_analysis": {
							Type:        schema.TypeString,
							Description: "Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.",
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
					},
				},
			},
			"radio_2": {
				Type:        schema.TypeList,
				Description: "Configuration options for radio 2.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"drma_manual_mode": {
							Type:        schema.TypeString,
							Description: "Radio mode to be used for DRMA manual mode (default = ncf).",
							Computed:    true,
						},
						"override_analysis": {
							Type:        schema.TypeString,
							Description: "Enable to override the WTP profile spectrum analysis configuration.",
							Computed:    true,
						},
						"override_band": {
							Type:        schema.TypeString,
							Description: "Enable to override the WTP profile band setting.",
							Computed:    true,
						},
						"override_channel": {
							Type:        schema.TypeString,
							Description: "Enable to override WTP profile channel settings.",
							Computed:    true,
						},
						"override_txpower": {
							Type:        schema.TypeString,
							Description: "Enable to override the WTP profile power level configuration.",
							Computed:    true,
						},
						"override_vaps": {
							Type:        schema.TypeString,
							Description: "Enable to override WTP profile Virtual Access Point (VAP) settings.",
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
						"spectrum_analysis": {
							Type:        schema.TypeString,
							Description: "Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.",
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
					},
				},
			},
			"radio_3": {
				Type:        schema.TypeList,
				Description: "Configuration options for radio 3.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"drma_manual_mode": {
							Type:        schema.TypeString,
							Description: "Radio mode to be used for DRMA manual mode (default = ncf).",
							Computed:    true,
						},
						"override_analysis": {
							Type:        schema.TypeString,
							Description: "Enable to override the WTP profile spectrum analysis configuration.",
							Computed:    true,
						},
						"override_band": {
							Type:        schema.TypeString,
							Description: "Enable to override the WTP profile band setting.",
							Computed:    true,
						},
						"override_channel": {
							Type:        schema.TypeString,
							Description: "Enable to override WTP profile channel settings.",
							Computed:    true,
						},
						"override_txpower": {
							Type:        schema.TypeString,
							Description: "Enable to override the WTP profile power level configuration.",
							Computed:    true,
						},
						"override_vaps": {
							Type:        schema.TypeString,
							Description: "Enable to override WTP profile Virtual Access Point (VAP) settings.",
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
						"spectrum_analysis": {
							Type:        schema.TypeString,
							Description: "Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.",
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
					},
				},
			},
			"radio_4": {
				Type:        schema.TypeList,
				Description: "Configuration options for radio 4.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
							Description: "WiFi band that Radio 4 operates on.",
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
						"drma_manual_mode": {
							Type:        schema.TypeString,
							Description: "Radio mode to be used for DRMA manual mode (default = ncf).",
							Computed:    true,
						},
						"override_analysis": {
							Type:        schema.TypeString,
							Description: "Enable to override the WTP profile spectrum analysis configuration.",
							Computed:    true,
						},
						"override_band": {
							Type:        schema.TypeString,
							Description: "Enable to override the WTP profile band setting.",
							Computed:    true,
						},
						"override_channel": {
							Type:        schema.TypeString,
							Description: "Enable to override WTP profile channel settings.",
							Computed:    true,
						},
						"override_txpower": {
							Type:        schema.TypeString,
							Description: "Enable to override the WTP profile power level configuration.",
							Computed:    true,
						},
						"override_vaps": {
							Type:        schema.TypeString,
							Description: "Enable to override WTP profile Virtual Access Point (VAP) settings.",
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
						"spectrum_analysis": {
							Type:        schema.TypeString,
							Description: "Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.",
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
					},
				},
			},
			"region": {
				Type:        schema.TypeString,
				Description: "Region name WTP is associated with.",
				Computed:    true,
			},
			"region_x": {
				Type:        schema.TypeString,
				Description: "Relative horizontal region coordinate (between 0 and 1).",
				Computed:    true,
			},
			"region_y": {
				Type:        schema.TypeString,
				Description: "Relative vertical region coordinate (between 0 and 1).",
				Computed:    true,
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
			"uuid": {
				Type:        schema.TypeString,
				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Computed:    true,
			},
			"wan_port_mode": {
				Type:        schema.TypeString,
				Description: "Enable/disable using the FortiAP WAN port as a LAN port.",
				Computed:    true,
			},
			"wtp_id": {
				Type:        schema.TypeString,
				Description: "WTP ID.",
				Required:    true,
			},
			"wtp_mode": {
				Type:        schema.TypeString,
				Description: "WTP, AP, or FortiAP operating mode; normal (by default) or remote. A tunnel mode SSID can be assigned to an AP in normal mode but not remote mode, while a local-bridge mode SSID can be assigned to an AP in either normal mode or remote mode.",
				Computed:    true,
			},
			"wtp_profile": {
				Type:        schema.TypeString,
				Description: "WTP profile name to apply to this WTP, AP or FortiAP.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWirelessControllerWtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("wtp_id")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadWirelessControllerWtp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerWtp dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerWtp(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
