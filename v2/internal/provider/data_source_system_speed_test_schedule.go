// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Speed test schedule for each interface.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemSpeedTestSchedule() *schema.Resource {
	return &schema.Resource{
		Description: "Speed test schedule for each interface.",

		ReadContext: dataSourceSystemSpeedTestScheduleRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"diffserv": {
				Type:        schema.TypeString,
				Description: "DSCP used for speed test.",
				Computed:    true,
			},
			"dynamic_server": {
				Type:        schema.TypeString,
				Description: "Enable/disable dynamic server option.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Interface name.",
				Required:    true,
			},
			"mode": {
				Type:        schema.TypeString,
				Description: "Protocol Auto(default), TCP or UDP used for speed test.",
				Computed:    true,
			},
			"schedules": {
				Type:        schema.TypeList,
				Description: "Schedules for the interface.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Name of a firewall recurring schedule.",
							Computed:    true,
						},
					},
				},
			},
			"server_name": {
				Type:        schema.TypeString,
				Description: "Speed test server name.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable scheduled speed test.",
				Computed:    true,
			},
			"update_inbandwidth": {
				Type:        schema.TypeString,
				Description: "Enable/disable bypassing interface's inbound bandwidth setting.",
				Computed:    true,
			},
			"update_inbandwidth_maximum": {
				Type:        schema.TypeInt,
				Description: "Maximum downloading bandwidth (kbps) to be used in a speed test.",
				Computed:    true,
			},
			"update_inbandwidth_minimum": {
				Type:        schema.TypeInt,
				Description: "Minimum downloading bandwidth (kbps) to be considered effective.",
				Computed:    true,
			},
			"update_outbandwidth": {
				Type:        schema.TypeString,
				Description: "Enable/disable bypassing interface's outbound bandwidth setting.",
				Computed:    true,
			},
			"update_outbandwidth_maximum": {
				Type:        schema.TypeInt,
				Description: "Maximum uploading bandwidth (kbps) to be used in a speed test.",
				Computed:    true,
			},
			"update_outbandwidth_minimum": {
				Type:        schema.TypeInt,
				Description: "Minimum uploading bandwidth (kbps) to be considered effective.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemSpeedTestScheduleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("interface")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadSystemSpeedTestSchedule(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSpeedTestSchedule dataSource: %v", err)
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

	diags := refreshObjectSystemSpeedTestSchedule(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
