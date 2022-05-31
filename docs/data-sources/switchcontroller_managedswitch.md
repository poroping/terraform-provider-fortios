---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_managedswitch"
description: |-
  Get information on a fortios Configure FortiSwitch devices that are managed by this FortiGate.
---

# Data Source: fortios_switchcontroller_managedswitch
Use this data source to get information on a fortios Configure FortiSwitch devices that are managed by this FortiGate.


## Example Usage

```hcl

```

## Argument Reference

* `switch_id` - (Required) Managed-switch id.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `access_profile` - FortiSwitch access profile.
* `delayed_restart_trigger` - Delayed restart triggered for this FortiSwitch.
* `description` - Description.
* `dhcp_server_access_list` - DHCP snooping server access list.
* `directly_connected` - Directly connected FortiSwitch.
* `dynamic_capability` - List of features this FortiSwitch supports (not configurable) that is sent to the FortiGate device for subsequent configuration initiated by the FortiGate device.
* `dynamically_discovered` - Dynamically discovered FortiSwitch.
* `firmware_provision` - Enable/disable provisioning of firmware to FortiSwitches on join connection.
* `firmware_provision_latest` - Enable/disable one-time automatic provisioning of the latest firmware version.
* `firmware_provision_version` - Firmware version to provision to this FortiSwitch on bootup (major.minor.build, i.e. 6.2.1234).
* `flow_identity` - Flow-tracking netflow ipfix switch identity in hex format(00000000-FFFFFFFF default=0).
* `fsw_wan1_admin` - FortiSwitch WAN1 admin status; enable to authorize the FortiSwitch as a managed switch.
* `fsw_wan1_peer` - FortiSwitch WAN1 peer port.
* `l3_discovered` - Layer 3 management discovered.
* `max_allowed_trunk_members` - FortiSwitch maximum allowed trunk members.
* `mclag_igmp_snooping_aware` - Enable/disable MCLAG IGMP-snooping awareness.
* `name` - Managed-switch name.
* `override_snmp_community` - Enable/disable overriding the global SNMP communities.
* `override_snmp_sysinfo` - Enable/disable overriding the global SNMP system information.
* `override_snmp_trap_threshold` - Enable/disable overriding the global SNMP trap threshold values.
* `override_snmp_user` - Enable/disable overriding the global SNMP users.
* `owner_vdom` - VDOM which owner of port belongs to.
* `poe_detection_type` - PoE detection type for FortiSwitch.
* `poe_lldp_detection` - Enable/disable PoE LLDP detection.
* `poe_pre_standard_detection` - Enable/disable PoE pre-standard detection.
* `pre_provisioned` - Pre-provisioned managed switch.
* `qos_drop_policy` - Set QoS drop-policy.
* `qos_red_probability` - Set QoS RED/WRED drop probability.
* `staged_image_version` - Staged image version for FortiSwitch.
* `switch_device_tag` - User definable label/tag.
* `switch_dhcp_opt43_key` - DHCP option43 key.
* `switch_id` - Managed-switch id.
* `switch_profile` - FortiSwitch profile.
* `tdr_supported` - TDR supported.
* `type` - Indication of switch type, physical or virtual.
* `version` - FortiSwitch version.
* `802_1x_settings` - Configuration method to edit FortiSwitch 802.1X global settings.The structure of `802_1x_settings` block is documented below.

The `802_1x_settings` block contains:

* `link_down_auth` - Authentication state to set if a link is down.
* `local_override` - Enable to override global 802.1X settings on individual FortiSwitches.
* `mab_reauth` - Enable or disable MAB reauthentication settings.
* `max_reauth_attempt` - Maximum number of authentication attempts (0 - 15, default = 3).
* `reauth_period` - Reauthentication time interval (1 - 1440 min, default = 60, 0 = disable).
* `tx_period` - 802.1X Tx period (seconds, default=30).
* `custom_command` - Configuration method to edit FortiSwitch commands to be pushed to this FortiSwitch device upon rebooting the FortiGate switch controller or the FortiSwitch.The structure of `custom_command` block is documented below.

The `custom_command` block contains:

* `command_entry` - List of FortiSwitch commands.
* `command_name` - Names of commands to be pushed to this FortiSwitch device, as configured under config switch-controller custom-command.
* `igmp_snooping` - Configure FortiSwitch IGMP snooping global settings.The structure of `igmp_snooping` block is documented below.

The `igmp_snooping` block contains:

