// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: OSPF6 interface configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceRouterOspf6Ospf6Interface() *schema.Resource {
	return &schema.Resource{
		Description: "OSPF6 interface configuration.",

		ReadContext: dataSourceRouterOspf6Ospf6InterfaceRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"area_id": {
				Type:        schema.TypeString,
				Description: "A.B.C.D, in IPv4 address format.",
				Computed:    true,
			},
			"authentication": {
				Type:        schema.TypeString,
				Description: "Authentication mode.",
				Computed:    true,
			},
			"bfd": {
				Type:        schema.TypeString,
				Description: "Enable/disable Bidirectional Forwarding Detection (BFD).",
				Computed:    true,
			},
			"cost": {
				Type:        schema.TypeInt,
				Description: "Cost of the interface, value range from 0 to 65535, 0 means auto-cost.",
				Computed:    true,
			},
			"dead_interval": {
				Type:        schema.TypeInt,
				Description: "Dead interval.",
				Computed:    true,
			},
			"hello_interval": {
				Type:        schema.TypeInt,
				Description: "Hello interval.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Configuration interface name.",
				Computed:    true,
			},
			"ipsec_auth_alg": {
				Type:        schema.TypeString,
				Description: "Authentication algorithm.",
				Computed:    true,
			},
			"ipsec_enc_alg": {
				Type:        schema.TypeString,
				Description: "Encryption algorithm.",
				Computed:    true,
			},
			"ipsec_keys": {
				Type:        schema.TypeList,
				Description: "IPsec authentication and encryption keys.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_key": {
							Type:        schema.TypeString,
							Description: "Authentication key.",
							Computed:    true,
							Sensitive:   true,
						},
						"enc_key": {
							Type:        schema.TypeString,
							Description: "Encryption key.",
							Computed:    true,
							Sensitive:   true,
						},
						"spi": {
							Type:        schema.TypeInt,
							Description: "Security Parameters Index.",
							Computed:    true,
						},
					},
				},
			},
			"key_rollover_interval": {
				Type:        schema.TypeInt,
				Description: "Key roll-over interval.",
				Computed:    true,
			},
			"mtu": {
				Type:        schema.TypeInt,
				Description: "MTU for OSPFv3 packets.",
				Computed:    true,
			},
			"mtu_ignore": {
				Type:        schema.TypeString,
				Description: "Enable/disable ignoring MTU field in DBD packets.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Interface entry name.",
				Required:    true,
			},
			"neighbor": {
				Type:        schema.TypeList,
				Description: "OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cost": {
							Type:        schema.TypeInt,
							Description: "Cost of the interface, value range from 0 to 65535, 0 means auto-cost.",
							Computed:    true,
						},
						"ip6": {
							Type:        schema.TypeString,
							Description: "IPv6 link local address of the neighbor.",
							Computed:    true,
						},
						"poll_interval": {
							Type:        schema.TypeInt,
							Description: "Poll interval time in seconds.",
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "Priority.",
							Computed:    true,
						},
					},
				},
			},
			"network_type": {
				Type:        schema.TypeString,
				Description: "Network type.",
				Computed:    true,
			},
			"priority": {
				Type:        schema.TypeInt,
				Description: "Priority.",
				Computed:    true,
			},
			"retransmit_interval": {
				Type:        schema.TypeInt,
				Description: "Retransmit interval.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable OSPF6 routing on this interface.",
				Computed:    true,
			},
			"transmit_delay": {
				Type:        schema.TypeInt,
				Description: "Transmit delay.",
				Computed:    true,
			},
		},
	}
}

func dataSourceRouterOspf6Ospf6InterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterOspf6Ospf6Interface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterOspf6Ospf6Interface dataSource: %v", err)
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

	diags := refreshObjectRouterOspf6Ospf6Interface(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
