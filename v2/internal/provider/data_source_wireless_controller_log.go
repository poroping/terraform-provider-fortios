// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure wireless controller event log filters.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceWirelessControllerLog() *schema.Resource {
	return &schema.Resource{
		Description: "Configure wireless controller event log filters.",

		ReadContext: dataSourceWirelessControllerLogRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"addrgrp_log": {
				Type:        schema.TypeString,
				Description: "Lowest severity level to log address group message.",
				Computed:    true,
			},
			"ble_log": {
				Type:        schema.TypeString,
				Description: "Lowest severity level to log BLE detection message.",
				Computed:    true,
			},
			"clb_log": {
				Type:        schema.TypeString,
				Description: "Lowest severity level to log client load balancing message.",
				Computed:    true,
			},
			"dhcp_starv_log": {
				Type:        schema.TypeString,
				Description: "Lowest severity level to log DHCP starvation event message.",
				Computed:    true,
			},
			"led_sched_log": {
				Type:        schema.TypeString,
				Description: "Lowest severity level to log LED schedule event message.",
				Computed:    true,
			},
			"radio_event_log": {
				Type:        schema.TypeString,
				Description: "Lowest severity level to log radio event message.",
				Computed:    true,
			},
			"rogue_event_log": {
				Type:        schema.TypeString,
				Description: "Lowest severity level to log rogue AP event message.",
				Computed:    true,
			},
			"sta_event_log": {
				Type:        schema.TypeString,
				Description: "Lowest severity level to log station event message.",
				Computed:    true,
			},
			"sta_locate_log": {
				Type:        schema.TypeString,
				Description: "Lowest severity level to log station locate message.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable wireless event logging.",
				Computed:    true,
			},
			"wids_log": {
				Type:        schema.TypeString,
				Description: "Lowest severity level to log WIDS message.",
				Computed:    true,
			},
			"wtp_event_log": {
				Type:        schema.TypeString,
				Description: "Lowest severity level to log WTP event message.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWirelessControllerLogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "WirelessControllerLog"

	o, err := c.Cmdb.ReadWirelessControllerLog(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerLog dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerLog(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
