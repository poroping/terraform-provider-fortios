// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure forward-server addresses.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWebProxyForwardServer() *schema.Resource {
	return &schema.Resource{
		Description: "Configure forward-server addresses.",

		ReadContext: dataSourceWebProxyForwardServerRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"addr_type": {
				Type:        schema.TypeString,
				Description: "Address type of the forwarding proxy server: IP or FQDN.",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"fqdn": {
				Type:        schema.TypeString,
				Description: "Forward server Fully Qualified Domain Name (FQDN).",
				Computed:    true,
			},
			"healthcheck": {
				Type:        schema.TypeString,
				Description: "Enable/disable forward server health checking. Attempts to connect through the remote forwarding server to a destination to verify that the forwarding server is operating normally.",
				Computed:    true,
			},
			"ip": {
				Type:        schema.TypeString,
				Description: "Forward proxy server IP address.",
				Computed:    true,
			},
			"monitor": {
				Type:        schema.TypeString,
				Description: "URL for forward server health check monitoring (default = http://www.google.com).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Server name.",
				Required:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "HTTP authentication password.",
				Computed:    true,
				Sensitive:   true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).",
				Computed:    true,
			},
			"server_down_option": {
				Type:        schema.TypeString,
				Description: "Action to take when the forward server is found to be down: block sessions until the server is back up or pass sessions to their destination.",
				Computed:    true,
			},
			"username": {
				Type:        schema.TypeString,
				Description: "HTTP authentication user name.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWebProxyForwardServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWebProxyForwardServer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebProxyForwardServer dataSource: %v", err)
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

	diags := refreshObjectWebProxyForwardServer(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
