// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: OSPF6 area configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceRouterOspf6Area() *schema.Resource {
	return &schema.Resource{
		Description: "OSPF6 area configuration.",

		CreateContext: resourceRouterOspf6AreaCreate,
		ReadContext:   resourceRouterOspf6AreaRead,
		UpdateContext: resourceRouterOspf6AreaUpdate,
		DeleteContext: resourceRouterOspf6AreaDelete,

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
				Type:         schema.TypeBool,
				Description:  "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:     true,
				Default:      false,
				RequiredWith: []string{"fosid"},
			},
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
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
			"fosid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Area entry IP address.",
				ForceNew:    true,
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
	}
}

func resourceRouterOspf6AreaCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "id"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating RouterOspf6Area resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterOspf6Area(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterOspf6Area(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterOspf6Area")
	}

	return resourceRouterOspf6AreaRead(ctx, d, meta)
}

func resourceRouterOspf6AreaUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterOspf6Area(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterOspf6Area(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterOspf6Area resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterOspf6Area")
	}

	return resourceRouterOspf6AreaRead(ctx, d, meta)
}

func resourceRouterOspf6AreaDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterOspf6Area(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterOspf6Area resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspf6AreaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterOspf6Area(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterOspf6Area resource: %v", err)
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

	diags := refreshObjectRouterOspf6Area(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectRouterOspf6Area(d *schema.ResourceData, o *models.RouterOspf6Area, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Authentication != nil {
		v := *o.Authentication

		if err = d.Set("authentication", v); err != nil {
			return diag.Errorf("error reading authentication: %v", err)
		}
	}

	if o.DefaultCost != nil {
		v := *o.DefaultCost

		if err = d.Set("default_cost", v); err != nil {
			return diag.Errorf("error reading default_cost: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.IpsecAuthAlg != nil {
		v := *o.IpsecAuthAlg

		if err = d.Set("ipsec_auth_alg", v); err != nil {
			return diag.Errorf("error reading ipsec_auth_alg: %v", err)
		}
	}

	if o.IpsecEncAlg != nil {
		v := *o.IpsecEncAlg

		if err = d.Set("ipsec_enc_alg", v); err != nil {
			return diag.Errorf("error reading ipsec_enc_alg: %v", err)
		}
	}

	if o.IpsecKeys != nil {
		if err = d.Set("ipsec_keys", flattenRouterOspf6AreaIpsecKeys(o.IpsecKeys, sort)); err != nil {
			return diag.Errorf("error reading ipsec_keys: %v", err)
		}
	}

	if o.KeyRolloverInterval != nil {
		v := *o.KeyRolloverInterval

		if err = d.Set("key_rollover_interval", v); err != nil {
			return diag.Errorf("error reading key_rollover_interval: %v", err)
		}
	}

	if o.NssaDefaultInformationOriginate != nil {
		v := *o.NssaDefaultInformationOriginate

		if err = d.Set("nssa_default_information_originate", v); err != nil {
			return diag.Errorf("error reading nssa_default_information_originate: %v", err)
		}
	}

	if o.NssaDefaultInformationOriginateMetric != nil {
		v := *o.NssaDefaultInformationOriginateMetric

		if err = d.Set("nssa_default_information_originate_metric", v); err != nil {
			return diag.Errorf("error reading nssa_default_information_originate_metric: %v", err)
		}
	}

	if o.NssaDefaultInformationOriginateMetricType != nil {
		v := *o.NssaDefaultInformationOriginateMetricType

		if err = d.Set("nssa_default_information_originate_metric_type", v); err != nil {
			return diag.Errorf("error reading nssa_default_information_originate_metric_type: %v", err)
		}
	}

	if o.NssaRedistribution != nil {
		v := *o.NssaRedistribution

		if err = d.Set("nssa_redistribution", v); err != nil {
			return diag.Errorf("error reading nssa_redistribution: %v", err)
		}
	}

	if o.NssaTranslatorRole != nil {
		v := *o.NssaTranslatorRole

		if err = d.Set("nssa_translator_role", v); err != nil {
			return diag.Errorf("error reading nssa_translator_role: %v", err)
		}
	}

	if o.Range != nil {
		if err = d.Set("range", flattenRouterOspf6AreaRange(o.Range, sort)); err != nil {
			return diag.Errorf("error reading range: %v", err)
		}
	}

	if o.StubType != nil {
		v := *o.StubType

		if err = d.Set("stub_type", v); err != nil {
			return diag.Errorf("error reading stub_type: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	if o.VirtualLink != nil {
		if err = d.Set("virtual_link", flattenRouterOspf6AreaVirtualLink(o.VirtualLink, sort)); err != nil {
			return diag.Errorf("error reading virtual_link: %v", err)
		}
	}

	return nil
}

func getObjectRouterOspf6Area(d *schema.ResourceData, sv string) (*models.RouterOspf6Area, diag.Diagnostics) {
	obj := models.RouterOspf6Area{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("authentication"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authentication", sv)
				diags = append(diags, e)
			}
			obj.Authentication = &v2
		}
	}
	if v1, ok := d.GetOk("default_cost"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_cost", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DefaultCost = &tmp
		}
	}
	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			obj.Id = &v2
		}
	}
	if v1, ok := d.GetOk("ipsec_auth_alg"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipsec_auth_alg", sv)
				diags = append(diags, e)
			}
			obj.IpsecAuthAlg = &v2
		}
	}
	if v1, ok := d.GetOk("ipsec_enc_alg"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipsec_enc_alg", sv)
				diags = append(diags, e)
			}
			obj.IpsecEncAlg = &v2
		}
	}
	if v, ok := d.GetOk("ipsec_keys"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ipsec_keys", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterOspf6AreaIpsecKeys(d, v, "ipsec_keys", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.IpsecKeys = t
		}
	} else if d.HasChange("ipsec_keys") {
		old, new := d.GetChange("ipsec_keys")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.IpsecKeys = &[]models.RouterOspf6AreaIpsecKeys{}
		}
	}
	if v1, ok := d.GetOk("key_rollover_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("key_rollover_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.KeyRolloverInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("nssa_default_information_originate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("nssa_default_information_originate", sv)
				diags = append(diags, e)
			}
			obj.NssaDefaultInformationOriginate = &v2
		}
	}
	if v1, ok := d.GetOk("nssa_default_information_originate_metric"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("nssa_default_information_originate_metric", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.NssaDefaultInformationOriginateMetric = &tmp
		}
	}
	if v1, ok := d.GetOk("nssa_default_information_originate_metric_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("nssa_default_information_originate_metric_type", sv)
				diags = append(diags, e)
			}
			obj.NssaDefaultInformationOriginateMetricType = &v2
		}
	}
	if v1, ok := d.GetOk("nssa_redistribution"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("nssa_redistribution", sv)
				diags = append(diags, e)
			}
			obj.NssaRedistribution = &v2
		}
	}
	if v1, ok := d.GetOk("nssa_translator_role"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("nssa_translator_role", sv)
				diags = append(diags, e)
			}
			obj.NssaTranslatorRole = &v2
		}
	}
	if v, ok := d.GetOk("range"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("range", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterOspf6AreaRange(d, v, "range", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Range = t
		}
	} else if d.HasChange("range") {
		old, new := d.GetChange("range")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Range = &[]models.RouterOspf6AreaRange{}
		}
	}
	if v1, ok := d.GetOk("stub_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("stub_type", sv)
				diags = append(diags, e)
			}
			obj.StubType = &v2
		}
	}
	if v1, ok := d.GetOk("type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("type", sv)
				diags = append(diags, e)
			}
			obj.Type = &v2
		}
	}
	if v, ok := d.GetOk("virtual_link"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("virtual_link", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterOspf6AreaVirtualLink(d, v, "virtual_link", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.VirtualLink = t
		}
	} else if d.HasChange("virtual_link") {
		old, new := d.GetChange("virtual_link")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.VirtualLink = &[]models.RouterOspf6AreaVirtualLink{}
		}
	}
	return &obj, diags
}
