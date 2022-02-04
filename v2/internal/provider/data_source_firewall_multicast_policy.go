// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure multicast NAT policies.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallMulticastPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure multicast NAT policies.",

		ReadContext: dataSourceFirewallMulticastPolicyRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"action": {
				Type:        schema.TypeString,
				Description: "Accept or deny traffic matching the policy.",
				Computed:    true,
			},
			"auto_asic_offload": {
				Type:        schema.TypeString,
				Description: "Enable/disable offloading policy traffic for hardware acceleration.",
				Computed:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"dnat": {
				Type:        schema.TypeString,
				Description: "IPv4 DNAT address used for multicast destination addresses.",
				Computed:    true,
			},
			"dstaddr": {
				Type:        schema.TypeList,
				Description: "Destination address objects.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Destination address objects.",
							Computed:    true,
						},
					},
				},
			},
			"dstintf": {
				Type:        schema.TypeString,
				Description: "Destination interface name.",
				Computed:    true,
			},
			"end_port": {
				Type:        schema.TypeInt,
				Description: "Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).",
				Computed:    true,
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "Policy ID ((0 - 4294967294).",
				Computed:    true,
			},
			"logtraffic": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging traffic accepted by this policy.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Policy name.",
				Computed:    true,
			},
			"protocol": {
				Type:        schema.TypeInt,
				Description: "Integer value for the protocol type as defined by IANA (0 - 255, default = 0).",
				Computed:    true,
			},
			"snat": {
				Type:        schema.TypeString,
				Description: "Enable/disable substitution of the outgoing interface IP address for the original source IP address (called source NAT or SNAT).",
				Computed:    true,
			},
			"snat_ip": {
				Type:        schema.TypeString,
				Description: "IPv4 address to be used as the source address for NATed traffic.",
				Computed:    true,
			},
			"srcaddr": {
				Type:        schema.TypeList,
				Description: "Source address objects.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Source address objects.",
							Computed:    true,
						},
					},
				},
			},
			"srcintf": {
				Type:        schema.TypeString,
				Description: "Source interface name.",
				Computed:    true,
			},
			"start_port": {
				Type:        schema.TypeInt,
				Description: "Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable this policy.",
				Computed:    true,
			},
			"uuid": {
				Type:        schema.TypeString,
				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallMulticastPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallMulticastPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallMulticastPolicy dataSource: %v", err)
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

	diags := refreshObjectFirewallMulticastPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
