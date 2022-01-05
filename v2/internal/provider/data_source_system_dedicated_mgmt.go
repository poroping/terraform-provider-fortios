// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure dedicated management.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemDedicatedMgmt() *schema.Resource {
	return &schema.Resource{
		Description: "Configure dedicated management.",

		ReadContext: dataSourceSystemDedicatedMgmtRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"default_gateway": {
				Type:        schema.TypeString,
				Description: "Default gateway for dedicated management interface.",
				Computed:    true,
			},
			"dhcp_end_ip": {
				Type:        schema.TypeString,
				Description: "DHCP end IP for dedicated management.",
				Computed:    true,
			},
			"dhcp_netmask": {
				Type:        schema.TypeString,
				Description: "DHCP netmask.",
				Computed:    true,
			},
			"dhcp_server": {
				Type:        schema.TypeString,
				Description: "Enable/disable DHCP server on management interface.",
				Computed:    true,
			},
			"dhcp_start_ip": {
				Type:        schema.TypeString,
				Description: "DHCP start IP for dedicated management.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Dedicated management interface.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable dedicated management.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemDedicatedMgmtRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemDedicatedMgmt"

	o, err := c.Cmdb.ReadSystemDedicatedMgmt(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemDedicatedMgmt dataSource: %v", err)
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

	diags := refreshObjectSystemDedicatedMgmt(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
