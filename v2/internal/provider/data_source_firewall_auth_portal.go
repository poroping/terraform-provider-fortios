// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure firewall authentication portals.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceFirewallAuthPortal() *schema.Resource {
	return &schema.Resource{
		Description: "Configure firewall authentication portals.",

		ReadContext: dataSourceFirewallAuthPortalRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"groups": {
				Type:        schema.TypeList,
				Description: "Firewall user groups permitted to authenticate through this portal. Separate group names with spaces.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Group name.",
							Computed:    true,
						},
					},
				},
			},
			"identity_based_route": {
				Type:        schema.TypeString,
				Description: "Name of the identity-based route that applies to this portal.",
				Computed:    true,
			},
			"portal_addr": {
				Type:        schema.TypeString,
				Description: "Address (or FQDN) of the authentication portal.",
				Computed:    true,
			},
			"portal_addr6": {
				Type:        schema.TypeString,
				Description: "IPv6 address (or FQDN) of authentication portal.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallAuthPortalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "FirewallAuthPortal"

	o, err := c.Cmdb.ReadFirewallAuthPortal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallAuthPortal dataSource: %v", err)
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

	diags := refreshObjectFirewallAuthPortal(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
