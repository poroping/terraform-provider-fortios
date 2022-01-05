// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure geographical location mapping for IP address(es) to override mappings from FortiGuard.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemGeoipOverride() *schema.Resource {
	return &schema.Resource{
		Description: "Configure geographical location mapping for IP address(es) to override mappings from FortiGuard.",

		ReadContext: dataSourceSystemGeoipOverrideRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"country_id": {
				Type:        schema.TypeString,
				Description: "Two character Country ID code.",
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description.",
				Computed:    true,
			},
			"ip_range": {
				Type:        schema.TypeList,
				Description: "Table of IP ranges assigned to country.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": {
							Type:        schema.TypeString,
							Description: "Ending IP address, inclusive, of the address range (format: xxx.xxx.xxx.xxx).",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID of individual entry in the IP range table.",
							Computed:    true,
						},
						"start_ip": {
							Type:        schema.TypeString,
							Description: "Starting IP address, inclusive, of the address range (format: xxx.xxx.xxx.xxx).",
							Computed:    true,
						},
					},
				},
			},
			"ip6_range": {
				Type:        schema.TypeList,
				Description: "Table of IPv6 ranges assigned to country.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": {
							Type:        schema.TypeString,
							Description: "Ending IP address, inclusive, of the address range (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID of individual entry in the IPv6 range table.",
							Computed:    true,
						},
						"start_ip": {
							Type:        schema.TypeString,
							Description: "Starting IP address, inclusive, of the address range (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Location name.",
				Required:    true,
			},
		},
	}
}

func dataSourceSystemGeoipOverrideRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemGeoipOverride(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemGeoipOverride dataSource: %v", err)
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

	diags := refreshObjectSystemGeoipOverride(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
