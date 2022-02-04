// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure PPTP.

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

func resourceVpnPptp() *schema.Resource {
	return &schema.Resource{
		Description: "Configure PPTP.",

		CreateContext: resourceVpnPptpCreate,
		ReadContext:   resourceVpnPptpRead,
		UpdateContext: resourceVpnPptpUpdate,
		DeleteContext: resourceVpnPptpDelete,

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
			"eip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "End IP.",
				Optional:    true,
				Computed:    true,
			},
			"ip_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"range", "usrgrp"}, false),

				Description: "IP assignment mode for PPTP client.",
				Optional:    true,
				Computed:    true,
			},
			"local_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Local IP to be used for peer's remote IP.",
				Optional:    true,
				Computed:    true,
			},
			"sip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Start IP.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiGate as a PPTP gateway.",
				Optional:    true,
				Computed:    true,
			},
			"usrgrp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "User group.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceVpnPptpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVpnPptp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateVpnPptp(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnPptp")
	}

	return resourceVpnPptpRead(ctx, d, meta)
}

func resourceVpnPptpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVpnPptp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateVpnPptp(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating VpnPptp resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnPptp")
	}

	return resourceVpnPptpRead(ctx, d, meta)
}

func resourceVpnPptpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectVpnPptp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateVpnPptp(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting VpnPptp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnPptpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnPptp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnPptp resource: %v", err)
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

	diags := refreshObjectVpnPptp(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectVpnPptp(d *schema.ResourceData, o *models.VpnPptp, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Eip != nil {
		v := *o.Eip

		if err = d.Set("eip", v); err != nil {
			return diag.Errorf("error reading eip: %v", err)
		}
	}

	if o.IpMode != nil {
		v := *o.IpMode

		if err = d.Set("ip_mode", v); err != nil {
			return diag.Errorf("error reading ip_mode: %v", err)
		}
	}

	if o.LocalIp != nil {
		v := *o.LocalIp

		if err = d.Set("local_ip", v); err != nil {
			return diag.Errorf("error reading local_ip: %v", err)
		}
	}

	if o.Sip != nil {
		v := *o.Sip

		if err = d.Set("sip", v); err != nil {
			return diag.Errorf("error reading sip: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Usrgrp != nil {
		v := *o.Usrgrp

		if err = d.Set("usrgrp", v); err != nil {
			return diag.Errorf("error reading usrgrp: %v", err)
		}
	}

	return nil
}

func getObjectVpnPptp(d *schema.ResourceData, sv string) (*models.VpnPptp, diag.Diagnostics) {
	obj := models.VpnPptp{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("eip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eip", sv)
				diags = append(diags, e)
			}
			obj.Eip = &v2
		}
	}
	if v1, ok := d.GetOk("ip_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip_mode", sv)
				diags = append(diags, e)
			}
			obj.IpMode = &v2
		}
	}
	if v1, ok := d.GetOk("local_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_ip", sv)
				diags = append(diags, e)
			}
			obj.LocalIp = &v2
		}
	}
	if v1, ok := d.GetOk("sip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sip", sv)
				diags = append(diags, e)
			}
			obj.Sip = &v2
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
	if v1, ok := d.GetOk("usrgrp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("usrgrp", sv)
				diags = append(diags, e)
			}
			obj.Usrgrp = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectVpnPptp(d *schema.ResourceData, sv string) (*models.VpnPptp, diag.Diagnostics) {
	obj := models.VpnPptp{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
