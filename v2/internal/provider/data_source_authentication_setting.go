// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure authentication setting.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceAuthenticationSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Configure authentication setting.",

		ReadContext: dataSourceAuthenticationSettingRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"active_auth_scheme": {
				Type:        schema.TypeString,
				Description: "Active authentication method (scheme name).",
				Computed:    true,
			},
			"auth_https": {
				Type:        schema.TypeString,
				Description: "Enable/disable redirecting HTTP user authentication to HTTPS.",
				Computed:    true,
			},
			"captive_portal": {
				Type:        schema.TypeString,
				Description: "Captive portal host name.",
				Computed:    true,
			},
			"captive_portal_ip": {
				Type:        schema.TypeString,
				Description: "Captive portal IP address.",
				Computed:    true,
			},
			"captive_portal_ip6": {
				Type:        schema.TypeString,
				Description: "Captive portal IPv6 address.",
				Computed:    true,
			},
			"captive_portal_port": {
				Type:        schema.TypeInt,
				Description: "Captive portal port number (1 - 65535, default = 7830).",
				Computed:    true,
			},
			"captive_portal_ssl_port": {
				Type:        schema.TypeInt,
				Description: "Captive portal SSL port number (1 - 65535, default = 7831).",
				Computed:    true,
			},
			"captive_portal_type": {
				Type:        schema.TypeString,
				Description: "Captive portal type.",
				Computed:    true,
			},
			"captive_portal6": {
				Type:        schema.TypeString,
				Description: "IPv6 captive portal host name.",
				Computed:    true,
			},
			"cert_auth": {
				Type:        schema.TypeString,
				Description: "Enable/disable redirecting certificate authentication to HTTPS portal.",
				Computed:    true,
			},
			"cert_captive_portal": {
				Type:        schema.TypeString,
				Description: "Certificate captive portal host name.",
				Computed:    true,
			},
			"cert_captive_portal_ip": {
				Type:        schema.TypeString,
				Description: "Certificate captive portal IP address.",
				Computed:    true,
			},
			"cert_captive_portal_port": {
				Type:        schema.TypeInt,
				Description: "Certificate captive portal port number (1 - 65535, default = 7832).",
				Computed:    true,
			},
			"cookie_max_age": {
				Type:        schema.TypeInt,
				Description: "Persistent web portal cookie maximum age in minutes (30 - 10080 (1 week), default = 480 (8 hours)).",
				Computed:    true,
			},
			"cookie_refresh_div": {
				Type:        schema.TypeInt,
				Description: "Refresh rate divider of persistent web portal cookie (default = 2). Refresh value = cookie-max-age/cookie-refresh-div.",
				Computed:    true,
			},
			"dev_range": {
				Type:        schema.TypeList,
				Description: "Address range for the IP based device query.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address name.",
							Computed:    true,
						},
					},
				},
			},
			"ip_auth_cookie": {
				Type:        schema.TypeString,
				Description: "Enable/disable persistent cookie on IP based web portal authentication (default = disable).",
				Computed:    true,
			},
			"persistent_cookie": {
				Type:        schema.TypeString,
				Description: "Enable/disable persistent cookie on web portal authentication (default = enable).",
				Computed:    true,
			},
			"sso_auth_scheme": {
				Type:        schema.TypeString,
				Description: "Single-Sign-On authentication method (scheme name).",
				Computed:    true,
			},
			"update_time": {
				Type:        schema.TypeString,
				Description: "Time of the last update.",
				Computed:    true,
			},
			"user_cert_ca": {
				Type:        schema.TypeList,
				Description: "CA certificate used for client certificate verification.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "CA certificate list.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAuthenticationSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "AuthenticationSetting"

	o, err := c.Cmdb.ReadAuthenticationSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading AuthenticationSetting dataSource: %v", err)
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

	diags := refreshObjectAuthenticationSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
