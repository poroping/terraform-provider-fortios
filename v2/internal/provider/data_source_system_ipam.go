// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IP address management services.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemIpam() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IP address management services.",

		ReadContext: dataSourceSystemIpamRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"pool_subnet": {
				Type:        schema.TypeString,
				Description: "Configure IPAM pool subnet, Class A - Class B subnet.",
				Computed:    true,
			},
			"pools": {
				Type:        schema.TypeList,
				Description: "Configure IPAM pools.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:        schema.TypeString,
							Description: "Description.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "IPAM pool name.",
							Computed:    true,
						},
						"subnet": {
							Type:        schema.TypeString,
							Description: "Configure IPAM pool subnet, Class A - Class B subnet.",
							Computed:    true,
						},
					},
				},
			},
			"rules": {
				Type:        schema.TypeList,
				Description: "Configure IPAM allocation rules.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:        schema.TypeString,
							Description: "Description.",
							Computed:    true,
						},
						"device": {
							Type:        schema.TypeList,
							Description: "Configure serial number or wildcard of FortiGate to match.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "FortiGate serial number or wildcard.",
										Computed:    true,
									},
								},
							},
						},
						"dhcp": {
							Type:        schema.TypeString,
							Description: "Enable/disable DHCP server for matching IPAM interfaces.",
							Computed:    true,
						},
						"interface": {
							Type:        schema.TypeList,
							Description: "Configure name or wildcard of interface to match.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Interface name or wildcard.",
										Computed:    true,
									},
								},
							},
						},
						"name": {
							Type:        schema.TypeString,
							Description: "IPAM rule name.",
							Computed:    true,
						},
						"pool": {
							Type:        schema.TypeList,
							Description: "Configure name of IPAM pool to use.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Ipam pool name.",
										Computed:    true,
									},
								},
							},
						},
						"role": {
							Type:        schema.TypeString,
							Description: "Configure role of interface to match.",
							Computed:    true,
						},
					},
				},
			},
			"server_type": {
				Type:        schema.TypeString,
				Description: "Configure the type of IPAM server to use.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable IP address management services.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemIpamRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemIpam"

	o, err := c.Cmdb.ReadSystemIpam(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemIpam dataSource: %v", err)
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

	diags := refreshObjectSystemIpam(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
