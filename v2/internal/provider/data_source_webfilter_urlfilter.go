// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure URL filter lists.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWebfilterUrlfilter() *schema.Resource {
	return &schema.Resource{
		Description: "Configure URL filter lists.",

		ReadContext: dataSourceWebfilterUrlfilterRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Optional comments.",
				Computed:    true,
			},
			"entries": {
				Type:        schema.TypeList,
				Description: "URL filter entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Action to take for URL filter matches.",
							Computed:    true,
						},
						"antiphish_action": {
							Type:        schema.TypeString,
							Description: "Action to take for AntiPhishing matches.",
							Computed:    true,
						},
						"dns_address_family": {
							Type:        schema.TypeString,
							Description: "Resolve IPv4 address, IPv6 address, or both from DNS server.",
							Computed:    true,
						},
						"exempt": {
							Type:        schema.TypeString,
							Description: "If action is set to exempt, select the security profile operations that exempt URLs skip. Separate multiple options with a space.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Id.",
							Computed:    true,
						},
						"referrer_host": {
							Type:        schema.TypeString,
							Description: "Referrer host name.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable this URL filter.",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "Filter type (simple, regex, or wildcard).",
							Computed:    true,
						},
						"url": {
							Type:        schema.TypeString,
							Description: "URL to be filtered.",
							Computed:    true,
						},
						"web_proxy_profile": {
							Type:        schema.TypeString,
							Description: "Web proxy profile.",
							Computed:    true,
						},
					},
				},
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "ID.",
				Computed:    true,
			},
			"ip_addr_block": {
				Type:        schema.TypeString,
				Description: "Enable/disable blocking URLs when the hostname appears as an IP address.",
				Computed:    true,
			},
			"ip4_mapped_ip6": {
				Type:        schema.TypeString,
				Description: "Enable/disable matching of IPv4 mapped IPv6 URLs.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of URL filter list.",
				Computed:    true,
			},
			"one_arm_ips_urlfilter": {
				Type:        schema.TypeString,
				Description: "Enable/disable DNS resolver for one-arm IPS URL filter operation.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWebfilterUrlfilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWebfilterUrlfilter(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebfilterUrlfilter dataSource: %v", err)
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

	diags := refreshObjectWebfilterUrlfilter(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
