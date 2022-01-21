---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_npu"
description: |-
  Configure NPU attributes.
---

## fortios_system_npu
Configure NPU attributes.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `capwap_offload` - Enable/disable offloading managed FortiAP and FortiLink CAPWAP sessions. Valid values: `enable` `disable` .
* `dedicated_management_affinity` - Affinity setting for management deamons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
* `dedicated_management_cpu` - Enable to dedicate one CPU for GUI and CLI connections when NPs are busy. Valid values: `enable` `disable` .
* `default_qos_type` - Set default QoS type. Valid values: `policing` `shaping` .
* `double_level_mcast_offload` - Enable double level mcast offload. Valid values: `enable` `disable` .
* `dse_timeout` -  DSE timeout in seconds (0-3600, default = 10).
* `fastpath` - Enable/disable NP6 offloading (also called fast path). Valid values: `disable` `enable` .
* `gtp_support` - Enable/Disable NP7 GTP support Valid values: `enable` `disable` .
* `hash_tbl_spread` - Enable/disable hash table entry spread (default enabled). Valid values: `enable` `disable` .
* `htab_dedi_queue_nr` - Set the number of dedicate queue for hash table messages.
* `htab_msg_queue` - Set hash table message queue mode. Valid values: `data` `idle` `dedicated` .
* `ippool_overload_high` -  High threshold for overload ippool port reuse (100%-2000%, default = 200).
* `ippool_overload_low` -  Low threshold for overload ippool port reuse (100%-2000%, default = 150).
* `ipsec_dec_subengine_mask` - IPsec decryption subengine mask (0x1 - 0xff, default 0xff).
* `ipsec_enc_subengine_mask` - IPsec encryption subengine mask (0x1 - 0xff, default 0xff).
* `ipsec_inbound_cache` - Enable/disable IPsec inbound cache for anti-replay. Valid values: `enable` `disable` .
* `ipsec_mtu_override` - Enable/disable NP6 IPsec MTU override. Valid values: `disable` `enable` .
* `ipsec_ob_np_sel` - IPsec NP selection for OB SA offloading. Valid values: `RR` `Packet` `Hash` .
* `ipsec_over_vlink` - Enable/disable IPSEC over vlink. Valid values: `enable` `disable` .
* `max_session_timeout` - Maximum time interval for refreshing NPU-offloaded sessions (10 - 1000 sec, default 40 sec).
* `mcast_session_accounting` - Enable/disable traffic accounting for each multicast session through TAE counter. Valid values: `tpe-based` `session-based` `disable` .
* `napi_break_interval` -  NAPI break interval (default 0).
* `np6_cps_optimization_mode` - Enable/disable NP6 connection per second (CPS) optimization mode. Valid values: `enable` `disable` .
* `pba_eim` - Configure option for PBA(non-overload)/EIM combination. Valid values: `disallow` `allow` .
* `per_session_accounting` - Set per-session accounting. Valid values: `traffic-log-only` `disable` `enable` .
* `policy_offload_level` - Firewall Policy Offload Level(DISABLE/DOS/FULL). Valid values: `disable` `dos-offload` .
* `qos_mode` - QoS mode on switch and NP. Valid values: `disable` `priority` `round-robin` .
* `rdp_offload` - Enable/disable rdp offload. Valid values: `enable` `disable` .
* `recover_np6_link` - Enable/disable internal link failure check and recovery after boot up. Valid values: `enable` `disable` .
* `session_acct_interval` - Session accounting update interval (1 - 10 sec, default 5 sec).
* `sse_backpressure` - Enable/disable sse backpressure. Valid values: `enable` `disable` .
* `strip_clear_text_padding` - Enable/disable stripping clear text padding. Valid values: `enable` `disable` .
* `strip_esp_padding` - Enable/disable stripping ESP padding. Valid values: `enable` `disable` .
* `sw_np_bandwidth` - Bandwidth from switch to NP. Valid values: `0G` `2G` `4G` `5G` `6G` .
* `tcp_rst_timeout` -  TCP RST timeout in seconds (0-3600, default = 5).
* `vlan_lookup_cache` - Enable/disable vlan lookup cache (default enabled). Valid values: `enable` `disable` .
* `dos_options` - NPU DoS configurations. The structure of `dos_options` block is documented below.

The `dos_options` block contains:

* `npu_dos_meter_mode` - Set DoS meter npu offloading mode. Valid values: `global` `local` .
* `npu_dos_tpe_mode` - Enable/Disable inserting DoS meter id to session table. Valid values: `enable` `disable` .
* `dsw_dts_profile` - Configure NPU DSW DTS profile. The structure of `dsw_dts_profile` block is documented below.