* `aging_time` - Maximum time to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
* `flood_unknown_multicast` - Enable/disable unknown multicast flooding.
* `local_override` - Enable/disable overriding the global IGMP snooping configuration.
* `vlans` - Configure IGMP snooping VLAN.The structure of `vlans` block is documented below.

The `vlans` block contains:

* `proxy` - IGMP snooping proxy for the VLAN interface.
* `querier` - Enable/disable IGMP snooping querier for the VLAN interface.
* `querier_addr` - IGMP snooping querier address.
* `version` - IGMP snooping querying version.
* `vlan_name` - List of FortiSwitch VLANs.
* `ip_source_guard` - IP source guard.The structure of `ip_source_guard` block is documented below.

The `ip_source_guard` block contains:

* `description` - Description.
* `port` - Ingress interface to which source guard is bound.
* `binding_entry` - IP and MAC address configuration.The structure of `binding_entry` block is documented below.

The `binding_entry` block contains:

* `entry_name` - Configure binding pair.
* `ip` - Source IP for this rule.
* `mac` - MAC address for this rule.
* `mirror` - Configuration method to edit FortiSwitch packet mirror.The structure of `mirror` block is documented below.

The `mirror` block contains:

* `dst` - Destination port.
* `name` - Mirror name.
* `status` - Active/inactive mirror configuration.
* `switching_packet` - Enable/disable switching functionality when mirroring.
* `src_egress` - Source egress interfaces.The structure of `src_egress` block is documented below.

The `src_egress` block contains:

* `name` - Interface name.
* `src_ingress` - Source ingress interfaces.The structure of `src_ingress` block is documented below.

The `src_ingress` block contains:

* `name` - Interface name.
* `ports` - Managed-switch port list.The structure of `ports` block is documented below.

The `ports` block contains:

* `access_mode` - Access mode of the port.
* `aggregator_mode` - LACP member select mode.
* `allowed_vlans_all` - Enable/disable all defined vlans on this port.
* `arp_inspection_trust` - Trusted or untrusted dynamic ARP inspection.
* `bundle` - Enable/disable Link Aggregation Group (LAG) bundling for non-FortiLink interfaces.
* `description` - Description for port.
* `dhcp_snoop_option82_trust` - Enable/disable allowance of DHCP with option-82 on untrusted interface.
* `dhcp_snooping` - Trusted or untrusted DHCP-snooping interface.
* `discard_mode` - Configure discard mode for port.
* `edge_port` - Enable/disable this interface as an edge port, bridging connections between workstations and/or computers.
* `export_to` - Export managed-switch port to a tenant VDOM.
* `export_to_pool` - Switch controller export port to pool-list.
* `fec_capable` - FEC capable.
* `fec_state` - State of forward error correction.
* `fgt_peer_device_name` - FGT peer device name.
* `fgt_peer_port_name` - FGT peer port name.
* `fiber_port` - Fiber-port.
* `flags` - Port properties flags.
* `flap_duration` - Period over which flap events are calculated (seconds).
* `flap_rate` - Number of stage change events needed within flap-duration.
* `flap_timeout` - Flap guard disabling protection (min).
* `flapguard` - Enable/disable flap guard.
* `flow_control` - Flow control direction.
* `fortilink_port` - FortiLink uplink port.
* `igmp_snooping` - Set IGMP snooping mode for the physical port interface.
* `igmps_flood_reports` - Enable/disable flooding of IGMP reports to this interface when igmp-snooping enabled.
* `igmps_flood_traffic` - Enable/disable flooding of IGMP snooping traffic to this interface.
* `ip_source_guard` - Enable/disable IP source guard.
* `isl_local_trunk_name` - ISL local trunk name.
* `isl_peer_device_name` - ISL peer device name.
* `isl_peer_port_name` - ISL peer port name.
* `lacp_speed` - End Link Aggregation Control Protocol (LACP) messages every 30 seconds (slow) or every second (fast).
* `learning_limit` - Limit the number of dynamic MAC addresses on this Port (1 - 128, 0 = no limit, default).
* `lldp_profile` - LLDP port TLV profile.
* `lldp_status` - LLDP transmit and receive status.
* `loop_guard` - Enable/disable loop-guard on this interface, an STP optimization used to prevent network loops.
* `loop_guard_timeout` - Loop-guard timeout (0 - 120 min, default = 45).
* `mac_addr` - Port/Trunk MAC.
* `matched_dpp_intf_tags` - Matched interface tags in the dynamic port policy.
* `matched_dpp_policy` - Matched child policy in the dynamic port policy.
* `max_bundle` - Maximum size of LAG bundle (1 - 24, default = 24).
* `mclag` - Enable/disable multi-chassis link aggregation (MCLAG).
* `mclag_icl_port` - MCLAG-ICL port.
* `media_type` - Media type.
* `member_withdrawal_behavior` - Port behavior after it withdraws because of loss of control packets.
* `min_bundle` - Minimum size of LAG bundle (1 - 24, default = 1).
* `mode` - LACP mode: ignore and do not send control messages, or negotiate 802.3ad aggregation passively or actively.
* `p2p_port` - General peer to peer tunnel port.
* `packet_sample_rate` - Packet sampling rate (0 - 99999 p/sec).
* `packet_sampler` - Enable/disable packet sampling on this interface.
* `pause_meter` - Configure ingress pause metering rate, in kbps (default = 0, disabled).
* `pause_meter_resume` - Resume threshold for resuming traffic on ingress port.
* `poe_capable` - PoE capable.
* `poe_max_power` - PoE maximum power.
* `poe_pre_standard_detection` - Enable/disable PoE pre-standard detection.
* `poe_standard` - PoE standard supported.
* `poe_status` - Enable/disable PoE status.
* `port_name` - Switch port name.
* `port_number` - Port number.
* `port_owner` - Switch port name.
* `port_policy` - Switch controller dynamic port policy from available options.
* `port_prefix_type` - Port prefix type.
* `port_security_policy` - Switch controller authentication policy to apply to this managed switch from available options.
* `port_selection_criteria` - Algorithm for aggregate port selection.
* `ptp_policy` - PTP policy configuration.
* `qos_policy` - Switch controller QoS policy from available options.
* `rpvst_port` - Enable/disable inter-operability with rapid PVST on this interface.
* `sample_direction` - Packet sampling direction.
* `sflow_counter_interval` - sFlow sampling counter polling interval in seconds (0 - 255).
* `speed` - Switch port speed; default and available settings depend on hardware.
* `stacking_port` - Stacking port.
* `status` - Switch port admin status: up or down.
* `sticky_mac` - Enable or disable sticky-mac on the interface.
* `storm_control_policy` - Switch controller storm control policy from available options.
* `stp_bpdu_guard` - Enable/disable STP BPDU guard on this interface.
* `stp_bpdu_guard_timeout` - BPDU Guard disabling protection (0 - 120 min).
* `stp_root_guard` - Enable/disable STP root guard on this interface.
* `stp_state` - Enable/disable Spanning Tree Protocol (STP) on this interface.
* `switch_id` - Switch id.
* `type` - Interface type: physical or trunk port.
* `vlan` - Assign switch ports to a VLAN.
* `allowed_vlans` - Configure switch port tagged VLANs.The structure of `allowed_vlans` block is documented below.

