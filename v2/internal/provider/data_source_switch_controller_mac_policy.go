// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure MAC policy to be applied on the managed FortiSwitch devices through NAC device.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSwitchControllerMacPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure MAC policy to be applied on the managed FortiSwitch devices through NAC device.",

		ReadContext: dataSourceSwitchControllerMacPolicyRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"bounce_port_link": {
				Type:        schema.TypeString,
				Description: "Enable/disable bouncing (administratively bring the link down, up) of a switch port where this mac-policy is applied.",
				Computed:    true,
			},
			"foscount": {
				Type:        schema.TypeString,
				Description: "Enable/disable packet count on the NAC device.",
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description for the MAC policy.",
				Computed:    true,
			},
			"drop": {
				Type:        schema.TypeString,
				Description: "Enable/disable dropping of NAC device traffic.",
				Computed:    true,
			},
			"fortilink": {
				Type:        schema.TypeString,
				Description: "FortiLink interface for which this MAC policy belongs to.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "MAC policy name.",
				Required:    true,
			},
			"traffic_policy": {
				Type:        schema.TypeString,
				Description: "Traffic policy to be applied when using this MAC policy.",
				Computed:    true,
			},
			"vlan": {
				Type:        schema.TypeString,
				Description: "Ingress traffic VLAN assignment for the MAC address matching this MAC policy.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSwitchControllerMacPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerMacPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerMacPolicy dataSource: %v", err)
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

	diags := refreshObjectSwitchControllerMacPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
