// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Designate cache-service for wan-optimization and webcache.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceWanoptCacheService() *schema.Resource {
	return &schema.Resource{
		Description: "Designate cache-service for wan-optimization and webcache.",

		ReadContext: dataSourceWanoptCacheServiceRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"acceptable_connections": {
				Type:        schema.TypeString,
				Description: "Set strategy when accepting cache collaboration connection.",
				Computed:    true,
			},
			"collaboration": {
				Type:        schema.TypeString,
				Description: "Enable/disable cache-collaboration between cache-service clusters.",
				Computed:    true,
			},
			"device_id": {
				Type:        schema.TypeString,
				Description: "Set identifier for this cache device.",
				Computed:    true,
			},
			"dst_peer": {
				Type:        schema.TypeList,
				Description: "Modify cache-service destination peer list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_type": {
							Type:        schema.TypeInt,
							Description: "Set authentication type for this peer.",
							Computed:    true,
						},
						"device_id": {
							Type:        schema.TypeString,
							Description: "Device ID of this peer.",
							Computed:    true,
						},
						"encode_type": {
							Type:        schema.TypeInt,
							Description: "Set encode type for this peer.",
							Computed:    true,
						},
						"ip": {
							Type:        schema.TypeString,
							Description: "Set cluster IP address of this peer.",
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "Set priority for this peer.",
							Computed:    true,
						},
					},
				},
			},
			"prefer_scenario": {
				Type:        schema.TypeString,
				Description: "Set the preferred cache behavior towards the balance between latency and hit-ratio.",
				Computed:    true,
			},
			"src_peer": {
				Type:        schema.TypeList,
				Description: "Modify cache-service source peer list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_type": {
							Type:        schema.TypeInt,
							Description: "Set authentication type for this peer.",
							Computed:    true,
						},
						"device_id": {
							Type:        schema.TypeString,
							Description: "Device ID of this peer.",
							Computed:    true,
						},
						"encode_type": {
							Type:        schema.TypeInt,
							Description: "Set encode type for this peer.",
							Computed:    true,
						},
						"ip": {
							Type:        schema.TypeString,
							Description: "Set cluster IP address of this peer.",
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "Set priority for this peer.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWanoptCacheServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "WanoptCacheService"

	o, err := c.Cmdb.ReadWanoptCacheService(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WanoptCacheService dataSource: %v", err)
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

	diags := refreshObjectWanoptCacheService(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
