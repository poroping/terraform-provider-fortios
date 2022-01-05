// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure proxy policies.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceFirewallProxyPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure proxy policies.",

		CreateContext: resourceFirewallProxyPolicyCreate,
		ReadContext:   resourceFirewallProxyPolicyRead,
		UpdateContext: resourceFirewallProxyPolicyUpdate,
		DeleteContext: resourceFirewallProxyPolicyDelete,

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
			"access_proxy": {
				Type:        schema.TypeList,
				Description: "IPv4 access proxy.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Access Proxy name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"access_proxy6": {
				Type:        schema.TypeList,
				Description: "IPv6 access proxy.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Access proxy name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"accept", "deny", "redirect"}, false),

				Description: "Accept or deny traffic matching the policy parameters.",
				Optional:    true,
				Computed:    true,
			},
			"application_list": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing Application list.",
				Optional:    true,
				Computed:    true,
			},
			"av_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing Antivirus profile.",
				Optional:    true,
				Computed:    true,
			},
			"cifs_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing CIFS profile.",
				Optional:    true,
				Computed:    true,
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),

				Description: "Optional comments.",
				Optional:    true,
				Computed:    true,
			},
			"decrypted_traffic_mirror": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Decrypted traffic mirror.",
				Optional:    true,
				Computed:    true,
			},
			"device_ownership": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "When enabled, the ownership enforcement will be done at policy level.",
				Optional:    true,
				Computed:    true,
			},
			"disclaimer": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "domain", "policy", "user"}, false),

				Description: "Web proxy disclaimer setting: by domain, policy, or user.",
				Optional:    true,
				Computed:    true,
			},
			"dlp_sensor": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing DLP sensor.",
				Optional:    true,
				Computed:    true,
			},
			"dstaddr": {
				Type:        schema.TypeList,
				Description: "Destination address objects.",
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
			"dstaddr_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "When enabled, destination addresses match against any address EXCEPT the specified destination addresses.",
				Optional:    true,
				Computed:    true,
			},
			"dstaddr6": {
				Type:        schema.TypeList,
				Description: "IPv6 destination address objects.",
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
			"dstintf": {
				Type:        schema.TypeList,
				Description: "Destination interface names.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"emailfilter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing email filter profile.",
				Optional:    true,
				Computed:    true,
			},
			"file_filter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing file-filter profile.",
				Optional:    true,
				Computed:    true,
			},
			"groups": {
				Type:        schema.TypeList,
				Description: "Names of group objects.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"http_tunnel_auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable HTTP tunnel authentication.",
				Optional:    true,
				Computed:    true,
			},
			"icap_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing ICAP profile.",
				Optional:    true,
				Computed:    true,
			},
			"internet_service": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.",
				Optional:    true,
				Computed:    true,
			},
			"internet_service_custom": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Custom Internet Service name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_custom_group": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service group name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Custom Internet Service group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_group": {
				Type:        schema.TypeList,
				Description: "Internet Service group name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Internet Service group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_id": {
				Type:        schema.TypeList,
				Description: "Internet Service ID.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Internet Service ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_name": {
				Type:        schema.TypeList,
				Description: "Internet Service name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Internet Service name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "When enabled, Internet Services match against any internet service EXCEPT the selected Internet Service.",
				Optional:    true,
				Computed:    true,
			},
			"ips_sensor": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing IPS sensor.",
				Optional:    true,
				Computed:    true,
			},
			"logtraffic": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"all", "utm", "disable"}, false),

				Description: "Enable/disable logging traffic through the policy.",
				Optional:    true,
				Computed:    true,
			},
			"logtraffic_start": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable policy log traffic start.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Policy name.",
				Optional:    true,
				Computed:    true,
			},
			"policyid": {
				Type: schema.TypeInt,

				Description: "Policy ID.",
				ForceNew:    true,
				Required:    true,
			},
			"poolname": {
				Type:        schema.TypeList,
				Description: "Name of IP pool object.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "IP pool name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"profile_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of profile group.",
				Optional:    true,
				Computed:    true,
			},
			"profile_protocol_options": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing Protocol options profile.",
				Optional:    true,
				Computed:    true,
			},
			"profile_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"single", "group"}, false),

				Description: "Determine whether the firewall policy allows security profile groups or single profiles only.",
				Optional:    true,
				Computed:    true,
			},
			"proxy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"explicit-web", "transparent-web", "ftp", "ssh", "ssh-tunnel", "access-proxy", "wanopt"}, false),

				Description: "Type of explicit proxy.",
				Optional:    true,
				Computed:    true,
			},
			"redirect_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),

				Description: "Redirect URL for further explicit web proxy processing.",
				Optional:    true,
				Computed:    true,
			},
			"replacemsg_override_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Authentication replacement message override group.",
				Optional:    true,
				Computed:    true,
			},
			"schedule": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of schedule object.",
				Optional:    true,
				Computed:    true,
			},
			"sctp_filter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing SCTP filter profile.",
				Optional:    true,
				Computed:    true,
			},
			"service": {
				Type:        schema.TypeList,
				Description: "Name of service objects.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Service name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"service_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "When enabled, services match against any service EXCEPT the specified destination services.",
				Optional:    true,
				Computed:    true,
			},
			"session_ttl": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 2764800),

				Description: "TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).",
				Optional:    true,
				Computed:    true,
			},
			"srcaddr": {
				Type:        schema.TypeList,
				Description: "Source address objects.",
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
			"srcaddr_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "When enabled, source addresses match against any address EXCEPT the specified source addresses.",
				Optional:    true,
				Computed:    true,
			},
			"srcaddr6": {
				Type:        schema.TypeList,
				Description: "IPv6 source address objects.",
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
			"srcintf": {
				Type:        schema.TypeList,
				Description: "Source interface names.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ssh_filter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing SSH filter profile.",
				Optional:    true,
				Computed:    true,
			},
			"ssh_policy_redirect": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Redirect SSH traffic to matching transparent proxy policy.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_ssh_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing SSL SSH profile.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the active status of the policy.",
				Optional:    true,
				Computed:    true,
			},
			"transparent": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to use the IP address of the client to connect to the server.",
				Optional:    true,
				Computed:    true,
			},
			"users": {
				Type:        schema.TypeList,
				Description: "Names of user objects.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"utm_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable the use of UTM profiles/sensors/lists.",
				Optional:    true,
				Computed:    true,
			},
			"uuid": {
				Type: schema.TypeString,

				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Optional:    true,
				Computed:    true,
			},
			"videofilter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing VideoFilter profile.",
				Optional:    true,
				Computed:    true,
			},
			"voip_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing VoIP profile.",
				Optional:    true,
				Computed:    true,
			},
			"waf_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing Web application firewall profile.",
				Optional:    true,
				Computed:    true,
			},
			"webcache": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable web caching.",
				Optional:    true,
				Computed:    true,
			},
			"webcache_https": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable web caching for HTTPS (Requires deep-inspection enabled in ssl-ssh-profile).",
				Optional:    true,
				Computed:    true,
			},
			"webfilter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing Web filter profile.",
				Optional:    true,
				Computed:    true,
			},
			"webproxy_forward_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Web proxy forward server name.",
				Optional:    true,
				Computed:    true,
			},
			"webproxy_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Name of web proxy profile.",
				Optional:    true,
				Computed:    true,
			},
			"ztna_ems_tag": {
				Type:        schema.TypeList,
				Description: "ZTNA EMS Tag names.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "EMS Tag name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ztna_tags_match_logic": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"or", "and"}, false),

				Description: "ZTNA tag matching logic.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallProxyPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "policyid"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating FirewallProxyPolicy resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallProxyPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallProxyPolicy(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallProxyPolicy")
	}

	return resourceFirewallProxyPolicyRead(ctx, d, meta)
}

func resourceFirewallProxyPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallProxyPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallProxyPolicy(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallProxyPolicy resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallProxyPolicy")
	}

	return resourceFirewallProxyPolicyRead(ctx, d, meta)
}

func resourceFirewallProxyPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallProxyPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallProxyPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallProxyPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallProxyPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallProxyPolicy resource: %v", err)
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

	diags := refreshObjectFirewallProxyPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallProxyPolicyAccessProxy(v *[]models.FirewallProxyPolicyAccessProxy, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallProxyPolicyAccessProxy6(v *[]models.FirewallProxyPolicyAccessProxy6, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallProxyPolicyDstaddr(v *[]models.FirewallProxyPolicyDstaddr, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallProxyPolicyDstaddr6(v *[]models.FirewallProxyPolicyDstaddr6, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallProxyPolicyDstintf(v *[]models.FirewallProxyPolicyDstintf, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallProxyPolicyGroups(v *[]models.FirewallProxyPolicyGroups, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallProxyPolicyInternetServiceCustom(v *[]models.FirewallProxyPolicyInternetServiceCustom, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallProxyPolicyInternetServiceCustomGroup(v *[]models.FirewallProxyPolicyInternetServiceCustomGroup, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallProxyPolicyInternetServiceGroup(v *[]models.FirewallProxyPolicyInternetServiceGroup, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallProxyPolicyInternetServiceId(v *[]models.FirewallProxyPolicyInternetServiceId, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
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

func flattenFirewallProxyPolicyInternetServiceName(v *[]models.FirewallProxyPolicyInternetServiceName, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallProxyPolicyPoolname(v *[]models.FirewallProxyPolicyPoolname, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallProxyPolicyService(v *[]models.FirewallProxyPolicyService, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallProxyPolicySrcaddr(v *[]models.FirewallProxyPolicySrcaddr, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallProxyPolicySrcaddr6(v *[]models.FirewallProxyPolicySrcaddr6, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallProxyPolicySrcintf(v *[]models.FirewallProxyPolicySrcintf, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallProxyPolicyUsers(v *[]models.FirewallProxyPolicyUsers, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallProxyPolicyZtnaEmsTag(v *[]models.FirewallProxyPolicyZtnaEmsTag, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func refreshObjectFirewallProxyPolicy(d *schema.ResourceData, o *models.FirewallProxyPolicy, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AccessProxy != nil {
		if err = d.Set("access_proxy", flattenFirewallProxyPolicyAccessProxy(o.AccessProxy, sort)); err != nil {
			return diag.Errorf("error reading access_proxy: %v", err)
		}
	}

	if o.AccessProxy6 != nil {
		if err = d.Set("access_proxy6", flattenFirewallProxyPolicyAccessProxy6(o.AccessProxy6, sort)); err != nil {
			return diag.Errorf("error reading access_proxy6: %v", err)
		}
	}

	if o.Action != nil {
		v := *o.Action

		if err = d.Set("action", v); err != nil {
			return diag.Errorf("error reading action: %v", err)
		}
	}

	if o.ApplicationList != nil {
		v := *o.ApplicationList

		if err = d.Set("application_list", v); err != nil {
			return diag.Errorf("error reading application_list: %v", err)
		}
	}

	if o.AvProfile != nil {
		v := *o.AvProfile

		if err = d.Set("av_profile", v); err != nil {
			return diag.Errorf("error reading av_profile: %v", err)
		}
	}

	if o.CifsProfile != nil {
		v := *o.CifsProfile

		if err = d.Set("cifs_profile", v); err != nil {
			return diag.Errorf("error reading cifs_profile: %v", err)
		}
	}

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.DecryptedTrafficMirror != nil {
		v := *o.DecryptedTrafficMirror

		if err = d.Set("decrypted_traffic_mirror", v); err != nil {
			return diag.Errorf("error reading decrypted_traffic_mirror: %v", err)
		}
	}

	if o.DeviceOwnership != nil {
		v := *o.DeviceOwnership

		if err = d.Set("device_ownership", v); err != nil {
			return diag.Errorf("error reading device_ownership: %v", err)
		}
	}

	if o.Disclaimer != nil {
		v := *o.Disclaimer

		if err = d.Set("disclaimer", v); err != nil {
			return diag.Errorf("error reading disclaimer: %v", err)
		}
	}

	if o.DlpSensor != nil {
		v := *o.DlpSensor

		if err = d.Set("dlp_sensor", v); err != nil {
			return diag.Errorf("error reading dlp_sensor: %v", err)
		}
	}

	if o.Dstaddr != nil {
		if err = d.Set("dstaddr", flattenFirewallProxyPolicyDstaddr(o.Dstaddr, sort)); err != nil {
			return diag.Errorf("error reading dstaddr: %v", err)
		}
	}

	if o.DstaddrNegate != nil {
		v := *o.DstaddrNegate

		if err = d.Set("dstaddr_negate", v); err != nil {
			return diag.Errorf("error reading dstaddr_negate: %v", err)
		}
	}

	if o.Dstaddr6 != nil {
		if err = d.Set("dstaddr6", flattenFirewallProxyPolicyDstaddr6(o.Dstaddr6, sort)); err != nil {
			return diag.Errorf("error reading dstaddr6: %v", err)
		}
	}

	if o.Dstintf != nil {
		if err = d.Set("dstintf", flattenFirewallProxyPolicyDstintf(o.Dstintf, sort)); err != nil {
			return diag.Errorf("error reading dstintf: %v", err)
		}
	}

	if o.EmailfilterProfile != nil {
		v := *o.EmailfilterProfile

		if err = d.Set("emailfilter_profile", v); err != nil {
			return diag.Errorf("error reading emailfilter_profile: %v", err)
		}
	}

	if o.FileFilterProfile != nil {
		v := *o.FileFilterProfile

		if err = d.Set("file_filter_profile", v); err != nil {
			return diag.Errorf("error reading file_filter_profile: %v", err)
		}
	}

	if o.Groups != nil {
		if err = d.Set("groups", flattenFirewallProxyPolicyGroups(o.Groups, sort)); err != nil {
			return diag.Errorf("error reading groups: %v", err)
		}
	}

	if o.HttpTunnelAuth != nil {
		v := *o.HttpTunnelAuth

		if err = d.Set("http_tunnel_auth", v); err != nil {
			return diag.Errorf("error reading http_tunnel_auth: %v", err)
		}
	}

	if o.IcapProfile != nil {
		v := *o.IcapProfile

		if err = d.Set("icap_profile", v); err != nil {
			return diag.Errorf("error reading icap_profile: %v", err)
		}
	}

	if o.InternetService != nil {
		v := *o.InternetService

		if err = d.Set("internet_service", v); err != nil {
			return diag.Errorf("error reading internet_service: %v", err)
		}
	}

	if o.InternetServiceCustom != nil {
		if err = d.Set("internet_service_custom", flattenFirewallProxyPolicyInternetServiceCustom(o.InternetServiceCustom, sort)); err != nil {
			return diag.Errorf("error reading internet_service_custom: %v", err)
		}
	}

	if o.InternetServiceCustomGroup != nil {
		if err = d.Set("internet_service_custom_group", flattenFirewallProxyPolicyInternetServiceCustomGroup(o.InternetServiceCustomGroup, sort)); err != nil {
			return diag.Errorf("error reading internet_service_custom_group: %v", err)
		}
	}

	if o.InternetServiceGroup != nil {
		if err = d.Set("internet_service_group", flattenFirewallProxyPolicyInternetServiceGroup(o.InternetServiceGroup, sort)); err != nil {
			return diag.Errorf("error reading internet_service_group: %v", err)
		}
	}

	if o.InternetServiceId != nil {
		if err = d.Set("internet_service_id", flattenFirewallProxyPolicyInternetServiceId(o.InternetServiceId, sort)); err != nil {
			return diag.Errorf("error reading internet_service_id: %v", err)
		}
	}

	if o.InternetServiceName != nil {
		if err = d.Set("internet_service_name", flattenFirewallProxyPolicyInternetServiceName(o.InternetServiceName, sort)); err != nil {
			return diag.Errorf("error reading internet_service_name: %v", err)
		}
	}

	if o.InternetServiceNegate != nil {
		v := *o.InternetServiceNegate

		if err = d.Set("internet_service_negate", v); err != nil {
			return diag.Errorf("error reading internet_service_negate: %v", err)
		}
	}

	if o.IpsSensor != nil {
		v := *o.IpsSensor

		if err = d.Set("ips_sensor", v); err != nil {
			return diag.Errorf("error reading ips_sensor: %v", err)
		}
	}

	if o.Logtraffic != nil {
		v := *o.Logtraffic

		if err = d.Set("logtraffic", v); err != nil {
			return diag.Errorf("error reading logtraffic: %v", err)
		}
	}

	if o.LogtrafficStart != nil {
		v := *o.LogtrafficStart

		if err = d.Set("logtraffic_start", v); err != nil {
			return diag.Errorf("error reading logtraffic_start: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Policyid != nil {
		v := *o.Policyid

		if err = d.Set("policyid", v); err != nil {
			return diag.Errorf("error reading policyid: %v", err)
		}
	}

	if o.Poolname != nil {
		if err = d.Set("poolname", flattenFirewallProxyPolicyPoolname(o.Poolname, sort)); err != nil {
			return diag.Errorf("error reading poolname: %v", err)
		}
	}

	if o.ProfileGroup != nil {
		v := *o.ProfileGroup

		if err = d.Set("profile_group", v); err != nil {
			return diag.Errorf("error reading profile_group: %v", err)
		}
	}

	if o.ProfileProtocolOptions != nil {
		v := *o.ProfileProtocolOptions

		if err = d.Set("profile_protocol_options", v); err != nil {
			return diag.Errorf("error reading profile_protocol_options: %v", err)
		}
	}

	if o.ProfileType != nil {
		v := *o.ProfileType

		if err = d.Set("profile_type", v); err != nil {
			return diag.Errorf("error reading profile_type: %v", err)
		}
	}

	if o.Proxy != nil {
		v := *o.Proxy

		if err = d.Set("proxy", v); err != nil {
			return diag.Errorf("error reading proxy: %v", err)
		}
	}

	if o.RedirectUrl != nil {
		v := *o.RedirectUrl

		if err = d.Set("redirect_url", v); err != nil {
			return diag.Errorf("error reading redirect_url: %v", err)
		}
	}

	if o.ReplacemsgOverrideGroup != nil {
		v := *o.ReplacemsgOverrideGroup

		if err = d.Set("replacemsg_override_group", v); err != nil {
			return diag.Errorf("error reading replacemsg_override_group: %v", err)
		}
	}

	if o.Schedule != nil {
		v := *o.Schedule

		if err = d.Set("schedule", v); err != nil {
			return diag.Errorf("error reading schedule: %v", err)
		}
	}

	if o.SctpFilterProfile != nil {
		v := *o.SctpFilterProfile

		if err = d.Set("sctp_filter_profile", v); err != nil {
			return diag.Errorf("error reading sctp_filter_profile: %v", err)
		}
	}

	if o.Service != nil {
		if err = d.Set("service", flattenFirewallProxyPolicyService(o.Service, sort)); err != nil {
			return diag.Errorf("error reading service: %v", err)
		}
	}

	if o.ServiceNegate != nil {
		v := *o.ServiceNegate

		if err = d.Set("service_negate", v); err != nil {
			return diag.Errorf("error reading service_negate: %v", err)
		}
	}

	if o.SessionTtl != nil {
		v := *o.SessionTtl

		if err = d.Set("session_ttl", v); err != nil {
			return diag.Errorf("error reading session_ttl: %v", err)
		}
	}

	if o.Srcaddr != nil {
		if err = d.Set("srcaddr", flattenFirewallProxyPolicySrcaddr(o.Srcaddr, sort)); err != nil {
			return diag.Errorf("error reading srcaddr: %v", err)
		}
	}

	if o.SrcaddrNegate != nil {
		v := *o.SrcaddrNegate

		if err = d.Set("srcaddr_negate", v); err != nil {
			return diag.Errorf("error reading srcaddr_negate: %v", err)
		}
	}

	if o.Srcaddr6 != nil {
		if err = d.Set("srcaddr6", flattenFirewallProxyPolicySrcaddr6(o.Srcaddr6, sort)); err != nil {
			return diag.Errorf("error reading srcaddr6: %v", err)
		}
	}

	if o.Srcintf != nil {
		if err = d.Set("srcintf", flattenFirewallProxyPolicySrcintf(o.Srcintf, sort)); err != nil {
			return diag.Errorf("error reading srcintf: %v", err)
		}
	}

	if o.SshFilterProfile != nil {
		v := *o.SshFilterProfile

		if err = d.Set("ssh_filter_profile", v); err != nil {
			return diag.Errorf("error reading ssh_filter_profile: %v", err)
		}
	}

	if o.SshPolicyRedirect != nil {
		v := *o.SshPolicyRedirect

		if err = d.Set("ssh_policy_redirect", v); err != nil {
			return diag.Errorf("error reading ssh_policy_redirect: %v", err)
		}
	}

	if o.SslSshProfile != nil {
		v := *o.SslSshProfile

		if err = d.Set("ssl_ssh_profile", v); err != nil {
			return diag.Errorf("error reading ssl_ssh_profile: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Transparent != nil {
		v := *o.Transparent

		if err = d.Set("transparent", v); err != nil {
			return diag.Errorf("error reading transparent: %v", err)
		}
	}

	if o.Users != nil {
		if err = d.Set("users", flattenFirewallProxyPolicyUsers(o.Users, sort)); err != nil {
			return diag.Errorf("error reading users: %v", err)
		}
	}

	if o.UtmStatus != nil {
		v := *o.UtmStatus

		if err = d.Set("utm_status", v); err != nil {
			return diag.Errorf("error reading utm_status: %v", err)
		}
	}

	if o.Uuid != nil {
		v := *o.Uuid

		if err = d.Set("uuid", v); err != nil {
			return diag.Errorf("error reading uuid: %v", err)
		}
	}

	if o.VideofilterProfile != nil {
		v := *o.VideofilterProfile

		if err = d.Set("videofilter_profile", v); err != nil {
			return diag.Errorf("error reading videofilter_profile: %v", err)
		}
	}

	if o.VoipProfile != nil {
		v := *o.VoipProfile

		if err = d.Set("voip_profile", v); err != nil {
			return diag.Errorf("error reading voip_profile: %v", err)
		}
	}

	if o.WafProfile != nil {
		v := *o.WafProfile

		if err = d.Set("waf_profile", v); err != nil {
			return diag.Errorf("error reading waf_profile: %v", err)
		}
	}

	if o.Webcache != nil {
		v := *o.Webcache

		if err = d.Set("webcache", v); err != nil {
			return diag.Errorf("error reading webcache: %v", err)
		}
	}

	if o.WebcacheHttps != nil {
		v := *o.WebcacheHttps

		if err = d.Set("webcache_https", v); err != nil {
			return diag.Errorf("error reading webcache_https: %v", err)
		}
	}

	if o.WebfilterProfile != nil {
		v := *o.WebfilterProfile

		if err = d.Set("webfilter_profile", v); err != nil {
			return diag.Errorf("error reading webfilter_profile: %v", err)
		}
	}

	if o.WebproxyForwardServer != nil {
		v := *o.WebproxyForwardServer

		if err = d.Set("webproxy_forward_server", v); err != nil {
			return diag.Errorf("error reading webproxy_forward_server: %v", err)
		}
	}

	if o.WebproxyProfile != nil {
		v := *o.WebproxyProfile

		if err = d.Set("webproxy_profile", v); err != nil {
			return diag.Errorf("error reading webproxy_profile: %v", err)
		}
	}

	if o.ZtnaEmsTag != nil {
		if err = d.Set("ztna_ems_tag", flattenFirewallProxyPolicyZtnaEmsTag(o.ZtnaEmsTag, sort)); err != nil {
			return diag.Errorf("error reading ztna_ems_tag: %v", err)
		}
	}

	if o.ZtnaTagsMatchLogic != nil {
		v := *o.ZtnaTagsMatchLogic

		if err = d.Set("ztna_tags_match_logic", v); err != nil {
			return diag.Errorf("error reading ztna_tags_match_logic: %v", err)
		}
	}

	return nil
}

func expandFirewallProxyPolicyAccessProxy(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProxyPolicyAccessProxy, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProxyPolicyAccessProxy

	for i := range l {
		tmp := models.FirewallProxyPolicyAccessProxy{}
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

func expandFirewallProxyPolicyAccessProxy6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProxyPolicyAccessProxy6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProxyPolicyAccessProxy6

	for i := range l {
		tmp := models.FirewallProxyPolicyAccessProxy6{}
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

func expandFirewallProxyPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProxyPolicyDstaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProxyPolicyDstaddr

	for i := range l {
		tmp := models.FirewallProxyPolicyDstaddr{}
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

func expandFirewallProxyPolicyDstaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProxyPolicyDstaddr6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProxyPolicyDstaddr6

	for i := range l {
		tmp := models.FirewallProxyPolicyDstaddr6{}
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

func expandFirewallProxyPolicyDstintf(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProxyPolicyDstintf, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProxyPolicyDstintf

	for i := range l {
		tmp := models.FirewallProxyPolicyDstintf{}
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

func expandFirewallProxyPolicyGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProxyPolicyGroups, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProxyPolicyGroups

	for i := range l {
		tmp := models.FirewallProxyPolicyGroups{}
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

func expandFirewallProxyPolicyInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProxyPolicyInternetServiceCustom, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProxyPolicyInternetServiceCustom

	for i := range l {
		tmp := models.FirewallProxyPolicyInternetServiceCustom{}
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

func expandFirewallProxyPolicyInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProxyPolicyInternetServiceCustomGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProxyPolicyInternetServiceCustomGroup

	for i := range l {
		tmp := models.FirewallProxyPolicyInternetServiceCustomGroup{}
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

func expandFirewallProxyPolicyInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProxyPolicyInternetServiceGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProxyPolicyInternetServiceGroup

	for i := range l {
		tmp := models.FirewallProxyPolicyInternetServiceGroup{}
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

func expandFirewallProxyPolicyInternetServiceId(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProxyPolicyInternetServiceId, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProxyPolicyInternetServiceId

	for i := range l {
		tmp := models.FirewallProxyPolicyInternetServiceId{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallProxyPolicyInternetServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProxyPolicyInternetServiceName, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProxyPolicyInternetServiceName

	for i := range l {
		tmp := models.FirewallProxyPolicyInternetServiceName{}
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

func expandFirewallProxyPolicyPoolname(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProxyPolicyPoolname, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProxyPolicyPoolname

	for i := range l {
		tmp := models.FirewallProxyPolicyPoolname{}
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

func expandFirewallProxyPolicyService(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProxyPolicyService, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProxyPolicyService

	for i := range l {
		tmp := models.FirewallProxyPolicyService{}
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

func expandFirewallProxyPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProxyPolicySrcaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProxyPolicySrcaddr

	for i := range l {
		tmp := models.FirewallProxyPolicySrcaddr{}
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

func expandFirewallProxyPolicySrcaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProxyPolicySrcaddr6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProxyPolicySrcaddr6

	for i := range l {
		tmp := models.FirewallProxyPolicySrcaddr6{}
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

func expandFirewallProxyPolicySrcintf(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProxyPolicySrcintf, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProxyPolicySrcintf

	for i := range l {
		tmp := models.FirewallProxyPolicySrcintf{}
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

func expandFirewallProxyPolicyUsers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProxyPolicyUsers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProxyPolicyUsers

	for i := range l {
		tmp := models.FirewallProxyPolicyUsers{}
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

func expandFirewallProxyPolicyZtnaEmsTag(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProxyPolicyZtnaEmsTag, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProxyPolicyZtnaEmsTag

	for i := range l {
		tmp := models.FirewallProxyPolicyZtnaEmsTag{}
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

func getObjectFirewallProxyPolicy(d *schema.ResourceData, sv string) (*models.FirewallProxyPolicy, diag.Diagnostics) {
	obj := models.FirewallProxyPolicy{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("access_proxy"); ok {
		if !utils.CheckVer(sv, "v7.0.0", "") {
			e := utils.AttributeVersionWarning("access_proxy", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProxyPolicyAccessProxy(d, v, "access_proxy", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.AccessProxy = t
		}
	} else if d.HasChange("access_proxy") {
		old, new := d.GetChange("access_proxy")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.AccessProxy = &[]models.FirewallProxyPolicyAccessProxy{}
		}
	}
	if v, ok := d.GetOk("access_proxy6"); ok {
		if !utils.CheckVer(sv, "v7.0.1", "") {
			e := utils.AttributeVersionWarning("access_proxy6", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProxyPolicyAccessProxy6(d, v, "access_proxy6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.AccessProxy6 = t
		}
	} else if d.HasChange("access_proxy6") {
		old, new := d.GetChange("access_proxy6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.AccessProxy6 = &[]models.FirewallProxyPolicyAccessProxy6{}
		}
	}
	if v1, ok := d.GetOk("action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("action", sv)
				diags = append(diags, e)
			}
			obj.Action = &v2
		}
	}
	if v1, ok := d.GetOk("application_list"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("application_list", sv)
				diags = append(diags, e)
			}
			obj.ApplicationList = &v2
		}
	}
	if v1, ok := d.GetOk("av_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("av_profile", sv)
				diags = append(diags, e)
			}
			obj.AvProfile = &v2
		}
	}
	if v1, ok := d.GetOk("cifs_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cifs_profile", sv)
				diags = append(diags, e)
			}
			obj.CifsProfile = &v2
		}
	}
	if v1, ok := d.GetOk("comments"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comments", sv)
				diags = append(diags, e)
			}
			obj.Comments = &v2
		}
	}
	if v1, ok := d.GetOk("decrypted_traffic_mirror"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("decrypted_traffic_mirror", sv)
				diags = append(diags, e)
			}
			obj.DecryptedTrafficMirror = &v2
		}
	}
	if v1, ok := d.GetOk("device_ownership"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("device_ownership", sv)
				diags = append(diags, e)
			}
			obj.DeviceOwnership = &v2
		}
	}
	if v1, ok := d.GetOk("disclaimer"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("disclaimer", sv)
				diags = append(diags, e)
			}
			obj.Disclaimer = &v2
		}
	}
	if v1, ok := d.GetOk("dlp_sensor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dlp_sensor", sv)
				diags = append(diags, e)
			}
			obj.DlpSensor = &v2
		}
	}
	if v, ok := d.GetOk("dstaddr"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dstaddr", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProxyPolicyDstaddr(d, v, "dstaddr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstaddr = t
		}
	} else if d.HasChange("dstaddr") {
		old, new := d.GetChange("dstaddr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstaddr = &[]models.FirewallProxyPolicyDstaddr{}
		}
	}
	if v1, ok := d.GetOk("dstaddr_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dstaddr_negate", sv)
				diags = append(diags, e)
			}
			obj.DstaddrNegate = &v2
		}
	}
	if v, ok := d.GetOk("dstaddr6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dstaddr6", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProxyPolicyDstaddr6(d, v, "dstaddr6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstaddr6 = t
		}
	} else if d.HasChange("dstaddr6") {
		old, new := d.GetChange("dstaddr6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstaddr6 = &[]models.FirewallProxyPolicyDstaddr6{}
		}
	}
	if v, ok := d.GetOk("dstintf"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dstintf", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProxyPolicyDstintf(d, v, "dstintf", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstintf = t
		}
	} else if d.HasChange("dstintf") {
		old, new := d.GetChange("dstintf")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstintf = &[]models.FirewallProxyPolicyDstintf{}
		}
	}
	if v1, ok := d.GetOk("emailfilter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("emailfilter_profile", sv)
				diags = append(diags, e)
			}
			obj.EmailfilterProfile = &v2
		}
	}
	if v1, ok := d.GetOk("file_filter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("file_filter_profile", sv)
				diags = append(diags, e)
			}
			obj.FileFilterProfile = &v2
		}
	}
	if v, ok := d.GetOk("groups"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("groups", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProxyPolicyGroups(d, v, "groups", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Groups = t
		}
	} else if d.HasChange("groups") {
		old, new := d.GetChange("groups")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Groups = &[]models.FirewallProxyPolicyGroups{}
		}
	}
	if v1, ok := d.GetOk("http_tunnel_auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_tunnel_auth", sv)
				diags = append(diags, e)
			}
			obj.HttpTunnelAuth = &v2
		}
	}
	if v1, ok := d.GetOk("icap_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("icap_profile", sv)
				diags = append(diags, e)
			}
			obj.IcapProfile = &v2
		}
	}
	if v1, ok := d.GetOk("internet_service"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("internet_service", sv)
				diags = append(diags, e)
			}
			obj.InternetService = &v2
		}
	}
	if v, ok := d.GetOk("internet_service_custom"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_custom", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProxyPolicyInternetServiceCustom(d, v, "internet_service_custom", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceCustom = t
		}
	} else if d.HasChange("internet_service_custom") {
		old, new := d.GetChange("internet_service_custom")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceCustom = &[]models.FirewallProxyPolicyInternetServiceCustom{}
		}
	}
	if v, ok := d.GetOk("internet_service_custom_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_custom_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProxyPolicyInternetServiceCustomGroup(d, v, "internet_service_custom_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceCustomGroup = t
		}
	} else if d.HasChange("internet_service_custom_group") {
		old, new := d.GetChange("internet_service_custom_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceCustomGroup = &[]models.FirewallProxyPolicyInternetServiceCustomGroup{}
		}
	}
	if v, ok := d.GetOk("internet_service_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProxyPolicyInternetServiceGroup(d, v, "internet_service_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceGroup = t
		}
	} else if d.HasChange("internet_service_group") {
		old, new := d.GetChange("internet_service_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceGroup = &[]models.FirewallProxyPolicyInternetServiceGroup{}
		}
	}
	if v, ok := d.GetOk("internet_service_id"); ok {
		if !utils.CheckVer(sv, "", "v6.4.0") {
			e := utils.AttributeVersionWarning("internet_service_id", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProxyPolicyInternetServiceId(d, v, "internet_service_id", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceId = t
		}
	} else if d.HasChange("internet_service_id") {
		old, new := d.GetChange("internet_service_id")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceId = &[]models.FirewallProxyPolicyInternetServiceId{}
		}
	}
	if v, ok := d.GetOk("internet_service_name"); ok {
		if !utils.CheckVer(sv, "v6.4.0", "") {
			e := utils.AttributeVersionWarning("internet_service_name", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProxyPolicyInternetServiceName(d, v, "internet_service_name", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceName = t
		}
	} else if d.HasChange("internet_service_name") {
		old, new := d.GetChange("internet_service_name")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceName = &[]models.FirewallProxyPolicyInternetServiceName{}
		}
	}
	if v1, ok := d.GetOk("internet_service_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("internet_service_negate", sv)
				diags = append(diags, e)
			}
			obj.InternetServiceNegate = &v2
		}
	}
	if v1, ok := d.GetOk("ips_sensor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ips_sensor", sv)
				diags = append(diags, e)
			}
			obj.IpsSensor = &v2
		}
	}
	if v1, ok := d.GetOk("logtraffic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("logtraffic", sv)
				diags = append(diags, e)
			}
			obj.Logtraffic = &v2
		}
	}
	if v1, ok := d.GetOk("logtraffic_start"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("logtraffic_start", sv)
				diags = append(diags, e)
			}
			obj.LogtrafficStart = &v2
		}
	}
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v1, ok := d.GetOk("policyid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("policyid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Policyid = &tmp
		}
	}
	if v, ok := d.GetOk("poolname"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("poolname", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProxyPolicyPoolname(d, v, "poolname", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Poolname = t
		}
	} else if d.HasChange("poolname") {
		old, new := d.GetChange("poolname")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Poolname = &[]models.FirewallProxyPolicyPoolname{}
		}
	}
	if v1, ok := d.GetOk("profile_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("profile_group", sv)
				diags = append(diags, e)
			}
			obj.ProfileGroup = &v2
		}
	}
	if v1, ok := d.GetOk("profile_protocol_options"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("profile_protocol_options", sv)
				diags = append(diags, e)
			}
			obj.ProfileProtocolOptions = &v2
		}
	}
	if v1, ok := d.GetOk("profile_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("profile_type", sv)
				diags = append(diags, e)
			}
			obj.ProfileType = &v2
		}
	}
	if v1, ok := d.GetOk("proxy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("proxy", sv)
				diags = append(diags, e)
			}
			obj.Proxy = &v2
		}
	}
	if v1, ok := d.GetOk("redirect_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("redirect_url", sv)
				diags = append(diags, e)
			}
			obj.RedirectUrl = &v2
		}
	}
	if v1, ok := d.GetOk("replacemsg_override_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("replacemsg_override_group", sv)
				diags = append(diags, e)
			}
			obj.ReplacemsgOverrideGroup = &v2
		}
	}
	if v1, ok := d.GetOk("schedule"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("schedule", sv)
				diags = append(diags, e)
			}
			obj.Schedule = &v2
		}
	}
	if v1, ok := d.GetOk("sctp_filter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("sctp_filter_profile", sv)
				diags = append(diags, e)
			}
			obj.SctpFilterProfile = &v2
		}
	}
	if v, ok := d.GetOk("service"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("service", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProxyPolicyService(d, v, "service", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Service = t
		}
	} else if d.HasChange("service") {
		old, new := d.GetChange("service")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Service = &[]models.FirewallProxyPolicyService{}
		}
	}
	if v1, ok := d.GetOk("service_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("service_negate", sv)
				diags = append(diags, e)
			}
			obj.ServiceNegate = &v2
		}
	}
	if v1, ok := d.GetOk("session_ttl"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("session_ttl", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SessionTtl = &tmp
		}
	}
	if v, ok := d.GetOk("srcaddr"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("srcaddr", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProxyPolicySrcaddr(d, v, "srcaddr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcaddr = t
		}
	} else if d.HasChange("srcaddr") {
		old, new := d.GetChange("srcaddr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcaddr = &[]models.FirewallProxyPolicySrcaddr{}
		}
	}
	if v1, ok := d.GetOk("srcaddr_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("srcaddr_negate", sv)
				diags = append(diags, e)
			}
			obj.SrcaddrNegate = &v2
		}
	}
	if v, ok := d.GetOk("srcaddr6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("srcaddr6", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProxyPolicySrcaddr6(d, v, "srcaddr6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcaddr6 = t
		}
	} else if d.HasChange("srcaddr6") {
		old, new := d.GetChange("srcaddr6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcaddr6 = &[]models.FirewallProxyPolicySrcaddr6{}
		}
	}
	if v, ok := d.GetOk("srcintf"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("srcintf", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProxyPolicySrcintf(d, v, "srcintf", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcintf = t
		}
	} else if d.HasChange("srcintf") {
		old, new := d.GetChange("srcintf")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcintf = &[]models.FirewallProxyPolicySrcintf{}
		}
	}
	if v1, ok := d.GetOk("ssh_filter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssh_filter_profile", sv)
				diags = append(diags, e)
			}
			obj.SshFilterProfile = &v2
		}
	}
	if v1, ok := d.GetOk("ssh_policy_redirect"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssh_policy_redirect", sv)
				diags = append(diags, e)
			}
			obj.SshPolicyRedirect = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_ssh_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_ssh_profile", sv)
				diags = append(diags, e)
			}
			obj.SslSshProfile = &v2
		}
	}
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	if v1, ok := d.GetOk("transparent"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("transparent", sv)
				diags = append(diags, e)
			}
			obj.Transparent = &v2
		}
	}
	if v, ok := d.GetOk("users"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("users", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProxyPolicyUsers(d, v, "users", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Users = t
		}
	} else if d.HasChange("users") {
		old, new := d.GetChange("users")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Users = &[]models.FirewallProxyPolicyUsers{}
		}
	}
	if v1, ok := d.GetOk("utm_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("utm_status", sv)
				diags = append(diags, e)
			}
			obj.UtmStatus = &v2
		}
	}
	if v1, ok := d.GetOk("uuid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("uuid", sv)
				diags = append(diags, e)
			}
			obj.Uuid = &v2
		}
	}
	if v1, ok := d.GetOk("videofilter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("videofilter_profile", sv)
				diags = append(diags, e)
			}
			obj.VideofilterProfile = &v2
		}
	}
	if v1, ok := d.GetOk("voip_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("voip_profile", sv)
				diags = append(diags, e)
			}
			obj.VoipProfile = &v2
		}
	}
	if v1, ok := d.GetOk("waf_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("waf_profile", sv)
				diags = append(diags, e)
			}
			obj.WafProfile = &v2
		}
	}
	if v1, ok := d.GetOk("webcache"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("webcache", sv)
				diags = append(diags, e)
			}
			obj.Webcache = &v2
		}
	}
	if v1, ok := d.GetOk("webcache_https"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("webcache_https", sv)
				diags = append(diags, e)
			}
			obj.WebcacheHttps = &v2
		}
	}
	if v1, ok := d.GetOk("webfilter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("webfilter_profile", sv)
				diags = append(diags, e)
			}
			obj.WebfilterProfile = &v2
		}
	}
	if v1, ok := d.GetOk("webproxy_forward_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("webproxy_forward_server", sv)
				diags = append(diags, e)
			}
			obj.WebproxyForwardServer = &v2
		}
	}
	if v1, ok := d.GetOk("webproxy_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("webproxy_profile", sv)
				diags = append(diags, e)
			}
			obj.WebproxyProfile = &v2
		}
	}
	if v, ok := d.GetOk("ztna_ems_tag"); ok {
		if !utils.CheckVer(sv, "v7.0.0", "") {
			e := utils.AttributeVersionWarning("ztna_ems_tag", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProxyPolicyZtnaEmsTag(d, v, "ztna_ems_tag", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ZtnaEmsTag = t
		}
	} else if d.HasChange("ztna_ems_tag") {
		old, new := d.GetChange("ztna_ems_tag")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ZtnaEmsTag = &[]models.FirewallProxyPolicyZtnaEmsTag{}
		}
	}
	if v1, ok := d.GetOk("ztna_tags_match_logic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("ztna_tags_match_logic", sv)
				diags = append(diags, e)
			}
			obj.ZtnaTagsMatchLogic = &v2
		}
	}
	return &obj, diags
}
