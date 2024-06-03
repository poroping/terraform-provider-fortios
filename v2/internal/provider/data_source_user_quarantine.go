// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure quarantine support.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceUserQuarantine() *schema.Resource {
	return &schema.Resource{
		Description: "Configure quarantine support.",

		ReadContext: dataSourceUserQuarantineRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"firewall_groups": {
				Type:        schema.TypeString,
				Description: "Firewall address group which includes all quarantine MAC address.",
				Computed:    true,
			},
			"quarantine": {
				Type:        schema.TypeString,
				Description: "Enable/disable quarantine.",
				Computed:    true,
			},
			"targets": {
				Type:        schema.TypeList,
				Description: "Quarantine entry to hold multiple MACs.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:        schema.TypeString,
							Description: "Description for the quarantine entry.",
							Computed:    true,
						},
						"entry": {
							Type:        schema.TypeString,
							Description: "Quarantine entry name.",
							Computed:    true,
						},
						"macs": {
							Type:        schema.TypeList,
							Description: "Quarantine MACs.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"description": {
										Type:        schema.TypeString,
										Description: "Description for the quarantine MAC.",
										Computed:    true,
									},
									"drop": {
										Type:        schema.TypeString,
										Description: "Enable/disable dropping of quarantined device traffic.",
										Computed:    true,
									},
									"mac": {
										Type:        schema.TypeString,
										Description: "Quarantine MAC.",
										Computed:    true,
									},
									"parent": {
										Type:        schema.TypeString,
										Description: "Parent entry name.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"traffic_policy": {
				Type:        schema.TypeString,
				Description: "Traffic policy for quarantined MACs.",
				Computed:    true,
			},
		},
	}
}

func dataSourceUserQuarantineRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "UserQuarantine"

	o, err := c.Cmdb.ReadUserQuarantine(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserQuarantine dataSource: %v", err)
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

	diags := refreshObjectUserQuarantine(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
