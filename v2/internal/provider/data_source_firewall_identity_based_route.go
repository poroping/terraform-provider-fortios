// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure identity based routing.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallIdentityBasedRoute() *schema.Resource {
	return &schema.Resource{
		Description: "Configure identity based routing.",

		ReadContext: dataSourceFirewallIdentityBasedRouteRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comments.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name.",
				Required:    true,
			},
			"rule": {
				Type:        schema.TypeList,
				Description: "Rule.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device": {
							Type:        schema.TypeString,
							Description: "Outgoing interface for the rule.",
							Computed:    true,
						},
						"gateway": {
							Type:        schema.TypeString,
							Description: "IPv4 address of the gateway (Format: xxx.xxx.xxx.xxx , Default: 0.0.0.0).",
							Computed:    true,
						},
						"groups": {
							Type:        schema.TypeList,
							Description: "Select one or more group(s) from available groups that are allowed to use this route. Separate group names with a space.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Group name.",
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Rule ID.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceFirewallIdentityBasedRouteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallIdentityBasedRoute(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallIdentityBasedRoute dataSource: %v", err)
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

	diags := refreshObjectFirewallIdentityBasedRoute(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
