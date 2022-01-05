// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPS rules.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceIpsRule() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPS rules.",

		ReadContext: dataSourceIpsRuleRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"action": {
				Type:        schema.TypeString,
				Description: "Action.",
				Computed:    true,
			},
			"application": {
				Type:        schema.TypeString,
				Description: "Vulnerable applications.",
				Computed:    true,
			},
			"date": {
				Type:        schema.TypeInt,
				Description: "Date.",
				Computed:    true,
			},
			"group": {
				Type:        schema.TypeString,
				Description: "Group.",
				Computed:    true,
			},
			"location": {
				Type:        schema.TypeString,
				Description: "Vulnerable location.",
				Computed:    true,
			},
			"log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging.",
				Computed:    true,
			},
			"log_packet": {
				Type:        schema.TypeString,
				Description: "Enable/disable packet logging.",
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
				Description: "Rule name.",
				Required:    true,
			},
			"os": {
				Type:        schema.TypeString,
				Description: "Vulnerable operation systems.",
				Computed:    true,
			},
			"rev": {
				Type:        schema.TypeInt,
				Description: "Revision.",
				Computed:    true,
			},
			"rule_id": {
				Type:        schema.TypeInt,
				Description: "Rule ID.",
				Computed:    true,
			},
			"service": {
				Type:        schema.TypeString,
				Description: "Vulnerable service.",
				Computed:    true,
			},
			"severity": {
				Type:        schema.TypeString,
				Description: "Severity.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable status.",
				Computed:    true,
			},
		},
	}
}

func dataSourceIpsRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadIpsRule(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading IpsRule dataSource: %v", err)
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

	diags := refreshObjectIpsRule(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
