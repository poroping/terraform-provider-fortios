// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv6 prefix lists.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceRouterPrefixList6() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv6 prefix lists.",

		ReadContext: dataSourceRouterPrefixList6Read,

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
				Description: "IPv6 prefix list rule.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Permit or deny packets that match this rule.",
							Computed:    true,
						},
						"flags": {
							Type:        schema.TypeInt,
							Description: "Flags.",
							Computed:    true,
						},
						"ge": {
							Type:        schema.TypeInt,
							Description: "Minimum prefix length to be matched (0 - 128).",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Rule ID.",
							Computed:    true,
						},
						"le": {
							Type:        schema.TypeInt,
							Description: "Maximum prefix length to be matched (0 - 128).",
							Computed:    true,
						},
						"prefix6": {
							Type:        schema.TypeString,
							Description: "IPv6 prefix to define regular filter criteria, such as \"any\" or subnets.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterPrefixList6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("name")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadRouterPrefixList6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterPrefixList6 dataSource: %v", err)
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

	diags := refreshObjectRouterPrefixList6(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
