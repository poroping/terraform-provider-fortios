// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure shared traffic shaper.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallShaperTrafficShaper() *schema.Resource {
	return &schema.Resource{
		Description: "Configure shared traffic shaper.",

		ReadContext: dataSourceFirewallShaperTrafficShaperRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"bandwidth_unit": {
				Type:        schema.TypeString,
				Description: "Unit of measurement for guaranteed and maximum bandwidth for this shaper (Kbps, Mbps or Gbps).",
				Computed:    true,
			},
			"diffserv": {
				Type:        schema.TypeString,
				Description: "Enable/disable changing the DiffServ setting applied to traffic accepted by this shaper.",
				Computed:    true,
			},
			"diffservcode": {
				Type:        schema.TypeString,
				Description: "DiffServ setting to be applied to traffic accepted by this shaper.",
				Computed:    true,
			},
			"dscp_marking_method": {
				Type:        schema.TypeString,
				Description: "Select DSCP marking method.",
				Computed:    true,
			},
			"exceed_bandwidth": {
				Type:        schema.TypeInt,
				Description: "Exceed bandwidth used for DSCP multi-stage marking. Units depend on the bandwidth-unit setting.",
				Computed:    true,
			},
			"exceed_class_id": {
				Type:        schema.TypeInt,
				Description: "Class ID for traffic in guaranteed-bandwidth and maximum-bandwidth.",
				Computed:    true,
			},
			"exceed_dscp": {
				Type:        schema.TypeString,
				Description: "DSCP mark for traffic in guaranteed-bandwidth and exceed-bandwidth.",
				Computed:    true,
			},
			"guaranteed_bandwidth": {
				Type:        schema.TypeInt,
				Description: "Amount of bandwidth guaranteed for this shaper (0 - 100000000). Units depend on the bandwidth-unit setting.",
				Computed:    true,
			},
			"maximum_bandwidth": {
				Type:        schema.TypeInt,
				Description: "Upper bandwidth limit enforced by this shaper (0 - 100000000). 0 means no limit. Units depend on the bandwidth-unit setting.",
				Computed:    true,
			},
			"maximum_dscp": {
				Type:        schema.TypeString,
				Description: "DSCP mark for traffic in exceed-bandwidth and maximum-bandwidth.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Traffic shaper name.",
				Required:    true,
			},
			"overhead": {
				Type:        schema.TypeInt,
				Description: "Per-packet size overhead used in rate computations.",
				Computed:    true,
			},
			"per_policy": {
				Type:        schema.TypeString,
				Description: "Enable/disable applying a separate shaper for each policy. For example, if enabled the guaranteed bandwidth is applied separately for each policy.",
				Computed:    true,
			},
			"priority": {
				Type:        schema.TypeString,
				Description: "Higher priority traffic is more likely to be forwarded without delays and without compromising the guaranteed bandwidth.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallShaperTrafficShaperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("name")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadFirewallShaperTrafficShaper(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallShaperTrafficShaper dataSource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] dataSource (%s) not found, removing from state", d.Id())
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

	d.SetId(mkey)

	return nil
}
