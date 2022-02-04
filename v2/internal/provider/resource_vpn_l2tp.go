// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure L2TP.

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

func resourceVpnL2tp() *schema.Resource {
	return &schema.Resource{
		Description: "Configure L2TP.",

		CreateContext: resourceVpnL2tpCreate,
		ReadContext:   resourceVpnL2tpRead,
		UpdateContext: resourceVpnL2tpUpdate,
		DeleteContext: resourceVpnL2tpDelete,

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
			"compress": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable data compression.",
				Optional:    true,
				Computed:    true,
			},
			"eip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "End IP.",
				Optional:    true,
				Computed:    true,
			},
			"enforce_ipsec": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPsec enforcement.",
				Optional:    true,
				Computed:    true,
			},
			"hello_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3600),

				Description: "L2TP hello message interval in seconds (0 - 3600 sec, default = 60).",
				Optional:    true,
				Computed:    true,
			},
			"lcp_echo_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32767),

				Description: "Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.",
				Optional:    true,
				Computed:    true,
			},
			"lcp_max_echo_fails": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32767),

				Description: "Maximum number of missed LCP echo messages before disconnect.",
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

				Description: "Enable/disable FortiGate as a L2TP gateway.",
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

func resourceVpnL2tpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVpnL2tp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateVpnL2tp(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnL2tp")
	}

	return resourceVpnL2tpRead(ctx, d, meta)
}

func resourceVpnL2tpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVpnL2tp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateVpnL2tp(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating VpnL2tp resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnL2tp")
	}

	return resourceVpnL2tpRead(ctx, d, meta)
}

func resourceVpnL2tpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectVpnL2tp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateVpnL2tp(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting VpnL2tp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnL2tpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnL2tp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnL2tp resource: %v", err)
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

	diags := refreshObjectVpnL2tp(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectVpnL2tp(d *schema.ResourceData, o *models.VpnL2tp, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Compress != nil {
		v := *o.Compress

		if err = d.Set("compress", v); err != nil {
			return diag.Errorf("error reading compress: %v", err)
		}
	}

	if o.Eip != nil {
		v := *o.Eip

		if err = d.Set("eip", v); err != nil {
			return diag.Errorf("error reading eip: %v", err)
		}
	}

	if o.EnforceIpsec != nil {
		v := *o.EnforceIpsec

		if err = d.Set("enforce_ipsec", v); err != nil {
			return diag.Errorf("error reading enforce_ipsec: %v", err)
		}
	}

	if o.HelloInterval != nil {
		v := *o.HelloInterval

		if err = d.Set("hello_interval", v); err != nil {
			return diag.Errorf("error reading hello_interval: %v", err)
		}
	}

	if o.LcpEchoInterval != nil {
		v := *o.LcpEchoInterval

		if err = d.Set("lcp_echo_interval", v); err != nil {
			return diag.Errorf("error reading lcp_echo_interval: %v", err)
		}
	}

	if o.LcpMaxEchoFails != nil {
		v := *o.LcpMaxEchoFails

		if err = d.Set("lcp_max_echo_fails", v); err != nil {
			return diag.Errorf("error reading lcp_max_echo_fails: %v", err)
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

func getObjectVpnL2tp(d *schema.ResourceData, sv string) (*models.VpnL2tp, diag.Diagnostics) {
	obj := models.VpnL2tp{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("compress"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("compress", sv)
				diags = append(diags, e)
			}
			obj.Compress = &v2
		}
	}
	if v1, ok := d.GetOk("eip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eip", sv)
				diags = append(diags, e)
			}
			obj.Eip = &v2
		}
	}
	if v1, ok := d.GetOk("enforce_ipsec"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("enforce_ipsec", sv)
				diags = append(diags, e)
			}
			obj.EnforceIpsec = &v2
		}
	}
	if v1, ok := d.GetOk("hello_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("hello_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HelloInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("lcp_echo_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("lcp_echo_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LcpEchoInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("lcp_max_echo_fails"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("lcp_max_echo_fails", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LcpMaxEchoFails = &tmp
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
func getEmptyObjectVpnL2tp(d *schema.ResourceData, sv string) (*models.VpnL2tp, diag.Diagnostics) {
	obj := models.VpnL2tp{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
