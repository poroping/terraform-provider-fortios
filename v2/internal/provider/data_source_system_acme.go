// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure ACME client.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemAcme() *schema.Resource {
	return &schema.Resource{
		Description: "Configure ACME client.",

		ReadContext: dataSourceSystemAcmeRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"accounts": {
				Type:        schema.TypeList,
				Description: "ACME accounts list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ca_url": {
							Type:        schema.TypeString,
							Description: "Account ca_url.",
							Computed:    true,
						},
						"email": {
							Type:        schema.TypeString,
							Description: "Account email.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeString,
							Description: "Account id.",
							Computed:    true,
						},
						"privatekey": {
							Type:        schema.TypeString,
							Description: "Account Private Key.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Account status.",
							Computed:    true,
						},
						"url": {
							Type:        schema.TypeString,
							Description: "Account url.",
							Computed:    true,
						},
					},
				},
			},
			"interface": {
				Type:        schema.TypeList,
				Description: "Interface(s) on which the ACME client will listen for challenges.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemAcmeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemAcme"

	o, err := c.Cmdb.ReadSystemAcme(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemAcme dataSource: %v", err)
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

	diags := refreshObjectSystemAcme(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
