// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemSdwanService() *schema.Resource {
	return &schema.Resource{
		Description: "Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN.",

		ReadContext: dataSourceSystemSdwanServiceRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"addr_mode": {
				Type:        schema.TypeString,
				Description: "Address mode (IPv4 or IPv6).",
				Computed:    true,
			},
			"agent_exclusive": {
				Type:        schema.TypeString,
				Description: "Set/unset the service as agent use exclusively.",
				Computed:    true,
			},
			"bandwidth_weight": {
				Type:        schema.TypeInt,
				Description: "Coefficient of reciprocal of available bidirectional bandwidth in the formula of custom-profile-1.",
				Computed:    true,
			},
			"default": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of SD-WAN as default service.",
				Computed:    true,
			},
			"dscp_forward": {
				Type:        schema.TypeString,
				Description: "Enable/disable forward traffic DSCP tag.",
				Computed:    true,
			},
			"dscp_forward_tag": {
				Type:        schema.TypeString,
				Description: "Forward traffic DSCP tag.",
				Computed:    true,
			},
			"dscp_reverse": {
				Type:        schema.TypeString,
				Description: "Enable/disable reverse traffic DSCP tag.",
				Computed:    true,
			},
			"dscp_reverse_tag": {
				Type:        schema.TypeString,
				Description: "Reverse traffic DSCP tag.",
				Computed:    true,
			},
			"dst": {
				Type:        schema.TypeList,
				Description: "Destination address name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address or address group name.",
							Computed:    true,
						},
					},
				},
			},
			"dst_negate": {
				Type:        schema.TypeString,
				Description: "Enable/disable negation of destination address match.",
				Computed:    true,
			},
			"dst6": {
				Type:        schema.TypeList,
				Description: "Destination address6 name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address6 or address6 group name.",
							Computed:    true,
						},
					},
				},
			},
			"end_port": {
				Type:        schema.TypeInt,
				Description: "End destination port number.",
				Computed:    true,
			},
			"gateway": {
				Type:        schema.TypeString,
				Description: "Enable/disable SD-WAN service gateway.",
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
			"hash_mode": {
				Type:        schema.TypeString,
				Description: "Hash algorithm for selected priority members for load balance mode.",
				Computed:    true,
			},
			"health_check": {
				Type:        schema.TypeList,
				Description: "Health check list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Health check name.",
							Computed:    true,
						},
					},
				},
			},
			"hold_down_time": {
				Type:        schema.TypeInt,
				Description: "Waiting period in seconds when switching from the back-up member to the primary member (0 - 10000000, default = 0).",
				Computed:    true,
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "SD-WAN rule ID (1 - 4000).",
				Computed:    true,
			},
			"input_device": {
				Type:        schema.TypeList,
				Description: "Source interface name.",
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
			"input_device_negate": {
				Type:        schema.TypeString,
				Description: "Enable/disable negation of input device match.",
				Computed:    true,
			},
			"input_zone": {
				Type:        schema.TypeList,
				Description: "Source input-zone name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Zone.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of Internet service for application-based load balancing.",
				Computed:    true,
			},
			"internet_service_app_ctrl": {
				Type:        schema.TypeList,
				Description: "Application control based Internet Service ID list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Application control based Internet Service ID.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_app_ctrl_category": {
				Type:        schema.TypeList,
				Description: "IDs of one or more application control categories.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Application control category ID.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_app_ctrl_group": {
				Type:        schema.TypeList,
				Description: "Application control based Internet Service group list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Application control based Internet Service group name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_custom": {
				Type:        schema.TypeList,
				Description: "Custom Internet service name list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Custom Internet service name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_custom_group": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service group list.",
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
				Description: "Internet Service group list.",
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
			"internet_service_name": {
				Type:        schema.TypeList,
				Description: "Internet service name list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Internet service name.",
							Computed:    true,
						},
					},
				},
			},
			"jitter_weight": {
				Type:        schema.TypeInt,
				Description: "Coefficient of jitter in the formula of custom-profile-1.",
				Computed:    true,
			},
			"latency_weight": {
				Type:        schema.TypeInt,
				Description: "Coefficient of latency in the formula of custom-profile-1.",
				Computed:    true,
			},
			"link_cost_factor": {
				Type:        schema.TypeString,
				Description: "Link cost factor.",
				Computed:    true,
			},
			"link_cost_threshold": {
				Type:        schema.TypeInt,
				Description: "Percentage threshold change of link cost values that will result in policy route regeneration (0 - 10000000, default = 10).",
				Computed:    true,
			},
			"minimum_sla_meet_members": {
				Type:        schema.TypeInt,
				Description: "Minimum number of members which meet SLA.",
				Computed:    true,
			},
			"mode": {
				Type:        schema.TypeString,
				Description: "Control how the SD-WAN rule sets the priority of interfaces in the SD-WAN.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "SD-WAN rule name.",
				Computed:    true,
			},
			"packet_loss_weight": {
				Type:        schema.TypeInt,
				Description: "Coefficient of packet-loss in the formula of custom-profile-1.",
				Computed:    true,
			},
			"passive_measurement": {
				Type:        schema.TypeString,
				Description: "Enable/disable passive measurement based on the service criteria.",
				Computed:    true,
			},
			"priority_members": {
				Type:        schema.TypeList,
				Description: "Member sequence number list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seq_num": {
							Type:        schema.TypeInt,
							Description: "Member sequence number.",
							Computed:    true,
						},
					},
				},
			},
			"priority_zone": {
				Type:        schema.TypeList,
				Description: "Priority zone name list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Priority zone name.",
							Computed:    true,
						},
					},
				},
			},
			"protocol": {
				Type:        schema.TypeInt,
				Description: "Protocol number.",
				Computed:    true,
			},
			"quality_link": {
				Type:        schema.TypeInt,
				Description: "Quality grade.",
				Computed:    true,
			},
			"role": {
				Type:        schema.TypeString,
				Description: "Service role to work with neighbor.",
				Computed:    true,
			},
			"route_tag": {
				Type:        schema.TypeInt,
				Description: "IPv4 route map route-tag.",
				Computed:    true,
			},
			"sla": {
				Type:        schema.TypeList,
				Description: "Service level agreement (SLA).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"health_check": {
							Type:        schema.TypeString,
							Description: "SD-WAN health-check.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "SLA ID.",
							Computed:    true,
						},
					},
				},
			},
			"sla_compare_method": {
				Type:        schema.TypeString,
				Description: "Method to compare SLA value for SLA mode.",
				Computed:    true,
			},
			"src": {
				Type:        schema.TypeList,
				Description: "Source address name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address or address group name.",
							Computed:    true,
						},
					},
				},
			},
			"src_negate": {
				Type:        schema.TypeString,
				Description: "Enable/disable negation of source address match.",
				Computed:    true,
			},
			"src6": {
				Type:        schema.TypeList,
				Description: "Source address6 name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address6 or address6 group name.",
							Computed:    true,
						},
					},
				},
			},
			"standalone_action": {
				Type:        schema.TypeString,
				Description: "Enable/disable service when selected neighbor role is standalone while service role is not standalone.",
				Computed:    true,
			},
			"start_port": {
				Type:        schema.TypeInt,
				Description: "Start destination port number.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable SD-WAN service.",
				Computed:    true,
			},
			"tie_break": {
				Type:        schema.TypeString,
				Description: "Method of selecting member if more than one meets the SLA.",
				Computed:    true,
			},
			"tos": {
				Type:        schema.TypeString,
				Description: "Type of service bit pattern.",
				Computed:    true,
			},
			"tos_mask": {
				Type:        schema.TypeString,
				Description: "Type of service evaluated bits.",
				Computed:    true,
			},
			"use_shortcut_sla": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of ADVPN shortcut for quality comparison.",
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

func dataSourceSystemSdwanServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemSdwanService(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSdwanService dataSource: %v", err)
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

	diags := refreshObjectSystemSdwanService(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
