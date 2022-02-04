// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure local users.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceUserLocal() *schema.Resource {
	return &schema.Resource{
		Description: "Configure local users.",

		ReadContext: dataSourceUserLocalRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"auth_concurrent_override": {
				Type:        schema.TypeString,
				Description: "Enable/disable overriding the policy-auth-concurrent under config system global.",
				Computed:    true,
			},
			"auth_concurrent_value": {
				Type:        schema.TypeInt,
				Description: "Maximum number of concurrent logins permitted from the same user.",
				Computed:    true,
			},
			"authtimeout": {
				Type:        schema.TypeInt,
				Description: "Time in minutes before the authentication timeout for a user is reached.",
				Computed:    true,
			},
			"email_to": {
				Type:        schema.TypeString,
				Description: "Two-factor recipient's email address.",
				Computed:    true,
			},
			"fortitoken": {
				Type:        schema.TypeString,
				Description: "Two-factor recipient's FortiToken serial number.",
				Computed:    true,
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "User ID.",
				Computed:    true,
			},
			"ldap_server": {
				Type:        schema.TypeString,
				Description: "Name of LDAP server with which the user must authenticate.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "User name.",
				Required:    true,
			},
			"passwd": {
				Type:        schema.TypeString,
				Description: "User's password.",
				Computed:    true,
				Sensitive:   true,
			},
			"passwd_policy": {
				Type:        schema.TypeString,
				Description: "Password policy to apply to this user, as defined in config user password-policy.",
				Computed:    true,
			},
			"passwd_time": {
				Type:        schema.TypeString,
				Description: "Time of the last password update.",
				Computed:    true,
			},
			"ppk_identity": {
				Type:        schema.TypeString,
				Description: "IKEv2 Postquantum Preshared Key Identity.",
				Computed:    true,
			},
			"ppk_secret": {
				Type:        schema.TypeString,
				Description: "IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).",
				Computed:    true,
				Sensitive:   true,
			},
			"radius_server": {
				Type:        schema.TypeString,
				Description: "Name of RADIUS server with which the user must authenticate.",
				Computed:    true,
			},
			"sms_custom_server": {
				Type:        schema.TypeString,
				Description: "Two-factor recipient's SMS server.",
				Computed:    true,
			},
			"sms_phone": {
				Type:        schema.TypeString,
				Description: "Two-factor recipient's mobile phone number.",
				Computed:    true,
			},
			"sms_server": {
				Type:        schema.TypeString,
				Description: "Send SMS through FortiGuard or other external server.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable allowing the local user to authenticate with the FortiGate unit.",
				Computed:    true,
			},
			"tacacs_server": {
				Type:        schema.TypeString,
				Description: "Name of TACACS+ server with which the user must authenticate.",
				Computed:    true,
			},
			"two_factor": {
				Type:        schema.TypeString,
				Description: "Enable/disable two-factor authentication.",
				Computed:    true,
			},
			"two_factor_authentication": {
				Type:        schema.TypeString,
				Description: "Authentication method by FortiToken Cloud.",
				Computed:    true,
			},
			"two_factor_notification": {
				Type:        schema.TypeString,
				Description: "Notification method for user activation by FortiToken Cloud.",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Authentication method.",
				Computed:    true,
			},
			"username_case_insensitivity": {
				Type:        schema.TypeString,
				Description: "Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent).",
				Computed:    true,
			},
			"username_case_sensitivity": {
				Type:        schema.TypeString,
				Description: "Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent).",
				Computed:    true,
			},
			"username_sensitivity": {
				Type:        schema.TypeString,
				Description: "Enable/disable case and accent sensitivity when performing username matching (accents are stripped and case is ignored when disabled).",
				Computed:    true,
			},
			"workstation": {
				Type:        schema.TypeString,
				Description: "Name of the remote user workstation, if you want to limit the user to authenticate only from a particular workstation.",
				Computed:    true,
			},
		},
	}
}

func dataSourceUserLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserLocal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserLocal dataSource: %v", err)
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

	diags := refreshObjectUserLocal(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
