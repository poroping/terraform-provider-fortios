// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure system probe response.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemProbeResponse() *schema.Resource {
	return &schema.Resource{
		Description: "Configure system probe response.",

		ReadContext: dataSourceSystemProbeResponseRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"http_probe_value": {
				Type:        schema.TypeString,
				Description: "Value to respond to the monitoring server.",
				Computed:    true,
			},
			"mode": {
				Type:        schema.TypeString,
				Description: "SLA response mode.",
				Computed:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "TWAMP responder password in authentication mode.",
				Computed:    true,
				Sensitive:   true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "Port number to response.",
				Computed:    true,
			},
			"security_mode": {
				Type:        schema.TypeString,
				Description: "TWAMP responder security mode.",
				Computed:    true,
			},
			"timeout": {
				Type:        schema.TypeInt,
				Description: "An inactivity timer for a twamp test session.",
				Computed:    true,
			},
			"ttl_mode": {
				Type:        schema.TypeString,
				Description: "Mode for TWAMP packet TTL modification.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemProbeResponseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemProbeResponse"

	o, err := c.Cmdb.ReadSystemProbeResponse(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemProbeResponse dataSource: %v", err)
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

	diags := refreshObjectSystemProbeResponse(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
