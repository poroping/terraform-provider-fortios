// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure application control lists.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceApplicationList() *schema.Resource {
	return &schema.Resource{
		Description: "Configure application control lists.",

		CreateContext: resourceApplicationListCreate,
		ReadContext:   resourceApplicationListRead,
		UpdateContext: resourceApplicationListUpdate,
		DeleteContext: resourceApplicationListDelete,

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
			"app_replacemsg": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable replacement messages for blocked applications.",
				Optional:    true,
				Computed:    true,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comments.",
				Optional:    true,
				Computed:    true,
			},
			"control_default_network_services": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable enforcement of protocols over selected ports.",
				Optional:    true,
				Computed:    true,
			},
			"deep_app_inspection": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable deep application inspection.",
				Optional:    true,
				Computed:    true,
			},
			"default_network_services": {
				Type:        schema.TypeList,
				Description: "Default network service entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Port number.",
							Optional:    true,
							Computed:    true,
						},
						"services": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Network protocols.",
							Optional:         true,
							Computed:         true,
						},
						"violation_action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"pass", "monitor", "block"}, false),

							Description: "Action for protocols not in the allowlist for selected port.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"enforce_default_app_port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable default application port enforcement for allowed applications.",
				Optional:    true,
				Computed:    true,
			},
			"entries": {
				Type:        schema.TypeList,
				Description: "Application list entries.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"pass", "block", "reset"}, false),

							Description: "Pass or block traffic, or reset connection for traffic from this application.",
							Optional:    true,
							Computed:    true,
						},
						"application": {
							Type:        schema.TypeList,
							Description: "ID of allowed applications.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type: schema.TypeInt,

										Description: "Application IDs.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"behavior": {
							Type: schema.TypeString,

							Description: "Application behavior filter.",
							Optional:    true,
							Computed:    true,
						},
						"category": {
							Type:        schema.TypeList,
							Description: "Category ID list.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type: schema.TypeInt,

										Description: "Application category ID.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"exclusion": {
							Type:        schema.TypeList,
							Description: "ID of excluded applications.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type: schema.TypeInt,

										Description: "Excluded application IDs.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"log": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable logging for this application list.",
							Optional:    true,
							Computed:    true,
						},
						"log_packet": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable packet logging.",
							Optional:    true,
							Computed:    true,
						},
						"parameters": {
							Type:        schema.TypeList,
							Description: "Application parameters.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type: schema.TypeInt,

										Description: "Parameter tuple ID.",
										Optional:    true,
										Computed:    true,
									},
									"members": {
										Type:        schema.TypeList,
										Description: "Parameter tuple members.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": {
													Type: schema.TypeInt,

													Description: "Parameter.",
													Optional:    true,
													Computed:    true,
												},
												"name": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 31),

													Description: "Parameter name.",
													Optional:    true,
													Computed:    true,
												},
												"value": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 199),

													Description: "Parameter value.",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"value": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),

										Description: "Parameter value.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"per_ip_shaper": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Per-IP traffic shaper.",
							Optional:    true,
							Computed:    true,
						},
						"popularity": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Application popularity filter (1 - 5, from least to most popular).",
							Optional:         true,
							Computed:         true,
						},
						"protocols": {
							Type: schema.TypeString,

							Description: "Application protocol filter.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "attacker"}, false),

							Description: "Quarantine method.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine_expiry": {
							Type: schema.TypeString,

							Description: "Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine_log": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable quarantine logging.",
							Optional:    true,
							Computed:    true,
						},
						"rate_count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Count of the rate.",
							Optional:    true,
							Computed:    true,
						},
						"rate_duration": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Duration (sec) of the rate.",
							Optional:    true,
							Computed:    true,
						},
						"rate_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"periodical", "continuous"}, false),

							Description: "Rate limit mode.",
							Optional:    true,
							Computed:    true,
						},
						"rate_track": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "src-ip", "dest-ip", "dhcp-client-mac", "dns-domain"}, false),

							Description: "Track the packet protocol field.",
							Optional:    true,
							Computed:    true,
						},
						"risk": {
							Type:        schema.TypeList,
							Description: "Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical).",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"level": {
										Type: schema.TypeInt,

										Description: "Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical).",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"session_ttl": {
							Type: schema.TypeInt,

							Description: "Session TTL (0 = default).",
							Optional:    true,
							Computed:    true,
						},
						"shaper": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Traffic shaper.",
							Optional:    true,
							Computed:    true,
						},
						"shaper_reverse": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Reverse traffic shaper.",
							Optional:    true,
							Computed:    true,
						},
						"sub_category": {
							Type:        schema.TypeList,
							Description: "Application Sub-category ID list.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type: schema.TypeInt,

										Description: "Application sub-category ID.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"technology": {
							Type: schema.TypeString,

							Description: "Application technology filter.",
							Optional:    true,
							Computed:    true,
						},
						"vendor": {
							Type: schema.TypeString,

							Description: "Application vendor filter.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"extended_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable extended logging.",
				Optional:    true,
				Computed:    true,
			},
			"force_inclusion_ssl_di_sigs": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable forced inclusion of SSL deep inspection signatures.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "List name.",
				ForceNew:    true,
				Required:    true,
			},
			"options": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Basic application protocol signatures allowed by default.",
				Optional:         true,
				Computed:         true,
			},
			"other_application_action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"pass", "block"}, false),

				Description: "Action for other applications.",
				Optional:    true,
				Computed:    true,
			},
			"other_application_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable logging for other applications.",
				Optional:    true,
				Computed:    true,
			},
			"p2p_black_list": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "P2P applications to be black listed.",
				Optional:         true,
				Computed:         true,
			},
			"p2p_block_list": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "P2P applications to be block listed.",
				Optional:         true,
				Computed:         true,
			},
			"replacemsg_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Replacement message group.",
				Optional:    true,
				Computed:    true,
			},
			"unknown_application_action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"pass", "block"}, false),

				Description: "Pass or block traffic from unknown applications.",
				Optional:    true,
				Computed:    true,
			},
			"unknown_application_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable logging for unknown applications.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceApplicationListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating ApplicationList resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectApplicationList(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateApplicationList(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ApplicationList")
	}

	return resourceApplicationListRead(ctx, d, meta)
}

func resourceApplicationListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectApplicationList(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateApplicationList(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating ApplicationList resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ApplicationList")
	}

	return resourceApplicationListRead(ctx, d, meta)
}

func resourceApplicationListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteApplicationList(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting ApplicationList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceApplicationListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadApplicationList(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading ApplicationList resource: %v", err)
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

	diags := refreshObjectApplicationList(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenApplicationListDefaultNetworkServices(v *[]models.ApplicationListDefaultNetworkServices, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Port; tmp != nil {
				v["port"] = *tmp
			}

			if tmp := cfg.Services; tmp != nil {
				v["services"] = *tmp
			}

			if tmp := cfg.ViolationAction; tmp != nil {
				v["violation_action"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenApplicationListEntries(v *[]models.ApplicationListEntries, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Application; tmp != nil {
				v["application"] = flattenApplicationListEntriesApplication(tmp, sort)
			}

			if tmp := cfg.Behavior; tmp != nil {
				v["behavior"] = *tmp
			}

			if tmp := cfg.Category; tmp != nil {
				v["category"] = flattenApplicationListEntriesCategory(tmp, sort)
			}

			if tmp := cfg.Exclusion; tmp != nil {
				v["exclusion"] = flattenApplicationListEntriesExclusion(tmp, sort)
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.LogPacket; tmp != nil {
				v["log_packet"] = *tmp
			}

			if tmp := cfg.Parameters; tmp != nil {
				v["parameters"] = flattenApplicationListEntriesParameters(tmp, sort)
			}

			if tmp := cfg.PerIpShaper; tmp != nil {
				v["per_ip_shaper"] = *tmp
			}

			if tmp := cfg.Popularity; tmp != nil {
				v["popularity"] = *tmp
			}

			if tmp := cfg.Protocols; tmp != nil {
				v["protocols"] = *tmp
			}

			if tmp := cfg.Quarantine; tmp != nil {
				v["quarantine"] = *tmp
			}

			if tmp := cfg.QuarantineExpiry; tmp != nil {
				v["quarantine_expiry"] = *tmp
			}

			if tmp := cfg.QuarantineLog; tmp != nil {
				v["quarantine_log"] = *tmp
			}

			if tmp := cfg.RateCount; tmp != nil {
				v["rate_count"] = *tmp
			}

			if tmp := cfg.RateDuration; tmp != nil {
				v["rate_duration"] = *tmp
			}

			if tmp := cfg.RateMode; tmp != nil {
				v["rate_mode"] = *tmp
			}

			if tmp := cfg.RateTrack; tmp != nil {
				v["rate_track"] = *tmp
			}

			if tmp := cfg.Risk; tmp != nil {
				v["risk"] = flattenApplicationListEntriesRisk(tmp, sort)
			}

			if tmp := cfg.SessionTtl; tmp != nil {
				v["session_ttl"] = *tmp
			}

			if tmp := cfg.Shaper; tmp != nil {
				v["shaper"] = *tmp
			}

			if tmp := cfg.ShaperReverse; tmp != nil {
				v["shaper_reverse"] = *tmp
			}

			if tmp := cfg.SubCategory; tmp != nil {
				v["sub_category"] = flattenApplicationListEntriesSubCategory(tmp, sort)
			}

			if tmp := cfg.Technology; tmp != nil {
				v["technology"] = *tmp
			}

			if tmp := cfg.Vendor; tmp != nil {
				v["vendor"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenApplicationListEntriesApplication(v *[]models.ApplicationListEntriesApplication, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenApplicationListEntriesCategory(v *[]models.ApplicationListEntriesCategory, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenApplicationListEntriesExclusion(v *[]models.ApplicationListEntriesExclusion, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenApplicationListEntriesParameters(v *[]models.ApplicationListEntriesParameters, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Members; tmp != nil {
				v["members"] = flattenApplicationListEntriesParametersMembers(tmp, sort)
			}

			if tmp := cfg.Value; tmp != nil {
				v["value"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenApplicationListEntriesParametersMembers(v *[]models.ApplicationListEntriesParametersMembers, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Value; tmp != nil {
				v["value"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenApplicationListEntriesRisk(v *[]models.ApplicationListEntriesRisk, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Level; tmp != nil {
				v["level"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "level")
	}

	return flat
}

func flattenApplicationListEntriesSubCategory(v *[]models.ApplicationListEntriesSubCategory, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectApplicationList(d *schema.ResourceData, o *models.ApplicationList, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AppReplacemsg != nil {
		v := *o.AppReplacemsg

		if err = d.Set("app_replacemsg", v); err != nil {
			return diag.Errorf("error reading app_replacemsg: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.ControlDefaultNetworkServices != nil {
		v := *o.ControlDefaultNetworkServices

		if err = d.Set("control_default_network_services", v); err != nil {
			return diag.Errorf("error reading control_default_network_services: %v", err)
		}
	}

	if o.DeepAppInspection != nil {
		v := *o.DeepAppInspection

		if err = d.Set("deep_app_inspection", v); err != nil {
			return diag.Errorf("error reading deep_app_inspection: %v", err)
		}
	}

	if o.DefaultNetworkServices != nil {
		if err = d.Set("default_network_services", flattenApplicationListDefaultNetworkServices(o.DefaultNetworkServices, sort)); err != nil {
			return diag.Errorf("error reading default_network_services: %v", err)
		}
	}

	if o.EnforceDefaultAppPort != nil {
		v := *o.EnforceDefaultAppPort

		if err = d.Set("enforce_default_app_port", v); err != nil {
			return diag.Errorf("error reading enforce_default_app_port: %v", err)
		}
	}

	if o.Entries != nil {
		if err = d.Set("entries", flattenApplicationListEntries(o.Entries, sort)); err != nil {
			return diag.Errorf("error reading entries: %v", err)
		}
	}

	if o.ExtendedLog != nil {
		v := *o.ExtendedLog

		if err = d.Set("extended_log", v); err != nil {
			return diag.Errorf("error reading extended_log: %v", err)
		}
	}

	if o.ForceInclusionSslDiSigs != nil {
		v := *o.ForceInclusionSslDiSigs

		if err = d.Set("force_inclusion_ssl_di_sigs", v); err != nil {
			return diag.Errorf("error reading force_inclusion_ssl_di_sigs: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Options != nil {
		v := *o.Options

		if err = d.Set("options", v); err != nil {
			return diag.Errorf("error reading options: %v", err)
		}
	}

	if o.OtherApplicationAction != nil {
		v := *o.OtherApplicationAction

		if err = d.Set("other_application_action", v); err != nil {
			return diag.Errorf("error reading other_application_action: %v", err)
		}
	}

	if o.OtherApplicationLog != nil {
		v := *o.OtherApplicationLog

		if err = d.Set("other_application_log", v); err != nil {
			return diag.Errorf("error reading other_application_log: %v", err)
		}
	}

	if o.P2pBlackList != nil {
		v := *o.P2pBlackList

		if err = d.Set("p2p_black_list", v); err != nil {
			return diag.Errorf("error reading p2p_black_list: %v", err)
		}
	}

	if o.P2pBlockList != nil {
		v := *o.P2pBlockList

		if err = d.Set("p2p_block_list", v); err != nil {
			return diag.Errorf("error reading p2p_block_list: %v", err)
		}
	}

	if o.ReplacemsgGroup != nil {
		v := *o.ReplacemsgGroup

		if err = d.Set("replacemsg_group", v); err != nil {
			return diag.Errorf("error reading replacemsg_group: %v", err)
		}
	}

	if o.UnknownApplicationAction != nil {
		v := *o.UnknownApplicationAction

		if err = d.Set("unknown_application_action", v); err != nil {
			return diag.Errorf("error reading unknown_application_action: %v", err)
		}
	}

	if o.UnknownApplicationLog != nil {
		v := *o.UnknownApplicationLog

		if err = d.Set("unknown_application_log", v); err != nil {
			return diag.Errorf("error reading unknown_application_log: %v", err)
		}
	}

	return nil
}

func expandApplicationListDefaultNetworkServices(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ApplicationListDefaultNetworkServices, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ApplicationListDefaultNetworkServices

	for i := range l {
		tmp := models.ApplicationListDefaultNetworkServices{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Port = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.services", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Services = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.violation_action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ViolationAction = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandApplicationListEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ApplicationListEntries, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ApplicationListEntries

	for i := range l {
		tmp := models.ApplicationListEntries{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.application", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandApplicationListEntriesApplication(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ApplicationListEntriesApplication
			// 	}
			tmp.Application = v2

		}

		pre_append = fmt.Sprintf("%s.%d.behavior", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Behavior = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.category", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandApplicationListEntriesCategory(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ApplicationListEntriesCategory
			// 	}
			tmp.Category = v2

		}

		pre_append = fmt.Sprintf("%s.%d.exclusion", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandApplicationListEntriesExclusion(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ApplicationListEntriesExclusion
			// 	}
			tmp.Exclusion = v2

		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log_packet", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogPacket = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.parameters", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandApplicationListEntriesParameters(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ApplicationListEntriesParameters
			// 	}
			tmp.Parameters = v2

		}

		pre_append = fmt.Sprintf("%s.%d.per_ip_shaper", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PerIpShaper = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.popularity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Popularity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.protocols", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Protocols = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Quarantine = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine_expiry", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.QuarantineExpiry = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine_log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.QuarantineLog = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rate_count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.RateCount = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rate_duration", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.RateDuration = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rate_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RateMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rate_track", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RateTrack = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.risk", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandApplicationListEntriesRisk(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ApplicationListEntriesRisk
			// 	}
			tmp.Risk = v2

		}

		pre_append = fmt.Sprintf("%s.%d.session_ttl", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.SessionTtl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.shaper", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Shaper = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.shaper_reverse", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ShaperReverse = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sub_category", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandApplicationListEntriesSubCategory(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ApplicationListEntriesSubCategory
			// 	}
			tmp.SubCategory = v2

		}

		pre_append = fmt.Sprintf("%s.%d.technology", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Technology = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vendor", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Vendor = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandApplicationListEntriesApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ApplicationListEntriesApplication, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ApplicationListEntriesApplication

	for i := range l {
		tmp := models.ApplicationListEntriesApplication{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandApplicationListEntriesCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ApplicationListEntriesCategory, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ApplicationListEntriesCategory

	for i := range l {
		tmp := models.ApplicationListEntriesCategory{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandApplicationListEntriesExclusion(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ApplicationListEntriesExclusion, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ApplicationListEntriesExclusion

	for i := range l {
		tmp := models.ApplicationListEntriesExclusion{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandApplicationListEntriesParameters(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ApplicationListEntriesParameters, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ApplicationListEntriesParameters

	for i := range l {
		tmp := models.ApplicationListEntriesParameters{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.members", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandApplicationListEntriesParametersMembers(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ApplicationListEntriesParametersMembers
			// 	}
			tmp.Members = v2

		}

		pre_append = fmt.Sprintf("%s.%d.value", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Value = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandApplicationListEntriesParametersMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ApplicationListEntriesParametersMembers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ApplicationListEntriesParametersMembers

	for i := range l {
		tmp := models.ApplicationListEntriesParametersMembers{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.value", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Value = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandApplicationListEntriesRisk(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ApplicationListEntriesRisk, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ApplicationListEntriesRisk

	for i := range l {
		tmp := models.ApplicationListEntriesRisk{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Level = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandApplicationListEntriesSubCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ApplicationListEntriesSubCategory, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ApplicationListEntriesSubCategory

	for i := range l {
		tmp := models.ApplicationListEntriesSubCategory{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectApplicationList(d *schema.ResourceData, sv string) (*models.ApplicationList, diag.Diagnostics) {
	obj := models.ApplicationList{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("app_replacemsg"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("app_replacemsg", sv)
				diags = append(diags, e)
			}
			obj.AppReplacemsg = &v2
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
	if v1, ok := d.GetOk("control_default_network_services"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("control_default_network_services", sv)
				diags = append(diags, e)
			}
			obj.ControlDefaultNetworkServices = &v2
		}
	}
	if v1, ok := d.GetOk("deep_app_inspection"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("deep_app_inspection", sv)
				diags = append(diags, e)
			}
			obj.DeepAppInspection = &v2
		}
	}
	if v, ok := d.GetOk("default_network_services"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("default_network_services", sv)
			diags = append(diags, e)
		}
		t, err := expandApplicationListDefaultNetworkServices(d, v, "default_network_services", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DefaultNetworkServices = t
		}
	} else if d.HasChange("default_network_services") {
		old, new := d.GetChange("default_network_services")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DefaultNetworkServices = &[]models.ApplicationListDefaultNetworkServices{}
		}
	}
	if v1, ok := d.GetOk("enforce_default_app_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("enforce_default_app_port", sv)
				diags = append(diags, e)
			}
			obj.EnforceDefaultAppPort = &v2
		}
	}
	if v, ok := d.GetOk("entries"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("entries", sv)
			diags = append(diags, e)
		}
		t, err := expandApplicationListEntries(d, v, "entries", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Entries = t
		}
	} else if d.HasChange("entries") {
		old, new := d.GetChange("entries")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Entries = &[]models.ApplicationListEntries{}
		}
	}
	if v1, ok := d.GetOk("extended_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("extended_log", sv)
				diags = append(diags, e)
			}
			obj.ExtendedLog = &v2
		}
	}
	if v1, ok := d.GetOk("force_inclusion_ssl_di_sigs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("force_inclusion_ssl_di_sigs", sv)
				diags = append(diags, e)
			}
			obj.ForceInclusionSslDiSigs = &v2
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
	if v1, ok := d.GetOk("options"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("options", sv)
				diags = append(diags, e)
			}
			obj.Options = &v2
		}
	}
	if v1, ok := d.GetOk("other_application_action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("other_application_action", sv)
				diags = append(diags, e)
			}
			obj.OtherApplicationAction = &v2
		}
	}
	if v1, ok := d.GetOk("other_application_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("other_application_log", sv)
				diags = append(diags, e)
			}
			obj.OtherApplicationLog = &v2
		}
	}
	if v1, ok := d.GetOk("p2p_black_list"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("p2p_black_list", sv)
				diags = append(diags, e)
			}
			obj.P2pBlackList = &v2
		}
	}
	if v1, ok := d.GetOk("p2p_block_list"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("p2p_block_list", sv)
				diags = append(diags, e)
			}
			obj.P2pBlockList = &v2
		}
	}
	if v1, ok := d.GetOk("replacemsg_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("replacemsg_group", sv)
				diags = append(diags, e)
			}
			obj.ReplacemsgGroup = &v2
		}
	}
	if v1, ok := d.GetOk("unknown_application_action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("unknown_application_action", sv)
				diags = append(diags, e)
			}
			obj.UnknownApplicationAction = &v2
		}
	}
	if v1, ok := d.GetOk("unknown_application_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("unknown_application_log", sv)
				diags = append(diags, e)
			}
			obj.UnknownApplicationLog = &v2
		}
	}
	return &obj, diags
}
