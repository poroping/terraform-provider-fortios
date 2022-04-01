// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure MS Exchange server entries.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceUserExchange() *schema.Resource {
	return &schema.Resource{
		Description: "Configure MS Exchange server entries.",

		ReadContext: dataSourceUserExchangeRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"auth_level": {
				Type:        schema.TypeString,
				Description: "Authentication security level used for the RPC protocol layer.",
				Computed:    true,
			},
			"auth_type": {
				Type:        schema.TypeString,
				Description: "Authentication security type used for the RPC protocol layer.",
				Computed:    true,
			},
			"auto_discover_kdc": {
				Type:        schema.TypeString,
				Description: "Enable/disable automatic discovery of KDC IP addresses.",
				Computed:    true,
			},
			"connect_protocol": {
				Type:        schema.TypeString,
				Description: "Connection protocol used to connect to MS Exchange service.",
				Computed:    true,
			},
			"domain_name": {
				Type:        schema.TypeString,
				Description: "MS Exchange server fully qualified domain name.",
				Computed:    true,
			},
			"http_auth_type": {
				Type:        schema.TypeString,
				Description: "Authentication security type used for the HTTP transport.",
				Computed:    true,
			},
			"ip": {
				Type:        schema.TypeString,
				Description: "Server IPv4 address.",
				Computed:    true,
			},
			"kdc_ip": {
				Type:        schema.TypeList,
				Description: "KDC IPv4 addresses for Kerberos authentication.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4": {
							Type:        schema.TypeString,
							Description: "KDC IPv4 addresses for Kerberos authentication.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "MS Exchange server entry name.",
				Required:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "Password for the specified username.",
				Computed:    true,
				Sensitive:   true,
			},
			"server_name": {
				Type:        schema.TypeString,
				Description: "MS Exchange server hostname.",
				Computed:    true,
			},
			"ssl_min_proto_version": {
				Type:        schema.TypeString,
				Description: "Minimum SSL/TLS protocol version for HTTPS transport (default is to follow system global setting).",
				Computed:    true,
			},
			"username": {
				Type:        schema.TypeString,
				Description: "User name used to sign in to the server. Must have proper permissions for service.",
				Computed:    true,
			},
		},
	}
}

func dataSourceUserExchangeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("name")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadUserExchange(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserExchange dataSource: %v", err)
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

	diags := refreshObjectUserExchange(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
