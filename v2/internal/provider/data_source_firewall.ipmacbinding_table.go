// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IP to MAC address pairs in the IP/MAC binding table.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallIpmacbindingTable() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IP to MAC address pairs in the IP/MAC binding table.",

		ReadContext: dataSourceFirewallIpmacbindingTableRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"ip": {
				Type:        schema.TypeString,
				Description: "IPv4 address portion of the pair (format: xxx.xxx.xxx.xxx).",
				Computed:    true,
			},
			"mac": {
				Type:        schema.TypeString,
				Description: "MAC address portion of the pair (format = xx:xx:xx:xx:xx:xx in hexadecimal).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the pair (optional, default = no name).",
				Computed:    true,
			},
			"seq_num": {
				Type:        schema.TypeInt,
				Description: "Entry number.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable this IP-mac binding pair.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallIpmacbindingTableRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("seq_num")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadFirewallIpmacbindingTable(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallIpmacbindingTable dataSource: %v", err)
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

	diags := refreshObjectFirewallIpmacbindingTable(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
