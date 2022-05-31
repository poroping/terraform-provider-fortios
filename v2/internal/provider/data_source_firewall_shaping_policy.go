// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure shaping policies.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallShapingPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure shaping policies.",

		ReadContext: dataSourceFirewallShapingPolicyRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"app_category": {
				Type:        schema.TypeList,
				Description: "IDs of one or more application categories that this shaper applies application control traffic shaping to.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Category IDs.",
							Computed:    true,
						},
					},
				},
			},
			"app_group": {
				Type:        schema.TypeList,
				Description: "One or more application group names.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Application group name.",
							Computed:    true,
						},
					},
				},
			},
			"application": {
				Type:        schema.TypeList,
				Description: "IDs of one or more applications that this shaper applies application control traffic shaping to.",
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
			"class_id": {
				Type:        schema.TypeInt,
				Description: "Traffic class ID.",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comments.",
				Computed:    true,
			},
			"diffserv_forward": {
				Type:        schema.TypeString,
				Description: "Enable to change packet's DiffServ values to the specified diffservcode-forward value.",
				Computed:    true,
			},
			"diffserv_reverse": {
				Type:        schema.TypeString,
				Description: "Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value.",
				Computed:    true,
			},
			"diffservcode_forward": {
				Type:        schema.TypeString,
				Description: "Change packet's DiffServ to this value.",
				Computed:    true,
			},
			"diffservcode_rev": {
				Type:        schema.TypeString,
				Description: "Change packet's reverse (reply) DiffServ to this value.",
				Computed:    true,
			},
			"dstaddr": {
				Type:        schema.TypeList,
				Description: "IPv4 destination address and address group names.",
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
				Description: "IPv6 destination address and address group names.",
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
			"dstintf": {
				Type:        schema.TypeList,
				Description: "One or more outgoing (egress) interfaces.",
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
			"groups": {
				Type:        schema.TypeList,
				Description: "Apply this traffic shaping policy to user groups that have authenticated with the FortiGate.",
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
				Description: "Shaping policy ID (0 - 4294967295).",
				Computed:    true,
			},
			"internet_service": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.",
				Computed:    true,
			},
			"internet_service_custom": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Custom Internet Service name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_custom_group": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service group name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Custom Internet Service group name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_group": {
				Type:        schema.TypeList,
				Description: "Internet Service group name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Internet Service group name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_id": {
				Type:        schema.TypeList,
				Description: "Internet Service ID.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Internet Service ID.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_name": {
				Type:        schema.TypeList,
				Description: "Internet Service ID.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Internet Service name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used.",
				Computed:    true,
			},
			"internet_service_src_custom": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service source name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Custom Internet Service name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_custom_group": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service source group name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Custom Internet Service group name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_group": {
				Type:        schema.TypeList,
				Description: "Internet Service source group name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Internet Service group name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_id": {
				Type:        schema.TypeList,
				Description: "Internet Service source ID.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Internet Service ID.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_name": {
				Type:        schema.TypeList,
				Description: "Internet Service source name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Internet Service name.",
							Computed:    true,
						},
					},
				},
			},
			"ip_version": {
				Type:        schema.TypeString,
				Description: "Apply this traffic shaping policy to IPv4 or IPv6 traffic.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Shaping policy name.",
				Computed:    true,
			},
			"per_ip_shaper": {
				Type:        schema.TypeString,
				Description: "Per-IP traffic shaper to apply with this policy.",
				Computed:    true,
			},
			"schedule": {
				Type:        schema.TypeString,
				Description: "Schedule name.",
				Computed:    true,
			},
			"service": {
				Type:        schema.TypeList,
				Description: "Service and service group names.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Service name.",
							Computed:    true,
						},
					},
				},
			},
			"srcaddr": {
				Type:        schema.TypeList,
				Description: "IPv4 source address and address group names.",
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
				Description: "IPv6 source address and address group names.",
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
				Description: "One or more incoming (ingress) interfaces.",
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
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable this traffic shaping policy.",
				Computed:    true,
			},
			"tos": {
				Type:        schema.TypeString,
				Description: "ToS (Type of Service) value used for comparison.",
				Computed:    true,
			},
			"tos_mask": {
				Type:        schema.TypeString,
				Description: "Non-zero bit positions are used for comparison while zero bit positions are ignored.",
				Computed:    true,
			},
			"tos_negate": {
				Type:        schema.TypeString,
				Description: "Enable negated TOS match.",
				Computed:    true,
			},
			"traffic_shaper": {
				Type:        schema.TypeString,
				Description: "Traffic shaper to apply to traffic forwarded by the firewall policy.",
				Computed:    true,
			},
			"traffic_shaper_reverse": {
				Type:        schema.TypeString,
				Description: "Traffic shaper to apply to response traffic received by the firewall policy.",
				Computed:    true,
			},
			"url_category": {
				Type:        schema.TypeList,
				Description: "IDs of one or more FortiGuard Web Filtering categories that this shaper applies traffic shaping to.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "URL category ID.",
							Computed:    true,
						},
					},
				},
			},
			"users": {
				Type:        schema.TypeList,
				Description: "Apply this traffic shaping policy to individual users that have authenticated with the FortiGate.",
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

func dataSourceFirewallShapingPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallShapingPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallShapingPolicy dataSource: %v", err)
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

	diags := refreshObjectFirewallShapingPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
