// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure VPN remote gateway.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceVpnIpsecPhase1() *schema.Resource {
	return &schema.Resource{
		Description: "Configure VPN remote gateway.",

		CreateContext: resourceVpnIpsecPhase1Create,
		ReadContext:   resourceVpnIpsecPhase1Read,
		UpdateContext: resourceVpnIpsecPhase1Update,
		DeleteContext: resourceVpnIpsecPhase1Delete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"allow_append": {
				Type:        schema.TypeBool,
				Description: "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"acct_verify": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable verification of RADIUS accounting record.",
				Optional:    true,
				Computed:    true,
			},
			"add_gw_route": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable automatically add a route to the remote gateway.",
				Optional:    true,
				Computed:    true,
			},
			"add_route": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable control addition of a route to peer destination selector.",
				Optional:    true,
				Computed:    true,
			},
			"assign_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable assignment of IP to IPsec interface via configuration method.",
				Optional:    true,
				Computed:    true,
			},
			"assign_ip_from": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"range", "usrgrp", "dhcp", "name"}, false),

				Description: "Method by which the IP address will be assigned.",
				Optional:    true,
				Computed:    true,
			},
			"authmethod": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"psk", "signature"}, false),

				Description: "Authentication method.",
				Optional:    true,
				Computed:    true,
			},
			"authmethod_remote": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"psk", "signature"}, false),

				Description: "Authentication method (remote side).",
				Optional:    true,
				Computed:    true,
			},
			"authpasswd": {
				Type: schema.TypeString,

				Description: "XAuth password (max 35 characters).",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"authusr": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "XAuth user name.",
				Optional:    true,
				Computed:    true,
			},
			"authusrgrp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Authentication user group.",
				Optional:    true,
				Computed:    true,
			},
			"auto_negotiate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable automatic initiation of IKE SA negotiation.",
				Optional:    true,
				Computed:    true,
			},
			"backup_gateway": {
				Type:        schema.TypeList,
				Description: "Instruct unity clients about the backup gateway address(es).",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address of backup gateway.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"banner": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1024),

				Description: "Message that unity client should display after connecting.",
				Optional:    true,
				Computed:    true,
			},
			"cert_id_validation": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable cross validation of peer ID and the identity in the peer's certificate as specified in RFC 4945.",
				Optional:    true,
				Computed:    true,
			},
			"certificate": {
				Type:        schema.TypeList,
				Description: "Names of up to 4 signed personal certificates.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Certificate name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"childless_ike": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable childless IKEv2 initiation (RFC 6023).",
				Optional:    true,
				Computed:    true,
			},
			"client_auto_negotiate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable allowing the VPN client to bring up the tunnel when there is no traffic.",
				Optional:    true,
				Computed:    true,
			},
			"client_keep_alive": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable allowing the VPN client to keep the tunnel up when there is no traffic.",
				Optional:    true,
				Computed:    true,
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_ra_giaddr": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Relay agent gateway IP address to use in the giaddr field of DHCP requests.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp6_ra_linkaddr": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Relay agent IPv6 link address to use in DHCP6 requests.",
				Optional:         true,
				Computed:         true,
			},
			"dhgrp": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "DH group.",
				Optional:         true,
				Computed:         true,
			},
			"digital_signature_auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IKEv2 Digital Signature Authentication (RFC 7427).",
				Optional:    true,
				Computed:    true,
			},
			"distance": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Distance for routes added by IKE (1 - 255).",
				Optional:    true,
				Computed:    true,
			},
			"dns_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"manual", "auto"}, false),

				Description: "DNS server mode.",
				Optional:    true,
				Computed:    true,
			},
			"domain": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Instruct unity clients about the single default DNS domain.",
				Optional:    true,
				Computed:    true,
			},
			"dpd": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "on-idle", "on-demand"}, false),

				Description: "Dead Peer Detection mode.",
				Optional:    true,
				Computed:    true,
			},
			"dpd_retrycount": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10),

				Description: "Number of DPD retry attempts.",
				Optional:    true,
				Computed:    true,
			},
			"dpd_retryinterval": {
				Type: schema.TypeString,

				Description: "DPD retry interval.",
				Optional:    true,
				Computed:    true,
			},
			"eap": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IKEv2 EAP authentication.",
				Optional:    true,
				Computed:    true,
			},
			"eap_exclude_peergrp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Peer group excluded from EAP authentication.",
				Optional:    true,
				Computed:    true,
			},
			"eap_identity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"use-id-payload", "send-request"}, false),

				Description: "IKEv2 EAP peer identity type.",
				Optional:    true,
				Computed:    true,
			},
			"enforce_unique_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "keep-new", "keep-old"}, false),

				Description: "Enable/disable peer ID uniqueness check.",
				Optional:    true,
				Computed:    true,
			},
			"esn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"require", "allow", "disable"}, false),

				Description: "Extended sequence number (ESN) negotiation.",
				Optional:    true,
				Computed:    true,
			},
			"fec_base": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 20),

				Description: "Number of base Forward Error Correction packets (1 - 20).",
				Optional:    true,
				Computed:    true,
			},
			"fec_codec": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"rs", "xor"}, false),

				Description: "Forward Error Correction encoding/decoding algorithm.",
				Optional:    true,
				Computed:    true,
			},
			"fec_egress": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Forward Error Correction for egress IPsec traffic.",
				Optional:    true,
				Computed:    true,
			},
			"fec_health_check": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "SD-WAN health check.",
				Optional:    true,
				Computed:    true,
			},
			"fec_ingress": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Forward Error Correction for ingress IPsec traffic.",
				Optional:    true,
				Computed:    true,
			},
			"fec_mapping_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Forward Error Correction (FEC) mapping profile.",
				Optional:    true,
				Computed:    true,
			},
			"fec_receive_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1000),

				Description: "Timeout in milliseconds before dropping Forward Error Correction packets (1 - 1000).",
				Optional:    true,
				Computed:    true,
			},
			"fec_redundant": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 5),

				Description: "Number of redundant Forward Error Correction packets (1 - 5 for reed-solomon, 1 for xor).",
				Optional:    true,
				Computed:    true,
			},
			"fec_send_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1000),

				Description: "Timeout in milliseconds before sending Forward Error Correction packets (1 - 1000).",
				Optional:    true,
				Computed:    true,
			},
			"forticlient_enforcement": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiClient enforcement.",
				Optional:    true,
				Computed:    true,
			},
			"fragmentation": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable fragment IKE message on re-transmission.",
				Optional:    true,
				Computed:    true,
			},
			"fragmentation_mtu": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(500, 16000),

				Description: "IKE fragmentation MTU (500 - 16000).",
				Optional:    true,
				Computed:    true,
			},
			"group_authentication": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IKEv2 IDi group authentication.",
				Optional:    true,
				Computed:    true,
			},
			"group_authentication_secret": {
				Type: schema.TypeString,

				Description: "Password for IKEv2 IDi group authentication.  (ASCII string or hexadecimal indicated by a leading 0x.)",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"ha_sync_esp_seqno": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable sequence number jump ahead for IPsec HA.",
				Optional:    true,
				Computed:    true,
			},
			"idle_timeout": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPsec tunnel idle timeout.",
				Optional:    true,
				Computed:    true,
			},
			"idle_timeoutinterval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 43200),

				Description: "IPsec tunnel idle timeout in minutes (5 - 43200).",
				Optional:    true,
				Computed:    true,
			},
			"ike_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"1", "2"}, false),

				Description: "IKE protocol version.",
				Optional:    true,
				Computed:    true,
			},
			"include_local_lan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable allow local LAN access on unity clients.",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Local physical, aggregate, or VLAN outgoing interface.",
				Optional:    true,
				Computed:    true,
			},
			"ip_delay_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 28800),

				Description: "IP address reuse delay interval in seconds (0 - 28800).",
				Optional:    true,
				Computed:    true,
			},
			"ipv4_dns_server1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IPv4 DNS server 1.",
				Optional:    true,
				Computed:    true,
			},
			"ipv4_dns_server2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IPv4 DNS server 2.",
				Optional:    true,
				Computed:    true,
			},
			"ipv4_dns_server3": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IPv4 DNS server 3.",
				Optional:    true,
				Computed:    true,
			},
			"ipv4_end_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "End of IPv4 range.",
				Optional:    true,
				Computed:    true,
			},
			"ipv4_exclude_range": {
				Type:        schema.TypeList,
				Description: "Configuration Method IPv4 exclude ranges.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "End of IPv4 exclusive range.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"start_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Start of IPv4 exclusive range.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ipv4_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "IPv4 address name.",
				Optional:    true,
				Computed:    true,
			},
			"ipv4_netmask": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4Netmask,

				Description: "IPv4 Netmask.",
				Optional:    true,
				Computed:    true,
			},
			"ipv4_split_exclude": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "IPv4 subnets that should not be sent over the IPsec tunnel.",
				Optional:    true,
				Computed:    true,
			},
			"ipv4_split_include": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "IPv4 split-include subnets.",
				Optional:    true,
				Computed:    true,
			},
			"ipv4_start_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Start of IPv4 range.",
				Optional:    true,
				Computed:    true,
			},
			"ipv4_wins_server1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "WINS server 1.",
				Optional:    true,
				Computed:    true,
			},
			"ipv4_wins_server2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "WINS server 2.",
				Optional:    true,
				Computed:    true,
			},
			"ipv6_dns_server1": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 DNS server 1.",
				Optional:         true,
				Computed:         true,
			},
			"ipv6_dns_server2": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 DNS server 2.",
				Optional:         true,
				Computed:         true,
			},
			"ipv6_dns_server3": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 DNS server 3.",
				Optional:         true,
				Computed:         true,
			},
			"ipv6_end_ip": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "End of IPv6 range.",
				Optional:         true,
				Computed:         true,
			},
			"ipv6_exclude_range": {
				Type:        schema.TypeList,
				Description: "Configuration method IPv6 exclude ranges.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "End of IPv6 exclusive range.",
							Optional:         true,
							Computed:         true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"start_ip": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "Start of IPv6 exclusive range.",
							Optional:         true,
							Computed:         true,
						},
					},
				},
			},
			"ipv6_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "IPv6 address name.",
				Optional:    true,
				Computed:    true,
			},
			"ipv6_prefix": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 128),

				Description: "IPv6 prefix.",
				Optional:    true,
				Computed:    true,
			},
			"ipv6_split_exclude": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "IPv6 subnets that should not be sent over the IPsec tunnel.",
				Optional:    true,
				Computed:    true,
			},
			"ipv6_split_include": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "IPv6 split-include subnets.",
				Optional:    true,
				Computed:    true,
			},
			"ipv6_start_ip": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Start of IPv6 range.",
				Optional:         true,
				Computed:         true,
			},
			"keepalive": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 900),

				Description: "NAT-T keep alive interval.",
				Optional:    true,
				Computed:    true,
			},
			"keylife": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(120, 172800),

				Description: "Time to wait in seconds before phase 1 encryption key expires.",
				Optional:    true,
				Computed:    true,
			},
			"local_gw": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Local VPN gateway.",
				Optional:    true,
				Computed:    true,
			},
			"localid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Local ID.",
				Optional:    true,
				Computed:    true,
			},
			"localid_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "fqdn", "user-fqdn", "keyid", "address", "asn1dn"}, false),

				Description: "Local ID type.",
				Optional:    true,
				Computed:    true,
			},
			"loopback_asymroute": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable asymmetric routing for IKE traffic on loopback interface.",
				Optional:    true,
				Computed:    true,
			},
			"mesh_selector_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "subnet", "host"}, false),

				Description: "Add selectors containing subsets of the configuration depending on traffic.",
				Optional:    true,
				Computed:    true,
			},
			"mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"aggressive", "main"}, false),

				Description: "ID protection mode used to establish a secure channel.",
				Optional:    true,
				Computed:    true,
			},
			"mode_cfg": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable configuration method.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "IPsec remote gateway name.",
				ForceNew:    true,
				Required:    true,
			},
			"nattraversal": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable", "forced"}, false),

				Description: "Enable/disable NAT traversal.",
				Optional:    true,
				Computed:    true,
			},
			"negotiate_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 300),

				Description: "IKE SA negotiation timeout in seconds (1 - 300).",
				Optional:    true,
				Computed:    true,
			},
			"network_id": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "VPN gateway network ID.",
				Optional:    true,
				Computed:    true,
			},
			"network_overlay": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable network overlays.",
				Optional:    true,
				Computed:    true,
			},
			"npu_offload": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable offloading NPU.",
				Optional:    true,
				Computed:    true,
			},
			"peer": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Accept this peer certificate.",
				Optional:    true,
				Computed:    true,
			},
			"peergrp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Accept this peer certificate group.",
				Optional:    true,
				Computed:    true,
			},
			"peerid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Accept this peer identity.",
				Optional:    true,
				Computed:    true,
			},
			"peertype": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"any", "one", "dialup", "peer", "peergrp"}, false),

				Description: "Accept this peer type.",
				Optional:    true,
				Computed:    true,
			},
			"ppk": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "allow", "require"}, false),

				Description: "Enable/disable IKEv2 Postquantum Preshared Key (PPK).",
				Optional:    true,
				Computed:    true,
			},
			"ppk_identity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "IKEv2 Postquantum Preshared Key Identity.",
				Optional:    true,
				Computed:    true,
			},
			"ppk_secret": {
				Type: schema.TypeString,

				Description: "IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"priority": {
				Type: schema.TypeInt,

				Description: "Priority for routes added by IKE (0 - 4294967295).",
				Optional:    true,
				Computed:    true,
			},
			"proposal": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Phase1 proposal.",
				Optional:         true,
				Computed:         true,
			},
			"psksecret": {
				Type: schema.TypeString,

				Description: "Pre-shared secret for PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"psksecret_remote": {
				Type: schema.TypeString,

				Description: "Pre-shared secret for remote side PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"reauth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable re-authentication upon IKE SA lifetime expiration.",
				Optional:    true,
				Computed:    true,
			},
			"rekey": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable phase1 rekey.",
				Optional:    true,
				Computed:    true,
			},
			"remote_gw": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Remote VPN gateway.",
				Optional:    true,
				Computed:    true,
			},
			"remotegw_ddns": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Domain name of remote gateway (eg. name.DDNS.com).",
				Optional:    true,
				Computed:    true,
			},
			"rsa_signature_format": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"pkcs1", "pss"}, false),

				Description: "Digital Signature Authentication RSA signature format.",
				Optional:    true,
				Computed:    true,
			},
			"save_password": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable saving XAuth username and password on VPN clients.",
				Optional:    true,
				Computed:    true,
			},
			"send_cert_chain": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable sending certificate chain.",
				Optional:    true,
				Computed:    true,
			},
			"signature_hash_alg": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Digital Signature Authentication hash algorithms.",
				Optional:         true,
				Computed:         true,
			},
			"split_include_service": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Split-include services.",
				Optional:    true,
				Computed:    true,
			},
			"suite_b": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "suite-b-gcm-128", "suite-b-gcm-256"}, false),

				Description: "Use Suite-B.",
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"static", "dynamic", "ddns"}, false),

				Description: "Remote gateway type.",
				Optional:    true,
				Computed:    true,
			},
			"unity_support": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable support for Cisco UNITY Configuration Method extensions.",
				Optional:    true,
				Computed:    true,
			},
			"usrgrp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "User group name for dialup peers.",
				Optional:    true,
				Computed:    true,
			},
			"wizard_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"custom", "dialup-forticlient", "dialup-ios", "dialup-android", "dialup-windows", "dialup-cisco", "static-fortigate", "dialup-fortigate", "static-cisco", "dialup-cisco-fw", "simplified-static-fortigate", "hub-fortigate-auto-discovery", "spoke-fortigate-auto-discovery"}, false),

				Description: "GUI VPN Wizard Type.",
				Optional:    true,
				Computed:    true,
			},
			"xauthtype": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "client", "pap", "chap", "auto"}, false),

				Description: "XAuth type.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceVpnIpsecPhase1Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	var diags diag.Diagnostics
	var err error
	// c.Retries = 1
	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	allow_append := false
	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}
	urlparams.AllowAppend = &allow_append

	mkey := ""
	key := "name"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating VpnIpsecPhase1 resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectVpnIpsecPhase1(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateVpnIpsecPhase1(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnIpsecPhase1")
	}

	return resourceVpnIpsecPhase1Read(ctx, d, meta)
}

