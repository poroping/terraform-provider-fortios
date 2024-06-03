// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Settings for TACACS+ accounting events filter.

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

func resourceLogTacacsaccountingFilter() *schema.Resource {
	return &schema.Resource{
		Description: "Settings for TACACS+ accounting events filter.",

		CreateContext: resourceLogTacacsaccountingFilterCreate,
		ReadContext:   resourceLogTacacsaccountingFilterRead,
		UpdateContext: resourceLogTacacsaccountingFilterUpdate,
		DeleteContext: resourceLogTacacsaccountingFilterDelete,

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
			"cli_cmd_audit": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable TACACS+ accounting for CLI commands audit.",
				Optional:    true,
				Computed:    true,
			},
			"config_change_audit": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable TACACS+ accounting for configuration change events audit.",
				Optional:    true,
				Computed:    true,
			},
			"login_audit": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable TACACS+ accounting for login events audit.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceLogTacacsaccountingFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogTacacsaccountingFilter(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateLogTacacsaccountingFilter(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogTacacsaccountingFilter")
	}

	return resourceLogTacacsaccountingFilterRead(ctx, d, meta)
}

func resourceLogTacacsaccountingFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogTacacsaccountingFilter(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateLogTacacsaccountingFilter(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating LogTacacsaccountingFilter resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogTacacsaccountingFilter")
	}

	return resourceLogTacacsaccountingFilterRead(ctx, d, meta)
}

func resourceLogTacacsaccountingFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectLogTacacsaccountingFilter(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateLogTacacsaccountingFilter(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting LogTacacsaccountingFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogTacacsaccountingFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadLogTacacsaccountingFilter(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogTacacsaccountingFilter resource: %v", err)
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

	diags := refreshObjectLogTacacsaccountingFilter(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectLogTacacsaccountingFilter(d *schema.ResourceData, o *models.LogTacacsaccountingFilter, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.CliCmdAudit != nil {
		v := *o.CliCmdAudit

		if err = d.Set("cli_cmd_audit", v); err != nil {
			return diag.Errorf("error reading cli_cmd_audit: %v", err)
		}
	}

	if o.ConfigChangeAudit != nil {
		v := *o.ConfigChangeAudit

		if err = d.Set("config_change_audit", v); err != nil {
			return diag.Errorf("error reading config_change_audit: %v", err)
		}
	}

	if o.LoginAudit != nil {
		v := *o.LoginAudit

		if err = d.Set("login_audit", v); err != nil {
			return diag.Errorf("error reading login_audit: %v", err)
		}
	}

	return nil
}

func getObjectLogTacacsaccountingFilter(d *schema.ResourceData, sv string) (*models.LogTacacsaccountingFilter, diag.Diagnostics) {
	obj := models.LogTacacsaccountingFilter{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("cli_cmd_audit"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cli_cmd_audit", sv)
				diags = append(diags, e)
			}
			obj.CliCmdAudit = &v2
		}
	}
	if v1, ok := d.GetOk("config_change_audit"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("config_change_audit", sv)
				diags = append(diags, e)
			}
			obj.ConfigChangeAudit = &v2
		}
	}
	if v1, ok := d.GetOk("login_audit"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("login_audit", sv)
				diags = append(diags, e)
			}
			obj.LoginAudit = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectLogTacacsaccountingFilter(d *schema.ResourceData, sv string) (*models.LogTacacsaccountingFilter, diag.Diagnostics) {
	obj := models.LogTacacsaccountingFilter{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
