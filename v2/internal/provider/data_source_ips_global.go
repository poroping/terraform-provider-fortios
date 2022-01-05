// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPS global parameter.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceIpsGlobal() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPS global parameter.",

		ReadContext: dataSourceIpsGlobalRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"anomaly_mode": {
				Type:        schema.TypeString,
				Description: "Global blocking mode for rate-based anomalies.",
				Computed:    true,
			},
			"cp_accel_mode": {
				Type:        schema.TypeString,
				Description: "IPS Pattern matching acceleration/offloading to CPx processors.",
				Computed:    true,
			},
			"database": {
				Type:        schema.TypeString,
				Description: "Regular or extended IPS database. Regular protects against the latest common and in-the-wild attacks. Extended includes protection from legacy attacks.",
				Computed:    true,
			},
			"deep_app_insp_db_limit": {
				Type:        schema.TypeInt,
				Description: "Limit on number of entries in deep application inspection database (1 - 2147483647, 0 = use recommended setting)",
				Computed:    true,
			},
			"deep_app_insp_timeout": {
				Type:        schema.TypeInt,
				Description: "Timeout for Deep application inspection (1 - 2147483647 sec., 0 = use recommended setting).",
				Computed:    true,
			},
			"engine_count": {
				Type:        schema.TypeInt,
				Description: "Number of IPS engines running. If set to the default value of 0, FortiOS sets the number to optimize performance depending on the number of CPU cores.",
				Computed:    true,
			},
			"exclude_signatures": {
				Type:        schema.TypeString,
				Description: "Excluded signatures.",
				Computed:    true,
			},
			"fail_open": {
				Type:        schema.TypeString,
				Description: "Enable to allow traffic if the IPS buffer is full. Default is disable and IPS traffic is blocked when the IPS buffer is full.",
				Computed:    true,
			},
			"intelligent_mode": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPS adaptive scanning (intelligent mode). Intelligent mode optimizes the scanning method for the type of traffic.",
				Computed:    true,
			},
			"ips_reserve_cpu": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPS daemon's use of CPUs other than CPU 0",
				Computed:    true,
			},
			"ngfw_max_scan_range": {
				Type:        schema.TypeInt,
				Description: "NGFW policy-mode app detection threshold.",
				Computed:    true,
			},
			"np_accel_mode": {
				Type:        schema.TypeString,
				Description: "Acceleration mode for IPS processing by NPx processors.",
				Computed:    true,
			},
			"packet_log_queue_depth": {
				Type:        schema.TypeInt,
				Description: "Packet/pcap log queue depth per IPS engine.",
				Computed:    true,
			},
			"session_limit_mode": {
				Type:        schema.TypeString,
				Description: "Method of counting concurrent sessions used by session limit anomalies. Choose between greater accuracy (accurate) or improved performance (heuristics).",
				Computed:    true,
			},
			"skype_client_public_ipaddr": {
				Type:        schema.TypeString,
				Description: "Public IP addresses of your network that receive Skype sessions. Helps identify Skype sessions. Separate IP addresses with commas.",
				Computed:    true,
			},
			"socket_size": {
				Type:        schema.TypeInt,
				Description: "IPS socket buffer size. Max and default value depend on available memory. Can be changed to tune performance.",
				Computed:    true,
			},
			"sync_session_ttl": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of kernel session TTL for IPS sessions.",
				Computed:    true,
			},
			"tls_active_probe": {
				Type:        schema.TypeList,
				Description: "TLS active probe configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface": {
							Type:        schema.TypeString,
							Description: "Specify outgoing interface to reach server.",
							Computed:    true,
						},
						"interface_select_method": {
							Type:        schema.TypeString,
							Description: "Specify how to select outgoing interface to reach server.",
							Computed:    true,
						},
						"source_ip": {
							Type:        schema.TypeString,
							Description: "Source IP address used for TLS active probe.",
							Computed:    true,
						},
						"source_ip6": {
							Type:        schema.TypeString,
							Description: "Source IPv6 address used for TLS active probe.",
							Computed:    true,
						},
						"vdom": {
							Type:        schema.TypeString,
							Description: "Virtual domain name for TLS active probe.",
							Computed:    true,
						},
					},
				},
			},
			"traffic_submit": {
				Type:        schema.TypeString,
				Description: "Enable/disable submitting attack data found by this FortiGate to FortiGuard.",
				Computed:    true,
			},
		},
	}
}

func dataSourceIpsGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadIpsGlobal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading IpsGlobal dataSource: %v", err)
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

	diags := refreshObjectIpsGlobal(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
