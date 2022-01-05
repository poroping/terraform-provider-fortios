// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure consolidated IPv4/IPv6 policies.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceFirewallConsolidatedPolicyList() *schema.Resource {
	return &schema.Resource{
		Description: "Configure consolidated IPv4/IPv6 policies.",

		ReadContext: dataSourceFirewallConsolidatedPolicyListRead,

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
			"policyidlist": {
				Type:        schema.TypeList,
				Description: "List of results",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func dataSourceFirewallConsolidatedPolicyListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	setfilter := ""
	filter := []string{}
	if v, ok := d.GetOk("filter"); ok {
		if s, ok := v.(string); ok {
			setfilter = s
			filter = []string{s}
		}
	}
	urlparams.Filter = &filter

	format := []string{"policyid"}
	urlparams.Format = &format

	o, err := c.Cmdb.ListFirewallConsolidatedPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallConsolidatedPolicy dataSource: %v", err)
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
		v2 := v.Policyid
		results = append(results, int(*v2))
	}

	d.Set("policyidlist", results)

	d.SetId("DataSourceFirewallConsolidatedPolicyList" + setfilter)

	return nil
}
