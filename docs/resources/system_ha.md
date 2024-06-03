---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ha"
description: |-
  Configure HA.
---

## fortios_system_ha
Configure HA.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `arps` - Number of gratuitous ARPs (1 - 60). Lower to reduce traffic. Higher to reduce failover time.
* `arps_interval` - Time between gratuitous ARPs  (1 - 20 sec). Lower to reduce failover time. Higher to reduce traffic.
* `authentication` - Enable/disable heartbeat message authentication. Valid values: `enable` `disable` .
* `cpu_threshold` - Dynamic weighted load balancing CPU usage weight and high and low thresholds.
* `encryption` - Enable/disable heartbeat message encryption. Valid values: `enable` `disable` .
* `failover_hold_time` - Time to wait before failover (0 - 300 sec, default = 0), to avoid flip.
* `ftp_proxy_threshold` - Dynamic weighted load balancing weight and high and low number of FTP proxy sessions.
* `gratuitous_arps` - Enable/disable gratuitous ARPs. Disable if link-failed-signal enabled. Valid values: `enable` `disable` .
* `group_id` - HA group ID  (0 - 1023;  or 0 - 7 when there are more than 2 vclusters). Must be the same for all members.
* `group_name` - Cluster group name. Must be the same for all members.
* `ha_direct` - Enable/disable using ha-mgmt interface for syslog, remote authentication (RADIUS), FortiAnalyzer, FortiSandbox, sFlow, and Netflow. Valid values: `enable` `disable` .
* `ha_eth_type` - HA heartbeat packet Ethertype (4-digit hex).
* `ha_mgmt_status` - Enable to reserve interfaces to manage individual cluster units. Valid values: `enable` `disable` .
* `ha_uptime_diff_margin` - Normally you would only reduce this value for failover testing.
* `hb_interval` - Time between sending heartbeat packets (1 - 20). Increase to reduce false positives.
* `hb_interval_in_milliseconds` - Number of milliseconds for each heartbeat interval: 100ms or 10ms. Valid values: `100ms` `10ms` .
* `hb_lost_threshold` - Number of lost heartbeats to signal a failure (1 - 60). Increase to reduce false positives.
* `hbdev` - Heartbeat interfaces. Must be the same for all members.
* `hc_eth_type` - Transparent mode HA heartbeat packet Ethertype (4-digit hex).
* `hello_holddown` - Time to wait before changing from hello to work state (5 - 300 sec).
* `http_proxy_threshold` - Dynamic weighted load balancing weight and high and low number of HTTP proxy sessions.
* `imap_proxy_threshold` - Dynamic weighted load balancing weight and high and low number of IMAP proxy sessions.
* `inter_cluster_session_sync` - Enable/disable synchronization of sessions among HA clusters. Valid values: `enable` `disable` .
* `key` - Key.
* `l2ep_eth_type` - Telnet session HA heartbeat packet Ethertype (4-digit hex).
* `link_failed_signal` - Enable to shut down all interfaces for 1 sec after a failover. Use if gratuitous ARPs do not update network. Valid values: `enable` `disable` .
* `load_balance_all` - Enable to load balance TCP sessions. Disable to load balance proxy sessions only. Valid values: `enable` `disable` .
* `logical_sn` - Enable/disable usage of the logical serial number. Valid values: `enable` `disable` .
* `memory_based_failover` - Enable/disable memory based failover. Valid values: `enable` `disable` .
* `memory_compatible_mode` - Enable/disable memory compatible mode. Valid values: `enable` `disable` .
* `memory_failover_flip_timeout` - Time to wait between subsequent memory based failovers in minutes (6 - 2147483647, default = 6).
* `memory_failover_monitor_period` - Duration of high memory usage before memory based failover is triggered in seconds (1 - 300, default = 60).
* `memory_failover_sample_rate` - Rate at which memory usage is sampled in order to measure memory usage in seconds (1 - 60, default = 1).
* `memory_failover_threshold` - Memory usage threshold to trigger memory based failover (0 means using conserve mode threshold in system.global).
* `memory_threshold` - Dynamic weighted load balancing memory usage weight and high and low thresholds.
* `mode` - HA mode. Must be the same for all members. FGSP requires standalone. Valid values: `standalone` `a-a` `a-p` .
* `monitor` - Interfaces to check for port monitoring (or link failure). This attribute must reference one of the following datasources: `system.interface.name` .
* `multicast_ttl` - HA multicast TTL on primary (5 - 3600 sec).
* `nntp_proxy_threshold` - Dynamic weighted load balancing weight and high and low number of NNTP proxy sessions.
* `override` - Enable and increase the priority of the unit that should always be primary (master). Valid values: `enable` `disable` .
* `override_wait_time` - Delay negotiating if override is enabled (0 - 3600 sec). Reduces how often the cluster negotiates.
* `password` - Cluster password. Must be the same for all members.
* `pingserver_failover_threshold` - Remote IP monitoring failover threshold (0 - 50).
* `pingserver_flip_timeout` - Time to wait in minutes before renegotiating after a remote IP monitoring failover.
* `pingserver_monitor_interface` - Interfaces to check for remote IP monitoring. This attribute must reference one of the following datasources: `system.interface.name` .
* `pingserver_secondary_force_reset` - Enable to force the cluster to negotiate after a remote IP monitoring failover. Valid values: `enable` `disable` .
* `pingserver_slave_force_reset` - Enable to force the cluster to negotiate after a remote IP monitoring failover. Valid values: `enable` `disable` .
* `pop3_proxy_threshold` - Dynamic weighted load balancing weight and high and low number of POP3 proxy sessions.
* `priority` - Increase the priority to select the primary unit (0 - 255).
* `route_hold` - Time to wait between routing table updates to the cluster (0 - 3600 sec).
* `route_ttl` - TTL for primary unit routes (5 - 3600 sec). Increase to maintain active routes during failover.
* `route_wait` - Time to wait before sending new routes to the cluster (0 - 3600 sec).
* `schedule` - Type of A-A load balancing. Use none if you have external load balancers. Valid values: `none` `leastconnection` `round-robin` `weight-round-robin` `random` `ip` `ipport` .
* `session_pickup` - Enable/disable session pickup. Enabling it can reduce session down time when fail over happens. Valid values: `enable` `disable` .
* `session_pickup_connectionless` - Enable/disable UDP and ICMP session sync. Valid values: `enable` `disable` .
* `session_pickup_delay` - Enable to sync sessions longer than 30 sec. Only longer lived sessions need to be synced. Valid values: `enable` `disable` .
* `session_pickup_expectation` - Enable/disable session helper expectation session sync for FGSP. Valid values: `enable` `disable` .
* `session_pickup_nat` - Enable/disable NAT session sync for FGSP. Valid values: `enable` `disable` .
* `session_sync_dev` - Offload session-sync process to kernel and sync sessions using connected interface(s) directly. This attribute must reference one of the following datasources: `system.interface.name` .
* `smtp_proxy_threshold` - Dynamic weighted load balancing weight and high and low number of SMTP proxy sessions.
* `ssd_failover` - Enable/disable automatic HA failover on SSD disk failure. Valid values: `enable` `disable` .
* `standalone_config_sync` - Enable/disable FGSP configuration synchronization. Valid values: `enable` `disable` .
* `standalone_mgmt_vdom` - Enable/disable standalone management VDOM. Valid values: `enable` `disable` .
* `sync_config` - Enable/disable configuration synchronization. Valid values: `enable` `disable` .
* `sync_packet_balance` - Enable/disable HA packet distribution to multiple CPUs. Valid values: `enable` `disable` .
* `unicast_gateway` - Default route gateway for unicast interface.
* `unicast_hb` - Enable/disable unicast heartbeat. Valid values: `enable` `disable` .
* `unicast_hb_netmask` - Unicast heartbeat netmask.
* `unicast_hb_peerip` - Unicast heartbeat peer IP.
* `unicast_status` - Enable/disable unicast connection. Valid values: `enable` `disable` .
* `uninterruptible_primary_wait` - Number of minutes the primary HA unit waits before the secondary HA unit is considered upgraded and the system is started before starting its own upgrade (15 - 300, default = 30).
* `uninterruptible_upgrade` - Enable to upgrade a cluster without blocking network traffic. Valid values: `enable` `disable` .
* `vcluster_id` - Cluster ID.
* `vcluster_status` - Enable/disable virtual cluster for virtual clustering. Valid values: `enable` `disable` .
* `vcluster2` - Enable/disable virtual cluster 2 for virtual clustering. Valid values: `enable` `disable` .
* `vdom` - VDOMs in virtual cluster 1.
* `weight` - Weight-round-robin weight for each cluster unit. Syntax <priority> <weight>.
* `ha_mgmt_interfaces` - Reserve interfaces to manage individual cluster units. The structure of `ha_mgmt_interfaces` block is documented below.

