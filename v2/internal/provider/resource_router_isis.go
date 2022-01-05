// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IS-IS.

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

func resourceRouterIsis() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IS-IS.",

		CreateContext: resourceRouterIsisCreate,
		ReadContext:   resourceRouterIsisRead,
		UpdateContext: resourceRouterIsisUpdate,
		DeleteContext: resourceRouterIsisDelete,

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
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"adjacency_check": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable adjacency check.",
				Optional:    true,
				Computed:    true,
			},
			"adjacency_check6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPv6 adjacency check.",
				Optional:    true,
				Computed:    true,
			},
			"adv_passive_only": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IS-IS advertisement of passive interfaces only.",
				Optional:    true,
				Computed:    true,
			},
			"adv_passive_only6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPv6 IS-IS advertisement of passive interfaces only.",
				Optional:    true,
				Computed:    true,
			},
			"auth_keychain_l1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Authentication key-chain for level 1 PDUs.",
				Optional:    true,
				Computed:    true,
			},
			"auth_keychain_l2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Authentication key-chain for level 2 PDUs.",
				Optional:    true,
				Computed:    true,
			},
			"auth_mode_l1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"password", "md5"}, false),

				Description: "Level 1 authentication mode.",
				Optional:    true,
				Computed:    true,
			},
			"auth_mode_l2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"password", "md5"}, false),

				Description: "Level 2 authentication mode.",
				Optional:    true,
				Computed:    true,
			},
			"auth_password_l1": {
				Type: schema.TypeString,

				Description: "Authentication password for level 1 PDUs.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"auth_password_l2": {
				Type: schema.TypeString,

				Description: "Authentication password for level 2 PDUs.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"auth_sendonly_l1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable level 1 authentication send-only.",
				Optional:    true,
				Computed:    true,
			},
			"auth_sendonly_l2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable level 2 authentication send-only.",
				Optional:    true,
				Computed:    true,
			},
			"default_originate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable distribution of default route information.",
				Optional:    true,
				Computed:    true,
			},
			"default_originate6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable distribution of default IPv6 route information.",
				Optional:    true,
				Computed:    true,
			},
			"dynamic_hostname": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable dynamic hostname.",
				Optional:    true,
				Computed:    true,
			},
			"ignore_lsp_errors": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable ignoring of LSP errors with bad checksums.",
				Optional:    true,
				Computed:    true,
			},
			"is_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"level-1-2", "level-1", "level-2-only"}, false),

				Description: "IS type.",
				Optional:    true,
				Computed:    true,
			},
			"isis_interface": {
				Type:        schema.TypeList,
				Description: "IS-IS interface configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_keychain_l1": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Authentication key-chain for level 1 PDUs.",
							Optional:    true,
							Computed:    true,
						},
						"auth_keychain_l2": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Authentication key-chain for level 2 PDUs.",
							Optional:    true,
							Computed:    true,
						},
						"auth_mode_l1": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"md5", "password"}, false),

							Description: "Level 1 authentication mode.",
							Optional:    true,
							Computed:    true,
						},
						"auth_mode_l2": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"md5", "password"}, false),

							Description: "Level 2 authentication mode.",
							Optional:    true,
							Computed:    true,
						},
						"auth_password_l1": {
							Type: schema.TypeString,

							Description: "Authentication password for level 1 PDUs.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"auth_password_l2": {
							Type: schema.TypeString,

							Description: "Authentication password for level 2 PDUs.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"auth_send_only_l1": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable authentication send-only for level 1 PDUs.",
							Optional:    true,
							Computed:    true,
						},
						"auth_send_only_l2": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable authentication send-only for level 2 PDUs.",
							Optional:    true,
							Computed:    true,
						},
						"circuit_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"level-1-2", "level-1", "level-2"}, false),

							Description: "IS-IS interface's circuit type",
							Optional:    true,
							Computed:    true,
						},
						"csnp_interval_l1": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Level 1 CSNP interval.",
							Optional:    true,
							Computed:    true,
						},
						"csnp_interval_l2": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Level 2 CSNP interval.",
							Optional:    true,
							Computed:    true,
						},
						"hello_interval_l1": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Level 1 hello interval.",
							Optional:    true,
							Computed:    true,
						},
						"hello_interval_l2": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Level 2 hello interval.",
							Optional:    true,
							Computed:    true,
						},
						"hello_multiplier_l1": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 100),

							Description: "Level 1 multiplier for Hello holding time.",
							Optional:    true,
							Computed:    true,
						},
						"hello_multiplier_l2": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 100),

							Description: "Level 2 multiplier for Hello holding time.",
							Optional:    true,
							Computed:    true,
						},
						"hello_padding": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable padding to IS-IS hello packets.",
							Optional:    true,
							Computed:    true,
						},
						"lsp_interval": {
							Type: schema.TypeInt,

							Description: "LSP transmission interval (milliseconds).",
							Optional:    true,
							Computed:    true,
						},
						"lsp_retransmit_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "LSP retransmission interval (sec).",
							Optional:    true,
							Computed:    true,
						},
						"mesh_group": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable IS-IS mesh group.",
							Optional:    true,
							Computed:    true,
						},
						"mesh_group_id": {
							Type: schema.TypeInt,

							Description: "Mesh group ID <0-4294967295>, 0: mesh-group blocked.",
							Optional:    true,
							Computed:    true,
						},
						"metric_l1": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 63),

							Description: "Level 1 metric for interface.",
							Optional:    true,
							Computed:    true,
						},
						"metric_l2": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 63),

							Description: "Level 2 metric for interface.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "IS-IS interface name.",
							Optional:    true,
							Computed:    true,
						},
						"network_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"broadcast", "point-to-point", "loopback"}, false),

							Description: "IS-IS interface's network type",
							Optional:    true,
							Computed:    true,
						},
						"priority_l1": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 127),

							Description: "Level 1 priority.",
							Optional:    true,
							Computed:    true,
						},
						"priority_l2": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 127),

							Description: "Level 2 priority.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable interface for IS-IS.",
							Optional:    true,
							Computed:    true,
						},
						"status6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable IPv6 interface for IS-IS.",
							Optional:    true,
							Computed:    true,
						},
						"wide_metric_l1": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 16777214),

							Description: "Level 1 wide metric for interface.",
							Optional:    true,
							Computed:    true,
						},
						"wide_metric_l2": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 16777214),

							Description: "Level 2 wide metric for interface.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"isis_net": {
				Type:        schema.TypeList,
				Description: "IS-IS net configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "isis-net ID.",
							Optional:    true,
							Computed:    true,
						},
						"net": {
							Type: schema.TypeString,

							Description: "IS-IS net xx.xxxx. ... .xxxx.xx.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"lsp_gen_interval_l1": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 120),

				Description: "Minimum interval for level 1 LSP regenerating.",
				Optional:    true,
				Computed:    true,
			},
			"lsp_gen_interval_l2": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 120),

				Description: "Minimum interval for level 2 LSP regenerating.",
				Optional:    true,
				Computed:    true,
			},
			"lsp_refresh_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "LSP refresh time in seconds.",
				Optional:    true,
				Computed:    true,
			},
			"max_lsp_lifetime": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(350, 65535),

				Description: "Maximum LSP lifetime in seconds.",
				Optional:    true,
				Computed:    true,
			},
			"metric_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"narrow", "wide", "transition", "narrow-transition", "narrow-transition-l1", "narrow-transition-l2", "wide-l1", "wide-l2", "wide-transition", "wide-transition-l1", "wide-transition-l2", "transition-l1", "transition-l2"}, false),

				Description: "Use old-style (ISO 10589) or new-style packet formats",
				Optional:    true,
				Computed:    true,
			},
			"overload_bit": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable signal other routers not to use us in SPF.",
				Optional:    true,
				Computed:    true,
			},
			"overload_bit_on_startup": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 86400),

				Description: "Overload-bit only temporarily after reboot.",
				Optional:    true,
				Computed:    true,
			},
			"overload_bit_suppress": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Suppress overload-bit for the specific prefixes.",
				Optional:         true,
				Computed:         true,
			},
			"redistribute": {
				Type:        schema.TypeList,
				Description: "IS-IS redistribute protocols.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"level": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"level-1-2", "level-1", "level-2"}, false),

							Description: "Level.",
							Optional:    true,
							Computed:    true,
						},
						"metric": {
							Type: schema.TypeInt,

							Description: "Metric.",
							Optional:    true,
							Computed:    true,
						},
						"metric_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"external", "internal"}, false),

							Description: "Metric type.",
							Optional:    true,
							Computed:    true,
						},
						"protocol": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Protocol name.",
							Optional:    true,
							Computed:    true,
						},
						"routemap": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Route map name.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Status.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"redistribute_l1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable redistribution of level 1 routes into level 2.",
				Optional:    true,
				Computed:    true,
			},
			"redistribute_l1_list": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Access-list for route redistribution from l1 to l2.",
				Optional:    true,
				Computed:    true,
			},
			"redistribute_l2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable redistribution of level 2 routes into level 1.",
				Optional:    true,
				Computed:    true,
			},
			"redistribute_l2_list": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Access-list for route redistribution from l2 to l1.",
				Optional:    true,
				Computed:    true,
			},
			"redistribute6": {
				Type:        schema.TypeList,
				Description: "IS-IS IPv6 redistribution for routing protocols.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"level": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"level-1-2", "level-1", "level-2"}, false),

							Description: "Level.",
							Optional:    true,
							Computed:    true,
						},
						"metric": {
							Type: schema.TypeInt,

							Description: "Metric.",
							Optional:    true,
							Computed:    true,
						},
						"metric_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"external", "internal"}, false),

							Description: "Metric type.",
							Optional:    true,
							Computed:    true,
						},
						"protocol": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Protocol name.",
							Optional:    true,
							Computed:    true,
						},
						"routemap": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Route map name.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable redistribution.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"redistribute6_l1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable redistribution of level 1 IPv6 routes into level 2.",
				Optional:    true,
				Computed:    true,
			},
			"redistribute6_l1_list": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Access-list for IPv6 route redistribution from l1 to l2.",
				Optional:    true,
				Computed:    true,
			},
			"redistribute6_l2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable redistribution of level 2 IPv6 routes into level 1.",
				Optional:    true,
				Computed:    true,
			},
			"redistribute6_l2_list": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Access-list for IPv6 route redistribution from l2 to l1.",
				Optional:    true,
				Computed:    true,
			},
			"spf_interval_exp_l1": {
				Type: schema.TypeString,

				Description: "Level 1 SPF calculation delay.",
				Optional:    true,
				Computed:    true,
			},
			"spf_interval_exp_l2": {
				Type: schema.TypeString,

				Description: "Level 2 SPF calculation delay.",
				Optional:    true,
				Computed:    true,
			},
			"summary_address": {
				Type:        schema.TypeList,
				Description: "IS-IS summary addresses.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Summary address entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"level": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"level-1-2", "level-1", "level-2"}, false),

							Description: "Level.",
							Optional:    true,
							Computed:    true,
						},
						"prefix": {
							Type:         schema.TypeString,
							ValidateFunc: validators.FortiValidateIPv4ClassnetAny,

							Description: "Prefix.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"summary_address6": {
				Type:        schema.TypeList,
				Description: "IS-IS IPv6 summary address.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Prefix entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"level": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"level-1-2", "level-1", "level-2"}, false),

							Description: "Level.",
							Optional:    true,
							Computed:    true,
						},
						"prefix6": {
							Type:             schema.TypeString,
							ValidateFunc:     validators.FortiValidateIPv6Prefix,
							DiffSuppressFunc: suppressors.DiffCidrEqual,
							Description:      "IPv6 prefix.",
							Optional:         true,
							Computed:         true,
						},
					},
				},
			},
		},
	}
}

func resourceRouterIsisCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterIsis(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterIsis(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterIsis")
	}

	return resourceRouterIsisRead(ctx, d, meta)
}

func resourceRouterIsisUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterIsis(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterIsis(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterIsis resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterIsis")
	}

	return resourceRouterIsisRead(ctx, d, meta)
}

func resourceRouterIsisDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectRouterIsis(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateRouterIsis(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterIsis resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterIsisRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterIsis(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterIsis resource: %v", err)
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

	diags := refreshObjectRouterIsis(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenRouterIsisIsisInterface(v *[]models.RouterIsisIsisInterface, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AuthKeychainL1; tmp != nil {
				v["auth_keychain_l1"] = *tmp
			}

			if tmp := cfg.AuthKeychainL2; tmp != nil {
				v["auth_keychain_l2"] = *tmp
			}

			if tmp := cfg.AuthModeL1; tmp != nil {
				v["auth_mode_l1"] = *tmp
			}

			if tmp := cfg.AuthModeL2; tmp != nil {
				v["auth_mode_l2"] = *tmp
			}

			if tmp := cfg.AuthPasswordL1; tmp != nil {
				v["auth_password_l1"] = *tmp
			}

			if tmp := cfg.AuthPasswordL2; tmp != nil {
				v["auth_password_l2"] = *tmp
			}

			if tmp := cfg.AuthSendOnlyL1; tmp != nil {
				v["auth_send_only_l1"] = *tmp
			}

			if tmp := cfg.AuthSendOnlyL2; tmp != nil {
				v["auth_send_only_l2"] = *tmp
			}

			if tmp := cfg.CircuitType; tmp != nil {
				v["circuit_type"] = *tmp
			}

			if tmp := cfg.CsnpIntervalL1; tmp != nil {
				v["csnp_interval_l1"] = *tmp
			}

			if tmp := cfg.CsnpIntervalL2; tmp != nil {
				v["csnp_interval_l2"] = *tmp
			}

			if tmp := cfg.HelloIntervalL1; tmp != nil {
				v["hello_interval_l1"] = *tmp
			}

			if tmp := cfg.HelloIntervalL2; tmp != nil {
				v["hello_interval_l2"] = *tmp
			}

			if tmp := cfg.HelloMultiplierL1; tmp != nil {
				v["hello_multiplier_l1"] = *tmp
			}

			if tmp := cfg.HelloMultiplierL2; tmp != nil {
				v["hello_multiplier_l2"] = *tmp
			}

			if tmp := cfg.HelloPadding; tmp != nil {
				v["hello_padding"] = *tmp
			}

			if tmp := cfg.LspInterval; tmp != nil {
				v["lsp_interval"] = *tmp
			}

			if tmp := cfg.LspRetransmitInterval; tmp != nil {
				v["lsp_retransmit_interval"] = *tmp
			}

			if tmp := cfg.MeshGroup; tmp != nil {
				v["mesh_group"] = *tmp
			}

			if tmp := cfg.MeshGroupId; tmp != nil {
				v["mesh_group_id"] = *tmp
			}

			if tmp := cfg.MetricL1; tmp != nil {
				v["metric_l1"] = *tmp
			}

			if tmp := cfg.MetricL2; tmp != nil {
				v["metric_l2"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.NetworkType; tmp != nil {
				v["network_type"] = *tmp
			}

			if tmp := cfg.PriorityL1; tmp != nil {
				v["priority_l1"] = *tmp
			}

			if tmp := cfg.PriorityL2; tmp != nil {
				v["priority_l2"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Status6; tmp != nil {
				v["status6"] = *tmp
			}

			if tmp := cfg.WideMetricL1; tmp != nil {
				v["wide_metric_l1"] = *tmp
			}

			if tmp := cfg.WideMetricL2; tmp != nil {
				v["wide_metric_l2"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenRouterIsisIsisNet(v *[]models.RouterIsisIsisNet, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Net; tmp != nil {
				v["net"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterIsisRedistribute(v *[]models.RouterIsisRedistribute, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Level; tmp != nil {
				v["level"] = *tmp
			}

			if tmp := cfg.Metric; tmp != nil {
				v["metric"] = *tmp
			}

			if tmp := cfg.MetricType; tmp != nil {
				v["metric_type"] = *tmp
			}

			if tmp := cfg.Protocol; tmp != nil {
				v["protocol"] = *tmp
			}

			if tmp := cfg.Routemap; tmp != nil {
				v["routemap"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "protocol")
	}

	return flat
}

func flattenRouterIsisRedistribute6(v *[]models.RouterIsisRedistribute6, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Level; tmp != nil {
				v["level"] = *tmp
			}

			if tmp := cfg.Metric; tmp != nil {
				v["metric"] = *tmp
			}

			if tmp := cfg.MetricType; tmp != nil {
				v["metric_type"] = *tmp
			}

			if tmp := cfg.Protocol; tmp != nil {
				v["protocol"] = *tmp
			}

			if tmp := cfg.Routemap; tmp != nil {
				v["routemap"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "protocol")
	}

	return flat
}

func flattenRouterIsisSummaryAddress(v *[]models.RouterIsisSummaryAddress, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Level; tmp != nil {
				v["level"] = *tmp
			}

			if tmp := cfg.Prefix; tmp != nil {
				v["prefix"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterIsisSummaryAddress6(v *[]models.RouterIsisSummaryAddress6, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Level; tmp != nil {
				v["level"] = *tmp
			}

			if tmp := cfg.Prefix6; tmp != nil {
				v["prefix6"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectRouterIsis(d *schema.ResourceData, o *models.RouterIsis, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AdjacencyCheck != nil {
		v := *o.AdjacencyCheck

		if err = d.Set("adjacency_check", v); err != nil {
			return diag.Errorf("error reading adjacency_check: %v", err)
		}
	}

	if o.AdjacencyCheck6 != nil {
		v := *o.AdjacencyCheck6

		if err = d.Set("adjacency_check6", v); err != nil {
			return diag.Errorf("error reading adjacency_check6: %v", err)
		}
	}

	if o.AdvPassiveOnly != nil {
		v := *o.AdvPassiveOnly

		if err = d.Set("adv_passive_only", v); err != nil {
			return diag.Errorf("error reading adv_passive_only: %v", err)
		}
	}

	if o.AdvPassiveOnly6 != nil {
		v := *o.AdvPassiveOnly6

		if err = d.Set("adv_passive_only6", v); err != nil {
			return diag.Errorf("error reading adv_passive_only6: %v", err)
		}
	}

	if o.AuthKeychainL1 != nil {
		v := *o.AuthKeychainL1

		if err = d.Set("auth_keychain_l1", v); err != nil {
			return diag.Errorf("error reading auth_keychain_l1: %v", err)
		}
	}

	if o.AuthKeychainL2 != nil {
		v := *o.AuthKeychainL2

		if err = d.Set("auth_keychain_l2", v); err != nil {
			return diag.Errorf("error reading auth_keychain_l2: %v", err)
		}
	}

	if o.AuthModeL1 != nil {
		v := *o.AuthModeL1

		if err = d.Set("auth_mode_l1", v); err != nil {
			return diag.Errorf("error reading auth_mode_l1: %v", err)
		}
	}

	if o.AuthModeL2 != nil {
		v := *o.AuthModeL2

		if err = d.Set("auth_mode_l2", v); err != nil {
			return diag.Errorf("error reading auth_mode_l2: %v", err)
		}
	}

	if o.AuthPasswordL1 != nil {
		v := *o.AuthPasswordL1

		if err = d.Set("auth_password_l1", v); err != nil {
			return diag.Errorf("error reading auth_password_l1: %v", err)
		}
	}

	if o.AuthPasswordL2 != nil {
		v := *o.AuthPasswordL2

		if err = d.Set("auth_password_l2", v); err != nil {
			return diag.Errorf("error reading auth_password_l2: %v", err)
		}
	}

	if o.AuthSendonlyL1 != nil {
		v := *o.AuthSendonlyL1

		if err = d.Set("auth_sendonly_l1", v); err != nil {
			return diag.Errorf("error reading auth_sendonly_l1: %v", err)
		}
	}

	if o.AuthSendonlyL2 != nil {
		v := *o.AuthSendonlyL2

		if err = d.Set("auth_sendonly_l2", v); err != nil {
			return diag.Errorf("error reading auth_sendonly_l2: %v", err)
		}
	}

	if o.DefaultOriginate != nil {
		v := *o.DefaultOriginate

		if err = d.Set("default_originate", v); err != nil {
			return diag.Errorf("error reading default_originate: %v", err)
		}
	}

	if o.DefaultOriginate6 != nil {
		v := *o.DefaultOriginate6

		if err = d.Set("default_originate6", v); err != nil {
			return diag.Errorf("error reading default_originate6: %v", err)
		}
	}

	if o.DynamicHostname != nil {
		v := *o.DynamicHostname

		if err = d.Set("dynamic_hostname", v); err != nil {
			return diag.Errorf("error reading dynamic_hostname: %v", err)
		}
	}

	if o.IgnoreLspErrors != nil {
		v := *o.IgnoreLspErrors

		if err = d.Set("ignore_lsp_errors", v); err != nil {
			return diag.Errorf("error reading ignore_lsp_errors: %v", err)
		}
	}

	if o.IsType != nil {
		v := *o.IsType

		if err = d.Set("is_type", v); err != nil {
			return diag.Errorf("error reading is_type: %v", err)
		}
	}

	if o.IsisInterface != nil {
		if err = d.Set("isis_interface", flattenRouterIsisIsisInterface(o.IsisInterface, sort)); err != nil {
			return diag.Errorf("error reading isis_interface: %v", err)
		}
	}

	if o.IsisNet != nil {
		if err = d.Set("isis_net", flattenRouterIsisIsisNet(o.IsisNet, sort)); err != nil {
			return diag.Errorf("error reading isis_net: %v", err)
		}
	}

	if o.LspGenIntervalL1 != nil {
		v := *o.LspGenIntervalL1

		if err = d.Set("lsp_gen_interval_l1", v); err != nil {
			return diag.Errorf("error reading lsp_gen_interval_l1: %v", err)
		}
	}

	if o.LspGenIntervalL2 != nil {
		v := *o.LspGenIntervalL2

		if err = d.Set("lsp_gen_interval_l2", v); err != nil {
			return diag.Errorf("error reading lsp_gen_interval_l2: %v", err)
		}
	}

	if o.LspRefreshInterval != nil {
		v := *o.LspRefreshInterval

		if err = d.Set("lsp_refresh_interval", v); err != nil {
			return diag.Errorf("error reading lsp_refresh_interval: %v", err)
		}
	}

	if o.MaxLspLifetime != nil {
		v := *o.MaxLspLifetime

		if err = d.Set("max_lsp_lifetime", v); err != nil {
			return diag.Errorf("error reading max_lsp_lifetime: %v", err)
		}
	}

	if o.MetricStyle != nil {
		v := *o.MetricStyle

		if err = d.Set("metric_style", v); err != nil {
			return diag.Errorf("error reading metric_style: %v", err)
		}
	}

	if o.OverloadBit != nil {
		v := *o.OverloadBit

		if err = d.Set("overload_bit", v); err != nil {
			return diag.Errorf("error reading overload_bit: %v", err)
		}
	}

	if o.OverloadBitOnStartup != nil {
		v := *o.OverloadBitOnStartup

		if err = d.Set("overload_bit_on_startup", v); err != nil {
			return diag.Errorf("error reading overload_bit_on_startup: %v", err)
		}
	}

	if o.OverloadBitSuppress != nil {
		v := *o.OverloadBitSuppress

		if err = d.Set("overload_bit_suppress", v); err != nil {
			return diag.Errorf("error reading overload_bit_suppress: %v", err)
		}
	}

	if o.Redistribute != nil {
		if err = d.Set("redistribute", flattenRouterIsisRedistribute(o.Redistribute, sort)); err != nil {
			return diag.Errorf("error reading redistribute: %v", err)
		}
	}

	if o.RedistributeL1 != nil {
		v := *o.RedistributeL1

		if err = d.Set("redistribute_l1", v); err != nil {
			return diag.Errorf("error reading redistribute_l1: %v", err)
		}
	}

	if o.RedistributeL1List != nil {
		v := *o.RedistributeL1List

		if err = d.Set("redistribute_l1_list", v); err != nil {
			return diag.Errorf("error reading redistribute_l1_list: %v", err)
		}
	}

	if o.RedistributeL2 != nil {
		v := *o.RedistributeL2

		if err = d.Set("redistribute_l2", v); err != nil {
			return diag.Errorf("error reading redistribute_l2: %v", err)
		}
	}

	if o.RedistributeL2List != nil {
		v := *o.RedistributeL2List

		if err = d.Set("redistribute_l2_list", v); err != nil {
			return diag.Errorf("error reading redistribute_l2_list: %v", err)
		}
	}

	if o.Redistribute6 != nil {
		if err = d.Set("redistribute6", flattenRouterIsisRedistribute6(o.Redistribute6, sort)); err != nil {
			return diag.Errorf("error reading redistribute6: %v", err)
		}
	}

	if o.Redistribute6L1 != nil {
		v := *o.Redistribute6L1

		if err = d.Set("redistribute6_l1", v); err != nil {
			return diag.Errorf("error reading redistribute6_l1: %v", err)
		}
	}

	if o.Redistribute6L1List != nil {
		v := *o.Redistribute6L1List

		if err = d.Set("redistribute6_l1_list", v); err != nil {
			return diag.Errorf("error reading redistribute6_l1_list: %v", err)
		}
	}

	if o.Redistribute6L2 != nil {
		v := *o.Redistribute6L2

		if err = d.Set("redistribute6_l2", v); err != nil {
			return diag.Errorf("error reading redistribute6_l2: %v", err)
		}
	}

	if o.Redistribute6L2List != nil {
		v := *o.Redistribute6L2List

		if err = d.Set("redistribute6_l2_list", v); err != nil {
			return diag.Errorf("error reading redistribute6_l2_list: %v", err)
		}
	}

	if o.SpfIntervalExpL1 != nil {
		v := *o.SpfIntervalExpL1

		if err = d.Set("spf_interval_exp_l1", v); err != nil {
			return diag.Errorf("error reading spf_interval_exp_l1: %v", err)
		}
	}

	if o.SpfIntervalExpL2 != nil {
		v := *o.SpfIntervalExpL2

		if err = d.Set("spf_interval_exp_l2", v); err != nil {
			return diag.Errorf("error reading spf_interval_exp_l2: %v", err)
		}
	}

	if o.SummaryAddress != nil {
		if err = d.Set("summary_address", flattenRouterIsisSummaryAddress(o.SummaryAddress, sort)); err != nil {
			return diag.Errorf("error reading summary_address: %v", err)
		}
	}

	if o.SummaryAddress6 != nil {
		if err = d.Set("summary_address6", flattenRouterIsisSummaryAddress6(o.SummaryAddress6, sort)); err != nil {
			return diag.Errorf("error reading summary_address6: %v", err)
		}
	}

	return nil
}

func expandRouterIsisIsisInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterIsisIsisInterface, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterIsisIsisInterface

	for i := range l {
		tmp := models.RouterIsisIsisInterface{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.auth_keychain_l1", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthKeychainL1 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auth_keychain_l2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthKeychainL2 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auth_mode_l1", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthModeL1 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auth_mode_l2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthModeL2 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auth_password_l1", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthPasswordL1 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auth_password_l2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthPasswordL2 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auth_send_only_l1", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthSendOnlyL1 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auth_send_only_l2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthSendOnlyL2 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.circuit_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CircuitType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.csnp_interval_l1", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.CsnpIntervalL1 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.csnp_interval_l2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.CsnpIntervalL2 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.hello_interval_l1", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.HelloIntervalL1 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.hello_interval_l2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.HelloIntervalL2 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.hello_multiplier_l1", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.HelloMultiplierL1 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.hello_multiplier_l2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.HelloMultiplierL2 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.hello_padding", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HelloPadding = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.lsp_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.LspInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.lsp_retransmit_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.LspRetransmitInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mesh_group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MeshGroup = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mesh_group_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.MeshGroupId = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.metric_l1", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.MetricL1 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.metric_l2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.MetricL2 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.network_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NetworkType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority_l1", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.PriorityL1 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority_l2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.PriorityL2 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.wide_metric_l1", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.WideMetricL1 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.wide_metric_l2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.WideMetricL2 = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterIsisIsisNet(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterIsisIsisNet, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterIsisIsisNet

	for i := range l {
		tmp := models.RouterIsisIsisNet{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.net", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Net = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterIsisRedistribute(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterIsisRedistribute, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterIsisRedistribute

	for i := range l {
		tmp := models.RouterIsisRedistribute{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Level = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.metric", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Metric = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.metric_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MetricType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Protocol = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.routemap", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Routemap = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterIsisRedistribute6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterIsisRedistribute6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterIsisRedistribute6

	for i := range l {
		tmp := models.RouterIsisRedistribute6{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Level = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.metric", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Metric = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.metric_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MetricType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Protocol = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.routemap", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Routemap = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterIsisSummaryAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterIsisSummaryAddress, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterIsisSummaryAddress

	for i := range l {
		tmp := models.RouterIsisSummaryAddress{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Level = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Prefix = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterIsisSummaryAddress6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterIsisSummaryAddress6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterIsisSummaryAddress6

	for i := range l {
		tmp := models.RouterIsisSummaryAddress6{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Level = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Prefix6 = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectRouterIsis(d *schema.ResourceData, sv string) (*models.RouterIsis, diag.Diagnostics) {
	obj := models.RouterIsis{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("adjacency_check"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("adjacency_check", sv)
				diags = append(diags, e)
			}
			obj.AdjacencyCheck = &v2
		}
	}
	if v1, ok := d.GetOk("adjacency_check6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("adjacency_check6", sv)
				diags = append(diags, e)
			}
			obj.AdjacencyCheck6 = &v2
		}
	}
	if v1, ok := d.GetOk("adv_passive_only"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("adv_passive_only", sv)
				diags = append(diags, e)
			}
			obj.AdvPassiveOnly = &v2
		}
	}
	if v1, ok := d.GetOk("adv_passive_only6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("adv_passive_only6", sv)
				diags = append(diags, e)
			}
			obj.AdvPassiveOnly6 = &v2
		}
	}
	if v1, ok := d.GetOk("auth_keychain_l1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_keychain_l1", sv)
				diags = append(diags, e)
			}
			obj.AuthKeychainL1 = &v2
		}
	}
	if v1, ok := d.GetOk("auth_keychain_l2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_keychain_l2", sv)
				diags = append(diags, e)
			}
			obj.AuthKeychainL2 = &v2
		}
	}
	if v1, ok := d.GetOk("auth_mode_l1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_mode_l1", sv)
				diags = append(diags, e)
			}
			obj.AuthModeL1 = &v2
		}
	}
	if v1, ok := d.GetOk("auth_mode_l2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_mode_l2", sv)
				diags = append(diags, e)
			}
			obj.AuthModeL2 = &v2
		}
	}
	if v1, ok := d.GetOk("auth_password_l1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_password_l1", sv)
				diags = append(diags, e)
			}
			obj.AuthPasswordL1 = &v2
		}
	}
	if v1, ok := d.GetOk("auth_password_l2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_password_l2", sv)
				diags = append(diags, e)
			}
			obj.AuthPasswordL2 = &v2
		}
	}
	if v1, ok := d.GetOk("auth_sendonly_l1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_sendonly_l1", sv)
				diags = append(diags, e)
			}
			obj.AuthSendonlyL1 = &v2
		}
	}
	if v1, ok := d.GetOk("auth_sendonly_l2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_sendonly_l2", sv)
				diags = append(diags, e)
			}
			obj.AuthSendonlyL2 = &v2
		}
	}
	if v1, ok := d.GetOk("default_originate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_originate", sv)
				diags = append(diags, e)
			}
			obj.DefaultOriginate = &v2
		}
	}
	if v1, ok := d.GetOk("default_originate6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_originate6", sv)
				diags = append(diags, e)
			}
			obj.DefaultOriginate6 = &v2
		}
	}
	if v1, ok := d.GetOk("dynamic_hostname"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dynamic_hostname", sv)
				diags = append(diags, e)
			}
			obj.DynamicHostname = &v2
		}
	}
	if v1, ok := d.GetOk("ignore_lsp_errors"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ignore_lsp_errors", sv)
				diags = append(diags, e)
			}
			obj.IgnoreLspErrors = &v2
		}
	}
	if v1, ok := d.GetOk("is_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("is_type", sv)
				diags = append(diags, e)
			}
			obj.IsType = &v2
		}
	}
	if v, ok := d.GetOk("isis_interface"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("isis_interface", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterIsisIsisInterface(d, v, "isis_interface", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.IsisInterface = t
		}
	} else if d.HasChange("isis_interface") {
		old, new := d.GetChange("isis_interface")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.IsisInterface = &[]models.RouterIsisIsisInterface{}
		}
	}
	if v, ok := d.GetOk("isis_net"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("isis_net", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterIsisIsisNet(d, v, "isis_net", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.IsisNet = t
		}
	} else if d.HasChange("isis_net") {
		old, new := d.GetChange("isis_net")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.IsisNet = &[]models.RouterIsisIsisNet{}
		}
	}
	if v1, ok := d.GetOk("lsp_gen_interval_l1"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lsp_gen_interval_l1", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LspGenIntervalL1 = &tmp
		}
	}
	if v1, ok := d.GetOk("lsp_gen_interval_l2"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lsp_gen_interval_l2", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LspGenIntervalL2 = &tmp
		}
	}
	if v1, ok := d.GetOk("lsp_refresh_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lsp_refresh_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LspRefreshInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("max_lsp_lifetime"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_lsp_lifetime", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxLspLifetime = &tmp
		}
	}
	if v1, ok := d.GetOk("metric_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("metric_style", sv)
				diags = append(diags, e)
			}
			obj.MetricStyle = &v2
		}
	}
	if v1, ok := d.GetOk("overload_bit"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("overload_bit", sv)
				diags = append(diags, e)
			}
			obj.OverloadBit = &v2
		}
	}
	if v1, ok := d.GetOk("overload_bit_on_startup"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("overload_bit_on_startup", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.OverloadBitOnStartup = &tmp
		}
	}
	if v1, ok := d.GetOk("overload_bit_suppress"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("overload_bit_suppress", sv)
				diags = append(diags, e)
			}
			obj.OverloadBitSuppress = &v2
		}
	}
	if v, ok := d.GetOk("redistribute"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("redistribute", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterIsisRedistribute(d, v, "redistribute", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Redistribute = t
		}
	} else if d.HasChange("redistribute") {
		old, new := d.GetChange("redistribute")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Redistribute = &[]models.RouterIsisRedistribute{}
		}
	}
	if v1, ok := d.GetOk("redistribute_l1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("redistribute_l1", sv)
				diags = append(diags, e)
			}
			obj.RedistributeL1 = &v2
		}
	}
	if v1, ok := d.GetOk("redistribute_l1_list"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("redistribute_l1_list", sv)
				diags = append(diags, e)
			}
			obj.RedistributeL1List = &v2
		}
	}
	if v1, ok := d.GetOk("redistribute_l2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("redistribute_l2", sv)
				diags = append(diags, e)
			}
			obj.RedistributeL2 = &v2
		}
	}
	if v1, ok := d.GetOk("redistribute_l2_list"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("redistribute_l2_list", sv)
				diags = append(diags, e)
			}
			obj.RedistributeL2List = &v2
		}
	}
	if v, ok := d.GetOk("redistribute6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("redistribute6", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterIsisRedistribute6(d, v, "redistribute6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Redistribute6 = t
		}
	} else if d.HasChange("redistribute6") {
		old, new := d.GetChange("redistribute6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Redistribute6 = &[]models.RouterIsisRedistribute6{}
		}
	}
	if v1, ok := d.GetOk("redistribute6_l1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("redistribute6_l1", sv)
				diags = append(diags, e)
			}
			obj.Redistribute6L1 = &v2
		}
	}
	if v1, ok := d.GetOk("redistribute6_l1_list"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("redistribute6_l1_list", sv)
				diags = append(diags, e)
			}
			obj.Redistribute6L1List = &v2
		}
	}
	if v1, ok := d.GetOk("redistribute6_l2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("redistribute6_l2", sv)
				diags = append(diags, e)
			}
			obj.Redistribute6L2 = &v2
		}
	}
	if v1, ok := d.GetOk("redistribute6_l2_list"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("redistribute6_l2_list", sv)
				diags = append(diags, e)
			}
			obj.Redistribute6L2List = &v2
		}
	}
	if v1, ok := d.GetOk("spf_interval_exp_l1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("spf_interval_exp_l1", sv)
				diags = append(diags, e)
			}
			obj.SpfIntervalExpL1 = &v2
		}
	}
	if v1, ok := d.GetOk("spf_interval_exp_l2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("spf_interval_exp_l2", sv)
				diags = append(diags, e)
			}
			obj.SpfIntervalExpL2 = &v2
		}
	}
	if v, ok := d.GetOk("summary_address"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("summary_address", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterIsisSummaryAddress(d, v, "summary_address", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SummaryAddress = t
		}
	} else if d.HasChange("summary_address") {
		old, new := d.GetChange("summary_address")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SummaryAddress = &[]models.RouterIsisSummaryAddress{}
		}
	}
	if v, ok := d.GetOk("summary_address6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("summary_address6", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterIsisSummaryAddress6(d, v, "summary_address6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SummaryAddress6 = t
		}
	} else if d.HasChange("summary_address6") {
		old, new := d.GetChange("summary_address6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SummaryAddress6 = &[]models.RouterIsisSummaryAddress6{}
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectRouterIsis(d *schema.ResourceData, sv string) (*models.RouterIsis, diag.Diagnostics) {
	obj := models.RouterIsis{}
	diags := diag.Diagnostics{}

	obj.IsisInterface = &[]models.RouterIsisIsisInterface{}
	obj.IsisNet = &[]models.RouterIsisIsisNet{}
	obj.Redistribute = &[]models.RouterIsisRedistribute{}
	obj.Redistribute6 = &[]models.RouterIsisRedistribute6{}
	obj.SummaryAddress = &[]models.RouterIsisSummaryAddress{}
	obj.SummaryAddress6 = &[]models.RouterIsisSummaryAddress6{}

	return &obj, diags
}
