// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPS URL filter settings.

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

func resourceWebfilterIpsUrlfilterSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPS URL filter settings.",

		CreateContext: resourceWebfilterIpsUrlfilterSettingCreate,
		ReadContext:   resourceWebfilterIpsUrlfilterSettingRead,
		UpdateContext: resourceWebfilterIpsUrlfilterSettingUpdate,
		DeleteContext: resourceWebfilterIpsUrlfilterSettingDelete,

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
			"device": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Interface for this route.",
				Optional:    true,
				Computed:    true,
			},
			"distance": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Administrative distance (1 - 255) for this route.",
				Optional:    true,
				Computed:    true,
			},
			"gateway": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Gateway IP address for this route.",
				Optional:    true,
				Computed:    true,
			},
			"geo_filter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Filter based on geographical location. Route will NOT be installed if the resolved IP address belongs to the country in the filter.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWebfilterIpsUrlfilterSettingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWebfilterIpsUrlfilterSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWebfilterIpsUrlfilterSetting(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebfilterIpsUrlfilterSetting")
	}

	return resourceWebfilterIpsUrlfilterSettingRead(ctx, d, meta)
}

func resourceWebfilterIpsUrlfilterSettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWebfilterIpsUrlfilterSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWebfilterIpsUrlfilterSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WebfilterIpsUrlfilterSetting resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebfilterIpsUrlfilterSetting")
	}

	return resourceWebfilterIpsUrlfilterSettingRead(ctx, d, meta)
}

func resourceWebfilterIpsUrlfilterSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectWebfilterIpsUrlfilterSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateWebfilterIpsUrlfilterSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WebfilterIpsUrlfilterSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterIpsUrlfilterSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWebfilterIpsUrlfilterSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebfilterIpsUrlfilterSetting resource: %v", err)
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

	diags := refreshObjectWebfilterIpsUrlfilterSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectWebfilterIpsUrlfilterSetting(d *schema.ResourceData, o *models.WebfilterIpsUrlfilterSetting, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Device != nil {
		v := *o.Device

		if err = d.Set("device", v); err != nil {
			return diag.Errorf("error reading device: %v", err)
		}
	}

	if o.Distance != nil {
		v := *o.Distance

		if err = d.Set("distance", v); err != nil {
			return diag.Errorf("error reading distance: %v", err)
		}
	}

	if o.Gateway != nil {
		v := *o.Gateway

		if err = d.Set("gateway", v); err != nil {
			return diag.Errorf("error reading gateway: %v", err)
		}
	}

	if o.GeoFilter != nil {
		v := *o.GeoFilter

		if err = d.Set("geo_filter", v); err != nil {
			return diag.Errorf("error reading geo_filter: %v", err)
		}
	}

	return nil
}

func getObjectWebfilterIpsUrlfilterSetting(d *schema.ResourceData, sv string) (*models.WebfilterIpsUrlfilterSetting, diag.Diagnostics) {
	obj := models.WebfilterIpsUrlfilterSetting{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("device"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("device", sv)
				diags = append(diags, e)
			}
			obj.Device = &v2
		}
	}
	if v1, ok := d.GetOk("distance"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("distance", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Distance = &tmp
		}
	}
	if v1, ok := d.GetOk("gateway"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gateway", sv)
				diags = append(diags, e)
			}
			obj.Gateway = &v2
		}
	}
	if v1, ok := d.GetOk("geo_filter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("geo_filter", sv)
				diags = append(diags, e)
			}
			obj.GeoFilter = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectWebfilterIpsUrlfilterSetting(d *schema.ResourceData, sv string) (*models.WebfilterIpsUrlfilterSetting, diag.Diagnostics) {
	obj := models.WebfilterIpsUrlfilterSetting{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
