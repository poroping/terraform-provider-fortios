---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_interface"
description: |-
  Configure interfaces.
---

## fortios_system_interface
Configure interfaces.

## Example Usage

```hcl
resource "fortios_system_interface" "example" {
  allow_append = true
  vdomparam    = "root"

  name = "TESTACCINT"
  type = "loopback"
  ip   = "169.254.72.255/32"
  vdom = "root"
}

resource "fortios_system_interface" "example2" {
  allow_append = true
  vdomparam    = "root"

  name = "TESTV6"
  type = "loopback"
  ip   = "169.254.71.255/32"
  vdom = "root"
  secondary_ip = "enable"

  ipv6 {
    ip6_address = "2001:beef::01/128"
  }

  secondaryip {
    ip = "3.3.3.3/32"
  }
}
```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `ac_name` - PPPoE server name.
* `aggregate` - Aggregate interface.
* `algorithm` - Frame distribution algorithm. Valid values: `L2` `L3` `L4` .
* `alias` - Alias will be displayed with the interface name to make it easier to distinguish.
* `allowaccess` - Permitted types of management access to this interface. Valid values: `ping` `https` `ssh` `snmp` `http` `telnet` `fgfm` `radius-acct` `probe-response` `fabric` `ftm` `speed-test` .
* `ap_discover` - Enable/disable automatic registration of unknown FortiAP devices. Valid values: `enable` `disable` .
* `arpforward` - Enable/disable ARP forwarding. Valid values: `enable` `disable` .
* `auth_cert` - HTTPS server certificate. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `auth_portal_addr` - Address of captive portal.
* `auth_type` - PPP authentication type to use. Valid values: `auto` `pap` `chap` `mschapv1` `mschapv2` .
* `auto_auth_extension_device` - Enable/disable automatic authorization of dedicated Fortinet extension device on this interface. Valid values: `enable` `disable` .
* `bandwidth_measure_time` - Bandwidth measure time.
* `bfd` - Bidirectional Forwarding Detection (BFD) settings. Valid values: `global` `enable` `disable` .
* `bfd_desired_min_tx` - BFD desired minimal transmit interval.
* `bfd_detect_mult` - BFD detection multiplier.
* `bfd_required_min_rx` - BFD required minimal receive interval.
* `broadcast_forticlient_discovery` - Enable/disable broadcasting FortiClient discovery messages. Valid values: `enable` `disable` .
* `broadcast_forward` - Enable/disable broadcast forwarding. Valid values: `enable` `disable` .
* `cli_conn_status` - CLI connection status.
* `color` - Color of icon on the GUI.
* `dedicated_to` - Configure interface for single purpose. Valid values: `none` `management` .
* `defaultgw` - Enable to get the gateway IP from the DHCP or PPPoE server. Valid values: `enable` `disable` .
* `description` - Description.
* `detected_peer_mtu` - MTU of detected peer (0 - 4294967295).
* `detectprotocol` - Protocols used to detect the server. Valid values: `ping` `tcp-echo` `udp-echo` .
* `detectserver` - Gateway's ping server for this IP.
* `device_identification` - Enable/disable passively gathering of device identity information about the devices on the network connected to this interface. Valid values: `enable` `disable` .
* `device_user_identification` - Enable/disable passive gathering of user identity information about users on this interface. Valid values: `enable` `disable` .
* `devindex` - Device Index.
* `dhcp_classless_route_addition` - Enable/disable addition of classless static routes retrieved from DHCP server. Valid values: `enable` `disable` .
* `dhcp_client_identifier` - DHCP client identifier.
* `dhcp_relay_agent_option` - Enable/disable DHCP relay agent option. Valid values: `enable` `disable` .
* `dhcp_relay_interface` - Specify outgoing interface to reach server. This attribute must reference one of the following datasources: `system.interface.name` .
* `dhcp_relay_interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `dhcp_relay_ip` - DHCP relay IP address.
* `dhcp_relay_link_selection` - DHCP relay link selection.
* `dhcp_relay_request_all_server` - Enable/disable sending of DHCP requests to all servers. Valid values: `disable` `enable` .
* `dhcp_relay_service` - Enable/disable allowing this interface to act as a DHCP relay. Valid values: `disable` `enable` .
* `dhcp_relay_type` - DHCP relay type (regular or IPsec). Valid values: `regular` `ipsec` .
* `dhcp_renew_time` - DHCP renew time in seconds (300-604800), 0 means use the renew time provided by the server.
* `disc_retry_timeout` - Time in seconds to wait before retrying to start a PPPoE discovery, 0 means no timeout.
* `disconnect_threshold` - Time in milliseconds to wait before sending a notification that this interface is down or disconnected.
* `distance` - Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
* `dns_server_override` - Enable/disable use DNS acquired by DHCP or PPPoE. Valid values: `enable` `disable` .
* `dns_server_protocol` - DNS transport protocols. Valid values: `cleartext` `dot` `doh` .
* `drop_fragment` - Enable/disable drop fragment packets. Valid values: `enable` `disable` .
* `drop_overlapped_fragment` - Enable/disable drop overlapped fragment packets. Valid values: `enable` `disable` .
* `eap_ca_cert` - EAP CA certificate name. This attribute must reference one of the following datasources: `certificate.ca.name` .
* `eap_identity` - EAP identity.
* `eap_method` - EAP method. Valid values: `tls` `peap` .
* `eap_password` - EAP password.
* `eap_supplicant` - Enable/disable EAP-Supplicant. Valid values: `enable` `disable` .
* `eap_user_cert` - EAP user certificate name. This attribute must reference one of the following datasources: `certificate.local.name` .
* `egress_cos` - Override outgoing CoS in user VLAN tag. Valid values: `disable` `cos0` `cos1` `cos2` `cos3` `cos4` `cos5` `cos6` `cos7` .
* `egress_shaping_profile` - Outgoing traffic shaping profile. This attribute must reference one of the following datasources: `firewall.shaping-profile.profile-name` .
* `eip` - External IP.
* `estimated_downstream_bandwidth` - Estimated maximum downstream bandwidth (kbps). Used to estimate link utilization.
* `estimated_upstream_bandwidth` - Estimated maximum upstream bandwidth (kbps). Used to estimate link utilization.
* `explicit_ftp_proxy` - Enable/disable the explicit FTP proxy on this interface. Valid values: `enable` `disable` .
* `explicit_web_proxy` - Enable/disable the explicit web proxy on this interface. Valid values: `enable` `disable` .
* `external` - Enable/disable identifying the interface as an external interface (which usually means it's connected to the Internet). Valid values: `enable` `disable` .
* `fail_action_on_extender` - Action on FortiExtender when interface fail. Valid values: `soft-restart` `hard-restart` `reboot` .
* `fail_alert_method` - Select link-failed-signal or link-down method to alert about a failed link. Valid values: `link-failed-signal` `link-down` .
* `fail_detect` - Enable/disable fail detection features for this interface. Valid values: `enable` `disable` .
* `fail_detect_option` - Options for detecting that this interface has failed. Valid values: `detectserver` `link-down` .
* `fortilink` - Enable FortiLink to dedicate this interface to manage other Fortinet devices. Valid values: `enable` `disable` .
* `fortilink_backup_link` - FortiLink split interface backup link.
* `fortilink_neighbor_detect` - Protocol for FortiGate neighbor discovery. Valid values: `lldp` `fortilink` .
* `fortilink_split_interface` - Enable/disable FortiLink split interface to connect member link to different FortiSwitch in stack for uplink redundancy. Valid values: `enable` `disable` .
* `fortilink_stacking` - Enable/disable FortiLink switch-stacking on this interface. Valid values: `enable` `disable` .
* `forward_domain` - Transparent mode forward domain.
* `gwdetect` - Enable/disable detect gateway alive for first. Valid values: `enable` `disable` .
* `ha_priority` - HA election priority for the PING server.
* `icmp_accept_redirect` - Enable/disable ICMP accept redirect. Valid values: `enable` `disable` .
* `icmp_send_redirect` - Enable/disable sending of ICMP redirects. Valid values: `enable` `disable` .
* `ident_accept` - Enable/disable authentication for this interface. Valid values: `enable` `disable` .
* `idle_timeout` - PPPoE auto disconnect after idle timeout seconds, 0 means no timeout.
* `ike_saml_server` - Configure IKE authentication SAML server. This attribute must reference one of the following datasources: `user.saml.name` .
* `inbandwidth` - Bandwidth limit for incoming traffic (0 - 16776000 kbps), 0 means unlimited.
* `ingress_cos` - Override incoming CoS in user VLAN tag on VLAN interface or assign a priority VLAN tag on physical interface. Valid values: `disable` `cos0` `cos1` `cos2` `cos3` `cos4` `cos5` `cos6` `cos7` .
* `ingress_shaping_profile` - Incoming traffic shaping profile. This attribute must reference one of the following datasources: `firewall.shaping-profile.profile-name` .
* `ingress_spillover_threshold` - Ingress Spillover threshold (0 - 16776000 kbps), 0 means unlimited.
* `interface` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `internal` - Implicitly created.
* `ip` - Interface IPv4 address and subnet mask, syntax: X.X.X.X/24.
* `ip_managed_by_fortiipam` - Enable/disable automatic IP address assignment of this interface by FortiIPAM. Valid values: `enable` `disable` .
* `ipmac` - Enable/disable IP/MAC binding. Valid values: `enable` `disable` .
* `ips_sniffer_mode` - Enable/disable the use of this interface as a one-armed sniffer. Valid values: `enable` `disable` .
* `ipunnumbered` - Unnumbered IP used for PPPoE interfaces for which no unique local address is provided.
* `l2forward` - Enable/disable l2 forwarding. Valid values: `enable` `disable` .
* `lacp_ha_slave` - LACP HA slave. Valid values: `enable` `disable` .
* `lacp_mode` - LACP mode. Valid values: `static` `passive` `active` .
* `lacp_speed` - How often the interface sends LACP messages. Valid values: `slow` `fast` .
* `lcp_echo_interval` - Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
* `lcp_max_echo_fails` - Maximum missed LCP echo messages before disconnect.
* `link_up_delay` - Number of milliseconds to wait before considering a link is up.
* `lldp_network_policy` - LLDP-MED network policy profile. This attribute must reference one of the following datasources: `system.lldp.network-policy.name` .
* `lldp_reception` - Enable/disable Link Layer Discovery Protocol (LLDP) reception. Valid values: `enable` `disable` `vdom` .
* `lldp_transmission` - Enable/disable Link Layer Discovery Protocol (LLDP) transmission. Valid values: `enable` `disable` `vdom` .
* `macaddr` - Change the interface's MAC address.
* `managed_subnetwork_size` - Number of IP addresses to be allocated by FortiIPAM and used by this FortiGate unit's DHCP server settings. Valid values: `32` `64` `128` `256` `512` `1024` `2048` `4096` `8192` `16384` `32768` `65536` .
* `management_ip` - High Availability in-band management IP address of this interface.
* `measured_downstream_bandwidth` - Measured downstream bandwidth (kbps).
* `measured_upstream_bandwidth` - Measured upstream bandwidth (kbps).
* `mediatype` - Select SFP media interface type Valid values: `serdes-sfp` `sgmii-sfp` `serdes-copper-sfp` .
* `min_links` - Minimum number of aggregated ports that must be up.
* `min_links_down` - Action to take when less than the configured minimum number of links are active. Valid values: `operational` `administrative` .
* `mode` - Addressing mode (static, DHCP, PPPoE). Valid values: `static` `dhcp` `pppoe` .
* `monitor_bandwidth` - Enable monitoring bandwidth on this interface. Valid values: `enable` `disable` .
* `mtu` - MTU value for this interface.
* `mtu_override` - Enable to set a custom MTU for this interface. Valid values: `enable` `disable` .
* `name` - Name.
* `ndiscforward` - Enable/disable NDISC forwarding. Valid values: `enable` `disable` .
* `netbios_forward` - Enable/disable NETBIOS forwarding. Valid values: `disable` `enable` .
* `netflow_sampler` - Enable/disable NetFlow on this interface and set the data that NetFlow collects (rx, tx, or both). Valid values: `disable` `tx` `rx` `both` .
* `outbandwidth` - Bandwidth limit for outgoing traffic (0 - 16776000 kbps), 0 means unlimited.
* `padt_retry_timeout` - PPPoE Active Discovery Terminate (PADT) used to terminate sessions after an idle time.
* `password` - PPPoE account's password.
* `ping_serv_status` - PING server status.
* `polling_interval` - sFlow polling interval in seconds (1 - 255).
* `pppoe_unnumbered_negotiate` - Enable/disable PPPoE unnumbered negotiation. Valid values: `enable` `disable` .
* `pptp_auth_type` - PPTP authentication type. Valid values: `auto` `pap` `chap` `mschapv1` `mschapv2` .
* `pptp_client` - Enable/disable PPTP client. Valid values: `enable` `disable` .
* `pptp_password` - PPTP password.
* `pptp_server_ip` - PPTP server IP address.
* `pptp_timeout` - Idle timer in minutes (0 for disabled).
* `pptp_user` - PPTP user name.
* `preserve_session_route` - Enable/disable preservation of session route when dirty. Valid values: `enable` `disable` .
* `priority` - Priority of learned routes.
* `priority_override` - Enable/disable fail back to higher priority port once recovered. Valid values: `enable` `disable` .
* `proxy_captive_portal` - Enable/disable proxy captive portal on this interface. Valid values: `enable` `disable` .
* `reachable_time` - IPv4 reachable time in milliseconds (30000 - 3600000, default = 30000).
* `redundant_interface` - Redundant interface.
* `remote_ip` - Remote IP address of tunnel.
* `replacemsg_override_group` - Replacement message override group.
* `ring_rx` - RX ring size.
* `ring_tx` - TX ring size.
* `role` - Interface role. Valid values: `lan` `wan` `dmz` `undefined` .
* `sample_direction` - Data that NetFlow collects (rx, tx, or both). Valid values: `tx` `rx` `both` .
* `sample_rate` - sFlow sample rate (10 - 99999).
* `secondary_ip` - Enable/disable adding a secondary IP to this interface. Valid values: `enable` `disable` .
* `security_exempt_list` - Name of security-exempt-list.
* `security_external_logout` - URL of external authentication logout server.
* `security_external_web` - URL of external authentication web server.
* `security_mac_auth_bypass` - Enable/disable MAC authentication bypass. Valid values: `mac-auth-only` `enable` `disable` .
* `security_mode` - Turn on captive portal authentication for this interface. Valid values: `none` `captive-portal` `802.1X` .
* `security_redirect_url` - URL redirection after disclaimer/authentication.
* `service_name` - PPPoE service name.
* `sflow_sampler` - Enable/disable sFlow on this interface. Valid values: `enable` `disable` .
* `snmp_index` - Permanent SNMP Index of the interface.
* `speed` - Interface speed. The default setting and the options available depend on the interface hardware. Valid values: `auto` `10full` `10half` `100full` `100half` `1000full` `1000auto` .
* `spillover_threshold` - Egress Spillover threshold (0 - 16776000 kbps), 0 means unlimited.
* `src_check` - Enable/disable source IP check. Valid values: `enable` `disable` .
* `status` - Bring the interface up or shut the interface down. Valid values: `up` `down` .
* `stpforward` - Enable/disable STP forwarding. Valid values: `enable` `disable` .
* `stpforward_mode` - Configure STP forwarding mode. Valid values: `rpl-all-ext-id` `rpl-bridge-ext-id` `rpl-nothing` .
* `subst` - Enable to always send packets from this interface to a destination MAC address. Valid values: `enable` `disable` .
* `substitute_dst_mac` - Destination MAC address that all packets are sent to from this interface.
* `swc_first_create` - Initial create for switch-controller VLANs.
* `swc_vlan` - Creation status for switch-controller VLANs.
* `switch` - Contained in switch.
* `switch_controller_access_vlan` - Block FortiSwitch port-to-port traffic. Valid values: `enable` `disable` .
* `switch_controller_arp_inspection` - Enable/disable FortiSwitch ARP inspection. Valid values: `enable` `disable` .
* `switch_controller_dhcp_snooping` - Switch controller DHCP snooping. Valid values: `enable` `disable` .
* `switch_controller_dhcp_snooping_option82` - Switch controller DHCP snooping option82. Valid values: `enable` `disable` .
* `switch_controller_dhcp_snooping_verify_mac` - Switch controller DHCP snooping verify MAC. Valid values: `enable` `disable` .
* `switch_controller_dynamic` - Integrated FortiLink settings for managed FortiSwitch. This attribute must reference one of the following datasources: `switch-controller.fortilink-settings.name` .
* `switch_controller_feature` - Interface's purpose when assigning traffic (read only). Valid values: `none` `default-vlan` `quarantine` `rspan` `voice` `video` `nac` `nac-segment` .
* `switch_controller_igmp_snooping` - Switch controller IGMP snooping. Valid values: `enable` `disable` .
* `switch_controller_igmp_snooping_fast_leave` - Switch controller IGMP snooping fast-leave. Valid values: `enable` `disable` .
* `switch_controller_igmp_snooping_proxy` - Switch controller IGMP snooping proxy. Valid values: `enable` `disable` .
* `switch_controller_iot_scanning` - Enable/disable managed FortiSwitch IoT scanning. Valid values: `enable` `disable` .
* `switch_controller_learning_limit` - Limit the number of dynamic MAC addresses on this VLAN (1 - 128, 0 = no limit, default).
* `switch_controller_mgmt_vlan` - VLAN to use for FortiLink management purposes.
* `switch_controller_nac` - Integrated FortiLink settings for managed FortiSwitch. This attribute must reference one of the following datasources: `switch-controller.fortilink-settings.name` .
* `switch_controller_rspan_mode` - Stop Layer2 MAC learning and interception of BPDUs and other packets on this interface. Valid values: `disable` `enable` .
* `switch_controller_source_ip` - Source IP address used in FortiLink over L3 connections. Valid values: `outbound` `fixed` .
* `switch_controller_traffic_policy` - Switch controller traffic policy for the VLAN. This attribute must reference one of the following datasources: `switch-controller.traffic-policy.name` .
* `system_id` - Define a system ID for the aggregate interface.
* `system_id_type` - Method in which system ID is generated. Valid values: `auto` `user` .
* `tcp_mss` - TCP maximum segment size. 0 means do not change segment size.
* `trust_ip_1` - Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
* `trust_ip_2` - Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
* `trust_ip_3` - Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
* `trust_ip6_1` - Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
* `trust_ip6_2` - Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
* `trust_ip6_3` - Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
* `type` - Interface type. Valid values: `physical` `vlan` `aggregate` `redundant` `tunnel` `vdom-link` `loopback` `switch` `vap-switch` `wl-mesh` `fext-wan` `vxlan` `geneve` `hdlc` `switch-vlan` `emac-vlan` `ssl` `lan-extension` .
* `username` - Username of the PPPoE account, provided by your ISP.
* `vdom` - Interface is in this virtual domain (VDOM). This attribute must reference one of the following datasources: `system.vdom.name` .
* `vindex` - Switch control interface VLAN ID.
* `vlan_protocol` - Ethernet protocol of VLAN. Valid values: `8021q` `8021ad` .
* `vlanforward` - Enable/disable traffic forwarding between VLANs on this interface. Valid values: `enable` `disable` .
* `vlanid` - VLAN ID (1 - 4094).
* `vrf` - Virtual Routing Forwarding ID.
* `vrrp_virtual_mac` - Enable/disable use of virtual MAC for VRRP. Valid values: `enable` `disable` .
* `wccp` - Enable/disable WCCP on this interface. Used for encapsulated WCCP communication between WCCP clients and servers. Valid values: `enable` `disable` .
* `weight` - Default weight for static routes (if route has no weight configured).
* `wins_ip` - WINS server IP.
* `client_options` - DHCP client options. The structure of `client_options` block is documented below.

The `client_options` block contains:

* `code` - DHCP client option code.
* `id` - ID.
* `ip` - DHCP option IPs.
* `type` - DHCP client option type. Valid values: `hex` `string` `ip` `fqdn` .
* `value` - DHCP client option value.
* `dhcp_snooping_server_list` - Configure DHCP server access list. The structure of `dhcp_snooping_server_list` block is documented below.

The `dhcp_snooping_server_list` block contains:

* `name` - DHCP server name.
* `server_ip` - IP address for DHCP server.
* `egress_queues` - Configure queues of NP port on egress path. The structure of `egress_queues` block is documented below.

The `egress_queues` block contains:

* `cos0` - CoS profile name for CoS 0. This attribute must reference one of the following datasources: `system.isf-queue-profile.name` .
* `cos1` - CoS profile name for CoS 1. This attribute must reference one of the following datasources: `system.isf-queue-profile.name` .
* `cos2` - CoS profile name for CoS 2. This attribute must reference one of the following datasources: `system.isf-queue-profile.name` .
* `cos3` - CoS profile name for CoS 3. This attribute must reference one of the following datasources: `system.isf-queue-profile.name` .
* `cos4` - CoS profile name for CoS 4. This attribute must reference one of the following datasources: `system.isf-queue-profile.name` .
* `cos5` - CoS profile name for CoS 5. This attribute must reference one of the following datasources: `system.isf-queue-profile.name` .
* `cos6` - CoS profile name for CoS 6. This attribute must reference one of the following datasources: `system.isf-queue-profile.name` .
* `cos7` - CoS profile name for CoS 7. This attribute must reference one of the following datasources: `system.isf-queue-profile.name` .
* `fail_alert_interfaces` - Names of the FortiGate interfaces to which the link failure alert is sent. The structure of `fail_alert_interfaces` block is documented below.

The `fail_alert_interfaces` block contains:

* `name` - Names of the non-virtual interface. This attribute must reference one of the following datasources: `system.interface.name` .
* `ipv6` - IPv6 of interface. The structure of `ipv6` block is documented below.

The `ipv6` block contains:

* `autoconf` - Enable/disable address auto config. Valid values: `enable` `disable` .
* `cli_conn6_status` - CLI IPv6 connection status.
* `dhcp6_client_options` - DHCPv6 client options. Valid values: `rapid` `iapd` `iana` .
* `dhcp6_information_request` - Enable/disable DHCPv6 information request. Valid values: `enable` `disable` .
* `dhcp6_prefix_delegation` - Enable/disable DHCPv6 prefix delegation. Valid values: `enable` `disable` .
* `dhcp6_prefix_hint` - DHCPv6 prefix that will be used as a hint to the upstream DHCPv6 server.
* `dhcp6_prefix_hint_plt` - DHCPv6 prefix hint preferred life time (sec), 0 means unlimited lease time.
* `dhcp6_prefix_hint_vlt` - DHCPv6 prefix hint valid life time (sec).
* `dhcp6_relay_ip` - DHCPv6 relay IP address.
* `dhcp6_relay_service` - Enable/disable DHCPv6 relay. Valid values: `disable` `enable` .
* `dhcp6_relay_type` - DHCPv6 relay type. Valid values: `regular` .
* `icmp6_send_redirect` - Enable/disable sending of ICMPv6 redirects. Valid values: `enable` `disable` .
* `interface_identifier` - IPv6 interface identifier.
* `ip6_address` - Primary IPv6 address prefix. Syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx.
* `ip6_allowaccess` - Allow management access to the interface. Valid values: `ping` `https` `ssh` `snmp` `http` `telnet` `fgfm` `fabric` .
* `ip6_default_life` - Default life (sec).
* `ip6_delegated_prefix_iaid` - IAID of obtained delegated-prefix from the upstream interface.
* `ip6_dns_server_override` - Enable/disable using the DNS server acquired by DHCP. Valid values: `enable` `disable` .
* `ip6_hop_limit` - Hop limit (0 means unspecified).
* `ip6_link_mtu` - IPv6 link MTU.
* `ip6_manage_flag` - Enable/disable the managed flag. Valid values: `enable` `disable` .
* `ip6_max_interval` - IPv6 maximum interval (4 to 1800 sec).
* `ip6_min_interval` - IPv6 minimum interval (3 to 1350 sec).
* `ip6_mode` - Addressing mode (static, DHCP, delegated). Valid values: `static` `dhcp` `pppoe` `delegated` .
* `ip6_other_flag` - Enable/disable the other IPv6 flag. Valid values: `enable` `disable` .
* `ip6_prefix_mode` - Assigning a prefix from DHCP or RA. Valid values: `dhcp6` `ra` .
* `ip6_reachable_time` - IPv6 reachable time (milliseconds; 0 means unspecified).
* `ip6_retrans_time` - IPv6 retransmit time (milliseconds; 0 means unspecified).
* `ip6_send_adv` - Enable/disable sending advertisements about the interface. Valid values: `enable` `disable` .
* `ip6_subnet` - Subnet to routing prefix. Syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx.
* `ip6_upstream_interface` - Interface name providing delegated information. This attribute must reference one of the following datasources: `system.interface.name` .
* `nd_cert` - Neighbor discovery certificate. This attribute must reference one of the following datasources: `certificate.local.name` .
* `nd_cga_modifier` - Neighbor discovery CGA modifier.
* `nd_mode` - Neighbor discovery mode. Valid values: `basic` `SEND-compatible` .
* `nd_security_level` - Neighbor discovery security level (0 - 7; 0 = least secure, default = 0).
* `nd_timestamp_delta` - Neighbor discovery timestamp delta value (1 - 3600 sec; default = 300).
* `nd_timestamp_fuzz` - Neighbor discovery timestamp fuzz factor (1 - 60 sec; default = 1).
* `ra_send_mtu` - Enable/disable sending link MTU in RA packet. Valid values: `enable` `disable` .
* `unique_autoconf_addr` - Enable/disable unique auto config address. Valid values: `enable` `disable` .
* `vrip6_link_local` - Link-local IPv6 address of virtual router.
* `vrrp_virtual_mac6` - Enable/disable virtual MAC for VRRP. Valid values: `enable` `disable` .
* `dhcp6_iapd_list` - DHCPv6 IA-PD list. The structure of `dhcp6_iapd_list` block is documented below.

The `dhcp6_iapd_list` block contains:

* `iaid` - Identity association identifier.
* `prefix_hint` - DHCPv6 prefix that will be used as a hint to the upstream DHCPv6 server.
* `prefix_hint_plt` - DHCPv6 prefix hint preferred life time (sec), 0 means unlimited lease time.
* `prefix_hint_vlt` - DHCPv6 prefix hint valid life time (sec).
* `ip6_delegated_prefix_list` - Advertised IPv6 delegated prefix list. The structure of `ip6_delegated_prefix_list` block is documented below.

The `ip6_delegated_prefix_list` block contains:

* `autonomous_flag` - Enable/disable the autonomous flag. Valid values: `enable` `disable` .
* `delegated_prefix_iaid` - IAID of obtained delegated-prefix from the upstream interface.
* `onlink_flag` - Enable/disable the onlink flag. Valid values: `enable` `disable` .
* `prefix_id` - Prefix ID.
* `rdnss` - Recursive DNS server option.
* `rdnss_service` - Recursive DNS service option. Valid values: `delegated` `default` `specify` .
* `subnet` - Add subnet ID to routing prefix.
* `upstream_interface` - Name of the interface that provides delegated information. This attribute must reference one of the following datasources: `system.interface.name` .
* `ip6_extra_addr` - Extra IPv6 address prefixes of interface. The structure of `ip6_extra_addr` block is documented below.

The `ip6_extra_addr` block contains:

* `prefix` - IPv6 address prefix.
* `ip6_prefix_list` - Advertised prefix list. The structure of `ip6_prefix_list` block is documented below.

The `ip6_prefix_list` block contains:

* `autonomous_flag` - Enable/disable the autonomous flag. Valid values: `enable` `disable` .
* `onlink_flag` - Enable/disable the onlink flag. Valid values: `enable` `disable` .
* `preferred_life_time` - Preferred life time (sec).
* `prefix` - IPv6 prefix.
* `rdnss` - Recursive DNS server option.
* `valid_life_time` - Valid life time (sec).
* `dnssl` - DNS search list option. The structure of `dnssl` block is documented below.

The `dnssl` block contains:

* `domain` - Domain name.
* `vrrp6` - IPv6 VRRP configuration. The structure of `vrrp6` block is documented below.

The `vrrp6` block contains:

* `accept_mode` - Enable/disable accept mode. Valid values: `enable` `disable` .
* `adv_interval` - Advertisement interval (1 - 255 seconds).
* `preempt` - Enable/disable preempt mode. Valid values: `enable` `disable` .
* `priority` - Priority of the virtual router (1 - 255).
* `start_time` - Startup time (1 - 255 seconds).
* `status` - Enable/disable VRRP. Valid values: `enable` `disable` .
* `vrdst6` - Monitor the route to this destination.
* `vrgrp` - VRRP group ID (1 - 65535).
* `vrid` - Virtual router identifier (1 - 255).
* `vrip6` - IPv6 address of the virtual router.
* `member` - Physical interfaces that belong to the aggregate or redundant interface. The structure of `member` block is documented below.

The `member` block contains:

* `interface_name` - Physical interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `secondaryip` - Second IP address of interface. The structure of `secondaryip` block is documented below.

The `secondaryip` block contains:

* `allowaccess` - Management access settings for the secondary IP address. Valid values: `ping` `https` `ssh` `snmp` `http` `telnet` `fgfm` `radius-acct` `probe-response` `fabric` `ftm` `speed-test` .
* `detectprotocol` - Protocols used to detect the server. Valid values: `ping` `tcp-echo` `udp-echo` .
* `detectserver` - Gateway's ping server for this IP.
* `gwdetect` - Enable/disable detect gateway alive for first. Valid values: `enable` `disable` .
* `ha_priority` - HA election priority for the PING server.
* `id` - ID.
* `ip` - Secondary IP address of the interface.
* `ping_serv_status` - PING server status.
* `security_groups` - User groups that can authenticate with the captive portal. The structure of `security_groups` block is documented below.

The `security_groups` block contains:

* `name` - Names of user groups that can authenticate with the captive portal. This attribute must reference one of the following datasources: `user.group.name` .
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.

The `tagging` block contains:

* `category` - Tag category. This attribute must reference one of the following datasources: `system.object-tagging.category` .
* `name` - Tagging entry name.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name. This attribute must reference one of the following datasources: `system.object-tagging.tags.name` .
* `vrrp` - VRRP configuration. The structure of `vrrp` block is documented below.

The `vrrp` block contains:

* `accept_mode` - Enable/disable accept mode. Valid values: `enable` `disable` .
* `adv_interval` - Advertisement interval (1 - 255 seconds).
* `ignore_default_route` - Enable/disable ignoring of default route when checking destination. Valid values: `enable` `disable` .
* `preempt` - Enable/disable preempt mode. Valid values: `enable` `disable` .
* `priority` - Priority of the virtual router (1 - 255).
* `start_time` - Startup time (1 - 255 seconds).
* `status` - Enable/disable this VRRP configuration. Valid values: `enable` `disable` .
* `version` - VRRP version. Valid values: `2` `3` .
* `vrdst` - Monitor the route to this destination.
* `vrdst_priority` - Priority of the virtual router when the virtual router destination becomes unreachable (0 - 254).
* `vrgrp` - VRRP group ID (1 - 65535).
* `vrid` - Virtual router identifier (1 - 255).
* `vrip` - IP address of the virtual router.
* `proxy_arp` - VRRP Proxy ARP configuration. The structure of `proxy_arp` block is documented below.

The `proxy_arp` block contains:

* `id` - ID.
* `ip` - Set IP addresses of proxy ARP.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_interface can be imported using:
```sh
terraform import fortios_system_interface.labelname {{mkey}}
```
