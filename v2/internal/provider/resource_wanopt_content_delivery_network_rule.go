// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure WAN optimization content delivery network rules.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceWanoptContentDeliveryNetworkRule() *schema.Resource {
	return &schema.Resource{
		Description: "Configure WAN optimization content delivery network rules.",

		CreateContext: resourceWanoptContentDeliveryNetworkRuleCreate,
		ReadContext:   resourceWanoptContentDeliveryNetworkRuleRead,
		UpdateContext: resourceWanoptContentDeliveryNetworkRuleUpdate,
		DeleteContext: resourceWanoptContentDeliveryNetworkRuleDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"allow_append": {
				Type:        schema.TypeBool,
				Description: "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"category": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"vcache", "youtube"}, false),

				Description: "Content delivery network rule category.",
				Optional:    true,
				Computed:    true,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment about this CDN-rule.",
				Optional:    true,
				Computed:    true,
			},
			"host_domain_name_suffix": {
				Type:        schema.TypeList,
				Description: "Suffix portion of the fully qualified domain name. For example, fortinet.com in \"www.fortinet.com\".",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Suffix portion of the fully qualified domain name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of table.",
				ForceNew:    true,
				Required:    true,
			},
			"request_cache_control": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable HTTP request cache control.",
				Optional:    true,
				Computed:    true,
			},
			"response_cache_control": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable HTTP response cache control.",
				Optional:    true,
				Computed:    true,
			},
			"response_expires": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable HTTP response cache expires.",
				Optional:    true,
				Computed:    true,
			},
			"rules": {
				Type:        schema.TypeList,
				Description: "WAN optimization content delivery network rule entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content_id": {
							Type:        schema.TypeList,
							Description: "Content ID settings.",
							Optional:    true, MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_direction": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"forward", "backward"}, false),

										Description: "Search direction from end-str match.",
										Optional:    true,
										Computed:    true,
									},
									"end_skip": {
										Type: schema.TypeInt,

										Description: "Number of characters in URL to skip after end-str has been matched.",
										Optional:    true,
										Computed:    true,
									},
									"end_str": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "String from which to end search.",
										Optional:    true,
										Computed:    true,
									},
									"range_str": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Name of content ID within the start string and end string.",
										Optional:    true,
										Computed:    true,
									},
									"start_direction": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"forward", "backward"}, false),

										Description: "Search direction from start-str match.",
										Optional:    true,
										Computed:    true,
									},
									"start_skip": {
										Type: schema.TypeInt,

										Description: "Number of characters in URL to skip after start-str has been matched.",
										Optional:    true,
										Computed:    true,
									},
									"start_str": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "String from which to start search.",
										Optional:    true,
										Computed:    true,
									},
									"target": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"path", "parameter", "referrer", "youtube-map", "youtube-id", "youku-id", "hls-manifest", "dash-manifest", "hls-fragment", "dash-fragment"}, false),

										Description: "Option in HTTP header or URL parameter to match.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"match_entries": {
							Type:        schema.TypeList,
							Description: "List of entries to match.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type: schema.TypeInt,

										Description: "Rule ID.",
										Optional:    true,
										Computed:    true,
									},
									"pattern": {
										Type:        schema.TypeList,
										Description: "Pattern string for matching target (Referrer or URL pattern). For example, a, a*c, *a*, a*c*e, and *.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"string": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 79),

													Description: "Pattern strings.",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"target": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"path", "parameter", "referrer", "youtube-map", "youtube-id", "youku-id"}, false),

										Description: "Option in HTTP header or URL parameter to match.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"match_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"all", "any"}, false),

							Description: "Match criteria for collecting content ID.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "WAN optimization content delivery network rule name.",
							Optional:    true,
							Computed:    true,
						},
						"skip_entries": {
							Type:        schema.TypeList,
							Description: "List of entries to skip.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type: schema.TypeInt,

										Description: "Rule ID.",
										Optional:    true,
										Computed:    true,
									},
									"pattern": {
										Type:        schema.TypeList,
										Description: "Pattern string for matching target (Referrer or URL pattern). For example, a, a*c, *a*, a*c*e, and *.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"string": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 79),

													Description: "Pattern strings.",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"target": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"path", "parameter", "referrer", "youtube-map", "youtube-id", "youku-id"}, false),

										Description: "Option in HTTP header or URL parameter to match.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"skip_rule_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"all", "any"}, false),

							Description: "Skip mode when evaluating skip-rules.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable WAN optimization content delivery network rules.",
				Optional:    true,
				Computed:    true,
			},
			"updateserver": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable update server.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWanoptContentDeliveryNetworkRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	var diags diag.Diagnostics
	var err error
	// c.Retries = 1
	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	allow_append := false
	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}
	urlparams.AllowAppend = &allow_append

	mkey := ""
	key := "name"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating WanoptContentDeliveryNetworkRule resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWanoptContentDeliveryNetworkRule(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWanoptContentDeliveryNetworkRule(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WanoptContentDeliveryNetworkRule")
	}

	return resourceWanoptContentDeliveryNetworkRuleRead(ctx, d, meta)
}

func resourceWanoptContentDeliveryNetworkRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWanoptContentDeliveryNetworkRule(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWanoptContentDeliveryNetworkRule(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WanoptContentDeliveryNetworkRule resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WanoptContentDeliveryNetworkRule")
	}

	return resourceWanoptContentDeliveryNetworkRuleRead(ctx, d, meta)
}

func resourceWanoptContentDeliveryNetworkRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWanoptContentDeliveryNetworkRule(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WanoptContentDeliveryNetworkRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptContentDeliveryNetworkRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	ptp := true
	urlparams.PlainTextPassword = &ptp

	o, err := c.Cmdb.ReadWanoptContentDeliveryNetworkRule(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WanoptContentDeliveryNetworkRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
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
	return nil
}

func flattenWanoptContentDeliveryNetworkRuleHostDomainNameSuffix(d *schema.ResourceData, v *[]models.WanoptContentDeliveryNetworkRuleHostDomainNameSuffix, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenWanoptContentDeliveryNetworkRuleRules(d *schema.ResourceData, v *[]models.WanoptContentDeliveryNetworkRuleRules, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ContentId; tmp != nil {
				v["content_id"] = flattenWanoptContentDeliveryNetworkRuleRulesContentId(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "content_id"), sort)
			}

			if tmp := cfg.MatchEntries; tmp != nil {
				v["match_entries"] = flattenWanoptContentDeliveryNetworkRuleRulesMatchEntries(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "match_entries"), sort)
			}

			if tmp := cfg.MatchMode; tmp != nil {
				v["match_mode"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.SkipEntries; tmp != nil {
				v["skip_entries"] = flattenWanoptContentDeliveryNetworkRuleRulesSkipEntries(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "skip_entries"), sort)
			}

			if tmp := cfg.SkipRuleMode; tmp != nil {
				v["skip_rule_mode"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentId(d *schema.ResourceData, v *models.WanoptContentDeliveryNetworkRuleRulesContentId, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.WanoptContentDeliveryNetworkRuleRulesContentId{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.EndDirection; tmp != nil {
				v["end_direction"] = *tmp
			}

			if tmp := cfg.EndSkip; tmp != nil {
				v["end_skip"] = *tmp
			}

			if tmp := cfg.EndStr; tmp != nil {
				v["end_str"] = *tmp
			}

			if tmp := cfg.RangeStr; tmp != nil {
				v["range_str"] = *tmp
			}

			if tmp := cfg.StartDirection; tmp != nil {
				v["start_direction"] = *tmp
			}

			if tmp := cfg.StartSkip; tmp != nil {
				v["start_skip"] = *tmp
			}

			if tmp := cfg.StartStr; tmp != nil {
				v["start_str"] = *tmp
			}

			if tmp := cfg.Target; tmp != nil {
				v["target"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenWanoptContentDeliveryNetworkRuleRulesMatchEntries(d *schema.ResourceData, v *[]models.WanoptContentDeliveryNetworkRuleRulesMatchEntries, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Pattern; tmp != nil {
				v["pattern"] = flattenWanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "pattern"), sort)
			}

			if tmp := cfg.Target; tmp != nil {
				v["target"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenWanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern(d *schema.ResourceData, v *[]models.WanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.String; tmp != nil {
				v["string"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "string")
	}

	return flat
}

func flattenWanoptContentDeliveryNetworkRuleRulesSkipEntries(d *schema.ResourceData, v *[]models.WanoptContentDeliveryNetworkRuleRulesSkipEntries, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Pattern; tmp != nil {
				v["pattern"] = flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "pattern"), sort)
			}

			if tmp := cfg.Target; tmp != nil {
				v["target"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern(d *schema.ResourceData, v *[]models.WanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.String; tmp != nil {
				v["string"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "string")
	}

	return flat
}

func refreshObjectWanoptContentDeliveryNetworkRule(d *schema.ResourceData, o *models.WanoptContentDeliveryNetworkRule, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Category != nil {
		v := *o.Category

		if err = d.Set("category", v); err != nil {
			return diag.Errorf("error reading category: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.HostDomainNameSuffix != nil {
		if err = d.Set("host_domain_name_suffix", flattenWanoptContentDeliveryNetworkRuleHostDomainNameSuffix(d, o.HostDomainNameSuffix, "host_domain_name_suffix", sort)); err != nil {
			return diag.Errorf("error reading host_domain_name_suffix: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.RequestCacheControl != nil {
		v := *o.RequestCacheControl

		if err = d.Set("request_cache_control", v); err != nil {
			return diag.Errorf("error reading request_cache_control: %v", err)
		}
	}

	if o.ResponseCacheControl != nil {
		v := *o.ResponseCacheControl

		if err = d.Set("response_cache_control", v); err != nil {
			return diag.Errorf("error reading response_cache_control: %v", err)
		}
	}

	if o.ResponseExpires != nil {
		v := *o.ResponseExpires

		if err = d.Set("response_expires", v); err != nil {
			return diag.Errorf("error reading response_expires: %v", err)
		}
	}

	if o.Rules != nil {
		if err = d.Set("rules", flattenWanoptContentDeliveryNetworkRuleRules(d, o.Rules, "rules", sort)); err != nil {
			return diag.Errorf("error reading rules: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Updateserver != nil {
		v := *o.Updateserver

		if err = d.Set("updateserver", v); err != nil {
			return diag.Errorf("error reading updateserver: %v", err)
		}
	}

	return nil
}

func expandWanoptContentDeliveryNetworkRuleHostDomainNameSuffix(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WanoptContentDeliveryNetworkRuleHostDomainNameSuffix, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WanoptContentDeliveryNetworkRuleHostDomainNameSuffix

	for i := range l {
		tmp := models.WanoptContentDeliveryNetworkRuleHostDomainNameSuffix{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWanoptContentDeliveryNetworkRuleRules(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WanoptContentDeliveryNetworkRuleRules, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WanoptContentDeliveryNetworkRuleRules

	for i := range l {
		tmp := models.WanoptContentDeliveryNetworkRuleRules{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.content_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWanoptContentDeliveryNetworkRuleRulesContentId(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WanoptContentDeliveryNetworkRuleRulesContentId
			// 	}
			tmp.ContentId = v2

		}

		pre_append = fmt.Sprintf("%s.%d.match_entries", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWanoptContentDeliveryNetworkRuleRulesMatchEntries(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WanoptContentDeliveryNetworkRuleRulesMatchEntries
			// 	}
			tmp.MatchEntries = v2

		}

		pre_append = fmt.Sprintf("%s.%d.match_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MatchMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.skip_entries", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWanoptContentDeliveryNetworkRuleRulesSkipEntries(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WanoptContentDeliveryNetworkRuleRulesSkipEntries
			// 	}
			tmp.SkipEntries = v2

		}

		pre_append = fmt.Sprintf("%s.%d.skip_rule_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SkipRuleMode = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentId(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.WanoptContentDeliveryNetworkRuleRulesContentId, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WanoptContentDeliveryNetworkRuleRulesContentId

	for i := range l {
		tmp := models.WanoptContentDeliveryNetworkRuleRulesContentId{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.end_direction", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EndDirection = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.end_skip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.EndSkip = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.end_str", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EndStr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.range_str", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RangeStr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.start_direction", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StartDirection = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.start_skip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.StartSkip = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.start_str", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StartStr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.target", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Target = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandWanoptContentDeliveryNetworkRuleRulesMatchEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WanoptContentDeliveryNetworkRuleRulesMatchEntries, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WanoptContentDeliveryNetworkRuleRulesMatchEntries

	for i := range l {
		tmp := models.WanoptContentDeliveryNetworkRuleRulesMatchEntries{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pattern", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern
			// 	}
			tmp.Pattern = v2

		}

		pre_append = fmt.Sprintf("%s.%d.target", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Target = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern

	for i := range l {
		tmp := models.WanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.String = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesSkipEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WanoptContentDeliveryNetworkRuleRulesSkipEntries, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WanoptContentDeliveryNetworkRuleRulesSkipEntries

	for i := range l {
		tmp := models.WanoptContentDeliveryNetworkRuleRulesSkipEntries{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pattern", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern
			// 	}
			tmp.Pattern = v2

		}

		pre_append = fmt.Sprintf("%s.%d.target", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Target = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern

	for i := range l {
		tmp := models.WanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.String = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWanoptContentDeliveryNetworkRule(d *schema.ResourceData, sv string) (*models.WanoptContentDeliveryNetworkRule, diag.Diagnostics) {
	obj := models.WanoptContentDeliveryNetworkRule{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("category"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("category", sv)
				diags = append(diags, e)
			}
			obj.Category = &v2
		}
	}
	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
		}
	}
	if v, ok := d.GetOk("host_domain_name_suffix"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("host_domain_name_suffix", sv)
			diags = append(diags, e)
		}
		t, err := expandWanoptContentDeliveryNetworkRuleHostDomainNameSuffix(d, v, "host_domain_name_suffix", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.HostDomainNameSuffix = t
		}
	} else if d.HasChange("host_domain_name_suffix") {
		old, new := d.GetChange("host_domain_name_suffix")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.HostDomainNameSuffix = &[]models.WanoptContentDeliveryNetworkRuleHostDomainNameSuffix{}
		}
	}
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v1, ok := d.GetOk("request_cache_control"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("request_cache_control", sv)
				diags = append(diags, e)
			}
			obj.RequestCacheControl = &v2
		}
	}
	if v1, ok := d.GetOk("response_cache_control"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("response_cache_control", sv)
				diags = append(diags, e)
			}
			obj.ResponseCacheControl = &v2
		}
	}
	if v1, ok := d.GetOk("response_expires"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("response_expires", sv)
				diags = append(diags, e)
			}
			obj.ResponseExpires = &v2
		}
	}
	if v, ok := d.GetOk("rules"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("rules", sv)
			diags = append(diags, e)
		}
		t, err := expandWanoptContentDeliveryNetworkRuleRules(d, v, "rules", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Rules = t
		}
	} else if d.HasChange("rules") {
		old, new := d.GetChange("rules")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Rules = &[]models.WanoptContentDeliveryNetworkRuleRules{}
		}
	}
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	if v1, ok := d.GetOk("updateserver"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("updateserver", sv)
				diags = append(diags, e)
			}
			obj.Updateserver = &v2
		}
	}
	return &obj, diags
}
