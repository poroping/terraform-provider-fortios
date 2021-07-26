// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN.

package fortios

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSystemsdwanService() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemsdwanServiceRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"addr_mode": {
				Type:        schema.TypeString,
				Description: "Address mode (IPv4 or IPv6).",
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

func dataSourceSystemsdwanServiceRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""
	key := "id"
	t := d.Get(key)
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("error describing SystemsdwanService: type error")
	}

	o, err := c.ReadSystemsdwanService(mkey, vdomparam, urlparams, 0)
	if err != nil {
		return fmt.Errorf("error describing SystemsdwanService: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemsdwanService(d, o)
	if err != nil {
		return fmt.Errorf("error describing SystemsdwanService from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemsdwanServiceAddrMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceBandwidthWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceDefault(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceDscpForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceDscpForwardTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceDscpReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceDscpReverseTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceDst(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenSystemsdwanServiceDstName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemsdwanServiceDstName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceDstNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceDst6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenSystemsdwanServiceDst6Name(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemsdwanServiceDst6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceEndPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceGroups(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenSystemsdwanServiceGroupsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemsdwanServiceGroupsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceHashMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceHealthCheck(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenSystemsdwanServiceHealthCheckName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemsdwanServiceHealthCheckName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceHoldDownTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceInputDevice(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenSystemsdwanServiceInputDeviceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemsdwanServiceInputDeviceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceInputDeviceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceInternetService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceInternetServiceAppCtrl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["id"] = dataSourceFlattenSystemsdwanServiceInternetServiceAppCtrlId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemsdwanServiceInternetServiceAppCtrlId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceInternetServiceAppCtrlGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenSystemsdwanServiceInternetServiceAppCtrlGroupName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemsdwanServiceInternetServiceAppCtrlGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenSystemsdwanServiceInternetServiceCustomName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemsdwanServiceInternetServiceCustomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceInternetServiceCustomGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenSystemsdwanServiceInternetServiceCustomGroupName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemsdwanServiceInternetServiceCustomGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceInternetServiceGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenSystemsdwanServiceInternetServiceGroupName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemsdwanServiceInternetServiceGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceInternetServiceName(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenSystemsdwanServiceInternetServiceNameName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemsdwanServiceInternetServiceNameName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceJitterWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceLatencyWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceLinkCostFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceLinkCostThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceMinimumSlaMeetMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServicePacketLossWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServicePriorityMembers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["seq_num"] = dataSourceFlattenSystemsdwanServicePriorityMembersSeqNum(i["seq-num"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemsdwanServicePriorityMembersSeqNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServicePriorityZone(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenSystemsdwanServicePriorityZoneName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemsdwanServicePriorityZoneName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceQualityLink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceRouteTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceSla(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["health_check"] = dataSourceFlattenSystemsdwanServiceSlaHealthCheck(i["health-check"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = dataSourceFlattenSystemsdwanServiceSlaId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemsdwanServiceSlaHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceSlaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceSlaCompareMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceSrc(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenSystemsdwanServiceSrcName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemsdwanServiceSrcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceSrcNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceSrc6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenSystemsdwanServiceSrc6Name(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemsdwanServiceSrc6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceStandaloneAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceStartPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceTieBreak(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceTos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceTosMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceUseShortcutSla(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanServiceUsers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenSystemsdwanServiceUsersName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemsdwanServiceUsersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemsdwanService(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("addr_mode", dataSourceFlattenSystemsdwanServiceAddrMode(o["addr-mode"], d, "addr_mode")); err != nil {
		if !fortiAPIPatch(o["addr-mode"]) {
			return fmt.Errorf("error reading addr_mode: %v", err)
		}
	}

	if err = d.Set("bandwidth_weight", dataSourceFlattenSystemsdwanServiceBandwidthWeight(o["bandwidth-weight"], d, "bandwidth_weight")); err != nil {
		if !fortiAPIPatch(o["bandwidth-weight"]) {
			return fmt.Errorf("error reading bandwidth_weight: %v", err)
		}
	}

	if err = d.Set("default", dataSourceFlattenSystemsdwanServiceDefault(o["default"], d, "default")); err != nil {
		if !fortiAPIPatch(o["default"]) {
			return fmt.Errorf("error reading default: %v", err)
		}
	}

	if err = d.Set("dscp_forward", dataSourceFlattenSystemsdwanServiceDscpForward(o["dscp-forward"], d, "dscp_forward")); err != nil {
		if !fortiAPIPatch(o["dscp-forward"]) {
			return fmt.Errorf("error reading dscp_forward: %v", err)
		}
	}

	if err = d.Set("dscp_forward_tag", dataSourceFlattenSystemsdwanServiceDscpForwardTag(o["dscp-forward-tag"], d, "dscp_forward_tag")); err != nil {
		if !fortiAPIPatch(o["dscp-forward-tag"]) {
			return fmt.Errorf("error reading dscp_forward_tag: %v", err)
		}
	}

	if err = d.Set("dscp_reverse", dataSourceFlattenSystemsdwanServiceDscpReverse(o["dscp-reverse"], d, "dscp_reverse")); err != nil {
		if !fortiAPIPatch(o["dscp-reverse"]) {
			return fmt.Errorf("error reading dscp_reverse: %v", err)
		}
	}

	if err = d.Set("dscp_reverse_tag", dataSourceFlattenSystemsdwanServiceDscpReverseTag(o["dscp-reverse-tag"], d, "dscp_reverse_tag")); err != nil {
		if !fortiAPIPatch(o["dscp-reverse-tag"]) {
			return fmt.Errorf("error reading dscp_reverse_tag: %v", err)
		}
	}

	if err = d.Set("dst", dataSourceFlattenSystemsdwanServiceDst(o["dst"], d, "dst")); err != nil {
		if !fortiAPIPatch(o["dst"]) {
			return fmt.Errorf("error reading dst: %v", err)
		}
	}

	if err = d.Set("dst_negate", dataSourceFlattenSystemsdwanServiceDstNegate(o["dst-negate"], d, "dst_negate")); err != nil {
		if !fortiAPIPatch(o["dst-negate"]) {
			return fmt.Errorf("error reading dst_negate: %v", err)
		}
	}

	if err = d.Set("dst6", dataSourceFlattenSystemsdwanServiceDst6(o["dst6"], d, "dst6")); err != nil {
		if !fortiAPIPatch(o["dst6"]) {
			return fmt.Errorf("error reading dst6: %v", err)
		}
	}

	if err = d.Set("end_port", dataSourceFlattenSystemsdwanServiceEndPort(o["end-port"], d, "end_port")); err != nil {
		if !fortiAPIPatch(o["end-port"]) {
			return fmt.Errorf("error reading end_port: %v", err)
		}
	}

	if err = d.Set("gateway", dataSourceFlattenSystemsdwanServiceGateway(o["gateway"], d, "gateway")); err != nil {
		if !fortiAPIPatch(o["gateway"]) {
			return fmt.Errorf("error reading gateway: %v", err)
		}
	}

	if err = d.Set("groups", dataSourceFlattenSystemsdwanServiceGroups(o["groups"], d, "groups")); err != nil {
		if !fortiAPIPatch(o["groups"]) {
			return fmt.Errorf("error reading groups: %v", err)
		}
	}

	if err = d.Set("hash_mode", dataSourceFlattenSystemsdwanServiceHashMode(o["hash-mode"], d, "hash_mode")); err != nil {
		if !fortiAPIPatch(o["hash-mode"]) {
			return fmt.Errorf("error reading hash_mode: %v", err)
		}
	}

	if err = d.Set("health_check", dataSourceFlattenSystemsdwanServiceHealthCheck(o["health-check"], d, "health_check")); err != nil {
		if !fortiAPIPatch(o["health-check"]) {
			return fmt.Errorf("error reading health_check: %v", err)
		}
	}

	if err = d.Set("hold_down_time", dataSourceFlattenSystemsdwanServiceHoldDownTime(o["hold-down-time"], d, "hold_down_time")); err != nil {
		if !fortiAPIPatch(o["hold-down-time"]) {
			return fmt.Errorf("error reading hold_down_time: %v", err)
		}
	}

	if err = d.Set("fosid", dataSourceFlattenSystemsdwanServiceId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("error reading fosid: %v", err)
		}
	}

	if err = d.Set("input_device", dataSourceFlattenSystemsdwanServiceInputDevice(o["input-device"], d, "input_device")); err != nil {
		if !fortiAPIPatch(o["input-device"]) {
			return fmt.Errorf("error reading input_device: %v", err)
		}
	}

	if err = d.Set("input_device_negate", dataSourceFlattenSystemsdwanServiceInputDeviceNegate(o["input-device-negate"], d, "input_device_negate")); err != nil {
		if !fortiAPIPatch(o["input-device-negate"]) {
			return fmt.Errorf("error reading input_device_negate: %v", err)
		}
	}

	if err = d.Set("internet_service", dataSourceFlattenSystemsdwanServiceInternetService(o["internet-service"], d, "internet_service")); err != nil {
		if !fortiAPIPatch(o["internet-service"]) {
			return fmt.Errorf("error reading internet_service: %v", err)
		}
	}

	if err = d.Set("internet_service_app_ctrl", dataSourceFlattenSystemsdwanServiceInternetServiceAppCtrl(o["internet-service-app-ctrl"], d, "internet_service_app_ctrl")); err != nil {
		if !fortiAPIPatch(o["internet-service-app-ctrl"]) {
			return fmt.Errorf("error reading internet_service_app_ctrl: %v", err)
		}
	}

	if err = d.Set("internet_service_app_ctrl_group", dataSourceFlattenSystemsdwanServiceInternetServiceAppCtrlGroup(o["internet-service-app-ctrl-group"], d, "internet_service_app_ctrl_group")); err != nil {
		if !fortiAPIPatch(o["internet-service-app-ctrl-group"]) {
			return fmt.Errorf("error reading internet_service_app_ctrl_group: %v", err)
		}
	}

	if err = d.Set("internet_service_custom", dataSourceFlattenSystemsdwanServiceInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom")); err != nil {
		if !fortiAPIPatch(o["internet-service-custom"]) {
			return fmt.Errorf("error reading internet_service_custom: %v", err)
		}
	}

	if err = d.Set("internet_service_custom_group", dataSourceFlattenSystemsdwanServiceInternetServiceCustomGroup(o["internet-service-custom-group"], d, "internet_service_custom_group")); err != nil {
		if !fortiAPIPatch(o["internet-service-custom-group"]) {
			return fmt.Errorf("error reading internet_service_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service_group", dataSourceFlattenSystemsdwanServiceInternetServiceGroup(o["internet-service-group"], d, "internet_service_group")); err != nil {
		if !fortiAPIPatch(o["internet-service-group"]) {
			return fmt.Errorf("error reading internet_service_group: %v", err)
		}
	}

	if err = d.Set("internet_service_name", dataSourceFlattenSystemsdwanServiceInternetServiceName(o["internet-service-name"], d, "internet_service_name")); err != nil {
		if !fortiAPIPatch(o["internet-service-name"]) {
			return fmt.Errorf("error reading internet_service_name: %v", err)
		}
	}

	if err = d.Set("jitter_weight", dataSourceFlattenSystemsdwanServiceJitterWeight(o["jitter-weight"], d, "jitter_weight")); err != nil {
		if !fortiAPIPatch(o["jitter-weight"]) {
			return fmt.Errorf("error reading jitter_weight: %v", err)
		}
	}

	if err = d.Set("latency_weight", dataSourceFlattenSystemsdwanServiceLatencyWeight(o["latency-weight"], d, "latency_weight")); err != nil {
		if !fortiAPIPatch(o["latency-weight"]) {
			return fmt.Errorf("error reading latency_weight: %v", err)
		}
	}

	if err = d.Set("link_cost_factor", dataSourceFlattenSystemsdwanServiceLinkCostFactor(o["link-cost-factor"], d, "link_cost_factor")); err != nil {
		if !fortiAPIPatch(o["link-cost-factor"]) {
			return fmt.Errorf("error reading link_cost_factor: %v", err)
		}
	}

	if err = d.Set("link_cost_threshold", dataSourceFlattenSystemsdwanServiceLinkCostThreshold(o["link-cost-threshold"], d, "link_cost_threshold")); err != nil {
		if !fortiAPIPatch(o["link-cost-threshold"]) {
			return fmt.Errorf("error reading link_cost_threshold: %v", err)
		}
	}

	if err = d.Set("minimum_sla_meet_members", dataSourceFlattenSystemsdwanServiceMinimumSlaMeetMembers(o["minimum-sla-meet-members"], d, "minimum_sla_meet_members")); err != nil {
		if !fortiAPIPatch(o["minimum-sla-meet-members"]) {
			return fmt.Errorf("error reading minimum_sla_meet_members: %v", err)
		}
	}

	if err = d.Set("mode", dataSourceFlattenSystemsdwanServiceMode(o["mode"], d, "mode")); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("error reading mode: %v", err)
		}
	}

	if err = d.Set("name", dataSourceFlattenSystemsdwanServiceName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("packet_loss_weight", dataSourceFlattenSystemsdwanServicePacketLossWeight(o["packet-loss-weight"], d, "packet_loss_weight")); err != nil {
		if !fortiAPIPatch(o["packet-loss-weight"]) {
			return fmt.Errorf("error reading packet_loss_weight: %v", err)
		}
	}

	if err = d.Set("priority_members", dataSourceFlattenSystemsdwanServicePriorityMembers(o["priority-members"], d, "priority_members")); err != nil {
		if !fortiAPIPatch(o["priority-members"]) {
			return fmt.Errorf("error reading priority_members: %v", err)
		}
	}

	if err = d.Set("priority_zone", dataSourceFlattenSystemsdwanServicePriorityZone(o["priority-zone"], d, "priority_zone")); err != nil {
		if !fortiAPIPatch(o["priority-zone"]) {
			return fmt.Errorf("error reading priority_zone: %v", err)
		}
	}

	if err = d.Set("protocol", dataSourceFlattenSystemsdwanServiceProtocol(o["protocol"], d, "protocol")); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("error reading protocol: %v", err)
		}
	}

	if err = d.Set("quality_link", dataSourceFlattenSystemsdwanServiceQualityLink(o["quality-link"], d, "quality_link")); err != nil {
		if !fortiAPIPatch(o["quality-link"]) {
			return fmt.Errorf("error reading quality_link: %v", err)
		}
	}

	if err = d.Set("role", dataSourceFlattenSystemsdwanServiceRole(o["role"], d, "role")); err != nil {
		if !fortiAPIPatch(o["role"]) {
			return fmt.Errorf("error reading role: %v", err)
		}
	}

	if err = d.Set("route_tag", dataSourceFlattenSystemsdwanServiceRouteTag(o["route-tag"], d, "route_tag")); err != nil {
		if !fortiAPIPatch(o["route-tag"]) {
			return fmt.Errorf("error reading route_tag: %v", err)
		}
	}

	if err = d.Set("sla", dataSourceFlattenSystemsdwanServiceSla(o["sla"], d, "sla")); err != nil {
		if !fortiAPIPatch(o["sla"]) {
			return fmt.Errorf("error reading sla: %v", err)
		}
	}

	if err = d.Set("sla_compare_method", dataSourceFlattenSystemsdwanServiceSlaCompareMethod(o["sla-compare-method"], d, "sla_compare_method")); err != nil {
		if !fortiAPIPatch(o["sla-compare-method"]) {
			return fmt.Errorf("error reading sla_compare_method: %v", err)
		}
	}

	if err = d.Set("src", dataSourceFlattenSystemsdwanServiceSrc(o["src"], d, "src")); err != nil {
		if !fortiAPIPatch(o["src"]) {
			return fmt.Errorf("error reading src: %v", err)
		}
	}

	if err = d.Set("src_negate", dataSourceFlattenSystemsdwanServiceSrcNegate(o["src-negate"], d, "src_negate")); err != nil {
		if !fortiAPIPatch(o["src-negate"]) {
			return fmt.Errorf("error reading src_negate: %v", err)
		}
	}

	if err = d.Set("src6", dataSourceFlattenSystemsdwanServiceSrc6(o["src6"], d, "src6")); err != nil {
		if !fortiAPIPatch(o["src6"]) {
			return fmt.Errorf("error reading src6: %v", err)
		}
	}

	if err = d.Set("standalone_action", dataSourceFlattenSystemsdwanServiceStandaloneAction(o["standalone-action"], d, "standalone_action")); err != nil {
		if !fortiAPIPatch(o["standalone-action"]) {
			return fmt.Errorf("error reading standalone_action: %v", err)
		}
	}

	if err = d.Set("start_port", dataSourceFlattenSystemsdwanServiceStartPort(o["start-port"], d, "start_port")); err != nil {
		if !fortiAPIPatch(o["start-port"]) {
			return fmt.Errorf("error reading start_port: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenSystemsdwanServiceStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("error reading status: %v", err)
		}
	}

	if err = d.Set("tie_break", dataSourceFlattenSystemsdwanServiceTieBreak(o["tie-break"], d, "tie_break")); err != nil {
		if !fortiAPIPatch(o["tie-break"]) {
			return fmt.Errorf("error reading tie_break: %v", err)
		}
	}

	if err = d.Set("tos", dataSourceFlattenSystemsdwanServiceTos(o["tos"], d, "tos")); err != nil {
		if !fortiAPIPatch(o["tos"]) {
			return fmt.Errorf("error reading tos: %v", err)
		}
	}

	if err = d.Set("tos_mask", dataSourceFlattenSystemsdwanServiceTosMask(o["tos-mask"], d, "tos_mask")); err != nil {
		if !fortiAPIPatch(o["tos-mask"]) {
			return fmt.Errorf("error reading tos_mask: %v", err)
		}
	}

	if err = d.Set("use_shortcut_sla", dataSourceFlattenSystemsdwanServiceUseShortcutSla(o["use-shortcut-sla"], d, "use_shortcut_sla")); err != nil {
		if !fortiAPIPatch(o["use-shortcut-sla"]) {
			return fmt.Errorf("error reading use_shortcut_sla: %v", err)
		}
	}

	if err = d.Set("users", dataSourceFlattenSystemsdwanServiceUsers(o["users"], d, "users")); err != nil {
		if !fortiAPIPatch(o["users"]) {
			return fmt.Errorf("error reading users: %v", err)
		}
	}

	return nil
}