func resourceVpnIpsecPhase1Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()
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

	obj, diags := getObjectVpnIpsecPhase1(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateVpnIpsecPhase1(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating VpnIpsecPhase1 resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnIpsecPhase1")
	}

	return resourceVpnIpsecPhase1Read(ctx, d, meta)
}

func resourceVpnIpsecPhase1Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

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

	err := c.Cmdb.DeleteVpnIpsecPhase1(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting VpnIpsecPhase1 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecPhase1Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

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

	ptp := true
	urlparams.PlainTextPassword = &ptp

	o, err := c.Cmdb.ReadVpnIpsecPhase1(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnIpsecPhase1 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	sort := false
	if v, ok := d.GetOk("dynamic_sort_subtable"); ok {
		if b, ok := v.(bool); ok {
			sort = b
		}
	}

	diags := refreshObjectVpnIpsecPhase1(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenVpnIpsecPhase1BackupGateway(v *[]models.VpnIpsecPhase1BackupGateway, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Address; tmp != nil {
				v["address"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "address")
	}

	return flat
}

func flattenVpnIpsecPhase1Certificate(v *[]models.VpnIpsecPhase1Certificate, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenVpnIpsecPhase1Ipv4ExcludeRange(v *[]models.VpnIpsecPhase1Ipv4ExcludeRange, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.EndIp; tmp != nil {
				v["end_ip"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.StartIp; tmp != nil {
				v["start_ip"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenVpnIpsecPhase1Ipv6ExcludeRange(v *[]models.VpnIpsecPhase1Ipv6ExcludeRange, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.EndIp; tmp != nil {
				v["end_ip"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.StartIp; tmp != nil {
				v["start_ip"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectVpnIpsecPhase1(d *schema.ResourceData, o *models.VpnIpsecPhase1, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AcctVerify != nil {
		v := *o.AcctVerify

		if err = d.Set("acct_verify", v); err != nil {
			return diag.Errorf("error reading acct_verify: %v", err)
		}
	}

	if o.AddGwRoute != nil {
		v := *o.AddGwRoute

		if err = d.Set("add_gw_route", v); err != nil {
			return diag.Errorf("error reading add_gw_route: %v", err)
		}
	}

	if o.AddRoute != nil {
		v := *o.AddRoute

		if err = d.Set("add_route", v); err != nil {
			return diag.Errorf("error reading add_route: %v", err)
		}
	}

	if o.AssignIp != nil {
		v := *o.AssignIp

		if err = d.Set("assign_ip", v); err != nil {
			return diag.Errorf("error reading assign_ip: %v", err)
		}
	}

	if o.AssignIpFrom != nil {
		v := *o.AssignIpFrom

		if err = d.Set("assign_ip_from", v); err != nil {
			return diag.Errorf("error reading assign_ip_from: %v", err)
		}
	}

	if o.Authmethod != nil {
		v := *o.Authmethod

		if err = d.Set("authmethod", v); err != nil {
			return diag.Errorf("error reading authmethod: %v", err)
		}
	}

	if o.AuthmethodRemote != nil {
		v := *o.AuthmethodRemote

		if err = d.Set("authmethod_remote", v); err != nil {
			return diag.Errorf("error reading authmethod_remote: %v", err)
		}
	}

	if o.Authpasswd != nil {
		v := *o.Authpasswd

		if v == "ENC XXXX" {
		} else if err = d.Set("authpasswd", v); err != nil {
			return diag.Errorf("error reading authpasswd: %v", err)
		}
	}

	if o.Authusr != nil {
		v := *o.Authusr

		if err = d.Set("authusr", v); err != nil {
			return diag.Errorf("error reading authusr: %v", err)
		}
	}

	if o.Authusrgrp != nil {
		v := *o.Authusrgrp

		if err = d.Set("authusrgrp", v); err != nil {
			return diag.Errorf("error reading authusrgrp: %v", err)
		}
	}

	if o.AutoNegotiate != nil {
		v := *o.AutoNegotiate

		if err = d.Set("auto_negotiate", v); err != nil {
			return diag.Errorf("error reading auto_negotiate: %v", err)
		}
	}

	if o.BackupGateway != nil {
		if err = d.Set("backup_gateway", flattenVpnIpsecPhase1BackupGateway(o.BackupGateway, sort)); err != nil {
			return diag.Errorf("error reading backup_gateway: %v", err)
		}
	}

	if o.Banner != nil {
		v := *o.Banner

		if err = d.Set("banner", v); err != nil {
			return diag.Errorf("error reading banner: %v", err)
		}
	}

	if o.CertIdValidation != nil {
		v := *o.CertIdValidation

		if err = d.Set("cert_id_validation", v); err != nil {
			return diag.Errorf("error reading cert_id_validation: %v", err)
		}
	}

	if o.Certificate != nil {
		if err = d.Set("certificate", flattenVpnIpsecPhase1Certificate(o.Certificate, sort)); err != nil {
			return diag.Errorf("error reading certificate: %v", err)
		}
	}

	if o.ChildlessIke != nil {
		v := *o.ChildlessIke

		if err = d.Set("childless_ike", v); err != nil {
			return diag.Errorf("error reading childless_ike: %v", err)
		}
	}

	if o.ClientAutoNegotiate != nil {
		v := *o.ClientAutoNegotiate

		if err = d.Set("client_auto_negotiate", v); err != nil {
			return diag.Errorf("error reading client_auto_negotiate: %v", err)
		}
	}

	if o.ClientKeepAlive != nil {
		v := *o.ClientKeepAlive

		if err = d.Set("client_keep_alive", v); err != nil {
			return diag.Errorf("error reading client_keep_alive: %v", err)
		}
	}

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.DhcpRaGiaddr != nil {
		v := *o.DhcpRaGiaddr

		if err = d.Set("dhcp_ra_giaddr", v); err != nil {
			return diag.Errorf("error reading dhcp_ra_giaddr: %v", err)
		}
	}

	if o.Dhcp6RaLinkaddr != nil {
		v := *o.Dhcp6RaLinkaddr

		if err = d.Set("dhcp6_ra_linkaddr", v); err != nil {
			return diag.Errorf("error reading dhcp6_ra_linkaddr: %v", err)
		}
	}

	if o.Dhgrp != nil {
		v := *o.Dhgrp

		if err = d.Set("dhgrp", v); err != nil {
			return diag.Errorf("error reading dhgrp: %v", err)
		}
	}

	if o.DigitalSignatureAuth != nil {
		v := *o.DigitalSignatureAuth

		if err = d.Set("digital_signature_auth", v); err != nil {
			return diag.Errorf("error reading digital_signature_auth: %v", err)
		}
	}

	if o.Distance != nil {
		v := *o.Distance

		if err = d.Set("distance", v); err != nil {
			return diag.Errorf("error reading distance: %v", err)
		}
	}

	if o.DnsMode != nil {
		v := *o.DnsMode

		if err = d.Set("dns_mode", v); err != nil {
			return diag.Errorf("error reading dns_mode: %v", err)
		}
	}

	if o.Domain != nil {
		v := *o.Domain

		if err = d.Set("domain", v); err != nil {
			return diag.Errorf("error reading domain: %v", err)
		}
	}

	if o.Dpd != nil {
		v := *o.Dpd

		if err = d.Set("dpd", v); err != nil {
			return diag.Errorf("error reading dpd: %v", err)
		}
	}

	if o.DpdRetrycount != nil {
		v := *o.DpdRetrycount

		if err = d.Set("dpd_retrycount", v); err != nil {
			return diag.Errorf("error reading dpd_retrycount: %v", err)
		}
	}

	if o.DpdRetryinterval != nil {
		v := *o.DpdRetryinterval

		if err = d.Set("dpd_retryinterval", v); err != nil {
			return diag.Errorf("error reading dpd_retryinterval: %v", err)
		}
	}

	if o.Eap != nil {
		v := *o.Eap

		if err = d.Set("eap", v); err != nil {
			return diag.Errorf("error reading eap: %v", err)
		}
	}

	if o.EapExcludePeergrp != nil {
		v := *o.EapExcludePeergrp

		if err = d.Set("eap_exclude_peergrp", v); err != nil {
			return diag.Errorf("error reading eap_exclude_peergrp: %v", err)
		}
	}

	if o.EapIdentity != nil {
		v := *o.EapIdentity

		if err = d.Set("eap_identity", v); err != nil {
			return diag.Errorf("error reading eap_identity: %v", err)
		}
	}

	if o.EnforceUniqueId != nil {
		v := *o.EnforceUniqueId

		if err = d.Set("enforce_unique_id", v); err != nil {
			return diag.Errorf("error reading enforce_unique_id: %v", err)
		}
	}

	if o.Esn != nil {
		v := *o.Esn

		if err = d.Set("esn", v); err != nil {
			return diag.Errorf("error reading esn: %v", err)
		}
	}

	if o.FecBase != nil {
		v := *o.FecBase

		if err = d.Set("fec_base", v); err != nil {
			return diag.Errorf("error reading fec_base: %v", err)
		}
	}

	if o.FecCodec != nil {
		v := *o.FecCodec

		if err = d.Set("fec_codec", v); err != nil {
			return diag.Errorf("error reading fec_codec: %v", err)
		}
	}

	if o.FecEgress != nil {
		v := *o.FecEgress

		if err = d.Set("fec_egress", v); err != nil {
			return diag.Errorf("error reading fec_egress: %v", err)
		}
	}

	if o.FecHealthCheck != nil {
		v := *o.FecHealthCheck

		if err = d.Set("fec_health_check", v); err != nil {
			return diag.Errorf("error reading fec_health_check: %v", err)
		}
	}

	if o.FecIngress != nil {
		v := *o.FecIngress

		if err = d.Set("fec_ingress", v); err != nil {
			return diag.Errorf("error reading fec_ingress: %v", err)
		}
	}

	if o.FecMappingProfile != nil {
		v := *o.FecMappingProfile

		if err = d.Set("fec_mapping_profile", v); err != nil {
			return diag.Errorf("error reading fec_mapping_profile: %v", err)
		}
	}

	if o.FecReceiveTimeout != nil {
		v := *o.FecReceiveTimeout

		if err = d.Set("fec_receive_timeout", v); err != nil {
			return diag.Errorf("error reading fec_receive_timeout: %v", err)
		}
	}

	if o.FecRedundant != nil {
		v := *o.FecRedundant

		if err = d.Set("fec_redundant", v); err != nil {
			return diag.Errorf("error reading fec_redundant: %v", err)
		}
	}

	if o.FecSendTimeout != nil {
		v := *o.FecSendTimeout

		if err = d.Set("fec_send_timeout", v); err != nil {
			return diag.Errorf("error reading fec_send_timeout: %v", err)
		}
	}

	if o.ForticlientEnforcement != nil {
		v := *o.ForticlientEnforcement

		if err = d.Set("forticlient_enforcement", v); err != nil {
			return diag.Errorf("error reading forticlient_enforcement: %v", err)
		}
	}

	if o.Fragmentation != nil {
		v := *o.Fragmentation

		if err = d.Set("fragmentation", v); err != nil {
			return diag.Errorf("error reading fragmentation: %v", err)
		}
	}

	if o.FragmentationMtu != nil {
		v := *o.FragmentationMtu

		if err = d.Set("fragmentation_mtu", v); err != nil {
			return diag.Errorf("error reading fragmentation_mtu: %v", err)
		}
	}

	if o.GroupAuthentication != nil {
		v := *o.GroupAuthentication

		if err = d.Set("group_authentication", v); err != nil {
			return diag.Errorf("error reading group_authentication: %v", err)
		}
	}

	if o.GroupAuthenticationSecret != nil {
		v := *o.GroupAuthenticationSecret

		if v == "ENC XXXX" {
		} else if err = d.Set("group_authentication_secret", v); err != nil {
			return diag.Errorf("error reading group_authentication_secret: %v", err)
		}
	}

	if o.HaSyncEspSeqno != nil {
		v := *o.HaSyncEspSeqno

		if err = d.Set("ha_sync_esp_seqno", v); err != nil {
			return diag.Errorf("error reading ha_sync_esp_seqno: %v", err)
		}
	}

	if o.IdleTimeout != nil {
		v := *o.IdleTimeout

		if err = d.Set("idle_timeout", v); err != nil {
			return diag.Errorf("error reading idle_timeout: %v", err)
		}
	}

	if o.IdleTimeoutinterval != nil {
		v := *o.IdleTimeoutinterval

		if err = d.Set("idle_timeoutinterval", v); err != nil {
			return diag.Errorf("error reading idle_timeoutinterval: %v", err)
		}
	}

	if o.IkeVersion != nil {
		v := *o.IkeVersion

		if err = d.Set("ike_version", v); err != nil {
			return diag.Errorf("error reading ike_version: %v", err)
		}
	}

	if o.IncludeLocalLan != nil {
		v := *o.IncludeLocalLan

		if err = d.Set("include_local_lan", v); err != nil {
			return diag.Errorf("error reading include_local_lan: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.IpDelayInterval != nil {
		v := *o.IpDelayInterval

		if err = d.Set("ip_delay_interval", v); err != nil {
			return diag.Errorf("error reading ip_delay_interval: %v", err)
		}
	}

	if o.Ipv4DnsServer1 != nil {
		v := *o.Ipv4DnsServer1

		if err = d.Set("ipv4_dns_server1", v); err != nil {
			return diag.Errorf("error reading ipv4_dns_server1: %v", err)
		}
	}

	if o.Ipv4DnsServer2 != nil {
		v := *o.Ipv4DnsServer2

		if err = d.Set("ipv4_dns_server2", v); err != nil {
			return diag.Errorf("error reading ipv4_dns_server2: %v", err)
		}
	}

	if o.Ipv4DnsServer3 != nil {
		v := *o.Ipv4DnsServer3

		if err = d.Set("ipv4_dns_server3", v); err != nil {
			return diag.Errorf("error reading ipv4_dns_server3: %v", err)
		}
	}

	if o.Ipv4EndIp != nil {
		v := *o.Ipv4EndIp

		if err = d.Set("ipv4_end_ip", v); err != nil {
			return diag.Errorf("error reading ipv4_end_ip: %v", err)
		}
	}

	if o.Ipv4ExcludeRange != nil {
		if err = d.Set("ipv4_exclude_range", flattenVpnIpsecPhase1Ipv4ExcludeRange(o.Ipv4ExcludeRange, sort)); err != nil {
			return diag.Errorf("error reading ipv4_exclude_range: %v", err)
		}
	}

	if o.Ipv4Name != nil {
		v := *o.Ipv4Name

		if err = d.Set("ipv4_name", v); err != nil {
			return diag.Errorf("error reading ipv4_name: %v", err)
		}
	}

	if o.Ipv4Netmask != nil {
		v := *o.Ipv4Netmask

		if err = d.Set("ipv4_netmask", v); err != nil {
			return diag.Errorf("error reading ipv4_netmask: %v", err)
		}
	}

	if o.Ipv4SplitExclude != nil {
		v := *o.Ipv4SplitExclude

		if err = d.Set("ipv4_split_exclude", v); err != nil {
			return diag.Errorf("error reading ipv4_split_exclude: %v", err)
		}
	}

	if o.Ipv4SplitInclude != nil {
		v := *o.Ipv4SplitInclude

		if err = d.Set("ipv4_split_include", v); err != nil {
			return diag.Errorf("error reading ipv4_split_include: %v", err)
		}
	}

	if o.Ipv4StartIp != nil {
		v := *o.Ipv4StartIp

		if err = d.Set("ipv4_start_ip", v); err != nil {
			return diag.Errorf("error reading ipv4_start_ip: %v", err)
		}
	}

	if o.Ipv4WinsServer1 != nil {
		v := *o.Ipv4WinsServer1

		if err = d.Set("ipv4_wins_server1", v); err != nil {
			return diag.Errorf("error reading ipv4_wins_server1: %v", err)
		}
	}

	if o.Ipv4WinsServer2 != nil {
		v := *o.Ipv4WinsServer2

		if err = d.Set("ipv4_wins_server2", v); err != nil {
			return diag.Errorf("error reading ipv4_wins_server2: %v", err)
		}
	}

	if o.Ipv6DnsServer1 != nil {
		v := *o.Ipv6DnsServer1

		if err = d.Set("ipv6_dns_server1", v); err != nil {
			return diag.Errorf("error reading ipv6_dns_server1: %v", err)
		}
	}

	if o.Ipv6DnsServer2 != nil {
		v := *o.Ipv6DnsServer2

		if err = d.Set("ipv6_dns_server2", v); err != nil {
			return diag.Errorf("error reading ipv6_dns_server2: %v", err)
		}
	}

	if o.Ipv6DnsServer3 != nil {
		v := *o.Ipv6DnsServer3

		if err = d.Set("ipv6_dns_server3", v); err != nil {
			return diag.Errorf("error reading ipv6_dns_server3: %v", err)
		}
	}

	if o.Ipv6EndIp != nil {
		v := *o.Ipv6EndIp

		if err = d.Set("ipv6_end_ip", v); err != nil {
			return diag.Errorf("error reading ipv6_end_ip: %v", err)
		}
	}

	if o.Ipv6ExcludeRange != nil {
		if err = d.Set("ipv6_exclude_range", flattenVpnIpsecPhase1Ipv6ExcludeRange(o.Ipv6ExcludeRange, sort)); err != nil {
			return diag.Errorf("error reading ipv6_exclude_range: %v", err)
		}
	}

	if o.Ipv6Name != nil {
		v := *o.Ipv6Name

		if err = d.Set("ipv6_name", v); err != nil {
			return diag.Errorf("error reading ipv6_name: %v", err)
		}
	}

	if o.Ipv6Prefix != nil {
		v := *o.Ipv6Prefix

		if err = d.Set("ipv6_prefix", v); err != nil {
			return diag.Errorf("error reading ipv6_prefix: %v", err)
		}
	}

	if o.Ipv6SplitExclude != nil {
		v := *o.Ipv6SplitExclude

		if err = d.Set("ipv6_split_exclude", v); err != nil {
			return diag.Errorf("error reading ipv6_split_exclude: %v", err)
		}
	}

	if o.Ipv6SplitInclude != nil {
		v := *o.Ipv6SplitInclude

		if err = d.Set("ipv6_split_include", v); err != nil {
			return diag.Errorf("error reading ipv6_split_include: %v", err)
		}
	}

	if o.Ipv6StartIp != nil {
		v := *o.Ipv6StartIp

		if err = d.Set("ipv6_start_ip", v); err != nil {
			return diag.Errorf("error reading ipv6_start_ip: %v", err)
		}
	}

	if o.Keepalive != nil {
		v := *o.Keepalive

		if err = d.Set("keepalive", v); err != nil {
			return diag.Errorf("error reading keepalive: %v", err)
		}
	}

	if o.Keylife != nil {
		v := *o.Keylife

		if err = d.Set("keylife", v); err != nil {
			return diag.Errorf("error reading keylife: %v", err)
		}
	}

	if o.LocalGw != nil {
		v := *o.LocalGw

		if err = d.Set("local_gw", v); err != nil {
			return diag.Errorf("error reading local_gw: %v", err)
		}
	}

	if o.Localid != nil {
		v := *o.Localid

		if err = d.Set("localid", v); err != nil {
			return diag.Errorf("error reading localid: %v", err)
		}
	}

	if o.LocalidType != nil {
		v := *o.LocalidType

		if err = d.Set("localid_type", v); err != nil {
			return diag.Errorf("error reading localid_type: %v", err)
		}
	}

	if o.LoopbackAsymroute != nil {
		v := *o.LoopbackAsymroute

		if err = d.Set("loopback_asymroute", v); err != nil {
			return diag.Errorf("error reading loopback_asymroute: %v", err)
		}
	}

	if o.MeshSelectorType != nil {
		v := *o.MeshSelectorType

		if err = d.Set("mesh_selector_type", v); err != nil {
			return diag.Errorf("error reading mesh_selector_type: %v", err)
		}
	}

	if o.Mode != nil {
		v := *o.Mode

		if err = d.Set("mode", v); err != nil {
			return diag.Errorf("error reading mode: %v", err)
		}
	}

	if o.ModeCfg != nil {
		v := *o.ModeCfg

		if err = d.Set("mode_cfg", v); err != nil {
			return diag.Errorf("error reading mode_cfg: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Nattraversal != nil {
		v := *o.Nattraversal

		if err = d.Set("nattraversal", v); err != nil {
			return diag.Errorf("error reading nattraversal: %v", err)
		}
	}

	if o.NegotiateTimeout != nil {
		v := *o.NegotiateTimeout

		if err = d.Set("negotiate_timeout", v); err != nil {
			return diag.Errorf("error reading negotiate_timeout: %v", err)
		}
	}

	if o.NetworkId != nil {
		v := *o.NetworkId

		if err = d.Set("network_id", v); err != nil {
			return diag.Errorf("error reading network_id: %v", err)
		}
	}

	if o.NetworkOverlay != nil {
		v := *o.NetworkOverlay

		if err = d.Set("network_overlay", v); err != nil {
			return diag.Errorf("error reading network_overlay: %v", err)
		}
	}

	if o.NpuOffload != nil {
		v := *o.NpuOffload

		if err = d.Set("npu_offload", v); err != nil {
			return diag.Errorf("error reading npu_offload: %v", err)
		}
	}

	if o.Peer != nil {
		v := *o.Peer

		if err = d.Set("peer", v); err != nil {
			return diag.Errorf("error reading peer: %v", err)
		}
	}

	if o.Peergrp != nil {
		v := *o.Peergrp

		if err = d.Set("peergrp", v); err != nil {
			return diag.Errorf("error reading peergrp: %v", err)
		}
	}

	if o.Peerid != nil {
		v := *o.Peerid

		if err = d.Set("peerid", v); err != nil {
			return diag.Errorf("error reading peerid: %v", err)
		}
	}

	if o.Peertype != nil {
		v := *o.Peertype

		if err = d.Set("peertype", v); err != nil {
			return diag.Errorf("error reading peertype: %v", err)
		}
	}

	if o.Ppk != nil {
		v := *o.Ppk

		if err = d.Set("ppk", v); err != nil {
			return diag.Errorf("error reading ppk: %v", err)
		}
	}

	if o.PpkIdentity != nil {
		v := *o.PpkIdentity

		if err = d.Set("ppk_identity", v); err != nil {
			return diag.Errorf("error reading ppk_identity: %v", err)
		}
	}

	if o.PpkSecret != nil {
		v := *o.PpkSecret

		if v == "ENC XXXX" {
		} else if err = d.Set("ppk_secret", v); err != nil {
			return diag.Errorf("error reading ppk_secret: %v", err)
		}
	}

	if o.Priority != nil {
		v := *o.Priority

		if err = d.Set("priority", v); err != nil {
			return diag.Errorf("error reading priority: %v", err)
		}
	}

	if o.Proposal != nil {
		v := *o.Proposal

		if err = d.Set("proposal", v); err != nil {
			return diag.Errorf("error reading proposal: %v", err)
		}
	}

	if o.Psksecret != nil {
		v := *o.Psksecret

		if v == "ENC XXXX" {
		} else if err = d.Set("psksecret", v); err != nil {
			return diag.Errorf("error reading psksecret: %v", err)
		}
	}

	if o.PsksecretRemote != nil {
		v := *o.PsksecretRemote

		if v == "ENC XXXX" {
		} else if err = d.Set("psksecret_remote", v); err != nil {
			return diag.Errorf("error reading psksecret_remote: %v", err)
		}
	}

	if o.Reauth != nil {
		v := *o.Reauth

		if err = d.Set("reauth", v); err != nil {
			return diag.Errorf("error reading reauth: %v", err)
		}
	}

	if o.Rekey != nil {
		v := *o.Rekey

		if err = d.Set("rekey", v); err != nil {
			return diag.Errorf("error reading rekey: %v", err)
		}
	}

	if o.RemoteGw != nil {
		v := *o.RemoteGw

		if err = d.Set("remote_gw", v); err != nil {
			return diag.Errorf("error reading remote_gw: %v", err)
		}
	}

	if o.RemotegwDdns != nil {
		v := *o.RemotegwDdns

		if err = d.Set("remotegw_ddns", v); err != nil {
			return diag.Errorf("error reading remotegw_ddns: %v", err)
		}
	}

	if o.RsaSignatureFormat != nil {
		v := *o.RsaSignatureFormat

		if err = d.Set("rsa_signature_format", v); err != nil {
			return diag.Errorf("error reading rsa_signature_format: %v", err)
		}
	}

	if o.SavePassword != nil {
		v := *o.SavePassword

		if err = d.Set("save_password", v); err != nil {
			return diag.Errorf("error reading save_password: %v", err)
		}
	}

	if o.SendCertChain != nil {
		v := *o.SendCertChain

		if err = d.Set("send_cert_chain", v); err != nil {
			return diag.Errorf("error reading send_cert_chain: %v", err)
		}
	}

	if o.SignatureHashAlg != nil {
		v := *o.SignatureHashAlg

		if err = d.Set("signature_hash_alg", v); err != nil {
			return diag.Errorf("error reading signature_hash_alg: %v", err)
		}
	}

	if o.SplitIncludeService != nil {
		v := *o.SplitIncludeService

		if err = d.Set("split_include_service", v); err != nil {
			return diag.Errorf("error reading split_include_service: %v", err)
		}
	}

	if o.SuiteB != nil {
		v := *o.SuiteB

		if err = d.Set("suite_b", v); err != nil {
			return diag.Errorf("error reading suite_b: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	if o.UnitySupport != nil {
		v := *o.UnitySupport

		if err = d.Set("unity_support", v); err != nil {
			return diag.Errorf("error reading unity_support: %v", err)
		}
	}

	if o.Usrgrp != nil {
		v := *o.Usrgrp

		if err = d.Set("usrgrp", v); err != nil {
			return diag.Errorf("error reading usrgrp: %v", err)
		}
	}

	if o.WizardType != nil {
		v := *o.WizardType

		if err = d.Set("wizard_type", v); err != nil {
			return diag.Errorf("error reading wizard_type: %v", err)
		}
	}

	if o.Xauthtype != nil {
		v := *o.Xauthtype

		if err = d.Set("xauthtype", v); err != nil {
			return diag.Errorf("error reading xauthtype: %v", err)
		}
	}

	return nil
}

func expandVpnIpsecPhase1BackupGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnIpsecPhase1BackupGateway, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnIpsecPhase1BackupGateway

	for i := range l {
		tmp := models.VpnIpsecPhase1BackupGateway{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Address = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandVpnIpsecPhase1Certificate(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnIpsecPhase1Certificate, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnIpsecPhase1Certificate

	for i := range l {
		tmp := models.VpnIpsecPhase1Certificate{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandVpnIpsecPhase1Ipv4ExcludeRange(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnIpsecPhase1Ipv4ExcludeRange, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnIpsecPhase1Ipv4ExcludeRange

	for i := range l {
		tmp := models.VpnIpsecPhase1Ipv4ExcludeRange{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.end_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EndIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.start_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StartIp = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandVpnIpsecPhase1Ipv6ExcludeRange(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnIpsecPhase1Ipv6ExcludeRange, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnIpsecPhase1Ipv6ExcludeRange

	for i := range l {
		tmp := models.VpnIpsecPhase1Ipv6ExcludeRange{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.end_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EndIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.start_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StartIp = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectVpnIpsecPhase1(d *schema.ResourceData, sv string) (*models.VpnIpsecPhase1, diag.Diagnostics) {
	obj := models.VpnIpsecPhase1{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("acct_verify"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("acct_verify", sv)
				diags = append(diags, e)
			}
			obj.AcctVerify = &v2
		}
	}
	if v1, ok := d.GetOk("add_gw_route"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("add_gw_route", sv)
				diags = append(diags, e)
			}
			obj.AddGwRoute = &v2
		}
	}
	if v1, ok := d.GetOk("add_route"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("add_route", sv)
				diags = append(diags, e)
			}
			obj.AddRoute = &v2
		}
	}
	if v1, ok := d.GetOk("assign_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("assign_ip", sv)
				diags = append(diags, e)
			}
			obj.AssignIp = &v2
		}
	}
	if v1, ok := d.GetOk("assign_ip_from"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("assign_ip_from", sv)
				diags = append(diags, e)
			}
			obj.AssignIpFrom = &v2
		}
	}
	if v1, ok := d.GetOk("authmethod"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authmethod", sv)
				diags = append(diags, e)
			}
			obj.Authmethod = &v2
		}
	}
	if v1, ok := d.GetOk("authmethod_remote"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authmethod_remote", sv)
				diags = append(diags, e)
			}
			obj.AuthmethodRemote = &v2
		}
	}
	if v1, ok := d.GetOk("authpasswd"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authpasswd", sv)
				diags = append(diags, e)
			}
			obj.Authpasswd = &v2
		}
	}
	if v1, ok := d.GetOk("authusr"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authusr", sv)
				diags = append(diags, e)
			}
			obj.Authusr = &v2
		}
	}
	if v1, ok := d.GetOk("authusrgrp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authusrgrp", sv)
				diags = append(diags, e)
			}
			obj.Authusrgrp = &v2
		}
	}
	if v1, ok := d.GetOk("auto_negotiate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_negotiate", sv)
				diags = append(diags, e)
			}
			obj.AutoNegotiate = &v2
		}
	}
	if v, ok := d.GetOk("backup_gateway"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("backup_gateway", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnIpsecPhase1BackupGateway(d, v, "backup_gateway", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.BackupGateway = t
		}
	} else if d.HasChange("backup_gateway") {
		old, new := d.GetChange("backup_gateway")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.BackupGateway = &[]models.VpnIpsecPhase1BackupGateway{}
		}
	}
	if v1, ok := d.GetOk("banner"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("banner", sv)
				diags = append(diags, e)
			}
			obj.Banner = &v2
		}
	}
	if v1, ok := d.GetOk("cert_id_validation"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cert_id_validation", sv)
				diags = append(diags, e)
			}
			obj.CertIdValidation = &v2
		}
	}
	if v, ok := d.GetOk("certificate"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("certificate", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnIpsecPhase1Certificate(d, v, "certificate", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Certificate = t
		}
	} else if d.HasChange("certificate") {
		old, new := d.GetChange("certificate")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Certificate = &[]models.VpnIpsecPhase1Certificate{}
		}
	}
	if v1, ok := d.GetOk("childless_ike"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("childless_ike", sv)
				diags = append(diags, e)
			}
			obj.ChildlessIke = &v2
		}
	}
	if v1, ok := d.GetOk("client_auto_negotiate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("client_auto_negotiate", sv)
				diags = append(diags, e)
			}
			obj.ClientAutoNegotiate = &v2
		}
	}
	if v1, ok := d.GetOk("client_keep_alive"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("client_keep_alive", sv)
				diags = append(diags, e)
			}
			obj.ClientKeepAlive = &v2
		}
	}
	if v1, ok := d.GetOk("comments"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comments", sv)
				diags = append(diags, e)
			}
			obj.Comments = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_ra_giaddr"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp_ra_giaddr", sv)
				diags = append(diags, e)
			}
			obj.DhcpRaGiaddr = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp6_ra_linkaddr"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp6_ra_linkaddr", sv)
				diags = append(diags, e)
			}
			obj.Dhcp6RaLinkaddr = &v2
		}
	}
	if v1, ok := d.GetOk("dhgrp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhgrp", sv)
				diags = append(diags, e)
			}
			obj.Dhgrp = &v2
		}
	}
	if v1, ok := d.GetOk("digital_signature_auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("digital_signature_auth", sv)
				diags = append(diags, e)
			}
			obj.DigitalSignatureAuth = &v2
		}
	}
	if v1, ok := d.GetOk("distance"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("distance", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Distance = &tmp
		}
	}
	if v1, ok := d.GetOk("dns_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_mode", sv)
				diags = append(diags, e)
			}
			obj.DnsMode = &v2
		}
	}
	if v1, ok := d.GetOk("domain"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("domain", sv)
				diags = append(diags, e)
			}
			obj.Domain = &v2
		}
	}
	if v1, ok := d.GetOk("dpd"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dpd", sv)
				diags = append(diags, e)
			}
			obj.Dpd = &v2
		}
	}
	if v1, ok := d.GetOk("dpd_retrycount"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dpd_retrycount", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DpdRetrycount = &tmp
		}
	}
	if v1, ok := d.GetOk("dpd_retryinterval"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dpd_retryinterval", sv)
				diags = append(diags, e)
			}
			obj.DpdRetryinterval = &v2
		}
	}
	if v1, ok := d.GetOk("eap"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eap", sv)
				diags = append(diags, e)
			}
			obj.Eap = &v2
		}
	}
	if v1, ok := d.GetOk("eap_exclude_peergrp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eap_exclude_peergrp", sv)
				diags = append(diags, e)
			}
			obj.EapExcludePeergrp = &v2
		}
	}
	if v1, ok := d.GetOk("eap_identity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eap_identity", sv)
				diags = append(diags, e)
			}
			obj.EapIdentity = &v2
		}
	}
	if v1, ok := d.GetOk("enforce_unique_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("enforce_unique_id", sv)
				diags = append(diags, e)
			}
			obj.EnforceUniqueId = &v2
		}
	}
	if v1, ok := d.GetOk("esn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("esn", sv)
				diags = append(diags, e)
			}
			obj.Esn = &v2
		}
	}
	if v1, ok := d.GetOk("fec_base"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fec_base", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FecBase = &tmp
		}
	}
	if v1, ok := d.GetOk("fec_codec"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.6", "") {
				e := utils.AttributeVersionWarning("fec_codec", sv)
				diags = append(diags, e)
			}
			obj.FecCodec = &v2
		}
	}
	if v1, ok := d.GetOk("fec_egress"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fec_egress", sv)
				diags = append(diags, e)
			}
			obj.FecEgress = &v2
		}
	}
	if v1, ok := d.GetOk("fec_health_check"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("fec_health_check", sv)
				diags = append(diags, e)
			}
			obj.FecHealthCheck = &v2
		}
	}
	if v1, ok := d.GetOk("fec_ingress"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fec_ingress", sv)
				diags = append(diags, e)
			}
			obj.FecIngress = &v2
		}
	}
	if v1, ok := d.GetOk("fec_mapping_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("fec_mapping_profile", sv)
				diags = append(diags, e)
			}
			obj.FecMappingProfile = &v2
		}
	}
	if v1, ok := d.GetOk("fec_receive_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fec_receive_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FecReceiveTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("fec_redundant"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fec_redundant", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FecRedundant = &tmp
		}
	}
	if v1, ok := d.GetOk("fec_send_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fec_send_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FecSendTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("forticlient_enforcement"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("forticlient_enforcement", sv)
				diags = append(diags, e)
			}
			obj.ForticlientEnforcement = &v2
		}
	}
	if v1, ok := d.GetOk("fragmentation"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fragmentation", sv)
				diags = append(diags, e)
			}
			obj.Fragmentation = &v2
		}
	}
	if v1, ok := d.GetOk("fragmentation_mtu"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fragmentation_mtu", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FragmentationMtu = &tmp
		}
	}
	if v1, ok := d.GetOk("group_authentication"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("group_authentication", sv)
				diags = append(diags, e)
			}
			obj.GroupAuthentication = &v2
		}
	}
	if v1, ok := d.GetOk("group_authentication_secret"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("group_authentication_secret", sv)
				diags = append(diags, e)
			}
			obj.GroupAuthenticationSecret = &v2
		}
	}
	if v1, ok := d.GetOk("ha_sync_esp_seqno"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ha_sync_esp_seqno", sv)
				diags = append(diags, e)
			}
			obj.HaSyncEspSeqno = &v2
		}
	}
	if v1, ok := d.GetOk("idle_timeout"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("idle_timeout", sv)
				diags = append(diags, e)
			}
			obj.IdleTimeout = &v2
		}
	}
	if v1, ok := d.GetOk("idle_timeoutinterval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("idle_timeoutinterval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IdleTimeoutinterval = &tmp
		}
	}
	if v1, ok := d.GetOk("ike_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ike_version", sv)
				diags = append(diags, e)
			}
			obj.IkeVersion = &v2
		}
	}
	if v1, ok := d.GetOk("include_local_lan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("include_local_lan", sv)
				diags = append(diags, e)
			}
			obj.IncludeLocalLan = &v2
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("ip_delay_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("ip_delay_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IpDelayInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("ipv4_dns_server1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv4_dns_server1", sv)
				diags = append(diags, e)
			}
			obj.Ipv4DnsServer1 = &v2
		}
	}
	if v1, ok := d.GetOk("ipv4_dns_server2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv4_dns_server2", sv)
				diags = append(diags, e)
			}
			obj.Ipv4DnsServer2 = &v2
		}
	}
	if v1, ok := d.GetOk("ipv4_dns_server3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv4_dns_server3", sv)
				diags = append(diags, e)
			}
			obj.Ipv4DnsServer3 = &v2
		}
	}
	if v1, ok := d.GetOk("ipv4_end_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv4_end_ip", sv)
				diags = append(diags, e)
			}
			obj.Ipv4EndIp = &v2
		}
	}
	if v, ok := d.GetOk("ipv4_exclude_range"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ipv4_exclude_range", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnIpsecPhase1Ipv4ExcludeRange(d, v, "ipv4_exclude_range", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Ipv4ExcludeRange = t
		}
	} else if d.HasChange("ipv4_exclude_range") {
		old, new := d.GetChange("ipv4_exclude_range")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Ipv4ExcludeRange = &[]models.VpnIpsecPhase1Ipv4ExcludeRange{}
		}
	}
	if v1, ok := d.GetOk("ipv4_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv4_name", sv)
				diags = append(diags, e)
			}
			obj.Ipv4Name = &v2
		}
	}
	if v1, ok := d.GetOk("ipv4_netmask"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv4_netmask", sv)
				diags = append(diags, e)
			}
			obj.Ipv4Netmask = &v2
		}
	}
	if v1, ok := d.GetOk("ipv4_split_exclude"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv4_split_exclude", sv)
				diags = append(diags, e)
			}
			obj.Ipv4SplitExclude = &v2
		}
	}
	if v1, ok := d.GetOk("ipv4_split_include"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv4_split_include", sv)
				diags = append(diags, e)
			}
			obj.Ipv4SplitInclude = &v2
		}
	}
	if v1, ok := d.GetOk("ipv4_start_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv4_start_ip", sv)
				diags = append(diags, e)
			}
			obj.Ipv4StartIp = &v2
		}
	}
	if v1, ok := d.GetOk("ipv4_wins_server1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv4_wins_server1", sv)
				diags = append(diags, e)
			}
			obj.Ipv4WinsServer1 = &v2
		}
	}
	if v1, ok := d.GetOk("ipv4_wins_server2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv4_wins_server2", sv)
				diags = append(diags, e)
			}
			obj.Ipv4WinsServer2 = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6_dns_server1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_dns_server1", sv)
				diags = append(diags, e)
			}
			obj.Ipv6DnsServer1 = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6_dns_server2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_dns_server2", sv)
				diags = append(diags, e)
			}
			obj.Ipv6DnsServer2 = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6_dns_server3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_dns_server3", sv)
				diags = append(diags, e)
			}
			obj.Ipv6DnsServer3 = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6_end_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_end_ip", sv)
				diags = append(diags, e)
			}
			obj.Ipv6EndIp = &v2
		}
	}
	if v, ok := d.GetOk("ipv6_exclude_range"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ipv6_exclude_range", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnIpsecPhase1Ipv6ExcludeRange(d, v, "ipv6_exclude_range", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Ipv6ExcludeRange = t
		}
	} else if d.HasChange("ipv6_exclude_range") {
		old, new := d.GetChange("ipv6_exclude_range")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Ipv6ExcludeRange = &[]models.VpnIpsecPhase1Ipv6ExcludeRange{}
		}
	}
	if v1, ok := d.GetOk("ipv6_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_name", sv)
				diags = append(diags, e)
			}
			obj.Ipv6Name = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6_prefix"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_prefix", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Ipv6Prefix = &tmp
		}
	}
	if v1, ok := d.GetOk("ipv6_split_exclude"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_split_exclude", sv)
				diags = append(diags, e)
			}
			obj.Ipv6SplitExclude = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6_split_include"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_split_include", sv)
				diags = append(diags, e)
			}
			obj.Ipv6SplitInclude = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6_start_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_start_ip", sv)
				diags = append(diags, e)
			}
			obj.Ipv6StartIp = &v2
		}
	}
	if v1, ok := d.GetOk("keepalive"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("keepalive", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Keepalive = &tmp
		}
	}
	if v1, ok := d.GetOk("keylife"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("keylife", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Keylife = &tmp
		}
	}
	if v1, ok := d.GetOk("local_gw"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_gw", sv)
				diags = append(diags, e)
			}
			obj.LocalGw = &v2
		}
	}
	if v1, ok := d.GetOk("localid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("localid", sv)
				diags = append(diags, e)
			}
			obj.Localid = &v2
		}
	}
	if v1, ok := d.GetOk("localid_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("localid_type", sv)
				diags = append(diags, e)
			}
			obj.LocalidType = &v2
		}
	}
	if v1, ok := d.GetOk("loopback_asymroute"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "v7.0.2") {
				e := utils.AttributeVersionWarning("loopback_asymroute", sv)
				diags = append(diags, e)
			}
			obj.LoopbackAsymroute = &v2
		}
	}
	if v1, ok := d.GetOk("mesh_selector_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mesh_selector_type", sv)
				diags = append(diags, e)
			}
			obj.MeshSelectorType = &v2
		}
	}
	if v1, ok := d.GetOk("mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mode", sv)
				diags = append(diags, e)
			}
			obj.Mode = &v2
		}
	}
	if v1, ok := d.GetOk("mode_cfg"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mode_cfg", sv)
				diags = append(diags, e)
			}
			obj.ModeCfg = &v2
		}
	}
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v1, ok := d.GetOk("nattraversal"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("nattraversal", sv)
				diags = append(diags, e)
			}
			obj.Nattraversal = &v2
		}
	}
	if v1, ok := d.GetOk("negotiate_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("negotiate_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.NegotiateTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("network_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v7.0.2") {
				e := utils.AttributeVersionWarning("network_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.NetworkId = &tmp
		}
	}
	if v1, ok := d.GetOk("network_overlay"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.2") {
				e := utils.AttributeVersionWarning("network_overlay", sv)
				diags = append(diags, e)
			}
			obj.NetworkOverlay = &v2
		}
	}
	if v1, ok := d.GetOk("npu_offload"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("npu_offload", sv)
				diags = append(diags, e)
			}
			obj.NpuOffload = &v2
		}
	}
	if v1, ok := d.GetOk("peer"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("peer", sv)
				diags = append(diags, e)
			}
			obj.Peer = &v2
		}
	}
	if v1, ok := d.GetOk("peergrp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("peergrp", sv)
				diags = append(diags, e)
			}
			obj.Peergrp = &v2
		}
	}
	if v1, ok := d.GetOk("peerid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("peerid", sv)
				diags = append(diags, e)
			}
			obj.Peerid = &v2
		}
	}
	if v1, ok := d.GetOk("peertype"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("peertype", sv)
				diags = append(diags, e)
			}
			obj.Peertype = &v2
		}
	}
	if v1, ok := d.GetOk("ppk"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ppk", sv)
				diags = append(diags, e)
			}
			obj.Ppk = &v2
		}
	}
	if v1, ok := d.GetOk("ppk_identity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ppk_identity", sv)
				diags = append(diags, e)
			}
			obj.PpkIdentity = &v2
		}
	}
	if v1, ok := d.GetOk("ppk_secret"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ppk_secret", sv)
				diags = append(diags, e)
			}
			obj.PpkSecret = &v2
		}
	}
	if v1, ok := d.GetOk("priority"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("priority", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Priority = &tmp
		}
	}
	if v1, ok := d.GetOk("proposal"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("proposal", sv)
				diags = append(diags, e)
			}
			obj.Proposal = &v2
		}
	}
	if v1, ok := d.GetOk("psksecret"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("psksecret", sv)
				diags = append(diags, e)
			}
			obj.Psksecret = &v2
		}
	}
	if v1, ok := d.GetOk("psksecret_remote"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("psksecret_remote", sv)
				diags = append(diags, e)
			}
			obj.PsksecretRemote = &v2
		}
	}
	if v1, ok := d.GetOk("reauth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("reauth", sv)
				diags = append(diags, e)
			}
			obj.Reauth = &v2
		}
	}
	if v1, ok := d.GetOk("rekey"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rekey", sv)
				diags = append(diags, e)
			}
			obj.Rekey = &v2
		}
	}
	if v1, ok := d.GetOk("remote_gw"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("remote_gw", sv)
				diags = append(diags, e)
			}
			obj.RemoteGw = &v2
		}
	}
	if v1, ok := d.GetOk("remotegw_ddns"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("remotegw_ddns", sv)
				diags = append(diags, e)
			}
			obj.RemotegwDdns = &v2
		}
	}
	if v1, ok := d.GetOk("rsa_signature_format"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rsa_signature_format", sv)
				diags = append(diags, e)
			}
			obj.RsaSignatureFormat = &v2
		}
	}
	if v1, ok := d.GetOk("save_password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("save_password", sv)
				diags = append(diags, e)
			}
			obj.SavePassword = &v2
		}
	}
	if v1, ok := d.GetOk("send_cert_chain"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("send_cert_chain", sv)
				diags = append(diags, e)
			}
			obj.SendCertChain = &v2
		}
	}
	if v1, ok := d.GetOk("signature_hash_alg"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("signature_hash_alg", sv)
				diags = append(diags, e)
			}
			obj.SignatureHashAlg = &v2
		}
	}
	if v1, ok := d.GetOk("split_include_service"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("split_include_service", sv)
				diags = append(diags, e)
			}
			obj.SplitIncludeService = &v2
		}
	}
	if v1, ok := d.GetOk("suite_b"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("suite_b", sv)
				diags = append(diags, e)
			}
			obj.SuiteB = &v2
		}
	}
	if v1, ok := d.GetOk("type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("type", sv)
				diags = append(diags, e)
			}
			obj.Type = &v2
		}
	}
	if v1, ok := d.GetOk("unity_support"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("unity_support", sv)
				diags = append(diags, e)
			}
			obj.UnitySupport = &v2
		}
	}
	if v1, ok := d.GetOk("usrgrp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("usrgrp", sv)
				diags = append(diags, e)
			}
			obj.Usrgrp = &v2
		}
	}
	if v1, ok := d.GetOk("wizard_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wizard_type", sv)
				diags = append(diags, e)
			}
			obj.WizardType = &v2
		}
	}
	if v1, ok := d.GetOk("xauthtype"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("xauthtype", sv)
				diags = append(diags, e)
			}
			obj.Xauthtype = &v2
		}
	}
	return &obj, diags
}
