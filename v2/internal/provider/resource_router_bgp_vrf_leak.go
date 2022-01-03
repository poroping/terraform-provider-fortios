// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
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

func resourceRouterBgpVrfLeak() *schema.Resource {
	return &schema.Resource{
		Description: "BGP VRF leaking table.",

		CreateContext: resourceRouterBgpVrfLeakCreate,
		ReadContext:   resourceRouterBgpVrfLeakRead,
		UpdateContext: resourceRouterBgpVrfLeakUpdate,
		DeleteContext: resourceRouterBgpVrfLeakDelete,

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
				Description: "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"target": {
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

							Description: "Target VRF ID <0 - 31>.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"vrf": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),

				Description: "Origin VRF ID <0 - 31>.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceRouterBgpVrfLeakCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating RouterBgpVrfLeak resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterBgpVrfLeak(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterBgpVrfLeak(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterBgpVrfLeak")
	}

	return resourceRouterBgpVrfLeakRead(ctx, d, meta)
}

func resourceRouterBgpVrfLeakUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterBgpVrfLeak(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterBgpVrfLeak(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterBgpVrfLeak resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterBgpVrfLeak")
	}

	return resourceRouterBgpVrfLeakRead(ctx, d, meta)
}

func resourceRouterBgpVrfLeakDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterBgpVrfLeak(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterBgpVrfLeak resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBgpVrfLeakRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterBgpVrfLeak(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterBgpVrfLeak resource: %v", err)
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

	diags := refreshObjectRouterBgpVrfLeak(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectRouterBgpVrfLeak(d *schema.ResourceData, o *models.RouterBgpVrfLeak, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Target != nil {
		if err = d.Set("target", flattenRouterBgpVrfLeakTarget(o.Target, sort)); err != nil {
			return diag.Errorf("error reading target: %v", err)
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

func getObjectRouterBgpVrfLeak(d *schema.ResourceData, sv string) (*models.RouterBgpVrfLeak, diag.Diagnostics) {
	obj := models.RouterBgpVrfLeak{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("target"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("target", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpVrfLeakTarget(d, v, "target", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Target = t
		}
	} else if d.HasChange("target") {
		old, new := d.GetChange("target")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Target = &[]models.RouterBgpVrfLeakTarget{}
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
