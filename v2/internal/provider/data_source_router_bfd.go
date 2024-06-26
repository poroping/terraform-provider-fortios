// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure BFD.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceRouterBfd() *schema.Resource {
	return &schema.Resource{
		Description: "Configure BFD.",

		ReadContext: dataSourceRouterBfdRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"multihop_template": {
				Type:        schema.TypeList,
				Description: "BFD multi-hop template table.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_mode": {
							Type:        schema.TypeString,
							Description: "Authentication mode.",
							Computed:    true,
						},
						"bfd_desired_min_tx": {
							Type:        schema.TypeInt,
							Description: "BFD desired minimal transmit interval (milliseconds).",
							Computed:    true,
						},
						"bfd_detect_mult": {
							Type:        schema.TypeInt,
							Description: "BFD detection multiplier.",
							Computed:    true,
						},
						"bfd_required_min_rx": {
							Type:        schema.TypeInt,
							Description: "BFD required minimal receive interval (milliseconds).",
							Computed:    true,
						},
						"dst": {
							Type:        schema.TypeString,
							Description: "Destination prefix.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"md5_key": {
							Type:        schema.TypeString,
							Description: "MD5 key of key ID 1.",
							Computed:    true,
							Sensitive:   true,
						},
						"src": {
							Type:        schema.TypeString,
							Description: "Source prefix.",
							Computed:    true,
						},
					},
				},
			},
			"neighbor": {
				Type:        schema.TypeList,
				Description: "Neighbor.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
						"ip": {
							Type:        schema.TypeString,
							Description: "IPv4 address of the BFD neighbor.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterBfdRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "RouterBfd"

	o, err := c.Cmdb.ReadRouterBfd(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterBfd dataSource: %v", err)
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

	diags := refreshObjectRouterBfd(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
