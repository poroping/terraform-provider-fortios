---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnipsec_phase1"
description: |-
  Get information on a fortios Configure VPN remote gateway.
---

# Data Source: fortios_vpnipsec_phase1
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
* `assign_ip` - Enable/disable assignment of IP to IPsec interface via configuration method.
* `assign_ip_from` - Method by which the IP address will be assigned.
* `authmethod` - Authentication method.
* `authmethod_remote` - Authentication method (remote side).
* `authpasswd` - XAuth password (max 35 characters).
* `authusr` - XAuth user name.
* `authusrgrp` - Authentication user group.
* `auto_negotiate` - Enable/disable automatic initiation of IKE SA negotiation.
* `banner` - Message that unity client should display after connecting.
* `cert_id_validation` - Enable/disable cross validation of peer ID and the identity in the peer's certificate as specified in RFC 4945.
* `childless_ike` - Enable/disable childless IKEv2 initiation (RFC 6023).
* `client_auto_negotiate` - Enable/disable allowing the VPN client to bring up the tunnel when there is no traffic.
* `client_keep_alive` - Enable/disable allowing the VPN client to keep the tunnel up when there is no traffic.
* `comments` - Comment.
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
* `enforce_unique_id` - Enable/disable peer ID uniqueness check.
* `esn` - Extended sequence number (ESN) negotiation.
* `fec_base` - Number of base Forward Error Correction packets (1 - 20).
* `fec_codec` - Forward Error Correction encoding/decoding algorithm.
* `fec_egress` - Enable/disable Forward Error Correction for egress IPsec traffic.
* `fec_health_check` - SD-WAN health check.
* `fec_ingress` - Enable/disable Forward Error Correction for ingress IPsec traffic.
* `fec_mapping_profile` - Forward Error Correction (FEC) mapping profile.
* `fec_receive_timeout` - Timeout in milliseconds before dropping Forward Error Correction packets (1 - 1000).
* `fec_redundant` - Number of redundant Forward Error Correction packets (1 - 5 for reed-solomon, 1 for xor).
* `fec_send_timeout` - Timeout in milliseconds before sending Forward Error Correction packets (1 - 1000).
* `forticlient_enforcement` - Enable/disable FortiClient enforcement.
* `fragmentation` - Enable/disable fragment IKE message on re-transmission.
* `fragmentation_mtu` - IKE fragmentation MTU (500 - 16000).
* `group_authentication` - Enable/disable IKEv2 IDi group authentication.
* `group_authentication_secret` - Password for IKEv2 IDi group authentication.  (ASCII string or hexadecimal indicated by a leading 0x.)
* `ha_sync_esp_seqno` - Enable/disable sequence number jump ahead for IPsec HA.
* `idle_timeout` - Enable/disable IPsec tunnel idle timeout.
* `idle_timeoutinterval` - IPsec tunnel idle timeout in minutes (5 - 43200).
* `ike_version` - IKE protocol version.
* `include_local_lan` - Enable/disable allow local LAN access on unity clients.
* `interface` - Local physical, aggregate, or VLAN outgoing interface.
* `ip_delay_interval` - IP address reuse delay interval in seconds (0 - 28800).
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
* `local_gw` - Local VPN gateway.
* `localid` - Local ID.
* `localid_type` - Local ID type.
* `loopback_asymroute` - Enable/disable asymmetric routing for IKE traffic on loopback interface.
* `mesh_selector_type` - Add selectors containing subsets of the configuration depending on traffic.
* `mode` - ID protection mode used to establish a secure channel.
* `mode_cfg` - Enable/disable configuration method.
* `name` - IPsec remote gateway name.
* `nattraversal` - Enable/disable NAT traversal.
* `negotiate_timeout` - IKE SA negotiation timeout in seconds (1 - 300).
* `network_id` - VPN gateway network ID.
* `network_overlay` - Enable/disable network overlays.
* `npu_offload` - Enable/disable offloading NPU.
* `peer` - Accept this peer certificate.
* `peergrp` - Accept this peer certificate group.
* `peerid` - Accept this peer identity.
* `peertype` - Accept this peer type.
* `ppk` - Enable/disable IKEv2 Postquantum Preshared Key (PPK).
* `ppk_identity` - IKEv2 Postquantum Preshared Key Identity.
* `ppk_secret` - IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
* `priority` - Priority for routes added by IKE (0 - 4294967295).
* `proposal` - Phase1 proposal.
* `psksecret` - Pre-shared secret for PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).
* `psksecret_remote` - Pre-shared secret for remote side PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).
* `reauth` - Enable/disable re-authentication upon IKE SA lifetime expiration.
* `rekey` - Enable/disable phase1 rekey.
* `remote_gw` - Remote VPN gateway.
* `remotegw_ddns` - Domain name of remote gateway (eg. name.DDNS.com).
* `rsa_signature_format` - Digital Signature Authentication RSA signature format.
* `save_password` - Enable/disable saving XAuth username and password on VPN clients.
* `send_cert_chain` - Enable/disable sending certificate chain.
* `signature_hash_alg` - Digital Signature Authentication hash algorithms.
* `split_include_service` - Split-include services.
* `suite_b` - Use Suite-B.
* `type` - Remote gateway type.
* `unity_support` - Enable/disable support for Cisco UNITY Configuration Method extensions.
* `usrgrp` - User group name for dialup peers.
* `wizard_type` - GUI VPN Wizard Type.
* `xauthtype` - XAuth type.
* `backup_gateway` - Instruct unity clients about the backup gateway address(es).The structure of `backup_gateway` block is documented below.

The `backup_gateway` block contains:

* `address` - Address of backup gateway.
* `certificate` - Names of up to 4 signed personal certificates.The structure of `certificate` block is documented below.

The `certificate` block contains:

* `name` - Certificate name.
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
