// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Authentication Schemes.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceAuthenticationScheme() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Authentication Schemes.",

		ReadContext: dataSourceAuthenticationSchemeRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"domain_controller": {
				Type:        schema.TypeString,
				Description: "Domain controller setting.",
				Computed:    true,
			},
			"ems_device_owner": {
				Type:        schema.TypeString,
				Description: "Enable/disable SSH public-key authentication with device owner (default = disable).",
				Computed:    true,
			},
			"fsso_agent_for_ntlm": {
				Type:        schema.TypeString,
				Description: "FSSO agent to use for NTLM authentication.",
				Computed:    true,
			},
			"fsso_guest": {
				Type:        schema.TypeString,
				Description: "Enable/disable user fsso-guest authentication (default = disable).",
				Computed:    true,
			},
			"kerberos_keytab": {
				Type:        schema.TypeString,
				Description: "Kerberos keytab setting.",
				Computed:    true,
			},
			"method": {
				Type:        schema.TypeString,
				Description: "Authentication methods (default = basic).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Authentication scheme name.",
				Required:    true,
			},
			"negotiate_ntlm": {
				Type:        schema.TypeString,
				Description: "Enable/disable negotiate authentication for NTLM (default = disable).",
				Computed:    true,
			},
			"require_tfa": {
				Type:        schema.TypeString,
				Description: "Enable/disable two-factor authentication (default = disable).",
				Computed:    true,
			},
			"saml_server": {
				Type:        schema.TypeString,
				Description: "SAML configuration.",
				Computed:    true,
			},
			"saml_timeout": {
				Type:        schema.TypeInt,
				Description: "SAML authentication timeout in seconds.",
				Computed:    true,
			},
			"ssh_ca": {
				Type:        schema.TypeString,
				Description: "SSH CA name.",
				Computed:    true,
			},
			"user_cert": {
				Type:        schema.TypeString,
				Description: "Enable/disable authentication with user certificate (default = disable).",
				Computed:    true,
			},
			"user_database": {
				Type:        schema.TypeList,
				Description: "Authentication server to contain user information; \"local\" (default) or \"123\" (for LDAP).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Authentication server name.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAuthenticationSchemeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadAuthenticationScheme(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading AuthenticationScheme dataSource: %v", err)
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

	diags := refreshObjectAuthenticationScheme(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