The `allowed_vlans` block contains:

* `vlan_name` - VLAN name.
* `export_tags` - Configure export tag(s) for FortiSwitch port when exported to a virtual port pool.The structure of `export_tags` block is documented below.

The `export_tags` block contains:

* `tag_name` - FortiSwitch port tag name when exported to a virtual port pool.
* `interface_tags` - Tag(s) associated with the interface for various features including virtual port pool, dynamic port policy.The structure of `interface_tags` block is documented below.

The `interface_tags` block contains:

* `tag_name` - FortiSwitch port tag name when exported to a virtual port pool or matched to dynamic port policy.
* `members` - Aggregated LAG bundle interfaces.The structure of `members` block is documented below.

The `members` block contains:

* `member_name` - Interface name from available options.
* `untagged_vlans` - Configure switch port untagged VLANs.The structure of `untagged_vlans` block is documented below.

The `untagged_vlans` block contains:

* `vlan_name` - VLAN name.
* `remote_log` - Configure logging by FortiSwitch device to a remote syslog server.The structure of `remote_log` block is documented below.

The `remote_log` block contains:

* `csv` - Enable/disable comma-separated value (CSV) strings.
* `facility` - Facility to log to remote syslog server.
* `name` - Remote log name.
* `port` - Remote syslog server listening port.
* `server` - IPv4 address of the remote syslog server.
* `severity` - Severity of logs to be transferred to remote log server.
* `status` - Enable/disable logging by FortiSwitch device to a remote syslog server.
* `snmp_community` - Configuration method to edit Simple Network Management Protocol (SNMP) communities.The structure of `snmp_community` block is documented below.

The `snmp_community` block contains:

