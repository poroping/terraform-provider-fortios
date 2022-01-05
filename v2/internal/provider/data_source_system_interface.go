// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure interfaces.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemInterface() *schema.Resource {
	return &schema.Resource{
		Description: "Configure interfaces.",

		ReadContext: dataSourceSystemInterfaceRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"ac_name": {
				Type:        schema.TypeString,
				Description: "PPPoE server name.",
				Computed:    true,
			},
			"aggregate": {
				Type:        schema.TypeString,
				Description: "Aggregate interface.",
				Computed:    true,
			},
			"algorithm": {
				Type:        schema.TypeString,
				Description: "Frame distribution algorithm.",
				Computed:    true,
			},
			"alias": {
				Type:        schema.TypeString,
				Description: "Alias will be displayed with the interface name to make it easier to distinguish.",
				Computed:    true,
			},
			"allowaccess": {
				Type:        schema.TypeString,
				Description: "Permitted types of management access to this interface.",
				Computed:    true,
			},
			"ap_discover": {
				Type:        schema.TypeString,
				Description: "Enable/disable automatic registration of unknown FortiAP devices.",
				Computed:    true,
			},
			"arpforward": {
				Type:        schema.TypeString,
				Description: "Enable/disable ARP forwarding.",
				Computed:    true,
			},
			"auth_type": {
				Type:        schema.TypeString,
				Description: "PPP authentication type to use.",
				Computed:    true,
			},
			"auto_auth_extension_device": {
				Type:        schema.TypeString,
				Description: "Enable/disable automatic authorization of dedicated Fortinet extension device on this interface.",
				Computed:    true,
			},
			"bandwidth_measure_time": {
				Type:        schema.TypeInt,
				Description: "Bandwidth measure time ",
				Computed:    true,
			},
			"bfd": {
				Type:        schema.TypeString,
				Description: "Bidirectional Forwarding Detection (BFD) settings.",
				Computed:    true,
			},
			"bfd_desired_min_tx": {
				Type:        schema.TypeInt,
				Description: "BFD desired minimal transmit interval.",
				Computed:    true,
			},
			"bfd_detect_mult": {
				Type:        schema.TypeInt,
				Description: "BFD detection multiplier.",
				Computed:    true,
			},
			"bfd_required_min_rx": {
				Type:        schema.TypeInt,
				Description: "BFD required minimal receive interval.",
				Computed:    true,
			},
			"broadcast_forticlient_discovery": {
				Type:        schema.TypeString,
				Description: "Enable/disable broadcasting FortiClient discovery messages.",
				Computed:    true,
			},
			"broadcast_forward": {
				Type:        schema.TypeString,
				Description: "Enable/disable broadcast forwarding.",
				Computed:    true,
			},
			"cli_conn_status": {
				Type:        schema.TypeInt,
				Description: "CLI connection status.",
				Computed:    true,
			},
			"client_options": {
				Type:        schema.TypeList,
				Description: "DHCP client options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type:        schema.TypeInt,
							Description: "DHCP client option code.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"ip": {
							Type:        schema.TypeString,
							Description: "DHCP option IPs.",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "DHCP client option type.",
							Computed:    true,
						},
						"value": {
							Type:        schema.TypeString,
							Description: "DHCP client option value.",
							Computed:    true,
						},
					},
				},
			},
			"color": {
				Type:        schema.TypeInt,
				Description: "Color of icon on the GUI.",
				Computed:    true,
			},
			"dedicated_to": {
				Type:        schema.TypeString,
				Description: "Configure interface for single purpose.",
				Computed:    true,
			},
			"defaultgw": {
				Type:        schema.TypeString,
				Description: "Enable to get the gateway IP from the DHCP or PPPoE server.",
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description.",
				Computed:    true,
			},
			"detected_peer_mtu": {
				Type:        schema.TypeInt,
				Description: "MTU of detected peer (0 - 4294967295).",
				Computed:    true,
			},
			"detectprotocol": {
				Type:        schema.TypeString,
				Description: "Protocols used to detect the server.",
				Computed:    true,
			},
			"detectserver": {
				Type:        schema.TypeString,
				Description: "Gateway's ping server for this IP.",
				Computed:    true,
			},
			"device_identification": {
				Type:        schema.TypeString,
				Description: "Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.",
				Computed:    true,
			},
			"device_user_identification": {
				Type:        schema.TypeString,
				Description: "Enable/disable passive gathering of user identity information about users on this interface.",
				Computed:    true,
			},
			"devindex": {
				Type:        schema.TypeInt,
				Description: "Device Index.",
				Computed:    true,
			},
			"dhcp_classless_route_addition": {
				Type:        schema.TypeString,
				Description: "Enable/disable addition of classless static routes retrieved from DHCP server.",
				Computed:    true,
			},
			"dhcp_client_identifier": {
				Type:        schema.TypeString,
				Description: "DHCP client identifier.",
				Computed:    true,
			},
			"dhcp_relay_agent_option": {
				Type:        schema.TypeString,
				Description: "Enable/disable DHCP relay agent option.",
				Computed:    true,
			},
			"dhcp_relay_interface": {
				Type:        schema.TypeString,
				Description: "Specify outgoing interface to reach server.",
				Computed:    true,
			},
			"dhcp_relay_interface_select_method": {
				Type:        schema.TypeString,
				Description: "Specify how to select outgoing interface to reach server.",
				Computed:    true,
			},
			"dhcp_relay_ip": {
				Type:        schema.TypeString,
				Description: "DHCP relay IP address.",
				Computed:    true,
			},
			"dhcp_relay_request_all_server": {
				Type:        schema.TypeString,
				Description: "Enable/disable sending of DHCP requests to all servers.",
				Computed:    true,
			},
			"dhcp_relay_service": {
				Type:        schema.TypeString,
				Description: "Enable/disable allowing this interface to act as a DHCP relay.",
				Computed:    true,
			},
			"dhcp_relay_type": {
				Type:        schema.TypeString,
				Description: "DHCP relay type (regular or IPsec).",
				Computed:    true,
			},
			"dhcp_renew_time": {
				Type:        schema.TypeInt,
				Description: "DHCP renew time in seconds (300-604800), 0 means use the renew time provided by the server.",
				Computed:    true,
			},
			"dhcp_snooping_server_list": {
				Type:        schema.TypeList,
				Description: "Configure DHCP server access list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "DHCP server name.",
							Computed:    true,
						},
						"server_ip": {
							Type:        schema.TypeString,
							Description: "IP address for DHCP server.",
							Computed:    true,
						},
					},
				},
			},
			"disc_retry_timeout": {
				Type:        schema.TypeInt,
				Description: "Time in seconds to wait before retrying to start a PPPoE discovery, 0 means no timeout.",
				Computed:    true,
			},
			"disconnect_threshold": {
				Type:        schema.TypeInt,
				Description: "Time in milliseconds to wait before sending a notification that this interface is down or disconnected.",
				Computed:    true,
			},
			"distance": {
				Type:        schema.TypeInt,
				Description: "Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.",
				Computed:    true,
			},
			"dns_server_override": {
				Type:        schema.TypeString,
				Description: "Enable/disable use DNS acquired by DHCP or PPPoE.",
				Computed:    true,
			},
			"drop_fragment": {
				Type:        schema.TypeString,
				Description: "Enable/disable drop fragment packets.",
				Computed:    true,
			},
			"drop_overlapped_fragment": {
				Type:        schema.TypeString,
				Description: "Enable/disable drop overlapped fragment packets.",
				Computed:    true,
			},
			"egress_cos": {
				Type:        schema.TypeString,
				Description: "Override outgoing CoS in user VLAN tag.",
				Computed:    true,
			},
			"egress_queues": {
				Type:        schema.TypeList,
				Description: "Configure queues of NP port on egress path.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cos0": {
							Type:        schema.TypeString,
							Description: "CoS profile name for CoS 0.",
							Computed:    true,
						},
						"cos1": {
							Type:        schema.TypeString,
							Description: "CoS profile name for CoS 1.",
							Computed:    true,
						},
						"cos2": {
							Type:        schema.TypeString,
							Description: "CoS profile name for CoS 2.",
							Computed:    true,
						},
						"cos3": {
							Type:        schema.TypeString,
							Description: "CoS profile name for CoS 3.",
							Computed:    true,
						},
						"cos4": {
							Type:        schema.TypeString,
							Description: "CoS profile name for CoS 4.",
							Computed:    true,
						},
						"cos5": {
							Type:        schema.TypeString,
							Description: "CoS profile name for CoS 5.",
							Computed:    true,
						},
						"cos6": {
							Type:        schema.TypeString,
							Description: "CoS profile name for CoS 6.",
							Computed:    true,
						},
						"cos7": {
							Type:        schema.TypeString,
							Description: "CoS profile name for CoS 7.",
							Computed:    true,
						},
					},
				},
			},
			"egress_shaping_profile": {
				Type:        schema.TypeString,
				Description: "Outgoing traffic shaping profile.",
				Computed:    true,
			},
			"eip": {
				Type:        schema.TypeString,
				Description: "External IP.",
				Computed:    true,
			},
			"estimated_downstream_bandwidth": {
				Type:        schema.TypeInt,
				Description: "Estimated maximum downstream bandwidth (kbps). Used to estimate link utilization.",
				Computed:    true,
			},
			"estimated_upstream_bandwidth": {
				Type:        schema.TypeInt,
				Description: "Estimated maximum upstream bandwidth (kbps). Used to estimate link utilization.",
				Computed:    true,
			},
			"explicit_ftp_proxy": {
				Type:        schema.TypeString,
				Description: "Enable/disable the explicit FTP proxy on this interface.",
				Computed:    true,
			},
			"explicit_web_proxy": {
				Type:        schema.TypeString,
				Description: "Enable/disable the explicit web proxy on this interface.",
				Computed:    true,
			},
			"external": {
				Type:        schema.TypeString,
				Description: "Enable/disable identifying the interface as an external interface (which usually means it's connected to the Internet).",
				Computed:    true,
			},
			"fail_action_on_extender": {
				Type:        schema.TypeString,
				Description: "Action on extender when interface fail .",
				Computed:    true,
			},
			"fail_alert_interfaces": {
				Type:        schema.TypeList,
				Description: "Names of the FortiGate interfaces to which the link failure alert is sent.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Names of the non-virtual interface.",
							Computed:    true,
						},
					},
				},
			},
			"fail_alert_method": {
				Type:        schema.TypeString,
				Description: "Select link-failed-signal or link-down method to alert about a failed link.",
				Computed:    true,
			},
			"fail_detect": {
				Type:        schema.TypeString,
				Description: "Enable/disable fail detection features for this interface.",
				Computed:    true,
			},
			"fail_detect_option": {
				Type:        schema.TypeString,
				Description: "Options for detecting that this interface has failed.",
				Computed:    true,
			},
			"fortilink": {
				Type:        schema.TypeString,
				Description: "Enable FortiLink to dedicate this interface to manage other Fortinet devices.",
				Computed:    true,
			},
			"fortilink_backup_link": {
				Type:        schema.TypeInt,
				Description: "fortilink split interface backup link.",
				Computed:    true,
			},
			"fortilink_neighbor_detect": {
				Type:        schema.TypeString,
				Description: "Protocol for FortiGate neighbor discovery.",
				Computed:    true,
			},
			"fortilink_split_interface": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiLink split interface to connect member link to different FortiSwitch in stack for uplink redundancy.",
				Computed:    true,
			},
			"fortilink_stacking": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiLink switch-stacking on this interface.",
				Computed:    true,
			},
			"forward_domain": {
				Type:        schema.TypeInt,
				Description: "Transparent mode forward domain.",
				Computed:    true,
			},
			"gwdetect": {
				Type:        schema.TypeString,
				Description: "Enable/disable detect gateway alive for first.",
				Computed:    true,
			},
			"ha_priority": {
				Type:        schema.TypeInt,
				Description: "HA election priority for the PING server.",
				Computed:    true,
			},
			"icmp_accept_redirect": {
				Type:        schema.TypeString,
				Description: "Enable/disable ICMP accept redirect.",
				Computed:    true,
			},
			"icmp_send_redirect": {
				Type:        schema.TypeString,
				Description: "Enable/disable sending of ICMP redirects.",
				Computed:    true,
			},
			"ident_accept": {
				Type:        schema.TypeString,
				Description: "Enable/disable authentication for this interface.",
				Computed:    true,
			},
			"idle_timeout": {
				Type:        schema.TypeInt,
				Description: "PPPoE auto disconnect after idle timeout seconds, 0 means no timeout.",
				Computed:    true,
			},
			"inbandwidth": {
				Type:        schema.TypeInt,
				Description: "Bandwidth limit for incoming traffic (0 - 16776000 kbps), 0 means unlimited.",
				Computed:    true,
			},
			"ingress_cos": {
				Type:        schema.TypeString,
				Description: "Override incoming CoS in user VLAN tag on VLAN interface or assign a priority VLAN tag on physical interface.",
				Computed:    true,
			},
			"ingress_shaping_profile": {
				Type:        schema.TypeString,
				Description: "Incoming traffic shaping profile.",
				Computed:    true,
			},
			"ingress_spillover_threshold": {
				Type:        schema.TypeInt,
				Description: "Ingress Spillover threshold (0 - 16776000 kbps), 0 means unlimited.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Interface name.",
				Computed:    true,
			},
			"internal": {
				Type:        schema.TypeInt,
				Description: "Implicitly created.",
				Computed:    true,
			},
			"ip": {
				Type:        schema.TypeString,
				Description: "Interface IPv4 address and subnet mask, syntax: X.X.X.X/24.",
				Computed:    true,
			},
			"ip_managed_by_fortiipam": {
				Type:        schema.TypeString,
				Description: "Enable/disable automatic IP address assignment of this interface by FortiIPAM.",
				Computed:    true,
			},
			"ipmac": {
				Type:        schema.TypeString,
				Description: "Enable/disable IP/MAC binding.",
				Computed:    true,
			},
			"ips_sniffer_mode": {
				Type:        schema.TypeString,
				Description: "Enable/disable the use of this interface as a one-armed sniffer.",
				Computed:    true,
			},
			"ipunnumbered": {
				Type:        schema.TypeString,
				Description: "Unnumbered IP used for PPPoE interfaces for which no unique local address is provided.",
				Computed:    true,
			},
			"ipv6": {
				Type:        schema.TypeList,
				Description: "IPv6 of interface.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"autoconf": {
							Type:        schema.TypeString,
							Description: "Enable/disable address auto config.",
							Computed:    true,
						},
						"cli_conn6_status": {
							Type:        schema.TypeInt,
							Description: "CLI IPv6 connection status.",
							Computed:    true,
						},
						"dhcp6_client_options": {
							Type:        schema.TypeString,
							Description: "DHCPv6 client options.",
							Computed:    true,
						},
						"dhcp6_iapd_list": {
							Type:        schema.TypeList,
							Description: "DHCPv6 IA-PD list",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"iaid": {
										Type:        schema.TypeInt,
										Description: "Identity association identifier.",
										Computed:    true,
									},
									"prefix_hint": {
										Type:        schema.TypeString,
										Description: "DHCPv6 prefix that will be used as a hint to the upstream DHCPv6 server.",
										Computed:    true,
									},
									"prefix_hint_plt": {
										Type:        schema.TypeInt,
										Description: "DHCPv6 prefix hint preferred life time (sec), 0 means unlimited lease time.",
										Computed:    true,
									},
									"prefix_hint_vlt": {
										Type:        schema.TypeInt,
										Description: "DHCPv6 prefix hint valid life time (sec).",
										Computed:    true,
									},
								},
							},
						},
						"dhcp6_information_request": {
							Type:        schema.TypeString,
							Description: "Enable/disable DHCPv6 information request.",
							Computed:    true,
						},
						"dhcp6_prefix_delegation": {
							Type:        schema.TypeString,
							Description: "Enable/disable DHCPv6 prefix delegation.",
							Computed:    true,
						},
						"dhcp6_prefix_hint": {
							Type:        schema.TypeString,
							Description: "DHCPv6 prefix that will be used as a hint to the upstream DHCPv6 server.",
							Computed:    true,
						},
						"dhcp6_prefix_hint_plt": {
							Type:        schema.TypeInt,
							Description: "DHCPv6 prefix hint preferred life time (sec), 0 means unlimited lease time.",
							Computed:    true,
						},
						"dhcp6_prefix_hint_vlt": {
							Type:        schema.TypeInt,
							Description: "DHCPv6 prefix hint valid life time (sec).",
							Computed:    true,
						},
						"dhcp6_relay_ip": {
							Type:        schema.TypeString,
							Description: "DHCPv6 relay IP address.",
							Computed:    true,
						},
						"dhcp6_relay_service": {
							Type:        schema.TypeString,
							Description: "Enable/disable DHCPv6 relay.",
							Computed:    true,
						},
						"dhcp6_relay_type": {
							Type:        schema.TypeString,
							Description: "DHCPv6 relay type.",
							Computed:    true,
						},
						"icmp6_send_redirect": {
							Type:        schema.TypeString,
							Description: "Enable/disable sending of ICMPv6 redirects.",
							Computed:    true,
						},
						"interface_identifier": {
							Type:        schema.TypeString,
							Description: "IPv6 interface identifier.",
							Computed:    true,
						},
						"ip6_address": {
							Type:        schema.TypeString,
							Description: "Primary IPv6 address prefix, syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx",
							Computed:    true,
						},
						"ip6_allowaccess": {
							Type:        schema.TypeString,
							Description: "Allow management access to the interface.",
							Computed:    true,
						},
						"ip6_default_life": {
							Type:        schema.TypeInt,
							Description: "Default life (sec).",
							Computed:    true,
						},
						"ip6_delegated_prefix_iaid": {
							Type:        schema.TypeInt,
							Description: "IAID of obtained delegated-prefix from the upstream interface.",
							Computed:    true,
						},
						"ip6_delegated_prefix_list": {
							Type:        schema.TypeList,
							Description: "Advertised IPv6 delegated prefix list.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"autonomous_flag": {
										Type:        schema.TypeString,
										Description: "Enable/disable the autonomous flag.",
										Computed:    true,
									},
									"delegated_prefix_iaid": {
										Type:        schema.TypeInt,
										Description: "IAID of obtained delegated-prefix from the upstream interface.",
										Computed:    true,
									},
									"onlink_flag": {
										Type:        schema.TypeString,
										Description: "Enable/disable the onlink flag.",
										Computed:    true,
									},
									"prefix_id": {
										Type:        schema.TypeInt,
										Description: "Prefix ID.",
										Computed:    true,
									},
									"rdnss": {
										Type:        schema.TypeString,
										Description: "Recursive DNS server option.",
										Computed:    true,
									},
									"rdnss_service": {
										Type:        schema.TypeString,
										Description: "Recursive DNS service option.",
										Computed:    true,
									},
									"subnet": {
										Type:        schema.TypeString,
										Description: " Add subnet ID to routing prefix.",
										Computed:    true,
									},
									"upstream_interface": {
										Type:        schema.TypeString,
										Description: "Name of the interface that provides delegated information.",
										Computed:    true,
									},
								},
							},
						},
						"ip6_dns_server_override": {
							Type:        schema.TypeString,
							Description: "Enable/disable using the DNS server acquired by DHCP.",
							Computed:    true,
						},
						"ip6_extra_addr": {
							Type:        schema.TypeList,
							Description: "Extra IPv6 address prefixes of interface.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": {
										Type:        schema.TypeString,
										Description: "IPv6 address prefix.",
										Computed:    true,
									},
								},
							},
						},
						"ip6_hop_limit": {
							Type:        schema.TypeInt,
							Description: "Hop limit (0 means unspecified).",
							Computed:    true,
						},
						"ip6_link_mtu": {
							Type:        schema.TypeInt,
							Description: "IPv6 link MTU.",
							Computed:    true,
						},
						"ip6_manage_flag": {
							Type:        schema.TypeString,
							Description: "Enable/disable the managed flag.",
							Computed:    true,
						},
						"ip6_max_interval": {
							Type:        schema.TypeInt,
							Description: "IPv6 maximum interval (4 to 1800 sec).",
							Computed:    true,
						},
						"ip6_min_interval": {
							Type:        schema.TypeInt,
							Description: "IPv6 minimum interval (3 to 1350 sec).",
							Computed:    true,
						},
						"ip6_mode": {
							Type:        schema.TypeString,
							Description: "Addressing mode (static, DHCP, delegated).",
							Computed:    true,
						},
						"ip6_other_flag": {
							Type:        schema.TypeString,
							Description: "Enable/disable the other IPv6 flag.",
							Computed:    true,
						},
						"ip6_prefix_list": {
							Type:        schema.TypeList,
							Description: "Advertised prefix list.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"autonomous_flag": {
										Type:        schema.TypeString,
										Description: "Enable/disable the autonomous flag.",
										Computed:    true,
									},
									"dnssl": {
										Type:        schema.TypeList,
										Description: "DNS search list option.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"domain": {
													Type:        schema.TypeString,
													Description: "Domain name.",
													Computed:    true,
												},
											},
										},
									},
									"onlink_flag": {
										Type:        schema.TypeString,
										Description: "Enable/disable the onlink flag.",
										Computed:    true,
									},
									"preferred_life_time": {
										Type:        schema.TypeInt,
										Description: "Preferred life time (sec).",
										Computed:    true,
									},
									"prefix": {
										Type:        schema.TypeString,
										Description: "IPv6 prefix.",
										Computed:    true,
									},
									"rdnss": {
										Type:        schema.TypeString,
										Description: "Recursive DNS server option.",
										Computed:    true,
									},
									"valid_life_time": {
										Type:        schema.TypeInt,
										Description: "Valid life time (sec).",
										Computed:    true,
									},
								},
							},
						},
						"ip6_prefix_mode": {
							Type:        schema.TypeString,
							Description: "Assigning a prefix from DHCP or RA.",
							Computed:    true,
						},
						"ip6_reachable_time": {
							Type:        schema.TypeInt,
							Description: "IPv6 reachable time (milliseconds; 0 means unspecified).",
							Computed:    true,
						},
						"ip6_retrans_time": {
							Type:        schema.TypeInt,
							Description: "IPv6 retransmit time (milliseconds; 0 means unspecified).",
							Computed:    true,
						},
						"ip6_send_adv": {
							Type:        schema.TypeString,
							Description: "Enable/disable sending advertisements about the interface.",
							Computed:    true,
						},
						"ip6_subnet": {
							Type:        schema.TypeString,
							Description: " Subnet to routing prefix, syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx",
							Computed:    true,
						},
						"ip6_upstream_interface": {
							Type:        schema.TypeString,
							Description: "Interface name providing delegated information.",
							Computed:    true,
						},
						"nd_cert": {
							Type:        schema.TypeString,
							Description: "Neighbor discovery certificate.",
							Computed:    true,
						},
						"nd_cga_modifier": {
							Type:        schema.TypeString,
							Description: "Neighbor discovery CGA modifier.",
							Computed:    true,
						},
						"nd_mode": {
							Type:        schema.TypeString,
							Description: "Neighbor discovery mode.",
							Computed:    true,
						},
						"nd_security_level": {
							Type:        schema.TypeInt,
							Description: "Neighbor discovery security level (0 - 7; 0 = least secure, default = 0).",
							Computed:    true,
						},
						"nd_timestamp_delta": {
							Type:        schema.TypeInt,
							Description: "Neighbor discovery timestamp delta value (1 - 3600 sec; default = 300).",
							Computed:    true,
						},
						"nd_timestamp_fuzz": {
							Type:        schema.TypeInt,
							Description: "Neighbor discovery timestamp fuzz factor (1 - 60 sec; default = 1).",
							Computed:    true,
						},
						"ra_send_mtu": {
							Type:        schema.TypeString,
							Description: "Enable/disable sending link MTU in RA packet.",
							Computed:    true,
						},
						"unique_autoconf_addr": {
							Type:        schema.TypeString,
							Description: "Enable/disable unique auto config address.",
							Computed:    true,
						},
						"vrip6_link_local": {
							Type:        schema.TypeString,
							Description: "Link-local IPv6 address of virtual router.",
							Computed:    true,
						},
						"vrrp_virtual_mac6": {
							Type:        schema.TypeString,
							Description: "Enable/disable virtual MAC for VRRP.",
							Computed:    true,
						},
						"vrrp6": {
							Type:        schema.TypeList,
							Description: "IPv6 VRRP configuration.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"accept_mode": {
										Type:        schema.TypeString,
										Description: "Enable/disable accept mode.",
										Computed:    true,
									},
									"adv_interval": {
										Type:        schema.TypeInt,
										Description: "Advertisement interval (1 - 255 seconds).",
										Computed:    true,
									},
									"preempt": {
										Type:        schema.TypeString,
										Description: "Enable/disable preempt mode.",
										Computed:    true,
									},
									"priority": {
										Type:        schema.TypeInt,
										Description: "Priority of the virtual router (1 - 255).",
										Computed:    true,
									},
									"start_time": {
										Type:        schema.TypeInt,
										Description: "Startup time (1 - 255 seconds).",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Enable/disable VRRP.",
										Computed:    true,
									},
									"vrdst6": {
										Type:        schema.TypeString,
										Description: "Monitor the route to this destination.",
										Computed:    true,
									},
									"vrgrp": {
										Type:        schema.TypeInt,
										Description: "VRRP group ID (1 - 65535).",
										Computed:    true,
									},
									"vrid": {
										Type:        schema.TypeInt,
										Description: "Virtual router identifier (1 - 255).",
										Computed:    true,
									},
									"vrip6": {
										Type:        schema.TypeString,
										Description: "IPv6 address of the virtual router.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"l2forward": {
				Type:        schema.TypeString,
				Description: "Enable/disable l2 forwarding.",
				Computed:    true,
			},
			"lacp_ha_slave": {
				Type:        schema.TypeString,
				Description: "LACP HA slave.",
				Computed:    true,
			},
			"lacp_mode": {
				Type:        schema.TypeString,
				Description: "LACP mode.",
				Computed:    true,
			},
			"lacp_speed": {
				Type:        schema.TypeString,
				Description: "How often the interface sends LACP messages.",
				Computed:    true,
			},
			"lcp_echo_interval": {
				Type:        schema.TypeInt,
				Description: "Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.",
				Computed:    true,
			},
			"lcp_max_echo_fails": {
				Type:        schema.TypeInt,
				Description: "Maximum missed LCP echo messages before disconnect.",
				Computed:    true,
			},
			"link_up_delay": {
				Type:        schema.TypeInt,
				Description: "Number of milliseconds to wait before considering a link is up.",
				Computed:    true,
			},
			"lldp_network_policy": {
				Type:        schema.TypeString,
				Description: "LLDP-MED network policy profile.",
				Computed:    true,
			},
			"lldp_reception": {
				Type:        schema.TypeString,
				Description: "Enable/disable Link Layer Discovery Protocol (LLDP) reception.",
				Computed:    true,
			},
			"lldp_transmission": {
				Type:        schema.TypeString,
				Description: "Enable/disable Link Layer Discovery Protocol (LLDP) transmission.",
				Computed:    true,
			},
			"macaddr": {
				Type:        schema.TypeString,
				Description: "Change the interface's MAC address.",
				Computed:    true,
			},
			"managed_subnetwork_size": {
				Type:        schema.TypeString,
				Description: "Number of IP addresses to be allocated by FortiIPAM and used by this FortiGate unit's DHCP server settings.",
				Computed:    true,
			},
			"management_ip": {
				Type:        schema.TypeString,
				Description: "High Availability in-band management IP address of this interface.",
				Computed:    true,
			},
			"measured_downstream_bandwidth": {
				Type:        schema.TypeInt,
				Description: "Measured downstream bandwidth (kbps). ",
				Computed:    true,
			},
			"measured_upstream_bandwidth": {
				Type:        schema.TypeInt,
				Description: "Measured upstream bandwidth (kbps). ",
				Computed:    true,
			},
			"mediatype": {
				Type:        schema.TypeString,
				Description: "Select SFP media interface type",
				Computed:    true,
			},
			"member": {
				Type:        schema.TypeList,
				Description: "Physical interfaces that belong to the aggregate or redundant interface.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": {
							Type:        schema.TypeString,
							Description: "Physical interface name.",
							Computed:    true,
						},
					},
				},
			},
			"min_links": {
				Type:        schema.TypeInt,
				Description: "Minimum number of aggregated ports that must be up.",
				Computed:    true,
			},
			"min_links_down": {
				Type:        schema.TypeString,
				Description: "Action to take when less than the configured minimum number of links are active.",
				Computed:    true,
			},
			"mode": {
				Type:        schema.TypeString,
				Description: "Addressing mode (static, DHCP, PPPoE).",
				Computed:    true,
			},
			"monitor_bandwidth": {
				Type:        schema.TypeString,
				Description: "Enable monitoring bandwidth on this interface.",
				Computed:    true,
			},
			"mtu": {
				Type:        schema.TypeInt,
				Description: "MTU value for this interface.",
				Computed:    true,
			},
			"mtu_override": {
				Type:        schema.TypeString,
				Description: "Enable to set a custom MTU for this interface.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name.",
				Required:    true,
			},
			"ndiscforward": {
				Type:        schema.TypeString,
				Description: "Enable/disable NDISC forwarding.",
				Computed:    true,
			},
			"netbios_forward": {
				Type:        schema.TypeString,
				Description: "Enable/disable NETBIOS forwarding.",
				Computed:    true,
			},
			"netflow_sampler": {
				Type:        schema.TypeString,
				Description: "Enable/disable NetFlow on this interface and set the data that NetFlow collects (rx, tx, or both).",
				Computed:    true,
			},
			"outbandwidth": {
				Type:        schema.TypeInt,
				Description: "Bandwidth limit for outgoing traffic (0 - 16776000 kbps), 0 means unlimited.",
				Computed:    true,
			},
			"padt_retry_timeout": {
				Type:        schema.TypeInt,
				Description: "PPPoE Active Discovery Terminate (PADT) used to terminate sessions after an idle time.",
				Computed:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "PPPoE account's password.",
				Computed:    true,
				Sensitive:   true,
			},
			"ping_serv_status": {
				Type:        schema.TypeInt,
				Description: "PING server status.",
				Computed:    true,
			},
			"polling_interval": {
				Type:        schema.TypeInt,
				Description: "sFlow polling interval (1 - 255 sec).",
				Computed:    true,
			},
			"pppoe_unnumbered_negotiate": {
				Type:        schema.TypeString,
				Description: "Enable/disable PPPoE unnumbered negotiation.",
				Computed:    true,
			},
			"pptp_auth_type": {
				Type:        schema.TypeString,
				Description: "PPTP authentication type.",
				Computed:    true,
			},
			"pptp_client": {
				Type:        schema.TypeString,
				Description: "Enable/disable PPTP client.",
				Computed:    true,
			},
			"pptp_password": {
				Type:        schema.TypeString,
				Description: "PPTP password.",
				Computed:    true,
				Sensitive:   true,
			},
			"pptp_server_ip": {
				Type:        schema.TypeString,
				Description: "PPTP server IP address.",
				Computed:    true,
			},
			"pptp_timeout": {
				Type:        schema.TypeInt,
				Description: "Idle timer in minutes (0 for disabled).",
				Computed:    true,
			},
			"pptp_user": {
				Type:        schema.TypeString,
				Description: "PPTP user name.",
				Computed:    true,
			},
			"preserve_session_route": {
				Type:        schema.TypeString,
				Description: "Enable/disable preservation of session route when dirty.",
				Computed:    true,
			},
			"priority": {
				Type:        schema.TypeInt,
				Description: "Priority of learned routes.",
				Computed:    true,
			},
			"priority_override": {
				Type:        schema.TypeString,
				Description: "Enable/disable fail back to higher priority port once recovered.",
				Computed:    true,
			},
			"proxy_captive_portal": {
				Type:        schema.TypeString,
				Description: "Enable/disable proxy captive portal on this interface.",
				Computed:    true,
			},
			"redundant_interface": {
				Type:        schema.TypeString,
				Description: "Redundant interface.",
				Computed:    true,
			},
			"remote_ip": {
				Type:        schema.TypeString,
				Description: "Remote IP address of tunnel.",
				Computed:    true,
			},
			"replacemsg_override_group": {
				Type:        schema.TypeString,
				Description: "Replacement message override group.",
				Computed:    true,
			},
			"ring_rx": {
				Type:        schema.TypeInt,
				Description: "RX ring size.",
				Computed:    true,
			},
			"ring_tx": {
				Type:        schema.TypeInt,
				Description: "TX ring size.",
				Computed:    true,
			},
			"role": {
				Type:        schema.TypeString,
				Description: "Interface role.",
				Computed:    true,
			},
			"sample_direction": {
				Type:        schema.TypeString,
				Description: "Data that NetFlow collects (rx, tx, or both).",
				Computed:    true,
			},
			"sample_rate": {
				Type:        schema.TypeInt,
				Description: "sFlow sample rate (10 - 99999).",
				Computed:    true,
			},
			"secondary_ip": {
				Type:        schema.TypeString,
				Description: "Enable/disable adding a secondary IP to this interface.",
				Computed:    true,
			},
			"secondaryip": {
				Type:        schema.TypeList,
				Description: "Second IP address of interface.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allowaccess": {
							Type:        schema.TypeString,
							Description: "Management access settings for the secondary IP address.",
							Computed:    true,
						},
						"detectprotocol": {
							Type:        schema.TypeString,
							Description: "Protocols used to detect the server.",
							Computed:    true,
						},
						"detectserver": {
							Type:        schema.TypeString,
							Description: "Gateway's ping server for this IP.",
							Computed:    true,
						},
						"gwdetect": {
							Type:        schema.TypeString,
							Description: "Enable/disable detect gateway alive for first.",
							Computed:    true,
						},
						"ha_priority": {
							Type:        schema.TypeInt,
							Description: "HA election priority for the PING server.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"ip": {
							Type:        schema.TypeString,
							Description: "Secondary IP address of the interface.",
							Computed:    true,
						},
						"ping_serv_status": {
							Type:        schema.TypeInt,
							Description: "PING server status.",
							Computed:    true,
						},
					},
				},
			},
			"security_exempt_list": {
				Type:        schema.TypeString,
				Description: "Name of security-exempt-list.",
				Computed:    true,
			},
			"security_external_logout": {
				Type:        schema.TypeString,
				Description: "URL of external authentication logout server.",
				Computed:    true,
			},
			"security_external_web": {
				Type:        schema.TypeString,
				Description: "URL of external authentication web server.",
				Computed:    true,
			},
			"security_groups": {
				Type:        schema.TypeList,
				Description: "User groups that can authenticate with the captive portal.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Names of user groups that can authenticate with the captive portal.",
							Computed:    true,
						},
					},
				},
			},
			"security_mac_auth_bypass": {
				Type:        schema.TypeString,
				Description: "Enable/disable MAC authentication bypass.",
				Computed:    true,
			},
			"security_mode": {
				Type:        schema.TypeString,
				Description: "Turn on captive portal authentication for this interface.",
				Computed:    true,
			},
			"security_redirect_url": {
				Type:        schema.TypeString,
				Description: "URL redirection after disclaimer/authentication.",
				Computed:    true,
			},
			"service_name": {
				Type:        schema.TypeString,
				Description: "PPPoE service name.",
				Computed:    true,
			},
			"sflow_sampler": {
				Type:        schema.TypeString,
				Description: "Enable/disable sFlow on this interface.",
				Computed:    true,
			},
			"snmp_index": {
				Type:        schema.TypeInt,
				Description: "Permanent SNMP Index of the interface.",
				Computed:    true,
			},
			"speed": {
				Type:        schema.TypeString,
				Description: "Interface speed. The default setting and the options available depend on the interface hardware.",
				Computed:    true,
			},
			"spillover_threshold": {
				Type:        schema.TypeInt,
				Description: "Egress Spillover threshold (0 - 16776000 kbps), 0 means unlimited.",
				Computed:    true,
			},
			"src_check": {
				Type:        schema.TypeString,
				Description: "Enable/disable source IP check.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Bring the interface up or shut the interface down.",
				Computed:    true,
			},
			"stpforward": {
				Type:        schema.TypeString,
				Description: "Enable/disable STP forwarding.",
				Computed:    true,
			},
			"stpforward_mode": {
				Type:        schema.TypeString,
				Description: "Configure STP forwarding mode.",
				Computed:    true,
			},
			"subst": {
				Type:        schema.TypeString,
				Description: "Enable to always send packets from this interface to a destination MAC address.",
				Computed:    true,
			},
			"substitute_dst_mac": {
				Type:        schema.TypeString,
				Description: "Destination MAC address that all packets are sent to from this interface.",
				Computed:    true,
			},
			"swc_first_create": {
				Type:        schema.TypeInt,
				Description: "Initial create for switch-controller VLANs.",
				Computed:    true,
			},
			"swc_vlan": {
				Type:        schema.TypeInt,
				Description: "Creation status for switch-controller VLANs.",
				Computed:    true,
			},
			"switch": {
				Type:        schema.TypeString,
				Description: "Contained in switch.",
				Computed:    true,
			},
			"switch_controller_access_vlan": {
				Type:        schema.TypeString,
				Description: "Block FortiSwitch port-to-port traffic.",
				Computed:    true,
			},
			"switch_controller_arp_inspection": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiSwitch ARP inspection.",
				Computed:    true,
			},
			"switch_controller_dhcp_snooping": {
				Type:        schema.TypeString,
				Description: "Switch controller DHCP snooping.",
				Computed:    true,
			},
			"switch_controller_dhcp_snooping_option82": {
				Type:        schema.TypeString,
				Description: "Switch controller DHCP snooping option82.",
				Computed:    true,
			},
			"switch_controller_dhcp_snooping_verify_mac": {
				Type:        schema.TypeString,
				Description: "Switch controller DHCP snooping verify MAC.",
				Computed:    true,
			},
			"switch_controller_dynamic": {
				Type:        schema.TypeString,
				Description: "Integrated FortiLink settings for managed FortiSwitch.",
				Computed:    true,
			},
			"switch_controller_feature": {
				Type:        schema.TypeString,
				Description: "Interface's purpose when assigning traffic (read only).",
				Computed:    true,
			},
			"switch_controller_igmp_snooping": {
				Type:        schema.TypeString,
				Description: "Switch controller IGMP snooping.",
				Computed:    true,
			},
			"switch_controller_igmp_snooping_fast_leave": {
				Type:        schema.TypeString,
				Description: "Switch controller IGMP snooping fast-leave.",
				Computed:    true,
			},
			"switch_controller_igmp_snooping_proxy": {
				Type:        schema.TypeString,
				Description: "Switch controller IGMP snooping proxy.",
				Computed:    true,
			},
			"switch_controller_iot_scanning": {
				Type:        schema.TypeString,
				Description: "Enable/disable managed FortiSwitch IoT scanning.",
				Computed:    true,
			},
			"switch_controller_learning_limit": {
				Type:        schema.TypeInt,
				Description: "Limit the number of dynamic MAC addresses on this VLAN (1 - 128, 0 = no limit, default).",
				Computed:    true,
			},
			"switch_controller_mgmt_vlan": {
				Type:        schema.TypeInt,
				Description: "VLAN to use for FortiLink management purposes.",
				Computed:    true,
			},
			"switch_controller_nac": {
				Type:        schema.TypeString,
				Description: "Integrated FortiLink settings for managed FortiSwitch.",
				Computed:    true,
			},
			"switch_controller_rspan_mode": {
				Type:        schema.TypeString,
				Description: "Stop Layer2 MAC learning and interception of BPDUs and other packets on this interface.",
				Computed:    true,
			},
			"switch_controller_source_ip": {
				Type:        schema.TypeString,
				Description: "Source IP address used in FortiLink over L3 connections.",
				Computed:    true,
			},
			"switch_controller_traffic_policy": {
				Type:        schema.TypeString,
				Description: "Switch controller traffic policy for the VLAN.",
				Computed:    true,
			},
			"system_id": {
				Type:        schema.TypeString,
				Description: "Define a system ID for the aggregate interface.",
				Computed:    true,
			},
			"system_id_type": {
				Type:        schema.TypeString,
				Description: "Method in which system ID is generated.",
				Computed:    true,
			},
			"tagging": {
				Type:        schema.TypeList,
				Description: "Config object tagging.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:        schema.TypeString,
							Description: "Tag category.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Tagging entry name.",
							Computed:    true,
						},
						"tags": {
							Type:        schema.TypeList,
							Description: "Tags.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Tag name.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"tcp_mss": {
				Type:        schema.TypeInt,
				Description: "TCP maximum segment size. 0 means do not change segment size.",
				Computed:    true,
			},
			"trust_ip_1": {
				Type:        schema.TypeString,
				Description: "Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).",
				Computed:    true,
			},
			"trust_ip_2": {
				Type:        schema.TypeString,
				Description: "Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).",
				Computed:    true,
			},
			"trust_ip_3": {
				Type:        schema.TypeString,
				Description: "Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).",
				Computed:    true,
			},
			"trust_ip6_1": {
				Type:        schema.TypeString,
				Description: "Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).",
				Computed:    true,
			},
			"trust_ip6_2": {
				Type:        schema.TypeString,
				Description: "Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).",
				Computed:    true,
			},
			"trust_ip6_3": {
				Type:        schema.TypeString,
				Description: "Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Interface type.",
				Computed:    true,
			},
			"username": {
				Type:        schema.TypeString,
				Description: "Username of the PPPoE account, provided by your ISP.",
				Computed:    true,
			},
			"vdom": {
				Type:        schema.TypeString,
				Description: "Interface is in this virtual domain (VDOM).",
				Computed:    true,
			},
			"vindex": {
				Type:        schema.TypeInt,
				Description: "Switch control interface VLAN ID.",
				Computed:    true,
			},
			"vlan_protocol": {
				Type:        schema.TypeString,
				Description: "Ethernet protocol of VLAN.",
				Computed:    true,
			},
			"vlanforward": {
				Type:        schema.TypeString,
				Description: "Enable/disable traffic forwarding between VLANs on this interface.",
				Computed:    true,
			},
			"vlanid": {
				Type:        schema.TypeInt,
				Description: "VLAN ID (1 - 4094).",
				Computed:    true,
			},
			"vrf": {
				Type:        schema.TypeInt,
				Description: "Virtual Routing Forwarding ID.",
				Computed:    true,
			},
			"vrrp": {
				Type:        schema.TypeList,
				Description: "VRRP configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accept_mode": {
							Type:        schema.TypeString,
							Description: "Enable/disable accept mode.",
							Computed:    true,
						},
						"adv_interval": {
							Type:        schema.TypeInt,
							Description: "Advertisement interval (1 - 255 seconds).",
							Computed:    true,
						},
						"ignore_default_route": {
							Type:        schema.TypeString,
							Description: "Enable/disable ignoring of default route when checking destination.",
							Computed:    true,
						},
						"preempt": {
							Type:        schema.TypeString,
							Description: "Enable/disable preempt mode.",
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "Priority of the virtual router (1 - 255).",
							Computed:    true,
						},
						"proxy_arp": {
							Type:        schema.TypeList,
							Description: "VRRP Proxy ARP configuration.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Description: "ID.",
										Computed:    true,
									},
									"ip": {
										Type:        schema.TypeString,
										Description: "Set IP addresses of proxy ARP.",
										Computed:    true,
									},
								},
							},
						},
						"start_time": {
							Type:        schema.TypeInt,
							Description: "Startup time (1 - 255 seconds).",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable this VRRP configuration.",
							Computed:    true,
						},
						"version": {
							Type:        schema.TypeString,
							Description: "VRRP version.",
							Computed:    true,
						},
						"vrdst": {
							Type:        schema.TypeString,
							Description: "Monitor the route to this destination.",
							Computed:    true,
						},
						"vrdst_priority": {
							Type:        schema.TypeInt,
							Description: "Priority of the virtual router when the virtual router destination becomes unreachable (0 - 254).",
							Computed:    true,
						},
						"vrgrp": {
							Type:        schema.TypeInt,
							Description: "VRRP group ID (1 - 65535).",
							Computed:    true,
						},
						"vrid": {
							Type:        schema.TypeInt,
							Description: "Virtual router identifier (1 - 255).",
							Computed:    true,
						},
						"vrip": {
							Type:        schema.TypeString,
							Description: "IP address of the virtual router.",
							Computed:    true,
						},
					},
				},
			},
			"vrrp_virtual_mac": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of virtual MAC for VRRP.",
				Computed:    true,
			},
			"wccp": {
				Type:        schema.TypeString,
				Description: "Enable/disable WCCP on this interface. Used for encapsulated WCCP communication between WCCP clients and servers.",
				Computed:    true,
			},
			"weight": {
				Type:        schema.TypeInt,
				Description: "Default weight for static routes (if route has no weight configured).",
				Computed:    true,
			},
			"wins_ip": {
				Type:        schema.TypeString,
				Description: "WINS server IP.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemInterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.Errorf("error reading SystemInterface dataSource: %v", err)
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

	diags := refreshObjectSystemInterface(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
