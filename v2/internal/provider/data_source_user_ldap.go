// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure LDAP server entries.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceUserLdap() *schema.Resource {
	return &schema.Resource{
		Description: "Configure LDAP server entries.",

		ReadContext: dataSourceUserLdapRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"account_key_filter": {
				Type:        schema.TypeString,
				Description: "Account key filter, using the UPN as the search filter.",
				Computed:    true,
			},
			"account_key_processing": {
				Type:        schema.TypeString,
				Description: "Account key processing operation, either keep or strip domain string of UPN in the token.",
				Computed:    true,
			},
			"account_key_upn_san": {
				Type:        schema.TypeString,
				Description: "Define SAN in certificate for user principle name matching.",
				Computed:    true,
			},
			"antiphish": {
				Type:        schema.TypeString,
				Description: "Enable/disable AntiPhishing credential backend.",
				Computed:    true,
			},
			"ca_cert": {
				Type:        schema.TypeString,
				Description: "CA certificate name.",
				Computed:    true,
			},
			"client_cert": {
				Type:        schema.TypeString,
				Description: "Client certificate name.",
				Computed:    true,
			},
			"client_cert_auth": {
				Type:        schema.TypeString,
				Description: "Enable/disable using client certificate for TLS authentication.",
				Computed:    true,
			},
			"cnid": {
				Type:        schema.TypeString,
				Description: "Common name identifier for the LDAP server. The common name identifier for most LDAP servers is \"cn\".",
				Computed:    true,
			},
			"dn": {
				Type:        schema.TypeString,
				Description: "Distinguished name used to look up entries on the LDAP server.",
				Computed:    true,
			},
			"group_filter": {
				Type:        schema.TypeString,
				Description: "Filter used for group matching.",
				Computed:    true,
			},
			"group_member_check": {
				Type:        schema.TypeString,
				Description: "Group member checking methods.",
				Computed:    true,
			},
			"group_object_filter": {
				Type:        schema.TypeString,
				Description: "Filter used for group searching.",
				Computed:    true,
			},
			"group_search_base": {
				Type:        schema.TypeString,
				Description: "Search base used for group searching.",
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
			"member_attr": {
				Type:        schema.TypeString,
				Description: "Name of attribute from which to get group membership.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "LDAP server entry name.",
				Required:    true,
			},
			"obtain_user_info": {
				Type:        schema.TypeString,
				Description: "Enable/disable obtaining of user information.",
				Computed:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "Password for initial binding.",
				Computed:    true,
				Sensitive:   true,
			},
			"password_attr": {
				Type:        schema.TypeString,
				Description: "Name of attribute to get password hash.",
				Computed:    true,
			},
			"password_expiry_warning": {
				Type:        schema.TypeString,
				Description: "Enable/disable password expiry warnings.",
				Computed:    true,
			},
			"password_renewal": {
				Type:        schema.TypeString,
				Description: "Enable/disable online password renewal.",
				Computed:    true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "Port to be used for communication with the LDAP server (default = 389).",
				Computed:    true,
			},
			"search_type": {
				Type:        schema.TypeString,
				Description: "Search type.",
				Computed:    true,
			},
			"secondary_server": {
				Type:        schema.TypeString,
				Description: "Secondary LDAP server CN domain name or IP.",
				Computed:    true,
			},
			"secure": {
				Type:        schema.TypeString,
				Description: "Port to be used for authentication.",
				Computed:    true,
			},
			"server": {
				Type:        schema.TypeString,
				Description: "LDAP server CN domain name or IP.",
				Computed:    true,
			},
			"server_identity_check": {
				Type:        schema.TypeString,
				Description: "Enable/disable LDAP server identity check (verify server domain name/IP address against the server certificate).",
				Computed:    true,
			},
			"source_ip": {
				Type:        schema.TypeString,
				Description: "FortiGate IP address to be used for communication with the LDAP server.",
				Computed:    true,
			},
			"source_port": {
				Type:        schema.TypeInt,
				Description: "Source port to be used for communication with the LDAP server.",
				Computed:    true,
			},
			"ssl_min_proto_version": {
				Type:        schema.TypeString,
				Description: "Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).",
				Computed:    true,
			},
			"tertiary_server": {
				Type:        schema.TypeString,
				Description: "Tertiary LDAP server CN domain name or IP.",
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
			"two_factor_filter": {
				Type:        schema.TypeString,
				Description: "Filter used to synchronize users to FortiToken Cloud.",
				Computed:    true,
			},
			"two_factor_notification": {
				Type:        schema.TypeString,
				Description: "Notification method for user activation by FortiToken Cloud.",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Authentication type for LDAP searches.",
				Computed:    true,
			},
			"user_info_exchange_server": {
				Type:        schema.TypeString,
				Description: "MS Exchange server from which to fetch user information.",
				Computed:    true,
			},
			"username": {
				Type:        schema.TypeString,
				Description: "Username (full DN) for initial binding.",
				Computed:    true,
			},
		},
	}
}

func dataSourceUserLdapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserLdap(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserLdap dataSource: %v", err)
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

	diags := refreshObjectUserLdap(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
