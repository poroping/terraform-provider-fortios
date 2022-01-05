// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure AntiSpam banned word list.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceEmailfilterBword() *schema.Resource {
	return &schema.Resource{
		Description: "Configure AntiSpam banned word list.",

		ReadContext: dataSourceEmailfilterBwordRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Optional comments.",
				Computed:    true,
			},
			"entries": {
				Type:        schema.TypeList,
				Description: "Spam filter banned word.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Mark spam or good.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Banned word entry ID.",
							Computed:    true,
						},
						"language": {
							Type:        schema.TypeString,
							Description: "Language for the banned word.",
							Computed:    true,
						},
						"pattern": {
							Type:        schema.TypeString,
							Description: "Pattern for the banned word.",
							Computed:    true,
						},
						"pattern_type": {
							Type:        schema.TypeString,
							Description: "Wildcard pattern or regular expression.",
							Computed:    true,
						},
						"score": {
							Type:        schema.TypeInt,
							Description: "Score value.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable status.",
							Computed:    true,
						},
						"where": {
							Type:        schema.TypeString,
							Description: "Component of the email to be scanned.",
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
			"name": {
				Type:        schema.TypeString,
				Description: "Name of table.",
				Computed:    true,
			},
		},
	}
}

func dataSourceEmailfilterBwordRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadEmailfilterBword(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading EmailfilterBword dataSource: %v", err)
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

	diags := refreshObjectEmailfilterBword(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
