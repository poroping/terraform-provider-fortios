// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: Configure IPv6 OSPF.

package fortios

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceRouterOspf6() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterOspf6Read,
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"abr_type": {
				Type:        schema.TypeString,
				Description: "Area border router type.",
				Computed:    true,
			},
			"area": {
				Type:        schema.TypeList,
				Description: "OSPF6 area configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:        schema.TypeString,
							Description: "Authentication mode.",
							Computed:    true,
						},
						"default_cost": {
							Type:        schema.TypeInt,
							Description: "Summary default cost of stub or NSSA area.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeString,
							Description: "Area entry IP address.",
							Computed:    true,
						},
						"ipsec_auth_alg": {
							Type:        schema.TypeString,
							Description: "Authentication algorithm.",
							Computed:    true,
						},
						"ipsec_enc_alg": {
							Type:        schema.TypeString,
							Description: "Encryption algorithm.",
							Computed:    true,
						},
						"ipsec_keys": {
							Type:        schema.TypeList,
							Description: "IPsec authentication and encryption keys.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_key": {
										Type:        schema.TypeString,
										Description: "Authentication key.",
										Computed:    true,
										Sensitive:   true,
									},
									"enc_key": {
										Type:        schema.TypeString,
										Description: "Encryption key.",
										Computed:    true,
										Sensitive:   true,
									},
									"spi": {
										Type:        schema.TypeInt,
										Description: "Security Parameters Index.",
										Computed:    true,
									},
								},
							},
						},
						"key_rollover_interval": {
							Type:        schema.TypeInt,
							Description: "Key roll-over interval.",
							Computed:    true,
						},
						"nssa_default_information_originate": {
							Type:        schema.TypeString,
							Description: "Enable/disable originate type 7 default into NSSA area.",
							Computed:    true,
						},
						"nssa_default_information_originate_metric": {
							Type:        schema.TypeInt,
							Description: "OSPFv3 default metric.",
							Computed:    true,
						},
						"nssa_default_information_originate_metric_type": {
							Type:        schema.TypeString,
							Description: "OSPFv3 metric type for default routes.",
							Computed:    true,
						},
						"nssa_redistribution": {
							Type:        schema.TypeString,
							Description: "Enable/disable redistribute into NSSA area.",
							Computed:    true,
						},
						"nssa_translator_role": {
							Type:        schema.TypeString,
							Description: "NSSA translator role type.",
							Computed:    true,
						},
						"range": {
							Type:        schema.TypeList,
							Description: "OSPF6 area range configuration.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"advertise": {
										Type:        schema.TypeString,
										Description: "Enable/disable advertise status.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "Range entry ID.",
										Computed:    true,
									},
									"prefix6": {
										Type:        schema.TypeString,
										Description: "IPv6 prefix.",
										Computed:    true,
									},
								},
							},
						},
						"stub_type": {
							Type:        schema.TypeString,
							Description: "Stub summary setting.",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "Area type setting.",
							Computed:    true,
						},
						"virtual_link": {
							Type:        schema.TypeList,
							Description: "OSPF6 virtual link configuration.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication": {
										Type:        schema.TypeString,
										Description: "Authentication mode.",
										Computed:    true,
									},
									"dead_interval": {
										Type:        schema.TypeInt,
										Description: "Dead interval.",
										Computed:    true,
									},
									"hello_interval": {
										Type:        schema.TypeInt,
										Description: "Hello interval.",
										Computed:    true,
									},
									"ipsec_auth_alg": {
										Type:        schema.TypeString,
										Description: "Authentication algorithm.",
										Computed:    true,
									},
									"ipsec_enc_alg": {
										Type:        schema.TypeString,
										Description: "Encryption algorithm.",
										Computed:    true,
									},
									"ipsec_keys": {
										Type:        schema.TypeList,
										Description: "IPsec authentication and encryption keys.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"auth_key": {
													Type:        schema.TypeString,
													Description: "Authentication key.",
													Computed:    true,
													Sensitive:   true,
												},
												"enc_key": {
													Type:        schema.TypeString,
													Description: "Encryption key.",
													Computed:    true,
													Sensitive:   true,
												},
												"spi": {
													Type:        schema.TypeInt,
													Description: "Security Parameters Index.",
													Computed:    true,
												},
											},
										},
									},
									"key_rollover_interval": {
										Type:        schema.TypeInt,
										Description: "Key roll-over interval.",
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "Virtual link entry name.",
										Computed:    true,
									},
									"peer": {
										Type:        schema.TypeString,
										Description: "A.B.C.D, peer router ID.",
										Computed:    true,
									},
									"retransmit_interval": {
										Type:        schema.TypeInt,
										Description: "Retransmit interval.",
										Computed:    true,
									},
									"transmit_delay": {
										Type:        schema.TypeInt,
										Description: "Transmit delay.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"auto_cost_ref_bandwidth": {
				Type:        schema.TypeInt,
				Description: "Reference bandwidth in terms of megabits per second.",
				Computed:    true,
			},
			"bfd": {
				Type:        schema.TypeString,
				Description: "Enable/disable Bidirectional Forwarding Detection (BFD).",
				Computed:    true,
			},
			"default_information_metric": {
				Type:        schema.TypeInt,
				Description: "Default information metric.",
				Computed:    true,
			},
			"default_information_metric_type": {
				Type:        schema.TypeString,
				Description: "Default information metric type.",
				Computed:    true,
			},
			"default_information_originate": {
				Type:        schema.TypeString,
				Description: "Enable/disable generation of default route.",
				Computed:    true,
			},
			"default_information_route_map": {
				Type:        schema.TypeString,
				Description: "Default information route map.",
				Computed:    true,
			},
			"default_metric": {
				Type:        schema.TypeInt,
				Description: "Default metric of redistribute routes.",
				Computed:    true,
			},
			"log_neighbour_changes": {
				Type:        schema.TypeString,
				Description: "Enable logging of OSPFv3 neighbour's changes",
				Computed:    true,
			},
			"ospf6_interface": {
				Type:        schema.TypeList,
				Description: "OSPF6 interface configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"area_id": {
							Type:        schema.TypeString,
							Description: "A.B.C.D, in IPv4 address format.",
							Computed:    true,
						},
						"authentication": {
							Type:        schema.TypeString,
							Description: "Authentication mode.",
							Computed:    true,
						},
						"bfd": {
							Type:        schema.TypeString,
							Description: "Enable/disable Bidirectional Forwarding Detection (BFD).",
							Computed:    true,
						},
						"cost": {
							Type:        schema.TypeInt,
							Description: "Cost of the interface, value range from 0 to 65535, 0 means auto-cost.",
							Computed:    true,
						},
						"dead_interval": {
							Type:        schema.TypeInt,
							Description: "Dead interval.",
							Computed:    true,
						},
						"hello_interval": {
							Type:        schema.TypeInt,
							Description: "Hello interval.",
							Computed:    true,
						},
						"interface": {
							Type:        schema.TypeString,
							Description: "Configuration interface name.",
							Computed:    true,
						},
						"ipsec_auth_alg": {
							Type:        schema.TypeString,
							Description: "Authentication algorithm.",
							Computed:    true,
						},
						"ipsec_enc_alg": {
							Type:        schema.TypeString,
							Description: "Encryption algorithm.",
							Computed:    true,
						},
						"ipsec_keys": {
							Type:        schema.TypeList,
							Description: "IPsec authentication and encryption keys.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_key": {
										Type:        schema.TypeString,
										Description: "Authentication key.",
										Computed:    true,
										Sensitive:   true,
									},
									"enc_key": {
										Type:        schema.TypeString,
										Description: "Encryption key.",
										Computed:    true,
										Sensitive:   true,
									},
									"spi": {
										Type:        schema.TypeInt,
										Description: "Security Parameters Index.",
										Computed:    true,
									},
								},
							},
						},
						"key_rollover_interval": {
							Type:        schema.TypeInt,
							Description: "Key roll-over interval.",
							Computed:    true,
						},
						"mtu": {
							Type:        schema.TypeInt,
							Description: "MTU for OSPFv3 packets.",
							Computed:    true,
						},
						"mtu_ignore": {
							Type:        schema.TypeString,
							Description: "Enable/disable ignoring MTU field in DBD packets.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Interface entry name.",
							Computed:    true,
						},
						"neighbor": {
							Type:        schema.TypeList,
							Description: "OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cost": {
										Type:        schema.TypeInt,
										Description: "Cost of the interface, value range from 0 to 65535, 0 means auto-cost.",
										Computed:    true,
									},
									"ip6": {
										Type:        schema.TypeString,
										Description: "IPv6 link local address of the neighbor.",
										Computed:    true,
									},
									"poll_interval": {
										Type:        schema.TypeInt,
										Description: "Poll interval time in seconds.",
										Computed:    true,
									},
									"priority": {
										Type:        schema.TypeInt,
										Description: "priority",
										Computed:    true,
									},
								},
							},
						},
						"network_type": {
							Type:        schema.TypeString,
							Description: "Network type.",
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "priority",
							Computed:    true,
						},
						"retransmit_interval": {
							Type:        schema.TypeInt,
							Description: "Retransmit interval.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable OSPF6 routing on this interface.",
							Computed:    true,
						},
						"transmit_delay": {
							Type:        schema.TypeInt,
							Description: "Transmit delay.",
							Computed:    true,
						},
					},
				},
			},
			"passive_interface": {
				Type:        schema.TypeList,
				Description: "Passive interface configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Passive interface name.",
							Computed:    true,
						},
					},
				},
			},
			"redistribute": {
				Type:        schema.TypeList,
				Description: "Redistribute configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metric": {
							Type:        schema.TypeInt,
							Description: "Redistribute metric setting.",
							Computed:    true,
						},
						"metric_type": {
							Type:        schema.TypeString,
							Description: "Metric type.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Redistribute name.",
							Computed:    true,
						},
						"routemap": {
							Type:        schema.TypeString,
							Description: "Route map name.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "status",
							Computed:    true,
						},
					},
				},
			},
			"router_id": {
				Type:        schema.TypeString,
				Description: "A.B.C.D, in IPv4 address format.",
				Computed:    true,
			},
			"spf_timers": {
				Type:        schema.TypeString,
				Description: "SPF calculation frequency.",
				Computed:    true,
			},
			"summary_address": {
				Type:        schema.TypeList,
				Description: "IPv6 address summary configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advertise": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertise status.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Summary address entry ID.",
							Computed:    true,
						},
						"prefix6": {
							Type:        schema.TypeString,
							Description: "IPv6 prefix.",
							Computed:    true,
						},
						"tag": {
							Type:        schema.TypeInt,
							Description: "Tag value.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterOspf6Read(d *schema.ResourceData, m interface{}) error {
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
	mkey = "RouterOspf6"

	o, err := c.ReadRouterOspf6(mkey, vdomparam, urlparams, 0)
	if err != nil {
		return fmt.Errorf("error describing RouterOspf6: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterOspf6(d, o)
	if err != nil {
		return fmt.Errorf("error describing RouterOspf6 from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterOspf6AbrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Area(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {

			tmp["authentication"] = dataSourceFlattenRouterOspf6AreaAuthentication(i["authentication"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_cost"
		if _, ok := i["default-cost"]; ok {

			tmp["default_cost"] = dataSourceFlattenRouterOspf6AreaDefaultCost(i["default-cost"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = dataSourceFlattenRouterOspf6AreaId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_auth_alg"
		if _, ok := i["ipsec-auth-alg"]; ok {

			tmp["ipsec_auth_alg"] = dataSourceFlattenRouterOspf6AreaIpsecAuthAlg(i["ipsec-auth-alg"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_enc_alg"
		if _, ok := i["ipsec-enc-alg"]; ok {

			tmp["ipsec_enc_alg"] = dataSourceFlattenRouterOspf6AreaIpsecEncAlg(i["ipsec-enc-alg"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_keys"
		if _, ok := i["ipsec-keys"]; ok {

			tmp["ipsec_keys"] = dataSourceFlattenRouterOspf6AreaIpsecKeys(i["ipsec-keys"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_rollover_interval"
		if _, ok := i["key-rollover-interval"]; ok {

			tmp["key_rollover_interval"] = dataSourceFlattenRouterOspf6AreaKeyRolloverInterval(i["key-rollover-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate"
		if _, ok := i["nssa-default-information-originate"]; ok {

			tmp["nssa_default_information_originate"] = dataSourceFlattenRouterOspf6AreaNssaDefaultInformationOriginate(i["nssa-default-information-originate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric"
		if _, ok := i["nssa-default-information-originate-metric"]; ok {

			tmp["nssa_default_information_originate_metric"] = dataSourceFlattenRouterOspf6AreaNssaDefaultInformationOriginateMetric(i["nssa-default-information-originate-metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric_type"
		if _, ok := i["nssa-default-information-originate-metric-type"]; ok {

			tmp["nssa_default_information_originate_metric_type"] = dataSourceFlattenRouterOspf6AreaNssaDefaultInformationOriginateMetricType(i["nssa-default-information-originate-metric-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_redistribution"
		if _, ok := i["nssa-redistribution"]; ok {

			tmp["nssa_redistribution"] = dataSourceFlattenRouterOspf6AreaNssaRedistribution(i["nssa-redistribution"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_translator_role"
		if _, ok := i["nssa-translator-role"]; ok {

			tmp["nssa_translator_role"] = dataSourceFlattenRouterOspf6AreaNssaTranslatorRole(i["nssa-translator-role"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := i["range"]; ok {

			tmp["range"] = dataSourceFlattenRouterOspf6AreaRange(i["range"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stub_type"
		if _, ok := i["stub-type"]; ok {

			tmp["stub_type"] = dataSourceFlattenRouterOspf6AreaStubType(i["stub-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {

			tmp["type"] = dataSourceFlattenRouterOspf6AreaType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_link"
		if _, ok := i["virtual-link"]; ok {

			tmp["virtual_link"] = dataSourceFlattenRouterOspf6AreaVirtualLink(i["virtual-link"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspf6AreaAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaDefaultCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaIpsecAuthAlg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaIpsecEncAlg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaIpsecKeys(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if _, ok := i["spi"]; ok {

			tmp["spi"] = dataSourceFlattenRouterOspf6AreaIpsecKeysSpi(i["spi"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspf6AreaIpsecKeysSpi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaKeyRolloverInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaNssaDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaNssaDefaultInformationOriginateMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaNssaDefaultInformationOriginateMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaNssaRedistribution(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaNssaTranslatorRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := i["advertise"]; ok {

			tmp["advertise"] = dataSourceFlattenRouterOspf6AreaRangeAdvertise(i["advertise"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = dataSourceFlattenRouterOspf6AreaRangeId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {

			tmp["prefix6"] = dataSourceFlattenRouterOspf6AreaRangePrefix6(i["prefix6"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspf6AreaRangeAdvertise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaRangePrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaStubType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLink(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {

			tmp["authentication"] = dataSourceFlattenRouterOspf6AreaVirtualLinkAuthentication(i["authentication"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := i["dead-interval"]; ok {

			tmp["dead_interval"] = dataSourceFlattenRouterOspf6AreaVirtualLinkDeadInterval(i["dead-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {

			tmp["hello_interval"] = dataSourceFlattenRouterOspf6AreaVirtualLinkHelloInterval(i["hello-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_auth_alg"
		if _, ok := i["ipsec-auth-alg"]; ok {

			tmp["ipsec_auth_alg"] = dataSourceFlattenRouterOspf6AreaVirtualLinkIpsecAuthAlg(i["ipsec-auth-alg"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_enc_alg"
		if _, ok := i["ipsec-enc-alg"]; ok {

			tmp["ipsec_enc_alg"] = dataSourceFlattenRouterOspf6AreaVirtualLinkIpsecEncAlg(i["ipsec-enc-alg"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_keys"
		if _, ok := i["ipsec-keys"]; ok {

			tmp["ipsec_keys"] = dataSourceFlattenRouterOspf6AreaVirtualLinkIpsecKeys(i["ipsec-keys"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_rollover_interval"
		if _, ok := i["key-rollover-interval"]; ok {

			tmp["key_rollover_interval"] = dataSourceFlattenRouterOspf6AreaVirtualLinkKeyRolloverInterval(i["key-rollover-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = dataSourceFlattenRouterOspf6AreaVirtualLinkName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := i["peer"]; ok {

			tmp["peer"] = dataSourceFlattenRouterOspf6AreaVirtualLinkPeer(i["peer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := i["retransmit-interval"]; ok {

			tmp["retransmit_interval"] = dataSourceFlattenRouterOspf6AreaVirtualLinkRetransmitInterval(i["retransmit-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := i["transmit-delay"]; ok {

			tmp["transmit_delay"] = dataSourceFlattenRouterOspf6AreaVirtualLinkTransmitDelay(i["transmit-delay"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkDeadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkIpsecAuthAlg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkIpsecEncAlg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkIpsecKeys(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if _, ok := i["spi"]; ok {

			tmp["spi"] = dataSourceFlattenRouterOspf6AreaVirtualLinkIpsecKeysSpi(i["spi"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkIpsecKeysSpi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkKeyRolloverInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkRetransmitInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkTransmitDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AutoCostRefBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Bfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6DefaultInformationMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6DefaultInformationMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6DefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6DefaultInformationRouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6DefaultMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6LogNeighbourChanges(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6Interface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
		pre_append = pre + "." + strconv.Itoa(con) + "." + "area_id"
		if _, ok := i["area-id"]; ok {

			tmp["area_id"] = dataSourceFlattenRouterOspf6Ospf6InterfaceAreaId(i["area-id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {

			tmp["authentication"] = dataSourceFlattenRouterOspf6Ospf6InterfaceAuthentication(i["authentication"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := i["bfd"]; ok {

			tmp["bfd"] = dataSourceFlattenRouterOspf6Ospf6InterfaceBfd(i["bfd"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {

			tmp["cost"] = dataSourceFlattenRouterOspf6Ospf6InterfaceCost(i["cost"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := i["dead-interval"]; ok {

			tmp["dead_interval"] = dataSourceFlattenRouterOspf6Ospf6InterfaceDeadInterval(i["dead-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {

			tmp["hello_interval"] = dataSourceFlattenRouterOspf6Ospf6InterfaceHelloInterval(i["hello-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {

			tmp["interface"] = dataSourceFlattenRouterOspf6Ospf6InterfaceInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_auth_alg"
		if _, ok := i["ipsec-auth-alg"]; ok {

			tmp["ipsec_auth_alg"] = dataSourceFlattenRouterOspf6Ospf6InterfaceIpsecAuthAlg(i["ipsec-auth-alg"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_enc_alg"
		if _, ok := i["ipsec-enc-alg"]; ok {

			tmp["ipsec_enc_alg"] = dataSourceFlattenRouterOspf6Ospf6InterfaceIpsecEncAlg(i["ipsec-enc-alg"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_keys"
		if _, ok := i["ipsec-keys"]; ok {

			tmp["ipsec_keys"] = dataSourceFlattenRouterOspf6Ospf6InterfaceIpsecKeys(i["ipsec-keys"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_rollover_interval"
		if _, ok := i["key-rollover-interval"]; ok {

			tmp["key_rollover_interval"] = dataSourceFlattenRouterOspf6Ospf6InterfaceKeyRolloverInterval(i["key-rollover-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu"
		if _, ok := i["mtu"]; ok {

			tmp["mtu"] = dataSourceFlattenRouterOspf6Ospf6InterfaceMtu(i["mtu"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu_ignore"
		if _, ok := i["mtu-ignore"]; ok {

			tmp["mtu_ignore"] = dataSourceFlattenRouterOspf6Ospf6InterfaceMtuIgnore(i["mtu-ignore"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = dataSourceFlattenRouterOspf6Ospf6InterfaceName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor"
		if _, ok := i["neighbor"]; ok {

			tmp["neighbor"] = dataSourceFlattenRouterOspf6Ospf6InterfaceNeighbor(i["neighbor"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_type"
		if _, ok := i["network-type"]; ok {

			tmp["network_type"] = dataSourceFlattenRouterOspf6Ospf6InterfaceNetworkType(i["network-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {

			tmp["priority"] = dataSourceFlattenRouterOspf6Ospf6InterfacePriority(i["priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := i["retransmit-interval"]; ok {

			tmp["retransmit_interval"] = dataSourceFlattenRouterOspf6Ospf6InterfaceRetransmitInterval(i["retransmit-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = dataSourceFlattenRouterOspf6Ospf6InterfaceStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := i["transmit-delay"]; ok {

			tmp["transmit_delay"] = dataSourceFlattenRouterOspf6Ospf6InterfaceTransmitDelay(i["transmit-delay"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceAreaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceDeadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceIpsecAuthAlg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceIpsecEncAlg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceIpsecKeys(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if _, ok := i["spi"]; ok {

			tmp["spi"] = dataSourceFlattenRouterOspf6Ospf6InterfaceIpsecKeysSpi(i["spi"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceIpsecKeysSpi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceKeyRolloverInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceMtuIgnore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceNeighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {

			tmp["cost"] = dataSourceFlattenRouterOspf6Ospf6InterfaceNeighborCost(i["cost"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {

			tmp["ip6"] = dataSourceFlattenRouterOspf6Ospf6InterfaceNeighborIp6(i["ip6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poll_interval"
		if _, ok := i["poll-interval"]; ok {

			tmp["poll_interval"] = dataSourceFlattenRouterOspf6Ospf6InterfaceNeighborPollInterval(i["poll-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {

			tmp["priority"] = dataSourceFlattenRouterOspf6Ospf6InterfaceNeighborPriority(i["priority"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceNeighborCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceNeighborIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceNeighborPollInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceNeighborPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceNetworkType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfacePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceRetransmitInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceTransmitDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6PassiveInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenRouterOspf6PassiveInterfaceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspf6PassiveInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Redistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := i["metric"]; ok {

			tmp["metric"] = dataSourceFlattenRouterOspf6RedistributeMetric(i["metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_type"
		if _, ok := i["metric-type"]; ok {

			tmp["metric_type"] = dataSourceFlattenRouterOspf6RedistributeMetricType(i["metric-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = dataSourceFlattenRouterOspf6RedistributeName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := i["routemap"]; ok {

			tmp["routemap"] = dataSourceFlattenRouterOspf6RedistributeRoutemap(i["routemap"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = dataSourceFlattenRouterOspf6RedistributeStatus(i["status"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspf6RedistributeMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6RedistributeMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6RedistributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6RedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6RedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6RouterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6SpfTimers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6SummaryAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := i["advertise"]; ok {

			tmp["advertise"] = dataSourceFlattenRouterOspf6SummaryAddressAdvertise(i["advertise"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = dataSourceFlattenRouterOspf6SummaryAddressId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {

			tmp["prefix6"] = dataSourceFlattenRouterOspf6SummaryAddressPrefix6(i["prefix6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if _, ok := i["tag"]; ok {

			tmp["tag"] = dataSourceFlattenRouterOspf6SummaryAddressTag(i["tag"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspf6SummaryAddressAdvertise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6SummaryAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6SummaryAddressPrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6SummaryAddressTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterOspf6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("abr_type", dataSourceFlattenRouterOspf6AbrType(o["abr-type"], d, "abr_type")); err != nil {
		if !fortiAPIPatch(o["abr-type"]) {
			return fmt.Errorf("error reading abr_type: %v", err)
		}
	}

	if err = d.Set("area", dataSourceFlattenRouterOspf6Area(o["area"], d, "area")); err != nil {
		if !fortiAPIPatch(o["area"]) {
			return fmt.Errorf("error reading area: %v", err)
		}
	}

	if err = d.Set("auto_cost_ref_bandwidth", dataSourceFlattenRouterOspf6AutoCostRefBandwidth(o["auto-cost-ref-bandwidth"], d, "auto_cost_ref_bandwidth")); err != nil {
		if !fortiAPIPatch(o["auto-cost-ref-bandwidth"]) {
			return fmt.Errorf("error reading auto_cost_ref_bandwidth: %v", err)
		}
	}

	if err = d.Set("bfd", dataSourceFlattenRouterOspf6Bfd(o["bfd"], d, "bfd")); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("error reading bfd: %v", err)
		}
	}

	if err = d.Set("default_information_metric", dataSourceFlattenRouterOspf6DefaultInformationMetric(o["default-information-metric"], d, "default_information_metric")); err != nil {
		if !fortiAPIPatch(o["default-information-metric"]) {
			return fmt.Errorf("error reading default_information_metric: %v", err)
		}
	}

	if err = d.Set("default_information_metric_type", dataSourceFlattenRouterOspf6DefaultInformationMetricType(o["default-information-metric-type"], d, "default_information_metric_type")); err != nil {
		if !fortiAPIPatch(o["default-information-metric-type"]) {
			return fmt.Errorf("error reading default_information_metric_type: %v", err)
		}
	}

	if err = d.Set("default_information_originate", dataSourceFlattenRouterOspf6DefaultInformationOriginate(o["default-information-originate"], d, "default_information_originate")); err != nil {
		if !fortiAPIPatch(o["default-information-originate"]) {
			return fmt.Errorf("error reading default_information_originate: %v", err)
		}
	}

	if err = d.Set("default_information_route_map", dataSourceFlattenRouterOspf6DefaultInformationRouteMap(o["default-information-route-map"], d, "default_information_route_map")); err != nil {
		if !fortiAPIPatch(o["default-information-route-map"]) {
			return fmt.Errorf("error reading default_information_route_map: %v", err)
		}
	}

	if err = d.Set("default_metric", dataSourceFlattenRouterOspf6DefaultMetric(o["default-metric"], d, "default_metric")); err != nil {
		if !fortiAPIPatch(o["default-metric"]) {
			return fmt.Errorf("error reading default_metric: %v", err)
		}
	}

	if err = d.Set("log_neighbour_changes", dataSourceFlattenRouterOspf6LogNeighbourChanges(o["log-neighbour-changes"], d, "log_neighbour_changes")); err != nil {
		if !fortiAPIPatch(o["log-neighbour-changes"]) {
			return fmt.Errorf("error reading log_neighbour_changes: %v", err)
		}
	}

	if err = d.Set("ospf6_interface", dataSourceFlattenRouterOspf6Ospf6Interface(o["ospf6-interface"], d, "ospf6_interface")); err != nil {
		if !fortiAPIPatch(o["ospf6-interface"]) {
			return fmt.Errorf("error reading ospf6_interface: %v", err)
		}
	}

	if err = d.Set("passive_interface", dataSourceFlattenRouterOspf6PassiveInterface(o["passive-interface"], d, "passive_interface")); err != nil {
		if !fortiAPIPatch(o["passive-interface"]) {
			return fmt.Errorf("error reading passive_interface: %v", err)
		}
	}

	if err = d.Set("redistribute", dataSourceFlattenRouterOspf6Redistribute(o["redistribute"], d, "redistribute")); err != nil {
		if !fortiAPIPatch(o["redistribute"]) {
			return fmt.Errorf("error reading redistribute: %v", err)
		}
	}

	if err = d.Set("router_id", dataSourceFlattenRouterOspf6RouterId(o["router-id"], d, "router_id")); err != nil {
		if !fortiAPIPatch(o["router-id"]) {
			return fmt.Errorf("error reading router_id: %v", err)
		}
	}

	if err = d.Set("spf_timers", dataSourceFlattenRouterOspf6SpfTimers(o["spf-timers"], d, "spf_timers")); err != nil {
		if !fortiAPIPatch(o["spf-timers"]) {
			return fmt.Errorf("error reading spf_timers: %v", err)
		}
	}

	if err = d.Set("summary_address", dataSourceFlattenRouterOspf6SummaryAddress(o["summary-address"], d, "summary_address")); err != nil {
		if !fortiAPIPatch(o["summary-address"]) {
			return fmt.Errorf("error reading summary_address: %v", err)
		}
	}

	return nil
}
