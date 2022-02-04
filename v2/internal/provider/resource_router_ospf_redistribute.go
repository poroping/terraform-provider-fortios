// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Redistribute configuration.

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

func resourceRouterOspfRedistribute() *schema.Resource {
	return &schema.Resource{
		Description: "Redistribute configuration.",

		CreateContext: resourceRouterOspfRedistributeCreate,
		ReadContext:   resourceRouterOspfRedistributeRead,
		UpdateContext: resourceRouterOspfRedistributeUpdate,
		DeleteContext: resourceRouterOspfRedistributeDelete,

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
				ForceNew:    true,
				Required:    true,
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
			"tag": {
				Type: schema.TypeInt,

				Description: "Tag value.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceRouterOspfRedistributeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "name"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating RouterOspfRedistribute resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterOspfRedistribute(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterOspfRedistribute(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterOspfRedistribute")
	}

	return resourceRouterOspfRedistributeRead(ctx, d, meta)
}

func resourceRouterOspfRedistributeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterOspfRedistribute(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterOspfRedistribute(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterOspfRedistribute resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterOspfRedistribute")
	}

	return resourceRouterOspfRedistributeRead(ctx, d, meta)
}

func resourceRouterOspfRedistributeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterOspfRedistribute(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterOspfRedistribute resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspfRedistributeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterOspfRedistribute(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterOspfRedistribute resource: %v", err)
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

	diags := refreshObjectRouterOspfRedistribute(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectRouterOspfRedistribute(d *schema.ResourceData, o *models.RouterOspfRedistribute, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Metric != nil {
		v := *o.Metric

		if err = d.Set("metric", v); err != nil {
			return diag.Errorf("error reading metric: %v", err)
		}
	}

	if o.MetricType != nil {
		v := *o.MetricType

		if err = d.Set("metric_type", v); err != nil {
			return diag.Errorf("error reading metric_type: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Routemap != nil {
		v := *o.Routemap

		if err = d.Set("routemap", v); err != nil {
			return diag.Errorf("error reading routemap: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Tag != nil {
		v := *o.Tag

		if err = d.Set("tag", v); err != nil {
			return diag.Errorf("error reading tag: %v", err)
		}
	}

	return nil
}

func getObjectRouterOspfRedistribute(d *schema.ResourceData, sv string) (*models.RouterOspfRedistribute, diag.Diagnostics) {
	obj := models.RouterOspfRedistribute{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("metric"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("metric", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Metric = &tmp
		}
	}
	if v1, ok := d.GetOk("metric_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("metric_type", sv)
				diags = append(diags, e)
			}
			obj.MetricType = &v2
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
	if v1, ok := d.GetOk("routemap"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("routemap", sv)
				diags = append(diags, e)
			}
			obj.Routemap = &v2
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
	if v1, ok := d.GetOk("tag"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tag", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Tag = &tmp
		}
	}
	return &obj, diags
}