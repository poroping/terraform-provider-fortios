// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.6,v7.0.2,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure NPU attributes.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemNpu() *schema.Resource {
	return &schema.Resource{
		Description: "Configure NPU attributes.",

		ReadContext: dataSourceSystemNpuRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"capwap_offload": {
				Type:        schema.TypeString,
				Description: "Enable/disable offloading managed FortiAP and FortiLink CAPWAP sessions.",
				Computed:    true,
			},
			"dedicated_management_affinity": {
				Type:        schema.TypeString,
				Description: "Affinity setting for management daemons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).",
				Computed:    true,
			},
			"dedicated_management_cpu": {
				Type:        schema.TypeString,
				Description: "Enable to dedicate one CPU for GUI and CLI connections when NPs are busy.",
				Computed:    true,
			},
			"default_qos_type": {
				Type:        schema.TypeString,
				Description: "Set default QoS type.",
				Computed:    true,
			},
			"dos_options": {
				Type:        schema.TypeList,
				Description: "NPU DoS configurations.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"npu_dos_meter_mode": {
							Type:        schema.TypeString,
							Description: "Set DoS meter NPU offloading mode.",
							Computed:    true,
						},
						"npu_dos_tpe_mode": {
							Type:        schema.TypeString,
							Description: "Enable/disable insertion of DoS meter ID to session table.",
							Computed:    true,
						},
					},
				},
			},
			"double_level_mcast_offload": {
				Type:        schema.TypeString,
				Description: "Enable double level mcast offload.",
				Computed:    true,
			},
			"dse_timeout": {
				Type:        schema.TypeInt,
				Description: " DSE timeout in seconds (0-3600, default = 10).",
				Computed:    true,
			},
			"dsw_dts_profile": {
				Type:        schema.TypeList,
				Description: "Configure NPU DSW DTS profile.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Set NPU DSW DTS profile action.",
							Computed:    true,
						},
						"min_limit": {
							Type:        schema.TypeInt,
							Description: "Set NPU DSW DTS profile min-limt.",
							Computed:    true,
						},
						"profile_id": {
							Type:        schema.TypeInt,
							Description: "Set NPU DSW DTS profile profile id.",
							Computed:    true,
						},
						"step": {
							Type:        schema.TypeInt,
							Description: "Set NPU DSW DTS profile step.",
							Computed:    true,
						},
					},
				},
			},
			"dsw_queue_dts_profile": {
				Type:        schema.TypeList,
				Description: "Configure NPU DSW Queue DTS profile.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iport": {
							Type:        schema.TypeString,
							Description: "Set NPU DSW DTS in port.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Name.",
							Computed:    true,
						},
						"oport": {
							Type:        schema.TypeString,
							Description: "Set NPU DSW DTS out port.",
							Computed:    true,
						},
						"profile_id": {
							Type:        schema.TypeInt,
							Description: "Set NPU DSW DTS profile ID.",
							Computed:    true,
						},
						"queue_select": {
							Type:        schema.TypeInt,
							Description: "Set NPU DSW DTS queue ID select (0 - reset to default).",
							Computed:    true,
						},
					},
				},
			},
			"fastpath": {
				Type:        schema.TypeString,
				Description: "Enable/disable NP6 offloading (also called fast path).",
				Computed:    true,
			},
			"fp_anomaly": {
				Type:        schema.TypeList,
				Description: "IPv4/IPv6 anomaly protection.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icmp_csum_err": {
							Type:        schema.TypeString,
							Description: "Invalid IPv4 ICMP checksum anomalies.",
							Computed:    true,
						},
						"icmp_frag": {
							Type:        schema.TypeString,
							Description: "Layer 3 fragmented packets that could be part of layer 4 ICMP anomalies.",
							Computed:    true,
						},
						"icmp_land": {
							Type:        schema.TypeString,
							Description: "ICMP land anomalies.",
							Computed:    true,
						},
						"ipv4_csum_err": {
							Type:        schema.TypeString,
							Description: "Invalid IPv4 IP checksum anomalies.",
							Computed:    true,
						},
						"ipv4_land": {
							Type:        schema.TypeString,
							Description: "Land anomalies.",
							Computed:    true,
						},
						"ipv4_optlsrr": {
							Type:        schema.TypeString,
							Description: "Loose source record route option anomalies.",
							Computed:    true,
						},
						"ipv4_optrr": {
							Type:        schema.TypeString,
							Description: "Record route option anomalies.",
							Computed:    true,
						},
						"ipv4_optsecurity": {
							Type:        schema.TypeString,
							Description: "Security option anomalies.",
							Computed:    true,
						},
						"ipv4_optssrr": {
							Type:        schema.TypeString,
							Description: "Strict source record route option anomalies.",
							Computed:    true,
						},
						"ipv4_optstream": {
							Type:        schema.TypeString,
							Description: "Stream option anomalies.",
							Computed:    true,
						},
						"ipv4_opttimestamp": {
							Type:        schema.TypeString,
							Description: "Timestamp option anomalies.",
							Computed:    true,
						},
						"ipv4_proto_err": {
							Type:        schema.TypeString,
							Description: "Invalid layer 4 protocol anomalies.",
							Computed:    true,
						},
						"ipv4_unknopt": {
							Type:        schema.TypeString,
							Description: "Unknown option anomalies.",
							Computed:    true,
						},
						"ipv6_daddr_err": {
							Type:        schema.TypeString,
							Description: "Destination address as unspecified or loopback address anomalies.",
							Computed:    true,
						},
						"ipv6_land": {
							Type:        schema.TypeString,
							Description: "Land anomalies.",
							Computed:    true,
						},
						"ipv6_optendpid": {
							Type:        schema.TypeString,
							Description: "End point identification anomalies.",
							Computed:    true,
						},
						"ipv6_opthomeaddr": {
							Type:        schema.TypeString,
							Description: "Home address option anomalies.",
							Computed:    true,
						},
						"ipv6_optinvld": {
							Type:        schema.TypeString,
							Description: "Invalid option anomalies.Invalid option anomalies.",
							Computed:    true,
						},
						"ipv6_optjumbo": {
							Type:        schema.TypeString,
							Description: "Jumbo options anomalies.",
							Computed:    true,
						},
						"ipv6_optnsap": {
							Type:        schema.TypeString,
							Description: "Network service access point address option anomalies.",
							Computed:    true,
						},
						"ipv6_optralert": {
							Type:        schema.TypeString,
							Description: "Router alert option anomalies.",
							Computed:    true,
						},
						"ipv6_opttunnel": {
							Type:        schema.TypeString,
							Description: "Tunnel encapsulation limit option anomalies.",
							Computed:    true,
						},
						"ipv6_proto_err": {
							Type:        schema.TypeString,
							Description: "Layer 4 invalid protocol anomalies.",
							Computed:    true,
						},
						"ipv6_saddr_err": {
							Type:        schema.TypeString,
							Description: "Source address as multicast anomalies.",
							Computed:    true,
						},
						"ipv6_unknopt": {
							Type:        schema.TypeString,
							Description: "Unknown option anomalies.",
							Computed:    true,
						},
						"sctp_csum_err": {
							Type:        schema.TypeString,
							Description: "Invalid IPv4 SCTP checksum anomalies.",
							Computed:    true,
						},
						"tcp_csum_err": {
							Type:        schema.TypeString,
							Description: "Invalid IPv4 TCP checksum anomalies.",
							Computed:    true,
						},
						"tcp_fin_noack": {
							Type:        schema.TypeString,
							Description: "TCP SYN flood with FIN flag set without ACK setting anomalies.",
							Computed:    true,
						},
						"tcp_fin_only": {
							Type:        schema.TypeString,
							Description: "TCP SYN flood with only FIN flag set anomalies.",
							Computed:    true,
						},
						"tcp_land": {
							Type:        schema.TypeString,
							Description: "TCP land anomalies.",
							Computed:    true,
						},
						"tcp_no_flag": {
							Type:        schema.TypeString,
							Description: "TCP SYN flood with no flag set anomalies.",
							Computed:    true,
						},
						"tcp_syn_data": {
							Type:        schema.TypeString,
							Description: "TCP SYN flood packets with data anomalies.",
							Computed:    true,
						},
						"tcp_syn_fin": {
							Type:        schema.TypeString,
							Description: "TCP SYN flood SYN/FIN flag set anomalies. ",
							Computed:    true,
						},
						"tcp_winnuke": {
							Type:        schema.TypeString,
							Description: "TCP WinNuke anomalies.",
							Computed:    true,
						},
						"udp_csum_err": {
							Type:        schema.TypeString,
							Description: "Invalid IPv4 UDP checksum anomalies.",
							Computed:    true,
						},
						"udp_land": {
							Type:        schema.TypeString,
							Description: "UDP land anomalies.",
							Computed:    true,
						},
					},
				},
			},
			"gtp_support": {
				Type:        schema.TypeString,
				Description: "Enable/Disable NP7 GTP support",
				Computed:    true,
			},
			"hash_tbl_spread": {
				Type:        schema.TypeString,
				Description: "Enable/disable hash table entry spread (default enabled).",
				Computed:    true,
			},
			"hpe": {
				Type:        schema.TypeList,
				Description: "Host protection engine configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all_protocol": {
							Type:        schema.TypeInt,
							Description: "Maximum packet rate of each host queue except high priority traffic(1K - 32M pps, default = 400K pps), set 0 to disable.",
							Computed:    true,
						},
						"arp_max": {
							Type:        schema.TypeInt,
							Description: "Maximum ARP packet rate (1K - 32M pps, default = 5K pps). Entry is valid when ARP is removed from high-priority traffic.",
							Computed:    true,
						},
						"enable_shaper": {
							Type:        schema.TypeString,
							Description: "Enable/Disable NPU Host Protection Engine (HPE) for packet type shaper.",
							Computed:    true,
						},
						"esp_max": {
							Type:        schema.TypeInt,
							Description: "Maximum ESP packet rate (1K - 32M pps, default = 5K pps).",
							Computed:    true,
						},
						"high_priority": {
							Type:        schema.TypeInt,
							Description: "Maximum packet rate for high priority traffic packets (1K - 32M pps, default = 400K pps).",
							Computed:    true,
						},
						"icmp_max": {
							Type:        schema.TypeInt,
							Description: "Maximum ICMP packet rate (1K - 32M pps, default = 5K pps).",
							Computed:    true,
						},
						"ip_frag_max": {
							Type:        schema.TypeInt,
							Description: "Maximum fragmented IP packet rate (1K - 32M pps, default = 5K pps).",
							Computed:    true,
						},
						"ip_others_max": {
							Type:        schema.TypeInt,
							Description: "Maximum IP packet rate for other packets (packet types that cannot be set with other options) (1K - 32G pps, default = 5K pps).",
							Computed:    true,
						},
						"l2_others_max": {
							Type:        schema.TypeInt,
							Description: "Maximum L2 packet rate for L2 packets that are not ARP packets (1K - 32M pps, default = 5K pps).",
							Computed:    true,
						},
						"sctp_max": {
							Type:        schema.TypeInt,
							Description: "Maximum SCTP packet rate (1K - 32M pps, default = 5K pps).",
							Computed:    true,
						},
						"tcp_max": {
							Type:        schema.TypeInt,
							Description: "Maximum TCP packet rate (1K - 32M pps, default = 40K pps).",
							Computed:    true,
						},
						"tcpfin_rst_max": {
							Type:        schema.TypeInt,
							Description: "Maximum TCP carries FIN or RST flags packet rate (1K - 32M pps, default = 40K pps).",
							Computed:    true,
						},
						"tcpsyn_ack_max": {
							Type:        schema.TypeInt,
							Description: "Maximum TCP carries SYN and ACK flags packet rate (1K - 32M pps, default = 40K pps).",
							Computed:    true,
						},
						"tcpsyn_max": {
							Type:        schema.TypeInt,
							Description: "Maximum TCP SYN packet rate (1K - 40M pps, default = 32K pps).",
							Computed:    true,
						},
						"udp_max": {
							Type:        schema.TypeInt,
							Description: "Maximum UDP packet rate (1K - 32M pps, default = 40K pps).",
							Computed:    true,
						},
					},
				},
			},
			"htab_dedi_queue_nr": {
				Type:        schema.TypeInt,
				Description: "Set the number of dedicate queue for hash table messages.",
				Computed:    true,
			},
			"htab_msg_queue": {
				Type:        schema.TypeString,
				Description: "Set hash table message queue mode.",
				Computed:    true,
			},
			"htx_icmp_csum_chk": {
				Type:        schema.TypeString,
				Description: "Set HTX icmp csum checking mode.",
				Computed:    true,
			},
			"inbound_dscp_copy_port": {
				Type:        schema.TypeList,
				Description: "Physical interfaces that support inbound-dscp-copy.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface": {
							Type:        schema.TypeString,
							Description: "Physical interface name.",
							Computed:    true,
						},
					},
				},
			},
			"ip_fragment_offload": {
				Type:        schema.TypeString,
				Description: "Enable/disable NP7 NPU IP fragment offload.",
				Computed:    true,
			},
			"ip_reassembly": {
				Type:        schema.TypeList,
				Description: "IP reassebmly engine configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_timeout": {
							Type:        schema.TypeInt,
							Description: "Maximum timeout value for IP reassembly (5 us - 600,000,000 us).",
							Computed:    true,
						},
						"min_timeout": {
							Type:        schema.TypeInt,
							Description: "Minimum timeout value for IP reassembly (5 us - 600,000,000 us).",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Set IP reassembly processing status.",
							Computed:    true,
						},
					},
				},
			},
			"ippool_overload_high": {
				Type:        schema.TypeInt,
				Description: " High threshold for overload ippool port reuse (100%-2000%, default = 200).",
				Computed:    true,
			},
			"ippool_overload_low": {
				Type:        schema.TypeInt,
				Description: " Low threshold for overload ippool port reuse (100%-2000%, default = 150).",
				Computed:    true,
			},
			"ipsec_dec_subengine_mask": {
				Type:        schema.TypeString,
				Description: "IPsec decryption subengine mask (0x1 - 0xff, default 0xff).",
				Computed:    true,
			},
			"ipsec_enc_subengine_mask": {
				Type:        schema.TypeString,
				Description: "IPsec encryption subengine mask (0x1 - 0xff, default 0xff).",
				Computed:    true,
			},
			"ipsec_inbound_cache": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPsec inbound cache for anti-replay.",
				Computed:    true,
			},
			"ipsec_mtu_override": {
				Type:        schema.TypeString,
				Description: "Enable/disable NP6 IPsec MTU override.",
				Computed:    true,
			},
			"ipsec_ob_np_sel": {
				Type:        schema.TypeString,
				Description: "IPsec NP selection for OB SA offloading.",
				Computed:    true,
			},
			"ipsec_over_vlink": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPSEC over vlink.",
				Computed:    true,
			},
			"isf_np_queues": {
				Type:        schema.TypeList,
				Description: "Configure queues of switch port connected to NP6 XAUI on ingress path.",
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
			"max_session_timeout": {
				Type:        schema.TypeInt,
				Description: "Maximum time interval for refreshing NPU-offloaded sessions (10 - 1000 sec, default 40 sec).",
				Computed:    true,
			},
			"mcast_session_accounting": {
				Type:        schema.TypeString,
				Description: "Enable/disable traffic accounting for each multicast session through TAE counter.",
				Computed:    true,
			},
			"napi_break_interval": {
				Type:        schema.TypeInt,
				Description: " NAPI break interval (default 0).",
				Computed:    true,
			},
			"np_queues": {
				Type:        schema.TypeList,
				Description: "Configure queue assignment on NP7.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ethernet_type": {
							Type:        schema.TypeList,
							Description: "Configure a NP7 QoS Ethernet Type.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Ethernet Type Name.",
										Computed:    true,
									},
									"queue": {
										Type:        schema.TypeInt,
										Description: "Queue Number.",
										Computed:    true,
									},
									"type": {
										Type:        schema.TypeString,
										Description: "Ethernet Type.",
										Computed:    true,
									},
									"weight": {
										Type:        schema.TypeInt,
										Description: "Class Weight.",
										Computed:    true,
									},
								},
							},
						},
						"ip_protocol": {
							Type:        schema.TypeList,
							Description: "Configure a NP7 QoS IP Protocol.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "IP Protocol Name.",
										Computed:    true,
									},
									"protocol": {
										Type:        schema.TypeInt,
										Description: "IP Protocol.",
										Computed:    true,
									},
									"queue": {
										Type:        schema.TypeInt,
										Description: "Queue Number.",
										Computed:    true,
									},
									"weight": {
										Type:        schema.TypeInt,
										Description: "Class Weight.",
										Computed:    true,
									},
								},
							},
						},
						"ip_service": {
							Type:        schema.TypeList,
							Description: "Configure a NP7 QoS IP Service.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dport": {
										Type:        schema.TypeInt,
										Description: "Destination port.",
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "IP service name.",
										Computed:    true,
									},
									"protocol": {
										Type:        schema.TypeInt,
										Description: "IP protocol.",
										Computed:    true,
									},
									"queue": {
										Type:        schema.TypeInt,
										Description: "Queue number.",
										Computed:    true,
									},
									"sport": {
										Type:        schema.TypeInt,
										Description: "Source port.",
										Computed:    true,
									},
									"weight": {
										Type:        schema.TypeInt,
										Description: "Class weight.",
										Computed:    true,
									},
								},
							},
						},
						"profile": {
							Type:        schema.TypeList,
							Description: "Configure a NP7 class profile.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cos0": {
										Type:        schema.TypeString,
										Description: "Queue number of CoS 0.",
										Computed:    true,
									},
									"cos1": {
										Type:        schema.TypeString,
										Description: "Queue number of CoS 1.",
										Computed:    true,
									},
									"cos2": {
										Type:        schema.TypeString,
										Description: "Queue number of CoS 2.",
										Computed:    true,
									},
									"cos3": {
										Type:        schema.TypeString,
										Description: "Queue number of CoS 3.",
										Computed:    true,
									},
									"cos4": {
										Type:        schema.TypeString,
										Description: "Queue number of CoS 4.",
										Computed:    true,
									},
									"cos5": {
										Type:        schema.TypeString,
										Description: "Queue number of CoS 5.",
										Computed:    true,
									},
									"cos6": {
										Type:        schema.TypeString,
										Description: "Queue number of CoS 6.",
										Computed:    true,
									},
									"cos7": {
										Type:        schema.TypeString,
										Description: "Queue number of CoS 7.",
										Computed:    true,
									},
									"dscp0": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 0.",
										Computed:    true,
									},
									"dscp1": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 1.",
										Computed:    true,
									},
									"dscp10": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 10.",
										Computed:    true,
									},
									"dscp11": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 11.",
										Computed:    true,
									},
									"dscp12": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 12.",
										Computed:    true,
									},
									"dscp13": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 13.",
										Computed:    true,
									},
									"dscp14": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 14.",
										Computed:    true,
									},
									"dscp15": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 15.",
										Computed:    true,
									},
									"dscp16": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 16.",
										Computed:    true,
									},
									"dscp17": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 17.",
										Computed:    true,
									},
									"dscp18": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 18.",
										Computed:    true,
									},
									"dscp19": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 19.",
										Computed:    true,
									},
									"dscp2": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 2.",
										Computed:    true,
									},
									"dscp20": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 20.",
										Computed:    true,
									},
									"dscp21": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 21.",
										Computed:    true,
									},
									"dscp22": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 22.",
										Computed:    true,
									},
									"dscp23": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 23.",
										Computed:    true,
									},
									"dscp24": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 24.",
										Computed:    true,
									},
									"dscp25": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 25.",
										Computed:    true,
									},
									"dscp26": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 26.",
										Computed:    true,
									},
									"dscp27": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 27.",
										Computed:    true,
									},
									"dscp28": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 28.",
										Computed:    true,
									},
									"dscp29": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 29.",
										Computed:    true,
									},
									"dscp3": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 3.",
										Computed:    true,
									},
									"dscp30": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 30.",
										Computed:    true,
									},
									"dscp31": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 31.",
										Computed:    true,
									},
									"dscp32": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 32.",
										Computed:    true,
									},
									"dscp33": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 33.",
										Computed:    true,
									},
									"dscp34": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 34.",
										Computed:    true,
									},
									"dscp35": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 35.",
										Computed:    true,
									},
									"dscp36": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 36.",
										Computed:    true,
									},
									"dscp37": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 37.",
										Computed:    true,
									},
									"dscp38": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 38.",
										Computed:    true,
									},
									"dscp39": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 39.",
										Computed:    true,
									},
									"dscp4": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 4.",
										Computed:    true,
									},
									"dscp40": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 40.",
										Computed:    true,
									},
									"dscp41": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 41.",
										Computed:    true,
									},
									"dscp42": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 42.",
										Computed:    true,
									},
									"dscp43": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 43.",
										Computed:    true,
									},
									"dscp44": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 44.",
										Computed:    true,
									},
									"dscp45": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 45.",
										Computed:    true,
									},
									"dscp46": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 46.",
										Computed:    true,
									},
									"dscp47": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 47.",
										Computed:    true,
									},
									"dscp48": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 48.",
										Computed:    true,
									},
									"dscp49": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 49.",
										Computed:    true,
									},
									"dscp5": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 5.",
										Computed:    true,
									},
									"dscp50": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 50.",
										Computed:    true,
									},
									"dscp51": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 51.",
										Computed:    true,
									},
									"dscp52": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 52.",
										Computed:    true,
									},
									"dscp53": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 53.",
										Computed:    true,
									},
									"dscp54": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 54.",
										Computed:    true,
									},
									"dscp55": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 55.",
										Computed:    true,
									},
									"dscp56": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 56.",
										Computed:    true,
									},
									"dscp57": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 57.",
										Computed:    true,
									},
									"dscp58": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 58.",
										Computed:    true,
									},
									"dscp59": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 59.",
										Computed:    true,
									},
									"dscp6": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 6.",
										Computed:    true,
									},
									"dscp60": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 60.",
										Computed:    true,
									},
									"dscp61": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 61.",
										Computed:    true,
									},
									"dscp62": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 62.",
										Computed:    true,
									},
									"dscp63": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 63.",
										Computed:    true,
									},
									"dscp7": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 7.",
										Computed:    true,
									},
									"dscp8": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 8.",
										Computed:    true,
									},
									"dscp9": {
										Type:        schema.TypeString,
										Description: "Queue number of DSCP 9.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "Profile ID.",
										Computed:    true,
									},
									"type": {
										Type:        schema.TypeString,
										Description: "Profile type.",
										Computed:    true,
									},
									"weight": {
										Type:        schema.TypeInt,
										Description: "Class weight.",
										Computed:    true,
									},
								},
							},
						},
						"scheduler": {
							Type:        schema.TypeList,
							Description: "Configure a NP7 QoS Scheduler.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mode": {
										Type:        schema.TypeString,
										Description: "Scheduler mode.",
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "Scheduler name.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"np6_cps_optimization_mode": {
				Type:        schema.TypeString,
				Description: "Enable/disable NP6 connection per second (CPS) optimization mode.",
				Computed:    true,
			},
			"npu_group_effective_scope": {
				Type:        schema.TypeInt,
				Description: "npu-group-effective-scope defines under which npu-group cmds such as list/purge will be excecuted. Default scope is for all four HS-ok groups. (0-3, default = 255).",
				Computed:    true,
			},
			"pba_eim": {
				Type:        schema.TypeString,
				Description: "Configure option for PBA(non-overload)/EIM combination.",
				Computed:    true,
			},
			"per_policy_accounting": {
				Type:        schema.TypeString,
				Description: "Set per-policy accounting.",
				Computed:    true,
			},
			"per_session_accounting": {
				Type:        schema.TypeString,
				Description: "Set per-session accounting.",
				Computed:    true,
			},
			"policy_offload_level": {
				Type:        schema.TypeString,
				Description: "Configure firewall policy offload level.",
				Computed:    true,
			},
			"port_cpu_map": {
				Type:        schema.TypeList,
				Description: "Configure NPU interface to CPU core mapping.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cpu_core": {
							Type:        schema.TypeString,
							Description: "The CPU core to map to an interface.",
							Computed:    true,
						},
						"interface": {
							Type:        schema.TypeString,
							Description: "The interface to map to a CPU core.",
							Computed:    true,
						},
					},
				},
			},
			"port_npu_map": {
				Type:        schema.TypeList,
				Description: "Configure port to NPU group mapping.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface": {
							Type:        schema.TypeString,
							Description: "Set NPU interface port for NPU group mapping.",
							Computed:    true,
						},
						"npu_group_index": {
							Type:        schema.TypeInt,
							Description: "Mapping NPU group index.",
							Computed:    true,
						},
					},
				},
			},
			"priority_protocol": {
				Type:        schema.TypeList,
				Description: "Configure NPU priority protocol.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bfd": {
							Type:        schema.TypeString,
							Description: "Enable/disable NPU BFD priority protocol.",
							Computed:    true,
						},
						"bgp": {
							Type:        schema.TypeString,
							Description: "Enable/disable NPU BGP priority protocol.",
							Computed:    true,
						},
						"slbc": {
							Type:        schema.TypeString,
							Description: "Enable/disable NPU SLBC priority protocol.",
							Computed:    true,
						},
					},
				},
			},
			"qos_mode": {
				Type:        schema.TypeString,
				Description: "QoS mode on switch and NP.",
				Computed:    true,
			},
			"qtm_buf_mode": {
				Type:        schema.TypeString,
				Description: "QTM channel configuration for packet buffer.",
				Computed:    true,
			},
			"rdp_offload": {
				Type:        schema.TypeString,
				Description: "Enable/disable rdp offload.",
				Computed:    true,
			},
			"recover_np6_link": {
				Type:        schema.TypeString,
				Description: "Enable/disable internal link failure check and recovery after boot up.",
				Computed:    true,
			},
			"session_acct_interval": {
				Type:        schema.TypeInt,
				Description: "Session accounting update interval (1 - 10 sec, default 5 sec).",
				Computed:    true,
			},
			"shaping_stats": {
				Type:        schema.TypeString,
				Description: "Enable/disable NP7 traffic shaping statistics (default = disable).",
				Computed:    true,
			},
			"sse_backpressure": {
				Type:        schema.TypeString,
				Description: "Enable/disable sse backpressure.",
				Computed:    true,
			},
			"strip_clear_text_padding": {
				Type:        schema.TypeString,
				Description: "Enable/disable stripping clear text padding.",
				Computed:    true,
			},
			"strip_esp_padding": {
				Type:        schema.TypeString,
				Description: "Enable/disable stripping ESP padding.",
				Computed:    true,
			},
			"sw_np_bandwidth": {
				Type:        schema.TypeString,
				Description: "Bandwidth from switch to NP.",
				Computed:    true,
			},
			"tcp_rst_timeout": {
				Type:        schema.TypeInt,
				Description: " TCP RST timeout in seconds (0-3600, default = 5).",
				Computed:    true,
			},
			"tcp_timeout_profile": {
				Type:        schema.TypeList,
				Description: "Configure TCP timeout profile.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"close_wait": {
							Type:        schema.TypeInt,
							Description: "Set close-wait timeout(seconds)",
							Computed:    true,
						},
						"fin_wait": {
							Type:        schema.TypeInt,
							Description: "Set fin-wait timeout(seconds)",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Timeout profile ID (5-47)",
							Computed:    true,
						},
						"syn_sent": {
							Type:        schema.TypeInt,
							Description: "Set syn-sent timeout(seconds)",
							Computed:    true,
						},
						"syn_wait": {
							Type:        schema.TypeInt,
							Description: "Set syn-wait timeout(seconds)",
							Computed:    true,
						},
						"tcp_idle": {
							Type:        schema.TypeInt,
							Description: "Set TCP establish timeout(seconds)",
							Computed:    true,
						},
						"time_wait": {
							Type:        schema.TypeInt,
							Description: "Set time-wait timeout(seconds)",
							Computed:    true,
						},
					},
				},
			},
			"udp_timeout_profile": {
				Type:        schema.TypeList,
				Description: "Configure UDP timeout profile.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Timeout profile ID (5-63)",
							Computed:    true,
						},
						"udp_idle": {
							Type:        schema.TypeInt,
							Description: "Set UDP idle timeout(seconds)",
							Computed:    true,
						},
					},
				},
			},
			"vlan_lookup_cache": {
				Type:        schema.TypeString,
				Description: "Enable/disable vlan lookup cache (default enabled).",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemNpuRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemNpu"

	o, err := c.Cmdb.ReadSystemNpu(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemNpu dataSource: %v", err)
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

	diags := refreshObjectSystemNpu(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
