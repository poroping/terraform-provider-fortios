// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure filters for local disk logging. Use these filters to determine the log messages to record according to severity and type.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceLogDiskFilter() *schema.Resource {
	return &schema.Resource{
		Description: "Configure filters for local disk logging. Use these filters to determine the log messages to record according to severity and type.",

		ReadContext: dataSourceLogDiskFilterRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"admin": {
				Type:        schema.TypeString,
				Description: "Enable/disable admin login/logout logging.",
				Computed:    true,
			},
			"anomaly": {
				Type:        schema.TypeString,
				Description: "Enable/disable anomaly logging.",
				Computed:    true,
			},
			"auth": {
				Type:        schema.TypeString,
				Description: "Enable/disable firewall authentication logging.",
				Computed:    true,
			},
			"cpu_memory_usage": {
				Type:        schema.TypeString,
				Description: "Enable/disable CPU & memory usage logging every 5 minutes.",
				Computed:    true,
			},
			"dhcp": {
				Type:        schema.TypeString,
				Description: "Enable/disable DHCP service messages logging.",
				Computed:    true,
			},
			"dlp_archive": {
				Type:        schema.TypeString,
				Description: "Enable/disable DLP archive logging.",
				Computed:    true,
			},
			"event": {
				Type:        schema.TypeString,
				Description: "Enable/disable event logging.",
				Computed:    true,
			},
			"filter": {
				Type:        schema.TypeString,
				Description: "Disk log filter.",
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
				Description: "Free style filters.",
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
			"ha": {
				Type:        schema.TypeString,
				Description: "Enable/disable HA logging.",
				Computed:    true,
			},
			"ipsec": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPsec negotiation messages logging.",
				Computed:    true,
			},
			"ldb_monitor": {
				Type:        schema.TypeString,
				Description: "Enable/disable VIP real server health monitoring logging.",
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
			"pattern": {
				Type:        schema.TypeString,
				Description: "Enable/disable pattern update logging.",
				Computed:    true,
			},
			"ppp": {
				Type:        schema.TypeString,
				Description: "Enable/disable L2TP/PPTP/PPPoE logging.",
				Computed:    true,
			},
			"radius": {
				Type:        schema.TypeString,
				Description: "Enable/disable RADIUS messages logging.",
				Computed:    true,
			},
			"severity": {
				Type:        schema.TypeString,
				Description: "Log to disk every message above and including this severity level.",
				Computed:    true,
			},
			"sniffer_traffic": {
				Type:        schema.TypeString,
				Description: "Enable/disable sniffer traffic logging.",
				Computed:    true,
			},
			"sslvpn_log_adm": {
				Type:        schema.TypeString,
				Description: "Enable/disable SSL administrator login logging.",
				Computed:    true,
			},
			"sslvpn_log_auth": {
				Type:        schema.TypeString,
				Description: "Enable/disable SSL user authentication logging.",
				Computed:    true,
			},
			"sslvpn_log_session": {
				Type:        schema.TypeString,
				Description: "Enable/disable SSL session logging.",
				Computed:    true,
			},
			"system": {
				Type:        schema.TypeString,
				Description: "Enable/disable system activity logging.",
				Computed:    true,
			},
			"vip_ssl": {
				Type:        schema.TypeString,
				Description: "Enable/disable VIP SSL logging.",
				Computed:    true,
			},
			"voip": {
				Type:        schema.TypeString,
				Description: "Enable/disable VoIP logging.",
				Computed:    true,
			},
			"wan_opt": {
				Type:        schema.TypeString,
				Description: "Enable/disable WAN optimization event logging.",
				Computed:    true,
			},
			"wireless_activity": {
				Type:        schema.TypeString,
				Description: "Enable/disable wireless activity event logging.",
				Computed:    true,
			},
			"ztna_traffic": {
				Type:        schema.TypeString,
				Description: "Enable/disable ztna traffic logging.",
				Computed:    true,
			},
		},
	}
}

func dataSourceLogDiskFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "LogDiskFilter"

	o, err := c.Cmdb.ReadLogDiskFilter(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogDiskFilter dataSource: %v", err)
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

	diags := refreshObjectLogDiskFilter(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
