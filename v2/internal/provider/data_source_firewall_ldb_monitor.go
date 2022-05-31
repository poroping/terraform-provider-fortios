// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure server load balancing health monitors.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallLdbMonitor() *schema.Resource {
	return &schema.Resource{
		Description: "Configure server load balancing health monitors.",

		ReadContext: dataSourceFirewallLdbMonitorRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"dns_match_ip": {
				Type:        schema.TypeString,
				Description: "Response IP expected from DNS server.",
				Computed:    true,
			},
			"dns_protocol": {
				Type:        schema.TypeString,
				Description: "Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP).",
				Computed:    true,
			},
			"dns_request_domain": {
				Type:        schema.TypeString,
				Description: "Fully qualified domain name to resolve for the DNS probe.",
				Computed:    true,
			},
			"http_get": {
				Type:        schema.TypeString,
				Description: "URL used to send a GET request to check the health of an HTTP server.",
				Computed:    true,
			},
			"http_match": {
				Type:        schema.TypeString,
				Description: "String to match the value expected in response to an HTTP-GET request.",
				Computed:    true,
			},
			"http_max_redirects": {
				Type:        schema.TypeInt,
				Description: "The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).",
				Computed:    true,
			},
			"interval": {
				Type:        schema.TypeInt,
				Description: "Time between health checks (5 - 65535 sec, default = 10).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Monitor name.",
				Required:    true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (0 - 65535, default = 0).",
				Computed:    true,
			},
			"retry": {
				Type:        schema.TypeInt,
				Description: "Number health check attempts before the server is considered down (1 - 255, default = 3).",
				Computed:    true,
			},
			"src_ip": {
				Type:        schema.TypeString,
				Description: "Source IP for ldb-monitor.",
				Computed:    true,
			},
			"timeout": {
				Type:        schema.TypeInt,
				Description: "Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Select the Monitor type used by the health check monitor to check the health of the server (PING | TCP | HTTP | HTTPS | DNS).",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallLdbMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallLdbMonitor(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallLdbMonitor dataSource: %v", err)
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

	diags := refreshObjectFirewallLdbMonitor(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
