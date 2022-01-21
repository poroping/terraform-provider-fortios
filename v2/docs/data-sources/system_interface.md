---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_interface"
description: |-
  Get information on a fortios Configure interfaces.
---

# Data Source: fortios_system_interface
Use this data source to get information on a fortios Configure interfaces.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `ac_name` - PPPoE server name.
* `aggregate` - Aggregate interface.
* `algorithm` - Frame distribution algorithm.
* `alias` - Alias will be displayed with the interface name to make it easier to distinguish.
* `allowaccess` - Permitted types of management access to this interface.
* `ap_discover` - Enable/disable automatic registration of unknown FortiAP devices.
* `arpforward` - Enable/disable ARP forwarding.
* `auth_type` - PPP authentication type to use.
* `auto_auth_extension_device` - Enable/disable automatic authorization of dedicated Fortinet extension device on this interface.
* `bandwidth_measure_time` - Bandwidth measure time 
* `bfd` - Bidirectional Forwarding Detection (BFD) settings.
* `bfd_desired_min_tx` - BFD desired minimal transmit interval.
* `bfd_detect_mult` - BFD detection multiplier.
* `bfd_required_min_rx` - BFD required minimal receive interval.
* `broadcast_forticlient_discovery` - Enable/disable broadcasting FortiClient discovery messages.
* `broadcast_forward` - Enable/disable broadcast forwarding.
* `cli_conn_status` - CLI connection status.
* `color` - Color of icon on the GUI.
* `dedicated_to` - Configure interface for single purpose.
* `defaultgw` - Enable to get the gateway IP from the DHCP or PPPoE server.
* `description` - Description.
* `detected_peer_mtu` - MTU of detected peer (0 - 4294967295).
* `detectprotocol` - Protocols used to detect the server.
* `detectserver` - Gateway's ping server for this IP.
* `device_identification` - Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.
* `device_user_identification` - Enable/disable passive gathering of user identity information about users on this interface.
* `devindex` - Device Index.
* `dhcp_classless_route_addition` - Enable/disable addition of classless static routes retrieved from DHCP server.
* `dhcp_client_identifier` - DHCP client identifier.
* `dhcp_relay_agent_option` - Enable/disable DHCP relay agent option.
* `dhcp_relay_interface` - Specify outgoing interface to reach server.
* `dhcp_relay_interface_select_method` - Specify how to select outgoing interface to reach server.
* `dhcp_relay_ip` - DHCP relay IP address.
* `dhcp_relay_request_all_server` - Enable/disable sending of DHCP requests to all servers.
* `dhcp_relay_service` - Enable/disable allowing this interface to act as a DHCP relay.
* `dhcp_relay_type` - DHCP relay type (regular or IPsec).
* `dhcp_renew_time` - DHCP renew time in seconds (300-604800), 0 means use the renew time provided by the server.
* `disc_retry_timeout` - Time in seconds to wait before retrying to start a PPPoE discovery, 0 means no timeout.
* `disconnect_threshold` - Time in milliseconds to wait before sending a notification that this interface is down or disconnected.
* `distance` - Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
* `dns_server_override` - Enable/disable use DNS acquired by DHCP or PPPoE.
* `drop_fragment` - Enable/disable drop fragment packets.
* `drop_overlapped_fragment` - Enable/disable drop overlapped fragment packets.
* `egress_cos` - Override outgoing CoS in user VLAN tag.
* `egress_shaping_profile` - Outgoing traffic shaping profile.
* `eip` - External IP.
* `estimated_downstream_bandwidth` - Estimated maximum downstream bandwidth (kbps). Used to estimate link utilization.
* `estimated_upstream_bandwidth` - Estimated maximum upstream bandwidth (kbps). Used to estimate link utilization.
* `explicit_ftp_proxy` - Enable/disable the explicit FTP proxy on this interface.
* `explicit_web_proxy` - Enable/disable the explicit web proxy on this interface.
* `external` - Enable/disable identifying the interface as an external interface (which usually means it's connected to the Internet).
* `fail_action_on_extender` - Action on extender when interface fail .
* `fail_alert_method` - Select link-failed-signal or link-down method to alert about a failed link.
* `fail_detect` - Enable/disable fail detection features for this interface.
* `fail_detect_option` - Options for detecting that this interface has failed.
* `fortilink` - Enable FortiLink to dedicate this interface to manage other Fortinet devices.
* `fortilink_backup_link` - fortilink split interface backup link.
* `fortilink_neighbor_detect` - Protocol for FortiGate neighbor discovery.
* `fortilink_split_interface` - Enable/disable FortiLink split interface to connect member link to different FortiSwitch in stack for uplink redundancy.
* `fortilink_stacking` - Enable/disable FortiLink switch-stacking on this interface.
* `forward_domain` - Transparent mode forward domain.
* `gwdetect` - Enable/disable detect gateway alive for first.
* `ha_priority` - HA election priority for the PING server.
* `icmp_accept_redirect` - Enable/disable ICMP accept redirect.
* `icmp_send_redirect` - Enable/disable sending of ICMP redirects.
* `ident_accept` - Enable/disable authentication for this interface.
* `idle_timeout` - PPPoE auto disconnect after idle timeout seconds, 0 means no timeout.
* `inbandwidth` - Bandwidth limit for incoming traffic (0 - 16776000 kbps), 0 means unlimited.
* `ingress_cos` - Override incoming CoS in user VLAN tag on VLAN interface or assign a priority VLAN tag on physical interface.
* `ingress_shaping_profile` - Incoming traffic shaping profile.
* `ingress_spillover_threshold` - Ingress Spillover threshold (0 - 16776000 kbps), 0 means unlimited.
* `interface` - Interface name.
* `internal` - Implicitly created.
* `ip` - Interface IPv4 address and subnet mask, syntax: X.X.X.X/24.
* `ip_managed_by_fortiipam` - Enable/disable automatic IP address assignment of this interface by FortiIPAM.
* `ipmac` - Enable/disable IP/MAC binding.
* `ips_sniffer_mode` - Enable/disable the use of this interface as a one-armed sniffer.
* `ipunnumbered` - Unnumbered IP used for PPPoE interfaces for which no unique local address is provided.
* `l2forward` - Enable/disable l2 forwarding.
* `lacp_ha_slave` - LACP HA slave.
* `lacp_mode` - LACP mode.
* `lacp_speed` - How often the interface sends LACP messages.
* `lcp_echo_interval` - Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
* `lcp_max_echo_fails` - Maximum missed LCP echo messages before disconnect.
* `link_up_delay` - Number of milliseconds to wait before considering a link is up.
* `lldp_network_policy` - LLDP-MED network policy profile.
* `lldp_reception` - Enable/disable Link Layer Discovery Protocol (LLDP) reception.
* `lldp_transmission` - Enable/disable Link Layer Discovery Protocol (LLDP) transmission.
* `macaddr` - Change the interface's MAC address.
* `managed_subnetwork_size` - Number of IP addresses to be allocated by FortiIPAM and used by this FortiGate unit's DHCP server settings.
* `management_ip` - High Availability in-band management IP address of this interface.
* `measured_downstream_bandwidth` - Measured downstream bandwidth (kbps). 
* `measured_upstream_bandwidth` - Measured upstream bandwidth (kbps). 
* `mediatype` - Select SFP media interface type
* `min_links` - Minimum number of aggregated ports that must be up.
* `min_links_down` - Action to take when less than the configured minimum number of links are active.
* `mode` - Addressing mode (static, DHCP, PPPoE).
* `monitor_bandwidth` - Enable monitoring bandwidth on this interface.
* `mtu` - MTU value for this interface.
* `mtu_override` - Enable to set a custom MTU for this interface.
* `name` - Name.
* `ndiscforward` - Enable/disable NDISC forwarding.
* `netbios_forward` - Enable/disable NETBIOS forwarding.
* `netflow_sampler` - Enable/disable NetFlow on this interface and set the data that NetFlow collects (rx, tx, or both).
* `outbandwidth` - Bandwidth limit for outgoing traffic (0 - 16776000 kbps), 0 means unlimited.
* `padt_retry_timeout` - PPPoE Active Discovery Terminate (PADT) used to terminate sessions after an idle time.
* `password` - PPPoE account's password.
* `ping_serv_status` - PING server status.
* `polling_interval` - sFlow polling interval (1 - 255 sec).
* `pppoe_unnumbered_negotiate` - Enable/disable PPPoE unnumbered negotiation.
* `pptp_auth_type` - PPTP authentication type.
* `pptp_client` - Enable/disable PPTP client.
* `pptp_password` - PPTP password.
* `pptp_server_ip` - PPTP server IP address.
* `pptp_timeout` - Idle timer in minutes (0 for disabled).
* `pptp_user` - PPTP user name.
* `preserve_session_route` - Enable/disable preservation of session route when dirty.
* `priority` - Priority of learned routes.
* `priority_override` - Enable/disable fail back to higher priority port once recovered.
* `proxy_captive_portal` - Enable/disable proxy captive portal on this interface.
* `redundant_interface` - Redundant interface.
* `remote_ip` - Remote IP address of tunnel.
* `replacemsg_override_group` - Replacement message override group.
* `ring_rx` - RX ring size.
* `ring_tx` - TX ring size.
* `role` - Interface role.
* `sample_direction` - Data that NetFlow collects (rx, tx, or both).
* `sample_rate` - sFlow sample rate (10 - 99999).
* `secondary_ip` - Enable/disable adding a secondary IP to this interface.
* `security_exempt_list` - Name of security-exempt-list.
* `security_external_logout` - URL of external authentication logout server.
* `security_external_web` - URL of external authentication web server.
* `security_mac_auth_bypass` - Enable/disable MAC authentication bypass.
* `security_mode` - Turn on captive portal authentication for this interface.
* `security_redirect_url` - URL redirection after disclaimer/authentication.
* `service_name` - PPPoE service name.
* `sflow_sampler` - Enable/disable sFlow on this interface.
* `snmp_index` - Permanent SNMP Index of the interface.
* `speed` - Interface speed. The default setting and the options available depend on the interface hardware.
* `spillover_threshold` - Egress Spillover threshold (0 - 16776000 kbps), 0 means unlimited.
* `src_check` - Enable/disable source IP check.
* `status` - Bring the interface up or shut the interface down.
* `stpforward` - Enable/disable STP forwarding.
* `stpforward_mode` - Configure STP forwarding mode.
* `subst` - Enable to always send packets from this interface to a destination MAC address.
* `substitute_dst_mac` - Destination MAC address that all packets are sent to from this interface.
* `swc_first_create` - Initial create for switch-controller VLANs.
* `swc_vlan` - Creation status for switch-controller VLANs.
* `switch` - Contained in switch.
* `switch_controller_access_vlan` - Block FortiSwitch port-to-port traffic.
* `switch_controller_arp_inspection` - Enable/disable FortiSwitch ARP inspection.
* `switch_controller_dhcp_snooping` - Switch controller DHCP snooping.
* `switch_controller_dhcp_snooping_option82` - Switch controller DHCP snooping option82.
* `switch_controller_dhcp_snooping_verify_mac` - Switch controller DHCP snooping verify MAC.
* `switch_controller_dynamic` - Integrated FortiLink settings for managed FortiSwitch.
* `switch_controller_feature` - Interface's purpose when assigning traffic (read only).
* `switch_controller_igmp_snooping` - Switch controller IGMP snooping.
* `switch_controller_igmp_snooping_fast_leave` - Switch controller IGMP snooping fast-leave.
* `switch_controller_igmp_snooping_proxy` - Switch controller IGMP snooping proxy.
* `switch_controller_iot_scanning` - Enable/disable managed FortiSwitch IoT scanning.
* `switch_controller_learning_limit` - Limit the number of dynamic MAC addresses on this VLAN (1 - 128, 0 = no limit, default).
* `switch_controller_mgmt_vlan` - VLAN to use for FortiLink management purposes.
* `switch_controller_nac` - Integrated FortiLink settings for managed FortiSwitch.
* `switch_controller_rspan_mode` - Stop Layer2 MAC learning and interception of BPDUs and other packets on this interface.
* `switch_controller_source_ip` - Source IP address used in FortiLink over L3 connections.
* `switch_controller_traffic_policy` - Switch controller traffic policy for the VLAN.
* `system_id` - Define a system ID for the aggregate interface.
* `system_id_type` - Method in which system ID is generated.
* `tcp_mss` - TCP maximum segment size. 0 means do not change segment size.
* `trust_ip_1` - Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
* `trust_ip_2` - Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
* `trust_ip_3` - Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
* `trust_ip6_1` - Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
* `trust_ip6_2` - Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
* `trust_ip6_3` - Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
* `type` - Interface type.
* `username` - Username of the PPPoE account, provided by your ISP.
* `vdom` - Interface is in this virtual domain (VDOM).
* `vindex` - Switch control interface VLAN ID.
* `vlan_protocol` - Ethernet protocol of VLAN.
* `vlanforward` - Enable/disable traffic forwarding between VLANs on this interface.
* `vlanid` - VLAN ID (1 - 4094).
* `vrf` - Virtual Routing Forwarding ID.
* `vrrp_virtual_mac` - Enable/disable use of virtual MAC for VRRP.
* `wccp` - Enable/disable WCCP on this interface. Used for encapsulated WCCP communication between WCCP clients and servers.
* `weight` - Default weight for static routes (if route has no weight configured).
* `wins_ip` - WINS server IP.
* `client_options` - DHCP client options.The structure of `client_options` block is documented below.

The `client_options` block contains:

* `code` - DHCP client option code.
* `id` - ID.
* `ip` - DHCP option IPs.
* `type` - DHCP client option type.
* `value` - DHCP client option value.
* `dhcp_snooping_server_list` - Configure DHCP server access list.The structure of `dhcp_snooping_server_list` block is documented below.

The `dhcp_snooping_server_list` block contains:

* `name` - DHCP server name.
* `server_ip` - IP address for DHCP server.
* `egress_queues` - Configure queues of NP port on egress path.The structure of `egress_queues` block is documented below.

The `egress_queues` block contains:

* `cos0` - CoS profile name for CoS 0.
* `cos1` - CoS profile name for CoS 1.
* `cos2` - CoS profile name for CoS 2.
* `cos3` - CoS profile name for CoS 3.
* `cos4` - CoS profile name for CoS 4.
* `cos5` - CoS profile name for CoS 5.
* `cos6` - CoS profile name for CoS 6.
* `cos7` - CoS profile name for CoS 7.
* `fail_alert_interfaces` - Names of the FortiGate interfaces to which the link failure alert is sent.The structure of `fail_alert_interfaces` block is documented below.

The `fail_alert_interfaces` block contains:

* `name` - Names of the non-virtual interface.
* `ipv6` - IPv6 of interface.The structure of `ipv6` block is documented below.

The `ipv6` block contains:

* `autoconf` - Enable/disable address auto config.
* `cli_conn6_status` - CLI IPv6 connection status.
* `dhcp6_client_options` - DHCPv6 client options.
* `dhcp6_information_request` - Enable/disable DHCPv6 information request.
* `dhcp6_prefix_delegation` - Enable/disable DHCPv6 prefix delegation.
* `dhcp6_prefix_hint` - DHCPv6 prefix that will be used as a hint to the upstream DHCPv6 server.
* `dhcp6_prefix_hint_plt` - DHCPv6 prefix hint preferred life time (sec), 0 means unlimited lease time.
* `dhcp6_prefix_hint_vlt` - DHCPv6 prefix hint valid life time (sec).
* `dhcp6_relay_ip` - DHCPv6 relay IP address.
* `dhcp6_relay_service` - Enable/disable DHCPv6 relay.
* `dhcp6_relay_type` - DHCPv6 relay type.
* `icmp6_send_redirect` - Enable/disable sending of ICMPv6 redirects.
* `interface_identifier` - IPv6 interface identifier.
* `ip6_address` - Primary IPv6 address prefix, syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx
* `ip6_allowaccess` - Allow management access to the interface.
* `ip6_default_life` - Default life (sec).
* `ip6_delegated_prefix_iaid` - IAID of obtained delegated-prefix from the upstream interface.
* `ip6_dns_server_override` - Enable/disable using the DNS server acquired by DHCP.
* `ip6_hop_limit` - Hop limit (0 means unspecified).
* `ip6_link_mtu` - IPv6 link MTU.
* `ip6_manage_flag` - Enable/disable the managed flag.
* `ip6_max_interval` - IPv6 maximum interval (4 to 1800 sec).
* `ip6_min_interval` - IPv6 minimum interval (3 to 1350 sec).
* `ip6_mode` - Addressing mode (static, DHCP, delegated).
* `ip6_other_flag` - Enable/disable the other IPv6 flag.
* `ip6_prefix_mode` - Assigning a prefix from DHCP or RA.
* `ip6_reachable_time` - IPv6 reachable time (milliseconds; 0 means unspecified).
* `ip6_retrans_time` - IPv6 retransmit time (milliseconds; 0 means unspecified).
* `ip6_send_adv` - Enable/disable sending advertisements about the interface.
* `ip6_subnet` -  Subnet to routing prefix, syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx
* `ip6_upstream_interface` - Interface name providing delegated information.
* `nd_cert` - Neighbor discovery certificate.
* `nd_cga_modifier` - Neighbor discovery CGA modifier.
* `nd_mode` - Neighbor discovery mode.
* `nd_security_level` - Neighbor discovery security level (0 - 7; 0 = least secure, default = 0).
* `nd_timestamp_delta` - Neighbor discovery timestamp delta value (1 - 3600 sec; default = 300).
* `nd_timestamp_fuzz` - Neighbor discovery timestamp fuzz factor (1 - 60 sec; default = 1).
* `ra_send_mtu` - Enable/disable sending link MTU in RA packet.
* `unique_autoconf_addr` - Enable/disable unique auto config address.
* `vrip6_link_local` - Link-local IPv6 address of virtual router.
* `vrrp_virtual_mac6` - Enable/disable virtual MAC for VRRP.
* `dhcp6_iapd_list` - DHCPv6 IA-PD listThe structure of `dhcp6_iapd_list` block is documented below.

The `dhcp6_iapd_list` block contains:

* `iaid` - Identity association identifier.
* `prefix_hint` - DHCPv6 prefix that will be used as a hint to the upstream DHCPv6 server.
* `prefix_hint_plt` - DHCPv6 prefix hint preferred life time (sec), 0 means unlimited lease time.
* `prefix_hint_vlt` - DHCPv6 prefix hint valid life time (sec).
* `ip6_delegated_prefix_list` - Advertised IPv6 delegated prefix list.The structure of `ip6_delegated_prefix_list` block is documented below.

The `ip6_delegated_prefix_list` block contains:

* `autonomous_flag` - Enable/disable the autonomous flag.
* `delegated_prefix_iaid` - IAID of obtained delegated-prefix from the upstream interface.
* `onlink_flag` - Enable/disable the onlink flag.
* `prefix_id` - Prefix ID.
* `rdnss` - Recursive DNS server option.
* `rdnss_service` - Recursive DNS service option.
* `subnet` -  Add subnet ID to routing prefix.
* `upstream_interface` - Name of the interface that provides delegated information.
* `ip6_extra_addr` - Extra IPv6 address prefixes of interface.The structure of `ip6_extra_addr` block is documented below.

The `ip6_extra_addr` block contains:

* `prefix` - IPv6 address prefix.
* `ip6_prefix_list` - Advertised prefix list.The structure of `ip6_prefix_list` block is documented below.

The `ip6_prefix_list` block contains:

* `autonomous_flag` - Enable/disable the autonomous flag.
* `onlink_flag` - Enable/disable the onlink flag.
* `preferred_life_time` - Preferred life time (sec).
* `prefix` - IPv6 prefix.
* `rdnss` - Recursive DNS server option.
* `valid_life_time` - Valid life time (sec).
* `dnssl` - DNS search list option.The structure of `dnssl` block is documented below.

The `dnssl` block contains:

* `domain` - Domain name.
* `vrrp6` - IPv6 VRRP configuration.The structure of `vrrp6` block is documented below.

The `vrrp6` block contains:

* `accept_mode` - Enable/disable accept mode.
* `adv_interval` - Advertisement interval (1 - 255 seconds).
* `preempt` - Enable/disable preempt mode.
* `priority` - Priority of the virtual router (1 - 255).
* `start_time` - Startup time (1 - 255 seconds).
* `status` - Enable/disable VRRP.
* `vrdst6` - Monitor the route to this destination.
* `vrgrp` - VRRP group ID (1 - 65535).
* `vrid` - Virtual router identifier (1 - 255).
* `vrip6` - IPv6 address of the virtual router.
* `member` - Physical interfaces that belong to the aggregate or redundant interface.The structure of `member` block is documented below.

The `member` block contains:

* `interface_name` - Physical interface name.
* `secondaryip` - Second IP address of interface.The structure of `secondaryip` block is documented below.

The `secondaryip` block contains:

* `allowaccess` - Management access settings for the secondary IP address.
* `detectprotocol` - Protocols used to detect the server.
* `detectserver` - Gateway's ping server for this IP.
* `gwdetect` - Enable/disable detect gateway alive for first.
* `ha_priority` - HA election priority for the PING server.
* `id` - ID.
* `ip` - Secondary IP address of the interface.
* `ping_serv_status` - PING server status.
* `security_groups` - User groups that can authenticate with the captive portal.The structure of `security_groups` block is documented below.

The `security_groups` block contains:

* `name` - Names of user groups that can authenticate with the captive portal.
* `tagging` - Config object tagging.The structure of `tagging` block is documented below.

The `tagging` block contains:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.
* `vrrp` - VRRP configuration.The structure of `vrrp` block is documented below.

The `vrrp` block contains:

* `accept_mode` - Enable/disable accept mode.
* `adv_interval` - Advertisement interval (1 - 255 seconds).
* `ignore_default_route` - Enable/disable ignoring of default route when checking destination.
* `preempt` - Enable/disable preempt mode.
* `priority` - Priority of the virtual router (1 - 255).
* `start_time` - Startup time (1 - 255 seconds).
* `status` - Enable/disable this VRRP configuration.
* `version` - VRRP version.
* `vrdst` - Monitor the route to this destination.
* `vrdst_priority` - Priority of the virtual router when the virtual router destination becomes unreachable (0 - 254).
* `vrgrp` - VRRP group ID (1 - 65535).
* `vrid` - Virtual router identifier (1 - 255).
* `vrip` - IP address of the virtual router.
* `proxy_arp` - VRRP Proxy ARP configuration.The structure of `proxy_arp` block is documented below.

The `proxy_arp` block contains:

* `id` - ID.
* `ip` - Set IP addresses of proxy ARP.
