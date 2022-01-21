// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure port policy to be applied on the managed FortiSwitch ports through NAC device.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSwitchControllerPortPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure port policy to be applied on the managed FortiSwitch ports through NAC device.",

		ReadContext: dataSourceSwitchControllerPortPolicyRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"802_1x": {
				Type:        schema.TypeString,
				Description: "802.1x security policy to be applied when using this port-policy.",
				Computed:    true,
			},
			"bounce_port_link": {
				Type:        schema.TypeString,
				Description: "Enable/disable bouncing (administratively bring the link down, up) of a switch port where this port policy is applied. Helps to clear and reassign VLAN from lldp-profile.",
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description for the port policy.",
				Computed:    true,
			},
			"fortilink": {
				Type:        schema.TypeString,
				Description: "FortiLink interface for which this port policy belongs to.",
				Computed:    true,
			},
			"lldp_profile": {
				Type:        schema.TypeString,
				Description: "LLDP profile to be applied when using this port-policy.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Port policy name.",
				Required:    true,
			},
			"qos_policy": {
				Type:        schema.TypeString,
				Description: "QoS policy to be applied when using this port-policy.",
				Computed:    true,
			},
			"vlan_policy": {
				Type:        schema.TypeString,
				Description: "VLAN policy to be applied when using this port-policy.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSwitchControllerPortPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerPortPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerPortPolicy dataSource: %v", err)
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

	diags := refreshObjectSwitchControllerPortPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
