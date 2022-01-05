// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiSwitch spanning tree protocol (STP).

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSwitchControllerStpSettings() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiSwitch spanning tree protocol (STP).",

		ReadContext: dataSourceSwitchControllerStpSettingsRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"forward_time": {
				Type:        schema.TypeInt,
				Description: "Period of time a port is in listening and learning state (4 - 30 sec, default = 15).",
				Computed:    true,
			},
			"hello_time": {
				Type:        schema.TypeInt,
				Description: "Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).",
				Computed:    true,
			},
			"max_age": {
				Type:        schema.TypeInt,
				Description: "Maximum time before a bridge port expires its configuration BPDU information (6 - 40 sec, default = 20).",
				Computed:    true,
			},
			"max_hops": {
				Type:        schema.TypeInt,
				Description: "Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of global STP settings configuration.",
				Computed:    true,
			},
			"pending_timer": {
				Type:        schema.TypeInt,
				Description: "Pending time (1 - 15 sec, default = 4).",
				Computed:    true,
			},
			"revision": {
				Type:        schema.TypeInt,
				Description: "STP revision number (0 - 65535).",
				Computed:    true,
			},
		},
	}
}

func dataSourceSwitchControllerStpSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerStpSettings(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerStpSettings dataSource: %v", err)
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

	diags := refreshObjectSwitchControllerStpSettings(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
