// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure additional port mappings for Internet Services.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceFirewallInternetServiceAppend() *schema.Resource {
	return &schema.Resource{
		Description: "Configure additional port mappings for Internet Services.",

		ReadContext: dataSourceFirewallInternetServiceAppendRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"append_port": {
				Type:        schema.TypeInt,
				Description: "Appending TCP/UDP/SCTP destination port (1 to 65535).",
				Computed:    true,
			},
			"match_port": {
				Type:        schema.TypeInt,
				Description: "Matching TCP/UDP/SCTP destination port (0 to 65535, 0 means any port).",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallInternetServiceAppendRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallInternetServiceAppend(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallInternetServiceAppend dataSource: %v", err)
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

	diags := refreshObjectFirewallInternetServiceAppend(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
