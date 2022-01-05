// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure application control lists.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceApplicationList() *schema.Resource {
	return &schema.Resource{
		Description: "Configure application control lists.",

		ReadContext: dataSourceApplicationListRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"app_replacemsg": {
				Type:        schema.TypeString,
				Description: "Enable/disable replacement messages for blocked applications.",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "comments",
				Computed:    true,
			},
			"control_default_network_services": {
				Type:        schema.TypeString,
				Description: "Enable/disable enforcement of protocols over selected ports.",
				Computed:    true,
			},
			"deep_app_inspection": {
				Type:        schema.TypeString,
				Description: "Enable/disable deep application inspection.",
				Computed:    true,
			},
			"default_network_services": {
				Type:        schema.TypeList,
				Description: "Default network service entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Entry ID.",
							Computed:    true,
						},
						"port": {
							Type:        schema.TypeInt,
							Description: "Port number.",
							Computed:    true,
						},
						"services": {
							Type:        schema.TypeString,
							Description: "Network protocols.",
							Computed:    true,
						},
						"violation_action": {
							Type:        schema.TypeString,
							Description: "Action for protocols not in the allowlist for selected port.",
							Computed:    true,
						},
					},
				},
			},
			"enforce_default_app_port": {
				Type:        schema.TypeString,
				Description: "Enable/disable default application port enforcement for allowed applications.",
				Computed:    true,
			},
			"entries": {
				Type:        schema.TypeList,
				Description: "Application list entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Pass or block traffic, or reset connection for traffic from this application.",
							Computed:    true,
						},
						"application": {
							Type:        schema.TypeList,
							Description: "ID of allowed applications.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Description: "Application IDs.",
										Computed:    true,
									},
								},
							},
						},
						"behavior": {
							Type:        schema.TypeString,
							Description: "Application behavior filter.",
							Computed:    true,
						},
						"category": {
							Type:        schema.TypeList,
							Description: "Category ID list.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Description: "Application category ID.",
										Computed:    true,
									},
								},
							},
						},
						"exclusion": {
							Type:        schema.TypeList,
							Description: "ID of excluded applications.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Description: "Excluded application IDs.",
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Entry ID.",
							Computed:    true,
						},
						"log": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging for this application list.",
							Computed:    true,
						},
						"log_packet": {
							Type:        schema.TypeString,
							Description: "Enable/disable packet logging.",
							Computed:    true,
						},
						"parameters": {
							Type:        schema.TypeList,
							Description: "Application parameters.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Description: "Parameter tuple ID.",
										Computed:    true,
									},
									"members": {
										Type:        schema.TypeList,
										Description: "Parameter tuple members.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": {
													Type:        schema.TypeInt,
													Description: "Parameter.",
													Computed:    true,
												},
												"name": {
													Type:        schema.TypeString,
													Description: "Parameter name.",
													Computed:    true,
												},
												"value": {
													Type:        schema.TypeString,
													Description: "Parameter value.",
													Computed:    true,
												},
											},
										},
									},
									"value": {
										Type:        schema.TypeString,
										Description: "Parameter value.",
										Computed:    true,
									},
								},
							},
						},
						"per_ip_shaper": {
							Type:        schema.TypeString,
							Description: "Per-IP traffic shaper.",
							Computed:    true,
						},
						"popularity": {
							Type:        schema.TypeString,
							Description: "Application popularity filter (1 - 5, from least to most popular).",
							Computed:    true,
						},
						"protocols": {
							Type:        schema.TypeString,
							Description: "Application protocol filter.",
							Computed:    true,
						},
						"quarantine": {
							Type:        schema.TypeString,
							Description: "Quarantine method.",
							Computed:    true,
						},
						"quarantine_expiry": {
							Type:        schema.TypeString,
							Description: "Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.",
							Computed:    true,
						},
						"quarantine_log": {
							Type:        schema.TypeString,
							Description: "Enable/disable quarantine logging.",
							Computed:    true,
						},
						"rate_count": {
							Type:        schema.TypeInt,
							Description: "Count of the rate.",
							Computed:    true,
						},
						"rate_duration": {
							Type:        schema.TypeInt,
							Description: "Duration (sec) of the rate.",
							Computed:    true,
						},
						"rate_mode": {
							Type:        schema.TypeString,
							Description: "Rate limit mode.",
							Computed:    true,
						},
						"rate_track": {
							Type:        schema.TypeString,
							Description: "Track the packet protocol field.",
							Computed:    true,
						},
						"risk": {
							Type:        schema.TypeList,
							Description: "Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical).",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"level": {
										Type:        schema.TypeInt,
										Description: "Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical).",
										Computed:    true,
									},
								},
							},
						},
						"session_ttl": {
							Type:        schema.TypeInt,
							Description: "Session TTL (0 = default).",
							Computed:    true,
						},
						"shaper": {
							Type:        schema.TypeString,
							Description: "Traffic shaper.",
							Computed:    true,
						},
						"shaper_reverse": {
							Type:        schema.TypeString,
							Description: "Reverse traffic shaper.",
							Computed:    true,
						},
						"sub_category": {
							Type:        schema.TypeList,
							Description: "Application Sub-category ID list.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Description: "Application sub-category ID.",
										Computed:    true,
									},
								},
							},
						},
						"technology": {
							Type:        schema.TypeString,
							Description: "Application technology filter.",
							Computed:    true,
						},
						"vendor": {
							Type:        schema.TypeString,
							Description: "Application vendor filter.",
							Computed:    true,
						},
					},
				},
			},
			"extended_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable extended logging.",
				Computed:    true,
			},
			"force_inclusion_ssl_di_sigs": {
				Type:        schema.TypeString,
				Description: "Enable/disable forced inclusion of SSL deep inspection signatures.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "List name.",
				Required:    true,
			},
			"options": {
				Type:        schema.TypeString,
				Description: "Basic application protocol signatures allowed by default.",
				Computed:    true,
			},
			"other_application_action": {
				Type:        schema.TypeString,
				Description: "Action for other applications.",
				Computed:    true,
			},
			"other_application_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging for other applications.",
				Computed:    true,
			},
			"p2p_black_list": {
				Type:        schema.TypeString,
				Description: "P2P applications to be black listed.",
				Computed:    true,
			},
			"p2p_block_list": {
				Type:        schema.TypeString,
				Description: "P2P applications to be blocklisted.",
				Computed:    true,
			},
			"replacemsg_group": {
				Type:        schema.TypeString,
				Description: "Replacement message group.",
				Computed:    true,
			},
			"unknown_application_action": {
				Type:        schema.TypeString,
				Description: "Pass or block traffic from unknown applications.",
				Computed:    true,
			},
			"unknown_application_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging for unknown applications.",
				Computed:    true,
			},
		},
	}
}

func dataSourceApplicationListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadApplicationList(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading ApplicationList dataSource: %v", err)
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

	diags := refreshObjectApplicationList(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
