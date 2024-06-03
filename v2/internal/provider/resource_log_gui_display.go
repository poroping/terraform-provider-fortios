// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure how log messages are displayed on the GUI.

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

func resourceLogGuiDisplay() *schema.Resource {
	return &schema.Resource{
		Description: "Configure how log messages are displayed on the GUI.",

		CreateContext: resourceLogGuiDisplayCreate,
		ReadContext:   resourceLogGuiDisplayRead,
		UpdateContext: resourceLogGuiDisplayUpdate,
		DeleteContext: resourceLogGuiDisplayDelete,

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
			"fortiview_unscanned_apps": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable showing unscanned traffic in FortiView application charts.",
				Optional:    true,
				Computed:    true,
			},
			"resolve_apps": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Resolve unknown applications on the GUI using Fortinet's remote application database.",
				Optional:    true,
				Computed:    true,
			},
			"resolve_hosts": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable resolving IP addresses to hostname in log messages on the GUI using reverse DNS lookup.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceLogGuiDisplayCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogGuiDisplay(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateLogGuiDisplay(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogGuiDisplay")
	}

	return resourceLogGuiDisplayRead(ctx, d, meta)
}

func resourceLogGuiDisplayUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogGuiDisplay(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateLogGuiDisplay(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating LogGuiDisplay resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogGuiDisplay")
	}

	return resourceLogGuiDisplayRead(ctx, d, meta)
}

func resourceLogGuiDisplayDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectLogGuiDisplay(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateLogGuiDisplay(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting LogGuiDisplay resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogGuiDisplayRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadLogGuiDisplay(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogGuiDisplay resource: %v", err)
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

	diags := refreshObjectLogGuiDisplay(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectLogGuiDisplay(d *schema.ResourceData, o *models.LogGuiDisplay, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.FortiviewUnscannedApps != nil {
		v := *o.FortiviewUnscannedApps

		if err = d.Set("fortiview_unscanned_apps", v); err != nil {
			return diag.Errorf("error reading fortiview_unscanned_apps: %v", err)
		}
	}

	if o.ResolveApps != nil {
		v := *o.ResolveApps

		if err = d.Set("resolve_apps", v); err != nil {
			return diag.Errorf("error reading resolve_apps: %v", err)
		}
	}

	if o.ResolveHosts != nil {
		v := *o.ResolveHosts

		if err = d.Set("resolve_hosts", v); err != nil {
			return diag.Errorf("error reading resolve_hosts: %v", err)
		}
	}

	return nil
}

func getObjectLogGuiDisplay(d *schema.ResourceData, sv string) (*models.LogGuiDisplay, diag.Diagnostics) {
	obj := models.LogGuiDisplay{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("fortiview_unscanned_apps"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fortiview_unscanned_apps", sv)
				diags = append(diags, e)
			}
			obj.FortiviewUnscannedApps = &v2
		}
	}
	if v1, ok := d.GetOk("resolve_apps"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("resolve_apps", sv)
				diags = append(diags, e)
			}
			obj.ResolveApps = &v2
		}
	}
	if v1, ok := d.GetOk("resolve_hosts"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("resolve_hosts", sv)
				diags = append(diags, e)
			}
			obj.ResolveHosts = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectLogGuiDisplay(d *schema.ResourceData, sv string) (*models.LogGuiDisplay, diag.Diagnostics) {
	obj := models.LogGuiDisplay{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
