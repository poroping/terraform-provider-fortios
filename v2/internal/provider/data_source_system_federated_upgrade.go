// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Coordinate federated upgrades within the Security Fabric.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemFederatedUpgrade() *schema.Resource {
	return &schema.Resource{
		Description: "Coordinate federated upgrades within the Security Fabric.",

		ReadContext: dataSourceSystemFederatedUpgradeRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"failure_device": {
				Type:        schema.TypeString,
				Description: "Serial number of the node to include.",
				Computed:    true,
			},
			"failure_reason": {
				Type:        schema.TypeString,
				Description: "Reason for upgrade failure.",
				Computed:    true,
			},
			"node_list": {
				Type:        schema.TypeList,
				Description: "Nodes which will be included in the upgrade.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"coordinating_fortigate": {
							Type:        schema.TypeString,
							Description: "The serial of the FortiGate that controls this device",
							Computed:    true,
						},
						"device_type": {
							Type:        schema.TypeString,
							Description: "What type of device this node represents.",
							Computed:    true,
						},
						"serial": {
							Type:        schema.TypeString,
							Description: "Serial number of the node to include.",
							Computed:    true,
						},
						"setup_time": {
							Type:        schema.TypeString,
							Description: "When the upgrade was configured. Format hh:mm yyyy/mm/dd UTC.",
							Computed:    true,
						},
						"time": {
							Type:        schema.TypeString,
							Description: "Scheduled time for the upgrade. Format hh:mm yyyy/mm/dd UTC.",
							Computed:    true,
						},
						"timing": {
							Type:        schema.TypeString,
							Description: "Whether the upgrade should be run immediately, or at a scheduled time.",
							Computed:    true,
						},
						"upgrade_path": {
							Type:        schema.TypeString,
							Description: "Image IDs to upgrade through.",
							Computed:    true,
						},
					},
				},
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Current status of the upgrade.",
				Computed:    true,
			},
			"upgrade_id": {
				Type:        schema.TypeInt,
				Description: "Unique identifier for this upgrade.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemFederatedUpgradeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemFederatedUpgrade"

	o, err := c.Cmdb.ReadSystemFederatedUpgrade(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemFederatedUpgrade dataSource: %v", err)
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

	diags := refreshObjectSystemFederatedUpgrade(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
