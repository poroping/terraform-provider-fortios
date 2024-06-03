// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure ICAP servers.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceIcapServer() *schema.Resource {
	return &schema.Resource{
		Description: "Configure ICAP servers.",

		ReadContext: dataSourceIcapServerRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"addr_type": {
				Type:        schema.TypeString,
				Description: "Address type of the remote ICAP server: IPv4, IPv6 or FQDN.",
				Computed:    true,
			},
			"fqdn": {
				Type:        schema.TypeString,
				Description: "ICAP remote server Fully Qualified Domain Name (FQDN).",
				Computed:    true,
			},
			"healthcheck": {
				Type:        schema.TypeString,
				Description: "Enable/disable ICAP remote server health checking. Attempts to connect to the remote ICAP server to verify that the server is operating normally.",
				Computed:    true,
			},
			"healthcheck_service": {
				Type:        schema.TypeString,
				Description: "ICAP Service name to use for health checks.",
				Computed:    true,
			},
			"ip_address": {
				Type:        schema.TypeString,
				Description: "IPv4 address of the ICAP server.",
				Computed:    true,
			},
			"ip_version": {
				Type:        schema.TypeString,
				Description: "IP version.",
				Computed:    true,
			},
			"ip6_address": {
				Type:        schema.TypeString,
				Description: "IPv6 address of the ICAP server.",
				Computed:    true,
			},
			"max_connections": {
				Type:        schema.TypeInt,
				Description: "Maximum number of concurrent connections to ICAP server (unlimited = 0, default = 100). Must not be less than wad-worker-count.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Server name.",
				Required:    true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "ICAP server port.",
				Computed:    true,
			},
			"secure": {
				Type:        schema.TypeString,
				Description: "Enable/disable secure connection to ICAP server.",
				Computed:    true,
			},
			"ssl_cert": {
				Type:        schema.TypeString,
				Description: "CA certificate name.",
				Computed:    true,
			},
		},
	}
}

func dataSourceIcapServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadIcapServer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading IcapServer dataSource: %v", err)
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

	diags := refreshObjectIcapServer(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
