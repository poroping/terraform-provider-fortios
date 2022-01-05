// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure a forward server group consisting or multiple forward servers. Supports failover and load balancing.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWebProxyForwardServerGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Configure a forward server group consisting or multiple forward servers. Supports failover and load balancing.",

		ReadContext: dataSourceWebProxyForwardServerGroupRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"affinity": {
				Type:        schema.TypeString,
				Description: "Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global).",
				Computed:    true,
			},
			"group_down_option": {
				Type:        schema.TypeString,
				Description: "Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination.",
				Computed:    true,
			},
			"ldb_method": {
				Type:        schema.TypeString,
				Description: "Load balance method: weighted or least-session.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.",
				Required:    true,
			},
			"server_list": {
				Type:        schema.TypeList,
				Description: "Add web forward servers to a list to form a server group. Optionally assign weights to each server.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Forward server name.",
							Computed:    true,
						},
						"weight": {
							Type:        schema.TypeInt,
							Description: "Optionally assign a weight of the forwarding server for weighted load balancing (1 - 100, default = 10)",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWebProxyForwardServerGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWebProxyForwardServerGroup(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebProxyForwardServerGroup dataSource: %v", err)
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

	diags := refreshObjectWebProxyForwardServerGroup(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
