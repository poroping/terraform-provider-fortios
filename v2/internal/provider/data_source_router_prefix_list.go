// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv4 prefix lists.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceRouterPrefixList() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv4 prefix lists.",

		ReadContext: dataSourceRouterPrefixListRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name.",
				Required:    true,
			},
			"rule": {
				Type:        schema.TypeList,
				Description: "IPv4 prefix list rule.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Permit or deny this IP address and netmask prefix.",
							Computed:    true,
						},
						"flags": {
							Type:        schema.TypeInt,
							Description: "Flags.",
							Computed:    true,
						},
						"ge": {
							Type:        schema.TypeInt,
							Description: "Minimum prefix length to be matched (0 - 32).",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Rule ID.",
							Computed:    true,
						},
						"le": {
							Type:        schema.TypeInt,
							Description: "Maximum prefix length to be matched (0 - 32).",
							Computed:    true,
						},
						"prefix": {
							Type:        schema.TypeString,
							Description: "IPv4 prefix to define regular filter criteria, such as \"any\" or subnets.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterPrefixListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterPrefixList(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterPrefixList dataSource: %v", err)
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

	diags := refreshObjectRouterPrefixList(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
