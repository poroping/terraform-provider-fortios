// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure DHCP servers.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemDhcpServer() *schema.Resource {
	return &schema.Resource{
		Description: "Configure DHCP servers.",

		ReadContext: dataSourceSystemDhcpServerRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"auto_configuration": {
				Type:        schema.TypeString,
				Description: "Enable/disable auto configuration.",
				Computed:    true,
			},
			"auto_managed_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of this DHCP server once this interface has been assigned an IP address from FortiIPAM.",
				Computed:    true,
			},
			"conflicted_ip_timeout": {
				Type:        schema.TypeInt,
				Description: "Time in seconds to wait after a conflicted IP address is removed from the DHCP range before it can be reused.",
				Computed:    true,
			},
			"ddns_auth": {
				Type:        schema.TypeString,
				Description: "DDNS authentication mode.",
				Computed:    true,
			},
			"ddns_key": {
				Type:        schema.TypeString,
				Description: "DDNS update key (base 64 encoding).",
				Computed:    true,
			},
			"ddns_keyname": {
				Type:        schema.TypeString,
				Description: "DDNS update key name.",
				Computed:    true,
			},
			"ddns_server_ip": {
				Type:        schema.TypeString,
				Description: "DDNS server IP.",
				Computed:    true,
			},
			"ddns_ttl": {
				Type:        schema.TypeInt,
				Description: "TTL.",
				Computed:    true,
			},
			"ddns_update": {
				Type:        schema.TypeString,
				Description: "Enable/disable DDNS update for DHCP.",
				Computed:    true,
			},
			"ddns_update_override": {
				Type:        schema.TypeString,
				Description: "Enable/disable DDNS update override for DHCP.",
				Computed:    true,
			},
			"ddns_zone": {
				Type:        schema.TypeString,
				Description: "Zone of your domain name (ex. DDNS.com).",
				Computed:    true,
			},
			"default_gateway": {
				Type:        schema.TypeString,
				Description: "Default gateway IP address assigned by the DHCP server.",
				Computed:    true,
			},
			"dhcp_settings_from_fortiipam": {
				Type:        schema.TypeString,
				Description: "Enable/disable populating of DHCP server settings from FortiIPAM.",
				Computed:    true,
			},
			"dns_server1": {
				Type:        schema.TypeString,
				Description: "DNS server 1.",
				Computed:    true,
			},
			"dns_server2": {
				Type:        schema.TypeString,
				Description: "DNS server 2.",
				Computed:    true,
			},
			"dns_server3": {
				Type:        schema.TypeString,
				Description: "DNS server 3.",
				Computed:    true,
			},
			"dns_server4": {
				Type:        schema.TypeString,
				Description: "DNS server 4.",
				Computed:    true,
			},
			"dns_service": {
				Type:        schema.TypeString,
				Description: "Options for assigning DNS servers to DHCP clients.",
				Computed:    true,
			},
			"domain": {
				Type:        schema.TypeString,
				Description: "Domain name suffix for the IP addresses that the DHCP server assigns to clients.",
				Computed:    true,
			},
			"exclude_range": {
				Type:        schema.TypeList,
				Description: "Exclude one or more ranges of IP addresses from being assigned to clients.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": {
							Type:        schema.TypeString,
							Description: "End of IP range.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"start_ip": {
							Type:        schema.TypeString,
							Description: "Start of IP range.",
							Computed:    true,
						},
					},
				},
			},
			"filename": {
				Type:        schema.TypeString,
				Description: "Name of the boot file on the TFTP server.",
				Computed:    true,
			},
			"forticlient_on_net_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiClient-On-Net service for this DHCP server.",
				Computed:    true,
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "ID.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "DHCP server can assign IP configurations to clients connected to this interface.",
				Computed:    true,
			},
			"ip_mode": {
				Type:        schema.TypeString,
				Description: "Method used to assign client IP.",
				Computed:    true,
			},
			"ip_range": {
				Type:        schema.TypeList,
				Description: "DHCP IP range configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": {
							Type:        schema.TypeString,
							Description: "End of IP range.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"start_ip": {
							Type:        schema.TypeString,
							Description: "Start of IP range.",
							Computed:    true,
						},
					},
				},
			},
			"ipsec_lease_hold": {
				Type:        schema.TypeInt,
				Description: "DHCP over IPsec leases expire this many seconds after tunnel down (0 to disable forced-expiry).",
				Computed:    true,
			},
			"lease_time": {
				Type:        schema.TypeInt,
				Description: "Lease time in seconds, 0 means unlimited.",
				Computed:    true,
			},
			"mac_acl_default_action": {
				Type:        schema.TypeString,
				Description: "MAC access control default action (allow or block assigning IP settings).",
				Computed:    true,
			},
			"netmask": {
				Type:        schema.TypeString,
				Description: "Netmask assigned by the DHCP server.",
				Computed:    true,
			},
			"next_server": {
				Type:        schema.TypeString,
				Description: "IP address of a server (for example, a TFTP sever) that DHCP clients can download a boot file from.",
				Computed:    true,
			},
			"ntp_server1": {
				Type:        schema.TypeString,
				Description: "NTP server 1.",
				Computed:    true,
			},
			"ntp_server2": {
				Type:        schema.TypeString,
				Description: "NTP server 2.",
				Computed:    true,
			},
			"ntp_server3": {
				Type:        schema.TypeString,
				Description: "NTP server 3.",
				Computed:    true,
			},
			"ntp_service": {
				Type:        schema.TypeString,
				Description: "Options for assigning Network Time Protocol (NTP) servers to DHCP clients.",
				Computed:    true,
			},
			"options": {
				Type:        schema.TypeList,
				Description: "DHCP options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type:        schema.TypeInt,
							Description: "DHCP option code.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"ip": {
							Type:        schema.TypeString,
							Description: "DHCP option IPs.",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "DHCP option type.",
							Computed:    true,
						},
						"value": {
							Type:        schema.TypeString,
							Description: "DHCP option value.",
							Computed:    true,
						},
					},
				},
			},
			"reserved_address": {
				Type:        schema.TypeList,
				Description: "Options for the DHCP server to assign IP settings to specific MAC addresses.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Options for the DHCP server to configure the client with the reserved MAC address.",
							Computed:    true,
						},
						"circuit_id": {
							Type:        schema.TypeString,
							Description: "Option 82 circuit-ID of the client that will get the reserved IP address.",
							Computed:    true,
						},
						"circuit_id_type": {
							Type:        schema.TypeString,
							Description: "DHCP option type.",
							Computed:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "Description.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"ip": {
							Type:        schema.TypeString,
							Description: "IP address to be reserved for the MAC address.",
							Computed:    true,
						},
						"mac": {
							Type:        schema.TypeString,
							Description: "MAC address of the client that will get the reserved IP address.",
							Computed:    true,
						},
						"remote_id": {
							Type:        schema.TypeString,
							Description: "Option 82 remote-ID of the client that will get the reserved IP address.",
							Computed:    true,
						},
						"remote_id_type": {
							Type:        schema.TypeString,
							Description: "DHCP option type.",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "DHCP reserved-address type.",
							Computed:    true,
						},
					},
				},
			},
			"server_type": {
				Type:        schema.TypeString,
				Description: "DHCP server can be a normal DHCP server or an IPsec DHCP server.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable this DHCP configuration.",
				Computed:    true,
			},
			"tftp_server": {
				Type:        schema.TypeList,
				Description: "One or more hostnames or IP addresses of the TFTP servers in quotes separated by spaces.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tftp_server": {
							Type:        schema.TypeString,
							Description: "TFTP server.",
							Computed:    true,
						},
					},
				},
			},
			"timezone": {
				Type:        schema.TypeString,
				Description: "Select the time zone to be assigned to DHCP clients.",
				Computed:    true,
			},
			"timezone_option": {
				Type:        schema.TypeString,
				Description: "Options for the DHCP server to set the client's time zone.",
				Computed:    true,
			},
			"vci_match": {
				Type:        schema.TypeString,
				Description: "Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served.",
				Computed:    true,
			},
			"vci_string": {
				Type:        schema.TypeList,
				Description: "One or more VCI strings in quotes separated by spaces.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vci_string": {
							Type:        schema.TypeString,
							Description: "VCI strings.",
							Computed:    true,
						},
					},
				},
			},
			"wifi_ac_service": {
				Type:        schema.TypeString,
				Description: "Options for assigning WiFi access controllers to DHCP clients.",
				Computed:    true,
			},
			"wifi_ac1": {
				Type:        schema.TypeString,
				Description: "WiFi Access Controller 1 IP address (DHCP option 138, RFC 5417).",
				Computed:    true,
			},
			"wifi_ac2": {
				Type:        schema.TypeString,
				Description: "WiFi Access Controller 2 IP address (DHCP option 138, RFC 5417).",
				Computed:    true,
			},
			"wifi_ac3": {
				Type:        schema.TypeString,
				Description: "WiFi Access Controller 3 IP address (DHCP option 138, RFC 5417).",
				Computed:    true,
			},
			"wins_server1": {
				Type:        schema.TypeString,
				Description: "WINS server 1.",
				Computed:    true,
			},
			"wins_server2": {
				Type:        schema.TypeString,
				Description: "WINS server 2.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemDhcpServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("fosid")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadSystemDhcpServer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemDhcpServer dataSource: %v", err)
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

	diags := refreshObjectSystemDhcpServer(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
