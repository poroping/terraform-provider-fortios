// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Settings for TACACS+ accounting.

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

func resourceLogTacacsaccounting3Setting() *schema.Resource {
	return &schema.Resource{
		Description: "Settings for TACACS+ accounting.",

		CreateContext: resourceLogTacacsaccounting3SettingCreate,
		ReadContext:   resourceLogTacacsaccounting3SettingRead,
		UpdateContext: resourceLogTacacsaccounting3SettingUpdate,
		DeleteContext: resourceLogTacacsaccounting3SettingDelete,

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

				Description: "Address of TACACS+ server.",
				Optional:    true,
				Computed:    true,
			},
			"server_key": {
				Type: schema.TypeString,

				Description: "Key to access the TACACS+ server.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable TACACS+ accounting.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceLogTacacsaccounting3SettingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogTacacsaccounting3Setting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateLogTacacsaccounting3Setting(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogTacacsaccounting3Setting")
	}

	return resourceLogTacacsaccounting3SettingRead(ctx, d, meta)
}

func resourceLogTacacsaccounting3SettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogTacacsaccounting3Setting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateLogTacacsaccounting3Setting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating LogTacacsaccounting3Setting resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogTacacsaccounting3Setting")
	}

	return resourceLogTacacsaccounting3SettingRead(ctx, d, meta)
}

func resourceLogTacacsaccounting3SettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectLogTacacsaccounting3Setting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateLogTacacsaccounting3Setting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting LogTacacsaccounting3Setting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogTacacsaccounting3SettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadLogTacacsaccounting3Setting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogTacacsaccounting3Setting resource: %v", err)
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

	diags := refreshObjectLogTacacsaccounting3Setting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectLogTacacsaccounting3Setting(d *schema.ResourceData, o *models.LogTacacsaccounting3Setting, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Server != nil {
		v := *o.Server

		if err = d.Set("server", v); err != nil {
			return diag.Errorf("error reading server: %v", err)
		}
	}

	if o.ServerKey != nil {
		v := *o.ServerKey

		if v == "ENC XXXX" {
		} else if err = d.Set("server_key", v); err != nil {
			return diag.Errorf("error reading server_key: %v", err)
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

func getObjectLogTacacsaccounting3Setting(d *schema.ResourceData, sv string) (*models.LogTacacsaccounting3Setting, diag.Diagnostics) {
	obj := models.LogTacacsaccounting3Setting{}
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
	if v1, ok := d.GetOk("server_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_key", sv)
				diags = append(diags, e)
			}
			obj.ServerKey = &v2
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
func getEmptyObjectLogTacacsaccounting3Setting(d *schema.ResourceData, sv string) (*models.LogTacacsaccounting3Setting, diag.Diagnostics) {
	obj := models.LogTacacsaccounting3Setting{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
