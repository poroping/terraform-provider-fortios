// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure RADIUS server entries.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceUserRadius() *schema.Resource {
	return &schema.Resource{
		Description: "Configure RADIUS server entries.",

		ReadContext: dataSourceUserRadiusRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"accounting_server": {
				Type:        schema.TypeList,
				Description: "Additional accounting servers.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "ID (0 - 4294967295).",
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
						"port": {
							Type:        schema.TypeInt,
							Description: "RADIUS accounting port number.",
							Computed:    true,
						},
						"secret": {
							Type:        schema.TypeString,
							Description: "Secret key.",
							Computed:    true,
							Sensitive:   true,
						},
						"server": {
							Type:        schema.TypeString,
							Description: "{<name_str|ip_str>} Server CN domain name or IP.",
							Computed:    true,
						},
						"source_ip": {
							Type:        schema.TypeString,
							Description: "Source IP address for communications to the RADIUS server.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Status.",
							Computed:    true,
						},
					},
				},
			},
			"acct_all_servers": {
				Type:        schema.TypeString,
				Description: "Enable/disable sending of accounting messages to all configured servers (default = disable).",
				Computed:    true,
			},
			"acct_interim_interval": {
				Type:        schema.TypeInt,
				Description: "Time in seconds between each accounting interim update message.",
				Computed:    true,
			},
			"all_usergroup": {
				Type:        schema.TypeString,
				Description: "Enable/disable automatically including this RADIUS server in all user groups.",
				Computed:    true,
			},
			"auth_type": {
				Type:        schema.TypeString,
				Description: "Authentication methods/protocols permitted for this RADIUS server.",
				Computed:    true,
			},
			"class": {
				Type:        schema.TypeList,
				Description: "Class attribute name(s).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Class name.",
							Computed:    true,
						},
					},
				},
			},
			"group_override_attr_type": {
				Type:        schema.TypeString,
				Description: "RADIUS attribute type to override user group information.",
				Computed:    true,
			},
			"h3c_compatibility": {
				Type:        schema.TypeString,
				Description: "Enable/disable compatibility with the H3C, a mechanism that performs security checking for authentication.",
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
			"name": {
				Type:        schema.TypeString,
				Description: "RADIUS server entry name.",
				Required:    true,
			},
			"nas_ip": {
				Type:        schema.TypeString,
				Description: "IP address used to communicate with the RADIUS server and used as NAS-IP-Address and Called-Station-ID attributes.",
				Computed:    true,
			},
			"password_encoding": {
				Type:        schema.TypeString,
				Description: "Password encoding.",
				Computed:    true,
			},
			"password_renewal": {
				Type:        schema.TypeString,
				Description: "Enable/disable password renewal.",
				Computed:    true,
			},
			"radius_coa": {
				Type:        schema.TypeString,
				Description: "Enable to allow a mechanism to change the attributes of an authentication, authorization, and accounting session after it is authenticated.",
				Computed:    true,
			},
			"radius_port": {
				Type:        schema.TypeInt,
				Description: "RADIUS service port number.",
				Computed:    true,
			},
			"rsso": {
				Type:        schema.TypeString,
				Description: "Enable/disable RADIUS based single sign on feature.",
				Computed:    true,
			},
			"rsso_context_timeout": {
				Type:        schema.TypeInt,
				Description: "Time in seconds before the logged out user is removed from the \"user context list\" of logged on users.",
				Computed:    true,
			},
			"rsso_endpoint_attribute": {
				Type:        schema.TypeString,
				Description: "RADIUS attributes used to extract the user end point identifer from the RADIUS Start record.",
				Computed:    true,
			},
			"rsso_endpoint_block_attribute": {
				Type:        schema.TypeString,
				Description: "RADIUS attributes used to block a user.",
				Computed:    true,
			},
			"rsso_ep_one_ip_only": {
				Type:        schema.TypeString,
				Description: "Enable/disable the replacement of old IP addresses with new ones for the same endpoint on RADIUS accounting Start messages.",
				Computed:    true,
			},
			"rsso_flush_ip_session": {
				Type:        schema.TypeString,
				Description: "Enable/disable flushing user IP sessions on RADIUS accounting Stop messages.",
				Computed:    true,
			},
			"rsso_log_flags": {
				Type:        schema.TypeString,
				Description: "Events to log.",
				Computed:    true,
			},
			"rsso_log_period": {
				Type:        schema.TypeInt,
				Description: "Time interval in seconds that group event log messages will be generated for dynamic profile events.",
				Computed:    true,
			},
			"rsso_radius_response": {
				Type:        schema.TypeString,
				Description: "Enable/disable sending RADIUS response packets after receiving Start and Stop records.",
				Computed:    true,
			},
			"rsso_radius_server_port": {
				Type:        schema.TypeInt,
				Description: "UDP port to listen on for RADIUS Start and Stop records.",
				Computed:    true,
			},
			"rsso_secret": {
				Type:        schema.TypeString,
				Description: "RADIUS secret used by the RADIUS accounting server.",
				Computed:    true,
				Sensitive:   true,
			},
			"rsso_validate_request_secret": {
				Type:        schema.TypeString,
				Description: "Enable/disable validating the RADIUS request shared secret in the Start or End record.",
				Computed:    true,
			},
			"secondary_secret": {
				Type:        schema.TypeString,
				Description: "Secret key to access the secondary server.",
				Computed:    true,
				Sensitive:   true,
			},
			"secondary_server": {
				Type:        schema.TypeString,
				Description: "{<name_str|ip_str>} secondary RADIUS CN domain name or IP.",
				Computed:    true,
			},
			"secret": {
				Type:        schema.TypeString,
				Description: "Pre-shared secret key used to access the primary RADIUS server.",
				Computed:    true,
				Sensitive:   true,
			},
			"server": {
				Type:        schema.TypeString,
				Description: "Primary RADIUS server CN domain name or IP address.",
				Computed:    true,
			},
			"source_ip": {
				Type:        schema.TypeString,
				Description: "Source IP address for communications to the RADIUS server.",
				Computed:    true,
			},
			"sso_attribute": {
				Type:        schema.TypeString,
				Description: "RADIUS attribute that contains the profile group name to be extracted from the RADIUS Start record.",
				Computed:    true,
			},
			"sso_attribute_key": {
				Type:        schema.TypeString,
				Description: "Key prefix for SSO group value in the SSO attribute.",
				Computed:    true,
			},
			"sso_attribute_value_override": {
				Type:        schema.TypeString,
				Description: "Enable/disable override old attribute value with new value for the same endpoint.",
				Computed:    true,
			},
			"switch_controller_acct_fast_framedip_detect": {
				Type:        schema.TypeInt,
				Description: "Switch controller accounting message Framed-IP detection from DHCP snooping (seconds, default=2).",
				Computed:    true,
			},
			"switch_controller_service_type": {
				Type:        schema.TypeString,
				Description: "RADIUS service type.",
				Computed:    true,
			},
			"tertiary_secret": {
				Type:        schema.TypeString,
				Description: "Secret key to access the tertiary server.",
				Computed:    true,
				Sensitive:   true,
			},
			"tertiary_server": {
				Type:        schema.TypeString,
				Description: "{<name_str|ip_str>} tertiary RADIUS CN domain name or IP.",
				Computed:    true,
			},
			"timeout": {
				Type:        schema.TypeInt,
				Description: "Time in seconds between re-sending authentication requests.",
				Computed:    true,
			},
			"use_management_vdom": {
				Type:        schema.TypeString,
				Description: "Enable/disable using management VDOM to send requests.",
				Computed:    true,
			},
			"username_case_sensitive": {
				Type:        schema.TypeString,
				Description: "Enable/disable case sensitive user names.",
				Computed:    true,
			},
		},
	}
}

func dataSourceUserRadiusRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserRadius(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserRadius dataSource: %v", err)
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

	diags := refreshObjectUserRadius(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
