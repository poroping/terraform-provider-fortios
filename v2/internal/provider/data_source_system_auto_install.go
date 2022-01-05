// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.6,v7.0.2 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure USB auto installation.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemAutoInstall() *schema.Resource {
	return &schema.Resource{
		Description: "Configure USB auto installation.",

		ReadContext: dataSourceSystemAutoInstallRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"auto_install_config": {
				Type:        schema.TypeString,
				Description: "Enable/disable auto install the config in USB disk.",
				Computed:    true,
			},
			"auto_install_image": {
				Type:        schema.TypeString,
				Description: "Enable/disable auto install the image in USB disk.",
				Computed:    true,
			},
			"default_config_file": {
				Type:        schema.TypeString,
				Description: "Default config file name in USB disk.",
				Computed:    true,
			},
			"default_image_file": {
				Type:        schema.TypeString,
				Description: "Default image file name in USB disk.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemAutoInstallRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemAutoInstall"

	o, err := c.Cmdb.ReadSystemAutoInstall(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemAutoInstall dataSource: %v", err)
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

	diags := refreshObjectSystemAutoInstall(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
