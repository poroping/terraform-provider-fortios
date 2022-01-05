// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure DHCPv6 servers.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemDhcp6Server() *schema.Resource {
	return &schema.Resource{
		Description: "Configure DHCPv6 servers.",

		ReadContext: dataSourceSystemDhcp6ServerRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"delegated_prefix_iaid": {
				Type:        schema.TypeInt,
				Description: "IAID of obtained delegated-prefix from the upstream interface.",
				Computed:    true,
			},
			"dns_search_list": {
				Type:        schema.TypeString,
				Description: "DNS search list options.",
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
				Description: " Options for assigning DNS servers to DHCPv6 clients.",
				Computed:    true,
			},
			"domain": {
				Type:        schema.TypeString,
				Description: "Domain name suffix for the IP addresses that the DHCP server assigns to clients.",
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
			"lease_time": {
				Type:        schema.TypeInt,
				Description: "Lease time in seconds, 0 means unlimited.",
				Computed:    true,
			},
			"option1": {
				Type:        schema.TypeString,
				Description: "Option 1.",
				Computed:    true,
			},
			"option2": {
				Type:        schema.TypeString,
				Description: "Option 2.",
				Computed:    true,
			},
			"option3": {
				Type:        schema.TypeString,
				Description: "Option 3.",
				Computed:    true,
			},
			"prefix_mode": {
				Type:        schema.TypeString,
				Description: "Assigning a prefix from a DHCPv6 client or RA.",
				Computed:    true,
			},
			"prefix_range": {
				Type:        schema.TypeList,
				Description: "DHCP prefix configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_prefix": {
							Type:        schema.TypeString,
							Description: "End of prefix range.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"prefix_length": {
							Type:        schema.TypeInt,
							Description: "Prefix length.",
							Computed:    true,
						},
						"start_prefix": {
							Type:        schema.TypeString,
							Description: "Start of prefix range.",
							Computed:    true,
						},
					},
				},
			},
			"rapid_commit": {
				Type:        schema.TypeString,
				Description: "Enable/disable allow/disallow rapid commit.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable this DHCPv6 configuration.",
				Computed:    true,
			},
			"subnet": {
				Type:        schema.TypeString,
				Description: "Subnet or subnet-id if the IP mode is delegated.",
				Computed:    true,
			},
			"upstream_interface": {
				Type:        schema.TypeString,
				Description: "Interface name from where delegated information is provided.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemDhcp6ServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemDhcp6Server(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemDhcp6Server dataSource: %v", err)
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

	diags := refreshObjectSystemDhcp6Server(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
