// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.6,v7.0.2 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure virtual hardware switch interfaces.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemVirtualSwitch() *schema.Resource {
	return &schema.Resource{
		Description: "Configure virtual hardware switch interfaces.",

		ReadContext: dataSourceSystemVirtualSwitchRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the virtual switch.",
				Required:    true,
			},
			"physical_switch": {
				Type:        schema.TypeString,
				Description: "Physical switch parent.",
				Computed:    true,
			},
			"port": {
				Type:        schema.TypeList,
				Description: "Configure member ports.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alias": {
							Type:        schema.TypeString,
							Description: "Alias.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Physical interface name.",
							Computed:    true,
						},
						"speed": {
							Type:        schema.TypeString,
							Description: "Interface speed.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Interface status.",
							Computed:    true,
						},
					},
				},
			},
			"span": {
				Type:        schema.TypeString,
				Description: "Enable/disable SPAN.",
				Computed:    true,
			},
			"span_dest_port": {
				Type:        schema.TypeString,
				Description: "SPAN destination port.",
				Computed:    true,
			},
			"span_direction": {
				Type:        schema.TypeString,
				Description: "SPAN direction.",
				Computed:    true,
			},
			"span_source_port": {
				Type:        schema.TypeString,
				Description: "SPAN source port.",
				Computed:    true,
			},
			"vlan": {
				Type:        schema.TypeInt,
				Description: "VLAN.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemVirtualSwitchRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemVirtualSwitch(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemVirtualSwitch dataSource: %v", err)
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

	diags := refreshObjectSystemVirtualSwitch(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