* `events` - SNMP notifications (traps) to send.
* `id` - SNMP community ID.
* `name` - SNMP community name.
* `query_v1_port` - SNMP v1 query port (default = 161).
* `query_v1_status` - Enable/disable SNMP v1 queries.
* `query_v2c_port` - SNMP v2c query port (default = 161).
* `query_v2c_status` - Enable/disable SNMP v2c queries.
* `status` - Enable/disable this SNMP community.
* `trap_v1_lport` - SNMP v2c trap local port (default = 162).
* `trap_v1_rport` - SNMP v2c trap remote port (default = 162).
* `trap_v1_status` - Enable/disable SNMP v1 traps.
* `trap_v2c_lport` - SNMP v2c trap local port (default = 162).
* `trap_v2c_rport` - SNMP v2c trap remote port (default = 162).
* `trap_v2c_status` - Enable/disable SNMP v2c traps.
* `hosts` - Configure IPv4 SNMP managers (hosts).The structure of `hosts` block is documented below.

The `hosts` block contains:

* `id` - Host entry ID.
* `ip` - IPv4 address of the SNMP manager (host).
* `snmp_sysinfo` - Configuration method to edit Simple Network Management Protocol (SNMP) system info.The structure of `snmp_sysinfo` block is documented below.

The `snmp_sysinfo` block contains:

* `contact_info` - Contact information.
* `description` - System description.
* `engine_id` - Local SNMP engine ID string (max 24 char).
* `location` - System location.
* `status` - Enable/disable SNMP.
* `snmp_trap_threshold` - Configuration method to edit Simple Network Management Protocol (SNMP) trap threshold values.The structure of `snmp_trap_threshold` block is documented below.

The `snmp_trap_threshold` block contains:

* `trap_high_cpu_threshold` - CPU usage when trap is sent.
* `trap_log_full_threshold` - Log disk usage when trap is sent.
* `trap_low_memory_threshold` - Memory usage when trap is sent.
* `snmp_user` - Configuration method to edit Simple Network Management Protocol (SNMP) users.The structure of `snmp_user` block is documented below.

The `snmp_user` block contains:

* `auth_proto` - Authentication protocol.
* `auth_pwd` - Password for authentication protocol.
* `name` - SNMP user name.
* `priv_proto` - Privacy (encryption) protocol.
* `priv_pwd` - Password for privacy (encryption) protocol.
* `queries` - Enable/disable SNMP queries for this user.
* `query_port` - SNMPv3 query port (default = 161).
* `security_level` - Security level for message authentication and encryption.
* `static_mac` - Configuration method to edit FortiSwitch Static and Sticky MAC.The structure of `static_mac` block is documented below.

The `static_mac` block contains:

* `description` - Description.
* `id` - ID.
* `interface` - Interface name.
* `mac` - MAC address.
* `type` - Type.
* `vlan` - Vlan.
* `storm_control` - Configuration method to edit FortiSwitch storm control for measuring traffic activity using data rates to prevent traffic disruption.The structure of `storm_control` block is documented below.

The `storm_control` block contains:

* `broadcast` - Enable/disable storm control to drop broadcast traffic.
* `local_override` - Enable to override global FortiSwitch storm control settings for this FortiSwitch.
* `rate` - Rate in packets per second at which storm traffic is controlled (1 - 10000000, default = 500). Storm control drops excess traffic data rates beyond this threshold.
* `unknown_multicast` - Enable/disable storm control to drop unknown multicast traffic.
* `unknown_unicast` - Enable/disable storm control to drop unknown unicast traffic.
* `stp_instance` - Configuration method to edit Spanning Tree Protocol (STP) instances.The structure of `stp_instance` block is documented below.

The `stp_instance` block contains:

* `id` - Instance ID.
* `priority` - Priority.
* `stp_settings` - Configuration method to edit Spanning Tree Protocol (STP) settings used to prevent bridge loops.The structure of `stp_settings` block is documented below.

The `stp_settings` block contains:

* `forward_time` - Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
* `hello_time` - Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
* `local_override` - Enable to configure local STP settings that override global STP settings.
* `max_age` - Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
* `max_hops` - Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
* `name` - Name of local STP settings configuration.
* `pending_timer` - Pending time (1 - 15 sec, default = 4).
* `revision` - STP revision number (0 - 65535).
* `switch_log` - Configuration method to edit FortiSwitch logging settings (logs are transferred to and inserted into the FortiGate event log).The structure of `switch_log` block is documented below.

The `switch_log` block contains:

* `local_override` - Enable to configure local logging settings that override global logging settings.
* `severity` - Severity of FortiSwitch logs that are added to the FortiGate event log.
* `status` - Enable/disable adding FortiSwitch logs to the FortiGate event log.
