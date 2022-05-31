// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure network visibility settings.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemNetworkVisibility() *schema.Resource {
	return &schema.Resource{
		Description: "Configure network visibility settings.",

		ReadContext: dataSourceSystemNetworkVisibilityRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"destination_hostname_visibility": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging of destination hostname visibility.",
				Computed:    true,
			},
			"destination_location": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging of destination geographical location visibility.",
				Computed:    true,
			},
			"destination_visibility": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging of destination visibility.",
				Computed:    true,
			},
			"hostname_limit": {
				Type:        schema.TypeInt,
				Description: "Limit of the number of hostname table entries (0 - 50000).",
				Computed:    true,
			},
			"hostname_ttl": {
				Type:        schema.TypeInt,
				Description: "TTL of hostname table entries (60 - 86400).",
				Computed:    true,
			},
			"source_location": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging of source geographical location visibility.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemNetworkVisibilityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemNetworkVisibility"

	o, err := c.Cmdb.ReadSystemNetworkVisibility(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemNetworkVisibility dataSource: %v", err)
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

	diags := refreshObjectSystemNetworkVisibility(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