The `dsw_dts_profile` block contains:

* `action` - Set NPU DSW DTS profile action. Valid values: `wait` `drop` `drop_tmr_0` `drop_tmr_1` `enque` `enque_0` `enque_1` .
* `min_limit` - Set NPU DSW DTS profile min-limt.
* `profile_id` - Set NPU DSW DTS profile profile id.
* `step` - Set NPU DSW DTS profile step.
* `dsw_queue_dts_profile` - Configure NPU DSW Queue DTS profile. The structure of `dsw_queue_dts_profile` block is documented below.

The `dsw_queue_dts_profile` block contains:

* `iport` - Set NPU DSW DTS in port. Valid values: `EIF0` `EIF1` `EIF2` `EIF3` `EIF4` `EIF5` `EIF6` `EIF7` `HTX0` `HTX1` `SSE0` `SSE1` `SSE2` `SSE3` `RLT` `DFR` `IPSECI` `IPSECO` `IPTI` `IPTO` `VEP0` `VEP2` `VEP4` `VEP6` `IVS` `L2TI1` `L2TO` `L2TI0` `PLE` `SPATH` `QTM` .
* `name` - Name.
* `oport` - Set NPU DSW DTS out port. Valid values: `EIF0` `EIF1` `EIF2` `EIF3` `EIF4` `EIF5` `EIF6` `EIF7` `HRX` `SSE0` `SSE1` `SSE2` `SSE3` `RLT` `DFR` `IPSECI` `IPSECO` `IPTI` `IPTO` `VEP0` `VEP2` `VEP4` `VEP6` `IVS` `L2TI1` `L2TO` `L2TI0` `PLE` `SYNK` `NSS` `TSK` `QTM` .
* `profile_id` - Set NPU DSW DTS profile id.
* `queue_select` - Set NPU DSW DTS queue id select(0 - reset to default).
* `fp_anomaly` - IPv4/IPv6 anomaly protection. The structure of `fp_anomaly` block is documented below.

The `fp_anomaly` block contains:

* `icmp_csum_err` - Invalid IPv4 ICMP checksum anomalies. Valid values: `drop` `trap-to-host` .
* `icmp_frag` - Layer 3 fragmented packets that could be part of layer 4 ICMP anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `icmp_land` - ICMP land anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `ipv4_csum_err` - Invalid IPv4 IP checksum anomalies. Valid values: `drop` `trap-to-host` .
* `ipv4_land` - Land anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `ipv4_optlsrr` - Loose source record route option anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `ipv4_optrr` - Record route option anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `ipv4_optsecurity` - Security option anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `ipv4_optssrr` - Strict source record route option anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `ipv4_optstream` - Stream option anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `ipv4_opttimestamp` - Timestamp option anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `ipv4_proto_err` - Invalid layer 4 protocol anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `ipv4_unknopt` - Unknown option anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `ipv6_daddr_err` - Destination address as unspecified or loopback address anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `ipv6_land` - Land anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `ipv6_optendpid` - End point identification anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `ipv6_opthomeaddr` - Home address option anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `ipv6_optinvld` - Invalid option anomalies.Invalid option anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `ipv6_optjumbo` - Jumbo options anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `ipv6_optnsap` - Network service access point address option anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `ipv6_optralert` - Router alert option anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `ipv6_opttunnel` - Tunnel encapsulation limit option anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `ipv6_proto_err` - Layer 4 invalid protocol anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `ipv6_saddr_err` - Source address as multicast anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `ipv6_unknopt` - Unknown option anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `tcp_csum_err` - Invalid IPv4 TCP checksum anomalies. Valid values: `drop` `trap-to-host` .
* `tcp_fin_noack` - TCP SYN flood with FIN flag set without ACK setting anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `tcp_fin_only` - TCP SYN flood with only FIN flag set anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `tcp_land` - TCP land anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `tcp_no_flag` - TCP SYN flood with no flag set anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `tcp_syn_data` - TCP SYN flood packets with data anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `tcp_syn_fin` - TCP SYN flood SYN/FIN flag set anomalies.  Valid values: `allow` `drop` `trap-to-host` .
* `tcp_winnuke` - TCP WinNuke anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `udp_csum_err` - Invalid IPv4 UDP checksum anomalies. Valid values: `drop` `trap-to-host` .
* `udp_land` - UDP land anomalies. Valid values: `allow` `drop` `trap-to-host` .
* `hpe` - Host protection engine configuration. The structure of `hpe` block is documented below.

