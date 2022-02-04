// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure WAN optimization profiles.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWanoptProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure WAN optimization profiles.",

		ReadContext: dataSourceWanoptProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"auth_group": {
				Type:        schema.TypeString,
				Description: "Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.",
				Computed:    true,
			},
			"cifs": {
				Type:        schema.TypeList,
				Description: "Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"byte_caching": {
							Type:        schema.TypeString,
							Description: "Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache.",
							Computed:    true,
						},
						"log_traffic": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging.",
							Computed:    true,
						},
						"port": {
							Type:        schema.TypeInt,
							Description: "Single port number or port number range for CIFS. Only packets with a destination port number that matches this port number or range are accepted by this profile.",
							Computed:    true,
						},
						"prefer_chunking": {
							Type:        schema.TypeString,
							Description: "Select dynamic or fixed-size data chunking for WAN Optimization.",
							Computed:    true,
						},
						"protocol_opt": {
							Type:        schema.TypeString,
							Description: "Select protocol specific optimization or generic TCP optimization.",
							Computed:    true,
						},
						"secure_tunnel": {
							Type:        schema.TypeString,
							Description: "Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810).",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable WAN Optimization.",
							Computed:    true,
						},
						"tunnel_sharing": {
							Type:        schema.TypeString,
							Description: "Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols.",
							Computed:    true,
						},
					},
				},
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"ftp": {
				Type:        schema.TypeList,
				Description: "Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"byte_caching": {
							Type:        schema.TypeString,
							Description: "Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache.",
							Computed:    true,
						},
						"log_traffic": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging.",
							Computed:    true,
						},
						"port": {
							Type:        schema.TypeInt,
							Description: "Single port number or port number range for FTP. Only packets with a destination port number that matches this port number or range are accepted by this profile.",
							Computed:    true,
						},
						"prefer_chunking": {
							Type:        schema.TypeString,
							Description: "Select dynamic or fixed-size data chunking for WAN Optimization.",
							Computed:    true,
						},
						"protocol_opt": {
							Type:        schema.TypeString,
							Description: "Select protocol specific optimization or generic TCP optimization.",
							Computed:    true,
						},
						"secure_tunnel": {
							Type:        schema.TypeString,
							Description: "Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810).",
							Computed:    true,
						},
						"ssl": {
							Type:        schema.TypeString,
							Description: "Enable/disable SSL/TLS offloading (hardware acceleration) for traffic in this tunnel.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable WAN Optimization.",
							Computed:    true,
						},
						"tunnel_sharing": {
							Type:        schema.TypeString,
							Description: "Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols.",
							Computed:    true,
						},
					},
				},
			},
			"http": {
				Type:        schema.TypeList,
				Description: "Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"byte_caching": {
							Type:        schema.TypeString,
							Description: "Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache.",
							Computed:    true,
						},
						"log_traffic": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging.",
							Computed:    true,
						},
						"port": {
							Type:        schema.TypeInt,
							Description: "Single port number or port number range for HTTP. Only packets with a destination port number that matches this port number or range are accepted by this profile.",
							Computed:    true,
						},
						"prefer_chunking": {
							Type:        schema.TypeString,
							Description: "Select dynamic or fixed-size data chunking for WAN Optimization.",
							Computed:    true,
						},
						"protocol_opt": {
							Type:        schema.TypeString,
							Description: "Select protocol specific optimization or generic TCP optimization.",
							Computed:    true,
						},
						"secure_tunnel": {
							Type:        schema.TypeString,
							Description: "Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810).",
							Computed:    true,
						},
						"ssl": {
							Type:        schema.TypeString,
							Description: "Enable/disable SSL/TLS offloading (hardware acceleration) for traffic in this tunnel.",
							Computed:    true,
						},
						"ssl_port": {
							Type:        schema.TypeInt,
							Description: "Port on which to expect HTTPS traffic for SSL/TLS offloading.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable WAN Optimization.",
							Computed:    true,
						},
						"tunnel_non_http": {
							Type:        schema.TypeString,
							Description: "Configure how to process non-HTTP traffic when a profile configured for HTTP traffic accepts a non-HTTP session. Can occur if an application sends non-HTTP traffic using an HTTP destination port.",
							Computed:    true,
						},
						"tunnel_sharing": {
							Type:        schema.TypeString,
							Description: "Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols.",
							Computed:    true,
						},
						"unknown_http_version": {
							Type:        schema.TypeString,
							Description: "How to handle HTTP sessions that do not comply with HTTP 0.9, 1.0, or 1.1.",
							Computed:    true,
						},
					},
				},
			},
			"mapi": {
				Type:        schema.TypeList,
				Description: "Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"byte_caching": {
							Type:        schema.TypeString,
							Description: "Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache.",
							Computed:    true,
						},
						"log_traffic": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging.",
							Computed:    true,
						},
						"port": {
							Type:        schema.TypeInt,
							Description: "Single port number or port number range for MAPI. Only packets with a destination port number that matches this port number or range are accepted by this profile.",
							Computed:    true,
						},
						"secure_tunnel": {
							Type:        schema.TypeString,
							Description: "Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810).",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable WAN Optimization.",
							Computed:    true,
						},
						"tunnel_sharing": {
							Type:        schema.TypeString,
							Description: "Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Profile name.",
				Required:    true,
			},
			"tcp": {
				Type:        schema.TypeList,
				Description: "Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"byte_caching": {
							Type:        schema.TypeString,
							Description: "Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache.",
							Computed:    true,
						},
						"byte_caching_opt": {
							Type:        schema.TypeString,
							Description: "Select whether TCP byte-caching uses system memory only or both memory and disk space.",
							Computed:    true,
						},
						"log_traffic": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging.",
							Computed:    true,
						},
						"port": {
							Type:        schema.TypeString,
							Description: "Port numbers or port number ranges for TCP. Only packets with a destination port number that matches this port number or range are accepted by this profile.",
							Computed:    true,
						},
						"secure_tunnel": {
							Type:        schema.TypeString,
							Description: "Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810).",
							Computed:    true,
						},
						"ssl": {
							Type:        schema.TypeString,
							Description: "Enable/disable SSL/TLS offloading (hardware acceleration) for traffic in this tunnel.",
							Computed:    true,
						},
						"ssl_port": {
							Type:        schema.TypeString,
							Description: "Port numbers or port number ranges on which to expect HTTPS traffic for SSL/TLS offloading.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable WAN Optimization.",
							Computed:    true,
						},
						"tunnel_sharing": {
							Type:        schema.TypeString,
							Description: "Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols.",
							Computed:    true,
						},
					},
				},
			},
			"transparent": {
				Type:        schema.TypeString,
				Description: "Enable/disable transparent mode.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWanoptProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWanoptProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WanoptProfile dataSource: %v", err)
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

	diags := refreshObjectWanoptProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
