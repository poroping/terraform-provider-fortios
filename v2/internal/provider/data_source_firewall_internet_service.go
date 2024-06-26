// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Show Internet Service application.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallInternetService() *schema.Resource {
	return &schema.Resource{
		Description: "Show Internet Service application.",

		ReadContext: dataSourceFirewallInternetServiceRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"database": {
				Type:        schema.TypeString,
				Description: "Database name this Internet Service belongs to.",
				Computed:    true,
			},
			"direction": {
				Type:        schema.TypeString,
				Description: "How this service may be used in a firewall policy (source, destination or both).",
				Computed:    true,
			},
			"extra_ip_range_number": {
				Type:        schema.TypeInt,
				Description: "Extra number of IPv4 ranges.",
				Computed:    true,
			},
			"extra_ip6_range_number": {
				Type:        schema.TypeInt,
				Description: "Extra number of IPv6 ranges.",
				Computed:    true,
			},
			"icon_id": {
				Type:        schema.TypeInt,
				Description: "Icon ID of Internet Service.",
				Computed:    true,
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "Internet Service ID.",
				Computed:    true,
			},
			"ip_number": {
				Type:        schema.TypeInt,
				Description: "Total number of IPv4 addresses.",
				Computed:    true,
			},
			"ip_range_number": {
				Type:        schema.TypeInt,
				Description: "Number of IPv4 ranges.",
				Computed:    true,
			},
			"ip6_range_number": {
				Type:        schema.TypeInt,
				Description: "Number of IPv6 ranges.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Internet Service name.",
				Computed:    true,
			},
			"obsolete": {
				Type:        schema.TypeInt,
				Description: "Indicates whether the Internet Service can be used.",
				Computed:    true,
			},
			"reputation": {
				Type:        schema.TypeInt,
				Description: "Reputation level of the Internet Service.",
				Computed:    true,
			},
			"singularity": {
				Type:        schema.TypeInt,
				Description: "Singular level of the Internet Service.",
				Computed:    true,
			},
			"sld_id": {
				Type:        schema.TypeInt,
				Description: "Second Level Domain.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallInternetServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallInternetService(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallInternetService dataSource: %v", err)
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

	diags := refreshObjectFirewallInternetService(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
