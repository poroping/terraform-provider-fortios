// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure shared traffic shaper.

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

func resourceFirewallShaperTrafficShaper() *schema.Resource {
	return &schema.Resource{
		Description: "Configure shared traffic shaper.",

		CreateContext: resourceFirewallShaperTrafficShaperCreate,
		ReadContext:   resourceFirewallShaperTrafficShaperRead,
		UpdateContext: resourceFirewallShaperTrafficShaperUpdate,
		DeleteContext: resourceFirewallShaperTrafficShaperDelete,

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

				Description: "Unit of measurement for guaranteed and maximum bandwidth for this shaper (Kbps, Mbps or Gbps).",
				Optional:    true,
				Computed:    true,
			},
			"diffserv": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable changing the DiffServ setting applied to traffic accepted by this shaper.",
				Optional:    true,
				Computed:    true,
			},
			"diffservcode": {
				Type: schema.TypeString,

				Description: "DiffServ setting to be applied to traffic accepted by this shaper.",
				Optional:    true,
				Computed:    true,
			},
			"dscp_marking_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"multi-stage", "static"}, false),

				Description: "Select DSCP marking method.",
				Optional:    true,
				Computed:    true,
			},
			"exceed_bandwidth": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),

				Description: "Exceed bandwidth used for DSCP multi-stage marking. Units depend on the bandwidth-unit setting.",
				Optional:    true,
				Computed:    true,
			},
			"exceed_class_id": {
				Type: schema.TypeInt,

				Description: "Class ID for traffic in [guaranteed-bandwidth, maximum-bandwidth].",
				Optional:    true,
				Computed:    true,
			},
			"exceed_dscp": {
				Type: schema.TypeString,

				Description: "DSCP mark for traffic in [guaranteed-bandwidth, exceed-bandwidth].",
				Optional:    true,
				Computed:    true,
			},
			"guaranteed_bandwidth": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),

				Description: "Amount of bandwidth guaranteed for this shaper (0 - 16776000). Units depend on the bandwidth-unit setting.",
				Optional:    true,
				Computed:    true,
			},
			"maximum_bandwidth": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),

				Description: "Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.",
				Optional:    true,
				Computed:    true,
			},
			"maximum_dscp": {
				Type: schema.TypeString,

				Description: "DSCP mark for traffic in [exceed-bandwidth, maximum-bandwidth].",
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
			"overhead": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),

				Description: "Per-packet size overhead used in rate computations.",
				Optional:    true,
				Computed:    true,
			},
			"per_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable applying a separate shaper for each policy. For example, if enabled the guaranteed bandwidth is applied separately for each policy.",
				Optional:    true,
				Computed:    true,
			},
			"priority": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"low", "medium", "high"}, false),

				Description: "Higher priority traffic is more likely to be forwarded without delays and without compromising the guaranteed bandwidth.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallShaperTrafficShaperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallShaperTrafficShaper resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallShaperTrafficShaper(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallShaperTrafficShaper(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallShaperTrafficShaper")
	}

	return resourceFirewallShaperTrafficShaperRead(ctx, d, meta)
}

func resourceFirewallShaperTrafficShaperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallShaperTrafficShaper(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallShaperTrafficShaper(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallShaperTrafficShaper resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallShaperTrafficShaper")
	}

	return resourceFirewallShaperTrafficShaperRead(ctx, d, meta)
}

func resourceFirewallShaperTrafficShaperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallShaperTrafficShaper(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallShaperTrafficShaper resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallShaperTrafficShaperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallShaperTrafficShaper(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallShaperTrafficShaper resource: %v", err)
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

	diags := refreshObjectFirewallShaperTrafficShaper(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectFirewallShaperTrafficShaper(d *schema.ResourceData, o *models.FirewallShaperTrafficShaper, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.BandwidthUnit != nil {
		v := *o.BandwidthUnit

		if err = d.Set("bandwidth_unit", v); err != nil {
			return diag.Errorf("error reading bandwidth_unit: %v", err)
		}
	}

	if o.Diffserv != nil {
		v := *o.Diffserv

		if err = d.Set("diffserv", v); err != nil {
			return diag.Errorf("error reading diffserv: %v", err)
		}
	}

	if o.Diffservcode != nil {
		v := *o.Diffservcode

		if err = d.Set("diffservcode", v); err != nil {
			return diag.Errorf("error reading diffservcode: %v", err)
		}
	}

	if o.DscpMarkingMethod != nil {
		v := *o.DscpMarkingMethod

		if err = d.Set("dscp_marking_method", v); err != nil {
			return diag.Errorf("error reading dscp_marking_method: %v", err)
		}
	}

	if o.ExceedBandwidth != nil {
		v := *o.ExceedBandwidth

		if err = d.Set("exceed_bandwidth", v); err != nil {
			return diag.Errorf("error reading exceed_bandwidth: %v", err)
		}
	}

	if o.ExceedClassId != nil {
		v := *o.ExceedClassId

		if err = d.Set("exceed_class_id", v); err != nil {
			return diag.Errorf("error reading exceed_class_id: %v", err)
		}
	}

	if o.ExceedDscp != nil {
		v := *o.ExceedDscp

		if err = d.Set("exceed_dscp", v); err != nil {
			return diag.Errorf("error reading exceed_dscp: %v", err)
		}
	}

	if o.GuaranteedBandwidth != nil {
		v := *o.GuaranteedBandwidth

		if err = d.Set("guaranteed_bandwidth", v); err != nil {
			return diag.Errorf("error reading guaranteed_bandwidth: %v", err)
		}
	}

	if o.MaximumBandwidth != nil {
		v := *o.MaximumBandwidth

		if err = d.Set("maximum_bandwidth", v); err != nil {
			return diag.Errorf("error reading maximum_bandwidth: %v", err)
		}
	}

	if o.MaximumDscp != nil {
		v := *o.MaximumDscp

		if err = d.Set("maximum_dscp", v); err != nil {
			return diag.Errorf("error reading maximum_dscp: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Overhead != nil {
		v := *o.Overhead

		if err = d.Set("overhead", v); err != nil {
			return diag.Errorf("error reading overhead: %v", err)
		}
	}

	if o.PerPolicy != nil {
		v := *o.PerPolicy

		if err = d.Set("per_policy", v); err != nil {
			return diag.Errorf("error reading per_policy: %v", err)
		}
	}

	if o.Priority != nil {
		v := *o.Priority

		if err = d.Set("priority", v); err != nil {
			return diag.Errorf("error reading priority: %v", err)
		}
	}

	return nil
}

func getObjectFirewallShaperTrafficShaper(d *schema.ResourceData, sv string) (*models.FirewallShaperTrafficShaper, diag.Diagnostics) {
	obj := models.FirewallShaperTrafficShaper{}
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
	if v1, ok := d.GetOk("diffserv"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("diffserv", sv)
				diags = append(diags, e)
			}
			obj.Diffserv = &v2
		}
	}
	if v1, ok := d.GetOk("diffservcode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("diffservcode", sv)
				diags = append(diags, e)
			}
			obj.Diffservcode = &v2
		}
	}
	if v1, ok := d.GetOk("dscp_marking_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dscp_marking_method", sv)
				diags = append(diags, e)
			}
			obj.DscpMarkingMethod = &v2
		}
	}
	if v1, ok := d.GetOk("exceed_bandwidth"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("exceed_bandwidth", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ExceedBandwidth = &tmp
		}
	}
	if v1, ok := d.GetOk("exceed_class_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("exceed_class_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ExceedClassId = &tmp
		}
	}
	if v1, ok := d.GetOk("exceed_dscp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("exceed_dscp", sv)
				diags = append(diags, e)
			}
			obj.ExceedDscp = &v2
		}
	}
	if v1, ok := d.GetOk("guaranteed_bandwidth"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("guaranteed_bandwidth", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.GuaranteedBandwidth = &tmp
		}
	}
	if v1, ok := d.GetOk("maximum_bandwidth"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("maximum_bandwidth", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaximumBandwidth = &tmp
		}
	}
	if v1, ok := d.GetOk("maximum_dscp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("maximum_dscp", sv)
				diags = append(diags, e)
			}
			obj.MaximumDscp = &v2
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
	if v1, ok := d.GetOk("overhead"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("overhead", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Overhead = &tmp
		}
	}
	if v1, ok := d.GetOk("per_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("per_policy", sv)
				diags = append(diags, e)
			}
			obj.PerPolicy = &v2
		}
	}
	if v1, ok := d.GetOk("priority"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("priority", sv)
				diags = append(diags, e)
			}
			obj.Priority = &v2
		}
	}
	return &obj, diags
}
