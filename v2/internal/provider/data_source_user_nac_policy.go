// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure NAC policy matching pattern to identify matching NAC devices.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceUserNacPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure NAC policy matching pattern to identify matching NAC devices.",

		ReadContext: dataSourceUserNacPolicyRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"category": {
				Type:        schema.TypeString,
				Description: "Category of NAC policy.",
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description for the NAC policy matching pattern.",
				Computed:    true,
			},
			"ems_tag": {
				Type:        schema.TypeString,
				Description: "NAC policy matching EMS tag.",
				Computed:    true,
			},
			"family": {
				Type:        schema.TypeString,
				Description: "NAC policy matching family.",
				Computed:    true,
			},
			"firewall_address": {
				Type:        schema.TypeString,
				Description: "Dynamic firewall address to associate MAC which match this policy.",
				Computed:    true,
			},
			"host": {
				Type:        schema.TypeString,
				Description: "NAC policy matching host.",
				Computed:    true,
			},
			"hw_vendor": {
				Type:        schema.TypeString,
				Description: "NAC policy matching hardware vendor.",
				Computed:    true,
			},
			"hw_version": {
				Type:        schema.TypeString,
				Description: "NAC policy matching hardware version.",
				Computed:    true,
			},
			"mac": {
				Type:        schema.TypeString,
				Description: "NAC policy matching MAC address.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "NAC policy name.",
				Required:    true,
			},
			"os": {
				Type:        schema.TypeString,
				Description: "NAC policy matching operating system.",
				Computed:    true,
			},
			"src": {
				Type:        schema.TypeString,
				Description: "NAC policy matching source.",
				Computed:    true,
			},
			"ssid_policy": {
				Type:        schema.TypeString,
				Description: "SSID policy to be applied on the matched NAC policy.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable NAC policy.",
				Computed:    true,
			},
			"sw_version": {
				Type:        schema.TypeString,
				Description: "NAC policy matching software version.",
				Computed:    true,
			},
			"switch_auto_auth": {
				Type:        schema.TypeString,
				Description: "NAC device auto authorization when discovered and nac-policy matched.",
				Computed:    true,
			},
			"switch_fortilink": {
				Type:        schema.TypeString,
				Description: "FortiLink interface for which this NAC policy belongs to.",
				Computed:    true,
			},
			"switch_group": {
				Type:        schema.TypeList,
				Description: "List of managed FortiSwitch groups on which NAC policy can be applied.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Managed FortiSwitch group name from available options.",
							Computed:    true,
						},
					},
				},
			},
			"switch_mac_policy": {
				Type:        schema.TypeString,
				Description: "switch-mac-policy action to be applied on the matched NAC policy.",
				Computed:    true,
			},
			"switch_port_policy": {
				Type:        schema.TypeString,
				Description: "switch-port-policy to be applied on the matched NAC policy.",
				Computed:    true,
			},
			"switch_scope": {
				Type:        schema.TypeList,
				Description: "List of managed FortiSwitches on which NAC policy can be applied.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"switch_id": {
							Type:        schema.TypeString,
							Description: "Managed FortiSwitch name from available options.",
							Computed:    true,
						},
					},
				},
			},
			"type": {
				Type:        schema.TypeString,
				Description: "NAC policy matching type.",
				Computed:    true,
			},
			"user": {
				Type:        schema.TypeString,
				Description: "NAC policy matching user.",
				Computed:    true,
			},
			"user_group": {
				Type:        schema.TypeString,
				Description: "NAC policy matching user group.",
				Computed:    true,
			},
		},
	}
}

func dataSourceUserNacPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserNacPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserNacPolicy dataSource: %v", err)
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

	diags := refreshObjectUserNacPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
