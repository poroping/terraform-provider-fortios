// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Report dataset configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceReportDataset() *schema.Resource {
	return &schema.Resource{
		Description: "Report dataset configuration.",

		ReadContext: dataSourceReportDatasetRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"field": {
				Type:        schema.TypeList,
				Description: "Fields.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"displayname": {
							Type:        schema.TypeString,
							Description: "Display name.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Field ID (1 to number of columns in SQL result).",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Name.",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "Field type.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name.",
				Required:    true,
			},
			"parameters": {
				Type:        schema.TypeList,
				Description: "Parameters.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_type": {
							Type:        schema.TypeString,
							Description: "Data type.",
							Computed:    true,
						},
						"display_name": {
							Type:        schema.TypeString,
							Description: "Display name.",
							Computed:    true,
						},
						"field": {
							Type:        schema.TypeString,
							Description: "SQL field name.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Parameter ID (1 to number of columns in SQL result).",
							Computed:    true,
						},
					},
				},
			},
			"policy": {
				Type:        schema.TypeInt,
				Description: "Used by monitor policy.",
				Computed:    true,
			},
			"query": {
				Type:        schema.TypeString,
				Description: "SQL query statement.",
				Computed:    true,
			},
		},
	}
}

func dataSourceReportDatasetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadReportDataset(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading ReportDataset dataSource: %v", err)
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

	diags := refreshObjectReportDataset(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
