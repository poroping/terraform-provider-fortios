// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure per-IP traffic shaper.

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

func resourceFirewallShaperPerIpShaper() *schema.Resource {
	return &schema.Resource{
		Description: "Configure per-IP traffic shaper.",

		CreateContext: resourceFirewallShaperPerIpShaperCreate,
		ReadContext:   resourceFirewallShaperPerIpShaperRead,
		UpdateContext: resourceFirewallShaperPerIpShaperUpdate,
		DeleteContext: resourceFirewallShaperPerIpShaperDelete,

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
			"allow_append": {
				Type:        schema.TypeBool,
				Description: "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"bandwidth_unit": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"kbps", "mbps", "gbps"}, false),

				Description: "Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps).",
				Optional:    true,
				Computed:    true,
			},
			"diffserv_forward": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper.",
				Optional:    true,
				Computed:    true,
			},
			"diffserv_reverse": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper.",
				Optional:    true,
				Computed:    true,
			},
			"diffservcode_forward": {
				Type: schema.TypeString,

				Description: "Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.",
				Optional:    true,
				Computed:    true,
			},
			"diffservcode_rev": {
				Type: schema.TypeString,

				Description: "Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.",
				Optional:    true,
				Computed:    true,
			},
			"max_bandwidth": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100000000),

				Description: "Upper bandwidth limit enforced by this shaper (0 - 100000000). 0 means no limit. Units depend on the bandwidth-unit setting.",
				Optional:    true,
				Computed:    true,
			},
			"max_concurrent_session": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2097000),

				Description: "Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.",
				Optional:    true,
				Computed:    true,
			},
			"max_concurrent_tcp_session": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2097000),

				Description: "Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.",
				Optional:    true,
				Computed:    true,
			},
			"max_concurrent_udp_session": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2097000),

				Description: "Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Traffic shaper name.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceFirewallShaperPerIpShaperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := ""
	key := "name"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating FirewallShaperPerIpShaper resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallShaperPerIpShaper(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallShaperPerIpShaper(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallShaperPerIpShaper")
	}

	return resourceFirewallShaperPerIpShaperRead(ctx, d, meta)
}

func resourceFirewallShaperPerIpShaperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallShaperPerIpShaper(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallShaperPerIpShaper(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallShaperPerIpShaper resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallShaperPerIpShaper")
	}

	return resourceFirewallShaperPerIpShaperRead(ctx, d, meta)
}

func resourceFirewallShaperPerIpShaperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallShaperPerIpShaper(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallShaperPerIpShaper resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallShaperPerIpShaperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallShaperPerIpShaper(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallShaperPerIpShaper resource: %v", err)
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

	diags := refreshObjectFirewallShaperPerIpShaper(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectFirewallShaperPerIpShaper(d *schema.ResourceData, o *models.FirewallShaperPerIpShaper, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.BandwidthUnit != nil {
		v := *o.BandwidthUnit

		if err = d.Set("bandwidth_unit", v); err != nil {
			return diag.Errorf("error reading bandwidth_unit: %v", err)
		}
	}

	if o.DiffservForward != nil {
		v := *o.DiffservForward

		if err = d.Set("diffserv_forward", v); err != nil {
			return diag.Errorf("error reading diffserv_forward: %v", err)
		}
	}

	if o.DiffservReverse != nil {
		v := *o.DiffservReverse

		if err = d.Set("diffserv_reverse", v); err != nil {
			return diag.Errorf("error reading diffserv_reverse: %v", err)
		}
	}

	if o.DiffservcodeForward != nil {
		v := *o.DiffservcodeForward

		if err = d.Set("diffservcode_forward", v); err != nil {
			return diag.Errorf("error reading diffservcode_forward: %v", err)
		}
	}

	if o.DiffservcodeRev != nil {
		v := *o.DiffservcodeRev

		if err = d.Set("diffservcode_rev", v); err != nil {
			return diag.Errorf("error reading diffservcode_rev: %v", err)
		}
	}

	if o.MaxBandwidth != nil {
		v := *o.MaxBandwidth

		if err = d.Set("max_bandwidth", v); err != nil {
			return diag.Errorf("error reading max_bandwidth: %v", err)
		}
	}

	if o.MaxConcurrentSession != nil {
		v := *o.MaxConcurrentSession

		if err = d.Set("max_concurrent_session", v); err != nil {
			return diag.Errorf("error reading max_concurrent_session: %v", err)
		}
	}

	if o.MaxConcurrentTcpSession != nil {
		v := *o.MaxConcurrentTcpSession

		if err = d.Set("max_concurrent_tcp_session", v); err != nil {
			return diag.Errorf("error reading max_concurrent_tcp_session: %v", err)
		}
	}

	if o.MaxConcurrentUdpSession != nil {
		v := *o.MaxConcurrentUdpSession

		if err = d.Set("max_concurrent_udp_session", v); err != nil {
			return diag.Errorf("error reading max_concurrent_udp_session: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	return nil
}

func getObjectFirewallShaperPerIpShaper(d *schema.ResourceData, sv string) (*models.FirewallShaperPerIpShaper, diag.Diagnostics) {
	obj := models.FirewallShaperPerIpShaper{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("bandwidth_unit"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bandwidth_unit", sv)
				diags = append(diags, e)
			}
			obj.BandwidthUnit = &v2
		}
	}
	if v1, ok := d.GetOk("diffserv_forward"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("diffserv_forward", sv)
				diags = append(diags, e)
			}
			obj.DiffservForward = &v2
		}
	}
	if v1, ok := d.GetOk("diffserv_reverse"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("diffserv_reverse", sv)
				diags = append(diags, e)
			}
			obj.DiffservReverse = &v2
		}
	}
	if v1, ok := d.GetOk("diffservcode_forward"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("diffservcode_forward", sv)
				diags = append(diags, e)
			}
			obj.DiffservcodeForward = &v2
		}
	}
	if v1, ok := d.GetOk("diffservcode_rev"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("diffservcode_rev", sv)
				diags = append(diags, e)
			}
			obj.DiffservcodeRev = &v2
		}
	}
	if v1, ok := d.GetOk("max_bandwidth"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_bandwidth", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxBandwidth = &tmp
		}
	}
	if v1, ok := d.GetOk("max_concurrent_session"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_concurrent_session", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxConcurrentSession = &tmp
		}
	}
	if v1, ok := d.GetOk("max_concurrent_tcp_session"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("max_concurrent_tcp_session", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxConcurrentTcpSession = &tmp
		}
	}
	if v1, ok := d.GetOk("max_concurrent_udp_session"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("max_concurrent_udp_session", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxConcurrentUdpSession = &tmp
		}
	}
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	return &obj, diags
}
