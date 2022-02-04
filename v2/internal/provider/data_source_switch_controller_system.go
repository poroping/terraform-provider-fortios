// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure system-wide switch controller settings.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSwitchControllerSystem() *schema.Resource {
	return &schema.Resource{
		Description: "Configure system-wide switch controller settings.",

		ReadContext: dataSourceSwitchControllerSystemRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"data_sync_interval": {
				Type:        schema.TypeInt,
				Description: "Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).",
				Computed:    true,
			},
			"dynamic_periodic_interval": {
				Type:        schema.TypeInt,
				Description: "Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).",
				Computed:    true,
			},
			"iot_holdoff": {
				Type:        schema.TypeInt,
				Description: "MAC entry's creation time. Time must be greater than this value for an entry to be created (0 - 10080 mins, default = 5 mins).",
				Computed:    true,
			},
			"iot_mac_idle": {
				Type:        schema.TypeInt,
				Description: "MAC entry's idle time. MAC entry is removed after this value (0 - 10080 mins, default = 1440 mins).",
				Computed:    true,
			},
			"iot_scan_interval": {
				Type:        schema.TypeInt,
				Description: "IoT scan interval (2 - 10080 mins, default = 60 mins, 0 = disable).",
				Computed:    true,
			},
			"iot_weight_threshold": {
				Type:        schema.TypeInt,
				Description: "MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).",
				Computed:    true,
			},
			"nac_periodic_interval": {
				Type:        schema.TypeInt,
				Description: "Periodic time interval to run NAC engine (5 - 60 sec, default = 15).",
				Computed:    true,
			},
			"parallel_process": {
				Type:        schema.TypeInt,
				Description: "Maximum number of parallel processes.",
				Computed:    true,
			},
			"parallel_process_override": {
				Type:        schema.TypeString,
				Description: "Enable/disable parallel process override.",
				Computed:    true,
			},
			"tunnel_mode": {
				Type:        schema.TypeString,
				Description: "Compatible/strict tunnel mode.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSwitchControllerSystemRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SwitchControllerSystem"

	o, err := c.Cmdb.ReadSwitchControllerSystem(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerSystem dataSource: %v", err)
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

	diags := refreshObjectSwitchControllerSystem(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
