// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure initial template for auto-generated VLAN interfaces.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSwitchControllerInitialConfigVlans() *schema.Resource {
	return &schema.Resource{
		Description: "Configure initial template for auto-generated VLAN interfaces.",

		ReadContext: dataSourceSwitchControllerInitialConfigVlansRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"default_vlan": {
				Type:        schema.TypeString,
				Description: "Default VLAN (native) assigned to all switch ports upon discovery.",
				Computed:    true,
			},
			"nac": {
				Type:        schema.TypeString,
				Description: "VLAN for NAC onboarding devices.",
				Computed:    true,
			},
			"nac_segment": {
				Type:        schema.TypeString,
				Description: "VLAN for NAC segment primary interface.",
				Computed:    true,
			},
			"quarantine": {
				Type:        schema.TypeString,
				Description: "VLAN for quarantined traffic.",
				Computed:    true,
			},
			"rspan": {
				Type:        schema.TypeString,
				Description: "VLAN for RSPAN/ERSPAN mirrored traffic.",
				Computed:    true,
			},
			"video": {
				Type:        schema.TypeString,
				Description: "VLAN dedicated for video devices.",
				Computed:    true,
			},
			"voice": {
				Type:        schema.TypeString,
				Description: "VLAN dedicated for voice devices.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSwitchControllerInitialConfigVlansRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SwitchControllerInitialConfigVlans"

	o, err := c.Cmdb.ReadSwitchControllerInitialConfigVlans(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerInitialConfigVlans dataSource: %v", err)
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

	diags := refreshObjectSwitchControllerInitialConfigVlans(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
