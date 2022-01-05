// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Access Proxy virtual hosts.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceFirewallAccessProxyVirtualHost() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Access Proxy virtual hosts.",

		ReadContext: dataSourceFirewallAccessProxyVirtualHostRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"host": {
				Type:        schema.TypeString,
				Description: "The host name.",
				Computed:    true,
			},
			"host_type": {
				Type:        schema.TypeString,
				Description: "Type of host pattern.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Virtual host name.",
				Required:    true,
			},
			"ssl_certificate": {
				Type:        schema.TypeString,
				Description: "SSL certificate for this host.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallAccessProxyVirtualHostRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallAccessProxyVirtualHost(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallAccessProxyVirtualHost dataSource: %v", err)
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

	diags := refreshObjectFirewallAccessProxyVirtualHost(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
