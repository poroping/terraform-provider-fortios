// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv6 OSPF.

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

func resourceRouterOspf6() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv6 OSPF.",

		CreateContext: resourceRouterOspf6Create,
		ReadContext:   resourceRouterOspf6Read,
		UpdateContext: resourceRouterOspf6Update,
		DeleteContext: resourceRouterOspf6Delete,

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
			"abr_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"cisco", "ibm", "standard"}, false),

				Description: "Area border router type.",
				Optional:    true,
				Computed:    true,
			},
			"area": {
				Type:        schema.TypeList,
				Description: "OSPF6 area configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "ah", "esp"}, false),

							Description: "Authentication mode.",
							Optional:    true,
							Computed:    true,
						},
						"default_cost": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 16777215),

							Description: "Summary default cost of stub or NSSA area.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Area entry IP address.",
							Optional:    true,
							Computed:    true,
						},
						"ipsec_auth_alg": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"md5", "sha1", "sha256", "sha384", "sha512"}, false),

							Description: "Authentication algorithm.",
							Optional:    true,
							Computed:    true,
						},
						"ipsec_enc_alg": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"null", "des", "3des", "aes128", "aes192", "aes256"}, false),

							Description: "Encryption algorithm.",
							Optional:    true,
							Computed:    true,
						},
						"ipsec_keys": {
							Type:        schema.TypeList,
							Description: "IPsec authentication and encryption keys.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_key": {
										Type: schema.TypeString,

										Description: "Authentication key.",
										Optional:    true,
										Computed:    true,
										Sensitive:   true,
									},
									"enc_key": {
										Type: schema.TypeString,

										Description: "Encryption key.",
										Optional:    true,
										Computed:    true,
										Sensitive:   true,
									},
									"spi": {
										Type: schema.TypeInt,

										Description: "Security Parameters Index.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"key_rollover_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(300, 216000),

							Description: "Key roll-over interval.",
							Optional:    true,
							Computed:    true,
						},
						"nssa_default_information_originate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable originate type 7 default into NSSA area.",
							Optional:    true,
							Computed:    true,
						},
						"nssa_default_information_originate_metric": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 16777214),

							Description: "OSPFv3 default metric.",
							Optional:    true,
							Computed:    true,
						},
						"nssa_default_information_originate_metric_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"1", "2"}, false),

							Description: "OSPFv3 metric type for default routes.",
							Optional:    true,
							Computed:    true,
						},
						"nssa_redistribution": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable redistribute into NSSA area.",
							Optional:    true,
							Computed:    true,
						},
						"nssa_translator_role": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"candidate", "never", "always"}, false),

							Description: "NSSA translator role type.",
							Optional:    true,
							Computed:    true,
						},
						"range": {
							Type:        schema.TypeList,
							Description: "OSPF6 area range configuration.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"advertise": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

										Description: "Enable/disable advertise status.",
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Type: schema.TypeInt,

										Description: "Range entry ID.",
										Optional:    true,
										Computed:    true,
									},
									"prefix6": {
										Type:             schema.TypeString,
										ValidateFunc:     validators.FortiValidateIPv6Network,
										DiffSuppressFunc: suppressors.DiffCidrEqual,
										Description:      "IPv6 prefix.",
										Optional:         true,
										Computed:         true,
									},
								},
							},
						},
						"stub_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"no-summary", "summary"}, false),

							Description: "Stub summary setting.",
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"regular", "nssa", "stub"}, false),

							Description: "Area type setting.",
							Optional:    true,
							Computed:    true,
						},
						"virtual_link": {
							Type:        schema.TypeList,
							Description: "OSPF6 virtual link configuration.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"none", "ah", "esp", "area"}, false),

										Description: "Authentication mode.",
										Optional:    true,
										Computed:    true,
									},
									"dead_interval": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),

										Description: "Dead interval.",
										Optional:    true,
										Computed:    true,
									},
									"hello_interval": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),

										Description: "Hello interval.",
										Optional:    true,
										Computed:    true,
									},
									"ipsec_auth_alg": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"md5", "sha1", "sha256", "sha384", "sha512"}, false),

										Description: "Authentication algorithm.",
										Optional:    true,
										Computed:    true,
									},
									"ipsec_enc_alg": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"null", "des", "3des", "aes128", "aes192", "aes256"}, false),

										Description: "Encryption algorithm.",
										Optional:    true,
										Computed:    true,
									},
									"ipsec_keys": {
										Type:        schema.TypeList,
										Description: "IPsec authentication and encryption keys.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"auth_key": {
													Type: schema.TypeString,

													Description: "Authentication key.",
													Optional:    true,
													Computed:    true,
													Sensitive:   true,
												},
												"enc_key": {
													Type: schema.TypeString,

													Description: "Encryption key.",
													Optional:    true,
													Computed:    true,
													Sensitive:   true,
												},
												"spi": {
													Type: schema.TypeInt,

													Description: "Security Parameters Index.",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"key_rollover_interval": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(300, 216000),

										Description: "Key roll-over interval.",
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Virtual link entry name.",
										Optional:    true,
										Computed:    true,
									},
									"peer": {
										Type:         schema.TypeString,
										ValidateFunc: validation.IsIPv4Address,

										Description: "A.B.C.D, peer router ID.",
										Optional:    true,
										Computed:    true,
									},
									"retransmit_interval": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),

										Description: "Retransmit interval.",
										Optional:    true,
										Computed:    true,
									},
									"transmit_delay": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),

										Description: "Transmit delay.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"auto_cost_ref_bandwidth": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1000000),

				Description: "Reference bandwidth in terms of megabits per second.",
				Optional:    true,
				Computed:    true,
			},
			"bfd": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Bidirectional Forwarding Detection (BFD).",
				Optional:    true,
				Computed:    true,
			},
			"default_information_metric": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16777214),

				Description: "Default information metric.",
				Optional:    true,
				Computed:    true,
			},
			"default_information_metric_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"1", "2"}, false),

				Description: "Default information metric type.",
				Optional:    true,
				Computed:    true,
			},
			"default_information_originate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "always", "disable"}, false),

				Description: "Enable/disable generation of default route.",
				Optional:    true,
				Computed:    true,
			},
			"default_information_route_map": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Default information route map.",
				Optional:    true,
				Computed:    true,
			},
			"default_metric": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16777214),

				Description: "Default metric of redistribute routes.",
				Optional:    true,
				Computed:    true,
			},
			"log_neighbour_changes": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable logging of OSPFv3 neighbour's changes",
				Optional:    true,
				Computed:    true,
			},
			"ospf6_interface": {
				Type:        schema.TypeList,
				Description: "OSPF6 interface configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"area_id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "A.B.C.D, in IPv4 address format.",
							Optional:    true,
							Computed:    true,
						},
						"authentication": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "ah", "esp", "area"}, false),

							Description: "Authentication mode.",
							Optional:    true,
							Computed:    true,
						},
						"bfd": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"global", "enable", "disable"}, false),

							Description: "Enable/disable Bidirectional Forwarding Detection (BFD).",
							Optional:    true,
							Computed:    true,
						},
						"cost": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Cost of the interface, value range from 0 to 65535, 0 means auto-cost.",
							Optional:    true,
							Computed:    true,
						},
						"dead_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Dead interval.",
							Optional:    true,
							Computed:    true,
						},
						"hello_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Hello interval.",
							Optional:    true,
							Computed:    true,
						},
						"interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Configuration interface name.",
							Optional:    true,
							Computed:    true,
						},
						"ipsec_auth_alg": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"md5", "sha1", "sha256", "sha384", "sha512"}, false),

							Description: "Authentication algorithm.",
							Optional:    true,
							Computed:    true,
						},
						"ipsec_enc_alg": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"null", "des", "3des", "aes128", "aes192", "aes256"}, false),

							Description: "Encryption algorithm.",
							Optional:    true,
							Computed:    true,
						},
						"ipsec_keys": {
							Type:        schema.TypeList,
							Description: "IPsec authentication and encryption keys.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_key": {
										Type: schema.TypeString,

										Description: "Authentication key.",
										Optional:    true,
										Computed:    true,
										Sensitive:   true,
									},
									"enc_key": {
										Type: schema.TypeString,

										Description: "Encryption key.",
										Optional:    true,
										Computed:    true,
										Sensitive:   true,
									},
									"spi": {
										Type: schema.TypeInt,

										Description: "Security Parameters Index.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"key_rollover_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(300, 216000),

							Description: "Key roll-over interval.",
							Optional:    true,
							Computed:    true,
						},
						"mtu": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(576, 65535),

							Description: "MTU for OSPFv3 packets.",
							Optional:    true,
							Computed:    true,
						},
						"mtu_ignore": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable ignoring MTU field in DBD packets.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Interface entry name.",
							Optional:    true,
							Computed:    true,
						},
						"neighbor": {
							Type:        schema.TypeList,
							Description: "OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cost": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),

										Description: "Cost of the interface, value range from 0 to 65535, 0 means auto-cost.",
										Optional:    true,
										Computed:    true,
									},
									"ip6": {
										Type:             schema.TypeString,
										ValidateFunc:     validation.IsIPv6Address,
										DiffSuppressFunc: suppressors.DiffIPEqual,
										Description:      "IPv6 link local address of the neighbor.",
										Optional:         true,
										Computed:         true,
									},
									"poll_interval": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),

										Description: "Poll interval time in seconds.",
										Optional:    true,
										Computed:    true,
									},
									"priority": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 255),

										Description: "priority",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"network_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"broadcast", "point-to-point", "non-broadcast", "point-to-multipoint", "point-to-multipoint-non-broadcast"}, false),

							Description: "Network type.",
							Optional:    true,
							Computed:    true,
						},
						"priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "priority",
							Optional:    true,
							Computed:    true,
						},
						"retransmit_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Retransmit interval.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable OSPF6 routing on this interface.",
							Optional:    true,
							Computed:    true,
						},
						"transmit_delay": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Transmit delay.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"passive_interface": {
				Type:        schema.TypeList,
				Description: "Passive interface configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Passive interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"redistribute": {
				Type:        schema.TypeList,
				Description: "Redistribute configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metric": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 16777214),

							Description: "Redistribute metric setting.",
							Optional:    true,
							Computed:    true,
						},
						"metric_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"1", "2"}, false),

							Description: "Metric type.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Redistribute name.",
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

							Description: "status",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"router_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "A.B.C.D, in IPv4 address format.",
				Optional:    true,
				Computed:    true,
			},
			"spf_timers": {
				Type: schema.TypeString,

				Description: "SPF calculation frequency.",
				Optional:    true,
				Computed:    true,
			},
			"summary_address": {
				Type:        schema.TypeList,
				Description: "IPv6 address summary configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advertise": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable advertise status.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Summary address entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"prefix6": {
							Type:             schema.TypeString,
							ValidateFunc:     validators.FortiValidateIPv6Network,
							DiffSuppressFunc: suppressors.DiffCidrEqual,
							Description:      "IPv6 prefix.",
							Optional:         true,
							Computed:         true,
						},
						"tag": {
							Type: schema.TypeInt,

							Description: "Tag value.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceRouterOspf6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterOspf6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterOspf6(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterOspf6")
	}

	return resourceRouterOspf6Read(ctx, d, meta)
}

