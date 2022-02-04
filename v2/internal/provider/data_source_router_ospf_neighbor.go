// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: OSPF neighbor configuration are used when OSPF runs on non-broadcast media.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceRouterOspfNeighbor() *schema.Resource {
	return &schema.Resource{
		Description: "OSPF neighbor configuration are used when OSPF runs on non-broadcast media.",

		ReadContext: dataSourceRouterOspfNeighborRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"cost": {
				Type:        schema.TypeInt,
				Description: "Cost of the interface, value range from 0 to 65535, 0 means auto-cost.",
				Computed:    true,
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "Neighbor entry ID.",
				Computed:    true,
			},
			"ip": {
				Type:        schema.TypeString,
				Description: "Interface IP address of the neighbor.",
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
	}
}

func dataSourceRouterOspfNeighborRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterOspfNeighbor(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterOspfNeighbor dataSource: %v", err)
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

	diags := refreshObjectRouterOspfNeighbor(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
