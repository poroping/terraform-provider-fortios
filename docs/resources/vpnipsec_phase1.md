---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnipsec_phase1"
description: |-
  Configure VPN remote gateway.
---

## fortios_vpnipsec_phase1
Configure VPN remote gateway.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `acct_verify` - Enable/disable verification of RADIUS accounting record. Valid values: `enable` `disable` .
* `add_gw_route` - Enable/disable automatically add a route to the remote gateway. Valid values: `enable` `disable` .
* `add_route` - Enable/disable control addition of a route to peer destination selector. Valid values: `disable` `enable` .
* `assign_ip` - Enable/disable assignment of IP to IPsec interface via configuration method. Valid values: `disable` `enable` .
* `assign_ip_from` - Method by which the IP address will be assigned. Valid values: `range` `usrgrp` `dhcp` `name` .
* `authmethod` - Authentication method. Valid values: `psk` `signature` .
* `authmethod_remote` - Authentication method (remote side). Valid values: `psk` `signature` .
* `authpasswd` - XAuth password (max 35 characters).
* `authusr` - XAuth user name.
* `authusrgrp` - Authentication user group. This attribute must reference one of the following datasources: `user.group.name` .
* `auto_negotiate` - Enable/disable automatic initiation of IKE SA negotiation. Valid values: `enable` `disable` .
* `azure_ad_autoconnect` - Enable/disable Azure AD Auto-Connect for FortiClient. Valid values: `enable` `disable` .
* `banner` - Message that unity client should display after connecting.
* `cert_id_validation` - Enable/disable cross validation of peer ID and the identity in the peer's certificate as specified in RFC 4945. Valid values: `enable` `disable` .
* `childless_ike` - Enable/disable childless IKEv2 initiation (RFC 6023). Valid values: `enable` `disable` .
* `client_auto_negotiate` - Enable/disable allowing the VPN client to bring up the tunnel when there is no traffic. Valid values: `disable` `enable` .
* `client_keep_alive` - Enable/disable allowing the VPN client to keep the tunnel up when there is no traffic. Valid values: `disable` `enable` .
* `comments` - Comment.
* `dhcp_ra_giaddr` - Relay agent gateway IP address to use in the giaddr field of DHCP requests.
* `dhcp6_ra_linkaddr` - Relay agent IPv6 link address to use in DHCP6 requests.
* `dhgrp` - DH group. Valid values: `1` `2` `5` `14` `15` `16` `17` `18` `19` `20` `21` `27` `28` `29` `30` `31` `32` .
* `digital_signature_auth` - Enable/disable IKEv2 Digital Signature Authentication (RFC 7427). Valid values: `enable` `disable` .
* `distance` - Distance for routes added by IKE (1 - 255).
* `dns_mode` - DNS server mode. Valid values: `manual` `auto` .
* `domain` - Instruct unity clients about the single default DNS domain.
* `dpd` - Dead Peer Detection mode. Valid values: `disable` `on-idle` `on-demand` .
* `dpd_retrycount` - Number of DPD retry attempts.
* `dpd_retryinterval` - DPD retry interval.
* `eap` - Enable/disable IKEv2 EAP authentication. Valid values: `enable` `disable` .
* `eap_exclude_peergrp` - Peer group excluded from EAP authentication. This attribute must reference one of the following datasources: `user.peergrp.name` .
* `eap_identity` - IKEv2 EAP peer identity type. Valid values: `use-id-payload` `send-request` .
* `enforce_unique_id` - Enable/disable peer ID uniqueness check. Valid values: `disable` `keep-new` `keep-old` .
* `esn` - Extended sequence number (ESN) negotiation. Valid values: `require` `allow` `disable` .
* `fec_base` - Number of base Forward Error Correction packets (1 - 20).
* `fec_codec` - Forward Error Correction encoding/decoding algorithm. Valid values: `rs` `xor` .
* `fec_egress` - Enable/disable Forward Error Correction for egress IPsec traffic. Valid values: `enable` `disable` .
* `fec_health_check` - SD-WAN health check. This attribute must reference one of the following datasources: `system.sdwan.health-check.name` .
* `fec_ingress` - Enable/disable Forward Error Correction for ingress IPsec traffic. Valid values: `enable` `disable` .
* `fec_mapping_profile` - Forward Error Correction (FEC) mapping profile.
* `fec_receive_timeout` - Timeout in milliseconds before dropping Forward Error Correction packets (1 - 1000).
* `fec_redundant` - Number of redundant Forward Error Correction packets (1 - 5 for reed-solomon, 1 for xor).
* `fec_send_timeout` - Timeout in milliseconds before sending Forward Error Correction packets (1 - 1000).
* `fgsp_sync` - Enable/disable IPsec syncing of tunnels for FGSP IPsec. Valid values: `enable` `disable` .
* `forticlient_enforcement` - Enable/disable FortiClient enforcement. Valid values: `enable` `disable` .
* `fragmentation` - Enable/disable fragment IKE message on re-transmission. Valid values: `enable` `disable` .
* `fragmentation_mtu` - IKE fragmentation MTU (500 - 16000).
* `group_authentication` - Enable/disable IKEv2 IDi group authentication. Valid values: `enable` `disable` .
* `group_authentication_secret` - Password for IKEv2 ID group authentication. ASCII string or hexadecimal indicated by a leading 0x.
* `ha_sync_esp_seqno` - Enable/disable sequence number jump ahead for IPsec HA. Valid values: `enable` `disable` .
* `idle_timeout` - Enable/disable IPsec tunnel idle timeout. Valid values: `enable` `disable` .
* `idle_timeoutinterval` - IPsec tunnel idle timeout in minutes (5 - 43200).
* `ike_version` - IKE protocol version. Valid values: `1` `2` .
* `inbound_dscp_copy` - Enable/disable copy the dscp in the ESP header to the inner IP Header. Valid values: `enable` `disable` .
* `include_local_lan` - Enable/disable allow local LAN access on unity clients. Valid values: `disable` `enable` .
* `interface` - Local physical, aggregate, or VLAN outgoing interface. This attribute must reference one of the following datasources: `system.interface.name` .
* `ip_delay_interval` - IP address reuse delay interval in seconds (0 - 28800).
* `ipv4_dns_server1` - IPv4 DNS server 1.
* `ipv4_dns_server2` - IPv4 DNS server 2.
* `ipv4_dns_server3` - IPv4 DNS server 3.
* `ipv4_end_ip` - End of IPv4 range.
* `ipv4_name` - IPv4 address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `ipv4_netmask` - IPv4 Netmask.
* `ipv4_split_exclude` - IPv4 subnets that should not be sent over the IPsec tunnel. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `ipv4_split_include` - IPv4 split-include subnets. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `ipv4_start_ip` - Start of IPv4 range.
* `ipv4_wins_server1` - WINS server 1.
* `ipv4_wins_server2` - WINS server 2.
* `ipv6_dns_server1` - IPv6 DNS server 1.
* `ipv6_dns_server2` - IPv6 DNS server 2.
* `ipv6_dns_server3` - IPv6 DNS server 3.
* `ipv6_end_ip` - End of IPv6 range.
* `ipv6_name` - IPv6 address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `ipv6_prefix` - IPv6 prefix.
* `ipv6_split_exclude` - IPv6 subnets that should not be sent over the IPsec tunnel. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `ipv6_split_include` - IPv6 split-include subnets. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `ipv6_start_ip` - Start of IPv6 range.
* `keepalive` - NAT-T keep alive interval.
* `keylife` - Time to wait in seconds before phase 1 encryption key expires.
* `local_gw` - Local VPN gateway.
* `localid` - Local ID.
* `localid_type` - Local ID type. Valid values: `auto` `fqdn` `user-fqdn` `keyid` `address` `asn1dn` .
* `loopback_asymroute` - Enable/disable asymmetric routing for IKE traffic on loopback interface. Valid values: `enable` `disable` .
* `mesh_selector_type` - Add selectors containing subsets of the configuration depending on traffic. Valid values: `disable` `subnet` `host` .
* `mode` - ID protection mode used to establish a secure channel. Valid values: `aggressive` `main` .
* `mode_cfg` - Enable/disable configuration method. Valid values: `disable` `enable` .
* `mode_cfg_allow_client_selector` - Enable/disable mode-cfg client to use custom phase2 selectors. Valid values: `disable` `enable` .
* `name` - IPsec remote gateway name.
* `nattraversal` - Enable/disable NAT traversal. Valid values: `enable` `disable` `forced` .
* `negotiate_timeout` - IKE SA negotiation timeout in seconds (1 - 300).
* `network_id` - VPN gateway network ID.
* `network_overlay` - Enable/disable network overlays. Valid values: `disable` `enable` .
* `npu_offload` - Enable/disable offloading NPU. Valid values: `enable` `disable` .
* `peer` - Accept this peer certificate. This attribute must reference one of the following datasources: `user.peer.name` .
* `peergrp` - Accept this peer certificate group. This attribute must reference one of the following datasources: `user.peergrp.name` .
* `peerid` - Accept this peer identity.
* `peertype` - Accept this peer type. Valid values: `any` `one` `dialup` `peer` `peergrp` .
* `ppk` - Enable/disable IKEv2 Postquantum Preshared Key (PPK). Valid values: `disable` `allow` `require` .
* `ppk_identity` - IKEv2 Postquantum Preshared Key Identity.
* `ppk_secret` - IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
* `priority` - Priority for routes added by IKE (1 - 65535).
* `proposal` - Phase1 proposal. Valid values: `des-md5` `des-sha1` `des-sha256` `des-sha384` `des-sha512` `3des-md5` `3des-sha1` `3des-sha256` `3des-sha384` `3des-sha512` `aes128-md5` `aes128-sha1` `aes128-sha256` `aes128-sha384` `aes128-sha512` `aes128gcm-prfsha1` `aes128gcm-prfsha256` `aes128gcm-prfsha384` `aes128gcm-prfsha512` `aes192-md5` `aes192-sha1` `aes192-sha256` `aes192-sha384` `aes192-sha512` `aes256-md5` `aes256-sha1` `aes256-sha256` `aes256-sha384` `aes256-sha512` `aes256gcm-prfsha1` `aes256gcm-prfsha256` `aes256gcm-prfsha384` `aes256gcm-prfsha512` `chacha20poly1305-prfsha1` `chacha20poly1305-prfsha256` `chacha20poly1305-prfsha384` `chacha20poly1305-prfsha512` `aria128-md5` `aria128-sha1` `aria128-sha256` `aria128-sha384` `aria128-sha512` `aria192-md5` `aria192-sha1` `aria192-sha256` `aria192-sha384` `aria192-sha512` `aria256-md5` `aria256-sha1` `aria256-sha256` `aria256-sha384` `aria256-sha512` `seed-md5` `seed-sha1` `seed-sha256` `seed-sha384` `seed-sha512` .
* `psksecret` - Pre-shared secret for PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).
* `psksecret_remote` - Pre-shared secret for remote side PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).
* `reauth` - Enable/disable re-authentication upon IKE SA lifetime expiration. Valid values: `disable` `enable` .
* `rekey` - Enable/disable phase1 rekey. Valid values: `enable` `disable` .
* `remote_gw` - Remote VPN gateway.
* `remotegw_ddns` - Domain name of remote gateway. For example, name.ddns.com.
* `rsa_signature_format` - Digital Signature Authentication RSA signature format. Valid values: `pkcs1` `pss` .
* `rsa_signature_hash_override` - Enable/disable IKEv2 RSA signature hash algorithm override. Valid values: `enable` `disable` .
* `save_password` - Enable/disable saving XAuth username and password on VPN clients. Valid values: `disable` `enable` .
* `send_cert_chain` - Enable/disable sending certificate chain. Valid values: `enable` `disable` .
* `signature_hash_alg` - Digital Signature Authentication hash algorithms. Valid values: `sha1` `sha2-256` `sha2-384` `sha2-512` .
* `split_include_service` - Split-include services. This attribute must reference one of the following datasources: `firewall.service.group.name` `firewall.service.custom.name` .
* `suite_b` - Use Suite-B. Valid values: `disable` `suite-b-gcm-128` `suite-b-gcm-256` .
* `type` - Remote gateway type. Valid values: `static` `dynamic` `ddns` .
* `unity_support` - Enable/disable support for Cisco UNITY Configuration Method extensions. Valid values: `disable` `enable` .
* `usrgrp` - User group name for dialup peers. This attribute must reference one of the following datasources: `user.group.name` .
* `wizard_type` - GUI VPN Wizard Type. Valid values: `custom` `dialup-forticlient` `dialup-ios` `dialup-android` `dialup-windows` `dialup-cisco` `static-fortigate` `dialup-fortigate` `static-cisco` `dialup-cisco-fw` `simplified-static-fortigate` `hub-fortigate-auto-discovery` `spoke-fortigate-auto-discovery` .
* `xauthtype` - XAuth type. Valid values: `disable` `client` `pap` `chap` `auto` .
* `backup_gateway` - Instruct unity clients about the backup gateway address(es). The structure of `backup_gateway` block is documented below.

