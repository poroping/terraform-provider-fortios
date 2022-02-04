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
)

func dataSourceFirewallIpmacbindingTableList() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IP to MAC address pairs in the IP/MAC binding table.",

		ReadContext: dataSourceFirewallIpmacbindingTableListRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"filter": {
				Type:        schema.TypeString,
				Description: "Filter to apply to query",
				Optional:    true,
				ForceNew:    true,
			},
			"seq_numlist": {
				Type:        schema.TypeList,
				Description: "List of results",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func dataSourceFirewallIpmacbindingTableListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	setfilter := ""
	filter := []string{}
	if v, ok := d.GetOk("filter"); ok {
		if s, ok := v.(string); ok {
			setfilter = s
			filter = []string{s}
		}
	}
	urlparams.Filter = &filter

	format := []string{"seq-num"}
	urlparams.Format = &format

	o, err := c.Cmdb.ListFirewallIpmacbindingTable(urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallIpmacbindingTable dataSource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] dataSource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if len(*o) == 0 {
		return nil
	}

	results := []int{}
	for _, v := range *o {
		v2 := v.SeqNum
		results = append(results, int(*v2))
	}

	d.Set("seq_numlist", results)

	d.SetId("DataSourceFirewallIpmacbindingTableList" + setfilter)

	return nil
}