The `hpe` block contains:

* `all_protocol` - Maximum packet rate of each host queue except high priority traffic(1K - 40M pps, default = 400K pps), set 0 to disable.
* `arp_max` - Maximum ARP packet rate (1K - 40M pps, default = 20K pps).
* `enable_shaper` - Enable/Disable NPU Host Protection Engine (HPE) for packet type shaper. Valid values: `disable` `enable` .
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
* `inbound_dscp_copy_port` - Physical interfaces that support inbound-dscp-copy. The structure of `inbound_dscp_copy_port` block is documented below.

The `inbound_dscp_copy_port` block contains:

* `interface` - Physical interface name.
* `ip_reassembly` - IP reassebmly engine configuration. The structure of `ip_reassembly` block is documented below.

The `ip_reassembly` block contains:

* `max_timeout` - Maximum timeout value for IP reassembly (5 us - 600,000,000 us).
* `min_timeout` - Minimum timeout value for IP reassembly (5 us - 600,000,000 us).
* `status` - Set IP reassembly processing status. Valid values: `disable` `enable` .
* `isf_np_queues` - Configure queues of switch port connected to NP6 XAUI on ingress path. The structure of `isf_np_queues` block is documented below.

The `isf_np_queues` block contains:

* `cos0` - CoS profile name for CoS 0. This attribute must reference one of the following datasources: `system.isf-queue-profile.name` .
* `cos1` - CoS profile name for CoS 1. This attribute must reference one of the following datasources: `system.isf-queue-profile.name` .
* `cos2` - CoS profile name for CoS 2. This attribute must reference one of the following datasources: `system.isf-queue-profile.name` .
* `cos3` - CoS profile name for CoS 3. This attribute must reference one of the following datasources: `system.isf-queue-profile.name` .
* `cos4` - CoS profile name for CoS 4. This attribute must reference one of the following datasources: `system.isf-queue-profile.name` .
* `cos5` - CoS profile name for CoS 5. This attribute must reference one of the following datasources: `system.isf-queue-profile.name` .
* `cos6` - CoS profile name for CoS 6. This attribute must reference one of the following datasources: `system.isf-queue-profile.name` .
* `cos7` - CoS profile name for CoS 7. This attribute must reference one of the following datasources: `system.isf-queue-profile.name` .
* `np_queues` - Configure queue assignment on NP7. The structure of `np_queues` block is documented below.

The `np_queues` block contains:

* `ethernet_type` - Configure a NP7 QoS Ethernet Type. The structure of `ethernet_type` block is documented below.

The `ethernet_type` block contains:

* `name` - Ethernet Type Name.
* `queue` - Queue Number.
* `type` - Ethernet Type.
* `weight` - Class Weight.
* `ip_protocol` - Configure a NP7 QoS IP Protocol. The structure of `ip_protocol` block is documented below.

The `ip_protocol` block contains:

* `name` - IP Protocol Name.
* `protocol` - IP Protocol.
* `queue` - Queue Number.
* `weight` - Class Weight.
* `ip_service` - Configure a NP7 QoS IP Service. The structure of `ip_service` block is documented below.

The `ip_service` block contains:

* `dport` - Destination Port.
* `name` - IP Service Name.
* `protocol` - IP Protocol.
* `queue` - Queue Number.
* `sport` - Source Port.
* `weight` - Class Weight.
* `profile` - Configure a NP7 class profile. The structure of `profile` block is documented below.

The `profile` block contains:

