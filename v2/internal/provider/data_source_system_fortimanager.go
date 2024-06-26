// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiManager.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemFortimanager() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiManager.",

		ReadContext: dataSourceSystemFortimanagerRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"central_management": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiManager central management.",
				Computed:    true,
			},
			"central_mgmt_auto_backup": {
				Type:        schema.TypeString,
				Description: "Enable/disable central management auto backup.",
				Computed:    true,
			},
			"central_mgmt_schedule_config_restore": {
				Type:        schema.TypeString,
				Description: "Enable/disable central management schedule config restore.",
				Computed:    true,
			},
			"central_mgmt_schedule_script_restore": {
				Type:        schema.TypeString,
				Description: "Enable/disable central management schedule script restore.",
				Computed:    true,
			},
			"ip": {
				Type:        schema.TypeString,
				Description: "IP address.",
				Computed:    true,
			},
			"ipsec": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiManager IPsec tunnel.",
				Computed:    true,
			},
			"vdom": {
				Type:        schema.TypeString,
				Description: "Virtual domain name.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemFortimanagerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemFortimanager"

	o, err := c.Cmdb.ReadSystemFortimanager(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemFortimanager dataSource: %v", err)
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

	diags := refreshObjectSystemFortimanager(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
