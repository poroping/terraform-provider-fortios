// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: VDOM wireless controller configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceWirelessControllerSetting() *schema.Resource {
	return &schema.Resource{
		Description: "VDOM wireless controller configuration.",

		ReadContext: dataSourceWirelessControllerSettingRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"account_id": {
				Type:        schema.TypeString,
				Description: "FortiCloud customer account ID.",
				Computed:    true,
			},
			"country": {
				Type:        schema.TypeString,
				Description: "Country or region in which the FortiGate is located. The country determines the 802.11 bands and channels that are available.",
				Computed:    true,
			},
			"darrp_optimize": {
				Type:        schema.TypeInt,
				Description: "Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).",
				Computed:    true,
			},
			"darrp_optimize_schedules": {
				Type:        schema.TypeList,
				Description: "Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Schedule name.",
							Computed:    true,
						},
					},
				},
			},
			"device_holdoff": {
				Type:        schema.TypeInt,
				Description: "Lower limit of creation time of device for identification in minutes (0 - 60, default = 5).",
				Computed:    true,
			},
			"device_idle": {
				Type:        schema.TypeInt,
				Description: "Upper limit of idle time of device for identification in minutes (0 - 14400, default = 1440).",
				Computed:    true,
			},
			"device_weight": {
				Type:        schema.TypeInt,
				Description: "Upper limit of confidence of device for identification (0 - 255, default = 1, 0 = disable).",
				Computed:    true,
			},
			"duplicate_ssid": {
				Type:        schema.TypeString,
				Description: "Enable/disable allowing Virtual Access Points (VAPs) to use the same SSID name in the same VDOM.",
				Computed:    true,
			},
			"fake_ssid_action": {
				Type:        schema.TypeString,
				Description: "Actions taken for detected fake SSID.",
				Computed:    true,
			},
			"fapc_compatibility": {
				Type:        schema.TypeString,
				Description: "Enable/disable FAP-C series compatibility.",
				Computed:    true,
			},
			"firmware_provision_on_authorization": {
				Type:        schema.TypeString,
				Description: "Enable/disable automatic provisioning of latest firmware on authorization.",
				Computed:    true,
			},
			"offending_ssid": {
				Type:        schema.TypeList,
				Description: "Configure offending SSID.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Actions taken for detected offending SSID.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"ssid_pattern": {
							Type:        schema.TypeString,
							Description: "Define offending SSID pattern (case insensitive). For example, word, word*, *word, wo*rd.",
							Computed:    true,
						},
					},
				},
			},
			"phishing_ssid_detect": {
				Type:        schema.TypeString,
				Description: "Enable/disable phishing SSID detection.",
				Computed:    true,
			},
			"wfa_compatibility": {
				Type:        schema.TypeString,
				Description: "Enable/disable WFA compatibility.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWirelessControllerSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "WirelessControllerSetting"

	o, err := c.Cmdb.ReadWirelessControllerSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerSetting dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
