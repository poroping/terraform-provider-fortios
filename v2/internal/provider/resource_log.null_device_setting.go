// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Settings for null device logging.

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

func resourceLogNullDeviceSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Settings for null device logging.",

		CreateContext: resourceLogNullDeviceSettingCreate,
		ReadContext:   resourceLogNullDeviceSettingRead,
		UpdateContext: resourceLogNullDeviceSettingUpdate,
		DeleteContext: resourceLogNullDeviceSettingDelete,

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
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable statistics collection for when no external logging destination, such as FortiAnalyzer, is present (data is not saved).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceLogNullDeviceSettingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogNullDeviceSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateLogNullDeviceSetting(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogNullDeviceSetting")
	}

	return resourceLogNullDeviceSettingRead(ctx, d, meta)
}

func resourceLogNullDeviceSettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogNullDeviceSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateLogNullDeviceSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating LogNullDeviceSetting resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogNullDeviceSetting")
	}

	return resourceLogNullDeviceSettingRead(ctx, d, meta)
}

func resourceLogNullDeviceSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectLogNullDeviceSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateLogNullDeviceSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting LogNullDeviceSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogNullDeviceSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadLogNullDeviceSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogNullDeviceSetting resource: %v", err)
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

	diags := refreshObjectLogNullDeviceSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectLogNullDeviceSetting(d *schema.ResourceData, o *models.LogNullDeviceSetting, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	return nil
}

func getObjectLogNullDeviceSetting(d *schema.ResourceData, sv string) (*models.LogNullDeviceSetting, diag.Diagnostics) {
	obj := models.LogNullDeviceSetting{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectLogNullDeviceSetting(d *schema.ResourceData, sv string) (*models.LogNullDeviceSetting, diag.Diagnostics) {
	obj := models.LogNullDeviceSetting{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
