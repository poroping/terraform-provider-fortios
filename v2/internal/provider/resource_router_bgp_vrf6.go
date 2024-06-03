// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: BGP IPv6 VRF leaking table.

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

func resourceRouterBgpVrf6() *schema.Resource {
	return &schema.Resource{
		Description: "BGP IPv6 VRF leaking table.",

		CreateContext: resourceRouterBgpVrf6Create,
		ReadContext:   resourceRouterBgpVrf6Read,
		UpdateContext: resourceRouterBgpVrf6Update,
		DeleteContext: resourceRouterBgpVrf6Delete,

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

func resourceRouterBgpVrf6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating RouterBgpVrf6 resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterBgpVrf6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterBgpVrf6(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterBgpVrf6")
	}

	return resourceRouterBgpVrf6Read(ctx, d, meta)
}

func resourceRouterBgpVrf6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterBgpVrf6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterBgpVrf6(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterBgpVrf6 resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterBgpVrf6")
	}

	return resourceRouterBgpVrf6Read(ctx, d, meta)
}

func resourceRouterBgpVrf6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterBgpVrf6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterBgpVrf6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBgpVrf6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterBgpVrf6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterBgpVrf6 resource: %v", err)
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

	diags := refreshObjectRouterBgpVrf6(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectRouterBgpVrf6(d *schema.ResourceData, o *models.RouterBgpVrf6, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.LeakTarget != nil {
		if err = d.Set("leak_target", flattenRouterBgpVrf6LeakTarget(d, o.LeakTarget, "leak_target", sort)); err != nil {
			return diag.Errorf("error reading leak_target: %v", err)
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

func getObjectRouterBgpVrf6(d *schema.ResourceData, sv string) (*models.RouterBgpVrf6, diag.Diagnostics) {
	obj := models.RouterBgpVrf6{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("leak_target"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("leak_target", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpVrf6LeakTarget(d, v, "leak_target", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.LeakTarget = t
		}
	} else if d.HasChange("leak_target") {
		old, new := d.GetChange("leak_target")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.LeakTarget = &[]models.RouterBgpVrf6LeakTarget{}
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