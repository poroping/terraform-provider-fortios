// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure YouTube channel filter.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceVideofilterYoutubeChannelFilter() *schema.Resource {
	return &schema.Resource{
		Description: "Configure YouTube channel filter.",

		ReadContext: dataSourceVideofilterYoutubeChannelFilterRead,

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
			"default_action": {
				Type:        schema.TypeString,
				Description: "YouTube channel filter default action.",
				Computed:    true,
			},
			"entries": {
				Type:        schema.TypeList,
				Description: "YouTube filter entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "YouTube channel filter action.",
							Computed:    true,
						},
						"channel_id": {
							Type:        schema.TypeString,
							Description: "Channel ID.",
							Computed:    true,
						},
						"comment": {
							Type:        schema.TypeString,
							Description: "Comment.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
					},
				},
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "ID.",
				Computed:    true,
			},
			"log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name.",
				Computed:    true,
			},
			"override_category": {
				Type:        schema.TypeString,
				Description: "Enable/disable overriding category filtering result.",
				Computed:    true,
			},
		},
	}
}

func dataSourceVideofilterYoutubeChannelFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("fosid")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadVideofilterYoutubeChannelFilter(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VideofilterYoutubeChannelFilter dataSource: %v", err)
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

	diags := refreshObjectVideofilterYoutubeChannelFilter(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
