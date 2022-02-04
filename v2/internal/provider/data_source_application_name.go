// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure application signatures.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceApplicationName() *schema.Resource {
	return &schema.Resource{
		Description: "Configure application signatures.",

		ReadContext: dataSourceApplicationNameRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"behavior": {
				Type:        schema.TypeString,
				Description: "Application behavior.",
				Computed:    true,
			},
			"category": {
				Type:        schema.TypeInt,
				Description: "Application category ID.",
				Computed:    true,
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "Application ID.",
				Computed:    true,
			},
			"metadata": {
				Type:        schema.TypeList,
				Description: "Meta data.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"metaid": {
							Type:        schema.TypeInt,
							Description: "Meta ID.",
							Computed:    true,
						},
						"valueid": {
							Type:        schema.TypeInt,
							Description: "Value ID.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Application name.",
				Required:    true,
			},
			"parameter": {
				Type:        schema.TypeString,
				Description: "Application parameter name.",
				Computed:    true,
			},
			"parameters": {
				Type:        schema.TypeList,
				Description: "Application parameters.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"defaultvalue": {
							Type:        schema.TypeString,
							Description: "Parameter default value.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Parameter name.",
							Computed:    true,
						},
					},
				},
			},
			"popularity": {
				Type:        schema.TypeInt,
				Description: "Application popularity.",
				Computed:    true,
			},
			"protocol": {
				Type:        schema.TypeString,
				Description: "Application protocol.",
				Computed:    true,
			},
			"risk": {
				Type:        schema.TypeInt,
				Description: "Application risk.",
				Computed:    true,
			},
			"sub_category": {
				Type:        schema.TypeInt,
				Description: "Application sub-category ID.",
				Computed:    true,
			},
			"technology": {
				Type:        schema.TypeString,
				Description: "Application technology.",
				Computed:    true,
			},
			"vendor": {
				Type:        schema.TypeString,
				Description: "Application vendor.",
				Computed:    true,
			},
			"weight": {
				Type:        schema.TypeInt,
				Description: "Application weight.",
				Computed:    true,
			},
		},
	}
}

func dataSourceApplicationNameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadApplicationName(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading ApplicationName dataSource: %v", err)
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

	diags := refreshObjectApplicationName(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
