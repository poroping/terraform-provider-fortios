// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: SSL-VPN host check software.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceVpnSslWebHostCheckSoftware() *schema.Resource {
	return &schema.Resource{
		Description: "SSL-VPN host check software.",

		ReadContext: dataSourceVpnSslWebHostCheckSoftwareRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"check_item_list": {
				Type:        schema.TypeList,
				Description: "Check item list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Action.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID (0 - 4294967295).",
							Computed:    true,
						},
						"md5s": {
							Type:        schema.TypeList,
							Description: "MD5 checksum.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Description: "Hex string of MD5 checksum.",
										Computed:    true,
									},
								},
							},
						},
						"target": {
							Type:        schema.TypeString,
							Description: "Target.",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "Type.",
							Computed:    true,
						},
						"version": {
							Type:        schema.TypeString,
							Description: "Version.",
							Computed:    true,
						},
					},
				},
			},
			"guid": {
				Type:        schema.TypeString,
				Description: "Globally unique ID.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name.",
				Required:    true,
			},
			"os_type": {
				Type:        schema.TypeString,
				Description: "OS type.",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Type.",
				Computed:    true,
			},
			"version": {
				Type:        schema.TypeString,
				Description: "Version.",
				Computed:    true,
			},
		},
	}
}

func dataSourceVpnSslWebHostCheckSoftwareRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnSslWebHostCheckSoftware(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnSslWebHostCheckSoftware dataSource: %v", err)
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

	diags := refreshObjectVpnSslWebHostCheckSoftware(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
