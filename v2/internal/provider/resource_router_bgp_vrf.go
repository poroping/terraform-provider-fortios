// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: BGP VRF leaking table.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceRouterBgpVrf() *schema.Resource {
	return &schema.Resource{
		Description: "BGP VRF leaking table.",

		CreateContext: resourceRouterBgpVrfCreate,
		ReadContext:   resourceRouterBgpVrfRead,
		UpdateContext: resourceRouterBgpVrfUpdate,
		DeleteContext: resourceRouterBgpVrfDelete,

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
			"export_rt": {
				Type:        schema.TypeList,
				Description: "List of export route target.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"route_target": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Attribute: AA|AA:NN.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"import_route_map": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Import route map.",
				Optional:    true,
				Computed:    true,
			},
			"import_rt": {
				Type:        schema.TypeList,
				Description: "List of import route target.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"route_target": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Attribute: AA|AA:NN.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"leak_target": {
				Type:        schema.TypeList,
				Description: "Target VRF table.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Interface which is used to leak routes to target VRF.",
							Optional:    true,
							Computed:    true,
						},
						"route_map": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Route map of VRF leaking.",
							Optional:    true,
							Computed:    true,
						},
						"vrf": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 7),

							Description: "Target VRF ID (0 - 251).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"rd": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Route Distinguisher: AA|AA:NN.",
				Optional:    true,
				Computed:    true,
			},
			"role": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"standalone", "ce", "pe"}, false),

				Description: "VRF role.",
				Optional:    true,
				Computed:    true,
			},
			"vrf": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),

				Description: "Origin VRF ID (0 - 251).",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceRouterBgpVrfCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "vrf"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating RouterBgpVrf resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterBgpVrf(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterBgpVrf(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterBgpVrf")
	}

	return resourceRouterBgpVrfRead(ctx, d, meta)
}

func resourceRouterBgpVrfUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterBgpVrf(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterBgpVrf(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterBgpVrf resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterBgpVrf")
	}

	return resourceRouterBgpVrfRead(ctx, d, meta)
}

func resourceRouterBgpVrfDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterBgpVrf(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterBgpVrf resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBgpVrfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterBgpVrf(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterBgpVrf resource: %v", err)
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

	diags := refreshObjectRouterBgpVrf(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectRouterBgpVrf(d *schema.ResourceData, o *models.RouterBgpVrf, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.ExportRt != nil {
		if err = d.Set("export_rt", flattenRouterBgpVrfExportRt(d, o.ExportRt, "export_rt", sort)); err != nil {
			return diag.Errorf("error reading export_rt: %v", err)
		}
	}

	if o.ImportRouteMap != nil {
		v := *o.ImportRouteMap

		if err = d.Set("import_route_map", v); err != nil {
			return diag.Errorf("error reading import_route_map: %v", err)
		}
	}

	if o.ImportRt != nil {
		if err = d.Set("import_rt", flattenRouterBgpVrfImportRt(d, o.ImportRt, "import_rt", sort)); err != nil {
			return diag.Errorf("error reading import_rt: %v", err)
		}
	}

	if o.LeakTarget != nil {
		if err = d.Set("leak_target", flattenRouterBgpVrfLeakTarget(d, o.LeakTarget, "leak_target", sort)); err != nil {
			return diag.Errorf("error reading leak_target: %v", err)
		}
	}

	if o.Rd != nil {
		v := *o.Rd

		if err = d.Set("rd", v); err != nil {
			return diag.Errorf("error reading rd: %v", err)
		}
	}

	if o.Role != nil {
		v := *o.Role

		if err = d.Set("role", v); err != nil {
			return diag.Errorf("error reading role: %v", err)
		}
	}

	if o.Vrf != nil {
		v := *o.Vrf

		if err = d.Set("vrf", v); err != nil {
			return diag.Errorf("error reading vrf: %v", err)
		}
	}

	return nil
}

func getObjectRouterBgpVrf(d *schema.ResourceData, sv string) (*models.RouterBgpVrf, diag.Diagnostics) {
	obj := models.RouterBgpVrf{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("export_rt"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("export_rt", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpVrfExportRt(d, v, "export_rt", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ExportRt = t
		}
	} else if d.HasChange("export_rt") {
		old, new := d.GetChange("export_rt")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ExportRt = &[]models.RouterBgpVrfExportRt{}
		}
	}
	if v1, ok := d.GetOk("import_route_map"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("import_route_map", sv)
				diags = append(diags, e)
			}
			obj.ImportRouteMap = &v2
		}
	}
	if v, ok := d.GetOk("import_rt"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("import_rt", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpVrfImportRt(d, v, "import_rt", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ImportRt = t
		}
	} else if d.HasChange("import_rt") {
		old, new := d.GetChange("import_rt")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ImportRt = &[]models.RouterBgpVrfImportRt{}
		}
	}
	if v, ok := d.GetOk("leak_target"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("leak_target", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpVrfLeakTarget(d, v, "leak_target", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.LeakTarget = t
		}
	} else if d.HasChange("leak_target") {
		old, new := d.GetChange("leak_target")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.LeakTarget = &[]models.RouterBgpVrfLeakTarget{}
		}
	}
	if v1, ok := d.GetOk("rd"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rd", sv)
				diags = append(diags, e)
			}
			obj.Rd = &v2
		}
	}
	if v1, ok := d.GetOk("role"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("role", sv)
				diags = append(diags, e)
			}
			obj.Role = &v2
		}
	}
	if v1, ok := d.GetOk("vrf"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vrf", sv)
				diags = append(diags, e)
			}
			obj.Vrf = &v2
		}
	}
	return &obj, diags
}
