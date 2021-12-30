// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN.

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

func resourceSystemsdwanService() *schema.Resource {
	return &schema.Resource{
		Description: "Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN.",

		CreateContext: resourceSystemsdwanServiceCreate,
		ReadContext:   resourceSystemsdwanServiceRead,
		UpdateContext: resourceSystemsdwanServiceUpdate,
		DeleteContext: resourceSystemsdwanServiceDelete,

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
				Description:  "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:     true,
				Default:      false,
				RequiredWith: []string{"fosid"},
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"addr_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ipv4", "ipv6"}, false),

				Description: "Address mode (IPv4 or IPv6).",
				Optional:    true,
				Computed:    true,
			},
			"bandwidth_weight": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10000000),

				Description: "Coefficient of reciprocal of available bidirectional bandwidth in the formula of custom-profile-1.",
				Optional:    true,
				Computed:    true,
			},
			"default": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of SD-WAN as default service.",
				Optional:    true,
				Computed:    true,
			},
			"dscp_forward": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable forward traffic DSCP tag.",
				Optional:    true,
				Computed:    true,
			},
			"dscp_forward_tag": {
				Type: schema.TypeString,

				Description: "Forward traffic DSCP tag.",
				Optional:    true,
				Computed:    true,
			},
			"dscp_reverse": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable reverse traffic DSCP tag.",
				Optional:    true,
				Computed:    true,
			},
			"dscp_reverse_tag": {
				Type: schema.TypeString,

				Description: "Reverse traffic DSCP tag.",
				Optional:    true,
				Computed:    true,
			},
			"dst": {
				Type:        schema.TypeList,
				Description: "Destination address name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address or address group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dst_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable negation of destination address match.",
				Optional:    true,
				Computed:    true,
			},
			"dst6": {
				Type:        schema.TypeList,
				Description: "Destination address6 name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address6 or address6 group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"end_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "End destination port number.",
				Optional:    true,
				Computed:    true,
			},
			"gateway": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SD-WAN service gateway.",
				Optional:    true,
				Computed:    true,
			},
			"groups": {
				Type:        schema.TypeList,
				Description: "User groups.",
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
			"hash_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"round-robin", "source-ip-based", "source-dest-ip-based", "inbandwidth", "outbandwidth", "bibandwidth"}, false),

				Description: "Hash algorithm for selected priority members for load balance mode.",
				Optional:    true,
				Computed:    true,
			},
			"health_check": {
				Type:        schema.TypeList,
				Description: "Health check list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Health check name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"hold_down_time": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10000000),

				Description: "Waiting period in seconds when switching from the back-up member to the primary member (0 - 10000000, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"fosid": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 4000),

				Description: "SD-WAN rule ID (1 - 4000).",
				Optional:    true,
				Computed:    true,
			},
			"input_device": {
				Type:        schema.TypeList,
				Description: "Source interface name.",
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
			"input_device_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable negation of input device match.",
				Optional:    true,
				Computed:    true,
			},
			"internet_service": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of Internet service for application-based load balancing.",
				Optional:    true,
				Computed:    true,
			},
			"internet_service_app_ctrl": {
				Type:        schema.TypeList,
				Description: "Application control based Internet Service ID list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Application control based Internet Service ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_app_ctrl_group": {
				Type:        schema.TypeList,
				Description: "Application control based Internet Service group list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Application control based Internet Service group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_custom": {
				Type:        schema.TypeList,
				Description: "Custom Internet service name list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Custom Internet service name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_custom_group": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service group list.",
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
				Description: "Internet Service group list.",
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
			"internet_service_name": {
				Type:        schema.TypeList,
				Description: "Internet service name list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Internet service name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"jitter_weight": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10000000),

				Description: "Coefficient of jitter in the formula of custom-profile-1.",
				Optional:    true,
				Computed:    true,
			},
			"latency_weight": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10000000),

				Description: "Coefficient of latency in the formula of custom-profile-1.",
				Optional:    true,
				Computed:    true,
			},
			"link_cost_factor": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"latency", "jitter", "packet-loss", "inbandwidth", "outbandwidth", "bibandwidth", "custom-profile-1"}, false),

				Description: "Link cost factor.",
				Optional:    true,
				Computed:    true,
			},
			"link_cost_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10000000),

				Description: "Percentage threshold change of link cost values that will result in policy route regeneration (0 - 10000000, default = 10).",
				Optional:    true,
				Computed:    true,
			},
			"minimum_sla_meet_members": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Minimum number of members which meet SLA.",
				Optional:    true,
				Computed:    true,
			},
			"mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "manual", "priority", "sla", "load-balance"}, false),

				Description: "Control how the SD-WAN rule sets the priority of interfaces in the SD-WAN.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "SD-WAN rule name.",
				Optional:    true,
				Computed:    true,
			},
			"packet_loss_weight": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10000000),

				Description: "Coefficient of packet-loss in the formula of custom-profile-1.",
				Optional:    true,
				Computed:    true,
			},
			"passive_measurement": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable passive measurement based on the service criteria.",
				Optional:    true,
				Computed:    true,
			},
			"priority_members": {
				Type:        schema.TypeList,
				Description: "Member sequence number list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seq_num": {
							Type: schema.TypeInt,

							Description: "Member sequence number.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"priority_zone": {
				Type:        schema.TypeList,
				Description: "Priority zone name list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Priority zone name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"protocol": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Protocol number.",
				Optional:    true,
				Computed:    true,
			},
			"quality_link": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Quality grade.",
				Optional:    true,
				Computed:    true,
			},
			"role": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"standalone", "primary", "secondary"}, false),

				Description: "Service role to work with neighbor.",
				Optional:    true,
				Computed:    true,
			},
			"route_tag": {
				Type: schema.TypeInt,

				Description: "IPv4 route map route-tag.",
				Optional:    true,
				Computed:    true,
			},
			"sla": {
				Type:        schema.TypeList,
				Description: "Service level agreement (SLA).",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"health_check": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "SD-WAN health-check.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "SLA ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"sla_compare_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"order", "number"}, false),

				Description: "Method to compare SLA value for SLA mode.",
				Optional:    true,
				Computed:    true,
			},
			"src": {
				Type:        schema.TypeList,
				Description: "Source address name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address or address group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"src_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable negation of source address match.",
				Optional:    true,
				Computed:    true,
			},
			"src6": {
				Type:        schema.TypeList,
				Description: "Source address6 name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address6 or address6 group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"standalone_action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable service when selected neighbor role is standalone while service role is not standalone.",
				Optional:    true,
				Computed:    true,
			},
			"start_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Start destination port number.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SD-WAN service.",
				Optional:    true,
				Computed:    true,
			},
			"tie_break": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"zone", "cfg-order", "fib-best-match"}, false),

				Description: "Method of selecting member if more than one meets the SLA.",
				Optional:    true,
				Computed:    true,
			},
			"tos": {
				Type: schema.TypeString,

				Description: "Type of service bit pattern.",
				Optional:    true,
				Computed:    true,
			},
			"tos_mask": {
				Type: schema.TypeString,

				Description: "Type of service evaluated bits.",
				Optional:    true,
				Computed:    true,
			},
			"use_shortcut_sla": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of ADVPN shortcut for quality comparison.",
				Optional:    true,
				Computed:    true,
			},
			"users": {
				Type:        schema.TypeList,
				Description: "User name.",
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

func resourceSystemsdwanServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemsdwanService resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemsdwanService(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemsdwanService(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemsdwanService")
	}

	return resourceSystemsdwanServiceRead(ctx, d, meta)
}

func resourceSystemsdwanServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemsdwanService(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemsdwanService(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemsdwanService resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemsdwanService")
	}

	return resourceSystemsdwanServiceRead(ctx, d, meta)
}

func resourceSystemsdwanServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemsdwanService(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemsdwanService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemsdwanServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemsdwanService(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemsdwanService resource: %v", err)
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

	diags := refreshObjectSystemsdwanService(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemsdwanServiceDst(v *[]models.SystemsdwanServiceDst, sort bool) interface{} {
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

func flattenSystemsdwanServiceDst6(v *[]models.SystemsdwanServiceDst6, sort bool) interface{} {
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

func flattenSystemsdwanServiceGroups(v *[]models.SystemsdwanServiceGroups, sort bool) interface{} {
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

func flattenSystemsdwanServiceHealthCheck(v *[]models.SystemsdwanServiceHealthCheck, sort bool) interface{} {
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

func flattenSystemsdwanServiceInputDevice(v *[]models.SystemsdwanServiceInputDevice, sort bool) interface{} {
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

func flattenSystemsdwanServiceInternetServiceAppCtrl(v *[]models.SystemsdwanServiceInternetServiceAppCtrl, sort bool) interface{} {
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

func flattenSystemsdwanServiceInternetServiceAppCtrlGroup(v *[]models.SystemsdwanServiceInternetServiceAppCtrlGroup, sort bool) interface{} {
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

func flattenSystemsdwanServiceInternetServiceCustom(v *[]models.SystemsdwanServiceInternetServiceCustom, sort bool) interface{} {
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

func flattenSystemsdwanServiceInternetServiceCustomGroup(v *[]models.SystemsdwanServiceInternetServiceCustomGroup, sort bool) interface{} {
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

func flattenSystemsdwanServiceInternetServiceGroup(v *[]models.SystemsdwanServiceInternetServiceGroup, sort bool) interface{} {
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

func flattenSystemsdwanServiceInternetServiceName(v *[]models.SystemsdwanServiceInternetServiceName, sort bool) interface{} {
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

func flattenSystemsdwanServicePriorityMembers(v *[]models.SystemsdwanServicePriorityMembers, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.SeqNum; tmp != nil {
				v["seq_num"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "seq_num")
	}

	return flat
}

func flattenSystemsdwanServicePriorityZone(v *[]models.SystemsdwanServicePriorityZone, sort bool) interface{} {
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

func flattenSystemsdwanServiceSla(v *[]models.SystemsdwanServiceSla, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.HealthCheck; tmp != nil {
				v["health_check"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "health_check")
	}

	return flat
}

func flattenSystemsdwanServiceSrc(v *[]models.SystemsdwanServiceSrc, sort bool) interface{} {
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

func flattenSystemsdwanServiceSrc6(v *[]models.SystemsdwanServiceSrc6, sort bool) interface{} {
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

func flattenSystemsdwanServiceUsers(v *[]models.SystemsdwanServiceUsers, sort bool) interface{} {
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

func refreshObjectSystemsdwanService(d *schema.ResourceData, o *models.SystemsdwanService, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AddrMode != nil {
		v := *o.AddrMode

		if err = d.Set("addr_mode", v); err != nil {
			return diag.Errorf("error reading addr_mode: %v", err)
		}
	}

	if o.BandwidthWeight != nil {
		v := *o.BandwidthWeight

		if err = d.Set("bandwidth_weight", v); err != nil {
			return diag.Errorf("error reading bandwidth_weight: %v", err)
		}
	}

	if o.Default != nil {
		v := *o.Default

		if err = d.Set("default", v); err != nil {
			return diag.Errorf("error reading default: %v", err)
		}
	}

	if o.DscpForward != nil {
		v := *o.DscpForward

		if err = d.Set("dscp_forward", v); err != nil {
			return diag.Errorf("error reading dscp_forward: %v", err)
		}
	}

	if o.DscpForwardTag != nil {
		v := *o.DscpForwardTag

		if err = d.Set("dscp_forward_tag", v); err != nil {
			return diag.Errorf("error reading dscp_forward_tag: %v", err)
		}
	}

	if o.DscpReverse != nil {
		v := *o.DscpReverse

		if err = d.Set("dscp_reverse", v); err != nil {
			return diag.Errorf("error reading dscp_reverse: %v", err)
		}
	}

	if o.DscpReverseTag != nil {
		v := *o.DscpReverseTag

		if err = d.Set("dscp_reverse_tag", v); err != nil {
			return diag.Errorf("error reading dscp_reverse_tag: %v", err)
		}
	}

	if o.Dst != nil {
		if err = d.Set("dst", flattenSystemsdwanServiceDst(o.Dst, sort)); err != nil {
			return diag.Errorf("error reading dst: %v", err)
		}
	}

	if o.DstNegate != nil {
		v := *o.DstNegate

		if err = d.Set("dst_negate", v); err != nil {
			return diag.Errorf("error reading dst_negate: %v", err)
		}
	}

	if o.Dst6 != nil {
		if err = d.Set("dst6", flattenSystemsdwanServiceDst6(o.Dst6, sort)); err != nil {
			return diag.Errorf("error reading dst6: %v", err)
		}
	}

	if o.EndPort != nil {
		v := *o.EndPort

		if err = d.Set("end_port", v); err != nil {
			return diag.Errorf("error reading end_port: %v", err)
		}
	}

	if o.Gateway != nil {
		v := *o.Gateway

		if err = d.Set("gateway", v); err != nil {
			return diag.Errorf("error reading gateway: %v", err)
		}
	}

	if o.Groups != nil {
		if err = d.Set("groups", flattenSystemsdwanServiceGroups(o.Groups, sort)); err != nil {
			return diag.Errorf("error reading groups: %v", err)
		}
	}

	if o.HashMode != nil {
		v := *o.HashMode

		if err = d.Set("hash_mode", v); err != nil {
			return diag.Errorf("error reading hash_mode: %v", err)
		}
	}

	if o.HealthCheck != nil {
		if err = d.Set("health_check", flattenSystemsdwanServiceHealthCheck(o.HealthCheck, sort)); err != nil {
			return diag.Errorf("error reading health_check: %v", err)
		}
	}

	if o.HoldDownTime != nil {
		v := *o.HoldDownTime

		if err = d.Set("hold_down_time", v); err != nil {
			return diag.Errorf("error reading hold_down_time: %v", err)
		}
	}

	if o.Fosid != nil {
		v := *o.Fosid

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.InputDevice != nil {
		if err = d.Set("input_device", flattenSystemsdwanServiceInputDevice(o.InputDevice, sort)); err != nil {
			return diag.Errorf("error reading input_device: %v", err)
		}
	}

	if o.InputDeviceNegate != nil {
		v := *o.InputDeviceNegate

		if err = d.Set("input_device_negate", v); err != nil {
			return diag.Errorf("error reading input_device_negate: %v", err)
		}
	}

	if o.InternetService != nil {
		v := *o.InternetService

		if err = d.Set("internet_service", v); err != nil {
			return diag.Errorf("error reading internet_service: %v", err)
		}
	}

	if o.InternetServiceAppCtrl != nil {
		if err = d.Set("internet_service_app_ctrl", flattenSystemsdwanServiceInternetServiceAppCtrl(o.InternetServiceAppCtrl, sort)); err != nil {
			return diag.Errorf("error reading internet_service_app_ctrl: %v", err)
		}
	}

	if o.InternetServiceAppCtrlGroup != nil {
		if err = d.Set("internet_service_app_ctrl_group", flattenSystemsdwanServiceInternetServiceAppCtrlGroup(o.InternetServiceAppCtrlGroup, sort)); err != nil {
			return diag.Errorf("error reading internet_service_app_ctrl_group: %v", err)
		}
	}

	if o.InternetServiceCustom != nil {
		if err = d.Set("internet_service_custom", flattenSystemsdwanServiceInternetServiceCustom(o.InternetServiceCustom, sort)); err != nil {
			return diag.Errorf("error reading internet_service_custom: %v", err)
		}
	}

	if o.InternetServiceCustomGroup != nil {
		if err = d.Set("internet_service_custom_group", flattenSystemsdwanServiceInternetServiceCustomGroup(o.InternetServiceCustomGroup, sort)); err != nil {
			return diag.Errorf("error reading internet_service_custom_group: %v", err)
		}
	}

	if o.InternetServiceGroup != nil {
		if err = d.Set("internet_service_group", flattenSystemsdwanServiceInternetServiceGroup(o.InternetServiceGroup, sort)); err != nil {
			return diag.Errorf("error reading internet_service_group: %v", err)
		}
	}

	if o.InternetServiceName != nil {
		if err = d.Set("internet_service_name", flattenSystemsdwanServiceInternetServiceName(o.InternetServiceName, sort)); err != nil {
			return diag.Errorf("error reading internet_service_name: %v", err)
		}
	}

	if o.JitterWeight != nil {
		v := *o.JitterWeight

		if err = d.Set("jitter_weight", v); err != nil {
			return diag.Errorf("error reading jitter_weight: %v", err)
		}
	}

	if o.LatencyWeight != nil {
		v := *o.LatencyWeight

		if err = d.Set("latency_weight", v); err != nil {
			return diag.Errorf("error reading latency_weight: %v", err)
		}
	}

	if o.LinkCostFactor != nil {
		v := *o.LinkCostFactor

		if err = d.Set("link_cost_factor", v); err != nil {
			return diag.Errorf("error reading link_cost_factor: %v", err)
		}
	}

	if o.LinkCostThreshold != nil {
		v := *o.LinkCostThreshold

		if err = d.Set("link_cost_threshold", v); err != nil {
			return diag.Errorf("error reading link_cost_threshold: %v", err)
		}
	}

	if o.MinimumSlaMeetMembers != nil {
		v := *o.MinimumSlaMeetMembers

		if err = d.Set("minimum_sla_meet_members", v); err != nil {
			return diag.Errorf("error reading minimum_sla_meet_members: %v", err)
		}
	}

	if o.Mode != nil {
		v := *o.Mode

		if err = d.Set("mode", v); err != nil {
			return diag.Errorf("error reading mode: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.PacketLossWeight != nil {
		v := *o.PacketLossWeight

		if err = d.Set("packet_loss_weight", v); err != nil {
			return diag.Errorf("error reading packet_loss_weight: %v", err)
		}
	}

	if o.PassiveMeasurement != nil {
		v := *o.PassiveMeasurement

		if err = d.Set("passive_measurement", v); err != nil {
			return diag.Errorf("error reading passive_measurement: %v", err)
		}
	}

	if o.PriorityMembers != nil {
		if err = d.Set("priority_members", flattenSystemsdwanServicePriorityMembers(o.PriorityMembers, sort)); err != nil {
			return diag.Errorf("error reading priority_members: %v", err)
		}
	}

	if o.PriorityZone != nil {
		if err = d.Set("priority_zone", flattenSystemsdwanServicePriorityZone(o.PriorityZone, sort)); err != nil {
			return diag.Errorf("error reading priority_zone: %v", err)
		}
	}

	if o.Protocol != nil {
		v := *o.Protocol

		if err = d.Set("protocol", v); err != nil {
			return diag.Errorf("error reading protocol: %v", err)
		}
	}

	if o.QualityLink != nil {
		v := *o.QualityLink

		if err = d.Set("quality_link", v); err != nil {
			return diag.Errorf("error reading quality_link: %v", err)
		}
	}

	if o.Role != nil {
		v := *o.Role

		if err = d.Set("role", v); err != nil {
			return diag.Errorf("error reading role: %v", err)
		}
	}

	if o.RouteTag != nil {
		v := *o.RouteTag

		if err = d.Set("route_tag", v); err != nil {
			return diag.Errorf("error reading route_tag: %v", err)
		}
	}

	if o.Sla != nil {
		if err = d.Set("sla", flattenSystemsdwanServiceSla(o.Sla, sort)); err != nil {
			return diag.Errorf("error reading sla: %v", err)
		}
	}

	if o.SlaCompareMethod != nil {
		v := *o.SlaCompareMethod

		if err = d.Set("sla_compare_method", v); err != nil {
			return diag.Errorf("error reading sla_compare_method: %v", err)
		}
	}

	if o.Src != nil {
		if err = d.Set("src", flattenSystemsdwanServiceSrc(o.Src, sort)); err != nil {
			return diag.Errorf("error reading src: %v", err)
		}
	}

	if o.SrcNegate != nil {
		v := *o.SrcNegate

		if err = d.Set("src_negate", v); err != nil {
			return diag.Errorf("error reading src_negate: %v", err)
		}
	}

	if o.Src6 != nil {
		if err = d.Set("src6", flattenSystemsdwanServiceSrc6(o.Src6, sort)); err != nil {
			return diag.Errorf("error reading src6: %v", err)
		}
	}

	if o.StandaloneAction != nil {
		v := *o.StandaloneAction

		if err = d.Set("standalone_action", v); err != nil {
			return diag.Errorf("error reading standalone_action: %v", err)
		}
	}

	if o.StartPort != nil {
		v := *o.StartPort

		if err = d.Set("start_port", v); err != nil {
			return diag.Errorf("error reading start_port: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.TieBreak != nil {
		v := *o.TieBreak

		if err = d.Set("tie_break", v); err != nil {
			return diag.Errorf("error reading tie_break: %v", err)
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

	if o.UseShortcutSla != nil {
		v := *o.UseShortcutSla

		if err = d.Set("use_shortcut_sla", v); err != nil {
			return diag.Errorf("error reading use_shortcut_sla: %v", err)
		}
	}

	if o.Users != nil {
		if err = d.Set("users", flattenSystemsdwanServiceUsers(o.Users, sort)); err != nil {
			return diag.Errorf("error reading users: %v", err)
		}
	}

	return nil
}

func expandSystemsdwanServiceDst(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanServiceDst, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanServiceDst

	for i := range l {
		tmp := models.SystemsdwanServiceDst{}
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

func expandSystemsdwanServiceDst6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanServiceDst6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanServiceDst6

	for i := range l {
		tmp := models.SystemsdwanServiceDst6{}
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

func expandSystemsdwanServiceGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanServiceGroups, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanServiceGroups

	for i := range l {
		tmp := models.SystemsdwanServiceGroups{}
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

func expandSystemsdwanServiceHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanServiceHealthCheck, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanServiceHealthCheck

	for i := range l {
		tmp := models.SystemsdwanServiceHealthCheck{}
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

func expandSystemsdwanServiceInputDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanServiceInputDevice, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanServiceInputDevice

	for i := range l {
		tmp := models.SystemsdwanServiceInputDevice{}
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

func expandSystemsdwanServiceInternetServiceAppCtrl(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanServiceInternetServiceAppCtrl, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanServiceInternetServiceAppCtrl

	for i := range l {
		tmp := models.SystemsdwanServiceInternetServiceAppCtrl{}
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

func expandSystemsdwanServiceInternetServiceAppCtrlGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanServiceInternetServiceAppCtrlGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanServiceInternetServiceAppCtrlGroup

	for i := range l {
		tmp := models.SystemsdwanServiceInternetServiceAppCtrlGroup{}
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

func expandSystemsdwanServiceInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanServiceInternetServiceCustom, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanServiceInternetServiceCustom

	for i := range l {
		tmp := models.SystemsdwanServiceInternetServiceCustom{}
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

func expandSystemsdwanServiceInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanServiceInternetServiceCustomGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanServiceInternetServiceCustomGroup

	for i := range l {
		tmp := models.SystemsdwanServiceInternetServiceCustomGroup{}
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

func expandSystemsdwanServiceInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanServiceInternetServiceGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanServiceInternetServiceGroup

	for i := range l {
		tmp := models.SystemsdwanServiceInternetServiceGroup{}
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

func expandSystemsdwanServiceInternetServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanServiceInternetServiceName, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanServiceInternetServiceName

	for i := range l {
		tmp := models.SystemsdwanServiceInternetServiceName{}
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

func expandSystemsdwanServicePriorityMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanServicePriorityMembers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanServicePriorityMembers

	for i := range l {
		tmp := models.SystemsdwanServicePriorityMembers{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.seq_num", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.SeqNum = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemsdwanServicePriorityZone(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanServicePriorityZone, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanServicePriorityZone

	for i := range l {
		tmp := models.SystemsdwanServicePriorityZone{}
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

func expandSystemsdwanServiceSla(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanServiceSla, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanServiceSla

	for i := range l {
		tmp := models.SystemsdwanServiceSla{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.health_check", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HealthCheck = &v2
			}
		}

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

func expandSystemsdwanServiceSrc(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanServiceSrc, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanServiceSrc

	for i := range l {
		tmp := models.SystemsdwanServiceSrc{}
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

func expandSystemsdwanServiceSrc6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanServiceSrc6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanServiceSrc6

	for i := range l {
		tmp := models.SystemsdwanServiceSrc6{}
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

func expandSystemsdwanServiceUsers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemsdwanServiceUsers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemsdwanServiceUsers

	for i := range l {
		tmp := models.SystemsdwanServiceUsers{}
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

func getObjectSystemsdwanService(d *schema.ResourceData, sv string) (*models.SystemsdwanService, diag.Diagnostics) {
	obj := models.SystemsdwanService{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("addr_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("addr_mode", sv)
				diags = append(diags, e)
			}
			obj.AddrMode = &v2
		}
	}
	if v1, ok := d.GetOk("bandwidth_weight"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bandwidth_weight", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.BandwidthWeight = &tmp
		}
	}
	if v1, ok := d.GetOk("default"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default", sv)
				diags = append(diags, e)
			}
			obj.Default = &v2
		}
	}
	if v1, ok := d.GetOk("dscp_forward"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dscp_forward", sv)
				diags = append(diags, e)
			}
			obj.DscpForward = &v2
		}
	}
	if v1, ok := d.GetOk("dscp_forward_tag"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dscp_forward_tag", sv)
				diags = append(diags, e)
			}
			obj.DscpForwardTag = &v2
		}
	}
	if v1, ok := d.GetOk("dscp_reverse"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dscp_reverse", sv)
				diags = append(diags, e)
			}
			obj.DscpReverse = &v2
		}
	}
	if v1, ok := d.GetOk("dscp_reverse_tag"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dscp_reverse_tag", sv)
				diags = append(diags, e)
			}
			obj.DscpReverseTag = &v2
		}
	}
	if v, ok := d.GetOk("dst"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dst", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanServiceDst(d, v, "dst", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dst = t
		}
	} else if d.HasChange("dst") {
		old, new := d.GetChange("dst")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dst = &[]models.SystemsdwanServiceDst{}
		}
	}
	if v1, ok := d.GetOk("dst_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dst_negate", sv)
				diags = append(diags, e)
			}
			obj.DstNegate = &v2
		}
	}
	if v, ok := d.GetOk("dst6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dst6", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanServiceDst6(d, v, "dst6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dst6 = t
		}
	} else if d.HasChange("dst6") {
		old, new := d.GetChange("dst6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dst6 = &[]models.SystemsdwanServiceDst6{}
		}
	}
	if v1, ok := d.GetOk("end_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("end_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EndPort = &tmp
		}
	}
	if v1, ok := d.GetOk("gateway"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gateway", sv)
				diags = append(diags, e)
			}
			obj.Gateway = &v2
		}
	}
	if v, ok := d.GetOk("groups"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("groups", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanServiceGroups(d, v, "groups", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Groups = t
		}
	} else if d.HasChange("groups") {
		old, new := d.GetChange("groups")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Groups = &[]models.SystemsdwanServiceGroups{}
		}
	}
	if v1, ok := d.GetOk("hash_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hash_mode", sv)
				diags = append(diags, e)
			}
			obj.HashMode = &v2
		}
	}
	if v, ok := d.GetOk("health_check"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("health_check", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanServiceHealthCheck(d, v, "health_check", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.HealthCheck = t
		}
	} else if d.HasChange("health_check") {
		old, new := d.GetChange("health_check")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.HealthCheck = &[]models.SystemsdwanServiceHealthCheck{}
		}
	}
	if v1, ok := d.GetOk("hold_down_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hold_down_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HoldDownTime = &tmp
		}
	}
	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Fosid = &tmp
		}
	}
	if v, ok := d.GetOk("input_device"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("input_device", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanServiceInputDevice(d, v, "input_device", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InputDevice = t
		}
	} else if d.HasChange("input_device") {
		old, new := d.GetChange("input_device")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InputDevice = &[]models.SystemsdwanServiceInputDevice{}
		}
	}
	if v1, ok := d.GetOk("input_device_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("input_device_negate", sv)
				diags = append(diags, e)
			}
			obj.InputDeviceNegate = &v2
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
	if v, ok := d.GetOk("internet_service_app_ctrl"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_app_ctrl", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanServiceInternetServiceAppCtrl(d, v, "internet_service_app_ctrl", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceAppCtrl = t
		}
	} else if d.HasChange("internet_service_app_ctrl") {
		old, new := d.GetChange("internet_service_app_ctrl")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceAppCtrl = &[]models.SystemsdwanServiceInternetServiceAppCtrl{}
		}
	}
	if v, ok := d.GetOk("internet_service_app_ctrl_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_app_ctrl_group", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanServiceInternetServiceAppCtrlGroup(d, v, "internet_service_app_ctrl_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceAppCtrlGroup = t
		}
	} else if d.HasChange("internet_service_app_ctrl_group") {
		old, new := d.GetChange("internet_service_app_ctrl_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceAppCtrlGroup = &[]models.SystemsdwanServiceInternetServiceAppCtrlGroup{}
		}
	}
	if v, ok := d.GetOk("internet_service_custom"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_custom", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanServiceInternetServiceCustom(d, v, "internet_service_custom", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceCustom = t
		}
	} else if d.HasChange("internet_service_custom") {
		old, new := d.GetChange("internet_service_custom")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceCustom = &[]models.SystemsdwanServiceInternetServiceCustom{}
		}
	}
	if v, ok := d.GetOk("internet_service_custom_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_custom_group", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanServiceInternetServiceCustomGroup(d, v, "internet_service_custom_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceCustomGroup = t
		}
	} else if d.HasChange("internet_service_custom_group") {
		old, new := d.GetChange("internet_service_custom_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceCustomGroup = &[]models.SystemsdwanServiceInternetServiceCustomGroup{}
		}
	}
	if v, ok := d.GetOk("internet_service_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_group", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanServiceInternetServiceGroup(d, v, "internet_service_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceGroup = t
		}
	} else if d.HasChange("internet_service_group") {
		old, new := d.GetChange("internet_service_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceGroup = &[]models.SystemsdwanServiceInternetServiceGroup{}
		}
	}
	if v, ok := d.GetOk("internet_service_name"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_name", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanServiceInternetServiceName(d, v, "internet_service_name", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceName = t
		}
	} else if d.HasChange("internet_service_name") {
		old, new := d.GetChange("internet_service_name")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceName = &[]models.SystemsdwanServiceInternetServiceName{}
		}
	}
	if v1, ok := d.GetOk("jitter_weight"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("jitter_weight", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.JitterWeight = &tmp
		}
	}
	if v1, ok := d.GetOk("latency_weight"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("latency_weight", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LatencyWeight = &tmp
		}
	}
	if v1, ok := d.GetOk("link_cost_factor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("link_cost_factor", sv)
				diags = append(diags, e)
			}
			obj.LinkCostFactor = &v2
		}
	}
	if v1, ok := d.GetOk("link_cost_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("link_cost_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LinkCostThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("minimum_sla_meet_members"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("minimum_sla_meet_members", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MinimumSlaMeetMembers = &tmp
		}
	}
	if v1, ok := d.GetOk("mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mode", sv)
				diags = append(diags, e)
			}
			obj.Mode = &v2
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
	if v1, ok := d.GetOk("packet_loss_weight"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("packet_loss_weight", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PacketLossWeight = &tmp
		}
	}
	if v1, ok := d.GetOk("passive_measurement"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("passive_measurement", sv)
				diags = append(diags, e)
			}
			obj.PassiveMeasurement = &v2
		}
	}
	if v, ok := d.GetOk("priority_members"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("priority_members", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanServicePriorityMembers(d, v, "priority_members", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.PriorityMembers = t
		}
	} else if d.HasChange("priority_members") {
		old, new := d.GetChange("priority_members")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.PriorityMembers = &[]models.SystemsdwanServicePriorityMembers{}
		}
	}
	if v, ok := d.GetOk("priority_zone"); ok {
		if !utils.CheckVer(sv, "v7.0.1", "") {
			e := utils.AttributeVersionWarning("priority_zone", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanServicePriorityZone(d, v, "priority_zone", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.PriorityZone = t
		}
	} else if d.HasChange("priority_zone") {
		old, new := d.GetChange("priority_zone")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.PriorityZone = &[]models.SystemsdwanServicePriorityZone{}
		}
	}
	if v1, ok := d.GetOk("protocol"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("protocol", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Protocol = &tmp
		}
	}
	if v1, ok := d.GetOk("quality_link"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("quality_link", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.QualityLink = &tmp
		}
	}
	if v1, ok := d.GetOk("role"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("role", sv)
				diags = append(diags, e)
			}
			obj.Role = &v2
		}
	}
	if v1, ok := d.GetOk("route_tag"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("route_tag", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RouteTag = &tmp
		}
	}
	if v, ok := d.GetOk("sla"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("sla", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanServiceSla(d, v, "sla", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Sla = t
		}
	} else if d.HasChange("sla") {
		old, new := d.GetChange("sla")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Sla = &[]models.SystemsdwanServiceSla{}
		}
	}
	if v1, ok := d.GetOk("sla_compare_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sla_compare_method", sv)
				diags = append(diags, e)
			}
			obj.SlaCompareMethod = &v2
		}
	}
	if v, ok := d.GetOk("src"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("src", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanServiceSrc(d, v, "src", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Src = t
		}
	} else if d.HasChange("src") {
		old, new := d.GetChange("src")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Src = &[]models.SystemsdwanServiceSrc{}
		}
	}
	if v1, ok := d.GetOk("src_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("src_negate", sv)
				diags = append(diags, e)
			}
			obj.SrcNegate = &v2
		}
	}
	if v, ok := d.GetOk("src6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("src6", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanServiceSrc6(d, v, "src6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Src6 = t
		}
	} else if d.HasChange("src6") {
		old, new := d.GetChange("src6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Src6 = &[]models.SystemsdwanServiceSrc6{}
		}
	}
	if v1, ok := d.GetOk("standalone_action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("standalone_action", sv)
				diags = append(diags, e)
			}
			obj.StandaloneAction = &v2
		}
	}
	if v1, ok := d.GetOk("start_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("start_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.StartPort = &tmp
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
	if v1, ok := d.GetOk("tie_break"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("tie_break", sv)
				diags = append(diags, e)
			}
			obj.TieBreak = &v2
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
	if v1, ok := d.GetOk("use_shortcut_sla"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.5", "") {
				e := utils.AttributeVersionWarning("use_shortcut_sla", sv)
				diags = append(diags, e)
			}
			obj.UseShortcutSla = &v2
		}
	}
	if v, ok := d.GetOk("users"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("users", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemsdwanServiceUsers(d, v, "users", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Users = t
		}
	} else if d.HasChange("users") {
		old, new := d.GetChange("users")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Users = &[]models.SystemsdwanServiceUsers{}
		}
	}
	return &obj, diags
}
