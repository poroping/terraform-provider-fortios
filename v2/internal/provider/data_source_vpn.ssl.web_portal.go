// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Portal.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceVpnSslWebPortal() *schema.Resource {
	return &schema.Resource{
		Description: "Portal.",

		ReadContext: dataSourceVpnSslWebPortalRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"allow_user_access": {
				Type:        schema.TypeString,
				Description: "Allow user access to SSL-VPN applications.",
				Computed:    true,
			},
			"auto_connect": {
				Type:        schema.TypeString,
				Description: "Enable/disable automatic connect by client when system is up.",
				Computed:    true,
			},
			"bookmark_group": {
				Type:        schema.TypeList,
				Description: "Portal bookmark group.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bookmarks": {
							Type:        schema.TypeList,
							Description: "Bookmark table.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_params": {
										Type:        schema.TypeString,
										Description: "Additional parameters.",
										Computed:    true,
									},
									"apptype": {
										Type:        schema.TypeString,
										Description: "Application type.",
										Computed:    true,
									},
									"color_depth": {
										Type:        schema.TypeString,
										Description: "Color depth per pixel.",
										Computed:    true,
									},
									"description": {
										Type:        schema.TypeString,
										Description: "Description.",
										Computed:    true,
									},
									"domain": {
										Type:        schema.TypeString,
										Description: "Login domain.",
										Computed:    true,
									},
									"folder": {
										Type:        schema.TypeString,
										Description: "Network shared file folder parameter.",
										Computed:    true,
									},
									"form_data": {
										Type:        schema.TypeList,
										Description: "Form data.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Description: "Name.",
													Computed:    true,
												},
												"value": {
													Type:        schema.TypeString,
													Description: "Value.",
													Computed:    true,
												},
											},
										},
									},
									"height": {
										Type:        schema.TypeInt,
										Description: "Screen height (range from 0 - 65535, default = 0).",
										Computed:    true,
									},
									"host": {
										Type:        schema.TypeString,
										Description: "Host name/IP parameter.",
										Computed:    true,
									},
									"keyboard_layout": {
										Type:        schema.TypeString,
										Description: "Keyboard layout.",
										Computed:    true,
									},
									"listening_port": {
										Type:        schema.TypeInt,
										Description: "Listening port (0 - 65535).",
										Computed:    true,
									},
									"load_balancing_info": {
										Type:        schema.TypeString,
										Description: "The load balancing information or cookie which should be provided to the connection broker.",
										Computed:    true,
									},
									"logon_password": {
										Type:        schema.TypeString,
										Description: "Logon password.",
										Computed:    true,
										Sensitive:   true,
									},
									"logon_user": {
										Type:        schema.TypeString,
										Description: "Logon user.",
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "Bookmark name.",
										Computed:    true,
									},
									"port": {
										Type:        schema.TypeInt,
										Description: "Remote port.",
										Computed:    true,
									},
									"preconnection_blob": {
										Type:        schema.TypeString,
										Description: "An arbitrary string which identifies the RDP source.",
										Computed:    true,
									},
									"preconnection_id": {
										Type:        schema.TypeInt,
										Description: "The numeric ID of the RDP source (0-4294967295).",
										Computed:    true,
									},
									"remote_port": {
										Type:        schema.TypeInt,
										Description: "Remote port (0 - 65535).",
										Computed:    true,
									},
									"restricted_admin": {
										Type:        schema.TypeString,
										Description: "Enable/disable restricted admin mode for RDP.",
										Computed:    true,
									},
									"security": {
										Type:        schema.TypeString,
										Description: "Security mode for RDP connection.",
										Computed:    true,
									},
									"send_preconnection_id": {
										Type:        schema.TypeString,
										Description: "Enable/disable sending of preconnection ID.",
										Computed:    true,
									},
									"server_layout": {
										Type:        schema.TypeString,
										Description: "Server side keyboard layout.",
										Computed:    true,
									},
									"show_status_window": {
										Type:        schema.TypeString,
										Description: "Enable/disable showing of status window.",
										Computed:    true,
									},
									"sso": {
										Type:        schema.TypeString,
										Description: "Single Sign-On.",
										Computed:    true,
									},
									"sso_credential": {
										Type:        schema.TypeString,
										Description: "Single sign-on credentials.",
										Computed:    true,
									},
									"sso_credential_sent_once": {
										Type:        schema.TypeString,
										Description: "Single sign-on credentials are only sent once to remote server.",
										Computed:    true,
									},
									"sso_password": {
										Type:        schema.TypeString,
										Description: "SSO password.",
										Computed:    true,
										Sensitive:   true,
									},
									"sso_username": {
										Type:        schema.TypeString,
										Description: "SSO user name.",
										Computed:    true,
									},
									"url": {
										Type:        schema.TypeString,
										Description: "URL parameter.",
										Computed:    true,
									},
									"width": {
										Type:        schema.TypeInt,
										Description: "Screen width (range from 0 - 65535, default = 0).",
										Computed:    true,
									},
								},
							},
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Bookmark group name.",
							Computed:    true,
						},
					},
				},
			},
			"clipboard": {
				Type:        schema.TypeString,
				Description: "Enable to support RDP/VPC clipboard functionality.",
				Computed:    true,
			},
			"custom_lang": {
				Type:        schema.TypeString,
				Description: "Change the web portal display language. Overrides config system global set language. You can use config system custom-language and execute system custom-language to add custom language files.",
				Computed:    true,
			},
			"customize_forticlient_download_url": {
				Type:        schema.TypeString,
				Description: "Enable support of customized download URL for FortiClient.",
				Computed:    true,
			},
			"default_window_height": {
				Type:        schema.TypeInt,
				Description: "Screen height (range from 0 - 65535, default = 768).",
				Computed:    true,
			},
			"default_window_width": {
				Type:        schema.TypeInt,
				Description: "Screen width (range from 0 - 65535, default = 1024).",
				Computed:    true,
			},
			"display_bookmark": {
				Type:        schema.TypeString,
				Description: "Enable to display the web portal bookmark widget.",
				Computed:    true,
			},
			"display_connection_tools": {
				Type:        schema.TypeString,
				Description: "Enable to display the web portal connection tools widget.",
				Computed:    true,
			},
			"display_history": {
				Type:        schema.TypeString,
				Description: "Enable to display the web portal user login history widget.",
				Computed:    true,
			},
			"display_status": {
				Type:        schema.TypeString,
				Description: "Enable to display the web portal status widget.",
				Computed:    true,
			},
			"dns_server1": {
				Type:        schema.TypeString,
				Description: "IPv4 DNS server 1.",
				Computed:    true,
			},
			"dns_server2": {
				Type:        schema.TypeString,
				Description: "IPv4 DNS server 2.",
				Computed:    true,
			},
			"dns_suffix": {
				Type:        schema.TypeString,
				Description: "DNS suffix.",
				Computed:    true,
			},
			"exclusive_routing": {
				Type:        schema.TypeString,
				Description: "Enable/disable all traffic go through tunnel only.",
				Computed:    true,
			},
			"forticlient_download": {
				Type:        schema.TypeString,
				Description: "Enable/disable download option for FortiClient.",
				Computed:    true,
			},
			"forticlient_download_method": {
				Type:        schema.TypeString,
				Description: "FortiClient download method.",
				Computed:    true,
			},
			"heading": {
				Type:        schema.TypeString,
				Description: "Web portal heading message.",
				Computed:    true,
			},
			"hide_sso_credential": {
				Type:        schema.TypeString,
				Description: "Enable to prevent SSO credential being sent to client.",
				Computed:    true,
			},
			"host_check": {
				Type:        schema.TypeString,
				Description: "Type of host checking performed on endpoints.",
				Computed:    true,
			},
			"host_check_interval": {
				Type:        schema.TypeInt,
				Description: "Periodic host check interval. Value of 0 means disabled and host checking only happens when the endpoint connects.",
				Computed:    true,
			},
			"host_check_policy": {
				Type:        schema.TypeList,
				Description: "One or more policies to require the endpoint to have specific security software.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Host check software list name.",
							Computed:    true,
						},
					},
				},
			},
			"ip_mode": {
				Type:        schema.TypeString,
				Description: "Method by which users of this SSL-VPN tunnel obtain IP addresses.",
				Computed:    true,
			},
			"ip_pools": {
				Type:        schema.TypeList,
				Description: "IPv4 firewall source address objects reserved for SSL-VPN tunnel mode clients.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address name.",
							Computed:    true,
						},
					},
				},
			},
			"ipv6_dns_server1": {
				Type:        schema.TypeString,
				Description: "IPv6 DNS server 1.",
				Computed:    true,
			},
			"ipv6_dns_server2": {
				Type:        schema.TypeString,
				Description: "IPv6 DNS server 2.",
				Computed:    true,
			},
			"ipv6_exclusive_routing": {
				Type:        schema.TypeString,
				Description: "Enable/disable all IPv6 traffic go through tunnel only.",
				Computed:    true,
			},
			"ipv6_pools": {
				Type:        schema.TypeList,
				Description: "IPv6 firewall source address objects reserved for SSL-VPN tunnel mode clients.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address name.",
							Computed:    true,
						},
					},
				},
			},
			"ipv6_service_restriction": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv6 tunnel service restriction.",
				Computed:    true,
			},
			"ipv6_split_tunneling": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv6 split tunneling.",
				Computed:    true,
			},
			"ipv6_split_tunneling_routing_address": {
				Type:        schema.TypeList,
				Description: "IPv6 SSL-VPN tunnel mode firewall address objects that override firewall policy destination addresses to control split-tunneling access.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address name.",
							Computed:    true,
						},
					},
				},
			},
			"ipv6_split_tunneling_routing_negate": {
				Type:        schema.TypeString,
				Description: "Enable to negate IPv6 split tunneling routing address.",
				Computed:    true,
			},
			"ipv6_tunnel_mode": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv6 SSL-VPN tunnel mode.",
				Computed:    true,
			},
			"ipv6_wins_server1": {
				Type:        schema.TypeString,
				Description: "IPv6 WINS server 1.",
				Computed:    true,
			},
			"ipv6_wins_server2": {
				Type:        schema.TypeString,
				Description: "IPv6 WINS server 2.",
				Computed:    true,
			},
			"keep_alive": {
				Type:        schema.TypeString,
				Description: "Enable/disable automatic reconnect for FortiClient connections.",
				Computed:    true,
			},
			"limit_user_logins": {
				Type:        schema.TypeString,
				Description: "Enable to limit each user to one SSL-VPN session at a time.",
				Computed:    true,
			},
			"mac_addr_action": {
				Type:        schema.TypeString,
				Description: "Client MAC address action.",
				Computed:    true,
			},
			"mac_addr_check": {
				Type:        schema.TypeString,
				Description: "Enable/disable MAC address host checking.",
				Computed:    true,
			},
			"mac_addr_check_rule": {
				Type:        schema.TypeList,
				Description: "Client MAC address check rule.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mac_addr_list": {
							Type:        schema.TypeList,
							Description: "Client MAC address list.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"addr": {
										Type:        schema.TypeString,
										Description: "Client MAC address.",
										Computed:    true,
									},
								},
							},
						},
						"mac_addr_mask": {
							Type:        schema.TypeInt,
							Description: "Client MAC address mask.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Client MAC address check rule name.",
							Computed:    true,
						},
					},
				},
			},
			"macos_forticlient_download_url": {
				Type:        schema.TypeString,
				Description: "Download URL for Mac FortiClient.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Portal name.",
				Required:    true,
			},
			"os_check": {
				Type:        schema.TypeString,
				Description: "Enable to let the FortiGate decide action based on client OS.",
				Computed:    true,
			},
			"os_check_list": {
				Type:        schema.TypeList,
				Description: "SSL-VPN OS checks.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "OS check options.",
							Computed:    true,
						},
						"latest_patch_level": {
							Type:        schema.TypeString,
							Description: "Latest OS patch level.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Name.",
							Computed:    true,
						},
						"tolerance": {
							Type:        schema.TypeInt,
							Description: "OS patch level tolerance.",
							Computed:    true,
						},
					},
				},
			},
			"prefer_ipv6_dns": {
				Type:        schema.TypeString,
				Description: "Prefer to query IPv6 DNS server first if enabled.",
				Computed:    true,
			},
			"redir_url": {
				Type:        schema.TypeString,
				Description: "Client login redirect URL.",
				Computed:    true,
			},
			"rewrite_ip_uri_ui": {
				Type:        schema.TypeString,
				Description: "Rewrite contents for URI contains IP and /ui/ (default = disable).",
				Computed:    true,
			},
			"save_password": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiClient saving the user's password.",
				Computed:    true,
			},
			"service_restriction": {
				Type:        schema.TypeString,
				Description: "Enable/disable tunnel service restriction.",
				Computed:    true,
			},
			"skip_check_for_browser": {
				Type:        schema.TypeString,
				Description: "Enable to skip host check for browser support.",
				Computed:    true,
			},
			"skip_check_for_unsupported_os": {
				Type:        schema.TypeString,
				Description: "Enable to skip host check if client OS does not support it.",
				Computed:    true,
			},
			"smb_max_version": {
				Type:        schema.TypeString,
				Description: "SMB maximum client protocol version.",
				Computed:    true,
			},
			"smb_min_version": {
				Type:        schema.TypeString,
				Description: "SMB minimum client protocol version.",
				Computed:    true,
			},
			"smb_ntlmv1_auth": {
				Type:        schema.TypeString,
				Description: "Enable support of NTLMv1 for Samba authentication.",
				Computed:    true,
			},
			"smbv1": {
				Type:        schema.TypeString,
				Description: "SMB version 1.",
				Computed:    true,
			},
			"split_dns": {
				Type:        schema.TypeList,
				Description: "Split DNS for SSL-VPN.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_server1": {
							Type:        schema.TypeString,
							Description: "DNS server 1.",
							Computed:    true,
						},
						"dns_server2": {
							Type:        schema.TypeString,
							Description: "DNS server 2.",
							Computed:    true,
						},
						"domains": {
							Type:        schema.TypeString,
							Description: "Split DNS domains used for SSL-VPN clients separated by comma.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"ipv6_dns_server1": {
							Type:        schema.TypeString,
							Description: "IPv6 DNS server 1.",
							Computed:    true,
						},
						"ipv6_dns_server2": {
							Type:        schema.TypeString,
							Description: "IPv6 DNS server 2.",
							Computed:    true,
						},
					},
				},
			},
			"split_tunneling": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv4 split tunneling.",
				Computed:    true,
			},
			"split_tunneling_routing_address": {
				Type:        schema.TypeList,
				Description: "IPv4 SSL-VPN tunnel mode firewall address objects that override firewall policy destination addresses to control split-tunneling access.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address name.",
							Computed:    true,
						},
					},
				},
			},
			"split_tunneling_routing_negate": {
				Type:        schema.TypeString,
				Description: "Enable to negate split tunneling routing address.",
				Computed:    true,
			},
			"theme": {
				Type:        schema.TypeString,
				Description: "Web portal color scheme.",
				Computed:    true,
			},
			"transform_backward_slashes": {
				Type:        schema.TypeString,
				Description: "Transform backward slashes to forward slashes in URLs.",
				Computed:    true,
			},
			"tunnel_mode": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv4 SSL-VPN tunnel mode.",
				Computed:    true,
			},
			"use_sdwan": {
				Type:        schema.TypeString,
				Description: "Use SD-WAN rules to get output interface.",
				Computed:    true,
			},
			"user_bookmark": {
				Type:        schema.TypeString,
				Description: "Enable to allow web portal users to create their own bookmarks.",
				Computed:    true,
			},
			"user_group_bookmark": {
				Type:        schema.TypeString,
				Description: "Enable to allow web portal users to create bookmarks for all users in the same user group.",
				Computed:    true,
			},
			"web_mode": {
				Type:        schema.TypeString,
				Description: "Enable/disable SSL-VPN web mode.",
				Computed:    true,
			},
			"windows_forticlient_download_url": {
				Type:        schema.TypeString,
				Description: "Download URL for Windows FortiClient.",
				Computed:    true,
			},
			"wins_server1": {
				Type:        schema.TypeString,
				Description: "IPv4 WINS server 1.",
				Computed:    true,
			},
			"wins_server2": {
				Type:        schema.TypeString,
				Description: "IPv4 WINS server 1.",
				Computed:    true,
			},
		},
	}
}

func dataSourceVpnSslWebPortalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnSslWebPortal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnSslWebPortal dataSource: %v", err)
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

	diags := refreshObjectVpnSslWebPortal(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
