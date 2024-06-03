// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiSwitch devices that are managed by this FortiGate.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSwitchControllerManagedSwitch() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiSwitch devices that are managed by this FortiGate.",

		ReadContext: dataSourceSwitchControllerManagedSwitchRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"802_1x_settings": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit FortiSwitch 802.1X global settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"link_down_auth": {
							Type:        schema.TypeString,
							Description: "Authentication state to set if a link is down.",
							Computed:    true,
						},
						"local_override": {
							Type:        schema.TypeString,
							Description: "Enable to override global 802.1X settings on individual FortiSwitches.",
							Computed:    true,
						},
						"mab_reauth": {
							Type:        schema.TypeString,
							Description: "Enable or disable MAB reauthentication settings.",
							Computed:    true,
						},
						"max_reauth_attempt": {
							Type:        schema.TypeInt,
							Description: "Maximum number of authentication attempts (0 - 15, default = 3).",
							Computed:    true,
						},
						"reauth_period": {
							Type:        schema.TypeInt,
							Description: "Reauthentication time interval (1 - 1440 min, default = 60, 0 = disable).",
							Computed:    true,
						},
						"tx_period": {
							Type:        schema.TypeInt,
							Description: "802.1X Tx period (seconds, default=30).",
							Computed:    true,
						},
					},
				},
			},
			"access_profile": {
				Type:        schema.TypeString,
				Description: "FortiSwitch access profile.",
				Computed:    true,
			},
			"custom_command": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit FortiSwitch commands to be pushed to this FortiSwitch device upon rebooting the FortiGate switch controller or the FortiSwitch.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command_entry": {
							Type:        schema.TypeString,
							Description: "List of FortiSwitch commands.",
							Computed:    true,
						},
						"command_name": {
							Type:        schema.TypeString,
							Description: "Names of commands to be pushed to this FortiSwitch device, as configured under config switch-controller custom-command.",
							Computed:    true,
						},
					},
				},
			},
			"delayed_restart_trigger": {
				Type:        schema.TypeInt,
				Description: "Delayed restart triggered for this FortiSwitch.",
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description.",
				Computed:    true,
			},
			"dhcp_server_access_list": {
				Type:        schema.TypeString,
				Description: "DHCP snooping server access list.",
				Computed:    true,
			},
			"dhcp_snooping_static_client": {
				Type:        schema.TypeList,
				Description: "Configure FortiSwitch DHCP snooping static clients.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:        schema.TypeString,
							Description: "Client static IP address.",
							Computed:    true,
						},
						"mac": {
							Type:        schema.TypeString,
							Description: "Client MAC address.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Client name.",
							Computed:    true,
						},
						"port": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
						"vlan": {
							Type:        schema.TypeString,
							Description: "VLAN name.",
							Computed:    true,
						},
					},
				},
			},
			"directly_connected": {
				Type:        schema.TypeInt,
				Description: "Directly connected FortiSwitch.",
				Computed:    true,
			},
			"dynamic_capability": {
				Type:        schema.TypeString,
				Description: "List of features this FortiSwitch supports (not configurable) that is sent to the FortiGate device for subsequent configuration initiated by the FortiGate device.",
				Computed:    true,
			},
			"dynamically_discovered": {
				Type:        schema.TypeInt,
				Description: "Dynamically discovered FortiSwitch.",
				Computed:    true,
			},
			"firmware_provision": {
				Type:        schema.TypeString,
				Description: "Enable/disable provisioning of firmware to FortiSwitches on join connection.",
				Computed:    true,
			},
			"firmware_provision_latest": {
				Type:        schema.TypeString,
				Description: "Enable/disable one-time automatic provisioning of the latest firmware version.",
				Computed:    true,
			},
			"firmware_provision_version": {
				Type:        schema.TypeString,
				Description: "Firmware version to provision to this FortiSwitch on bootup (major.minor.build, i.e. 6.2.1234).",
				Computed:    true,
			},
			"flow_identity": {
				Type:        schema.TypeString,
				Description: "Flow-tracking netflow ipfix switch identity in hex format(00000000-FFFFFFFF default=0).",
				Computed:    true,
			},
			"fsw_wan1_admin": {
				Type:        schema.TypeString,
				Description: "FortiSwitch WAN1 admin status; enable to authorize the FortiSwitch as a managed switch.",
				Computed:    true,
			},
			"fsw_wan1_peer": {
				Type:        schema.TypeString,
				Description: "FortiSwitch WAN1 peer port.",
				Computed:    true,
			},
			"igmp_snooping": {
				Type:        schema.TypeList,
				Description: "Configure FortiSwitch IGMP snooping global settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aging_time": {
							Type:        schema.TypeInt,
							Description: "Maximum time to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).",
							Computed:    true,
						},
						"flood_unknown_multicast": {
							Type:        schema.TypeString,
							Description: "Enable/disable unknown multicast flooding.",
							Computed:    true,
						},
						"local_override": {
							Type:        schema.TypeString,
							Description: "Enable/disable overriding the global IGMP snooping configuration.",
							Computed:    true,
						},
						"vlans": {
							Type:        schema.TypeList,
							Description: "Configure IGMP snooping VLAN.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"proxy": {
										Type:        schema.TypeString,
										Description: "IGMP snooping proxy for the VLAN interface.",
										Computed:    true,
									},
									"querier": {
										Type:        schema.TypeString,
										Description: "Enable/disable IGMP snooping querier for the VLAN interface.",
										Computed:    true,
									},
									"querier_addr": {
										Type:        schema.TypeString,
										Description: "IGMP snooping querier address.",
										Computed:    true,
									},
									"version": {
										Type:        schema.TypeInt,
										Description: "IGMP snooping querying version.",
										Computed:    true,
									},
									"vlan_name": {
										Type:        schema.TypeString,
										Description: "List of FortiSwitch VLANs.",
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
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"binding_entry": {
							Type:        schema.TypeList,
							Description: "IP and MAC address configuration.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"entry_name": {
										Type:        schema.TypeString,
										Description: "Configure binding pair.",
										Computed:    true,
									},
									"ip": {
										Type:        schema.TypeString,
										Description: "Source IP for this rule.",
										Computed:    true,
									},
									"mac": {
										Type:        schema.TypeString,
										Description: "MAC address for this rule.",
										Computed:    true,
									},
								},
							},
						},
						"description": {
							Type:        schema.TypeString,
							Description: "Description.",
							Computed:    true,
						},
						"port": {
							Type:        schema.TypeString,
							Description: "Ingress interface to which source guard is bound.",
							Computed:    true,
						},
					},
				},
			},
			"l3_discovered": {
				Type:        schema.TypeInt,
				Description: "Layer 3 management discovered.",
				Computed:    true,
			},
			"max_allowed_trunk_members": {
				Type:        schema.TypeInt,
				Description: "FortiSwitch maximum allowed trunk members.",
				Computed:    true,
			},
			"mclag_igmp_snooping_aware": {
				Type:        schema.TypeString,
				Description: "Enable/disable MCLAG IGMP-snooping awareness.",
				Computed:    true,
			},
			"mirror": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit FortiSwitch packet mirror.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst": {
							Type:        schema.TypeString,
							Description: "Destination port.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Mirror name.",
							Computed:    true,
						},
						"src_egress": {
							Type:        schema.TypeList,
							Description: "Source egress interfaces.",
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
						"src_ingress": {
							Type:        schema.TypeList,
							Description: "Source ingress interfaces.",
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
						"status": {
							Type:        schema.TypeString,
							Description: "Active/inactive mirror configuration.",
							Computed:    true,
						},
						"switching_packet": {
							Type:        schema.TypeString,
							Description: "Enable/disable switching functionality when mirroring.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Managed-switch name.",
				Computed:    true,
			},
			"override_snmp_community": {
				Type:        schema.TypeString,
				Description: "Enable/disable overriding the global SNMP communities.",
				Computed:    true,
			},
			"override_snmp_sysinfo": {
				Type:        schema.TypeString,
				Description: "Enable/disable overriding the global SNMP system information.",
				Computed:    true,
			},
			"override_snmp_trap_threshold": {
				Type:        schema.TypeString,
				Description: "Enable/disable overriding the global SNMP trap threshold values.",
				Computed:    true,
			},
			"override_snmp_user": {
				Type:        schema.TypeString,
				Description: "Enable/disable overriding the global SNMP users.",
				Computed:    true,
			},
			"owner_vdom": {
				Type:        schema.TypeString,
				Description: "VDOM which owner of port belongs to.",
				Computed:    true,
			},
			"poe_detection_type": {
				Type:        schema.TypeInt,
				Description: "PoE detection type for FortiSwitch.",
				Computed:    true,
			},
			"poe_lldp_detection": {
				Type:        schema.TypeString,
				Description: "Enable/disable PoE LLDP detection.",
				Computed:    true,
			},
			"poe_pre_standard_detection": {
				Type:        schema.TypeString,
				Description: "Enable/disable PoE pre-standard detection.",
				Computed:    true,
			},
			"ports": {
				Type:        schema.TypeList,
				Description: "Managed-switch port list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_mode": {
							Type:        schema.TypeString,
							Description: "Access mode of the port.",
							Computed:    true,
						},
						"aggregator_mode": {
							Type:        schema.TypeString,
							Description: "LACP member select mode.",
							Computed:    true,
						},
						"allowed_vlans": {
							Type:        schema.TypeList,
							Description: "Configure switch port tagged VLANs.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vlan_name": {
										Type:        schema.TypeString,
										Description: "VLAN name.",
										Computed:    true,
									},
								},
							},
						},
						"allowed_vlans_all": {
							Type:        schema.TypeString,
							Description: "Enable/disable all defined vlans on this port.",
							Computed:    true,
						},
						"arp_inspection_trust": {
							Type:        schema.TypeString,
							Description: "Trusted or untrusted dynamic ARP inspection.",
							Computed:    true,
						},
						"bundle": {
							Type:        schema.TypeString,
							Description: "Enable/disable Link Aggregation Group (LAG) bundling for non-FortiLink interfaces.",
							Computed:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "Description for port.",
							Computed:    true,
						},
						"dhcp_snoop_option82_trust": {
							Type:        schema.TypeString,
							Description: "Enable/disable allowance of DHCP with option-82 on untrusted interface.",
							Computed:    true,
						},
						"dhcp_snooping": {
							Type:        schema.TypeString,
							Description: "Trusted or untrusted DHCP-snooping interface.",
							Computed:    true,
						},
						"discard_mode": {
							Type:        schema.TypeString,
							Description: "Configure discard mode for port.",
							Computed:    true,
						},
						"edge_port": {
							Type:        schema.TypeString,
							Description: "Enable/disable this interface as an edge port, bridging connections between workstations and/or computers.",
							Computed:    true,
						},
						"export_tags": {
							Type:        schema.TypeList,
							Description: "Configure export tag(s) for FortiSwitch port when exported to a virtual port pool.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag_name": {
										Type:        schema.TypeString,
										Description: "FortiSwitch port tag name when exported to a virtual port pool.",
										Computed:    true,
									},
								},
							},
						},
						"export_to": {
							Type:        schema.TypeString,
							Description: "Export managed-switch port to a tenant VDOM.",
							Computed:    true,
						},
						"export_to_pool": {
							Type:        schema.TypeString,
							Description: "Switch controller export port to pool-list.",
							Computed:    true,
						},
						"fec_capable": {
							Type:        schema.TypeInt,
							Description: "FEC capable.",
							Computed:    true,
						},
						"fec_state": {
							Type:        schema.TypeString,
							Description: "State of forward error correction.",
							Computed:    true,
						},
						"fgt_peer_device_name": {
							Type:        schema.TypeString,
							Description: "FGT peer device name.",
							Computed:    true,
						},
						"fgt_peer_port_name": {
							Type:        schema.TypeString,
							Description: "FGT peer port name.",
							Computed:    true,
						},
						"fiber_port": {
							Type:        schema.TypeInt,
							Description: "Fiber-port.",
							Computed:    true,
						},
						"flags": {
							Type:        schema.TypeInt,
							Description: "Port properties flags.",
							Computed:    true,
						},
						"flap_duration": {
							Type:        schema.TypeInt,
							Description: "Period over which flap events are calculated (seconds).",
							Computed:    true,
						},
						"flap_rate": {
							Type:        schema.TypeInt,
							Description: "Number of stage change events needed within flap-duration.",
							Computed:    true,
						},
						"flap_timeout": {
							Type:        schema.TypeInt,
							Description: "Flap guard disabling protection (min).",
							Computed:    true,
						},
						"flapguard": {
							Type:        schema.TypeString,
							Description: "Enable/disable flap guard.",
							Computed:    true,
						},
						"flow_control": {
							Type:        schema.TypeString,
							Description: "Flow control direction.",
							Computed:    true,
						},
						"fortilink_port": {
							Type:        schema.TypeInt,
							Description: "FortiLink uplink port.",
							Computed:    true,
						},
						"igmp_snooping": {
							Type:        schema.TypeString,
							Description: "Set IGMP snooping mode for the physical port interface.",
							Computed:    true,
						},
						"igmp_snooping_flood_reports": {
							Type:        schema.TypeString,
							Description: "Enable/disable flooding of IGMP reports to this interface when igmp-snooping enabled.",
							Computed:    true,
						},
						"igmps_flood_reports": {
							Type:        schema.TypeString,
							Description: "Enable/disable flooding of IGMP reports to this interface when igmp-snooping enabled.",
							Computed:    true,
						},
						"igmps_flood_traffic": {
							Type:        schema.TypeString,
							Description: "Enable/disable flooding of IGMP snooping traffic to this interface.",
							Computed:    true,
						},
						"interface_tags": {
							Type:        schema.TypeList,
							Description: "Tag(s) associated with the interface for various features including virtual port pool, dynamic port policy.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag_name": {
										Type:        schema.TypeString,
										Description: "FortiSwitch port tag name when exported to a virtual port pool or matched to dynamic port policy.",
										Computed:    true,
									},
								},
							},
						},
						"ip_source_guard": {
							Type:        schema.TypeString,
							Description: "Enable/disable IP source guard.",
							Computed:    true,
						},
						"isl_local_trunk_name": {
							Type:        schema.TypeString,
							Description: "ISL local trunk name.",
							Computed:    true,
						},
						"isl_peer_device_name": {
							Type:        schema.TypeString,
							Description: "ISL peer device name.",
							Computed:    true,
						},
						"isl_peer_port_name": {
							Type:        schema.TypeString,
							Description: "ISL peer port name.",
							Computed:    true,
						},
						"lacp_speed": {
							Type:        schema.TypeString,
							Description: "End Link Aggregation Control Protocol (LACP) messages every 30 seconds (slow) or every second (fast).",
							Computed:    true,
						},
						"learning_limit": {
							Type:        schema.TypeInt,
							Description: "Limit the number of dynamic MAC addresses on this Port (1 - 128, 0 = no limit, default).",
							Computed:    true,
						},
						"lldp_profile": {
							Type:        schema.TypeString,
							Description: "LLDP port TLV profile.",
							Computed:    true,
						},
						"lldp_status": {
							Type:        schema.TypeString,
							Description: "LLDP transmit and receive status.",
							Computed:    true,
						},
						"loop_guard": {
							Type:        schema.TypeString,
							Description: "Enable/disable loop-guard on this interface, an STP optimization used to prevent network loops.",
							Computed:    true,
						},
						"loop_guard_timeout": {
							Type:        schema.TypeInt,
							Description: "Loop-guard timeout (0 - 120 min, default = 45).",
							Computed:    true,
						},
						"mac_addr": {
							Type:        schema.TypeString,
							Description: "Port/Trunk MAC.",
							Computed:    true,
						},
						"matched_dpp_intf_tags": {
							Type:        schema.TypeString,
							Description: "Matched interface tags in the dynamic port policy.",
							Computed:    true,
						},
						"matched_dpp_policy": {
							Type:        schema.TypeString,
							Description: "Matched child policy in the dynamic port policy.",
							Computed:    true,
						},
						"max_bundle": {
							Type:        schema.TypeInt,
							Description: "Maximum size of LAG bundle (1 - 24, default = 24).",
							Computed:    true,
						},
						"mcast_snooping_flood_traffic": {
							Type:        schema.TypeString,
							Description: "Enable/disable flooding of IGMP snooping traffic to this interface.",
							Computed:    true,
						},
						"mclag": {
							Type:        schema.TypeString,
							Description: "Enable/disable multi-chassis link aggregation (MCLAG).",
							Computed:    true,
						},
						"mclag_icl_port": {
							Type:        schema.TypeInt,
							Description: "MCLAG-ICL port.",
							Computed:    true,
						},
						"media_type": {
							Type:        schema.TypeString,
							Description: "Media type.",
							Computed:    true,
						},
						"member_withdrawal_behavior": {
							Type:        schema.TypeString,
							Description: "Port behavior after it withdraws because of loss of control packets.",
							Computed:    true,
						},
						"members": {
							Type:        schema.TypeList,
							Description: "Aggregated LAG bundle interfaces.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"member_name": {
										Type:        schema.TypeString,
										Description: "Interface name from available options.",
										Computed:    true,
									},
								},
							},
						},
						"min_bundle": {
							Type:        schema.TypeInt,
							Description: "Minimum size of LAG bundle (1 - 24, default = 1).",
							Computed:    true,
						},
						"mode": {
							Type:        schema.TypeString,
							Description: "LACP mode: ignore and do not send control messages, or negotiate 802.3ad aggregation passively or actively.",
							Computed:    true,
						},
						"p2p_port": {
							Type:        schema.TypeInt,
							Description: "General peer to peer tunnel port.",
							Computed:    true,
						},
						"packet_sample_rate": {
							Type:        schema.TypeInt,
							Description: "Packet sampling rate (0 - 99999 p/sec).",
							Computed:    true,
						},
						"packet_sampler": {
							Type:        schema.TypeString,
							Description: "Enable/disable packet sampling on this interface.",
							Computed:    true,
						},
						"pause_meter": {
							Type:        schema.TypeInt,
							Description: "Configure ingress pause metering rate, in kbps (default = 0, disabled).",
							Computed:    true,
						},
						"pause_meter_resume": {
							Type:        schema.TypeString,
							Description: "Resume threshold for resuming traffic on ingress port.",
							Computed:    true,
						},
						"poe_capable": {
							Type:        schema.TypeInt,
							Description: "PoE capable.",
							Computed:    true,
						},
						"poe_max_power": {
							Type:        schema.TypeString,
							Description: "PoE maximum power.",
							Computed:    true,
						},
						"poe_mode_bt_cabable": {
							Type:        schema.TypeInt,
							Description: "PoE mode IEEE 802.3BT capable.",
							Computed:    true,
						},
						"poe_port_mode": {
							Type:        schema.TypeString,
							Description: "Configure PoE port mode.",
							Computed:    true,
						},
						"poe_port_power": {
							Type:        schema.TypeString,
							Description: "Configure PoE port power.",
							Computed:    true,
						},
						"poe_port_priority": {
							Type:        schema.TypeString,
							Description: "Configure PoE port priority.",
							Computed:    true,
						},
						"poe_pre_standard_detection": {
							Type:        schema.TypeString,
							Description: "Enable/disable PoE pre-standard detection.",
							Computed:    true,
						},
						"poe_standard": {
							Type:        schema.TypeString,
							Description: "PoE standard supported.",
							Computed:    true,
						},
						"poe_status": {
							Type:        schema.TypeString,
							Description: "Enable/disable PoE status.",
							Computed:    true,
						},
						"port_name": {
							Type:        schema.TypeString,
							Description: "Switch port name.",
							Computed:    true,
						},
						"port_number": {
							Type:        schema.TypeInt,
							Description: "Port number.",
							Computed:    true,
						},
						"port_owner": {
							Type:        schema.TypeString,
							Description: "Switch port name.",
							Computed:    true,
						},
						"port_policy": {
							Type:        schema.TypeString,
							Description: "Switch controller dynamic port policy from available options.",
							Computed:    true,
						},
						"port_prefix_type": {
							Type:        schema.TypeInt,
							Description: "Port prefix type.",
							Computed:    true,
						},
						"port_security_policy": {
							Type:        schema.TypeString,
							Description: "Switch controller authentication policy to apply to this managed switch from available options.",
							Computed:    true,
						},
						"port_selection_criteria": {
							Type:        schema.TypeString,
							Description: "Algorithm for aggregate port selection.",
							Computed:    true,
						},
						"ptp_policy": {
							Type:        schema.TypeString,
							Description: "PTP policy configuration.",
							Computed:    true,
						},
						"qos_policy": {
							Type:        schema.TypeString,
							Description: "Switch controller QoS policy from available options.",
							Computed:    true,
						},
						"rpvst_port": {
							Type:        schema.TypeString,
							Description: "Enable/disable inter-operability with rapid PVST on this interface.",
							Computed:    true,
						},
						"sample_direction": {
							Type:        schema.TypeString,
							Description: "Packet sampling direction.",
							Computed:    true,
						},
						"sflow_counter_interval": {
							Type:        schema.TypeInt,
							Description: "sFlow sampling counter polling interval in seconds (0 - 255).",
							Computed:    true,
						},
						"speed": {
							Type:        schema.TypeString,
							Description: "Switch port speed; default and available settings depend on hardware.",
							Computed:    true,
						},
						"stacking_port": {
							Type:        schema.TypeInt,
							Description: "Stacking port.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Switch port admin status: up or down.",
							Computed:    true,
						},
						"sticky_mac": {
							Type:        schema.TypeString,
							Description: "Enable or disable sticky-mac on the interface.",
							Computed:    true,
						},
						"storm_control_policy": {
							Type:        schema.TypeString,
							Description: "Switch controller storm control policy from available options.",
							Computed:    true,
						},
						"stp_bpdu_guard": {
							Type:        schema.TypeString,
							Description: "Enable/disable STP BPDU guard on this interface.",
							Computed:    true,
						},
						"stp_bpdu_guard_timeout": {
							Type:        schema.TypeInt,
							Description: "BPDU Guard disabling protection (0 - 120 min).",
							Computed:    true,
						},
						"stp_root_guard": {
							Type:        schema.TypeString,
							Description: "Enable/disable STP root guard on this interface.",
							Computed:    true,
						},
						"stp_state": {
							Type:        schema.TypeString,
							Description: "Enable/disable Spanning Tree Protocol (STP) on this interface.",
							Computed:    true,
						},
						"switch_id": {
							Type:        schema.TypeString,
							Description: "Switch id.",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "Interface type: physical or trunk port.",
							Computed:    true,
						},
						"untagged_vlans": {
							Type:        schema.TypeList,
							Description: "Configure switch port untagged VLANs.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vlan_name": {
										Type:        schema.TypeString,
										Description: "VLAN name.",
										Computed:    true,
									},
								},
							},
						},
						"vlan": {
							Type:        schema.TypeString,
							Description: "Assign switch ports to a VLAN.",
							Computed:    true,
						},
					},
				},
			},
			"pre_provisioned": {
				Type:        schema.TypeInt,
				Description: "Pre-provisioned managed switch.",
				Computed:    true,
			},
			"qos_drop_policy": {
				Type:        schema.TypeString,
				Description: "Set QoS drop-policy.",
				Computed:    true,
			},
			"qos_red_probability": {
				Type:        schema.TypeInt,
				Description: "Set QoS RED/WRED drop probability.",
				Computed:    true,
			},
			"remote_log": {
				Type:        schema.TypeList,
				Description: "Configure logging by FortiSwitch device to a remote syslog server.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"csv": {
							Type:        schema.TypeString,
							Description: "Enable/disable comma-separated value (CSV) strings.",
							Computed:    true,
						},
						"facility": {
							Type:        schema.TypeString,
							Description: "Facility to log to remote syslog server.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Remote log name.",
							Computed:    true,
						},
						"port": {
							Type:        schema.TypeInt,
							Description: "Remote syslog server listening port.",
							Computed:    true,
						},
						"server": {
							Type:        schema.TypeString,
							Description: "IPv4 address of the remote syslog server.",
							Computed:    true,
						},
						"severity": {
							Type:        schema.TypeString,
							Description: "Severity of logs to be transferred to remote log server.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging by FortiSwitch device to a remote syslog server.",
							Computed:    true,
						},
					},
				},
			},
			"snmp_community": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit Simple Network Management Protocol (SNMP) communities.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"events": {
							Type:        schema.TypeString,
							Description: "SNMP notifications (traps) to send.",
							Computed:    true,
						},
						"hosts": {
							Type:        schema.TypeList,
							Description: "Configure IPv4 SNMP managers (hosts).",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Description: "Host entry ID.",
										Computed:    true,
									},
									"ip": {
										Type:        schema.TypeString,
										Description: "IPv4 address of the SNMP manager (host).",
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "SNMP community ID.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "SNMP community name.",
							Computed:    true,
						},
						"query_v1_port": {
							Type:        schema.TypeInt,
							Description: "SNMP v1 query port (default = 161).",
							Computed:    true,
						},
						"query_v1_status": {
							Type:        schema.TypeString,
							Description: "Enable/disable SNMP v1 queries.",
							Computed:    true,
						},
						"query_v2c_port": {
							Type:        schema.TypeInt,
							Description: "SNMP v2c query port (default = 161).",
							Computed:    true,
						},
						"query_v2c_status": {
							Type:        schema.TypeString,
							Description: "Enable/disable SNMP v2c queries.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable this SNMP community.",
							Computed:    true,
						},
						"trap_v1_lport": {
							Type:        schema.TypeInt,
							Description: "SNMP v2c trap local port (default = 162).",
							Computed:    true,
						},
						"trap_v1_rport": {
							Type:        schema.TypeInt,
							Description: "SNMP v2c trap remote port (default = 162).",
							Computed:    true,
						},
						"trap_v1_status": {
							Type:        schema.TypeString,
							Description: "Enable/disable SNMP v1 traps.",
							Computed:    true,
						},
						"trap_v2c_lport": {
							Type:        schema.TypeInt,
							Description: "SNMP v2c trap local port (default = 162).",
							Computed:    true,
						},
						"trap_v2c_rport": {
							Type:        schema.TypeInt,
							Description: "SNMP v2c trap remote port (default = 162).",
							Computed:    true,
						},
						"trap_v2c_status": {
							Type:        schema.TypeString,
							Description: "Enable/disable SNMP v2c traps.",
							Computed:    true,
						},
					},
				},
			},
			"snmp_sysinfo": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit Simple Network Management Protocol (SNMP) system info.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"contact_info": {
							Type:        schema.TypeString,
							Description: "Contact information.",
							Computed:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "System description.",
							Computed:    true,
						},
						"engine_id": {
							Type:        schema.TypeString,
							Description: "Local SNMP engine ID string (max 24 char).",
							Computed:    true,
						},
						"location": {
							Type:        schema.TypeString,
							Description: "System location.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable SNMP.",
							Computed:    true,
						},
					},
				},
			},
			"snmp_trap_threshold": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit Simple Network Management Protocol (SNMP) trap threshold values.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trap_high_cpu_threshold": {
							Type:        schema.TypeInt,
							Description: "CPU usage when trap is sent.",
							Computed:    true,
						},
						"trap_log_full_threshold": {
							Type:        schema.TypeInt,
							Description: "Log disk usage when trap is sent.",
							Computed:    true,
						},
						"trap_low_memory_threshold": {
							Type:        schema.TypeInt,
							Description: "Memory usage when trap is sent.",
							Computed:    true,
						},
					},
				},
			},
			"snmp_user": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit Simple Network Management Protocol (SNMP) users.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_proto": {
							Type:        schema.TypeString,
							Description: "Authentication protocol.",
							Computed:    true,
						},
						"auth_pwd": {
							Type:        schema.TypeString,
							Description: "Password for authentication protocol.",
							Computed:    true,
							Sensitive:   true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "SNMP user name.",
							Computed:    true,
						},
						"priv_proto": {
							Type:        schema.TypeString,
							Description: "Privacy (encryption) protocol.",
							Computed:    true,
						},
						"priv_pwd": {
							Type:        schema.TypeString,
							Description: "Password for privacy (encryption) protocol.",
							Computed:    true,
							Sensitive:   true,
						},
						"queries": {
							Type:        schema.TypeString,
							Description: "Enable/disable SNMP queries for this user.",
							Computed:    true,
						},
						"query_port": {
							Type:        schema.TypeInt,
							Description: "SNMPv3 query port (default = 161).",
							Computed:    true,
						},
						"security_level": {
							Type:        schema.TypeString,
							Description: "Security level for message authentication and encryption.",
							Computed:    true,
						},
					},
				},
			},
			"staged_image_version": {
				Type:        schema.TypeString,
				Description: "Staged image version for FortiSwitch.",
				Computed:    true,
			},
			"static_mac": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit FortiSwitch Static and Sticky MAC.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:        schema.TypeString,
							Description: "Description.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"interface": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
						"mac": {
							Type:        schema.TypeString,
							Description: "MAC address.",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "Type.",
							Computed:    true,
						},
						"vlan": {
							Type:        schema.TypeString,
							Description: "Vlan.",
							Computed:    true,
						},
					},
				},
			},
			"storm_control": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit FortiSwitch storm control for measuring traffic activity using data rates to prevent traffic disruption.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"broadcast": {
							Type:        schema.TypeString,
							Description: "Enable/disable storm control to drop broadcast traffic.",
							Computed:    true,
						},
						"local_override": {
							Type:        schema.TypeString,
							Description: "Enable to override global FortiSwitch storm control settings for this FortiSwitch.",
							Computed:    true,
						},
						"rate": {
							Type:        schema.TypeInt,
							Description: "Rate in packets per second at which storm traffic is controlled (1 - 10000000, default = 500). Storm control drops excess traffic data rates beyond this threshold.",
							Computed:    true,
						},
						"unknown_multicast": {
							Type:        schema.TypeString,
							Description: "Enable/disable storm control to drop unknown multicast traffic.",
							Computed:    true,
						},
						"unknown_unicast": {
							Type:        schema.TypeString,
							Description: "Enable/disable storm control to drop unknown unicast traffic.",
							Computed:    true,
						},
					},
				},
			},
			"stp_instance": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit Spanning Tree Protocol (STP) instances.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "Instance ID.",
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeString,
							Description: "Priority.",
							Computed:    true,
						},
					},
				},
			},
			"stp_settings": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit Spanning Tree Protocol (STP) settings used to prevent bridge loops.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"forward_time": {
							Type:        schema.TypeInt,
							Description: "Period of time a port is in listening and learning state (4 - 30 sec, default = 15).",
							Computed:    true,
						},
						"hello_time": {
							Type:        schema.TypeInt,
							Description: "Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).",
							Computed:    true,
						},
						"local_override": {
							Type:        schema.TypeString,
							Description: "Enable to configure local STP settings that override global STP settings.",
							Computed:    true,
						},
						"max_age": {
							Type:        schema.TypeInt,
							Description: "Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).",
							Computed:    true,
						},
						"max_hops": {
							Type:        schema.TypeInt,
							Description: "Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Name of local STP settings configuration.",
							Computed:    true,
						},
						"pending_timer": {
							Type:        schema.TypeInt,
							Description: "Pending time (1 - 15 sec, default = 4).",
							Computed:    true,
						},
						"revision": {
							Type:        schema.TypeInt,
							Description: "STP revision number (0 - 65535).",
							Computed:    true,
						},
					},
				},
			},
			"switch_device_tag": {
				Type:        schema.TypeString,
				Description: "User definable label/tag.",
				Computed:    true,
			},
			"switch_dhcp_opt43_key": {
				Type:        schema.TypeString,
				Description: "DHCP option43 key.",
				Computed:    true,
			},
			"switch_id": {
				Type:        schema.TypeString,
				Description: "Managed-switch id.",
				Required:    true,
			},
			"switch_log": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit FortiSwitch logging settings (logs are transferred to and inserted into the FortiGate event log).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"local_override": {
							Type:        schema.TypeString,
							Description: "Enable to configure local logging settings that override global logging settings.",
							Computed:    true,
						},
						"severity": {
							Type:        schema.TypeString,
							Description: "Severity of FortiSwitch logs that are added to the FortiGate event log.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable adding FortiSwitch logs to the FortiGate event log.",
							Computed:    true,
						},
					},
				},
			},
			"switch_profile": {
				Type:        schema.TypeString,
				Description: "FortiSwitch profile.",
				Computed:    true,
			},
			"tdr_supported": {
				Type:        schema.TypeString,
				Description: "TDR supported.",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Indication of switch type, physical or virtual.",
				Computed:    true,
			},
			"version": {
				Type:        schema.TypeInt,
				Description: "FortiSwitch version.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSwitchControllerManagedSwitchRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("switch_id")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadSwitchControllerManagedSwitch(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerManagedSwitch dataSource: %v", err)
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

	diags := refreshObjectSwitchControllerManagedSwitch(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
