// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure shaping policies.

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

func resourceFirewallShapingPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure shaping policies.",

		CreateContext: resourceFirewallShapingPolicyCreate,
		ReadContext:   resourceFirewallShapingPolicyRead,
		UpdateContext: resourceFirewallShapingPolicyUpdate,
		DeleteContext: resourceFirewallShapingPolicyDelete,

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
				Type:         schema.TypeBool,
				Description:  "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:     true,
				Default:      false,
				RequiredWith: []string{"fosid"},
			},
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"app_category": {
				Type:        schema.TypeList,
				Description: "IDs of one or more application categories that this shaper applies application control traffic shaping to.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Category IDs.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"app_group": {
				Type:        schema.TypeList,
				Description: "One or more application group names.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Application group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"application": {
				Type:        schema.TypeList,
				Description: "IDs of one or more applications that this shaper applies application control traffic shaping to.",
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
			"class_id": {
				Type: schema.TypeInt,

				Description: "Traffic class ID.",
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
			"diffserv_forward": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to change packet's DiffServ values to the specified diffservcode-forward value.",
				Optional:    true,
				Computed:    true,
			},
			"diffserv_reverse": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value.",
				Optional:    true,
				Computed:    true,
			},
			"diffservcode_forward": {
				Type: schema.TypeString,

				Description: "Change packet's DiffServ to this value.",
				Optional:    true,
				Computed:    true,
			},
			"diffservcode_rev": {
				Type: schema.TypeString,

				Description: "Change packet's reverse (reply) DiffServ to this value.",
				Optional:    true,
				Computed:    true,
			},
			"dstaddr": {
				Type:        schema.TypeList,
				Description: "IPv4 destination address and address group names.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dstaddr6": {
				Type:        schema.TypeList,
				Description: "IPv6 destination address and address group names.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dstintf": {
				Type:        schema.TypeList,
				Description: "One or more outgoing (egress) interfaces.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"groups": {
				Type:        schema.TypeList,
				Description: "Apply this traffic shaping policy to user groups that have authenticated with the FortiGate.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "Shaping policy ID (0 - 4294967295).",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"internet_service": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.",
				Optional:    true,
				Computed:    true,
			},
			"internet_service_custom": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Custom Internet Service name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_custom_group": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service group name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Custom Internet Service group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_group": {
				Type:        schema.TypeList,
				Description: "Internet Service group name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Internet Service group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_id": {
				Type:        schema.TypeList,
				Description: "Internet Service ID.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Internet Service ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_name": {
				Type:        schema.TypeList,
				Description: "Internet Service ID.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Internet Service name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used.",
				Optional:    true,
				Computed:    true,
			},
			"internet_service_src_custom": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service source name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Custom Internet Service name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_custom_group": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service source group name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Custom Internet Service group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_group": {
				Type:        schema.TypeList,
				Description: "Internet Service source group name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Internet Service group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_id": {
				Type:        schema.TypeList,
				Description: "Internet Service source ID.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Internet Service ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_name": {
				Type:        schema.TypeList,
				Description: "Internet Service source name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Internet Service name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ip_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"4", "6"}, false),

				Description: "Apply this traffic shaping policy to IPv4 or IPv6 traffic.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Shaping policy name.",
				Optional:    true,
				Computed:    true,
			},
			"per_ip_shaper": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Per-IP traffic shaper to apply with this policy.",
				Optional:    true,
				Computed:    true,
			},
			"schedule": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Schedule name.",
				Optional:    true,
				Computed:    true,
			},
			"service": {
				Type:        schema.TypeList,
				Description: "Service and service group names.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Service name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"srcaddr": {
				Type:        schema.TypeList,
				Description: "IPv4 source address and address group names.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"srcaddr6": {
				Type:        schema.TypeList,
				Description: "IPv6 source address and address group names.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"srcintf": {
				Type:        schema.TypeList,
				Description: "One or more incoming (ingress) interfaces.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable this traffic shaping policy.",
				Optional:    true,
				Computed:    true,
			},
			"tos": {
				Type: schema.TypeString,

				Description: "ToS (Type of Service) value used for comparison.",
				Optional:    true,
				Computed:    true,
			},
			"tos_mask": {
				Type: schema.TypeString,

				Description: "Non-zero bit positions are used for comparison while zero bit positions are ignored.",
				Optional:    true,
				Computed:    true,
			},
			"tos_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable negated TOS match.",
				Optional:    true,
				Computed:    true,
			},
			"traffic_shaper": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Traffic shaper to apply to traffic forwarded by the firewall policy.",
				Optional:    true,
				Computed:    true,
			},
			"traffic_shaper_reverse": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Traffic shaper to apply to response traffic received by the firewall policy.",
				Optional:    true,
				Computed:    true,
			},
			"url_category": {
				Type:        schema.TypeList,
				Description: "IDs of one or more FortiGuard Web Filtering categories that this shaper applies traffic shaping to.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "URL category ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"users": {
				Type:        schema.TypeList,
				Description: "Apply this traffic shaping policy to individual users that have authenticated with the FortiGate.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "User name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceFirewallShapingPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "id"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating FirewallShapingPolicy resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallShapingPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallShapingPolicy(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallShapingPolicy")
	}

	return resourceFirewallShapingPolicyRead(ctx, d, meta)
}

func resourceFirewallShapingPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallShapingPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallShapingPolicy(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallShapingPolicy resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallShapingPolicy")
	}

	return resourceFirewallShapingPolicyRead(ctx, d, meta)
}

func resourceFirewallShapingPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallShapingPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallShapingPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallShapingPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallShapingPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallShapingPolicy resource: %v", err)
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

	diags := refreshObjectFirewallShapingPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallShapingPolicyAppCategory(v *[]models.FirewallShapingPolicyAppCategory, sort bool) interface{} {
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

func flattenFirewallShapingPolicyAppGroup(v *[]models.FirewallShapingPolicyAppGroup, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallShapingPolicyApplication(v *[]models.FirewallShapingPolicyApplication, sort bool) interface{} {
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

func flattenFirewallShapingPolicyDstaddr(v *[]models.FirewallShapingPolicyDstaddr, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallShapingPolicyDstaddr6(v *[]models.FirewallShapingPolicyDstaddr6, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallShapingPolicyDstintf(v *[]models.FirewallShapingPolicyDstintf, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallShapingPolicyGroups(v *[]models.FirewallShapingPolicyGroups, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallShapingPolicyInternetServiceCustom(v *[]models.FirewallShapingPolicyInternetServiceCustom, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallShapingPolicyInternetServiceCustomGroup(v *[]models.FirewallShapingPolicyInternetServiceCustomGroup, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallShapingPolicyInternetServiceGroup(v *[]models.FirewallShapingPolicyInternetServiceGroup, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallShapingPolicyInternetServiceId(v *[]models.FirewallShapingPolicyInternetServiceId, sort bool) interface{} {
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

func flattenFirewallShapingPolicyInternetServiceName(v *[]models.FirewallShapingPolicyInternetServiceName, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallShapingPolicyInternetServiceSrcCustom(v *[]models.FirewallShapingPolicyInternetServiceSrcCustom, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallShapingPolicyInternetServiceSrcCustomGroup(v *[]models.FirewallShapingPolicyInternetServiceSrcCustomGroup, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallShapingPolicyInternetServiceSrcGroup(v *[]models.FirewallShapingPolicyInternetServiceSrcGroup, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallShapingPolicyInternetServiceSrcId(v *[]models.FirewallShapingPolicyInternetServiceSrcId, sort bool) interface{} {
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

func flattenFirewallShapingPolicyInternetServiceSrcName(v *[]models.FirewallShapingPolicyInternetServiceSrcName, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallShapingPolicyService(v *[]models.FirewallShapingPolicyService, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallShapingPolicySrcaddr(v *[]models.FirewallShapingPolicySrcaddr, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallShapingPolicySrcaddr6(v *[]models.FirewallShapingPolicySrcaddr6, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallShapingPolicySrcintf(v *[]models.FirewallShapingPolicySrcintf, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallShapingPolicyUrlCategory(v *[]models.FirewallShapingPolicyUrlCategory, sort bool) interface{} {
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

func flattenFirewallShapingPolicyUsers(v *[]models.FirewallShapingPolicyUsers, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func refreshObjectFirewallShapingPolicy(d *schema.ResourceData, o *models.FirewallShapingPolicy, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AppCategory != nil {
		if err = d.Set("app_category", flattenFirewallShapingPolicyAppCategory(o.AppCategory, sort)); err != nil {
			return diag.Errorf("error reading app_category: %v", err)
		}
	}

	if o.AppGroup != nil {
		if err = d.Set("app_group", flattenFirewallShapingPolicyAppGroup(o.AppGroup, sort)); err != nil {
			return diag.Errorf("error reading app_group: %v", err)
		}
	}

	if o.Application != nil {
		if err = d.Set("application", flattenFirewallShapingPolicyApplication(o.Application, sort)); err != nil {
			return diag.Errorf("error reading application: %v", err)
		}
	}

	if o.ClassId != nil {
		v := *o.ClassId

		if err = d.Set("class_id", v); err != nil {
			return diag.Errorf("error reading class_id: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.DiffservForward != nil {
		v := *o.DiffservForward

		if err = d.Set("diffserv_forward", v); err != nil {
			return diag.Errorf("error reading diffserv_forward: %v", err)
		}
	}

	if o.DiffservReverse != nil {
		v := *o.DiffservReverse

		if err = d.Set("diffserv_reverse", v); err != nil {
			return diag.Errorf("error reading diffserv_reverse: %v", err)
		}
	}

	if o.DiffservcodeForward != nil {
		v := *o.DiffservcodeForward

		if err = d.Set("diffservcode_forward", v); err != nil {
			return diag.Errorf("error reading diffservcode_forward: %v", err)
		}
	}

	if o.DiffservcodeRev != nil {
		v := *o.DiffservcodeRev

		if err = d.Set("diffservcode_rev", v); err != nil {
			return diag.Errorf("error reading diffservcode_rev: %v", err)
		}
	}

	if o.Dstaddr != nil {
		if err = d.Set("dstaddr", flattenFirewallShapingPolicyDstaddr(o.Dstaddr, sort)); err != nil {
			return diag.Errorf("error reading dstaddr: %v", err)
		}
	}

	if o.Dstaddr6 != nil {
		if err = d.Set("dstaddr6", flattenFirewallShapingPolicyDstaddr6(o.Dstaddr6, sort)); err != nil {
			return diag.Errorf("error reading dstaddr6: %v", err)
		}
	}

	if o.Dstintf != nil {
		if err = d.Set("dstintf", flattenFirewallShapingPolicyDstintf(o.Dstintf, sort)); err != nil {
			return diag.Errorf("error reading dstintf: %v", err)
		}
	}

	if o.Groups != nil {
		if err = d.Set("groups", flattenFirewallShapingPolicyGroups(o.Groups, sort)); err != nil {
			return diag.Errorf("error reading groups: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.InternetService != nil {
		v := *o.InternetService

		if err = d.Set("internet_service", v); err != nil {
			return diag.Errorf("error reading internet_service: %v", err)
		}
	}

	if o.InternetServiceCustom != nil {
		if err = d.Set("internet_service_custom", flattenFirewallShapingPolicyInternetServiceCustom(o.InternetServiceCustom, sort)); err != nil {
			return diag.Errorf("error reading internet_service_custom: %v", err)
		}
	}

	if o.InternetServiceCustomGroup != nil {
		if err = d.Set("internet_service_custom_group", flattenFirewallShapingPolicyInternetServiceCustomGroup(o.InternetServiceCustomGroup, sort)); err != nil {
			return diag.Errorf("error reading internet_service_custom_group: %v", err)
		}
	}

	if o.InternetServiceGroup != nil {
		if err = d.Set("internet_service_group", flattenFirewallShapingPolicyInternetServiceGroup(o.InternetServiceGroup, sort)); err != nil {
			return diag.Errorf("error reading internet_service_group: %v", err)
		}
	}

	if o.InternetServiceId != nil {
		if err = d.Set("internet_service_id", flattenFirewallShapingPolicyInternetServiceId(o.InternetServiceId, sort)); err != nil {
			return diag.Errorf("error reading internet_service_id: %v", err)
		}
	}

	if o.InternetServiceName != nil {
		if err = d.Set("internet_service_name", flattenFirewallShapingPolicyInternetServiceName(o.InternetServiceName, sort)); err != nil {
			return diag.Errorf("error reading internet_service_name: %v", err)
		}
	}

	if o.InternetServiceSrc != nil {
		v := *o.InternetServiceSrc

		if err = d.Set("internet_service_src", v); err != nil {
			return diag.Errorf("error reading internet_service_src: %v", err)
		}
	}

	if o.InternetServiceSrcCustom != nil {
		if err = d.Set("internet_service_src_custom", flattenFirewallShapingPolicyInternetServiceSrcCustom(o.InternetServiceSrcCustom, sort)); err != nil {
			return diag.Errorf("error reading internet_service_src_custom: %v", err)
		}
	}

	if o.InternetServiceSrcCustomGroup != nil {
		if err = d.Set("internet_service_src_custom_group", flattenFirewallShapingPolicyInternetServiceSrcCustomGroup(o.InternetServiceSrcCustomGroup, sort)); err != nil {
			return diag.Errorf("error reading internet_service_src_custom_group: %v", err)
		}
	}

	if o.InternetServiceSrcGroup != nil {
		if err = d.Set("internet_service_src_group", flattenFirewallShapingPolicyInternetServiceSrcGroup(o.InternetServiceSrcGroup, sort)); err != nil {
			return diag.Errorf("error reading internet_service_src_group: %v", err)
		}
	}

	if o.InternetServiceSrcId != nil {
		if err = d.Set("internet_service_src_id", flattenFirewallShapingPolicyInternetServiceSrcId(o.InternetServiceSrcId, sort)); err != nil {
			return diag.Errorf("error reading internet_service_src_id: %v", err)
		}
	}

	if o.InternetServiceSrcName != nil {
		if err = d.Set("internet_service_src_name", flattenFirewallShapingPolicyInternetServiceSrcName(o.InternetServiceSrcName, sort)); err != nil {
			return diag.Errorf("error reading internet_service_src_name: %v", err)
		}
	}

	if o.IpVersion != nil {
		v := *o.IpVersion

		if err = d.Set("ip_version", v); err != nil {
			return diag.Errorf("error reading ip_version: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.PerIpShaper != nil {
		v := *o.PerIpShaper

		if err = d.Set("per_ip_shaper", v); err != nil {
			return diag.Errorf("error reading per_ip_shaper: %v", err)
		}
	}

	if o.Schedule != nil {
		v := *o.Schedule

		if err = d.Set("schedule", v); err != nil {
			return diag.Errorf("error reading schedule: %v", err)
		}
	}

	if o.Service != nil {
		if err = d.Set("service", flattenFirewallShapingPolicyService(o.Service, sort)); err != nil {
			return diag.Errorf("error reading service: %v", err)
		}
	}

	if o.Srcaddr != nil {
		if err = d.Set("srcaddr", flattenFirewallShapingPolicySrcaddr(o.Srcaddr, sort)); err != nil {
			return diag.Errorf("error reading srcaddr: %v", err)
		}
	}

	if o.Srcaddr6 != nil {
		if err = d.Set("srcaddr6", flattenFirewallShapingPolicySrcaddr6(o.Srcaddr6, sort)); err != nil {
			return diag.Errorf("error reading srcaddr6: %v", err)
		}
	}

	if o.Srcintf != nil {
		if err = d.Set("srcintf", flattenFirewallShapingPolicySrcintf(o.Srcintf, sort)); err != nil {
			return diag.Errorf("error reading srcintf: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Tos != nil {
		v := *o.Tos

		if err = d.Set("tos", v); err != nil {
			return diag.Errorf("error reading tos: %v", err)
		}
	}

	if o.TosMask != nil {
		v := *o.TosMask

		if err = d.Set("tos_mask", v); err != nil {
			return diag.Errorf("error reading tos_mask: %v", err)
		}
	}

	if o.TosNegate != nil {
		v := *o.TosNegate

		if err = d.Set("tos_negate", v); err != nil {
			return diag.Errorf("error reading tos_negate: %v", err)
		}
	}

	if o.TrafficShaper != nil {
		v := *o.TrafficShaper

		if err = d.Set("traffic_shaper", v); err != nil {
			return diag.Errorf("error reading traffic_shaper: %v", err)
		}
	}

	if o.TrafficShaperReverse != nil {
		v := *o.TrafficShaperReverse

		if err = d.Set("traffic_shaper_reverse", v); err != nil {
			return diag.Errorf("error reading traffic_shaper_reverse: %v", err)
		}
	}

	if o.UrlCategory != nil {
		if err = d.Set("url_category", flattenFirewallShapingPolicyUrlCategory(o.UrlCategory, sort)); err != nil {
			return diag.Errorf("error reading url_category: %v", err)
		}
	}

	if o.Users != nil {
		if err = d.Set("users", flattenFirewallShapingPolicyUsers(o.Users, sort)); err != nil {
			return diag.Errorf("error reading users: %v", err)
		}
	}

	return nil
}

func expandFirewallShapingPolicyAppCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicyAppCategory, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicyAppCategory

	for i := range l {
		tmp := models.FirewallShapingPolicyAppCategory{}
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

func expandFirewallShapingPolicyAppGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicyAppGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicyAppGroup

	for i := range l {
		tmp := models.FirewallShapingPolicyAppGroup{}
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

func expandFirewallShapingPolicyApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicyApplication, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicyApplication

	for i := range l {
		tmp := models.FirewallShapingPolicyApplication{}
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

func expandFirewallShapingPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicyDstaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicyDstaddr

	for i := range l {
		tmp := models.FirewallShapingPolicyDstaddr{}
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

func expandFirewallShapingPolicyDstaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicyDstaddr6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicyDstaddr6

	for i := range l {
		tmp := models.FirewallShapingPolicyDstaddr6{}
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

func expandFirewallShapingPolicyDstintf(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicyDstintf, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicyDstintf

	for i := range l {
		tmp := models.FirewallShapingPolicyDstintf{}
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

func expandFirewallShapingPolicyGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicyGroups, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicyGroups

	for i := range l {
		tmp := models.FirewallShapingPolicyGroups{}
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

func expandFirewallShapingPolicyInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicyInternetServiceCustom, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicyInternetServiceCustom

	for i := range l {
		tmp := models.FirewallShapingPolicyInternetServiceCustom{}
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

func expandFirewallShapingPolicyInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicyInternetServiceCustomGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicyInternetServiceCustomGroup

	for i := range l {
		tmp := models.FirewallShapingPolicyInternetServiceCustomGroup{}
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

func expandFirewallShapingPolicyInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicyInternetServiceGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicyInternetServiceGroup

	for i := range l {
		tmp := models.FirewallShapingPolicyInternetServiceGroup{}
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

func expandFirewallShapingPolicyInternetServiceId(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicyInternetServiceId, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicyInternetServiceId

	for i := range l {
		tmp := models.FirewallShapingPolicyInternetServiceId{}
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

func expandFirewallShapingPolicyInternetServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicyInternetServiceName, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicyInternetServiceName

	for i := range l {
		tmp := models.FirewallShapingPolicyInternetServiceName{}
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

func expandFirewallShapingPolicyInternetServiceSrcCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicyInternetServiceSrcCustom, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicyInternetServiceSrcCustom

	for i := range l {
		tmp := models.FirewallShapingPolicyInternetServiceSrcCustom{}
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

func expandFirewallShapingPolicyInternetServiceSrcCustomGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicyInternetServiceSrcCustomGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicyInternetServiceSrcCustomGroup

	for i := range l {
		tmp := models.FirewallShapingPolicyInternetServiceSrcCustomGroup{}
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

func expandFirewallShapingPolicyInternetServiceSrcGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicyInternetServiceSrcGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicyInternetServiceSrcGroup

	for i := range l {
		tmp := models.FirewallShapingPolicyInternetServiceSrcGroup{}
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

func expandFirewallShapingPolicyInternetServiceSrcId(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicyInternetServiceSrcId, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicyInternetServiceSrcId

	for i := range l {
		tmp := models.FirewallShapingPolicyInternetServiceSrcId{}
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

func expandFirewallShapingPolicyInternetServiceSrcName(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicyInternetServiceSrcName, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicyInternetServiceSrcName

	for i := range l {
		tmp := models.FirewallShapingPolicyInternetServiceSrcName{}
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

func expandFirewallShapingPolicyService(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicyService, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicyService

	for i := range l {
		tmp := models.FirewallShapingPolicyService{}
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

func expandFirewallShapingPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicySrcaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicySrcaddr

	for i := range l {
		tmp := models.FirewallShapingPolicySrcaddr{}
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

func expandFirewallShapingPolicySrcaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicySrcaddr6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicySrcaddr6

	for i := range l {
		tmp := models.FirewallShapingPolicySrcaddr6{}
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

func expandFirewallShapingPolicySrcintf(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicySrcintf, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicySrcintf

	for i := range l {
		tmp := models.FirewallShapingPolicySrcintf{}
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

func expandFirewallShapingPolicyUrlCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicyUrlCategory, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicyUrlCategory

	for i := range l {
		tmp := models.FirewallShapingPolicyUrlCategory{}
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

func expandFirewallShapingPolicyUsers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallShapingPolicyUsers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallShapingPolicyUsers

	for i := range l {
		tmp := models.FirewallShapingPolicyUsers{}
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

func getObjectFirewallShapingPolicy(d *schema.ResourceData, sv string) (*models.FirewallShapingPolicy, diag.Diagnostics) {
	obj := models.FirewallShapingPolicy{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("app_category"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("app_category", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicyAppCategory(d, v, "app_category", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.AppCategory = t
		}
	} else if d.HasChange("app_category") {
		old, new := d.GetChange("app_category")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.AppCategory = &[]models.FirewallShapingPolicyAppCategory{}
		}
	}
	if v, ok := d.GetOk("app_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("app_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicyAppGroup(d, v, "app_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.AppGroup = t
		}
	} else if d.HasChange("app_group") {
		old, new := d.GetChange("app_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.AppGroup = &[]models.FirewallShapingPolicyAppGroup{}
		}
	}
	if v, ok := d.GetOk("application"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("application", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicyApplication(d, v, "application", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Application = t
		}
	} else if d.HasChange("application") {
		old, new := d.GetChange("application")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Application = &[]models.FirewallShapingPolicyApplication{}
		}
	}
	if v1, ok := d.GetOk("class_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("class_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ClassId = &tmp
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
	if v1, ok := d.GetOk("diffserv_forward"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("diffserv_forward", sv)
				diags = append(diags, e)
			}
			obj.DiffservForward = &v2
		}
	}
	if v1, ok := d.GetOk("diffserv_reverse"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("diffserv_reverse", sv)
				diags = append(diags, e)
			}
			obj.DiffservReverse = &v2
		}
	}
	if v1, ok := d.GetOk("diffservcode_forward"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("diffservcode_forward", sv)
				diags = append(diags, e)
			}
			obj.DiffservcodeForward = &v2
		}
	}
	if v1, ok := d.GetOk("diffservcode_rev"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("diffservcode_rev", sv)
				diags = append(diags, e)
			}
			obj.DiffservcodeRev = &v2
		}
	}
	if v, ok := d.GetOk("dstaddr"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dstaddr", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicyDstaddr(d, v, "dstaddr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstaddr = t
		}
	} else if d.HasChange("dstaddr") {
		old, new := d.GetChange("dstaddr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstaddr = &[]models.FirewallShapingPolicyDstaddr{}
		}
	}
	if v, ok := d.GetOk("dstaddr6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dstaddr6", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicyDstaddr6(d, v, "dstaddr6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstaddr6 = t
		}
	} else if d.HasChange("dstaddr6") {
		old, new := d.GetChange("dstaddr6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstaddr6 = &[]models.FirewallShapingPolicyDstaddr6{}
		}
	}
	if v, ok := d.GetOk("dstintf"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dstintf", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicyDstintf(d, v, "dstintf", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstintf = t
		}
	} else if d.HasChange("dstintf") {
		old, new := d.GetChange("dstintf")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstintf = &[]models.FirewallShapingPolicyDstintf{}
		}
	}
	if v, ok := d.GetOk("groups"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("groups", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicyGroups(d, v, "groups", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Groups = t
		}
	} else if d.HasChange("groups") {
		old, new := d.GetChange("groups")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Groups = &[]models.FirewallShapingPolicyGroups{}
		}
	}
	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Id = &tmp
		}
	}
	if v1, ok := d.GetOk("internet_service"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("internet_service", sv)
				diags = append(diags, e)
			}
			obj.InternetService = &v2
		}
	}
	if v, ok := d.GetOk("internet_service_custom"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_custom", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicyInternetServiceCustom(d, v, "internet_service_custom", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceCustom = t
		}
	} else if d.HasChange("internet_service_custom") {
		old, new := d.GetChange("internet_service_custom")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceCustom = &[]models.FirewallShapingPolicyInternetServiceCustom{}
		}
	}
	if v, ok := d.GetOk("internet_service_custom_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_custom_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicyInternetServiceCustomGroup(d, v, "internet_service_custom_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceCustomGroup = t
		}
	} else if d.HasChange("internet_service_custom_group") {
		old, new := d.GetChange("internet_service_custom_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceCustomGroup = &[]models.FirewallShapingPolicyInternetServiceCustomGroup{}
		}
	}
	if v, ok := d.GetOk("internet_service_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicyInternetServiceGroup(d, v, "internet_service_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceGroup = t
		}
	} else if d.HasChange("internet_service_group") {
		old, new := d.GetChange("internet_service_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceGroup = &[]models.FirewallShapingPolicyInternetServiceGroup{}
		}
	}
	if v, ok := d.GetOk("internet_service_id"); ok {
		if !utils.CheckVer(sv, "", "v6.4.0") {
			e := utils.AttributeVersionWarning("internet_service_id", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicyInternetServiceId(d, v, "internet_service_id", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceId = t
		}
	} else if d.HasChange("internet_service_id") {
		old, new := d.GetChange("internet_service_id")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceId = &[]models.FirewallShapingPolicyInternetServiceId{}
		}
	}
	if v, ok := d.GetOk("internet_service_name"); ok {
		if !utils.CheckVer(sv, "v6.4.0", "") {
			e := utils.AttributeVersionWarning("internet_service_name", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicyInternetServiceName(d, v, "internet_service_name", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceName = t
		}
	} else if d.HasChange("internet_service_name") {
		old, new := d.GetChange("internet_service_name")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceName = &[]models.FirewallShapingPolicyInternetServiceName{}
		}
	}
	if v1, ok := d.GetOk("internet_service_src"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("internet_service_src", sv)
				diags = append(diags, e)
			}
			obj.InternetServiceSrc = &v2
		}
	}
	if v, ok := d.GetOk("internet_service_src_custom"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_src_custom", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicyInternetServiceSrcCustom(d, v, "internet_service_src_custom", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceSrcCustom = t
		}
	} else if d.HasChange("internet_service_src_custom") {
		old, new := d.GetChange("internet_service_src_custom")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceSrcCustom = &[]models.FirewallShapingPolicyInternetServiceSrcCustom{}
		}
	}
	if v, ok := d.GetOk("internet_service_src_custom_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_src_custom_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicyInternetServiceSrcCustomGroup(d, v, "internet_service_src_custom_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceSrcCustomGroup = t
		}
	} else if d.HasChange("internet_service_src_custom_group") {
		old, new := d.GetChange("internet_service_src_custom_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceSrcCustomGroup = &[]models.FirewallShapingPolicyInternetServiceSrcCustomGroup{}
		}
	}
	if v, ok := d.GetOk("internet_service_src_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_src_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicyInternetServiceSrcGroup(d, v, "internet_service_src_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceSrcGroup = t
		}
	} else if d.HasChange("internet_service_src_group") {
		old, new := d.GetChange("internet_service_src_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceSrcGroup = &[]models.FirewallShapingPolicyInternetServiceSrcGroup{}
		}
	}
	if v, ok := d.GetOk("internet_service_src_id"); ok {
		if !utils.CheckVer(sv, "", "v6.4.0") {
			e := utils.AttributeVersionWarning("internet_service_src_id", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicyInternetServiceSrcId(d, v, "internet_service_src_id", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceSrcId = t
		}
	} else if d.HasChange("internet_service_src_id") {
		old, new := d.GetChange("internet_service_src_id")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceSrcId = &[]models.FirewallShapingPolicyInternetServiceSrcId{}
		}
	}
	if v, ok := d.GetOk("internet_service_src_name"); ok {
		if !utils.CheckVer(sv, "v6.4.0", "") {
			e := utils.AttributeVersionWarning("internet_service_src_name", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicyInternetServiceSrcName(d, v, "internet_service_src_name", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceSrcName = t
		}
	} else if d.HasChange("internet_service_src_name") {
		old, new := d.GetChange("internet_service_src_name")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceSrcName = &[]models.FirewallShapingPolicyInternetServiceSrcName{}
		}
	}
	if v1, ok := d.GetOk("ip_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip_version", sv)
				diags = append(diags, e)
			}
			obj.IpVersion = &v2
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
	if v1, ok := d.GetOk("per_ip_shaper"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("per_ip_shaper", sv)
				diags = append(diags, e)
			}
			obj.PerIpShaper = &v2
		}
	}
	if v1, ok := d.GetOk("schedule"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("schedule", sv)
				diags = append(diags, e)
			}
			obj.Schedule = &v2
		}
	}
	if v, ok := d.GetOk("service"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("service", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicyService(d, v, "service", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Service = t
		}
	} else if d.HasChange("service") {
		old, new := d.GetChange("service")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Service = &[]models.FirewallShapingPolicyService{}
		}
	}
	if v, ok := d.GetOk("srcaddr"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("srcaddr", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicySrcaddr(d, v, "srcaddr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcaddr = t
		}
	} else if d.HasChange("srcaddr") {
		old, new := d.GetChange("srcaddr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcaddr = &[]models.FirewallShapingPolicySrcaddr{}
		}
	}
	if v, ok := d.GetOk("srcaddr6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("srcaddr6", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicySrcaddr6(d, v, "srcaddr6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcaddr6 = t
		}
	} else if d.HasChange("srcaddr6") {
		old, new := d.GetChange("srcaddr6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcaddr6 = &[]models.FirewallShapingPolicySrcaddr6{}
		}
	}
	if v, ok := d.GetOk("srcintf"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("srcintf", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicySrcintf(d, v, "srcintf", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcintf = t
		}
	} else if d.HasChange("srcintf") {
		old, new := d.GetChange("srcintf")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcintf = &[]models.FirewallShapingPolicySrcintf{}
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
	if v1, ok := d.GetOk("tos"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tos", sv)
				diags = append(diags, e)
			}
			obj.Tos = &v2
		}
	}
	if v1, ok := d.GetOk("tos_mask"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tos_mask", sv)
				diags = append(diags, e)
			}
			obj.TosMask = &v2
		}
	}
	if v1, ok := d.GetOk("tos_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tos_negate", sv)
				diags = append(diags, e)
			}
			obj.TosNegate = &v2
		}
	}
	if v1, ok := d.GetOk("traffic_shaper"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("traffic_shaper", sv)
				diags = append(diags, e)
			}
			obj.TrafficShaper = &v2
		}
	}
	if v1, ok := d.GetOk("traffic_shaper_reverse"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("traffic_shaper_reverse", sv)
				diags = append(diags, e)
			}
			obj.TrafficShaperReverse = &v2
		}
	}
	if v, ok := d.GetOk("url_category"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("url_category", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicyUrlCategory(d, v, "url_category", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.UrlCategory = t
		}
	} else if d.HasChange("url_category") {
		old, new := d.GetChange("url_category")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.UrlCategory = &[]models.FirewallShapingPolicyUrlCategory{}
		}
	}
	if v, ok := d.GetOk("users"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("users", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallShapingPolicyUsers(d, v, "users", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Users = t
		}
	} else if d.HasChange("users") {
		old, new := d.GetChange("users")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Users = &[]models.FirewallShapingPolicyUsers{}
		}
	}
	return &obj, diags
}
