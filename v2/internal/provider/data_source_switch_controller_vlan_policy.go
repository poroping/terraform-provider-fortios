// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure VLAN policy to be applied on the managed FortiSwitch ports through dynamic-port-policy.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSwitchControllerVlanPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure VLAN policy to be applied on the managed FortiSwitch ports through dynamic-port-policy.",

		ReadContext: dataSourceSwitchControllerVlanPolicyRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"allowed_vlans": {
				Type:        schema.TypeList,
				Description: "Allowed VLANs to be applied when using this VLAN policy.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan_name": {
							Type:        schema.TypeString,
							Description: "VLAN name.",
							Computed:    true,
						},
					},
				},
			},
			"allowed_vlans_all": {
				Type:        schema.TypeString,
				Description: "Enable/disable all defined VLANs when using this VLAN policy.",
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description for the VLAN policy.",
				Computed:    true,
			},
			"discard_mode": {
				Type:        schema.TypeString,
				Description: "Discard mode to be applied when using this VLAN policy.",
				Computed:    true,
			},
			"fortilink": {
				Type:        schema.TypeString,
				Description: "FortiLink interface for which this VLAN policy belongs to.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "VLAN policy name.",
				Required:    true,
			},
			"untagged_vlans": {
				Type:        schema.TypeList,
				Description: "Untagged VLANs to be applied when using this VLAN policy.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan_name": {
							Type:        schema.TypeString,
							Description: "VLAN name.",
							Computed:    true,
						},
					},
				},
			},
			"vlan": {
				Type:        schema.TypeString,
				Description: "Native VLAN to be applied when using this VLAN policy.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSwitchControllerVlanPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerVlanPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerVlanPolicy dataSource: %v", err)
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

	diags := refreshObjectSwitchControllerVlanPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