The `backup_gateway` block contains:

* `address` - Address of backup gateway.
* `certificate` - Names of up to 4 signed personal certificates. The structure of `certificate` block is documented below.

The `certificate` block contains:

* `name` - Certificate name. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `internal_domain_list` - List of domains for which the client directs DNS queries to the internal DNS servers for resolution.  DNS servers are configured in the mode-cfg settings.  One or more internal domain names in quotes separated by spaces, like "abc.com xyz.com 123.com" The structure of `internal_domain_list` block is documented below.

The `internal_domain_list` block contains:

* `domain_name` - Domain name.
* `ipv4_exclude_range` - Configuration Method IPv4 exclude ranges. The structure of `ipv4_exclude_range` block is documented below.

The `ipv4_exclude_range` block contains:

* `end_ip` - End of IPv4 exclusive range.
* `id` - ID.
* `start_ip` - Start of IPv4 exclusive range.
* `ipv6_exclude_range` - Configuration method IPv6 exclude ranges. The structure of `ipv6_exclude_range` block is documented below.

The `ipv6_exclude_range` block contains:

* `end_ip` - End of IPv6 exclusive range.
* `id` - ID.
* `start_ip` - Start of IPv6 exclusive range.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_vpnipsec_phase1 can be imported using:
```sh
terraform import fortios_vpnipsec_phase1.labelname {{mkey}}
```
