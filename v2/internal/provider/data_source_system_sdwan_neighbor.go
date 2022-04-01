// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemSdwanNeighbor() *schema.Resource {
	return &schema.Resource{
		Description: "Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status.",

		ReadContext: dataSourceSystemSdwanNeighborRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"health_check": {
				Type:        schema.TypeString,
				Description: "SD-WAN health-check name.",
				Computed:    true,
			},
			"ip": {
				Type:        schema.TypeString,
				Description: "IP/IPv6 address of neighbor.",
				Required:    true,
			},
			"member": {
				Type:        schema.TypeList,
				Description: "Member sequence number list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seq_num": {
							Type:        schema.TypeInt,
							Description: "Member sequence number.",
							Computed:    true,
						},
					},
				},
			},
			"minimum_sla_meet_members": {
				Type:        schema.TypeInt,
				Description: "Minimum number of members which meet SLA when the neighbor is preferred.",
				Computed:    true,
			},
			"mode": {
				Type:        schema.TypeString,
				Description: "What metric to select the neighbor.",
				Computed:    true,
			},
			"role": {
				Type:        schema.TypeString,
				Description: "Role of neighbor.",
				Computed:    true,
			},
			"sla_id": {
				Type:        schema.TypeInt,
				Description: "SLA ID.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemSdwanNeighborRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("ip")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadSystemSdwanNeighbor(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSdwanNeighbor dataSource: %v", err)
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

	diags := refreshObjectSystemSdwanNeighbor(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