The `ha_mgmt_interfaces` block contains:

* `dst` - Default route destination for reserved HA management interface.
* `gateway` - Default route gateway for reserved HA management interface.
* `gateway6` - Default IPv6 gateway for reserved HA management interface.
* `id` - Table ID.
* `interface` - Interface to reserve for HA management. This attribute must reference one of the following datasources: `system.interface.name` .
* `secondary_vcluster` - Configure virtual cluster 2. The structure of `secondary_vcluster` block is documented below.

The `secondary_vcluster` block contains:

* `monitor` - Interfaces to check for port monitoring (or link failure). This attribute must reference one of the following datasources: `system.interface.name` .
* `override` - Enable and increase the priority of the unit that should always be primary. Valid values: `enable` `disable` .
* `override_wait_time` - Delay negotiating if override is enabled (0 - 3600 sec). Reduces how often the cluster negotiates.
* `pingserver_failover_threshold` - Remote IP monitoring failover threshold (0 - 50).
* `pingserver_monitor_interface` - Interfaces to check for remote IP monitoring. This attribute must reference one of the following datasources: `system.interface.name` .
* `pingserver_secondary_force_reset` - Enable to force the cluster to negotiate after a remote IP monitoring failover. Valid values: `enable` `disable` .
* `pingserver_slave_force_reset` - Enable to force the cluster to negotiate after a remote IP monitoring failover. Valid values: `enable` `disable` .
* `priority` - Increase the priority to select the primary unit (0 - 255).
* `vcluster_id` - Cluster ID.
* `vdom` - VDOMs in virtual cluster 2.
* `unicast_peers` - Number of unicast peers. The structure of `unicast_peers` block is documented below.

