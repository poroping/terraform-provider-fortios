// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure per-IP traffic shaper.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallShaperPerIpShaper() *schema.Resource {
	return &schema.Resource{
		Description: "Configure per-IP traffic shaper.",

		ReadContext: dataSourceFirewallShaperPerIpShaperRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"bandwidth_unit": {
				Type:        schema.TypeString,
				Description: "Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps).",
				Computed:    true,
			},
			"diffserv_forward": {
				Type:        schema.TypeString,
				Description: "Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper.",
				Computed:    true,
			},
			"diffserv_reverse": {
				Type:        schema.TypeString,
				Description: "Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper.",
				Computed:    true,
			},
			"diffservcode_forward": {
				Type:        schema.TypeString,
				Description: "Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.",
				Computed:    true,
			},
			"diffservcode_rev": {
				Type:        schema.TypeString,
				Description: "Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.",
				Computed:    true,
			},
			"max_bandwidth": {
				Type:        schema.TypeInt,
				Description: "Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.",
				Computed:    true,
			},
			"max_concurrent_session": {
				Type:        schema.TypeInt,
				Description: "Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.",
				Computed:    true,
			},
			"max_concurrent_tcp_session": {
				Type:        schema.TypeInt,
				Description: "Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.",
				Computed:    true,
			},
			"max_concurrent_udp_session": {
				Type:        schema.TypeInt,
				Description: "Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Traffic shaper name.",
				Required:    true,
			},
		},
	}
}

func dataSourceFirewallShaperPerIpShaperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallShaperPerIpShaper(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallShaperPerIpShaper dataSource: %v", err)
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

	diags := refreshObjectFirewallShaperPerIpShaper(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
