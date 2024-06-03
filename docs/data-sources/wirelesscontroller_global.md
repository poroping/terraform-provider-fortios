---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_global"
description: |-
  Get information on a fortios Configure wireless controller global settings.
---

# Data Source: fortios_wirelesscontroller_global
Use this data source to get information on a fortios Configure wireless controller global settings.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `acd_process_count` - Configure the number cw_acd daemons for multi-core CPU support (default = 0).
* `ap_log_server` - Enable/disable configuring FortiGate to redirect wireless event log messages or FortiAPs to send UTM log messages to a syslog server (default = disable).
* `ap_log_server_ip` - IP address that FortiGate or FortiAPs send log messages to.
* `ap_log_server_port` - Port that FortiGate or FortiAPs send log messages to.
* `control_message_offload` - Configure CAPWAP control message data channel offload.
* `data_ethernet_ii` - Configure the wireless controller to use Ethernet II or 802.3 frames with 802.3 data tunnel mode (default = enable).
* `discovery_mc_addr` - Multicast IP address for AP discovery (default = 244.0.1.140).
* `fiapp_eth_type` - Ethernet type for Fortinet Inter-Access Point Protocol (IAPP), or IEEE 802.11f, packets (0 - 65535, default = 5252).
* `image_download` - Enable/disable WTP image download at join time.
* `ipsec_base_ip` - Base IP address for IPsec VPN tunnels between the access points and the wireless controller (default = 169.254.0.1).
* `link_aggregation` - Enable/disable calculating the CAPWAP transmit hash to load balance sessions to link aggregation nodes (default = disable).
* `location` - Description of the location of the wireless controller.
* `max_clients` - Maximum number of clients that can connect simultaneously (default = 0, meaning no limitation).
* `max_retransmit` - Maximum number of tunnel packet retransmissions (0 - 64, default = 3).
* `mesh_eth_type` - Mesh Ethernet identifier included in backhaul packets (0 - 65535, default = 8755).
* `nac_interval` - Interval in seconds between two WiFi network access control (NAC) checks (10 - 600, default = 120).
* `name` - Name of the wireless controller.
* `rogue_scan_mac_adjacency` - Maximum numerical difference between an AP's Ethernet and wireless MAC values to match for rogue detection (0 - 31, default = 7).
* `tunnel_mode` - Compatible/strict tunnel mode.
* `wtp_share` - Enable/disable sharing of WTPs between VDOMs.
