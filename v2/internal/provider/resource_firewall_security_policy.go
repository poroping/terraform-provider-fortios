// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure NGFW IPv4/IPv6 application policies.

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

func resourceFirewallSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure NGFW IPv4/IPv6 application policies.",

		CreateContext: resourceFirewallSecurityPolicyCreate,
		ReadContext:   resourceFirewallSecurityPolicyRead,
		UpdateContext: resourceFirewallSecurityPolicyUpdate,
		DeleteContext: resourceFirewallSecurityPolicyDelete,

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
			"action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"accept", "deny"}, false),

				Description: "Policy action (accept/deny).",
				Optional:    true,
				Computed:    true,
			},
			"app_category": {
				Type:        schema.TypeList,
				Description: "Application category ID list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Category IDs.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"app_group": {
				Type:        schema.TypeList,
				Description: "Application group names.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Application group names.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"application": {
				Type:        schema.TypeList,
				Description: "Application ID list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Application IDs.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
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

				Description: "Comment.",
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
			"dnsfilter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing DNS filter profile.",
				Optional:    true,
				Computed:    true,
			},
			"dstaddr": {
				Type:        schema.TypeList,
				Description: "Destination IPv4 address name and address group names.",
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

				Description: "When enabled dstaddr/dstaddr6 specifies what the destination address must NOT be.",
				Optional:    true,
				Computed:    true,
			},
			"dstaddr4": {
				Type:        schema.TypeList,
				Description: "Destination IPv4 address name and address group names.",
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
			"dstaddr6": {
				Type:        schema.TypeList,
				Description: "Destination IPv6 address name and address group names.",
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
				Description: "Outgoing (egress) interface.",
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
			"enforce_default_app_port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable default application port enforcement for allowed applications.",
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
			"fsso_groups": {
				Type:        schema.TypeList,
				Description: "Names of FSSO groups.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),

							Description: "Names of FSSO groups.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"groups": {
				Type:        schema.TypeList,
				Description: "Names of user groups that can authenticate with this policy.",
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

				Description: "When enabled internet-service specifies what the service must NOT be.",
				Optional:    true,
				Computed:    true,
			},
			"internet_service_src": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used.",
				Optional:    true,
				Computed:    true,
			},
			"internet_service_src_custom": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service source name.",
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
			"internet_service_src_custom_group": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service source group name.",
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
			"internet_service_src_group": {
				Type:        schema.TypeList,
				Description: "Internet Service source group name.",
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
			"internet_service_src_id": {
				Type:        schema.TypeList,
				Description: "Internet Service source ID.",
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
			"internet_service_src_name": {
				Type:        schema.TypeList,
				Description: "Internet Service source name.",
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
			"internet_service_src_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "When enabled internet-service-src specifies what the service must NOT be.",
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
			"learning_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to allow everything, but log all of the meaningful data for security information gathering. A learning report will be generated.",
				Optional:    true,
				Computed:    true,
			},
			"logtraffic": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"all", "utm", "disable"}, false),

				Description: "Enable or disable logging. Log all sessions or security profile sessions.",
				Optional:    true,
				Computed:    true,
			},
			"logtraffic_start": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Record logs when a session starts.",
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
			"nat46": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable NAT46.",
				Optional:    true,
				Computed:    true,
			},
			"nat64": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable NAT64.",
				Optional:    true,
				Computed:    true,
			},
			"policyid": {
				Type: schema.TypeInt,

				Description: "Policy ID.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
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
			"schedule": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Schedule name.",
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
			"send_deny_packet": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable to send a reply when a session is denied or blocked by a firewall policy.",
				Optional:    true,
				Computed:    true,
			},
			"service": {
				Type:        schema.TypeList,
				Description: "Service and service group names.",
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

				Description: "When enabled service specifies what the service must NOT be.",
				Optional:    true,
				Computed:    true,
			},
			"srcaddr": {
				Type:        schema.TypeList,
				Description: "Source IPv4 address name and address group names.",
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

				Description: "When enabled srcaddr/srcaddr6 specifies what the source address must NOT be.",
				Optional:    true,
				Computed:    true,
			},
			"srcaddr4": {
				Type:        schema.TypeList,
				Description: "Source IPv4 address name and address group names.",
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
			"srcaddr6": {
				Type:        schema.TypeList,
				Description: "Source IPv6 address name and address group names.",
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
				Description: "Incoming (ingress) interface.",
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

				Description: "Enable or disable this policy.",
				Optional:    true,
				Computed:    true,
			},
			"url_category": {
				Type:        schema.TypeList,
				Description: "URL category ID list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "URL category ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"users": {
				Type:        schema.TypeList,
				Description: "Names of individual users that can authenticate with this policy.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "User name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
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
			"webfilter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing Web filter profile.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallSecurityPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallSecurityPolicy resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallSecurityPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallSecurityPolicy(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallSecurityPolicy")
	}

	return resourceFirewallSecurityPolicyRead(ctx, d, meta)
}

func resourceFirewallSecurityPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallSecurityPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallSecurityPolicy(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallSecurityPolicy resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallSecurityPolicy")
	}

	return resourceFirewallSecurityPolicyRead(ctx, d, meta)
}

func resourceFirewallSecurityPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallSecurityPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallSecurityPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSecurityPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallSecurityPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallSecurityPolicy resource: %v", err)
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

	diags := refreshObjectFirewallSecurityPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallSecurityPolicyAppCategory(v *[]models.FirewallSecurityPolicyAppCategory, sort bool) interface{} {
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

func flattenFirewallSecurityPolicyAppGroup(v *[]models.FirewallSecurityPolicyAppGroup, sort bool) interface{} {
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

func flattenFirewallSecurityPolicyApplication(v *[]models.FirewallSecurityPolicyApplication, sort bool) interface{} {
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

func flattenFirewallSecurityPolicyDstaddr(v *[]models.FirewallSecurityPolicyDstaddr, sort bool) interface{} {
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

func flattenFirewallSecurityPolicyDstaddr4(v *[]models.FirewallSecurityPolicyDstaddr4, sort bool) interface{} {
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

func flattenFirewallSecurityPolicyDstaddr6(v *[]models.FirewallSecurityPolicyDstaddr6, sort bool) interface{} {
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

func flattenFirewallSecurityPolicyDstintf(v *[]models.FirewallSecurityPolicyDstintf, sort bool) interface{} {
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

func flattenFirewallSecurityPolicyFssoGroups(v *[]models.FirewallSecurityPolicyFssoGroups, sort bool) interface{} {
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

func flattenFirewallSecurityPolicyGroups(v *[]models.FirewallSecurityPolicyGroups, sort bool) interface{} {
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

func flattenFirewallSecurityPolicyInternetServiceCustom(v *[]models.FirewallSecurityPolicyInternetServiceCustom, sort bool) interface{} {
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

func flattenFirewallSecurityPolicyInternetServiceCustomGroup(v *[]models.FirewallSecurityPolicyInternetServiceCustomGroup, sort bool) interface{} {
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

func flattenFirewallSecurityPolicyInternetServiceGroup(v *[]models.FirewallSecurityPolicyInternetServiceGroup, sort bool) interface{} {
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

func flattenFirewallSecurityPolicyInternetServiceId(v *[]models.FirewallSecurityPolicyInternetServiceId, sort bool) interface{} {
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

func flattenFirewallSecurityPolicyInternetServiceName(v *[]models.FirewallSecurityPolicyInternetServiceName, sort bool) interface{} {
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

func flattenFirewallSecurityPolicyInternetServiceSrcCustom(v *[]models.FirewallSecurityPolicyInternetServiceSrcCustom, sort bool) interface{} {
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

func flattenFirewallSecurityPolicyInternetServiceSrcCustomGroup(v *[]models.FirewallSecurityPolicyInternetServiceSrcCustomGroup, sort bool) interface{} {
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

func flattenFirewallSecurityPolicyInternetServiceSrcGroup(v *[]models.FirewallSecurityPolicyInternetServiceSrcGroup, sort bool) interface{} {
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

func flattenFirewallSecurityPolicyInternetServiceSrcId(v *[]models.FirewallSecurityPolicyInternetServiceSrcId, sort bool) interface{} {
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

func flattenFirewallSecurityPolicyInternetServiceSrcName(v *[]models.FirewallSecurityPolicyInternetServiceSrcName, sort bool) interface{} {
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

func flattenFirewallSecurityPolicyService(v *[]models.FirewallSecurityPolicyService, sort bool) interface{} {
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

func flattenFirewallSecurityPolicySrcaddr(v *[]models.FirewallSecurityPolicySrcaddr, sort bool) interface{} {
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

func flattenFirewallSecurityPolicySrcaddr4(v *[]models.FirewallSecurityPolicySrcaddr4, sort bool) interface{} {
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

func flattenFirewallSecurityPolicySrcaddr6(v *[]models.FirewallSecurityPolicySrcaddr6, sort bool) interface{} {
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

func flattenFirewallSecurityPolicySrcintf(v *[]models.FirewallSecurityPolicySrcintf, sort bool) interface{} {
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

func flattenFirewallSecurityPolicyUrlCategory(v *[]models.FirewallSecurityPolicyUrlCategory, sort bool) interface{} {
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

func flattenFirewallSecurityPolicyUsers(v *[]models.FirewallSecurityPolicyUsers, sort bool) interface{} {
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

func refreshObjectFirewallSecurityPolicy(d *schema.ResourceData, o *models.FirewallSecurityPolicy, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Action != nil {
		v := *o.Action

		if err = d.Set("action", v); err != nil {
			return diag.Errorf("error reading action: %v", err)
		}
	}

	if o.AppCategory != nil {
		if err = d.Set("app_category", flattenFirewallSecurityPolicyAppCategory(o.AppCategory, sort)); err != nil {
			return diag.Errorf("error reading app_category: %v", err)
		}
	}

	if o.AppGroup != nil {
		if err = d.Set("app_group", flattenFirewallSecurityPolicyAppGroup(o.AppGroup, sort)); err != nil {
			return diag.Errorf("error reading app_group: %v", err)
		}
	}

	if o.Application != nil {
		if err = d.Set("application", flattenFirewallSecurityPolicyApplication(o.Application, sort)); err != nil {
			return diag.Errorf("error reading application: %v", err)
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

	if o.DlpSensor != nil {
		v := *o.DlpSensor

		if err = d.Set("dlp_sensor", v); err != nil {
			return diag.Errorf("error reading dlp_sensor: %v", err)
		}
	}

	if o.DnsfilterProfile != nil {
		v := *o.DnsfilterProfile

		if err = d.Set("dnsfilter_profile", v); err != nil {
			return diag.Errorf("error reading dnsfilter_profile: %v", err)
		}
	}

	if o.Dstaddr != nil {
		if err = d.Set("dstaddr", flattenFirewallSecurityPolicyDstaddr(o.Dstaddr, sort)); err != nil {
			return diag.Errorf("error reading dstaddr: %v", err)
		}
	}

	if o.DstaddrNegate != nil {
		v := *o.DstaddrNegate

		if err = d.Set("dstaddr_negate", v); err != nil {
			return diag.Errorf("error reading dstaddr_negate: %v", err)
		}
	}

	if o.Dstaddr4 != nil {
		if err = d.Set("dstaddr4", flattenFirewallSecurityPolicyDstaddr4(o.Dstaddr4, sort)); err != nil {
			return diag.Errorf("error reading dstaddr4: %v", err)
		}
	}

	if o.Dstaddr6 != nil {
		if err = d.Set("dstaddr6", flattenFirewallSecurityPolicyDstaddr6(o.Dstaddr6, sort)); err != nil {
			return diag.Errorf("error reading dstaddr6: %v", err)
		}
	}

	if o.Dstintf != nil {
		if err = d.Set("dstintf", flattenFirewallSecurityPolicyDstintf(o.Dstintf, sort)); err != nil {
			return diag.Errorf("error reading dstintf: %v", err)
		}
	}

	if o.EmailfilterProfile != nil {
		v := *o.EmailfilterProfile

		if err = d.Set("emailfilter_profile", v); err != nil {
			return diag.Errorf("error reading emailfilter_profile: %v", err)
		}
	}

	if o.EnforceDefaultAppPort != nil {
		v := *o.EnforceDefaultAppPort

		if err = d.Set("enforce_default_app_port", v); err != nil {
			return diag.Errorf("error reading enforce_default_app_port: %v", err)
		}
	}

	if o.FileFilterProfile != nil {
		v := *o.FileFilterProfile

		if err = d.Set("file_filter_profile", v); err != nil {
			return diag.Errorf("error reading file_filter_profile: %v", err)
		}
	}

	if o.FssoGroups != nil {
		if err = d.Set("fsso_groups", flattenFirewallSecurityPolicyFssoGroups(o.FssoGroups, sort)); err != nil {
			return diag.Errorf("error reading fsso_groups: %v", err)
		}
	}

	if o.Groups != nil {
		if err = d.Set("groups", flattenFirewallSecurityPolicyGroups(o.Groups, sort)); err != nil {
			return diag.Errorf("error reading groups: %v", err)
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
		if err = d.Set("internet_service_custom", flattenFirewallSecurityPolicyInternetServiceCustom(o.InternetServiceCustom, sort)); err != nil {
			return diag.Errorf("error reading internet_service_custom: %v", err)
		}
	}

	if o.InternetServiceCustomGroup != nil {
		if err = d.Set("internet_service_custom_group", flattenFirewallSecurityPolicyInternetServiceCustomGroup(o.InternetServiceCustomGroup, sort)); err != nil {
			return diag.Errorf("error reading internet_service_custom_group: %v", err)
		}
	}

	if o.InternetServiceGroup != nil {
		if err = d.Set("internet_service_group", flattenFirewallSecurityPolicyInternetServiceGroup(o.InternetServiceGroup, sort)); err != nil {
			return diag.Errorf("error reading internet_service_group: %v", err)
		}
	}

	if o.InternetServiceId != nil {
		if err = d.Set("internet_service_id", flattenFirewallSecurityPolicyInternetServiceId(o.InternetServiceId, sort)); err != nil {
			return diag.Errorf("error reading internet_service_id: %v", err)
		}
	}

	if o.InternetServiceName != nil {
		if err = d.Set("internet_service_name", flattenFirewallSecurityPolicyInternetServiceName(o.InternetServiceName, sort)); err != nil {
			return diag.Errorf("error reading internet_service_name: %v", err)
		}
	}

	if o.InternetServiceNegate != nil {
		v := *o.InternetServiceNegate

		if err = d.Set("internet_service_negate", v); err != nil {
			return diag.Errorf("error reading internet_service_negate: %v", err)
		}
	}

	if o.InternetServiceSrc != nil {
		v := *o.InternetServiceSrc

		if err = d.Set("internet_service_src", v); err != nil {
			return diag.Errorf("error reading internet_service_src: %v", err)
		}
	}

	if o.InternetServiceSrcCustom != nil {
		if err = d.Set("internet_service_src_custom", flattenFirewallSecurityPolicyInternetServiceSrcCustom(o.InternetServiceSrcCustom, sort)); err != nil {
			return diag.Errorf("error reading internet_service_src_custom: %v", err)
		}
	}

	if o.InternetServiceSrcCustomGroup != nil {
		if err = d.Set("internet_service_src_custom_group", flattenFirewallSecurityPolicyInternetServiceSrcCustomGroup(o.InternetServiceSrcCustomGroup, sort)); err != nil {
			return diag.Errorf("error reading internet_service_src_custom_group: %v", err)
		}
	}

	if o.InternetServiceSrcGroup != nil {
		if err = d.Set("internet_service_src_group", flattenFirewallSecurityPolicyInternetServiceSrcGroup(o.InternetServiceSrcGroup, sort)); err != nil {
			return diag.Errorf("error reading internet_service_src_group: %v", err)
		}
	}

	if o.InternetServiceSrcId != nil {
		if err = d.Set("internet_service_src_id", flattenFirewallSecurityPolicyInternetServiceSrcId(o.InternetServiceSrcId, sort)); err != nil {
			return diag.Errorf("error reading internet_service_src_id: %v", err)
		}
	}

	if o.InternetServiceSrcName != nil {
		if err = d.Set("internet_service_src_name", flattenFirewallSecurityPolicyInternetServiceSrcName(o.InternetServiceSrcName, sort)); err != nil {
			return diag.Errorf("error reading internet_service_src_name: %v", err)
		}
	}

	if o.InternetServiceSrcNegate != nil {
		v := *o.InternetServiceSrcNegate

		if err = d.Set("internet_service_src_negate", v); err != nil {
			return diag.Errorf("error reading internet_service_src_negate: %v", err)
		}
	}

	if o.IpsSensor != nil {
		v := *o.IpsSensor

		if err = d.Set("ips_sensor", v); err != nil {
			return diag.Errorf("error reading ips_sensor: %v", err)
		}
	}

	if o.LearningMode != nil {
		v := *o.LearningMode

		if err = d.Set("learning_mode", v); err != nil {
			return diag.Errorf("error reading learning_mode: %v", err)
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

	if o.Nat46 != nil {
		v := *o.Nat46

		if err = d.Set("nat46", v); err != nil {
			return diag.Errorf("error reading nat46: %v", err)
		}
	}

	if o.Nat64 != nil {
		v := *o.Nat64

		if err = d.Set("nat64", v); err != nil {
			return diag.Errorf("error reading nat64: %v", err)
		}
	}

	if o.Policyid != nil {
		v := *o.Policyid

		if err = d.Set("policyid", v); err != nil {
			return diag.Errorf("error reading policyid: %v", err)
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

	if o.SendDenyPacket != nil {
		v := *o.SendDenyPacket

		if err = d.Set("send_deny_packet", v); err != nil {
			return diag.Errorf("error reading send_deny_packet: %v", err)
		}
	}

	if o.Service != nil {
		if err = d.Set("service", flattenFirewallSecurityPolicyService(o.Service, sort)); err != nil {
			return diag.Errorf("error reading service: %v", err)
		}
	}

	if o.ServiceNegate != nil {
		v := *o.ServiceNegate

		if err = d.Set("service_negate", v); err != nil {
			return diag.Errorf("error reading service_negate: %v", err)
		}
	}

	if o.Srcaddr != nil {
		if err = d.Set("srcaddr", flattenFirewallSecurityPolicySrcaddr(o.Srcaddr, sort)); err != nil {
			return diag.Errorf("error reading srcaddr: %v", err)
		}
	}

	if o.SrcaddrNegate != nil {
		v := *o.SrcaddrNegate

		if err = d.Set("srcaddr_negate", v); err != nil {
			return diag.Errorf("error reading srcaddr_negate: %v", err)
		}
	}

	if o.Srcaddr4 != nil {
		if err = d.Set("srcaddr4", flattenFirewallSecurityPolicySrcaddr4(o.Srcaddr4, sort)); err != nil {
			return diag.Errorf("error reading srcaddr4: %v", err)
		}
	}

	if o.Srcaddr6 != nil {
		if err = d.Set("srcaddr6", flattenFirewallSecurityPolicySrcaddr6(o.Srcaddr6, sort)); err != nil {
			return diag.Errorf("error reading srcaddr6: %v", err)
		}
	}

	if o.Srcintf != nil {
		if err = d.Set("srcintf", flattenFirewallSecurityPolicySrcintf(o.Srcintf, sort)); err != nil {
			return diag.Errorf("error reading srcintf: %v", err)
		}
	}

	if o.SshFilterProfile != nil {
		v := *o.SshFilterProfile

		if err = d.Set("ssh_filter_profile", v); err != nil {
			return diag.Errorf("error reading ssh_filter_profile: %v", err)
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

	if o.UrlCategory != nil {
		if err = d.Set("url_category", flattenFirewallSecurityPolicyUrlCategory(o.UrlCategory, sort)); err != nil {
			return diag.Errorf("error reading url_category: %v", err)
		}
	}

	if o.Users != nil {
		if err = d.Set("users", flattenFirewallSecurityPolicyUsers(o.Users, sort)); err != nil {
			return diag.Errorf("error reading users: %v", err)
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

	if o.WebfilterProfile != nil {
		v := *o.WebfilterProfile

		if err = d.Set("webfilter_profile", v); err != nil {
			return diag.Errorf("error reading webfilter_profile: %v", err)
		}
	}

	return nil
}

func expandFirewallSecurityPolicyAppCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicyAppCategory, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicyAppCategory

	for i := range l {
		tmp := models.FirewallSecurityPolicyAppCategory{}
		var pre_append string

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

func expandFirewallSecurityPolicyAppGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicyAppGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicyAppGroup

	for i := range l {
		tmp := models.FirewallSecurityPolicyAppGroup{}
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

func expandFirewallSecurityPolicyApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicyApplication, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicyApplication

	for i := range l {
		tmp := models.FirewallSecurityPolicyApplication{}
		var pre_append string

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

func expandFirewallSecurityPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicyDstaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicyDstaddr

	for i := range l {
		tmp := models.FirewallSecurityPolicyDstaddr{}
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

func expandFirewallSecurityPolicyDstaddr4(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicyDstaddr4, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicyDstaddr4

	for i := range l {
		tmp := models.FirewallSecurityPolicyDstaddr4{}
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

func expandFirewallSecurityPolicyDstaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicyDstaddr6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicyDstaddr6

	for i := range l {
		tmp := models.FirewallSecurityPolicyDstaddr6{}
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

func expandFirewallSecurityPolicyDstintf(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicyDstintf, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicyDstintf

	for i := range l {
		tmp := models.FirewallSecurityPolicyDstintf{}
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

func expandFirewallSecurityPolicyFssoGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicyFssoGroups, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicyFssoGroups

	for i := range l {
		tmp := models.FirewallSecurityPolicyFssoGroups{}
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

func expandFirewallSecurityPolicyGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicyGroups, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicyGroups

	for i := range l {
		tmp := models.FirewallSecurityPolicyGroups{}
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

func expandFirewallSecurityPolicyInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicyInternetServiceCustom, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicyInternetServiceCustom

	for i := range l {
		tmp := models.FirewallSecurityPolicyInternetServiceCustom{}
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

func expandFirewallSecurityPolicyInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicyInternetServiceCustomGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicyInternetServiceCustomGroup

	for i := range l {
		tmp := models.FirewallSecurityPolicyInternetServiceCustomGroup{}
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

func expandFirewallSecurityPolicyInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicyInternetServiceGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicyInternetServiceGroup

	for i := range l {
		tmp := models.FirewallSecurityPolicyInternetServiceGroup{}
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

func expandFirewallSecurityPolicyInternetServiceId(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicyInternetServiceId, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicyInternetServiceId

	for i := range l {
		tmp := models.FirewallSecurityPolicyInternetServiceId{}
		var pre_append string

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

func expandFirewallSecurityPolicyInternetServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicyInternetServiceName, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicyInternetServiceName

	for i := range l {
		tmp := models.FirewallSecurityPolicyInternetServiceName{}
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

func expandFirewallSecurityPolicyInternetServiceSrcCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicyInternetServiceSrcCustom, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicyInternetServiceSrcCustom

	for i := range l {
		tmp := models.FirewallSecurityPolicyInternetServiceSrcCustom{}
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

func expandFirewallSecurityPolicyInternetServiceSrcCustomGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicyInternetServiceSrcCustomGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicyInternetServiceSrcCustomGroup

	for i := range l {
		tmp := models.FirewallSecurityPolicyInternetServiceSrcCustomGroup{}
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

func expandFirewallSecurityPolicyInternetServiceSrcGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicyInternetServiceSrcGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicyInternetServiceSrcGroup

	for i := range l {
		tmp := models.FirewallSecurityPolicyInternetServiceSrcGroup{}
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

func expandFirewallSecurityPolicyInternetServiceSrcId(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicyInternetServiceSrcId, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicyInternetServiceSrcId

	for i := range l {
		tmp := models.FirewallSecurityPolicyInternetServiceSrcId{}
		var pre_append string

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

func expandFirewallSecurityPolicyInternetServiceSrcName(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicyInternetServiceSrcName, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicyInternetServiceSrcName

	for i := range l {
		tmp := models.FirewallSecurityPolicyInternetServiceSrcName{}
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

func expandFirewallSecurityPolicyService(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicyService, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicyService

	for i := range l {
		tmp := models.FirewallSecurityPolicyService{}
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

func expandFirewallSecurityPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicySrcaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicySrcaddr

	for i := range l {
		tmp := models.FirewallSecurityPolicySrcaddr{}
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

func expandFirewallSecurityPolicySrcaddr4(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicySrcaddr4, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicySrcaddr4

	for i := range l {
		tmp := models.FirewallSecurityPolicySrcaddr4{}
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

func expandFirewallSecurityPolicySrcaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicySrcaddr6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicySrcaddr6

	for i := range l {
		tmp := models.FirewallSecurityPolicySrcaddr6{}
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

func expandFirewallSecurityPolicySrcintf(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicySrcintf, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicySrcintf

	for i := range l {
		tmp := models.FirewallSecurityPolicySrcintf{}
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

func expandFirewallSecurityPolicyUrlCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicyUrlCategory, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicyUrlCategory

	for i := range l {
		tmp := models.FirewallSecurityPolicyUrlCategory{}
		var pre_append string

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

func expandFirewallSecurityPolicyUsers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSecurityPolicyUsers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSecurityPolicyUsers

	for i := range l {
		tmp := models.FirewallSecurityPolicyUsers{}
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

func getObjectFirewallSecurityPolicy(d *schema.ResourceData, sv string) (*models.FirewallSecurityPolicy, diag.Diagnostics) {
	obj := models.FirewallSecurityPolicy{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("action", sv)
				diags = append(diags, e)
			}
			obj.Action = &v2
		}
	}
	if v, ok := d.GetOk("app_category"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("app_category", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicyAppCategory(d, v, "app_category", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.AppCategory = t
		}
	} else if d.HasChange("app_category") {
		old, new := d.GetChange("app_category")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.AppCategory = &[]models.FirewallSecurityPolicyAppCategory{}
		}
	}
	if v, ok := d.GetOk("app_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("app_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicyAppGroup(d, v, "app_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.AppGroup = t
		}
	} else if d.HasChange("app_group") {
		old, new := d.GetChange("app_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.AppGroup = &[]models.FirewallSecurityPolicyAppGroup{}
		}
	}
	if v, ok := d.GetOk("application"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("application", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicyApplication(d, v, "application", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Application = t
		}
	} else if d.HasChange("application") {
		old, new := d.GetChange("application")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Application = &[]models.FirewallSecurityPolicyApplication{}
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
	if v1, ok := d.GetOk("dlp_sensor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dlp_sensor", sv)
				diags = append(diags, e)
			}
			obj.DlpSensor = &v2
		}
	}
	if v1, ok := d.GetOk("dnsfilter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dnsfilter_profile", sv)
				diags = append(diags, e)
			}
			obj.DnsfilterProfile = &v2
		}
	}
	if v, ok := d.GetOk("dstaddr"); ok {
		if !utils.CheckVer(sv, "v6.4.0", "") {
			e := utils.AttributeVersionWarning("dstaddr", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicyDstaddr(d, v, "dstaddr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstaddr = t
		}
	} else if d.HasChange("dstaddr") {
		old, new := d.GetChange("dstaddr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstaddr = &[]models.FirewallSecurityPolicyDstaddr{}
		}
	}
	if v1, ok := d.GetOk("dstaddr_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("dstaddr_negate", sv)
				diags = append(diags, e)
			}
			obj.DstaddrNegate = &v2
		}
	}
	if v, ok := d.GetOk("dstaddr4"); ok {
		if !utils.CheckVer(sv, "", "v6.4.0") {
			e := utils.AttributeVersionWarning("dstaddr4", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicyDstaddr4(d, v, "dstaddr4", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstaddr4 = t
		}
	} else if d.HasChange("dstaddr4") {
		old, new := d.GetChange("dstaddr4")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstaddr4 = &[]models.FirewallSecurityPolicyDstaddr4{}
		}
	}
	if v, ok := d.GetOk("dstaddr6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dstaddr6", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicyDstaddr6(d, v, "dstaddr6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstaddr6 = t
		}
	} else if d.HasChange("dstaddr6") {
		old, new := d.GetChange("dstaddr6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstaddr6 = &[]models.FirewallSecurityPolicyDstaddr6{}
		}
	}
	if v, ok := d.GetOk("dstintf"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dstintf", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicyDstintf(d, v, "dstintf", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstintf = t
		}
	} else if d.HasChange("dstintf") {
		old, new := d.GetChange("dstintf")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstintf = &[]models.FirewallSecurityPolicyDstintf{}
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
	if v1, ok := d.GetOk("enforce_default_app_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("enforce_default_app_port", sv)
				diags = append(diags, e)
			}
			obj.EnforceDefaultAppPort = &v2
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
	if v, ok := d.GetOk("fsso_groups"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("fsso_groups", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicyFssoGroups(d, v, "fsso_groups", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.FssoGroups = t
		}
	} else if d.HasChange("fsso_groups") {
		old, new := d.GetChange("fsso_groups")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.FssoGroups = &[]models.FirewallSecurityPolicyFssoGroups{}
		}
	}
	if v, ok := d.GetOk("groups"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("groups", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicyGroups(d, v, "groups", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Groups = t
		}
	} else if d.HasChange("groups") {
		old, new := d.GetChange("groups")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Groups = &[]models.FirewallSecurityPolicyGroups{}
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
		t, err := expandFirewallSecurityPolicyInternetServiceCustom(d, v, "internet_service_custom", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceCustom = t
		}
	} else if d.HasChange("internet_service_custom") {
		old, new := d.GetChange("internet_service_custom")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceCustom = &[]models.FirewallSecurityPolicyInternetServiceCustom{}
		}
	}
	if v, ok := d.GetOk("internet_service_custom_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_custom_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicyInternetServiceCustomGroup(d, v, "internet_service_custom_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceCustomGroup = t
		}
	} else if d.HasChange("internet_service_custom_group") {
		old, new := d.GetChange("internet_service_custom_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceCustomGroup = &[]models.FirewallSecurityPolicyInternetServiceCustomGroup{}
		}
	}
	if v, ok := d.GetOk("internet_service_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicyInternetServiceGroup(d, v, "internet_service_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceGroup = t
		}
	} else if d.HasChange("internet_service_group") {
		old, new := d.GetChange("internet_service_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceGroup = &[]models.FirewallSecurityPolicyInternetServiceGroup{}
		}
	}
	if v, ok := d.GetOk("internet_service_id"); ok {
		if !utils.CheckVer(sv, "", "v6.4.0") {
			e := utils.AttributeVersionWarning("internet_service_id", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicyInternetServiceId(d, v, "internet_service_id", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceId = t
		}
	} else if d.HasChange("internet_service_id") {
		old, new := d.GetChange("internet_service_id")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceId = &[]models.FirewallSecurityPolicyInternetServiceId{}
		}
	}
	if v, ok := d.GetOk("internet_service_name"); ok {
		if !utils.CheckVer(sv, "v6.4.0", "") {
			e := utils.AttributeVersionWarning("internet_service_name", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicyInternetServiceName(d, v, "internet_service_name", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceName = t
		}
	} else if d.HasChange("internet_service_name") {
		old, new := d.GetChange("internet_service_name")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceName = &[]models.FirewallSecurityPolicyInternetServiceName{}
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
	if v1, ok := d.GetOk("internet_service_src"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("internet_service_src", sv)
				diags = append(diags, e)
			}
			obj.InternetServiceSrc = &v2
		}
	}
	if v, ok := d.GetOk("internet_service_src_custom"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_src_custom", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicyInternetServiceSrcCustom(d, v, "internet_service_src_custom", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceSrcCustom = t
		}
	} else if d.HasChange("internet_service_src_custom") {
		old, new := d.GetChange("internet_service_src_custom")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceSrcCustom = &[]models.FirewallSecurityPolicyInternetServiceSrcCustom{}
		}
	}
	if v, ok := d.GetOk("internet_service_src_custom_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_src_custom_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicyInternetServiceSrcCustomGroup(d, v, "internet_service_src_custom_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceSrcCustomGroup = t
		}
	} else if d.HasChange("internet_service_src_custom_group") {
		old, new := d.GetChange("internet_service_src_custom_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceSrcCustomGroup = &[]models.FirewallSecurityPolicyInternetServiceSrcCustomGroup{}
		}
	}
	if v, ok := d.GetOk("internet_service_src_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_src_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicyInternetServiceSrcGroup(d, v, "internet_service_src_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceSrcGroup = t
		}
	} else if d.HasChange("internet_service_src_group") {
		old, new := d.GetChange("internet_service_src_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceSrcGroup = &[]models.FirewallSecurityPolicyInternetServiceSrcGroup{}
		}
	}
	if v, ok := d.GetOk("internet_service_src_id"); ok {
		if !utils.CheckVer(sv, "", "v6.4.0") {
			e := utils.AttributeVersionWarning("internet_service_src_id", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicyInternetServiceSrcId(d, v, "internet_service_src_id", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceSrcId = t
		}
	} else if d.HasChange("internet_service_src_id") {
		old, new := d.GetChange("internet_service_src_id")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceSrcId = &[]models.FirewallSecurityPolicyInternetServiceSrcId{}
		}
	}
	if v, ok := d.GetOk("internet_service_src_name"); ok {
		if !utils.CheckVer(sv, "v6.4.0", "") {
			e := utils.AttributeVersionWarning("internet_service_src_name", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicyInternetServiceSrcName(d, v, "internet_service_src_name", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceSrcName = t
		}
	} else if d.HasChange("internet_service_src_name") {
		old, new := d.GetChange("internet_service_src_name")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceSrcName = &[]models.FirewallSecurityPolicyInternetServiceSrcName{}
		}
	}
	if v1, ok := d.GetOk("internet_service_src_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("internet_service_src_negate", sv)
				diags = append(diags, e)
			}
			obj.InternetServiceSrcNegate = &v2
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
	if v1, ok := d.GetOk("learning_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("learning_mode", sv)
				diags = append(diags, e)
			}
			obj.LearningMode = &v2
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
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("logtraffic_start", sv)
				diags = append(diags, e)
			}
			obj.LogtrafficStart = &v2
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
	if v1, ok := d.GetOk("nat46"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("nat46", sv)
				diags = append(diags, e)
			}
			obj.Nat46 = &v2
		}
	}
	if v1, ok := d.GetOk("nat64"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("nat64", sv)
				diags = append(diags, e)
			}
			obj.Nat64 = &v2
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
	if v1, ok := d.GetOk("send_deny_packet"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("send_deny_packet", sv)
				diags = append(diags, e)
			}
			obj.SendDenyPacket = &v2
		}
	}
	if v, ok := d.GetOk("service"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("service", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicyService(d, v, "service", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Service = t
		}
	} else if d.HasChange("service") {
		old, new := d.GetChange("service")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Service = &[]models.FirewallSecurityPolicyService{}
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
	if v, ok := d.GetOk("srcaddr"); ok {
		if !utils.CheckVer(sv, "v6.4.0", "") {
			e := utils.AttributeVersionWarning("srcaddr", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicySrcaddr(d, v, "srcaddr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcaddr = t
		}
	} else if d.HasChange("srcaddr") {
		old, new := d.GetChange("srcaddr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcaddr = &[]models.FirewallSecurityPolicySrcaddr{}
		}
	}
	if v1, ok := d.GetOk("srcaddr_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("srcaddr_negate", sv)
				diags = append(diags, e)
			}
			obj.SrcaddrNegate = &v2
		}
	}
	if v, ok := d.GetOk("srcaddr4"); ok {
		if !utils.CheckVer(sv, "", "v6.4.0") {
			e := utils.AttributeVersionWarning("srcaddr4", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicySrcaddr4(d, v, "srcaddr4", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcaddr4 = t
		}
	} else if d.HasChange("srcaddr4") {
		old, new := d.GetChange("srcaddr4")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcaddr4 = &[]models.FirewallSecurityPolicySrcaddr4{}
		}
	}
	if v, ok := d.GetOk("srcaddr6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("srcaddr6", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicySrcaddr6(d, v, "srcaddr6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcaddr6 = t
		}
	} else if d.HasChange("srcaddr6") {
		old, new := d.GetChange("srcaddr6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcaddr6 = &[]models.FirewallSecurityPolicySrcaddr6{}
		}
	}
	if v, ok := d.GetOk("srcintf"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("srcintf", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicySrcintf(d, v, "srcintf", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcintf = t
		}
	} else if d.HasChange("srcintf") {
		old, new := d.GetChange("srcintf")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcintf = &[]models.FirewallSecurityPolicySrcintf{}
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
	if v, ok := d.GetOk("url_category"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("url_category", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicyUrlCategory(d, v, "url_category", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.UrlCategory = t
		}
	} else if d.HasChange("url_category") {
		old, new := d.GetChange("url_category")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.UrlCategory = &[]models.FirewallSecurityPolicyUrlCategory{}
		}
	}
	if v, ok := d.GetOk("users"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("users", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSecurityPolicyUsers(d, v, "users", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Users = t
		}
	} else if d.HasChange("users") {
		old, new := d.GetChange("users")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Users = &[]models.FirewallSecurityPolicyUsers{}
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
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("voip_profile", sv)
				diags = append(diags, e)
			}
			obj.VoipProfile = &v2
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
	return &obj, diags
}
