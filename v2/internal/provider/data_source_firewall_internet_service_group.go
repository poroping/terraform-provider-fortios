// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure group of Internet Service.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallInternetServiceGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Configure group of Internet Service.",

		ReadContext: dataSourceFirewallInternetServiceGroupRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"direction": {
				Type:        schema.TypeString,
				Description: "How this service may be used (source, destination or both).",
				Computed:    true,
			},
			"member": {
				Type:        schema.TypeList,
				Description: "Internet Service group member.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Internet Service ID.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Internet Service name.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Internet Service group name.",
				Required:    true,
			},
		},
	}
}

func dataSourceFirewallInternetServiceGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallInternetServiceGroup(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallInternetServiceGroup dataSource: %v", err)
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

	diags := refreshObjectFirewallInternetServiceGroup(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
