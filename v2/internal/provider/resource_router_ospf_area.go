// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: OSPF area configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceRouterOspfArea() *schema.Resource {
	return &schema.Resource{
		Description: "OSPF area configuration.",

		CreateContext: resourceRouterOspfAreaCreate,
		ReadContext:   resourceRouterOspfAreaRead,
		UpdateContext: resourceRouterOspfAreaUpdate,
		DeleteContext: resourceRouterOspfAreaDelete,

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
				ValidateFunc: validation.StringInSlice([]string{"none", "text", "message-digest"}, false),

				Description: "Authentication type.",
				Optional:    true,
				Computed:    true,
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"default_cost": {
				Type: schema.TypeInt,

				Description: "Summary default cost of stub or NSSA area.",
				Optional:    true,
				Computed:    true,
			},
			"filter_list": {
				Type:        schema.TypeList,
				Description: "OSPF area filter-list configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"direction": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),

							Description: "Direction.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Filter list entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"list": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Access-list or prefix-list name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"fosid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Area entry IP address.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"nssa_default_information_originate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "always", "disable"}, false),

				Description: "Redistribute, advertise, or do not originate Type-7 default route into NSSA area.",
				Optional:    true,
				Computed:    true,
			},
			"nssa_default_information_originate_metric": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16777214),

				Description: "OSPF default metric.",
				Optional:    true,
				Computed:    true,
			},
			"nssa_default_information_originate_metric_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"1", "2"}, false),

				Description: "OSPF metric type for default routes.",
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
				Description: "OSPF area range configuration.",
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
						"prefix": {
							Type:         schema.TypeString,
							ValidateFunc: validators.FortiValidateIPv4ClassnetAny,

							Description: "Prefix.",
							Optional:    true,
							Computed:    true,
						},
						"substitute": {
							Type:         schema.TypeString,
							ValidateFunc: validators.FortiValidateIPv4ClassnetAny,

							Description: "Substitute prefix.",
							Optional:    true,
							Computed:    true,
						},
						"substitute_status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable substitute status.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"shortcut": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable", "default"}, false),

				Description: "Enable/disable shortcut option.",
				Optional:    true,
				Computed:    true,
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
				Description: "OSPF virtual link configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "text", "message-digest"}, false),

							Description: "Authentication type.",
							Optional:    true,
							Computed:    true,
						},
						"authentication_key": {
							Type: schema.TypeString,

							Description: "Authentication key.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
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
						"keychain": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Message-digest key-chain name.",
							Optional:    true,
							Computed:    true,
						},
						"md5_keychain": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Authentication MD5 key-chain name.",
							Optional:    true,
							Computed:    true,
						},
						"md5_keys": {
							Type:        schema.TypeList,
							Description: "MD5 key.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 255),

										Description: "Key ID (1 - 255).",
										Optional:    true,
										Computed:    true,
									},
									"key_string": {
										Type: schema.TypeString,

										Description: "Password for the key.",
										Optional:    true,
										Computed:    true,
										Sensitive:   true,
									},
								},
							},
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

							Description: "Peer IP.",
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

func resourceRouterOspfAreaCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating RouterOspfArea resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterOspfArea(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterOspfArea(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterOspfArea")
	}

	return resourceRouterOspfAreaRead(ctx, d, meta)
}

func resourceRouterOspfAreaUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterOspfArea(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterOspfArea(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterOspfArea resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterOspfArea")
	}

	return resourceRouterOspfAreaRead(ctx, d, meta)
}

func resourceRouterOspfAreaDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterOspfArea(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterOspfArea resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspfAreaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterOspfArea(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterOspfArea resource: %v", err)
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

	diags := refreshObjectRouterOspfArea(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectRouterOspfArea(d *schema.ResourceData, o *models.RouterOspfArea, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Authentication != nil {
		v := *o.Authentication

		if err = d.Set("authentication", v); err != nil {
			return diag.Errorf("error reading authentication: %v", err)
		}
	}

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.DefaultCost != nil {
		v := *o.DefaultCost

		if err = d.Set("default_cost", v); err != nil {
			return diag.Errorf("error reading default_cost: %v", err)
		}
	}

	if o.FilterList != nil {
		if err = d.Set("filter_list", flattenRouterOspfAreaFilterList(o.FilterList, sort)); err != nil {
			return diag.Errorf("error reading filter_list: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
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
		if err = d.Set("range", flattenRouterOspfAreaRange(o.Range, sort)); err != nil {
			return diag.Errorf("error reading range: %v", err)
		}
	}

	if o.Shortcut != nil {
		v := *o.Shortcut

		if err = d.Set("shortcut", v); err != nil {
			return diag.Errorf("error reading shortcut: %v", err)
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
		if err = d.Set("virtual_link", flattenRouterOspfAreaVirtualLink(o.VirtualLink, sort)); err != nil {
			return diag.Errorf("error reading virtual_link: %v", err)
		}
	}

	return nil
}

func getObjectRouterOspfArea(d *schema.ResourceData, sv string) (*models.RouterOspfArea, diag.Diagnostics) {
	obj := models.RouterOspfArea{}
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
	if v1, ok := d.GetOk("comments"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("comments", sv)
				diags = append(diags, e)
			}
			obj.Comments = &v2
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
	if v, ok := d.GetOk("filter_list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("filter_list", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterOspfAreaFilterList(d, v, "filter_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.FilterList = t
		}
	} else if d.HasChange("filter_list") {
		old, new := d.GetChange("filter_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.FilterList = &[]models.RouterOspfAreaFilterList{}
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
		t, err := expandRouterOspfAreaRange(d, v, "range", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Range = t
		}
	} else if d.HasChange("range") {
		old, new := d.GetChange("range")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Range = &[]models.RouterOspfAreaRange{}
		}
	}
	if v1, ok := d.GetOk("shortcut"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("shortcut", sv)
				diags = append(diags, e)
			}
			obj.Shortcut = &v2
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
		t, err := expandRouterOspfAreaVirtualLink(d, v, "virtual_link", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.VirtualLink = t
		}
	} else if d.HasChange("virtual_link") {
		old, new := d.GetChange("virtual_link")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.VirtualLink = &[]models.RouterOspfAreaVirtualLink{}
		}
	}
	return &obj, diags
}
