// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure user authentication setting.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceUserSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Configure user authentication setting.",

		ReadContext: dataSourceUserSettingRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"auth_blackout_time": {
				Type:        schema.TypeInt,
				Description: "Time in seconds an IP address is denied access after failing to authenticate five times within one minute.",
				Computed:    true,
			},
			"auth_ca_cert": {
				Type:        schema.TypeString,
				Description: "HTTPS CA certificate for policy authentication.",
				Computed:    true,
			},
			"auth_cert": {
				Type:        schema.TypeString,
				Description: "HTTPS server certificate for policy authentication.",
				Computed:    true,
			},
			"auth_http_basic": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of HTTP basic authentication for identity-based firewall policies.",
				Computed:    true,
			},
			"auth_invalid_max": {
				Type:        schema.TypeInt,
				Description: "Maximum number of failed authentication attempts before the user is blocked.",
				Computed:    true,
			},
			"auth_lockout_duration": {
				Type:        schema.TypeInt,
				Description: "Lockout period in seconds after too many login failures.",
				Computed:    true,
			},
			"auth_lockout_threshold": {
				Type:        schema.TypeInt,
				Description: "Maximum number of failed login attempts before login lockout is triggered.",
				Computed:    true,
			},
			"auth_on_demand": {
				Type:        schema.TypeString,
				Description: "Always/implicitly trigger firewall authentication on demand.",
				Computed:    true,
			},
			"auth_portal_timeout": {
				Type:        schema.TypeInt,
				Description: "Time in minutes before captive portal user have to re-authenticate (1 - 30 min, default 3 min).",
				Computed:    true,
			},
			"auth_ports": {
				Type:        schema.TypeList,
				Description: "Set up non-standard ports for authentication with HTTP, HTTPS, FTP, and TELNET.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"port": {
							Type:        schema.TypeInt,
							Description: "Non-standard port for firewall user authentication.",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "Service type.",
							Computed:    true,
						},
					},
				},
			},
			"auth_secure_http": {
				Type:        schema.TypeString,
				Description: "Enable/disable redirecting HTTP user authentication to more secure HTTPS.",
				Computed:    true,
			},
			"auth_src_mac": {
				Type:        schema.TypeString,
				Description: "Enable/disable source MAC for user identity.",
				Computed:    true,
			},
			"auth_ssl_allow_renegotiation": {
				Type:        schema.TypeString,
				Description: "Allow/forbid SSL re-negotiation for HTTPS authentication.",
				Computed:    true,
			},
			"auth_ssl_max_proto_version": {
				Type:        schema.TypeString,
				Description: "Maximum supported protocol version for SSL/TLS connections (default is no limit).",
				Computed:    true,
			},
			"auth_ssl_min_proto_version": {
				Type:        schema.TypeString,
				Description: "Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).",
				Computed:    true,
			},
			"auth_ssl_sigalgs": {
				Type:        schema.TypeString,
				Description: "Set signature algorithms related to HTTPS authentication (affects TLS version <= 1.2 only, default is to enable all).",
				Computed:    true,
			},
			"auth_timeout": {
				Type:        schema.TypeInt,
				Description: "Time in minutes before the firewall user authentication timeout requires the user to re-authenticate.",
				Computed:    true,
			},
			"auth_timeout_type": {
				Type:        schema.TypeString,
				Description: "Control if authenticated users have to login again after a hard timeout, after an idle timeout, or after a session timeout.",
				Computed:    true,
			},
			"auth_type": {
				Type:        schema.TypeString,
				Description: "Supported firewall policy authentication protocols/methods.",
				Computed:    true,
			},
			"per_policy_disclaimer": {
				Type:        schema.TypeString,
				Description: "Enable/disable per policy disclaimer.",
				Computed:    true,
			},
			"radius_ses_timeout_act": {
				Type:        schema.TypeString,
				Description: "Set the RADIUS session timeout to a hard timeout or to ignore RADIUS server session timeouts.",
				Computed:    true,
			},
		},
	}
}

func dataSourceUserSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "UserSetting"

	o, err := c.Cmdb.ReadUserSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserSetting dataSource: %v", err)
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

	diags := refreshObjectUserSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
