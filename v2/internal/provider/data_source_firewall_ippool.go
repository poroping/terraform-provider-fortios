// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv4 IP pools.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallIppool() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv4 IP pools.",

		ReadContext: dataSourceFirewallIppoolRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"add_nat64_route": {
				Type:        schema.TypeString,
				Description: "Enable/disable adding NAT64 route.",
				Computed:    true,
			},
			"arp_intf": {
				Type:        schema.TypeString,
				Description: "Select an interface from available options that will reply to ARP requests. (If blank, any is selected).",
				Computed:    true,
			},
			"arp_reply": {
				Type:        schema.TypeString,
				Description: "Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable).",
				Computed:    true,
			},
			"associated_interface": {
				Type:        schema.TypeString,
				Description: "Associated interface name.",
				Computed:    true,
			},
			"block_size": {
				Type:        schema.TypeInt,
				Description: "Number of addresses in a block (64 - 4096, default = 128).",
				Computed:    true,
			},
			"cgn_block_size": {
				Type:        schema.TypeInt,
				Description: "Number of ports in a block(64 to 4096 in unit of 64, default = 128).",
				Computed:    true,
			},
			"cgn_client_endip": {
				Type:        schema.TypeString,
				Description: "Final client IPv4 address (inclusive) (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).",
				Computed:    true,
			},
			"cgn_client_ipv6shift": {
				Type:        schema.TypeInt,
				Description: "IPv6 shift for fixed-allocation.(default 0)",
				Computed:    true,
			},
			"cgn_client_startip": {
				Type:        schema.TypeString,
				Description: "First client IPv4 address (inclusive) (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).",
				Computed:    true,
			},
			"cgn_fixedalloc": {
				Type:        schema.TypeString,
				Description: "Enable/disable fixed-allocation mode.",
				Computed:    true,
			},
			"cgn_overload": {
				Type:        schema.TypeString,
				Description: "Enable/disable overload mode.",
				Computed:    true,
			},
			"cgn_port_end": {
				Type:        schema.TypeInt,
				Description: "Ending public port can be allocated. ",
				Computed:    true,
			},
			"cgn_port_start": {
				Type:        schema.TypeInt,
				Description: "Starting public port can be allocated. ",
				Computed:    true,
			},
			"cgn_spa": {
				Type:        schema.TypeString,
				Description: "Enable/disable single port allocation mode.",
				Computed:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"endip": {
				Type:        schema.TypeString,
				Description: "Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).",
				Computed:    true,
			},
			"endport": {
				Type:        schema.TypeInt,
				Description: "Final port number (inclusive) in the range for the address pool (Default: 65533).",
				Computed:    true,
			},
			"exclude_ip": {
				Type:        schema.TypeList,
				Description: "Exclude IPs x.x.x.x.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet": {
							Type:        schema.TypeString,
							Description: "Exclude IPs",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "IP pool name.",
				Required:    true,
			},
			"nat64": {
				Type:        schema.TypeString,
				Description: "Enable/disable NAT64.",
				Computed:    true,
			},
			"num_blocks_per_user": {
				Type:        schema.TypeInt,
				Description: "Number of addresses blocks that can be used by a user (1 to 128, default = 8).",
				Computed:    true,
			},
			"pba_timeout": {
				Type:        schema.TypeInt,
				Description: "Port block allocation timeout (seconds).",
				Computed:    true,
			},
			"permit_any_host": {
				Type:        schema.TypeString,
				Description: "Enable/disable full cone NAT.",
				Computed:    true,
			},
			"port_per_user": {
				Type:        schema.TypeInt,
				Description: "Number of port for each user (32 - 60416, default = 0, which is auto).",
				Computed:    true,
			},
			"source_endip": {
				Type:        schema.TypeString,
				Description: "Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).",
				Computed:    true,
			},
			"source_startip": {
				Type:        schema.TypeString,
				Description: "First IPv4 address (inclusive) in the range of the source addresses to be translated (format = xxx.xxx.xxx.xxx, default = 0.0.0.0).",
				Computed:    true,
			},
			"startip": {
				Type:        schema.TypeString,
				Description: "First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).",
				Computed:    true,
			},
			"startport": {
				Type:        schema.TypeInt,
				Description: "First port number (inclusive) in the range for the address pool (Default: 5117).",
				Computed:    true,
			},
			"subnet_broadcast_in_ippool": {
				Type:        schema.TypeString,
				Description: "Enable/disable inclusion of the subnetwork address and broadcast IP address in the NAT64 IP pool.",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "IP pool type (overload, one-to-one, fixed port range, or port block allocation).",
				Computed:    true,
			},
			"utilization_alarm_clear": {
				Type:        schema.TypeInt,
				Description: "Pool utilization alarm clear threshold (40-100).",
				Computed:    true,
			},
			"utilization_alarm_raise": {
				Type:        schema.TypeInt,
				Description: "Pool utilization alarm raise threshold (50-100).",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallIppoolRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallIppool(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallIppool dataSource: %v", err)
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

	diags := refreshObjectFirewallIppool(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
