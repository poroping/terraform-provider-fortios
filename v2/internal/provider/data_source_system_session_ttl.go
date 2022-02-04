// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure global session TTL timers for this FortiGate.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemSessionTtl() *schema.Resource {
	return &schema.Resource{
		Description: "Configure global session TTL timers for this FortiGate.",

		ReadContext: dataSourceSystemSessionTtlRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"default": {
				Type:        schema.TypeString,
				Description: "Default timeout.",
				Computed:    true,
			},
			"port": {
				Type:        schema.TypeList,
				Description: "Session TTL port.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_port": {
							Type:        schema.TypeInt,
							Description: "End port number.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Table entry ID.",
							Computed:    true,
						},
						"protocol": {
							Type:        schema.TypeInt,
							Description: "Protocol (0 - 255).",
							Computed:    true,
						},
						"start_port": {
							Type:        schema.TypeInt,
							Description: "Start port number.",
							Computed:    true,
						},
						"timeout": {
							Type:        schema.TypeString,
							Description: "Session timeout (TTL).",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemSessionTtlRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemSessionTtl"

	o, err := c.Cmdb.ReadSystemSessionTtl(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSessionTtl dataSource: %v", err)
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

	diags := refreshObjectSystemSessionTtl(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
