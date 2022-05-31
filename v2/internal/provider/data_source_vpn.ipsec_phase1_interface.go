// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure VPN remote gateway.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceVpnIpsecPhase1Interface() *schema.Resource {
	return &schema.Resource{
		Description: "Configure VPN remote gateway.",

		ReadContext: dataSourceVpnIpsecPhase1InterfaceRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"acct_verify": {
				Type:        schema.TypeString,
				Description: "Enable/disable verification of RADIUS accounting record.",
				Computed:    true,
			},
			"add_gw_route": {
				Type:        schema.TypeString,
				Description: "Enable/disable automatically add a route to the remote gateway.",
				Computed:    true,
			},
			"add_route": {
				Type:        schema.TypeString,
				Description: "Enable/disable control addition of a route to peer destination selector.",
				Computed:    true,
			},
			"aggregate_member": {
				Type:        schema.TypeString,
				Description: "Enable/disable use as an aggregate member.",
				Computed:    true,
			},
			"aggregate_weight": {
				Type:        schema.TypeInt,
				Description: "Link weight for aggregate.",
				Computed:    true,
			},
			"assign_ip": {
				Type:        schema.TypeString,
				Description: "Enable/disable assignment of IP to IPsec interface via configuration method.",
				Computed:    true,
			},
			"assign_ip_from": {
				Type:        schema.TypeString,
				Description: "Method by which the IP address will be assigned.",
				Computed:    true,
			},
			"authmethod": {
				Type:        schema.TypeString,
				Description: "Authentication method.",
				Computed:    true,
			},
			"authmethod_remote": {
				Type:        schema.TypeString,
				Description: "Authentication method (remote side).",
				Computed:    true,
			},
			"authpasswd": {
				Type:        schema.TypeString,
				Description: "XAuth password (max 35 characters).",
				Computed:    true,
				Sensitive:   true,
			},
			"authusr": {
				Type:        schema.TypeString,
				Description: "XAuth user name.",
				Computed:    true,
			},
			"authusrgrp": {
				Type:        schema.TypeString,
				Description: "Authentication user group.",
				Computed:    true,
			},
			"auto_discovery_forwarder": {
				Type:        schema.TypeString,
				Description: "Enable/disable forwarding auto-discovery short-cut messages.",
				Computed:    true,
			},
			"auto_discovery_offer_interval": {
				Type:        schema.TypeInt,
				Description: "Interval between shortcut offer messages in seconds (1 - 300, default = 5).",
				Computed:    true,
			},
			"auto_discovery_psk": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of pre-shared secrets for authentication of auto-discovery tunnels.",
				Computed:    true,
			},
			"auto_discovery_receiver": {
				Type:        schema.TypeString,
				Description: "Enable/disable accepting auto-discovery short-cut messages.",
				Computed:    true,
			},
			"auto_discovery_sender": {
				Type:        schema.TypeString,
				Description: "Enable/disable sending auto-discovery short-cut messages.",
				Computed:    true,
			},
			"auto_discovery_shortcuts": {
				Type:        schema.TypeString,
				Description: "Control deletion of child short-cut tunnels when the parent tunnel goes down.",
				Computed:    true,
			},
			"auto_negotiate": {
				Type:        schema.TypeString,
				Description: "Enable/disable automatic initiation of IKE SA negotiation.",
				Computed:    true,
			},
			"backup_gateway": {
				Type:        schema.TypeList,
				Description: "Instruct unity clients about the backup gateway address(es).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Description: "Address of backup gateway.",
							Computed:    true,
						},
					},
				},
			},
			"banner": {
				Type:        schema.TypeString,
				Description: "Message that unity client should display after connecting.",
				Computed:    true,
			},
			"cert_id_validation": {
				Type:        schema.TypeString,
				Description: "Enable/disable cross validation of peer ID and the identity in the peer's certificate as specified in RFC 4945.",
				Computed:    true,
			},
			"certificate": {
				Type:        schema.TypeList,
				Description: "The names of up to 4 signed personal certificates.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Certificate name.",
							Computed:    true,
						},
					},
				},
			},
			"childless_ike": {
				Type:        schema.TypeString,
				Description: "Enable/disable childless IKEv2 initiation (RFC 6023).",
				Computed:    true,
			},
			"client_auto_negotiate": {
				Type:        schema.TypeString,
				Description: "Enable/disable allowing the VPN client to bring up the tunnel when there is no traffic.",
				Computed:    true,
			},
			"client_keep_alive": {
				Type:        schema.TypeString,
				Description: "Enable/disable allowing the VPN client to keep the tunnel up when there is no traffic.",
				Computed:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"default_gw": {
				Type:        schema.TypeString,
				Description: "IPv4 address of default route gateway to use for traffic exiting the interface.",
				Computed:    true,
			},
			"default_gw_priority": {
				Type:        schema.TypeInt,
				Description: "Priority for default gateway route. A higher priority number signifies a less preferred route.",
				Computed:    true,
			},
			"dhcp_ra_giaddr": {
				Type:        schema.TypeString,
				Description: "Relay agent gateway IP address to use in the giaddr field of DHCP requests.",
				Computed:    true,
			},
			"dhcp6_ra_linkaddr": {
				Type:        schema.TypeString,
				Description: "Relay agent IPv6 link address to use in DHCP6 requests.",
				Computed:    true,
			},
			"dhgrp": {
				Type:        schema.TypeString,
				Description: "DH group.",
				Computed:    true,
			},
			"digital_signature_auth": {
				Type:        schema.TypeString,
				Description: "Enable/disable IKEv2 Digital Signature Authentication (RFC 7427).",
				Computed:    true,
			},
			"distance": {
				Type:        schema.TypeInt,
				Description: "Distance for routes added by IKE (1 - 255).",
				Computed:    true,
			},
			"dns_mode": {
				Type:        schema.TypeString,
				Description: "DNS server mode.",
				Computed:    true,
			},
			"domain": {
				Type:        schema.TypeString,
				Description: "Instruct unity clients about the single default DNS domain.",
				Computed:    true,
			},
			"dpd": {
				Type:        schema.TypeString,
				Description: "Dead Peer Detection mode.",
				Computed:    true,
			},
			"dpd_retrycount": {
				Type:        schema.TypeInt,
				Description: "Number of DPD retry attempts.",
				Computed:    true,
			},
			"dpd_retryinterval": {
				Type:        schema.TypeString,
				Description: "DPD retry interval.",
				Computed:    true,
			},
			"eap": {
				Type:        schema.TypeString,
				Description: "Enable/disable IKEv2 EAP authentication.",
				Computed:    true,
			},
			"eap_exclude_peergrp": {
				Type:        schema.TypeString,
				Description: "Peer group excluded from EAP authentication.",
				Computed:    true,
			},
			"eap_identity": {
				Type:        schema.TypeString,
				Description: "IKEv2 EAP peer identity type.",
				Computed:    true,
			},
			"encap_local_gw4": {
				Type:        schema.TypeString,
				Description: "Local IPv4 address of GRE/VXLAN tunnel.",
				Computed:    true,
			},
			"encap_local_gw6": {
				Type:        schema.TypeString,
				Description: "Local IPv6 address of GRE/VXLAN tunnel.",
				Computed:    true,
			},
			"encap_remote_gw4": {
				Type:        schema.TypeString,
				Description: "Remote IPv4 address of GRE/VXLAN tunnel.",
				Computed:    true,
			},
			"encap_remote_gw6": {
				Type:        schema.TypeString,
				Description: "Remote IPv6 address of GRE/VXLAN tunnel.",
				Computed:    true,
			},
			"encapsulation": {
				Type:        schema.TypeString,
				Description: "Enable/disable GRE/VXLAN/VPNID encapsulation.",
				Computed:    true,
			},
			"encapsulation_address": {
				Type:        schema.TypeString,
				Description: "Source for GRE/VXLAN tunnel address.",
				Computed:    true,
			},
			"enforce_unique_id": {
				Type:        schema.TypeString,
				Description: "Enable/disable peer ID uniqueness check.",
				Computed:    true,
			},
			"esn": {
				Type:        schema.TypeString,
				Description: "Extended sequence number (ESN) negotiation.",
				Computed:    true,
			},
			"exchange_interface_ip": {
				Type:        schema.TypeString,
				Description: "Enable/disable exchange of IPsec interface IP address.",
				Computed:    true,
			},
			"exchange_ip_addr4": {
				Type:        schema.TypeString,
				Description: "IPv4 address to exchange with peers.",
				Computed:    true,
			},
			"exchange_ip_addr6": {
				Type:        schema.TypeString,
				Description: "IPv6 address to exchange with peers.",
				Computed:    true,
			},
			"fec_base": {
				Type:        schema.TypeInt,
				Description: "Number of base Forward Error Correction packets (1 - 20).",
				Computed:    true,
			},
			"fec_codec": {
				Type:        schema.TypeString,
				Description: "Forward Error Correction encoding/decoding algorithm.",
				Computed:    true,
			},
			"fec_egress": {
				Type:        schema.TypeString,
				Description: "Enable/disable Forward Error Correction for egress IPsec traffic.",
				Computed:    true,
			},
			"fec_health_check": {
				Type:        schema.TypeString,
				Description: "SD-WAN health check.",
				Computed:    true,
			},
			"fec_ingress": {
				Type:        schema.TypeString,
				Description: "Enable/disable Forward Error Correction for ingress IPsec traffic.",
				Computed:    true,
			},
			"fec_mapping_profile": {
				Type:        schema.TypeString,
				Description: "Forward Error Correction (FEC) mapping profile.",
				Computed:    true,
			},
			"fec_receive_timeout": {
				Type:        schema.TypeInt,
				Description: "Timeout in milliseconds before dropping Forward Error Correction packets (1 - 1000).",
				Computed:    true,
			},
			"fec_redundant": {
				Type:        schema.TypeInt,
				Description: "Number of redundant Forward Error Correction packets (1 - 5 for reed-solomon, 1 for xor).",
				Computed:    true,
			},
			"fec_send_timeout": {
				Type:        schema.TypeInt,
				Description: "Timeout in milliseconds before sending Forward Error Correction packets (1 - 1000).",
				Computed:    true,
			},
			"forticlient_enforcement": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiClient enforcement.",
				Computed:    true,
			},
			"fragmentation": {
				Type:        schema.TypeString,
				Description: "Enable/disable fragment IKE message on re-transmission.",
				Computed:    true,
			},
			"fragmentation_mtu": {
				Type:        schema.TypeInt,
				Description: "IKE fragmentation MTU (500 - 16000).",
				Computed:    true,
			},
			"group_authentication": {
				Type:        schema.TypeString,
				Description: "Enable/disable IKEv2 IDi group authentication.",
				Computed:    true,
			},
			"group_authentication_secret": {
				Type:        schema.TypeString,
				Description: "Password for IKEv2 ID group authentication. ASCII string or hexadecimal indicated by a leading 0x.",
				Computed:    true,
				Sensitive:   true,
			},
			"ha_sync_esp_seqno": {
				Type:        schema.TypeString,
				Description: "Enable/disable sequence number jump ahead for IPsec HA.",
				Computed:    true,
			},
			"idle_timeout": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPsec tunnel idle timeout.",
				Computed:    true,
			},
			"idle_timeoutinterval": {
				Type:        schema.TypeInt,
				Description: "IPsec tunnel idle timeout in minutes (5 - 43200).",
				Computed:    true,
			},
			"ike_version": {
				Type:        schema.TypeString,
				Description: "IKE protocol version.",
				Computed:    true,
			},
			"include_local_lan": {
				Type:        schema.TypeString,
				Description: "Enable/disable allow local LAN access on unity clients.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Local physical, aggregate, or VLAN outgoing interface.",
				Computed:    true,
			},
			"ip_delay_interval": {
				Type:        schema.TypeInt,
				Description: "IP address reuse delay interval in seconds (0 - 28800).",
				Computed:    true,
			},
			"ip_fragmentation": {
				Type:        schema.TypeString,
				Description: "Determine whether IP packets are fragmented before or after IPsec encapsulation.",
				Computed:    true,
			},
			"ip_version": {
				Type:        schema.TypeString,
				Description: "IP version to use for VPN interface.",
				Computed:    true,
			},
			"ipv4_dns_server1": {
				Type:        schema.TypeString,
				Description: "IPv4 DNS server 1.",
				Computed:    true,
			},
			"ipv4_dns_server2": {
				Type:        schema.TypeString,
				Description: "IPv4 DNS server 2.",
				Computed:    true,
			},
			"ipv4_dns_server3": {
				Type:        schema.TypeString,
				Description: "IPv4 DNS server 3.",
				Computed:    true,
			},
			"ipv4_end_ip": {
				Type:        schema.TypeString,
				Description: "End of IPv4 range.",
				Computed:    true,
			},
			"ipv4_exclude_range": {
				Type:        schema.TypeList,
				Description: "Configuration Method IPv4 exclude ranges.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": {
							Type:        schema.TypeString,
							Description: "End of IPv4 exclusive range.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"start_ip": {
							Type:        schema.TypeString,
							Description: "Start of IPv4 exclusive range.",
							Computed:    true,
						},
					},
				},
			},
			"ipv4_name": {
				Type:        schema.TypeString,
				Description: "IPv4 address name.",
				Computed:    true,
			},
			"ipv4_netmask": {
				Type:        schema.TypeString,
				Description: "IPv4 Netmask.",
				Computed:    true,
			},
			"ipv4_split_exclude": {
				Type:        schema.TypeString,
				Description: "IPv4 subnets that should not be sent over the IPsec tunnel.",
				Computed:    true,
			},
			"ipv4_split_include": {
				Type:        schema.TypeString,
				Description: "IPv4 split-include subnets.",
				Computed:    true,
			},
			"ipv4_start_ip": {
				Type:        schema.TypeString,
				Description: "Start of IPv4 range.",
				Computed:    true,
			},
			"ipv4_wins_server1": {
				Type:        schema.TypeString,
				Description: "WINS server 1.",
				Computed:    true,
			},
			"ipv4_wins_server2": {
				Type:        schema.TypeString,
				Description: "WINS server 2.",
				Computed:    true,
			},
			"ipv6_dns_server1": {
				Type:        schema.TypeString,
				Description: "IPv6 DNS server 1.",
				Computed:    true,
			},
			"ipv6_dns_server2": {
				Type:        schema.TypeString,
				Description: "IPv6 DNS server 2.",
				Computed:    true,
			},
			"ipv6_dns_server3": {
				Type:        schema.TypeString,
				Description: "IPv6 DNS server 3.",
				Computed:    true,
			},
			"ipv6_end_ip": {
				Type:        schema.TypeString,
				Description: "End of IPv6 range.",
				Computed:    true,
			},
			"ipv6_exclude_range": {
				Type:        schema.TypeList,
				Description: "Configuration method IPv6 exclude ranges.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": {
							Type:        schema.TypeString,
							Description: "End of IPv6 exclusive range.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"start_ip": {
							Type:        schema.TypeString,
							Description: "Start of IPv6 exclusive range.",
							Computed:    true,
						},
					},
				},
			},
			"ipv6_name": {
				Type:        schema.TypeString,
				Description: "IPv6 address name.",
				Computed:    true,
			},
			"ipv6_prefix": {
				Type:        schema.TypeInt,
				Description: "IPv6 prefix.",
				Computed:    true,
			},
			"ipv6_split_exclude": {
				Type:        schema.TypeString,
				Description: "IPv6 subnets that should not be sent over the IPsec tunnel.",
				Computed:    true,
			},
			"ipv6_split_include": {
				Type:        schema.TypeString,
				Description: "IPv6 split-include subnets.",
				Computed:    true,
			},
			"ipv6_start_ip": {
				Type:        schema.TypeString,
				Description: "Start of IPv6 range.",
				Computed:    true,
			},
			"keepalive": {
				Type:        schema.TypeInt,
				Description: "NAT-T keep alive interval.",
				Computed:    true,
			},
			"keylife": {
				Type:        schema.TypeInt,
				Description: "Time to wait in seconds before phase 1 encryption key expires.",
				Computed:    true,
			},
			"local_gw": {
				Type:        schema.TypeString,
				Description: "IPv4 address of the local gateway's external interface.",
				Computed:    true,
			},
			"local_gw6": {
				Type:        schema.TypeString,
				Description: "IPv6 address of the local gateway's external interface.",
				Computed:    true,
			},
			"localid": {
				Type:        schema.TypeString,
				Description: "Local ID.",
				Computed:    true,
			},
			"localid_type": {
				Type:        schema.TypeString,
				Description: "Local ID type.",
				Computed:    true,
			},
			"loopback_asymroute": {
				Type:        schema.TypeString,
				Description: "Enable/disable asymmetric routing for IKE traffic on loopback interface.",
				Computed:    true,
			},
			"mesh_selector_type": {
				Type:        schema.TypeString,
				Description: "Add selectors containing subsets of the configuration depending on traffic.",
				Computed:    true,
			},
			"mode": {
				Type:        schema.TypeString,
				Description: "The ID protection mode used to establish a secure channel.",
				Computed:    true,
			},
			"mode_cfg": {
				Type:        schema.TypeString,
				Description: "Enable/disable configuration method.",
				Computed:    true,
			},
			"mode_cfg_allow_client_selector": {
				Type:        schema.TypeString,
				Description: "Enable/disable mode-cfg client to use custom phase2 selectors.",
				Computed:    true,
			},
			"monitor": {
				Type:        schema.TypeString,
				Description: "IPsec interface as backup for primary interface.",
				Computed:    true,
			},
			"monitor_hold_down_delay": {
				Type:        schema.TypeInt,
				Description: "Time to wait in seconds before recovery once primary re-establishes.",
				Computed:    true,
			},
			"monitor_hold_down_time": {
				Type:        schema.TypeString,
				Description: "Time of day at which to fail back to primary after it re-establishes.",
				Computed:    true,
			},
			"monitor_hold_down_type": {
				Type:        schema.TypeString,
				Description: "Recovery time method when primary interface re-establishes.",
				Computed:    true,
			},
			"monitor_hold_down_weekday": {
				Type:        schema.TypeString,
				Description: "Day of the week to recover once primary re-establishes.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "IPsec remote gateway name.",
				Required:    true,
			},
			"nattraversal": {
				Type:        schema.TypeString,
				Description: "Enable/disable NAT traversal.",
				Computed:    true,
			},
			"negotiate_timeout": {
				Type:        schema.TypeInt,
				Description: "IKE SA negotiation timeout in seconds (1 - 300).",
				Computed:    true,
			},
			"net_device": {
				Type:        schema.TypeString,
				Description: "Enable/disable kernel device creation.",
				Computed:    true,
			},
			"network_id": {
				Type:        schema.TypeInt,
				Description: "VPN gateway network ID.",
				Computed:    true,
			},
			"network_overlay": {
				Type:        schema.TypeString,
				Description: "Enable/disable network overlays.",
				Computed:    true,
			},
			"npu_offload": {
				Type:        schema.TypeString,
				Description: "Enable/disable offloading NPU.",
				Computed:    true,
			},
			"passive_mode": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPsec passive mode for static tunnels.",
				Computed:    true,
			},
			"peer": {
				Type:        schema.TypeString,
				Description: "Accept this peer certificate.",
				Computed:    true,
			},
			"peergrp": {
				Type:        schema.TypeString,
				Description: "Accept this peer certificate group.",
				Computed:    true,
			},
			"peerid": {
				Type:        schema.TypeString,
				Description: "Accept this peer identity.",
				Computed:    true,
			},
			"peertype": {
				Type:        schema.TypeString,
				Description: "Accept this peer type.",
				Computed:    true,
			},
			"ppk": {
				Type:        schema.TypeString,
				Description: "Enable/disable IKEv2 Postquantum Preshared Key (PPK).",
				Computed:    true,
			},
			"ppk_identity": {
				Type:        schema.TypeString,
				Description: "IKEv2 Postquantum Preshared Key Identity.",
				Computed:    true,
			},
			"ppk_secret": {
				Type:        schema.TypeString,
				Description: "IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).",
				Computed:    true,
				Sensitive:   true,
			},
			"priority": {
				Type:        schema.TypeInt,
				Description: "Priority for routes added by IKE (1 - 65535).",
				Computed:    true,
			},
			"proposal": {
				Type:        schema.TypeString,
				Description: "Phase1 proposal.",
				Computed:    true,
			},
			"psksecret": {
				Type:        schema.TypeString,
				Description: "Pre-shared secret for PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).",
				Computed:    true,
				Sensitive:   true,
			},
			"psksecret_remote": {
				Type:        schema.TypeString,
				Description: "Pre-shared secret for remote side PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).",
				Computed:    true,
				Sensitive:   true,
			},
			"reauth": {
				Type:        schema.TypeString,
				Description: "Enable/disable re-authentication upon IKE SA lifetime expiration.",
				Computed:    true,
			},
			"rekey": {
				Type:        schema.TypeString,
				Description: "Enable/disable phase1 rekey.",
				Computed:    true,
			},
			"remote_gw": {
				Type:        schema.TypeString,
				Description: "IPv4 address of the remote gateway's external interface.",
				Computed:    true,
			},
			"remote_gw6": {
				Type:        schema.TypeString,
				Description: "IPv6 address of the remote gateway's external interface.",
				Computed:    true,
			},
			"remotegw_ddns": {
				Type:        schema.TypeString,
				Description: "Domain name of remote gateway. For example, name.ddns.com.",
				Computed:    true,
			},
			"rsa_signature_format": {
				Type:        schema.TypeString,
				Description: "Digital Signature Authentication RSA signature format.",
				Computed:    true,
			},
			"save_password": {
				Type:        schema.TypeString,
				Description: "Enable/disable saving XAuth username and password on VPN clients.",
				Computed:    true,
			},
			"send_cert_chain": {
				Type:        schema.TypeString,
				Description: "Enable/disable sending certificate chain.",
				Computed:    true,
			},
			"signature_hash_alg": {
				Type:        schema.TypeString,
				Description: "Digital Signature Authentication hash algorithms.",
				Computed:    true,
			},
			"split_include_service": {
				Type:        schema.TypeString,
				Description: "Split-include services.",
				Computed:    true,
			},
			"suite_b": {
				Type:        schema.TypeString,
				Description: "Use Suite-B.",
				Computed:    true,
			},
			"tunnel_search": {
				Type:        schema.TypeString,
				Description: "Tunnel search method for when the interface is shared.",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Remote gateway type.",
				Computed:    true,
			},
			"unity_support": {
				Type:        schema.TypeString,
				Description: "Enable/disable support for Cisco UNITY Configuration Method extensions.",
				Computed:    true,
			},
			"usrgrp": {
				Type:        schema.TypeString,
				Description: "User group name for dialup peers.",
				Computed:    true,
			},
			"vni": {
				Type:        schema.TypeInt,
				Description: "VNI of VXLAN tunnel.",
				Computed:    true,
			},
			"wizard_type": {
				Type:        schema.TypeString,
				Description: "GUI VPN Wizard Type.",
				Computed:    true,
			},
			"xauthtype": {
				Type:        schema.TypeString,
				Description: "XAuth type.",
				Computed:    true,
			},
		},
	}
}

func dataSourceVpnIpsecPhase1InterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	// c.Retries = 1

	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	i := d.Get("name")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadVpnIpsecPhase1Interface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnIpsecPhase1Interface dataSource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] dataSource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	sort := false
	if v, ok := d.GetOk("dynamic_sort_subtable"); ok {
		if b, ok := v.(bool); ok {
			sort = b
		}
	}

	diags := refreshObjectVpnIpsecPhase1Interface(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
