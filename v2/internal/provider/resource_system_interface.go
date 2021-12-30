// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure interfaces.

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
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceSystemInterface() *schema.Resource {
	return &schema.Resource{
		Description: "Configure interfaces.",

		CreateContext: resourceSystemInterfaceCreate,
		ReadContext:   resourceSystemInterfaceRead,
		UpdateContext: resourceSystemInterfaceUpdate,
		DeleteContext: resourceSystemInterfaceDelete,

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
			"allow_append": {
				Type:        schema.TypeBool,
				Description: "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ac_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "PPPoE server name.",
				Optional:    true,
				Computed:    true,
			},
			"aggregate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Aggregate interface.",
				Optional:    true,
				Computed:    true,
			},
			"algorithm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"L2", "L3", "L4"}, false),

				Description: "Frame distribution algorithm.",
				Optional:    true,
				Computed:    true,
			},
			"alias": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 25),

				Description: "Alias will be displayed with the interface name to make it easier to distinguish.",
				Optional:    true,
				Computed:    true,
			},
			"allowaccess": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Permitted types of management access to this interface.",
				Optional:         true,
				Computed:         true,
			},
			"ap_discover": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable automatic registration of unknown FortiAP devices.",
				Optional:    true,
				Computed:    true,
			},
			"arpforward": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable ARP forwarding.",
				Optional:    true,
				Computed:    true,
			},
			"auth_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "pap", "chap", "mschapv1", "mschapv2"}, false),

				Description: "PPP authentication type to use.",
				Optional:    true,
				Computed:    true,
			},
			"auto_auth_extension_device": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable automatic authorization of dedicated Fortinet extension device on this interface.",
				Optional:    true,
				Computed:    true,
			},
			"bandwidth_measure_time": {
				Type: schema.TypeInt,

				Description: "Bandwidth measure time ",
				Optional:    true,
				Computed:    true,
			},
			"bfd": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"global", "enable", "disable"}, false),

				Description: "Bidirectional Forwarding Detection (BFD) settings.",
				Optional:    true,
				Computed:    true,
			},
			"bfd_desired_min_tx": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100000),

				Description: "BFD desired minimal transmit interval.",
				Optional:    true,
				Computed:    true,
			},
			"bfd_detect_mult": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 50),

				Description: "BFD detection multiplier.",
				Optional:    true,
				Computed:    true,
			},
			"bfd_required_min_rx": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100000),

				Description: "BFD required minimal receive interval.",
				Optional:    true,
				Computed:    true,
			},
			"broadcast_forticlient_discovery": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable broadcasting FortiClient discovery messages.",
				Optional:    true,
				Computed:    true,
			},
			"broadcast_forward": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable broadcast forwarding.",
				Optional:    true,
				Computed:    true,
			},
			"cli_conn_status": {
				Type: schema.TypeInt,

				Description: "CLI connection status.",
				Optional:    true,
				Computed:    true,
			},
			"client_options": {
				Type:        schema.TypeList,
				Description: "DHCP client options.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "DHCP client option code.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"ip": {
							Type: schema.TypeString,

							Description: "DHCP option IPs.",
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"hex", "string", "ip", "fqdn"}, false),

							Description: "DHCP client option type.",
							Optional:    true,
							Computed:    true,
						},
						"value": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 312),

							Description: "DHCP client option value.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"color": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),

				Description: "Color of icon on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"dedicated_to": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "management"}, false),

				Description: "Configure interface for single purpose.",
				Optional:    true,
				Computed:    true,
			},
			"defaultgw": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to get the gateway IP from the DHCP or PPPoE server.",
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Description.",
				Optional:    true,
				Computed:    true,
			},
			"detected_peer_mtu": {
				Type: schema.TypeInt,

				Description: "MTU of detected peer (0 - 4294967295).",
				Optional:    true,
				Computed:    true,
			},
			"detectprotocol": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Protocols used to detect the server.",
				Optional:         true,
				Computed:         true,
			},
			"detectserver": {
				Type: schema.TypeString,

				Description: "Gateway's ping server for this IP.",
				Optional:    true,
				Computed:    true,
			},
			"device_identification": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.",
				Optional:    true,
				Computed:    true,
			},
			"device_user_identification": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable passive gathering of user identity information about users on this interface.",
				Optional:    true,
				Computed:    true,
			},
			"devindex": {
				Type: schema.TypeInt,

				Description: "Device Index.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_classless_route_addition": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable addition of classless static routes retrieved from DHCP server.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_client_identifier": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 48),

				Description: "DHCP client identifier.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_relay_agent_option": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable DHCP relay agent option.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_relay_interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Specify outgoing interface to reach server.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_relay_interface_select_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "sdwan", "specify"}, false),

				Description: "Specify how to select outgoing interface to reach server.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_relay_ip": {
				Type: schema.TypeString,

				Description: "DHCP relay IP address.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_relay_request_all_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable sending of DHCP requests to all servers.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_relay_service": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable allowing this interface to act as a DHCP relay.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_relay_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"regular", "ipsec"}, false),

				Description: "DHCP relay type (regular or IPsec).",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_renew_time": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 604800),

				Description: "DHCP renew time in seconds (300-604800), 0 means use the renew time provided by the server.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_snooping_server_list": {
				Type:        schema.TypeList,
				Description: "Configure DHCP server access list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "DHCP server name.",
							Optional:    true,
							Computed:    true,
						},
						"server_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "IP address for DHCP server.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"disc_retry_timeout": {
				Type: schema.TypeInt,

				Description: "Time in seconds to wait before retrying to start a PPPoE discovery, 0 means no timeout.",
				Optional:    true,
				Computed:    true,
			},
			"disconnect_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10000),

				Description: "Time in milliseconds to wait before sending a notification that this interface is down or disconnected.",
				Optional:    true,
				Computed:    true,
			},
			"distance": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.",
				Optional:    true,
				Computed:    true,
			},
			"dns_server_override": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use DNS acquired by DHCP or PPPoE.",
				Optional:    true,
				Computed:    true,
			},
			"drop_fragment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable drop fragment packets.",
				Optional:    true,
				Computed:    true,
			},
			"drop_overlapped_fragment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable drop overlapped fragment packets.",
				Optional:    true,
				Computed:    true,
			},
			"egress_cos": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "cos0", "cos1", "cos2", "cos3", "cos4", "cos5", "cos6", "cos7"}, false),

				Description: "Override outgoing CoS in user VLAN tag.",
				Optional:    true,
				Computed:    true,
			},
			"egress_queues": {
				Type:        schema.TypeList,
				Description: "Configure queues of NP port on egress path.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cos0": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "CoS profile name for CoS 0.",
							Optional:    true,
							Computed:    true,
						},
						"cos1": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "CoS profile name for CoS 1.",
							Optional:    true,
							Computed:    true,
						},
						"cos2": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "CoS profile name for CoS 2.",
							Optional:    true,
							Computed:    true,
						},
						"cos3": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "CoS profile name for CoS 3.",
							Optional:    true,
							Computed:    true,
						},
						"cos4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "CoS profile name for CoS 4.",
							Optional:    true,
							Computed:    true,
						},
						"cos5": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "CoS profile name for CoS 5.",
							Optional:    true,
							Computed:    true,
						},
						"cos6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "CoS profile name for CoS 6.",
							Optional:    true,
							Computed:    true,
						},
						"cos7": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "CoS profile name for CoS 7.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"egress_shaping_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Outgoing traffic shaping profile.",
				Optional:    true,
				Computed:    true,
			},
			"eip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "External IP.",
				Optional:    true,
				Computed:    true,
			},
			"estimated_downstream_bandwidth": {
				Type: schema.TypeInt,

				Description: "Estimated maximum downstream bandwidth (kbps). Used to estimate link utilization.",
				Optional:    true,
				Computed:    true,
			},
			"estimated_upstream_bandwidth": {
				Type: schema.TypeInt,

				Description: "Estimated maximum upstream bandwidth (kbps). Used to estimate link utilization.",
				Optional:    true,
				Computed:    true,
			},
			"explicit_ftp_proxy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the explicit FTP proxy on this interface.",
				Optional:    true,
				Computed:    true,
			},
			"explicit_web_proxy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the explicit web proxy on this interface.",
				Optional:    true,
				Computed:    true,
			},
			"external": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable identifying the interface as an external interface (which usually means it's connected to the Internet).",
				Optional:    true,
				Computed:    true,
			},
			"fail_action_on_extender": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"soft-restart", "hard-restart", "reboot"}, false),

				Description: "Action on extender when interface fail .",
				Optional:    true,
				Computed:    true,
			},
			"fail_alert_interfaces": {
				Type:        schema.TypeList,
				Description: "Names of the FortiGate interfaces to which the link failure alert is sent.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Names of the non-virtual interface.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"fail_alert_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"link-failed-signal", "link-down"}, false),

				Description: "Select link-failed-signal or link-down method to alert about a failed link.",
				Optional:    true,
				Computed:    true,
			},
			"fail_detect": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable fail detection features for this interface.",
				Optional:    true,
				Computed:    true,
			},
			"fail_detect_option": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Options for detecting that this interface has failed.",
				Optional:         true,
				Computed:         true,
			},
			"fortilink": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable FortiLink to dedicate this interface to manage other Fortinet devices.",
				Optional:    true,
				Computed:    true,
			},
			"fortilink_backup_link": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "fortilink split interface backup link.",
				Optional:    true,
				Computed:    true,
			},
			"fortilink_neighbor_detect": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"lldp", "fortilink"}, false),

				Description: "Protocol for FortiGate neighbor discovery.",
				Optional:    true,
				Computed:    true,
			},
			"fortilink_split_interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiLink split interface to connect member link to different FortiSwitch in stack for uplink redundancy.",
				Optional:    true,
				Computed:    true,
			},
			"fortilink_stacking": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiLink switch-stacking on this interface.",
				Optional:    true,
				Computed:    true,
			},
			"forward_domain": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2147483647),

				Description: "Transparent mode forward domain.",
				Optional:    true,
				Computed:    true,
			},
			"gwdetect": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable detect gateway alive for first.",
				Optional:    true,
				Computed:    true,
			},
			"ha_priority": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 50),

				Description: "HA election priority for the PING server.",
				Optional:    true,
				Computed:    true,
			},
			"icmp_accept_redirect": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable ICMP accept redirect.",
				Optional:    true,
				Computed:    true,
			},
			"icmp_send_redirect": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable sending of ICMP redirects.",
				Optional:    true,
				Computed:    true,
			},
			"ident_accept": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable authentication for this interface.",
				Optional:    true,
				Computed:    true,
			},
			"idle_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32767),

				Description: "PPPoE auto disconnect after idle timeout seconds, 0 means no timeout.",
				Optional:    true,
				Computed:    true,
			},
			"inbandwidth": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),

				Description: "Bandwidth limit for incoming traffic (0 - 16776000 kbps), 0 means unlimited.",
				Optional:    true,
				Computed:    true,
			},
			"ingress_cos": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "cos0", "cos1", "cos2", "cos3", "cos4", "cos5", "cos6", "cos7"}, false),

				Description: "Override incoming CoS in user VLAN tag on VLAN interface or assign a priority VLAN tag on physical interface.",
				Optional:    true,
				Computed:    true,
			},
			"ingress_shaping_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Incoming traffic shaping profile.",
				Optional:    true,
				Computed:    true,
			},
			"ingress_spillover_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),

				Description: "Ingress Spillover threshold (0 - 16776000 kbps), 0 means unlimited.",
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
			"internal": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Implicitly created.",
				Optional:    true,
				Computed:    true,
			},
			"ip": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4ClassnetHost,

				Description: "Interface IPv4 address and subnet mask, syntax: X.X.X.X/24.",
				Optional:    true,
				Computed:    true,
			},
			"ip_managed_by_fortiipam": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable automatic IP address assignment of this interface by FortiIPAM.",
				Optional:    true,
				Computed:    true,
			},
			"ipmac": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IP/MAC binding.",
				Optional:    true,
				Computed:    true,
			},
			"ips_sniffer_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the use of this interface as a one-armed sniffer.",
				Optional:    true,
				Computed:    true,
			},
			"ipunnumbered": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Unnumbered IP used for PPPoE interfaces for which no unique local address is provided.",
				Optional:    true,
				Computed:    true,
			},
			"ipv6": {
				Type:        schema.TypeList,
				Description: "IPv6 of interface.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"autoconf": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable address auto config.",
							Optional:    true,
							Computed:    true,
						},
						"cli_conn6_status": {
							Type: schema.TypeInt,

							Description: "CLI IPv6 connection status.",
							Optional:    true,
							Computed:    true,
						},
						"dhcp6_client_options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "DHCPv6 client options.",
							Optional:         true,
							Computed:         true,
						},
						"dhcp6_iapd_list": {
							Type:        schema.TypeList,
							Description: "DHCPv6 IA-PD list",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"iaid": {
										Type: schema.TypeInt,

										Description: "Identity association identifier.",
										Optional:    true,
										Computed:    true,
									},
									"prefix_hint": {
										Type:             schema.TypeString,
										ValidateFunc:     validators.FortiValidateIPv6Network,
										DiffSuppressFunc: suppressors.DiffCidrEqual,
										Description:      "DHCPv6 prefix that will be used as a hint to the upstream DHCPv6 server.",
										Optional:         true,
										Computed:         true,
									},
									"prefix_hint_plt": {
										Type: schema.TypeInt,

										Description: "DHCPv6 prefix hint preferred life time (sec), 0 means unlimited lease time.",
										Optional:    true,
										Computed:    true,
									},
									"prefix_hint_vlt": {
										Type: schema.TypeInt,

										Description: "DHCPv6 prefix hint valid life time (sec).",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"dhcp6_information_request": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable DHCPv6 information request.",
							Optional:    true,
							Computed:    true,
						},
						"dhcp6_prefix_delegation": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable DHCPv6 prefix delegation.",
							Optional:    true,
							Computed:    true,
						},
						"dhcp6_prefix_hint": {
							Type:             schema.TypeString,
							ValidateFunc:     validators.FortiValidateIPv6Network,
							DiffSuppressFunc: suppressors.DiffCidrEqual,
							Description:      "DHCPv6 prefix that will be used as a hint to the upstream DHCPv6 server.",
							Optional:         true,
							Computed:         true,
						},
						"dhcp6_prefix_hint_plt": {
							Type: schema.TypeInt,

							Description: "DHCPv6 prefix hint preferred life time (sec), 0 means unlimited lease time.",
							Optional:    true,
							Computed:    true,
						},
						"dhcp6_prefix_hint_vlt": {
							Type: schema.TypeInt,

							Description: "DHCPv6 prefix hint valid life time (sec).",
							Optional:    true,
							Computed:    true,
						},
						"dhcp6_relay_ip": {
							Type: schema.TypeString,

							Description: "DHCPv6 relay IP address.",
							Optional:    true,
							Computed:    true,
						},
						"dhcp6_relay_service": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable DHCPv6 relay.",
							Optional:    true,
							Computed:    true,
						},
						"dhcp6_relay_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"regular"}, false),

							Description: "DHCPv6 relay type.",
							Optional:    true,
							Computed:    true,
						},
						"icmp6_send_redirect": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sending of ICMPv6 redirects.",
							Optional:    true,
							Computed:    true,
						},
						"interface_identifier": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "IPv6 interface identifier.",
							Optional:         true,
							Computed:         true,
						},
						"ip6_address": {
							Type:             schema.TypeString,
							ValidateFunc:     validators.FortiValidateIPv6Prefix,
							DiffSuppressFunc: suppressors.DiffCidrEqual,
							Description:      "Primary IPv6 address prefix, syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx",
							Optional:         true,
							Computed:         true,
						},
						"ip6_allowaccess": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Allow management access to the interface.",
							Optional:         true,
							Computed:         true,
						},
						"ip6_default_life": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 9000),

							Description: "Default life (sec).",
							Optional:    true,
							Computed:    true,
						},
						"ip6_delegated_prefix_iaid": {
							Type: schema.TypeInt,

							Description: "IAID of obtained delegated-prefix from the upstream interface.",
							Optional:    true,
							Computed:    true,
						},
						"ip6_delegated_prefix_list": {
							Type:        schema.TypeList,
							Description: "Advertised IPv6 delegated prefix list.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"autonomous_flag": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable the autonomous flag.",
										Optional:    true,
										Computed:    true,
									},
									"delegated_prefix_iaid": {
										Type: schema.TypeInt,

										Description: "IAID of obtained delegated-prefix from the upstream interface.",
										Optional:    true,
										Computed:    true,
									},
									"onlink_flag": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable the onlink flag.",
										Optional:    true,
										Computed:    true,
									},
									"prefix_id": {
										Type: schema.TypeInt,

										Description: "Prefix ID.",
										Optional:    true,
										Computed:    true,
									},
									"rdnss": {
										Type: schema.TypeString,

										Description: "Recursive DNS server option.",
										Optional:    true,
										Computed:    true,
									},
									"rdnss_service": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"delegated", "default", "specify"}, false),

										Description: "Recursive DNS service option.",
										Optional:    true,
										Computed:    true,
									},
									"subnet": {
										Type:             schema.TypeString,
										ValidateFunc:     validators.FortiValidateIPv6Network,
										DiffSuppressFunc: suppressors.DiffCidrEqual,
										Description:      " Add subnet ID to routing prefix.",
										Optional:         true,
										Computed:         true,
									},
									"upstream_interface": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 15),

										Description: "Name of the interface that provides delegated information.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"ip6_dns_server_override": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable using the DNS server acquired by DHCP.",
							Optional:    true,
							Computed:    true,
						},
						"ip6_extra_addr": {
							Type:        schema.TypeList,
							Description: "Extra IPv6 address prefixes of interface.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": {
										Type:             schema.TypeString,
										ValidateFunc:     validators.FortiValidateIPv6Prefix,
										DiffSuppressFunc: suppressors.DiffCidrEqual,
										Description:      "IPv6 address prefix.",
										Optional:         true,
										Computed:         true,
									},
								},
							},
						},
						"ip6_hop_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Hop limit (0 means unspecified).",
							Optional:    true,
							Computed:    true,
						},
						"ip6_link_mtu": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1280, 16000),

							Description: "IPv6 link MTU.",
							Optional:    true,
							Computed:    true,
						},
						"ip6_manage_flag": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the managed flag.",
							Optional:    true,
							Computed:    true,
						},
						"ip6_max_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(4, 1800),

							Description: "IPv6 maximum interval (4 to 1800 sec).",
							Optional:    true,
							Computed:    true,
						},
						"ip6_min_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(3, 1350),

							Description: "IPv6 minimum interval (3 to 1350 sec).",
							Optional:    true,
							Computed:    true,
						},
						"ip6_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"static", "dhcp", "pppoe", "delegated"}, false),

							Description: "Addressing mode (static, DHCP, delegated).",
							Optional:    true,
							Computed:    true,
						},
						"ip6_other_flag": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the other IPv6 flag.",
							Optional:    true,
							Computed:    true,
						},
						"ip6_prefix_list": {
							Type:        schema.TypeList,
							Description: "Advertised prefix list.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"autonomous_flag": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable the autonomous flag.",
										Optional:    true,
										Computed:    true,
									},
									"dnssl": {
										Type:        schema.TypeList,
										Description: "DNS search list option.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"domain": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 79),

													Description: "Domain name.",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"onlink_flag": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable the onlink flag.",
										Optional:    true,
										Computed:    true,
									},
									"preferred_life_time": {
										Type: schema.TypeInt,

										Description: "Preferred life time (sec).",
										Optional:    true,
										Computed:    true,
									},
									"prefix": {
										Type:             schema.TypeString,
										ValidateFunc:     validators.FortiValidateIPv6Network,
										DiffSuppressFunc: suppressors.DiffCidrEqual,
										Description:      "IPv6 prefix.",
										Optional:         true,
										Computed:         true,
									},
									"rdnss": {
										Type: schema.TypeString,

										Description: "Recursive DNS server option.",
										Optional:    true,
										Computed:    true,
									},
									"valid_life_time": {
										Type: schema.TypeInt,

										Description: "Valid life time (sec).",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"ip6_prefix_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"dhcp6", "ra"}, false),

							Description: "Assigning a prefix from DHCP or RA.",
							Optional:    true,
							Computed:    true,
						},
						"ip6_reachable_time": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 3600000),

							Description: "IPv6 reachable time (milliseconds; 0 means unspecified).",
							Optional:    true,
							Computed:    true,
						},
						"ip6_retrans_time": {
							Type: schema.TypeInt,

							Description: "IPv6 retransmit time (milliseconds; 0 means unspecified).",
							Optional:    true,
							Computed:    true,
						},
						"ip6_send_adv": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sending advertisements about the interface.",
							Optional:    true,
							Computed:    true,
						},
						"ip6_subnet": {
							Type:             schema.TypeString,
							ValidateFunc:     validators.FortiValidateIPv6Prefix,
							DiffSuppressFunc: suppressors.DiffCidrEqual,
							Description:      " Subnet to routing prefix, syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx",
							Optional:         true,
							Computed:         true,
						},
						"ip6_upstream_interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Interface name providing delegated information.",
							Optional:    true,
							Computed:    true,
						},
						"nd_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Neighbor discovery certificate.",
							Optional:    true,
							Computed:    true,
						},
						"nd_cga_modifier": {
							Type: schema.TypeString,

							Description: "Neighbor discovery CGA modifier.",
							Optional:    true,
							Computed:    true,
						},
						"nd_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"basic", "SEND-compatible"}, false),

							Description: "Neighbor discovery mode.",
							Optional:    true,
							Computed:    true,
						},
						"nd_security_level": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),

							Description: "Neighbor discovery security level (0 - 7; 0 = least secure, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"nd_timestamp_delta": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 3600),

							Description: "Neighbor discovery timestamp delta value (1 - 3600 sec; default = 300).",
							Optional:    true,
							Computed:    true,
						},
						"nd_timestamp_fuzz": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 60),

							Description: "Neighbor discovery timestamp fuzz factor (1 - 60 sec; default = 1).",
							Optional:    true,
							Computed:    true,
						},
						"ra_send_mtu": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable sending link MTU in RA packet.",
							Optional:    true,
							Computed:    true,
						},
						"unique_autoconf_addr": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable unique auto config address.",
							Optional:    true,
							Computed:    true,
						},
						"vrip6_link_local": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "Link-local IPv6 address of virtual router.",
							Optional:         true,
							Computed:         true,
						},
						"vrrp_virtual_mac6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable virtual MAC for VRRP.",
							Optional:    true,
							Computed:    true,
						},
						"vrrp6": {
							Type:        schema.TypeList,
							Description: "IPv6 VRRP configuration.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"accept_mode": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable accept mode.",
										Optional:    true,
										Computed:    true,
									},
									"adv_interval": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 255),

										Description: "Advertisement interval (1 - 255 seconds).",
										Optional:    true,
										Computed:    true,
									},
									"preempt": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable preempt mode.",
										Optional:    true,
										Computed:    true,
									},
									"priority": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 255),

										Description: "Priority of the virtual router (1 - 255).",
										Optional:    true,
										Computed:    true,
									},
									"start_time": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 255),

										Description: "Startup time (1 - 255 seconds).",
										Optional:    true,
										Computed:    true,
									},
									"status": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable VRRP.",
										Optional:    true,
										Computed:    true,
									},
									"vrdst6": {
										Type:             schema.TypeString,
										ValidateFunc:     validation.IsIPv6Address,
										DiffSuppressFunc: suppressors.DiffIPEqual,
										Description:      "Monitor the route to this destination.",
										Optional:         true,
										Computed:         true,
									},
									"vrgrp": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),

										Description: "VRRP group ID (1 - 65535).",
										Optional:    true,
										Computed:    true,
									},
									"vrid": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 255),

										Description: "Virtual router identifier (1 - 255).",
										Optional:    true,
										Computed:    true,
									},
									"vrip6": {
										Type:             schema.TypeString,
										ValidateFunc:     validation.IsIPv6Address,
										DiffSuppressFunc: suppressors.DiffIPEqual,
										Description:      "IPv6 address of the virtual router.",
										Optional:         true,
										Computed:         true,
									},
								},
							},
						},
					},
				},
			},
			"l2forward": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable l2 forwarding.",
				Optional:    true,
				Computed:    true,
			},
			"lacp_ha_slave": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "LACP HA slave.",
				Optional:    true,
				Computed:    true,
			},
			"lacp_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"static", "passive", "active"}, false),

				Description: "LACP mode.",
				Optional:    true,
				Computed:    true,
			},
			"lacp_speed": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"slow", "fast"}, false),

				Description: "How often the interface sends LACP messages.",
				Optional:    true,
				Computed:    true,
			},
			"lcp_echo_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32767),

				Description: "Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.",
				Optional:    true,
				Computed:    true,
			},
			"lcp_max_echo_fails": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32767),

				Description: "Maximum missed LCP echo messages before disconnect.",
				Optional:    true,
				Computed:    true,
			},
			"link_up_delay": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(50, 3600000),

				Description: "Number of milliseconds to wait before considering a link is up.",
				Optional:    true,
				Computed:    true,
			},
			"lldp_network_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "LLDP-MED network policy profile.",
				Optional:    true,
				Computed:    true,
			},
			"lldp_reception": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable", "vdom"}, false),

				Description: "Enable/disable Link Layer Discovery Protocol (LLDP) reception.",
				Optional:    true,
				Computed:    true,
			},
			"lldp_transmission": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable", "vdom"}, false),

				Description: "Enable/disable Link Layer Discovery Protocol (LLDP) transmission.",
				Optional:    true,
				Computed:    true,
			},
			"macaddr": {
				Type: schema.TypeString,

				Description: "Change the interface's MAC address.",
				Optional:    true,
				Computed:    true,
			},
			"managed_subnetwork_size": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"32", "64", "128", "256", "512", "1024", "2048", "4096", "8192", "16384", "32768", "65536"}, false),

				Description: "Number of IP addresses to be allocated by FortiIPAM and used by this FortiGate unit's DHCP server settings.",
				Optional:    true,
				Computed:    true,
			},
			"management_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4ClassnetHost,

				Description: "High Availability in-band management IP address of this interface.",
				Optional:    true,
				Computed:    true,
			},
			"measured_downstream_bandwidth": {
				Type: schema.TypeInt,

				Description: "Measured downstream bandwidth (kbps). ",
				Optional:    true,
				Computed:    true,
			},
			"measured_upstream_bandwidth": {
				Type: schema.TypeInt,

				Description: "Measured upstream bandwidth (kbps). ",
				Optional:    true,
				Computed:    true,
			},
			"mediatype": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"serdes-sfp", "sgmii-sfp", "serdes-copper-sfp"}, false),

				Description: "Select SFP media interface type",
				Optional:    true,
				Computed:    true,
			},
			"member": {
				Type:        schema.TypeList,
				Description: "Physical interfaces that belong to the aggregate or redundant interface.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Physical interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"min_links": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 32),

				Description: "Minimum number of aggregated ports that must be up.",
				Optional:    true,
				Computed:    true,
			},
			"min_links_down": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"operational", "administrative"}, false),

				Description: "Action to take when less than the configured minimum number of links are active.",
				Optional:    true,
				Computed:    true,
			},
			"mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"static", "dhcp", "pppoe"}, false),

				Description: "Addressing mode (static, DHCP, PPPoE).",
				Optional:    true,
				Computed:    true,
			},
			"monitor_bandwidth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable monitoring bandwidth on this interface.",
				Optional:    true,
				Computed:    true,
			},
			"mtu": {
				Type: schema.TypeInt,

				Description: "MTU value for this interface.",
				Optional:    true,
				Computed:    true,
			},
			"mtu_override": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to set a custom MTU for this interface.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Name.",
				ForceNew:    true,
				Required:    true,
			},
			"ndiscforward": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable NDISC forwarding.",
				Optional:    true,
				Computed:    true,
			},
			"netbios_forward": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable NETBIOS forwarding.",
				Optional:    true,
				Computed:    true,
			},
			"netflow_sampler": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "tx", "rx", "both"}, false),

				Description: "Enable/disable NetFlow on this interface and set the data that NetFlow collects (rx, tx, or both).",
				Optional:    true,
				Computed:    true,
			},
			"outbandwidth": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),

				Description: "Bandwidth limit for outgoing traffic (0 - 16776000 kbps), 0 means unlimited.",
				Optional:    true,
				Computed:    true,
			},
			"padt_retry_timeout": {
				Type: schema.TypeInt,

				Description: "PPPoE Active Discovery Terminate (PADT) used to terminate sessions after an idle time.",
				Optional:    true,
				Computed:    true,
			},
			"password": {
				Type: schema.TypeString,

				Description: "PPPoE account's password.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"ping_serv_status": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "PING server status.",
				Optional:    true,
				Computed:    true,
			},
			"polling_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "sFlow polling interval (1 - 255 sec).",
				Optional:    true,
				Computed:    true,
			},
			"pppoe_unnumbered_negotiate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable PPPoE unnumbered negotiation.",
				Optional:    true,
				Computed:    true,
			},
			"pptp_auth_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "pap", "chap", "mschapv1", "mschapv2"}, false),

				Description: "PPTP authentication type.",
				Optional:    true,
				Computed:    true,
			},
			"pptp_client": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable PPTP client.",
				Optional:    true,
				Computed:    true,
			},
			"pptp_password": {
				Type: schema.TypeString,

				Description: "PPTP password.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"pptp_server_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "PPTP server IP address.",
				Optional:    true,
				Computed:    true,
			},
			"pptp_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Idle timer in minutes (0 for disabled).",
				Optional:    true,
				Computed:    true,
			},
			"pptp_user": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "PPTP user name.",
				Optional:    true,
				Computed:    true,
			},
			"preserve_session_route": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable preservation of session route when dirty.",
				Optional:    true,
				Computed:    true,
			},
			"priority": {
				Type: schema.TypeInt,

				Description: "Priority of learned routes.",
				Optional:    true,
				Computed:    true,
			},
			"priority_override": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable fail back to higher priority port once recovered.",
				Optional:    true,
				Computed:    true,
			},
			"proxy_captive_portal": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable proxy captive portal on this interface.",
				Optional:    true,
				Computed:    true,
			},
			"redundant_interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Redundant interface.",
				Optional:    true,
				Computed:    true,
			},
			"remote_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4ClassnetHost,

				Description: "Remote IP address of tunnel.",
				Optional:    true,
				Computed:    true,
			},
			"replacemsg_override_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Replacement message override group.",
				Optional:    true,
				Computed:    true,
			},
			"ring_rx": {
				Type: schema.TypeInt,

				Description: "RX ring size.",
				Optional:    true,
				Computed:    true,
			},
			"ring_tx": {
				Type: schema.TypeInt,

				Description: "TX ring size.",
				Optional:    true,
				Computed:    true,
			},
			"role": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"lan", "wan", "dmz", "undefined"}, false),

				Description: "Interface role.",
				Optional:    true,
				Computed:    true,
			},
			"sample_direction": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"tx", "rx", "both"}, false),

				Description: "Data that NetFlow collects (rx, tx, or both).",
				Optional:    true,
				Computed:    true,
			},
			"sample_rate": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 99999),

				Description: "sFlow sample rate (10 - 99999).",
				Optional:    true,
				Computed:    true,
			},
			"secondary_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable adding a secondary IP to this interface.",
				Optional:    true,
				Computed:    true,
			},
			"secondaryip": {
				Type:        schema.TypeList,
				Description: "Second IP address of interface.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allowaccess": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Management access settings for the secondary IP address.",
							Optional:         true,
							Computed:         true,
						},
						"detectprotocol": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Protocols used to detect the server.",
							Optional:         true,
							Computed:         true,
						},
						"detectserver": {
							Type: schema.TypeString,

							Description: "Gateway's ping server for this IP.",
							Optional:    true,
							Computed:    true,
						},
						"gwdetect": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable detect gateway alive for first.",
							Optional:    true,
							Computed:    true,
						},
						"ha_priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 50),

							Description: "HA election priority for the PING server.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"ip": {
							Type:         schema.TypeString,
							ValidateFunc: validators.FortiValidateIPv4ClassnetHost,

							Description: "Secondary IP address of the interface.",
							Optional:    true,
							Computed:    true,
						},
						"ping_serv_status": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "PING server status.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"security_exempt_list": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of security-exempt-list.",
				Optional:    true,
				Computed:    true,
			},
			"security_external_logout": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "URL of external authentication logout server.",
				Optional:    true,
				Computed:    true,
			},
			"security_external_web": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),

				Description: "URL of external authentication web server.",
				Optional:    true,
				Computed:    true,
			},
			"security_groups": {
				Type:        schema.TypeList,
				Description: "User groups that can authenticate with the captive portal.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Names of user groups that can authenticate with the captive portal.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"security_mac_auth_bypass": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"mac-auth-only", "enable", "disable"}, false),

				Description: "Enable/disable MAC authentication bypass.",
				Optional:    true,
				Computed:    true,
			},
			"security_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "captive-portal", "802.1X"}, false),

				Description: "Turn on captive portal authentication for this interface.",
				Optional:    true,
				Computed:    true,
			},
			"security_redirect_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),

				Description: "URL redirection after disclaimer/authentication.",
				Optional:    true,
				Computed:    true,
			},
			"service_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "PPPoE service name.",
				Optional:    true,
				Computed:    true,
			},
			"sflow_sampler": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable sFlow on this interface.",
				Optional:    true,
				Computed:    true,
			},
			"snmp_index": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 2147483647),

				Description: "Permanent SNMP Index of the interface.",
				Optional:    true,
				Computed:    true,
			},
			"speed": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "10full", "10half", "100full", "100half", "1000full", "1000half", "1000auto"}, false),

				Description: "Interface speed. The default setting and the options available depend on the interface hardware.",
				Optional:    true,
				Computed:    true,
			},
			"spillover_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),

				Description: "Egress Spillover threshold (0 - 16776000 kbps), 0 means unlimited.",
				Optional:    true,
				Computed:    true,
			},
			"src_check": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable source IP check.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"up", "down"}, false),

				Description: "Bring the interface up or shut the interface down.",
				Optional:    true,
				Computed:    true,
			},
			"stpforward": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable STP forwarding.",
				Optional:    true,
				Computed:    true,
			},
			"stpforward_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"rpl-all-ext-id", "rpl-bridge-ext-id", "rpl-nothing"}, false),

				Description: "Configure STP forwarding mode.",
				Optional:    true,
				Computed:    true,
			},
			"subst": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to always send packets from this interface to a destination MAC address.",
				Optional:    true,
				Computed:    true,
			},
			"substitute_dst_mac": {
				Type: schema.TypeString,

				Description: "Destination MAC address that all packets are sent to from this interface.",
				Optional:    true,
				Computed:    true,
			},
			"swc_first_create": {
				Type: schema.TypeInt,

				Description: "Initial create for switch-controller VLANs.",
				Optional:    true,
				Computed:    true,
			},
			"swc_vlan": {
				Type: schema.TypeInt,

				Description: "Creation status for switch-controller VLANs.",
				Optional:    true,
				Computed:    true,
			},
			"switch": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Contained in switch.",
				Optional:    true,
				Computed:    true,
			},
			"switch_controller_access_vlan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Block FortiSwitch port-to-port traffic.",
				Optional:    true,
				Computed:    true,
			},
			"switch_controller_arp_inspection": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiSwitch ARP inspection.",
				Optional:    true,
				Computed:    true,
			},
			"switch_controller_dhcp_snooping": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Switch controller DHCP snooping.",
				Optional:    true,
				Computed:    true,
			},
			"switch_controller_dhcp_snooping_option82": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Switch controller DHCP snooping option82.",
				Optional:    true,
				Computed:    true,
			},
			"switch_controller_dhcp_snooping_verify_mac": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Switch controller DHCP snooping verify MAC.",
				Optional:    true,
				Computed:    true,
			},
			"switch_controller_dynamic": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Integrated FortiLink settings for managed FortiSwitch.",
				Optional:    true,
				Computed:    true,
			},
			"switch_controller_feature": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "default-vlan", "quarantine", "rspan", "voice", "video", "nac", "nac-segment"}, false),

				Description: "Interface's purpose when assigning traffic (read only).",
				Optional:    true,
				Computed:    true,
			},
			"switch_controller_igmp_snooping": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Switch controller IGMP snooping.",
				Optional:    true,
				Computed:    true,
			},
			"switch_controller_igmp_snooping_fast_leave": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Switch controller IGMP snooping fast-leave.",
				Optional:    true,
				Computed:    true,
			},
			"switch_controller_igmp_snooping_proxy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Switch controller IGMP snooping proxy.",
				Optional:    true,
				Computed:    true,
			},
			"switch_controller_iot_scanning": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable managed FortiSwitch IoT scanning.",
				Optional:    true,
				Computed:    true,
			},
			"switch_controller_learning_limit": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),

				Description: "Limit the number of dynamic MAC addresses on this VLAN (1 - 128, 0 = no limit, default).",
				Optional:    true,
				Computed:    true,
			},
			"switch_controller_mgmt_vlan": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 4094),

				Description: "VLAN to use for FortiLink management purposes.",
				Optional:    true,
				Computed:    true,
			},
			"switch_controller_nac": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Integrated FortiLink settings for managed FortiSwitch.",
				Optional:    true,
				Computed:    true,
			},
			"switch_controller_rspan_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Stop Layer2 MAC learning and interception of BPDUs and other packets on this interface.",
				Optional:    true,
				Computed:    true,
			},
			"switch_controller_source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"outbound", "fixed"}, false),

				Description: "Source IP address used in FortiLink over L3 connections.",
				Optional:    true,
				Computed:    true,
			},
			"switch_controller_traffic_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Switch controller traffic policy for the VLAN.",
				Optional:    true,
				Computed:    true,
			},
			"system_id": {
				Type: schema.TypeString,

				Description: "Define a system ID for the aggregate interface.",
				Optional:    true,
				Computed:    true,
			},
			"system_id_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "user"}, false),

				Description: "Method in which system ID is generated.",
				Optional:    true,
				Computed:    true,
			},
			"tagging": {
				Type:        schema.TypeList,
				Description: "Config object tagging.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Tag category.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Tagging entry name.",
							Optional:    true,
							Computed:    true,
						},
						"tags": {
							Type:        schema.TypeList,
							Description: "Tags.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Tag name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"tcp_mss": {
				Type: schema.TypeInt,

				Description: "TCP maximum segment size. 0 means do not change segment size.",
				Optional:    true,
				Computed:    true,
			},
			"trust_ip_1": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4ClassnetAny,

				Description: "Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).",
				Optional:    true,
				Computed:    true,
			},
			"trust_ip_2": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4ClassnetAny,

				Description: "Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).",
				Optional:    true,
				Computed:    true,
			},
			"trust_ip_3": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4ClassnetAny,

				Description: "Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).",
				Optional:    true,
				Computed:    true,
			},
			"trust_ip6_1": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Prefix,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).",
				Optional:         true,
				Computed:         true,
			},
			"trust_ip6_2": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Prefix,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).",
				Optional:         true,
				Computed:         true,
			},
			"trust_ip6_3": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Prefix,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).",
				Optional:         true,
				Computed:         true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"physical", "vlan", "aggregate", "redundant", "tunnel", "vdom-link", "loopback", "switch", "vap-switch", "wl-mesh", "fext-wan", "vxlan", "geneve", "hdlc", "switch-vlan", "emac-vlan", "ssl", "lan-extension"}, false),

				Description: "Interface type.",
				Optional:    true,
				Computed:    true,
			},
			"username": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "Username of the PPPoE account, provided by your ISP.",
				Optional:    true,
				Computed:    true,
			},
			"vdom": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "Interface is in this virtual domain (VDOM).",
				Optional:    true,
				Computed:    true,
			},
			"vindex": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Switch control interface VLAN ID.",
				Optional:    true,
				Computed:    true,
			},
			"vlan_protocol": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"8021q", "8021ad"}, false),

				Description: "Ethernet protocol of VLAN.",
				Optional:    true,
				Computed:    true,
			},
			"vlanforward": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable traffic forwarding between VLANs on this interface.",
				Optional:    true,
				Computed:    true,
			},
			"vlanid": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 4094),

				Description: "VLAN ID (1 - 4094).",
				Optional:    true,
				Computed:    true,
			},
			"vrf": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 31),

				Description: "Virtual Routing Forwarding ID.",
				Optional:    true,
				Computed:    true,
			},
			"vrrp": {
				Type:        schema.TypeList,
				Description: "VRRP configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accept_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable accept mode.",
							Optional:    true,
							Computed:    true,
						},
						"adv_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),

							Description: "Advertisement interval (1 - 255 seconds).",
							Optional:    true,
							Computed:    true,
						},
						"ignore_default_route": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable ignoring of default route when checking destination.",
							Optional:    true,
							Computed:    true,
						},
						"preempt": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable preempt mode.",
							Optional:    true,
							Computed:    true,
						},
						"priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),

							Description: "Priority of the virtual router (1 - 255).",
							Optional:    true,
							Computed:    true,
						},
						"proxy_arp": {
							Type:        schema.TypeList,
							Description: "VRRP Proxy ARP configuration.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type: schema.TypeInt,

										Description: "ID.",
										Optional:    true,
										Computed:    true,
									},
									"ip": {
										Type: schema.TypeString,

										Description: "Set IP addresses of proxy ARP.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"start_time": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),

							Description: "Startup time (1 - 255 seconds).",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable this VRRP configuration.",
							Optional:    true,
							Computed:    true,
						},
						"version": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"2", "3"}, false),

							Description: "VRRP version.",
							Optional:    true,
							Computed:    true,
						},
						"vrdst": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Monitor the route to this destination.",
							Optional:    true,
							Computed:    true,
						},
						"vrdst_priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 254),

							Description: "Priority of the virtual router when the virtual router destination becomes unreachable (0 - 254).",
							Optional:    true,
							Computed:    true,
						},
						"vrgrp": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "VRRP group ID (1 - 65535).",
							Optional:    true,
							Computed:    true,
						},
						"vrid": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),

							Description: "Virtual router identifier (1 - 255).",
							Optional:    true,
							Computed:    true,
						},
						"vrip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "IP address of the virtual router.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"vrrp_virtual_mac": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of virtual MAC for VRRP.",
				Optional:    true,
				Computed:    true,
			},
			"wccp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable WCCP on this interface. Used for encapsulated WCCP communication between WCCP clients and servers.",
				Optional:    true,
				Computed:    true,
			},
			"weight": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Default weight for static routes (if route has no weight configured).",
				Optional:    true,
				Computed:    true,
			},
			"wins_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "WINS server IP.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemInterfaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := ""
	key := "name"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating SystemInterface resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemInterface(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemInterface(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemInterface")
	}

	return resourceSystemInterfaceRead(ctx, d, meta)
}

func resourceSystemInterfaceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemInterface(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemInterface(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemInterface resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemInterface")
	}

	return resourceSystemInterfaceRead(ctx, d, meta)
}

func resourceSystemInterfaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemInterface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemInterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemInterface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemInterface resource: %v", err)
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

	diags := refreshObjectSystemInterface(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemInterfaceClientOptions(v *[]models.SystemInterfaceClientOptions, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Code; tmp != nil {
				v["code"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			if tmp := cfg.Value; tmp != nil {
				v["value"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemInterfaceDhcpSnoopingServerList(v *[]models.SystemInterfaceDhcpSnoopingServerList, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.ServerIp; tmp != nil {
				v["server_ip"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemInterfaceEgressQueues(v *[]models.SystemInterfaceEgressQueues, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Cos0; tmp != nil {
				v["cos0"] = *tmp
			}

			if tmp := cfg.Cos1; tmp != nil {
				v["cos1"] = *tmp
			}

			if tmp := cfg.Cos2; tmp != nil {
				v["cos2"] = *tmp
			}

			if tmp := cfg.Cos3; tmp != nil {
				v["cos3"] = *tmp
			}

			if tmp := cfg.Cos4; tmp != nil {
				v["cos4"] = *tmp
			}

			if tmp := cfg.Cos5; tmp != nil {
				v["cos5"] = *tmp
			}

			if tmp := cfg.Cos6; tmp != nil {
				v["cos6"] = *tmp
			}

			if tmp := cfg.Cos7; tmp != nil {
				v["cos7"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemInterfaceFailAlertInterfaces(v *[]models.SystemInterfaceFailAlertInterfaces, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenSystemInterfaceIpv6(v *[]models.SystemInterfaceIpv6, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Autoconf; tmp != nil {
				v["autoconf"] = *tmp
			}

			if tmp := cfg.CliConn6Status; tmp != nil {
				v["cli_conn6_status"] = *tmp
			}

			if tmp := cfg.Dhcp6ClientOptions; tmp != nil {
				v["dhcp6_client_options"] = *tmp
			}

			if tmp := cfg.Dhcp6IapdList; tmp != nil {
				v["dhcp6_iapd_list"] = flattenSystemInterfaceIpv6Dhcp6IapdList(tmp, sort)
			}

			if tmp := cfg.Dhcp6InformationRequest; tmp != nil {
				v["dhcp6_information_request"] = *tmp
			}

			if tmp := cfg.Dhcp6PrefixDelegation; tmp != nil {
				v["dhcp6_prefix_delegation"] = *tmp
			}

			if tmp := cfg.Dhcp6PrefixHint; tmp != nil {
				v["dhcp6_prefix_hint"] = *tmp
			}

			if tmp := cfg.Dhcp6PrefixHintPlt; tmp != nil {
				v["dhcp6_prefix_hint_plt"] = *tmp
			}

			if tmp := cfg.Dhcp6PrefixHintVlt; tmp != nil {
				v["dhcp6_prefix_hint_vlt"] = *tmp
			}

			if tmp := cfg.Dhcp6RelayIp; tmp != nil {
				v["dhcp6_relay_ip"] = *tmp
			}

			if tmp := cfg.Dhcp6RelayService; tmp != nil {
				v["dhcp6_relay_service"] = *tmp
			}

			if tmp := cfg.Dhcp6RelayType; tmp != nil {
				v["dhcp6_relay_type"] = *tmp
			}

			if tmp := cfg.Icmp6SendRedirect; tmp != nil {
				v["icmp6_send_redirect"] = *tmp
			}

			if tmp := cfg.InterfaceIdentifier; tmp != nil {
				v["interface_identifier"] = *tmp
			}

			if tmp := cfg.Ip6Address; tmp != nil {
				v["ip6_address"] = *tmp
			}

			if tmp := cfg.Ip6Allowaccess; tmp != nil {
				v["ip6_allowaccess"] = *tmp
			}

			if tmp := cfg.Ip6DefaultLife; tmp != nil {
				v["ip6_default_life"] = *tmp
			}

			if tmp := cfg.Ip6DelegatedPrefixIaid; tmp != nil {
				v["ip6_delegated_prefix_iaid"] = *tmp
			}

			if tmp := cfg.Ip6DelegatedPrefixList; tmp != nil {
				v["ip6_delegated_prefix_list"] = flattenSystemInterfaceIpv6Ip6DelegatedPrefixList(tmp, sort)
			}

			if tmp := cfg.Ip6DnsServerOverride; tmp != nil {
				v["ip6_dns_server_override"] = *tmp
			}

			if tmp := cfg.Ip6ExtraAddr; tmp != nil {
				v["ip6_extra_addr"] = flattenSystemInterfaceIpv6Ip6ExtraAddr(tmp, sort)
			}

			if tmp := cfg.Ip6HopLimit; tmp != nil {
				v["ip6_hop_limit"] = *tmp
			}

			if tmp := cfg.Ip6LinkMtu; tmp != nil {
				v["ip6_link_mtu"] = *tmp
			}

			if tmp := cfg.Ip6ManageFlag; tmp != nil {
				v["ip6_manage_flag"] = *tmp
			}

			if tmp := cfg.Ip6MaxInterval; tmp != nil {
				v["ip6_max_interval"] = *tmp
			}

			if tmp := cfg.Ip6MinInterval; tmp != nil {
				v["ip6_min_interval"] = *tmp
			}

			if tmp := cfg.Ip6Mode; tmp != nil {
				v["ip6_mode"] = *tmp
			}

			if tmp := cfg.Ip6OtherFlag; tmp != nil {
				v["ip6_other_flag"] = *tmp
			}

			if tmp := cfg.Ip6PrefixList; tmp != nil {
				v["ip6_prefix_list"] = flattenSystemInterfaceIpv6Ip6PrefixList(tmp, sort)
			}

			if tmp := cfg.Ip6PrefixMode; tmp != nil {
				v["ip6_prefix_mode"] = *tmp
			}

			if tmp := cfg.Ip6ReachableTime; tmp != nil {
				v["ip6_reachable_time"] = *tmp
			}

			if tmp := cfg.Ip6RetransTime; tmp != nil {
				v["ip6_retrans_time"] = *tmp
			}

			if tmp := cfg.Ip6SendAdv; tmp != nil {
				v["ip6_send_adv"] = *tmp
			}

			if tmp := cfg.Ip6Subnet; tmp != nil {
				v["ip6_subnet"] = *tmp
			}

			if tmp := cfg.Ip6UpstreamInterface; tmp != nil {
				v["ip6_upstream_interface"] = *tmp
			}

			if tmp := cfg.NdCert; tmp != nil {
				v["nd_cert"] = *tmp
			}

			if tmp := cfg.NdCgaModifier; tmp != nil {
				v["nd_cga_modifier"] = *tmp
			}

			if tmp := cfg.NdMode; tmp != nil {
				v["nd_mode"] = *tmp
			}

			if tmp := cfg.NdSecurityLevel; tmp != nil {
				v["nd_security_level"] = *tmp
			}

			if tmp := cfg.NdTimestampDelta; tmp != nil {
				v["nd_timestamp_delta"] = *tmp
			}

			if tmp := cfg.NdTimestampFuzz; tmp != nil {
				v["nd_timestamp_fuzz"] = *tmp
			}

			if tmp := cfg.RaSendMtu; tmp != nil {
				v["ra_send_mtu"] = *tmp
			}

			if tmp := cfg.UniqueAutoconfAddr; tmp != nil {
				v["unique_autoconf_addr"] = *tmp
			}

			if tmp := cfg.Vrip6_link_local; tmp != nil {
				v["vrip6_link_local"] = *tmp
			}

			if tmp := cfg.VrrpVirtualMac6; tmp != nil {
				v["vrrp_virtual_mac6"] = *tmp
			}

			if tmp := cfg.Vrrp6; tmp != nil {
				v["vrrp6"] = flattenSystemInterfaceIpv6Vrrp6(tmp, sort)
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemInterfaceIpv6Dhcp6IapdList(v *[]models.SystemInterfaceIpv6Dhcp6IapdList, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Iaid; tmp != nil {
				v["iaid"] = *tmp
			}

			if tmp := cfg.PrefixHint; tmp != nil {
				v["prefix_hint"] = *tmp
			}

			if tmp := cfg.PrefixHintPlt; tmp != nil {
				v["prefix_hint_plt"] = *tmp
			}

			if tmp := cfg.PrefixHintVlt; tmp != nil {
				v["prefix_hint_vlt"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "iaid")
	}

	return flat
}

func flattenSystemInterfaceIpv6Ip6DelegatedPrefixList(v *[]models.SystemInterfaceIpv6Ip6DelegatedPrefixList, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AutonomousFlag; tmp != nil {
				v["autonomous_flag"] = *tmp
			}

			if tmp := cfg.DelegatedPrefixIaid; tmp != nil {
				v["delegated_prefix_iaid"] = *tmp
			}

			if tmp := cfg.OnlinkFlag; tmp != nil {
				v["onlink_flag"] = *tmp
			}

			if tmp := cfg.PrefixId; tmp != nil {
				v["prefix_id"] = *tmp
			}

			if tmp := cfg.Rdnss; tmp != nil {
				v["rdnss"] = *tmp
			}

			if tmp := cfg.RdnssService; tmp != nil {
				v["rdnss_service"] = *tmp
			}

			if tmp := cfg.Subnet; tmp != nil {
				v["subnet"] = *tmp
			}

			if tmp := cfg.UpstreamInterface; tmp != nil {
				v["upstream_interface"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "prefix_id")
	}

	return flat
}

func flattenSystemInterfaceIpv6Ip6ExtraAddr(v *[]models.SystemInterfaceIpv6Ip6ExtraAddr, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Prefix; tmp != nil {
				v["prefix"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "prefix")
	}

	return flat
}

func flattenSystemInterfaceIpv6Ip6PrefixList(v *[]models.SystemInterfaceIpv6Ip6PrefixList, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AutonomousFlag; tmp != nil {
				v["autonomous_flag"] = *tmp
			}

			if tmp := cfg.Dnssl; tmp != nil {
				v["dnssl"] = flattenSystemInterfaceIpv6Ip6PrefixListDnssl(tmp, sort)
			}

			if tmp := cfg.OnlinkFlag; tmp != nil {
				v["onlink_flag"] = *tmp
			}

			if tmp := cfg.PreferredLifeTime; tmp != nil {
				v["preferred_life_time"] = *tmp
			}

			if tmp := cfg.Prefix; tmp != nil {
				v["prefix"] = *tmp
			}

			if tmp := cfg.Rdnss; tmp != nil {
				v["rdnss"] = *tmp
			}

			if tmp := cfg.ValidLifeTime; tmp != nil {
				v["valid_life_time"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "prefix")
	}

	return flat
}

func flattenSystemInterfaceIpv6Ip6PrefixListDnssl(v *[]models.SystemInterfaceIpv6Ip6PrefixListDnssl, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Domain; tmp != nil {
				v["domain"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "domain")
	}

	return flat
}

func flattenSystemInterfaceIpv6Vrrp6(v *[]models.SystemInterfaceIpv6Vrrp6, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AcceptMode; tmp != nil {
				v["accept_mode"] = *tmp
			}

			if tmp := cfg.AdvInterval; tmp != nil {
				v["adv_interval"] = *tmp
			}

			if tmp := cfg.Preempt; tmp != nil {
				v["preempt"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			if tmp := cfg.StartTime; tmp != nil {
				v["start_time"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Vrdst6; tmp != nil {
				v["vrdst6"] = *tmp
			}

			if tmp := cfg.Vrgrp; tmp != nil {
				v["vrgrp"] = *tmp
			}

			if tmp := cfg.Vrid; tmp != nil {
				v["vrid"] = *tmp
			}

			if tmp := cfg.Vrip6; tmp != nil {
				v["vrip6"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "vrid")
	}

	return flat
}

func flattenSystemInterfaceMember(v *[]models.SystemInterfaceMember, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.InterfaceName; tmp != nil {
				v["interface_name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "interface_name")
	}

	return flat
}

func flattenSystemInterfaceSecondaryip(v *[]models.SystemInterfaceSecondaryip, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Allowaccess; tmp != nil {
				v["allowaccess"] = *tmp
			}

			if tmp := cfg.Detectprotocol; tmp != nil {
				v["detectprotocol"] = *tmp
			}

			if tmp := cfg.Detectserver; tmp != nil {
				v["detectserver"] = *tmp
			}

			if tmp := cfg.Gwdetect; tmp != nil {
				v["gwdetect"] = *tmp
			}

			if tmp := cfg.HaPriority; tmp != nil {
				v["ha_priority"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			if tmp := cfg.PingServStatus; tmp != nil {
				v["ping_serv_status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemInterfaceSecurityGroups(v *[]models.SystemInterfaceSecurityGroups, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenSystemInterfaceTagging(v *[]models.SystemInterfaceTagging, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Category; tmp != nil {
				v["category"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Tags; tmp != nil {
				v["tags"] = flattenSystemInterfaceTaggingTags(tmp, sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemInterfaceTaggingTags(v *[]models.SystemInterfaceTaggingTags, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenSystemInterfaceVrrp(v *[]models.SystemInterfaceVrrp, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AcceptMode; tmp != nil {
				v["accept_mode"] = *tmp
			}

			if tmp := cfg.AdvInterval; tmp != nil {
				v["adv_interval"] = *tmp
			}

			if tmp := cfg.IgnoreDefaultRoute; tmp != nil {
				v["ignore_default_route"] = *tmp
			}

			if tmp := cfg.Preempt; tmp != nil {
				v["preempt"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			if tmp := cfg.ProxyArp; tmp != nil {
				v["proxy_arp"] = flattenSystemInterfaceVrrpProxyArp(tmp, sort)
			}

			if tmp := cfg.StartTime; tmp != nil {
				v["start_time"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Version; tmp != nil {
				v["version"] = *tmp
			}

			if tmp := cfg.Vrdst; tmp != nil {
				v["vrdst"] = *tmp
			}

			if tmp := cfg.VrdstPriority; tmp != nil {
				v["vrdst_priority"] = *tmp
			}

			if tmp := cfg.Vrgrp; tmp != nil {
				v["vrgrp"] = *tmp
			}

			if tmp := cfg.Vrid; tmp != nil {
				v["vrid"] = *tmp
			}

			if tmp := cfg.Vrip; tmp != nil {
				v["vrip"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "vrid")
	}

	return flat
}

func flattenSystemInterfaceVrrpProxyArp(v *[]models.SystemInterfaceVrrpProxyArp, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectSystemInterface(d *schema.ResourceData, o *models.SystemInterface, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AcName != nil {
		v := *o.AcName

		if err = d.Set("ac_name", v); err != nil {
			return diag.Errorf("error reading ac_name: %v", err)
		}
	}

	if o.Aggregate != nil {
		v := *o.Aggregate

		if err = d.Set("aggregate", v); err != nil {
			return diag.Errorf("error reading aggregate: %v", err)
		}
	}

	if o.Algorithm != nil {
		v := *o.Algorithm

		if err = d.Set("algorithm", v); err != nil {
			return diag.Errorf("error reading algorithm: %v", err)
		}
	}

	if o.Alias != nil {
		v := *o.Alias

		if err = d.Set("alias", v); err != nil {
			return diag.Errorf("error reading alias: %v", err)
		}
	}

	if o.Allowaccess != nil {
		v := *o.Allowaccess

		if err = d.Set("allowaccess", v); err != nil {
			return diag.Errorf("error reading allowaccess: %v", err)
		}
	}

	if o.ApDiscover != nil {
		v := *o.ApDiscover

		if err = d.Set("ap_discover", v); err != nil {
			return diag.Errorf("error reading ap_discover: %v", err)
		}
	}

	if o.Arpforward != nil {
		v := *o.Arpforward

		if err = d.Set("arpforward", v); err != nil {
			return diag.Errorf("error reading arpforward: %v", err)
		}
	}

	if o.AuthType != nil {
		v := *o.AuthType

		if err = d.Set("auth_type", v); err != nil {
			return diag.Errorf("error reading auth_type: %v", err)
		}
	}

	if o.AutoAuthExtensionDevice != nil {
		v := *o.AutoAuthExtensionDevice

		if err = d.Set("auto_auth_extension_device", v); err != nil {
			return diag.Errorf("error reading auto_auth_extension_device: %v", err)
		}
	}

	if o.BandwidthMeasureTime != nil {
		v := *o.BandwidthMeasureTime

		if err = d.Set("bandwidth_measure_time", v); err != nil {
			return diag.Errorf("error reading bandwidth_measure_time: %v", err)
		}
	}

	if o.Bfd != nil {
		v := *o.Bfd

		if err = d.Set("bfd", v); err != nil {
			return diag.Errorf("error reading bfd: %v", err)
		}
	}

	if o.BfdDesiredMinTx != nil {
		v := *o.BfdDesiredMinTx

		if err = d.Set("bfd_desired_min_tx", v); err != nil {
			return diag.Errorf("error reading bfd_desired_min_tx: %v", err)
		}
	}

	if o.BfdDetectMult != nil {
		v := *o.BfdDetectMult

		if err = d.Set("bfd_detect_mult", v); err != nil {
			return diag.Errorf("error reading bfd_detect_mult: %v", err)
		}
	}

	if o.BfdRequiredMinRx != nil {
		v := *o.BfdRequiredMinRx

		if err = d.Set("bfd_required_min_rx", v); err != nil {
			return diag.Errorf("error reading bfd_required_min_rx: %v", err)
		}
	}

	if o.BroadcastForticlientDiscovery != nil {
		v := *o.BroadcastForticlientDiscovery

		if err = d.Set("broadcast_forticlient_discovery", v); err != nil {
			return diag.Errorf("error reading broadcast_forticlient_discovery: %v", err)
		}
	}

	if o.BroadcastForward != nil {
		v := *o.BroadcastForward

		if err = d.Set("broadcast_forward", v); err != nil {
			return diag.Errorf("error reading broadcast_forward: %v", err)
		}
	}

	if o.CliConnStatus != nil {
		v := *o.CliConnStatus

		if err = d.Set("cli_conn_status", v); err != nil {
			return diag.Errorf("error reading cli_conn_status: %v", err)
		}
	}

	if o.ClientOptions != nil {
		if err = d.Set("client_options", flattenSystemInterfaceClientOptions(o.ClientOptions, sort)); err != nil {
			return diag.Errorf("error reading client_options: %v", err)
		}
	}

	if o.Color != nil {
		v := *o.Color

		if err = d.Set("color", v); err != nil {
			return diag.Errorf("error reading color: %v", err)
		}
	}

	if o.DedicatedTo != nil {
		v := *o.DedicatedTo

		if err = d.Set("dedicated_to", v); err != nil {
			return diag.Errorf("error reading dedicated_to: %v", err)
		}
	}

	if o.Defaultgw != nil {
		v := *o.Defaultgw

		if err = d.Set("defaultgw", v); err != nil {
			return diag.Errorf("error reading defaultgw: %v", err)
		}
	}

	if o.Description != nil {
		v := *o.Description

		if err = d.Set("description", v); err != nil {
			return diag.Errorf("error reading description: %v", err)
		}
	}

	if o.DetectedPeerMtu != nil {
		v := *o.DetectedPeerMtu

		if err = d.Set("detected_peer_mtu", v); err != nil {
			return diag.Errorf("error reading detected_peer_mtu: %v", err)
		}
	}

	if o.Detectprotocol != nil {
		v := *o.Detectprotocol

		if err = d.Set("detectprotocol", v); err != nil {
			return diag.Errorf("error reading detectprotocol: %v", err)
		}
	}

	if o.Detectserver != nil {
		v := *o.Detectserver

		if err = d.Set("detectserver", v); err != nil {
			return diag.Errorf("error reading detectserver: %v", err)
		}
	}

	if o.DeviceIdentification != nil {
		v := *o.DeviceIdentification

		if err = d.Set("device_identification", v); err != nil {
			return diag.Errorf("error reading device_identification: %v", err)
		}
	}

	if o.DeviceUserIdentification != nil {
		v := *o.DeviceUserIdentification

		if err = d.Set("device_user_identification", v); err != nil {
			return diag.Errorf("error reading device_user_identification: %v", err)
		}
	}

	if o.Devindex != nil {
		v := *o.Devindex

		if err = d.Set("devindex", v); err != nil {
			return diag.Errorf("error reading devindex: %v", err)
		}
	}

	if o.DhcpClasslessRouteAddition != nil {
		v := *o.DhcpClasslessRouteAddition

		if err = d.Set("dhcp_classless_route_addition", v); err != nil {
			return diag.Errorf("error reading dhcp_classless_route_addition: %v", err)
		}
	}

	if o.DhcpClientIdentifier != nil {
		v := *o.DhcpClientIdentifier

		if err = d.Set("dhcp_client_identifier", v); err != nil {
			return diag.Errorf("error reading dhcp_client_identifier: %v", err)
		}
	}

	if o.DhcpRelayAgentOption != nil {
		v := *o.DhcpRelayAgentOption

		if err = d.Set("dhcp_relay_agent_option", v); err != nil {
			return diag.Errorf("error reading dhcp_relay_agent_option: %v", err)
		}
	}

	if o.DhcpRelayInterface != nil {
		v := *o.DhcpRelayInterface

		if err = d.Set("dhcp_relay_interface", v); err != nil {
			return diag.Errorf("error reading dhcp_relay_interface: %v", err)
		}
	}

	if o.DhcpRelayInterfaceSelectMethod != nil {
		v := *o.DhcpRelayInterfaceSelectMethod

		if err = d.Set("dhcp_relay_interface_select_method", v); err != nil {
			return diag.Errorf("error reading dhcp_relay_interface_select_method: %v", err)
		}
	}

	if o.DhcpRelayIp != nil {
		v := *o.DhcpRelayIp

		if err = d.Set("dhcp_relay_ip", v); err != nil {
			return diag.Errorf("error reading dhcp_relay_ip: %v", err)
		}
	}

	if o.DhcpRelayRequestAllServer != nil {
		v := *o.DhcpRelayRequestAllServer

		if err = d.Set("dhcp_relay_request_all_server", v); err != nil {
			return diag.Errorf("error reading dhcp_relay_request_all_server: %v", err)
		}
	}

	if o.DhcpRelayService != nil {
		v := *o.DhcpRelayService

		if err = d.Set("dhcp_relay_service", v); err != nil {
			return diag.Errorf("error reading dhcp_relay_service: %v", err)
		}
	}

	if o.DhcpRelayType != nil {
		v := *o.DhcpRelayType

		if err = d.Set("dhcp_relay_type", v); err != nil {
			return diag.Errorf("error reading dhcp_relay_type: %v", err)
		}
	}

	if o.DhcpRenewTime != nil {
		v := *o.DhcpRenewTime

		if err = d.Set("dhcp_renew_time", v); err != nil {
			return diag.Errorf("error reading dhcp_renew_time: %v", err)
		}
	}

	if o.DhcpSnoopingServerList != nil {
		if err = d.Set("dhcp_snooping_server_list", flattenSystemInterfaceDhcpSnoopingServerList(o.DhcpSnoopingServerList, sort)); err != nil {
			return diag.Errorf("error reading dhcp_snooping_server_list: %v", err)
		}
	}

	if o.DiscRetryTimeout != nil {
		v := *o.DiscRetryTimeout

		if err = d.Set("disc_retry_timeout", v); err != nil {
			return diag.Errorf("error reading disc_retry_timeout: %v", err)
		}
	}

	if o.DisconnectThreshold != nil {
		v := *o.DisconnectThreshold

		if err = d.Set("disconnect_threshold", v); err != nil {
			return diag.Errorf("error reading disconnect_threshold: %v", err)
		}
	}

	if o.Distance != nil {
		v := *o.Distance

		if err = d.Set("distance", v); err != nil {
			return diag.Errorf("error reading distance: %v", err)
		}
	}

	if o.DnsServerOverride != nil {
		v := *o.DnsServerOverride

		if err = d.Set("dns_server_override", v); err != nil {
			return diag.Errorf("error reading dns_server_override: %v", err)
		}
	}

	if o.DropFragment != nil {
		v := *o.DropFragment

		if err = d.Set("drop_fragment", v); err != nil {
			return diag.Errorf("error reading drop_fragment: %v", err)
		}
	}

	if o.DropOverlappedFragment != nil {
		v := *o.DropOverlappedFragment

		if err = d.Set("drop_overlapped_fragment", v); err != nil {
			return diag.Errorf("error reading drop_overlapped_fragment: %v", err)
		}
	}

	if o.EgressCos != nil {
		v := *o.EgressCos

		if err = d.Set("egress_cos", v); err != nil {
			return diag.Errorf("error reading egress_cos: %v", err)
		}
	}

	if o.EgressQueues != nil {
		if err = d.Set("egress_queues", flattenSystemInterfaceEgressQueues(o.EgressQueues, sort)); err != nil {
			return diag.Errorf("error reading egress_queues: %v", err)
		}
	}

	if o.EgressShapingProfile != nil {
		v := *o.EgressShapingProfile

		if err = d.Set("egress_shaping_profile", v); err != nil {
			return diag.Errorf("error reading egress_shaping_profile: %v", err)
		}
	}

	if o.Eip != nil {
		v := *o.Eip

		if err = d.Set("eip", v); err != nil {
			return diag.Errorf("error reading eip: %v", err)
		}
	}

	if o.EstimatedDownstreamBandwidth != nil {
		v := *o.EstimatedDownstreamBandwidth

		if err = d.Set("estimated_downstream_bandwidth", v); err != nil {
			return diag.Errorf("error reading estimated_downstream_bandwidth: %v", err)
		}
	}

	if o.EstimatedUpstreamBandwidth != nil {
		v := *o.EstimatedUpstreamBandwidth

		if err = d.Set("estimated_upstream_bandwidth", v); err != nil {
			return diag.Errorf("error reading estimated_upstream_bandwidth: %v", err)
		}
	}

	if o.ExplicitFtpProxy != nil {
		v := *o.ExplicitFtpProxy

		if err = d.Set("explicit_ftp_proxy", v); err != nil {
			return diag.Errorf("error reading explicit_ftp_proxy: %v", err)
		}
	}

	if o.ExplicitWebProxy != nil {
		v := *o.ExplicitWebProxy

		if err = d.Set("explicit_web_proxy", v); err != nil {
			return diag.Errorf("error reading explicit_web_proxy: %v", err)
		}
	}

	if o.External != nil {
		v := *o.External

		if err = d.Set("external", v); err != nil {
			return diag.Errorf("error reading external: %v", err)
		}
	}

	if o.FailActionOnExtender != nil {
		v := *o.FailActionOnExtender

		if err = d.Set("fail_action_on_extender", v); err != nil {
			return diag.Errorf("error reading fail_action_on_extender: %v", err)
		}
	}

	if o.FailAlertInterfaces != nil {
		if err = d.Set("fail_alert_interfaces", flattenSystemInterfaceFailAlertInterfaces(o.FailAlertInterfaces, sort)); err != nil {
			return diag.Errorf("error reading fail_alert_interfaces: %v", err)
		}
	}

	if o.FailAlertMethod != nil {
		v := *o.FailAlertMethod

		if err = d.Set("fail_alert_method", v); err != nil {
			return diag.Errorf("error reading fail_alert_method: %v", err)
		}
	}

	if o.FailDetect != nil {
		v := *o.FailDetect

		if err = d.Set("fail_detect", v); err != nil {
			return diag.Errorf("error reading fail_detect: %v", err)
		}
	}

	if o.FailDetectOption != nil {
		v := *o.FailDetectOption

		if err = d.Set("fail_detect_option", v); err != nil {
			return diag.Errorf("error reading fail_detect_option: %v", err)
		}
	}

	if o.Fortilink != nil {
		v := *o.Fortilink

		if err = d.Set("fortilink", v); err != nil {
			return diag.Errorf("error reading fortilink: %v", err)
		}
	}

	if o.FortilinkBackupLink != nil {
		v := *o.FortilinkBackupLink

		if err = d.Set("fortilink_backup_link", v); err != nil {
			return diag.Errorf("error reading fortilink_backup_link: %v", err)
		}
	}

	if o.FortilinkNeighborDetect != nil {
		v := *o.FortilinkNeighborDetect

		if err = d.Set("fortilink_neighbor_detect", v); err != nil {
			return diag.Errorf("error reading fortilink_neighbor_detect: %v", err)
		}
	}

	if o.FortilinkSplitInterface != nil {
		v := *o.FortilinkSplitInterface

		if err = d.Set("fortilink_split_interface", v); err != nil {
			return diag.Errorf("error reading fortilink_split_interface: %v", err)
		}
	}

	if o.FortilinkStacking != nil {
		v := *o.FortilinkStacking

		if err = d.Set("fortilink_stacking", v); err != nil {
			return diag.Errorf("error reading fortilink_stacking: %v", err)
		}
	}

	if o.ForwardDomain != nil {
		v := *o.ForwardDomain

		if err = d.Set("forward_domain", v); err != nil {
			return diag.Errorf("error reading forward_domain: %v", err)
		}
	}

	if o.Gwdetect != nil {
		v := *o.Gwdetect

		if err = d.Set("gwdetect", v); err != nil {
			return diag.Errorf("error reading gwdetect: %v", err)
		}
	}

	if o.HaPriority != nil {
		v := *o.HaPriority

		if err = d.Set("ha_priority", v); err != nil {
			return diag.Errorf("error reading ha_priority: %v", err)
		}
	}

	if o.IcmpAcceptRedirect != nil {
		v := *o.IcmpAcceptRedirect

		if err = d.Set("icmp_accept_redirect", v); err != nil {
			return diag.Errorf("error reading icmp_accept_redirect: %v", err)
		}
	}

	if o.IcmpSendRedirect != nil {
		v := *o.IcmpSendRedirect

		if err = d.Set("icmp_send_redirect", v); err != nil {
			return diag.Errorf("error reading icmp_send_redirect: %v", err)
		}
	}

	if o.IdentAccept != nil {
		v := *o.IdentAccept

		if err = d.Set("ident_accept", v); err != nil {
			return diag.Errorf("error reading ident_accept: %v", err)
		}
	}

	if o.IdleTimeout != nil {
		v := *o.IdleTimeout

		if err = d.Set("idle_timeout", v); err != nil {
			return diag.Errorf("error reading idle_timeout: %v", err)
		}
	}

	if o.Inbandwidth != nil {
		v := *o.Inbandwidth

		if err = d.Set("inbandwidth", v); err != nil {
			return diag.Errorf("error reading inbandwidth: %v", err)
		}
	}

	if o.IngressCos != nil {
		v := *o.IngressCos

		if err = d.Set("ingress_cos", v); err != nil {
			return diag.Errorf("error reading ingress_cos: %v", err)
		}
	}

	if o.IngressShapingProfile != nil {
		v := *o.IngressShapingProfile

		if err = d.Set("ingress_shaping_profile", v); err != nil {
			return diag.Errorf("error reading ingress_shaping_profile: %v", err)
		}
	}

	if o.IngressSpilloverThreshold != nil {
		v := *o.IngressSpilloverThreshold

		if err = d.Set("ingress_spillover_threshold", v); err != nil {
			return diag.Errorf("error reading ingress_spillover_threshold: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.Internal != nil {
		v := *o.Internal

		if err = d.Set("internal", v); err != nil {
			return diag.Errorf("error reading internal: %v", err)
		}
	}

	if o.Ip != nil {
		v := *o.Ip
		if current, ok := d.GetOk("ip"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("ip", v); err != nil {
			return diag.Errorf("error reading ip: %v", err)
		}
	}

	if o.IpManagedByFortiipam != nil {
		v := *o.IpManagedByFortiipam

		if err = d.Set("ip_managed_by_fortiipam", v); err != nil {
			return diag.Errorf("error reading ip_managed_by_fortiipam: %v", err)
		}
	}

	if o.Ipmac != nil {
		v := *o.Ipmac

		if err = d.Set("ipmac", v); err != nil {
			return diag.Errorf("error reading ipmac: %v", err)
		}
	}

	if o.IpsSnifferMode != nil {
		v := *o.IpsSnifferMode

		if err = d.Set("ips_sniffer_mode", v); err != nil {
			return diag.Errorf("error reading ips_sniffer_mode: %v", err)
		}
	}

	if o.Ipunnumbered != nil {
		v := *o.Ipunnumbered

		if err = d.Set("ipunnumbered", v); err != nil {
			return diag.Errorf("error reading ipunnumbered: %v", err)
		}
	}

	if o.Ipv6 != nil {
		if err = d.Set("ipv6", flattenSystemInterfaceIpv6(o.Ipv6, sort)); err != nil {
			return diag.Errorf("error reading ipv6: %v", err)
		}
	}

	if o.L2forward != nil {
		v := *o.L2forward

		if err = d.Set("l2forward", v); err != nil {
			return diag.Errorf("error reading l2forward: %v", err)
		}
	}

	if o.LacpHaSlave != nil {
		v := *o.LacpHaSlave

		if err = d.Set("lacp_ha_slave", v); err != nil {
			return diag.Errorf("error reading lacp_ha_slave: %v", err)
		}
	}

	if o.LacpMode != nil {
		v := *o.LacpMode

		if err = d.Set("lacp_mode", v); err != nil {
			return diag.Errorf("error reading lacp_mode: %v", err)
		}
	}

	if o.LacpSpeed != nil {
		v := *o.LacpSpeed

		if err = d.Set("lacp_speed", v); err != nil {
			return diag.Errorf("error reading lacp_speed: %v", err)
		}
	}

	if o.LcpEchoInterval != nil {
		v := *o.LcpEchoInterval

		if err = d.Set("lcp_echo_interval", v); err != nil {
			return diag.Errorf("error reading lcp_echo_interval: %v", err)
		}
	}

	if o.LcpMaxEchoFails != nil {
		v := *o.LcpMaxEchoFails

		if err = d.Set("lcp_max_echo_fails", v); err != nil {
			return diag.Errorf("error reading lcp_max_echo_fails: %v", err)
		}
	}

	if o.LinkUpDelay != nil {
		v := *o.LinkUpDelay

		if err = d.Set("link_up_delay", v); err != nil {
			return diag.Errorf("error reading link_up_delay: %v", err)
		}
	}

	if o.LldpNetworkPolicy != nil {
		v := *o.LldpNetworkPolicy

		if err = d.Set("lldp_network_policy", v); err != nil {
			return diag.Errorf("error reading lldp_network_policy: %v", err)
		}
	}

	if o.LldpReception != nil {
		v := *o.LldpReception

		if err = d.Set("lldp_reception", v); err != nil {
			return diag.Errorf("error reading lldp_reception: %v", err)
		}
	}

	if o.LldpTransmission != nil {
		v := *o.LldpTransmission

		if err = d.Set("lldp_transmission", v); err != nil {
			return diag.Errorf("error reading lldp_transmission: %v", err)
		}
	}

	if o.Macaddr != nil {
		v := *o.Macaddr

		if err = d.Set("macaddr", v); err != nil {
			return diag.Errorf("error reading macaddr: %v", err)
		}
	}

	if o.ManagedSubnetworkSize != nil {
		v := *o.ManagedSubnetworkSize

		if err = d.Set("managed_subnetwork_size", v); err != nil {
			return diag.Errorf("error reading managed_subnetwork_size: %v", err)
		}
	}

	if o.ManagementIp != nil {
		v := *o.ManagementIp
		if current, ok := d.GetOk("management_ip"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("management_ip", v); err != nil {
			return diag.Errorf("error reading management_ip: %v", err)
		}
	}

	if o.MeasuredDownstreamBandwidth != nil {
		v := *o.MeasuredDownstreamBandwidth

		if err = d.Set("measured_downstream_bandwidth", v); err != nil {
			return diag.Errorf("error reading measured_downstream_bandwidth: %v", err)
		}
	}

	if o.MeasuredUpstreamBandwidth != nil {
		v := *o.MeasuredUpstreamBandwidth

		if err = d.Set("measured_upstream_bandwidth", v); err != nil {
			return diag.Errorf("error reading measured_upstream_bandwidth: %v", err)
		}
	}

	if o.Mediatype != nil {
		v := *o.Mediatype

		if err = d.Set("mediatype", v); err != nil {
			return diag.Errorf("error reading mediatype: %v", err)
		}
	}

	if o.Member != nil {
		if err = d.Set("member", flattenSystemInterfaceMember(o.Member, sort)); err != nil {
			return diag.Errorf("error reading member: %v", err)
		}
	}

	if o.MinLinks != nil {
		v := *o.MinLinks

		if err = d.Set("min_links", v); err != nil {
			return diag.Errorf("error reading min_links: %v", err)
		}
	}

	if o.MinLinksDown != nil {
		v := *o.MinLinksDown

		if err = d.Set("min_links_down", v); err != nil {
			return diag.Errorf("error reading min_links_down: %v", err)
		}
	}

	if o.Mode != nil {
		v := *o.Mode

		if err = d.Set("mode", v); err != nil {
			return diag.Errorf("error reading mode: %v", err)
		}
	}

	if o.MonitorBandwidth != nil {
		v := *o.MonitorBandwidth

		if err = d.Set("monitor_bandwidth", v); err != nil {
			return diag.Errorf("error reading monitor_bandwidth: %v", err)
		}
	}

	if o.Mtu != nil {
		v := *o.Mtu

		if err = d.Set("mtu", v); err != nil {
			return diag.Errorf("error reading mtu: %v", err)
		}
	}

	if o.MtuOverride != nil {
		v := *o.MtuOverride

		if err = d.Set("mtu_override", v); err != nil {
			return diag.Errorf("error reading mtu_override: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Ndiscforward != nil {
		v := *o.Ndiscforward

		if err = d.Set("ndiscforward", v); err != nil {
			return diag.Errorf("error reading ndiscforward: %v", err)
		}
	}

	if o.NetbiosForward != nil {
		v := *o.NetbiosForward

		if err = d.Set("netbios_forward", v); err != nil {
			return diag.Errorf("error reading netbios_forward: %v", err)
		}
	}

	if o.NetflowSampler != nil {
		v := *o.NetflowSampler

		if err = d.Set("netflow_sampler", v); err != nil {
			return diag.Errorf("error reading netflow_sampler: %v", err)
		}
	}

	if o.Outbandwidth != nil {
		v := *o.Outbandwidth

		if err = d.Set("outbandwidth", v); err != nil {
			return diag.Errorf("error reading outbandwidth: %v", err)
		}
	}

	if o.PadtRetryTimeout != nil {
		v := *o.PadtRetryTimeout

		if err = d.Set("padt_retry_timeout", v); err != nil {
			return diag.Errorf("error reading padt_retry_timeout: %v", err)
		}
	}

	if o.Password != nil {
		v := *o.Password

		if err = d.Set("password", v); err != nil {
			return diag.Errorf("error reading password: %v", err)
		}
	}

	if o.PingServStatus != nil {
		v := *o.PingServStatus

		if err = d.Set("ping_serv_status", v); err != nil {
			return diag.Errorf("error reading ping_serv_status: %v", err)
		}
	}

	if o.PollingInterval != nil {
		v := *o.PollingInterval

		if err = d.Set("polling_interval", v); err != nil {
			return diag.Errorf("error reading polling_interval: %v", err)
		}
	}

	if o.PppoeUnnumberedNegotiate != nil {
		v := *o.PppoeUnnumberedNegotiate

		if err = d.Set("pppoe_unnumbered_negotiate", v); err != nil {
			return diag.Errorf("error reading pppoe_unnumbered_negotiate: %v", err)
		}
	}

	if o.PptpAuthType != nil {
		v := *o.PptpAuthType

		if err = d.Set("pptp_auth_type", v); err != nil {
			return diag.Errorf("error reading pptp_auth_type: %v", err)
		}
	}

	if o.PptpClient != nil {
		v := *o.PptpClient

		if err = d.Set("pptp_client", v); err != nil {
			return diag.Errorf("error reading pptp_client: %v", err)
		}
	}

	if o.PptpPassword != nil {
		v := *o.PptpPassword

		if err = d.Set("pptp_password", v); err != nil {
			return diag.Errorf("error reading pptp_password: %v", err)
		}
	}

	if o.PptpServerIp != nil {
		v := *o.PptpServerIp

		if err = d.Set("pptp_server_ip", v); err != nil {
			return diag.Errorf("error reading pptp_server_ip: %v", err)
		}
	}

	if o.PptpTimeout != nil {
		v := *o.PptpTimeout

		if err = d.Set("pptp_timeout", v); err != nil {
			return diag.Errorf("error reading pptp_timeout: %v", err)
		}
	}

	if o.PptpUser != nil {
		v := *o.PptpUser

		if err = d.Set("pptp_user", v); err != nil {
			return diag.Errorf("error reading pptp_user: %v", err)
		}
	}

	if o.PreserveSessionRoute != nil {
		v := *o.PreserveSessionRoute

		if err = d.Set("preserve_session_route", v); err != nil {
			return diag.Errorf("error reading preserve_session_route: %v", err)
		}
	}

	if o.Priority != nil {
		v := *o.Priority

		if err = d.Set("priority", v); err != nil {
			return diag.Errorf("error reading priority: %v", err)
		}
	}

	if o.PriorityOverride != nil {
		v := *o.PriorityOverride

		if err = d.Set("priority_override", v); err != nil {
			return diag.Errorf("error reading priority_override: %v", err)
		}
	}

	if o.ProxyCaptivePortal != nil {
		v := *o.ProxyCaptivePortal

		if err = d.Set("proxy_captive_portal", v); err != nil {
			return diag.Errorf("error reading proxy_captive_portal: %v", err)
		}
	}

	if o.RedundantInterface != nil {
		v := *o.RedundantInterface

		if err = d.Set("redundant_interface", v); err != nil {
			return diag.Errorf("error reading redundant_interface: %v", err)
		}
	}

	if o.RemoteIp != nil {
		v := *o.RemoteIp
		if current, ok := d.GetOk("remote_ip"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("remote_ip", v); err != nil {
			return diag.Errorf("error reading remote_ip: %v", err)
		}
	}

	if o.ReplacemsgOverrideGroup != nil {
		v := *o.ReplacemsgOverrideGroup

		if err = d.Set("replacemsg_override_group", v); err != nil {
			return diag.Errorf("error reading replacemsg_override_group: %v", err)
		}
	}

	if o.RingRx != nil {
		v := *o.RingRx

		if err = d.Set("ring_rx", v); err != nil {
			return diag.Errorf("error reading ring_rx: %v", err)
		}
	}

	if o.RingTx != nil {
		v := *o.RingTx

		if err = d.Set("ring_tx", v); err != nil {
			return diag.Errorf("error reading ring_tx: %v", err)
		}
	}

	if o.Role != nil {
		v := *o.Role

		if err = d.Set("role", v); err != nil {
			return diag.Errorf("error reading role: %v", err)
		}
	}

	if o.SampleDirection != nil {
		v := *o.SampleDirection

		if err = d.Set("sample_direction", v); err != nil {
			return diag.Errorf("error reading sample_direction: %v", err)
		}
	}

	if o.SampleRate != nil {
		v := *o.SampleRate

		if err = d.Set("sample_rate", v); err != nil {
			return diag.Errorf("error reading sample_rate: %v", err)
		}
	}

	if o.SecondaryIP != nil {
		v := *o.SecondaryIP

		if err = d.Set("secondary_ip", v); err != nil {
			return diag.Errorf("error reading secondary_ip: %v", err)
		}
	}

	if o.Secondaryip != nil {
		if err = d.Set("secondaryip", flattenSystemInterfaceSecondaryip(o.Secondaryip, sort)); err != nil {
			return diag.Errorf("error reading secondaryip: %v", err)
		}
	}

	if o.SecurityExemptList != nil {
		v := *o.SecurityExemptList

		if err = d.Set("security_exempt_list", v); err != nil {
			return diag.Errorf("error reading security_exempt_list: %v", err)
		}
	}

	if o.SecurityExternalLogout != nil {
		v := *o.SecurityExternalLogout

		if err = d.Set("security_external_logout", v); err != nil {
			return diag.Errorf("error reading security_external_logout: %v", err)
		}
	}

	if o.SecurityExternalWeb != nil {
		v := *o.SecurityExternalWeb

		if err = d.Set("security_external_web", v); err != nil {
			return diag.Errorf("error reading security_external_web: %v", err)
		}
	}

	if o.SecurityGroups != nil {
		if err = d.Set("security_groups", flattenSystemInterfaceSecurityGroups(o.SecurityGroups, sort)); err != nil {
			return diag.Errorf("error reading security_groups: %v", err)
		}
	}

	if o.SecurityMacAuthBypass != nil {
		v := *o.SecurityMacAuthBypass

		if err = d.Set("security_mac_auth_bypass", v); err != nil {
			return diag.Errorf("error reading security_mac_auth_bypass: %v", err)
		}
	}

	if o.SecurityMode != nil {
		v := *o.SecurityMode

		if err = d.Set("security_mode", v); err != nil {
			return diag.Errorf("error reading security_mode: %v", err)
		}
	}

	if o.SecurityRedirectUrl != nil {
		v := *o.SecurityRedirectUrl

		if err = d.Set("security_redirect_url", v); err != nil {
			return diag.Errorf("error reading security_redirect_url: %v", err)
		}
	}

	if o.ServiceName != nil {
		v := *o.ServiceName

		if err = d.Set("service_name", v); err != nil {
			return diag.Errorf("error reading service_name: %v", err)
		}
	}

	if o.SflowSampler != nil {
		v := *o.SflowSampler

		if err = d.Set("sflow_sampler", v); err != nil {
			return diag.Errorf("error reading sflow_sampler: %v", err)
		}
	}

	if o.SnmpIndex != nil {
		v := *o.SnmpIndex

		if err = d.Set("snmp_index", v); err != nil {
			return diag.Errorf("error reading snmp_index: %v", err)
		}
	}

	if o.Speed != nil {
		v := *o.Speed

		if err = d.Set("speed", v); err != nil {
			return diag.Errorf("error reading speed: %v", err)
		}
	}

	if o.SpilloverThreshold != nil {
		v := *o.SpilloverThreshold

		if err = d.Set("spillover_threshold", v); err != nil {
			return diag.Errorf("error reading spillover_threshold: %v", err)
		}
	}

	if o.SrcCheck != nil {
		v := *o.SrcCheck

		if err = d.Set("src_check", v); err != nil {
			return diag.Errorf("error reading src_check: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Stpforward != nil {
		v := *o.Stpforward

		if err = d.Set("stpforward", v); err != nil {
			return diag.Errorf("error reading stpforward: %v", err)
		}
	}

	if o.StpforwardMode != nil {
		v := *o.StpforwardMode

		if err = d.Set("stpforward_mode", v); err != nil {
			return diag.Errorf("error reading stpforward_mode: %v", err)
		}
	}

	if o.Subst != nil {
		v := *o.Subst

		if err = d.Set("subst", v); err != nil {
			return diag.Errorf("error reading subst: %v", err)
		}
	}

	if o.SubstituteDstMac != nil {
		v := *o.SubstituteDstMac

		if err = d.Set("substitute_dst_mac", v); err != nil {
			return diag.Errorf("error reading substitute_dst_mac: %v", err)
		}
	}

	if o.SwcFirstCreate != nil {
		v := *o.SwcFirstCreate

		if err = d.Set("swc_first_create", v); err != nil {
			return diag.Errorf("error reading swc_first_create: %v", err)
		}
	}

	if o.SwcVlan != nil {
		v := *o.SwcVlan

		if err = d.Set("swc_vlan", v); err != nil {
			return diag.Errorf("error reading swc_vlan: %v", err)
		}
	}

	if o.Switch != nil {
		v := *o.Switch

		if err = d.Set("switch", v); err != nil {
			return diag.Errorf("error reading switch: %v", err)
		}
	}

	if o.SwitchControllerAccessVlan != nil {
		v := *o.SwitchControllerAccessVlan

		if err = d.Set("switch_controller_access_vlan", v); err != nil {
			return diag.Errorf("error reading switch_controller_access_vlan: %v", err)
		}
	}

	if o.SwitchControllerArpInspection != nil {
		v := *o.SwitchControllerArpInspection

		if err = d.Set("switch_controller_arp_inspection", v); err != nil {
			return diag.Errorf("error reading switch_controller_arp_inspection: %v", err)
		}
	}

	if o.SwitchControllerDhcpSnooping != nil {
		v := *o.SwitchControllerDhcpSnooping

		if err = d.Set("switch_controller_dhcp_snooping", v); err != nil {
			return diag.Errorf("error reading switch_controller_dhcp_snooping: %v", err)
		}
	}

	if o.SwitchControllerDhcpSnoopingOption82 != nil {
		v := *o.SwitchControllerDhcpSnoopingOption82

		if err = d.Set("switch_controller_dhcp_snooping_option82", v); err != nil {
			return diag.Errorf("error reading switch_controller_dhcp_snooping_option82: %v", err)
		}
	}

	if o.SwitchControllerDhcpSnoopingVerifyMac != nil {
		v := *o.SwitchControllerDhcpSnoopingVerifyMac

		if err = d.Set("switch_controller_dhcp_snooping_verify_mac", v); err != nil {
			return diag.Errorf("error reading switch_controller_dhcp_snooping_verify_mac: %v", err)
		}
	}

	if o.SwitchControllerDynamic != nil {
		v := *o.SwitchControllerDynamic

		if err = d.Set("switch_controller_dynamic", v); err != nil {
			return diag.Errorf("error reading switch_controller_dynamic: %v", err)
		}
	}

	if o.SwitchControllerFeature != nil {
		v := *o.SwitchControllerFeature

		if err = d.Set("switch_controller_feature", v); err != nil {
			return diag.Errorf("error reading switch_controller_feature: %v", err)
		}
	}

	if o.SwitchControllerIgmpSnooping != nil {
		v := *o.SwitchControllerIgmpSnooping

		if err = d.Set("switch_controller_igmp_snooping", v); err != nil {
			return diag.Errorf("error reading switch_controller_igmp_snooping: %v", err)
		}
	}

	if o.SwitchControllerIgmpSnoopingFastLeave != nil {
		v := *o.SwitchControllerIgmpSnoopingFastLeave

		if err = d.Set("switch_controller_igmp_snooping_fast_leave", v); err != nil {
			return diag.Errorf("error reading switch_controller_igmp_snooping_fast_leave: %v", err)
		}
	}

	if o.SwitchControllerIgmpSnoopingProxy != nil {
		v := *o.SwitchControllerIgmpSnoopingProxy

		if err = d.Set("switch_controller_igmp_snooping_proxy", v); err != nil {
			return diag.Errorf("error reading switch_controller_igmp_snooping_proxy: %v", err)
		}
	}

	if o.SwitchControllerIotScanning != nil {
		v := *o.SwitchControllerIotScanning

		if err = d.Set("switch_controller_iot_scanning", v); err != nil {
			return diag.Errorf("error reading switch_controller_iot_scanning: %v", err)
		}
	}

	if o.SwitchControllerLearningLimit != nil {
		v := *o.SwitchControllerLearningLimit

		if err = d.Set("switch_controller_learning_limit", v); err != nil {
			return diag.Errorf("error reading switch_controller_learning_limit: %v", err)
		}
	}

	if o.SwitchControllerMgmtVlan != nil {
		v := *o.SwitchControllerMgmtVlan

		if err = d.Set("switch_controller_mgmt_vlan", v); err != nil {
			return diag.Errorf("error reading switch_controller_mgmt_vlan: %v", err)
		}
	}

	if o.SwitchControllerNac != nil {
		v := *o.SwitchControllerNac

		if err = d.Set("switch_controller_nac", v); err != nil {
			return diag.Errorf("error reading switch_controller_nac: %v", err)
		}
	}

	if o.SwitchControllerRspanMode != nil {
		v := *o.SwitchControllerRspanMode

		if err = d.Set("switch_controller_rspan_mode", v); err != nil {
			return diag.Errorf("error reading switch_controller_rspan_mode: %v", err)
		}
	}

	if o.SwitchControllerSourceIp != nil {
		v := *o.SwitchControllerSourceIp

		if err = d.Set("switch_controller_source_ip", v); err != nil {
			return diag.Errorf("error reading switch_controller_source_ip: %v", err)
		}
	}

	if o.SwitchControllerTrafficPolicy != nil {
		v := *o.SwitchControllerTrafficPolicy

		if err = d.Set("switch_controller_traffic_policy", v); err != nil {
			return diag.Errorf("error reading switch_controller_traffic_policy: %v", err)
		}
	}

	if o.SystemId != nil {
		v := *o.SystemId

		if err = d.Set("system_id", v); err != nil {
			return diag.Errorf("error reading system_id: %v", err)
		}
	}

	if o.SystemIdType != nil {
		v := *o.SystemIdType

		if err = d.Set("system_id_type", v); err != nil {
			return diag.Errorf("error reading system_id_type: %v", err)
		}
	}

	if o.Tagging != nil {
		if err = d.Set("tagging", flattenSystemInterfaceTagging(o.Tagging, sort)); err != nil {
			return diag.Errorf("error reading tagging: %v", err)
		}
	}

	if o.TcpMss != nil {
		v := *o.TcpMss

		if err = d.Set("tcp_mss", v); err != nil {
			return diag.Errorf("error reading tcp_mss: %v", err)
		}
	}

	if o.TrustIp1 != nil {
		v := *o.TrustIp1
		if current, ok := d.GetOk("trust_ip_1"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("trust_ip_1", v); err != nil {
			return diag.Errorf("error reading trust_ip_1: %v", err)
		}
	}

	if o.TrustIp2 != nil {
		v := *o.TrustIp2
		if current, ok := d.GetOk("trust_ip_2"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("trust_ip_2", v); err != nil {
			return diag.Errorf("error reading trust_ip_2: %v", err)
		}
	}

	if o.TrustIp3 != nil {
		v := *o.TrustIp3
		if current, ok := d.GetOk("trust_ip_3"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("trust_ip_3", v); err != nil {
			return diag.Errorf("error reading trust_ip_3: %v", err)
		}
	}

	if o.TrustIp61 != nil {
		v := *o.TrustIp61

		if err = d.Set("trust_ip6_1", v); err != nil {
			return diag.Errorf("error reading trust_ip6_1: %v", err)
		}
	}

	if o.TrustIp62 != nil {
		v := *o.TrustIp62

		if err = d.Set("trust_ip6_2", v); err != nil {
			return diag.Errorf("error reading trust_ip6_2: %v", err)
		}
	}

	if o.TrustIp63 != nil {
		v := *o.TrustIp63

		if err = d.Set("trust_ip6_3", v); err != nil {
			return diag.Errorf("error reading trust_ip6_3: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	if o.Username != nil {
		v := *o.Username

		if err = d.Set("username", v); err != nil {
			return diag.Errorf("error reading username: %v", err)
		}
	}

	if o.Vdom != nil {
		v := *o.Vdom

		if err = d.Set("vdom", v); err != nil {
			return diag.Errorf("error reading vdom: %v", err)
		}
	}

	if o.Vindex != nil {
		v := *o.Vindex

		if err = d.Set("vindex", v); err != nil {
			return diag.Errorf("error reading vindex: %v", err)
		}
	}

	if o.VlanProtocol != nil {
		v := *o.VlanProtocol

		if err = d.Set("vlan_protocol", v); err != nil {
			return diag.Errorf("error reading vlan_protocol: %v", err)
		}
	}

	if o.Vlanforward != nil {
		v := *o.Vlanforward

		if err = d.Set("vlanforward", v); err != nil {
			return diag.Errorf("error reading vlanforward: %v", err)
		}
	}

	if o.Vlanid != nil {
		v := *o.Vlanid

		if err = d.Set("vlanid", v); err != nil {
			return diag.Errorf("error reading vlanid: %v", err)
		}
	}

	if o.Vrf != nil {
		v := *o.Vrf

		if err = d.Set("vrf", v); err != nil {
			return diag.Errorf("error reading vrf: %v", err)
		}
	}

	if o.Vrrp != nil {
		if err = d.Set("vrrp", flattenSystemInterfaceVrrp(o.Vrrp, sort)); err != nil {
			return diag.Errorf("error reading vrrp: %v", err)
		}
	}

	if o.VrrpVirtualMac != nil {
		v := *o.VrrpVirtualMac

		if err = d.Set("vrrp_virtual_mac", v); err != nil {
			return diag.Errorf("error reading vrrp_virtual_mac: %v", err)
		}
	}

	if o.Wccp != nil {
		v := *o.Wccp

		if err = d.Set("wccp", v); err != nil {
			return diag.Errorf("error reading wccp: %v", err)
		}
	}

	if o.Weight != nil {
		v := *o.Weight

		if err = d.Set("weight", v); err != nil {
			return diag.Errorf("error reading weight: %v", err)
		}
	}

	if o.WinsIp != nil {
		v := *o.WinsIp

		if err = d.Set("wins_ip", v); err != nil {
			return diag.Errorf("error reading wins_ip: %v", err)
		}
	}

	return nil
}

func expandSystemInterfaceClientOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemInterfaceClientOptions, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemInterfaceClientOptions

	for i := range l {
		tmp := models.SystemInterfaceClientOptions{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.code", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Code = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.value", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Value = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemInterfaceDhcpSnoopingServerList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemInterfaceDhcpSnoopingServerList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemInterfaceDhcpSnoopingServerList

	for i := range l {
		tmp := models.SystemInterfaceDhcpSnoopingServerList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.server_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ServerIp = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemInterfaceEgressQueues(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemInterfaceEgressQueues, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemInterfaceEgressQueues

	for i := range l {
		tmp := models.SystemInterfaceEgressQueues{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.cos0", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Cos0 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cos1", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Cos1 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cos2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Cos2 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cos3", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Cos3 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cos4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Cos4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cos5", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Cos5 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cos6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Cos6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cos7", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Cos7 = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemInterfaceFailAlertInterfaces(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemInterfaceFailAlertInterfaces, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemInterfaceFailAlertInterfaces

	for i := range l {
		tmp := models.SystemInterfaceFailAlertInterfaces{}
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

func expandSystemInterfaceIpv6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemInterfaceIpv6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemInterfaceIpv6

	for i := range l {
		tmp := models.SystemInterfaceIpv6{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.autoconf", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Autoconf = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cli_conn6_status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.CliConn6Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dhcp6_client_options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dhcp6ClientOptions = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dhcp6_iapd_list", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemInterfaceIpv6Dhcp6IapdList(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemInterfaceIpv6Dhcp6IapdList
			// 	}
			tmp.Dhcp6IapdList = v2

		}

		pre_append = fmt.Sprintf("%s.%d.dhcp6_information_request", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dhcp6InformationRequest = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dhcp6_prefix_delegation", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dhcp6PrefixDelegation = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dhcp6_prefix_hint", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dhcp6PrefixHint = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dhcp6_prefix_hint_plt", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Dhcp6PrefixHintPlt = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dhcp6_prefix_hint_vlt", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Dhcp6PrefixHintVlt = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dhcp6_relay_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dhcp6RelayIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dhcp6_relay_service", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dhcp6RelayService = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dhcp6_relay_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dhcp6RelayType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.icmp6_send_redirect", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Icmp6SendRedirect = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface_identifier", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InterfaceIdentifier = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip6_address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip6Address = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip6_allowaccess", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip6Allowaccess = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip6_default_life", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Ip6DefaultLife = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip6_delegated_prefix_iaid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Ip6DelegatedPrefixIaid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip6_delegated_prefix_list", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemInterfaceIpv6Ip6DelegatedPrefixList(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemInterfaceIpv6Ip6DelegatedPrefixList
			// 	}
			tmp.Ip6DelegatedPrefixList = v2

		}

		pre_append = fmt.Sprintf("%s.%d.ip6_dns_server_override", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip6DnsServerOverride = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip6_extra_addr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemInterfaceIpv6Ip6ExtraAddr(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemInterfaceIpv6Ip6ExtraAddr
			// 	}
			tmp.Ip6ExtraAddr = v2

		}

		pre_append = fmt.Sprintf("%s.%d.ip6_hop_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Ip6HopLimit = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip6_link_mtu", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Ip6LinkMtu = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip6_manage_flag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip6ManageFlag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip6_max_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Ip6MaxInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip6_min_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Ip6MinInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip6_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip6Mode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip6_other_flag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip6OtherFlag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip6_prefix_list", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemInterfaceIpv6Ip6PrefixList(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemInterfaceIpv6Ip6PrefixList
			// 	}
			tmp.Ip6PrefixList = v2

		}

		pre_append = fmt.Sprintf("%s.%d.ip6_prefix_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip6PrefixMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip6_reachable_time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Ip6ReachableTime = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip6_retrans_time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Ip6RetransTime = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip6_send_adv", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip6SendAdv = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip6_subnet", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip6Subnet = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip6_upstream_interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip6UpstreamInterface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nd_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NdCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nd_cga_modifier", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NdCgaModifier = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nd_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NdMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nd_security_level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.NdSecurityLevel = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nd_timestamp_delta", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.NdTimestampDelta = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nd_timestamp_fuzz", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.NdTimestampFuzz = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ra_send_mtu", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RaSendMtu = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unique_autoconf_addr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UniqueAutoconfAddr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vrip6_link_local", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Vrip6_link_local = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vrrp_virtual_mac6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VrrpVirtualMac6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vrrp6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemInterfaceIpv6Vrrp6(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemInterfaceIpv6Vrrp6
			// 	}
			tmp.Vrrp6 = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemInterfaceIpv6Dhcp6IapdList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemInterfaceIpv6Dhcp6IapdList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemInterfaceIpv6Dhcp6IapdList

	for i := range l {
		tmp := models.SystemInterfaceIpv6Dhcp6IapdList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.iaid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Iaid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix_hint", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PrefixHint = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix_hint_plt", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.PrefixHintPlt = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix_hint_vlt", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.PrefixHintVlt = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemInterfaceIpv6Ip6DelegatedPrefixList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemInterfaceIpv6Ip6DelegatedPrefixList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemInterfaceIpv6Ip6DelegatedPrefixList

	for i := range l {
		tmp := models.SystemInterfaceIpv6Ip6DelegatedPrefixList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.autonomous_flag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AutonomousFlag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.delegated_prefix_iaid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.DelegatedPrefixIaid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.onlink_flag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OnlinkFlag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.PrefixId = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rdnss", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Rdnss = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rdnss_service", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RdnssService = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.subnet", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Subnet = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.upstream_interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UpstreamInterface = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemInterfaceIpv6Ip6ExtraAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemInterfaceIpv6Ip6ExtraAddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemInterfaceIpv6Ip6ExtraAddr

	for i := range l {
		tmp := models.SystemInterfaceIpv6Ip6ExtraAddr{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.prefix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Prefix = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemInterfaceIpv6Ip6PrefixList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemInterfaceIpv6Ip6PrefixList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemInterfaceIpv6Ip6PrefixList

	for i := range l {
		tmp := models.SystemInterfaceIpv6Ip6PrefixList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.autonomous_flag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AutonomousFlag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dnssl", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemInterfaceIpv6Ip6PrefixListDnssl(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemInterfaceIpv6Ip6PrefixListDnssl
			// 	}
			tmp.Dnssl = v2

		}

		pre_append = fmt.Sprintf("%s.%d.onlink_flag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OnlinkFlag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.preferred_life_time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.PreferredLifeTime = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Prefix = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rdnss", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Rdnss = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.valid_life_time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.ValidLifeTime = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemInterfaceIpv6Ip6PrefixListDnssl(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemInterfaceIpv6Ip6PrefixListDnssl, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemInterfaceIpv6Ip6PrefixListDnssl

	for i := range l {
		tmp := models.SystemInterfaceIpv6Ip6PrefixListDnssl{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.domain", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Domain = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemInterfaceIpv6Vrrp6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemInterfaceIpv6Vrrp6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemInterfaceIpv6Vrrp6

	for i := range l {
		tmp := models.SystemInterfaceIpv6Vrrp6{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.accept_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AcceptMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.adv_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.AdvInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.preempt", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Preempt = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Priority = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.start_time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.StartTime = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vrdst6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Vrdst6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vrgrp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Vrgrp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vrid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Vrid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vrip6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Vrip6 = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemInterfaceMember(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemInterfaceMember, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemInterfaceMember

	for i := range l {
		tmp := models.SystemInterfaceMember{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.interface_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InterfaceName = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemInterfaceSecondaryip(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemInterfaceSecondaryip, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemInterfaceSecondaryip

	for i := range l {
		tmp := models.SystemInterfaceSecondaryip{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.allowaccess", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Allowaccess = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.detectprotocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Detectprotocol = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.detectserver", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Detectserver = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.gwdetect", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Gwdetect = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ha_priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.HaPriority = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ping_serv_status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.PingServStatus = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemInterfaceSecurityGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemInterfaceSecurityGroups, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemInterfaceSecurityGroups

	for i := range l {
		tmp := models.SystemInterfaceSecurityGroups{}
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

func expandSystemInterfaceTagging(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemInterfaceTagging, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemInterfaceTagging

	for i := range l {
		tmp := models.SystemInterfaceTagging{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.category", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Category = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tags", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemInterfaceTaggingTags(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemInterfaceTaggingTags
			// 	}
			tmp.Tags = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemInterfaceTaggingTags(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemInterfaceTaggingTags, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemInterfaceTaggingTags

	for i := range l {
		tmp := models.SystemInterfaceTaggingTags{}
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

func expandSystemInterfaceVrrp(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemInterfaceVrrp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemInterfaceVrrp

	for i := range l {
		tmp := models.SystemInterfaceVrrp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.accept_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AcceptMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.adv_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.AdvInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ignore_default_route", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IgnoreDefaultRoute = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.preempt", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Preempt = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Priority = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.proxy_arp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemInterfaceVrrpProxyArp(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemInterfaceVrrpProxyArp
			// 	}
			tmp.ProxyArp = v2

		}

		pre_append = fmt.Sprintf("%s.%d.start_time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.StartTime = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Version = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vrdst", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Vrdst = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vrdst_priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.VrdstPriority = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vrgrp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Vrgrp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vrid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Vrid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vrip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Vrip = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemInterfaceVrrpProxyArp(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemInterfaceVrrpProxyArp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemInterfaceVrrpProxyArp

	for i := range l {
		tmp := models.SystemInterfaceVrrpProxyArp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemInterface(d *schema.ResourceData, sv string) (*models.SystemInterface, diag.Diagnostics) {
	obj := models.SystemInterface{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("ac_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ac_name", sv)
				diags = append(diags, e)
			}
			obj.AcName = &v2
		}
	}
	if v1, ok := d.GetOk("aggregate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("aggregate", sv)
				diags = append(diags, e)
			}
			obj.Aggregate = &v2
		}
	}
	if v1, ok := d.GetOk("algorithm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("algorithm", sv)
				diags = append(diags, e)
			}
			obj.Algorithm = &v2
		}
	}
	if v1, ok := d.GetOk("alias"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("alias", sv)
				diags = append(diags, e)
			}
			obj.Alias = &v2
		}
	}
	if v1, ok := d.GetOk("allowaccess"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("allowaccess", sv)
				diags = append(diags, e)
			}
			obj.Allowaccess = &v2
		}
	}
	if v1, ok := d.GetOk("ap_discover"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ap_discover", sv)
				diags = append(diags, e)
			}
			obj.ApDiscover = &v2
		}
	}
	if v1, ok := d.GetOk("arpforward"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("arpforward", sv)
				diags = append(diags, e)
			}
			obj.Arpforward = &v2
		}
	}
	if v1, ok := d.GetOk("auth_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_type", sv)
				diags = append(diags, e)
			}
			obj.AuthType = &v2
		}
	}
	if v1, ok := d.GetOk("auto_auth_extension_device"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_auth_extension_device", sv)
				diags = append(diags, e)
			}
			obj.AutoAuthExtensionDevice = &v2
		}
	}
	if v1, ok := d.GetOk("bandwidth_measure_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("bandwidth_measure_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.BandwidthMeasureTime = &tmp
		}
	}
	if v1, ok := d.GetOk("bfd"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bfd", sv)
				diags = append(diags, e)
			}
			obj.Bfd = &v2
		}
	}
	if v1, ok := d.GetOk("bfd_desired_min_tx"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bfd_desired_min_tx", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.BfdDesiredMinTx = &tmp
		}
	}
	if v1, ok := d.GetOk("bfd_detect_mult"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bfd_detect_mult", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.BfdDetectMult = &tmp
		}
	}
	if v1, ok := d.GetOk("bfd_required_min_rx"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bfd_required_min_rx", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.BfdRequiredMinRx = &tmp
		}
	}
	if v1, ok := d.GetOk("broadcast_forticlient_discovery"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("broadcast_forticlient_discovery", sv)
				diags = append(diags, e)
			}
			obj.BroadcastForticlientDiscovery = &v2
		}
	}
	if v1, ok := d.GetOk("broadcast_forward"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("broadcast_forward", sv)
				diags = append(diags, e)
			}
			obj.BroadcastForward = &v2
		}
	}
	if v1, ok := d.GetOk("cli_conn_status"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cli_conn_status", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CliConnStatus = &tmp
		}
	}
	if v, ok := d.GetOk("client_options"); ok {
		if !utils.CheckVer(sv, "v6.4.2", "") {
			e := utils.AttributeVersionWarning("client_options", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemInterfaceClientOptions(d, v, "client_options", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ClientOptions = t
		}
	} else if d.HasChange("client_options") {
		old, new := d.GetChange("client_options")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ClientOptions = &[]models.SystemInterfaceClientOptions{}
		}
	}
	if v1, ok := d.GetOk("color"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("color", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Color = &tmp
		}
	}
	if v1, ok := d.GetOk("dedicated_to"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dedicated_to", sv)
				diags = append(diags, e)
			}
			obj.DedicatedTo = &v2
		}
	}
	if v1, ok := d.GetOk("defaultgw"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("defaultgw", sv)
				diags = append(diags, e)
			}
			obj.Defaultgw = &v2
		}
	}
	if v1, ok := d.GetOk("description"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("description", sv)
				diags = append(diags, e)
			}
			obj.Description = &v2
		}
	}
	if v1, ok := d.GetOk("detected_peer_mtu"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("detected_peer_mtu", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DetectedPeerMtu = &tmp
		}
	}
	if v1, ok := d.GetOk("detectprotocol"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("detectprotocol", sv)
				diags = append(diags, e)
			}
			obj.Detectprotocol = &v2
		}
	}
	if v1, ok := d.GetOk("detectserver"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("detectserver", sv)
				diags = append(diags, e)
			}
			obj.Detectserver = &v2
		}
	}
	if v1, ok := d.GetOk("device_identification"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("device_identification", sv)
				diags = append(diags, e)
			}
			obj.DeviceIdentification = &v2
		}
	}
	if v1, ok := d.GetOk("device_user_identification"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("device_user_identification", sv)
				diags = append(diags, e)
			}
			obj.DeviceUserIdentification = &v2
		}
	}
	if v1, ok := d.GetOk("devindex"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("devindex", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Devindex = &tmp
		}
	}
	if v1, ok := d.GetOk("dhcp_classless_route_addition"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("dhcp_classless_route_addition", sv)
				diags = append(diags, e)
			}
			obj.DhcpClasslessRouteAddition = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_client_identifier"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp_client_identifier", sv)
				diags = append(diags, e)
			}
			obj.DhcpClientIdentifier = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_relay_agent_option"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp_relay_agent_option", sv)
				diags = append(diags, e)
			}
			obj.DhcpRelayAgentOption = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_relay_interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp_relay_interface", sv)
				diags = append(diags, e)
			}
			obj.DhcpRelayInterface = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_relay_interface_select_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp_relay_interface_select_method", sv)
				diags = append(diags, e)
			}
			obj.DhcpRelayInterfaceSelectMethod = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_relay_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp_relay_ip", sv)
				diags = append(diags, e)
			}
			obj.DhcpRelayIp = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_relay_request_all_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("dhcp_relay_request_all_server", sv)
				diags = append(diags, e)
			}
			obj.DhcpRelayRequestAllServer = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_relay_service"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp_relay_service", sv)
				diags = append(diags, e)
			}
			obj.DhcpRelayService = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_relay_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp_relay_type", sv)
				diags = append(diags, e)
			}
			obj.DhcpRelayType = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_renew_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp_renew_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DhcpRenewTime = &tmp
		}
	}
	if v, ok := d.GetOk("dhcp_snooping_server_list"); ok {
		if !utils.CheckVer(sv, "v7.0.1", "") {
			e := utils.AttributeVersionWarning("dhcp_snooping_server_list", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemInterfaceDhcpSnoopingServerList(d, v, "dhcp_snooping_server_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DhcpSnoopingServerList = t
		}
	} else if d.HasChange("dhcp_snooping_server_list") {
		old, new := d.GetChange("dhcp_snooping_server_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DhcpSnoopingServerList = &[]models.SystemInterfaceDhcpSnoopingServerList{}
		}
	}
	if v1, ok := d.GetOk("disc_retry_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("disc_retry_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DiscRetryTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("disconnect_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("disconnect_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DisconnectThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("distance"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("distance", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Distance = &tmp
		}
	}
	if v1, ok := d.GetOk("dns_server_override"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_server_override", sv)
				diags = append(diags, e)
			}
			obj.DnsServerOverride = &v2
		}
	}
	if v1, ok := d.GetOk("drop_fragment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("drop_fragment", sv)
				diags = append(diags, e)
			}
			obj.DropFragment = &v2
		}
	}
	if v1, ok := d.GetOk("drop_overlapped_fragment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("drop_overlapped_fragment", sv)
				diags = append(diags, e)
			}
			obj.DropOverlappedFragment = &v2
		}
	}
	if v1, ok := d.GetOk("egress_cos"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("egress_cos", sv)
				diags = append(diags, e)
			}
			obj.EgressCos = &v2
		}
	}
	if v, ok := d.GetOk("egress_queues"); ok {
		if !utils.CheckVer(sv, "", "v6.4.2") {
			e := utils.AttributeVersionWarning("egress_queues", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemInterfaceEgressQueues(d, v, "egress_queues", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.EgressQueues = t
		}
	} else if d.HasChange("egress_queues") {
		old, new := d.GetChange("egress_queues")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.EgressQueues = &[]models.SystemInterfaceEgressQueues{}
		}
	}
	if v1, ok := d.GetOk("egress_shaping_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("egress_shaping_profile", sv)
				diags = append(diags, e)
			}
			obj.EgressShapingProfile = &v2
		}
	}
	if v1, ok := d.GetOk("eip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("eip", sv)
				diags = append(diags, e)
			}
			obj.Eip = &v2
		}
	}
	if v1, ok := d.GetOk("estimated_downstream_bandwidth"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("estimated_downstream_bandwidth", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EstimatedDownstreamBandwidth = &tmp
		}
	}
	if v1, ok := d.GetOk("estimated_upstream_bandwidth"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("estimated_upstream_bandwidth", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EstimatedUpstreamBandwidth = &tmp
		}
	}
	if v1, ok := d.GetOk("explicit_ftp_proxy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("explicit_ftp_proxy", sv)
				diags = append(diags, e)
			}
			obj.ExplicitFtpProxy = &v2
		}
	}
	if v1, ok := d.GetOk("explicit_web_proxy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("explicit_web_proxy", sv)
				diags = append(diags, e)
			}
			obj.ExplicitWebProxy = &v2
		}
	}
	if v1, ok := d.GetOk("external"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("external", sv)
				diags = append(diags, e)
			}
			obj.External = &v2
		}
	}
	if v1, ok := d.GetOk("fail_action_on_extender"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fail_action_on_extender", sv)
				diags = append(diags, e)
			}
			obj.FailActionOnExtender = &v2
		}
	}
	if v, ok := d.GetOk("fail_alert_interfaces"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("fail_alert_interfaces", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemInterfaceFailAlertInterfaces(d, v, "fail_alert_interfaces", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.FailAlertInterfaces = t
		}
	} else if d.HasChange("fail_alert_interfaces") {
		old, new := d.GetChange("fail_alert_interfaces")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.FailAlertInterfaces = &[]models.SystemInterfaceFailAlertInterfaces{}
		}
	}
	if v1, ok := d.GetOk("fail_alert_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fail_alert_method", sv)
				diags = append(diags, e)
			}
			obj.FailAlertMethod = &v2
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
	if v1, ok := d.GetOk("fail_detect_option"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fail_detect_option", sv)
				diags = append(diags, e)
			}
			obj.FailDetectOption = &v2
		}
	}
	if v1, ok := d.GetOk("fortilink"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fortilink", sv)
				diags = append(diags, e)
			}
			obj.Fortilink = &v2
		}
	}
	if v1, ok := d.GetOk("fortilink_backup_link"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fortilink_backup_link", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FortilinkBackupLink = &tmp
		}
	}
	if v1, ok := d.GetOk("fortilink_neighbor_detect"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fortilink_neighbor_detect", sv)
				diags = append(diags, e)
			}
			obj.FortilinkNeighborDetect = &v2
		}
	}
	if v1, ok := d.GetOk("fortilink_split_interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fortilink_split_interface", sv)
				diags = append(diags, e)
			}
			obj.FortilinkSplitInterface = &v2
		}
	}
	if v1, ok := d.GetOk("fortilink_stacking"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("fortilink_stacking", sv)
				diags = append(diags, e)
			}
			obj.FortilinkStacking = &v2
		}
	}
	if v1, ok := d.GetOk("forward_domain"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("forward_domain", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ForwardDomain = &tmp
		}
	}
	if v1, ok := d.GetOk("gwdetect"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gwdetect", sv)
				diags = append(diags, e)
			}
			obj.Gwdetect = &v2
		}
	}
	if v1, ok := d.GetOk("ha_priority"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ha_priority", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HaPriority = &tmp
		}
	}
	if v1, ok := d.GetOk("icmp_accept_redirect"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("icmp_accept_redirect", sv)
				diags = append(diags, e)
			}
			obj.IcmpAcceptRedirect = &v2
		}
	}
	if v1, ok := d.GetOk("icmp_send_redirect"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("icmp_send_redirect", sv)
				diags = append(diags, e)
			}
			obj.IcmpSendRedirect = &v2
		}
	}
	if v1, ok := d.GetOk("ident_accept"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ident_accept", sv)
				diags = append(diags, e)
			}
			obj.IdentAccept = &v2
		}
	}
	if v1, ok := d.GetOk("idle_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("idle_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IdleTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("inbandwidth"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("inbandwidth", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Inbandwidth = &tmp
		}
	}
	if v1, ok := d.GetOk("ingress_cos"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("ingress_cos", sv)
				diags = append(diags, e)
			}
			obj.IngressCos = &v2
		}
	}
	if v1, ok := d.GetOk("ingress_shaping_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ingress_shaping_profile", sv)
				diags = append(diags, e)
			}
			obj.IngressShapingProfile = &v2
		}
	}
	if v1, ok := d.GetOk("ingress_spillover_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ingress_spillover_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IngressSpilloverThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("internal"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("internal", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Internal = &tmp
		}
	}
	if v1, ok := d.GetOk("ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip", sv)
				diags = append(diags, e)
			}
			obj.Ip = &v2
		}
	}
	if v1, ok := d.GetOk("ip_managed_by_fortiipam"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("ip_managed_by_fortiipam", sv)
				diags = append(diags, e)
			}
			obj.IpManagedByFortiipam = &v2
		}
	}
	if v1, ok := d.GetOk("ipmac"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipmac", sv)
				diags = append(diags, e)
			}
			obj.Ipmac = &v2
		}
	}
	if v1, ok := d.GetOk("ips_sniffer_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ips_sniffer_mode", sv)
				diags = append(diags, e)
			}
			obj.IpsSnifferMode = &v2
		}
	}
	if v1, ok := d.GetOk("ipunnumbered"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipunnumbered", sv)
				diags = append(diags, e)
			}
			obj.Ipunnumbered = &v2
		}
	}
	if v, ok := d.GetOk("ipv6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ipv6", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemInterfaceIpv6(d, v, "ipv6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Ipv6 = t
		}
	} else if d.HasChange("ipv6") {
		old, new := d.GetChange("ipv6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Ipv6 = &[]models.SystemInterfaceIpv6{}
		}
	}
	if v1, ok := d.GetOk("l2forward"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("l2forward", sv)
				diags = append(diags, e)
			}
			obj.L2forward = &v2
		}
	}
	if v1, ok := d.GetOk("lacp_ha_slave"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lacp_ha_slave", sv)
				diags = append(diags, e)
			}
			obj.LacpHaSlave = &v2
		}
	}
	if v1, ok := d.GetOk("lacp_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lacp_mode", sv)
				diags = append(diags, e)
			}
			obj.LacpMode = &v2
		}
	}
	if v1, ok := d.GetOk("lacp_speed"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lacp_speed", sv)
				diags = append(diags, e)
			}
			obj.LacpSpeed = &v2
		}
	}
	if v1, ok := d.GetOk("lcp_echo_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lcp_echo_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LcpEchoInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("lcp_max_echo_fails"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lcp_max_echo_fails", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LcpMaxEchoFails = &tmp
		}
	}
	if v1, ok := d.GetOk("link_up_delay"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("link_up_delay", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LinkUpDelay = &tmp
		}
	}
	if v1, ok := d.GetOk("lldp_network_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lldp_network_policy", sv)
				diags = append(diags, e)
			}
			obj.LldpNetworkPolicy = &v2
		}
	}
	if v1, ok := d.GetOk("lldp_reception"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lldp_reception", sv)
				diags = append(diags, e)
			}
			obj.LldpReception = &v2
		}
	}
	if v1, ok := d.GetOk("lldp_transmission"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lldp_transmission", sv)
				diags = append(diags, e)
			}
			obj.LldpTransmission = &v2
		}
	}
	if v1, ok := d.GetOk("macaddr"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("macaddr", sv)
				diags = append(diags, e)
			}
			obj.Macaddr = &v2
		}
	}
	if v1, ok := d.GetOk("managed_subnetwork_size"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("managed_subnetwork_size", sv)
				diags = append(diags, e)
			}
			obj.ManagedSubnetworkSize = &v2
		}
	}
	if v1, ok := d.GetOk("management_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("management_ip", sv)
				diags = append(diags, e)
			}
			obj.ManagementIp = &v2
		}
	}
	if v1, ok := d.GetOk("measured_downstream_bandwidth"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("measured_downstream_bandwidth", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MeasuredDownstreamBandwidth = &tmp
		}
	}
	if v1, ok := d.GetOk("measured_upstream_bandwidth"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("measured_upstream_bandwidth", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MeasuredUpstreamBandwidth = &tmp
		}
	}
	if v1, ok := d.GetOk("mediatype"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("mediatype", sv)
				diags = append(diags, e)
			}
			obj.Mediatype = &v2
		}
	}
	if v, ok := d.GetOk("member"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("member", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemInterfaceMember(d, v, "member", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Member = t
		}
	} else if d.HasChange("member") {
		old, new := d.GetChange("member")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Member = &[]models.SystemInterfaceMember{}
		}
	}
	if v1, ok := d.GetOk("min_links"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("min_links", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MinLinks = &tmp
		}
	}
	if v1, ok := d.GetOk("min_links_down"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("min_links_down", sv)
				diags = append(diags, e)
			}
			obj.MinLinksDown = &v2
		}
	}
	if v1, ok := d.GetOk("mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mode", sv)
				diags = append(diags, e)
			}
			obj.Mode = &v2
		}
	}
	if v1, ok := d.GetOk("monitor_bandwidth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("monitor_bandwidth", sv)
				diags = append(diags, e)
			}
			obj.MonitorBandwidth = &v2
		}
	}
	if v1, ok := d.GetOk("mtu"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mtu", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Mtu = &tmp
		}
	}
	if v1, ok := d.GetOk("mtu_override"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mtu_override", sv)
				diags = append(diags, e)
			}
			obj.MtuOverride = &v2
		}
	}
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v1, ok := d.GetOk("ndiscforward"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ndiscforward", sv)
				diags = append(diags, e)
			}
			obj.Ndiscforward = &v2
		}
	}
	if v1, ok := d.GetOk("netbios_forward"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("netbios_forward", sv)
				diags = append(diags, e)
			}
			obj.NetbiosForward = &v2
		}
	}
	if v1, ok := d.GetOk("netflow_sampler"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("netflow_sampler", sv)
				diags = append(diags, e)
			}
			obj.NetflowSampler = &v2
		}
	}
	if v1, ok := d.GetOk("outbandwidth"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("outbandwidth", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Outbandwidth = &tmp
		}
	}
	if v1, ok := d.GetOk("padt_retry_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("padt_retry_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PadtRetryTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("password", sv)
				diags = append(diags, e)
			}
			obj.Password = &v2
		}
	}
	if v1, ok := d.GetOk("ping_serv_status"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ping_serv_status", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PingServStatus = &tmp
		}
	}
	if v1, ok := d.GetOk("polling_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("polling_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PollingInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("pppoe_unnumbered_negotiate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pppoe_unnumbered_negotiate", sv)
				diags = append(diags, e)
			}
			obj.PppoeUnnumberedNegotiate = &v2
		}
	}
	if v1, ok := d.GetOk("pptp_auth_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pptp_auth_type", sv)
				diags = append(diags, e)
			}
			obj.PptpAuthType = &v2
		}
	}
	if v1, ok := d.GetOk("pptp_client"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pptp_client", sv)
				diags = append(diags, e)
			}
			obj.PptpClient = &v2
		}
	}
	if v1, ok := d.GetOk("pptp_password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pptp_password", sv)
				diags = append(diags, e)
			}
			obj.PptpPassword = &v2
		}
	}
	if v1, ok := d.GetOk("pptp_server_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pptp_server_ip", sv)
				diags = append(diags, e)
			}
			obj.PptpServerIp = &v2
		}
	}
	if v1, ok := d.GetOk("pptp_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pptp_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PptpTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("pptp_user"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pptp_user", sv)
				diags = append(diags, e)
			}
			obj.PptpUser = &v2
		}
	}
	if v1, ok := d.GetOk("preserve_session_route"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("preserve_session_route", sv)
				diags = append(diags, e)
			}
			obj.PreserveSessionRoute = &v2
		}
	}
	if v1, ok := d.GetOk("priority"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("priority", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Priority = &tmp
		}
	}
	if v1, ok := d.GetOk("priority_override"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("priority_override", sv)
				diags = append(diags, e)
			}
			obj.PriorityOverride = &v2
		}
	}
	if v1, ok := d.GetOk("proxy_captive_portal"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("proxy_captive_portal", sv)
				diags = append(diags, e)
			}
			obj.ProxyCaptivePortal = &v2
		}
	}
	if v1, ok := d.GetOk("redundant_interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("redundant_interface", sv)
				diags = append(diags, e)
			}
			obj.RedundantInterface = &v2
		}
	}
	if v1, ok := d.GetOk("remote_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("remote_ip", sv)
				diags = append(diags, e)
			}
			obj.RemoteIp = &v2
		}
	}
	if v1, ok := d.GetOk("replacemsg_override_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("replacemsg_override_group", sv)
				diags = append(diags, e)
			}
			obj.ReplacemsgOverrideGroup = &v2
		}
	}
	if v1, ok := d.GetOk("ring_rx"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("ring_rx", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RingRx = &tmp
		}
	}
	if v1, ok := d.GetOk("ring_tx"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("ring_tx", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RingTx = &tmp
		}
	}
	if v1, ok := d.GetOk("role"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("role", sv)
				diags = append(diags, e)
			}
			obj.Role = &v2
		}
	}
	if v1, ok := d.GetOk("sample_direction"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sample_direction", sv)
				diags = append(diags, e)
			}
			obj.SampleDirection = &v2
		}
	}
	if v1, ok := d.GetOk("sample_rate"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sample_rate", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SampleRate = &tmp
		}
	}
	if v1, ok := d.GetOk("secondary_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("secondary_ip", sv)
				diags = append(diags, e)
			}
			obj.SecondaryIP = &v2
		}
	}
	if v, ok := d.GetOk("secondaryip"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("secondaryip", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemInterfaceSecondaryip(d, v, "secondaryip", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Secondaryip = t
		}
	} else if d.HasChange("secondaryip") {
		old, new := d.GetChange("secondaryip")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Secondaryip = &[]models.SystemInterfaceSecondaryip{}
		}
	}
	if v1, ok := d.GetOk("security_exempt_list"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("security_exempt_list", sv)
				diags = append(diags, e)
			}
			obj.SecurityExemptList = &v2
		}
	}
	if v1, ok := d.GetOk("security_external_logout"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("security_external_logout", sv)
				diags = append(diags, e)
			}
			obj.SecurityExternalLogout = &v2
		}
	}
	if v1, ok := d.GetOk("security_external_web"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("security_external_web", sv)
				diags = append(diags, e)
			}
			obj.SecurityExternalWeb = &v2
		}
	}
	if v, ok := d.GetOk("security_groups"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("security_groups", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemInterfaceSecurityGroups(d, v, "security_groups", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SecurityGroups = t
		}
	} else if d.HasChange("security_groups") {
		old, new := d.GetChange("security_groups")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SecurityGroups = &[]models.SystemInterfaceSecurityGroups{}
		}
	}
	if v1, ok := d.GetOk("security_mac_auth_bypass"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("security_mac_auth_bypass", sv)
				diags = append(diags, e)
			}
			obj.SecurityMacAuthBypass = &v2
		}
	}
	if v1, ok := d.GetOk("security_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("security_mode", sv)
				diags = append(diags, e)
			}
			obj.SecurityMode = &v2
		}
	}
	if v1, ok := d.GetOk("security_redirect_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("security_redirect_url", sv)
				diags = append(diags, e)
			}
			obj.SecurityRedirectUrl = &v2
		}
	}
	if v1, ok := d.GetOk("service_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("service_name", sv)
				diags = append(diags, e)
			}
			obj.ServiceName = &v2
		}
	}
	if v1, ok := d.GetOk("sflow_sampler"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sflow_sampler", sv)
				diags = append(diags, e)
			}
			obj.SflowSampler = &v2
		}
	}
	if v1, ok := d.GetOk("snmp_index"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("snmp_index", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SnmpIndex = &tmp
		}
	}
	if v1, ok := d.GetOk("speed"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("speed", sv)
				diags = append(diags, e)
			}
			obj.Speed = &v2
		}
	}
	if v1, ok := d.GetOk("spillover_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("spillover_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SpilloverThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("src_check"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("src_check", sv)
				diags = append(diags, e)
			}
			obj.SrcCheck = &v2
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
	if v1, ok := d.GetOk("stpforward"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("stpforward", sv)
				diags = append(diags, e)
			}
			obj.Stpforward = &v2
		}
	}
	if v1, ok := d.GetOk("stpforward_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("stpforward_mode", sv)
				diags = append(diags, e)
			}
			obj.StpforwardMode = &v2
		}
	}
	if v1, ok := d.GetOk("subst"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("subst", sv)
				diags = append(diags, e)
			}
			obj.Subst = &v2
		}
	}
	if v1, ok := d.GetOk("substitute_dst_mac"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("substitute_dst_mac", sv)
				diags = append(diags, e)
			}
			obj.SubstituteDstMac = &v2
		}
	}
	if v1, ok := d.GetOk("swc_first_create"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("swc_first_create", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SwcFirstCreate = &tmp
		}
	}
	if v1, ok := d.GetOk("swc_vlan"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("swc_vlan", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SwcVlan = &tmp
		}
	}
	if v1, ok := d.GetOk("switch"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("switch", sv)
				diags = append(diags, e)
			}
			obj.Switch = &v2
		}
	}
	if v1, ok := d.GetOk("switch_controller_access_vlan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("switch_controller_access_vlan", sv)
				diags = append(diags, e)
			}
			obj.SwitchControllerAccessVlan = &v2
		}
	}
	if v1, ok := d.GetOk("switch_controller_arp_inspection"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("switch_controller_arp_inspection", sv)
				diags = append(diags, e)
			}
			obj.SwitchControllerArpInspection = &v2
		}
	}
	if v1, ok := d.GetOk("switch_controller_dhcp_snooping"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("switch_controller_dhcp_snooping", sv)
				diags = append(diags, e)
			}
			obj.SwitchControllerDhcpSnooping = &v2
		}
	}
	if v1, ok := d.GetOk("switch_controller_dhcp_snooping_option82"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("switch_controller_dhcp_snooping_option82", sv)
				diags = append(diags, e)
			}
			obj.SwitchControllerDhcpSnoopingOption82 = &v2
		}
	}
	if v1, ok := d.GetOk("switch_controller_dhcp_snooping_verify_mac"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("switch_controller_dhcp_snooping_verify_mac", sv)
				diags = append(diags, e)
			}
			obj.SwitchControllerDhcpSnoopingVerifyMac = &v2
		}
	}
	if v1, ok := d.GetOk("switch_controller_dynamic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("switch_controller_dynamic", sv)
				diags = append(diags, e)
			}
			obj.SwitchControllerDynamic = &v2
		}
	}
	if v1, ok := d.GetOk("switch_controller_feature"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("switch_controller_feature", sv)
				diags = append(diags, e)
			}
			obj.SwitchControllerFeature = &v2
		}
	}
	if v1, ok := d.GetOk("switch_controller_igmp_snooping"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("switch_controller_igmp_snooping", sv)
				diags = append(diags, e)
			}
			obj.SwitchControllerIgmpSnooping = &v2
		}
	}
	if v1, ok := d.GetOk("switch_controller_igmp_snooping_fast_leave"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("switch_controller_igmp_snooping_fast_leave", sv)
				diags = append(diags, e)
			}
			obj.SwitchControllerIgmpSnoopingFastLeave = &v2
		}
	}
	if v1, ok := d.GetOk("switch_controller_igmp_snooping_proxy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("switch_controller_igmp_snooping_proxy", sv)
				diags = append(diags, e)
			}
			obj.SwitchControllerIgmpSnoopingProxy = &v2
		}
	}
	if v1, ok := d.GetOk("switch_controller_iot_scanning"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("switch_controller_iot_scanning", sv)
				diags = append(diags, e)
			}
			obj.SwitchControllerIotScanning = &v2
		}
	}
	if v1, ok := d.GetOk("switch_controller_learning_limit"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("switch_controller_learning_limit", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SwitchControllerLearningLimit = &tmp
		}
	}
	if v1, ok := d.GetOk("switch_controller_mgmt_vlan"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("switch_controller_mgmt_vlan", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SwitchControllerMgmtVlan = &tmp
		}
	}
	if v1, ok := d.GetOk("switch_controller_nac"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("switch_controller_nac", sv)
				diags = append(diags, e)
			}
			obj.SwitchControllerNac = &v2
		}
	}
	if v1, ok := d.GetOk("switch_controller_rspan_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("switch_controller_rspan_mode", sv)
				diags = append(diags, e)
			}
			obj.SwitchControllerRspanMode = &v2
		}
	}
	if v1, ok := d.GetOk("switch_controller_source_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("switch_controller_source_ip", sv)
				diags = append(diags, e)
			}
			obj.SwitchControllerSourceIp = &v2
		}
	}
	if v1, ok := d.GetOk("switch_controller_traffic_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("switch_controller_traffic_policy", sv)
				diags = append(diags, e)
			}
			obj.SwitchControllerTrafficPolicy = &v2
		}
	}
	if v1, ok := d.GetOk("system_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.7", "v7.0.0") {
				e := utils.AttributeVersionWarning("system_id", sv)
				diags = append(diags, e)
			}
			obj.SystemId = &v2
		}
	}
	if v1, ok := d.GetOk("system_id_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.7", "v7.0.0") {
				e := utils.AttributeVersionWarning("system_id_type", sv)
				diags = append(diags, e)
			}
			obj.SystemIdType = &v2
		}
	}
	if v, ok := d.GetOk("tagging"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("tagging", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemInterfaceTagging(d, v, "tagging", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Tagging = t
		}
	} else if d.HasChange("tagging") {
		old, new := d.GetChange("tagging")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Tagging = &[]models.SystemInterfaceTagging{}
		}
	}
	if v1, ok := d.GetOk("tcp_mss"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tcp_mss", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TcpMss = &tmp
		}
	}
	if v1, ok := d.GetOk("trust_ip_1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trust_ip_1", sv)
				diags = append(diags, e)
			}
			obj.TrustIp1 = &v2
		}
	}
	if v1, ok := d.GetOk("trust_ip_2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trust_ip_2", sv)
				diags = append(diags, e)
			}
			obj.TrustIp2 = &v2
		}
	}
	if v1, ok := d.GetOk("trust_ip_3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trust_ip_3", sv)
				diags = append(diags, e)
			}
			obj.TrustIp3 = &v2
		}
	}
	if v1, ok := d.GetOk("trust_ip6_1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trust_ip6_1", sv)
				diags = append(diags, e)
			}
			obj.TrustIp61 = &v2
		}
	}
	if v1, ok := d.GetOk("trust_ip6_2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trust_ip6_2", sv)
				diags = append(diags, e)
			}
			obj.TrustIp62 = &v2
		}
	}
	if v1, ok := d.GetOk("trust_ip6_3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trust_ip6_3", sv)
				diags = append(diags, e)
			}
			obj.TrustIp63 = &v2
		}
	}
	if v1, ok := d.GetOk("type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("type", sv)
				diags = append(diags, e)
			}
			obj.Type = &v2
		}
	}
	if v1, ok := d.GetOk("username"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("username", sv)
				diags = append(diags, e)
			}
			obj.Username = &v2
		}
	}
	if v1, ok := d.GetOk("vdom"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vdom", sv)
				diags = append(diags, e)
			}
			obj.Vdom = &v2
		}
	}
	if v1, ok := d.GetOk("vindex"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vindex", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Vindex = &tmp
		}
	}
	if v1, ok := d.GetOk("vlan_protocol"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("vlan_protocol", sv)
				diags = append(diags, e)
			}
			obj.VlanProtocol = &v2
		}
	}
	if v1, ok := d.GetOk("vlanforward"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vlanforward", sv)
				diags = append(diags, e)
			}
			obj.Vlanforward = &v2
		}
	}
	if v1, ok := d.GetOk("vlanid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vlanid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Vlanid = &tmp
		}
	}
	if v1, ok := d.GetOk("vrf"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vrf", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Vrf = &tmp
		}
	}
	if v, ok := d.GetOk("vrrp"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("vrrp", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemInterfaceVrrp(d, v, "vrrp", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Vrrp = t
		}
	} else if d.HasChange("vrrp") {
		old, new := d.GetChange("vrrp")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Vrrp = &[]models.SystemInterfaceVrrp{}
		}
	}
	if v1, ok := d.GetOk("vrrp_virtual_mac"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vrrp_virtual_mac", sv)
				diags = append(diags, e)
			}
			obj.VrrpVirtualMac = &v2
		}
	}
	if v1, ok := d.GetOk("wccp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wccp", sv)
				diags = append(diags, e)
			}
			obj.Wccp = &v2
		}
	}
	if v1, ok := d.GetOk("weight"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("weight", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Weight = &tmp
		}
	}
	if v1, ok := d.GetOk("wins_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wins_ip", sv)
				diags = append(diags, e)
			}
			obj.WinsIp = &v2
		}
	}
	return &obj, diags
}
