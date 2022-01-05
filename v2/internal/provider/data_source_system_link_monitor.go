// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Link Health Monitor.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemLinkMonitor() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Link Health Monitor.",

		ReadContext: dataSourceSystemLinkMonitorRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"addr_mode": {
				Type:        schema.TypeString,
				Description: "Address mode (IPv4 or IPv6).",
				Computed:    true,
			},
			"class_id": {
				Type:        schema.TypeInt,
				Description: "Traffic class ID.",
				Computed:    true,
			},
			"diffservcode": {
				Type:        schema.TypeString,
				Description: "Differentiated services code point (DSCP) in the IP header of the probe packet.",
				Computed:    true,
			},
			"fail_weight": {
				Type:        schema.TypeInt,
				Description: "Threshold weight to trigger link failure alert.",
				Computed:    true,
			},
			"failtime": {
				Type:        schema.TypeInt,
				Description: "Number of retry attempts before the server is considered down (1 - 10, default = 5)",
				Computed:    true,
			},
			"gateway_ip": {
				Type:        schema.TypeString,
				Description: "Gateway IP address used to probe the server.",
				Computed:    true,
			},
			"gateway_ip6": {
				Type:        schema.TypeString,
				Description: "Gateway IPv6 address used to probe the server.",
				Computed:    true,
			},
			"ha_priority": {
				Type:        schema.TypeInt,
				Description: "HA election priority (1 - 50).",
				Computed:    true,
			},
			"http_agent": {
				Type:        schema.TypeString,
				Description: "String in the http-agent field in the HTTP header.",
				Computed:    true,
			},
			"http_get": {
				Type:        schema.TypeString,
				Description: "If you are monitoring an HTML server you can send an HTTP-GET request with a custom string. Use this option to define the string.",
				Computed:    true,
			},
			"http_match": {
				Type:        schema.TypeString,
				Description: "String that you expect to see in the HTTP-GET requests of the traffic to be monitored.",
				Computed:    true,
			},
			"interval": {
				Type:        schema.TypeInt,
				Description: "Detection interval in milliseconds (500 - 3600 * 1000 msec, default = 500).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Link monitor name.",
				Required:    true,
			},
			"packet_size": {
				Type:        schema.TypeInt,
				Description: "Packet size of a twamp test session,",
				Computed:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "Twamp controller password in authentication mode",
				Computed:    true,
				Sensitive:   true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "Port number of the traffic to be used to monitor the server.",
				Computed:    true,
			},
			"probe_count": {
				Type:        schema.TypeInt,
				Description: "Number of most recent probes that should be used to calculate latency and jitter (5 - 30, default = 30).",
				Computed:    true,
			},
			"probe_timeout": {
				Type:        schema.TypeInt,
				Description: "Time to wait before a probe packet is considered lost (500 - 5000 msec, default = 500).",
				Computed:    true,
			},
			"protocol": {
				Type:        schema.TypeString,
				Description: "Protocols used to monitor the server.",
				Computed:    true,
			},
			"recoverytime": {
				Type:        schema.TypeInt,
				Description: "Number of successful responses received before server is considered recovered (1 - 10, default = 5).",
				Computed:    true,
			},
			"route": {
				Type:        schema.TypeList,
				Description: "Subnet to monitor.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet": {
							Type:        schema.TypeString,
							Description: "IP and netmask (x.x.x.x/y).",
							Computed:    true,
						},
					},
				},
			},
			"security_mode": {
				Type:        schema.TypeString,
				Description: "Twamp controller security mode.",
				Computed:    true,
			},
			"server": {
				Type:        schema.TypeList,
				Description: "IP address of the server(s) to be monitored.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Description: "Server address.",
							Computed:    true,
						},
					},
				},
			},
			"server_config": {
				Type:        schema.TypeString,
				Description: "Mode of server configuration.",
				Computed:    true,
			},
			"server_list": {
				Type:        schema.TypeList,
				Description: "Servers for link-monitor to monitor.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst": {
							Type:        schema.TypeString,
							Description: "IP address of the server to be monitored.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Server ID.",
							Computed:    true,
						},
						"port": {
							Type:        schema.TypeInt,
							Description: "Port number of the traffic to be used to monitor the server.",
							Computed:    true,
						},
						"protocol": {
							Type:        schema.TypeString,
							Description: "Protocols used to monitor the server.",
							Computed:    true,
						},
						"weight": {
							Type:        schema.TypeInt,
							Description: "Weight of the monitor to this dst (0 - 255).",
							Computed:    true,
						},
					},
				},
			},
			"service_detection": {
				Type:        schema.TypeString,
				Description: "Only use monitor to read quality values. If enabled, static routes and cascade interfaces will not be updated.",
				Computed:    true,
			},
			"source_ip": {
				Type:        schema.TypeString,
				Description: "Source IP address used in packet to the server.",
				Computed:    true,
			},
			"source_ip6": {
				Type:        schema.TypeString,
				Description: "Source IPv6 address used in packet to the server.",
				Computed:    true,
			},
			"srcintf": {
				Type:        schema.TypeString,
				Description: "Interface that receives the traffic to be monitored.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable this link monitor.",
				Computed:    true,
			},
			"update_cascade_interface": {
				Type:        schema.TypeString,
				Description: "Enable/disable update cascade interface.",
				Computed:    true,
			},
			"update_policy_route": {
				Type:        schema.TypeString,
				Description: "Enable/disable updating the policy route.",
				Computed:    true,
			},
			"update_static_route": {
				Type:        schema.TypeString,
				Description: "Enable/disable updating the static route.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemLinkMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemLinkMonitor(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemLinkMonitor dataSource: %v", err)
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

	diags := refreshObjectSystemLinkMonitor(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
