// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure redundant Internet connections with multiple outbound links and health-check profiles.

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

func resourceSystemSdwan() *schema.Resource {
	return &schema.Resource{
		Description: "Configure redundant Internet connections with multiple outbound links and health-check profiles.",

		CreateContext: resourceSystemSdwanCreate,
		ReadContext:   resourceSystemSdwanRead,
		UpdateContext: resourceSystemSdwanUpdate,
		DeleteContext: resourceSystemSdwanDelete,

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
			"duplication": {
				Type:        schema.TypeList,
				Description: "Create SD-WAN duplication rule.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dstaddr": {
							Type:        schema.TypeList,
							Description: "Destination address or address group names.",
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
						"dstaddr6": {
							Type:        schema.TypeList,
							Description: "Destination address6 or address6 group names.",
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
						"dstintf": {
							Type:        schema.TypeList,
							Description: "Outgoing (egress) interfaces or zones.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Interface, zone or SDWAN zone name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),

							Description: "Duplication rule ID (1 - 255).",
							Optional:    true,
							Computed:    true,
						},
						"packet_de_duplication": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable discarding of packets that have been duplicated.",
							Optional:    true,
							Computed:    true,
						},
						"packet_duplication": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "force", "on-demand"}, false),

							Description: "Configure packet duplication method.",
							Optional:    true,
							Computed:    true,
						},
						"service": {
							Type:        schema.TypeList,
							Description: "Service and service group name.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Service and service group name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"service_id": {
							Type:        schema.TypeList,
							Description: "SD-WAN service rule ID list.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type: schema.TypeInt,

										Description: "SD-WAN service rule ID.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"srcaddr": {
							Type:        schema.TypeList,
							Description: "Source address or address group names.",
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
						"srcaddr6": {
							Type:        schema.TypeList,
							Description: "Source address6 or address6 group names.",
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
						"srcintf": {
							Type:        schema.TypeList,
							Description: "Incoming (ingress) interfaces or zones.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Interface, zone or SDWAN zone name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"duplication_max_num": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 4),

				Description: "Maximum number of interface members a packet is duplicated in the SD-WAN zone (2 - 4, default = 2; if set to 3, the original packet plus 2 more copies are created).",
				Optional:    true,
				Computed:    true,
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
						"detect_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"active", "passive", "prefer-passive"}, false),

							Description: "The mode determining how to detect the server.",
							Optional:    true,
							Computed:    true,
						},
						"diffservcode": {
							Type: schema.TypeString,

							Description: "Differentiated services code point (DSCP) in the IP header of the probe packet.",
							Optional:    true,
							Computed:    true,
						},
						"dns_match_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Response IP expected from DNS server if the protocol is DNS.",
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
						"ftp_file": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 254),

							Description: "Full path and file name on the FTP server to download for FTP health-check to probe.",
							Optional:    true,
							Computed:    true,
						},
						"ftp_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"passive", "port"}, false),

							Description: "FTP mode.",
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

							Description: "Packet size of a TWAMP test session.",
							Optional:    true,
							Computed:    true,
						},
						"password": {
							Type: schema.TypeString,

							Description: "TWAMP controller password in authentication mode.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Port number used to communicate with the server over the selected protocol (0 - 65535, default = 0, auto select. http, tcp-connect: 80, udp-echo, tcp-echo: 7, dns: 53, ftp: 21, twamp: 862).",
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
							ValidateFunc: validation.IntBetween(500, 3600000),

							Description: "Time to wait before a probe packet is considered lost (500 - 3600*1000 msec, default = 500).",
							Optional:    true,
							Computed:    true,
						},
						"protocol": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ping", "tcp-echo", "udp-echo", "http", "twamp", "dns", "tcp-connect", "ftp"}, false),

							Description: "Protocol used to determine if the FortiGate can communicate with the server.",
							Optional:    true,
							Computed:    true,
						},
						"quality_measured_method": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"half-open", "half-close"}, false),

							Description: "Method to measure the quality of tcp-connect.",
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
							Type:             schema.TypeString,
							ValidateFunc:     validation.StringLenBetween(0, 79),
							DiffSuppressFunc: suppressors.DiffMultiStringEqual,
							Description:      "IP address or FQDN name of the server.",
							Optional:         true,
							Computed:         true,
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
						"user": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),

							Description: "The user name to access probe server.",
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
				Description: "FortiGate interfaces added to the SD-WAN.",
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
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Priority of the interface for IPv4 (1 - 65535, default = 1). Used for SD-WAN rules or priority rules.",
							Optional:    true,
							Computed:    true,
						},
						"priority6": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Priority of the interface for IPv6 (1 - 65535, default = 1024). Used for SD-WAN rules or priority rules.",
							Optional:    true,
							Computed:    true,
						},
						"seq_num": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 512),

							Description: "Sequence number(1-512).",
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
						"zone": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Zone name.",
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
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"sla", "speedtest"}, false),

							Description: "What metric to select the neighbor.",
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
						"hash_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"round-robin", "source-ip-based", "source-dest-ip-based", "inbandwidth", "outbandwidth", "bibandwidth"}, false),

							Description: "Hash algorithm for selected priority members for load balance mode.",
							Optional:    true,
							Computed:    true,
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

							Description: "SD-WAN rule ID (1 - 4000).",
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
						"minimum_sla_meet_members": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Minimum number of members which meet SLA.",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"auto", "manual", "priority", "sla", "load-balance"}, false),

							Description: "Control how the SD-WAN rule sets the priority of interfaces in the SD-WAN.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "SD-WAN rule name.",
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
						"passive_measurement": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable passive measurement based on the service criteria.",
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
						"priority_zone": {
							Type:        schema.TypeList,
							Description: "Priority zone name list.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Priority zone name.",
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

										Description: "SD-WAN health-check.",
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
						"tie_break": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"zone", "cfg-order", "fib-best-match"}, false),

							Description: "Method of selecting member if more than one meets the SLA.",
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
						"use_shortcut_sla": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable use of ADVPN shortcut for quality comparison.",
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
			"speedtest_bypass_routing": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable bypass routing when speedtest on a SD-WAN member.",
				Optional:    true,
				Computed:    true,
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
						"service_sla_tie_break": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"cfg-order", "fib-best-match"}, false),

							Description: "Method of selecting member if more than one meets the SLA.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemSdwanCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemSdwan(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemSdwan(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSdwan")
	}

	return resourceSystemSdwanRead(ctx, d, meta)
}

func resourceSystemSdwanUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemSdwan(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemSdwan(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemSdwan resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSdwan")
	}

	return resourceSystemSdwanRead(ctx, d, meta)
}

func resourceSystemSdwanDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemSdwan(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemSdwan(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemSdwan resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdwanRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemSdwan(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSdwan resource: %v", err)
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

	diags := refreshObjectSystemSdwan(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemSdwanDuplication(d *schema.ResourceData, v *[]models.SystemSdwanDuplication, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Dstaddr; tmp != nil {
				v["dstaddr"] = flattenSystemSdwanDuplicationDstaddr(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "dstaddr"), sort)
			}

			if tmp := cfg.Dstaddr6; tmp != nil {
				v["dstaddr6"] = flattenSystemSdwanDuplicationDstaddr6(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "dstaddr6"), sort)
			}

			if tmp := cfg.Dstintf; tmp != nil {
				v["dstintf"] = flattenSystemSdwanDuplicationDstintf(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "dstintf"), sort)
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.PacketDeDuplication; tmp != nil {
				v["packet_de_duplication"] = *tmp
			}

			if tmp := cfg.PacketDuplication; tmp != nil {
				v["packet_duplication"] = *tmp
			}

			if tmp := cfg.Service; tmp != nil {
				v["service"] = flattenSystemSdwanDuplicationService(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "service"), sort)
			}

			if tmp := cfg.ServiceId; tmp != nil {
				v["service_id"] = flattenSystemSdwanDuplicationServiceId(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "service_id"), sort)
			}

			if tmp := cfg.Srcaddr; tmp != nil {
				v["srcaddr"] = flattenSystemSdwanDuplicationSrcaddr(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "srcaddr"), sort)
			}

			if tmp := cfg.Srcaddr6; tmp != nil {
				v["srcaddr6"] = flattenSystemSdwanDuplicationSrcaddr6(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "srcaddr6"), sort)
			}

			if tmp := cfg.Srcintf; tmp != nil {
				v["srcintf"] = flattenSystemSdwanDuplicationSrcintf(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "srcintf"), sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemSdwanDuplicationDstaddr(d *schema.ResourceData, v *[]models.SystemSdwanDuplicationDstaddr, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanDuplicationDstaddr6(d *schema.ResourceData, v *[]models.SystemSdwanDuplicationDstaddr6, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanDuplicationDstintf(d *schema.ResourceData, v *[]models.SystemSdwanDuplicationDstintf, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanDuplicationService(d *schema.ResourceData, v *[]models.SystemSdwanDuplicationService, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanDuplicationServiceId(d *schema.ResourceData, v *[]models.SystemSdwanDuplicationServiceId, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanDuplicationSrcaddr(d *schema.ResourceData, v *[]models.SystemSdwanDuplicationSrcaddr, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanDuplicationSrcaddr6(d *schema.ResourceData, v *[]models.SystemSdwanDuplicationSrcaddr6, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanDuplicationSrcintf(d *schema.ResourceData, v *[]models.SystemSdwanDuplicationSrcintf, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanFailAlertInterfaces(d *schema.ResourceData, v *[]models.SystemSdwanFailAlertInterfaces, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanHealthCheck(d *schema.ResourceData, v *[]models.SystemSdwanHealthCheck, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AddrMode; tmp != nil {
				v["addr_mode"] = *tmp
			}

			if tmp := cfg.DetectMode; tmp != nil {
				v["detect_mode"] = *tmp
			}

			if tmp := cfg.Diffservcode; tmp != nil {
				v["diffservcode"] = *tmp
			}

			if tmp := cfg.DnsMatchIp; tmp != nil {
				v["dns_match_ip"] = *tmp
			}

			if tmp := cfg.DnsRequestDomain; tmp != nil {
				v["dns_request_domain"] = *tmp
			}

			if tmp := cfg.Failtime; tmp != nil {
				v["failtime"] = *tmp
			}

			if tmp := cfg.FtpFile; tmp != nil {
				v["ftp_file"] = *tmp
			}

			if tmp := cfg.FtpMode; tmp != nil {
				v["ftp_mode"] = *tmp
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
				v["members"] = flattenSystemSdwanHealthCheckMembers(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "members"), sort)
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

			if tmp := cfg.QualityMeasuredMethod; tmp != nil {
				v["quality_measured_method"] = *tmp
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
				v["sla"] = flattenSystemSdwanHealthCheckSla(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "sla"), sort)
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

			if tmp := cfg.User; tmp != nil {
				v["user"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemSdwanHealthCheckMembers(d *schema.ResourceData, v *[]models.SystemSdwanHealthCheckMembers, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanHealthCheckSla(d *schema.ResourceData, v *[]models.SystemSdwanHealthCheckSla, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanMembers(d *schema.ResourceData, v *[]models.SystemSdwanMembers, prefix string, sort bool) interface{} {
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

			if tmp := cfg.Priority6; tmp != nil {
				v["priority6"] = *tmp
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

			if tmp := cfg.Zone; tmp != nil {
				v["zone"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "seq_num")
	}

	return flat
}

func flattenSystemSdwanNeighbor(d *schema.ResourceData, v *[]models.SystemSdwanNeighbor, prefix string, sort bool) interface{} {
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

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
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

func flattenSystemSdwanService(d *schema.ResourceData, v *[]models.SystemSdwanService, prefix string, sort bool) interface{} {
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
				v["dst"] = flattenSystemSdwanServiceDst(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "dst"), sort)
			}

			if tmp := cfg.DstNegate; tmp != nil {
				v["dst_negate"] = *tmp
			}

			if tmp := cfg.Dst6; tmp != nil {
				v["dst6"] = flattenSystemSdwanServiceDst6(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "dst6"), sort)
			}

			if tmp := cfg.EndPort; tmp != nil {
				v["end_port"] = *tmp
			}

			if tmp := cfg.Gateway; tmp != nil {
				v["gateway"] = *tmp
			}

			if tmp := cfg.Groups; tmp != nil {
				v["groups"] = flattenSystemSdwanServiceGroups(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "groups"), sort)
			}

			if tmp := cfg.HashMode; tmp != nil {
				v["hash_mode"] = *tmp
			}

			if tmp := cfg.HealthCheck; tmp != nil {
				v["health_check"] = flattenSystemSdwanServiceHealthCheck(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "health_check"), sort)
			}

			if tmp := cfg.HoldDownTime; tmp != nil {
				v["hold_down_time"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.InputDevice; tmp != nil {
				v["input_device"] = flattenSystemSdwanServiceInputDevice(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "input_device"), sort)
			}

			if tmp := cfg.InputDeviceNegate; tmp != nil {
				v["input_device_negate"] = *tmp
			}

			if tmp := cfg.InternetService; tmp != nil {
				v["internet_service"] = *tmp
			}

			if tmp := cfg.InternetServiceAppCtrl; tmp != nil {
				v["internet_service_app_ctrl"] = flattenSystemSdwanServiceInternetServiceAppCtrl(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "internet_service_app_ctrl"), sort)
			}

			if tmp := cfg.InternetServiceAppCtrlGroup; tmp != nil {
				v["internet_service_app_ctrl_group"] = flattenSystemSdwanServiceInternetServiceAppCtrlGroup(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "internet_service_app_ctrl_group"), sort)
			}

			if tmp := cfg.InternetServiceCustom; tmp != nil {
				v["internet_service_custom"] = flattenSystemSdwanServiceInternetServiceCustom(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "internet_service_custom"), sort)
			}

			if tmp := cfg.InternetServiceCustomGroup; tmp != nil {
				v["internet_service_custom_group"] = flattenSystemSdwanServiceInternetServiceCustomGroup(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "internet_service_custom_group"), sort)
			}

			if tmp := cfg.InternetServiceGroup; tmp != nil {
				v["internet_service_group"] = flattenSystemSdwanServiceInternetServiceGroup(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "internet_service_group"), sort)
			}

			if tmp := cfg.InternetServiceName; tmp != nil {
				v["internet_service_name"] = flattenSystemSdwanServiceInternetServiceName(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "internet_service_name"), sort)
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

			if tmp := cfg.MinimumSlaMeetMembers; tmp != nil {
				v["minimum_sla_meet_members"] = *tmp
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

			if tmp := cfg.PassiveMeasurement; tmp != nil {
				v["passive_measurement"] = *tmp
			}

			if tmp := cfg.PriorityMembers; tmp != nil {
				v["priority_members"] = flattenSystemSdwanServicePriorityMembers(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "priority_members"), sort)
			}

			if tmp := cfg.PriorityZone; tmp != nil {
				v["priority_zone"] = flattenSystemSdwanServicePriorityZone(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "priority_zone"), sort)
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
				v["sla"] = flattenSystemSdwanServiceSla(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "sla"), sort)
			}

			if tmp := cfg.SlaCompareMethod; tmp != nil {
				v["sla_compare_method"] = *tmp
			}

			if tmp := cfg.Src; tmp != nil {
				v["src"] = flattenSystemSdwanServiceSrc(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "src"), sort)
			}

			if tmp := cfg.SrcNegate; tmp != nil {
				v["src_negate"] = *tmp
			}

			if tmp := cfg.Src6; tmp != nil {
				v["src6"] = flattenSystemSdwanServiceSrc6(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "src6"), sort)
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

			if tmp := cfg.TieBreak; tmp != nil {
				v["tie_break"] = *tmp
			}

			if tmp := cfg.Tos; tmp != nil {
				v["tos"] = *tmp
			}

			if tmp := cfg.TosMask; tmp != nil {
				v["tos_mask"] = *tmp
			}

			if tmp := cfg.UseShortcutSla; tmp != nil {
				v["use_shortcut_sla"] = *tmp
			}

			if tmp := cfg.Users; tmp != nil {
				v["users"] = flattenSystemSdwanServiceUsers(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "users"), sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemSdwanServiceDst(d *schema.ResourceData, v *[]models.SystemSdwanServiceDst, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanServiceDst6(d *schema.ResourceData, v *[]models.SystemSdwanServiceDst6, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanServiceGroups(d *schema.ResourceData, v *[]models.SystemSdwanServiceGroups, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanServiceHealthCheck(d *schema.ResourceData, v *[]models.SystemSdwanServiceHealthCheck, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanServiceInputDevice(d *schema.ResourceData, v *[]models.SystemSdwanServiceInputDevice, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanServiceInternetServiceAppCtrl(d *schema.ResourceData, v *[]models.SystemSdwanServiceInternetServiceAppCtrl, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanServiceInternetServiceAppCtrlGroup(d *schema.ResourceData, v *[]models.SystemSdwanServiceInternetServiceAppCtrlGroup, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanServiceInternetServiceCustom(d *schema.ResourceData, v *[]models.SystemSdwanServiceInternetServiceCustom, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanServiceInternetServiceCustomGroup(d *schema.ResourceData, v *[]models.SystemSdwanServiceInternetServiceCustomGroup, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanServiceInternetServiceGroup(d *schema.ResourceData, v *[]models.SystemSdwanServiceInternetServiceGroup, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanServiceInternetServiceName(d *schema.ResourceData, v *[]models.SystemSdwanServiceInternetServiceName, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanServicePriorityMembers(d *schema.ResourceData, v *[]models.SystemSdwanServicePriorityMembers, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanServicePriorityZone(d *schema.ResourceData, v *[]models.SystemSdwanServicePriorityZone, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanServiceSla(d *schema.ResourceData, v *[]models.SystemSdwanServiceSla, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanServiceSrc(d *schema.ResourceData, v *[]models.SystemSdwanServiceSrc, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanServiceSrc6(d *schema.ResourceData, v *[]models.SystemSdwanServiceSrc6, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanServiceUsers(d *schema.ResourceData, v *[]models.SystemSdwanServiceUsers, prefix string, sort bool) interface{} {
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

func flattenSystemSdwanZone(d *schema.ResourceData, v *[]models.SystemSdwanZone, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.ServiceSlaTieBreak; tmp != nil {
				v["service_sla_tie_break"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectSystemSdwan(d *schema.ResourceData, o *models.SystemSdwan, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Duplication != nil {
		if err = d.Set("duplication", flattenSystemSdwanDuplication(d, o.Duplication, "duplication", sort)); err != nil {
			return diag.Errorf("error reading duplication: %v", err)
		}
	}

	if o.DuplicationMaxNum != nil {
		v := *o.DuplicationMaxNum

		if err = d.Set("duplication_max_num", v); err != nil {
			return diag.Errorf("error reading duplication_max_num: %v", err)
		}
	}

	if o.FailAlertInterfaces != nil {
		if err = d.Set("fail_alert_interfaces", flattenSystemSdwanFailAlertInterfaces(d, o.FailAlertInterfaces, "fail_alert_interfaces", sort)); err != nil {
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
		if err = d.Set("health_check", flattenSystemSdwanHealthCheck(d, o.HealthCheck, "health_check", sort)); err != nil {
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
		if err = d.Set("members", flattenSystemSdwanMembers(d, o.Members, "members", sort)); err != nil {
			return diag.Errorf("error reading members: %v", err)
		}
	}

	if o.Neighbor != nil {
		if err = d.Set("neighbor", flattenSystemSdwanNeighbor(d, o.Neighbor, "neighbor", sort)); err != nil {
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
		if err = d.Set("service", flattenSystemSdwanService(d, o.Service, "service", sort)); err != nil {
			return diag.Errorf("error reading service: %v", err)
		}
	}

	if o.SpeedtestBypassRouting != nil {
		v := *o.SpeedtestBypassRouting

		if err = d.Set("speedtest_bypass_routing", v); err != nil {
			return diag.Errorf("error reading speedtest_bypass_routing: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Zone != nil {
		if err = d.Set("zone", flattenSystemSdwanZone(d, o.Zone, "zone", sort)); err != nil {
			return diag.Errorf("error reading zone: %v", err)
		}
	}

	return nil
}

func expandSystemSdwanDuplication(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanDuplication, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanDuplication

	for i := range l {
		tmp := models.SystemSdwanDuplication{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dstaddr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemSdwanDuplicationDstaddr(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanDuplicationDstaddr
			// 	}
			tmp.Dstaddr = v2

		}

		pre_append = fmt.Sprintf("%s.%d.dstaddr6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemSdwanDuplicationDstaddr6(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanDuplicationDstaddr6
			// 	}
			tmp.Dstaddr6 = v2

		}

		pre_append = fmt.Sprintf("%s.%d.dstintf", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemSdwanDuplicationDstintf(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanDuplicationDstintf
			// 	}
			tmp.Dstintf = v2

		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.packet_de_duplication", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PacketDeDuplication = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.packet_duplication", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PacketDuplication = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.service", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemSdwanDuplicationService(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanDuplicationService
			// 	}
			tmp.Service = v2

		}

		pre_append = fmt.Sprintf("%s.%d.service_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemSdwanDuplicationServiceId(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanDuplicationServiceId
			// 	}
			tmp.ServiceId = v2

		}

		pre_append = fmt.Sprintf("%s.%d.srcaddr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemSdwanDuplicationSrcaddr(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanDuplicationSrcaddr
			// 	}
			tmp.Srcaddr = v2

		}

		pre_append = fmt.Sprintf("%s.%d.srcaddr6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemSdwanDuplicationSrcaddr6(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanDuplicationSrcaddr6
			// 	}
			tmp.Srcaddr6 = v2

		}

		pre_append = fmt.Sprintf("%s.%d.srcintf", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemSdwanDuplicationSrcintf(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanDuplicationSrcintf
			// 	}
			tmp.Srcintf = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemSdwanDuplicationDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanDuplicationDstaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanDuplicationDstaddr

	for i := range l {
		tmp := models.SystemSdwanDuplicationDstaddr{}
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

func expandSystemSdwanDuplicationDstaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanDuplicationDstaddr6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanDuplicationDstaddr6

	for i := range l {
		tmp := models.SystemSdwanDuplicationDstaddr6{}
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

func expandSystemSdwanDuplicationDstintf(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanDuplicationDstintf, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanDuplicationDstintf

	for i := range l {
		tmp := models.SystemSdwanDuplicationDstintf{}
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

func expandSystemSdwanDuplicationService(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanDuplicationService, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanDuplicationService

	for i := range l {
		tmp := models.SystemSdwanDuplicationService{}
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

func expandSystemSdwanDuplicationServiceId(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanDuplicationServiceId, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanDuplicationServiceId

	for i := range l {
		tmp := models.SystemSdwanDuplicationServiceId{}
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

func expandSystemSdwanDuplicationSrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanDuplicationSrcaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanDuplicationSrcaddr

	for i := range l {
		tmp := models.SystemSdwanDuplicationSrcaddr{}
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

func expandSystemSdwanDuplicationSrcaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanDuplicationSrcaddr6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanDuplicationSrcaddr6

	for i := range l {
		tmp := models.SystemSdwanDuplicationSrcaddr6{}
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

func expandSystemSdwanDuplicationSrcintf(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanDuplicationSrcintf, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanDuplicationSrcintf

	for i := range l {
		tmp := models.SystemSdwanDuplicationSrcintf{}
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

func expandSystemSdwanFailAlertInterfaces(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanFailAlertInterfaces, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanFailAlertInterfaces

	for i := range l {
		tmp := models.SystemSdwanFailAlertInterfaces{}
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

func expandSystemSdwanHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanHealthCheck, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanHealthCheck

	for i := range l {
		tmp := models.SystemSdwanHealthCheck{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.addr_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AddrMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.detect_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DetectMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.diffservcode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Diffservcode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dns_match_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DnsMatchIp = &v2
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

		pre_append = fmt.Sprintf("%s.%d.ftp_file", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FtpFile = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ftp_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FtpMode = &v2
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
			v2, _ := expandSystemSdwanHealthCheckMembers(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanHealthCheckMembers
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

		pre_append = fmt.Sprintf("%s.%d.quality_measured_method", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.QualityMeasuredMethod = &v2
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
			v2, _ := expandSystemSdwanHealthCheckSla(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanHealthCheckSla
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

		pre_append = fmt.Sprintf("%s.%d.user", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.User = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemSdwanHealthCheckMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanHealthCheckMembers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanHealthCheckMembers

	for i := range l {
		tmp := models.SystemSdwanHealthCheckMembers{}
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

func expandSystemSdwanHealthCheckSla(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanHealthCheckSla, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanHealthCheckSla

	for i := range l {
		tmp := models.SystemSdwanHealthCheckSla{}
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

func expandSystemSdwanMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanMembers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanMembers

	for i := range l {
		tmp := models.SystemSdwanMembers{}
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

		pre_append = fmt.Sprintf("%s.%d.priority6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Priority6 = &v3
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

		pre_append = fmt.Sprintf("%s.%d.zone", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Zone = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemSdwanNeighbor(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanNeighbor, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanNeighbor

	for i := range l {
		tmp := models.SystemSdwanNeighbor{}
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

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
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

func expandSystemSdwanService(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanService, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanService

	for i := range l {
		tmp := models.SystemSdwanService{}
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
			v2, _ := expandSystemSdwanServiceDst(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanServiceDst
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
			v2, _ := expandSystemSdwanServiceDst6(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanServiceDst6
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
			v2, _ := expandSystemSdwanServiceGroups(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanServiceGroups
			// 	}
			tmp.Groups = v2

		}

		pre_append = fmt.Sprintf("%s.%d.hash_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HashMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.health_check", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemSdwanServiceHealthCheck(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanServiceHealthCheck
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
			v2, _ := expandSystemSdwanServiceInputDevice(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanServiceInputDevice
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
			v2, _ := expandSystemSdwanServiceInternetServiceAppCtrl(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanServiceInternetServiceAppCtrl
			// 	}
			tmp.InternetServiceAppCtrl = v2

		}

		pre_append = fmt.Sprintf("%s.%d.internet_service_app_ctrl_group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemSdwanServiceInternetServiceAppCtrlGroup(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanServiceInternetServiceAppCtrlGroup
			// 	}
			tmp.InternetServiceAppCtrlGroup = v2

		}

		pre_append = fmt.Sprintf("%s.%d.internet_service_custom", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemSdwanServiceInternetServiceCustom(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanServiceInternetServiceCustom
			// 	}
			tmp.InternetServiceCustom = v2

		}

		pre_append = fmt.Sprintf("%s.%d.internet_service_custom_group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemSdwanServiceInternetServiceCustomGroup(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanServiceInternetServiceCustomGroup
			// 	}
			tmp.InternetServiceCustomGroup = v2

		}

		pre_append = fmt.Sprintf("%s.%d.internet_service_group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemSdwanServiceInternetServiceGroup(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanServiceInternetServiceGroup
			// 	}
			tmp.InternetServiceGroup = v2

		}

		pre_append = fmt.Sprintf("%s.%d.internet_service_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemSdwanServiceInternetServiceName(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanServiceInternetServiceName
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

		pre_append = fmt.Sprintf("%s.%d.minimum_sla_meet_members", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MinimumSlaMeetMembers = &v3
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

		pre_append = fmt.Sprintf("%s.%d.passive_measurement", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PassiveMeasurement = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority_members", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemSdwanServicePriorityMembers(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanServicePriorityMembers
			// 	}
			tmp.PriorityMembers = v2

		}

		pre_append = fmt.Sprintf("%s.%d.priority_zone", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemSdwanServicePriorityZone(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanServicePriorityZone
			// 	}
			tmp.PriorityZone = v2

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
			v2, _ := expandSystemSdwanServiceSla(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanServiceSla
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
			v2, _ := expandSystemSdwanServiceSrc(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanServiceSrc
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
			v2, _ := expandSystemSdwanServiceSrc6(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanServiceSrc6
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

		pre_append = fmt.Sprintf("%s.%d.tie_break", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TieBreak = &v2
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

		pre_append = fmt.Sprintf("%s.%d.use_shortcut_sla", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UseShortcutSla = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.users", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemSdwanServiceUsers(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemSdwanServiceUsers
			// 	}
			tmp.Users = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemSdwanServiceDst(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanServiceDst, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanServiceDst

	for i := range l {
		tmp := models.SystemSdwanServiceDst{}
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

func expandSystemSdwanServiceDst6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanServiceDst6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanServiceDst6

	for i := range l {
		tmp := models.SystemSdwanServiceDst6{}
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

func expandSystemSdwanServiceGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanServiceGroups, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanServiceGroups

	for i := range l {
		tmp := models.SystemSdwanServiceGroups{}
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

func expandSystemSdwanServiceHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanServiceHealthCheck, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanServiceHealthCheck

	for i := range l {
		tmp := models.SystemSdwanServiceHealthCheck{}
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

func expandSystemSdwanServiceInputDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanServiceInputDevice, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanServiceInputDevice

	for i := range l {
		tmp := models.SystemSdwanServiceInputDevice{}
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

func expandSystemSdwanServiceInternetServiceAppCtrl(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanServiceInternetServiceAppCtrl, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanServiceInternetServiceAppCtrl

	for i := range l {
		tmp := models.SystemSdwanServiceInternetServiceAppCtrl{}
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

func expandSystemSdwanServiceInternetServiceAppCtrlGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanServiceInternetServiceAppCtrlGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanServiceInternetServiceAppCtrlGroup

	for i := range l {
		tmp := models.SystemSdwanServiceInternetServiceAppCtrlGroup{}
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

func expandSystemSdwanServiceInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanServiceInternetServiceCustom, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanServiceInternetServiceCustom

	for i := range l {
		tmp := models.SystemSdwanServiceInternetServiceCustom{}
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

func expandSystemSdwanServiceInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanServiceInternetServiceCustomGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanServiceInternetServiceCustomGroup

	for i := range l {
		tmp := models.SystemSdwanServiceInternetServiceCustomGroup{}
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

func expandSystemSdwanServiceInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanServiceInternetServiceGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanServiceInternetServiceGroup

	for i := range l {
		tmp := models.SystemSdwanServiceInternetServiceGroup{}
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

func expandSystemSdwanServiceInternetServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanServiceInternetServiceName, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanServiceInternetServiceName

	for i := range l {
		tmp := models.SystemSdwanServiceInternetServiceName{}
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

func expandSystemSdwanServicePriorityMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanServicePriorityMembers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanServicePriorityMembers

	for i := range l {
		tmp := models.SystemSdwanServicePriorityMembers{}
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

func expandSystemSdwanServicePriorityZone(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanServicePriorityZone, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanServicePriorityZone

	for i := range l {
		tmp := models.SystemSdwanServicePriorityZone{}
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

func expandSystemSdwanServiceSla(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanServiceSla, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanServiceSla

	for i := range l {
		tmp := models.SystemSdwanServiceSla{}
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

func expandSystemSdwanServiceSrc(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanServiceSrc, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanServiceSrc

	for i := range l {
		tmp := models.SystemSdwanServiceSrc{}
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

func expandSystemSdwanServiceSrc6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanServiceSrc6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanServiceSrc6

	for i := range l {
		tmp := models.SystemSdwanServiceSrc6{}
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

func expandSystemSdwanServiceUsers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanServiceUsers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanServiceUsers

	for i := range l {
		tmp := models.SystemSdwanServiceUsers{}
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

func expandSystemSdwanZone(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSdwanZone, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSdwanZone

	for i := range l {
		tmp := models.SystemSdwanZone{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.service_sla_tie_break", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ServiceSlaTieBreak = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemSdwan(d *schema.ResourceData, sv string) (*models.SystemSdwan, diag.Diagnostics) {
	obj := models.SystemSdwan{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("duplication"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("duplication", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSdwanDuplication(d, v, "duplication", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Duplication = t
		}
	} else if d.HasChange("duplication") {
		old, new := d.GetChange("duplication")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Duplication = &[]models.SystemSdwanDuplication{}
		}
	}
	if v1, ok := d.GetOk("duplication_max_num"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("duplication_max_num", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DuplicationMaxNum = &tmp
		}
	}
	if v, ok := d.GetOk("fail_alert_interfaces"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("fail_alert_interfaces", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSdwanFailAlertInterfaces(d, v, "fail_alert_interfaces", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.FailAlertInterfaces = t
		}
	} else if d.HasChange("fail_alert_interfaces") {
		old, new := d.GetChange("fail_alert_interfaces")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.FailAlertInterfaces = &[]models.SystemSdwanFailAlertInterfaces{}
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
		t, err := expandSystemSdwanHealthCheck(d, v, "health_check", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.HealthCheck = t
		}
	} else if d.HasChange("health_check") {
		old, new := d.GetChange("health_check")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.HealthCheck = &[]models.SystemSdwanHealthCheck{}
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
		t, err := expandSystemSdwanMembers(d, v, "members", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Members = t
		}
	} else if d.HasChange("members") {
		old, new := d.GetChange("members")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Members = &[]models.SystemSdwanMembers{}
		}
	}
	if v, ok := d.GetOk("neighbor"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("neighbor", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSdwanNeighbor(d, v, "neighbor", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Neighbor = t
		}
	} else if d.HasChange("neighbor") {
		old, new := d.GetChange("neighbor")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Neighbor = &[]models.SystemSdwanNeighbor{}
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
		t, err := expandSystemSdwanService(d, v, "service", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Service = t
		}
	} else if d.HasChange("service") {
		old, new := d.GetChange("service")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Service = &[]models.SystemSdwanService{}
		}
	}
	if v1, ok := d.GetOk("speedtest_bypass_routing"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("speedtest_bypass_routing", sv)
				diags = append(diags, e)
			}
			obj.SpeedtestBypassRouting = &v2
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
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("zone", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSdwanZone(d, v, "zone", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Zone = t
		}
	} else if d.HasChange("zone") {
		old, new := d.GetChange("zone")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Zone = &[]models.SystemSdwanZone{}
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemSdwan(d *schema.ResourceData, sv string) (*models.SystemSdwan, diag.Diagnostics) {
	obj := models.SystemSdwan{}
	diags := diag.Diagnostics{}

	obj.Duplication = &[]models.SystemSdwanDuplication{}
	obj.FailAlertInterfaces = &[]models.SystemSdwanFailAlertInterfaces{}
	obj.HealthCheck = &[]models.SystemSdwanHealthCheck{}
	obj.Members = &[]models.SystemSdwanMembers{}
	obj.Neighbor = &[]models.SystemSdwanNeighbor{}
	obj.Service = &[]models.SystemSdwanService{}
	obj.Zone = &[]models.SystemSdwanZone{}

	return &obj, diags
}
