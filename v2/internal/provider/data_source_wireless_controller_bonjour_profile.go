// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Bonjour profiles. Bonjour is Apple's zero configuration networking protocol. Bonjour profiles allow APs and FortiAPs to connnect to networks using Bonjour.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceWirelessControllerBonjourProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Bonjour profiles. Bonjour is Apple's zero configuration networking protocol. Bonjour profiles allow APs and FortiAPs to connnect to networks using Bonjour.",

		ReadContext: dataSourceWirelessControllerBonjourProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Bonjour profile name.",
				Required:    true,
			},
			"policy_list": {
				Type:        schema.TypeList,
				Description: "Bonjour policy list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:        schema.TypeString,
							Description: "Description.",
							Computed:    true,
						},
						"from_vlan": {
							Type:        schema.TypeString,
							Description: "VLAN ID from which the Bonjour service is advertised (0 - 4094, default = 0).",
							Computed:    true,
						},
						"policy_id": {
							Type:        schema.TypeInt,
							Description: "Policy ID.",
							Computed:    true,
						},
						"services": {
							Type:        schema.TypeString,
							Description: "Bonjour services for the VLAN connecting to the Bonjour network.",
							Computed:    true,
						},
						"to_vlan": {
							Type:        schema.TypeString,
							Description: "VLAN ID to which the Bonjour service is made available (0 - 4094, default = all).",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWirelessControllerBonjourProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerBonjourProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerBonjourProfile dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerBonjourProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
