// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiGuard Web Filter administrative overrides.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceWebfilterOverride() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiGuard Web Filter administrative overrides.",

		ReadContext: dataSourceWebfilterOverrideRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"expires": {
				Type:        schema.TypeString,
				Description: "Override expiration date and time, from 5 minutes to 365 from now (format: yyyy/mm/dd hh:mm:ss).",
				Computed:    true,
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "Override rule ID.",
				Computed:    true,
			},
			"initiator": {
				Type:        schema.TypeString,
				Description: "Initiating user of override (read-only setting).",
				Computed:    true,
			},
			"ip": {
				Type:        schema.TypeString,
				Description: "IPv4 address which the override applies.",
				Computed:    true,
			},
			"ip6": {
				Type:        schema.TypeString,
				Description: "IPv6 address which the override applies.",
				Computed:    true,
			},
			"new_profile": {
				Type:        schema.TypeString,
				Description: "Name of the new web filter profile used by the override.",
				Computed:    true,
			},
			"old_profile": {
				Type:        schema.TypeString,
				Description: "Name of the web filter profile which the override applies.",
				Computed:    true,
			},
			"scope": {
				Type:        schema.TypeString,
				Description: "Override either the specific user, user group, IPv4 address, or IPv6 address.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable override rule.",
				Computed:    true,
			},
			"user": {
				Type:        schema.TypeString,
				Description: "Name of the user which the override applies.",
				Computed:    true,
			},
			"user_group": {
				Type:        schema.TypeString,
				Description: "Specify the user group for which the override applies.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWebfilterOverrideRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWebfilterOverride(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebfilterOverride dataSource: %v", err)
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

	diags := refreshObjectWebfilterOverride(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
