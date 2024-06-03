// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure DHCP servers.

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

func resourceSystemDhcpServer() *schema.Resource {
	return &schema.Resource{
		Description: "Configure DHCP servers.",

		CreateContext: resourceSystemDhcpServerCreate,
		ReadContext:   resourceSystemDhcpServerRead,
		UpdateContext: resourceSystemDhcpServerUpdate,
		DeleteContext: resourceSystemDhcpServerDelete,

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
				Type:         schema.TypeBool,
				Description:  "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:     true,
				Default:      false,
				RequiredWith: []string{"fosid"},
			},
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"auto_configuration": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable auto configuration.",
				Optional:    true,
				Computed:    true,
			},
			"auto_managed_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable use of this DHCP server once this interface has been assigned an IP address from FortiIPAM.",
				Optional:    true,
				Computed:    true,
			},
			"conflicted_ip_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 8640000),

				Description: "Time in seconds to wait after a conflicted IP address is removed from the DHCP range before it can be reused.",
				Optional:    true,
				Computed:    true,
			},
			"ddns_auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "tsig"}, false),

				Description: "DDNS authentication mode.",
				Optional:    true,
				Computed:    true,
			},
			"ddns_key": {
				Type: schema.TypeString,

				Description: "DDNS update key (base 64 encoding).",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"ddns_keyname": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "DDNS update key name.",
				Optional:    true,
				Computed:    true,
			},
			"ddns_server_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "DDNS server IP.",
				Optional:    true,
				Computed:    true,
			},
			"ddns_ttl": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 86400),

				Description: "TTL.",
				Optional:    true,
				Computed:    true,
			},
			"ddns_update": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable DDNS update for DHCP.",
				Optional:    true,
				Computed:    true,
			},
			"ddns_update_override": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable DDNS update override for DHCP.",
				Optional:    true,
				Computed:    true,
			},
			"ddns_zone": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "Zone of your domain name (ex. DDNS.com).",
				Optional:    true,
				Computed:    true,
			},
			"default_gateway": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Default gateway IP address assigned by the DHCP server.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_settings_from_fortiipam": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable populating of DHCP server settings from FortiIPAM.",
				Optional:    true,
				Computed:    true,
			},
			"dns_server1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "DNS server 1.",
				Optional:    true,
				Computed:    true,
			},
			"dns_server2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "DNS server 2.",
				Optional:    true,
				Computed:    true,
			},
			"dns_server3": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "DNS server 3.",
				Optional:    true,
				Computed:    true,
			},
			"dns_server4": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "DNS server 4.",
				Optional:    true,
				Computed:    true,
			},
			"dns_service": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"local", "default", "specify"}, false),

				Description: "Options for assigning DNS servers to DHCP clients.",
				Optional:    true,
				Computed:    true,
			},
			"domain": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Domain name suffix for the IP addresses that the DHCP server assigns to clients.",
				Optional:    true,
				Computed:    true,
			},
			"exclude_range": {
				Type:        schema.TypeList,
				Description: "Exclude one or more ranges of IP addresses from being assigned to clients.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "End of IP range.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"lease_time": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(300, 8640000),

							Description: "Lease time in seconds, 0 means default lease time.",
							Optional:    true,
							Computed:    true,
						},
						"start_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Start of IP range.",
							Optional:    true,
							Computed:    true,
						},
						"uci_match": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable user class identifier (UCI) matching. When enabled only DHCP requests with a matching UCI are served with this range.",
							Optional:    true,
							Computed:    true,
						},
						"uci_string": {
							Type:        schema.TypeList,
							Description: "One or more UCI strings in quotes separated by spaces.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uci_string": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),

										Description: "UCI strings.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"vci_match": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this range.",
							Optional:    true,
							Computed:    true,
						},
						"vci_string": {
							Type:        schema.TypeList,
							Description: "One or more VCI strings in quotes separated by spaces.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vci_string": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),

										Description: "VCI strings.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"filename": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Name of the boot file on the TFTP server.",
				Optional:    true,
				Computed:    true,
			},
			"forticlient_on_net_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable FortiClient-On-Net service for this DHCP server.",
				Optional:    true,
				Computed:    true,
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "ID.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "DHCP server can assign IP configurations to clients connected to this interface.",
				Optional:    true,
				Computed:    true,
			},
			"ip_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"range", "usrgrp"}, false),

				Description: "Method used to assign client IP.",
				Optional:    true,
				Computed:    true,
			},
			"ip_range": {
				Type:        schema.TypeList,
				Description: "DHCP IP range configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "End of IP range.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"lease_time": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(300, 8640000),

							Description: "Lease time in seconds, 0 means default lease time.",
							Optional:    true,
							Computed:    true,
						},
						"start_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Start of IP range.",
							Optional:    true,
							Computed:    true,
						},
						"uci_match": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable user class identifier (UCI) matching. When enabled only DHCP requests with a matching UCI are served with this range.",
							Optional:    true,
							Computed:    true,
						},
						"uci_string": {
							Type:        schema.TypeList,
							Description: "One or more UCI strings in quotes separated by spaces.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uci_string": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),

										Description: "UCI strings.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"vci_match": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this range.",
							Optional:    true,
							Computed:    true,
						},
						"vci_string": {
							Type:        schema.TypeList,
							Description: "One or more VCI strings in quotes separated by spaces.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vci_string": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),

										Description: "VCI strings.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"ipsec_lease_hold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 8640000),

				Description: "DHCP over IPsec leases expire this many seconds after tunnel down (0 to disable forced-expiry).",
				Optional:    true,
				Computed:    true,
			},
			"lease_time": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 8640000),

				Description: "Lease time in seconds, 0 means unlimited.",
				Optional:    true,
				Computed:    true,
			},
			"mac_acl_default_action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"assign", "block"}, false),

				Description: "MAC access control default action (allow or block assigning IP settings).",
				Optional:    true,
				Computed:    true,
			},
			"netmask": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4Netmask,

				Description: "Netmask assigned by the DHCP server.",
				Optional:    true,
				Computed:    true,
			},
			"next_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address of a server (for example, a TFTP sever) that DHCP clients can download a boot file from.",
				Optional:    true,
				Computed:    true,
			},
			"ntp_server1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "NTP server 1.",
				Optional:    true,
				Computed:    true,
			},
			"ntp_server2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "NTP server 2.",
				Optional:    true,
				Computed:    true,
			},
			"ntp_server3": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "NTP server 3.",
				Optional:    true,
				Computed:    true,
			},
			"ntp_service": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"local", "default", "specify"}, false),

				Description: "Options for assigning Network Time Protocol (NTP) servers to DHCP clients.",
				Optional:    true,
				Computed:    true,
			},
			"options": {
				Type:        schema.TypeList,
				Description: "DHCP options.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "DHCP option code.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"ip": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffMultiStringEqual,
							Description:      "DHCP option IPs.",
							Optional:         true,
							Computed:         true,
						},
						"type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"hex", "string", "ip", "fqdn"}, false),

							Description: "DHCP option type.",
							Optional:    true,
							Computed:    true,
						},
						"uci_match": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable user class identifier (UCI) matching. When enabled only DHCP requests with a matching UCI are served with this option.",
							Optional:    true,
							Computed:    true,
						},
						"uci_string": {
							Type:        schema.TypeList,
							Description: "One or more UCI strings in quotes separated by spaces.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uci_string": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),

										Description: "UCI strings.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"value": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 312),

							Description: "DHCP option value.",
							Optional:    true,
							Computed:    true,
						},
						"vci_match": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this option.",
							Optional:    true,
							Computed:    true,
						},
						"vci_string": {
							Type:        schema.TypeList,
							Description: "One or more VCI strings in quotes separated by spaces.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vci_string": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),

										Description: "VCI strings.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"reserved_address": {
				Type:        schema.TypeList,
				Description: "Options for the DHCP server to assign IP settings to specific MAC addresses.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"assign", "block", "reserved"}, false),

							Description: "Options for the DHCP server to configure the client with the reserved MAC address.",
							Optional:    true,
							Computed:    true,
						},
						"circuit_id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 312),

							Description: "Option 82 circuit-ID of the client that will get the reserved IP address.",
							Optional:    true,
							Computed:    true,
						},
						"circuit_id_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"hex", "string"}, false),

							Description: "DHCP option type.",
							Optional:    true,
							Computed:    true,
						},
						"description": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Description.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "IP address to be reserved for the MAC address.",
							Optional:    true,
							Computed:    true,
						},
						"mac": {
							Type: schema.TypeString,

							Description: "MAC address of the client that will get the reserved IP address.",
							Optional:    true,
							Computed:    true,
						},
						"remote_id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 312),

							Description: "Option 82 remote-ID of the client that will get the reserved IP address.",
							Optional:    true,
							Computed:    true,
						},
						"remote_id_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"hex", "string"}, false),

							Description: "DHCP option type.",
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"mac", "option82"}, false),

							Description: "DHCP reserved-address type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"server_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"regular", "ipsec"}, false),

				Description: "DHCP server can be a normal DHCP server or an IPsec DHCP server.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable this DHCP configuration.",
				Optional:    true,
				Computed:    true,
			},
			"tftp_server": {
				Type:        schema.TypeList,
				Description: "One or more hostnames or IP addresses of the TFTP servers in quotes separated by spaces.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tftp_server": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "TFTP server.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"timezone": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"01", "02", "03", "04", "05", "81", "06", "07", "08", "09", "10", "11", "12", "13", "74", "14", "77", "15", "87", "16", "17", "18", "19", "20", "75", "21", "22", "23", "24", "80", "79", "25", "26", "27", "28", "78", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "83", "84", "40", "85", "39", "41", "42", "43", "44", "45", "46", "47", "51", "48", "49", "50", "52", "53", "54", "55", "56", "57", "58", "59", "60", "61", "62", "63", "64", "65", "66", "67", "68", "69", "70", "71", "72", "00", "82", "73", "86", "76"}, false),

				Description: "Select the time zone to be assigned to DHCP clients.",
				Optional:    true,
				Computed:    true,
			},
			"timezone_option": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "default", "specify"}, false),

				Description: "Options for the DHCP server to set the client's time zone.",
				Optional:    true,
				Computed:    true,
			},
			"vci_match": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served.",
				Optional:    true,
				Computed:    true,
			},
			"vci_string": {
				Type:        schema.TypeList,
				Description: "One or more VCI strings in quotes separated by spaces.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vci_string": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "VCI strings.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"wifi_ac_service": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"specify", "local"}, false),

				Description: "Options for assigning WiFi access controllers to DHCP clients.",
				Optional:    true,
				Computed:    true,
			},
			"wifi_ac1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "WiFi Access Controller 1 IP address (DHCP option 138, RFC 5417).",
				Optional:    true,
				Computed:    true,
			},
			"wifi_ac2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "WiFi Access Controller 2 IP address (DHCP option 138, RFC 5417).",
				Optional:    true,
				Computed:    true,
			},
			"wifi_ac3": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "WiFi Access Controller 3 IP address (DHCP option 138, RFC 5417).",
				Optional:    true,
				Computed:    true,
			},
			"wins_server1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "WINS server 1.",
				Optional:    true,
				Computed:    true,
			},
			"wins_server2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "WINS server 2.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemDhcpServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "id"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating SystemDhcpServer resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemDhcpServer(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemDhcpServer(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemDhcpServer")
	}

	return resourceSystemDhcpServerRead(ctx, d, meta)
}

func resourceSystemDhcpServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemDhcpServer(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemDhcpServer(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemDhcpServer resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemDhcpServer")
	}

	return resourceSystemDhcpServerRead(ctx, d, meta)
}

func resourceSystemDhcpServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemDhcpServer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemDhcpServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDhcpServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemDhcpServer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemDhcpServer resource: %v", err)
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

	diags := refreshObjectSystemDhcpServer(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemDhcpServerExcludeRange(d *schema.ResourceData, v *[]models.SystemDhcpServerExcludeRange, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.EndIp; tmp != nil {
				v["end_ip"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.LeaseTime; tmp != nil {
				v["lease_time"] = *tmp
			}

			if tmp := cfg.StartIp; tmp != nil {
				v["start_ip"] = *tmp
			}

			if tmp := cfg.UciMatch; tmp != nil {
				v["uci_match"] = *tmp
			}

			if tmp := cfg.UciString; tmp != nil {
				v["uci_string"] = flattenSystemDhcpServerExcludeRangeUciString(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "uci_string"), sort)
			}

			if tmp := cfg.VciMatch; tmp != nil {
				v["vci_match"] = *tmp
			}

			if tmp := cfg.VciString; tmp != nil {
				v["vci_string"] = flattenSystemDhcpServerExcludeRangeVciString(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "vci_string"), sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemDhcpServerExcludeRangeUciString(d *schema.ResourceData, v *[]models.SystemDhcpServerExcludeRangeUciString, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.UciString; tmp != nil {
				v["uci_string"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "uci_string")
	}

	return flat
}

func flattenSystemDhcpServerExcludeRangeVciString(d *schema.ResourceData, v *[]models.SystemDhcpServerExcludeRangeVciString, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.VciString; tmp != nil {
				v["vci_string"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "vci_string")
	}

	return flat
}

func flattenSystemDhcpServerIpRange(d *schema.ResourceData, v *[]models.SystemDhcpServerIpRange, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.EndIp; tmp != nil {
				v["end_ip"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.LeaseTime; tmp != nil {
				v["lease_time"] = *tmp
			}

			if tmp := cfg.StartIp; tmp != nil {
				v["start_ip"] = *tmp
			}

			if tmp := cfg.UciMatch; tmp != nil {
				v["uci_match"] = *tmp
			}

			if tmp := cfg.UciString; tmp != nil {
				v["uci_string"] = flattenSystemDhcpServerIpRangeUciString(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "uci_string"), sort)
			}

			if tmp := cfg.VciMatch; tmp != nil {
				v["vci_match"] = *tmp
			}

			if tmp := cfg.VciString; tmp != nil {
				v["vci_string"] = flattenSystemDhcpServerIpRangeVciString(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "vci_string"), sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemDhcpServerIpRangeUciString(d *schema.ResourceData, v *[]models.SystemDhcpServerIpRangeUciString, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.UciString; tmp != nil {
				v["uci_string"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "uci_string")
	}

	return flat
}

func flattenSystemDhcpServerIpRangeVciString(d *schema.ResourceData, v *[]models.SystemDhcpServerIpRangeVciString, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.VciString; tmp != nil {
				v["vci_string"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "vci_string")
	}

	return flat
}

func flattenSystemDhcpServerOptions(d *schema.ResourceData, v *[]models.SystemDhcpServerOptions, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Code; tmp != nil {
				v["code"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			if tmp := cfg.UciMatch; tmp != nil {
				v["uci_match"] = *tmp
			}

			if tmp := cfg.UciString; tmp != nil {
				v["uci_string"] = flattenSystemDhcpServerOptionsUciString(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "uci_string"), sort)
			}

			if tmp := cfg.Value; tmp != nil {
				v["value"] = *tmp
			}

			if tmp := cfg.VciMatch; tmp != nil {
				v["vci_match"] = *tmp
			}

			if tmp := cfg.VciString; tmp != nil {
				v["vci_string"] = flattenSystemDhcpServerOptionsVciString(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "vci_string"), sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemDhcpServerOptionsUciString(d *schema.ResourceData, v *[]models.SystemDhcpServerOptionsUciString, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.UciString; tmp != nil {
				v["uci_string"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "uci_string")
	}

	return flat
}

func flattenSystemDhcpServerOptionsVciString(d *schema.ResourceData, v *[]models.SystemDhcpServerOptionsVciString, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.VciString; tmp != nil {
				v["vci_string"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "vci_string")
	}

	return flat
}

func flattenSystemDhcpServerReservedAddress(d *schema.ResourceData, v *[]models.SystemDhcpServerReservedAddress, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.CircuitId; tmp != nil {
				v["circuit_id"] = *tmp
			}

			if tmp := cfg.CircuitIdType; tmp != nil {
				v["circuit_id_type"] = *tmp
			}

			if tmp := cfg.Description; tmp != nil {
				v["description"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			if tmp := cfg.Mac; tmp != nil {
				v["mac"] = *tmp
			}

			if tmp := cfg.RemoteId; tmp != nil {
				v["remote_id"] = *tmp
			}

			if tmp := cfg.RemoteIdType; tmp != nil {
				v["remote_id_type"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemDhcpServerTftpServer(d *schema.ResourceData, v *[]models.SystemDhcpServerTftpServer, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.TftpServer; tmp != nil {
				v["tftp_server"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "tftp_server")
	}

	return flat
}

func flattenSystemDhcpServerVciString(d *schema.ResourceData, v *[]models.SystemDhcpServerVciString, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.VciString; tmp != nil {
				v["vci_string"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "vci_string")
	}

	return flat
}

func refreshObjectSystemDhcpServer(d *schema.ResourceData, o *models.SystemDhcpServer, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AutoConfiguration != nil {
		v := *o.AutoConfiguration

		if err = d.Set("auto_configuration", v); err != nil {
			return diag.Errorf("error reading auto_configuration: %v", err)
		}
	}

	if o.AutoManagedStatus != nil {
		v := *o.AutoManagedStatus

		if err = d.Set("auto_managed_status", v); err != nil {
			return diag.Errorf("error reading auto_managed_status: %v", err)
		}
	}

	if o.ConflictedIpTimeout != nil {
		v := *o.ConflictedIpTimeout

		if err = d.Set("conflicted_ip_timeout", v); err != nil {
			return diag.Errorf("error reading conflicted_ip_timeout: %v", err)
		}
	}

	if o.DdnsAuth != nil {
		v := *o.DdnsAuth

		if err = d.Set("ddns_auth", v); err != nil {
			return diag.Errorf("error reading ddns_auth: %v", err)
		}
	}

	if o.DdnsKey != nil {
		v := *o.DdnsKey

		if v == "ENC XXXX" {
		} else if err = d.Set("ddns_key", v); err != nil {
			return diag.Errorf("error reading ddns_key: %v", err)
		}
	}

	if o.DdnsKeyname != nil {
		v := *o.DdnsKeyname

		if err = d.Set("ddns_keyname", v); err != nil {
			return diag.Errorf("error reading ddns_keyname: %v", err)
		}
	}

	if o.DdnsServerIp != nil {
		v := *o.DdnsServerIp

		if err = d.Set("ddns_server_ip", v); err != nil {
			return diag.Errorf("error reading ddns_server_ip: %v", err)
		}
	}

	if o.DdnsTtl != nil {
		v := *o.DdnsTtl

		if err = d.Set("ddns_ttl", v); err != nil {
			return diag.Errorf("error reading ddns_ttl: %v", err)
		}
	}

	if o.DdnsUpdate != nil {
		v := *o.DdnsUpdate

		if err = d.Set("ddns_update", v); err != nil {
			return diag.Errorf("error reading ddns_update: %v", err)
		}
	}

	if o.DdnsUpdateOverride != nil {
		v := *o.DdnsUpdateOverride

		if err = d.Set("ddns_update_override", v); err != nil {
			return diag.Errorf("error reading ddns_update_override: %v", err)
		}
	}

	if o.DdnsZone != nil {
		v := *o.DdnsZone

		if err = d.Set("ddns_zone", v); err != nil {
			return diag.Errorf("error reading ddns_zone: %v", err)
		}
	}

	if o.DefaultGateway != nil {
		v := *o.DefaultGateway

		if err = d.Set("default_gateway", v); err != nil {
			return diag.Errorf("error reading default_gateway: %v", err)
		}
	}

	if o.DhcpSettingsFromFortiipam != nil {
		v := *o.DhcpSettingsFromFortiipam

		if err = d.Set("dhcp_settings_from_fortiipam", v); err != nil {
			return diag.Errorf("error reading dhcp_settings_from_fortiipam: %v", err)
		}
	}

	if o.DnsServer1 != nil {
		v := *o.DnsServer1

		if err = d.Set("dns_server1", v); err != nil {
			return diag.Errorf("error reading dns_server1: %v", err)
		}
	}

	if o.DnsServer2 != nil {
		v := *o.DnsServer2

		if err = d.Set("dns_server2", v); err != nil {
			return diag.Errorf("error reading dns_server2: %v", err)
		}
	}

	if o.DnsServer3 != nil {
		v := *o.DnsServer3

		if err = d.Set("dns_server3", v); err != nil {
			return diag.Errorf("error reading dns_server3: %v", err)
		}
	}

	if o.DnsServer4 != nil {
		v := *o.DnsServer4

		if err = d.Set("dns_server4", v); err != nil {
			return diag.Errorf("error reading dns_server4: %v", err)
		}
	}

	if o.DnsService != nil {
		v := *o.DnsService

		if err = d.Set("dns_service", v); err != nil {
			return diag.Errorf("error reading dns_service: %v", err)
		}
	}

	if o.Domain != nil {
		v := *o.Domain

		if err = d.Set("domain", v); err != nil {
			return diag.Errorf("error reading domain: %v", err)
		}
	}

	if o.ExcludeRange != nil {
		if err = d.Set("exclude_range", flattenSystemDhcpServerExcludeRange(d, o.ExcludeRange, "exclude_range", sort)); err != nil {
			return diag.Errorf("error reading exclude_range: %v", err)
		}
	}

	if o.Filename != nil {
		v := *o.Filename

		if err = d.Set("filename", v); err != nil {
			return diag.Errorf("error reading filename: %v", err)
		}
	}

	if o.ForticlientOnNetStatus != nil {
		v := *o.ForticlientOnNetStatus

		if err = d.Set("forticlient_on_net_status", v); err != nil {
			return diag.Errorf("error reading forticlient_on_net_status: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.IpMode != nil {
		v := *o.IpMode

		if err = d.Set("ip_mode", v); err != nil {
			return diag.Errorf("error reading ip_mode: %v", err)
		}
	}

	if o.IpRange != nil {
		if err = d.Set("ip_range", flattenSystemDhcpServerIpRange(d, o.IpRange, "ip_range", sort)); err != nil {
			return diag.Errorf("error reading ip_range: %v", err)
		}
	}

	if o.IpsecLeaseHold != nil {
		v := *o.IpsecLeaseHold

		if err = d.Set("ipsec_lease_hold", v); err != nil {
			return diag.Errorf("error reading ipsec_lease_hold: %v", err)
		}
	}

	if o.LeaseTime != nil {
		v := *o.LeaseTime

		if err = d.Set("lease_time", v); err != nil {
			return diag.Errorf("error reading lease_time: %v", err)
		}
	}

	if o.MacAclDefaultAction != nil {
		v := *o.MacAclDefaultAction

		if err = d.Set("mac_acl_default_action", v); err != nil {
			return diag.Errorf("error reading mac_acl_default_action: %v", err)
		}
	}

	if o.Netmask != nil {
		v := *o.Netmask

		if err = d.Set("netmask", v); err != nil {
			return diag.Errorf("error reading netmask: %v", err)
		}
	}

	if o.NextServer != nil {
		v := *o.NextServer

		if err = d.Set("next_server", v); err != nil {
			return diag.Errorf("error reading next_server: %v", err)
		}
	}

	if o.NtpServer1 != nil {
		v := *o.NtpServer1

		if err = d.Set("ntp_server1", v); err != nil {
			return diag.Errorf("error reading ntp_server1: %v", err)
		}
	}

	if o.NtpServer2 != nil {
		v := *o.NtpServer2

		if err = d.Set("ntp_server2", v); err != nil {
			return diag.Errorf("error reading ntp_server2: %v", err)
		}
	}

	if o.NtpServer3 != nil {
		v := *o.NtpServer3

		if err = d.Set("ntp_server3", v); err != nil {
			return diag.Errorf("error reading ntp_server3: %v", err)
		}
	}

	if o.NtpService != nil {
		v := *o.NtpService

		if err = d.Set("ntp_service", v); err != nil {
			return diag.Errorf("error reading ntp_service: %v", err)
		}
	}

	if o.Options != nil {
		if err = d.Set("options", flattenSystemDhcpServerOptions(d, o.Options, "options", sort)); err != nil {
			return diag.Errorf("error reading options: %v", err)
		}
	}

	if o.ReservedAddress != nil {
		if err = d.Set("reserved_address", flattenSystemDhcpServerReservedAddress(d, o.ReservedAddress, "reserved_address", sort)); err != nil {
			return diag.Errorf("error reading reserved_address: %v", err)
		}
	}

	if o.ServerType != nil {
		v := *o.ServerType

		if err = d.Set("server_type", v); err != nil {
			return diag.Errorf("error reading server_type: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.TftpServer != nil {
		if err = d.Set("tftp_server", flattenSystemDhcpServerTftpServer(d, o.TftpServer, "tftp_server", sort)); err != nil {
			return diag.Errorf("error reading tftp_server: %v", err)
		}
	}

	if o.Timezone != nil {
		v := *o.Timezone

		if err = d.Set("timezone", v); err != nil {
			return diag.Errorf("error reading timezone: %v", err)
		}
	}

	if o.TimezoneOption != nil {
		v := *o.TimezoneOption

		if err = d.Set("timezone_option", v); err != nil {
			return diag.Errorf("error reading timezone_option: %v", err)
		}
	}

	if o.VciMatch != nil {
		v := *o.VciMatch

		if err = d.Set("vci_match", v); err != nil {
			return diag.Errorf("error reading vci_match: %v", err)
		}
	}

	if o.VciString != nil {
		if err = d.Set("vci_string", flattenSystemDhcpServerVciString(d, o.VciString, "vci_string", sort)); err != nil {
			return diag.Errorf("error reading vci_string: %v", err)
		}
	}

	if o.WifiAcService != nil {
		v := *o.WifiAcService

		if err = d.Set("wifi_ac_service", v); err != nil {
			return diag.Errorf("error reading wifi_ac_service: %v", err)
		}
	}

	if o.WifiAc1 != nil {
		v := *o.WifiAc1

		if err = d.Set("wifi_ac1", v); err != nil {
			return diag.Errorf("error reading wifi_ac1: %v", err)
		}
	}

	if o.WifiAc2 != nil {
		v := *o.WifiAc2

		if err = d.Set("wifi_ac2", v); err != nil {
			return diag.Errorf("error reading wifi_ac2: %v", err)
		}
	}

	if o.WifiAc3 != nil {
		v := *o.WifiAc3

		if err = d.Set("wifi_ac3", v); err != nil {
			return diag.Errorf("error reading wifi_ac3: %v", err)
		}
	}

	if o.WinsServer1 != nil {
		v := *o.WinsServer1

		if err = d.Set("wins_server1", v); err != nil {
			return diag.Errorf("error reading wins_server1: %v", err)
		}
	}

	if o.WinsServer2 != nil {
		v := *o.WinsServer2

		if err = d.Set("wins_server2", v); err != nil {
			return diag.Errorf("error reading wins_server2: %v", err)
		}
	}

	return nil
}

func expandSystemDhcpServerExcludeRange(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemDhcpServerExcludeRange, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemDhcpServerExcludeRange

	for i := range l {
		tmp := models.SystemDhcpServerExcludeRange{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.end_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EndIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.lease_time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.LeaseTime = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.start_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StartIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uci_match", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UciMatch = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uci_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemDhcpServerExcludeRangeUciString(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemDhcpServerExcludeRangeUciString
			// 	}
			tmp.UciString = v2

		}

		pre_append = fmt.Sprintf("%s.%d.vci_match", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VciMatch = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vci_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemDhcpServerExcludeRangeVciString(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemDhcpServerExcludeRangeVciString
			// 	}
			tmp.VciString = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemDhcpServerExcludeRangeUciString(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemDhcpServerExcludeRangeUciString, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemDhcpServerExcludeRangeUciString

	for i := range l {
		tmp := models.SystemDhcpServerExcludeRangeUciString{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.uci_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UciString = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemDhcpServerExcludeRangeVciString(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemDhcpServerExcludeRangeVciString, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemDhcpServerExcludeRangeVciString

	for i := range l {
		tmp := models.SystemDhcpServerExcludeRangeVciString{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.vci_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VciString = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemDhcpServerIpRange(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemDhcpServerIpRange, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemDhcpServerIpRange

	for i := range l {
		tmp := models.SystemDhcpServerIpRange{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.end_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EndIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.lease_time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.LeaseTime = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.start_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StartIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uci_match", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UciMatch = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uci_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemDhcpServerIpRangeUciString(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemDhcpServerIpRangeUciString
			// 	}
			tmp.UciString = v2

		}

		pre_append = fmt.Sprintf("%s.%d.vci_match", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VciMatch = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vci_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemDhcpServerIpRangeVciString(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemDhcpServerIpRangeVciString
			// 	}
			tmp.VciString = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemDhcpServerIpRangeUciString(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemDhcpServerIpRangeUciString, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemDhcpServerIpRangeUciString

	for i := range l {
		tmp := models.SystemDhcpServerIpRangeUciString{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.uci_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UciString = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemDhcpServerIpRangeVciString(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemDhcpServerIpRangeVciString, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemDhcpServerIpRangeVciString

	for i := range l {
		tmp := models.SystemDhcpServerIpRangeVciString{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.vci_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VciString = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemDhcpServerOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemDhcpServerOptions, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemDhcpServerOptions

	for i := range l {
		tmp := models.SystemDhcpServerOptions{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.code", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Code = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uci_match", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UciMatch = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uci_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemDhcpServerOptionsUciString(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemDhcpServerOptionsUciString
			// 	}
			tmp.UciString = v2

		}

		pre_append = fmt.Sprintf("%s.%d.value", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Value = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vci_match", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VciMatch = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vci_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemDhcpServerOptionsVciString(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemDhcpServerOptionsVciString
			// 	}
			tmp.VciString = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemDhcpServerOptionsUciString(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemDhcpServerOptionsUciString, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemDhcpServerOptionsUciString

	for i := range l {
		tmp := models.SystemDhcpServerOptionsUciString{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.uci_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UciString = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemDhcpServerOptionsVciString(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemDhcpServerOptionsVciString, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemDhcpServerOptionsVciString

	for i := range l {
		tmp := models.SystemDhcpServerOptionsVciString{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.vci_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VciString = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemDhcpServerReservedAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemDhcpServerReservedAddress, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemDhcpServerReservedAddress

	for i := range l {
		tmp := models.SystemDhcpServerReservedAddress{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.circuit_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CircuitId = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.circuit_id_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CircuitIdType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.description", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Description = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mac", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mac = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.remote_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RemoteId = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.remote_id_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RemoteIdType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemDhcpServerTftpServer(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemDhcpServerTftpServer, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemDhcpServerTftpServer

	for i := range l {
		tmp := models.SystemDhcpServerTftpServer{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.tftp_server", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TftpServer = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemDhcpServerVciString(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemDhcpServerVciString, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemDhcpServerVciString

	for i := range l {
		tmp := models.SystemDhcpServerVciString{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.vci_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VciString = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemDhcpServer(d *schema.ResourceData, sv string) (*models.SystemDhcpServer, diag.Diagnostics) {
	obj := models.SystemDhcpServer{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("auto_configuration"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_configuration", sv)
				diags = append(diags, e)
			}
			obj.AutoConfiguration = &v2
		}
	}
	if v1, ok := d.GetOk("auto_managed_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("auto_managed_status", sv)
				diags = append(diags, e)
			}
			obj.AutoManagedStatus = &v2
		}
	}
	if v1, ok := d.GetOk("conflicted_ip_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("conflicted_ip_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ConflictedIpTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("ddns_auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ddns_auth", sv)
				diags = append(diags, e)
			}
			obj.DdnsAuth = &v2
		}
	}
	if v1, ok := d.GetOk("ddns_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ddns_key", sv)
				diags = append(diags, e)
			}
			obj.DdnsKey = &v2
		}
	}
	if v1, ok := d.GetOk("ddns_keyname"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ddns_keyname", sv)
				diags = append(diags, e)
			}
			obj.DdnsKeyname = &v2
		}
	}
	if v1, ok := d.GetOk("ddns_server_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ddns_server_ip", sv)
				diags = append(diags, e)
			}
			obj.DdnsServerIp = &v2
		}
	}
	if v1, ok := d.GetOk("ddns_ttl"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ddns_ttl", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DdnsTtl = &tmp
		}
	}
	if v1, ok := d.GetOk("ddns_update"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ddns_update", sv)
				diags = append(diags, e)
			}
			obj.DdnsUpdate = &v2
		}
	}
	if v1, ok := d.GetOk("ddns_update_override"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ddns_update_override", sv)
				diags = append(diags, e)
			}
			obj.DdnsUpdateOverride = &v2
		}
	}
	if v1, ok := d.GetOk("ddns_zone"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ddns_zone", sv)
				diags = append(diags, e)
			}
			obj.DdnsZone = &v2
		}
	}
	if v1, ok := d.GetOk("default_gateway"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_gateway", sv)
				diags = append(diags, e)
			}
			obj.DefaultGateway = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_settings_from_fortiipam"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("dhcp_settings_from_fortiipam", sv)
				diags = append(diags, e)
			}
			obj.DhcpSettingsFromFortiipam = &v2
		}
	}
	if v1, ok := d.GetOk("dns_server1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_server1", sv)
				diags = append(diags, e)
			}
			obj.DnsServer1 = &v2
		}
	}
	if v1, ok := d.GetOk("dns_server2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_server2", sv)
				diags = append(diags, e)
			}
			obj.DnsServer2 = &v2
		}
	}
	if v1, ok := d.GetOk("dns_server3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_server3", sv)
				diags = append(diags, e)
			}
			obj.DnsServer3 = &v2
		}
	}
	if v1, ok := d.GetOk("dns_server4"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_server4", sv)
				diags = append(diags, e)
			}
			obj.DnsServer4 = &v2
		}
	}
	if v1, ok := d.GetOk("dns_service"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_service", sv)
				diags = append(diags, e)
			}
			obj.DnsService = &v2
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
	if v, ok := d.GetOk("exclude_range"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("exclude_range", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemDhcpServerExcludeRange(d, v, "exclude_range", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ExcludeRange = t
		}
	} else if d.HasChange("exclude_range") {
		old, new := d.GetChange("exclude_range")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ExcludeRange = &[]models.SystemDhcpServerExcludeRange{}
		}
	}
	if v1, ok := d.GetOk("filename"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("filename", sv)
				diags = append(diags, e)
			}
			obj.Filename = &v2
		}
	}
	if v1, ok := d.GetOk("forticlient_on_net_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("forticlient_on_net_status", sv)
				diags = append(diags, e)
			}
			obj.ForticlientOnNetStatus = &v2
		}
	}
	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Id = &tmp
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
	if v1, ok := d.GetOk("ip_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip_mode", sv)
				diags = append(diags, e)
			}
			obj.IpMode = &v2
		}
	}
	if v, ok := d.GetOk("ip_range"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ip_range", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemDhcpServerIpRange(d, v, "ip_range", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.IpRange = t
		}
	} else if d.HasChange("ip_range") {
		old, new := d.GetChange("ip_range")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.IpRange = &[]models.SystemDhcpServerIpRange{}
		}
	}
	if v1, ok := d.GetOk("ipsec_lease_hold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipsec_lease_hold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IpsecLeaseHold = &tmp
		}
	}
	if v1, ok := d.GetOk("lease_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lease_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LeaseTime = &tmp
		}
	}
	if v1, ok := d.GetOk("mac_acl_default_action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mac_acl_default_action", sv)
				diags = append(diags, e)
			}
			obj.MacAclDefaultAction = &v2
		}
	}
	if v1, ok := d.GetOk("netmask"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("netmask", sv)
				diags = append(diags, e)
			}
			obj.Netmask = &v2
		}
	}
	if v1, ok := d.GetOk("next_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("next_server", sv)
				diags = append(diags, e)
			}
			obj.NextServer = &v2
		}
	}
	if v1, ok := d.GetOk("ntp_server1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ntp_server1", sv)
				diags = append(diags, e)
			}
			obj.NtpServer1 = &v2
		}
	}
	if v1, ok := d.GetOk("ntp_server2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ntp_server2", sv)
				diags = append(diags, e)
			}
			obj.NtpServer2 = &v2
		}
	}
	if v1, ok := d.GetOk("ntp_server3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ntp_server3", sv)
				diags = append(diags, e)
			}
			obj.NtpServer3 = &v2
		}
	}
	if v1, ok := d.GetOk("ntp_service"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ntp_service", sv)
				diags = append(diags, e)
			}
			obj.NtpService = &v2
		}
	}
	if v, ok := d.GetOk("options"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("options", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemDhcpServerOptions(d, v, "options", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Options = t
		}
	} else if d.HasChange("options") {
		old, new := d.GetChange("options")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Options = &[]models.SystemDhcpServerOptions{}
		}
	}
	if v, ok := d.GetOk("reserved_address"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("reserved_address", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemDhcpServerReservedAddress(d, v, "reserved_address", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ReservedAddress = t
		}
	} else if d.HasChange("reserved_address") {
		old, new := d.GetChange("reserved_address")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ReservedAddress = &[]models.SystemDhcpServerReservedAddress{}
		}
	}
	if v1, ok := d.GetOk("server_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_type", sv)
				diags = append(diags, e)
			}
			obj.ServerType = &v2
		}
	}
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	if v, ok := d.GetOk("tftp_server"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("tftp_server", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemDhcpServerTftpServer(d, v, "tftp_server", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.TftpServer = t
		}
	} else if d.HasChange("tftp_server") {
		old, new := d.GetChange("tftp_server")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.TftpServer = &[]models.SystemDhcpServerTftpServer{}
		}
	}
	if v1, ok := d.GetOk("timezone"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("timezone", sv)
				diags = append(diags, e)
			}
			obj.Timezone = &v2
		}
	}
	if v1, ok := d.GetOk("timezone_option"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("timezone_option", sv)
				diags = append(diags, e)
			}
			obj.TimezoneOption = &v2
		}
	}
	if v1, ok := d.GetOk("vci_match"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vci_match", sv)
				diags = append(diags, e)
			}
			obj.VciMatch = &v2
		}
	}
	if v, ok := d.GetOk("vci_string"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("vci_string", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemDhcpServerVciString(d, v, "vci_string", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.VciString = t
		}
	} else if d.HasChange("vci_string") {
		old, new := d.GetChange("vci_string")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.VciString = &[]models.SystemDhcpServerVciString{}
		}
	}
	if v1, ok := d.GetOk("wifi_ac_service"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wifi_ac_service", sv)
				diags = append(diags, e)
			}
			obj.WifiAcService = &v2
		}
	}
	if v1, ok := d.GetOk("wifi_ac1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wifi_ac1", sv)
				diags = append(diags, e)
			}
			obj.WifiAc1 = &v2
		}
	}
	if v1, ok := d.GetOk("wifi_ac2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wifi_ac2", sv)
				diags = append(diags, e)
			}
			obj.WifiAc2 = &v2
		}
	}
	if v1, ok := d.GetOk("wifi_ac3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wifi_ac3", sv)
				diags = append(diags, e)
			}
			obj.WifiAc3 = &v2
		}
	}
	if v1, ok := d.GetOk("wins_server1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wins_server1", sv)
				diags = append(diags, e)
			}
			obj.WinsServer1 = &v2
		}
	}
	if v1, ok := d.GetOk("wins_server2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wins_server2", sv)
				diags = append(diags, e)
			}
			obj.WinsServer2 = &v2
		}
	}
	return &obj, diags
}