func resourceRouterOspf6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterOspf6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterOspf6(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterOspf6 resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterOspf6")
	}

	return resourceRouterOspf6Read(ctx, d, meta)
}

func resourceRouterOspf6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectRouterOspf6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateRouterOspf6(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterOspf6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspf6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterOspf6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterOspf6 resource: %v", err)
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

	diags := refreshObjectRouterOspf6(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenRouterOspf6Area(v *[]models.RouterOspf6Area, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Authentication; tmp != nil {
				v["authentication"] = *tmp
			}

			if tmp := cfg.DefaultCost; tmp != nil {
				v["default_cost"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.IpsecAuthAlg; tmp != nil {
				v["ipsec_auth_alg"] = *tmp
			}

			if tmp := cfg.IpsecEncAlg; tmp != nil {
				v["ipsec_enc_alg"] = *tmp
			}

			if tmp := cfg.IpsecKeys; tmp != nil {
				v["ipsec_keys"] = flattenRouterOspf6AreaIpsecKeys(tmp, sort)
			}

			if tmp := cfg.KeyRolloverInterval; tmp != nil {
				v["key_rollover_interval"] = *tmp
			}

			if tmp := cfg.NssaDefaultInformationOriginate; tmp != nil {
				v["nssa_default_information_originate"] = *tmp
			}

			if tmp := cfg.NssaDefaultInformationOriginateMetric; tmp != nil {
				v["nssa_default_information_originate_metric"] = *tmp
			}

			if tmp := cfg.NssaDefaultInformationOriginateMetricType; tmp != nil {
				v["nssa_default_information_originate_metric_type"] = *tmp
			}

			if tmp := cfg.NssaRedistribution; tmp != nil {
				v["nssa_redistribution"] = *tmp
			}

			if tmp := cfg.NssaTranslatorRole; tmp != nil {
				v["nssa_translator_role"] = *tmp
			}

			if tmp := cfg.Range; tmp != nil {
				v["range"] = flattenRouterOspf6AreaRange(tmp, sort)
			}

			if tmp := cfg.StubType; tmp != nil {
				v["stub_type"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			if tmp := cfg.VirtualLink; tmp != nil {
				v["virtual_link"] = flattenRouterOspf6AreaVirtualLink(tmp, sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterOspf6AreaIpsecKeys(v *[]models.RouterOspf6AreaIpsecKeys, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AuthKey; tmp != nil {
				v["auth_key"] = *tmp
			}

			if tmp := cfg.EncKey; tmp != nil {
				v["enc_key"] = *tmp
			}

			if tmp := cfg.Spi; tmp != nil {
				v["spi"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "spi")
	}

	return flat
}

func flattenRouterOspf6AreaRange(v *[]models.RouterOspf6AreaRange, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Advertise; tmp != nil {
				v["advertise"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
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

func flattenRouterOspf6AreaVirtualLink(v *[]models.RouterOspf6AreaVirtualLink, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Authentication; tmp != nil {
				v["authentication"] = *tmp
			}

			if tmp := cfg.DeadInterval; tmp != nil {
				v["dead_interval"] = *tmp
			}

			if tmp := cfg.HelloInterval; tmp != nil {
				v["hello_interval"] = *tmp
			}

			if tmp := cfg.IpsecAuthAlg; tmp != nil {
				v["ipsec_auth_alg"] = *tmp
			}

			if tmp := cfg.IpsecEncAlg; tmp != nil {
				v["ipsec_enc_alg"] = *tmp
			}

			if tmp := cfg.IpsecKeys; tmp != nil {
				v["ipsec_keys"] = flattenRouterOspf6AreaVirtualLinkIpsecKeys(tmp, sort)
			}

			if tmp := cfg.KeyRolloverInterval; tmp != nil {
				v["key_rollover_interval"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Peer; tmp != nil {
				v["peer"] = *tmp
			}

			if tmp := cfg.RetransmitInterval; tmp != nil {
				v["retransmit_interval"] = *tmp
			}

			if tmp := cfg.TransmitDelay; tmp != nil {
				v["transmit_delay"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenRouterOspf6AreaVirtualLinkIpsecKeys(v *[]models.RouterOspf6AreaVirtualLinkIpsecKeys, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AuthKey; tmp != nil {
				v["auth_key"] = *tmp
			}

			if tmp := cfg.EncKey; tmp != nil {
				v["enc_key"] = *tmp
			}

			if tmp := cfg.Spi; tmp != nil {
				v["spi"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "spi")
	}

	return flat
}

func flattenRouterOspf6Ospf6Interface(v *[]models.RouterOspf6Ospf6Interface, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AreaId; tmp != nil {
				v["area_id"] = *tmp
			}

			if tmp := cfg.Authentication; tmp != nil {
				v["authentication"] = *tmp
			}

			if tmp := cfg.Bfd; tmp != nil {
				v["bfd"] = *tmp
			}

			if tmp := cfg.Cost; tmp != nil {
				v["cost"] = *tmp
			}

			if tmp := cfg.DeadInterval; tmp != nil {
				v["dead_interval"] = *tmp
			}

			if tmp := cfg.HelloInterval; tmp != nil {
				v["hello_interval"] = *tmp
			}

			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			if tmp := cfg.IpsecAuthAlg; tmp != nil {
				v["ipsec_auth_alg"] = *tmp
			}

			if tmp := cfg.IpsecEncAlg; tmp != nil {
				v["ipsec_enc_alg"] = *tmp
			}

			if tmp := cfg.IpsecKeys; tmp != nil {
				v["ipsec_keys"] = flattenRouterOspf6Ospf6InterfaceIpsecKeys(tmp, sort)
			}

			if tmp := cfg.KeyRolloverInterval; tmp != nil {
				v["key_rollover_interval"] = *tmp
			}

			if tmp := cfg.Mtu; tmp != nil {
				v["mtu"] = *tmp
			}

			if tmp := cfg.MtuIgnore; tmp != nil {
				v["mtu_ignore"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Neighbor; tmp != nil {
				v["neighbor"] = flattenRouterOspf6Ospf6InterfaceNeighbor(tmp, sort)
			}

			if tmp := cfg.NetworkType; tmp != nil {
				v["network_type"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			if tmp := cfg.RetransmitInterval; tmp != nil {
				v["retransmit_interval"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.TransmitDelay; tmp != nil {
				v["transmit_delay"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenRouterOspf6Ospf6InterfaceIpsecKeys(v *[]models.RouterOspf6Ospf6InterfaceIpsecKeys, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AuthKey; tmp != nil {
				v["auth_key"] = *tmp
			}

			if tmp := cfg.EncKey; tmp != nil {
				v["enc_key"] = *tmp
			}

			if tmp := cfg.Spi; tmp != nil {
				v["spi"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "spi")
	}

	return flat
}

func flattenRouterOspf6Ospf6InterfaceNeighbor(v *[]models.RouterOspf6Ospf6InterfaceNeighbor, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Cost; tmp != nil {
				v["cost"] = *tmp
			}

			if tmp := cfg.Ip6; tmp != nil {
				v["ip6"] = *tmp
			}

			if tmp := cfg.PollInterval; tmp != nil {
				v["poll_interval"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "ip6")
	}

	return flat
}

func flattenRouterOspf6PassiveInterface(v *[]models.RouterOspf6PassiveInterface, sort bool) interface{} {
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

func flattenRouterOspf6Redistribute(v *[]models.RouterOspf6Redistribute, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Metric; tmp != nil {
				v["metric"] = *tmp
			}

			if tmp := cfg.MetricType; tmp != nil {
				v["metric_type"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
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
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenRouterOspf6SummaryAddress(v *[]models.RouterOspf6SummaryAddress, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Advertise; tmp != nil {
				v["advertise"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Prefix6; tmp != nil {
				v["prefix6"] = *tmp
			}

			if tmp := cfg.Tag; tmp != nil {
				v["tag"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectRouterOspf6(d *schema.ResourceData, o *models.RouterOspf6, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AbrType != nil {
		v := *o.AbrType

		if err = d.Set("abr_type", v); err != nil {
			return diag.Errorf("error reading abr_type: %v", err)
		}
	}

	if o.Area != nil {
		if err = d.Set("area", flattenRouterOspf6Area(o.Area, sort)); err != nil {
			return diag.Errorf("error reading area: %v", err)
		}
	}

	if o.AutoCostRefBandwidth != nil {
		v := *o.AutoCostRefBandwidth

		if err = d.Set("auto_cost_ref_bandwidth", v); err != nil {
			return diag.Errorf("error reading auto_cost_ref_bandwidth: %v", err)
		}
	}

	if o.Bfd != nil {
		v := *o.Bfd

		if err = d.Set("bfd", v); err != nil {
			return diag.Errorf("error reading bfd: %v", err)
		}
	}

	if o.DefaultInformationMetric != nil {
		v := *o.DefaultInformationMetric

		if err = d.Set("default_information_metric", v); err != nil {
			return diag.Errorf("error reading default_information_metric: %v", err)
		}
	}

	if o.DefaultInformationMetricType != nil {
		v := *o.DefaultInformationMetricType

		if err = d.Set("default_information_metric_type", v); err != nil {
			return diag.Errorf("error reading default_information_metric_type: %v", err)
		}
	}

	if o.DefaultInformationOriginate != nil {
		v := *o.DefaultInformationOriginate

		if err = d.Set("default_information_originate", v); err != nil {
			return diag.Errorf("error reading default_information_originate: %v", err)
		}
	}

	if o.DefaultInformationRouteMap != nil {
		v := *o.DefaultInformationRouteMap

		if err = d.Set("default_information_route_map", v); err != nil {
			return diag.Errorf("error reading default_information_route_map: %v", err)
		}
	}

	if o.DefaultMetric != nil {
		v := *o.DefaultMetric

		if err = d.Set("default_metric", v); err != nil {
			return diag.Errorf("error reading default_metric: %v", err)
		}
	}

	if o.LogNeighbourChanges != nil {
		v := *o.LogNeighbourChanges

		if err = d.Set("log_neighbour_changes", v); err != nil {
			return diag.Errorf("error reading log_neighbour_changes: %v", err)
		}
	}

	if o.Ospf6Interface != nil {
		if err = d.Set("ospf6_interface", flattenRouterOspf6Ospf6Interface(o.Ospf6Interface, sort)); err != nil {
			return diag.Errorf("error reading ospf6_interface: %v", err)
		}
	}

	if o.PassiveInterface != nil {
		if err = d.Set("passive_interface", flattenRouterOspf6PassiveInterface(o.PassiveInterface, sort)); err != nil {
			return diag.Errorf("error reading passive_interface: %v", err)
		}
	}

	if o.Redistribute != nil {
		if err = d.Set("redistribute", flattenRouterOspf6Redistribute(o.Redistribute, sort)); err != nil {
			return diag.Errorf("error reading redistribute: %v", err)
		}
	}

	if o.RouterId != nil {
		v := *o.RouterId

		if err = d.Set("router_id", v); err != nil {
			return diag.Errorf("error reading router_id: %v", err)
		}
	}

	if o.SpfTimers != nil {
		v := *o.SpfTimers

		if err = d.Set("spf_timers", v); err != nil {
			return diag.Errorf("error reading spf_timers: %v", err)
		}
	}

	if o.SummaryAddress != nil {
		if err = d.Set("summary_address", flattenRouterOspf6SummaryAddress(o.SummaryAddress, sort)); err != nil {
			return diag.Errorf("error reading summary_address: %v", err)
		}
	}

	return nil
}

func expandRouterOspf6Area(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspf6Area, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspf6Area

	for i := range l {
		tmp := models.RouterOspf6Area{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.authentication", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Authentication = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.default_cost", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.DefaultCost = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipsec_auth_alg", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IpsecAuthAlg = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipsec_enc_alg", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IpsecEncAlg = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipsec_keys", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterOspf6AreaIpsecKeys(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterOspf6AreaIpsecKeys
			// 	}
			tmp.IpsecKeys = v2

		}

		pre_append = fmt.Sprintf("%s.%d.key_rollover_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.KeyRolloverInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nssa_default_information_originate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NssaDefaultInformationOriginate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nssa_default_information_originate_metric", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.NssaDefaultInformationOriginateMetric = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nssa_default_information_originate_metric_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NssaDefaultInformationOriginateMetricType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nssa_redistribution", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NssaRedistribution = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nssa_translator_role", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NssaTranslatorRole = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.range", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterOspf6AreaRange(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterOspf6AreaRange
			// 	}
			tmp.Range = v2

		}

		pre_append = fmt.Sprintf("%s.%d.stub_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StubType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.virtual_link", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterOspf6AreaVirtualLink(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterOspf6AreaVirtualLink
			// 	}
			tmp.VirtualLink = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterOspf6AreaIpsecKeys(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspf6AreaIpsecKeys, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspf6AreaIpsecKeys

	for i := range l {
		tmp := models.RouterOspf6AreaIpsecKeys{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.auth_key", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthKey = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.enc_key", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EncKey = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.spi", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Spi = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterOspf6AreaRange(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspf6AreaRange, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspf6AreaRange

	for i := range l {
		tmp := models.RouterOspf6AreaRange{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.advertise", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Advertise = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
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

func expandRouterOspf6AreaVirtualLink(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspf6AreaVirtualLink, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspf6AreaVirtualLink

	for i := range l {
		tmp := models.RouterOspf6AreaVirtualLink{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.authentication", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Authentication = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dead_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.DeadInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.hello_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.HelloInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipsec_auth_alg", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IpsecAuthAlg = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipsec_enc_alg", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IpsecEncAlg = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipsec_keys", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterOspf6AreaVirtualLinkIpsecKeys(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterOspf6AreaVirtualLinkIpsecKeys
			// 	}
			tmp.IpsecKeys = v2

		}

		pre_append = fmt.Sprintf("%s.%d.key_rollover_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.KeyRolloverInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.peer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Peer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.retransmit_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.RetransmitInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.transmit_delay", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.TransmitDelay = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterOspf6AreaVirtualLinkIpsecKeys(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspf6AreaVirtualLinkIpsecKeys, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspf6AreaVirtualLinkIpsecKeys

	for i := range l {
		tmp := models.RouterOspf6AreaVirtualLinkIpsecKeys{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.auth_key", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthKey = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.enc_key", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EncKey = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.spi", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Spi = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterOspf6Ospf6Interface(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspf6Ospf6Interface, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspf6Ospf6Interface

	for i := range l {
		tmp := models.RouterOspf6Ospf6Interface{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.area_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AreaId = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.authentication", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Authentication = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bfd", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Bfd = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cost", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Cost = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dead_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.DeadInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.hello_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.HelloInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipsec_auth_alg", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IpsecAuthAlg = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipsec_enc_alg", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IpsecEncAlg = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipsec_keys", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterOspf6Ospf6InterfaceIpsecKeys(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterOspf6Ospf6InterfaceIpsecKeys
			// 	}
			tmp.IpsecKeys = v2

		}

		pre_append = fmt.Sprintf("%s.%d.key_rollover_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.KeyRolloverInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mtu", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Mtu = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mtu_ignore", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MtuIgnore = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.neighbor", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterOspf6Ospf6InterfaceNeighbor(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterOspf6Ospf6InterfaceNeighbor
			// 	}
			tmp.Neighbor = v2

		}

		pre_append = fmt.Sprintf("%s.%d.network_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NetworkType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Priority = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.retransmit_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.RetransmitInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.transmit_delay", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.TransmitDelay = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterOspf6Ospf6InterfaceIpsecKeys(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspf6Ospf6InterfaceIpsecKeys, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspf6Ospf6InterfaceIpsecKeys

	for i := range l {
		tmp := models.RouterOspf6Ospf6InterfaceIpsecKeys{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.auth_key", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthKey = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.enc_key", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EncKey = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.spi", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Spi = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterOspf6Ospf6InterfaceNeighbor(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspf6Ospf6InterfaceNeighbor, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspf6Ospf6InterfaceNeighbor

	for i := range l {
		tmp := models.RouterOspf6Ospf6InterfaceNeighbor{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.cost", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Cost = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.poll_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.PollInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Priority = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterOspf6PassiveInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspf6PassiveInterface, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspf6PassiveInterface

	for i := range l {
		tmp := models.RouterOspf6PassiveInterface{}
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

func expandRouterOspf6Redistribute(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspf6Redistribute, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspf6Redistribute

	for i := range l {
		tmp := models.RouterOspf6Redistribute{}
		var pre_append string

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

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
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

func expandRouterOspf6SummaryAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspf6SummaryAddress, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspf6SummaryAddress

	for i := range l {
		tmp := models.RouterOspf6SummaryAddress{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.advertise", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Advertise = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Prefix6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Tag = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectRouterOspf6(d *schema.ResourceData, sv string) (*models.RouterOspf6, diag.Diagnostics) {
	obj := models.RouterOspf6{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("abr_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("abr_type", sv)
				diags = append(diags, e)
			}
			obj.AbrType = &v2
		}
	}
	if v, ok := d.GetOk("area"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("area", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterOspf6Area(d, v, "area", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Area = t
		}
	} else if d.HasChange("area") {
		old, new := d.GetChange("area")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Area = &[]models.RouterOspf6Area{}
		}
	}
	if v1, ok := d.GetOk("auto_cost_ref_bandwidth"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_cost_ref_bandwidth", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AutoCostRefBandwidth = &tmp
		}
	}
	if v1, ok := d.GetOk("bfd"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bfd", sv)
				diags = append(diags, e)
			}
			obj.Bfd = &v2
		}
	}
	if v1, ok := d.GetOk("default_information_metric"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_information_metric", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DefaultInformationMetric = &tmp
		}
	}
	if v1, ok := d.GetOk("default_information_metric_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_information_metric_type", sv)
				diags = append(diags, e)
			}
			obj.DefaultInformationMetricType = &v2
		}
	}
	if v1, ok := d.GetOk("default_information_originate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_information_originate", sv)
				diags = append(diags, e)
			}
			obj.DefaultInformationOriginate = &v2
		}
	}
	if v1, ok := d.GetOk("default_information_route_map"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_information_route_map", sv)
				diags = append(diags, e)
			}
			obj.DefaultInformationRouteMap = &v2
		}
	}
	if v1, ok := d.GetOk("default_metric"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_metric", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DefaultMetric = &tmp
		}
	}
	if v1, ok := d.GetOk("log_neighbour_changes"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("log_neighbour_changes", sv)
				diags = append(diags, e)
			}
			obj.LogNeighbourChanges = &v2
		}
	}
	if v, ok := d.GetOk("ospf6_interface"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ospf6_interface", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterOspf6Ospf6Interface(d, v, "ospf6_interface", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Ospf6Interface = t
		}
	} else if d.HasChange("ospf6_interface") {
		old, new := d.GetChange("ospf6_interface")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Ospf6Interface = &[]models.RouterOspf6Ospf6Interface{}
		}
	}
	if v, ok := d.GetOk("passive_interface"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("passive_interface", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterOspf6PassiveInterface(d, v, "passive_interface", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.PassiveInterface = t
		}
	} else if d.HasChange("passive_interface") {
		old, new := d.GetChange("passive_interface")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.PassiveInterface = &[]models.RouterOspf6PassiveInterface{}
		}
	}
	if v, ok := d.GetOk("redistribute"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("redistribute", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterOspf6Redistribute(d, v, "redistribute", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Redistribute = t
		}
	} else if d.HasChange("redistribute") {
		old, new := d.GetChange("redistribute")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Redistribute = &[]models.RouterOspf6Redistribute{}
		}
	}
	if v1, ok := d.GetOk("router_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("router_id", sv)
				diags = append(diags, e)
			}
			obj.RouterId = &v2
		}
	}
	if v1, ok := d.GetOk("spf_timers"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("spf_timers", sv)
				diags = append(diags, e)
			}
			obj.SpfTimers = &v2
		}
	}
	if v, ok := d.GetOk("summary_address"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("summary_address", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterOspf6SummaryAddress(d, v, "summary_address", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SummaryAddress = t
		}
	} else if d.HasChange("summary_address") {
		old, new := d.GetChange("summary_address")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SummaryAddress = &[]models.RouterOspf6SummaryAddress{}
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectRouterOspf6(d *schema.ResourceData, sv string) (*models.RouterOspf6, diag.Diagnostics) {
	obj := models.RouterOspf6{}
	diags := diag.Diagnostics{}

	obj.Area = &[]models.RouterOspf6Area{}
	obj.Ospf6Interface = &[]models.RouterOspf6Ospf6Interface{}
	obj.PassiveInterface = &[]models.RouterOspf6PassiveInterface{}
	obj.Redistribute = &[]models.RouterOspf6Redistribute{}
	obj.SummaryAddress = &[]models.RouterOspf6SummaryAddress{}

	return &obj, diags
}
