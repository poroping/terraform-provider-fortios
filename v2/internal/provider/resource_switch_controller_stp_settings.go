// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiSwitch spanning tree protocol (STP).

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

func resourceSwitchControllerStpSettings() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiSwitch spanning tree protocol (STP).",

		CreateContext: resourceSwitchControllerStpSettingsCreate,
		ReadContext:   resourceSwitchControllerStpSettingsRead,
		UpdateContext: resourceSwitchControllerStpSettingsUpdate,
		DeleteContext: resourceSwitchControllerStpSettingsDelete,

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
			"forward_time": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(4, 30),

				Description: "Period of time a port is in listening and learning state (4 - 30 sec, default = 15).",
				Optional:    true,
				Computed:    true,
			},
			"hello_time": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),

				Description: "Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).",
				Optional:    true,
				Computed:    true,
			},
			"max_age": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(6, 40),

				Description: "Maximum time before a bridge port expires its configuration BPDU information (6 - 40 sec, default = 20).",
				Optional:    true,
				Computed:    true,
			},
			"max_hops": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 40),

				Description: "Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "Name of global STP settings configuration.",
				Optional:    true,
				Computed:    true,
			},
			"pending_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 15),

				Description: "Pending time (1 - 15 sec, default = 4).",
				Optional:    true,
				Computed:    true,
			},
			"revision": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "STP revision number (0 - 65535).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSwitchControllerStpSettingsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerStpSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSwitchControllerStpSettings(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerStpSettings")
	}

	return resourceSwitchControllerStpSettingsRead(ctx, d, meta)
}

func resourceSwitchControllerStpSettingsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerStpSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSwitchControllerStpSettings(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SwitchControllerStpSettings resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerStpSettings")
	}

	return resourceSwitchControllerStpSettingsRead(ctx, d, meta)
}

func resourceSwitchControllerStpSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSwitchControllerStpSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSwitchControllerStpSettings(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SwitchControllerStpSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerStpSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerStpSettings(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerStpSettings resource: %v", err)
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

	diags := refreshObjectSwitchControllerStpSettings(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSwitchControllerStpSettings(d *schema.ResourceData, o *models.SwitchControllerStpSettings, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.ForwardTime != nil {
		v := *o.ForwardTime

		if err = d.Set("forward_time", v); err != nil {
			return diag.Errorf("error reading forward_time: %v", err)
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

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.PendingTimer != nil {
		v := *o.PendingTimer

		if err = d.Set("pending_timer", v); err != nil {
			return diag.Errorf("error reading pending_timer: %v", err)
		}
	}

	if o.Revision != nil {
		v := *o.Revision

		if err = d.Set("revision", v); err != nil {
			return diag.Errorf("error reading revision: %v", err)
		}
	}

	return nil
}

func getObjectSwitchControllerStpSettings(d *schema.ResourceData, sv string) (*models.SwitchControllerStpSettings, diag.Diagnostics) {
	obj := models.SwitchControllerStpSettings{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("forward_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("forward_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ForwardTime = &tmp
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
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v1, ok := d.GetOk("pending_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pending_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PendingTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("revision"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("revision", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Revision = &tmp
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSwitchControllerStpSettings(d *schema.ResourceData, sv string) (*models.SwitchControllerStpSettings, diag.Diagnostics) {
	obj := models.SwitchControllerStpSettings{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
