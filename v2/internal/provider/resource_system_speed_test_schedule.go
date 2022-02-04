// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Speed test schedule for each interface.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceSystemSpeedTestSchedule() *schema.Resource {
	return &schema.Resource{
		Description: "Speed test schedule for each interface.",

		CreateContext: resourceSystemSpeedTestScheduleCreate,
		ReadContext:   resourceSystemSpeedTestScheduleRead,
		UpdateContext: resourceSystemSpeedTestScheduleUpdate,
		DeleteContext: resourceSystemSpeedTestScheduleDelete,

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
			"diffserv": {
				Type: schema.TypeString,

				Description: "DSCP used for speed test.",
				Optional:    true,
				Computed:    true,
			},
			"dynamic_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable dynamic server option.",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Interface name.",
				ForceNew:    true,
				Required:    true,
			},
			"schedules": {
				Type:        schema.TypeList,
				Description: "Schedules for the interface.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),

							Description: "Name of a firewall recurring schedule.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"server_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Speed test server name.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable scheduled speed test.",
				Optional:    true,
				Computed:    true,
			},
			"update_inbandwidth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable bypassing interface's inbound bandwidth setting.",
				Optional:    true,
				Computed:    true,
			},
			"update_inbandwidth_maximum": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),

				Description: "Maximum downloading bandwidth (kbps) to be used in a speed test.",
				Optional:    true,
				Computed:    true,
			},
			"update_inbandwidth_minimum": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),

				Description: "Minimum downloading bandwidth (kbps) to be considered effective.",
				Optional:    true,
				Computed:    true,
			},
			"update_outbandwidth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable bypassing interface's outbound bandwidth setting.",
				Optional:    true,
				Computed:    true,
			},
			"update_outbandwidth_maximum": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),

				Description: "Maximum uploading bandwidth (kbps) to be used in a speed test.",
				Optional:    true,
				Computed:    true,
			},
			"update_outbandwidth_minimum": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),

				Description: "Minimum uploading bandwidth (kbps) to be considered effective.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemSpeedTestScheduleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "interface"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating SystemSpeedTestSchedule resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemSpeedTestSchedule(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemSpeedTestSchedule(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSpeedTestSchedule")
	}

	return resourceSystemSpeedTestScheduleRead(ctx, d, meta)
}

func resourceSystemSpeedTestScheduleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemSpeedTestSchedule(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemSpeedTestSchedule(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemSpeedTestSchedule resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSpeedTestSchedule")
	}

	return resourceSystemSpeedTestScheduleRead(ctx, d, meta)
}

func resourceSystemSpeedTestScheduleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemSpeedTestSchedule(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemSpeedTestSchedule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSpeedTestScheduleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemSpeedTestSchedule(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSpeedTestSchedule resource: %v", err)
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

	diags := refreshObjectSystemSpeedTestSchedule(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemSpeedTestScheduleSchedules(v *[]models.SystemSpeedTestScheduleSchedules, sort bool) interface{} {
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

func refreshObjectSystemSpeedTestSchedule(d *schema.ResourceData, o *models.SystemSpeedTestSchedule, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Diffserv != nil {
		v := *o.Diffserv

		if err = d.Set("diffserv", v); err != nil {
			return diag.Errorf("error reading diffserv: %v", err)
		}
	}

	if o.DynamicServer != nil {
		v := *o.DynamicServer

		if err = d.Set("dynamic_server", v); err != nil {
			return diag.Errorf("error reading dynamic_server: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.Schedules != nil {
		if err = d.Set("schedules", flattenSystemSpeedTestScheduleSchedules(o.Schedules, sort)); err != nil {
			return diag.Errorf("error reading schedules: %v", err)
		}
	}

	if o.ServerName != nil {
		v := *o.ServerName

		if err = d.Set("server_name", v); err != nil {
			return diag.Errorf("error reading server_name: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.UpdateInbandwidth != nil {
		v := *o.UpdateInbandwidth

		if err = d.Set("update_inbandwidth", v); err != nil {
			return diag.Errorf("error reading update_inbandwidth: %v", err)
		}
	}

	if o.UpdateInbandwidthMaximum != nil {
		v := *o.UpdateInbandwidthMaximum

		if err = d.Set("update_inbandwidth_maximum", v); err != nil {
			return diag.Errorf("error reading update_inbandwidth_maximum: %v", err)
		}
	}

	if o.UpdateInbandwidthMinimum != nil {
		v := *o.UpdateInbandwidthMinimum

		if err = d.Set("update_inbandwidth_minimum", v); err != nil {
			return diag.Errorf("error reading update_inbandwidth_minimum: %v", err)
		}
	}

	if o.UpdateOutbandwidth != nil {
		v := *o.UpdateOutbandwidth

		if err = d.Set("update_outbandwidth", v); err != nil {
			return diag.Errorf("error reading update_outbandwidth: %v", err)
		}
	}

	if o.UpdateOutbandwidthMaximum != nil {
		v := *o.UpdateOutbandwidthMaximum

		if err = d.Set("update_outbandwidth_maximum", v); err != nil {
			return diag.Errorf("error reading update_outbandwidth_maximum: %v", err)
		}
	}

	if o.UpdateOutbandwidthMinimum != nil {
		v := *o.UpdateOutbandwidthMinimum

		if err = d.Set("update_outbandwidth_minimum", v); err != nil {
			return diag.Errorf("error reading update_outbandwidth_minimum: %v", err)
		}
	}

	return nil
}

func expandSystemSpeedTestScheduleSchedules(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSpeedTestScheduleSchedules, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSpeedTestScheduleSchedules

	for i := range l {
		tmp := models.SystemSpeedTestScheduleSchedules{}
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

func getObjectSystemSpeedTestSchedule(d *schema.ResourceData, sv string) (*models.SystemSpeedTestSchedule, diag.Diagnostics) {
	obj := models.SystemSpeedTestSchedule{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("diffserv"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("diffserv", sv)
				diags = append(diags, e)
			}
			obj.Diffserv = &v2
		}
	}
	if v1, ok := d.GetOk("dynamic_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("dynamic_server", sv)
				diags = append(diags, e)
			}
			obj.DynamicServer = &v2
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v, ok := d.GetOk("schedules"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("schedules", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSpeedTestScheduleSchedules(d, v, "schedules", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Schedules = t
		}
	} else if d.HasChange("schedules") {
		old, new := d.GetChange("schedules")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Schedules = &[]models.SystemSpeedTestScheduleSchedules{}
		}
	}
	if v1, ok := d.GetOk("server_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_name", sv)
				diags = append(diags, e)
			}
			obj.ServerName = &v2
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
	if v1, ok := d.GetOk("update_inbandwidth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("update_inbandwidth", sv)
				diags = append(diags, e)
			}
			obj.UpdateInbandwidth = &v2
		}
	}
	if v1, ok := d.GetOk("update_inbandwidth_maximum"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("update_inbandwidth_maximum", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.UpdateInbandwidthMaximum = &tmp
		}
	}
	if v1, ok := d.GetOk("update_inbandwidth_minimum"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("update_inbandwidth_minimum", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.UpdateInbandwidthMinimum = &tmp
		}
	}
	if v1, ok := d.GetOk("update_outbandwidth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("update_outbandwidth", sv)
				diags = append(diags, e)
			}
			obj.UpdateOutbandwidth = &v2
		}
	}
	if v1, ok := d.GetOk("update_outbandwidth_maximum"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("update_outbandwidth_maximum", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.UpdateOutbandwidthMaximum = &tmp
		}
	}
	if v1, ok := d.GetOk("update_outbandwidth_minimum"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("update_outbandwidth_minimum", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.UpdateOutbandwidthMinimum = &tmp
		}
	}
	return &obj, diags
}
