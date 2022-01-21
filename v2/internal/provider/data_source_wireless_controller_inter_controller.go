// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure inter wireless controller operation.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceWirelessControllerInterController() *schema.Resource {
	return &schema.Resource{
		Description: "Configure inter wireless controller operation.",

		ReadContext: dataSourceWirelessControllerInterControllerRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"fast_failover_max": {
				Type:        schema.TypeInt,
				Description: "Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).",
				Computed:    true,
			},
			"fast_failover_wait": {
				Type:        schema.TypeInt,
				Description: "Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).",
				Computed:    true,
			},
			"inter_controller_key": {
				Type:        schema.TypeString,
				Description: "Secret key for inter-controller communications.",
				Computed:    true,
				Sensitive:   true,
			},
			"inter_controller_mode": {
				Type:        schema.TypeString,
				Description: "Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable).",
				Computed:    true,
			},
			"inter_controller_peer": {
				Type:        schema.TypeList,
				Description: "Fast failover peer wireless controller list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"peer_ip": {
							Type:        schema.TypeString,
							Description: "Peer wireless controller's IP address.",
							Computed:    true,
						},
						"peer_port": {
							Type:        schema.TypeInt,
							Description: "Port used by the wireless controller's for inter-controller communications (1024 - 49150, default = 5246).",
							Computed:    true,
						},
						"peer_priority": {
							Type:        schema.TypeString,
							Description: "Peer wireless controller's priority (primary or secondary, default = primary).",
							Computed:    true,
						},
					},
				},
			},
			"inter_controller_pri": {
				Type:        schema.TypeString,
				Description: "Configure inter-controller's priority (primary or secondary, default = primary).",
				Computed:    true,
			},
		},
	}
}

func dataSourceWirelessControllerInterControllerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "WirelessControllerInterController"

	o, err := c.Cmdb.ReadWirelessControllerInterController(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerInterController dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerInterController(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
