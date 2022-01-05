// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure firewall application groups.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceApplicationGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Configure firewall application groups.",

		ReadContext: dataSourceApplicationGroupRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"application": {
				Type:        schema.TypeList,
				Description: "Application ID list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Application IDs.",
							Computed:    true,
						},
					},
				},
			},
			"behavior": {
				Type:        schema.TypeString,
				Description: "Application behavior filter.",
				Computed:    true,
			},
			"category": {
				Type:        schema.TypeList,
				Description: "Application category ID list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Category IDs.",
							Computed:    true,
						},
					},
				},
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Application group name.",
				Required:    true,
			},
			"popularity": {
				Type:        schema.TypeString,
				Description: "Application popularity filter (1 - 5, from least to most popular).",
				Computed:    true,
			},
			"protocols": {
				Type:        schema.TypeString,
				Description: "Application protocol filter.",
				Computed:    true,
			},
			"risk": {
				Type:        schema.TypeList,
				Description: "Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"level": {
							Type:        schema.TypeInt,
							Description: "Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical).",
							Computed:    true,
						},
					},
				},
			},
			"technology": {
				Type:        schema.TypeString,
				Description: "Application technology filter.",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Application group type.",
				Computed:    true,
			},
			"vendor": {
				Type:        schema.TypeString,
				Description: "Application vendor filter.",
				Computed:    true,
			},
		},
	}
}

func dataSourceApplicationGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadApplicationGroup(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading ApplicationGroup dataSource: %v", err)
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

	diags := refreshObjectApplicationGroup(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
