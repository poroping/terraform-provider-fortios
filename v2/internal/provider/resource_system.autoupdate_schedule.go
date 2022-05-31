// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure update schedule.

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

func resourceSystemAutoupdateSchedule() *schema.Resource {
	return &schema.Resource{
		Description: "Configure update schedule.",

		CreateContext: resourceSystemAutoupdateScheduleCreate,
		ReadContext:   resourceSystemAutoupdateScheduleRead,
		UpdateContext: resourceSystemAutoupdateScheduleUpdate,
		DeleteContext: resourceSystemAutoupdateScheduleDelete,

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
			"day": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}, false),

				Description: "Update day.",
				Optional:    true,
				Computed:    true,
			},
			"frequency": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"every", "daily", "weekly", "automatic"}, false),

				Description: "Update frequency.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable scheduled updates.",
				Optional:    true,
				Computed:    true,
			},
			"time": {
				Type: schema.TypeString,

				Description: "Update time.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemAutoupdateScheduleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemAutoupdateSchedule(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemAutoupdateSchedule(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemAutoupdateSchedule")
	}

	return resourceSystemAutoupdateScheduleRead(ctx, d, meta)
}

func resourceSystemAutoupdateScheduleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemAutoupdateSchedule(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemAutoupdateSchedule(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemAutoupdateSchedule resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemAutoupdateSchedule")
	}

	return resourceSystemAutoupdateScheduleRead(ctx, d, meta)
}

func resourceSystemAutoupdateScheduleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemAutoupdateSchedule(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemAutoupdateSchedule(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemAutoupdateSchedule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutoupdateScheduleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemAutoupdateSchedule(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemAutoupdateSchedule resource: %v", err)
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

	diags := refreshObjectSystemAutoupdateSchedule(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemAutoupdateSchedule(d *schema.ResourceData, o *models.SystemAutoupdateSchedule, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Day != nil {
		v := *o.Day

		if err = d.Set("day", v); err != nil {
			return diag.Errorf("error reading day: %v", err)
		}
	}

	if o.Frequency != nil {
		v := *o.Frequency

		if err = d.Set("frequency", v); err != nil {
			return diag.Errorf("error reading frequency: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Time != nil {
		v := *o.Time

		if err = d.Set("time", v); err != nil {
			return diag.Errorf("error reading time: %v", err)
		}
	}

	return nil
}

func getObjectSystemAutoupdateSchedule(d *schema.ResourceData, sv string) (*models.SystemAutoupdateSchedule, diag.Diagnostics) {
	obj := models.SystemAutoupdateSchedule{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("day"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("day", sv)
				diags = append(diags, e)
			}
			obj.Day = &v2
		}
	}
	if v1, ok := d.GetOk("frequency"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("frequency", sv)
				diags = append(diags, e)
			}
			obj.Frequency = &v2
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
	if v1, ok := d.GetOk("time"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("time", sv)
				diags = append(diags, e)
			}
			obj.Time = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemAutoupdateSchedule(d *schema.ResourceData, sv string) (*models.SystemAutoupdateSchedule, diag.Diagnostics) {
	obj := models.SystemAutoupdateSchedule{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
