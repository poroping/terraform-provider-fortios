// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure WAN optimization content delivery network rules.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWanoptContentDeliveryNetworkRule() *schema.Resource {
	return &schema.Resource{
		Description: "Configure WAN optimization content delivery network rules.",

		ReadContext: dataSourceWanoptContentDeliveryNetworkRuleRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"category": {
				Type:        schema.TypeString,
				Description: "Content delivery network rule category.",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment about this CDN-rule.",
				Computed:    true,
			},
			"host_domain_name_suffix": {
				Type:        schema.TypeList,
				Description: "Suffix portion of the fully qualified domain name. For example, fortinet.com in \"www.fortinet.com\".",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Suffix portion of the fully qualified domain name.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of table.",
				Required:    true,
			},
			"request_cache_control": {
				Type:        schema.TypeString,
				Description: "Enable/disable HTTP request cache control.",
				Computed:    true,
			},
			"response_cache_control": {
				Type:        schema.TypeString,
				Description: "Enable/disable HTTP response cache control.",
				Computed:    true,
			},
			"response_expires": {
				Type:        schema.TypeString,
				Description: "Enable/disable HTTP response cache expires.",
				Computed:    true,
			},
			"rules": {
				Type:        schema.TypeList,
				Description: "WAN optimization content delivery network rule entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content_id": {
							Type:        schema.TypeList,
							Description: "Content ID settings.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_direction": {
										Type:        schema.TypeString,
										Description: "Search direction from end-str match.",
										Computed:    true,
									},
									"end_skip": {
										Type:        schema.TypeInt,
										Description: "Number of characters in URL to skip after end-str has been matched.",
										Computed:    true,
									},
									"end_str": {
										Type:        schema.TypeString,
										Description: "String from which to end search.",
										Computed:    true,
									},
									"range_str": {
										Type:        schema.TypeString,
										Description: "Name of content ID within the start string and end string.",
										Computed:    true,
									},
									"start_direction": {
										Type:        schema.TypeString,
										Description: "Search direction from start-str match.",
										Computed:    true,
									},
									"start_skip": {
										Type:        schema.TypeInt,
										Description: "Number of characters in URL to skip after start-str has been matched.",
										Computed:    true,
									},
									"start_str": {
										Type:        schema.TypeString,
										Description: "String from which to start search.",
										Computed:    true,
									},
									"target": {
										Type:        schema.TypeString,
										Description: "Option in HTTP header or URL parameter to match.",
										Computed:    true,
									},
								},
							},
						},
						"match_entries": {
							Type:        schema.TypeList,
							Description: "List of entries to match.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Description: "Rule ID.",
										Computed:    true,
									},
									"pattern": {
										Type:        schema.TypeList,
										Description: "Pattern string for matching target (Referrer or URL pattern). For example, a, a*c, *a*, a*c*e, and *.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"string": {
													Type:        schema.TypeString,
													Description: "Pattern strings.",
													Computed:    true,
												},
											},
										},
									},
									"target": {
										Type:        schema.TypeString,
										Description: "Option in HTTP header or URL parameter to match.",
										Computed:    true,
									},
								},
							},
						},
						"match_mode": {
							Type:        schema.TypeString,
							Description: "Match criteria for collecting content ID.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "WAN optimization content delivery network rule name.",
							Computed:    true,
						},
						"skip_entries": {
							Type:        schema.TypeList,
							Description: "List of entries to skip.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Description: "Rule ID.",
										Computed:    true,
									},
									"pattern": {
										Type:        schema.TypeList,
										Description: "Pattern string for matching target (Referrer or URL pattern). For example, a, a*c, *a*, a*c*e, and *.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"string": {
													Type:        schema.TypeString,
													Description: "Pattern strings.",
													Computed:    true,
												},
											},
										},
									},
									"target": {
										Type:        schema.TypeString,
										Description: "Option in HTTP header or URL parameter to match.",
										Computed:    true,
									},
								},
							},
						},
						"skip_rule_mode": {
							Type:        schema.TypeString,
							Description: "Skip mode when evaluating skip-rules.",
							Computed:    true,
						},
					},
				},
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable WAN optimization content delivery network rules.",
				Computed:    true,
			},
			"updateserver": {
				Type:        schema.TypeString,
				Description: "Enable/disable update server.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWanoptContentDeliveryNetworkRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWanoptContentDeliveryNetworkRule(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WanoptContentDeliveryNetworkRule dataSource: %v", err)
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

	diags := refreshObjectWanoptContentDeliveryNetworkRule(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
