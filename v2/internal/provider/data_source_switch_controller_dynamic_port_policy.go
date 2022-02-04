// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Dynamic port policy to be applied on the managed FortiSwitch ports through DPP device.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSwitchControllerDynamicPortPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Dynamic port policy to be applied on the managed FortiSwitch ports through DPP device.",

		ReadContext: dataSourceSwitchControllerDynamicPortPolicyRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description for the Dynamic port policy.",
				Computed:    true,
			},
			"fortilink": {
				Type:        schema.TypeString,
				Description: "FortiLink interface for which this Dynamic port policy belongs to.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Dynamic port policy name.",
				Required:    true,
			},
			"policy": {
				Type:        schema.TypeList,
				Description: "Port policies with matching criteria and actions.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"802_1x": {
							Type:        schema.TypeString,
							Description: "802.1x security policy to be applied when using this policy.",
							Computed:    true,
						},
						"bounce_port_link": {
							Type:        schema.TypeString,
							Description: "Enable/disable bouncing (administratively bring the link down, up) of a switch port where this policy is applied. Helps to clear and reassign VLAN from lldp-profile.",
							Computed:    true,
						},
						"category": {
							Type:        schema.TypeString,
							Description: "Category of Dynamic port policy.",
							Computed:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "Description for the policy.",
							Computed:    true,
						},
						"family": {
							Type:        schema.TypeString,
							Description: "Match policy based on family.",
							Computed:    true,
						},
						"host": {
							Type:        schema.TypeString,
							Description: "Match policy based on host.",
							Computed:    true,
						},
						"hw_vendor": {
							Type:        schema.TypeString,
							Description: "Match policy based on hardware vendor.",
							Computed:    true,
						},
						"interface_tags": {
							Type:        schema.TypeList,
							Description: "Match policy based on the FortiSwitch interface object tags.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag_name": {
										Type:        schema.TypeString,
										Description: "FortiSwitch port tag name.",
										Computed:    true,
									},
								},
							},
						},
						"lldp_profile": {
							Type:        schema.TypeString,
							Description: "LLDP profile to be applied when using this policy.",
							Computed:    true,
						},
						"mac": {
							Type:        schema.TypeString,
							Description: "Match policy based on MAC address.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Policy name.",
							Computed:    true,
						},
						"qos_policy": {
							Type:        schema.TypeString,
							Description: "QoS policy to be applied when using this policy.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable policy.",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "Match policy based on type.",
							Computed:    true,
						},
						"vlan_policy": {
							Type:        schema.TypeString,
							Description: "VLAN policy to be applied when using this policy.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSwitchControllerDynamicPortPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerDynamicPortPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerDynamicPortPolicy dataSource: %v", err)
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

	diags := refreshObjectSwitchControllerDynamicPortPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
