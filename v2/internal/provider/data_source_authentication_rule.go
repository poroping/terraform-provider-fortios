// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Authentication Rules.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceAuthenticationRule() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Authentication Rules.",

		ReadContext: dataSourceAuthenticationRuleRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"active_auth_method": {
				Type:        schema.TypeString,
				Description: "Select an active authentication method.",
				Computed:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"dstaddr": {
				Type:        schema.TypeList,
				Description: "Select an IPv4 destination address from available options. Required for web proxy authentication.",
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
			"dstaddr6": {
				Type:        schema.TypeList,
				Description: "Select an IPv6 destination address from available options. Required for web proxy authentication.",
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
			"ip_based": {
				Type:        schema.TypeString,
				Description: "Enable/disable IP-based authentication. When enabled, previously authenticated users from the same IP address will be exempted.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Authentication rule name.",
				Required:    true,
			},
			"protocol": {
				Type:        schema.TypeString,
				Description: "Authentication is required for the selected protocol (default = HTTP).",
				Computed:    true,
			},
			"srcaddr": {
				Type:        schema.TypeList,
				Description: "Authentication is required for the selected IPv4 source address.",
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
			"srcaddr6": {
				Type:        schema.TypeList,
				Description: "Authentication is required for the selected IPv6 source address.",
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
			"srcintf": {
				Type:        schema.TypeList,
				Description: "Incoming (ingress) interface.",
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
			"sso_auth_method": {
				Type:        schema.TypeString,
				Description: "Select a single-sign on (SSO) authentication method.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable this authentication rule.",
				Computed:    true,
			},
			"transaction_based": {
				Type:        schema.TypeString,
				Description: "Enable/disable transaction based authentication (default = disable).",
				Computed:    true,
			},
			"web_auth_cookie": {
				Type:        schema.TypeString,
				Description: "Enable/disable Web authentication cookies (default = disable).",
				Computed:    true,
			},
			"web_portal": {
				Type:        schema.TypeString,
				Description: "Enable/disable web portal for proxy transparent policy (default = enable).",
				Computed:    true,
			},
		},
	}
}

func dataSourceAuthenticationRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadAuthenticationRule(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading AuthenticationRule dataSource: %v", err)
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

	diags := refreshObjectAuthenticationRule(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
