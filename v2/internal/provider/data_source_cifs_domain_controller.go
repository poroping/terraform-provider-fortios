// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Define known domain controller servers.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceCifsDomainController() *schema.Resource {
	return &schema.Resource{
		Description: "Define known domain controller servers.",

		ReadContext: dataSourceCifsDomainControllerRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"domain_name": {
				Type:        schema.TypeString,
				Description: "Fully qualified domain name (FQDN). E.g. 'EXAMPLE.COM'.",
				Computed:    true,
			},
			"ip": {
				Type:        schema.TypeString,
				Description: "IPv4 server address.",
				Computed:    true,
			},
			"ip6": {
				Type:        schema.TypeString,
				Description: "IPv6 server address.",
				Computed:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "Password for specified username.",
				Computed:    true,
				Sensitive:   true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "Port number of service. Port number 0 indicates automatic discovery.",
				Computed:    true,
			},
			"server_name": {
				Type:        schema.TypeString,
				Description: "Name of the server to connect to.",
				Required:    true,
			},
			"username": {
				Type:        schema.TypeString,
				Description: "User name to sign in with. Must have proper permissions for service.",
				Computed:    true,
			},
		},
	}
}

func dataSourceCifsDomainControllerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadCifsDomainController(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading CifsDomainController dataSource: %v", err)
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

	diags := refreshObjectCifsDomainController(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
