// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure WiFi quality of service (QoS) profiles.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceWirelessControllerQosProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure WiFi quality of service (QoS) profiles.",

		ReadContext: dataSourceWirelessControllerQosProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"bandwidth_admission_control": {
				Type:        schema.TypeString,
				Description: "Enable/disable WMM bandwidth admission control.",
				Computed:    true,
			},
			"bandwidth_capacity": {
				Type:        schema.TypeInt,
				Description: "Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).",
				Computed:    true,
			},
			"burst": {
				Type:        schema.TypeString,
				Description: "Enable/disable client rate burst.",
				Computed:    true,
			},
			"call_admission_control": {
				Type:        schema.TypeString,
				Description: "Enable/disable WMM call admission control.",
				Computed:    true,
			},
			"call_capacity": {
				Type:        schema.TypeInt,
				Description: "Maximum number of Voice over WLAN (VoWLAN) phones allowed (0 - 60, default = 10).",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"downlink": {
				Type:        schema.TypeInt,
				Description: "Maximum downlink bandwidth for Virtual Access Points (VAPs) (0 - 2097152 Kbps, default = 0, 0 means no limit).",
				Computed:    true,
			},
			"downlink_sta": {
				Type:        schema.TypeInt,
				Description: "Maximum downlink bandwidth for clients (0 - 2097152 Kbps, default = 0, 0 means no limit).",
				Computed:    true,
			},
			"dscp_wmm_be": {
				Type:        schema.TypeList,
				Description: "DSCP mapping for best effort access (default = 0 24).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "DSCP WMM mapping numbers (0 - 63).",
							Computed:    true,
						},
					},
				},
			},
			"dscp_wmm_bk": {
				Type:        schema.TypeList,
				Description: "DSCP mapping for background access (default = 8 16).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "DSCP WMM mapping numbers (0 - 63).",
							Computed:    true,
						},
					},
				},
			},
			"dscp_wmm_mapping": {
				Type:        schema.TypeString,
				Description: "Enable/disable Differentiated Services Code Point (DSCP) mapping.",
				Computed:    true,
			},
			"dscp_wmm_vi": {
				Type:        schema.TypeList,
				Description: "DSCP mapping for video access (default = 32 40).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "DSCP WMM mapping numbers (0 - 63).",
							Computed:    true,
						},
					},
				},
			},
			"dscp_wmm_vo": {
				Type:        schema.TypeList,
				Description: "DSCP mapping for voice access (default = 48 56).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "DSCP WMM mapping numbers (0 - 63).",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "WiFi QoS profile name.",
				Required:    true,
			},
			"uplink": {
				Type:        schema.TypeInt,
				Description: "Maximum uplink bandwidth for Virtual Access Points (VAPs) (0 - 2097152 Kbps, default = 0, 0 means no limit).",
				Computed:    true,
			},
			"uplink_sta": {
				Type:        schema.TypeInt,
				Description: "Maximum uplink bandwidth for clients (0 - 2097152 Kbps, default = 0, 0 means no limit).",
				Computed:    true,
			},
			"wmm": {
				Type:        schema.TypeString,
				Description: "Enable/disable WiFi multi-media (WMM) control.",
				Computed:    true,
			},
			"wmm_be_dscp": {
				Type:        schema.TypeInt,
				Description: "DSCP marking for best effort access (default = 0).",
				Computed:    true,
			},
			"wmm_bk_dscp": {
				Type:        schema.TypeInt,
				Description: "DSCP marking for background access (default = 8).",
				Computed:    true,
			},
			"wmm_dscp_marking": {
				Type:        schema.TypeString,
				Description: "Enable/disable WMM Differentiated Services Code Point (DSCP) marking.",
				Computed:    true,
			},
			"wmm_uapsd": {
				Type:        schema.TypeString,
				Description: "Enable/disable WMM Unscheduled Automatic Power Save Delivery (U-APSD) power save mode.",
				Computed:    true,
			},
			"wmm_vi_dscp": {
				Type:        schema.TypeInt,
				Description: "DSCP marking for video access (default = 32).",
				Computed:    true,
			},
			"wmm_vo_dscp": {
				Type:        schema.TypeInt,
				Description: "DSCP marking for voice access (default = 48).",
				Computed:    true,
			},
		},
	}
}

func dataSourceWirelessControllerQosProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerQosProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerQosProfile dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerQosProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
