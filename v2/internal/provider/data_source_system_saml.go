// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Global settings for SAML authentication.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemSaml() *schema.Resource {
	return &schema.Resource{
		Description: "Global settings for SAML authentication.",

		ReadContext: dataSourceSystemSamlRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"binding_protocol": {
				Type:        schema.TypeString,
				Description: "IdP Binding protocol.",
				Computed:    true,
			},
			"cert": {
				Type:        schema.TypeString,
				Description: "Certificate to sign SAML messages.",
				Computed:    true,
			},
			"default_login_page": {
				Type:        schema.TypeString,
				Description: "Choose default login page.",
				Computed:    true,
			},
			"default_profile": {
				Type:        schema.TypeString,
				Description: "Default profile for new SSO admin.",
				Computed:    true,
			},
			"entity_id": {
				Type:        schema.TypeString,
				Description: "SP entity ID.",
				Computed:    true,
			},
			"idp_cert": {
				Type:        schema.TypeString,
				Description: "IDP certificate name.",
				Computed:    true,
			},
			"idp_entity_id": {
				Type:        schema.TypeString,
				Description: "IDP entity ID.",
				Computed:    true,
			},
			"idp_single_logout_url": {
				Type:        schema.TypeString,
				Description: "IDP single logout URL.",
				Computed:    true,
			},
			"idp_single_sign_on_url": {
				Type:        schema.TypeString,
				Description: "IDP single sign-on URL.",
				Computed:    true,
			},
			"life": {
				Type:        schema.TypeInt,
				Description: "Length of the range of time when the assertion is valid (in minutes).",
				Computed:    true,
			},
			"portal_url": {
				Type:        schema.TypeString,
				Description: "SP portal URL.",
				Computed:    true,
			},
			"role": {
				Type:        schema.TypeString,
				Description: "SAML role.",
				Computed:    true,
			},
			"server_address": {
				Type:        schema.TypeString,
				Description: "Server address.",
				Computed:    true,
			},
			"service_providers": {
				Type:        schema.TypeList,
				Description: "Authorized service providers.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"assertion_attributes": {
							Type:        schema.TypeList,
							Description: "Customized SAML attributes to send along with assertion.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Name.",
										Computed:    true,
									},
									"type": {
										Type:        schema.TypeString,
										Description: "Type.",
										Computed:    true,
									},
								},
							},
						},
						"idp_entity_id": {
							Type:        schema.TypeString,
							Description: "IDP entity ID.",
							Computed:    true,
						},
						"idp_single_logout_url": {
							Type:        schema.TypeString,
							Description: "IDP single logout URL.",
							Computed:    true,
						},
						"idp_single_sign_on_url": {
							Type:        schema.TypeString,
							Description: "IDP single sign-on URL.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Name.",
							Computed:    true,
						},
						"prefix": {
							Type:        schema.TypeString,
							Description: "Prefix.",
							Computed:    true,
						},
						"sp_binding_protocol": {
							Type:        schema.TypeString,
							Description: "SP binding protocol.",
							Computed:    true,
						},
						"sp_cert": {
							Type:        schema.TypeString,
							Description: "SP certificate name.",
							Computed:    true,
						},
						"sp_entity_id": {
							Type:        schema.TypeString,
							Description: "SP entity ID.",
							Computed:    true,
						},
						"sp_portal_url": {
							Type:        schema.TypeString,
							Description: "SP portal URL.",
							Computed:    true,
						},
						"sp_single_logout_url": {
							Type:        schema.TypeString,
							Description: "SP single logout URL.",
							Computed:    true,
						},
						"sp_single_sign_on_url": {
							Type:        schema.TypeString,
							Description: "SP single sign-on URL.",
							Computed:    true,
						},
					},
				},
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
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable SAML authentication (default = disable).",
				Computed:    true,
			},
			"tolerance": {
				Type:        schema.TypeInt,
				Description: "Tolerance to the range of time when the assertion is valid (in minutes).",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemSamlRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemSaml"

	o, err := c.Cmdb.ReadSystemSaml(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSaml dataSource: %v", err)
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

	diags := refreshObjectSystemSaml(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
