// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiGuard services.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemFortiguard() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiGuard services.",

		ReadContext: dataSourceSystemFortiguardRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"antispam_cache": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiGuard antispam request caching. Uses a small amount of memory but improves performance.",
				Computed:    true,
			},
			"antispam_cache_mpercent": {
				Type:        schema.TypeInt,
				Description: "Maximum percent of FortiGate memory the antispam cache is allowed to use (1 - 15%).",
				Computed:    true,
			},
			"antispam_cache_ttl": {
				Type:        schema.TypeInt,
				Description: "Time-to-live for antispam cache entries in seconds (300 - 86400). Lower times reduce the cache size. Higher times may improve performance since the cache will have more entries.",
				Computed:    true,
			},
			"antispam_expiration": {
				Type:        schema.TypeInt,
				Description: "Expiration date of the FortiGuard antispam contract.",
				Computed:    true,
			},
			"antispam_force_off": {
				Type:        schema.TypeString,
				Description: "Enable/disable turning off the FortiGuard antispam service.",
				Computed:    true,
			},
			"antispam_license": {
				Type:        schema.TypeInt,
				Description: "Interval of time between license checks for the FortiGuard antispam contract.",
				Computed:    true,
			},
			"antispam_timeout": {
				Type:        schema.TypeInt,
				Description: "Antispam query time out (1 - 30 sec, default = 7).",
				Computed:    true,
			},
			"anycast_sdns_server_ip": {
				Type:        schema.TypeString,
				Description: "IP address of the FortiGuard anycast DNS rating server.",
				Computed:    true,
			},
			"anycast_sdns_server_port": {
				Type:        schema.TypeInt,
				Description: "Port to connect to on the FortiGuard anycast DNS rating server.",
				Computed:    true,
			},
			"auto_join_forticloud": {
				Type:        schema.TypeString,
				Description: "Automatically connect to and login to FortiCloud.",
				Computed:    true,
			},
			"ddns_server_ip": {
				Type:        schema.TypeString,
				Description: "IP address of the FortiDDNS server.",
				Computed:    true,
			},
			"ddns_server_ip6": {
				Type:        schema.TypeString,
				Description: "IPv6 address of the FortiDDNS server.",
				Computed:    true,
			},
			"ddns_server_port": {
				Type:        schema.TypeInt,
				Description: "Port used to communicate with FortiDDNS servers.",
				Computed:    true,
			},
			"fortiguard_anycast": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of FortiGuard's Anycast network.",
				Computed:    true,
			},
			"fortiguard_anycast_source": {
				Type:        schema.TypeString,
				Description: "Configure which of Fortinet's servers to provide FortiGuard services in FortiGuard's anycast network. Default is Fortinet.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Specify outgoing interface to reach server.",
				Computed:    true,
			},
			"interface_select_method": {
				Type:        schema.TypeString,
				Description: "Specify how to select outgoing interface to reach server.",
				Computed:    true,
			},
			"load_balance_servers": {
				Type:        schema.TypeInt,
				Description: "Number of servers to alternate between as first FortiGuard option.",
				Computed:    true,
			},
			"outbreak_prevention_cache": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiGuard Virus Outbreak Prevention cache.",
				Computed:    true,
			},
			"outbreak_prevention_cache_mpercent": {
				Type:        schema.TypeInt,
				Description: "Maximum percent of memory FortiGuard Virus Outbreak Prevention cache can use (1 - 15%, default = 2).",
				Computed:    true,
			},
			"outbreak_prevention_cache_ttl": {
				Type:        schema.TypeInt,
				Description: "Time-to-live for FortiGuard Virus Outbreak Prevention cache entries (300 - 86400 sec, default = 300).",
				Computed:    true,
			},
			"outbreak_prevention_expiration": {
				Type:        schema.TypeInt,
				Description: "Expiration date of FortiGuard Virus Outbreak Prevention contract.",
				Computed:    true,
			},
			"outbreak_prevention_force_off": {
				Type:        schema.TypeString,
				Description: "Turn off FortiGuard Virus Outbreak Prevention service.",
				Computed:    true,
			},
			"outbreak_prevention_license": {
				Type:        schema.TypeInt,
				Description: "Interval of time between license checks for FortiGuard Virus Outbreak Prevention contract.",
				Computed:    true,
			},
			"outbreak_prevention_timeout": {
				Type:        schema.TypeInt,
				Description: "FortiGuard Virus Outbreak Prevention time out (1 - 30 sec, default = 7).",
				Computed:    true,
			},
			"persistent_connection": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of persistent connection to receive update notification from FortiGuard.",
				Computed:    true,
			},
			"port": {
				Type:        schema.TypeString,
				Description: "Port used to communicate with the FortiGuard servers.",
				Computed:    true,
			},
			"protocol": {
				Type:        schema.TypeString,
				Description: "Protocol used to communicate with the FortiGuard servers.",
				Computed:    true,
			},
			"proxy_password": {
				Type:        schema.TypeString,
				Description: "Proxy user password.",
				Computed:    true,
				Sensitive:   true,
			},
			"proxy_server_ip": {
				Type:        schema.TypeString,
				Description: "IP address of the proxy server.",
				Computed:    true,
			},
			"proxy_server_port": {
				Type:        schema.TypeInt,
				Description: "Port used to communicate with the proxy server.",
				Computed:    true,
			},
			"proxy_username": {
				Type:        schema.TypeString,
				Description: "Proxy user name.",
				Computed:    true,
			},
			"sandbox_region": {
				Type:        schema.TypeString,
				Description: "Cloud sandbox region.",
				Computed:    true,
			},
			"sdns_options": {
				Type:        schema.TypeString,
				Description: "Customization options for the FortiGuard DNS service.",
				Computed:    true,
			},
			"sdns_server_ip": {
				Type:        schema.TypeString,
				Description: "IP address of the FortiGuard DNS rating server.",
				Computed:    true,
			},
			"sdns_server_port": {
				Type:        schema.TypeInt,
				Description: "Port to connect to on the FortiGuard DNS rating server.",
				Computed:    true,
			},
			"source_ip": {
				Type:        schema.TypeString,
				Description: "Source IPv4 address used to communicate with FortiGuard.",
				Computed:    true,
			},
			"source_ip6": {
				Type:        schema.TypeString,
				Description: "Source IPv6 address used to communicate with FortiGuard.",
				Computed:    true,
			},
			"update_build_proxy": {
				Type:        schema.TypeString,
				Description: "Enable/disable proxy dictionary rebuild.",
				Computed:    true,
			},
			"update_extdb": {
				Type:        schema.TypeString,
				Description: "Enable/disable external resource update.",
				Computed:    true,
			},
			"update_ffdb": {
				Type:        schema.TypeString,
				Description: "Enable/disable Internet Service Database update.",
				Computed:    true,
			},
			"update_server_location": {
				Type:        schema.TypeString,
				Description: "Location from which to receive FortiGuard updates.",
				Computed:    true,
			},
			"update_uwdb": {
				Type:        schema.TypeString,
				Description: "Enable/disable allowlist update.",
				Computed:    true,
			},
			"videofilter_expiration": {
				Type:        schema.TypeInt,
				Description: "Expiration date of the FortiGuard video filter contract.",
				Computed:    true,
			},
			"videofilter_license": {
				Type:        schema.TypeInt,
				Description: "Interval of time between license checks for the FortiGuard video filter contract.",
				Computed:    true,
			},
			"webfilter_cache": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiGuard web filter caching.",
				Computed:    true,
			},
			"webfilter_cache_ttl": {
				Type:        schema.TypeInt,
				Description: "Time-to-live for web filter cache entries in seconds (300 - 86400).",
				Computed:    true,
			},
			"webfilter_expiration": {
				Type:        schema.TypeInt,
				Description: "Expiration date of the FortiGuard web filter contract.",
				Computed:    true,
			},
			"webfilter_force_off": {
				Type:        schema.TypeString,
				Description: "Enable/disable turning off the FortiGuard web filtering service.",
				Computed:    true,
			},
			"webfilter_license": {
				Type:        schema.TypeInt,
				Description: "Interval of time between license checks for the FortiGuard web filter contract.",
				Computed:    true,
			},
			"webfilter_timeout": {
				Type:        schema.TypeInt,
				Description: "Web filter query time out (1 - 30 sec, default = 7).",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemFortiguardRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemFortiguard(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemFortiguard dataSource: %v", err)
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

	diags := refreshObjectSystemFortiguard(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
