// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure router settings.

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

func resourceRouterSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Configure router settings.",

		CreateContext: resourceRouterSettingCreate,
		ReadContext:   resourceRouterSettingRead,
		UpdateContext: resourceRouterSettingUpdate,
		DeleteContext: resourceRouterSettingDelete,

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
			"hostname": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 14),

				Description: "Hostname for this virtual domain router.",
				Optional:    true,
				Computed:    true,
			},
			"show_filter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Prefix-list as filter for showing routes.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceRouterSettingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterSetting(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterSetting")
	}

	return resourceRouterSettingRead(ctx, d, meta)
}

func resourceRouterSettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterSetting resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterSetting")
	}

	return resourceRouterSettingRead(ctx, d, meta)
}

func resourceRouterSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectRouterSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateRouterSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterSetting resource: %v", err)
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

	diags := refreshObjectRouterSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectRouterSetting(d *schema.ResourceData, o *models.RouterSetting, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Hostname != nil {
		v := *o.Hostname

		if err = d.Set("hostname", v); err != nil {
			return diag.Errorf("error reading hostname: %v", err)
		}
	}

	if o.ShowFilter != nil {
		v := *o.ShowFilter

		if err = d.Set("show_filter", v); err != nil {
			return diag.Errorf("error reading show_filter: %v", err)
		}
	}

	return nil
}

func getObjectRouterSetting(d *schema.ResourceData, sv string) (*models.RouterSetting, diag.Diagnostics) {
	obj := models.RouterSetting{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("hostname"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hostname", sv)
				diags = append(diags, e)
			}
			obj.Hostname = &v2
		}
	}
	if v1, ok := d.GetOk("show_filter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("show_filter", sv)
				diags = append(diags, e)
			}
			obj.ShowFilter = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectRouterSetting(d *schema.ResourceData, sv string) (*models.RouterSetting, diag.Diagnostics) {
	obj := models.RouterSetting{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
