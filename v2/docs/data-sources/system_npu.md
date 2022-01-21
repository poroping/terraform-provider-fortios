---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_npu"
description: |-
  Get information on a fortios Configure NPU attributes.
---

# Data Source: fortios_system_npu
Use this data source to get information on a fortios Configure NPU attributes.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `capwap_offload` - Enable/disable offloading managed FortiAP and FortiLink CAPWAP sessions.
* `dedicated_management_affinity` - Affinity setting for management deamons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
* `dedicated_management_cpu` - Enable to dedicate one CPU for GUI and CLI connections when NPs are busy.
* `default_qos_type` - Set default QoS type.
* `double_level_mcast_offload` - Enable double level mcast offload.
* `dse_timeout` -  DSE timeout in seconds (0-3600, default = 10).
* `fastpath` - Enable/disable NP6 offloading (also called fast path).
* `gtp_support` - Enable/Disable NP7 GTP support
* `hash_tbl_spread` - Enable/disable hash table entry spread (default enabled).
* `htab_dedi_queue_nr` - Set the number of dedicate queue for hash table messages.
* `htab_msg_queue` - Set hash table message queue mode.
* `ippool_overload_high` -  High threshold for overload ippool port reuse (100%-2000%, default = 200).
* `ippool_overload_low` -  Low threshold for overload ippool port reuse (100%-2000%, default = 150).
* `ipsec_dec_subengine_mask` - IPsec decryption subengine mask (0x1 - 0xff, default 0xff).
* `ipsec_enc_subengine_mask` - IPsec encryption subengine mask (0x1 - 0xff, default 0xff).
* `ipsec_inbound_cache` - Enable/disable IPsec inbound cache for anti-replay.
* `ipsec_mtu_override` - Enable/disable NP6 IPsec MTU override.
* `ipsec_ob_np_sel` - IPsec NP selection for OB SA offloading.
* `ipsec_over_vlink` - Enable/disable IPSEC over vlink.
* `max_session_timeout` - Maximum time interval for refreshing NPU-offloaded sessions (10 - 1000 sec, default 40 sec).
* `mcast_session_accounting` - Enable/disable traffic accounting for each multicast session through TAE counter.
* `napi_break_interval` -  NAPI break interval (default 0).
* `np6_cps_optimization_mode` - Enable/disable NP6 connection per second (CPS) optimization mode.
* `pba_eim` - Configure option for PBA(non-overload)/EIM combination.
* `per_session_accounting` - Set per-session accounting.
* `policy_offload_level` - Firewall Policy Offload Level(DISABLE/DOS/FULL).
* `qos_mode` - QoS mode on switch and NP.
* `rdp_offload` - Enable/disable rdp offload.
* `recover_np6_link` - Enable/disable internal link failure check and recovery after boot up.
* `session_acct_interval` - Session accounting update interval (1 - 10 sec, default 5 sec).
* `sse_backpressure` - Enable/disable sse backpressure.
* `strip_clear_text_padding` - Enable/disable stripping clear text padding.
* `strip_esp_padding` - Enable/disable stripping ESP padding.
* `sw_np_bandwidth` - Bandwidth from switch to NP.
* `tcp_rst_timeout` -  TCP RST timeout in seconds (0-3600, default = 5).
* `vlan_lookup_cache` - Enable/disable vlan lookup cache (default enabled).
* `dos_options` - NPU DoS configurations.The structure of `dos_options` block is documented below.

The `dos_options` block contains:

* `npu_dos_meter_mode` - Set DoS meter npu offloading mode.
* `npu_dos_tpe_mode` - Enable/Disable inserting DoS meter id to session table.
* `dsw_dts_profile` - Configure NPU DSW DTS profile.The structure of `dsw_dts_profile` block is documented below.

The `dsw_dts_profile` block contains:

* `action` - Set NPU DSW DTS profile action.
* `min_limit` - Set NPU DSW DTS profile min-limt.
* `profile_id` - Set NPU DSW DTS profile profile id.
* `step` - Set NPU DSW DTS profile step.
* `dsw_queue_dts_profile` - Configure NPU DSW Queue DTS profile.The structure of `dsw_queue_dts_profile` block is documented below.

The `dsw_queue_dts_profile` block contains:

