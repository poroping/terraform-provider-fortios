// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure log event filters.

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

func resourceLogEventfilter() *schema.Resource {
	return &schema.Resource{
		Description: "Configure log event filters.",

		CreateContext: resourceLogEventfilterCreate,
		ReadContext:   resourceLogEventfilterRead,
		UpdateContext: resourceLogEventfilterUpdate,
		DeleteContext: resourceLogEventfilterDelete,

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
			"cifs": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable CIFS logging.",
				Optional:    true,
				Computed:    true,
			},
			"connector": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SDN connector logging.",
				Optional:    true,
				Computed:    true,
			},
			"endpoint": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable endpoint event logging.",
				Optional:    true,
				Computed:    true,
			},
			"event": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable event logging.",
				Optional:    true,
				Computed:    true,
			},
			"fortiextender": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiExtender logging.",
				Optional:    true,
				Computed:    true,
			},
			"ha": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable ha event logging.",
				Optional:    true,
				Computed:    true,
			},
			"rest_api": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable REST API logging.",
				Optional:    true,
				Computed:    true,
			},
			"router": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable router event logging.",
				Optional:    true,
				Computed:    true,
			},
			"sdwan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SD-WAN logging.",
				Optional:    true,
				Computed:    true,
			},
			"security_rating": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Security Rating result logging.",
				Optional:    true,
				Computed:    true,
			},
			"switch_controller": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Switch-Controller logging.",
				Optional:    true,
				Computed:    true,
			},
			"system": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable system event logging.",
				Optional:    true,
				Computed:    true,
			},
			"user": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable user authentication event logging.",
				Optional:    true,
				Computed:    true,
			},
			"vpn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable VPN event logging.",
				Optional:    true,
				Computed:    true,
			},
			"wan_opt": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable WAN optimization event logging.",
				Optional:    true,
				Computed:    true,
			},
			"wireless_activity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable wireless event logging.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceLogEventfilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogEventfilter(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateLogEventfilter(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogEventfilter")
	}

	return resourceLogEventfilterRead(ctx, d, meta)
}

func resourceLogEventfilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogEventfilter(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateLogEventfilter(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating LogEventfilter resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogEventfilter")
	}

	return resourceLogEventfilterRead(ctx, d, meta)
}

func resourceLogEventfilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectLogEventfilter(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateLogEventfilter(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting LogEventfilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogEventfilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadLogEventfilter(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogEventfilter resource: %v", err)
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

	diags := refreshObjectLogEventfilter(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectLogEventfilter(d *schema.ResourceData, o *models.LogEventfilter, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Cifs != nil {
		v := *o.Cifs

		if err = d.Set("cifs", v); err != nil {
			return diag.Errorf("error reading cifs: %v", err)
		}
	}

	if o.Connector != nil {
		v := *o.Connector

		if err = d.Set("connector", v); err != nil {
			return diag.Errorf("error reading connector: %v", err)
		}
	}

	if o.Endpoint != nil {
		v := *o.Endpoint

		if err = d.Set("endpoint", v); err != nil {
			return diag.Errorf("error reading endpoint: %v", err)
		}
	}

	if o.Event != nil {
		v := *o.Event

		if err = d.Set("event", v); err != nil {
			return diag.Errorf("error reading event: %v", err)
		}
	}

	if o.Fortiextender != nil {
		v := *o.Fortiextender

		if err = d.Set("fortiextender", v); err != nil {
			return diag.Errorf("error reading fortiextender: %v", err)
		}
	}

	if o.Ha != nil {
		v := *o.Ha

		if err = d.Set("ha", v); err != nil {
			return diag.Errorf("error reading ha: %v", err)
		}
	}

	if o.RestApi != nil {
		v := *o.RestApi

		if err = d.Set("rest_api", v); err != nil {
			return diag.Errorf("error reading rest_api: %v", err)
		}
	}

	if o.Router != nil {
		v := *o.Router

		if err = d.Set("router", v); err != nil {
			return diag.Errorf("error reading router: %v", err)
		}
	}

	if o.Sdwan != nil {
		v := *o.Sdwan

		if err = d.Set("sdwan", v); err != nil {
			return diag.Errorf("error reading sdwan: %v", err)
		}
	}

	if o.SecurityRating != nil {
		v := *o.SecurityRating

		if err = d.Set("security_rating", v); err != nil {
			return diag.Errorf("error reading security_rating: %v", err)
		}
	}

	if o.SwitchController != nil {
		v := *o.SwitchController

		if err = d.Set("switch_controller", v); err != nil {
			return diag.Errorf("error reading switch_controller: %v", err)
		}
	}

	if o.System != nil {
		v := *o.System

		if err = d.Set("system", v); err != nil {
			return diag.Errorf("error reading system: %v", err)
		}
	}

	if o.User != nil {
		v := *o.User

		if err = d.Set("user", v); err != nil {
			return diag.Errorf("error reading user: %v", err)
		}
	}

	if o.Vpn != nil {
		v := *o.Vpn

		if err = d.Set("vpn", v); err != nil {
			return diag.Errorf("error reading vpn: %v", err)
		}
	}

	if o.WanOpt != nil {
		v := *o.WanOpt

		if err = d.Set("wan_opt", v); err != nil {
			return diag.Errorf("error reading wan_opt: %v", err)
		}
	}

	if o.WirelessActivity != nil {
		v := *o.WirelessActivity

		if err = d.Set("wireless_activity", v); err != nil {
			return diag.Errorf("error reading wireless_activity: %v", err)
		}
	}

	return nil
}

func getObjectLogEventfilter(d *schema.ResourceData, sv string) (*models.LogEventfilter, diag.Diagnostics) {
	obj := models.LogEventfilter{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("cifs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("cifs", sv)
				diags = append(diags, e)
			}
			obj.Cifs = &v2
		}
	}
	if v1, ok := d.GetOk("connector"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("connector", sv)
				diags = append(diags, e)
			}
			obj.Connector = &v2
		}
	}
	if v1, ok := d.GetOk("endpoint"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("endpoint", sv)
				diags = append(diags, e)
			}
			obj.Endpoint = &v2
		}
	}
	if v1, ok := d.GetOk("event"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("event", sv)
				diags = append(diags, e)
			}
			obj.Event = &v2
		}
	}
	if v1, ok := d.GetOk("fortiextender"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fortiextender", sv)
				diags = append(diags, e)
			}
			obj.Fortiextender = &v2
		}
	}
	if v1, ok := d.GetOk("ha"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ha", sv)
				diags = append(diags, e)
			}
			obj.Ha = &v2
		}
	}
	if v1, ok := d.GetOk("rest_api"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.4", "") {
				e := utils.AttributeVersionWarning("rest_api", sv)
				diags = append(diags, e)
			}
			obj.RestApi = &v2
		}
	}
	if v1, ok := d.GetOk("router"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("router", sv)
				diags = append(diags, e)
			}
			obj.Router = &v2
		}
	}
	if v1, ok := d.GetOk("sdwan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("sdwan", sv)
				diags = append(diags, e)
			}
			obj.Sdwan = &v2
		}
	}
	if v1, ok := d.GetOk("security_rating"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("security_rating", sv)
				diags = append(diags, e)
			}
			obj.SecurityRating = &v2
		}
	}
	if v1, ok := d.GetOk("switch_controller"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("switch_controller", sv)
				diags = append(diags, e)
			}
			obj.SwitchController = &v2
		}
	}
	if v1, ok := d.GetOk("system"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("system", sv)
				diags = append(diags, e)
			}
			obj.System = &v2
		}
	}
	if v1, ok := d.GetOk("user"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("user", sv)
				diags = append(diags, e)
			}
			obj.User = &v2
		}
	}
	if v1, ok := d.GetOk("vpn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vpn", sv)
				diags = append(diags, e)
			}
			obj.Vpn = &v2
		}
	}
	if v1, ok := d.GetOk("wan_opt"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wan_opt", sv)
				diags = append(diags, e)
			}
			obj.WanOpt = &v2
		}
	}
	if v1, ok := d.GetOk("wireless_activity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wireless_activity", sv)
				diags = append(diags, e)
			}
			obj.WirelessActivity = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectLogEventfilter(d *schema.ResourceData, sv string) (*models.LogEventfilter, diag.Diagnostics) {
	obj := models.LogEventfilter{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
