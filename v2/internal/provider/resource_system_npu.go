// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.6,v7.0.2,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure NPU attributes.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceSystemNpu() *schema.Resource {
	return &schema.Resource{
		Description: "Configure NPU attributes.",

		CreateContext: resourceSystemNpuCreate,
		ReadContext:   resourceSystemNpuRead,
		UpdateContext: resourceSystemNpuUpdate,
		DeleteContext: resourceSystemNpuDelete,

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
			"capwap_offload": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable offloading managed FortiAP and FortiLink CAPWAP sessions.",
				Optional:    true,
				Computed:    true,
			},
			"dedicated_management_affinity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Affinity setting for management daemons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).",
				Optional:    true,
				Computed:    true,
			},
			"dedicated_management_cpu": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to dedicate one CPU for GUI and CLI connections when NPs are busy.",
				Optional:    true,
				Computed:    true,
			},
			"default_qos_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"policing", "shaping", "policing-enhanced"}, false),

				Description: "Set default QoS type.",
				Optional:    true,
				Computed:    true,
			},
			"dos_options": {
				Type:        schema.TypeList,
				Description: "NPU DoS configurations.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"npu_dos_meter_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"global", "local"}, false),

							Description: "Set DoS meter NPU offloading mode.",
							Optional:    true,
							Computed:    true,
						},
						"npu_dos_tpe_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable insertion of DoS meter ID to session table.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"double_level_mcast_offload": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable double level mcast offload.",
				Optional:    true,
				Computed:    true,
			},
			"dse_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3600),

				Description: " DSE timeout in seconds (0-3600, default = 10).",
				Optional:    true,
				Computed:    true,
			},
			"dsw_dts_profile": {
				Type:        schema.TypeList,
				Description: "Configure NPU DSW DTS profile.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"wait", "drop", "drop_tmr_0", "drop_tmr_1", "enque", "enque_0", "enque_1"}, false),

							Description: "Set NPU DSW DTS profile action.",
							Optional:    true,
							Computed:    true,
						},
						"min_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(32, 2048),

							Description: "Set NPU DSW DTS profile min-limt.",
							Optional:    true,
							Computed:    true,
						},
						"profile_id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 32),

							Description: "Set NPU DSW DTS profile profile id.",
							Optional:    true,
							Computed:    true,
						},
						"step": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 64),

							Description: "Set NPU DSW DTS profile step.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dsw_queue_dts_profile": {
				Type:        schema.TypeList,
				Description: "Configure NPU DSW Queue DTS profile.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iport": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"eif0", "eif1", "eif2", "eif3", "eif4", "eif5", "eif6", "eif7", "htx0", "htx1", "sse0", "sse1", "sse2", "sse3", "rlt", "dfr", "ipseci", "ipseco", "ipti", "ipto", "vep0", "vep2", "vep4", "vep6", "ivs", "l2ti1", "l2to", "l2ti0", "ple", "spath", "qtm"}, false),

							Description: "Set NPU DSW DTS in port.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Name.",
							Optional:    true,
							Computed:    true,
						},
						"oport": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"eif0", "eif1", "eif2", "eif3", "eif4", "eif5", "eif6", "eif7", "hrx", "sse0", "sse1", "sse2", "sse3", "rlt", "dfr", "ipseci", "ipseco", "ipti", "ipto", "vep0", "vep2", "vep4", "vep6", "ivs", "l2ti1", "l2to", "l2ti0", "ple", "sync", "nss", "tsk", "qtm"}, false),

							Description: "Set NPU DSW DTS out port.",
							Optional:    true,
							Computed:    true,
						},
						"profile_id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 32),

							Description: "Set NPU DSW DTS profile ID.",
							Optional:    true,
							Computed:    true,
						},
						"queue_select": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4095),

							Description: "Set NPU DSW DTS queue ID select (0 - reset to default).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"fastpath": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable NP6 offloading (also called fast path).",
				Optional:    true,
				Computed:    true,
			},
			"fp_anomaly": {
				Type:        schema.TypeList,
				Description: "IPv4/IPv6 anomaly protection.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icmp_csum_err": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"drop", "trap-to-host"}, false),

							Description: "Invalid IPv4 ICMP checksum anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"icmp_frag": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "Layer 3 fragmented packets that could be part of layer 4 ICMP anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"icmp_land": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "ICMP land anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"ipv4_csum_err": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"drop", "trap-to-host"}, false),

							Description: "Invalid IPv4 IP checksum anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"ipv4_land": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "Land anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"ipv4_optlsrr": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "Loose source record route option anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"ipv4_optrr": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "Record route option anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"ipv4_optsecurity": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "Security option anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"ipv4_optssrr": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "Strict source record route option anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"ipv4_optstream": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "Stream option anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"ipv4_opttimestamp": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "Timestamp option anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"ipv4_proto_err": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "Invalid layer 4 protocol anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"ipv4_unknopt": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "Unknown option anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"ipv6_daddr_err": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "Destination address as unspecified or loopback address anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"ipv6_land": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "Land anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"ipv6_optendpid": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "End point identification anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"ipv6_opthomeaddr": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "Home address option anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"ipv6_optinvld": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "Invalid option anomalies.Invalid option anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"ipv6_optjumbo": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "Jumbo options anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"ipv6_optnsap": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "Network service access point address option anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"ipv6_optralert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "Router alert option anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"ipv6_opttunnel": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "Tunnel encapsulation limit option anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"ipv6_proto_err": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "Layer 4 invalid protocol anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"ipv6_saddr_err": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "Source address as multicast anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"ipv6_unknopt": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "Unknown option anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"sctp_csum_err": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "Invalid IPv4 SCTP checksum anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"tcp_csum_err": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"drop", "trap-to-host"}, false),

							Description: "Invalid IPv4 TCP checksum anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"tcp_fin_noack": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "TCP SYN flood with FIN flag set without ACK setting anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"tcp_fin_only": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "TCP SYN flood with only FIN flag set anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"tcp_land": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "TCP land anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"tcp_no_flag": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "TCP SYN flood with no flag set anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"tcp_syn_data": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "TCP SYN flood packets with data anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"tcp_syn_fin": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "TCP SYN flood SYN/FIN flag set anomalies. ",
							Optional:    true,
							Computed:    true,
						},
						"tcp_winnuke": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "TCP WinNuke anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"udp_csum_err": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"drop", "trap-to-host"}, false),

							Description: "Invalid IPv4 UDP checksum anomalies.",
							Optional:    true,
							Computed:    true,
						},
						"udp_land": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "drop", "trap-to-host"}, false),

							Description: "UDP land anomalies.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"gtp_support": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/Disable NP7 GTP support",
				Optional:    true,
				Computed:    true,
			},
			"hash_tbl_spread": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable hash table entry spread (default enabled).",
				Optional:    true,
				Computed:    true,
			},
			"hpe": {
				Type:        schema.TypeList,
				Description: "Host protection engine configuration.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all_protocol": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 32000000),

							Description: "Maximum packet rate of each host queue except high priority traffic(1K - 32M pps, default = 400K pps), set 0 to disable.",
							Optional:    true,
							Computed:    true,
						},
						"arp_max": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),

							Description: "Maximum ARP packet rate (1K - 32M pps, default = 5K pps). Entry is valid when ARP is removed from high-priority traffic.",
							Optional:    true,
							Computed:    true,
						},
						"enable_shaper": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/Disable NPU Host Protection Engine (HPE) for packet type shaper.",
							Optional:    true,
							Computed:    true,
						},
						"esp_max": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),

							Description: "Maximum ESP packet rate (1K - 32M pps, default = 5K pps).",
							Optional:    true,
							Computed:    true,
						},
						"high_priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),

							Description: "Maximum packet rate for high priority traffic packets (1K - 32M pps, default = 400K pps).",
							Optional:    true,
							Computed:    true,
						},
						"icmp_max": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),

							Description: "Maximum ICMP packet rate (1K - 32M pps, default = 5K pps).",
							Optional:    true,
							Computed:    true,
						},
						"ip_frag_max": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),

							Description: "Maximum fragmented IP packet rate (1K - 32M pps, default = 5K pps).",
							Optional:    true,
							Computed:    true,
						},
						"ip_others_max": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),

							Description: "Maximum IP packet rate for other packets (packet types that cannot be set with other options) (1K - 32G pps, default = 5K pps).",
							Optional:    true,
							Computed:    true,
						},
						"l2_others_max": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),

							Description: "Maximum L2 packet rate for L2 packets that are not ARP packets (1K - 32M pps, default = 5K pps).",
							Optional:    true,
							Computed:    true,
						},
						"sctp_max": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),

							Description: "Maximum SCTP packet rate (1K - 32M pps, default = 5K pps).",
							Optional:    true,
							Computed:    true,
						},
						"tcp_max": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),

							Description: "Maximum TCP packet rate (1K - 32M pps, default = 40K pps).",
							Optional:    true,
							Computed:    true,
						},
						"tcpfin_rst_max": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),

							Description: "Maximum TCP carries FIN or RST flags packet rate (1K - 32M pps, default = 40K pps).",
							Optional:    true,
							Computed:    true,
						},
						"tcpsyn_ack_max": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),

							Description: "Maximum TCP carries SYN and ACK flags packet rate (1K - 32M pps, default = 40K pps).",
							Optional:    true,
							Computed:    true,
						},
						"tcpsyn_max": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),

							Description: "Maximum TCP SYN packet rate (1K - 40M pps, default = 32K pps).",
							Optional:    true,
							Computed:    true,
						},
						"udp_max": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),

							Description: "Maximum UDP packet rate (1K - 32M pps, default = 40K pps).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"htab_dedi_queue_nr": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 2),

				Description: "Set the number of dedicate queue for hash table messages.",
				Optional:    true,
				Computed:    true,
			},
			"htab_msg_queue": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"data", "idle", "dedicated"}, false),

				Description: "Set hash table message queue mode.",
				Optional:    true,
				Computed:    true,
			},
			"htx_icmp_csum_chk": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"drop", "pass"}, false),

				Description: "Set HTX icmp csum checking mode.",
				Optional:    true,
				Computed:    true,
			},
			"inbound_dscp_copy_port": {
				Type:        schema.TypeList,
				Description: "Physical interfaces that support inbound-dscp-copy.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Physical interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ip_fragment_offload": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable NP7 NPU IP fragment offload.",
				Optional:    true,
				Computed:    true,
			},
			"ip_reassembly": {
				Type:        schema.TypeList,
				Description: "IP reassebmly engine configuration.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_timeout": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 600000000),

							Description: "Maximum timeout value for IP reassembly (5 us - 600,000,000 us).",
							Optional:    true,
							Computed:    true,
						},
						"min_timeout": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 600000000),

							Description: "Minimum timeout value for IP reassembly (5 us - 600,000,000 us).",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Set IP reassembly processing status.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ippool_overload_high": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(100, 2000),

				Description: " High threshold for overload ippool port reuse (100%-2000%, default = 200).",
				Optional:    true,
				Computed:    true,
			},
			"ippool_overload_low": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(100, 2000),

				Description: " Low threshold for overload ippool port reuse (100%-2000%, default = 150).",
				Optional:    true,
				Computed:    true,
			},
			"ipsec_dec_subengine_mask": {
				Type: schema.TypeString,

				Description: "IPsec decryption subengine mask (0x1 - 0xff, default 0xff).",
				Optional:    true,
				Computed:    true,
			},
			"ipsec_enc_subengine_mask": {
				Type: schema.TypeString,

				Description: "IPsec encryption subengine mask (0x1 - 0xff, default 0xff).",
				Optional:    true,
				Computed:    true,
			},
			"ipsec_inbound_cache": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPsec inbound cache for anti-replay.",
				Optional:    true,
				Computed:    true,
			},
			"ipsec_mtu_override": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable NP6 IPsec MTU override.",
				Optional:    true,
				Computed:    true,
			},
			"ipsec_ob_np_sel": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"rr", "Packet", "Hash"}, false),

				Description: "IPsec NP selection for OB SA offloading.",
				Optional:    true,
				Computed:    true,
			},
			"ipsec_over_vlink": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPSEC over vlink.",
				Optional:    true,
				Computed:    true,
			},
			"isf_np_queues": {
				Type:        schema.TypeList,
				Description: "Configure queues of switch port connected to NP6 XAUI on ingress path.",
				Optional:    true, MaxItems: 1,
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
			"max_session_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 1000),

				Description: "Maximum time interval for refreshing NPU-offloaded sessions (10 - 1000 sec, default 40 sec).",
				Optional:    true,
				Computed:    true,
			},
			"mcast_session_accounting": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"tpe-based", "session-based", "disable"}, false),

				Description: "Enable/disable traffic accounting for each multicast session through TAE counter.",
				Optional:    true,
				Computed:    true,
			},
			"napi_break_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: " NAPI break interval (default 0).",
				Optional:    true,
				Computed:    true,
			},
			"np_queues": {
				Type:        schema.TypeList,
				Description: "Configure queue assignment on NP7.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ethernet_type": {
							Type:        schema.TypeList,
							Description: "Configure a NP7 QoS Ethernet Type.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Ethernet Type Name.",
										Optional:    true,
										Computed:    true,
									},
									"queue": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 11),

										Description: "Queue Number.",
										Optional:    true,
										Computed:    true,
									},
									"type": {
										Type: schema.TypeString,

										Description: "Ethernet Type.",
										Optional:    true,
										Computed:    true,
									},
									"weight": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 15),

										Description: "Class Weight.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"ip_protocol": {
							Type:        schema.TypeList,
							Description: "Configure a NP7 QoS IP Protocol.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "IP Protocol Name.",
										Optional:    true,
										Computed:    true,
									},
									"protocol": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 255),

										Description: "IP Protocol.",
										Optional:    true,
										Computed:    true,
									},
									"queue": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 11),

										Description: "Queue Number.",
										Optional:    true,
										Computed:    true,
									},
									"weight": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 15),

										Description: "Class Weight.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"ip_service": {
							Type:        schema.TypeList,
							Description: "Configure a NP7 QoS IP Service.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dport": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),

										Description: "Destination port.",
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "IP service name.",
										Optional:    true,
										Computed:    true,
									},
									"protocol": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 255),

										Description: "IP protocol.",
										Optional:    true,
										Computed:    true,
									},
									"queue": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 11),

										Description: "Queue number.",
										Optional:    true,
										Computed:    true,
									},
									"sport": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),

										Description: "Source port.",
										Optional:    true,
										Computed:    true,
									},
									"weight": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 15),

										Description: "Class weight.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"profile": {
							Type:        schema.TypeList,
							Description: "Configure a NP7 class profile.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cos0": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of CoS 0.",
										Optional:    true,
										Computed:    true,
									},
									"cos1": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of CoS 1.",
										Optional:    true,
										Computed:    true,
									},
									"cos2": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of CoS 2.",
										Optional:    true,
										Computed:    true,
									},
									"cos3": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of CoS 3.",
										Optional:    true,
										Computed:    true,
									},
									"cos4": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of CoS 4.",
										Optional:    true,
										Computed:    true,
									},
									"cos5": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of CoS 5.",
										Optional:    true,
										Computed:    true,
									},
									"cos6": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of CoS 6.",
										Optional:    true,
										Computed:    true,
									},
									"cos7": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of CoS 7.",
										Optional:    true,
										Computed:    true,
									},
									"dscp0": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 0.",
										Optional:    true,
										Computed:    true,
									},
									"dscp1": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 1.",
										Optional:    true,
										Computed:    true,
									},
									"dscp10": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 10.",
										Optional:    true,
										Computed:    true,
									},
									"dscp11": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 11.",
										Optional:    true,
										Computed:    true,
									},
									"dscp12": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 12.",
										Optional:    true,
										Computed:    true,
									},
									"dscp13": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 13.",
										Optional:    true,
										Computed:    true,
									},
									"dscp14": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 14.",
										Optional:    true,
										Computed:    true,
									},
									"dscp15": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 15.",
										Optional:    true,
										Computed:    true,
									},
									"dscp16": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 16.",
										Optional:    true,
										Computed:    true,
									},
									"dscp17": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 17.",
										Optional:    true,
										Computed:    true,
									},
									"dscp18": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 18.",
										Optional:    true,
										Computed:    true,
									},
									"dscp19": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 19.",
										Optional:    true,
										Computed:    true,
									},
									"dscp2": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 2.",
										Optional:    true,
										Computed:    true,
									},
									"dscp20": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 20.",
										Optional:    true,
										Computed:    true,
									},
									"dscp21": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 21.",
										Optional:    true,
										Computed:    true,
									},
									"dscp22": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 22.",
										Optional:    true,
										Computed:    true,
									},
									"dscp23": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 23.",
										Optional:    true,
										Computed:    true,
									},
									"dscp24": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 24.",
										Optional:    true,
										Computed:    true,
									},
									"dscp25": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 25.",
										Optional:    true,
										Computed:    true,
									},
									"dscp26": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 26.",
										Optional:    true,
										Computed:    true,
									},
									"dscp27": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 27.",
										Optional:    true,
										Computed:    true,
									},
									"dscp28": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 28.",
										Optional:    true,
										Computed:    true,
									},
									"dscp29": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 29.",
										Optional:    true,
										Computed:    true,
									},
									"dscp3": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 3.",
										Optional:    true,
										Computed:    true,
									},
									"dscp30": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 30.",
										Optional:    true,
										Computed:    true,
									},
									"dscp31": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 31.",
										Optional:    true,
										Computed:    true,
									},
									"dscp32": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 32.",
										Optional:    true,
										Computed:    true,
									},
									"dscp33": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 33.",
										Optional:    true,
										Computed:    true,
									},
									"dscp34": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 34.",
										Optional:    true,
										Computed:    true,
									},
									"dscp35": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 35.",
										Optional:    true,
										Computed:    true,
									},
									"dscp36": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 36.",
										Optional:    true,
										Computed:    true,
									},
									"dscp37": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 37.",
										Optional:    true,
										Computed:    true,
									},
									"dscp38": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 38.",
										Optional:    true,
										Computed:    true,
									},
									"dscp39": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 39.",
										Optional:    true,
										Computed:    true,
									},
									"dscp4": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 4.",
										Optional:    true,
										Computed:    true,
									},
									"dscp40": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 40.",
										Optional:    true,
										Computed:    true,
									},
									"dscp41": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 41.",
										Optional:    true,
										Computed:    true,
									},
									"dscp42": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 42.",
										Optional:    true,
										Computed:    true,
									},
									"dscp43": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 43.",
										Optional:    true,
										Computed:    true,
									},
									"dscp44": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 44.",
										Optional:    true,
										Computed:    true,
									},
									"dscp45": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 45.",
										Optional:    true,
										Computed:    true,
									},
									"dscp46": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 46.",
										Optional:    true,
										Computed:    true,
									},
									"dscp47": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 47.",
										Optional:    true,
										Computed:    true,
									},
									"dscp48": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 48.",
										Optional:    true,
										Computed:    true,
									},
									"dscp49": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 49.",
										Optional:    true,
										Computed:    true,
									},
									"dscp5": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 5.",
										Optional:    true,
										Computed:    true,
									},
									"dscp50": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 50.",
										Optional:    true,
										Computed:    true,
									},
									"dscp51": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 51.",
										Optional:    true,
										Computed:    true,
									},
									"dscp52": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 52.",
										Optional:    true,
										Computed:    true,
									},
									"dscp53": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 53.",
										Optional:    true,
										Computed:    true,
									},
									"dscp54": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 54.",
										Optional:    true,
										Computed:    true,
									},
									"dscp55": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 55.",
										Optional:    true,
										Computed:    true,
									},
									"dscp56": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 56.",
										Optional:    true,
										Computed:    true,
									},
									"dscp57": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 57.",
										Optional:    true,
										Computed:    true,
									},
									"dscp58": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 58.",
										Optional:    true,
										Computed:    true,
									},
									"dscp59": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 59.",
										Optional:    true,
										Computed:    true,
									},
									"dscp6": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 6.",
										Optional:    true,
										Computed:    true,
									},
									"dscp60": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 60.",
										Optional:    true,
										Computed:    true,
									},
									"dscp61": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 61.",
										Optional:    true,
										Computed:    true,
									},
									"dscp62": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 62.",
										Optional:    true,
										Computed:    true,
									},
									"dscp63": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 63.",
										Optional:    true,
										Computed:    true,
									},
									"dscp7": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 7.",
										Optional:    true,
										Computed:    true,
									},
									"dscp8": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 8.",
										Optional:    true,
										Computed:    true,
									},
									"dscp9": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"queue0", "queue1", "queue2", "queue3", "queue4", "queue5", "queue6", "queue7"}, false),

										Description: "Queue number of DSCP 9.",
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 255),

										Description: "Profile ID.",
										Optional:    true,
										Computed:    true,
									},
									"type": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"cos", "dscp"}, false),

										Description: "Profile type.",
										Optional:    true,
										Computed:    true,
									},
									"weight": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 15),

										Description: "Class weight.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"scheduler": {
							Type:        schema.TypeList,
							Description: "Configure a NP7 QoS Scheduler.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mode": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"none", "priority", "round-robin"}, false),

										Description: "Scheduler mode.",
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Scheduler name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"np6_cps_optimization_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable NP6 connection per second (CPS) optimization mode.",
				Optional:    true,
				Computed:    true,
			},
			"npu_group_effective_scope": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "npu-group-effective-scope defines under which npu-group cmds such as list/purge will be excecuted. Default scope is for all four HS-ok groups. (0-3, default = 255).",
				Optional:    true,
				Computed:    true,
			},
			"pba_eim": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disallow", "allow"}, false),

				Description: "Configure option for PBA(non-overload)/EIM combination.",
				Optional:    true,
				Computed:    true,
			},
			"per_policy_accounting": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Set per-policy accounting.",
				Optional:    true,
				Computed:    true,
			},
			"per_session_accounting": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"traffic-log-only", "disable", "enable"}, false),

				Description: "Set per-session accounting.",
				Optional:    true,
				Computed:    true,
			},
			"policy_offload_level": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "dos-offload"}, false),

				Description: "Configure firewall policy offload level.",
				Optional:    true,
				Computed:    true,
			},
			"port_cpu_map": {
				Type:        schema.TypeList,
				Description: "Configure NPU interface to CPU core mapping.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cpu_core": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),

							Description: "The CPU core to map to an interface.",
							Optional:    true,
							Computed:    true,
						},
						"interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "The interface to map to a CPU core.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"port_npu_map": {
				Type:        schema.TypeList,
				Description: "Configure port to NPU group mapping.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Set NPU interface port for NPU group mapping.",
							Optional:    true,
							Computed:    true,
						},
						"npu_group_index": {
							Type: schema.TypeInt,

							Description: "Mapping NPU group index.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"priority_protocol": {
				Type:        schema.TypeList,
				Description: "Configure NPU priority protocol.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bfd": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable NPU BFD priority protocol.",
							Optional:    true,
							Computed:    true,
						},
						"bgp": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable NPU BGP priority protocol.",
							Optional:    true,
							Computed:    true,
						},
						"slbc": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable NPU SLBC priority protocol.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"qos_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "priority", "round-robin"}, false),

				Description: "QoS mode on switch and NP.",
				Optional:    true,
				Computed:    true,
			},
			"qtm_buf_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"6ch", "4ch"}, false),

				Description: "QTM channel configuration for packet buffer.",
				Optional:    true,
				Computed:    true,
			},
			"rdp_offload": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable rdp offload.",
				Optional:    true,
				Computed:    true,
			},
			"recover_np6_link": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable internal link failure check and recovery after boot up.",
				Optional:    true,
				Computed:    true,
			},
			"session_acct_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),

				Description: "Session accounting update interval (1 - 10 sec, default 5 sec).",
				Optional:    true,
				Computed:    true,
			},
			"shaping_stats": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable NP7 traffic shaping statistics (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"sse_backpressure": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable sse backpressure.",
				Optional:    true,
				Computed:    true,
			},
			"strip_clear_text_padding": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable stripping clear text padding.",
				Optional:    true,
				Computed:    true,
			},
			"strip_esp_padding": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable stripping ESP padding.",
				Optional:    true,
				Computed:    true,
			},
			"sw_np_bandwidth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"0G", "2G", "4G", "5G", "6G"}, false),

				Description: "Bandwidth from switch to NP.",
				Optional:    true,
				Computed:    true,
			},
			"tcp_rst_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16777215),

				Description: " TCP RST timeout in seconds (0-3600, default = 5).",
				Optional:    true,
				Computed:    true,
			},
			"tcp_timeout_profile": {
				Type:        schema.TypeList,
				Description: "Configure TCP timeout profile.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"close_wait": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 86400),

							Description: "Set close-wait timeout(seconds)",
							Optional:    true,
							Computed:    true,
						},
						"fin_wait": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 86400),

							Description: "Set fin-wait timeout(seconds)",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 47),

							Description: "Timeout profile ID (5-47)",
							Optional:    true,
							Computed:    true,
						},
						"syn_sent": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 86400),

							Description: "Set syn-sent timeout(seconds)",
							Optional:    true,
							Computed:    true,
						},
						"syn_wait": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 86400),

							Description: "Set syn-wait timeout(seconds)",
							Optional:    true,
							Computed:    true,
						},
						"tcp_idle": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 86400),

							Description: "Set TCP establish timeout(seconds)",
							Optional:    true,
							Computed:    true,
						},
						"time_wait": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 300),

							Description: "Set time-wait timeout(seconds)",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"udp_timeout_profile": {
				Type:        schema.TypeList,
				Description: "Configure UDP timeout profile.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 63),

							Description: "Timeout profile ID (5-63)",
							Optional:    true,
							Computed:    true,
						},
						"udp_idle": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 86400),

							Description: "Set UDP idle timeout(seconds)",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"vlan_lookup_cache": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable vlan lookup cache (default enabled).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemNpuCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemNpu(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemNpu(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemNpu")
	}

	return resourceSystemNpuRead(ctx, d, meta)
}

func resourceSystemNpuUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemNpu(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemNpu(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemNpu resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemNpu")
	}

	return resourceSystemNpuRead(ctx, d, meta)
}

func resourceSystemNpuDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemNpu(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemNpu(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemNpu resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNpuRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemNpu(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemNpu resource: %v", err)
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

	diags := refreshObjectSystemNpu(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemNpuDosOptions(d *schema.ResourceData, v *models.SystemNpuDosOptions, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemNpuDosOptions{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.NpuDosMeterMode; tmp != nil {
				v["npu_dos_meter_mode"] = *tmp
			}

			if tmp := cfg.NpuDosTpeMode; tmp != nil {
				v["npu_dos_tpe_mode"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemNpuDswDtsProfile(d *schema.ResourceData, v *[]models.SystemNpuDswDtsProfile, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.MinLimit; tmp != nil {
				v["min_limit"] = *tmp
			}

			if tmp := cfg.ProfileId; tmp != nil {
				v["profile_id"] = *tmp
			}

			if tmp := cfg.Step; tmp != nil {
				v["step"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "profile_id")
	}

	return flat
}

func flattenSystemNpuDswQueueDtsProfile(d *schema.ResourceData, v *[]models.SystemNpuDswQueueDtsProfile, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Iport; tmp != nil {
				v["iport"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Oport; tmp != nil {
				v["oport"] = *tmp
			}

			if tmp := cfg.ProfileId; tmp != nil {
				v["profile_id"] = *tmp
			}

			if tmp := cfg.QueueSelect; tmp != nil {
				v["queue_select"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemNpuFpAnomaly(d *schema.ResourceData, v *models.SystemNpuFpAnomaly, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemNpuFpAnomaly{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.IcmpCsumErr; tmp != nil {
				v["icmp_csum_err"] = *tmp
			}

			if tmp := cfg.IcmpFrag; tmp != nil {
				v["icmp_frag"] = *tmp
			}

			if tmp := cfg.IcmpLand; tmp != nil {
				v["icmp_land"] = *tmp
			}

			if tmp := cfg.Ipv4CsumErr; tmp != nil {
				v["ipv4_csum_err"] = *tmp
			}

			if tmp := cfg.Ipv4Land; tmp != nil {
				v["ipv4_land"] = *tmp
			}

			if tmp := cfg.Ipv4Optlsrr; tmp != nil {
				v["ipv4_optlsrr"] = *tmp
			}

			if tmp := cfg.Ipv4Optrr; tmp != nil {
				v["ipv4_optrr"] = *tmp
			}

			if tmp := cfg.Ipv4Optsecurity; tmp != nil {
				v["ipv4_optsecurity"] = *tmp
			}

			if tmp := cfg.Ipv4Optssrr; tmp != nil {
				v["ipv4_optssrr"] = *tmp
			}

			if tmp := cfg.Ipv4Optstream; tmp != nil {
				v["ipv4_optstream"] = *tmp
			}

			if tmp := cfg.Ipv4Opttimestamp; tmp != nil {
				v["ipv4_opttimestamp"] = *tmp
			}

			if tmp := cfg.Ipv4ProtoErr; tmp != nil {
				v["ipv4_proto_err"] = *tmp
			}

			if tmp := cfg.Ipv4Unknopt; tmp != nil {
				v["ipv4_unknopt"] = *tmp
			}

			if tmp := cfg.Ipv6DaddrErr; tmp != nil {
				v["ipv6_daddr_err"] = *tmp
			}

			if tmp := cfg.Ipv6Land; tmp != nil {
				v["ipv6_land"] = *tmp
			}

			if tmp := cfg.Ipv6Optendpid; tmp != nil {
				v["ipv6_optendpid"] = *tmp
			}

			if tmp := cfg.Ipv6Opthomeaddr; tmp != nil {
				v["ipv6_opthomeaddr"] = *tmp
			}

			if tmp := cfg.Ipv6Optinvld; tmp != nil {
				v["ipv6_optinvld"] = *tmp
			}

			if tmp := cfg.Ipv6Optjumbo; tmp != nil {
				v["ipv6_optjumbo"] = *tmp
			}

			if tmp := cfg.Ipv6Optnsap; tmp != nil {
				v["ipv6_optnsap"] = *tmp
			}

			if tmp := cfg.Ipv6Optralert; tmp != nil {
				v["ipv6_optralert"] = *tmp
			}

			if tmp := cfg.Ipv6Opttunnel; tmp != nil {
				v["ipv6_opttunnel"] = *tmp
			}

			if tmp := cfg.Ipv6ProtoErr; tmp != nil {
				v["ipv6_proto_err"] = *tmp
			}

			if tmp := cfg.Ipv6SaddrErr; tmp != nil {
				v["ipv6_saddr_err"] = *tmp
			}

			if tmp := cfg.Ipv6Unknopt; tmp != nil {
				v["ipv6_unknopt"] = *tmp
			}

			if tmp := cfg.SctpCsumErr; tmp != nil {
				v["sctp_csum_err"] = *tmp
			}

			if tmp := cfg.TcpCsumErr; tmp != nil {
				v["tcp_csum_err"] = *tmp
			}

			if tmp := cfg.TcpFinNoack; tmp != nil {
				v["tcp_fin_noack"] = *tmp
			}

			if tmp := cfg.TcpFinOnly; tmp != nil {
				v["tcp_fin_only"] = *tmp
			}

			if tmp := cfg.TcpLand; tmp != nil {
				v["tcp_land"] = *tmp
			}

			if tmp := cfg.TcpNoFlag; tmp != nil {
				v["tcp_no_flag"] = *tmp
			}

			if tmp := cfg.TcpSynData; tmp != nil {
				v["tcp_syn_data"] = *tmp
			}

			if tmp := cfg.TcpSynFin; tmp != nil {
				v["tcp_syn_fin"] = *tmp
			}

			if tmp := cfg.TcpWinnuke; tmp != nil {
				v["tcp_winnuke"] = *tmp
			}

			if tmp := cfg.UdpCsumErr; tmp != nil {
				v["udp_csum_err"] = *tmp
			}

			if tmp := cfg.UdpLand; tmp != nil {
				v["udp_land"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemNpuHpe(d *schema.ResourceData, v *models.SystemNpuHpe, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemNpuHpe{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AllProtocol; tmp != nil {
				v["all_protocol"] = *tmp
			}

			if tmp := cfg.ArpMax; tmp != nil {
				v["arp_max"] = *tmp
			}

			if tmp := cfg.EnableShaper; tmp != nil {
				v["enable_shaper"] = *tmp
			}

			if tmp := cfg.EspMax; tmp != nil {
				v["esp_max"] = *tmp
			}

			if tmp := cfg.HighPriority; tmp != nil {
				v["high_priority"] = *tmp
			}

			if tmp := cfg.IcmpMax; tmp != nil {
				v["icmp_max"] = *tmp
			}

			if tmp := cfg.IpFragMax; tmp != nil {
				v["ip_frag_max"] = *tmp
			}

			if tmp := cfg.IpOthersMax; tmp != nil {
				v["ip_others_max"] = *tmp
			}

			if tmp := cfg.L2OthersMax; tmp != nil {
				v["l2_others_max"] = *tmp
			}

			if tmp := cfg.SctpMax; tmp != nil {
				v["sctp_max"] = *tmp
			}

			if tmp := cfg.TcpMax; tmp != nil {
				v["tcp_max"] = *tmp
			}

			if tmp := cfg.TcpfinRstMax; tmp != nil {
				v["tcpfin_rst_max"] = *tmp
			}

			if tmp := cfg.TcpsynAckMax; tmp != nil {
				v["tcpsyn_ack_max"] = *tmp
			}

			if tmp := cfg.TcpsynMax; tmp != nil {
				v["tcpsyn_max"] = *tmp
			}

			if tmp := cfg.UdpMax; tmp != nil {
				v["udp_max"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemNpuInboundDscpCopyPort(d *schema.ResourceData, v *[]models.SystemNpuInboundDscpCopyPort, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "interface")
	}

	return flat
}

func flattenSystemNpuIpReassembly(d *schema.ResourceData, v *models.SystemNpuIpReassembly, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemNpuIpReassembly{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.MaxTimeout; tmp != nil {
				v["max_timeout"] = *tmp
			}

			if tmp := cfg.MinTimeout; tmp != nil {
				v["min_timeout"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemNpuIsfNpQueues(d *schema.ResourceData, v *models.SystemNpuIsfNpQueues, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemNpuIsfNpQueues{*v}
		for i, cfg := range v2 {
			_ = i
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

func flattenSystemNpuNpQueues(d *schema.ResourceData, v *models.SystemNpuNpQueues, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemNpuNpQueues{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.EthernetType; tmp != nil {
				v["ethernet_type"] = flattenSystemNpuNpQueuesEthernetType(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "ethernet_type"), sort)
			}

			if tmp := cfg.IpProtocol; tmp != nil {
				v["ip_protocol"] = flattenSystemNpuNpQueuesIpProtocol(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "ip_protocol"), sort)
			}

			if tmp := cfg.IpService; tmp != nil {
				v["ip_service"] = flattenSystemNpuNpQueuesIpService(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "ip_service"), sort)
			}

			if tmp := cfg.Profile; tmp != nil {
				v["profile"] = flattenSystemNpuNpQueuesProfile(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "profile"), sort)
			}

			if tmp := cfg.Scheduler; tmp != nil {
				v["scheduler"] = flattenSystemNpuNpQueuesScheduler(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "scheduler"), sort)
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemNpuNpQueuesEthernetType(d *schema.ResourceData, v *[]models.SystemNpuNpQueuesEthernetType, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Queue; tmp != nil {
				v["queue"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			if tmp := cfg.Weight; tmp != nil {
				v["weight"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemNpuNpQueuesIpProtocol(d *schema.ResourceData, v *[]models.SystemNpuNpQueuesIpProtocol, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Protocol; tmp != nil {
				v["protocol"] = *tmp
			}

			if tmp := cfg.Queue; tmp != nil {
				v["queue"] = *tmp
			}

			if tmp := cfg.Weight; tmp != nil {
				v["weight"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemNpuNpQueuesIpService(d *schema.ResourceData, v *[]models.SystemNpuNpQueuesIpService, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Dport; tmp != nil {
				v["dport"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Protocol; tmp != nil {
				v["protocol"] = *tmp
			}

			if tmp := cfg.Queue; tmp != nil {
				v["queue"] = *tmp
			}

			if tmp := cfg.Sport; tmp != nil {
				v["sport"] = *tmp
			}

			if tmp := cfg.Weight; tmp != nil {
				v["weight"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemNpuNpQueuesProfile(d *schema.ResourceData, v *[]models.SystemNpuNpQueuesProfile, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
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

			if tmp := cfg.Dscp0; tmp != nil {
				v["dscp0"] = *tmp
			}

			if tmp := cfg.Dscp1; tmp != nil {
				v["dscp1"] = *tmp
			}

			if tmp := cfg.Dscp10; tmp != nil {
				v["dscp10"] = *tmp
			}

			if tmp := cfg.Dscp11; tmp != nil {
				v["dscp11"] = *tmp
			}

			if tmp := cfg.Dscp12; tmp != nil {
				v["dscp12"] = *tmp
			}

			if tmp := cfg.Dscp13; tmp != nil {
				v["dscp13"] = *tmp
			}

			if tmp := cfg.Dscp14; tmp != nil {
				v["dscp14"] = *tmp
			}

			if tmp := cfg.Dscp15; tmp != nil {
				v["dscp15"] = *tmp
			}

			if tmp := cfg.Dscp16; tmp != nil {
				v["dscp16"] = *tmp
			}

			if tmp := cfg.Dscp17; tmp != nil {
				v["dscp17"] = *tmp
			}

			if tmp := cfg.Dscp18; tmp != nil {
				v["dscp18"] = *tmp
			}

			if tmp := cfg.Dscp19; tmp != nil {
				v["dscp19"] = *tmp
			}

			if tmp := cfg.Dscp2; tmp != nil {
				v["dscp2"] = *tmp
			}

			if tmp := cfg.Dscp20; tmp != nil {
				v["dscp20"] = *tmp
			}

			if tmp := cfg.Dscp21; tmp != nil {
				v["dscp21"] = *tmp
			}

			if tmp := cfg.Dscp22; tmp != nil {
				v["dscp22"] = *tmp
			}

			if tmp := cfg.Dscp23; tmp != nil {
				v["dscp23"] = *tmp
			}

			if tmp := cfg.Dscp24; tmp != nil {
				v["dscp24"] = *tmp
			}

			if tmp := cfg.Dscp25; tmp != nil {
				v["dscp25"] = *tmp
			}

			if tmp := cfg.Dscp26; tmp != nil {
				v["dscp26"] = *tmp
			}

			if tmp := cfg.Dscp27; tmp != nil {
				v["dscp27"] = *tmp
			}

			if tmp := cfg.Dscp28; tmp != nil {
				v["dscp28"] = *tmp
			}

			if tmp := cfg.Dscp29; tmp != nil {
				v["dscp29"] = *tmp
			}

			if tmp := cfg.Dscp3; tmp != nil {
				v["dscp3"] = *tmp
			}

			if tmp := cfg.Dscp30; tmp != nil {
				v["dscp30"] = *tmp
			}

			if tmp := cfg.Dscp31; tmp != nil {
				v["dscp31"] = *tmp
			}

			if tmp := cfg.Dscp32; tmp != nil {
				v["dscp32"] = *tmp
			}

			if tmp := cfg.Dscp33; tmp != nil {
				v["dscp33"] = *tmp
			}

			if tmp := cfg.Dscp34; tmp != nil {
				v["dscp34"] = *tmp
			}

			if tmp := cfg.Dscp35; tmp != nil {
				v["dscp35"] = *tmp
			}

			if tmp := cfg.Dscp36; tmp != nil {
				v["dscp36"] = *tmp
			}

			if tmp := cfg.Dscp37; tmp != nil {
				v["dscp37"] = *tmp
			}

			if tmp := cfg.Dscp38; tmp != nil {
				v["dscp38"] = *tmp
			}

			if tmp := cfg.Dscp39; tmp != nil {
				v["dscp39"] = *tmp
			}

			if tmp := cfg.Dscp4; tmp != nil {
				v["dscp4"] = *tmp
			}

			if tmp := cfg.Dscp40; tmp != nil {
				v["dscp40"] = *tmp
			}

			if tmp := cfg.Dscp41; tmp != nil {
				v["dscp41"] = *tmp
			}

			if tmp := cfg.Dscp42; tmp != nil {
				v["dscp42"] = *tmp
			}

			if tmp := cfg.Dscp43; tmp != nil {
				v["dscp43"] = *tmp
			}

			if tmp := cfg.Dscp44; tmp != nil {
				v["dscp44"] = *tmp
			}

			if tmp := cfg.Dscp45; tmp != nil {
				v["dscp45"] = *tmp
			}

			if tmp := cfg.Dscp46; tmp != nil {
				v["dscp46"] = *tmp
			}

			if tmp := cfg.Dscp47; tmp != nil {
				v["dscp47"] = *tmp
			}

			if tmp := cfg.Dscp48; tmp != nil {
				v["dscp48"] = *tmp
			}

			if tmp := cfg.Dscp49; tmp != nil {
				v["dscp49"] = *tmp
			}

			if tmp := cfg.Dscp5; tmp != nil {
				v["dscp5"] = *tmp
			}

			if tmp := cfg.Dscp50; tmp != nil {
				v["dscp50"] = *tmp
			}

			if tmp := cfg.Dscp51; tmp != nil {
				v["dscp51"] = *tmp
			}

			if tmp := cfg.Dscp52; tmp != nil {
				v["dscp52"] = *tmp
			}

			if tmp := cfg.Dscp53; tmp != nil {
				v["dscp53"] = *tmp
			}

			if tmp := cfg.Dscp54; tmp != nil {
				v["dscp54"] = *tmp
			}

			if tmp := cfg.Dscp55; tmp != nil {
				v["dscp55"] = *tmp
			}

			if tmp := cfg.Dscp56; tmp != nil {
				v["dscp56"] = *tmp
			}

			if tmp := cfg.Dscp57; tmp != nil {
				v["dscp57"] = *tmp
			}

			if tmp := cfg.Dscp58; tmp != nil {
				v["dscp58"] = *tmp
			}

			if tmp := cfg.Dscp59; tmp != nil {
				v["dscp59"] = *tmp
			}

			if tmp := cfg.Dscp6; tmp != nil {
				v["dscp6"] = *tmp
			}

			if tmp := cfg.Dscp60; tmp != nil {
				v["dscp60"] = *tmp
			}

			if tmp := cfg.Dscp61; tmp != nil {
				v["dscp61"] = *tmp
			}

			if tmp := cfg.Dscp62; tmp != nil {
				v["dscp62"] = *tmp
			}

			if tmp := cfg.Dscp63; tmp != nil {
				v["dscp63"] = *tmp
			}

			if tmp := cfg.Dscp7; tmp != nil {
				v["dscp7"] = *tmp
			}

			if tmp := cfg.Dscp8; tmp != nil {
				v["dscp8"] = *tmp
			}

			if tmp := cfg.Dscp9; tmp != nil {
				v["dscp9"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			if tmp := cfg.Weight; tmp != nil {
				v["weight"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemNpuNpQueuesScheduler(d *schema.ResourceData, v *[]models.SystemNpuNpQueuesScheduler, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

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

func flattenSystemNpuPortCpuMap(d *schema.ResourceData, v *[]models.SystemNpuPortCpuMap, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.CpuCore; tmp != nil {
				v["cpu_core"] = *tmp
			}

			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "interface")
	}

	return flat
}

func flattenSystemNpuPortNpuMap(d *schema.ResourceData, v *[]models.SystemNpuPortNpuMap, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			if tmp := cfg.NpuGroupIndex; tmp != nil {
				v["npu_group_index"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "interface")
	}

	return flat
}

func flattenSystemNpuPriorityProtocol(d *schema.ResourceData, v *models.SystemNpuPriorityProtocol, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemNpuPriorityProtocol{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Bfd; tmp != nil {
				v["bfd"] = *tmp
			}

			if tmp := cfg.Bgp; tmp != nil {
				v["bgp"] = *tmp
			}

			if tmp := cfg.Slbc; tmp != nil {
				v["slbc"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemNpuTcpTimeoutProfile(d *schema.ResourceData, v *[]models.SystemNpuTcpTimeoutProfile, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.CloseWait; tmp != nil {
				v["close_wait"] = *tmp
			}

			if tmp := cfg.FinWait; tmp != nil {
				v["fin_wait"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.SynSent; tmp != nil {
				v["syn_sent"] = *tmp
			}

			if tmp := cfg.SynWait; tmp != nil {
				v["syn_wait"] = *tmp
			}

			if tmp := cfg.TcpIdle; tmp != nil {
				v["tcp_idle"] = *tmp
			}

			if tmp := cfg.TimeWait; tmp != nil {
				v["time_wait"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemNpuUdpTimeoutProfile(d *schema.ResourceData, v *[]models.SystemNpuUdpTimeoutProfile, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.UdpIdle; tmp != nil {
				v["udp_idle"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectSystemNpu(d *schema.ResourceData, o *models.SystemNpu, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.CapwapOffload != nil {
		v := *o.CapwapOffload

		if err = d.Set("capwap_offload", v); err != nil {
			return diag.Errorf("error reading capwap_offload: %v", err)
		}
	}

	if o.DedicatedManagementAffinity != nil {
		v := *o.DedicatedManagementAffinity

		if err = d.Set("dedicated_management_affinity", v); err != nil {
			return diag.Errorf("error reading dedicated_management_affinity: %v", err)
		}
	}

	if o.DedicatedManagementCpu != nil {
		v := *o.DedicatedManagementCpu

		if err = d.Set("dedicated_management_cpu", v); err != nil {
			return diag.Errorf("error reading dedicated_management_cpu: %v", err)
		}
	}

	if o.DefaultQosType != nil {
		v := *o.DefaultQosType

		if err = d.Set("default_qos_type", v); err != nil {
			return diag.Errorf("error reading default_qos_type: %v", err)
		}
	}

	if _, ok := d.GetOk("dos_options"); ok {
		if o.DosOptions != nil {
			if err = d.Set("dos_options", flattenSystemNpuDosOptions(d, o.DosOptions, "dos_options", sort)); err != nil {
				return diag.Errorf("error reading dos_options: %v", err)
			}
		}
	}

	if o.DoubleLevelMcastOffload != nil {
		v := *o.DoubleLevelMcastOffload

		if err = d.Set("double_level_mcast_offload", v); err != nil {
			return diag.Errorf("error reading double_level_mcast_offload: %v", err)
		}
	}

	if o.DseTimeout != nil {
		v := *o.DseTimeout

		if err = d.Set("dse_timeout", v); err != nil {
			return diag.Errorf("error reading dse_timeout: %v", err)
		}
	}

	if o.DswDtsProfile != nil {
		if err = d.Set("dsw_dts_profile", flattenSystemNpuDswDtsProfile(d, o.DswDtsProfile, "dsw_dts_profile", sort)); err != nil {
			return diag.Errorf("error reading dsw_dts_profile: %v", err)
		}
	}

	if o.DswQueueDtsProfile != nil {
		if err = d.Set("dsw_queue_dts_profile", flattenSystemNpuDswQueueDtsProfile(d, o.DswQueueDtsProfile, "dsw_queue_dts_profile", sort)); err != nil {
			return diag.Errorf("error reading dsw_queue_dts_profile: %v", err)
		}
	}

	if o.Fastpath != nil {
		v := *o.Fastpath

		if err = d.Set("fastpath", v); err != nil {
			return diag.Errorf("error reading fastpath: %v", err)
		}
	}

	if _, ok := d.GetOk("fp_anomaly"); ok {
		if o.FpAnomaly != nil {
			if err = d.Set("fp_anomaly", flattenSystemNpuFpAnomaly(d, o.FpAnomaly, "fp_anomaly", sort)); err != nil {
				return diag.Errorf("error reading fp_anomaly: %v", err)
			}
		}
	}

	if o.GtpSupport != nil {
		v := *o.GtpSupport

		if err = d.Set("gtp_support", v); err != nil {
			return diag.Errorf("error reading gtp_support: %v", err)
		}
	}

	if o.HashTblSpread != nil {
		v := *o.HashTblSpread

		if err = d.Set("hash_tbl_spread", v); err != nil {
			return diag.Errorf("error reading hash_tbl_spread: %v", err)
		}
	}

	if _, ok := d.GetOk("hpe"); ok {
		if o.Hpe != nil {
			if err = d.Set("hpe", flattenSystemNpuHpe(d, o.Hpe, "hpe", sort)); err != nil {
				return diag.Errorf("error reading hpe: %v", err)
			}
		}
	}

	if o.HtabDediQueueNr != nil {
		v := *o.HtabDediQueueNr

		if err = d.Set("htab_dedi_queue_nr", v); err != nil {
			return diag.Errorf("error reading htab_dedi_queue_nr: %v", err)
		}
	}

	if o.HtabMsgQueue != nil {
		v := *o.HtabMsgQueue

		if err = d.Set("htab_msg_queue", v); err != nil {
			return diag.Errorf("error reading htab_msg_queue: %v", err)
		}
	}

	if o.HtxIcmpCsumChk != nil {
		v := *o.HtxIcmpCsumChk

		if err = d.Set("htx_icmp_csum_chk", v); err != nil {
			return diag.Errorf("error reading htx_icmp_csum_chk: %v", err)
		}
	}

	if o.InboundDscpCopyPort != nil {
		if err = d.Set("inbound_dscp_copy_port", flattenSystemNpuInboundDscpCopyPort(d, o.InboundDscpCopyPort, "inbound_dscp_copy_port", sort)); err != nil {
			return diag.Errorf("error reading inbound_dscp_copy_port: %v", err)
		}
	}

	if o.IpFragmentOffload != nil {
		v := *o.IpFragmentOffload

		if err = d.Set("ip_fragment_offload", v); err != nil {
			return diag.Errorf("error reading ip_fragment_offload: %v", err)
		}
	}

	if _, ok := d.GetOk("ip_reassembly"); ok {
		if o.IpReassembly != nil {
			if err = d.Set("ip_reassembly", flattenSystemNpuIpReassembly(d, o.IpReassembly, "ip_reassembly", sort)); err != nil {
				return diag.Errorf("error reading ip_reassembly: %v", err)
			}
		}
	}

	if o.IppoolOverloadHigh != nil {
		v := *o.IppoolOverloadHigh

		if err = d.Set("ippool_overload_high", v); err != nil {
			return diag.Errorf("error reading ippool_overload_high: %v", err)
		}
	}

	if o.IppoolOverloadLow != nil {
		v := *o.IppoolOverloadLow

		if err = d.Set("ippool_overload_low", v); err != nil {
			return diag.Errorf("error reading ippool_overload_low: %v", err)
		}
	}

	if o.IpsecDecSubengineMask != nil {
		v := *o.IpsecDecSubengineMask

		if err = d.Set("ipsec_dec_subengine_mask", v); err != nil {
			return diag.Errorf("error reading ipsec_dec_subengine_mask: %v", err)
		}
	}

	if o.IpsecEncSubengineMask != nil {
		v := *o.IpsecEncSubengineMask

		if err = d.Set("ipsec_enc_subengine_mask", v); err != nil {
			return diag.Errorf("error reading ipsec_enc_subengine_mask: %v", err)
		}
	}

	if o.IpsecInboundCache != nil {
		v := *o.IpsecInboundCache

		if err = d.Set("ipsec_inbound_cache", v); err != nil {
			return diag.Errorf("error reading ipsec_inbound_cache: %v", err)
		}
	}

	if o.IpsecMtuOverride != nil {
		v := *o.IpsecMtuOverride

		if err = d.Set("ipsec_mtu_override", v); err != nil {
			return diag.Errorf("error reading ipsec_mtu_override: %v", err)
		}
	}

	if o.IpsecObNpSel != nil {
		v := *o.IpsecObNpSel

		if err = d.Set("ipsec_ob_np_sel", v); err != nil {
			return diag.Errorf("error reading ipsec_ob_np_sel: %v", err)
		}
	}

	if o.IpsecOverVlink != nil {
		v := *o.IpsecOverVlink

		if err = d.Set("ipsec_over_vlink", v); err != nil {
			return diag.Errorf("error reading ipsec_over_vlink: %v", err)
		}
	}

	if _, ok := d.GetOk("isf_np_queues"); ok {
		if o.IsfNpQueues != nil {
			if err = d.Set("isf_np_queues", flattenSystemNpuIsfNpQueues(d, o.IsfNpQueues, "isf_np_queues", sort)); err != nil {
				return diag.Errorf("error reading isf_np_queues: %v", err)
			}
		}
	}

	if o.MaxSessionTimeout != nil {
		v := *o.MaxSessionTimeout

		if err = d.Set("max_session_timeout", v); err != nil {
			return diag.Errorf("error reading max_session_timeout: %v", err)
		}
	}

	if o.McastSessionAccounting != nil {
		v := *o.McastSessionAccounting

		if err = d.Set("mcast_session_accounting", v); err != nil {
			return diag.Errorf("error reading mcast_session_accounting: %v", err)
		}
	}

	if o.NapiBreakInterval != nil {
		v := *o.NapiBreakInterval

		if err = d.Set("napi_break_interval", v); err != nil {
			return diag.Errorf("error reading napi_break_interval: %v", err)
		}
	}

	if _, ok := d.GetOk("np_queues"); ok {
		if o.NpQueues != nil {
			if err = d.Set("np_queues", flattenSystemNpuNpQueues(d, o.NpQueues, "np_queues", sort)); err != nil {
				return diag.Errorf("error reading np_queues: %v", err)
			}
		}
	}

	if o.Np6CpsOptimizationMode != nil {
		v := *o.Np6CpsOptimizationMode

		if err = d.Set("np6_cps_optimization_mode", v); err != nil {
			return diag.Errorf("error reading np6_cps_optimization_mode: %v", err)
		}
	}

	if o.NpuGroupEffectiveScope != nil {
		v := *o.NpuGroupEffectiveScope

		if err = d.Set("npu_group_effective_scope", v); err != nil {
			return diag.Errorf("error reading npu_group_effective_scope: %v", err)
		}
	}

	if o.PbaEim != nil {
		v := *o.PbaEim

		if err = d.Set("pba_eim", v); err != nil {
			return diag.Errorf("error reading pba_eim: %v", err)
		}
	}

	if o.PerPolicyAccounting != nil {
		v := *o.PerPolicyAccounting

		if err = d.Set("per_policy_accounting", v); err != nil {
			return diag.Errorf("error reading per_policy_accounting: %v", err)
		}
	}

	if o.PerSessionAccounting != nil {
		v := *o.PerSessionAccounting

		if err = d.Set("per_session_accounting", v); err != nil {
			return diag.Errorf("error reading per_session_accounting: %v", err)
		}
	}

	if o.PolicyOffloadLevel != nil {
		v := *o.PolicyOffloadLevel

		if err = d.Set("policy_offload_level", v); err != nil {
			return diag.Errorf("error reading policy_offload_level: %v", err)
		}
	}

	if o.PortCpuMap != nil {
		if err = d.Set("port_cpu_map", flattenSystemNpuPortCpuMap(d, o.PortCpuMap, "port_cpu_map", sort)); err != nil {
			return diag.Errorf("error reading port_cpu_map: %v", err)
		}
	}

	if o.PortNpuMap != nil {
		if err = d.Set("port_npu_map", flattenSystemNpuPortNpuMap(d, o.PortNpuMap, "port_npu_map", sort)); err != nil {
			return diag.Errorf("error reading port_npu_map: %v", err)
		}
	}

	if _, ok := d.GetOk("priority_protocol"); ok {
		if o.PriorityProtocol != nil {
			if err = d.Set("priority_protocol", flattenSystemNpuPriorityProtocol(d, o.PriorityProtocol, "priority_protocol", sort)); err != nil {
				return diag.Errorf("error reading priority_protocol: %v", err)
			}
		}
	}

	if o.QosMode != nil {
		v := *o.QosMode

		if err = d.Set("qos_mode", v); err != nil {
			return diag.Errorf("error reading qos_mode: %v", err)
		}
	}

	if o.QtmBufMode != nil {
		v := *o.QtmBufMode

		if err = d.Set("qtm_buf_mode", v); err != nil {
			return diag.Errorf("error reading qtm_buf_mode: %v", err)
		}
	}

	if o.RdpOffload != nil {
		v := *o.RdpOffload

		if err = d.Set("rdp_offload", v); err != nil {
			return diag.Errorf("error reading rdp_offload: %v", err)
		}
	}

	if o.RecoverNp6Link != nil {
		v := *o.RecoverNp6Link

		if err = d.Set("recover_np6_link", v); err != nil {
			return diag.Errorf("error reading recover_np6_link: %v", err)
		}
	}

	if o.SessionAcctInterval != nil {
		v := *o.SessionAcctInterval

		if err = d.Set("session_acct_interval", v); err != nil {
			return diag.Errorf("error reading session_acct_interval: %v", err)
		}
	}

	if o.ShapingStats != nil {
		v := *o.ShapingStats

		if err = d.Set("shaping_stats", v); err != nil {
			return diag.Errorf("error reading shaping_stats: %v", err)
		}
	}

	if o.SseBackpressure != nil {
		v := *o.SseBackpressure

		if err = d.Set("sse_backpressure", v); err != nil {
			return diag.Errorf("error reading sse_backpressure: %v", err)
		}
	}

	if o.StripClearTextPadding != nil {
		v := *o.StripClearTextPadding

		if err = d.Set("strip_clear_text_padding", v); err != nil {
			return diag.Errorf("error reading strip_clear_text_padding: %v", err)
		}
	}

	if o.StripEspPadding != nil {
		v := *o.StripEspPadding

		if err = d.Set("strip_esp_padding", v); err != nil {
			return diag.Errorf("error reading strip_esp_padding: %v", err)
		}
	}

	if o.SwNpBandwidth != nil {
		v := *o.SwNpBandwidth

		if err = d.Set("sw_np_bandwidth", v); err != nil {
			return diag.Errorf("error reading sw_np_bandwidth: %v", err)
		}
	}

	if o.TcpRstTimeout != nil {
		v := *o.TcpRstTimeout

		if err = d.Set("tcp_rst_timeout", v); err != nil {
			return diag.Errorf("error reading tcp_rst_timeout: %v", err)
		}
	}

	if o.TcpTimeoutProfile != nil {
		if err = d.Set("tcp_timeout_profile", flattenSystemNpuTcpTimeoutProfile(d, o.TcpTimeoutProfile, "tcp_timeout_profile", sort)); err != nil {
			return diag.Errorf("error reading tcp_timeout_profile: %v", err)
		}
	}

	if o.UdpTimeoutProfile != nil {
		if err = d.Set("udp_timeout_profile", flattenSystemNpuUdpTimeoutProfile(d, o.UdpTimeoutProfile, "udp_timeout_profile", sort)); err != nil {
			return diag.Errorf("error reading udp_timeout_profile: %v", err)
		}
	}

	if o.VlanLookupCache != nil {
		v := *o.VlanLookupCache

		if err = d.Set("vlan_lookup_cache", v); err != nil {
			return diag.Errorf("error reading vlan_lookup_cache: %v", err)
		}
	}

	return nil
}

func expandSystemNpuDosOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemNpuDosOptions, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNpuDosOptions

	for i := range l {
		tmp := models.SystemNpuDosOptions{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.npu_dos_meter_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NpuDosMeterMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.npu_dos_tpe_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NpuDosTpeMode = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemNpuDswDtsProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemNpuDswDtsProfile, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNpuDswDtsProfile

	for i := range l {
		tmp := models.SystemNpuDswDtsProfile{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.min_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MinLimit = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.profile_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ProfileId = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.step", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Step = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemNpuDswQueueDtsProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemNpuDswQueueDtsProfile, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNpuDswQueueDtsProfile

	for i := range l {
		tmp := models.SystemNpuDswQueueDtsProfile{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.iport", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Iport = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.oport", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Oport = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.profile_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ProfileId = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.queue_select", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.QueueSelect = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemNpuFpAnomaly(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemNpuFpAnomaly, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNpuFpAnomaly

	for i := range l {
		tmp := models.SystemNpuFpAnomaly{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.icmp_csum_err", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IcmpCsumErr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.icmp_frag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IcmpFrag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.icmp_land", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IcmpLand = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv4_csum_err", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv4CsumErr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv4_land", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv4Land = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv4_optlsrr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv4Optlsrr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv4_optrr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv4Optrr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv4_optsecurity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv4Optsecurity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv4_optssrr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv4Optssrr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv4_optstream", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv4Optstream = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv4_opttimestamp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv4Opttimestamp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv4_proto_err", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv4ProtoErr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv4_unknopt", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv4Unknopt = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv6_daddr_err", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv6DaddrErr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv6_land", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv6Land = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv6_optendpid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv6Optendpid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv6_opthomeaddr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv6Opthomeaddr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv6_optinvld", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv6Optinvld = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv6_optjumbo", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv6Optjumbo = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv6_optnsap", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv6Optnsap = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv6_optralert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv6Optralert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv6_opttunnel", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv6Opttunnel = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv6_proto_err", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv6ProtoErr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv6_saddr_err", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv6SaddrErr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv6_unknopt", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv6Unknopt = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sctp_csum_err", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SctpCsumErr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_csum_err", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TcpCsumErr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_fin_noack", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TcpFinNoack = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_fin_only", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TcpFinOnly = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_land", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TcpLand = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_no_flag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TcpNoFlag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_syn_data", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TcpSynData = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_syn_fin", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TcpSynFin = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_winnuke", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TcpWinnuke = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.udp_csum_err", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UdpCsumErr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.udp_land", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UdpLand = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemNpuHpe(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemNpuHpe, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNpuHpe

	for i := range l {
		tmp := models.SystemNpuHpe{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.all_protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AllProtocol = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.arp_max", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ArpMax = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.enable_shaper", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EnableShaper = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.esp_max", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.EspMax = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.high_priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.HighPriority = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.icmp_max", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.IcmpMax = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip_frag_max", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.IpFragMax = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip_others_max", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.IpOthersMax = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.l2_others_max", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.L2OthersMax = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sctp_max", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SctpMax = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_max", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.TcpMax = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcpfin_rst_max", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.TcpfinRstMax = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcpsyn_ack_max", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.TcpsynAckMax = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcpsyn_max", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.TcpsynMax = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.udp_max", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UdpMax = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemNpuInboundDscpCopyPort(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemNpuInboundDscpCopyPort, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNpuInboundDscpCopyPort

	for i := range l {
		tmp := models.SystemNpuInboundDscpCopyPort{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemNpuIpReassembly(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemNpuIpReassembly, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNpuIpReassembly

	for i := range l {
		tmp := models.SystemNpuIpReassembly{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.max_timeout", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaxTimeout = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.min_timeout", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MinTimeout = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemNpuIsfNpQueues(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemNpuIsfNpQueues, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNpuIsfNpQueues

	for i := range l {
		tmp := models.SystemNpuIsfNpQueues{}
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
	return &result[0], nil
}

func expandSystemNpuNpQueues(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemNpuNpQueues, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNpuNpQueues

	for i := range l {
		tmp := models.SystemNpuNpQueues{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.ethernet_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemNpuNpQueuesEthernetType(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemNpuNpQueuesEthernetType
			// 	}
			tmp.EthernetType = v2

		}

		pre_append = fmt.Sprintf("%s.%d.ip_protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemNpuNpQueuesIpProtocol(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemNpuNpQueuesIpProtocol
			// 	}
			tmp.IpProtocol = v2

		}

		pre_append = fmt.Sprintf("%s.%d.ip_service", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemNpuNpQueuesIpService(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemNpuNpQueuesIpService
			// 	}
			tmp.IpService = v2

		}

		pre_append = fmt.Sprintf("%s.%d.profile", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemNpuNpQueuesProfile(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemNpuNpQueuesProfile
			// 	}
			tmp.Profile = v2

		}

		pre_append = fmt.Sprintf("%s.%d.scheduler", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemNpuNpQueuesScheduler(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemNpuNpQueuesScheduler
			// 	}
			tmp.Scheduler = v2

		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemNpuNpQueuesEthernetType(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemNpuNpQueuesEthernetType, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNpuNpQueuesEthernetType

	for i := range l {
		tmp := models.SystemNpuNpQueuesEthernetType{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.queue", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Queue = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
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

func expandSystemNpuNpQueuesIpProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemNpuNpQueuesIpProtocol, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNpuNpQueuesIpProtocol

	for i := range l {
		tmp := models.SystemNpuNpQueuesIpProtocol{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Protocol = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.queue", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Queue = &v3
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

func expandSystemNpuNpQueuesIpService(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemNpuNpQueuesIpService, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNpuNpQueuesIpService

	for i := range l {
		tmp := models.SystemNpuNpQueuesIpService{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dport", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Dport = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Protocol = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.queue", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Queue = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sport", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Sport = &v3
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

func expandSystemNpuNpQueuesProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemNpuNpQueuesProfile, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNpuNpQueuesProfile

	for i := range l {
		tmp := models.SystemNpuNpQueuesProfile{}
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

		pre_append = fmt.Sprintf("%s.%d.dscp0", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp0 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp1", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp1 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp10", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp10 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp11", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp11 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp12", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp12 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp13", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp13 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp14", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp14 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp15", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp15 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp16", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp16 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp17", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp17 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp18", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp18 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp19", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp19 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp2 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp20", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp20 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp21", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp21 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp22", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp22 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp23", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp23 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp24", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp24 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp25", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp25 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp26", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp26 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp27", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp27 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp28", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp28 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp29", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp29 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp3", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp3 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp30", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp30 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp31", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp31 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp32", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp32 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp33", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp33 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp34", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp34 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp35", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp35 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp36", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp36 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp37", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp37 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp38", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp38 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp39", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp39 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp40", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp40 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp41", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp41 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp42", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp42 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp43", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp43 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp44", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp44 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp45", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp45 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp46", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp46 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp47", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp47 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp48", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp48 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp49", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp49 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp5", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp5 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp50", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp50 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp51", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp51 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp52", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp52 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp53", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp53 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp54", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp54 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp55", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp55 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp56", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp56 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp57", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp57 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp58", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp58 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp59", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp59 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp60", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp60 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp61", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp61 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp62", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp62 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp63", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp63 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp7", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp7 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp8", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp8 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dscp9", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dscp9 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
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

func expandSystemNpuNpQueuesScheduler(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemNpuNpQueuesScheduler, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNpuNpQueuesScheduler

	for i := range l {
		tmp := models.SystemNpuNpQueuesScheduler{}
		var pre_append string

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

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemNpuPortCpuMap(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemNpuPortCpuMap, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNpuPortCpuMap

	for i := range l {
		tmp := models.SystemNpuPortCpuMap{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.cpu_core", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CpuCore = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemNpuPortNpuMap(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemNpuPortNpuMap, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNpuPortNpuMap

	for i := range l {
		tmp := models.SystemNpuPortNpuMap{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.npu_group_index", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.NpuGroupIndex = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemNpuPriorityProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemNpuPriorityProtocol, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNpuPriorityProtocol

	for i := range l {
		tmp := models.SystemNpuPriorityProtocol{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.bfd", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Bfd = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bgp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Bgp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.slbc", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Slbc = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemNpuTcpTimeoutProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemNpuTcpTimeoutProfile, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNpuTcpTimeoutProfile

	for i := range l {
		tmp := models.SystemNpuTcpTimeoutProfile{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.close_wait", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.CloseWait = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fin_wait", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.FinWait = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.syn_sent", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SynSent = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.syn_wait", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SynWait = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_idle", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.TcpIdle = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.time_wait", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.TimeWait = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemNpuUdpTimeoutProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemNpuUdpTimeoutProfile, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNpuUdpTimeoutProfile

	for i := range l {
		tmp := models.SystemNpuUdpTimeoutProfile{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.udp_idle", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UdpIdle = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemNpu(d *schema.ResourceData, sv string) (*models.SystemNpu, diag.Diagnostics) {
	obj := models.SystemNpu{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("capwap_offload"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("capwap_offload", sv)
				diags = append(diags, e)
			}
			obj.CapwapOffload = &v2
		}
	}
	if v1, ok := d.GetOk("dedicated_management_affinity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("dedicated_management_affinity", sv)
				diags = append(diags, e)
			}
			obj.DedicatedManagementAffinity = &v2
		}
	}
	if v1, ok := d.GetOk("dedicated_management_cpu"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dedicated_management_cpu", sv)
				diags = append(diags, e)
			}
			obj.DedicatedManagementCpu = &v2
		}
	}
	if v1, ok := d.GetOk("default_qos_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
				e := utils.AttributeVersionWarning("default_qos_type", sv)
				diags = append(diags, e)
			}
			obj.DefaultQosType = &v2
		}
	}
	if v, ok := d.GetOk("dos_options"); ok {
		if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
			e := utils.AttributeVersionWarning("dos_options", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemNpuDosOptions(d, v, "dos_options", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DosOptions = t
		}
	} else if d.HasChange("dos_options") {
		old, new := d.GetChange("dos_options")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DosOptions = &models.SystemNpuDosOptions{}
		}
	}
	if v1, ok := d.GetOk("double_level_mcast_offload"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
				e := utils.AttributeVersionWarning("double_level_mcast_offload", sv)
				diags = append(diags, e)
			}
			obj.DoubleLevelMcastOffload = &v2
		}
	}
	if v1, ok := d.GetOk("dse_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
				e := utils.AttributeVersionWarning("dse_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DseTimeout = &tmp
		}
	}
	if v, ok := d.GetOk("dsw_dts_profile"); ok {
		if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
			e := utils.AttributeVersionWarning("dsw_dts_profile", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemNpuDswDtsProfile(d, v, "dsw_dts_profile", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DswDtsProfile = t
		}
	} else if d.HasChange("dsw_dts_profile") {
		old, new := d.GetChange("dsw_dts_profile")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DswDtsProfile = &[]models.SystemNpuDswDtsProfile{}
		}
	}
	if v, ok := d.GetOk("dsw_queue_dts_profile"); ok {
		if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
			e := utils.AttributeVersionWarning("dsw_queue_dts_profile", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemNpuDswQueueDtsProfile(d, v, "dsw_queue_dts_profile", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DswQueueDtsProfile = t
		}
	} else if d.HasChange("dsw_queue_dts_profile") {
		old, new := d.GetChange("dsw_queue_dts_profile")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DswQueueDtsProfile = &[]models.SystemNpuDswQueueDtsProfile{}
		}
	}
	if v1, ok := d.GetOk("fastpath"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.6") {
				e := utils.AttributeVersionWarning("fastpath", sv)
				diags = append(diags, e)
			}
			obj.Fastpath = &v2
		}
	}
	if v, ok := d.GetOk("fp_anomaly"); ok {
		if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
			e := utils.AttributeVersionWarning("fp_anomaly", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemNpuFpAnomaly(d, v, "fp_anomaly", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.FpAnomaly = t
		}
	} else if d.HasChange("fp_anomaly") {
		old, new := d.GetChange("fp_anomaly")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.FpAnomaly = &models.SystemNpuFpAnomaly{}
		}
	}
	if v1, ok := d.GetOk("gtp_support"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
				e := utils.AttributeVersionWarning("gtp_support", sv)
				diags = append(diags, e)
			}
			obj.GtpSupport = &v2
		}
	}
	if v1, ok := d.GetOk("hash_tbl_spread"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
				e := utils.AttributeVersionWarning("hash_tbl_spread", sv)
				diags = append(diags, e)
			}
			obj.HashTblSpread = &v2
		}
	}
	if v, ok := d.GetOk("hpe"); ok {
		if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
			e := utils.AttributeVersionWarning("hpe", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemNpuHpe(d, v, "hpe", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Hpe = t
		}
	} else if d.HasChange("hpe") {
		old, new := d.GetChange("hpe")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Hpe = &models.SystemNpuHpe{}
		}
	}
	if v1, ok := d.GetOk("htab_dedi_queue_nr"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
				e := utils.AttributeVersionWarning("htab_dedi_queue_nr", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HtabDediQueueNr = &tmp
		}
	}
	if v1, ok := d.GetOk("htab_msg_queue"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
				e := utils.AttributeVersionWarning("htab_msg_queue", sv)
				diags = append(diags, e)
			}
			obj.HtabMsgQueue = &v2
		}
	}
	if v1, ok := d.GetOk("htx_icmp_csum_chk"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("htx_icmp_csum_chk", sv)
				diags = append(diags, e)
			}
			obj.HtxIcmpCsumChk = &v2
		}
	}
	if v, ok := d.GetOk("inbound_dscp_copy_port"); ok {
		if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
			e := utils.AttributeVersionWarning("inbound_dscp_copy_port", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemNpuInboundDscpCopyPort(d, v, "inbound_dscp_copy_port", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InboundDscpCopyPort = t
		}
	} else if d.HasChange("inbound_dscp_copy_port") {
		old, new := d.GetChange("inbound_dscp_copy_port")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InboundDscpCopyPort = &[]models.SystemNpuInboundDscpCopyPort{}
		}
	}
	if v1, ok := d.GetOk("ip_fragment_offload"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("ip_fragment_offload", sv)
				diags = append(diags, e)
			}
			obj.IpFragmentOffload = &v2
		}
	}
	if v, ok := d.GetOk("ip_reassembly"); ok {
		if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
			e := utils.AttributeVersionWarning("ip_reassembly", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemNpuIpReassembly(d, v, "ip_reassembly", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.IpReassembly = t
		}
	} else if d.HasChange("ip_reassembly") {
		old, new := d.GetChange("ip_reassembly")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.IpReassembly = &models.SystemNpuIpReassembly{}
		}
	}
	if v1, ok := d.GetOk("ippool_overload_high"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
				e := utils.AttributeVersionWarning("ippool_overload_high", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IppoolOverloadHigh = &tmp
		}
	}
	if v1, ok := d.GetOk("ippool_overload_low"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
				e := utils.AttributeVersionWarning("ippool_overload_low", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IppoolOverloadLow = &tmp
		}
	}
	if v1, ok := d.GetOk("ipsec_dec_subengine_mask"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.6") {
				e := utils.AttributeVersionWarning("ipsec_dec_subengine_mask", sv)
				diags = append(diags, e)
			}
			obj.IpsecDecSubengineMask = &v2
		}
	}
	if v1, ok := d.GetOk("ipsec_enc_subengine_mask"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.6") {
				e := utils.AttributeVersionWarning("ipsec_enc_subengine_mask", sv)
				diags = append(diags, e)
			}
			obj.IpsecEncSubengineMask = &v2
		}
	}
	if v1, ok := d.GetOk("ipsec_inbound_cache"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.6") {
				e := utils.AttributeVersionWarning("ipsec_inbound_cache", sv)
				diags = append(diags, e)
			}
			obj.IpsecInboundCache = &v2
		}
	}
	if v1, ok := d.GetOk("ipsec_mtu_override"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.6") {
				e := utils.AttributeVersionWarning("ipsec_mtu_override", sv)
				diags = append(diags, e)
			}
			obj.IpsecMtuOverride = &v2
		}
	}
	if v1, ok := d.GetOk("ipsec_ob_np_sel"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
				e := utils.AttributeVersionWarning("ipsec_ob_np_sel", sv)
				diags = append(diags, e)
			}
			obj.IpsecObNpSel = &v2
		}
	}
	if v1, ok := d.GetOk("ipsec_over_vlink"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.6") {
				e := utils.AttributeVersionWarning("ipsec_over_vlink", sv)
				diags = append(diags, e)
			}
			obj.IpsecOverVlink = &v2
		}
	}
	if v, ok := d.GetOk("isf_np_queues"); ok {
		if !utils.CheckVer(sv, "", "v6.4.6") {
			e := utils.AttributeVersionWarning("isf_np_queues", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemNpuIsfNpQueues(d, v, "isf_np_queues", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.IsfNpQueues = t
		}
	} else if d.HasChange("isf_np_queues") {
		old, new := d.GetChange("isf_np_queues")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.IsfNpQueues = &models.SystemNpuIsfNpQueues{}
		}
	}
	if v1, ok := d.GetOk("max_session_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
				e := utils.AttributeVersionWarning("max_session_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxSessionTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("mcast_session_accounting"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.6") {
				e := utils.AttributeVersionWarning("mcast_session_accounting", sv)
				diags = append(diags, e)
			}
			obj.McastSessionAccounting = &v2
		}
	}
	if v1, ok := d.GetOk("napi_break_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
				e := utils.AttributeVersionWarning("napi_break_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.NapiBreakInterval = &tmp
		}
	}
	if v, ok := d.GetOk("np_queues"); ok {
		if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
			e := utils.AttributeVersionWarning("np_queues", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemNpuNpQueues(d, v, "np_queues", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.NpQueues = t
		}
	} else if d.HasChange("np_queues") {
		old, new := d.GetChange("np_queues")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.NpQueues = &models.SystemNpuNpQueues{}
		}
	}
	if v1, ok := d.GetOk("np6_cps_optimization_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.6") {
				e := utils.AttributeVersionWarning("np6_cps_optimization_mode", sv)
				diags = append(diags, e)
			}
			obj.Np6CpsOptimizationMode = &v2
		}
	}
	if v1, ok := d.GetOk("npu_group_effective_scope"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("npu_group_effective_scope", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.NpuGroupEffectiveScope = &tmp
		}
	}
	if v1, ok := d.GetOk("pba_eim"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
				e := utils.AttributeVersionWarning("pba_eim", sv)
				diags = append(diags, e)
			}
			obj.PbaEim = &v2
		}
	}
	if v1, ok := d.GetOk("per_policy_accounting"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("per_policy_accounting", sv)
				diags = append(diags, e)
			}
			obj.PerPolicyAccounting = &v2
		}
	}
	if v1, ok := d.GetOk("per_session_accounting"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
				e := utils.AttributeVersionWarning("per_session_accounting", sv)
				diags = append(diags, e)
			}
			obj.PerSessionAccounting = &v2
		}
	}
	if v1, ok := d.GetOk("policy_offload_level"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
				e := utils.AttributeVersionWarning("policy_offload_level", sv)
				diags = append(diags, e)
			}
			obj.PolicyOffloadLevel = &v2
		}
	}
	if v, ok := d.GetOk("port_cpu_map"); ok {
		if !utils.CheckVer(sv, "", "v6.4.6") {
			e := utils.AttributeVersionWarning("port_cpu_map", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemNpuPortCpuMap(d, v, "port_cpu_map", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.PortCpuMap = t
		}
	} else if d.HasChange("port_cpu_map") {
		old, new := d.GetChange("port_cpu_map")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.PortCpuMap = &[]models.SystemNpuPortCpuMap{}
		}
	}
	if v, ok := d.GetOk("port_npu_map"); ok {
		if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
			e := utils.AttributeVersionWarning("port_npu_map", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemNpuPortNpuMap(d, v, "port_npu_map", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.PortNpuMap = t
		}
	} else if d.HasChange("port_npu_map") {
		old, new := d.GetChange("port_npu_map")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.PortNpuMap = &[]models.SystemNpuPortNpuMap{}
		}
	}
	if v, ok := d.GetOk("priority_protocol"); ok {
		if !utils.CheckVer(sv, "", "v6.4.6") {
			e := utils.AttributeVersionWarning("priority_protocol", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemNpuPriorityProtocol(d, v, "priority_protocol", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.PriorityProtocol = t
		}
	} else if d.HasChange("priority_protocol") {
		old, new := d.GetChange("priority_protocol")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.PriorityProtocol = &models.SystemNpuPriorityProtocol{}
		}
	}
	if v1, ok := d.GetOk("qos_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.6") {
				e := utils.AttributeVersionWarning("qos_mode", sv)
				diags = append(diags, e)
			}
			obj.QosMode = &v2
		}
	}
	if v1, ok := d.GetOk("qtm_buf_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("qtm_buf_mode", sv)
				diags = append(diags, e)
			}
			obj.QtmBufMode = &v2
		}
	}
	if v1, ok := d.GetOk("rdp_offload"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.6") {
				e := utils.AttributeVersionWarning("rdp_offload", sv)
				diags = append(diags, e)
			}
			obj.RdpOffload = &v2
		}
	}
	if v1, ok := d.GetOk("recover_np6_link"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.6") {
				e := utils.AttributeVersionWarning("recover_np6_link", sv)
				diags = append(diags, e)
			}
			obj.RecoverNp6Link = &v2
		}
	}
	if v1, ok := d.GetOk("session_acct_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
				e := utils.AttributeVersionWarning("session_acct_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SessionAcctInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("shaping_stats"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.8", "") {
				e := utils.AttributeVersionWarning("shaping_stats", sv)
				diags = append(diags, e)
			}
			obj.ShapingStats = &v2
		}
	}
	if v1, ok := d.GetOk("sse_backpressure"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.6") {
				e := utils.AttributeVersionWarning("sse_backpressure", sv)
				diags = append(diags, e)
			}
			obj.SseBackpressure = &v2
		}
	}
	if v1, ok := d.GetOk("strip_clear_text_padding"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.6") {
				e := utils.AttributeVersionWarning("strip_clear_text_padding", sv)
				diags = append(diags, e)
			}
			obj.StripClearTextPadding = &v2
		}
	}
	if v1, ok := d.GetOk("strip_esp_padding"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.6") {
				e := utils.AttributeVersionWarning("strip_esp_padding", sv)
				diags = append(diags, e)
			}
			obj.StripEspPadding = &v2
		}
	}
	if v1, ok := d.GetOk("sw_np_bandwidth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.6") {
				e := utils.AttributeVersionWarning("sw_np_bandwidth", sv)
				diags = append(diags, e)
			}
			obj.SwNpBandwidth = &v2
		}
	}
	if v1, ok := d.GetOk("tcp_rst_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
				e := utils.AttributeVersionWarning("tcp_rst_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TcpRstTimeout = &tmp
		}
	}
	if v, ok := d.GetOk("tcp_timeout_profile"); ok {
		if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
			e := utils.AttributeVersionWarning("tcp_timeout_profile", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemNpuTcpTimeoutProfile(d, v, "tcp_timeout_profile", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.TcpTimeoutProfile = t
		}
	} else if d.HasChange("tcp_timeout_profile") {
		old, new := d.GetChange("tcp_timeout_profile")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.TcpTimeoutProfile = &[]models.SystemNpuTcpTimeoutProfile{}
		}
	}
	if v, ok := d.GetOk("udp_timeout_profile"); ok {
		if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
			e := utils.AttributeVersionWarning("udp_timeout_profile", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemNpuUdpTimeoutProfile(d, v, "udp_timeout_profile", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.UdpTimeoutProfile = t
		}
	} else if d.HasChange("udp_timeout_profile") {
		old, new := d.GetChange("udp_timeout_profile")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.UdpTimeoutProfile = &[]models.SystemNpuUdpTimeoutProfile{}
		}
	}
	if v1, ok := d.GetOk("vlan_lookup_cache"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.6", "v7.0.2") {
				e := utils.AttributeVersionWarning("vlan_lookup_cache", sv)
				diags = append(diags, e)
			}
			obj.VlanLookupCache = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemNpu(d *schema.ResourceData, sv string) (*models.SystemNpu, diag.Diagnostics) {
	obj := models.SystemNpu{}
	diags := diag.Diagnostics{}

	obj.DosOptions = &models.SystemNpuDosOptions{}
	obj.DswDtsProfile = &[]models.SystemNpuDswDtsProfile{}
	obj.DswQueueDtsProfile = &[]models.SystemNpuDswQueueDtsProfile{}
	obj.FpAnomaly = &models.SystemNpuFpAnomaly{}
	obj.Hpe = &models.SystemNpuHpe{}
	obj.InboundDscpCopyPort = &[]models.SystemNpuInboundDscpCopyPort{}
	obj.IpReassembly = &models.SystemNpuIpReassembly{}
	obj.IsfNpQueues = &models.SystemNpuIsfNpQueues{}
	obj.NpQueues = &models.SystemNpuNpQueues{}
	obj.PortCpuMap = &[]models.SystemNpuPortCpuMap{}
	obj.PortNpuMap = &[]models.SystemNpuPortNpuMap{}
	obj.PriorityProtocol = &models.SystemNpuPriorityProtocol{}
	obj.TcpTimeoutProfile = &[]models.SystemNpuTcpTimeoutProfile{}
	obj.UdpTimeoutProfile = &[]models.SystemNpuUdpTimeoutProfile{}

	return &obj, diags
}
