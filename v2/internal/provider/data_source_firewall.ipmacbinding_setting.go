// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IP to MAC binding settings.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceFirewallIpmacbindingSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IP to MAC binding settings.",

		ReadContext: dataSourceFirewallIpmacbindingSettingRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"bindthroughfw": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of IP/MAC binding to filter packets that would normally go through the firewall.",
				Computed:    true,
			},
			"bindtofw": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of IP/MAC binding to filter packets that would normally go to the firewall.",
				Computed:    true,
			},
			"undefinedhost": {
				Type:        schema.TypeString,
				Description: "Select action to take on packets with IP/MAC addresses not in the binding list (default = block).",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallIpmacbindingSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "FirewallIpmacbindingSetting"

	o, err := c.Cmdb.ReadFirewallIpmacbindingSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallIpmacbindingSetting dataSource: %v", err)
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

	diags := refreshObjectFirewallIpmacbindingSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
