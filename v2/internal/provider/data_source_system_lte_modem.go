// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.6,v7.0.2 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure USB LTE/WIMAX devices.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemLteModem() *schema.Resource {
	return &schema.Resource{
		Description: "Configure USB LTE/WIMAX devices.",

		ReadContext: dataSourceSystemLteModemRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"apn": {
				Type:        schema.TypeString,
				Description: "Login APN string for PDP-IP packet data calls.",
				Computed:    true,
			},
			"authtype": {
				Type:        schema.TypeString,
				Description: "Authentication type for PDP-IP packet data calls.",
				Computed:    true,
			},
			"extra_init": {
				Type:        schema.TypeString,
				Description: "Extra initialization string for USB LTE/WIMAX devices.",
				Computed:    true,
			},
			"holddown_timer": {
				Type:        schema.TypeInt,
				Description: "Hold down timer (10 - 60 sec).",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "The interface that the modem is acting as a redundant interface for.",
				Computed:    true,
			},
			"mode": {
				Type:        schema.TypeString,
				Description: "Modem operation mode.",
				Computed:    true,
			},
			"modem_port": {
				Type:        schema.TypeInt,
				Description: "Modem port index (0 - 20).",
				Computed:    true,
			},
			"passwd": {
				Type:        schema.TypeString,
				Description: "Authentication password for PDP-IP packet data calls.",
				Computed:    true,
				Sensitive:   true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable USB LTE/WIMAX device.",
				Computed:    true,
			},
			"username": {
				Type:        schema.TypeString,
				Description: "Authentication username for PDP-IP packet data calls.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemLteModemRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemLteModem"

	o, err := c.Cmdb.ReadSystemLteModem(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemLteModem dataSource: %v", err)
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

	diags := refreshObjectSystemLteModem(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
