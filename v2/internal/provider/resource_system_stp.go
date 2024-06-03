// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.6,v7.0.2,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Spanning Tree Protocol (STP).

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

func resourceSystemStp() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Spanning Tree Protocol (STP).",

		CreateContext: resourceSystemStpCreate,
		ReadContext:   resourceSystemStpRead,
		UpdateContext: resourceSystemStpUpdate,
		DeleteContext: resourceSystemStpDelete,

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
			"forward_delay": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(4, 30),

				Description: "Forward delay (4 - 30 sec, default = 15).",
				Optional:    true,
				Computed:    true,
			},
			"hello_time": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),

				Description: "Hello time (1 - 10 sec, default = 2).",
				Optional:    true,
				Computed:    true,
			},
			"max_age": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(6, 40),

				Description: "Maximum packet age (6 - 40 sec, default = 20).",
				Optional:    true,
				Computed:    true,
			},
			"max_hops": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 40),

				Description: "Maximum number of hops (1 - 40, default = 20).",
				Optional:    true,
				Computed:    true,
			},
			"switch_priority": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"0", "4096", "8192", "12288", "16384", "20480", "24576", "28672", "32768", "36864", "40960", "45056", "49152", "53248", "57344"}, false),

				Description: "STP switch priority; the lower the number the higher the priority (select from 0, 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, and 57344).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemStpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemStp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemStp(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemStp")
	}

	return resourceSystemStpRead(ctx, d, meta)
}

func resourceSystemStpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemStp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemStp(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemStp resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemStp")
	}

	return resourceSystemStpRead(ctx, d, meta)
}

func resourceSystemStpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemStp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemStp(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemStp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemStpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemStp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemStp resource: %v", err)
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

	diags := refreshObjectSystemStp(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemStp(d *schema.ResourceData, o *models.SystemStp, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.ForwardDelay != nil {
		v := *o.ForwardDelay

		if err = d.Set("forward_delay", v); err != nil {
			return diag.Errorf("error reading forward_delay: %v", err)
		}
	}

	if o.HelloTime != nil {
		v := *o.HelloTime

		if err = d.Set("hello_time", v); err != nil {
			return diag.Errorf("error reading hello_time: %v", err)
		}
	}

	if o.MaxAge != nil {
		v := *o.MaxAge

		if err = d.Set("max_age", v); err != nil {
			return diag.Errorf("error reading max_age: %v", err)
		}
	}

	if o.MaxHops != nil {
		v := *o.MaxHops

		if err = d.Set("max_hops", v); err != nil {
			return diag.Errorf("error reading max_hops: %v", err)
		}
	}

	if o.SwitchPriority != nil {
		v := *o.SwitchPriority

		if err = d.Set("switch_priority", v); err != nil {
			return diag.Errorf("error reading switch_priority: %v", err)
		}
	}

	return nil
}

func getObjectSystemStp(d *schema.ResourceData, sv string) (*models.SystemStp, diag.Diagnostics) {
	obj := models.SystemStp{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("forward_delay"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("forward_delay", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ForwardDelay = &tmp
		}
	}
	if v1, ok := d.GetOk("hello_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hello_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HelloTime = &tmp
		}
	}
	if v1, ok := d.GetOk("list"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("list", sv)
				diags = append(diags, e)
			}
			obj.List = &v2
		}
	}
	if v1, ok := d.GetOk("max_age"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_age", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxAge = &tmp
		}
	}
	if v1, ok := d.GetOk("max_hops"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_hops", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxHops = &tmp
		}
	}
	if v1, ok := d.GetOk("switch_priority"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("switch_priority", sv)
				diags = append(diags, e)
			}
			obj.SwitchPriority = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemStp(d *schema.ResourceData, sv string) (*models.SystemStp, diag.Diagnostics) {
	obj := models.SystemStp{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
