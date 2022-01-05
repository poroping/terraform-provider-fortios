// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure redundant internet connections using SD-WAN (formerly virtual WAN link).

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemVirtualWanLink() *schema.Resource {
	return &schema.Resource{
		Description: "Configure redundant internet connections using SD-WAN (formerly virtual WAN link).",

		ReadContext: dataSourceSystemVirtualWanLinkRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"fail_alert_interfaces": {
				Type:        schema.TypeList,
				Description: "Physical interfaces that will be alerted.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Physical interface name.",
							Computed:    true,
						},
					},
				},
			},
			"fail_detect": {
				Type:        schema.TypeString,
				Description: "Enable/disable SD-WAN Internet connection status checking (failure detection).",
				Computed:    true,
			},
			"health_check": {
				Type:        schema.TypeList,
				Description: "SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr_mode": {
							Type:        schema.TypeString,
							Description: "Address mode (IPv4 or IPv6).",
							Computed:    true,
						},
						"diffservcode": {
							Type:        schema.TypeString,
							Description: "Differentiated services code point (DSCP) in the IP header of the probe packet.",
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
							Computed:    true,
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
							Description: "Port number used to communicate with the server over the selected protocol.",
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
							Description: "Time to wait before a probe packet is considered lost (500 - 5000 msec, default = 500).",
							Computed:    true,
						},
						"protocol": {
							Type:        schema.TypeString,
							Description: "Protocol used to determine if the FortiGate can communicate with the server.",
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
					},
				},
			},
			"load_balance_mode": {
				Type:        schema.TypeString,
				Description: "Algorithm or mode to use for load balancing Internet traffic to SD-WAN members.",
				Computed:    true,
			},
			"members": {
				Type:        schema.TypeList,
				Description: "FortiGate interfaces added to the virtual-wan-link.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comment": {
							Type:        schema.TypeString,
							Description: "Comments.",
							Computed:    true,
						},
						"cost": {
							Type:        schema.TypeInt,
							Description: "Cost of this interface for services in SLA mode (0 - 4294967295, default = 0).",
							Computed:    true,
						},
						"gateway": {
							Type:        schema.TypeString,
							Description: "The default gateway for this interface. Usually the default gateway of the Internet service provider that this interface is connected to.",
							Computed:    true,
						},
						"gateway6": {
							Type:        schema.TypeString,
							Description: "IPv6 gateway.",
							Computed:    true,
						},
						"ingress_spillover_threshold": {
							Type:        schema.TypeInt,
							Description: "Ingress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.",
							Computed:    true,
						},
						"interface": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "Priority of the interface (0 - 4294967295). Used for SD-WAN rules or priority rules.",
							Computed:    true,
						},
						"seq_num": {
							Type:        schema.TypeInt,
							Description: "Sequence number(1-255).",
							Computed:    true,
						},
						"source": {
							Type:        schema.TypeString,
							Description: "Source IP address used in the health-check packet to the server.",
							Computed:    true,
						},
						"source6": {
							Type:        schema.TypeString,
							Description: "Source IPv6 address used in the health-check packet to the server.",
							Computed:    true,
						},
						"spillover_threshold": {
							Type:        schema.TypeInt,
							Description: "Egress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable this interface in the SD-WAN.",
							Computed:    true,
						},
						"volume_ratio": {
							Type:        schema.TypeInt,
							Description: "Measured volume ratio (this value / sum of all values = percentage of link volume, 1 - 255).",
							Computed:    true,
						},
						"weight": {
							Type:        schema.TypeInt,
							Description: "Weight of this interface for weighted load balancing. (1 - 255) More traffic is directed to interfaces with higher weights.",
							Computed:    true,
						},
					},
				},
			},
			"neighbor": {
				Type:        schema.TypeList,
				Description: "Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"health_check": {
							Type:        schema.TypeString,
							Description: "SD-WAN health-check name.",
							Computed:    true,
						},
						"ip": {
							Type:        schema.TypeString,
							Description: "IP/IPv6 address of neighbor.",
							Computed:    true,
						},
						"member": {
							Type:        schema.TypeInt,
							Description: "Member sequence number.",
							Computed:    true,
						},
						"role": {
							Type:        schema.TypeString,
							Description: "Role of neighbor.",
							Computed:    true,
						},
						"sla_id": {
							Type:        schema.TypeInt,
							Description: "SLA ID.",
							Computed:    true,
						},
					},
				},
			},
			"neighbor_hold_boot_time": {
				Type:        schema.TypeInt,
				Description: "Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).",
				Computed:    true,
			},
			"neighbor_hold_down": {
				Type:        schema.TypeString,
				Description: "Enable/disable hold switching from the secondary neighbor to the primary neighbor.",
				Computed:    true,
			},
			"neighbor_hold_down_time": {
				Type:        schema.TypeInt,
				Description: "Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).",
				Computed:    true,
			},
			"service": {
				Type:        schema.TypeList,
				Description: "Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr_mode": {
							Type:        schema.TypeString,
							Description: "Address mode (IPv4 or IPv6).",
							Computed:    true,
						},
						"bandwidth_weight": {
							Type:        schema.TypeInt,
							Description: "Coefficient of reciprocal of available bidirectional bandwidth in the formula of custom-profile-1.",
							Computed:    true,
						},
						"default": {
							Type:        schema.TypeString,
							Description: "Enable/disable use of SD-WAN as default service.",
							Computed:    true,
						},
						"dscp_forward": {
							Type:        schema.TypeString,
							Description: "Enable/disable forward traffic DSCP tag.",
							Computed:    true,
						},
						"dscp_forward_tag": {
							Type:        schema.TypeString,
							Description: "Forward traffic DSCP tag.",
							Computed:    true,
						},
						"dscp_reverse": {
							Type:        schema.TypeString,
							Description: "Enable/disable reverse traffic DSCP tag.",
							Computed:    true,
						},
						"dscp_reverse_tag": {
							Type:        schema.TypeString,
							Description: "Reverse traffic DSCP tag.",
							Computed:    true,
						},
						"dst": {
							Type:        schema.TypeList,
							Description: "Destination address name.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Address or address group name.",
										Computed:    true,
									},
								},
							},
						},
						"dst_negate": {
							Type:        schema.TypeString,
							Description: "Enable/disable negation of destination address match.",
							Computed:    true,
						},
						"dst6": {
							Type:        schema.TypeList,
							Description: "Destination address6 name.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Address6 or address6 group name.",
										Computed:    true,
									},
								},
							},
						},
						"end_port": {
							Type:        schema.TypeInt,
							Description: "End destination port number.",
							Computed:    true,
						},
						"gateway": {
							Type:        schema.TypeString,
							Description: "Enable/disable SD-WAN service gateway.",
							Computed:    true,
						},
						"groups": {
							Type:        schema.TypeList,
							Description: "User groups.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Group name.",
										Computed:    true,
									},
								},
							},
						},
						"health_check": {
							Type:        schema.TypeList,
							Description: "Health check list.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Health check name.",
										Computed:    true,
									},
								},
							},
						},
						"hold_down_time": {
							Type:        schema.TypeInt,
							Description: "Waiting period in seconds when switching from the back-up member to the primary member (0 - 10000000, default = 0).",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Priority rule ID (1 - 4000).",
							Computed:    true,
						},
						"input_device": {
							Type:        schema.TypeList,
							Description: "Source interface name.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Interface name.",
										Computed:    true,
									},
								},
							},
						},
						"input_device_negate": {
							Type:        schema.TypeString,
							Description: "Enable/disable negation of input device match.",
							Computed:    true,
						},
						"internet_service": {
							Type:        schema.TypeString,
							Description: "Enable/disable use of Internet service for application-based load balancing.",
							Computed:    true,
						},
						"internet_service_app_ctrl": {
							Type:        schema.TypeList,
							Description: "Application control based Internet Service ID list.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Description: "Application control based Internet Service ID.",
										Computed:    true,
									},
								},
							},
						},
						"internet_service_app_ctrl_group": {
							Type:        schema.TypeList,
							Description: "Application control based Internet Service group list.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Application control based Internet Service group name.",
										Computed:    true,
									},
								},
							},
						},
						"internet_service_custom": {
							Type:        schema.TypeList,
							Description: "Custom Internet service name list.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Custom Internet service name.",
										Computed:    true,
									},
								},
							},
						},
						"internet_service_custom_group": {
							Type:        schema.TypeList,
							Description: "Custom Internet Service group list.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Custom Internet Service group name.",
										Computed:    true,
									},
								},
							},
						},
						"internet_service_group": {
							Type:        schema.TypeList,
							Description: "Internet Service group list.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Internet Service group name.",
										Computed:    true,
									},
								},
							},
						},
						"internet_service_id": {
							Type:        schema.TypeList,
							Description: "Internet service ID list.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Description: "Internet service ID.",
										Computed:    true,
									},
								},
							},
						},
						"internet_service_name": {
							Type:        schema.TypeList,
							Description: "Internet service name list.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Internet service name.",
										Computed:    true,
									},
								},
							},
						},
						"jitter_weight": {
							Type:        schema.TypeInt,
							Description: "Coefficient of jitter in the formula of custom-profile-1.",
							Computed:    true,
						},
						"latency_weight": {
							Type:        schema.TypeInt,
							Description: "Coefficient of latency in the formula of custom-profile-1.",
							Computed:    true,
						},
						"link_cost_factor": {
							Type:        schema.TypeString,
							Description: "Link cost factor.",
							Computed:    true,
						},
						"link_cost_threshold": {
							Type:        schema.TypeInt,
							Description: "Percentage threshold change of link cost values that will result in policy route regeneration (0 - 10000000, default = 10).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "Control how the priority rule sets the priority of interfaces in the SD-WAN.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Priority rule name.",
							Computed:    true,
						},
						"packet_loss_weight": {
							Type:        schema.TypeInt,
							Description: "Coefficient of packet-loss in the formula of custom-profile-1.",
							Computed:    true,
						},
						"priority_members": {
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
						"protocol": {
							Type:        schema.TypeInt,
							Description: "Protocol number.",
							Computed:    true,
						},
						"quality_link": {
							Type:        schema.TypeInt,
							Description: "Quality grade.",
							Computed:    true,
						},
						"role": {
							Type:        schema.TypeString,
							Description: "Service role to work with neighbor.",
							Computed:    true,
						},
						"route_tag": {
							Type:        schema.TypeInt,
							Description: "IPv4 route map route-tag.",
							Computed:    true,
						},
						"sla": {
							Type:        schema.TypeList,
							Description: "Service level agreement (SLA).",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"health_check": {
										Type:        schema.TypeString,
										Description: "Virtual WAN Link health-check.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "SLA ID.",
										Computed:    true,
									},
								},
							},
						},
						"sla_compare_method": {
							Type:        schema.TypeString,
							Description: "Method to compare SLA value for SLA mode.",
							Computed:    true,
						},
						"src": {
							Type:        schema.TypeList,
							Description: "Source address name.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Address or address group name.",
										Computed:    true,
									},
								},
							},
						},
						"src_negate": {
							Type:        schema.TypeString,
							Description: "Enable/disable negation of source address match.",
							Computed:    true,
						},
						"src6": {
							Type:        schema.TypeList,
							Description: "Source address6 name.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Address6 or address6 group name.",
										Computed:    true,
									},
								},
							},
						},
						"standalone_action": {
							Type:        schema.TypeString,
							Description: "Enable/disable service when selected neighbor role is standalone while service role is not standalone.",
							Computed:    true,
						},
						"start_port": {
							Type:        schema.TypeInt,
							Description: "Start destination port number.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable SD-WAN service.",
							Computed:    true,
						},
						"tos": {
							Type:        schema.TypeString,
							Description: "Type of service bit pattern.",
							Computed:    true,
						},
						"tos_mask": {
							Type:        schema.TypeString,
							Description: "Type of service evaluated bits.",
							Computed:    true,
						},
						"users": {
							Type:        schema.TypeList,
							Description: "User name.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "User name.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable SD-WAN.",
				Computed:    true,
			},
			"zone": {
				Type:        schema.TypeList,
				Description: "Configure SD-WAN zones.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Zone name.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemVirtualWanLinkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemVirtualWanLink(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemVirtualWanLink dataSource: %v", err)
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

	diags := refreshObjectSystemVirtualWanLink(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
