// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: SNMP system info configuration.

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

func resourceSystemSnmpSysinfo() *schema.Resource {
	return &schema.Resource{
		Description: "SNMP system info configuration.",

		CreateContext: resourceSystemSnmpSysinfoCreate,
		ReadContext:   resourceSystemSnmpSysinfoRead,
		UpdateContext: resourceSystemSnmpSysinfoUpdate,
		DeleteContext: resourceSystemSnmpSysinfoDelete,

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
			"contact_info": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Contact information.",
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "System description.",
				Optional:    true,
				Computed:    true,
			},
			"engine_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 54),

				Description: "Local SNMP engineID string (maximum 27 characters).",
				Optional:    true,
				Computed:    true,
			},
			"engine_id_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"text", "hex", "mac"}, false),

				Description: "Local SNMP engineID type (text/hex/mac).",
				Optional:    true,
				Computed:    true,
			},
			"location": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "System location.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SNMP.",
				Optional:    true,
				Computed:    true,
			},
			"trap_high_cpu_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),

				Description: "CPU usage when trap is sent.",
				Optional:    true,
				Computed:    true,
			},
			"trap_log_full_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),

				Description: "Log disk usage when trap is sent.",
				Optional:    true,
				Computed:    true,
			},
			"trap_low_memory_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),

				Description: "Memory usage when trap is sent.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemSnmpSysinfoCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemSnmpSysinfo(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemSnmpSysinfo(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSnmpSysinfo")
	}

	return resourceSystemSnmpSysinfoRead(ctx, d, meta)
}

func resourceSystemSnmpSysinfoUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemSnmpSysinfo(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemSnmpSysinfo(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemSnmpSysinfo resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSnmpSysinfo")
	}

	return resourceSystemSnmpSysinfoRead(ctx, d, meta)
}

func resourceSystemSnmpSysinfoDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemSnmpSysinfo(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemSnmpSysinfo(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemSnmpSysinfo resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpSysinfoRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemSnmpSysinfo(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSnmpSysinfo resource: %v", err)
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

	diags := refreshObjectSystemSnmpSysinfo(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemSnmpSysinfo(d *schema.ResourceData, o *models.SystemSnmpSysinfo, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.ContactInfo != nil {
		v := *o.ContactInfo

		if err = d.Set("contact_info", v); err != nil {
			return diag.Errorf("error reading contact_info: %v", err)
		}
	}

	if o.Description != nil {
		v := *o.Description

		if err = d.Set("description", v); err != nil {
			return diag.Errorf("error reading description: %v", err)
		}
	}

	if o.EngineId != nil {
		v := *o.EngineId

		if err = d.Set("engine_id", v); err != nil {
			return diag.Errorf("error reading engine_id: %v", err)
		}
	}

	if o.EngineIdType != nil {
		v := *o.EngineIdType

		if err = d.Set("engine_id_type", v); err != nil {
			return diag.Errorf("error reading engine_id_type: %v", err)
		}
	}

	if o.Location != nil {
		v := *o.Location

		if err = d.Set("location", v); err != nil {
			return diag.Errorf("error reading location: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.TrapHighCpuThreshold != nil {
		v := *o.TrapHighCpuThreshold

		if err = d.Set("trap_high_cpu_threshold", v); err != nil {
			return diag.Errorf("error reading trap_high_cpu_threshold: %v", err)
		}
	}

	if o.TrapLogFullThreshold != nil {
		v := *o.TrapLogFullThreshold

		if err = d.Set("trap_log_full_threshold", v); err != nil {
			return diag.Errorf("error reading trap_log_full_threshold: %v", err)
		}
	}

	if o.TrapLowMemoryThreshold != nil {
		v := *o.TrapLowMemoryThreshold

		if err = d.Set("trap_low_memory_threshold", v); err != nil {
			return diag.Errorf("error reading trap_low_memory_threshold: %v", err)
		}
	}

	return nil
}

func getObjectSystemSnmpSysinfo(d *schema.ResourceData, sv string) (*models.SystemSnmpSysinfo, diag.Diagnostics) {
	obj := models.SystemSnmpSysinfo{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("contact_info"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("contact_info", sv)
				diags = append(diags, e)
			}
			obj.ContactInfo = &v2
		}
	}
	if v1, ok := d.GetOk("description"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("description", sv)
				diags = append(diags, e)
			}
			obj.Description = &v2
		}
	}
	if v1, ok := d.GetOk("engine_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("engine_id", sv)
				diags = append(diags, e)
			}
			obj.EngineId = &v2
		}
	}
	if v1, ok := d.GetOk("engine_id_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("engine_id_type", sv)
				diags = append(diags, e)
			}
			obj.EngineIdType = &v2
		}
	}
	if v1, ok := d.GetOk("location"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("location", sv)
				diags = append(diags, e)
			}
			obj.Location = &v2
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
	if v1, ok := d.GetOk("trap_high_cpu_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trap_high_cpu_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TrapHighCpuThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("trap_log_full_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trap_log_full_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TrapLogFullThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("trap_low_memory_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trap_low_memory_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TrapLowMemoryThreshold = &tmp
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemSnmpSysinfo(d *schema.ResourceData, sv string) (*models.SystemSnmpSysinfo, diag.Diagnostics) {
	obj := models.SystemSnmpSysinfo{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
