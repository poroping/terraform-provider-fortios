// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Authentication rule for SSL-VPN.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceVpnSslSettingsAuthenticationRule() *schema.Resource {
	return &schema.Resource{
		Description: "Authentication rule for SSL-VPN.",

		ReadContext: dataSourceVpnSslSettingsAuthenticationRuleRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"auth": {
				Type:        schema.TypeString,
				Description: "SSL-VPN authentication method restriction.",
				Computed:    true,
			},
			"cipher": {
				Type:        schema.TypeString,
				Description: "SSL-VPN cipher strength.",
				Computed:    true,
			},
			"client_cert": {
				Type:        schema.TypeString,
				Description: "Enable/disable SSL-VPN client certificate restrictive.",
				Computed:    true,
			},
			"groups": {
				Type:        schema.TypeList,
				Description: "User groups.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Group name.",
							Computed:    true,
						},
					},
				},
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "ID (0 - 4294967295).",
				Computed:    true,
			},
			"portal": {
				Type:        schema.TypeString,
				Description: "SSL-VPN portal.",
				Computed:    true,
			},
			"realm": {
				Type:        schema.TypeString,
				Description: "SSL-VPN realm.",
				Computed:    true,
			},
			"source_address": {
				Type:        schema.TypeList,
				Description: "Source address of incoming traffic.",
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
			"source_address_negate": {
				Type:        schema.TypeString,
				Description: "Enable/disable negated source address match.",
				Computed:    true,
			},
			"source_address6": {
				Type:        schema.TypeList,
				Description: "IPv6 source address of incoming traffic.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "IPv6 address name.",
							Computed:    true,
						},
					},
				},
			},
			"source_address6_negate": {
				Type:        schema.TypeString,
				Description: "Enable/disable negated source IPv6 address match.",
				Computed:    true,
			},
			"source_interface": {
				Type:        schema.TypeList,
				Description: "SSL-VPN source interface of incoming traffic.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
					},
				},
			},
			"user_peer": {
				Type:        schema.TypeString,
				Description: "Name of user peer.",
				Computed:    true,
			},
			"users": {
				Type:        schema.TypeList,
				Description: "User name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "User name.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceVpnSslSettingsAuthenticationRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("fosid")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadVpnSslSettingsAuthenticationRule(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnSslSettingsAuthenticationRule dataSource: %v", err)
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

	diags := refreshObjectVpnSslSettingsAuthenticationRule(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
