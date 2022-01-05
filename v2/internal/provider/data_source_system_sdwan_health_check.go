// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemSdwanHealthCheck() *schema.Resource {
	return &schema.Resource{
		Description: "SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it.",

		ReadContext: dataSourceSystemSdwanHealthCheckRead,

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
			"detect_mode": {
				Type:        schema.TypeString,
				Description: "The mode determining how to detect the server.",
				Computed:    true,
			},
			"diffservcode": {
				Type:        schema.TypeString,
				Description: "Differentiated services code point (DSCP) in the IP header of the probe packet.",
				Computed:    true,
			},
			"dns_match_ip": {
				Type:        schema.TypeString,
				Description: "Response IP expected from DNS server if the protocol is DNS.",
				Computed:    true,
			},
			"dns_request_domain": {
				Type:        schema.TypeString,
				Description: "Fully qualified domain name to resolve for the DNS probe.",
				Computed:    true,
			},
			"failtime": {
				Type:        schema.TypeInt,
				Description: "Number of failures before server is considered lost (1 - 3600, default = 5).",
				Computed:    true,
			},
			"ftp_file": {
				Type:        schema.TypeString,
				Description: "Full path and file name on the FTP server to download for FTP health-check to probe.",
				Computed:    true,
			},
			"ftp_mode": {
				Type:        schema.TypeString,
				Description: "FTP mode.",
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
				Description: "URL used to communicate with the server if the protocol if the protocol is HTTP.",
				Computed:    true,
			},
			"http_match": {
				Type:        schema.TypeString,
				Description: "Response string expected from the server if the protocol is HTTP.",
				Computed:    true,
			},
			"interval": {
				Type:        schema.TypeInt,
				Description: "Status check interval in milliseconds, or the time between attempting to connect to the server (500 - 3600*1000 msec, default = 500).",
				Computed:    true,
			},
			"members": {
				Type:        schema.TypeList,
				Description: "Member sequence number list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seq_num": {
							Type:        schema.TypeInt,
							Description: "Member sequence number.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Status check or health check name.",
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
				Description: "Port number used to communicate with the server over the selected protocol (0-65535, default = 0, auto select. http, tcp-connect: 80, udp-echo, tcp-echo: 7, dns: 53, ftp: 21, twamp: 862).",
				Computed:    true,
			},
			"probe_count": {
				Type:        schema.TypeInt,
				Description: "Number of most recent probes that should be used to calculate latency and jitter (5 - 30, default = 30).",
				Computed:    true,
			},
			"probe_packets": {
				Type:        schema.TypeString,
				Description: "Enable/disable transmission of probe packets.",
				Computed:    true,
			},
			"probe_timeout": {
				Type:        schema.TypeInt,
				Description: "Time to wait before a probe packet is considered lost (500 - 3600*1000 msec, default = 500).",
				Computed:    true,
			},
			"protocol": {
				Type:        schema.TypeString,
				Description: "Protocol used to determine if the FortiGate can communicate with the server.",
				Computed:    true,
			},
			"quality_measured_method": {
				Type:        schema.TypeString,
				Description: "Method to measure the quality of tcp-connect.",
				Computed:    true,
			},
			"recoverytime": {
				Type:        schema.TypeInt,
				Description: "Number of successful responses received before server is considered recovered (1 - 3600, default = 5).",
				Computed:    true,
			},
			"security_mode": {
				Type:        schema.TypeString,
				Description: "Twamp controller security mode.",
				Computed:    true,
			},
			"server": {
				Type:        schema.TypeString,
				Description: "IP address or FQDN name of the server.",
				Computed:    true,
			},
			"sla": {
				Type:        schema.TypeList,
				Description: "Service level agreement (SLA).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "SLA ID.",
							Computed:    true,
						},
						"jitter_threshold": {
							Type:        schema.TypeInt,
							Description: "Jitter for SLA to make decision in milliseconds. (0 - 10000000, default = 5).",
							Computed:    true,
						},
						"latency_threshold": {
							Type:        schema.TypeInt,
							Description: "Latency for SLA to make decision in milliseconds. (0 - 10000000, default = 5).",
							Computed:    true,
						},
						"link_cost_factor": {
							Type:        schema.TypeString,
							Description: "Criteria on which to base link selection.",
							Computed:    true,
						},
						"packetloss_threshold": {
							Type:        schema.TypeInt,
							Description: "Packet loss for SLA to make decision in percentage. (0 - 100, default = 0).",
							Computed:    true,
						},
					},
				},
			},
			"sla_fail_log_period": {
				Type:        schema.TypeInt,
				Description: "Time interval in seconds that SLA fail log messages will be generated (0 - 3600, default = 0).",
				Computed:    true,
			},
			"sla_pass_log_period": {
				Type:        schema.TypeInt,
				Description: "Time interval in seconds that SLA pass log messages will be generated (0 - 3600, default = 0).",
				Computed:    true,
			},
			"system_dns": {
				Type:        schema.TypeString,
				Description: "Enable/disable system DNS as the probe server.",
				Computed:    true,
			},
			"threshold_alert_jitter": {
				Type:        schema.TypeInt,
				Description: "Alert threshold for jitter (ms, default = 0).",
				Computed:    true,
			},
			"threshold_alert_latency": {
				Type:        schema.TypeInt,
				Description: "Alert threshold for latency (ms, default = 0).",
				Computed:    true,
			},
			"threshold_alert_packetloss": {
				Type:        schema.TypeInt,
				Description: "Alert threshold for packet loss (percentage, default = 0).",
				Computed:    true,
			},
			"threshold_warning_jitter": {
				Type:        schema.TypeInt,
				Description: "Warning threshold for jitter (ms, default = 0).",
				Computed:    true,
			},
			"threshold_warning_latency": {
				Type:        schema.TypeInt,
				Description: "Warning threshold for latency (ms, default = 0).",
				Computed:    true,
			},
			"threshold_warning_packetloss": {
				Type:        schema.TypeInt,
				Description: "Warning threshold for packet loss (percentage, default = 0).",
				Computed:    true,
			},
			"update_cascade_interface": {
				Type:        schema.TypeString,
				Description: "Enable/disable update cascade interface.",
				Computed:    true,
			},
			"update_static_route": {
				Type:        schema.TypeString,
				Description: "Enable/disable updating the static route.",
				Computed:    true,
			},
			"user": {
				Type:        schema.TypeString,
				Description: "The user name to access probe server.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemSdwanHealthCheckRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemSdwanHealthCheck(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSdwanHealthCheck dataSource: %v", err)
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

	diags := refreshObjectSystemSdwanHealthCheck(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
