// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure redundant internet connections using SD-WAN (formerly virtual WAN link).

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceSystemVirtualWanLink() *schema.Resource {
	return &schema.Resource{
		Description: "Configure redundant internet connections using SD-WAN (formerly virtual WAN link).",

		CreateContext: resourceSystemVirtualWanLinkCreate,
		ReadContext:   resourceSystemVirtualWanLinkRead,
		UpdateContext: resourceSystemVirtualWanLinkUpdate,
		DeleteContext: resourceSystemVirtualWanLinkDelete,

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
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"fail_alert_interfaces": {
				Type:        schema.TypeList,
				Description: "Physical interfaces that will be alerted.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Physical interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"fail_detect": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SD-WAN Internet connection status checking (failure detection).",
				Optional:    true,
				Computed:    true,
			},
			"health_check": {
				Type:        schema.TypeList,
				Description: "SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ipv4", "ipv6"}, false),

							Description: "Address mode (IPv4 or IPv6).",
							Optional:    true,
							Computed:    true,
						},
						"diffservcode": {
							Type: schema.TypeString,

							Description: "Differentiated services code point (DSCP) in the IP header of the probe packet.",
							Optional:    true,
							Computed:    true,
						},
						"dns_request_domain": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Fully qualified domain name to resolve for the DNS probe.",
							Optional:    true,
							Computed:    true,
						},
						"failtime": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 3600),

							Description: "Number of failures before server is considered lost (1 - 3600, default = 5).",
							Optional:    true,
							Computed:    true,
						},
						"ha_priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 50),

							Description: "HA election priority (1 - 50).",
							Optional:    true,
							Computed:    true,
						},
						"http_agent": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1024),

							Description: "String in the http-agent field in the HTTP header.",
							Optional:    true,
							Computed:    true,
						},
						"http_get": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1024),

							Description: "URL used to communicate with the server if the protocol if the protocol is HTTP.",
							Optional:    true,
							Computed:    true,
						},
						"http_match": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1024),

							Description: "Response string expected from the server if the protocol is HTTP.",
							Optional:    true,
							Computed:    true,
						},
						"interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(500, 3600000),

							Description: "Status check interval in milliseconds, or the time between attempting to connect to the server (500 - 3600*1000 msec, default = 500).",
							Optional:    true,
							Computed:    true,
						},
						"members": {
							Type:        schema.TypeList,
							Description: "Member sequence number list.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"seq_num": {
										Type: schema.TypeInt,

										Description: "Member sequence number.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Status check or health check name.",
							Optional:    true,
							Computed:    true,
						},
						"packet_size": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(64, 1024),

							Description: "Packet size of a twamp test session,",
							Optional:    true,
							Computed:    true,
						},
						"password": {
							Type: schema.TypeString,

							Description: "Twamp controller password in authentication mode",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Port number used to communicate with the server over the selected protocol.",
							Optional:    true,
							Computed:    true,
						},
						"probe_count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 30),

							Description: "Number of most recent probes that should be used to calculate latency and jitter (5 - 30, default = 30).",
							Optional:    true,
							Computed:    true,
						},
						"probe_packets": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable transmission of probe packets.",
							Optional:    true,
							Computed:    true,
						},
						"probe_timeout": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(500, 5000),

							Description: "Time to wait before a probe packet is considered lost (500 - 5000 msec, default = 500).",
							Optional:    true,
							Computed:    true,
						},
						"protocol": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ping", "tcp-echo", "udp-echo", "http", "twamp", "ping6", "dns"}, false),

							Description: "Protocol used to determine if the FortiGate can communicate with the server.",
							Optional:    true,
							Computed:    true,
						},
						"recoverytime": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 3600),

							Description: "Number of successful responses received before server is considered recovered (1 - 3600, default = 5).",
							Optional:    true,
							Computed:    true,
						},
						"security_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "authentication"}, false),

							Description: "Twamp controller security mode.",
							Optional:    true,
							Computed:    true,
						},
						"server": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "IP address or FQDN name of the server.",
							Optional:    true,
							Computed:    true,
						},
						"sla": {
							Type:        schema.TypeList,
							Description: "Service level agreement (SLA).",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 32),

										Description: "SLA ID.",
										Optional:    true,
										Computed:    true,
									},
									"jitter_threshold": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 10000000),

										Description: "Jitter for SLA to make decision in milliseconds. (0 - 10000000, default = 5).",
										Optional:    true,
										Computed:    true,
									},
									"latency_threshold": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 10000000),

										Description: "Latency for SLA to make decision in milliseconds. (0 - 10000000, default = 5).",
										Optional:    true,
										Computed:    true,
									},
									"link_cost_factor": {
										Type: schema.TypeString,

										DiffSuppressFunc: suppressors.DiffFakeListEqual,
										Description:      "Criteria on which to base link selection.",
										Optional:         true,
										Computed:         true,
									},
									"packetloss_threshold": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 100),

										Description: "Packet loss for SLA to make decision in percentage. (0 - 100, default = 0).",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"sla_fail_log_period": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 3600),

							Description: "Time interval in seconds that SLA fail log messages will be generated (0 - 3600, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"sla_pass_log_period": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 3600),

							Description: "Time interval in seconds that SLA pass log messages will be generated (0 - 3600, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"system_dns": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable system DNS as the probe server.",
							Optional:    true,
							Computed:    true,
						},
						"threshold_alert_jitter": {
							Type: schema.TypeInt,

							Description: "Alert threshold for jitter (ms, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"threshold_alert_latency": {
							Type: schema.TypeInt,

							Description: "Alert threshold for latency (ms, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"threshold_alert_packetloss": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),

							Description: "Alert threshold for packet loss (percentage, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"threshold_warning_jitter": {
							Type: schema.TypeInt,

							Description: "Warning threshold for jitter (ms, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"threshold_warning_latency": {
							Type: schema.TypeInt,

							Description: "Warning threshold for latency (ms, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"threshold_warning_packetloss": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),

							Description: "Warning threshold for packet loss (percentage, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"update_cascade_interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable update cascade interface.",
							Optional:    true,
							Computed:    true,
						},
						"update_static_route": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable updating the static route.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"load_balance_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"source-ip-based", "weight-based", "usage-based", "source-dest-ip-based", "measured-volume-based"}, false),

				Description: "Algorithm or mode to use for load balancing Internet traffic to SD-WAN members.",
				Optional:    true,
				Computed:    true,
			},
			"members": {
				Type:        schema.TypeList,
				Description: "FortiGate interfaces added to the virtual-wan-link.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comment": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Comments.",
							Optional:    true,
							Computed:    true,
						},
						"cost": {
							Type: schema.TypeInt,

							Description: "Cost of this interface for services in SLA mode (0 - 4294967295, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"gateway": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "The default gateway for this interface. Usually the default gateway of the Internet service provider that this interface is connected to.",
							Optional:    true,
							Computed:    true,
						},
						"gateway6": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "IPv6 gateway.",
							Optional:         true,
							Computed:         true,
						},
						"ingress_spillover_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 16776000),

							Description: "Ingress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.",
							Optional:    true,
							Computed:    true,
						},
						"interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
						"priority": {
							Type: schema.TypeInt,

							Description: "Priority of the interface (0 - 4294967295). Used for SD-WAN rules or priority rules.",
							Optional:    true,
							Computed:    true,
						},
						"seq_num": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Sequence number(1-255).",
							Optional:    true,
							Computed:    true,
						},
						"source": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Source IP address used in the health-check packet to the server.",
							Optional:    true,
							Computed:    true,
						},
						"source6": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "Source IPv6 address used in the health-check packet to the server.",
							Optional:         true,
							Computed:         true,
						},
						"spillover_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 16776000),

							Description: "Egress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable this interface in the SD-WAN.",
							Optional:    true,
							Computed:    true,
						},
						"volume_ratio": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),

							Description: "Measured volume ratio (this value / sum of all values = percentage of link volume, 1 - 255).",
							Optional:    true,
							Computed:    true,
						},
						"weight": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),

							Description: "Weight of this interface for weighted load balancing. (1 - 255) More traffic is directed to interfaces with higher weights.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"neighbor": {
				Type:        schema.TypeList,
				Description: "Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"health_check": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "SD-WAN health-check name.",
							Optional:    true,
							Computed:    true,
						},
						"ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 45),

							Description: "IP/IPv6 address of neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"member": {
							Type: schema.TypeInt,

							Description: "Member sequence number.",
							Optional:    true,
							Computed:    true,
						},
						"role": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"standalone", "primary", "secondary"}, false),

							Description: "Role of neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"sla_id": {
							Type: schema.TypeInt,

							Description: "SLA ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"neighbor_hold_boot_time": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10000000),

				Description: "Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"neighbor_hold_down": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable hold switching from the secondary neighbor to the primary neighbor.",
				Optional:    true,
				Computed:    true,
			},
			"neighbor_hold_down_time": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10000000),

				Description: "Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"service": {
				Type:        schema.TypeList,
				Description: "Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ipv4", "ipv6"}, false),

							Description: "Address mode (IPv4 or IPv6).",
							Optional:    true,
							Computed:    true,
						},
						"bandwidth_weight": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10000000),

							Description: "Coefficient of reciprocal of available bidirectional bandwidth in the formula of custom-profile-1.",
							Optional:    true,
							Computed:    true,
						},
						"default": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable use of SD-WAN as default service.",
							Optional:    true,
							Computed:    true,
						},
						"dscp_forward": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable forward traffic DSCP tag.",
							Optional:    true,
							Computed:    true,
						},
						"dscp_forward_tag": {
							Type: schema.TypeString,

							Description: "Forward traffic DSCP tag.",
							Optional:    true,
							Computed:    true,
						},
						"dscp_reverse": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable reverse traffic DSCP tag.",
							Optional:    true,
							Computed:    true,
						},
						"dscp_reverse_tag": {
							Type: schema.TypeString,

							Description: "Reverse traffic DSCP tag.",
							Optional:    true,
							Computed:    true,
						},
						"dst": {
							Type:        schema.TypeList,
							Description: "Destination address name.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Address or address group name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"dst_negate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable negation of destination address match.",
							Optional:    true,
							Computed:    true,
						},
						"dst6": {
							Type:        schema.TypeList,
							Description: "Destination address6 name.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Address6 or address6 group name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"end_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "End destination port number.",
							Optional:    true,
							Computed:    true,
						},
						"gateway": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable SD-WAN service gateway.",
							Optional:    true,
							Computed:    true,
						},
						"groups": {
							Type:        schema.TypeList,
							Description: "User groups.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Group name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"health_check": {
							Type:        schema.TypeList,
							Description: "Health check list.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Health check name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"hold_down_time": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10000000),

							Description: "Waiting period in seconds when switching from the back-up member to the primary member (0 - 10000000, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4000),

							Description: "Priority rule ID (1 - 4000).",
							Optional:    true,
							Computed:    true,
						},
						"input_device": {
							Type:        schema.TypeList,
							Description: "Source interface name.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Interface name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"input_device_negate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable negation of input device match.",
							Optional:    true,
							Computed:    true,
						},
						"internet_service": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable use of Internet service for application-based load balancing.",
							Optional:    true,
							Computed:    true,
						},
						"internet_service_app_ctrl": {
							Type:        schema.TypeList,
							Description: "Application control based Internet Service ID list.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type: schema.TypeInt,

										Description: "Application control based Internet Service ID.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"internet_service_app_ctrl_group": {
							Type:        schema.TypeList,
							Description: "Application control based Internet Service group list.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Application control based Internet Service group name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"internet_service_custom": {
							Type:        schema.TypeList,
							Description: "Custom Internet service name list.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Custom Internet service name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"internet_service_custom_group": {
							Type:        schema.TypeList,
							Description: "Custom Internet Service group list.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Custom Internet Service group name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"internet_service_group": {
							Type:        schema.TypeList,
							Description: "Internet Service group list.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Internet Service group name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"internet_service_id": {
							Type:        schema.TypeList,
							Description: "Internet service ID list.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type: schema.TypeInt,

										Description: "Internet service ID.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"internet_service_name": {
							Type:        schema.TypeList,
							Description: "Internet service name list.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Internet service name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"jitter_weight": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10000000),

							Description: "Coefficient of jitter in the formula of custom-profile-1.",
							Optional:    true,
							Computed:    true,
						},
						"latency_weight": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10000000),

							Description: "Coefficient of latency in the formula of custom-profile-1.",
							Optional:    true,
							Computed:    true,
						},
						"link_cost_factor": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"latency", "jitter", "packet-loss", "inbandwidth", "outbandwidth", "bibandwidth", "custom-profile-1"}, false),

							Description: "Link cost factor.",
							Optional:    true,
							Computed:    true,
						},
						"link_cost_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10000000),

							Description: "Percentage threshold change of link cost values that will result in policy route regeneration (0 - 10000000, default = 10).",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"auto", "manual", "priority", "sla", "load-balance"}, false),

							Description: "Control how the priority rule sets the priority of interfaces in the SD-WAN.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Priority rule name.",
							Optional:    true,
							Computed:    true,
						},
						"packet_loss_weight": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10000000),

							Description: "Coefficient of packet-loss in the formula of custom-profile-1.",
							Optional:    true,
							Computed:    true,
						},
						"priority_members": {
							Type:        schema.TypeList,
							Description: "Member sequence number list.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"seq_num": {
										Type: schema.TypeInt,

										Description: "Member sequence number.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"protocol": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Protocol number.",
							Optional:    true,
							Computed:    true,
						},
						"quality_link": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Quality grade.",
							Optional:    true,
							Computed:    true,
						},
						"role": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"standalone", "primary", "secondary"}, false),

							Description: "Service role to work with neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"route_tag": {
							Type: schema.TypeInt,

							Description: "IPv4 route map route-tag.",
							Optional:    true,
							Computed:    true,
						},
						"sla": {
							Type:        schema.TypeList,
							Description: "Service level agreement (SLA).",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"health_check": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Virtual WAN Link health-check.",
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Type: schema.TypeInt,

										Description: "SLA ID.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"sla_compare_method": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"order", "number"}, false),

							Description: "Method to compare SLA value for SLA mode.",
							Optional:    true,
							Computed:    true,
						},
						"src": {
							Type:        schema.TypeList,
							Description: "Source address name.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Address or address group name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"src_negate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable negation of source address match.",
							Optional:    true,
							Computed:    true,
						},
						"src6": {
							Type:        schema.TypeList,
							Description: "Source address6 name.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Address6 or address6 group name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"standalone_action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable service when selected neighbor role is standalone while service role is not standalone.",
							Optional:    true,
							Computed:    true,
						},
						"start_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Start destination port number.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable SD-WAN service.",
							Optional:    true,
							Computed:    true,
						},
						"tos": {
							Type: schema.TypeString,

							Description: "Type of service bit pattern.",
							Optional:    true,
							Computed:    true,
						},
						"tos_mask": {
							Type: schema.TypeString,

							Description: "Type of service evaluated bits.",
							Optional:    true,
							Computed:    true,
						},
						"users": {
							Type:        schema.TypeList,
							Description: "User name.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "User name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable SD-WAN.",
				Optional:    true,
				Computed:    true,
			},
			"zone": {
				Type:        schema.TypeList,
				Description: "Configure SD-WAN zones.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Zone name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemVirtualWanLinkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemVirtualWanLink(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemVirtualWanLink(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemVirtualWanLink")
	}

	return resourceSystemVirtualWanLinkRead(ctx, d, meta)
}

func resourceSystemVirtualWanLinkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemVirtualWanLink(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemVirtualWanLink(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemVirtualWanLink resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemVirtualWanLink")
	}

	return resourceSystemVirtualWanLinkRead(ctx, d, meta)
}

func resourceSystemVirtualWanLinkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemVirtualWanLink(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemVirtualWanLink(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemVirtualWanLink resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVirtualWanLinkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemVirtualWanLink(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemVirtualWanLink resource: %v", err)
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

	diags := refreshObjectSystemVirtualWanLink(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemVirtualWanLinkFailAlertInterfaces(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkFailAlertInterfaces, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemVirtualWanLinkHealthCheck(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkHealthCheck, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AddrMode; tmp != nil {
				v["addr_mode"] = *tmp
			}

			if tmp := cfg.Diffservcode; tmp != nil {
				v["diffservcode"] = *tmp
			}

			if tmp := cfg.DnsRequestDomain; tmp != nil {
				v["dns_request_domain"] = *tmp
			}

			if tmp := cfg.Failtime; tmp != nil {
				v["failtime"] = *tmp
			}

			if tmp := cfg.HaPriority; tmp != nil {
				v["ha_priority"] = *tmp
			}

			if tmp := cfg.HttpAgent; tmp != nil {
				v["http_agent"] = *tmp
			}

			if tmp := cfg.HttpGet; tmp != nil {
				v["http_get"] = *tmp
			}

			if tmp := cfg.HttpMatch; tmp != nil {
				v["http_match"] = *tmp
			}

			if tmp := cfg.Interval; tmp != nil {
				v["interval"] = *tmp
			}

			if tmp := cfg.Members; tmp != nil {
				v["members"] = flattenSystemVirtualWanLinkHealthCheckMembers(d, tmp, prefix+"members", sort)
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.PacketSize; tmp != nil {
				v["packet_size"] = *tmp
			}

			if tmp := cfg.Password; tmp != nil {
				v["password"] = *tmp
			}

			if tmp := cfg.Port; tmp != nil {
				v["port"] = *tmp
			}

			if tmp := cfg.ProbeCount; tmp != nil {
				v["probe_count"] = *tmp
			}

			if tmp := cfg.ProbePackets; tmp != nil {
				v["probe_packets"] = *tmp
			}

			if tmp := cfg.ProbeTimeout; tmp != nil {
				v["probe_timeout"] = *tmp
			}

			if tmp := cfg.Protocol; tmp != nil {
				v["protocol"] = *tmp
			}

			if tmp := cfg.Recoverytime; tmp != nil {
				v["recoverytime"] = *tmp
			}

			if tmp := cfg.SecurityMode; tmp != nil {
				v["security_mode"] = *tmp
			}

			if tmp := cfg.Server; tmp != nil {
				v["server"] = *tmp
			}

			if tmp := cfg.Sla; tmp != nil {
				v["sla"] = flattenSystemVirtualWanLinkHealthCheckSla(d, tmp, prefix+"sla", sort)
			}

			if tmp := cfg.SlaFailLogPeriod; tmp != nil {
				v["sla_fail_log_period"] = *tmp
			}

			if tmp := cfg.SlaPassLogPeriod; tmp != nil {
				v["sla_pass_log_period"] = *tmp
			}

			if tmp := cfg.SystemDns; tmp != nil {
				v["system_dns"] = *tmp
			}

			if tmp := cfg.ThresholdAlertJitter; tmp != nil {
				v["threshold_alert_jitter"] = *tmp
			}

			if tmp := cfg.ThresholdAlertLatency; tmp != nil {
				v["threshold_alert_latency"] = *tmp
			}

			if tmp := cfg.ThresholdAlertPacketloss; tmp != nil {
				v["threshold_alert_packetloss"] = *tmp
			}

			if tmp := cfg.ThresholdWarningJitter; tmp != nil {
				v["threshold_warning_jitter"] = *tmp
			}

			if tmp := cfg.ThresholdWarningLatency; tmp != nil {
				v["threshold_warning_latency"] = *tmp
			}

			if tmp := cfg.ThresholdWarningPacketloss; tmp != nil {
				v["threshold_warning_packetloss"] = *tmp
			}

			if tmp := cfg.UpdateCascadeInterface; tmp != nil {
				v["update_cascade_interface"] = *tmp
			}

			if tmp := cfg.UpdateStaticRoute; tmp != nil {
				v["update_static_route"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemVirtualWanLinkHealthCheckMembers(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkHealthCheckMembers, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.SeqNum; tmp != nil {
				v["seq_num"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "seq_num")
	}

	return flat
}

func flattenSystemVirtualWanLinkHealthCheckSla(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkHealthCheckSla, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.JitterThreshold; tmp != nil {
				v["jitter_threshold"] = *tmp
			}

			if tmp := cfg.LatencyThreshold; tmp != nil {
				v["latency_threshold"] = *tmp
			}

			if tmp := cfg.LinkCostFactor; tmp != nil {
				v["link_cost_factor"] = *tmp
			}

			if tmp := cfg.PacketlossThreshold; tmp != nil {
				v["packetloss_threshold"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemVirtualWanLinkMembers(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkMembers, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Comment; tmp != nil {
				v["comment"] = *tmp
			}

			if tmp := cfg.Cost; tmp != nil {
				v["cost"] = *tmp
			}

			if tmp := cfg.Gateway; tmp != nil {
				v["gateway"] = *tmp
			}

			if tmp := cfg.Gateway6; tmp != nil {
				v["gateway6"] = *tmp
			}

			if tmp := cfg.IngressSpilloverThreshold; tmp != nil {
				v["ingress_spillover_threshold"] = *tmp
			}

			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			if tmp := cfg.SeqNum; tmp != nil {
				v["seq_num"] = *tmp
			}

			if tmp := cfg.Source; tmp != nil {
				v["source"] = *tmp
			}

			if tmp := cfg.Source6; tmp != nil {
				v["source6"] = *tmp
			}

			if tmp := cfg.SpilloverThreshold; tmp != nil {
				v["spillover_threshold"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.VolumeRatio; tmp != nil {
				v["volume_ratio"] = *tmp
			}

			if tmp := cfg.Weight; tmp != nil {
				v["weight"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "seq_num")
	}

	return flat
}

func flattenSystemVirtualWanLinkNeighbor(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkNeighbor, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.HealthCheck; tmp != nil {
				v["health_check"] = *tmp
			}

			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			if tmp := cfg.Member; tmp != nil {
				v["member"] = *tmp
			}

			if tmp := cfg.Role; tmp != nil {
				v["role"] = *tmp
			}

			if tmp := cfg.SlaId; tmp != nil {
				v["sla_id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "ip")
	}

	return flat
}

func flattenSystemVirtualWanLinkService(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkService, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AddrMode; tmp != nil {
				v["addr_mode"] = *tmp
			}

			if tmp := cfg.BandwidthWeight; tmp != nil {
				v["bandwidth_weight"] = *tmp
			}

			if tmp := cfg.Default; tmp != nil {
				v["default"] = *tmp
			}

			if tmp := cfg.DscpForward; tmp != nil {
				v["dscp_forward"] = *tmp
			}

			if tmp := cfg.DscpForwardTag; tmp != nil {
				v["dscp_forward_tag"] = *tmp
			}

			if tmp := cfg.DscpReverse; tmp != nil {
				v["dscp_reverse"] = *tmp
			}

			if tmp := cfg.DscpReverseTag; tmp != nil {
				v["dscp_reverse_tag"] = *tmp
			}

			if tmp := cfg.Dst; tmp != nil {
				v["dst"] = flattenSystemVirtualWanLinkServiceDst(d, tmp, prefix+"dst", sort)
			}

			if tmp := cfg.DstNegate; tmp != nil {
				v["dst_negate"] = *tmp
			}

			if tmp := cfg.Dst6; tmp != nil {
				v["dst6"] = flattenSystemVirtualWanLinkServiceDst6(d, tmp, prefix+"dst6", sort)
			}

			if tmp := cfg.EndPort; tmp != nil {
				v["end_port"] = *tmp
			}

			if tmp := cfg.Gateway; tmp != nil {
				v["gateway"] = *tmp
			}

			if tmp := cfg.Groups; tmp != nil {
				v["groups"] = flattenSystemVirtualWanLinkServiceGroups(d, tmp, prefix+"groups", sort)
			}

			if tmp := cfg.HealthCheck; tmp != nil {
				v["health_check"] = flattenSystemVirtualWanLinkServiceHealthCheck(d, tmp, prefix+"health_check", sort)
			}

			if tmp := cfg.HoldDownTime; tmp != nil {
				v["hold_down_time"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.InputDevice; tmp != nil {
				v["input_device"] = flattenSystemVirtualWanLinkServiceInputDevice(d, tmp, prefix+"input_device", sort)
			}

			if tmp := cfg.InputDeviceNegate; tmp != nil {
				v["input_device_negate"] = *tmp
			}

			if tmp := cfg.InternetService; tmp != nil {
				v["internet_service"] = *tmp
			}

			if tmp := cfg.InternetServiceAppCtrl; tmp != nil {
				v["internet_service_app_ctrl"] = flattenSystemVirtualWanLinkServiceInternetServiceAppCtrl(d, tmp, prefix+"internet_service_app_ctrl", sort)
			}

			if tmp := cfg.InternetServiceAppCtrlGroup; tmp != nil {
				v["internet_service_app_ctrl_group"] = flattenSystemVirtualWanLinkServiceInternetServiceAppCtrlGroup(d, tmp, prefix+"internet_service_app_ctrl_group", sort)
			}

			if tmp := cfg.InternetServiceCustom; tmp != nil {
				v["internet_service_custom"] = flattenSystemVirtualWanLinkServiceInternetServiceCustom(d, tmp, prefix+"internet_service_custom", sort)
			}

			if tmp := cfg.InternetServiceCustomGroup; tmp != nil {
				v["internet_service_custom_group"] = flattenSystemVirtualWanLinkServiceInternetServiceCustomGroup(d, tmp, prefix+"internet_service_custom_group", sort)
			}

			if tmp := cfg.InternetServiceGroup; tmp != nil {
				v["internet_service_group"] = flattenSystemVirtualWanLinkServiceInternetServiceGroup(d, tmp, prefix+"internet_service_group", sort)
			}

			if tmp := cfg.InternetServiceId; tmp != nil {
				v["internet_service_id"] = flattenSystemVirtualWanLinkServiceInternetServiceId(d, tmp, prefix+"internet_service_id", sort)
			}

			if tmp := cfg.InternetServiceName; tmp != nil {
				v["internet_service_name"] = flattenSystemVirtualWanLinkServiceInternetServiceName(d, tmp, prefix+"internet_service_name", sort)
			}

			if tmp := cfg.JitterWeight; tmp != nil {
				v["jitter_weight"] = *tmp
			}

			if tmp := cfg.LatencyWeight; tmp != nil {
				v["latency_weight"] = *tmp
			}

			if tmp := cfg.LinkCostFactor; tmp != nil {
				v["link_cost_factor"] = *tmp
			}

			if tmp := cfg.LinkCostThreshold; tmp != nil {
				v["link_cost_threshold"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.PacketLossWeight; tmp != nil {
				v["packet_loss_weight"] = *tmp
			}

			if tmp := cfg.PriorityMembers; tmp != nil {
				v["priority_members"] = flattenSystemVirtualWanLinkServicePriorityMembers(d, tmp, prefix+"priority_members", sort)
			}

			if tmp := cfg.Protocol; tmp != nil {
				v["protocol"] = *tmp
			}

			if tmp := cfg.QualityLink; tmp != nil {
				v["quality_link"] = *tmp
			}

			if tmp := cfg.Role; tmp != nil {
				v["role"] = *tmp
			}

			if tmp := cfg.RouteTag; tmp != nil {
				v["route_tag"] = *tmp
			}

			if tmp := cfg.Sla; tmp != nil {
				v["sla"] = flattenSystemVirtualWanLinkServiceSla(d, tmp, prefix+"sla", sort)
			}

			if tmp := cfg.SlaCompareMethod; tmp != nil {
				v["sla_compare_method"] = *tmp
			}

			if tmp := cfg.Src; tmp != nil {
				v["src"] = flattenSystemVirtualWanLinkServiceSrc(d, tmp, prefix+"src", sort)
			}

			if tmp := cfg.SrcNegate; tmp != nil {
				v["src_negate"] = *tmp
			}

			if tmp := cfg.Src6; tmp != nil {
				v["src6"] = flattenSystemVirtualWanLinkServiceSrc6(d, tmp, prefix+"src6", sort)
			}

			if tmp := cfg.StandaloneAction; tmp != nil {
				v["standalone_action"] = *tmp
			}

			if tmp := cfg.StartPort; tmp != nil {
				v["start_port"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Tos; tmp != nil {
				v["tos"] = *tmp
			}

			if tmp := cfg.TosMask; tmp != nil {
				v["tos_mask"] = *tmp
			}

			if tmp := cfg.Users; tmp != nil {
				v["users"] = flattenSystemVirtualWanLinkServiceUsers(d, tmp, prefix+"users", sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemVirtualWanLinkServiceDst(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkServiceDst, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemVirtualWanLinkServiceDst6(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkServiceDst6, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemVirtualWanLinkServiceGroups(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkServiceGroups, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemVirtualWanLinkServiceHealthCheck(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkServiceHealthCheck, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemVirtualWanLinkServiceInputDevice(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkServiceInputDevice, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemVirtualWanLinkServiceInternetServiceAppCtrl(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkServiceInternetServiceAppCtrl, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemVirtualWanLinkServiceInternetServiceAppCtrlGroup(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkServiceInternetServiceAppCtrlGroup, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemVirtualWanLinkServiceInternetServiceCustom(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkServiceInternetServiceCustom, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemVirtualWanLinkServiceInternetServiceCustomGroup(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkServiceInternetServiceCustomGroup, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemVirtualWanLinkServiceInternetServiceGroup(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkServiceInternetServiceGroup, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemVirtualWanLinkServiceInternetServiceId(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkServiceInternetServiceId, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemVirtualWanLinkServiceInternetServiceName(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkServiceInternetServiceName, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemVirtualWanLinkServicePriorityMembers(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkServicePriorityMembers, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.SeqNum; tmp != nil {
				v["seq_num"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "seq_num")
	}

	return flat
}

func flattenSystemVirtualWanLinkServiceSla(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkServiceSla, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.HealthCheck; tmp != nil {
				v["health_check"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "health_check")
	}

	return flat
}

func flattenSystemVirtualWanLinkServiceSrc(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkServiceSrc, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemVirtualWanLinkServiceSrc6(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkServiceSrc6, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemVirtualWanLinkServiceUsers(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkServiceUsers, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemVirtualWanLinkZone(d *schema.ResourceData, v *[]models.SystemVirtualWanLinkZone, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectSystemVirtualWanLink(d *schema.ResourceData, o *models.SystemVirtualWanLink, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.FailAlertInterfaces != nil {
		if err = d.Set("fail_alert_interfaces", flattenSystemVirtualWanLinkFailAlertInterfaces(d, o.FailAlertInterfaces, "fail_alert_interfaces", sort)); err != nil {
			return diag.Errorf("error reading fail_alert_interfaces: %v", err)
		}
	}

	if o.FailDetect != nil {
		v := *o.FailDetect

		if err = d.Set("fail_detect", v); err != nil {
			return diag.Errorf("error reading fail_detect: %v", err)
		}
	}

	if o.HealthCheck != nil {
		if err = d.Set("health_check", flattenSystemVirtualWanLinkHealthCheck(d, o.HealthCheck, "health_check", sort)); err != nil {
			return diag.Errorf("error reading health_check: %v", err)
		}
	}

	if o.LoadBalanceMode != nil {
		v := *o.LoadBalanceMode

		if err = d.Set("load_balance_mode", v); err != nil {
			return diag.Errorf("error reading load_balance_mode: %v", err)
		}
	}

	if o.Members != nil {
		if err = d.Set("members", flattenSystemVirtualWanLinkMembers(d, o.Members, "members", sort)); err != nil {
			return diag.Errorf("error reading members: %v", err)
		}
	}

	if o.Neighbor != nil {
		if err = d.Set("neighbor", flattenSystemVirtualWanLinkNeighbor(d, o.Neighbor, "neighbor", sort)); err != nil {
			return diag.Errorf("error reading neighbor: %v", err)
		}
	}

	if o.NeighborHoldBootTime != nil {
		v := *o.NeighborHoldBootTime

		if err = d.Set("neighbor_hold_boot_time", v); err != nil {
			return diag.Errorf("error reading neighbor_hold_boot_time: %v", err)
		}
	}

	if o.NeighborHoldDown != nil {
		v := *o.NeighborHoldDown

		if err = d.Set("neighbor_hold_down", v); err != nil {
			return diag.Errorf("error reading neighbor_hold_down: %v", err)
		}
	}

	if o.NeighborHoldDownTime != nil {
		v := *o.NeighborHoldDownTime

		if err = d.Set("neighbor_hold_down_time", v); err != nil {
			return diag.Errorf("error reading neighbor_hold_down_time: %v", err)
		}
	}

	if o.Service != nil {
		if err = d.Set("service", flattenSystemVirtualWanLinkService(d, o.Service, "service", sort)); err != nil {
			return diag.Errorf("error reading service: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Zone != nil {
		if err = d.Set("zone", flattenSystemVirtualWanLinkZone(d, o.Zone, "zone", sort)); err != nil {
			return diag.Errorf("error reading zone: %v", err)
		}
	}

	return nil
}

func expandSystemVirtualWanLinkFailAlertInterfaces(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkFailAlertInterfaces, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkFailAlertInterfaces

	for i := range l {
		tmp := models.SystemVirtualWanLinkFailAlertInterfaces{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkHealthCheck, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkHealthCheck

	for i := range l {
		tmp := models.SystemVirtualWanLinkHealthCheck{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.addr_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AddrMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.diffservcode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Diffservcode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dns_request_domain", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DnsRequestDomain = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.failtime", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Failtime = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ha_priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.HaPriority = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.http_agent", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HttpAgent = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.http_get", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HttpGet = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.http_match", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HttpMatch = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Interval = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.members", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemVirtualWanLinkHealthCheckMembers(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemVirtualWanLinkHealthCheckMembers
			// 	}
			tmp.Members = v2

		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.packet_size", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.PacketSize = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.password", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Password = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Port = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.probe_count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ProbeCount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.probe_packets", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ProbePackets = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.probe_timeout", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ProbeTimeout = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Protocol = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.recoverytime", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Recoverytime = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.security_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SecurityMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.server", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Server = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sla", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemVirtualWanLinkHealthCheckSla(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemVirtualWanLinkHealthCheckSla
			// 	}
			tmp.Sla = v2

		}

		pre_append = fmt.Sprintf("%s.%d.sla_fail_log_period", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SlaFailLogPeriod = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sla_pass_log_period", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SlaPassLogPeriod = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.system_dns", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SystemDns = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.threshold_alert_jitter", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ThresholdAlertJitter = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.threshold_alert_latency", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ThresholdAlertLatency = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.threshold_alert_packetloss", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ThresholdAlertPacketloss = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.threshold_warning_jitter", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ThresholdWarningJitter = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.threshold_warning_latency", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ThresholdWarningLatency = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.threshold_warning_packetloss", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ThresholdWarningPacketloss = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.update_cascade_interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UpdateCascadeInterface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.update_static_route", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UpdateStaticRoute = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkHealthCheckMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkHealthCheckMembers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkHealthCheckMembers

	for i := range l {
		tmp := models.SystemVirtualWanLinkHealthCheckMembers{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.seq_num", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SeqNum = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkHealthCheckSla(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkHealthCheckSla, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkHealthCheckSla

	for i := range l {
		tmp := models.SystemVirtualWanLinkHealthCheckSla{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.jitter_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.JitterThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.latency_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.LatencyThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.link_cost_factor", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LinkCostFactor = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.packetloss_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.PacketlossThreshold = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkMembers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkMembers

	for i := range l {
		tmp := models.SystemVirtualWanLinkMembers{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.comment", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Comment = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cost", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Cost = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.gateway", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Gateway = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.gateway6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Gateway6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ingress_spillover_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.IngressSpilloverThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Priority = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.seq_num", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SeqNum = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.source", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Source = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.source6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Source6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.spillover_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SpilloverThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.volume_ratio", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.VolumeRatio = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.weight", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Weight = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkNeighbor(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkNeighbor, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkNeighbor

	for i := range l {
		tmp := models.SystemVirtualWanLinkNeighbor{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.health_check", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HealthCheck = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.member", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Member = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.role", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Role = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sla_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SlaId = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkService(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkService, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkService

	for i := range l {
		tmp := models.SystemVirtualWanLinkService{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.addr_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AddrMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bandwidth_weight", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.BandwidthWeight = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.default", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Default = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp_forward", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DscpForward = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp_forward_tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DscpForwardTag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp_reverse", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DscpReverse = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp_reverse_tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DscpReverseTag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dst", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemVirtualWanLinkServiceDst(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemVirtualWanLinkServiceDst
			// 	}
			tmp.Dst = v2

		}

		pre_append = fmt.Sprintf("%s.%d.dst_negate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DstNegate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dst6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemVirtualWanLinkServiceDst6(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemVirtualWanLinkServiceDst6
			// 	}
			tmp.Dst6 = v2

		}

		pre_append = fmt.Sprintf("%s.%d.end_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.EndPort = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.gateway", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Gateway = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.groups", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemVirtualWanLinkServiceGroups(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemVirtualWanLinkServiceGroups
			// 	}
			tmp.Groups = v2

		}

		pre_append = fmt.Sprintf("%s.%d.health_check", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemVirtualWanLinkServiceHealthCheck(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemVirtualWanLinkServiceHealthCheck
			// 	}
			tmp.HealthCheck = v2

		}

		pre_append = fmt.Sprintf("%s.%d.hold_down_time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.HoldDownTime = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.input_device", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemVirtualWanLinkServiceInputDevice(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemVirtualWanLinkServiceInputDevice
			// 	}
			tmp.InputDevice = v2

		}

		pre_append = fmt.Sprintf("%s.%d.input_device_negate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InputDeviceNegate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.internet_service", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InternetService = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.internet_service_app_ctrl", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemVirtualWanLinkServiceInternetServiceAppCtrl(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemVirtualWanLinkServiceInternetServiceAppCtrl
			// 	}
			tmp.InternetServiceAppCtrl = v2

		}

		pre_append = fmt.Sprintf("%s.%d.internet_service_app_ctrl_group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemVirtualWanLinkServiceInternetServiceAppCtrlGroup(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemVirtualWanLinkServiceInternetServiceAppCtrlGroup
			// 	}
			tmp.InternetServiceAppCtrlGroup = v2

		}

		pre_append = fmt.Sprintf("%s.%d.internet_service_custom", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemVirtualWanLinkServiceInternetServiceCustom(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemVirtualWanLinkServiceInternetServiceCustom
			// 	}
			tmp.InternetServiceCustom = v2

		}

		pre_append = fmt.Sprintf("%s.%d.internet_service_custom_group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemVirtualWanLinkServiceInternetServiceCustomGroup(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemVirtualWanLinkServiceInternetServiceCustomGroup
			// 	}
			tmp.InternetServiceCustomGroup = v2

		}

		pre_append = fmt.Sprintf("%s.%d.internet_service_group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemVirtualWanLinkServiceInternetServiceGroup(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemVirtualWanLinkServiceInternetServiceGroup
			// 	}
			tmp.InternetServiceGroup = v2

		}

		pre_append = fmt.Sprintf("%s.%d.internet_service_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemVirtualWanLinkServiceInternetServiceId(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemVirtualWanLinkServiceInternetServiceId
			// 	}
			tmp.InternetServiceId = v2

		}

		pre_append = fmt.Sprintf("%s.%d.internet_service_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemVirtualWanLinkServiceInternetServiceName(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemVirtualWanLinkServiceInternetServiceName
			// 	}
			tmp.InternetServiceName = v2

		}

		pre_append = fmt.Sprintf("%s.%d.jitter_weight", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.JitterWeight = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.latency_weight", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.LatencyWeight = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.link_cost_factor", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LinkCostFactor = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.link_cost_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.LinkCostThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.packet_loss_weight", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.PacketLossWeight = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority_members", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemVirtualWanLinkServicePriorityMembers(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemVirtualWanLinkServicePriorityMembers
			// 	}
			tmp.PriorityMembers = v2

		}

		pre_append = fmt.Sprintf("%s.%d.protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Protocol = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quality_link", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.QualityLink = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.role", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Role = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.RouteTag = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sla", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemVirtualWanLinkServiceSla(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemVirtualWanLinkServiceSla
			// 	}
			tmp.Sla = v2

		}

		pre_append = fmt.Sprintf("%s.%d.sla_compare_method", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SlaCompareMethod = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.src", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemVirtualWanLinkServiceSrc(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemVirtualWanLinkServiceSrc
			// 	}
			tmp.Src = v2

		}

		pre_append = fmt.Sprintf("%s.%d.src_negate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SrcNegate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.src6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemVirtualWanLinkServiceSrc6(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemVirtualWanLinkServiceSrc6
			// 	}
			tmp.Src6 = v2

		}

		pre_append = fmt.Sprintf("%s.%d.standalone_action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StandaloneAction = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.start_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.StartPort = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tos", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Tos = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tos_mask", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TosMask = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.users", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemVirtualWanLinkServiceUsers(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemVirtualWanLinkServiceUsers
			// 	}
			tmp.Users = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkServiceDst(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkServiceDst, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkServiceDst

	for i := range l {
		tmp := models.SystemVirtualWanLinkServiceDst{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkServiceDst6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkServiceDst6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkServiceDst6

	for i := range l {
		tmp := models.SystemVirtualWanLinkServiceDst6{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkServiceGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkServiceGroups, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkServiceGroups

	for i := range l {
		tmp := models.SystemVirtualWanLinkServiceGroups{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkServiceHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkServiceHealthCheck, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkServiceHealthCheck

	for i := range l {
		tmp := models.SystemVirtualWanLinkServiceHealthCheck{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkServiceInputDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkServiceInputDevice, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkServiceInputDevice

	for i := range l {
		tmp := models.SystemVirtualWanLinkServiceInputDevice{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceAppCtrl(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkServiceInternetServiceAppCtrl, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkServiceInternetServiceAppCtrl

	for i := range l {
		tmp := models.SystemVirtualWanLinkServiceInternetServiceAppCtrl{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceAppCtrlGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkServiceInternetServiceAppCtrlGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkServiceInternetServiceAppCtrlGroup

	for i := range l {
		tmp := models.SystemVirtualWanLinkServiceInternetServiceAppCtrlGroup{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkServiceInternetServiceCustom, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkServiceInternetServiceCustom

	for i := range l {
		tmp := models.SystemVirtualWanLinkServiceInternetServiceCustom{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkServiceInternetServiceCustomGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkServiceInternetServiceCustomGroup

	for i := range l {
		tmp := models.SystemVirtualWanLinkServiceInternetServiceCustomGroup{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkServiceInternetServiceGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkServiceInternetServiceGroup

	for i := range l {
		tmp := models.SystemVirtualWanLinkServiceInternetServiceGroup{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceId(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkServiceInternetServiceId, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkServiceInternetServiceId

	for i := range l {
		tmp := models.SystemVirtualWanLinkServiceInternetServiceId{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkServiceInternetServiceName, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkServiceInternetServiceName

	for i := range l {
		tmp := models.SystemVirtualWanLinkServiceInternetServiceName{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkServicePriorityMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkServicePriorityMembers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkServicePriorityMembers

	for i := range l {
		tmp := models.SystemVirtualWanLinkServicePriorityMembers{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.seq_num", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SeqNum = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkServiceSla(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkServiceSla, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkServiceSla

	for i := range l {
		tmp := models.SystemVirtualWanLinkServiceSla{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.health_check", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HealthCheck = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkServiceSrc(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkServiceSrc, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkServiceSrc

	for i := range l {
		tmp := models.SystemVirtualWanLinkServiceSrc{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkServiceSrc6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkServiceSrc6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkServiceSrc6

	for i := range l {
		tmp := models.SystemVirtualWanLinkServiceSrc6{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkServiceUsers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkServiceUsers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkServiceUsers

	for i := range l {
		tmp := models.SystemVirtualWanLinkServiceUsers{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemVirtualWanLinkZone(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemVirtualWanLinkZone, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemVirtualWanLinkZone

	for i := range l {
		tmp := models.SystemVirtualWanLinkZone{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemVirtualWanLink(d *schema.ResourceData, sv string) (*models.SystemVirtualWanLink, diag.Diagnostics) {
	obj := models.SystemVirtualWanLink{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("fail_alert_interfaces"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("fail_alert_interfaces", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemVirtualWanLinkFailAlertInterfaces(d, v, "fail_alert_interfaces", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.FailAlertInterfaces = t
		}
	} else if d.HasChange("fail_alert_interfaces") {
		old, new := d.GetChange("fail_alert_interfaces")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.FailAlertInterfaces = &[]models.SystemVirtualWanLinkFailAlertInterfaces{}
		}
	}
	if v1, ok := d.GetOk("fail_detect"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fail_detect", sv)
				diags = append(diags, e)
			}
			obj.FailDetect = &v2
		}
	}
	if v, ok := d.GetOk("health_check"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("health_check", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemVirtualWanLinkHealthCheck(d, v, "health_check", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.HealthCheck = t
		}
	} else if d.HasChange("health_check") {
		old, new := d.GetChange("health_check")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.HealthCheck = &[]models.SystemVirtualWanLinkHealthCheck{}
		}
	}
	if v1, ok := d.GetOk("load_balance_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("load_balance_mode", sv)
				diags = append(diags, e)
			}
			obj.LoadBalanceMode = &v2
		}
	}
	if v, ok := d.GetOk("members"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("members", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemVirtualWanLinkMembers(d, v, "members", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Members = t
		}
	} else if d.HasChange("members") {
		old, new := d.GetChange("members")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Members = &[]models.SystemVirtualWanLinkMembers{}
		}
	}
	if v, ok := d.GetOk("neighbor"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("neighbor", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemVirtualWanLinkNeighbor(d, v, "neighbor", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Neighbor = t
		}
	} else if d.HasChange("neighbor") {
		old, new := d.GetChange("neighbor")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Neighbor = &[]models.SystemVirtualWanLinkNeighbor{}
		}
	}
	if v1, ok := d.GetOk("neighbor_hold_boot_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("neighbor_hold_boot_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.NeighborHoldBootTime = &tmp
		}
	}
	if v1, ok := d.GetOk("neighbor_hold_down"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("neighbor_hold_down", sv)
				diags = append(diags, e)
			}
			obj.NeighborHoldDown = &v2
		}
	}
	if v1, ok := d.GetOk("neighbor_hold_down_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("neighbor_hold_down_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.NeighborHoldDownTime = &tmp
		}
	}
	if v, ok := d.GetOk("service"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("service", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemVirtualWanLinkService(d, v, "service", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Service = t
		}
	} else if d.HasChange("service") {
		old, new := d.GetChange("service")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Service = &[]models.SystemVirtualWanLinkService{}
		}
	}
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	if v, ok := d.GetOk("zone"); ok {
		if !utils.CheckVer(sv, "", "v6.4.0") {
			e := utils.AttributeVersionWarning("zone", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemVirtualWanLinkZone(d, v, "zone", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Zone = t
		}
	} else if d.HasChange("zone") {
		old, new := d.GetChange("zone")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Zone = &[]models.SystemVirtualWanLinkZone{}
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemVirtualWanLink(d *schema.ResourceData, sv string) (*models.SystemVirtualWanLink, diag.Diagnostics) {
	obj := models.SystemVirtualWanLink{}
	diags := diag.Diagnostics{}

	obj.FailAlertInterfaces = &[]models.SystemVirtualWanLinkFailAlertInterfaces{}
	obj.HealthCheck = &[]models.SystemVirtualWanLinkHealthCheck{}
	obj.Members = &[]models.SystemVirtualWanLinkMembers{}
	obj.Neighbor = &[]models.SystemVirtualWanLinkNeighbor{}
	obj.Service = &[]models.SystemVirtualWanLinkService{}
	obj.Zone = &[]models.SystemVirtualWanLinkZone{}

	return &obj, diags
}
