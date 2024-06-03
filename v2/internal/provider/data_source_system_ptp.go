// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure system PTP information.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemPtp() *schema.Resource {
	return &schema.Resource{
		Description: "Configure system PTP information.",

		ReadContext: dataSourceSystemPtpRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"delay_mechanism": {
				Type:        schema.TypeString,
				Description: "End to end delay detection or peer to peer delay detection.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "PTP client will reply through this interface.",
				Computed:    true,
			},
			"mode": {
				Type:        schema.TypeString,
				Description: "Multicast transmission or hybrid transmission.",
				Computed:    true,
			},
			"request_interval": {
				Type:        schema.TypeInt,
				Description: "The delay request value is the logarithmic mean interval in seconds between the delay request messages sent by the slave to the master.",
				Computed:    true,
			},
			"server_interface": {
				Type:        schema.TypeList,
				Description: "FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delay_mechanism": {
							Type:        schema.TypeString,
							Description: "End to end delay detection or peer to peer delay detection.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"server_interface_name": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
					},
				},
			},
			"server_mode": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiGate PTP server mode. Your FortiGate becomes an PTP server for other devices on your network.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable setting the FortiGate system time by synchronizing with an PTP Server.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemPtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemPtp"

	o, err := c.Cmdb.ReadSystemPtp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemPtp dataSource: %v", err)
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

	diags := refreshObjectSystemPtp(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