* `iport` - Set NPU DSW DTS in port.
* `name` - Name.
* `oport` - Set NPU DSW DTS out port.
* `profile_id` - Set NPU DSW DTS profile id.
* `queue_select` - Set NPU DSW DTS queue id select(0 - reset to default).
* `fp_anomaly` - IPv4/IPv6 anomaly protection.The structure of `fp_anomaly` block is documented below.

The `fp_anomaly` block contains:

* `icmp_csum_err` - Invalid IPv4 ICMP checksum anomalies.
* `icmp_frag` - Layer 3 fragmented packets that could be part of layer 4 ICMP anomalies.
* `icmp_land` - ICMP land anomalies.
* `ipv4_csum_err` - Invalid IPv4 IP checksum anomalies.
* `ipv4_land` - Land anomalies.
* `ipv4_optlsrr` - Loose source record route option anomalies.
* `ipv4_optrr` - Record route option anomalies.
* `ipv4_optsecurity` - Security option anomalies.
* `ipv4_optssrr` - Strict source record route option anomalies.
* `ipv4_optstream` - Stream option anomalies.
* `ipv4_opttimestamp` - Timestamp option anomalies.
* `ipv4_proto_err` - Invalid layer 4 protocol anomalies.
* `ipv4_unknopt` - Unknown option anomalies.
* `ipv6_daddr_err` - Destination address as unspecified or loopback address anomalies.
* `ipv6_land` - Land anomalies.
* `ipv6_optendpid` - End point identification anomalies.
* `ipv6_opthomeaddr` - Home address option anomalies.
* `ipv6_optinvld` - Invalid option anomalies.Invalid option anomalies.
* `ipv6_optjumbo` - Jumbo options anomalies.
* `ipv6_optnsap` - Network service access point address option anomalies.
* `ipv6_optralert` - Router alert option anomalies.
* `ipv6_opttunnel` - Tunnel encapsulation limit option anomalies.
* `ipv6_proto_err` - Layer 4 invalid protocol anomalies.
* `ipv6_saddr_err` - Source address as multicast anomalies.
* `ipv6_unknopt` - Unknown option anomalies.
* `tcp_csum_err` - Invalid IPv4 TCP checksum anomalies.
* `tcp_fin_noack` - TCP SYN flood with FIN flag set without ACK setting anomalies.
* `tcp_fin_only` - TCP SYN flood with only FIN flag set anomalies.
* `tcp_land` - TCP land anomalies.
* `tcp_no_flag` - TCP SYN flood with no flag set anomalies.
* `tcp_syn_data` - TCP SYN flood packets with data anomalies.
* `tcp_syn_fin` - TCP SYN flood SYN/FIN flag set anomalies. 
* `tcp_winnuke` - TCP WinNuke anomalies.
* `udp_csum_err` - Invalid IPv4 UDP checksum anomalies.
* `udp_land` - UDP land anomalies.
* `hpe` - Host protection engine configuration.The structure of `hpe` block is documented below.

The `hpe` block contains:

* `all_protocol` - Maximum packet rate of each host queue except high priority traffic(1K - 40M pps, default = 400K pps), set 0 to disable.
* `arp_max` - Maximum ARP packet rate (1K - 40M pps, default = 20K pps).
* `enable_shaper` - Enable/Disable NPU Host Protection Engine (HPE) for packet type shaper.
* `esp_max` - Maximum ESP packet rate (1K - 40M pps, default = 20K pps).
* `high_priority` - Maximum packet rate for high priority traffic packets (1K - 40M pps, default = 400K pps).
* `icmp_max` - Maximum ICMP packet rate (1K - 40M pps, default = 20K pps).
* `ip_frag_max` - Maximum fragmented IP packet rate (1K - 40M pps, default = 20K pps).
* `ip_others_max` - Maximum IP packet rate for other packets (packet types that cannot be set with other options) (1K - 1G pps, default = 20K pps).
* `l2_others_max` - Maximum L2 packet rate for L2 packets that are not ARP packets (1K - 40M pps, default = 20K pps).
* `sctp_max` - Maximum SCTP packet rate (1K - 40M pps, default = 20K pps).
* `tcp_max` - Maximum TCP packet rate (1K - 40M pps, default = 40K pps).
* `tcpfin_rst_max` - Maximum TCP carries FIN or RST flags packet rate (1K - 40M pps, default = 40K pps).
* `tcpsyn_ack_max` - Maximum TCP carries SYN and ACK flags packet rate (1K - 40M pps, default = 40K pps).
* `tcpsyn_max` - Maximum TCP SYN packet rate (1K - 40M pps, default = 40K pps).
* `udp_max` - Maximum UDP packet rate (1K - 40M pps, default = 40K pps).
* `inbound_dscp_copy_port` - Physical interfaces that support inbound-dscp-copy.The structure of `inbound_dscp_copy_port` block is documented below.

