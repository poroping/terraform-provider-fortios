// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure/list NAC devices learned on the managed FortiSwitch ports which matches NAC policy.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSwitchControllerNacDevice() *schema.Resource {
	return &schema.Resource{
		Description: "Configure/list NAC devices learned on the managed FortiSwitch ports which matches NAC policy.",

		ReadContext: dataSourceSwitchControllerNacDeviceRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description for the learned NAC device.",
				Computed:    true,
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "Device ID.",
				Computed:    true,
			},
			"last_known_port": {
				Type:        schema.TypeString,
				Description: "Managed FortiSwitch port where NAC device is last learned.",
				Computed:    true,
			},
			"last_known_switch": {
				Type:        schema.TypeString,
				Description: "Managed FortiSwitch where NAC device is last learned.",
				Computed:    true,
			},
			"last_seen": {
				Type:        schema.TypeInt,
				Description: "Device last seen.",
				Computed:    true,
			},
			"mac": {
				Type:        schema.TypeString,
				Description: "MAC address of the learned NAC device.",
				Computed:    true,
			},
			"mac_policy": {
				Type:        schema.TypeString,
				Description: "MAC policy to be applied on this learned NAC device.",
				Computed:    true,
			},
			"matched_nac_policy": {
				Type:        schema.TypeString,
				Description: "Matched NAC policy for the learned NAC device.",
				Computed:    true,
			},
			"port_policy": {
				Type:        schema.TypeString,
				Description: "Port policy to be applied on this learned NAC device.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Status of the learned NAC device. Set enable to authorize the NAC device.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSwitchControllerNacDeviceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("fosid")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadSwitchControllerNacDevice(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerNacDevice dataSource: %v", err)
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

	diags := refreshObjectSwitchControllerNacDevice(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
