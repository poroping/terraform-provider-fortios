// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: SAML server entry configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceUserSaml() *schema.Resource {
	return &schema.Resource{
		Description: "SAML server entry configuration.",

		ReadContext: dataSourceUserSamlRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"adfs_claim": {
				Type:        schema.TypeString,
				Description: "Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable).",
				Computed:    true,
			},
			"cert": {
				Type:        schema.TypeString,
				Description: "Certificate to sign SAML messages.",
				Computed:    true,
			},
			"digest_method": {
				Type:        schema.TypeString,
				Description: "Digest Method Algorithm. (default = sha1).",
				Computed:    true,
			},
			"entity_id": {
				Type:        schema.TypeString,
				Description: "SP entity ID.",
				Computed:    true,
			},
			"group_claim_type": {
				Type:        schema.TypeString,
				Description: "Group claim in assertion statement.",
				Computed:    true,
			},
			"group_name": {
				Type:        schema.TypeString,
				Description: "Group name in assertion statement.",
				Computed:    true,
			},
			"idp_cert": {
				Type:        schema.TypeString,
				Description: "IDP Certificate name.",
				Computed:    true,
			},
			"idp_entity_id": {
				Type:        schema.TypeString,
				Description: "IDP entity ID.",
				Computed:    true,
			},
			"idp_single_logout_url": {
				Type:        schema.TypeString,
				Description: "IDP single logout url.",
				Computed:    true,
			},
			"idp_single_sign_on_url": {
				Type:        schema.TypeString,
				Description: "IDP single sign-on URL.",
				Computed:    true,
			},
			"limit_relaystate": {
				Type:        schema.TypeString,
				Description: "Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "SAML server entry name.",
				Required:    true,
			},
			"single_logout_url": {
				Type:        schema.TypeString,
				Description: "SP single logout URL.",
				Computed:    true,
			},
			"single_sign_on_url": {
				Type:        schema.TypeString,
				Description: "SP single sign-on URL.",
				Computed:    true,
			},
			"user_claim_type": {
				Type:        schema.TypeString,
				Description: "User name claim in assertion statement.",
				Computed:    true,
			},
			"user_name": {
				Type:        schema.TypeString,
				Description: "User name in assertion statement.",
				Computed:    true,
			},
		},
	}
}

func dataSourceUserSamlRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserSaml(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserSaml dataSource: %v", err)
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

	diags := refreshObjectUserSaml(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
