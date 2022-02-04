// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure UTM (Unified Threat Management) profile.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWirelessControllerUtmProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure UTM (Unified Threat Management) profile.",

		ReadContext: dataSourceWirelessControllerUtmProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"antivirus_profile": {
				Type:        schema.TypeString,
				Description: "AntiVirus profile name.",
				Computed:    true,
			},
			"application_list": {
				Type:        schema.TypeString,
				Description: "Application control list name.",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"ips_sensor": {
				Type:        schema.TypeString,
				Description: "IPS sensor name.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "UTM profile name.",
				Required:    true,
			},
			"scan_botnet_connections": {
				Type:        schema.TypeString,
				Description: "Block or monitor connections to Botnet servers or disable Botnet scanning.",
				Computed:    true,
			},
			"utm_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable UTM logging.",
				Computed:    true,
			},
			"webfilter_profile": {
				Type:        schema.TypeString,
				Description: "WebFilter profile name.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWirelessControllerUtmProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("name")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadWirelessControllerUtmProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerUtmProfile dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerUtmProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