The `unicast_peers` block contains:

* `id` - Table ID.
* `peer_ip` - Unicast peer IP.
* `vcluster` - Virtual cluster table. The structure of `vcluster` block is documented below.

The `vcluster` block contains:

* `monitor` - Interfaces to check for port monitoring (or link failure). This attribute must reference one of the following datasources: `system.interface.name` .
* `override` - Enable and increase the priority of the unit that should always be primary (master). Valid values: `enable` `disable` .
* `override_wait_time` - Delay negotiating if override is enabled (0 - 3600 sec). Reduces how often the cluster negotiates.
* `pingserver_failover_threshold` - Remote IP monitoring failover threshold (0 - 50).
* `pingserver_monitor_interface` - Interfaces to check for remote IP monitoring. This attribute must reference one of the following datasources: `system.interface.name` .
* `pingserver_secondary_force_reset` - Enable to force the cluster to negotiate after a remote IP monitoring failover. Valid values: `enable` `disable` .
* `pingserver_slave_force_reset` - Enable to force the cluster to negotiate after a remote IP monitoring failover. Valid values: `enable` `disable` .
* `priority` - Increase the priority to select the primary unit (0 - 255).
* `vcluster_id` - ID.
* `vdom` - Virtual domain(s) in the virtual cluster. The structure of `vdom` block is documented below.

The `vdom` block contains:

* `name` - Virtual domain name. This attribute must reference one of the following datasources: `system.vdom.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_ha can be imported using:
```sh
terraform import fortios_system_ha.labelname {{mkey}}
```
