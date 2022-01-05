// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure DNS domain filter profile.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceDnsfilterProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure DNS domain filter profile.",

		ReadContext: dataSourceDnsfilterProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"block_action": {
				Type:        schema.TypeString,
				Description: "Action to take for blocked domains.",
				Computed:    true,
			},
			"block_botnet": {
				Type:        schema.TypeString,
				Description: "Enable/disable blocking botnet C&C DNS lookups.",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"dns_translation": {
				Type:        schema.TypeList,
				Description: "DNS translation settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr_type": {
							Type:        schema.TypeString,
							Description: "DNS translation type (IPv4 or IPv6).",
							Computed:    true,
						},
						"dst": {
							Type:        schema.TypeString,
							Description: "IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.",
							Computed:    true,
						},
						"dst6": {
							Type:        schema.TypeString,
							Description: "IPv6 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src6.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"netmask": {
							Type:        schema.TypeString,
							Description: "If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.",
							Computed:    true,
						},
						"prefix": {
							Type:        schema.TypeInt,
							Description: "If src6 and dst6 are subnets rather than single IP addresses, enter the prefix for both src6 and dst6 (1 - 128, default = 128).",
							Computed:    true,
						},
						"src": {
							Type:        schema.TypeString,
							Description: "IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.",
							Computed:    true,
						},
						"src6": {
							Type:        schema.TypeString,
							Description: "IPv6 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst6.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable this DNS translation entry.",
							Computed:    true,
						},
					},
				},
			},
			"domain_filter": {
				Type:        schema.TypeList,
				Description: "Domain filter settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_filter_table": {
							Type:        schema.TypeInt,
							Description: "DNS domain filter table ID.",
							Computed:    true,
						},
					},
				},
			},
			"external_ip_blocklist": {
				Type:        schema.TypeList,
				Description: "One or more external IP block lists.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "External domain block list name.",
							Computed:    true,
						},
					},
				},
			},
			"ftgd_dns": {
				Type:        schema.TypeList,
				Description: "FortiGuard DNS Filter settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filters": {
							Type:        schema.TypeList,
							Description: "FortiGuard DNS domain filters.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Description: "Action to take for DNS requests matching the category.",
										Computed:    true,
									},
									"category": {
										Type:        schema.TypeInt,
										Description: "Category number.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "ID number.",
										Computed:    true,
									},
									"log": {
										Type:        schema.TypeString,
										Description: "Enable/disable DNS filter logging for this DNS profile.",
										Computed:    true,
									},
								},
							},
						},
						"options": {
							Type:        schema.TypeString,
							Description: "FortiGuard DNS filter options.",
							Computed:    true,
						},
					},
				},
			},
			"log_all_domain": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging of all domains visited (detailed DNS logging).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Profile name.",
				Required:    true,
			},
			"redirect_portal": {
				Type:        schema.TypeString,
				Description: "IPv4 address of the SDNS redirect portal.",
				Computed:    true,
			},
			"redirect_portal6": {
				Type:        schema.TypeString,
				Description: "IPv6 address of the SDNS redirect portal.",
				Computed:    true,
			},
			"safe_search": {
				Type:        schema.TypeString,
				Description: "Enable/disable Google, Bing, YouTube, Qwant, DuckDuckGo safe search.",
				Computed:    true,
			},
			"sdns_domain_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable domain filtering and botnet domain logging.",
				Computed:    true,
			},
			"sdns_ftgd_err_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiGuard SDNS rating error logging.",
				Computed:    true,
			},
			"youtube_restrict": {
				Type:        schema.TypeString,
				Description: "Set safe search for YouTube restriction level.",
				Computed:    true,
			},
		},
	}
}

func dataSourceDnsfilterProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadDnsfilterProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading DnsfilterProfile dataSource: %v", err)
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

	diags := refreshObjectDnsfilterProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
