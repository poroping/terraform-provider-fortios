// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiSwitch devices that are managed by this FortiGate.

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

func resourceSwitchControllerManagedSwitch() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiSwitch devices that are managed by this FortiGate.",

		CreateContext: resourceSwitchControllerManagedSwitchCreate,
		ReadContext:   resourceSwitchControllerManagedSwitchRead,
		UpdateContext: resourceSwitchControllerManagedSwitchUpdate,
		DeleteContext: resourceSwitchControllerManagedSwitchDelete,

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
			"802_1x_settings": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit FortiSwitch 802.1X global settings.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"link_down_auth": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"set-unauth", "no-action"}, false),

							Description: "Authentication state to set if a link is down.",
							Optional:    true,
							Computed:    true,
						},
						"local_override": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to override global 802.1X settings on individual FortiSwitches.",
							Optional:    true,
							Computed:    true,
						},
						"max_reauth_attempt": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 15),

							Description: "Maximum number of authentication attempts (0 - 15, default = 3).",
							Optional:    true,
							Computed:    true,
						},
						"reauth_period": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1440),

							Description: "Reauthentication time interval (1 - 1440 min, default = 60, 0 = disable).",
							Optional:    true,
							Computed:    true,
						},
						"tx_period": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(4, 60),

							Description: "802.1X Tx period (seconds, default=30).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"access_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "FortiSwitch access profile.",
				Optional:    true,
				Computed:    true,
			},
			"custom_command": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit FortiSwitch commands to be pushed to this FortiSwitch device upon rebooting the FortiGate switch controller or the FortiSwitch.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command_entry": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "List of FortiSwitch commands.",
							Optional:    true,
							Computed:    true,
						},
						"command_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Names of commands to be pushed to this FortiSwitch device, as configured under config switch-controller custom-command.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"delayed_restart_trigger": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Delayed restart triggered for this FortiSwitch.",
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Description.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_server_access_list": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"global", "enable", "disable"}, false),

				Description: "DHCP snooping server access list.",
				Optional:    true,
				Computed:    true,
			},
			"directly_connected": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1),

				Description: "Directly connected FortiSwitch.",
				Optional:    true,
				Computed:    true,
			},
			"dynamic_capability": {
				Type: schema.TypeString,

				Description: "List of features this FortiSwitch supports (not configurable) that is sent to the FortiGate device for subsequent configuration initiated by the FortiGate device.",
				Optional:    true,
				Computed:    true,
			},
			"dynamically_discovered": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1),

				Description: "Dynamically discovered FortiSwitch.",
				Optional:    true,
				Computed:    true,
			},
			"firmware_provision": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable provisioning of firmware to FortiSwitches on join connection.",
				Optional:    true,
				Computed:    true,
			},
			"firmware_provision_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Firmware version to provision to this FortiSwitch on bootup (major.minor.build, i.e. 6.2.1234).",
				Optional:    true,
				Computed:    true,
			},
			"flow_identity": {
				Type: schema.TypeString,

				Description: "Flow-tracking netflow ipfix switch identity in hex format(00000000-FFFFFFFF default=0).",
				Optional:    true,
				Computed:    true,
			},
			"fsw_wan1_admin": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"discovered", "disable", "enable"}, false),

				Description: "FortiSwitch WAN1 admin status; enable to authorize the FortiSwitch as a managed switch.",
				Optional:    true,
				Computed:    true,
			},
			"fsw_wan1_peer": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Fortiswitch WAN1 peer port.",
				Optional:    true,
				Computed:    true,
			},
			"igmp_snooping": {
				Type:        schema.TypeList,
				Description: "Configure FortiSwitch IGMP snooping global settings.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aging_time": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(15, 3600),

							Description: "Maximum time to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).",
							Optional:    true,
							Computed:    true,
						},
						"flood_unknown_multicast": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable unknown multicast flooding.",
							Optional:    true,
							Computed:    true,
						},
						"local_override": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable overriding the global IGMP snooping configuration.",
							Optional:    true,
							Computed:    true,
						},
						"vlans": {
							Type:        schema.TypeList,
							Description: "Configure IGMP snooping VLAN.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"proxy": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"disable", "enable", "global"}, false),

										Description: "IGMP snooping proxy for the VLAN interface.",
										Optional:    true,
										Computed:    true,
									},
									"querier": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

										Description: "Enable/disable IGMP snooping querier for the VLAN interface.",
										Optional:    true,
										Computed:    true,
									},
									"querier_addr": {
										Type:         schema.TypeString,
										ValidateFunc: validation.IsIPv4Address,

										Description: "IGMP snooping querier address.",
										Optional:    true,
										Computed:    true,
									},
									"version": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(2, 3),

										Description: "IGMP snooping querier version.",
										Optional:    true,
										Computed:    true,
									},
									"vlan_name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 15),

										Description: "List of FortiSwitch VLANs.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"ip_source_guard": {
				Type:        schema.TypeList,
				Description: "IP source guard.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"binding_entry": {
							Type:        schema.TypeList,
							Description: "IP and MAC address configuration.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"entry_name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 16),

										Description: "Configure binding pair.",
										Optional:    true,
										Computed:    true,
									},
									"ip": {
										Type:         schema.TypeString,
										ValidateFunc: validation.IsIPv4Address,

										Description: "Source IP for this rule.",
										Optional:    true,
										Computed:    true,
									},
									"mac": {
										Type: schema.TypeString,

										Description: "MAC address for this rule.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"description": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Description.",
							Optional:    true,
							Computed:    true,
						},
						"port": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Ingress interface to which source guard is bound.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"l3_discovered": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1),

				Description: "Layer 3 management discovered.",
				Optional:    true,
				Computed:    true,
			},
			"max_allowed_trunk_members": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "FortiSwitch maximum allowed trunk members.",
				Optional:    true,
				Computed:    true,
			},
			"mclag_igmp_snooping_aware": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable MCLAG IGMP-snooping awareness.",
				Optional:    true,
				Computed:    true,
			},
			"mirror": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit FortiSwitch packet mirror.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Destination port.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Mirror name.",
							Optional:    true,
							Computed:    true,
						},
						"src_egress": {
							Type:        schema.TypeList,
							Description: "Source egress interfaces.",
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
						"src_ingress": {
							Type:        schema.TypeList,
							Description: "Source ingress interfaces.",
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
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"active", "inactive"}, false),

							Description: "Active/inactive mirror configuration.",
							Optional:    true,
							Computed:    true,
						},
						"switching_packet": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable switching functionality when mirroring.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Managed-switch name.",
				Optional:    true,
				Computed:    true,
			},
			"override_snmp_community": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable overriding the global SNMP communities.",
				Optional:    true,
				Computed:    true,
			},
			"override_snmp_sysinfo": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable overriding the global SNMP system information.",
				Optional:    true,
				Computed:    true,
			},
			"override_snmp_trap_threshold": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable overriding the global SNMP trap threshold values.",
				Optional:    true,
				Computed:    true,
			},
			"override_snmp_user": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable overriding the global SNMP users.",
				Optional:    true,
				Computed:    true,
			},
			"owner_vdom": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "VDOM which owner of port belongs to.",
				Optional:    true,
				Computed:    true,
			},
			"poe_detection_type": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "PoE detection type for FortiSwitch.",
				Optional:    true,
				Computed:    true,
			},
			"poe_lldp_detection": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable PoE LLDP detection.",
				Optional:    true,
				Computed:    true,
			},
			"poe_pre_standard_detection": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable PoE pre-standard detection.",
				Optional:    true,
				Computed:    true,
			},
			"ports": {
				Type:        schema.TypeList,
				Description: "Managed-switch port list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"dynamic", "nac", "static"}, false),

							Description: "Access mode of the port.",
							Optional:    true,
							Computed:    true,
						},
						"aggregator_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bandwidth", "count"}, false),

							Description: "LACP member select mode.",
							Optional:    true,
							Computed:    true,
						},
						"allowed_vlans": {
							Type:        schema.TypeList,
							Description: "Configure switch port tagged vlans",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vlan_name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "VLAN name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"allowed_vlans_all": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable all defined vlans on this port.",
							Optional:    true,
							Computed:    true,
						},
						"arp_inspection_trust": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"untrusted", "trusted"}, false),

							Description: "Trusted or untrusted dynamic ARP inspection.",
							Optional:    true,
							Computed:    true,
						},
						"bundle": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable Link Aggregation Group (LAG) bundling for non-FortiLink interfaces.",
							Optional:    true,
							Computed:    true,
						},
						"description": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Description for port.",
							Optional:    true,
							Computed:    true,
						},
						"dhcp_snoop_option82_trust": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable allowance of DHCP with option-82 on untrusted interface.",
							Optional:    true,
							Computed:    true,
						},
						"dhcp_snooping": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"untrusted", "trusted"}, false),

							Description: "Trusted or untrusted DHCP-snooping interface.",
							Optional:    true,
							Computed:    true,
						},
						"discard_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "all-untagged", "all-tagged"}, false),

							Description: "Configure discard mode for port.",
							Optional:    true,
							Computed:    true,
						},
						"edge_port": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable this interface as an edge port, bridging connections between workstations and/or computers.",
							Optional:    true,
							Computed:    true,
						},
						"export_tags": {
							Type:        schema.TypeList,
							Description: "Configure export tag(s) for FortiSwitch port when exported to a virtual port pool.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag_name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),

										Description: "FortiSwitch port tag name when exported to a virtual port pool.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"export_to": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),

							Description: "Export managed-switch port to a tenant VDOM.",
							Optional:    true,
							Computed:    true,
						},
						"export_to_pool": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Switch controller export port to pool-list.",
							Optional:    true,
							Computed:    true,
						},
						"fec_capable": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),

							Description: "FEC capable.",
							Optional:    true,
							Computed:    true,
						},
						"fec_state": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disabled", "cl74", "cl91"}, false),

							Description: "State of forward error correction.",
							Optional:    true,
							Computed:    true,
						},
						"fgt_peer_device_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 16),

							Description: "FGT peer device name.",
							Optional:    true,
							Computed:    true,
						},
						"fgt_peer_port_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "FGT peer port name.",
							Optional:    true,
							Computed:    true,
						},
						"fiber_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),

							Description: "Fiber-port.",
							Optional:    true,
							Computed:    true,
						},
						"flags": {
							Type: schema.TypeInt,

							Description: "Port properties flags.",
							Optional:    true,
							Computed:    true,
						},
						"flow_control": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "tx", "rx", "both"}, false),

							Description: "Flow control direction.",
							Optional:    true,
							Computed:    true,
						},
						"fortilink_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),

							Description: "FortiLink uplink port.",
							Optional:    true,
							Computed:    true,
						},
						"igmp_snooping": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Set IGMP snooping mode for the physical port interface.",
							Optional:    true,
							Computed:    true,
						},
						"igmps_flood_reports": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable flooding of IGMP reports to this interface when igmp-snooping enabled.",
							Optional:    true,
							Computed:    true,
						},
						"igmps_flood_traffic": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable flooding of IGMP snooping traffic to this interface.",
							Optional:    true,
							Computed:    true,
						},
						"interface_tags": {
							Type:        schema.TypeList,
							Description: "Tag(s) associated with the interface for various features including virtual port pool, dynamic port policy.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag_name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),

										Description: "FortiSwitch port tag name when exported to a virtual port pool or matched to dynamic port policy.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"ip_source_guard": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable IP source guard.",
							Optional:    true,
							Computed:    true,
						},
						"isl_local_trunk_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "ISL local trunk name.",
							Optional:    true,
							Computed:    true,
						},
						"isl_peer_device_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 16),

							Description: "ISL peer device name.",
							Optional:    true,
							Computed:    true,
						},
						"isl_peer_port_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "ISL peer port name.",
							Optional:    true,
							Computed:    true,
						},
						"lacp_speed": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"slow", "fast"}, false),

							Description: "end Link Aggregation Control Protocol (LACP) messages every 30 seconds (slow) or every second (fast).",
							Optional:    true,
							Computed:    true,
						},
						"learning_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 128),

							Description: "Limit the number of dynamic MAC addresses on this Port (1 - 128, 0 = no limit, default).",
							Optional:    true,
							Computed:    true,
						},
						"lldp_profile": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "LLDP port TLV profile.",
							Optional:    true,
							Computed:    true,
						},
						"lldp_status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "rx-only", "tx-only", "tx-rx"}, false),

							Description: "LLDP transmit and receive status.",
							Optional:    true,
							Computed:    true,
						},
						"loop_guard": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enabled", "disabled"}, false),

							Description: "Enable/disable loop-guard on this interface, an STP optimization used to prevent network loops.",
							Optional:    true,
							Computed:    true,
						},
						"loop_guard_timeout": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 120),

							Description: "Loop-guard timeout (0 - 120 min, default = 45).",
							Optional:    true,
							Computed:    true,
						},
						"mac_addr": {
							Type: schema.TypeString,

							Description: "Port/Trunk MAC.",
							Optional:    true,
							Computed:    true,
						},
						"matched_dpp_intf_tags": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Matched interface tags in the dynamic port policy.",
							Optional:    true,
							Computed:    true,
						},
						"matched_dpp_policy": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Matched child policy in the dynamic port policy.",
							Optional:    true,
							Computed:    true,
						},
						"max_bundle": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 24),

							Description: "Maximum size of LAG bundle (1 - 24, default = 24)",
							Optional:    true,
							Computed:    true,
						},
						"mclag": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable multi-chassis link aggregation (MCLAG).",
							Optional:    true,
							Computed:    true,
						},
						"mclag_icl_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),

							Description: "MCLAG-ICL port.",
							Optional:    true,
							Computed:    true,
						},
						"media_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),

							Description: "Media type.",
							Optional:    true,
							Computed:    true,
						},
						"member_withdrawal_behavior": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"forward", "block"}, false),

							Description: "Port behavior after it withdraws because of loss of control packets.",
							Optional:    true,
							Computed:    true,
						},
						"members": {
							Type:        schema.TypeList,
							Description: "Aggregated LAG bundle interfaces.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"member_name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Interface name from available options.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"min_bundle": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 24),

							Description: "Minimum size of LAG bundle (1 - 24, default = 1)",
							Optional:    true,
							Computed:    true,
						},
						"mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"static", "lacp-passive", "lacp-active"}, false),

							Description: "LACP mode: ignore and do not send control messages, or negotiate 802.3ad aggregation passively or actively.",
							Optional:    true,
							Computed:    true,
						},
						"p2p_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),

							Description: "General peer to peer tunnel port.",
							Optional:    true,
							Computed:    true,
						},
						"packet_sample_rate": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 99999),

							Description: "Packet sampling rate (0 - 99999 p/sec).",
							Optional:    true,
							Computed:    true,
						},
						"packet_sampler": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enabled", "disabled"}, false),

							Description: "Enable/disable packet sampling on this interface.",
							Optional:    true,
							Computed:    true,
						},
						"pause_meter": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(128, 2147483647),

							Description: "Configure ingress pause metering rate, in kbps (default = 0, disabled).",
							Optional:    true,
							Computed:    true,
						},
						"pause_meter_resume": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"75%", "50%", "25%"}, false),

							Description: "Resume threshold for resuming traffic on ingress port.",
							Optional:    true,
							Computed:    true,
						},
						"poe_capable": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),

							Description: "PoE capable.",
							Optional:    true,
							Computed:    true,
						},
						"poe_max_power": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "PoE maximum power.",
							Optional:    true,
							Computed:    true,
						},
						"poe_pre_standard_detection": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable PoE pre-standard detection.",
							Optional:    true,
							Computed:    true,
						},
						"poe_standard": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "PoE standard supported.",
							Optional:    true,
							Computed:    true,
						},
						"poe_status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable PoE status.",
							Optional:    true,
							Computed:    true,
						},
						"port_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Switch port name.",
							Optional:    true,
							Computed:    true,
						},
						"port_number": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 64),

							Description: "Port number.",
							Optional:    true,
							Computed:    true,
						},
						"port_owner": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Switch port name.",
							Optional:    true,
							Computed:    true,
						},
						"port_policy": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Switch controller dynamic port policy from available options.",
							Optional:    true,
							Computed:    true,
						},
						"port_prefix_type": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),

							Description: "Port prefix type.",
							Optional:    true,
							Computed:    true,
						},
						"port_security_policy": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),

							Description: "Switch controller authentication policy to apply to this managed switch from available options.",
							Optional:    true,
							Computed:    true,
						},
						"port_selection_criteria": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"src-mac", "dst-mac", "src-dst-mac", "src-ip", "dst-ip", "src-dst-ip"}, false),

							Description: "Algorithm for aggregate port selection.",
							Optional:    true,
							Computed:    true,
						},
						"ptp_policy": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "PTP policy configuration.",
							Optional:    true,
							Computed:    true,
						},
						"qos_policy": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Switch controller QoS policy from available options.",
							Optional:    true,
							Computed:    true,
						},
						"rpvst_port": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disabled", "enabled"}, false),

							Description: "Enable/disable inter-operability with rapid PVST on this interface.",
							Optional:    true,
							Computed:    true,
						},
						"sample_direction": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"tx", "rx", "both"}, false),

							Description: "Packet sampling direction.",
							Optional:    true,
							Computed:    true,
						},
						"sflow_counter_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "sFlow sampling counter polling interval (0 - 255 sec).",
							Optional:    true,
							Computed:    true,
						},
						"speed": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"10half", "10full", "100half", "100full", "1000auto", "1000fiber", "1000full", "10000", "40000", "auto", "auto-module", "100FX-half", "100FX-full", "100000full", "2500auto", "25000full", "50000full", "10000cr", "10000sr", "100000sr4", "100000cr4", "25000cr4", "25000sr4", "5000full"}, false),

							Description: "Switch port speed; default and available settings depend on hardware.",
							Optional:    true,
							Computed:    true,
						},
						"stacking_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),

							Description: "Stacking port.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"up", "down"}, false),

							Description: "Switch port admin status: up or down.",
							Optional:    true,
							Computed:    true,
						},
						"sticky_mac": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable or disable sticky-mac on the interface.",
							Optional:    true,
							Computed:    true,
						},
						"storm_control_policy": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Switch controller storm control policy from available options.",
							Optional:    true,
							Computed:    true,
						},
						"stp_bpdu_guard": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enabled", "disabled"}, false),

							Description: "Enable/disable STP BPDU guard on this interface.",
							Optional:    true,
							Computed:    true,
						},
						"stp_bpdu_guard_timeout": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 120),

							Description: "BPDU Guard disabling protection (0 - 120 min).",
							Optional:    true,
							Computed:    true,
						},
						"stp_root_guard": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enabled", "disabled"}, false),

							Description: "Enable/disable STP root guard on this interface.",
							Optional:    true,
							Computed:    true,
						},
						"stp_state": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enabled", "disabled"}, false),

							Description: "Enable/disable Spanning Tree Protocol (STP) on this interface.",
							Optional:    true,
							Computed:    true,
						},
						"switch_id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 16),

							Description: "Switch id.",
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"physical", "trunk"}, false),

							Description: "Interface type: physical or trunk port.",
							Optional:    true,
							Computed:    true,
						},
						"untagged_vlans": {
							Type:        schema.TypeList,
							Description: "Configure switch port untagged vlans",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vlan_name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "VLAN name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"vlan": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Assign switch ports to a VLAN.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"pre_provisioned": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Pre-provisioned managed switch.",
				Optional:    true,
				Computed:    true,
			},
			"qos_drop_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"taildrop", "random-early-detection"}, false),

				Description: "Set QoS drop-policy.",
				Optional:    true,
				Computed:    true,
			},
			"qos_red_probability": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),

				Description: "Set QoS RED/WRED drop probability.",
				Optional:    true,
				Computed:    true,
			},
			"remote_log": {
				Type:        schema.TypeList,
				Description: "Configure logging by FortiSwitch device to a remote syslog server.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"csv": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable comma-separated value (CSV) strings.",
							Optional:    true,
							Computed:    true,
						},
						"facility": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"kernel", "user", "mail", "daemon", "auth", "syslog", "lpr", "news", "uucp", "cron", "authpriv", "ftp", "ntp", "audit", "alert", "clock", "local0", "local1", "local2", "local3", "local4", "local5", "local6", "local7"}, false),

							Description: "Facility to log to remote syslog server.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Remote log name.",
							Optional:    true,
							Computed:    true,
						},
						"port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Remote syslog server listening port.",
							Optional:    true,
							Computed:    true,
						},
						"server": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "IPv4 address of the remote syslog server.",
							Optional:    true,
							Computed:    true,
						},
						"severity": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"emergency", "alert", "critical", "error", "warning", "notification", "information", "debug"}, false),

							Description: "Severity of logs to be transferred to remote log server.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable logging by FortiSwitch device to a remote syslog server.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"snmp_community": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit Simple Network Management Protocol (SNMP) communities.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"events": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "SNMP notifications (traps) to send.",
							Optional:         true,
							Computed:         true,
						},
						"hosts": {
							Type:        schema.TypeList,
							Description: "Configure IPv4 SNMP managers (hosts).",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type: schema.TypeInt,

										Description: "Host entry ID.",
										Optional:    true,
										Computed:    true,
									},
									"ip": {
										Type: schema.TypeString,

										Description: "IPv4 address of the SNMP manager (host).",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type: schema.TypeInt,

							Description: "SNMP community ID.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "SNMP community name.",
							Optional:    true,
							Computed:    true,
						},
						"query_v1_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "SNMP v1 query port (default = 161).",
							Optional:    true,
							Computed:    true,
						},
						"query_v1_status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable SNMP v1 queries.",
							Optional:    true,
							Computed:    true,
						},
						"query_v2c_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "SNMP v2c query port (default = 161).",
							Optional:    true,
							Computed:    true,
						},
						"query_v2c_status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable SNMP v2c queries.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable this SNMP community.",
							Optional:    true,
							Computed:    true,
						},
						"trap_v1_lport": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "SNMP v2c trap local port (default = 162).",
							Optional:    true,
							Computed:    true,
						},
						"trap_v1_rport": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "SNMP v2c trap remote port (default = 162).",
							Optional:    true,
							Computed:    true,
						},
						"trap_v1_status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable SNMP v1 traps.",
							Optional:    true,
							Computed:    true,
						},
						"trap_v2c_lport": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "SNMP v2c trap local port (default = 162).",
							Optional:    true,
							Computed:    true,
						},
						"trap_v2c_rport": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "SNMP v2c trap remote port (default = 162).",
							Optional:    true,
							Computed:    true,
						},
						"trap_v2c_status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable SNMP v2c traps.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"snmp_sysinfo": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit Simple Network Management Protocol (SNMP) system info.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"contact_info": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Contact information.",
							Optional:    true,
							Computed:    true,
						},
						"description": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "System description.",
							Optional:    true,
							Computed:    true,
						},
						"engine_id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 24),

							Description: "Local SNMP engine ID string (max 24 char).",
							Optional:    true,
							Computed:    true,
						},
						"location": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "System location.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable SNMP.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"snmp_trap_threshold": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit Simple Network Management Protocol (SNMP) trap threshold values.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trap_high_cpu_threshold": {
							Type: schema.TypeInt,

							Description: "CPU usage when trap is sent.",
							Optional:    true,
							Computed:    true,
						},
						"trap_log_full_threshold": {
							Type: schema.TypeInt,

							Description: "Log disk usage when trap is sent.",
							Optional:    true,
							Computed:    true,
						},
						"trap_low_memory_threshold": {
							Type: schema.TypeInt,

							Description: "Memory usage when trap is sent.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"snmp_user": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit Simple Network Management Protocol (SNMP) users.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_proto": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"md5", "sha1", "sha224", "sha256", "sha384", "sha512"}, false),

							Description: "Authentication protocol.",
							Optional:    true,
							Computed:    true,
						},
						"auth_pwd": {
							Type: schema.TypeString,

							Description: "Password for authentication protocol.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32),

							Description: "SNMP user name.",
							Optional:    true,
							Computed:    true,
						},
						"priv_proto": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"aes128", "aes192", "aes192c", "aes256", "aes256c", "des"}, false),

							Description: "Privacy (encryption) protocol.",
							Optional:    true,
							Computed:    true,
						},
						"priv_pwd": {
							Type: schema.TypeString,

							Description: "Password for privacy (encryption) protocol.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"queries": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable SNMP queries for this user.",
							Optional:    true,
							Computed:    true,
						},
						"query_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "SNMPv3 query port (default = 161).",
							Optional:    true,
							Computed:    true,
						},
						"security_level": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"no-auth-no-priv", "auth-no-priv", "auth-priv"}, false),

							Description: "Security level for message authentication and encryption.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"staged_image_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Staged image version for FortiSwitch.",
				Optional:    true,
				Computed:    true,
			},
			"static_mac": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit FortiSwitch Static and Sticky MAC.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Description.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Id",
							Optional:    true,
							Computed:    true,
						},
						"interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
						"mac": {
							Type: schema.TypeString,

							Description: "MAC address.",
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"static", "sticky"}, false),

							Description: "Type.",
							Optional:    true,
							Computed:    true,
						},
						"vlan": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Vlan.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"storm_control": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit FortiSwitch storm control for measuring traffic activity using data rates to prevent traffic disruption.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"broadcast": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable storm control to drop broadcast traffic.",
							Optional:    true,
							Computed:    true,
						},
						"local_override": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to override global FortiSwitch storm control settings for this FortiSwitch.",
							Optional:    true,
							Computed:    true,
						},
						"rate": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 10000000),

							Description: "Rate in packets per second at which storm traffic is controlled (1 - 10000000, default = 500). Storm control drops excess traffic data rates beyond this threshold.",
							Optional:    true,
							Computed:    true,
						},
						"unknown_multicast": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable storm control to drop unknown multicast traffic.",
							Optional:    true,
							Computed:    true,
						},
						"unknown_unicast": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable storm control to drop unknown unicast traffic.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"stp_instance": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit Spanning Tree Protocol (STP) instances.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 2),

							Description: "Instance ID.",
							Optional:    true,
							Computed:    true,
						},
						"priority": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"0", "4096", "8192", "12288", "16384", "20480", "24576", "28672", "32768", "36864", "40960", "45056", "49152", "53248", "57344", "61440"}, false),

							Description: "Priority.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"stp_settings": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit Spanning Tree Protocol (STP) settings used to prevent bridge loops.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"forward_time": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(4, 30),

							Description: "Period of time a port is in listening and learning state (4 - 30 sec, default = 15).",
							Optional:    true,
							Computed:    true,
						},
						"hello_time": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 10),

							Description: "Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).",
							Optional:    true,
							Computed:    true,
						},
						"local_override": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to configure local STP settings that override global STP settings.",
							Optional:    true,
							Computed:    true,
						},
						"max_age": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(6, 40),

							Description: "Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).",
							Optional:    true,
							Computed:    true,
						},
						"max_hops": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 40),

							Description: "Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),

							Description: "Name of local STP settings configuration.",
							Optional:    true,
							Computed:    true,
						},
						"pending_timer": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 15),

							Description: "Pending time (1 - 15 sec, default = 4).",
							Optional:    true,
							Computed:    true,
						},
						"revision": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "STP revision number (0 - 65535).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"switch_device_tag": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),

				Description: "User definable label/tag.",
				Optional:    true,
				Computed:    true,
			},
			"switch_dhcp_opt43_key": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "DHCP option43 key.",
				Optional:    true,
				Computed:    true,
			},
			"switch_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 16),

				Description: "Managed-switch id.",
				ForceNew:    true,
				Required:    true,
			},
			"switch_log": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit FortiSwitch logging settings (logs are transferred to and inserted into the FortiGate event log).",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"local_override": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to configure local logging settings that override global logging settings.",
							Optional:    true,
							Computed:    true,
						},
						"severity": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"emergency", "alert", "critical", "error", "warning", "notification", "information", "debug"}, false),

							Description: "Severity of FortiSwitch logs that are added to the FortiGate event log.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable adding FortiSwitch logs to the FortiGate event log.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"switch_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "FortiSwitch profile.",
				Optional:    true,
				Computed:    true,
			},
			"tdr_supported": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "TDR supported.",
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"virtual", "physical"}, false),

				Description: "Indication of switch type, physical or virtual.",
				Optional:    true,
				Computed:    true,
			},
			"version": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "FortiSwitch version.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSwitchControllerManagedSwitchCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "switch_id"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating SwitchControllerManagedSwitch resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSwitchControllerManagedSwitch(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSwitchControllerManagedSwitch(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerManagedSwitch")
	}

	return resourceSwitchControllerManagedSwitchRead(ctx, d, meta)
}

func resourceSwitchControllerManagedSwitchUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerManagedSwitch(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSwitchControllerManagedSwitch(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SwitchControllerManagedSwitch resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerManagedSwitch")
	}

	return resourceSwitchControllerManagedSwitchRead(ctx, d, meta)
}

func resourceSwitchControllerManagedSwitchDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSwitchControllerManagedSwitch(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SwitchControllerManagedSwitch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerManagedSwitchRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerManagedSwitch(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerManagedSwitch resource: %v", err)
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

	diags := refreshObjectSwitchControllerManagedSwitch(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSwitchControllerManagedSwitch8021XSettings(v *[]models.SwitchControllerManagedSwitch8021XSettings, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.LinkDownAuth; tmp != nil {
				v["link_down_auth"] = *tmp
			}

			if tmp := cfg.LocalOverride; tmp != nil {
				v["local_override"] = *tmp
			}

			if tmp := cfg.MaxReauthAttempt; tmp != nil {
				v["max_reauth_attempt"] = *tmp
			}

			if tmp := cfg.ReauthPeriod; tmp != nil {
				v["reauth_period"] = *tmp
			}

			if tmp := cfg.TxPeriod; tmp != nil {
				v["tx_period"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSwitchControllerManagedSwitchCustomCommand(v *[]models.SwitchControllerManagedSwitchCustomCommand, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.CommandEntry; tmp != nil {
				v["command_entry"] = *tmp
			}

			if tmp := cfg.CommandName; tmp != nil {
				v["command_name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "command_entry")
	}

	return flat
}

func flattenSwitchControllerManagedSwitchIgmpSnooping(v *[]models.SwitchControllerManagedSwitchIgmpSnooping, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AgingTime; tmp != nil {
				v["aging_time"] = *tmp
			}

			if tmp := cfg.FloodUnknownMulticast; tmp != nil {
				v["flood_unknown_multicast"] = *tmp
			}

			if tmp := cfg.LocalOverride; tmp != nil {
				v["local_override"] = *tmp
			}

			if tmp := cfg.Vlans; tmp != nil {
				v["vlans"] = flattenSwitchControllerManagedSwitchIgmpSnoopingVlans(tmp, sort)
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlans(v *[]models.SwitchControllerManagedSwitchIgmpSnoopingVlans, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Proxy; tmp != nil {
				v["proxy"] = *tmp
			}

			if tmp := cfg.Querier; tmp != nil {
				v["querier"] = *tmp
			}

			if tmp := cfg.QuerierAddr; tmp != nil {
				v["querier_addr"] = *tmp
			}

			if tmp := cfg.Version; tmp != nil {
				v["version"] = *tmp
			}

			if tmp := cfg.VlanName; tmp != nil {
				v["vlan_name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "vlan_name")
	}

	return flat
}

func flattenSwitchControllerManagedSwitchIpSourceGuard(v *[]models.SwitchControllerManagedSwitchIpSourceGuard, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.BindingEntry; tmp != nil {
				v["binding_entry"] = flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntry(tmp, sort)
			}

			if tmp := cfg.Description; tmp != nil {
				v["description"] = *tmp
			}

			if tmp := cfg.Port; tmp != nil {
				v["port"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "port")
	}

	return flat
}

func flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntry(v *[]models.SwitchControllerManagedSwitchIpSourceGuardBindingEntry, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.EntryName; tmp != nil {
				v["entry_name"] = *tmp
			}

			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			if tmp := cfg.Mac; tmp != nil {
				v["mac"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "entry_name")
	}

	return flat
}

func flattenSwitchControllerManagedSwitchMirror(v *[]models.SwitchControllerManagedSwitchMirror, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Dst; tmp != nil {
				v["dst"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.SrcEgress; tmp != nil {
				v["src_egress"] = flattenSwitchControllerManagedSwitchMirrorSrcEgress(tmp, sort)
			}

			if tmp := cfg.SrcIngress; tmp != nil {
				v["src_ingress"] = flattenSwitchControllerManagedSwitchMirrorSrcIngress(tmp, sort)
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.SwitchingPacket; tmp != nil {
				v["switching_packet"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSwitchControllerManagedSwitchMirrorSrcEgress(v *[]models.SwitchControllerManagedSwitchMirrorSrcEgress, sort bool) interface{} {
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

func flattenSwitchControllerManagedSwitchMirrorSrcIngress(v *[]models.SwitchControllerManagedSwitchMirrorSrcIngress, sort bool) interface{} {
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

func flattenSwitchControllerManagedSwitchPorts(v *[]models.SwitchControllerManagedSwitchPorts, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AccessMode; tmp != nil {
				v["access_mode"] = *tmp
			}

			if tmp := cfg.AggregatorMode; tmp != nil {
				v["aggregator_mode"] = *tmp
			}

			if tmp := cfg.AllowedVlans; tmp != nil {
				v["allowed_vlans"] = flattenSwitchControllerManagedSwitchPortsAllowedVlans(tmp, sort)
			}

			if tmp := cfg.AllowedVlansAll; tmp != nil {
				v["allowed_vlans_all"] = *tmp
			}

			if tmp := cfg.ArpInspectionTrust; tmp != nil {
				v["arp_inspection_trust"] = *tmp
			}

			if tmp := cfg.Bundle; tmp != nil {
				v["bundle"] = *tmp
			}

			if tmp := cfg.Description; tmp != nil {
				v["description"] = *tmp
			}

			if tmp := cfg.DhcpSnoopOption82Trust; tmp != nil {
				v["dhcp_snoop_option82_trust"] = *tmp
			}

			if tmp := cfg.DhcpSnooping; tmp != nil {
				v["dhcp_snooping"] = *tmp
			}

			if tmp := cfg.DiscardMode; tmp != nil {
				v["discard_mode"] = *tmp
			}

			if tmp := cfg.EdgePort; tmp != nil {
				v["edge_port"] = *tmp
			}

			if tmp := cfg.ExportTags; tmp != nil {
				v["export_tags"] = flattenSwitchControllerManagedSwitchPortsExportTags(tmp, sort)
			}

			if tmp := cfg.ExportTo; tmp != nil {
				v["export_to"] = *tmp
			}

			if tmp := cfg.ExportToPool; tmp != nil {
				v["export_to_pool"] = *tmp
			}

			if tmp := cfg.FecCapable; tmp != nil {
				v["fec_capable"] = *tmp
			}

			if tmp := cfg.FecState; tmp != nil {
				v["fec_state"] = *tmp
			}

			if tmp := cfg.FgtPeerDeviceName; tmp != nil {
				v["fgt_peer_device_name"] = *tmp
			}

			if tmp := cfg.FgtPeerPortName; tmp != nil {
				v["fgt_peer_port_name"] = *tmp
			}

			if tmp := cfg.FiberPort; tmp != nil {
				v["fiber_port"] = *tmp
			}

			if tmp := cfg.Flags; tmp != nil {
				v["flags"] = *tmp
			}

			if tmp := cfg.FlowControl; tmp != nil {
				v["flow_control"] = *tmp
			}

			if tmp := cfg.FortilinkPort; tmp != nil {
				v["fortilink_port"] = *tmp
			}

			if tmp := cfg.IgmpSnooping; tmp != nil {
				v["igmp_snooping"] = *tmp
			}

			if tmp := cfg.IgmpsFloodReports; tmp != nil {
				v["igmps_flood_reports"] = *tmp
			}

			if tmp := cfg.IgmpsFloodTraffic; tmp != nil {
				v["igmps_flood_traffic"] = *tmp
			}

			if tmp := cfg.InterfaceTags; tmp != nil {
				v["interface_tags"] = flattenSwitchControllerManagedSwitchPortsInterfaceTags(tmp, sort)
			}

			if tmp := cfg.IpSourceGuard; tmp != nil {
				v["ip_source_guard"] = *tmp
			}

			if tmp := cfg.IslLocalTrunkName; tmp != nil {
				v["isl_local_trunk_name"] = *tmp
			}

			if tmp := cfg.IslPeerDeviceName; tmp != nil {
				v["isl_peer_device_name"] = *tmp
			}

			if tmp := cfg.IslPeerPortName; tmp != nil {
				v["isl_peer_port_name"] = *tmp
			}

			if tmp := cfg.LacpSpeed; tmp != nil {
				v["lacp_speed"] = *tmp
			}

			if tmp := cfg.LearningLimit; tmp != nil {
				v["learning_limit"] = *tmp
			}

			if tmp := cfg.LldpProfile; tmp != nil {
				v["lldp_profile"] = *tmp
			}

			if tmp := cfg.LldpStatus; tmp != nil {
				v["lldp_status"] = *tmp
			}

			if tmp := cfg.LoopGuard; tmp != nil {
				v["loop_guard"] = *tmp
			}

			if tmp := cfg.LoopGuardTimeout; tmp != nil {
				v["loop_guard_timeout"] = *tmp
			}

			if tmp := cfg.MacAddr; tmp != nil {
				v["mac_addr"] = *tmp
			}

			if tmp := cfg.MatchedDppIntfTags; tmp != nil {
				v["matched_dpp_intf_tags"] = *tmp
			}

			if tmp := cfg.MatchedDppPolicy; tmp != nil {
				v["matched_dpp_policy"] = *tmp
			}

			if tmp := cfg.MaxBundle; tmp != nil {
				v["max_bundle"] = *tmp
			}

			if tmp := cfg.Mclag; tmp != nil {
				v["mclag"] = *tmp
			}

			if tmp := cfg.MclagIclPort; tmp != nil {
				v["mclag_icl_port"] = *tmp
			}

			if tmp := cfg.MediaType; tmp != nil {
				v["media_type"] = *tmp
			}

			if tmp := cfg.MemberWithdrawalBehavior; tmp != nil {
				v["member_withdrawal_behavior"] = *tmp
			}

			if tmp := cfg.Members; tmp != nil {
				v["members"] = flattenSwitchControllerManagedSwitchPortsMembers(tmp, sort)
			}

			if tmp := cfg.MinBundle; tmp != nil {
				v["min_bundle"] = *tmp
			}

			if tmp := cfg.Mode; tmp != nil {
				v["mode"] = *tmp
			}

			if tmp := cfg.P2pPort; tmp != nil {
				v["p2p_port"] = *tmp
			}

			if tmp := cfg.PacketSampleRate; tmp != nil {
				v["packet_sample_rate"] = *tmp
			}

			if tmp := cfg.PacketSampler; tmp != nil {
				v["packet_sampler"] = *tmp
			}

			if tmp := cfg.PauseMeter; tmp != nil {
				v["pause_meter"] = *tmp
			}

			if tmp := cfg.PauseMeterResume; tmp != nil {
				v["pause_meter_resume"] = *tmp
			}

			if tmp := cfg.PoeCapable; tmp != nil {
				v["poe_capable"] = *tmp
			}

			if tmp := cfg.PoeMaxPower; tmp != nil {
				v["poe_max_power"] = *tmp
			}

			if tmp := cfg.PoePreStandardDetection; tmp != nil {
				v["poe_pre_standard_detection"] = *tmp
			}

			if tmp := cfg.PoeStandard; tmp != nil {
				v["poe_standard"] = *tmp
			}

			if tmp := cfg.PoeStatus; tmp != nil {
				v["poe_status"] = *tmp
			}

			if tmp := cfg.PortName; tmp != nil {
				v["port_name"] = *tmp
			}

			if tmp := cfg.PortNumber; tmp != nil {
				v["port_number"] = *tmp
			}

			if tmp := cfg.PortOwner; tmp != nil {
				v["port_owner"] = *tmp
			}

			if tmp := cfg.PortPolicy; tmp != nil {
				v["port_policy"] = *tmp
			}

			if tmp := cfg.PortPrefixType; tmp != nil {
				v["port_prefix_type"] = *tmp
			}

			if tmp := cfg.PortSecurityPolicy; tmp != nil {
				v["port_security_policy"] = *tmp
			}

			if tmp := cfg.PortSelectionCriteria; tmp != nil {
				v["port_selection_criteria"] = *tmp
			}

			if tmp := cfg.PtpPolicy; tmp != nil {
				v["ptp_policy"] = *tmp
			}

			if tmp := cfg.QosPolicy; tmp != nil {
				v["qos_policy"] = *tmp
			}

			if tmp := cfg.RpvstPort; tmp != nil {
				v["rpvst_port"] = *tmp
			}

			if tmp := cfg.SampleDirection; tmp != nil {
				v["sample_direction"] = *tmp
			}

			if tmp := cfg.SflowCounterInterval; tmp != nil {
				v["sflow_counter_interval"] = *tmp
			}

			if tmp := cfg.Speed; tmp != nil {
				v["speed"] = *tmp
			}

			if tmp := cfg.StackingPort; tmp != nil {
				v["stacking_port"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.StickyMac; tmp != nil {
				v["sticky_mac"] = *tmp
			}

			if tmp := cfg.StormControlPolicy; tmp != nil {
				v["storm_control_policy"] = *tmp
			}

			if tmp := cfg.StpBpduGuard; tmp != nil {
				v["stp_bpdu_guard"] = *tmp
			}

			if tmp := cfg.StpBpduGuardTimeout; tmp != nil {
				v["stp_bpdu_guard_timeout"] = *tmp
			}

			if tmp := cfg.StpRootGuard; tmp != nil {
				v["stp_root_guard"] = *tmp
			}

			if tmp := cfg.StpState; tmp != nil {
				v["stp_state"] = *tmp
			}

			if tmp := cfg.SwitchId; tmp != nil {
				v["switch_id"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			if tmp := cfg.UntaggedVlans; tmp != nil {
				v["untagged_vlans"] = flattenSwitchControllerManagedSwitchPortsUntaggedVlans(tmp, sort)
			}

			if tmp := cfg.Vlan; tmp != nil {
				v["vlan"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "port_name")
	}

	return flat
}

func flattenSwitchControllerManagedSwitchPortsAllowedVlans(v *[]models.SwitchControllerManagedSwitchPortsAllowedVlans, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.VlanName; tmp != nil {
				v["vlan_name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "vlan_name")
	}

	return flat
}

func flattenSwitchControllerManagedSwitchPortsExportTags(v *[]models.SwitchControllerManagedSwitchPortsExportTags, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.TagName; tmp != nil {
				v["tag_name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "tag_name")
	}

	return flat
}

func flattenSwitchControllerManagedSwitchPortsInterfaceTags(v *[]models.SwitchControllerManagedSwitchPortsInterfaceTags, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.TagName; tmp != nil {
				v["tag_name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "tag_name")
	}

	return flat
}

func flattenSwitchControllerManagedSwitchPortsMembers(v *[]models.SwitchControllerManagedSwitchPortsMembers, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.MemberName; tmp != nil {
				v["member_name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "member_name")
	}

	return flat
}

func flattenSwitchControllerManagedSwitchPortsUntaggedVlans(v *[]models.SwitchControllerManagedSwitchPortsUntaggedVlans, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.VlanName; tmp != nil {
				v["vlan_name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "vlan_name")
	}

	return flat
}

func flattenSwitchControllerManagedSwitchRemoteLog(v *[]models.SwitchControllerManagedSwitchRemoteLog, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Csv; tmp != nil {
				v["csv"] = *tmp
			}

			if tmp := cfg.Facility; tmp != nil {
				v["facility"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Port; tmp != nil {
				v["port"] = *tmp
			}

			if tmp := cfg.Server; tmp != nil {
				v["server"] = *tmp
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSwitchControllerManagedSwitchSnmpCommunity(v *[]models.SwitchControllerManagedSwitchSnmpCommunity, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Events; tmp != nil {
				v["events"] = *tmp
			}

			if tmp := cfg.Hosts; tmp != nil {
				v["hosts"] = flattenSwitchControllerManagedSwitchSnmpCommunityHosts(tmp, sort)
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.QueryV1Port; tmp != nil {
				v["query_v1_port"] = *tmp
			}

			if tmp := cfg.QueryV1Status; tmp != nil {
				v["query_v1_status"] = *tmp
			}

			if tmp := cfg.QueryV2cPort; tmp != nil {
				v["query_v2c_port"] = *tmp
			}

			if tmp := cfg.QueryV2cStatus; tmp != nil {
				v["query_v2c_status"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.TrapV1Lport; tmp != nil {
				v["trap_v1_lport"] = *tmp
			}

			if tmp := cfg.TrapV1Rport; tmp != nil {
				v["trap_v1_rport"] = *tmp
			}

			if tmp := cfg.TrapV1Status; tmp != nil {
				v["trap_v1_status"] = *tmp
			}

			if tmp := cfg.TrapV2cLport; tmp != nil {
				v["trap_v2c_lport"] = *tmp
			}

			if tmp := cfg.TrapV2cRport; tmp != nil {
				v["trap_v2c_rport"] = *tmp
			}

			if tmp := cfg.TrapV2cStatus; tmp != nil {
				v["trap_v2c_status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSwitchControllerManagedSwitchSnmpCommunityHosts(v *[]models.SwitchControllerManagedSwitchSnmpCommunityHosts, sort bool) interface{} {
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

func flattenSwitchControllerManagedSwitchSnmpSysinfo(v *[]models.SwitchControllerManagedSwitchSnmpSysinfo, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.ContactInfo; tmp != nil {
				v["contact_info"] = *tmp
			}

			if tmp := cfg.Description; tmp != nil {
				v["description"] = *tmp
			}

			if tmp := cfg.EngineId; tmp != nil {
				v["engine_id"] = *tmp
			}

			if tmp := cfg.Location; tmp != nil {
				v["location"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSwitchControllerManagedSwitchSnmpTrapThreshold(v *[]models.SwitchControllerManagedSwitchSnmpTrapThreshold, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.TrapHighCpuThreshold; tmp != nil {
				v["trap_high_cpu_threshold"] = *tmp
			}

			if tmp := cfg.TrapLogFullThreshold; tmp != nil {
				v["trap_log_full_threshold"] = *tmp
			}

			if tmp := cfg.TrapLowMemoryThreshold; tmp != nil {
				v["trap_low_memory_threshold"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSwitchControllerManagedSwitchSnmpUser(v *[]models.SwitchControllerManagedSwitchSnmpUser, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AuthProto; tmp != nil {
				v["auth_proto"] = *tmp
			}

			if tmp := cfg.AuthPwd; tmp != nil {
				v["auth_pwd"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.PrivProto; tmp != nil {
				v["priv_proto"] = *tmp
			}

			if tmp := cfg.PrivPwd; tmp != nil {
				v["priv_pwd"] = *tmp
			}

			if tmp := cfg.Queries; tmp != nil {
				v["queries"] = *tmp
			}

			if tmp := cfg.QueryPort; tmp != nil {
				v["query_port"] = *tmp
			}

			if tmp := cfg.SecurityLevel; tmp != nil {
				v["security_level"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSwitchControllerManagedSwitchStaticMac(v *[]models.SwitchControllerManagedSwitchStaticMac, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Description; tmp != nil {
				v["description"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			if tmp := cfg.Mac; tmp != nil {
				v["mac"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			if tmp := cfg.Vlan; tmp != nil {
				v["vlan"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSwitchControllerManagedSwitchStormControl(v *[]models.SwitchControllerManagedSwitchStormControl, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Broadcast; tmp != nil {
				v["broadcast"] = *tmp
			}

			if tmp := cfg.LocalOverride; tmp != nil {
				v["local_override"] = *tmp
			}

			if tmp := cfg.Rate; tmp != nil {
				v["rate"] = *tmp
			}

			if tmp := cfg.UnknownMulticast; tmp != nil {
				v["unknown_multicast"] = *tmp
			}

			if tmp := cfg.UnknownUnicast; tmp != nil {
				v["unknown_unicast"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSwitchControllerManagedSwitchStpInstance(v *[]models.SwitchControllerManagedSwitchStpInstance, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSwitchControllerManagedSwitchStpSettings(v *[]models.SwitchControllerManagedSwitchStpSettings, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.ForwardTime; tmp != nil {
				v["forward_time"] = *tmp
			}

			if tmp := cfg.HelloTime; tmp != nil {
				v["hello_time"] = *tmp
			}

			if tmp := cfg.LocalOverride; tmp != nil {
				v["local_override"] = *tmp
			}

			if tmp := cfg.MaxAge; tmp != nil {
				v["max_age"] = *tmp
			}

			if tmp := cfg.MaxHops; tmp != nil {
				v["max_hops"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.PendingTimer; tmp != nil {
				v["pending_timer"] = *tmp
			}

			if tmp := cfg.Revision; tmp != nil {
				v["revision"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSwitchControllerManagedSwitchSwitchLog(v *[]models.SwitchControllerManagedSwitchSwitchLog, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.LocalOverride; tmp != nil {
				v["local_override"] = *tmp
			}

			if tmp := cfg.Severity; tmp != nil {
				v["severity"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func refreshObjectSwitchControllerManagedSwitch(d *schema.ResourceData, o *models.SwitchControllerManagedSwitch, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.The8021XSettings != nil {
		if err = d.Set("802_1x_settings", flattenSwitchControllerManagedSwitch8021XSettings(o.The8021XSettings, sort)); err != nil {
			return diag.Errorf("error reading 802_1x_settings: %v", err)
		}
	}

	if o.AccessProfile != nil {
		v := *o.AccessProfile

		if err = d.Set("access_profile", v); err != nil {
			return diag.Errorf("error reading access_profile: %v", err)
		}
	}

	if o.CustomCommand != nil {
		if err = d.Set("custom_command", flattenSwitchControllerManagedSwitchCustomCommand(o.CustomCommand, sort)); err != nil {
			return diag.Errorf("error reading custom_command: %v", err)
		}
	}

	if o.DelayedRestartTrigger != nil {
		v := *o.DelayedRestartTrigger

		if err = d.Set("delayed_restart_trigger", v); err != nil {
			return diag.Errorf("error reading delayed_restart_trigger: %v", err)
		}
	}

	if o.Description != nil {
		v := *o.Description

		if err = d.Set("description", v); err != nil {
			return diag.Errorf("error reading description: %v", err)
		}
	}

	if o.DhcpServerAccessList != nil {
		v := *o.DhcpServerAccessList

		if err = d.Set("dhcp_server_access_list", v); err != nil {
			return diag.Errorf("error reading dhcp_server_access_list: %v", err)
		}
	}

	if o.DirectlyConnected != nil {
		v := *o.DirectlyConnected

		if err = d.Set("directly_connected", v); err != nil {
			return diag.Errorf("error reading directly_connected: %v", err)
		}
	}

	if o.DynamicCapability != nil {
		v := *o.DynamicCapability

		if err = d.Set("dynamic_capability", v); err != nil {
			return diag.Errorf("error reading dynamic_capability: %v", err)
		}
	}

	if o.DynamicallyDiscovered != nil {
		v := *o.DynamicallyDiscovered

		if err = d.Set("dynamically_discovered", v); err != nil {
			return diag.Errorf("error reading dynamically_discovered: %v", err)
		}
	}

	if o.FirmwareProvision != nil {
		v := *o.FirmwareProvision

		if err = d.Set("firmware_provision", v); err != nil {
			return diag.Errorf("error reading firmware_provision: %v", err)
		}
	}

	if o.FirmwareProvisionVersion != nil {
		v := *o.FirmwareProvisionVersion

		if err = d.Set("firmware_provision_version", v); err != nil {
			return diag.Errorf("error reading firmware_provision_version: %v", err)
		}
	}

	if o.FlowIdentity != nil {
		v := *o.FlowIdentity

		if err = d.Set("flow_identity", v); err != nil {
			return diag.Errorf("error reading flow_identity: %v", err)
		}
	}

	if o.FswWan1Admin != nil {
		v := *o.FswWan1Admin

		if err = d.Set("fsw_wan1_admin", v); err != nil {
			return diag.Errorf("error reading fsw_wan1_admin: %v", err)
		}
	}

	if o.FswWan1Peer != nil {
		v := *o.FswWan1Peer

		if err = d.Set("fsw_wan1_peer", v); err != nil {
			return diag.Errorf("error reading fsw_wan1_peer: %v", err)
		}
	}

	if o.IgmpSnooping != nil {
		if err = d.Set("igmp_snooping", flattenSwitchControllerManagedSwitchIgmpSnooping(o.IgmpSnooping, sort)); err != nil {
			return diag.Errorf("error reading igmp_snooping: %v", err)
		}
	}

	if o.IpSourceGuard != nil {
		if err = d.Set("ip_source_guard", flattenSwitchControllerManagedSwitchIpSourceGuard(o.IpSourceGuard, sort)); err != nil {
			return diag.Errorf("error reading ip_source_guard: %v", err)
		}
	}

	if o.L3Discovered != nil {
		v := *o.L3Discovered

		if err = d.Set("l3_discovered", v); err != nil {
			return diag.Errorf("error reading l3_discovered: %v", err)
		}
	}

	if o.MaxAllowedTrunkMembers != nil {
		v := *o.MaxAllowedTrunkMembers

		if err = d.Set("max_allowed_trunk_members", v); err != nil {
			return diag.Errorf("error reading max_allowed_trunk_members: %v", err)
		}
	}

	if o.MclagIgmpSnoopingAware != nil {
		v := *o.MclagIgmpSnoopingAware

		if err = d.Set("mclag_igmp_snooping_aware", v); err != nil {
			return diag.Errorf("error reading mclag_igmp_snooping_aware: %v", err)
		}
	}

	if o.Mirror != nil {
		if err = d.Set("mirror", flattenSwitchControllerManagedSwitchMirror(o.Mirror, sort)); err != nil {
			return diag.Errorf("error reading mirror: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.OverrideSnmpCommunity != nil {
		v := *o.OverrideSnmpCommunity

		if err = d.Set("override_snmp_community", v); err != nil {
			return diag.Errorf("error reading override_snmp_community: %v", err)
		}
	}

	if o.OverrideSnmpSysinfo != nil {
		v := *o.OverrideSnmpSysinfo

		if err = d.Set("override_snmp_sysinfo", v); err != nil {
			return diag.Errorf("error reading override_snmp_sysinfo: %v", err)
		}
	}

	if o.OverrideSnmpTrapThreshold != nil {
		v := *o.OverrideSnmpTrapThreshold

		if err = d.Set("override_snmp_trap_threshold", v); err != nil {
			return diag.Errorf("error reading override_snmp_trap_threshold: %v", err)
		}
	}

	if o.OverrideSnmpUser != nil {
		v := *o.OverrideSnmpUser

		if err = d.Set("override_snmp_user", v); err != nil {
			return diag.Errorf("error reading override_snmp_user: %v", err)
		}
	}

	if o.OwnerVdom != nil {
		v := *o.OwnerVdom

		if err = d.Set("owner_vdom", v); err != nil {
			return diag.Errorf("error reading owner_vdom: %v", err)
		}
	}

	if o.PoeDetectionType != nil {
		v := *o.PoeDetectionType

		if err = d.Set("poe_detection_type", v); err != nil {
			return diag.Errorf("error reading poe_detection_type: %v", err)
		}
	}

	if o.PoeLldpDetection != nil {
		v := *o.PoeLldpDetection

		if err = d.Set("poe_lldp_detection", v); err != nil {
			return diag.Errorf("error reading poe_lldp_detection: %v", err)
		}
	}

	if o.PoePreStandardDetection != nil {
		v := *o.PoePreStandardDetection

		if err = d.Set("poe_pre_standard_detection", v); err != nil {
			return diag.Errorf("error reading poe_pre_standard_detection: %v", err)
		}
	}

	if o.Ports != nil {
		if err = d.Set("ports", flattenSwitchControllerManagedSwitchPorts(o.Ports, sort)); err != nil {
			return diag.Errorf("error reading ports: %v", err)
		}
	}

	if o.PreProvisioned != nil {
		v := *o.PreProvisioned

		if err = d.Set("pre_provisioned", v); err != nil {
			return diag.Errorf("error reading pre_provisioned: %v", err)
		}
	}

	if o.QosDropPolicy != nil {
		v := *o.QosDropPolicy

		if err = d.Set("qos_drop_policy", v); err != nil {
			return diag.Errorf("error reading qos_drop_policy: %v", err)
		}
	}

	if o.QosRedProbability != nil {
		v := *o.QosRedProbability

		if err = d.Set("qos_red_probability", v); err != nil {
			return diag.Errorf("error reading qos_red_probability: %v", err)
		}
	}

	if o.RemoteLog != nil {
		if err = d.Set("remote_log", flattenSwitchControllerManagedSwitchRemoteLog(o.RemoteLog, sort)); err != nil {
			return diag.Errorf("error reading remote_log: %v", err)
		}
	}

	if o.SnmpCommunity != nil {
		if err = d.Set("snmp_community", flattenSwitchControllerManagedSwitchSnmpCommunity(o.SnmpCommunity, sort)); err != nil {
			return diag.Errorf("error reading snmp_community: %v", err)
		}
	}

	if o.SnmpSysinfo != nil {
		if err = d.Set("snmp_sysinfo", flattenSwitchControllerManagedSwitchSnmpSysinfo(o.SnmpSysinfo, sort)); err != nil {
			return diag.Errorf("error reading snmp_sysinfo: %v", err)
		}
	}

	if o.SnmpTrapThreshold != nil {
		if err = d.Set("snmp_trap_threshold", flattenSwitchControllerManagedSwitchSnmpTrapThreshold(o.SnmpTrapThreshold, sort)); err != nil {
			return diag.Errorf("error reading snmp_trap_threshold: %v", err)
		}
	}

	if o.SnmpUser != nil {
		if err = d.Set("snmp_user", flattenSwitchControllerManagedSwitchSnmpUser(o.SnmpUser, sort)); err != nil {
			return diag.Errorf("error reading snmp_user: %v", err)
		}
	}

	if o.StagedImageVersion != nil {
		v := *o.StagedImageVersion

		if err = d.Set("staged_image_version", v); err != nil {
			return diag.Errorf("error reading staged_image_version: %v", err)
		}
	}

	if o.StaticMac != nil {
		if err = d.Set("static_mac", flattenSwitchControllerManagedSwitchStaticMac(o.StaticMac, sort)); err != nil {
			return diag.Errorf("error reading static_mac: %v", err)
		}
	}

	if o.StormControl != nil {
		if err = d.Set("storm_control", flattenSwitchControllerManagedSwitchStormControl(o.StormControl, sort)); err != nil {
			return diag.Errorf("error reading storm_control: %v", err)
		}
	}

	if o.StpInstance != nil {
		if err = d.Set("stp_instance", flattenSwitchControllerManagedSwitchStpInstance(o.StpInstance, sort)); err != nil {
			return diag.Errorf("error reading stp_instance: %v", err)
		}
	}

	if o.StpSettings != nil {
		if err = d.Set("stp_settings", flattenSwitchControllerManagedSwitchStpSettings(o.StpSettings, sort)); err != nil {
			return diag.Errorf("error reading stp_settings: %v", err)
		}
	}

	if o.SwitchDeviceTag != nil {
		v := *o.SwitchDeviceTag

		if err = d.Set("switch_device_tag", v); err != nil {
			return diag.Errorf("error reading switch_device_tag: %v", err)
		}
	}

	if o.SwitchDhcp_opt43_key != nil {
		v := *o.SwitchDhcp_opt43_key

		if err = d.Set("switch_dhcp_opt43_key", v); err != nil {
			return diag.Errorf("error reading switch_dhcp_opt43_key: %v", err)
		}
	}

	if o.SwitchId != nil {
		v := *o.SwitchId

		if err = d.Set("switch_id", v); err != nil {
			return diag.Errorf("error reading switch_id: %v", err)
		}
	}

	if o.SwitchLog != nil {
		if err = d.Set("switch_log", flattenSwitchControllerManagedSwitchSwitchLog(o.SwitchLog, sort)); err != nil {
			return diag.Errorf("error reading switch_log: %v", err)
		}
	}

	if o.SwitchProfile != nil {
		v := *o.SwitchProfile

		if err = d.Set("switch_profile", v); err != nil {
			return diag.Errorf("error reading switch_profile: %v", err)
		}
	}

	if o.TdrSupported != nil {
		v := *o.TdrSupported

		if err = d.Set("tdr_supported", v); err != nil {
			return diag.Errorf("error reading tdr_supported: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	if o.Version != nil {
		v := *o.Version

		if err = d.Set("version", v); err != nil {
			return diag.Errorf("error reading version: %v", err)
		}
	}

	return nil
}

func expandSwitchControllerManagedSwitch8021XSettings(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitch8021XSettings, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitch8021XSettings

	for i := range l {
		tmp := models.SwitchControllerManagedSwitch8021XSettings{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.link_down_auth", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LinkDownAuth = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.local_override", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LocalOverride = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_reauth_attempt", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.MaxReauthAttempt = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.reauth_period", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.ReauthPeriod = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tx_period", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.TxPeriod = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerManagedSwitchCustomCommand(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchCustomCommand, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchCustomCommand

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchCustomCommand{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.command_entry", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CommandEntry = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.command_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CommandName = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerManagedSwitchIgmpSnooping(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchIgmpSnooping, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchIgmpSnooping

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchIgmpSnooping{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.aging_time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.AgingTime = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.flood_unknown_multicast", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FloodUnknownMulticast = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.local_override", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LocalOverride = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlans", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSwitchControllerManagedSwitchIgmpSnoopingVlans(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SwitchControllerManagedSwitchIgmpSnoopingVlans
			// 	}
			tmp.Vlans = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlans(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchIgmpSnoopingVlans, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchIgmpSnoopingVlans

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchIgmpSnoopingVlans{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.proxy", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Proxy = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.querier", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Querier = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.querier_addr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.QuerierAddr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Version = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VlanName = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerManagedSwitchIpSourceGuard(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchIpSourceGuard, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchIpSourceGuard

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchIpSourceGuard{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.binding_entry", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSwitchControllerManagedSwitchIpSourceGuardBindingEntry(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SwitchControllerManagedSwitchIpSourceGuardBindingEntry
			// 	}
			tmp.BindingEntry = v2

		}

		pre_append = fmt.Sprintf("%s.%d.description", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Description = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Port = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerManagedSwitchIpSourceGuardBindingEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchIpSourceGuardBindingEntry, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchIpSourceGuardBindingEntry

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchIpSourceGuardBindingEntry{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.entry_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EntryName = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mac", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mac = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerManagedSwitchMirror(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchMirror, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchMirror

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchMirror{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dst", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dst = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.src_egress", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSwitchControllerManagedSwitchMirrorSrcEgress(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SwitchControllerManagedSwitchMirrorSrcEgress
			// 	}
			tmp.SrcEgress = v2

		}

		pre_append = fmt.Sprintf("%s.%d.src_ingress", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSwitchControllerManagedSwitchMirrorSrcIngress(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SwitchControllerManagedSwitchMirrorSrcIngress
			// 	}
			tmp.SrcIngress = v2

		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.switching_packet", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SwitchingPacket = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerManagedSwitchMirrorSrcEgress(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchMirrorSrcEgress, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchMirrorSrcEgress

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchMirrorSrcEgress{}
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

func expandSwitchControllerManagedSwitchMirrorSrcIngress(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchMirrorSrcIngress, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchMirrorSrcIngress

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchMirrorSrcIngress{}
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

func expandSwitchControllerManagedSwitchPorts(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchPorts, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchPorts

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchPorts{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.access_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AccessMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.aggregator_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AggregatorMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.allowed_vlans", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSwitchControllerManagedSwitchPortsAllowedVlans(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SwitchControllerManagedSwitchPortsAllowedVlans
			// 	}
			tmp.AllowedVlans = v2

		}

		pre_append = fmt.Sprintf("%s.%d.allowed_vlans_all", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AllowedVlansAll = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.arp_inspection_trust", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ArpInspectionTrust = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bundle", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Bundle = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.description", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Description = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dhcp_snoop_option82_trust", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DhcpSnoopOption82Trust = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dhcp_snooping", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DhcpSnooping = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.discard_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DiscardMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.edge_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EdgePort = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.export_tags", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSwitchControllerManagedSwitchPortsExportTags(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SwitchControllerManagedSwitchPortsExportTags
			// 	}
			tmp.ExportTags = v2

		}

		pre_append = fmt.Sprintf("%s.%d.export_to", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExportTo = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.export_to_pool", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExportToPool = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fec_capable", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.FecCapable = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fec_state", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FecState = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fgt_peer_device_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FgtPeerDeviceName = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fgt_peer_port_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FgtPeerPortName = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fiber_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.FiberPort = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.flags", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Flags = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.flow_control", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FlowControl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortilink_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.FortilinkPort = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.igmp_snooping", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IgmpSnooping = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.igmps_flood_reports", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IgmpsFloodReports = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.igmps_flood_traffic", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IgmpsFloodTraffic = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface_tags", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSwitchControllerManagedSwitchPortsInterfaceTags(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SwitchControllerManagedSwitchPortsInterfaceTags
			// 	}
			tmp.InterfaceTags = v2

		}

		pre_append = fmt.Sprintf("%s.%d.ip_source_guard", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IpSourceGuard = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.isl_local_trunk_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IslLocalTrunkName = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.isl_peer_device_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IslPeerDeviceName = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.isl_peer_port_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IslPeerPortName = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.lacp_speed", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LacpSpeed = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.learning_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.LearningLimit = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.lldp_profile", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LldpProfile = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.lldp_status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LldpStatus = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.loop_guard", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LoopGuard = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.loop_guard_timeout", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.LoopGuardTimeout = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mac_addr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MacAddr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.matched_dpp_intf_tags", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MatchedDppIntfTags = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.matched_dpp_policy", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MatchedDppPolicy = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_bundle", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.MaxBundle = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mclag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mclag = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mclag_icl_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.MclagIclPort = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.media_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MediaType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.member_withdrawal_behavior", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MemberWithdrawalBehavior = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.members", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSwitchControllerManagedSwitchPortsMembers(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SwitchControllerManagedSwitchPortsMembers
			// 	}
			tmp.Members = v2

		}

		pre_append = fmt.Sprintf("%s.%d.min_bundle", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.MinBundle = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.p2p_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.P2pPort = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.packet_sample_rate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.PacketSampleRate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.packet_sampler", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PacketSampler = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pause_meter", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.PauseMeter = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pause_meter_resume", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PauseMeterResume = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.poe_capable", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.PoeCapable = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.poe_max_power", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PoeMaxPower = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.poe_pre_standard_detection", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PoePreStandardDetection = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.poe_standard", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PoeStandard = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.poe_status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PoeStatus = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PortName = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port_number", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.PortNumber = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port_owner", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PortOwner = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port_policy", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PortPolicy = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port_prefix_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.PortPrefixType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port_security_policy", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PortSecurityPolicy = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port_selection_criteria", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PortSelectionCriteria = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ptp_policy", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PtpPolicy = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.qos_policy", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.QosPolicy = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rpvst_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RpvstPort = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sample_direction", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SampleDirection = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sflow_counter_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.SflowCounterInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.speed", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Speed = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.stacking_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.StackingPort = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sticky_mac", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StickyMac = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.storm_control_policy", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StormControlPolicy = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.stp_bpdu_guard", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StpBpduGuard = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.stp_bpdu_guard_timeout", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.StpBpduGuardTimeout = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.stp_root_guard", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StpRootGuard = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.stp_state", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StpState = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.switch_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SwitchId = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.untagged_vlans", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSwitchControllerManagedSwitchPortsUntaggedVlans(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SwitchControllerManagedSwitchPortsUntaggedVlans
			// 	}
			tmp.UntaggedVlans = v2

		}

		pre_append = fmt.Sprintf("%s.%d.vlan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Vlan = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerManagedSwitchPortsAllowedVlans(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchPortsAllowedVlans, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchPortsAllowedVlans

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchPortsAllowedVlans{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.vlan_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VlanName = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerManagedSwitchPortsExportTags(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchPortsExportTags, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchPortsExportTags

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchPortsExportTags{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.tag_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TagName = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerManagedSwitchPortsInterfaceTags(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchPortsInterfaceTags, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchPortsInterfaceTags

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchPortsInterfaceTags{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.tag_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TagName = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerManagedSwitchPortsMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchPortsMembers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchPortsMembers

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchPortsMembers{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.member_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MemberName = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerManagedSwitchPortsUntaggedVlans(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchPortsUntaggedVlans, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchPortsUntaggedVlans

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchPortsUntaggedVlans{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.vlan_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VlanName = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerManagedSwitchRemoteLog(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchRemoteLog, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchRemoteLog

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchRemoteLog{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.csv", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Csv = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.facility", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Facility = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Port = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.server", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Server = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
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
	return &result, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchSnmpCommunity, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchSnmpCommunity

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchSnmpCommunity{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.events", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Events = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.hosts", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSwitchControllerManagedSwitchSnmpCommunityHosts(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SwitchControllerManagedSwitchSnmpCommunityHosts
			// 	}
			tmp.Hosts = v2

		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.query_v1_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.QueryV1Port = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.query_v1_status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.QueryV1Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.query_v2c_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.QueryV2cPort = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.query_v2c_status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.QueryV2cStatus = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.trap_v1_lport", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.TrapV1Lport = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.trap_v1_rport", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.TrapV1Rport = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.trap_v1_status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TrapV1Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.trap_v2c_lport", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.TrapV2cLport = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.trap_v2c_rport", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.TrapV2cRport = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.trap_v2c_status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TrapV2cStatus = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityHosts(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchSnmpCommunityHosts, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchSnmpCommunityHosts

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchSnmpCommunityHosts{}
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

func expandSwitchControllerManagedSwitchSnmpSysinfo(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchSnmpSysinfo, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchSnmpSysinfo

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchSnmpSysinfo{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.contact_info", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ContactInfo = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.description", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Description = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.engine_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EngineId = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.location", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Location = &v2
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
	return &result, nil
}

func expandSwitchControllerManagedSwitchSnmpTrapThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchSnmpTrapThreshold, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchSnmpTrapThreshold

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchSnmpTrapThreshold{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.trap_high_cpu_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.TrapHighCpuThreshold = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.trap_log_full_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.TrapLogFullThreshold = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.trap_low_memory_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.TrapLowMemoryThreshold = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerManagedSwitchSnmpUser(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchSnmpUser, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchSnmpUser

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchSnmpUser{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.auth_proto", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthProto = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auth_pwd", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthPwd = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priv_proto", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PrivProto = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priv_pwd", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PrivPwd = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.queries", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Queries = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.query_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.QueryPort = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.security_level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SecurityLevel = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerManagedSwitchStaticMac(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchStaticMac, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchStaticMac

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchStaticMac{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.description", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Description = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mac", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mac = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Vlan = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerManagedSwitchStormControl(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchStormControl, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchStormControl

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchStormControl{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.broadcast", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Broadcast = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.local_override", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LocalOverride = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Rate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unknown_multicast", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnknownMulticast = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unknown_unicast", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnknownUnicast = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerManagedSwitchStpInstance(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchStpInstance, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchStpInstance

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchStpInstance{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Priority = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerManagedSwitchStpSettings(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchStpSettings, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchStpSettings

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchStpSettings{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.forward_time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.ForwardTime = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.hello_time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.HelloTime = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.local_override", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LocalOverride = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_age", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.MaxAge = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_hops", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.MaxHops = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pending_timer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.PendingTimer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.revision", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Revision = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerManagedSwitchSwitchLog(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerManagedSwitchSwitchLog, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerManagedSwitchSwitchLog

	for i := range l {
		tmp := models.SwitchControllerManagedSwitchSwitchLog{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.local_override", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LocalOverride = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.severity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Severity = &v2
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
	return &result, nil
}

func getObjectSwitchControllerManagedSwitch(d *schema.ResourceData, sv string) (*models.SwitchControllerManagedSwitch, diag.Diagnostics) {
	obj := models.SwitchControllerManagedSwitch{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("802_1x_settings"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("802_1x_settings", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerManagedSwitch8021XSettings(d, v, "802_1x_settings", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.The8021XSettings = t
		}
	} else if d.HasChange("802_1x_settings") {
		old, new := d.GetChange("802_1x_settings")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.The8021XSettings = &[]models.SwitchControllerManagedSwitch8021XSettings{}
		}
	}
	if v1, ok := d.GetOk("access_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("access_profile", sv)
				diags = append(diags, e)
			}
			obj.AccessProfile = &v2
		}
	}
	if v, ok := d.GetOk("custom_command"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("custom_command", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerManagedSwitchCustomCommand(d, v, "custom_command", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.CustomCommand = t
		}
	} else if d.HasChange("custom_command") {
		old, new := d.GetChange("custom_command")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.CustomCommand = &[]models.SwitchControllerManagedSwitchCustomCommand{}
		}
	}
	if v1, ok := d.GetOk("delayed_restart_trigger"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("delayed_restart_trigger", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DelayedRestartTrigger = &tmp
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
	if v1, ok := d.GetOk("dhcp_server_access_list"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("dhcp_server_access_list", sv)
				diags = append(diags, e)
			}
			obj.DhcpServerAccessList = &v2
		}
	}
	if v1, ok := d.GetOk("directly_connected"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("directly_connected", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DirectlyConnected = &tmp
		}
	}
	if v1, ok := d.GetOk("dynamic_capability"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dynamic_capability", sv)
				diags = append(diags, e)
			}
			obj.DynamicCapability = &v2
		}
	}
	if v1, ok := d.GetOk("dynamically_discovered"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dynamically_discovered", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DynamicallyDiscovered = &tmp
		}
	}
	if v1, ok := d.GetOk("firmware_provision"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("firmware_provision", sv)
				diags = append(diags, e)
			}
			obj.FirmwareProvision = &v2
		}
	}
	if v1, ok := d.GetOk("firmware_provision_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("firmware_provision_version", sv)
				diags = append(diags, e)
			}
			obj.FirmwareProvisionVersion = &v2
		}
	}
	if v1, ok := d.GetOk("flow_identity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("flow_identity", sv)
				diags = append(diags, e)
			}
			obj.FlowIdentity = &v2
		}
	}
	if v1, ok := d.GetOk("fsw_wan1_admin"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fsw_wan1_admin", sv)
				diags = append(diags, e)
			}
			obj.FswWan1Admin = &v2
		}
	}
	if v1, ok := d.GetOk("fsw_wan1_peer"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fsw_wan1_peer", sv)
				diags = append(diags, e)
			}
			obj.FswWan1Peer = &v2
		}
	}
	if v, ok := d.GetOk("igmp_snooping"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("igmp_snooping", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerManagedSwitchIgmpSnooping(d, v, "igmp_snooping", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.IgmpSnooping = t
		}
	} else if d.HasChange("igmp_snooping") {
		old, new := d.GetChange("igmp_snooping")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.IgmpSnooping = &[]models.SwitchControllerManagedSwitchIgmpSnooping{}
		}
	}
	if v, ok := d.GetOk("ip_source_guard"); ok {
		if !utils.CheckVer(sv, "v6.4.0", "") {
			e := utils.AttributeVersionWarning("ip_source_guard", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerManagedSwitchIpSourceGuard(d, v, "ip_source_guard", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.IpSourceGuard = t
		}
	} else if d.HasChange("ip_source_guard") {
		old, new := d.GetChange("ip_source_guard")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.IpSourceGuard = &[]models.SwitchControllerManagedSwitchIpSourceGuard{}
		}
	}
	if v1, ok := d.GetOk("l3_discovered"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("l3_discovered", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.L3Discovered = &tmp
		}
	}
	if v1, ok := d.GetOk("max_allowed_trunk_members"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("max_allowed_trunk_members", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxAllowedTrunkMembers = &tmp
		}
	}
	if v1, ok := d.GetOk("mclag_igmp_snooping_aware"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mclag_igmp_snooping_aware", sv)
				diags = append(diags, e)
			}
			obj.MclagIgmpSnoopingAware = &v2
		}
	}
	if v, ok := d.GetOk("mirror"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("mirror", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerManagedSwitchMirror(d, v, "mirror", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Mirror = t
		}
	} else if d.HasChange("mirror") {
		old, new := d.GetChange("mirror")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Mirror = &[]models.SwitchControllerManagedSwitchMirror{}
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
	if v1, ok := d.GetOk("override_snmp_community"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("override_snmp_community", sv)
				diags = append(diags, e)
			}
			obj.OverrideSnmpCommunity = &v2
		}
	}
	if v1, ok := d.GetOk("override_snmp_sysinfo"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("override_snmp_sysinfo", sv)
				diags = append(diags, e)
			}
			obj.OverrideSnmpSysinfo = &v2
		}
	}
	if v1, ok := d.GetOk("override_snmp_trap_threshold"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("override_snmp_trap_threshold", sv)
				diags = append(diags, e)
			}
			obj.OverrideSnmpTrapThreshold = &v2
		}
	}
	if v1, ok := d.GetOk("override_snmp_user"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("override_snmp_user", sv)
				diags = append(diags, e)
			}
			obj.OverrideSnmpUser = &v2
		}
	}
	if v1, ok := d.GetOk("owner_vdom"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("owner_vdom", sv)
				diags = append(diags, e)
			}
			obj.OwnerVdom = &v2
		}
	}
	if v1, ok := d.GetOk("poe_detection_type"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("poe_detection_type", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PoeDetectionType = &tmp
		}
	}
	if v1, ok := d.GetOk("poe_lldp_detection"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("poe_lldp_detection", sv)
				diags = append(diags, e)
			}
			obj.PoeLldpDetection = &v2
		}
	}
	if v1, ok := d.GetOk("poe_pre_standard_detection"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("poe_pre_standard_detection", sv)
				diags = append(diags, e)
			}
			obj.PoePreStandardDetection = &v2
		}
	}
	if v, ok := d.GetOk("ports"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ports", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerManagedSwitchPorts(d, v, "ports", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Ports = t
		}
	} else if d.HasChange("ports") {
		old, new := d.GetChange("ports")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Ports = &[]models.SwitchControllerManagedSwitchPorts{}
		}
	}
	if v1, ok := d.GetOk("pre_provisioned"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pre_provisioned", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PreProvisioned = &tmp
		}
	}
	if v1, ok := d.GetOk("qos_drop_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("qos_drop_policy", sv)
				diags = append(diags, e)
			}
			obj.QosDropPolicy = &v2
		}
	}
	if v1, ok := d.GetOk("qos_red_probability"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("qos_red_probability", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.QosRedProbability = &tmp
		}
	}
	if v, ok := d.GetOk("remote_log"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("remote_log", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerManagedSwitchRemoteLog(d, v, "remote_log", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.RemoteLog = t
		}
	} else if d.HasChange("remote_log") {
		old, new := d.GetChange("remote_log")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.RemoteLog = &[]models.SwitchControllerManagedSwitchRemoteLog{}
		}
	}
	if v, ok := d.GetOk("snmp_community"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("snmp_community", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerManagedSwitchSnmpCommunity(d, v, "snmp_community", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SnmpCommunity = t
		}
	} else if d.HasChange("snmp_community") {
		old, new := d.GetChange("snmp_community")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SnmpCommunity = &[]models.SwitchControllerManagedSwitchSnmpCommunity{}
		}
	}
	if v, ok := d.GetOk("snmp_sysinfo"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("snmp_sysinfo", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerManagedSwitchSnmpSysinfo(d, v, "snmp_sysinfo", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SnmpSysinfo = t
		}
	} else if d.HasChange("snmp_sysinfo") {
		old, new := d.GetChange("snmp_sysinfo")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SnmpSysinfo = &[]models.SwitchControllerManagedSwitchSnmpSysinfo{}
		}
	}
	if v, ok := d.GetOk("snmp_trap_threshold"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("snmp_trap_threshold", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerManagedSwitchSnmpTrapThreshold(d, v, "snmp_trap_threshold", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SnmpTrapThreshold = t
		}
	} else if d.HasChange("snmp_trap_threshold") {
		old, new := d.GetChange("snmp_trap_threshold")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SnmpTrapThreshold = &[]models.SwitchControllerManagedSwitchSnmpTrapThreshold{}
		}
	}
	if v, ok := d.GetOk("snmp_user"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("snmp_user", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerManagedSwitchSnmpUser(d, v, "snmp_user", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SnmpUser = t
		}
	} else if d.HasChange("snmp_user") {
		old, new := d.GetChange("snmp_user")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SnmpUser = &[]models.SwitchControllerManagedSwitchSnmpUser{}
		}
	}
	if v1, ok := d.GetOk("staged_image_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("staged_image_version", sv)
				diags = append(diags, e)
			}
			obj.StagedImageVersion = &v2
		}
	}
	if v, ok := d.GetOk("static_mac"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("static_mac", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerManagedSwitchStaticMac(d, v, "static_mac", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.StaticMac = t
		}
	} else if d.HasChange("static_mac") {
		old, new := d.GetChange("static_mac")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.StaticMac = &[]models.SwitchControllerManagedSwitchStaticMac{}
		}
	}
	if v, ok := d.GetOk("storm_control"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("storm_control", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerManagedSwitchStormControl(d, v, "storm_control", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.StormControl = t
		}
	} else if d.HasChange("storm_control") {
		old, new := d.GetChange("storm_control")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.StormControl = &[]models.SwitchControllerManagedSwitchStormControl{}
		}
	}
	if v, ok := d.GetOk("stp_instance"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("stp_instance", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerManagedSwitchStpInstance(d, v, "stp_instance", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.StpInstance = t
		}
	} else if d.HasChange("stp_instance") {
		old, new := d.GetChange("stp_instance")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.StpInstance = &[]models.SwitchControllerManagedSwitchStpInstance{}
		}
	}
	if v, ok := d.GetOk("stp_settings"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("stp_settings", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerManagedSwitchStpSettings(d, v, "stp_settings", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.StpSettings = t
		}
	} else if d.HasChange("stp_settings") {
		old, new := d.GetChange("stp_settings")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.StpSettings = &[]models.SwitchControllerManagedSwitchStpSettings{}
		}
	}
	if v1, ok := d.GetOk("switch_device_tag"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("switch_device_tag", sv)
				diags = append(diags, e)
			}
			obj.SwitchDeviceTag = &v2
		}
	}
	if v1, ok := d.GetOk("switch_dhcp_opt43_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("switch_dhcp_opt43_key", sv)
				diags = append(diags, e)
			}
			obj.SwitchDhcp_opt43_key = &v2
		}
	}
	if v1, ok := d.GetOk("switch_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("switch_id", sv)
				diags = append(diags, e)
			}
			obj.SwitchId = &v2
		}
	}
	if v, ok := d.GetOk("switch_log"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("switch_log", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerManagedSwitchSwitchLog(d, v, "switch_log", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SwitchLog = t
		}
	} else if d.HasChange("switch_log") {
		old, new := d.GetChange("switch_log")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SwitchLog = &[]models.SwitchControllerManagedSwitchSwitchLog{}
		}
	}
	if v1, ok := d.GetOk("switch_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("switch_profile", sv)
				diags = append(diags, e)
			}
			obj.SwitchProfile = &v2
		}
	}
	if v1, ok := d.GetOk("tdr_supported"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("tdr_supported", sv)
				diags = append(diags, e)
			}
			obj.TdrSupported = &v2
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
	if v1, ok := d.GetOk("version"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("version", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Version = &tmp
		}
	}
	return &obj, diags
}
