// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Settings for WebTrends.

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

func resourceLogWebtrendsSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Settings for WebTrends.",

		CreateContext: resourceLogWebtrendsSettingCreate,
		ReadContext:   resourceLogWebtrendsSettingRead,
		UpdateContext: resourceLogWebtrendsSettingUpdate,
		DeleteContext: resourceLogWebtrendsSettingDelete,

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
			"server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Address of the remote WebTrends server.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging to WebTrends.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceLogWebtrendsSettingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogWebtrendsSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateLogWebtrendsSetting(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogWebtrendsSetting")
	}

	return resourceLogWebtrendsSettingRead(ctx, d, meta)
}

func resourceLogWebtrendsSettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogWebtrendsSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateLogWebtrendsSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating LogWebtrendsSetting resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogWebtrendsSetting")
	}

	return resourceLogWebtrendsSettingRead(ctx, d, meta)
}

func resourceLogWebtrendsSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectLogWebtrendsSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateLogWebtrendsSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting LogWebtrendsSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogWebtrendsSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadLogWebtrendsSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogWebtrendsSetting resource: %v", err)
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

	diags := refreshObjectLogWebtrendsSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectLogWebtrendsSetting(d *schema.ResourceData, o *models.LogWebtrendsSetting, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Server != nil {
		v := *o.Server

		if err = d.Set("server", v); err != nil {
			return diag.Errorf("error reading server: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	return nil
}

func getObjectLogWebtrendsSetting(d *schema.ResourceData, sv string) (*models.LogWebtrendsSetting, diag.Diagnostics) {
	obj := models.LogWebtrendsSetting{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server", sv)
				diags = append(diags, e)
			}
			obj.Server = &v2
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
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectLogWebtrendsSetting(d *schema.ResourceData, sv string) (*models.LogWebtrendsSetting, diag.Diagnostics) {
	obj := models.LogWebtrendsSetting{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
