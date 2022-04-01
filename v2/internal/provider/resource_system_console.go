// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure console.

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

func resourceSystemConsole() *schema.Resource {
	return &schema.Resource{
		Description: "Configure console.",

		CreateContext: resourceSystemConsoleCreate,
		ReadContext:   resourceSystemConsoleRead,
		UpdateContext: resourceSystemConsoleUpdate,
		DeleteContext: resourceSystemConsoleDelete,

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
			"baudrate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"9600", "19200", "38400", "57600", "115200"}, false),

				Description: "Console baud rate.",
				Optional:    true,
				Computed:    true,
			},
			"fortiexplorer": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable access for FortiExplorer.",
				Optional:    true,
				Computed:    true,
			},
			"login": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable serial console and FortiExplorer.",
				Optional:    true,
				Computed:    true,
			},
			"mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"batch", "line"}, false),

				Description: "Console mode.",
				Optional:    true,
				Computed:    true,
			},
			"output": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"standard", "more"}, false),

				Description: "Console output mode.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemConsoleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemConsole(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemConsole(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemConsole")
	}

	return resourceSystemConsoleRead(ctx, d, meta)
}

func resourceSystemConsoleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemConsole(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemConsole(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemConsole resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemConsole")
	}

	return resourceSystemConsoleRead(ctx, d, meta)
}

func resourceSystemConsoleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemConsole(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemConsole(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemConsole resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemConsoleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemConsole(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemConsole resource: %v", err)
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

	diags := refreshObjectSystemConsole(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemConsole(d *schema.ResourceData, o *models.SystemConsole, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Baudrate != nil {
		v := *o.Baudrate

		if err = d.Set("baudrate", v); err != nil {
			return diag.Errorf("error reading baudrate: %v", err)
		}
	}

	if o.Fortiexplorer != nil {
		v := *o.Fortiexplorer

		if err = d.Set("fortiexplorer", v); err != nil {
			return diag.Errorf("error reading fortiexplorer: %v", err)
		}
	}

	if o.Login != nil {
		v := *o.Login

		if err = d.Set("login", v); err != nil {
			return diag.Errorf("error reading login: %v", err)
		}
	}

	if o.Mode != nil {
		v := *o.Mode

		if err = d.Set("mode", v); err != nil {
			return diag.Errorf("error reading mode: %v", err)
		}
	}

	if o.Output != nil {
		v := *o.Output

		if err = d.Set("output", v); err != nil {
			return diag.Errorf("error reading output: %v", err)
		}
	}

	return nil
}

func getObjectSystemConsole(d *schema.ResourceData, sv string) (*models.SystemConsole, diag.Diagnostics) {
	obj := models.SystemConsole{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("baudrate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("baudrate", sv)
				diags = append(diags, e)
			}
			obj.Baudrate = &v2
		}
	}
	if v1, ok := d.GetOk("fortiexplorer"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("fortiexplorer", sv)
				diags = append(diags, e)
			}
			obj.Fortiexplorer = &v2
		}
	}
	if v1, ok := d.GetOk("login"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("login", sv)
				diags = append(diags, e)
			}
			obj.Login = &v2
		}
	}
	if v1, ok := d.GetOk("mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mode", sv)
				diags = append(diags, e)
			}
			obj.Mode = &v2
		}
	}
	if v1, ok := d.GetOk("output"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("output", sv)
				diags = append(diags, e)
			}
			obj.Output = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemConsole(d *schema.ResourceData, sv string) (*models.SystemConsole, diag.Diagnostics) {
	obj := models.SystemConsole{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
