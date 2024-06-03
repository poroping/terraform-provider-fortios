// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure peer users.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceUserPeer() *schema.Resource {
	return &schema.Resource{
		Description: "Configure peer users.",

		ReadContext: dataSourceUserPeerRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"ca": {
				Type:        schema.TypeString,
				Description: "Name of the CA certificate.",
				Computed:    true,
			},
			"cn": {
				Type:        schema.TypeString,
				Description: "Peer certificate common name.",
				Computed:    true,
			},
			"cn_type": {
				Type:        schema.TypeString,
				Description: "Peer certificate common name type.",
				Computed:    true,
			},
			"ldap_mode": {
				Type:        schema.TypeString,
				Description: "Mode for LDAP peer authentication.",
				Computed:    true,
			},
			"ldap_password": {
				Type:        schema.TypeString,
				Description: "Password for LDAP server bind.",
				Computed:    true,
				Sensitive:   true,
			},
			"ldap_server": {
				Type:        schema.TypeString,
				Description: "Name of an LDAP server defined under the user ldap command. Performs client access rights check.",
				Computed:    true,
			},
			"ldap_username": {
				Type:        schema.TypeString,
				Description: "Username for LDAP server bind.",
				Computed:    true,
			},
			"mandatory_ca_verify": {
				Type:        schema.TypeString,
				Description: "Determine what happens to the peer if the CA certificate is not installed. Disable to automatically consider the peer certificate as valid.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Peer name.",
				Required:    true,
			},
			"ocsp_override_server": {
				Type:        schema.TypeString,
				Description: "Online Certificate Status Protocol (OCSP) server for certificate retrieval.",
				Computed:    true,
			},
			"passwd": {
				Type:        schema.TypeString,
				Description: "Peer's password used for two-factor authentication.",
				Computed:    true,
				Sensitive:   true,
			},
			"subject": {
				Type:        schema.TypeString,
				Description: "Peer certificate name constraints.",
				Computed:    true,
			},
			"two_factor": {
				Type:        schema.TypeString,
				Description: "Enable/disable two-factor authentication, applying certificate and password-based authentication.",
				Computed:    true,
			},
		},
	}
}

func dataSourceUserPeerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserPeer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserPeer dataSource: %v", err)
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

	diags := refreshObjectUserPeer(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
