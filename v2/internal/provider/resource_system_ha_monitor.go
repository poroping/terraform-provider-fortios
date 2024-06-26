// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure HA monitor.

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

func resourceSystemHaMonitor() *schema.Resource {
	return &schema.Resource{
		Description: "Configure HA monitor.",

		CreateContext: resourceSystemHaMonitorCreate,
		ReadContext:   resourceSystemHaMonitorRead,
		UpdateContext: resourceSystemHaMonitorUpdate,
		DeleteContext: resourceSystemHaMonitorDelete,

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
			"monitor_vlan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable monitor VLAN interfaces.",
				Optional:    true,
				Computed:    true,
			},
			"vlan_hb_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),

				Description: "Configure heartbeat interval (seconds).",
				Optional:    true,
				Computed:    true,
			},
			"vlan_hb_lost_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 60),

				Description: "VLAN lost heartbeat threshold (1 - 60).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemHaMonitorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemHaMonitor(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemHaMonitor(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemHaMonitor")
	}

	return resourceSystemHaMonitorRead(ctx, d, meta)
}

func resourceSystemHaMonitorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemHaMonitor(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemHaMonitor(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemHaMonitor resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemHaMonitor")
	}

	return resourceSystemHaMonitorRead(ctx, d, meta)
}

func resourceSystemHaMonitorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemHaMonitor(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemHaMonitor(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemHaMonitor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemHaMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemHaMonitor(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemHaMonitor resource: %v", err)
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

	diags := refreshObjectSystemHaMonitor(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemHaMonitor(d *schema.ResourceData, o *models.SystemHaMonitor, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.MonitorVlan != nil {
		v := *o.MonitorVlan

		if err = d.Set("monitor_vlan", v); err != nil {
			return diag.Errorf("error reading monitor_vlan: %v", err)
		}
	}

	if o.VlanHbInterval != nil {
		v := *o.VlanHbInterval

		if err = d.Set("vlan_hb_interval", v); err != nil {
			return diag.Errorf("error reading vlan_hb_interval: %v", err)
		}
	}

	if o.VlanHbLostThreshold != nil {
		v := *o.VlanHbLostThreshold

		if err = d.Set("vlan_hb_lost_threshold", v); err != nil {
			return diag.Errorf("error reading vlan_hb_lost_threshold: %v", err)
		}
	}

	return nil
}

func getObjectSystemHaMonitor(d *schema.ResourceData, sv string) (*models.SystemHaMonitor, diag.Diagnostics) {
	obj := models.SystemHaMonitor{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("monitor_vlan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("monitor_vlan", sv)
				diags = append(diags, e)
			}
			obj.MonitorVlan = &v2
		}
	}
	if v1, ok := d.GetOk("vlan_hb_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vlan_hb_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.VlanHbInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("vlan_hb_lost_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vlan_hb_lost_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.VlanHbLostThreshold = &tmp
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemHaMonitor(d *schema.ResourceData, sv string) (*models.SystemHaMonitor, diag.Diagnostics) {
	obj := models.SystemHaMonitor{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
