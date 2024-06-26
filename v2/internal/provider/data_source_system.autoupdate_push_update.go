// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure push updates.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemAutoupdatePushUpdate() *schema.Resource {
	return &schema.Resource{
		Description: "Configure push updates.",

		ReadContext: dataSourceSystemAutoupdatePushUpdateRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"address": {
				Type:        schema.TypeString,
				Description: "IPv4 or IPv6 address used by FortiGuard servers to send push updates to this FortiGate.",
				Computed:    true,
			},
			"override": {
				Type:        schema.TypeString,
				Description: "Enable/disable push update override server.",
				Computed:    true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "Push update override port. (Do not overlap with other service ports)",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable push updates.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemAutoupdatePushUpdateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemAutoupdatePushUpdate"

	o, err := c.Cmdb.ReadSystemAutoupdatePushUpdate(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemAutoupdatePushUpdate dataSource: %v", err)
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

	diags := refreshObjectSystemAutoupdatePushUpdate(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
