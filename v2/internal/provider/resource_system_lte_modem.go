// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.6,v7.0.2 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure USB LTE/WIMAX devices.

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

func resourceSystemLteModem() *schema.Resource {
	return &schema.Resource{
		Description: "Configure USB LTE/WIMAX devices.",

		CreateContext: resourceSystemLteModemCreate,
		ReadContext:   resourceSystemLteModemRead,
		UpdateContext: resourceSystemLteModemUpdate,
		DeleteContext: resourceSystemLteModemDelete,

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
			"apn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Login APN string for PDP-IP packet data calls.",
				Optional:    true,
				Computed:    true,
			},
			"authtype": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "pap", "chap"}, false),

				Description: "Authentication type for PDP-IP packet data calls.",
				Optional:    true,
				Computed:    true,
			},
			"extra_init": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Extra initialization string for USB LTE/WIMAX devices.",
				Optional:    true,
				Computed:    true,
			},
			"holddown_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 60),

				Description: "Hold down timer (10 - 60 sec).",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "The interface that the modem is acting as a redundant interface for.",
				Optional:    true,
				Computed:    true,
			},
			"mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"standalone", "redundant"}, false),

				Description: "Modem operation mode.",
				Optional:    true,
				Computed:    true,
			},
			"modem_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 20),

				Description: "Modem port index (0 - 20).",
				Optional:    true,
				Computed:    true,
			},
			"passwd": {
				Type: schema.TypeString,

				Description: "Authentication password for PDP-IP packet data calls.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable USB LTE/WIMAX device.",
				Optional:    true,
				Computed:    true,
			},
			"username": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Authentication username for PDP-IP packet data calls.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemLteModemCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemLteModem(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemLteModem(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemLteModem")
	}

	return resourceSystemLteModemRead(ctx, d, meta)
}

func resourceSystemLteModemUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemLteModem(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemLteModem(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemLteModem resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemLteModem")
	}

	return resourceSystemLteModemRead(ctx, d, meta)
}

func resourceSystemLteModemDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemLteModem(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemLteModem(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemLteModem resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLteModemRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemLteModem(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemLteModem resource: %v", err)
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

	diags := refreshObjectSystemLteModem(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemLteModem(d *schema.ResourceData, o *models.SystemLteModem, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Apn != nil {
		v := *o.Apn

		if err = d.Set("apn", v); err != nil {
			return diag.Errorf("error reading apn: %v", err)
		}
	}

	if o.Authtype != nil {
		v := *o.Authtype

		if err = d.Set("authtype", v); err != nil {
			return diag.Errorf("error reading authtype: %v", err)
		}
	}

	if o.ExtraInit != nil {
		v := *o.ExtraInit

		if err = d.Set("extra_init", v); err != nil {
			return diag.Errorf("error reading extra_init: %v", err)
		}
	}

	if o.HolddownTimer != nil {
		v := *o.HolddownTimer

		if err = d.Set("holddown_timer", v); err != nil {
			return diag.Errorf("error reading holddown_timer: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.Mode != nil {
		v := *o.Mode

		if err = d.Set("mode", v); err != nil {
			return diag.Errorf("error reading mode: %v", err)
		}
	}

	if o.ModemPort != nil {
		v := *o.ModemPort

		if err = d.Set("modem_port", v); err != nil {
			return diag.Errorf("error reading modem_port: %v", err)
		}
	}

	if o.Passwd != nil {
		v := *o.Passwd

		if err = d.Set("passwd", v); err != nil {
			return diag.Errorf("error reading passwd: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Username != nil {
		v := *o.Username

		if err = d.Set("username", v); err != nil {
			return diag.Errorf("error reading username: %v", err)
		}
	}

	return nil
}

func getObjectSystemLteModem(d *schema.ResourceData, sv string) (*models.SystemLteModem, diag.Diagnostics) {
	obj := models.SystemLteModem{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("apn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("apn", sv)
				diags = append(diags, e)
			}
			obj.Apn = &v2
		}
	}
	if v1, ok := d.GetOk("authtype"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authtype", sv)
				diags = append(diags, e)
			}
			obj.Authtype = &v2
		}
	}
	if v1, ok := d.GetOk("extra_init"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("extra_init", sv)
				diags = append(diags, e)
			}
			obj.ExtraInit = &v2
		}
	}
	if v1, ok := d.GetOk("holddown_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("holddown_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HolddownTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
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
	if v1, ok := d.GetOk("modem_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("modem_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ModemPort = &tmp
		}
	}
	if v1, ok := d.GetOk("passwd"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("passwd", sv)
				diags = append(diags, e)
			}
			obj.Passwd = &v2
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
	if v1, ok := d.GetOk("username"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("username", sv)
				diags = append(diags, e)
			}
			obj.Username = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemLteModem(d *schema.ResourceData, sv string) (*models.SystemLteModem, diag.Diagnostics) {
	obj := models.SystemLteModem{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
