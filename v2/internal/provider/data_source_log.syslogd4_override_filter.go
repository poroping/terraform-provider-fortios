// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Override filters for remote system server.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceLogSyslogd4OverrideFilter() *schema.Resource {
	return &schema.Resource{
		Description: "Override filters for remote system server.",

		ReadContext: dataSourceLogSyslogd4OverrideFilterRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"anomaly": {
				Type:        schema.TypeString,
				Description: "Enable/disable anomaly logging.",
				Computed:    true,
			},
			"filter": {
				Type:        schema.TypeString,
				Description: "Syslog 4 filter.",
				Computed:    true,
			},
			"filter_type": {
				Type:        schema.TypeString,
				Description: "Include/exclude logs that match the filter.",
				Computed:    true,
			},
			"forward_traffic": {
				Type:        schema.TypeString,
				Description: "Enable/disable forward traffic logging.",
				Computed:    true,
			},
			"free_style": {
				Type:        schema.TypeList,
				Description: "Free Style Filters",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:        schema.TypeString,
							Description: "Log category.",
							Computed:    true,
						},
						"filter": {
							Type:        schema.TypeString,
							Description: "Free style filter string.",
							Computed:    true,
						},
						"filter_type": {
							Type:        schema.TypeString,
							Description: "Include/exclude logs that match the filter.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Entry ID.",
							Computed:    true,
						},
					},
				},
			},
			"gtp": {
				Type:        schema.TypeString,
				Description: "Enable/disable GTP messages logging.",
				Computed:    true,
			},
			"local_traffic": {
				Type:        schema.TypeString,
				Description: "Enable/disable local in or out traffic logging.",
				Computed:    true,
			},
			"multicast_traffic": {
				Type:        schema.TypeString,
				Description: "Enable/disable multicast traffic logging.",
				Computed:    true,
			},
			"severity": {
				Type:        schema.TypeString,
				Description: "Lowest severity level to log.",
				Computed:    true,
			},
			"sniffer_traffic": {
				Type:        schema.TypeString,
				Description: "Enable/disable sniffer traffic logging.",
				Computed:    true,
			},
			"voip": {
				Type:        schema.TypeString,
				Description: "Enable/disable VoIP logging.",
				Computed:    true,
			},
		},
	}
}

func dataSourceLogSyslogd4OverrideFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadLogSyslogd4OverrideFilter(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogSyslogd4OverrideFilter dataSource: %v", err)
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

	diags := refreshObjectLogSyslogd4OverrideFilter(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
