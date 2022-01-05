// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure web proxy address group.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceFirewallProxyAddrgrp() *schema.Resource {
	return &schema.Resource{
		Description: "Configure web proxy address group.",

		ReadContext: dataSourceFirewallProxyAddrgrpRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"color": {
				Type:        schema.TypeInt,
				Description: "Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Optional comments.",
				Computed:    true,
			},
			"member": {
				Type:        schema.TypeList,
				Description: "Members of address group.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address name.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Address group name.",
				Required:    true,
			},
			"tagging": {
				Type:        schema.TypeList,
				Description: "Config object tagging.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:        schema.TypeString,
							Description: "Tag category.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Tagging entry name.",
							Computed:    true,
						},
						"tags": {
							Type:        schema.TypeList,
							Description: "Tags.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Tag name.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Source or destination address group type.",
				Computed:    true,
			},
			"uuid": {
				Type:        schema.TypeString,
				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Computed:    true,
			},
			"visibility": {
				Type:        schema.TypeString,
				Description: "Enable/disable visibility of the object in the GUI.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallProxyAddrgrpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallProxyAddrgrp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallProxyAddrgrp dataSource: %v", err)
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

	diags := refreshObjectFirewallProxyAddrgrp(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
