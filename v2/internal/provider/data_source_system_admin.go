// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure admin users.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemAdmin() *schema.Resource {
	return &schema.Resource{
		Description: "Configure admin users.",

		ReadContext: dataSourceSystemAdminRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"accprofile": {
				Type:        schema.TypeString,
				Description: "Access profile for this administrator. Access profiles control administrator access to FortiGate features.",
				Computed:    true,
			},
			"accprofile_override": {
				Type:        schema.TypeString,
				Description: "Enable to use the name of an access profile provided by the remote authentication server to control the FortiGate features that this administrator can access.",
				Computed:    true,
			},
			"allow_remove_admin_session": {
				Type:        schema.TypeString,
				Description: "Enable/disable allow admin session to be removed by privileged admin users.",
				Computed:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"email_to": {
				Type:        schema.TypeString,
				Description: "This administrator's email address.",
				Computed:    true,
			},
			"force_password_change": {
				Type:        schema.TypeString,
				Description: "Enable/disable force password change on next login.",
				Computed:    true,
			},
			"fortitoken": {
				Type:        schema.TypeString,
				Description: "This administrator's FortiToken serial number.",
				Computed:    true,
			},
			"guest_auth": {
				Type:        schema.TypeString,
				Description: "Enable/disable guest authentication.",
				Computed:    true,
			},
			"guest_lang": {
				Type:        schema.TypeString,
				Description: "Guest management portal language.",
				Computed:    true,
			},
			"guest_usergroups": {
				Type:        schema.TypeList,
				Description: "Select guest user groups.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Select guest user groups.",
							Computed:    true,
						},
					},
				},
			},
			"ip6_trusthost1": {
				Type:        schema.TypeString,
				Description: "Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.",
				Computed:    true,
			},
			"ip6_trusthost10": {
				Type:        schema.TypeString,
				Description: "Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.",
				Computed:    true,
			},
			"ip6_trusthost2": {
				Type:        schema.TypeString,
				Description: "Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.",
				Computed:    true,
			},
			"ip6_trusthost3": {
				Type:        schema.TypeString,
				Description: "Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.",
				Computed:    true,
			},
			"ip6_trusthost4": {
				Type:        schema.TypeString,
				Description: "Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.",
				Computed:    true,
			},
			"ip6_trusthost5": {
				Type:        schema.TypeString,
				Description: "Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.",
				Computed:    true,
			},
			"ip6_trusthost6": {
				Type:        schema.TypeString,
				Description: "Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.",
				Computed:    true,
			},
			"ip6_trusthost7": {
				Type:        schema.TypeString,
				Description: "Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.",
				Computed:    true,
			},
			"ip6_trusthost8": {
				Type:        schema.TypeString,
				Description: "Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.",
				Computed:    true,
			},
			"ip6_trusthost9": {
				Type:        schema.TypeString,
				Description: "Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "User name.",
				Required:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "Admin user password.",
				Computed:    true,
				Sensitive:   true,
			},
			"password_expire": {
				Type:        schema.TypeString,
				Description: "Password expire time.",
				Computed:    true,
			},
			"peer_auth": {
				Type:        schema.TypeString,
				Description: "Set to enable peer certificate authentication (for HTTPS admin access).",
				Computed:    true,
			},
			"peer_group": {
				Type:        schema.TypeString,
				Description: "Name of peer group defined under config user group which has PKI members. Used for peer certificate authentication (for HTTPS admin access).",
				Computed:    true,
			},
			"radius_vdom_override": {
				Type:        schema.TypeString,
				Description: "Enable to use the names of VDOMs provided by the remote authentication server to control the VDOMs that this administrator can access.",
				Computed:    true,
			},
			"remote_auth": {
				Type:        schema.TypeString,
				Description: "Enable/disable authentication using a remote RADIUS, LDAP, or TACACS+ server.",
				Computed:    true,
			},
			"remote_group": {
				Type:        schema.TypeString,
				Description: "User group name used for remote auth.",
				Computed:    true,
			},
			"schedule": {
				Type:        schema.TypeString,
				Description: "Firewall schedule used to restrict when the administrator can log in. No schedule means no restrictions.",
				Computed:    true,
			},
			"sms_custom_server": {
				Type:        schema.TypeString,
				Description: "Custom SMS server to send SMS messages to.",
				Computed:    true,
			},
			"sms_phone": {
				Type:        schema.TypeString,
				Description: "Phone number on which the administrator receives SMS messages.",
				Computed:    true,
			},
			"sms_server": {
				Type:        schema.TypeString,
				Description: "Send SMS messages using the FortiGuard SMS server or a custom server.",
				Computed:    true,
			},
			"ssh_certificate": {
				Type:        schema.TypeString,
				Description: "Select the certificate to be used by the FortiGate for authentication with an SSH client.",
				Computed:    true,
			},
			"ssh_public_key1": {
				Type:        schema.TypeString,
				Description: "Public key of an SSH client. The client is authenticated without being asked for credentials. Create the public-private key pair in the SSH client application.",
				Computed:    true,
			},
			"ssh_public_key2": {
				Type:        schema.TypeString,
				Description: "Public key of an SSH client. The client is authenticated without being asked for credentials. Create the public-private key pair in the SSH client application.",
				Computed:    true,
			},
			"ssh_public_key3": {
				Type:        schema.TypeString,
				Description: "Public key of an SSH client. The client is authenticated without being asked for credentials. Create the public-private key pair in the SSH client application.",
				Computed:    true,
			},
			"trusthost1": {
				Type:        schema.TypeString,
				Description: "Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.",
				Computed:    true,
			},
			"trusthost10": {
				Type:        schema.TypeString,
				Description: "Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.",
				Computed:    true,
			},
			"trusthost2": {
				Type:        schema.TypeString,
				Description: "Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.",
				Computed:    true,
			},
			"trusthost3": {
				Type:        schema.TypeString,
				Description: "Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.",
				Computed:    true,
			},
			"trusthost4": {
				Type:        schema.TypeString,
				Description: "Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.",
				Computed:    true,
			},
			"trusthost5": {
				Type:        schema.TypeString,
				Description: "Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.",
				Computed:    true,
			},
			"trusthost6": {
				Type:        schema.TypeString,
				Description: "Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.",
				Computed:    true,
			},
			"trusthost7": {
				Type:        schema.TypeString,
				Description: "Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.",
				Computed:    true,
			},
			"trusthost8": {
				Type:        schema.TypeString,
				Description: "Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.",
				Computed:    true,
			},
			"trusthost9": {
				Type:        schema.TypeString,
				Description: "Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.",
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
			"vdom": {
				Type:        schema.TypeList,
				Description: "Virtual domain(s) that the administrator can access.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Virtual domain name.",
							Computed:    true,
						},
					},
				},
			},
			"wildcard": {
				Type:        schema.TypeString,
				Description: "Enable/disable wildcard RADIUS authentication.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemAdminRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemAdmin(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemAdmin dataSource: %v", err)
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

	diags := refreshObjectSystemAdmin(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