The `inbound_dscp_copy_port` block contains:

* `interface` - Physical interface name.
* `ip_reassembly` - IP reassebmly engine configuration.The structure of `ip_reassembly` block is documented below.

The `ip_reassembly` block contains:

* `max_timeout` - Maximum timeout value for IP reassembly (5 us - 600,000,000 us).
* `min_timeout` - Minimum timeout value for IP reassembly (5 us - 600,000,000 us).
* `status` - Set IP reassembly processing status.
* `isf_np_queues` - Configure queues of switch port connected to NP6 XAUI on ingress path.The structure of `isf_np_queues` block is documented below.

The `isf_np_queues` block contains:

* `cos0` - CoS profile name for CoS 0.
* `cos1` - CoS profile name for CoS 1.
* `cos2` - CoS profile name for CoS 2.
* `cos3` - CoS profile name for CoS 3.
* `cos4` - CoS profile name for CoS 4.
* `cos5` - CoS profile name for CoS 5.
* `cos6` - CoS profile name for CoS 6.
* `cos7` - CoS profile name for CoS 7.
* `np_queues` - Configure queue assignment on NP7.The structure of `np_queues` block is documented below.

The `np_queues` block contains:

* `ethernet_type` - Configure a NP7 QoS Ethernet Type.The structure of `ethernet_type` block is documented below.

The `ethernet_type` block contains:

* `name` - Ethernet Type Name.
* `queue` - Queue Number.
* `type` - Ethernet Type.
* `weight` - Class Weight.
* `ip_protocol` - Configure a NP7 QoS IP Protocol.The structure of `ip_protocol` block is documented below.

The `ip_protocol` block contains:

* `name` - IP Protocol Name.
* `protocol` - IP Protocol.
* `queue` - Queue Number.
* `weight` - Class Weight.
* `ip_service` - Configure a NP7 QoS IP Service.The structure of `ip_service` block is documented below.

The `ip_service` block contains:

* `dport` - Destination Port.
* `name` - IP Service Name.
* `protocol` - IP Protocol.
* `queue` - Queue Number.
* `sport` - Source Port.
* `weight` - Class Weight.
* `profile` - Configure a NP7 class profile.The structure of `profile` block is documented below.

The `profile` block contains:

