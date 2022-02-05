// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Portal.

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
)

func resourceVpnSslWebPortal() *schema.Resource {
	return &schema.Resource{
		Description: "Portal.",

		CreateContext: resourceVpnSslWebPortalCreate,
		ReadContext:   resourceVpnSslWebPortalRead,
		UpdateContext: resourceVpnSslWebPortalUpdate,
		DeleteContext: resourceVpnSslWebPortalDelete,

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
			"allow_user_access": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Allow user access to SSL-VPN applications.",
				Optional:         true,
				Computed:         true,
			},
			"auto_connect": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable automatic connect by client when system is up.",
				Optional:    true,
				Computed:    true,
			},
			"bookmark_group": {
				Type:        schema.TypeList,
				Description: "Portal bookmark group.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bookmarks": {
							Type:        schema.TypeList,
							Description: "Bookmark table.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_params": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 128),

										Description: "Additional parameters.",
										Optional:    true,
										Computed:    true,
									},
									"apptype": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"ftp", "rdp", "sftp", "smb", "ssh", "telnet", "vnc", "web"}, false),

										Description: "Application type.",
										Optional:    true,
										Computed:    true,
									},
									"color_depth": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"32", "16", "8"}, false),

										Description: "Color depth per pixel.",
										Optional:    true,
										Computed:    true,
									},
									"description": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 128),

										Description: "Description.",
										Optional:    true,
										Computed:    true,
									},
									"domain": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 128),

										Description: "Login domain.",
										Optional:    true,
										Computed:    true,
									},
									"folder": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 128),

										Description: "Network shared file folder parameter.",
										Optional:    true,
										Computed:    true,
									},
									"form_data": {
										Type:        schema.TypeList,
										Description: "Form data.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 35),

													Description: "Name.",
													Optional:    true,
													Computed:    true,
												},
												"value": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 63),

													Description: "Value.",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"height": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(480, 65535),

										Description: "Screen height (range from 480 - 65535, default = 768).",
										Optional:    true,
										Computed:    true,
									},
									"host": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 128),

										Description: "Host name/IP parameter.",
										Optional:    true,
										Computed:    true,
									},
									"keyboard_layout": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"ar-101", "ar-102", "ar-102-azerty", "can-mul", "cz", "cz-qwerty", "cz-pr", "da", "nl", "de", "de-ch", "de-ibm", "en-uk", "en-uk-ext", "en-us", "en-us-dvorak", "es", "es-var", "fi", "fi-sami", "fr", "fr-ca", "fr-ch", "fr-be", "hr", "hu", "hu-101", "it", "it-142", "ja", "ko", "lt", "lt-ibm", "lt-std", "lav-std", "lav-leg", "mk", "mk-std", "no", "no-sami", "pol-214", "pol-pr", "pt", "pt-br", "pt-br-abnt2", "ru", "ru-mne", "ru-t", "sl", "sv", "sv-sami", "tuk", "tur-f", "tur-q", "zh-sym-sg-us", "zh-sym-us", "zh-tr-hk", "zh-tr-mo", "zh-tr-us"}, false),

										Description: "Keyboard layout.",
										Optional:    true,
										Computed:    true,
									},
									"listening_port": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),

										Description: "Listening port (0 - 65535).",
										Optional:    true,
										Computed:    true,
									},
									"load_balancing_info": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 511),

										Description: "The load balancing information or cookie which should be provided to the connection broker.",
										Optional:    true,
										Computed:    true,
									},
									"logon_password": {
										Type: schema.TypeString,

										Description: "Logon password.",
										Optional:    true,
										Computed:    true,
										Sensitive:   true,
									},
									"logon_user": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Logon user.",
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Bookmark name.",
										Optional:    true,
										Computed:    true,
									},
									"port": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),

										Description: "Remote port.",
										Optional:    true,
										Computed:    true,
									},
									"preconnection_blob": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 511),

										Description: "An arbitrary string which identifies the RDP source.",
										Optional:    true,
										Computed:    true,
									},
									"preconnection_id": {
										Type: schema.TypeInt,

										Description: "The numeric ID of the RDP source (0-4294967295).",
										Optional:    true,
										Computed:    true,
									},
									"remote_port": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),

										Description: "Remote port (0 - 65535).",
										Optional:    true,
										Computed:    true,
									},
									"restricted_admin": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable restricted admin mode for RDP.",
										Optional:    true,
										Computed:    true,
									},
									"security": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"rdp", "nla", "tls", "any"}, false),

										Description: "Security mode for RDP connection.",
										Optional:    true,
										Computed:    true,
									},
									"send_preconnection_id": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable sending of preconnection ID.",
										Optional:    true,
										Computed:    true,
									},
									"server_layout": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"de-de-qwertz", "en-gb-qwerty", "en-us-qwerty", "es-es-qwerty", "fr-ca-qwerty", "fr-fr-azerty", "fr-ch-qwertz", "it-it-qwerty", "ja-jp-qwerty", "pt-br-qwerty", "sv-se-qwerty", "tr-tr-qwerty", "failsafe"}, false),

										Description: "Server side keyboard layout.",
										Optional:    true,
										Computed:    true,
									},
									"show_status_window": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable showing of status window.",
										Optional:    true,
										Computed:    true,
									},
									"sso": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"disable", "static", "auto"}, false),

										Description: "Single Sign-On.",
										Optional:    true,
										Computed:    true,
									},
									"sso_credential": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"sslvpn-login", "alternative"}, false),

										Description: "Single sign-on credentials.",
										Optional:    true,
										Computed:    true,
									},
									"sso_credential_sent_once": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Single sign-on credentials are only sent once to remote server.",
										Optional:    true,
										Computed:    true,
									},
									"sso_password": {
										Type: schema.TypeString,

										Description: "SSO password.",
										Optional:    true,
										Computed:    true,
										Sensitive:   true,
									},
									"sso_username": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "SSO user name.",
										Optional:    true,
										Computed:    true,
									},
									"url": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 128),

										Description: "URL parameter.",
										Optional:    true,
										Computed:    true,
									},
									"width": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(640, 65535),

										Description: "Screen width (range from 640 - 65535, default = 1024).",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Bookmark group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"clipboard": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to support RDP/VPC clipboard functionality.",
				Optional:    true,
				Computed:    true,
			},
			"custom_lang": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Change the web portal display language. Overrides config system global set language. You can use config system custom-language and execute system custom-language to add custom language files.",
				Optional:    true,
				Computed:    true,
			},
			"customize_forticlient_download_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable support of customized download URL for FortiClient.",
				Optional:    true,
				Computed:    true,
			},
			"display_bookmark": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to display the web portal bookmark widget.",
				Optional:    true,
				Computed:    true,
			},
			"display_connection_tools": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to display the web portal connection tools widget.",
				Optional:    true,
				Computed:    true,
			},
			"display_history": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to display the web portal user login history widget.",
				Optional:    true,
				Computed:    true,
			},
			"display_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to display the web portal status widget.",
				Optional:    true,
				Computed:    true,
			},
			"dns_server1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IPv4 DNS server 1.",
				Optional:    true,
				Computed:    true,
			},
			"dns_server2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IPv4 DNS server 2.",
				Optional:    true,
				Computed:    true,
			},
			"dns_suffix": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 253),

				Description: "DNS suffix.",
				Optional:    true,
				Computed:    true,
			},
			"exclusive_routing": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable all traffic go through tunnel only.",
				Optional:    true,
				Computed:    true,
			},
			"forticlient_download": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable download option for FortiClient.",
				Optional:    true,
				Computed:    true,
			},
			"forticlient_download_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"direct", "ssl-vpn"}, false),

				Description: "FortiClient download method.",
				Optional:    true,
				Computed:    true,
			},
			"heading": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "Web portal heading message.",
				Optional:    true,
				Computed:    true,
			},
			"hide_sso_credential": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to prevent SSO credential being sent to client.",
				Optional:    true,
				Computed:    true,
			},
			"host_check": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "av", "fw", "av-fw", "custom"}, false),

				Description: "Type of host checking performed on endpoints.",
				Optional:    true,
				Computed:    true,
			},
			"host_check_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(120, 259200),

				Description: "Periodic host check interval. Value of 0 means disabled and host checking only happens when the endpoint connects.",
				Optional:    true,
				Computed:    true,
			},
			"host_check_policy": {
				Type:        schema.TypeList,
				Description: "One or more policies to require the endpoint to have specific security software.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Host check software list name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ip_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"range", "user-group"}, false),

				Description: "Method by which users of this SSL-VPN tunnel obtain IP addresses.",
				Optional:    true,
				Computed:    true,
			},
			"ip_pools": {
				Type:        schema.TypeList,
				Description: "IPv4 firewall source address objects reserved for SSL-VPN tunnel mode clients.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ipv6_dns_server1": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 DNS server 1.",
				Optional:         true,
				Computed:         true,
			},
			"ipv6_dns_server2": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 DNS server 2.",
				Optional:         true,
				Computed:         true,
			},
			"ipv6_exclusive_routing": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable all IPv6 traffic go through tunnel only.",
				Optional:    true,
				Computed:    true,
			},
			"ipv6_pools": {
				Type:        schema.TypeList,
				Description: "IPv6 firewall source address objects reserved for SSL-VPN tunnel mode clients.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ipv6_service_restriction": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPv6 tunnel service restriction.",
				Optional:    true,
				Computed:    true,
			},
			"ipv6_split_tunneling": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPv6 split tunneling.",
				Optional:    true,
				Computed:    true,
			},
			"ipv6_split_tunneling_routing_address": {
				Type:        schema.TypeList,
				Description: "IPv6 SSL-VPN tunnel mode firewall address objects that override firewall policy destination addresses to control split-tunneling access.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ipv6_split_tunneling_routing_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to negate IPv6 split tunneling routing address.",
				Optional:    true,
				Computed:    true,
			},
			"ipv6_tunnel_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPv6 SSL-VPN tunnel mode.",
				Optional:    true,
				Computed:    true,
			},
			"ipv6_wins_server1": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 WINS server 1.",
				Optional:         true,
				Computed:         true,
			},
			"ipv6_wins_server2": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 WINS server 2.",
				Optional:         true,
				Computed:         true,
			},
			"keep_alive": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable automatic reconnect for FortiClient connections.",
				Optional:    true,
				Computed:    true,
			},
			"limit_user_logins": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to limit each user to one SSL-VPN session at a time.",
				Optional:    true,
				Computed:    true,
			},
			"mac_addr_action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"allow", "deny"}, false),

				Description: "Client MAC address action.",
				Optional:    true,
				Computed:    true,
			},
			"mac_addr_check": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable MAC address host checking.",
				Optional:    true,
				Computed:    true,
			},
			"mac_addr_check_rule": {
				Type:        schema.TypeList,
				Description: "Client MAC address check rule.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mac_addr_list": {
							Type:        schema.TypeList,
							Description: "Client MAC address list.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"addr": {
										Type: schema.TypeString,

										Description: "Client MAC address.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"mac_addr_mask": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 48),

							Description: "Client MAC address mask.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Client MAC address check rule name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"macos_forticlient_download_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),

				Description: "Download URL for Mac FortiClient.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Portal name.",
				ForceNew:    true,
				Required:    true,
			},
			"os_check": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to let the FortiGate decide action based on client OS.",
				Optional:    true,
				Computed:    true,
			},
			"os_check_list": {
				Type:        schema.TypeList,
				Description: "SSL-VPN OS checks.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"deny", "allow", "check-up-to-date"}, false),

							Description: "OS check options.",
							Optional:    true,
							Computed:    true,
						},
						"latest_patch_level": {
							Type: schema.TypeString,

							Description: "Latest OS patch level.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Name.",
							Optional:    true,
							Computed:    true,
						},
						"tolerance": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "OS patch level tolerance.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"prefer_ipv6_dns": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Prefer to query IPv6 DNS server first if enabled.",
				Optional:    true,
				Computed:    true,
			},
			"redir_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Client login redirect URL.",
				Optional:    true,
				Computed:    true,
			},
			"rewrite_ip_uri_ui": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Rewrite contents for URI contains IP and /ui/ (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"save_password": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiClient saving the user's password.",
				Optional:    true,
				Computed:    true,
			},
			"service_restriction": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable tunnel service restriction.",
				Optional:    true,
				Computed:    true,
			},
			"skip_check_for_browser": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to skip host check for browser support.",
				Optional:    true,
				Computed:    true,
			},
			"skip_check_for_unsupported_os": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to skip host check if client OS does not support it.",
				Optional:    true,
				Computed:    true,
			},
			"smb_max_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"smbv1", "smbv2", "smbv3"}, false),

				Description: "SMB maximum client protocol version.",
				Optional:    true,
				Computed:    true,
			},
			"smb_min_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"smbv1", "smbv2", "smbv3"}, false),

				Description: "SMB minimum client protocol version.",
				Optional:    true,
				Computed:    true,
			},
			"smb_ntlmv1_auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable support of NTLMv1 for Samba authentication.",
				Optional:    true,
				Computed:    true,
			},
			"smbv1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "SMB version 1.",
				Optional:    true,
				Computed:    true,
			},
			"split_dns": {
				Type:        schema.TypeList,
				Description: "Split DNS for SSL-VPN.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_server1": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "DNS server 1.",
							Optional:    true,
							Computed:    true,
						},
						"dns_server2": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "DNS server 2.",
							Optional:    true,
							Computed:    true,
						},
						"domains": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1024),

							Description: "Split DNS domains used for SSL-VPN clients separated by comma.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"ipv6_dns_server1": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "IPv6 DNS server 1.",
							Optional:         true,
							Computed:         true,
						},
						"ipv6_dns_server2": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "IPv6 DNS server 2.",
							Optional:         true,
							Computed:         true,
						},
					},
				},
			},
			"split_tunneling": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPv4 split tunneling.",
				Optional:    true,
				Computed:    true,
			},
			"split_tunneling_routing_address": {
				Type:        schema.TypeList,
				Description: "IPv4 SSL-VPN tunnel mode firewall address objects that override firewall policy destination addresses to control split-tunneling access.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"split_tunneling_routing_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to negate split tunneling routing address.",
				Optional:    true,
				Computed:    true,
			},
			"theme": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"jade", "neutrino", "mariner", "graphite", "melongene", "dark-matter", "onyx", "eclipse"}, false),

				Description: "Web portal color scheme.",
				Optional:    true,
				Computed:    true,
			},
			"transform_backward_slashes": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Transform backward slashes to forward slashes in URLs.",
				Optional:    true,
				Computed:    true,
			},
			"tunnel_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPv4 SSL-VPN tunnel mode.",
				Optional:    true,
				Computed:    true,
			},
			"use_sdwan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Use SD-WAN rules to get output interface.",
				Optional:    true,
				Computed:    true,
			},
			"user_bookmark": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to allow web portal users to create their own bookmarks.",
				Optional:    true,
				Computed:    true,
			},
			"user_group_bookmark": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to allow web portal users to create bookmarks for all users in the same user group.",
				Optional:    true,
				Computed:    true,
			},
			"web_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SSL-VPN web mode.",
				Optional:    true,
				Computed:    true,
			},
			"windows_forticlient_download_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),

				Description: "Download URL for Windows FortiClient.",
				Optional:    true,
				Computed:    true,
			},
			"wins_server1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IPv4 WINS server 1.",
				Optional:    true,
				Computed:    true,
			},
			"wins_server2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IPv4 WINS server 1.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceVpnSslWebPortalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating VpnSslWebPortal resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectVpnSslWebPortal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateVpnSslWebPortal(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnSslWebPortal")
	}

	return resourceVpnSslWebPortalRead(ctx, d, meta)
}

func resourceVpnSslWebPortalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVpnSslWebPortal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateVpnSslWebPortal(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating VpnSslWebPortal resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnSslWebPortal")
	}

	return resourceVpnSslWebPortalRead(ctx, d, meta)
}

func resourceVpnSslWebPortalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteVpnSslWebPortal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting VpnSslWebPortal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslWebPortalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnSslWebPortal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnSslWebPortal resource: %v", err)
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

	diags := refreshObjectVpnSslWebPortal(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenVpnSslWebPortalBookmarkGroup(d *schema.ResourceData, v *[]models.VpnSslWebPortalBookmarkGroup, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Bookmarks; tmp != nil {
				v["bookmarks"] = flattenVpnSslWebPortalBookmarkGroupBookmarks(d, tmp, prefix+"bookmarks", sort)
			}

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

func flattenVpnSslWebPortalBookmarkGroupBookmarks(d *schema.ResourceData, v *[]models.VpnSslWebPortalBookmarkGroupBookmarks, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AdditionalParams; tmp != nil {
				v["additional_params"] = *tmp
			}

			if tmp := cfg.Apptype; tmp != nil {
				v["apptype"] = *tmp
			}

			if tmp := cfg.ColorDepth; tmp != nil {
				v["color_depth"] = *tmp
			}

			if tmp := cfg.Description; tmp != nil {
				v["description"] = *tmp
			}

			if tmp := cfg.Domain; tmp != nil {
				v["domain"] = *tmp
			}

			if tmp := cfg.Folder; tmp != nil {
				v["folder"] = *tmp
			}

			if tmp := cfg.FormData; tmp != nil {
				v["form_data"] = flattenVpnSslWebPortalBookmarkGroupBookmarksFormData(d, tmp, prefix+"form_data", sort)
			}

			if tmp := cfg.Height; tmp != nil {
				v["height"] = *tmp
			}

			if tmp := cfg.Host; tmp != nil {
				v["host"] = *tmp
			}

			if tmp := cfg.KeyboardLayout; tmp != nil {
				v["keyboard_layout"] = *tmp
			}

			if tmp := cfg.ListeningPort; tmp != nil {
				v["listening_port"] = *tmp
			}

			if tmp := cfg.LoadBalancingInfo; tmp != nil {
				v["load_balancing_info"] = *tmp
			}

			if tmp := cfg.LogonPassword; tmp != nil {
				v["logon_password"] = *tmp
			}

			if tmp := cfg.LogonUser; tmp != nil {
				v["logon_user"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Port; tmp != nil {
				v["port"] = *tmp
			}

			if tmp := cfg.PreconnectionBlob; tmp != nil {
				v["preconnection_blob"] = *tmp
			}

			if tmp := cfg.PreconnectionId; tmp != nil {
				v["preconnection_id"] = *tmp
			}

			if tmp := cfg.RemotePort; tmp != nil {
				v["remote_port"] = *tmp
			}

			if tmp := cfg.RestrictedAdmin; tmp != nil {
				v["restricted_admin"] = *tmp
			}

			if tmp := cfg.Security; tmp != nil {
				v["security"] = *tmp
			}

			if tmp := cfg.SendPreconnectionId; tmp != nil {
				v["send_preconnection_id"] = *tmp
			}

			if tmp := cfg.ServerLayout; tmp != nil {
				v["server_layout"] = *tmp
			}

			if tmp := cfg.ShowStatusWindow; tmp != nil {
				v["show_status_window"] = *tmp
			}

			if tmp := cfg.Sso; tmp != nil {
				v["sso"] = *tmp
			}

			if tmp := cfg.SsoCredential; tmp != nil {
				v["sso_credential"] = *tmp
			}

			if tmp := cfg.SsoCredentialSentOnce; tmp != nil {
				v["sso_credential_sent_once"] = *tmp
			}

			if tmp := cfg.SsoPassword; tmp != nil {
				v["sso_password"] = *tmp
			}

			if tmp := cfg.SsoUsername; tmp != nil {
				v["sso_username"] = *tmp
			}

			if tmp := cfg.Url; tmp != nil {
				v["url"] = *tmp
			}

			if tmp := cfg.Width; tmp != nil {
				v["width"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksFormData(d *schema.ResourceData, v *[]models.VpnSslWebPortalBookmarkGroupBookmarksFormData, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Value; tmp != nil {
				v["value"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenVpnSslWebPortalHostCheckPolicy(d *schema.ResourceData, v *[]models.VpnSslWebPortalHostCheckPolicy, prefix string, sort bool) interface{} {
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

func flattenVpnSslWebPortalIpPools(d *schema.ResourceData, v *[]models.VpnSslWebPortalIpPools, prefix string, sort bool) interface{} {
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

func flattenVpnSslWebPortalIpv6Pools(d *schema.ResourceData, v *[]models.VpnSslWebPortalIpv6Pools, prefix string, sort bool) interface{} {
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

func flattenVpnSslWebPortalIpv6SplitTunnelingRoutingAddress(d *schema.ResourceData, v *[]models.VpnSslWebPortalIpv6SplitTunnelingRoutingAddress, prefix string, sort bool) interface{} {
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

func flattenVpnSslWebPortalMacAddrCheckRule(d *schema.ResourceData, v *[]models.VpnSslWebPortalMacAddrCheckRule, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.MacAddrList; tmp != nil {
				v["mac_addr_list"] = flattenVpnSslWebPortalMacAddrCheckRuleMacAddrList(d, tmp, prefix+"mac_addr_list", sort)
			}

			if tmp := cfg.MacAddrMask; tmp != nil {
				v["mac_addr_mask"] = *tmp
			}

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

func flattenVpnSslWebPortalMacAddrCheckRuleMacAddrList(d *schema.ResourceData, v *[]models.VpnSslWebPortalMacAddrCheckRuleMacAddrList, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Addr; tmp != nil {
				v["addr"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "addr")
	}

	return flat
}

func flattenVpnSslWebPortalOsCheckList(d *schema.ResourceData, v *[]models.VpnSslWebPortalOsCheckList, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.LatestPatchLevel; tmp != nil {
				v["latest_patch_level"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Tolerance; tmp != nil {
				v["tolerance"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenVpnSslWebPortalSplitDns(d *schema.ResourceData, v *[]models.VpnSslWebPortalSplitDns, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.DnsServer1; tmp != nil {
				v["dns_server1"] = *tmp
			}

			if tmp := cfg.DnsServer2; tmp != nil {
				v["dns_server2"] = *tmp
			}

			if tmp := cfg.Domains; tmp != nil {
				v["domains"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Ipv6DnsServer1; tmp != nil {
				v["ipv6_dns_server1"] = *tmp
			}

			if tmp := cfg.Ipv6DnsServer2; tmp != nil {
				v["ipv6_dns_server2"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenVpnSslWebPortalSplitTunnelingRoutingAddress(d *schema.ResourceData, v *[]models.VpnSslWebPortalSplitTunnelingRoutingAddress, prefix string, sort bool) interface{} {
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

func refreshObjectVpnSslWebPortal(d *schema.ResourceData, o *models.VpnSslWebPortal, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AllowUserAccess != nil {
		v := *o.AllowUserAccess

		if err = d.Set("allow_user_access", v); err != nil {
			return diag.Errorf("error reading allow_user_access: %v", err)
		}
	}

	if o.AutoConnect != nil {
		v := *o.AutoConnect

		if err = d.Set("auto_connect", v); err != nil {
			return diag.Errorf("error reading auto_connect: %v", err)
		}
	}

	if o.BookmarkGroup != nil {
		if err = d.Set("bookmark_group", flattenVpnSslWebPortalBookmarkGroup(d, o.BookmarkGroup, "bookmark_group", sort)); err != nil {
			return diag.Errorf("error reading bookmark_group: %v", err)
		}
	}

	if o.Clipboard != nil {
		v := *o.Clipboard

		if err = d.Set("clipboard", v); err != nil {
			return diag.Errorf("error reading clipboard: %v", err)
		}
	}

	if o.CustomLang != nil {
		v := *o.CustomLang

		if err = d.Set("custom_lang", v); err != nil {
			return diag.Errorf("error reading custom_lang: %v", err)
		}
	}

	if o.CustomizeForticlientDownloadUrl != nil {
		v := *o.CustomizeForticlientDownloadUrl

		if err = d.Set("customize_forticlient_download_url", v); err != nil {
			return diag.Errorf("error reading customize_forticlient_download_url: %v", err)
		}
	}

	if o.DisplayBookmark != nil {
		v := *o.DisplayBookmark

		if err = d.Set("display_bookmark", v); err != nil {
			return diag.Errorf("error reading display_bookmark: %v", err)
		}
	}

	if o.DisplayConnectionTools != nil {
		v := *o.DisplayConnectionTools

		if err = d.Set("display_connection_tools", v); err != nil {
			return diag.Errorf("error reading display_connection_tools: %v", err)
		}
	}

	if o.DisplayHistory != nil {
		v := *o.DisplayHistory

		if err = d.Set("display_history", v); err != nil {
			return diag.Errorf("error reading display_history: %v", err)
		}
	}

	if o.DisplayStatus != nil {
		v := *o.DisplayStatus

		if err = d.Set("display_status", v); err != nil {
			return diag.Errorf("error reading display_status: %v", err)
		}
	}

	if o.DnsServer1 != nil {
		v := *o.DnsServer1

		if err = d.Set("dns_server1", v); err != nil {
			return diag.Errorf("error reading dns_server1: %v", err)
		}
	}

	if o.DnsServer2 != nil {
		v := *o.DnsServer2

		if err = d.Set("dns_server2", v); err != nil {
			return diag.Errorf("error reading dns_server2: %v", err)
		}
	}

	if o.DnsSuffix != nil {
		v := *o.DnsSuffix

		if err = d.Set("dns_suffix", v); err != nil {
			return diag.Errorf("error reading dns_suffix: %v", err)
		}
	}

	if o.ExclusiveRouting != nil {
		v := *o.ExclusiveRouting

		if err = d.Set("exclusive_routing", v); err != nil {
			return diag.Errorf("error reading exclusive_routing: %v", err)
		}
	}

	if o.ForticlientDownload != nil {
		v := *o.ForticlientDownload

		if err = d.Set("forticlient_download", v); err != nil {
			return diag.Errorf("error reading forticlient_download: %v", err)
		}
	}

	if o.ForticlientDownloadMethod != nil {
		v := *o.ForticlientDownloadMethod

		if err = d.Set("forticlient_download_method", v); err != nil {
			return diag.Errorf("error reading forticlient_download_method: %v", err)
		}
	}

	if o.Heading != nil {
		v := *o.Heading

		if err = d.Set("heading", v); err != nil {
			return diag.Errorf("error reading heading: %v", err)
		}
	}

	if o.HideSsoCredential != nil {
		v := *o.HideSsoCredential

		if err = d.Set("hide_sso_credential", v); err != nil {
			return diag.Errorf("error reading hide_sso_credential: %v", err)
		}
	}

	if o.HostCheck != nil {
		v := *o.HostCheck

		if err = d.Set("host_check", v); err != nil {
			return diag.Errorf("error reading host_check: %v", err)
		}
	}

	if o.HostCheckInterval != nil {
		v := *o.HostCheckInterval

		if err = d.Set("host_check_interval", v); err != nil {
			return diag.Errorf("error reading host_check_interval: %v", err)
		}
	}

	if o.HostCheckPolicy != nil {
		if err = d.Set("host_check_policy", flattenVpnSslWebPortalHostCheckPolicy(d, o.HostCheckPolicy, "host_check_policy", sort)); err != nil {
			return diag.Errorf("error reading host_check_policy: %v", err)
		}
	}

	if o.IpMode != nil {
		v := *o.IpMode

		if err = d.Set("ip_mode", v); err != nil {
			return diag.Errorf("error reading ip_mode: %v", err)
		}
	}

	if o.IpPools != nil {
		if err = d.Set("ip_pools", flattenVpnSslWebPortalIpPools(d, o.IpPools, "ip_pools", sort)); err != nil {
			return diag.Errorf("error reading ip_pools: %v", err)
		}
	}

	if o.Ipv6DnsServer1 != nil {
		v := *o.Ipv6DnsServer1

		if err = d.Set("ipv6_dns_server1", v); err != nil {
			return diag.Errorf("error reading ipv6_dns_server1: %v", err)
		}
	}

	if o.Ipv6DnsServer2 != nil {
		v := *o.Ipv6DnsServer2

		if err = d.Set("ipv6_dns_server2", v); err != nil {
			return diag.Errorf("error reading ipv6_dns_server2: %v", err)
		}
	}

	if o.Ipv6ExclusiveRouting != nil {
		v := *o.Ipv6ExclusiveRouting

		if err = d.Set("ipv6_exclusive_routing", v); err != nil {
			return diag.Errorf("error reading ipv6_exclusive_routing: %v", err)
		}
	}

	if o.Ipv6Pools != nil {
		if err = d.Set("ipv6_pools", flattenVpnSslWebPortalIpv6Pools(d, o.Ipv6Pools, "ipv6_pools", sort)); err != nil {
			return diag.Errorf("error reading ipv6_pools: %v", err)
		}
	}

	if o.Ipv6ServiceRestriction != nil {
		v := *o.Ipv6ServiceRestriction

		if err = d.Set("ipv6_service_restriction", v); err != nil {
			return diag.Errorf("error reading ipv6_service_restriction: %v", err)
		}
	}

	if o.Ipv6SplitTunneling != nil {
		v := *o.Ipv6SplitTunneling

		if err = d.Set("ipv6_split_tunneling", v); err != nil {
			return diag.Errorf("error reading ipv6_split_tunneling: %v", err)
		}
	}

	if o.Ipv6SplitTunnelingRoutingAddress != nil {
		if err = d.Set("ipv6_split_tunneling_routing_address", flattenVpnSslWebPortalIpv6SplitTunnelingRoutingAddress(d, o.Ipv6SplitTunnelingRoutingAddress, "ipv6_split_tunneling_routing_address", sort)); err != nil {
			return diag.Errorf("error reading ipv6_split_tunneling_routing_address: %v", err)
		}
	}

	if o.Ipv6SplitTunnelingRoutingNegate != nil {
		v := *o.Ipv6SplitTunnelingRoutingNegate

		if err = d.Set("ipv6_split_tunneling_routing_negate", v); err != nil {
			return diag.Errorf("error reading ipv6_split_tunneling_routing_negate: %v", err)
		}
	}

	if o.Ipv6TunnelMode != nil {
		v := *o.Ipv6TunnelMode

		if err = d.Set("ipv6_tunnel_mode", v); err != nil {
			return diag.Errorf("error reading ipv6_tunnel_mode: %v", err)
		}
	}

	if o.Ipv6WinsServer1 != nil {
		v := *o.Ipv6WinsServer1

		if err = d.Set("ipv6_wins_server1", v); err != nil {
			return diag.Errorf("error reading ipv6_wins_server1: %v", err)
		}
	}

	if o.Ipv6WinsServer2 != nil {
		v := *o.Ipv6WinsServer2

		if err = d.Set("ipv6_wins_server2", v); err != nil {
			return diag.Errorf("error reading ipv6_wins_server2: %v", err)
		}
	}

	if o.KeepAlive != nil {
		v := *o.KeepAlive

		if err = d.Set("keep_alive", v); err != nil {
			return diag.Errorf("error reading keep_alive: %v", err)
		}
	}

	if o.LimitUserLogins != nil {
		v := *o.LimitUserLogins

		if err = d.Set("limit_user_logins", v); err != nil {
			return diag.Errorf("error reading limit_user_logins: %v", err)
		}
	}

	if o.MacAddrAction != nil {
		v := *o.MacAddrAction

		if err = d.Set("mac_addr_action", v); err != nil {
			return diag.Errorf("error reading mac_addr_action: %v", err)
		}
	}

	if o.MacAddrCheck != nil {
		v := *o.MacAddrCheck

		if err = d.Set("mac_addr_check", v); err != nil {
			return diag.Errorf("error reading mac_addr_check: %v", err)
		}
	}

	if o.MacAddrCheckRule != nil {
		if err = d.Set("mac_addr_check_rule", flattenVpnSslWebPortalMacAddrCheckRule(d, o.MacAddrCheckRule, "mac_addr_check_rule", sort)); err != nil {
			return diag.Errorf("error reading mac_addr_check_rule: %v", err)
		}
	}

	if o.MacosForticlientDownloadUrl != nil {
		v := *o.MacosForticlientDownloadUrl

		if err = d.Set("macos_forticlient_download_url", v); err != nil {
			return diag.Errorf("error reading macos_forticlient_download_url: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.OsCheck != nil {
		v := *o.OsCheck

		if err = d.Set("os_check", v); err != nil {
			return diag.Errorf("error reading os_check: %v", err)
		}
	}

	if o.OsCheckList != nil {
		if err = d.Set("os_check_list", flattenVpnSslWebPortalOsCheckList(d, o.OsCheckList, "os_check_list", sort)); err != nil {
			return diag.Errorf("error reading os_check_list: %v", err)
		}
	}

	if o.PreferIpv6Dns != nil {
		v := *o.PreferIpv6Dns

		if err = d.Set("prefer_ipv6_dns", v); err != nil {
			return diag.Errorf("error reading prefer_ipv6_dns: %v", err)
		}
	}

	if o.RedirUrl != nil {
		v := *o.RedirUrl

		if err = d.Set("redir_url", v); err != nil {
			return diag.Errorf("error reading redir_url: %v", err)
		}
	}

	if o.RewriteIpUriUi != nil {
		v := *o.RewriteIpUriUi

		if err = d.Set("rewrite_ip_uri_ui", v); err != nil {
			return diag.Errorf("error reading rewrite_ip_uri_ui: %v", err)
		}
	}

	if o.SavePassword != nil {
		v := *o.SavePassword

		if err = d.Set("save_password", v); err != nil {
			return diag.Errorf("error reading save_password: %v", err)
		}
	}

	if o.ServiceRestriction != nil {
		v := *o.ServiceRestriction

		if err = d.Set("service_restriction", v); err != nil {
			return diag.Errorf("error reading service_restriction: %v", err)
		}
	}

	if o.SkipCheckForBrowser != nil {
		v := *o.SkipCheckForBrowser

		if err = d.Set("skip_check_for_browser", v); err != nil {
			return diag.Errorf("error reading skip_check_for_browser: %v", err)
		}
	}

	if o.SkipCheckForUnsupportedOs != nil {
		v := *o.SkipCheckForUnsupportedOs

		if err = d.Set("skip_check_for_unsupported_os", v); err != nil {
			return diag.Errorf("error reading skip_check_for_unsupported_os: %v", err)
		}
	}

	if o.SmbMaxVersion != nil {
		v := *o.SmbMaxVersion

		if err = d.Set("smb_max_version", v); err != nil {
			return diag.Errorf("error reading smb_max_version: %v", err)
		}
	}

	if o.SmbMinVersion != nil {
		v := *o.SmbMinVersion

		if err = d.Set("smb_min_version", v); err != nil {
			return diag.Errorf("error reading smb_min_version: %v", err)
		}
	}

	if o.SmbNtlmv1Auth != nil {
		v := *o.SmbNtlmv1Auth

		if err = d.Set("smb_ntlmv1_auth", v); err != nil {
			return diag.Errorf("error reading smb_ntlmv1_auth: %v", err)
		}
	}

	if o.Smbv1 != nil {
		v := *o.Smbv1

		if err = d.Set("smbv1", v); err != nil {
			return diag.Errorf("error reading smbv1: %v", err)
		}
	}

	if o.SplitDns != nil {
		if err = d.Set("split_dns", flattenVpnSslWebPortalSplitDns(d, o.SplitDns, "split_dns", sort)); err != nil {
			return diag.Errorf("error reading split_dns: %v", err)
		}
	}

	if o.SplitTunneling != nil {
		v := *o.SplitTunneling

		if err = d.Set("split_tunneling", v); err != nil {
			return diag.Errorf("error reading split_tunneling: %v", err)
		}
	}

	if o.SplitTunnelingRoutingAddress != nil {
		if err = d.Set("split_tunneling_routing_address", flattenVpnSslWebPortalSplitTunnelingRoutingAddress(d, o.SplitTunnelingRoutingAddress, "split_tunneling_routing_address", sort)); err != nil {
			return diag.Errorf("error reading split_tunneling_routing_address: %v", err)
		}
	}

	if o.SplitTunnelingRoutingNegate != nil {
		v := *o.SplitTunnelingRoutingNegate

		if err = d.Set("split_tunneling_routing_negate", v); err != nil {
			return diag.Errorf("error reading split_tunneling_routing_negate: %v", err)
		}
	}

	if o.Theme != nil {
		v := *o.Theme

		if err = d.Set("theme", v); err != nil {
			return diag.Errorf("error reading theme: %v", err)
		}
	}

	if o.TransformBackwardSlashes != nil {
		v := *o.TransformBackwardSlashes

		if err = d.Set("transform_backward_slashes", v); err != nil {
			return diag.Errorf("error reading transform_backward_slashes: %v", err)
		}
	}

	if o.TunnelMode != nil {
		v := *o.TunnelMode

		if err = d.Set("tunnel_mode", v); err != nil {
			return diag.Errorf("error reading tunnel_mode: %v", err)
		}
	}

	if o.UseSdwan != nil {
		v := *o.UseSdwan

		if err = d.Set("use_sdwan", v); err != nil {
			return diag.Errorf("error reading use_sdwan: %v", err)
		}
	}

	if o.UserBookmark != nil {
		v := *o.UserBookmark

		if err = d.Set("user_bookmark", v); err != nil {
			return diag.Errorf("error reading user_bookmark: %v", err)
		}
	}

	if o.UserGroupBookmark != nil {
		v := *o.UserGroupBookmark

		if err = d.Set("user_group_bookmark", v); err != nil {
			return diag.Errorf("error reading user_group_bookmark: %v", err)
		}
	}

	if o.WebMode != nil {
		v := *o.WebMode

		if err = d.Set("web_mode", v); err != nil {
			return diag.Errorf("error reading web_mode: %v", err)
		}
	}

	if o.WindowsForticlientDownloadUrl != nil {
		v := *o.WindowsForticlientDownloadUrl

		if err = d.Set("windows_forticlient_download_url", v); err != nil {
			return diag.Errorf("error reading windows_forticlient_download_url: %v", err)
		}
	}

	if o.WinsServer1 != nil {
		v := *o.WinsServer1

		if err = d.Set("wins_server1", v); err != nil {
			return diag.Errorf("error reading wins_server1: %v", err)
		}
	}

	if o.WinsServer2 != nil {
		v := *o.WinsServer2

		if err = d.Set("wins_server2", v); err != nil {
			return diag.Errorf("error reading wins_server2: %v", err)
		}
	}

	return nil
}

func expandVpnSslWebPortalBookmarkGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslWebPortalBookmarkGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslWebPortalBookmarkGroup

	for i := range l {
		tmp := models.VpnSslWebPortalBookmarkGroup{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.bookmarks", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandVpnSslWebPortalBookmarkGroupBookmarks(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.VpnSslWebPortalBookmarkGroupBookmarks
			// 	}
			tmp.Bookmarks = v2

		}

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

func expandVpnSslWebPortalBookmarkGroupBookmarks(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslWebPortalBookmarkGroupBookmarks, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslWebPortalBookmarkGroupBookmarks

	for i := range l {
		tmp := models.VpnSslWebPortalBookmarkGroupBookmarks{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.additional_params", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AdditionalParams = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.apptype", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Apptype = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.color_depth", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ColorDepth = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.description", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Description = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.domain", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Domain = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.folder", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Folder = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.form_data", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandVpnSslWebPortalBookmarkGroupBookmarksFormData(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.VpnSslWebPortalBookmarkGroupBookmarksFormData
			// 	}
			tmp.FormData = v2

		}

		pre_append = fmt.Sprintf("%s.%d.height", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Height = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.host", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Host = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keyboard_layout", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeyboardLayout = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.listening_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ListeningPort = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.load_balancing_info", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LoadBalancingInfo = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.logon_password", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogonPassword = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.logon_user", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogonUser = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Port = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.preconnection_blob", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PreconnectionBlob = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.preconnection_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.PreconnectionId = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.remote_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.RemotePort = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.restricted_admin", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RestrictedAdmin = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.security", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Security = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.send_preconnection_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SendPreconnectionId = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.server_layout", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ServerLayout = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.show_status_window", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ShowStatusWindow = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sso", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Sso = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sso_credential", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SsoCredential = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sso_credential_sent_once", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SsoCredentialSentOnce = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sso_password", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SsoPassword = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sso_username", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SsoUsername = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.url", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Url = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.width", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Width = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksFormData(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslWebPortalBookmarkGroupBookmarksFormData, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslWebPortalBookmarkGroupBookmarksFormData

	for i := range l {
		tmp := models.VpnSslWebPortalBookmarkGroupBookmarksFormData{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.value", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Value = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandVpnSslWebPortalHostCheckPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslWebPortalHostCheckPolicy, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslWebPortalHostCheckPolicy

	for i := range l {
		tmp := models.VpnSslWebPortalHostCheckPolicy{}
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

func expandVpnSslWebPortalIpPools(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslWebPortalIpPools, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslWebPortalIpPools

	for i := range l {
		tmp := models.VpnSslWebPortalIpPools{}
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

func expandVpnSslWebPortalIpv6Pools(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslWebPortalIpv6Pools, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslWebPortalIpv6Pools

	for i := range l {
		tmp := models.VpnSslWebPortalIpv6Pools{}
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

func expandVpnSslWebPortalIpv6SplitTunnelingRoutingAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslWebPortalIpv6SplitTunnelingRoutingAddress, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslWebPortalIpv6SplitTunnelingRoutingAddress

	for i := range l {
		tmp := models.VpnSslWebPortalIpv6SplitTunnelingRoutingAddress{}
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

func expandVpnSslWebPortalMacAddrCheckRule(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslWebPortalMacAddrCheckRule, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslWebPortalMacAddrCheckRule

	for i := range l {
		tmp := models.VpnSslWebPortalMacAddrCheckRule{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.mac_addr_list", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandVpnSslWebPortalMacAddrCheckRuleMacAddrList(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.VpnSslWebPortalMacAddrCheckRuleMacAddrList
			// 	}
			tmp.MacAddrList = v2

		}

		pre_append = fmt.Sprintf("%s.%d.mac_addr_mask", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MacAddrMask = &v3
			}
		}

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

func expandVpnSslWebPortalMacAddrCheckRuleMacAddrList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslWebPortalMacAddrCheckRuleMacAddrList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslWebPortalMacAddrCheckRuleMacAddrList

	for i := range l {
		tmp := models.VpnSslWebPortalMacAddrCheckRuleMacAddrList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.addr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Addr = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandVpnSslWebPortalOsCheckList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslWebPortalOsCheckList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslWebPortalOsCheckList

	for i := range l {
		tmp := models.VpnSslWebPortalOsCheckList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.latest_patch_level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LatestPatchLevel = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tolerance", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Tolerance = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandVpnSslWebPortalSplitDns(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslWebPortalSplitDns, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslWebPortalSplitDns

	for i := range l {
		tmp := models.VpnSslWebPortalSplitDns{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dns_server1", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DnsServer1 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dns_server2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DnsServer2 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.domains", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Domains = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv6_dns_server1", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv6DnsServer1 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv6_dns_server2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv6DnsServer2 = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandVpnSslWebPortalSplitTunnelingRoutingAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslWebPortalSplitTunnelingRoutingAddress, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslWebPortalSplitTunnelingRoutingAddress

	for i := range l {
		tmp := models.VpnSslWebPortalSplitTunnelingRoutingAddress{}
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

func getObjectVpnSslWebPortal(d *schema.ResourceData, sv string) (*models.VpnSslWebPortal, diag.Diagnostics) {
	obj := models.VpnSslWebPortal{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("allow_user_access"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("allow_user_access", sv)
				diags = append(diags, e)
			}
			obj.AllowUserAccess = &v2
		}
	}
	if v1, ok := d.GetOk("auto_connect"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_connect", sv)
				diags = append(diags, e)
			}
			obj.AutoConnect = &v2
		}
	}
	if v, ok := d.GetOk("bookmark_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("bookmark_group", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnSslWebPortalBookmarkGroup(d, v, "bookmark_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.BookmarkGroup = t
		}
	} else if d.HasChange("bookmark_group") {
		old, new := d.GetChange("bookmark_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.BookmarkGroup = &[]models.VpnSslWebPortalBookmarkGroup{}
		}
	}
	if v1, ok := d.GetOk("clipboard"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.7", "v7.0.0") {
				e := utils.AttributeVersionWarning("clipboard", sv)
				diags = append(diags, e)
			}
			obj.Clipboard = &v2
		}
	}
	if v1, ok := d.GetOk("custom_lang"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("custom_lang", sv)
				diags = append(diags, e)
			}
			obj.CustomLang = &v2
		}
	}
	if v1, ok := d.GetOk("customize_forticlient_download_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("customize_forticlient_download_url", sv)
				diags = append(diags, e)
			}
			obj.CustomizeForticlientDownloadUrl = &v2
		}
	}
	if v1, ok := d.GetOk("display_bookmark"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("display_bookmark", sv)
				diags = append(diags, e)
			}
			obj.DisplayBookmark = &v2
		}
	}
	if v1, ok := d.GetOk("display_connection_tools"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("display_connection_tools", sv)
				diags = append(diags, e)
			}
			obj.DisplayConnectionTools = &v2
		}
	}
	if v1, ok := d.GetOk("display_history"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("display_history", sv)
				diags = append(diags, e)
			}
			obj.DisplayHistory = &v2
		}
	}
	if v1, ok := d.GetOk("display_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("display_status", sv)
				diags = append(diags, e)
			}
			obj.DisplayStatus = &v2
		}
	}
	if v1, ok := d.GetOk("dns_server1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_server1", sv)
				diags = append(diags, e)
			}
			obj.DnsServer1 = &v2
		}
	}
	if v1, ok := d.GetOk("dns_server2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_server2", sv)
				diags = append(diags, e)
			}
			obj.DnsServer2 = &v2
		}
	}
	if v1, ok := d.GetOk("dns_suffix"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_suffix", sv)
				diags = append(diags, e)
			}
			obj.DnsSuffix = &v2
		}
	}
	if v1, ok := d.GetOk("exclusive_routing"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("exclusive_routing", sv)
				diags = append(diags, e)
			}
			obj.ExclusiveRouting = &v2
		}
	}
	if v1, ok := d.GetOk("forticlient_download"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("forticlient_download", sv)
				diags = append(diags, e)
			}
			obj.ForticlientDownload = &v2
		}
	}
	if v1, ok := d.GetOk("forticlient_download_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("forticlient_download_method", sv)
				diags = append(diags, e)
			}
			obj.ForticlientDownloadMethod = &v2
		}
	}
	if v1, ok := d.GetOk("heading"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("heading", sv)
				diags = append(diags, e)
			}
			obj.Heading = &v2
		}
	}
	if v1, ok := d.GetOk("hide_sso_credential"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hide_sso_credential", sv)
				diags = append(diags, e)
			}
			obj.HideSsoCredential = &v2
		}
	}
	if v1, ok := d.GetOk("host_check"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("host_check", sv)
				diags = append(diags, e)
			}
			obj.HostCheck = &v2
		}
	}
	if v1, ok := d.GetOk("host_check_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("host_check_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HostCheckInterval = &tmp
		}
	}
	if v, ok := d.GetOk("host_check_policy"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("host_check_policy", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnSslWebPortalHostCheckPolicy(d, v, "host_check_policy", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.HostCheckPolicy = t
		}
	} else if d.HasChange("host_check_policy") {
		old, new := d.GetChange("host_check_policy")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.HostCheckPolicy = &[]models.VpnSslWebPortalHostCheckPolicy{}
		}
	}
	if v1, ok := d.GetOk("ip_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip_mode", sv)
				diags = append(diags, e)
			}
			obj.IpMode = &v2
		}
	}
	if v, ok := d.GetOk("ip_pools"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ip_pools", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnSslWebPortalIpPools(d, v, "ip_pools", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.IpPools = t
		}
	} else if d.HasChange("ip_pools") {
		old, new := d.GetChange("ip_pools")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.IpPools = &[]models.VpnSslWebPortalIpPools{}
		}
	}
	if v1, ok := d.GetOk("ipv6_dns_server1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_dns_server1", sv)
				diags = append(diags, e)
			}
			obj.Ipv6DnsServer1 = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6_dns_server2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_dns_server2", sv)
				diags = append(diags, e)
			}
			obj.Ipv6DnsServer2 = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6_exclusive_routing"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_exclusive_routing", sv)
				diags = append(diags, e)
			}
			obj.Ipv6ExclusiveRouting = &v2
		}
	}
	if v, ok := d.GetOk("ipv6_pools"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ipv6_pools", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnSslWebPortalIpv6Pools(d, v, "ipv6_pools", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Ipv6Pools = t
		}
	} else if d.HasChange("ipv6_pools") {
		old, new := d.GetChange("ipv6_pools")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Ipv6Pools = &[]models.VpnSslWebPortalIpv6Pools{}
		}
	}
	if v1, ok := d.GetOk("ipv6_service_restriction"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_service_restriction", sv)
				diags = append(diags, e)
			}
			obj.Ipv6ServiceRestriction = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6_split_tunneling"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_split_tunneling", sv)
				diags = append(diags, e)
			}
			obj.Ipv6SplitTunneling = &v2
		}
	}
	if v, ok := d.GetOk("ipv6_split_tunneling_routing_address"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ipv6_split_tunneling_routing_address", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnSslWebPortalIpv6SplitTunnelingRoutingAddress(d, v, "ipv6_split_tunneling_routing_address", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Ipv6SplitTunnelingRoutingAddress = t
		}
	} else if d.HasChange("ipv6_split_tunneling_routing_address") {
		old, new := d.GetChange("ipv6_split_tunneling_routing_address")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Ipv6SplitTunnelingRoutingAddress = &[]models.VpnSslWebPortalIpv6SplitTunnelingRoutingAddress{}
		}
	}
	if v1, ok := d.GetOk("ipv6_split_tunneling_routing_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("ipv6_split_tunneling_routing_negate", sv)
				diags = append(diags, e)
			}
			obj.Ipv6SplitTunnelingRoutingNegate = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6_tunnel_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_tunnel_mode", sv)
				diags = append(diags, e)
			}
			obj.Ipv6TunnelMode = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6_wins_server1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_wins_server1", sv)
				diags = append(diags, e)
			}
			obj.Ipv6WinsServer1 = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6_wins_server2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_wins_server2", sv)
				diags = append(diags, e)
			}
			obj.Ipv6WinsServer2 = &v2
		}
	}
	if v1, ok := d.GetOk("keep_alive"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("keep_alive", sv)
				diags = append(diags, e)
			}
			obj.KeepAlive = &v2
		}
	}
	if v1, ok := d.GetOk("limit_user_logins"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("limit_user_logins", sv)
				diags = append(diags, e)
			}
			obj.LimitUserLogins = &v2
		}
	}
	if v1, ok := d.GetOk("mac_addr_action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mac_addr_action", sv)
				diags = append(diags, e)
			}
			obj.MacAddrAction = &v2
		}
	}
	if v1, ok := d.GetOk("mac_addr_check"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mac_addr_check", sv)
				diags = append(diags, e)
			}
			obj.MacAddrCheck = &v2
		}
	}
	if v, ok := d.GetOk("mac_addr_check_rule"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("mac_addr_check_rule", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnSslWebPortalMacAddrCheckRule(d, v, "mac_addr_check_rule", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.MacAddrCheckRule = t
		}
	} else if d.HasChange("mac_addr_check_rule") {
		old, new := d.GetChange("mac_addr_check_rule")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.MacAddrCheckRule = &[]models.VpnSslWebPortalMacAddrCheckRule{}
		}
	}
	if v1, ok := d.GetOk("macos_forticlient_download_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("macos_forticlient_download_url", sv)
				diags = append(diags, e)
			}
			obj.MacosForticlientDownloadUrl = &v2
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
	if v1, ok := d.GetOk("os_check"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("os_check", sv)
				diags = append(diags, e)
			}
			obj.OsCheck = &v2
		}
	}
	if v, ok := d.GetOk("os_check_list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("os_check_list", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnSslWebPortalOsCheckList(d, v, "os_check_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.OsCheckList = t
		}
	} else if d.HasChange("os_check_list") {
		old, new := d.GetChange("os_check_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.OsCheckList = &[]models.VpnSslWebPortalOsCheckList{}
		}
	}
	if v1, ok := d.GetOk("prefer_ipv6_dns"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("prefer_ipv6_dns", sv)
				diags = append(diags, e)
			}
			obj.PreferIpv6Dns = &v2
		}
	}
	if v1, ok := d.GetOk("redir_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("redir_url", sv)
				diags = append(diags, e)
			}
			obj.RedirUrl = &v2
		}
	}
	if v1, ok := d.GetOk("rewrite_ip_uri_ui"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("rewrite_ip_uri_ui", sv)
				diags = append(diags, e)
			}
			obj.RewriteIpUriUi = &v2
		}
	}
	if v1, ok := d.GetOk("save_password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("save_password", sv)
				diags = append(diags, e)
			}
			obj.SavePassword = &v2
		}
	}
	if v1, ok := d.GetOk("service_restriction"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("service_restriction", sv)
				diags = append(diags, e)
			}
			obj.ServiceRestriction = &v2
		}
	}
	if v1, ok := d.GetOk("skip_check_for_browser"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("skip_check_for_browser", sv)
				diags = append(diags, e)
			}
			obj.SkipCheckForBrowser = &v2
		}
	}
	if v1, ok := d.GetOk("skip_check_for_unsupported_os"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("skip_check_for_unsupported_os", sv)
				diags = append(diags, e)
			}
			obj.SkipCheckForUnsupportedOs = &v2
		}
	}
	if v1, ok := d.GetOk("smb_max_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("smb_max_version", sv)
				diags = append(diags, e)
			}
			obj.SmbMaxVersion = &v2
		}
	}
	if v1, ok := d.GetOk("smb_min_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("smb_min_version", sv)
				diags = append(diags, e)
			}
			obj.SmbMinVersion = &v2
		}
	}
	if v1, ok := d.GetOk("smb_ntlmv1_auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("smb_ntlmv1_auth", sv)
				diags = append(diags, e)
			}
			obj.SmbNtlmv1Auth = &v2
		}
	}
	if v1, ok := d.GetOk("smbv1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("smbv1", sv)
				diags = append(diags, e)
			}
			obj.Smbv1 = &v2
		}
	}
	if v, ok := d.GetOk("split_dns"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("split_dns", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnSslWebPortalSplitDns(d, v, "split_dns", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SplitDns = t
		}
	} else if d.HasChange("split_dns") {
		old, new := d.GetChange("split_dns")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SplitDns = &[]models.VpnSslWebPortalSplitDns{}
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
	if v, ok := d.GetOk("split_tunneling_routing_address"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("split_tunneling_routing_address", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnSslWebPortalSplitTunnelingRoutingAddress(d, v, "split_tunneling_routing_address", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SplitTunnelingRoutingAddress = t
		}
	} else if d.HasChange("split_tunneling_routing_address") {
		old, new := d.GetChange("split_tunneling_routing_address")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SplitTunnelingRoutingAddress = &[]models.VpnSslWebPortalSplitTunnelingRoutingAddress{}
		}
	}
	if v1, ok := d.GetOk("split_tunneling_routing_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("split_tunneling_routing_negate", sv)
				diags = append(diags, e)
			}
			obj.SplitTunnelingRoutingNegate = &v2
		}
	}
	if v1, ok := d.GetOk("theme"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("theme", sv)
				diags = append(diags, e)
			}
			obj.Theme = &v2
		}
	}
	if v1, ok := d.GetOk("transform_backward_slashes"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("transform_backward_slashes", sv)
				diags = append(diags, e)
			}
			obj.TransformBackwardSlashes = &v2
		}
	}
	if v1, ok := d.GetOk("tunnel_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tunnel_mode", sv)
				diags = append(diags, e)
			}
			obj.TunnelMode = &v2
		}
	}
	if v1, ok := d.GetOk("use_sdwan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("use_sdwan", sv)
				diags = append(diags, e)
			}
			obj.UseSdwan = &v2
		}
	}
	if v1, ok := d.GetOk("user_bookmark"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("user_bookmark", sv)
				diags = append(diags, e)
			}
			obj.UserBookmark = &v2
		}
	}
	if v1, ok := d.GetOk("user_group_bookmark"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("user_group_bookmark", sv)
				diags = append(diags, e)
			}
			obj.UserGroupBookmark = &v2
		}
	}
	if v1, ok := d.GetOk("web_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("web_mode", sv)
				diags = append(diags, e)
			}
			obj.WebMode = &v2
		}
	}
	if v1, ok := d.GetOk("windows_forticlient_download_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("windows_forticlient_download_url", sv)
				diags = append(diags, e)
			}
			obj.WindowsForticlientDownloadUrl = &v2
		}
	}
	if v1, ok := d.GetOk("wins_server1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wins_server1", sv)
				diags = append(diags, e)
			}
			obj.WinsServer1 = &v2
		}
	}
	if v1, ok := d.GetOk("wins_server2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wins_server2", sv)
				diags = append(diags, e)
			}
			obj.WinsServer2 = &v2
		}
	}
	return &obj, diags
}
