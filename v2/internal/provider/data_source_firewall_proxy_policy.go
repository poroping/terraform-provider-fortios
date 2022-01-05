// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure proxy policies.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallProxyPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure proxy policies.",

		ReadContext: dataSourceFirewallProxyPolicyRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"access_proxy": {
				Type:        schema.TypeList,
				Description: "IPv4 access proxy.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Access Proxy name.",
							Computed:    true,
						},
					},
				},
			},
			"access_proxy6": {
				Type:        schema.TypeList,
				Description: "IPv6 access proxy.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Access proxy name.",
							Computed:    true,
						},
					},
				},
			},
			"action": {
				Type:        schema.TypeString,
				Description: "Accept or deny traffic matching the policy parameters.",
				Computed:    true,
			},
			"application_list": {
				Type:        schema.TypeString,
				Description: "Name of an existing Application list.",
				Computed:    true,
			},
			"av_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing Antivirus profile.",
				Computed:    true,
			},
			"cifs_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing CIFS profile.",
				Computed:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Optional comments.",
				Computed:    true,
			},
			"decrypted_traffic_mirror": {
				Type:        schema.TypeString,
				Description: "Decrypted traffic mirror.",
				Computed:    true,
			},
			"device_ownership": {
				Type:        schema.TypeString,
				Description: "When enabled, the ownership enforcement will be done at policy level.",
				Computed:    true,
			},
			"disclaimer": {
				Type:        schema.TypeString,
				Description: "Web proxy disclaimer setting: by domain, policy, or user.",
				Computed:    true,
			},
			"dlp_sensor": {
				Type:        schema.TypeString,
				Description: "Name of an existing DLP sensor.",
				Computed:    true,
			},
			"dstaddr": {
				Type:        schema.TypeList,
				Description: "Destination address objects.",
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
			"dstaddr_negate": {
				Type:        schema.TypeString,
				Description: "When enabled, destination addresses match against any address EXCEPT the specified destination addresses.",
				Computed:    true,
			},
			"dstaddr6": {
				Type:        schema.TypeList,
				Description: "IPv6 destination address objects.",
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
			"dstintf": {
				Type:        schema.TypeList,
				Description: "Destination interface names.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
					},
				},
			},
			"emailfilter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing email filter profile.",
				Computed:    true,
			},
			"file_filter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing file-filter profile.",
				Computed:    true,
			},
			"groups": {
				Type:        schema.TypeList,
				Description: "Names of group objects.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Group name.",
							Computed:    true,
						},
					},
				},
			},
			"http_tunnel_auth": {
				Type:        schema.TypeString,
				Description: "Enable/disable HTTP tunnel authentication.",
				Computed:    true,
			},
			"icap_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing ICAP profile.",
				Computed:    true,
			},
			"internet_service": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.",
				Computed:    true,
			},
			"internet_service_custom": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Custom Internet Service name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_custom_group": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service group name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Custom Internet Service group name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_group": {
				Type:        schema.TypeList,
				Description: "Internet Service group name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Internet Service group name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_id": {
				Type:        schema.TypeList,
				Description: "Internet Service ID.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Internet Service ID.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_name": {
				Type:        schema.TypeList,
				Description: "Internet Service name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Internet Service name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_negate": {
				Type:        schema.TypeString,
				Description: "When enabled, Internet Services match against any internet service EXCEPT the selected Internet Service.",
				Computed:    true,
			},
			"ips_sensor": {
				Type:        schema.TypeString,
				Description: "Name of an existing IPS sensor.",
				Computed:    true,
			},
			"logtraffic": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging traffic through the policy.",
				Computed:    true,
			},
			"logtraffic_start": {
				Type:        schema.TypeString,
				Description: "Enable/disable policy log traffic start.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Policy name.",
				Computed:    true,
			},
			"policyid": {
				Type:        schema.TypeInt,
				Description: "Policy ID.",
				Required:    true,
			},
			"poolname": {
				Type:        schema.TypeList,
				Description: "Name of IP pool object.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "IP pool name.",
							Computed:    true,
						},
					},
				},
			},
			"profile_group": {
				Type:        schema.TypeString,
				Description: "Name of profile group.",
				Computed:    true,
			},
			"profile_protocol_options": {
				Type:        schema.TypeString,
				Description: "Name of an existing Protocol options profile.",
				Computed:    true,
			},
			"profile_type": {
				Type:        schema.TypeString,
				Description: "Determine whether the firewall policy allows security profile groups or single profiles only.",
				Computed:    true,
			},
			"proxy": {
				Type:        schema.TypeString,
				Description: "Type of explicit proxy.",
				Computed:    true,
			},
			"redirect_url": {
				Type:        schema.TypeString,
				Description: "Redirect URL for further explicit web proxy processing.",
				Computed:    true,
			},
			"replacemsg_override_group": {
				Type:        schema.TypeString,
				Description: "Authentication replacement message override group.",
				Computed:    true,
			},
			"schedule": {
				Type:        schema.TypeString,
				Description: "Name of schedule object.",
				Computed:    true,
			},
			"sctp_filter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing SCTP filter profile.",
				Computed:    true,
			},
			"service": {
				Type:        schema.TypeList,
				Description: "Name of service objects.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Service name.",
							Computed:    true,
						},
					},
				},
			},
			"service_negate": {
				Type:        schema.TypeString,
				Description: "When enabled, services match against any service EXCEPT the specified destination services.",
				Computed:    true,
			},
			"session_ttl": {
				Type:        schema.TypeInt,
				Description: "TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).",
				Computed:    true,
			},
			"srcaddr": {
				Type:        schema.TypeList,
				Description: "Source address objects.",
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
			"srcaddr_negate": {
				Type:        schema.TypeString,
				Description: "When enabled, source addresses match against any address EXCEPT the specified source addresses.",
				Computed:    true,
			},
			"srcaddr6": {
				Type:        schema.TypeList,
				Description: "IPv6 source address objects.",
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
			"srcintf": {
				Type:        schema.TypeList,
				Description: "Source interface names.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
					},
				},
			},
			"ssh_filter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing SSH filter profile.",
				Computed:    true,
			},
			"ssh_policy_redirect": {
				Type:        schema.TypeString,
				Description: "Redirect SSH traffic to matching transparent proxy policy.",
				Computed:    true,
			},
			"ssl_ssh_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing SSL SSH profile.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable the active status of the policy.",
				Computed:    true,
			},
			"transparent": {
				Type:        schema.TypeString,
				Description: "Enable to use the IP address of the client to connect to the server.",
				Computed:    true,
			},
			"users": {
				Type:        schema.TypeList,
				Description: "Names of user objects.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Group name.",
							Computed:    true,
						},
					},
				},
			},
			"utm_status": {
				Type:        schema.TypeString,
				Description: "Enable the use of UTM profiles/sensors/lists.",
				Computed:    true,
			},
			"uuid": {
				Type:        schema.TypeString,
				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Computed:    true,
			},
			"videofilter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing VideoFilter profile.",
				Computed:    true,
			},
			"voip_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing VoIP profile.",
				Computed:    true,
			},
			"waf_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing Web application firewall profile.",
				Computed:    true,
			},
			"webcache": {
				Type:        schema.TypeString,
				Description: "Enable/disable web caching.",
				Computed:    true,
			},
			"webcache_https": {
				Type:        schema.TypeString,
				Description: "Enable/disable web caching for HTTPS (Requires deep-inspection enabled in ssl-ssh-profile).",
				Computed:    true,
			},
			"webfilter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing Web filter profile.",
				Computed:    true,
			},
			"webproxy_forward_server": {
				Type:        schema.TypeString,
				Description: "Web proxy forward server name.",
				Computed:    true,
			},
			"webproxy_profile": {
				Type:        schema.TypeString,
				Description: "Name of web proxy profile.",
				Computed:    true,
			},
			"ztna_ems_tag": {
				Type:        schema.TypeList,
				Description: "ZTNA EMS Tag names.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "EMS Tag name.",
							Computed:    true,
						},
					},
				},
			},
			"ztna_tags_match_logic": {
				Type:        schema.TypeString,
				Description: "ZTNA tag matching logic.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallProxyPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("policyid")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadFirewallProxyPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallProxyPolicy dataSource: %v", err)
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

	diags := refreshObjectFirewallProxyPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