* `cos0` - Queue number of CoS 0.
* `cos1` - Queue number of CoS 1.
* `cos2` - Queue number of CoS 2.
* `cos3` - Queue number of CoS 3.
* `cos4` - Queue number of CoS 4.
* `cos5` - Queue number of CoS 5.
* `cos6` - Queue number of CoS 6.
* `cos7` - Queue number of CoS 7.
* `dscp0` - Queue number of DSCP 0.
* `dscp1` - Queue number of DSCP 1.
* `dscp10` - Queue number of DSCP 10.
* `dscp11` - Queue number of DSCP 11.
* `dscp12` - Queue number of DSCP 12.
* `dscp13` - Queue number of DSCP 13.
* `dscp14` - Queue number of DSCP 14.
* `dscp15` - Queue number of DSCP 15.
* `dscp16` - Queue number of DSCP 16.
* `dscp17` - Queue number of DSCP 17.
* `dscp18` - Queue number of DSCP 18.
* `dscp19` - Queue number of DSCP 19.
* `dscp2` - Queue number of DSCP 2.
* `dscp20` - Queue number of DSCP 20.
* `dscp21` - Queue number of DSCP 21.
* `dscp22` - Queue number of DSCP 22.
* `dscp23` - Queue number of DSCP 23.
* `dscp24` - Queue number of DSCP 24.
* `dscp25` - Queue number of DSCP 25.
* `dscp26` - Queue number of DSCP 26.
* `dscp27` - Queue number of DSCP 27.
* `dscp28` - Queue number of DSCP 28.
* `dscp29` - Queue number of DSCP 29.
* `dscp3` - Queue number of DSCP 3.
* `dscp30` - Queue number of DSCP 30.
* `dscp31` - Queue number of DSCP 31.
* `dscp32` - Queue number of DSCP 32.
* `dscp33` - Queue number of DSCP 33.
* `dscp34` - Queue number of DSCP 34.
* `dscp35` - Queue number of DSCP 35.
* `dscp36` - Queue number of DSCP 36.
* `dscp37` - Queue number of DSCP 37.
* `dscp38` - Queue number of DSCP 38.
* `dscp39` - Queue number of DSCP 39.
* `dscp4` - Queue number of DSCP 4.
* `dscp40` - Queue number of DSCP 40.
* `dscp41` - Queue number of DSCP 41.
* `dscp42` - Queue number of DSCP 42.
* `dscp43` - Queue number of DSCP 43.
* `dscp44` - Queue number of DSCP 44.
* `dscp45` - Queue number of DSCP 45.
* `dscp46` - Queue number of DSCP 46.
* `dscp47` - Queue number of DSCP 47.
* `dscp48` - Queue number of DSCP 48.
* `dscp49` - Queue number of DSCP 49.
* `dscp5` - Queue number of DSCP 5.
* `dscp50` - Queue number of DSCP 50.
* `dscp51` - Queue number of DSCP 51.
* `dscp52` - Queue number of DSCP 52.
* `dscp53` - Queue number of DSCP 53.
* `dscp54` - Queue number of DSCP 54.
* `dscp55` - Queue number of DSCP 55.
* `dscp56` - Queue number of DSCP 56.
* `dscp57` - Queue number of DSCP 57.
* `dscp58` - Queue number of DSCP 58.
* `dscp59` - Queue number of DSCP 59.
* `dscp6` - Queue number of DSCP 6.
* `dscp60` - Queue number of DSCP 60.
* `dscp61` - Queue number of DSCP 61.
* `dscp62` - Queue number of DSCP 62.
* `dscp63` - Queue number of DSCP 63.
* `dscp7` - Queue number of DSCP 7.
* `dscp8` - Queue number of DSCP 8.
* `dscp9` - Queue number of DSCP 9.
* `id` - Profile ID.
* `type` - Profile type.
* `weight` - Class weight.
* `scheduler` - Configure a NP7 QoS Scheduler.The structure of `scheduler` block is documented below.

The `scheduler` block contains:

* `mode` - Scheduler Mode.
* `name` - Scheduler Name.
* `port_cpu_map` - Configure NPU interface to CPU core mapping.The structure of `port_cpu_map` block is documented below.

The `port_cpu_map` block contains:

* `cpu_core` - The CPU core to map to an interface.
* `interface` - The interface to map to a CPU core.
* `port_npu_map` - Configure port to NPU group mapping.The structure of `port_npu_map` block is documented below.

The `port_npu_map` block contains:

* `interface` - Set npu interface port to NPU group map.
* `npu_group_index` - Mapping NPU group index.
* `priority_protocol` - Configure NPU priority protocol.The structure of `priority_protocol` block is documented below.

The `priority_protocol` block contains:

* `bfd` - Enable/disable NPU BFD priority protocol.
* `bgp` - Enable/disable NPU BGP priority protocol.
* `slbc` - Enable/disable NPU SLBC priority protocol.
* `tcp_timeout_profile` - Configure TCP timeout profile.The structure of `tcp_timeout_profile` block is documented below.

The `tcp_timeout_profile` block contains:

* `close_wait` - Set close-wait timeout(seconds)
* `fin_wait` - Set fin-wait timeout(seconds)
* `id` - Timeout profile ID (5-47)
* `syn_sent` - Set syn-sent timeout(seconds)
* `syn_wait` - Set syn-wait timeout(seconds)
* `tcp_idle` - Set TCP establish timeout(seconds)
* `time_wait` - Set time-wait timeout(seconds)
* `udp_timeout_profile` - Configure UDP timeout profile.The structure of `udp_timeout_profile` block is documented below.

The `udp_timeout_profile` block contains:

* `id` - Timeout profile ID (5-63)
* `udp_idle` - Set UDP idle timeout(seconds)
