// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure log event filters.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceLogEventfilter() *schema.Resource {
	return &schema.Resource{
		Description: "Configure log event filters.",

		ReadContext: dataSourceLogEventfilterRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"cifs": {
				Type:        schema.TypeString,
				Description: "Enable/disable CIFS logging.",
				Computed:    true,
			},
			"connector": {
				Type:        schema.TypeString,
				Description: "Enable/disable SDN connector logging.",
				Computed:    true,
			},
			"endpoint": {
				Type:        schema.TypeString,
				Description: "Enable/disable endpoint event logging.",
				Computed:    true,
			},
			"event": {
				Type:        schema.TypeString,
				Description: "Enable/disable event logging.",
				Computed:    true,
			},
			"fortiextender": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiExtender logging.",
				Computed:    true,
			},
			"ha": {
				Type:        schema.TypeString,
				Description: "Enable/disable ha event logging.",
				Computed:    true,
			},
			"rest_api": {
				Type:        schema.TypeString,
				Description: "Enable/disable REST API logging.",
				Computed:    true,
			},
			"router": {
				Type:        schema.TypeString,
				Description: "Enable/disable router event logging.",
				Computed:    true,
			},
			"sdwan": {
				Type:        schema.TypeString,
				Description: "Enable/disable SD-WAN logging.",
				Computed:    true,
			},
			"security_rating": {
				Type:        schema.TypeString,
				Description: "Enable/disable Security Rating result logging.",
				Computed:    true,
			},
			"switch_controller": {
				Type:        schema.TypeString,
				Description: "Enable/disable Switch-Controller logging.",
				Computed:    true,
			},
			"system": {
				Type:        schema.TypeString,
				Description: "Enable/disable system event logging.",
				Computed:    true,
			},
			"user": {
				Type:        schema.TypeString,
				Description: "Enable/disable user authentication event logging.",
				Computed:    true,
			},
			"vpn": {
				Type:        schema.TypeString,
				Description: "Enable/disable VPN event logging.",
				Computed:    true,
			},
			"wan_opt": {
				Type:        schema.TypeString,
				Description: "Enable/disable WAN optimization event logging.",
				Computed:    true,
			},
			"wireless_activity": {
				Type:        schema.TypeString,
				Description: "Enable/disable wireless event logging.",
				Computed:    true,
			},
		},
	}
}

func dataSourceLogEventfilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "LogEventfilter"

	o, err := c.Cmdb.ReadLogEventfilter(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogEventfilter dataSource: %v", err)
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

	diags := refreshObjectLogEventfilter(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
