// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Global configuration objects that can be configured independently across different ha peers for all VDOMs or for the defined VDOM scope.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemVdomException() *schema.Resource {
	return &schema.Resource{
		Description: "Global configuration objects that can be configured independently across different ha peers for all VDOMs or for the defined VDOM scope.",

		ReadContext: dataSourceSystemVdomExceptionRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "Index <1-4096>.",
				Computed:    true,
			},
			"object": {
				Type:        schema.TypeString,
				Description: "Name of the configuration object that can be configured independently for all VDOMs.",
				Computed:    true,
			},
			"scope": {
				Type:        schema.TypeString,
				Description: "Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration.",
				Computed:    true,
			},
			"vdom": {
				Type:        schema.TypeList,
				Description: "Names of the VDOMs.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "VDOM name.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemVdomExceptionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemVdomException(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemVdomException dataSource: %v", err)
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

	diags := refreshObjectSystemVdomException(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
