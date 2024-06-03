---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnipsec_phase1interface"
description: |-
  Get information on a fortios Configure VPN remote gateway.
---

# Data Source: fortios_vpnipsec_phase1interface
Use this data source to get information on a fortios Configure VPN remote gateway.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) IPsec remote gateway name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `acct_verify` - Enable/disable verification of RADIUS accounting record.
* `add_gw_route` - Enable/disable automatically add a route to the remote gateway.
* `add_route` - Enable/disable control addition of a route to peer destination selector.
* `aggregate_member` - Enable/disable use as an aggregate member.
* `aggregate_weight` - Link weight for aggregate.
* `assign_ip` - Enable/disable assignment of IP to IPsec interface via configuration method.
* `assign_ip_from` - Method by which the IP address will be assigned.
* `authmethod` - Authentication method.
* `authmethod_remote` - Authentication method (remote side).
* `authpasswd` - XAuth password (max 35 characters).
* `authusr` - XAuth user name.
* `authusrgrp` - Authentication user group.
* `auto_discovery_crossover` - Allow/block set-up of short-cut tunnels between different network IDs.
* `auto_discovery_forwarder` - Enable/disable forwarding auto-discovery short-cut messages.
* `auto_discovery_offer_interval` - Interval between shortcut offer messages in seconds (1 - 300, default = 5).
* `auto_discovery_psk` - Enable/disable use of pre-shared secrets for authentication of auto-discovery tunnels.
* `auto_discovery_receiver` - Enable/disable accepting auto-discovery short-cut messages.
* `auto_discovery_sender` - Enable/disable sending auto-discovery short-cut messages.
* `auto_discovery_shortcuts` - Control deletion of child short-cut tunnels when the parent tunnel goes down.
* `auto_negotiate` - Enable/disable automatic initiation of IKE SA negotiation.
* `azure_ad_autoconnect` - Enable/disable Azure AD Auto-Connect for FortiClient.
* `banner` - Message that unity client should display after connecting.
* `cert_id_validation` - Enable/disable cross validation of peer ID and the identity in the peer's certificate as specified in RFC 4945.
* `cert_trust_store` - CA certificate trust store.
* `childless_ike` - Enable/disable childless IKEv2 initiation (RFC 6023).
* `client_auto_negotiate` - Enable/disable allowing the VPN client to bring up the tunnel when there is no traffic.
* `client_keep_alive` - Enable/disable allowing the VPN client to keep the tunnel up when there is no traffic.
* `comments` - Comment.
* `default_gw` - IPv4 address of default route gateway to use for traffic exiting the interface.
* `default_gw_priority` - Priority for default gateway route. A higher priority number signifies a less preferred route.
* `dev_id` - Device ID carried by the device ID notification.
* `dev_id_notification` - Enable/disable device ID notification.
* `dhcp_ra_giaddr` - Relay agent gateway IP address to use in the giaddr field of DHCP requests.
* `dhcp6_ra_linkaddr` - Relay agent IPv6 link address to use in DHCP6 requests.
* `dhgrp` - DH group.
* `digital_signature_auth` - Enable/disable IKEv2 Digital Signature Authentication (RFC 7427).
* `distance` - Distance for routes added by IKE (1 - 255).
* `dns_mode` - DNS server mode.
* `domain` - Instruct unity clients about the single default DNS domain.
* `dpd` - Dead Peer Detection mode.
* `dpd_retrycount` - Number of DPD retry attempts.
* `dpd_retryinterval` - DPD retry interval.
* `eap` - Enable/disable IKEv2 EAP authentication.
* `eap_exclude_peergrp` - Peer group excluded from EAP authentication.
* `eap_identity` - IKEv2 EAP peer identity type.
* `ems_sn_check` - Enable/disable verification of EMS serial number.
* `encap_local_gw4` - Local IPv4 address of GRE/VXLAN tunnel.
* `encap_local_gw6` - Local IPv6 address of GRE/VXLAN tunnel.
* `encap_remote_gw4` - Remote IPv4 address of GRE/VXLAN tunnel.
* `encap_remote_gw6` - Remote IPv6 address of GRE/VXLAN tunnel.
* `encapsulation` - Enable/disable GRE/VXLAN/VPNID encapsulation.
* `encapsulation_address` - Source for GRE/VXLAN tunnel address.
* `enforce_unique_id` - Enable/disable peer ID uniqueness check.
* `esn` - Extended sequence number (ESN) negotiation.
* `exchange_fgt_device_id` - Enable/disable device identifier exchange with peer FortiGate units for use of VPN monitor data by FortiManager.
* `exchange_interface_ip` - Enable/disable exchange of IPsec interface IP address.
* `exchange_ip_addr4` - IPv4 address to exchange with peers.
* `exchange_ip_addr6` - IPv6 address to exchange with peers.
* `fec_base` - Number of base Forward Error Correction packets (1 - 20).
* `fec_codec` - Forward Error Correction encoding/decoding algorithm.
* `fec_egress` - Enable/disable Forward Error Correction for egress IPsec traffic.
* `fec_health_check` - SD-WAN health check.
* `fec_ingress` - Enable/disable Forward Error Correction for ingress IPsec traffic.
* `fec_mapping_profile` - Forward Error Correction (FEC) mapping profile.
* `fec_receive_timeout` - Timeout in milliseconds before dropping Forward Error Correction packets (1 - 1000).
* `fec_redundant` - Number of redundant Forward Error Correction packets (1 - 5 for reed-solomon, 1 for xor).
* `fec_send_timeout` - Timeout in milliseconds before sending Forward Error Correction packets (1 - 1000).
* `fgsp_sync` - Enable/disable IPsec syncing of tunnels for FGSP IPsec.
* `forticlient_enforcement` - Enable/disable FortiClient enforcement.
* `fragmentation` - Enable/disable fragment IKE message on re-transmission.
* `fragmentation_mtu` - IKE fragmentation MTU (500 - 16000).
* `group_authentication` - Enable/disable IKEv2 IDi group authentication.
* `group_authentication_secret` - Password for IKEv2 ID group authentication. ASCII string or hexadecimal indicated by a leading 0x.
* `ha_sync_esp_seqno` - Enable/disable sequence number jump ahead for IPsec HA.
* `idle_timeout` - Enable/disable IPsec tunnel idle timeout.
* `idle_timeoutinterval` - IPsec tunnel idle timeout in minutes (5 - 43200).
* `ike_version` - IKE protocol version.
* `inbound_dscp_copy` - Enable/disable copy the dscp in the ESP header to the inner IP Header.
* `include_local_lan` - Enable/disable allow local LAN access on unity clients.
* `interface` - Local physical, aggregate, or VLAN outgoing interface.
* `ip_delay_interval` - IP address reuse delay interval in seconds (0 - 28800).
* `ip_fragmentation` - Determine whether IP packets are fragmented before or after IPsec encapsulation.
* `ip_version` - IP version to use for VPN interface.
* `ipv4_dns_server1` - IPv4 DNS server 1.
* `ipv4_dns_server2` - IPv4 DNS server 2.
* `ipv4_dns_server3` - IPv4 DNS server 3.
* `ipv4_end_ip` - End of IPv4 range.
* `ipv4_name` - IPv4 address name.
* `ipv4_netmask` - IPv4 Netmask.
* `ipv4_split_exclude` - IPv4 subnets that should not be sent over the IPsec tunnel.
* `ipv4_split_include` - IPv4 split-include subnets.
* `ipv4_start_ip` - Start of IPv4 range.
* `ipv4_wins_server1` - WINS server 1.
* `ipv4_wins_server2` - WINS server 2.
* `ipv6_dns_server1` - IPv6 DNS server 1.
* `ipv6_dns_server2` - IPv6 DNS server 2.
* `ipv6_dns_server3` - IPv6 DNS server 3.
* `ipv6_end_ip` - End of IPv6 range.
* `ipv6_name` - IPv6 address name.
* `ipv6_prefix` - IPv6 prefix.
* `ipv6_split_exclude` - IPv6 subnets that should not be sent over the IPsec tunnel.
* `ipv6_split_include` - IPv6 split-include subnets.
* `ipv6_start_ip` - Start of IPv6 range.
* `keepalive` - NAT-T keep alive interval.
* `keylife` - Time to wait in seconds before phase 1 encryption key expires.
* `link_cost` - VPN tunnel underlay link cost.
* `local_gw` - IPv4 address of the local gateway's external interface.
* `local_gw6` - IPv6 address of the local gateway's external interface.
* `localid` - Local ID.
* `localid_type` - Local ID type.
* `loopback_asymroute` - Enable/disable asymmetric routing for IKE traffic on loopback interface.
* `mesh_selector_type` - Add selectors containing subsets of the configuration depending on traffic.
* `mode` - The ID protection mode used to establish a secure channel.
* `mode_cfg` - Enable/disable configuration method.
* `mode_cfg_allow_client_selector` - Enable/disable mode-cfg client to use custom phase2 selectors.
* `monitor` - IPsec interface as backup for primary interface.
* `monitor_hold_down_delay` - Time to wait in seconds before recovery once primary re-establishes.
* `monitor_hold_down_time` - Time of day at which to fail back to primary after it re-establishes.
* `monitor_hold_down_type` - Recovery time method when primary interface re-establishes.
* `monitor_hold_down_weekday` - Day of the week to recover once primary re-establishes.
* `name` - IPsec remote gateway name.
* `nattraversal` - Enable/disable NAT traversal.
* `negotiate_timeout` - IKE SA negotiation timeout in seconds (1 - 300).
* `net_device` - Enable/disable kernel device creation.
* `network_id` - VPN gateway network ID.
* `network_overlay` - Enable/disable network overlays.
* `npu_offload` - Enable/disable offloading NPU.
* `packet_redistribution` - Enable/disable packet distribution (RPS) on the IPsec interface.
* `passive_mode` - Enable/disable IPsec passive mode for static tunnels.
* `peer` - Accept this peer certificate.
* `peergrp` - Accept this peer certificate group.
* `peerid` - Accept this peer identity.
* `peertype` - Accept this peer type.
* `ppk` - Enable/disable IKEv2 Postquantum Preshared Key (PPK).
* `ppk_identity` - IKEv2 Postquantum Preshared Key Identity.
* `ppk_secret` - IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
* `priority` - Priority for routes added by IKE (1 - 65535).
* `proposal` - Phase1 proposal.
* `psksecret` - Pre-shared secret for PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).
* `psksecret_remote` - Pre-shared secret for remote side PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).
* `reauth` - Enable/disable re-authentication upon IKE SA lifetime expiration.
* `rekey` - Enable/disable phase1 rekey.
* `remote_gw` - IPv4 address of the remote gateway's external interface.
* `remote_gw_country` - IPv4 addresses associated to a specific country.
* `remote_gw_end_ip` - Last IPv4 address in the range.
* `remote_gw_match` - Set type of IPv4 remote gateway address matching.
* `remote_gw_start_ip` - First IPv4 address in the range.
* `remote_gw_subnet` - IPv4 address and subnet mask.
* `remote_gw6` - IPv6 address of the remote gateway's external interface.
* `remote_gw6_country` - IPv6 addresses associated to a specific country.
* `remote_gw6_end_ip` - Last IPv6 address in the range.
* `remote_gw6_match` - Set type of IPv6 remote gateway address matching.
* `remote_gw6_start_ip` - First IPv6 address in the range.
* `remote_gw6_subnet` - IPv6 address and prefix.
* `remotegw_ddns` - Domain name of remote gateway. For example, name.ddns.com.
* `rsa_signature_format` - Digital Signature Authentication RSA signature format.
* `rsa_signature_hash_override` - Enable/disable IKEv2 RSA signature hash algorithm override.
* `save_password` - Enable/disable saving XAuth username and password on VPN clients.
* `send_cert_chain` - Enable/disable sending certificate chain.
* `signature_hash_alg` - Digital Signature Authentication hash algorithms.
* `split_include_service` - Split-include services.
* `suite_b` - Use Suite-B.
* `tunnel_search` - Tunnel search method for when the interface is shared.
* `type` - Remote gateway type.
* `unity_support` - Enable/disable support for Cisco UNITY Configuration Method extensions.
* `usrgrp` - User group name for dialup peers.
* `vni` - VNI of VXLAN tunnel.
* `wizard_type` - GUI VPN Wizard Type.
* `xauthtype` - XAuth type.
* `backup_gateway` - Instruct unity clients about the backup gateway address(es).The structure of `backup_gateway` block is documented below.

The `backup_gateway` block contains:

* `address` - Address of backup gateway.
* `certificate` - The names of up to 4 signed personal certificates.The structure of `certificate` block is documented below.

The `certificate` block contains:

* `name` - Certificate name.
* `internal_domain_list` - List of domains for which the client directs DNS queries to the internal DNS servers for resolution.  DNS servers are configured in the mode-cfg settings.  One or more internal domain names in quotes separated by spaces, like "abc.com xyz.com 123.com"The structure of `internal_domain_list` block is documented below.

The `internal_domain_list` block contains:

* `domain_name` - Domain name.
* `ipv4_exclude_range` - Configuration Method IPv4 exclude ranges.The structure of `ipv4_exclude_range` block is documented below.

The `ipv4_exclude_range` block contains:

* `end_ip` - End of IPv4 exclusive range.
* `id` - ID.
* `start_ip` - Start of IPv4 exclusive range.
* `ipv6_exclude_range` - Configuration method IPv6 exclude ranges.The structure of `ipv6_exclude_range` block is documented below.

The `ipv6_exclude_range` block contains:

* `end_ip` - End of IPv6 exclusive range.
* `id` - ID.
* `start_ip` - Start of IPv6 exclusive range.
