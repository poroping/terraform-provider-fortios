// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure WAN optimization profiles.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceWanoptProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure WAN optimization profiles.",

		CreateContext: resourceWanoptProfileCreate,
		ReadContext:   resourceWanoptProfileRead,
		UpdateContext: resourceWanoptProfileUpdate,
		DeleteContext: resourceWanoptProfileDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"allow_append": {
				Type:        schema.TypeBool,
				Description: "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"auth_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.",
				Optional:    true,
				Computed:    true,
			},
			"cifs": {
				Type:        schema.TypeList,
				Description: "Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"byte_caching": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache.",
							Optional:    true,
							Computed:    true,
						},
						"log_traffic": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable logging.",
							Optional:    true,
							Computed:    true,
						},
						"port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Single port number or port number range for CIFS. Only packets with a destination port number that matches this port number or range are accepted by this profile.",
							Optional:    true,
							Computed:    true,
						},
						"prefer_chunking": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"dynamic", "fix"}, false),

							Description: "Select dynamic or fixed-size data chunking for WAN Optimization.",
							Optional:    true,
							Computed:    true,
						},
						"protocol_opt": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"protocol", "tcp"}, false),

							Description: "Select protocol specific optimization or generic TCP optimization.",
							Optional:    true,
							Computed:    true,
						},
						"secure_tunnel": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810).",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable WAN Optimization.",
							Optional:    true,
							Computed:    true,
						},
						"tunnel_sharing": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"shared", "express-shared", "private"}, false),

							Description: "Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"ftp": {
				Type:        schema.TypeList,
				Description: "Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"byte_caching": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache.",
							Optional:    true,
							Computed:    true,
						},
						"log_traffic": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable logging.",
							Optional:    true,
							Computed:    true,
						},
						"port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Single port number or port number range for FTP. Only packets with a destination port number that matches this port number or range are accepted by this profile.",
							Optional:    true,
							Computed:    true,
						},
						"prefer_chunking": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"dynamic", "fix"}, false),

							Description: "Select dynamic or fixed-size data chunking for WAN Optimization.",
							Optional:    true,
							Computed:    true,
						},
						"protocol_opt": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"protocol", "tcp"}, false),

							Description: "Select protocol specific optimization or generic TCP optimization.",
							Optional:    true,
							Computed:    true,
						},
						"secure_tunnel": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810).",
							Optional:    true,
							Computed:    true,
						},
						"ssl": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable SSL/TLS offloading (hardware acceleration) for traffic in this tunnel.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable WAN Optimization.",
							Optional:    true,
							Computed:    true,
						},
						"tunnel_sharing": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"shared", "express-shared", "private"}, false),

							Description: "Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"http": {
				Type:        schema.TypeList,
				Description: "Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"byte_caching": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache.",
							Optional:    true,
							Computed:    true,
						},
						"log_traffic": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable logging.",
							Optional:    true,
							Computed:    true,
						},
						"port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Single port number or port number range for HTTP. Only packets with a destination port number that matches this port number or range are accepted by this profile.",
							Optional:    true,
							Computed:    true,
						},
						"prefer_chunking": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"dynamic", "fix"}, false),

							Description: "Select dynamic or fixed-size data chunking for WAN Optimization.",
							Optional:    true,
							Computed:    true,
						},
						"protocol_opt": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"protocol", "tcp"}, false),

							Description: "Select protocol specific optimization or generic TCP optimization.",
							Optional:    true,
							Computed:    true,
						},
						"secure_tunnel": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810).",
							Optional:    true,
							Computed:    true,
						},
						"ssl": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable SSL/TLS offloading (hardware acceleration) for traffic in this tunnel.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Port on which to expect HTTPS traffic for SSL/TLS offloading.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable WAN Optimization.",
							Optional:    true,
							Computed:    true,
						},
						"tunnel_non_http": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Configure how to process non-HTTP traffic when a profile configured for HTTP traffic accepts a non-HTTP session. Can occur if an application sends non-HTTP traffic using an HTTP destination port.",
							Optional:    true,
							Computed:    true,
						},
						"tunnel_sharing": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"shared", "express-shared", "private"}, false),

							Description: "Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols.",
							Optional:    true,
							Computed:    true,
						},
						"unknown_http_version": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"reject", "tunnel", "best-effort"}, false),

							Description: "How to handle HTTP sessions that do not comply with HTTP 0.9, 1.0, or 1.1.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"mapi": {
				Type:        schema.TypeList,
				Description: "Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"byte_caching": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache.",
							Optional:    true,
							Computed:    true,
						},
						"log_traffic": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable logging.",
							Optional:    true,
							Computed:    true,
						},
						"port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Single port number or port number range for MAPI. Only packets with a destination port number that matches this port number or range are accepted by this profile.",
							Optional:    true,
							Computed:    true,
						},
						"secure_tunnel": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810).",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable WAN Optimization.",
							Optional:    true,
							Computed:    true,
						},
						"tunnel_sharing": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"shared", "express-shared", "private"}, false),

							Description: "Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Profile name.",
				ForceNew:    true,
				Required:    true,
			},
			"tcp": {
				Type:        schema.TypeList,
				Description: "Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"byte_caching": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache.",
							Optional:    true,
							Computed:    true,
						},
						"byte_caching_opt": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"mem-only", "mem-disk"}, false),

							Description: "Select whether TCP byte-caching uses system memory only or both memory and disk space.",
							Optional:    true,
							Computed:    true,
						},
						"log_traffic": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable logging.",
							Optional:    true,
							Computed:    true,
						},
						"port": {
							Type: schema.TypeString,

							Description: "Port numbers or port number ranges for TCP. Only packets with a destination port number that matches this port number or range are accepted by this profile.",
							Optional:    true,
							Computed:    true,
						},
						"secure_tunnel": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810).",
							Optional:    true,
							Computed:    true,
						},
						"ssl": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable SSL/TLS offloading (hardware acceleration) for traffic in this tunnel.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_port": {
							Type: schema.TypeString,

							Description: "Port numbers or port number ranges on which to expect HTTPS traffic for SSL/TLS offloading.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable WAN Optimization.",
							Optional:    true,
							Computed:    true,
						},
						"tunnel_sharing": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"shared", "express-shared", "private"}, false),

							Description: "Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"transparent": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable transparent mode.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWanoptProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	var diags diag.Diagnostics
	var err error
	// c.Retries = 1
	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	allow_append := false
	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}
	urlparams.AllowAppend = &allow_append

	mkey := ""
	key := "name"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating WanoptProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWanoptProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWanoptProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WanoptProfile")
	}

	return resourceWanoptProfileRead(ctx, d, meta)
}

func resourceWanoptProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWanoptProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWanoptProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WanoptProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WanoptProfile")
	}

	return resourceWanoptProfileRead(ctx, d, meta)
}

func resourceWanoptProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWanoptProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WanoptProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	ptp := true
	urlparams.PlainTextPassword = &ptp

	o, err := c.Cmdb.ReadWanoptProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WanoptProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
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
	return nil
}

func flattenWanoptProfileCifs(d *schema.ResourceData, v *models.WanoptProfileCifs, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.WanoptProfileCifs{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ByteCaching; tmp != nil {
				v["byte_caching"] = *tmp
			}

			if tmp := cfg.LogTraffic; tmp != nil {
				v["log_traffic"] = *tmp
			}

			if tmp := cfg.Port; tmp != nil {
				v["port"] = *tmp
			}

			if tmp := cfg.PreferChunking; tmp != nil {
				v["prefer_chunking"] = *tmp
			}

			if tmp := cfg.ProtocolOpt; tmp != nil {
				v["protocol_opt"] = *tmp
			}

			if tmp := cfg.SecureTunnel; tmp != nil {
				v["secure_tunnel"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.TunnelSharing; tmp != nil {
				v["tunnel_sharing"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWanoptProfileFtp(d *schema.ResourceData, v *models.WanoptProfileFtp, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.WanoptProfileFtp{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ByteCaching; tmp != nil {
				v["byte_caching"] = *tmp
			}

			if tmp := cfg.LogTraffic; tmp != nil {
				v["log_traffic"] = *tmp
			}

			if tmp := cfg.Port; tmp != nil {
				v["port"] = *tmp
			}

			if tmp := cfg.PreferChunking; tmp != nil {
				v["prefer_chunking"] = *tmp
			}

			if tmp := cfg.ProtocolOpt; tmp != nil {
				v["protocol_opt"] = *tmp
			}

			if tmp := cfg.SecureTunnel; tmp != nil {
				v["secure_tunnel"] = *tmp
			}

			if tmp := cfg.Ssl; tmp != nil {
				v["ssl"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.TunnelSharing; tmp != nil {
				v["tunnel_sharing"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWanoptProfileHttp(d *schema.ResourceData, v *models.WanoptProfileHttp, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.WanoptProfileHttp{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ByteCaching; tmp != nil {
				v["byte_caching"] = *tmp
			}

			if tmp := cfg.LogTraffic; tmp != nil {
				v["log_traffic"] = *tmp
			}

			if tmp := cfg.Port; tmp != nil {
				v["port"] = *tmp
			}

			if tmp := cfg.PreferChunking; tmp != nil {
				v["prefer_chunking"] = *tmp
			}

			if tmp := cfg.ProtocolOpt; tmp != nil {
				v["protocol_opt"] = *tmp
			}

			if tmp := cfg.SecureTunnel; tmp != nil {
				v["secure_tunnel"] = *tmp
			}

			if tmp := cfg.Ssl; tmp != nil {
				v["ssl"] = *tmp
			}

			if tmp := cfg.SslPort; tmp != nil {
				v["ssl_port"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.TunnelNonHttp; tmp != nil {
				v["tunnel_non_http"] = *tmp
			}

			if tmp := cfg.TunnelSharing; tmp != nil {
				v["tunnel_sharing"] = *tmp
			}

			if tmp := cfg.UnknownHttpVersion; tmp != nil {
				v["unknown_http_version"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWanoptProfileMapi(d *schema.ResourceData, v *models.WanoptProfileMapi, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.WanoptProfileMapi{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ByteCaching; tmp != nil {
				v["byte_caching"] = *tmp
			}

			if tmp := cfg.LogTraffic; tmp != nil {
				v["log_traffic"] = *tmp
			}

			if tmp := cfg.Port; tmp != nil {
				v["port"] = *tmp
			}

			if tmp := cfg.SecureTunnel; tmp != nil {
				v["secure_tunnel"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.TunnelSharing; tmp != nil {
				v["tunnel_sharing"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWanoptProfileTcp(d *schema.ResourceData, v *models.WanoptProfileTcp, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.WanoptProfileTcp{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ByteCaching; tmp != nil {
				v["byte_caching"] = *tmp
			}

			if tmp := cfg.ByteCachingOpt; tmp != nil {
				v["byte_caching_opt"] = *tmp
			}

			if tmp := cfg.LogTraffic; tmp != nil {
				v["log_traffic"] = *tmp
			}

			if tmp := cfg.Port; tmp != nil {
				v["port"] = *tmp
			}

			if tmp := cfg.SecureTunnel; tmp != nil {
				v["secure_tunnel"] = *tmp
			}

			if tmp := cfg.Ssl; tmp != nil {
				v["ssl"] = *tmp
			}

			if tmp := cfg.SslPort; tmp != nil {
				v["ssl_port"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.TunnelSharing; tmp != nil {
				v["tunnel_sharing"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func refreshObjectWanoptProfile(d *schema.ResourceData, o *models.WanoptProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AuthGroup != nil {
		v := *o.AuthGroup

		if err = d.Set("auth_group", v); err != nil {
			return diag.Errorf("error reading auth_group: %v", err)
		}
	}

	if _, ok := d.GetOk("cifs"); ok {
		if o.Cifs != nil {
			if err = d.Set("cifs", flattenWanoptProfileCifs(d, o.Cifs, "cifs", sort)); err != nil {
				return diag.Errorf("error reading cifs: %v", err)
			}
		}
	}

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if _, ok := d.GetOk("ftp"); ok {
		if o.Ftp != nil {
			if err = d.Set("ftp", flattenWanoptProfileFtp(d, o.Ftp, "ftp", sort)); err != nil {
				return diag.Errorf("error reading ftp: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("http"); ok {
		if o.Http != nil {
			if err = d.Set("http", flattenWanoptProfileHttp(d, o.Http, "http", sort)); err != nil {
				return diag.Errorf("error reading http: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("mapi"); ok {
		if o.Mapi != nil {
			if err = d.Set("mapi", flattenWanoptProfileMapi(d, o.Mapi, "mapi", sort)); err != nil {
				return diag.Errorf("error reading mapi: %v", err)
			}
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if _, ok := d.GetOk("tcp"); ok {
		if o.Tcp != nil {
			if err = d.Set("tcp", flattenWanoptProfileTcp(d, o.Tcp, "tcp", sort)); err != nil {
				return diag.Errorf("error reading tcp: %v", err)
			}
		}
	}

	if o.Transparent != nil {
		v := *o.Transparent

		if err = d.Set("transparent", v); err != nil {
			return diag.Errorf("error reading transparent: %v", err)
		}
	}

	return nil
}

func expandWanoptProfileCifs(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.WanoptProfileCifs, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WanoptProfileCifs

	for i := range l {
		tmp := models.WanoptProfileCifs{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.byte_caching", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ByteCaching = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log_traffic", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogTraffic = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Port = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefer_chunking", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PreferChunking = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.protocol_opt", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ProtocolOpt = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.secure_tunnel", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SecureTunnel = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tunnel_sharing", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TunnelSharing = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandWanoptProfileFtp(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.WanoptProfileFtp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WanoptProfileFtp

	for i := range l {
		tmp := models.WanoptProfileFtp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.byte_caching", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ByteCaching = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log_traffic", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogTraffic = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Port = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefer_chunking", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PreferChunking = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.protocol_opt", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ProtocolOpt = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.secure_tunnel", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SecureTunnel = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ssl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tunnel_sharing", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TunnelSharing = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandWanoptProfileHttp(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.WanoptProfileHttp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WanoptProfileHttp

	for i := range l {
		tmp := models.WanoptProfileHttp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.byte_caching", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ByteCaching = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log_traffic", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogTraffic = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Port = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefer_chunking", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PreferChunking = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.protocol_opt", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ProtocolOpt = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.secure_tunnel", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SecureTunnel = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ssl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SslPort = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tunnel_non_http", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TunnelNonHttp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tunnel_sharing", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TunnelSharing = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unknown_http_version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnknownHttpVersion = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandWanoptProfileMapi(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.WanoptProfileMapi, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WanoptProfileMapi

	for i := range l {
		tmp := models.WanoptProfileMapi{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.byte_caching", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ByteCaching = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log_traffic", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogTraffic = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Port = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.secure_tunnel", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SecureTunnel = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tunnel_sharing", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TunnelSharing = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandWanoptProfileTcp(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.WanoptProfileTcp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WanoptProfileTcp

	for i := range l {
		tmp := models.WanoptProfileTcp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.byte_caching", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ByteCaching = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.byte_caching_opt", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ByteCachingOpt = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log_traffic", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogTraffic = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Port = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.secure_tunnel", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SecureTunnel = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ssl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslPort = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tunnel_sharing", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TunnelSharing = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func getObjectWanoptProfile(d *schema.ResourceData, sv string) (*models.WanoptProfile, diag.Diagnostics) {
	obj := models.WanoptProfile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("auth_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_group", sv)
				diags = append(diags, e)
			}
			obj.AuthGroup = &v2
		}
	}
	if v, ok := d.GetOk("cifs"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("cifs", sv)
			diags = append(diags, e)
		}
		t, err := expandWanoptProfileCifs(d, v, "cifs", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Cifs = t
		}
	} else if d.HasChange("cifs") {
		old, new := d.GetChange("cifs")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Cifs = &models.WanoptProfileCifs{}
		}
	}
	if v1, ok := d.GetOk("comments"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comments", sv)
				diags = append(diags, e)
			}
			obj.Comments = &v2
		}
	}
	if v, ok := d.GetOk("ftp"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ftp", sv)
			diags = append(diags, e)
		}
		t, err := expandWanoptProfileFtp(d, v, "ftp", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Ftp = t
		}
	} else if d.HasChange("ftp") {
		old, new := d.GetChange("ftp")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Ftp = &models.WanoptProfileFtp{}
		}
	}
	if v, ok := d.GetOk("http"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("http", sv)
			diags = append(diags, e)
		}
		t, err := expandWanoptProfileHttp(d, v, "http", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Http = t
		}
	} else if d.HasChange("http") {
		old, new := d.GetChange("http")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Http = &models.WanoptProfileHttp{}
		}
	}
	if v, ok := d.GetOk("mapi"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("mapi", sv)
			diags = append(diags, e)
		}
		t, err := expandWanoptProfileMapi(d, v, "mapi", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Mapi = t
		}
	} else if d.HasChange("mapi") {
		old, new := d.GetChange("mapi")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Mapi = &models.WanoptProfileMapi{}
		}
	}
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v, ok := d.GetOk("tcp"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("tcp", sv)
			diags = append(diags, e)
		}
		t, err := expandWanoptProfileTcp(d, v, "tcp", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Tcp = t
		}
	} else if d.HasChange("tcp") {
		old, new := d.GetChange("tcp")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Tcp = &models.WanoptProfileTcp{}
		}
	}
	if v1, ok := d.GetOk("transparent"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("transparent", sv)
				diags = append(diags, e)
			}
			obj.Transparent = &v2
		}
	}
	return &obj, diags
}
