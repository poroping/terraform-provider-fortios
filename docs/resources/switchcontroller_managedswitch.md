---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_managedswitch"
description: |-
  Configure FortiSwitch devices that are managed by this FortiGate.
---

## fortios_switchcontroller_managedswitch
Configure FortiSwitch devices that are managed by this FortiGate.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `switch_id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `access_profile` - FortiSwitch access profile. This attribute must reference one of the following datasources: `switch-controller.security-policy.local-access.name` .
* `delayed_restart_trigger` - Delayed restart triggered for this FortiSwitch.
* `description` - Description.
* `dhcp_server_access_list` - DHCP snooping server access list. Valid values: `global` `enable` `disable` .
* `directly_connected` - Directly connected FortiSwitch.
* `dynamic_capability` - List of features this FortiSwitch supports (not configurable) that is sent to the FortiGate device for subsequent configuration initiated by the FortiGate device.
* `dynamically_discovered` - Dynamically discovered FortiSwitch.
* `firmware_provision` - Enable/disable provisioning of firmware to FortiSwitches on join connection. Valid values: `enable` `disable` .
* `firmware_provision_version` - Firmware version to provision to this FortiSwitch on bootup (major.minor.build, i.e. 6.2.1234).
* `flow_identity` - Flow-tracking netflow ipfix switch identity in hex format(00000000-FFFFFFFF default=0).
* `fsw_wan1_admin` - FortiSwitch WAN1 admin status; enable to authorize the FortiSwitch as a managed switch. Valid values: `discovered` `disable` `enable` .
* `fsw_wan1_peer` - Fortiswitch WAN1 peer port. This attribute must reference one of the following datasources: `system.interface.name` .
* `l3_discovered` - Layer 3 management discovered.
* `max_allowed_trunk_members` - FortiSwitch maximum allowed trunk members.
* `mclag_igmp_snooping_aware` - Enable/disable MCLAG IGMP-snooping awareness. Valid values: `enable` `disable` .
* `name` - Managed-switch name.
* `override_snmp_community` - Enable/disable overriding the global SNMP communities. Valid values: `enable` `disable` .
* `override_snmp_sysinfo` - Enable/disable overriding the global SNMP system information. Valid values: `disable` `enable` .
* `override_snmp_trap_threshold` - Enable/disable overriding the global SNMP trap threshold values. Valid values: `enable` `disable` .
* `override_snmp_user` - Enable/disable overriding the global SNMP users. Valid values: `enable` `disable` .
* `owner_vdom` - VDOM which owner of port belongs to.
* `poe_detection_type` - PoE detection type for FortiSwitch.
* `poe_lldp_detection` - Enable/disable PoE LLDP detection. Valid values: `enable` `disable` .
* `poe_pre_standard_detection` - Enable/disable PoE pre-standard detection. Valid values: `enable` `disable` .
* `pre_provisioned` - Pre-provisioned managed switch.
* `qos_drop_policy` - Set QoS drop-policy. Valid values: `taildrop` `random-early-detection` .
* `qos_red_probability` - Set QoS RED/WRED drop probability.
* `staged_image_version` - Staged image version for FortiSwitch.
* `switch_device_tag` - User definable label/tag.
* `switch_dhcp_opt43_key` - DHCP option43 key.
* `switch_id` - Managed-switch id.
* `switch_profile` - FortiSwitch profile. This attribute must reference one of the following datasources: `switch-controller.switch-profile.name` .
* `tdr_supported` - TDR supported.
* `type` - Indication of switch type, physical or virtual. Valid values: `virtual` `physical` .
* `version` - FortiSwitch version.
* `802_1x_settings` - Configuration method to edit FortiSwitch 802.1X global settings. The structure of `802_1x_settings` block is documented below.

The `802_1x_settings` block contains:

* `link_down_auth` - Authentication state to set if a link is down. Valid values: `set-unauth` `no-action` .
* `local_override` - Enable to override global 802.1X settings on individual FortiSwitches. Valid values: `enable` `disable` .
* `max_reauth_attempt` - Maximum number of authentication attempts (0 - 15, default = 3).
* `reauth_period` - Reauthentication time interval (1 - 1440 min, default = 60, 0 = disable).
* `tx_period` - 802.1X Tx period (seconds, default=30).
* `custom_command` - Configuration method to edit FortiSwitch commands to be pushed to this FortiSwitch device upon rebooting the FortiGate switch controller or the FortiSwitch. The structure of `custom_command` block is documented below.

The `custom_command` block contains:

* `command_entry` - List of FortiSwitch commands.
* `command_name` - Names of commands to be pushed to this FortiSwitch device, as configured under config switch-controller custom-command. This attribute must reference one of the following datasources: `switch-controller.custom-command.command-name` .
* `igmp_snooping` - Configure FortiSwitch IGMP snooping global settings. The structure of `igmp_snooping` block is documented below.

The `igmp_snooping` block contains:

* `aging_time` - Maximum time to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
* `flood_unknown_multicast` - Enable/disable unknown multicast flooding. Valid values: `enable` `disable` .
* `local_override` - Enable/disable overriding the global IGMP snooping configuration. Valid values: `enable` `disable` .
* `vlans` - Configure IGMP snooping VLAN. The structure of `vlans` block is documented below.

The `vlans` block contains:

* `proxy` - IGMP snooping proxy for the VLAN interface. Valid values: `disable` `enable` `global` .
* `querier` - Enable/disable IGMP snooping querier for the VLAN interface. Valid values: `disable` `enable` .
* `querier_addr` - IGMP snooping querier address.
* `version` - IGMP snooping querier version.
* `vlan_name` - List of FortiSwitch VLANs. This attribute must reference one of the following datasources: `system.interface.name` .
* `ip_source_guard` - IP source guard. The structure of `ip_source_guard` block is documented below.

The `ip_source_guard` block contains:

* `description` - Description.
* `port` - Ingress interface to which source guard is bound.
* `binding_entry` - IP and MAC address configuration. The structure of `binding_entry` block is documented below.

The `binding_entry` block contains:

* `entry_name` - Configure binding pair.
* `ip` - Source IP for this rule.
* `mac` - MAC address for this rule.
* `mirror` - Configuration method to edit FortiSwitch packet mirror. The structure of `mirror` block is documented below.

The `mirror` block contains:

* `dst` - Destination port.
* `name` - Mirror name.
* `status` - Active/inactive mirror configuration. Valid values: `active` `inactive` .
* `switching_packet` - Enable/disable switching functionality when mirroring. Valid values: `enable` `disable` .
* `src_egress` - Source egress interfaces. The structure of `src_egress` block is documented below.

The `src_egress` block contains:

* `name` - Interface name.
* `src_ingress` - Source ingress interfaces. The structure of `src_ingress` block is documented below.

The `src_ingress` block contains:

* `name` - Interface name.
* `ports` - Managed-switch port list. The structure of `ports` block is documented below.

The `ports` block contains:

* `access_mode` - Access mode of the port. Valid values: `dynamic` `nac` `static` .
* `aggregator_mode` - LACP member select mode. Valid values: `bandwidth` `count` .
* `allowed_vlans_all` - Enable/disable all defined vlans on this port. Valid values: `enable` `disable` .
* `arp_inspection_trust` - Trusted or untrusted dynamic ARP inspection. Valid values: `untrusted` `trusted` .
* `bundle` - Enable/disable Link Aggregation Group (LAG) bundling for non-FortiLink interfaces. Valid values: `enable` `disable` .
* `description` - Description for port.
* `dhcp_snoop_option82_trust` - Enable/disable allowance of DHCP with option-82 on untrusted interface. Valid values: `enable` `disable` .
* `dhcp_snooping` - Trusted or untrusted DHCP-snooping interface. Valid values: `untrusted` `trusted` .
* `discard_mode` - Configure discard mode for port. Valid values: `none` `all-untagged` `all-tagged` .
* `edge_port` - Enable/disable this interface as an edge port, bridging connections between workstations and/or computers. Valid values: `enable` `disable` .
* `export_to` - Export managed-switch port to a tenant VDOM. This attribute must reference one of the following datasources: `system.vdom.name` .
* `export_to_pool` - Switch controller export port to pool-list. This attribute must reference one of the following datasources: `switch-controller.virtual-port-pool.name` .
* `fec_capable` - FEC capable.
* `fec_state` - State of forward error correction. Valid values: `disabled` `cl74` `cl91` .
* `fgt_peer_device_name` - FGT peer device name.
* `fgt_peer_port_name` - FGT peer port name.
* `fiber_port` - Fiber-port.
* `flags` - Port properties flags.
* `flow_control` - Flow control direction. Valid values: `disable` `tx` `rx` `both` .
* `fortilink_port` - FortiLink uplink port.
* `igmp_snooping` - Set IGMP snooping mode for the physical port interface. Valid values: `enable` `disable` .
* `igmps_flood_reports` - Enable/disable flooding of IGMP reports to this interface when igmp-snooping enabled. Valid values: `enable` `disable` .
* `igmps_flood_traffic` - Enable/disable flooding of IGMP snooping traffic to this interface. Valid values: `enable` `disable` .
* `ip_source_guard` - Enable/disable IP source guard. Valid values: `disable` `enable` .
* `isl_local_trunk_name` - ISL local trunk name.
* `isl_peer_device_name` - ISL peer device name.
* `isl_peer_port_name` - ISL peer port name.
* `lacp_speed` - end Link Aggregation Control Protocol (LACP) messages every 30 seconds (slow) or every second (fast). Valid values: `slow` `fast` .
* `learning_limit` - Limit the number of dynamic MAC addresses on this Port (1 - 128, 0 = no limit, default).
* `lldp_profile` - LLDP port TLV profile. This attribute must reference one of the following datasources: `switch-controller.lldp-profile.name` .
* `lldp_status` - LLDP transmit and receive status. Valid values: `disable` `rx-only` `tx-only` `tx-rx` .
* `loop_guard` - Enable/disable loop-guard on this interface, an STP optimization used to prevent network loops. Valid values: `enabled` `disabled` .
* `loop_guard_timeout` - Loop-guard timeout (0 - 120 min, default = 45).
* `mac_addr` - Port/Trunk MAC.
* `matched_dpp_intf_tags` - Matched interface tags in the dynamic port policy.
* `matched_dpp_policy` - Matched child policy in the dynamic port policy.
* `max_bundle` - Maximum size of LAG bundle (1 - 24, default = 24)
* `mclag` - Enable/disable multi-chassis link aggregation (MCLAG). Valid values: `enable` `disable` .
* `mclag_icl_port` - MCLAG-ICL port.
* `media_type` - Media type.
* `member_withdrawal_behavior` - Port behavior after it withdraws because of loss of control packets. Valid values: `forward` `block` .
* `min_bundle` - Minimum size of LAG bundle (1 - 24, default = 1)
* `mode` - LACP mode: ignore and do not send control messages, or negotiate 802.3ad aggregation passively or actively. Valid values: `static` `lacp-passive` `lacp-active` .
* `p2p_port` - General peer to peer tunnel port.
* `packet_sample_rate` - Packet sampling rate (0 - 99999 p/sec).
* `packet_sampler` - Enable/disable packet sampling on this interface. Valid values: `enabled` `disabled` .
* `pause_meter` - Configure ingress pause metering rate, in kbps (default = 0, disabled).
* `pause_meter_resume` - Resume threshold for resuming traffic on ingress port. Valid values: `75%` `50%` `25%` .
* `poe_capable` - PoE capable.
* `poe_max_power` - PoE maximum power.
* `poe_pre_standard_detection` - Enable/disable PoE pre-standard detection. Valid values: `enable` `disable` .
* `poe_standard` - PoE standard supported.
* `poe_status` - Enable/disable PoE status. Valid values: `enable` `disable` .
* `port_name` - Switch port name.
* `port_number` - Port number.
* `port_owner` - Switch port name.
* `port_policy` - Switch controller dynamic port policy from available options. This attribute must reference one of the following datasources: `switch-controller.dynamic-port-policy.name` .
* `port_prefix_type` - Port prefix type.
* `port_security_policy` - Switch controller authentication policy to apply to this managed switch from available options. This attribute must reference one of the following datasources: `switch-controller.security-policy.802-1X.name` .
* `port_selection_criteria` - Algorithm for aggregate port selection. Valid values: `src-mac` `dst-mac` `src-dst-mac` `src-ip` `dst-ip` `src-dst-ip` .
* `ptp_policy` - PTP policy configuration. This attribute must reference one of the following datasources: `switch-controller.ptp.policy.name` .
* `qos_policy` - Switch controller QoS policy from available options. This attribute must reference one of the following datasources: `switch-controller.qos.qos-policy.name` .
* `rpvst_port` - Enable/disable inter-operability with rapid PVST on this interface. Valid values: `disabled` `enabled` .
* `sample_direction` - Packet sampling direction. Valid values: `tx` `rx` `both` .
* `sflow_counter_interval` - sFlow sampling counter polling interval (0 - 255 sec).
* `speed` - Switch port speed; default and available settings depend on hardware. Valid values: `10half` `10full` `100half` `100full` `1000auto` `1000fiber` `1000full` `10000` `40000` `auto` `auto-module` `100FX-half` `100FX-full` `100000full` `2500auto` `25000full` `50000full` `10000cr` `10000sr` `100000sr4` `100000cr4` `25000cr4` `25000sr4` `5000full` .
* `stacking_port` - Stacking port.
* `status` - Switch port admin status: up or down. Valid values: `up` `down` .
* `sticky_mac` - Enable or disable sticky-mac on the interface. Valid values: `enable` `disable` .
* `storm_control_policy` - Switch controller storm control policy from available options. This attribute must reference one of the following datasources: `switch-controller.storm-control-policy.name` .
* `stp_bpdu_guard` - Enable/disable STP BPDU guard on this interface. Valid values: `enabled` `disabled` .
* `stp_bpdu_guard_timeout` - BPDU Guard disabling protection (0 - 120 min).
* `stp_root_guard` - Enable/disable STP root guard on this interface. Valid values: `enabled` `disabled` .
* `stp_state` - Enable/disable Spanning Tree Protocol (STP) on this interface. Valid values: `enabled` `disabled` .
* `switch_id` - Switch id.
* `type` - Interface type: physical or trunk port. Valid values: `physical` `trunk` .
* `vlan` - Assign switch ports to a VLAN. This attribute must reference one of the following datasources: `system.interface.name` .
* `allowed_vlans` - Configure switch port tagged vlans The structure of `allowed_vlans` block is documented below.

The `allowed_vlans` block contains:

* `vlan_name` - VLAN name. This attribute must reference one of the following datasources: `system.interface.name` .
* `export_tags` - Configure export tag(s) for FortiSwitch port when exported to a virtual port pool. The structure of `export_tags` block is documented below.

The `export_tags` block contains:

* `tag_name` - FortiSwitch port tag name when exported to a virtual port pool. This attribute must reference one of the following datasources: `switch-controller.switch-interface-tag.name` .
* `interface_tags` - Tag(s) associated with the interface for various features including virtual port pool, dynamic port policy. The structure of `interface_tags` block is documented below.

The `interface_tags` block contains:

* `tag_name` - FortiSwitch port tag name when exported to a virtual port pool or matched to dynamic port policy. This attribute must reference one of the following datasources: `switch-controller.switch-interface-tag.name` .
* `members` - Aggregated LAG bundle interfaces. The structure of `members` block is documented below.

The `members` block contains:

* `member_name` - Interface name from available options.
* `untagged_vlans` - Configure switch port untagged vlans The structure of `untagged_vlans` block is documented below.

The `untagged_vlans` block contains:

* `vlan_name` - VLAN name. This attribute must reference one of the following datasources: `system.interface.name` .
* `remote_log` - Configure logging by FortiSwitch device to a remote syslog server. The structure of `remote_log` block is documented below.

The `remote_log` block contains:

* `csv` - Enable/disable comma-separated value (CSV) strings. Valid values: `enable` `disable` .
* `facility` - Facility to log to remote syslog server. Valid values: `kernel` `user` `mail` `daemon` `auth` `syslog` `lpr` `news` `uucp` `cron` `authpriv` `ftp` `ntp` `audit` `alert` `clock` `local0` `local1` `local2` `local3` `local4` `local5` `local6` `local7` .
* `name` - Remote log name.
* `port` - Remote syslog server listening port.
* `server` - IPv4 address of the remote syslog server.
* `severity` - Severity of logs to be transferred to remote log server. Valid values: `emergency` `alert` `critical` `error` `warning` `notification` `information` `debug` .
* `status` - Enable/disable logging by FortiSwitch device to a remote syslog server. Valid values: `enable` `disable` .
* `snmp_community` - Configuration method to edit Simple Network Management Protocol (SNMP) communities. The structure of `snmp_community` block is documented below.

The `snmp_community` block contains:

* `events` - SNMP notifications (traps) to send. Valid values: `cpu-high` `mem-low` `log-full` `intf-ip` `ent-conf-change` .
* `id` - SNMP community ID.
* `name` - SNMP community name.
* `query_v1_port` - SNMP v1 query port (default = 161).
* `query_v1_status` - Enable/disable SNMP v1 queries. Valid values: `disable` `enable` .
* `query_v2c_port` - SNMP v2c query port (default = 161).
* `query_v2c_status` - Enable/disable SNMP v2c queries. Valid values: `disable` `enable` .
* `status` - Enable/disable this SNMP community. Valid values: `disable` `enable` .
* `trap_v1_lport` - SNMP v2c trap local port (default = 162).
* `trap_v1_rport` - SNMP v2c trap remote port (default = 162).
* `trap_v1_status` - Enable/disable SNMP v1 traps. Valid values: `disable` `enable` .
* `trap_v2c_lport` - SNMP v2c trap local port (default = 162).
* `trap_v2c_rport` - SNMP v2c trap remote port (default = 162).
* `trap_v2c_status` - Enable/disable SNMP v2c traps. Valid values: `disable` `enable` .
* `hosts` - Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.

The `hosts` block contains:

* `id` - Host entry ID.
* `ip` - IPv4 address of the SNMP manager (host).
* `snmp_sysinfo` - Configuration method to edit Simple Network Management Protocol (SNMP) system info. The structure of `snmp_sysinfo` block is documented below.

The `snmp_sysinfo` block contains:

* `contact_info` - Contact information.
* `description` - System description.
* `engine_id` - Local SNMP engine ID string (max 24 char).
* `location` - System location.
* `status` - Enable/disable SNMP. Valid values: `disable` `enable` .
* `snmp_trap_threshold` - Configuration method to edit Simple Network Management Protocol (SNMP) trap threshold values. The structure of `snmp_trap_threshold` block is documented below.

The `snmp_trap_threshold` block contains:

* `trap_high_cpu_threshold` - CPU usage when trap is sent.
* `trap_log_full_threshold` - Log disk usage when trap is sent.
* `trap_low_memory_threshold` - Memory usage when trap is sent.
* `snmp_user` - Configuration method to edit Simple Network Management Protocol (SNMP) users. The structure of `snmp_user` block is documented below.

The `snmp_user` block contains:

* `auth_proto` - Authentication protocol. Valid values: `md5` `sha1` `sha224` `sha256` `sha384` `sha512` .
* `auth_pwd` - Password for authentication protocol.
* `name` - SNMP user name.
* `priv_proto` - Privacy (encryption) protocol. Valid values: `aes128` `aes192` `aes192c` `aes256` `aes256c` `des` .
* `priv_pwd` - Password for privacy (encryption) protocol.
* `queries` - Enable/disable SNMP queries for this user. Valid values: `disable` `enable` .
* `query_port` - SNMPv3 query port (default = 161).
* `security_level` - Security level for message authentication and encryption. Valid values: `no-auth-no-priv` `auth-no-priv` `auth-priv` .
* `static_mac` - Configuration method to edit FortiSwitch Static and Sticky MAC. The structure of `static_mac` block is documented below.

The `static_mac` block contains:

* `description` - Description.
* `id` - Id
* `interface` - Interface name.
* `mac` - MAC address.
* `type` - Type. Valid values: `static` `sticky` .
* `vlan` - Vlan. This attribute must reference one of the following datasources: `system.interface.name` .
* `storm_control` - Configuration method to edit FortiSwitch storm control for measuring traffic activity using data rates to prevent traffic disruption. The structure of `storm_control` block is documented below.

The `storm_control` block contains:

* `broadcast` - Enable/disable storm control to drop broadcast traffic. Valid values: `enable` `disable` .
* `local_override` - Enable to override global FortiSwitch storm control settings for this FortiSwitch. Valid values: `enable` `disable` .
* `rate` - Rate in packets per second at which storm traffic is controlled (1 - 10000000, default = 500). Storm control drops excess traffic data rates beyond this threshold.
* `unknown_multicast` - Enable/disable storm control to drop unknown multicast traffic. Valid values: `enable` `disable` .
* `unknown_unicast` - Enable/disable storm control to drop unknown unicast traffic. Valid values: `enable` `disable` .
* `stp_instance` - Configuration method to edit Spanning Tree Protocol (STP) instances. The structure of `stp_instance` block is documented below.

The `stp_instance` block contains:

* `id` - Instance ID.
* `priority` - Priority. Valid values: `0` `4096` `8192` `12288` `16384` `20480` `24576` `28672` `32768` `36864` `40960` `45056` `49152` `53248` `57344` `61440` .
* `stp_settings` - Configuration method to edit Spanning Tree Protocol (STP) settings used to prevent bridge loops. The structure of `stp_settings` block is documented below.

The `stp_settings` block contains:

* `forward_time` - Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
* `hello_time` - Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
* `local_override` - Enable to configure local STP settings that override global STP settings. Valid values: `enable` `disable` .
* `max_age` - Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
* `max_hops` - Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
* `name` - Name of local STP settings configuration.
* `pending_timer` - Pending time (1 - 15 sec, default = 4).
* `revision` - STP revision number (0 - 65535).
* `switch_log` - Configuration method to edit FortiSwitch logging settings (logs are transferred to and inserted into the FortiGate event log). The structure of `switch_log` block is documented below.

The `switch_log` block contains:

* `local_override` - Enable to configure local logging settings that override global logging settings. Valid values: `enable` `disable` .
* `severity` - Severity of FortiSwitch logs that are added to the FortiGate event log. Valid values: `emergency` `alert` `critical` `error` `warning` `notification` `information` `debug` .
* `status` - Enable/disable adding FortiSwitch logs to the FortiGate event log. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_switchcontroller_managedswitch can be imported using:
```sh
terraform import fortios_switchcontroller_managedswitch.labelname {{mkey}}
```
