// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure key-chain.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceRouterKeyChain() *schema.Resource {
	return &schema.Resource{
		Description: "Configure key-chain.",

		ReadContext: dataSourceRouterKeyChainRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"key": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit key settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accept_lifetime": {
							Type:        schema.TypeString,
							Description: "Lifetime of received authentication key (format: hh:mm:ss day month year).",
							Computed:    true,
						},
						"algorithm": {
							Type:        schema.TypeString,
							Description: "Cryptographic algorithm.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeString,
							Description: "Key ID (0 - 2147483647).",
							Computed:    true,
						},
						"key_string": {
							Type:        schema.TypeString,
							Description: "Password for the key (maximum = 64 characters).",
							Computed:    true,
							Sensitive:   true,
						},
						"send_lifetime": {
							Type:        schema.TypeString,
							Description: "Lifetime of sent authentication key (format: hh:mm:ss day month year).",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Key-chain name.",
				Required:    true,
			},
		},
	}
}

func dataSourceRouterKeyChainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterKeyChain(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterKeyChain dataSource: %v", err)
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

	diags := refreshObjectRouterKeyChain(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
