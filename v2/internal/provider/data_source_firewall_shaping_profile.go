// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure shaping profiles.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceFirewallShapingProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure shaping profiles.",

		ReadContext: dataSourceFirewallShapingProfileRead,

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
			"default_class_id": {
				Type:        schema.TypeInt,
				Description: "Default class ID to handle unclassified packets (including all local traffic).",
				Computed:    true,
			},
			"profile_name": {
				Type:        schema.TypeString,
				Description: "Shaping profile name.",
				Required:    true,
			},
			"shaping_entries": {
				Type:        schema.TypeList,
				Description: "Define shaping entries of this shaping profile.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"burst_in_msec": {
							Type:        schema.TypeInt,
							Description: "Number of bytes that can be burst at maximum-bandwidth speed. Formula: burst = maximum-bandwidth*burst-in-msec.",
							Computed:    true,
						},
						"cburst_in_msec": {
							Type:        schema.TypeInt,
							Description: "Number of bytes that can be burst as fast as the interface can transmit. Formula: cburst = maximum-bandwidth*cburst-in-msec.",
							Computed:    true,
						},
						"class_id": {
							Type:        schema.TypeInt,
							Description: "Class ID.",
							Computed:    true,
						},
						"guaranteed_bandwidth_percentage": {
							Type:        schema.TypeInt,
							Description: "Guaranteed bandwith in percentage.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID number.",
							Computed:    true,
						},
						"limit": {
							Type:        schema.TypeInt,
							Description: "Hard limit on the real queue size in packets.",
							Computed:    true,
						},
						"max": {
							Type:        schema.TypeInt,
							Description: "Average queue size in packets at which RED drop probability is maximal.",
							Computed:    true,
						},
						"maximum_bandwidth_percentage": {
							Type:        schema.TypeInt,
							Description: "Maximum bandwith in percentage.",
							Computed:    true,
						},
						"min": {
							Type:        schema.TypeInt,
							Description: "Average queue size in packets at which RED drop becomes a possibility.",
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeString,
							Description: "Priority.",
							Computed:    true,
						},
						"red_probability": {
							Type:        schema.TypeInt,
							Description: "Maximum probability (in percentage) for RED marking.",
							Computed:    true,
						},
					},
				},
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Select shaping profile type: policing / queuing.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallShapingProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallShapingProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallShapingProfile dataSource: %v", err)
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

	diags := refreshObjectFirewallShapingProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
