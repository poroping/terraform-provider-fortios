// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv6 multicast.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceRouterMulticast6() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv6 multicast.",

		ReadContext: dataSourceRouterMulticast6Read,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"interface": {
				Type:        schema.TypeList,
				Description: "Protocol Independent Multicast (PIM) interfaces.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hello_holdtime": {
							Type:        schema.TypeInt,
							Description: "Time before old neighbor information expires in seconds (1 - 65535, default = 105).",
							Computed:    true,
						},
						"hello_interval": {
							Type:        schema.TypeInt,
							Description: "Interval between sending PIM hello messages in seconds (1 - 65535, default = 30).",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
					},
				},
			},
			"multicast_pmtu": {
				Type:        schema.TypeString,
				Description: "Enable/disable PMTU for IPv6 multicast.",
				Computed:    true,
			},
			"multicast_routing": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv6 multicast routing.",
				Computed:    true,
			},
			"pim_sm_global": {
				Type:        schema.TypeList,
				Description: "PIM sparse-mode global settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"register_rate_limit": {
							Type:        schema.TypeInt,
							Description: "Limit of packets/sec per source registered through this RP (0 means unlimited).",
							Computed:    true,
						},
						"rp_address": {
							Type:        schema.TypeList,
							Description: "Statically configured RP addresses.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Description: "ID of the entry.",
										Computed:    true,
									},
									"ip6_address": {
										Type:        schema.TypeString,
										Description: "RP router IPv6 address.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterMulticast6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "RouterMulticast6"

	o, err := c.Cmdb.ReadRouterMulticast6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterMulticast6 dataSource: %v", err)
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

	diags := refreshObjectRouterMulticast6(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
