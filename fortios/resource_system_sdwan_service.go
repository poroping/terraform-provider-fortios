// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemsdwanService() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemsdwanServiceCreate,
		Read:   resourceSystemsdwanServiceRead,
		Update: resourceSystemsdwanServiceUpdate,
		Delete: resourceSystemsdwanServiceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"batchid": {
				Type:        schema.TypeInt,
				Description: "Associate with batch. From 6.4.x+. Currently a WIP and broken.",
				Optional:    true,
				Default:     0,
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
				ValidateFunc: fortiValidateEnum([]string{"ipv4", "ipv6"}),

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
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable use of SD-WAN as default service.",
				Optional:    true,
				Computed:    true,
			},
			"dscp_forward": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

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
				ValidateFunc: fortiValidateEnableDisable(),

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
				ValidateFunc: fortiValidateEnableDisable(),

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
				ValidateFunc: fortiValidateEnableDisable(),

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
				ValidateFunc: fortiValidateEnum([]string{"round-robin", "source-ip-based", "source-dest-ip-based", "inbandwidth", "outbandwidth", "bibandwidth"}),

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
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable negation of input device match.",
				Optional:    true,
				Computed:    true,
			},
			"internet_service": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

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
				ValidateFunc: fortiValidateEnum([]string{"latency", "jitter", "packet-loss", "inbandwidth", "outbandwidth", "bibandwidth", "custom-profile-1"}),

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
				ValidateFunc: fortiValidateEnum([]string{"auto", "manual", "priority", "sla", "load-balance"}),

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
				ValidateFunc: fortiValidateEnum([]string{"standalone", "primary", "secondary"}),

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
				ValidateFunc: fortiValidateEnum([]string{"order", "number"}),

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
				ValidateFunc: fortiValidateEnableDisable(),

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
				ValidateFunc: fortiValidateEnableDisable(),

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
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable SD-WAN service.",
				Optional:    true,
				Computed:    true,
			},
			"tie_break": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"zone", "cfg-order", "fib-best-match"}),

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
				ValidateFunc: fortiValidateEnableDisable(),

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

func resourceSystemsdwanServiceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	allow_append := false

	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}

	urlparams["allow_append"] = []string{strconv.FormatBool(allow_append)}

	key := "id"
	mkey := ""
	if v, ok := d.GetOk(key); ok {
		if s, ok := v.(string); ok {
			mkey = s
		}
	}

	obj, err := getObjectSystemsdwanService(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating SystemsdwanService resource while getting object: %v", err)
	}

	if mkey == "" && allow_append {
		return fmt.Errorf("error creating SystemsdwanService resource: %q must be set if \"allow_append\" is true", key)
	}

	o := make(map[string]interface{})
	if mkey != "" && allow_append {
		o, err = c.UpdateSystemsdwanService(obj, mkey, vdomparam, urlparams, batchid)
	} else {
		o, err = c.CreateSystemsdwanService(obj, vdomparam, urlparams, batchid)
	}

	if err != nil {
		return fmt.Errorf("error creating SystemsdwanService resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemsdwanService")
	}

	return resourceSystemsdwanServiceRead(d, m)
}

func resourceSystemsdwanServiceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	obj, err := getObjectSystemsdwanService(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating SystemsdwanService resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemsdwanService(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating SystemsdwanService resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemsdwanService")
	}

	return resourceSystemsdwanServiceRead(d, m)
}

func resourceSystemsdwanServiceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	err := c.DeleteSystemsdwanService(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting SystemsdwanService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemsdwanServiceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	o, err := c.ReadSystemsdwanService(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading SystemsdwanService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemsdwanService(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading SystemsdwanService resource from API: %v", err)
	}
	return nil
}

func flattenSystemsdwanServiceAddrMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceBandwidthWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceDefault(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceDscpForward(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceDscpForwardTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceDscpReverse(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceDscpReverseTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceDst(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemsdwanServiceDstName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemsdwanServiceDstName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceDstNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceDst6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemsdwanServiceDst6Name(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemsdwanServiceDst6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceEndPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceGateway(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceGroups(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemsdwanServiceGroupsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemsdwanServiceGroupsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceHashMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceHealthCheck(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemsdwanServiceHealthCheckName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemsdwanServiceHealthCheckName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceHoldDownTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceInputDevice(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemsdwanServiceInputDeviceName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemsdwanServiceInputDeviceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceInputDeviceNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceInternetService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceInternetServiceAppCtrl(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenSystemsdwanServiceInternetServiceAppCtrlId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemsdwanServiceInternetServiceAppCtrlId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceInternetServiceAppCtrlGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemsdwanServiceInternetServiceAppCtrlGroupName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemsdwanServiceInternetServiceAppCtrlGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemsdwanServiceInternetServiceCustomName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemsdwanServiceInternetServiceCustomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceInternetServiceCustomGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemsdwanServiceInternetServiceCustomGroupName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemsdwanServiceInternetServiceCustomGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceInternetServiceGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemsdwanServiceInternetServiceGroupName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemsdwanServiceInternetServiceGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceInternetServiceName(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemsdwanServiceInternetServiceNameName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemsdwanServiceInternetServiceNameName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceJitterWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceLatencyWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceLinkCostFactor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceLinkCostThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceMinimumSlaMeetMembers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServicePacketLossWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServicePriorityMembers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := i["seq-num"]; ok {

			tmp["seq_num"] = flattenSystemsdwanServicePriorityMembersSeqNum(i["seq-num"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "seq-num", d)
	return result
}

func flattenSystemsdwanServicePriorityMembersSeqNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServicePriorityZone(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemsdwanServicePriorityZoneName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemsdwanServicePriorityZoneName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceQualityLink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceRole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceRouteTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceSla(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {

			tmp["health_check"] = flattenSystemsdwanServiceSlaHealthCheck(i["health-check"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenSystemsdwanServiceSlaId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "health-check", d)
	return result
}

func flattenSystemsdwanServiceSlaHealthCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceSlaId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceSlaCompareMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceSrc(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemsdwanServiceSrcName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemsdwanServiceSrcName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceSrcNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceSrc6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemsdwanServiceSrc6Name(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemsdwanServiceSrc6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceStandaloneAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceStartPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceTieBreak(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceTos(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceTosMask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceUseShortcutSla(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanServiceUsers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemsdwanServiceUsersName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemsdwanServiceUsersName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemsdwanService(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("addr_mode", flattenSystemsdwanServiceAddrMode(o["addr-mode"], d, "addr_mode", sv)); err != nil {
		if !fortiAPIPatch(o["addr-mode"]) {
			return fmt.Errorf("error reading addr_mode: %v", err)
		}
	}

	if err = d.Set("bandwidth_weight", flattenSystemsdwanServiceBandwidthWeight(o["bandwidth-weight"], d, "bandwidth_weight", sv)); err != nil {
		if !fortiAPIPatch(o["bandwidth-weight"]) {
			return fmt.Errorf("error reading bandwidth_weight: %v", err)
		}
	}

	if err = d.Set("default", flattenSystemsdwanServiceDefault(o["default"], d, "default", sv)); err != nil {
		if !fortiAPIPatch(o["default"]) {
			return fmt.Errorf("error reading default: %v", err)
		}
	}

	if err = d.Set("dscp_forward", flattenSystemsdwanServiceDscpForward(o["dscp-forward"], d, "dscp_forward", sv)); err != nil {
		if !fortiAPIPatch(o["dscp-forward"]) {
			return fmt.Errorf("error reading dscp_forward: %v", err)
		}
	}

	if err = d.Set("dscp_forward_tag", flattenSystemsdwanServiceDscpForwardTag(o["dscp-forward-tag"], d, "dscp_forward_tag", sv)); err != nil {
		if !fortiAPIPatch(o["dscp-forward-tag"]) {
			return fmt.Errorf("error reading dscp_forward_tag: %v", err)
		}
	}

	if err = d.Set("dscp_reverse", flattenSystemsdwanServiceDscpReverse(o["dscp-reverse"], d, "dscp_reverse", sv)); err != nil {
		if !fortiAPIPatch(o["dscp-reverse"]) {
			return fmt.Errorf("error reading dscp_reverse: %v", err)
		}
	}

	if err = d.Set("dscp_reverse_tag", flattenSystemsdwanServiceDscpReverseTag(o["dscp-reverse-tag"], d, "dscp_reverse_tag", sv)); err != nil {
		if !fortiAPIPatch(o["dscp-reverse-tag"]) {
			return fmt.Errorf("error reading dscp_reverse_tag: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dst", flattenSystemsdwanServiceDst(o["dst"], d, "dst", sv)); err != nil {
			if !fortiAPIPatch(o["dst"]) {
				return fmt.Errorf("error reading dst: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dst"); ok {
			if err = d.Set("dst", flattenSystemsdwanServiceDst(o["dst"], d, "dst", sv)); err != nil {
				if !fortiAPIPatch(o["dst"]) {
					return fmt.Errorf("error reading dst: %v", err)
				}
			}
		}
	}

	if err = d.Set("dst_negate", flattenSystemsdwanServiceDstNegate(o["dst-negate"], d, "dst_negate", sv)); err != nil {
		if !fortiAPIPatch(o["dst-negate"]) {
			return fmt.Errorf("error reading dst_negate: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dst6", flattenSystemsdwanServiceDst6(o["dst6"], d, "dst6", sv)); err != nil {
			if !fortiAPIPatch(o["dst6"]) {
				return fmt.Errorf("error reading dst6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dst6"); ok {
			if err = d.Set("dst6", flattenSystemsdwanServiceDst6(o["dst6"], d, "dst6", sv)); err != nil {
				if !fortiAPIPatch(o["dst6"]) {
					return fmt.Errorf("error reading dst6: %v", err)
				}
			}
		}
	}

	if err = d.Set("end_port", flattenSystemsdwanServiceEndPort(o["end-port"], d, "end_port", sv)); err != nil {
		if !fortiAPIPatch(o["end-port"]) {
			return fmt.Errorf("error reading end_port: %v", err)
		}
	}

	if err = d.Set("gateway", flattenSystemsdwanServiceGateway(o["gateway"], d, "gateway", sv)); err != nil {
		if !fortiAPIPatch(o["gateway"]) {
			return fmt.Errorf("error reading gateway: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("groups", flattenSystemsdwanServiceGroups(o["groups"], d, "groups", sv)); err != nil {
			if !fortiAPIPatch(o["groups"]) {
				return fmt.Errorf("error reading groups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("groups"); ok {
			if err = d.Set("groups", flattenSystemsdwanServiceGroups(o["groups"], d, "groups", sv)); err != nil {
				if !fortiAPIPatch(o["groups"]) {
					return fmt.Errorf("error reading groups: %v", err)
				}
			}
		}
	}

	if err = d.Set("hash_mode", flattenSystemsdwanServiceHashMode(o["hash-mode"], d, "hash_mode", sv)); err != nil {
		if !fortiAPIPatch(o["hash-mode"]) {
			return fmt.Errorf("error reading hash_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("health_check", flattenSystemsdwanServiceHealthCheck(o["health-check"], d, "health_check", sv)); err != nil {
			if !fortiAPIPatch(o["health-check"]) {
				return fmt.Errorf("error reading health_check: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("health_check"); ok {
			if err = d.Set("health_check", flattenSystemsdwanServiceHealthCheck(o["health-check"], d, "health_check", sv)); err != nil {
				if !fortiAPIPatch(o["health-check"]) {
					return fmt.Errorf("error reading health_check: %v", err)
				}
			}
		}
	}

	if err = d.Set("hold_down_time", flattenSystemsdwanServiceHoldDownTime(o["hold-down-time"], d, "hold_down_time", sv)); err != nil {
		if !fortiAPIPatch(o["hold-down-time"]) {
			return fmt.Errorf("error reading hold_down_time: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemsdwanServiceId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("error reading fosid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("input_device", flattenSystemsdwanServiceInputDevice(o["input-device"], d, "input_device", sv)); err != nil {
			if !fortiAPIPatch(o["input-device"]) {
				return fmt.Errorf("error reading input_device: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("input_device"); ok {
			if err = d.Set("input_device", flattenSystemsdwanServiceInputDevice(o["input-device"], d, "input_device", sv)); err != nil {
				if !fortiAPIPatch(o["input-device"]) {
					return fmt.Errorf("error reading input_device: %v", err)
				}
			}
		}
	}

	if err = d.Set("input_device_negate", flattenSystemsdwanServiceInputDeviceNegate(o["input-device-negate"], d, "input_device_negate", sv)); err != nil {
		if !fortiAPIPatch(o["input-device-negate"]) {
			return fmt.Errorf("error reading input_device_negate: %v", err)
		}
	}

	if err = d.Set("internet_service", flattenSystemsdwanServiceInternetService(o["internet-service"], d, "internet_service", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service"]) {
			return fmt.Errorf("error reading internet_service: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_app_ctrl", flattenSystemsdwanServiceInternetServiceAppCtrl(o["internet-service-app-ctrl"], d, "internet_service_app_ctrl", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-app-ctrl"]) {
				return fmt.Errorf("error reading internet_service_app_ctrl: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_app_ctrl"); ok {
			if err = d.Set("internet_service_app_ctrl", flattenSystemsdwanServiceInternetServiceAppCtrl(o["internet-service-app-ctrl"], d, "internet_service_app_ctrl", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-app-ctrl"]) {
					return fmt.Errorf("error reading internet_service_app_ctrl: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_app_ctrl_group", flattenSystemsdwanServiceInternetServiceAppCtrlGroup(o["internet-service-app-ctrl-group"], d, "internet_service_app_ctrl_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-app-ctrl-group"]) {
				return fmt.Errorf("error reading internet_service_app_ctrl_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_app_ctrl_group"); ok {
			if err = d.Set("internet_service_app_ctrl_group", flattenSystemsdwanServiceInternetServiceAppCtrlGroup(o["internet-service-app-ctrl-group"], d, "internet_service_app_ctrl_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-app-ctrl-group"]) {
					return fmt.Errorf("error reading internet_service_app_ctrl_group: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_custom", flattenSystemsdwanServiceInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-custom"]) {
				return fmt.Errorf("error reading internet_service_custom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_custom"); ok {
			if err = d.Set("internet_service_custom", flattenSystemsdwanServiceInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-custom"]) {
					return fmt.Errorf("error reading internet_service_custom: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_custom_group", flattenSystemsdwanServiceInternetServiceCustomGroup(o["internet-service-custom-group"], d, "internet_service_custom_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-custom-group"]) {
				return fmt.Errorf("error reading internet_service_custom_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_custom_group"); ok {
			if err = d.Set("internet_service_custom_group", flattenSystemsdwanServiceInternetServiceCustomGroup(o["internet-service-custom-group"], d, "internet_service_custom_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-custom-group"]) {
					return fmt.Errorf("error reading internet_service_custom_group: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_group", flattenSystemsdwanServiceInternetServiceGroup(o["internet-service-group"], d, "internet_service_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-group"]) {
				return fmt.Errorf("error reading internet_service_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_group"); ok {
			if err = d.Set("internet_service_group", flattenSystemsdwanServiceInternetServiceGroup(o["internet-service-group"], d, "internet_service_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-group"]) {
					return fmt.Errorf("error reading internet_service_group: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_name", flattenSystemsdwanServiceInternetServiceName(o["internet-service-name"], d, "internet_service_name", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-name"]) {
				return fmt.Errorf("error reading internet_service_name: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_name"); ok {
			if err = d.Set("internet_service_name", flattenSystemsdwanServiceInternetServiceName(o["internet-service-name"], d, "internet_service_name", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-name"]) {
					return fmt.Errorf("error reading internet_service_name: %v", err)
				}
			}
		}
	}

	if err = d.Set("jitter_weight", flattenSystemsdwanServiceJitterWeight(o["jitter-weight"], d, "jitter_weight", sv)); err != nil {
		if !fortiAPIPatch(o["jitter-weight"]) {
			return fmt.Errorf("error reading jitter_weight: %v", err)
		}
	}

	if err = d.Set("latency_weight", flattenSystemsdwanServiceLatencyWeight(o["latency-weight"], d, "latency_weight", sv)); err != nil {
		if !fortiAPIPatch(o["latency-weight"]) {
			return fmt.Errorf("error reading latency_weight: %v", err)
		}
	}

	if err = d.Set("link_cost_factor", flattenSystemsdwanServiceLinkCostFactor(o["link-cost-factor"], d, "link_cost_factor", sv)); err != nil {
		if !fortiAPIPatch(o["link-cost-factor"]) {
			return fmt.Errorf("error reading link_cost_factor: %v", err)
		}
	}

	if err = d.Set("link_cost_threshold", flattenSystemsdwanServiceLinkCostThreshold(o["link-cost-threshold"], d, "link_cost_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["link-cost-threshold"]) {
			return fmt.Errorf("error reading link_cost_threshold: %v", err)
		}
	}

	if err = d.Set("minimum_sla_meet_members", flattenSystemsdwanServiceMinimumSlaMeetMembers(o["minimum-sla-meet-members"], d, "minimum_sla_meet_members", sv)); err != nil {
		if !fortiAPIPatch(o["minimum-sla-meet-members"]) {
			return fmt.Errorf("error reading minimum_sla_meet_members: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemsdwanServiceMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("error reading mode: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemsdwanServiceName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("packet_loss_weight", flattenSystemsdwanServicePacketLossWeight(o["packet-loss-weight"], d, "packet_loss_weight", sv)); err != nil {
		if !fortiAPIPatch(o["packet-loss-weight"]) {
			return fmt.Errorf("error reading packet_loss_weight: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("priority_members", flattenSystemsdwanServicePriorityMembers(o["priority-members"], d, "priority_members", sv)); err != nil {
			if !fortiAPIPatch(o["priority-members"]) {
				return fmt.Errorf("error reading priority_members: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("priority_members"); ok {
			if err = d.Set("priority_members", flattenSystemsdwanServicePriorityMembers(o["priority-members"], d, "priority_members", sv)); err != nil {
				if !fortiAPIPatch(o["priority-members"]) {
					return fmt.Errorf("error reading priority_members: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("priority_zone", flattenSystemsdwanServicePriorityZone(o["priority-zone"], d, "priority_zone", sv)); err != nil {
			if !fortiAPIPatch(o["priority-zone"]) {
				return fmt.Errorf("error reading priority_zone: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("priority_zone"); ok {
			if err = d.Set("priority_zone", flattenSystemsdwanServicePriorityZone(o["priority-zone"], d, "priority_zone", sv)); err != nil {
				if !fortiAPIPatch(o["priority-zone"]) {
					return fmt.Errorf("error reading priority_zone: %v", err)
				}
			}
		}
	}

	if err = d.Set("protocol", flattenSystemsdwanServiceProtocol(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("error reading protocol: %v", err)
		}
	}

	if err = d.Set("quality_link", flattenSystemsdwanServiceQualityLink(o["quality-link"], d, "quality_link", sv)); err != nil {
		if !fortiAPIPatch(o["quality-link"]) {
			return fmt.Errorf("error reading quality_link: %v", err)
		}
	}

	if err = d.Set("role", flattenSystemsdwanServiceRole(o["role"], d, "role", sv)); err != nil {
		if !fortiAPIPatch(o["role"]) {
			return fmt.Errorf("error reading role: %v", err)
		}
	}

	if err = d.Set("route_tag", flattenSystemsdwanServiceRouteTag(o["route-tag"], d, "route_tag", sv)); err != nil {
		if !fortiAPIPatch(o["route-tag"]) {
			return fmt.Errorf("error reading route_tag: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("sla", flattenSystemsdwanServiceSla(o["sla"], d, "sla", sv)); err != nil {
			if !fortiAPIPatch(o["sla"]) {
				return fmt.Errorf("error reading sla: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sla"); ok {
			if err = d.Set("sla", flattenSystemsdwanServiceSla(o["sla"], d, "sla", sv)); err != nil {
				if !fortiAPIPatch(o["sla"]) {
					return fmt.Errorf("error reading sla: %v", err)
				}
			}
		}
	}

	if err = d.Set("sla_compare_method", flattenSystemsdwanServiceSlaCompareMethod(o["sla-compare-method"], d, "sla_compare_method", sv)); err != nil {
		if !fortiAPIPatch(o["sla-compare-method"]) {
			return fmt.Errorf("error reading sla_compare_method: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("src", flattenSystemsdwanServiceSrc(o["src"], d, "src", sv)); err != nil {
			if !fortiAPIPatch(o["src"]) {
				return fmt.Errorf("error reading src: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("src"); ok {
			if err = d.Set("src", flattenSystemsdwanServiceSrc(o["src"], d, "src", sv)); err != nil {
				if !fortiAPIPatch(o["src"]) {
					return fmt.Errorf("error reading src: %v", err)
				}
			}
		}
	}

	if err = d.Set("src_negate", flattenSystemsdwanServiceSrcNegate(o["src-negate"], d, "src_negate", sv)); err != nil {
		if !fortiAPIPatch(o["src-negate"]) {
			return fmt.Errorf("error reading src_negate: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("src6", flattenSystemsdwanServiceSrc6(o["src6"], d, "src6", sv)); err != nil {
			if !fortiAPIPatch(o["src6"]) {
				return fmt.Errorf("error reading src6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("src6"); ok {
			if err = d.Set("src6", flattenSystemsdwanServiceSrc6(o["src6"], d, "src6", sv)); err != nil {
				if !fortiAPIPatch(o["src6"]) {
					return fmt.Errorf("error reading src6: %v", err)
				}
			}
		}
	}

	if err = d.Set("standalone_action", flattenSystemsdwanServiceStandaloneAction(o["standalone-action"], d, "standalone_action", sv)); err != nil {
		if !fortiAPIPatch(o["standalone-action"]) {
			return fmt.Errorf("error reading standalone_action: %v", err)
		}
	}

	if err = d.Set("start_port", flattenSystemsdwanServiceStartPort(o["start-port"], d, "start_port", sv)); err != nil {
		if !fortiAPIPatch(o["start-port"]) {
			return fmt.Errorf("error reading start_port: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemsdwanServiceStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("error reading status: %v", err)
		}
	}

	if err = d.Set("tie_break", flattenSystemsdwanServiceTieBreak(o["tie-break"], d, "tie_break", sv)); err != nil {
		if !fortiAPIPatch(o["tie-break"]) {
			return fmt.Errorf("error reading tie_break: %v", err)
		}
	}

	if err = d.Set("tos", flattenSystemsdwanServiceTos(o["tos"], d, "tos", sv)); err != nil {
		if !fortiAPIPatch(o["tos"]) {
			return fmt.Errorf("error reading tos: %v", err)
		}
	}

	if err = d.Set("tos_mask", flattenSystemsdwanServiceTosMask(o["tos-mask"], d, "tos_mask", sv)); err != nil {
		if !fortiAPIPatch(o["tos-mask"]) {
			return fmt.Errorf("error reading tos_mask: %v", err)
		}
	}

	if err = d.Set("use_shortcut_sla", flattenSystemsdwanServiceUseShortcutSla(o["use-shortcut-sla"], d, "use_shortcut_sla", sv)); err != nil {
		if !fortiAPIPatch(o["use-shortcut-sla"]) {
			return fmt.Errorf("error reading use_shortcut_sla: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("users", flattenSystemsdwanServiceUsers(o["users"], d, "users", sv)); err != nil {
			if !fortiAPIPatch(o["users"]) {
				return fmt.Errorf("error reading users: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("users"); ok {
			if err = d.Set("users", flattenSystemsdwanServiceUsers(o["users"], d, "users", sv)); err != nil {
				if !fortiAPIPatch(o["users"]) {
					return fmt.Errorf("error reading users: %v", err)
				}
			}
		}
	}

	return nil
}

func expandSystemsdwanServiceAddrMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceBandwidthWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceDefault(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceDscpForward(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceDscpForwardTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceDscpReverse(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceDscpReverseTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceDst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemsdwanServiceDstName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemsdwanServiceDstName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceDstNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceDst6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemsdwanServiceDst6Name(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemsdwanServiceDst6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceEndPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemsdwanServiceGroupsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemsdwanServiceGroupsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceHashMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemsdwanServiceHealthCheckName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemsdwanServiceHealthCheckName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceHoldDownTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceInputDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemsdwanServiceInputDeviceName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemsdwanServiceInputDeviceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceInputDeviceNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceInternetService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceInternetServiceAppCtrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandSystemsdwanServiceInternetServiceAppCtrlId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemsdwanServiceInternetServiceAppCtrlId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceInternetServiceAppCtrlGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemsdwanServiceInternetServiceAppCtrlGroupName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemsdwanServiceInternetServiceAppCtrlGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemsdwanServiceInternetServiceCustomName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemsdwanServiceInternetServiceCustomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemsdwanServiceInternetServiceCustomGroupName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemsdwanServiceInternetServiceCustomGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemsdwanServiceInternetServiceGroupName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemsdwanServiceInternetServiceGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceInternetServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemsdwanServiceInternetServiceNameName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemsdwanServiceInternetServiceNameName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceJitterWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceLatencyWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceLinkCostFactor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceLinkCostThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceMinimumSlaMeetMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServicePacketLossWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServicePriorityMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["seq-num"], _ = expandSystemsdwanServicePriorityMembersSeqNum(d, i["seq_num"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemsdwanServicePriorityMembersSeqNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServicePriorityZone(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemsdwanServicePriorityZoneName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemsdwanServicePriorityZoneName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceQualityLink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceRole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceRouteTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceSla(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["health-check"], _ = expandSystemsdwanServiceSlaHealthCheck(d, i["health_check"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandSystemsdwanServiceSlaId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemsdwanServiceSlaHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceSlaId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceSlaCompareMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceSrc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemsdwanServiceSrcName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemsdwanServiceSrcName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceSrcNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceSrc6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemsdwanServiceSrc6Name(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemsdwanServiceSrc6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceStandaloneAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceStartPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceTieBreak(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceTos(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceTosMask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceUseShortcutSla(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanServiceUsers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemsdwanServiceUsersName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemsdwanServiceUsersName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemsdwanService(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_mode"); ok {
		t, err := expandSystemsdwanServiceAddrMode(d, v, "addr_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-mode"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_weight"); ok {
		t, err := expandSystemsdwanServiceBandwidthWeight(d, v, "bandwidth_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-weight"] = t
		}
	}

	if v, ok := d.GetOk("default"); ok {
		t, err := expandSystemsdwanServiceDefault(d, v, "default", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default"] = t
		}
	}

	if v, ok := d.GetOk("dscp_forward"); ok {
		t, err := expandSystemsdwanServiceDscpForward(d, v, "dscp_forward", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-forward"] = t
		}
	}

	if v, ok := d.GetOk("dscp_forward_tag"); ok {
		t, err := expandSystemsdwanServiceDscpForwardTag(d, v, "dscp_forward_tag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-forward-tag"] = t
		}
	}

	if v, ok := d.GetOk("dscp_reverse"); ok {
		t, err := expandSystemsdwanServiceDscpReverse(d, v, "dscp_reverse", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-reverse"] = t
		}
	}

	if v, ok := d.GetOk("dscp_reverse_tag"); ok {
		t, err := expandSystemsdwanServiceDscpReverseTag(d, v, "dscp_reverse_tag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-reverse-tag"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok {
		t, err := expandSystemsdwanServiceDst(d, v, "dst", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	} else if d.HasChange("dst") {
		old, new := d.GetChange("dst")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["dst"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("dst_negate"); ok {
		t, err := expandSystemsdwanServiceDstNegate(d, v, "dst_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-negate"] = t
		}
	}

	if v, ok := d.GetOk("dst6"); ok {
		t, err := expandSystemsdwanServiceDst6(d, v, "dst6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst6"] = t
		}
	} else if d.HasChange("dst6") {
		old, new := d.GetChange("dst6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["dst6"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("end_port"); ok {
		t, err := expandSystemsdwanServiceEndPort(d, v, "end_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-port"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok {
		t, err := expandSystemsdwanServiceGateway(d, v, "gateway", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok {
		t, err := expandSystemsdwanServiceGroups(d, v, "groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	} else if d.HasChange("groups") {
		old, new := d.GetChange("groups")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["groups"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("hash_mode"); ok {
		t, err := expandSystemsdwanServiceHashMode(d, v, "hash_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hash-mode"] = t
		}
	}

	if v, ok := d.GetOk("health_check"); ok {
		t, err := expandSystemsdwanServiceHealthCheck(d, v, "health_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check"] = t
		}
	} else if d.HasChange("health_check") {
		old, new := d.GetChange("health_check")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["health-check"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("hold_down_time"); ok {
		t, err := expandSystemsdwanServiceHoldDownTime(d, v, "hold_down_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hold-down-time"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandSystemsdwanServiceId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("input_device"); ok {
		t, err := expandSystemsdwanServiceInputDevice(d, v, "input_device", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["input-device"] = t
		}
	} else if d.HasChange("input_device") {
		old, new := d.GetChange("input_device")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["input-device"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("input_device_negate"); ok {
		t, err := expandSystemsdwanServiceInputDeviceNegate(d, v, "input_device_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["input-device-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service"); ok {
		t, err := expandSystemsdwanServiceInternetService(d, v, "internet_service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_app_ctrl"); ok {
		t, err := expandSystemsdwanServiceInternetServiceAppCtrl(d, v, "internet_service_app_ctrl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-app-ctrl"] = t
		}
	} else if d.HasChange("internet_service_app_ctrl") {
		old, new := d.GetChange("internet_service_app_ctrl")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["internet-service-app-ctrl"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("internet_service_app_ctrl_group"); ok {
		t, err := expandSystemsdwanServiceInternetServiceAppCtrlGroup(d, v, "internet_service_app_ctrl_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-app-ctrl-group"] = t
		}
	} else if d.HasChange("internet_service_app_ctrl_group") {
		old, new := d.GetChange("internet_service_app_ctrl_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["internet-service-app-ctrl-group"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("internet_service_custom"); ok {
		t, err := expandSystemsdwanServiceInternetServiceCustom(d, v, "internet_service_custom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom"] = t
		}
	} else if d.HasChange("internet_service_custom") {
		old, new := d.GetChange("internet_service_custom")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["internet-service-custom"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("internet_service_custom_group"); ok {
		t, err := expandSystemsdwanServiceInternetServiceCustomGroup(d, v, "internet_service_custom_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom-group"] = t
		}
	} else if d.HasChange("internet_service_custom_group") {
		old, new := d.GetChange("internet_service_custom_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["internet-service-custom-group"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("internet_service_group"); ok {
		t, err := expandSystemsdwanServiceInternetServiceGroup(d, v, "internet_service_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-group"] = t
		}
	} else if d.HasChange("internet_service_group") {
		old, new := d.GetChange("internet_service_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["internet-service-group"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("internet_service_name"); ok {
		t, err := expandSystemsdwanServiceInternetServiceName(d, v, "internet_service_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-name"] = t
		}
	} else if d.HasChange("internet_service_name") {
		old, new := d.GetChange("internet_service_name")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["internet-service-name"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("jitter_weight"); ok {
		t, err := expandSystemsdwanServiceJitterWeight(d, v, "jitter_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jitter-weight"] = t
		}
	}

	if v, ok := d.GetOk("latency_weight"); ok {
		t, err := expandSystemsdwanServiceLatencyWeight(d, v, "latency_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["latency-weight"] = t
		}
	}

	if v, ok := d.GetOk("link_cost_factor"); ok {
		t, err := expandSystemsdwanServiceLinkCostFactor(d, v, "link_cost_factor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-cost-factor"] = t
		}
	}

	if v, ok := d.GetOk("link_cost_threshold"); ok {
		t, err := expandSystemsdwanServiceLinkCostThreshold(d, v, "link_cost_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-cost-threshold"] = t
		}
	}

	if v, ok := d.GetOk("minimum_sla_meet_members"); ok {
		t, err := expandSystemsdwanServiceMinimumSlaMeetMembers(d, v, "minimum_sla_meet_members", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minimum-sla-meet-members"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandSystemsdwanServiceMode(d, v, "mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemsdwanServiceName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("packet_loss_weight"); ok {
		t, err := expandSystemsdwanServicePacketLossWeight(d, v, "packet_loss_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-loss-weight"] = t
		}
	}

	if v, ok := d.GetOk("priority_members"); ok {
		t, err := expandSystemsdwanServicePriorityMembers(d, v, "priority_members", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-members"] = t
		}
	} else if d.HasChange("priority_members") {
		old, new := d.GetChange("priority_members")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["priority-members"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("priority_zone"); ok {
		t, err := expandSystemsdwanServicePriorityZone(d, v, "priority_zone", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-zone"] = t
		}
	} else if d.HasChange("priority_zone") {
		old, new := d.GetChange("priority_zone")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["priority-zone"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("protocol"); ok {
		t, err := expandSystemsdwanServiceProtocol(d, v, "protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("quality_link"); ok {
		t, err := expandSystemsdwanServiceQualityLink(d, v, "quality_link", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quality-link"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok {
		t, err := expandSystemsdwanServiceRole(d, v, "role", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("route_tag"); ok {
		t, err := expandSystemsdwanServiceRouteTag(d, v, "route_tag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-tag"] = t
		}
	}

	if v, ok := d.GetOk("sla"); ok {
		t, err := expandSystemsdwanServiceSla(d, v, "sla", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla"] = t
		}
	} else if d.HasChange("sla") {
		old, new := d.GetChange("sla")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["sla"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("sla_compare_method"); ok {
		t, err := expandSystemsdwanServiceSlaCompareMethod(d, v, "sla_compare_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-compare-method"] = t
		}
	}

	if v, ok := d.GetOk("src"); ok {
		t, err := expandSystemsdwanServiceSrc(d, v, "src", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	} else if d.HasChange("src") {
		old, new := d.GetChange("src")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["src"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("src_negate"); ok {
		t, err := expandSystemsdwanServiceSrcNegate(d, v, "src_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-negate"] = t
		}
	}

	if v, ok := d.GetOk("src6"); ok {
		t, err := expandSystemsdwanServiceSrc6(d, v, "src6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src6"] = t
		}
	} else if d.HasChange("src6") {
		old, new := d.GetChange("src6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["src6"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("standalone_action"); ok {
		t, err := expandSystemsdwanServiceStandaloneAction(d, v, "standalone_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["standalone-action"] = t
		}
	}

	if v, ok := d.GetOk("start_port"); ok {
		t, err := expandSystemsdwanServiceStartPort(d, v, "start_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-port"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemsdwanServiceStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tie_break"); ok {
		t, err := expandSystemsdwanServiceTieBreak(d, v, "tie_break", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tie-break"] = t
		}
	}

	if v, ok := d.GetOk("tos"); ok {
		t, err := expandSystemsdwanServiceTos(d, v, "tos", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos"] = t
		}
	}

	if v, ok := d.GetOk("tos_mask"); ok {
		t, err := expandSystemsdwanServiceTosMask(d, v, "tos_mask", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-mask"] = t
		}
	}

	if v, ok := d.GetOk("use_shortcut_sla"); ok {
		t, err := expandSystemsdwanServiceUseShortcutSla(d, v, "use_shortcut_sla", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-shortcut-sla"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok {
		t, err := expandSystemsdwanServiceUsers(d, v, "users", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	} else if d.HasChange("users") {
		old, new := d.GetChange("users")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["users"] = make([]struct{}, 0)
		}
	}

	return &obj, nil
}