* `cos0` - Queue number of CoS 0. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `cos1` - Queue number of CoS 1. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `cos2` - Queue number of CoS 2. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `cos3` - Queue number of CoS 3. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `cos4` - Queue number of CoS 4. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `cos5` - Queue number of CoS 5. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `cos6` - Queue number of CoS 6. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `cos7` - Queue number of CoS 7. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp0` - Queue number of DSCP 0. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp1` - Queue number of DSCP 1. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp10` - Queue number of DSCP 10. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp11` - Queue number of DSCP 11. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp12` - Queue number of DSCP 12. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp13` - Queue number of DSCP 13. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp14` - Queue number of DSCP 14. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp15` - Queue number of DSCP 15. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp16` - Queue number of DSCP 16. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp17` - Queue number of DSCP 17. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp18` - Queue number of DSCP 18. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp19` - Queue number of DSCP 19. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp2` - Queue number of DSCP 2. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp20` - Queue number of DSCP 20. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp21` - Queue number of DSCP 21. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp22` - Queue number of DSCP 22. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp23` - Queue number of DSCP 23. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp24` - Queue number of DSCP 24. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp25` - Queue number of DSCP 25. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp26` - Queue number of DSCP 26. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp27` - Queue number of DSCP 27. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp28` - Queue number of DSCP 28. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp29` - Queue number of DSCP 29. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp3` - Queue number of DSCP 3. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp30` - Queue number of DSCP 30. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp31` - Queue number of DSCP 31. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp32` - Queue number of DSCP 32. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp33` - Queue number of DSCP 33. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp34` - Queue number of DSCP 34. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp35` - Queue number of DSCP 35. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp36` - Queue number of DSCP 36. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp37` - Queue number of DSCP 37. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp38` - Queue number of DSCP 38. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp39` - Queue number of DSCP 39. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp4` - Queue number of DSCP 4. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp40` - Queue number of DSCP 40. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp41` - Queue number of DSCP 41. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp42` - Queue number of DSCP 42. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp43` - Queue number of DSCP 43. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp44` - Queue number of DSCP 44. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp45` - Queue number of DSCP 45. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp46` - Queue number of DSCP 46. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp47` - Queue number of DSCP 47. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp48` - Queue number of DSCP 48. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp49` - Queue number of DSCP 49. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp5` - Queue number of DSCP 5. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp50` - Queue number of DSCP 50. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp51` - Queue number of DSCP 51. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp52` - Queue number of DSCP 52. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp53` - Queue number of DSCP 53. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp54` - Queue number of DSCP 54. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp55` - Queue number of DSCP 55. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp56` - Queue number of DSCP 56. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp57` - Queue number of DSCP 57. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp58` - Queue number of DSCP 58. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp59` - Queue number of DSCP 59. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp6` - Queue number of DSCP 6. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp60` - Queue number of DSCP 60. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp61` - Queue number of DSCP 61. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp62` - Queue number of DSCP 62. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp63` - Queue number of DSCP 63. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp7` - Queue number of DSCP 7. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp8` - Queue number of DSCP 8. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `dscp9` - Queue number of DSCP 9. Valid values: `queue0` `queue1` `queue2` `queue3` `queue4` `queue5` `queue6` `queue7` .
* `id` - Profile ID.
* `type` - Profile type. Valid values: `cos` `dscp` .
* `weight` - Class weight.
* `scheduler` - Configure a NP7 QoS Scheduler. The structure of `scheduler` block is documented below.

The `scheduler` block contains:

* `mode` - Scheduler Mode. Valid values: `none` `priority` `round-robin` .
* `name` - Scheduler Name.
* `port_cpu_map` - Configure NPU interface to CPU core mapping. The structure of `port_cpu_map` block is documented below.

The `port_cpu_map` block contains:

* `cpu_core` - The CPU core to map to an interface.
* `interface` - The interface to map to a CPU core.
* `port_npu_map` - Configure port to NPU group mapping. The structure of `port_npu_map` block is documented below.

The `port_npu_map` block contains:

* `interface` - Set npu interface port to NPU group map.
* `npu_group_index` - Mapping NPU group index.
* `priority_protocol` - Configure NPU priority protocol. The structure of `priority_protocol` block is documented below.

The `priority_protocol` block contains:

* `bfd` - Enable/disable NPU BFD priority protocol. Valid values: `enable` `disable` .
* `bgp` - Enable/disable NPU BGP priority protocol. Valid values: `enable` `disable` .
* `slbc` - Enable/disable NPU SLBC priority protocol. Valid values: `enable` `disable` .
* `tcp_timeout_profile` - Configure TCP timeout profile. The structure of `tcp_timeout_profile` block is documented below.

The `tcp_timeout_profile` block contains:

* `close_wait` - Set close-wait timeout(seconds)
* `fin_wait` - Set fin-wait timeout(seconds)
* `id` - Timeout profile ID (5-47)
* `syn_sent` - Set syn-sent timeout(seconds)
* `syn_wait` - Set syn-wait timeout(seconds)
* `tcp_idle` - Set TCP establish timeout(seconds)
* `time_wait` - Set time-wait timeout(seconds)
* `udp_timeout_profile` - Configure UDP timeout profile. The structure of `udp_timeout_profile` block is documented below.

The `udp_timeout_profile` block contains:

* `id` - Timeout profile ID (5-63)
* `udp_idle` - Set UDP idle timeout(seconds)

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_npu can be imported using:
```sh
terraform import fortios_system_npu.labelname {{mkey}}
```
