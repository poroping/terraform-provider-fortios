// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Trigger for automation stitches.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemAutomationTrigger() *schema.Resource {
	return &schema.Resource{
		Description: "Trigger for automation stitches.",

		ReadContext: dataSourceSystemAutomationTriggerRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description.",
				Computed:    true,
			},
			"event_type": {
				Type:        schema.TypeString,
				Description: "Event type.",
				Computed:    true,
			},
			"fabric_event_name": {
				Type:        schema.TypeString,
				Description: "Fabric connector event handler name.",
				Computed:    true,
			},
			"fabric_event_severity": {
				Type:        schema.TypeString,
				Description: "Fabric connector event severity.",
				Computed:    true,
			},
			"faz_event_name": {
				Type:        schema.TypeString,
				Description: "FortiAnalyzer event handler name.",
				Computed:    true,
			},
			"faz_event_severity": {
				Type:        schema.TypeString,
				Description: "FortiAnalyzer event severity.",
				Computed:    true,
			},
			"faz_event_tags": {
				Type:        schema.TypeString,
				Description: "FortiAnalyzer event tags.",
				Computed:    true,
			},
			"fields": {
				Type:        schema.TypeList,
				Description: "Customized trigger field settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Entry ID.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Name.",
							Computed:    true,
						},
						"value": {
							Type:        schema.TypeString,
							Description: "Value.",
							Computed:    true,
						},
					},
				},
			},
			"ioc_level": {
				Type:        schema.TypeString,
				Description: "IOC threat level.",
				Computed:    true,
			},
			"license_type": {
				Type:        schema.TypeString,
				Description: "License type.",
				Computed:    true,
			},
			"logid": {
				Type:        schema.TypeList,
				Description: "Log IDs to trigger event.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Log ID.",
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
			"report_type": {
				Type:        schema.TypeString,
				Description: "Security Rating report.",
				Computed:    true,
			},
			"serial": {
				Type:        schema.TypeString,
				Description: "Fabric connector serial number.",
				Computed:    true,
			},
			"trigger_datetime": {
				Type:        schema.TypeString,
				Description: "Trigger date and time (YYYY-MM-DD HH:MM:SS).",
				Computed:    true,
			},
			"trigger_day": {
				Type:        schema.TypeInt,
				Description: "Day within a month to trigger.",
				Computed:    true,
			},
			"trigger_frequency": {
				Type:        schema.TypeString,
				Description: "Scheduled trigger frequency (default = daily).",
				Computed:    true,
			},
			"trigger_hour": {
				Type:        schema.TypeInt,
				Description: "Hour of the day on which to trigger (0 - 23, default = 1).",
				Computed:    true,
			},
			"trigger_minute": {
				Type:        schema.TypeInt,
				Description: "Minute of the hour on which to trigger (0 - 59, default = 0).",
				Computed:    true,
			},
			"trigger_type": {
				Type:        schema.TypeString,
				Description: "Trigger type.",
				Computed:    true,
			},
			"trigger_weekday": {
				Type:        schema.TypeString,
				Description: "Day of week for trigger.",
				Computed:    true,
			},
			"vdom": {
				Type:        schema.TypeList,
				Description: "Virtual domain(s) that this trigger is valid for.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Virtual domain name.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemAutomationTriggerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemAutomationTrigger(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemAutomationTrigger dataSource: %v", err)
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

	diags := refreshObjectSystemAutomationTrigger(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
