// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: OSPF area configuration.

package fortios

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceRouterospfArea() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterospfAreaRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"authentication": {
				Type:        schema.TypeString,
				Description: "Authentication type.",
				Computed:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"default_cost": {
				Type:        schema.TypeInt,
				Description: "Summary default cost of stub or NSSA area.",
				Computed:    true,
			},
			"filter_list": {
				Type:        schema.TypeList,
				Description: "OSPF area filter-list configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"direction": {
							Type:        schema.TypeString,
							Description: "Direction.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Filter list entry ID.",
							Computed:    true,
						},
						"list": {
							Type:        schema.TypeString,
							Description: "Access-list or prefix-list name.",
							Computed:    true,
						},
					},
				},
			},
			"fosid": {
				Type:        schema.TypeString,
				Description: "Area entry IP address.",
				Computed:    true,
			},
			"nssa_default_information_originate": {
				Type:        schema.TypeString,
				Description: "Redistribute, advertise, or do not originate Type-7 default route into NSSA area.",
				Computed:    true,
			},
			"nssa_default_information_originate_metric": {
				Type:        schema.TypeInt,
				Description: "OSPF default metric.",
				Computed:    true,
			},
			"nssa_default_information_originate_metric_type": {
				Type:        schema.TypeString,
				Description: "OSPF metric type for default routes.",
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
				Description: "OSPF area range configuration.",
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
						"prefix": {
							Type:        schema.TypeString,
							Description: "Prefix.",
							Computed:    true,
						},
						"substitute": {
							Type:        schema.TypeString,
							Description: "Substitute prefix.",
							Computed:    true,
						},
						"substitute_status": {
							Type:        schema.TypeString,
							Description: "Enable/disable substitute status.",
							Computed:    true,
						},
					},
				},
			},
			"shortcut": {
				Type:        schema.TypeString,
				Description: "Enable/disable shortcut option.",
				Computed:    true,
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
				Description: "OSPF virtual link configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:        schema.TypeString,
							Description: "Authentication type.",
							Computed:    true,
						},
						"authentication_key": {
							Type:        schema.TypeString,
							Description: "Authentication key.",
							Computed:    true,
							Sensitive:   true,
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
						"keychain": {
							Type:        schema.TypeString,
							Description: "Message-digest key-chain name.",
							Computed:    true,
						},
						"md5_keys": {
							Type:        schema.TypeList,
							Description: "MD5 key.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Description: "Key ID (1 - 255).",
										Computed:    true,
									},
									"key_string": {
										Type:        schema.TypeString,
										Description: "Password for the key.",
										Computed:    true,
										Sensitive:   true,
									},
								},
							},
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Virtual link entry name.",
							Computed:    true,
						},
						"peer": {
							Type:        schema.TypeString,
							Description: "Peer IP.",
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
	}
}

func dataSourceRouterospfAreaRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("error describing RouterospfArea: type error")
	}

	o, err := c.ReadRouterospfArea(mkey, vdomparam, urlparams, 0)
	if err != nil {
		return fmt.Errorf("error describing RouterospfArea: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterospfArea(d, o)
	if err != nil {
		return fmt.Errorf("error describing RouterospfArea from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterospfAreaAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaDefaultCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaFilterList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {

			tmp["direction"] = dataSourceFlattenRouterospfAreaFilterListDirection(i["direction"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = dataSourceFlattenRouterospfAreaFilterListId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "list"
		if _, ok := i["list"]; ok {

			tmp["list"] = dataSourceFlattenRouterospfAreaFilterListList(i["list"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterospfAreaFilterListDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaFilterListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaFilterListList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaNssaDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaNssaDefaultInformationOriginateMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaNssaDefaultInformationOriginateMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaNssaRedistribution(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaNssaTranslatorRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["advertise"] = dataSourceFlattenRouterospfAreaRangeAdvertise(i["advertise"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = dataSourceFlattenRouterospfAreaRangeId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {

			tmp["prefix"] = dataSourceFlattenRouterospfAreaRangePrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute"
		if _, ok := i["substitute"]; ok {

			tmp["substitute"] = dataSourceFlattenRouterospfAreaRangeSubstitute(i["substitute"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute_status"
		if _, ok := i["substitute-status"]; ok {

			tmp["substitute_status"] = dataSourceFlattenRouterospfAreaRangeSubstituteStatus(i["substitute-status"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterospfAreaRangeAdvertise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaRangePrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}
	return v
}

func dataSourceFlattenRouterospfAreaRangeSubstitute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}
	return v
}

func dataSourceFlattenRouterospfAreaRangeSubstituteStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaShortcut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaStubType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaVirtualLink(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["authentication"] = dataSourceFlattenRouterospfAreaVirtualLinkAuthentication(i["authentication"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := i["dead-interval"]; ok {

			tmp["dead_interval"] = dataSourceFlattenRouterospfAreaVirtualLinkDeadInterval(i["dead-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {

			tmp["hello_interval"] = dataSourceFlattenRouterospfAreaVirtualLinkHelloInterval(i["hello-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keychain"
		if _, ok := i["keychain"]; ok {

			tmp["keychain"] = dataSourceFlattenRouterospfAreaVirtualLinkKeychain(i["keychain"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keys"
		if _, ok := i["md5-keys"]; ok {

			tmp["md5_keys"] = dataSourceFlattenRouterospfAreaVirtualLinkMd5Keys(i["md5-keys"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = dataSourceFlattenRouterospfAreaVirtualLinkName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := i["peer"]; ok {

			tmp["peer"] = dataSourceFlattenRouterospfAreaVirtualLinkPeer(i["peer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := i["retransmit-interval"]; ok {

			tmp["retransmit_interval"] = dataSourceFlattenRouterospfAreaVirtualLinkRetransmitInterval(i["retransmit-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := i["transmit-delay"]; ok {

			tmp["transmit_delay"] = dataSourceFlattenRouterospfAreaVirtualLinkTransmitDelay(i["transmit-delay"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterospfAreaVirtualLinkAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaVirtualLinkDeadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaVirtualLinkHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaVirtualLinkKeychain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaVirtualLinkMd5Keys(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["id"] = dataSourceFlattenRouterospfAreaVirtualLinkMd5KeysId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterospfAreaVirtualLinkMd5KeysId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaVirtualLinkName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaVirtualLinkPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaVirtualLinkRetransmitInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterospfAreaVirtualLinkTransmitDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterospfArea(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("authentication", dataSourceFlattenRouterospfAreaAuthentication(o["authentication"], d, "authentication")); err != nil {
		if !fortiAPIPatch(o["authentication"]) {
			return fmt.Errorf("error reading authentication: %v", err)
		}
	}

	if err = d.Set("comments", dataSourceFlattenRouterospfAreaComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("error reading comments: %v", err)
		}
	}

	if err = d.Set("default_cost", dataSourceFlattenRouterospfAreaDefaultCost(o["default-cost"], d, "default_cost")); err != nil {
		if !fortiAPIPatch(o["default-cost"]) {
			return fmt.Errorf("error reading default_cost: %v", err)
		}
	}

	if err = d.Set("filter_list", dataSourceFlattenRouterospfAreaFilterList(o["filter-list"], d, "filter_list")); err != nil {
		if !fortiAPIPatch(o["filter-list"]) {
			return fmt.Errorf("error reading filter_list: %v", err)
		}
	}

	if err = d.Set("fosid", dataSourceFlattenRouterospfAreaId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("error reading fosid: %v", err)
		}
	}

	if err = d.Set("nssa_default_information_originate", dataSourceFlattenRouterospfAreaNssaDefaultInformationOriginate(o["nssa-default-information-originate"], d, "nssa_default_information_originate")); err != nil {
		if !fortiAPIPatch(o["nssa-default-information-originate"]) {
			return fmt.Errorf("error reading nssa_default_information_originate: %v", err)
		}
	}

	if err = d.Set("nssa_default_information_originate_metric", dataSourceFlattenRouterospfAreaNssaDefaultInformationOriginateMetric(o["nssa-default-information-originate-metric"], d, "nssa_default_information_originate_metric")); err != nil {
		if !fortiAPIPatch(o["nssa-default-information-originate-metric"]) {
			return fmt.Errorf("error reading nssa_default_information_originate_metric: %v", err)
		}
	}

	if err = d.Set("nssa_default_information_originate_metric_type", dataSourceFlattenRouterospfAreaNssaDefaultInformationOriginateMetricType(o["nssa-default-information-originate-metric-type"], d, "nssa_default_information_originate_metric_type")); err != nil {
		if !fortiAPIPatch(o["nssa-default-information-originate-metric-type"]) {
			return fmt.Errorf("error reading nssa_default_information_originate_metric_type: %v", err)
		}
	}

	if err = d.Set("nssa_redistribution", dataSourceFlattenRouterospfAreaNssaRedistribution(o["nssa-redistribution"], d, "nssa_redistribution")); err != nil {
		if !fortiAPIPatch(o["nssa-redistribution"]) {
			return fmt.Errorf("error reading nssa_redistribution: %v", err)
		}
	}

	if err = d.Set("nssa_translator_role", dataSourceFlattenRouterospfAreaNssaTranslatorRole(o["nssa-translator-role"], d, "nssa_translator_role")); err != nil {
		if !fortiAPIPatch(o["nssa-translator-role"]) {
			return fmt.Errorf("error reading nssa_translator_role: %v", err)
		}
	}

	if err = d.Set("range", dataSourceFlattenRouterospfAreaRange(o["range"], d, "range")); err != nil {
		if !fortiAPIPatch(o["range"]) {
			return fmt.Errorf("error reading range: %v", err)
		}
	}

	if err = d.Set("shortcut", dataSourceFlattenRouterospfAreaShortcut(o["shortcut"], d, "shortcut")); err != nil {
		if !fortiAPIPatch(o["shortcut"]) {
			return fmt.Errorf("error reading shortcut: %v", err)
		}
	}

	if err = d.Set("stub_type", dataSourceFlattenRouterospfAreaStubType(o["stub-type"], d, "stub_type")); err != nil {
		if !fortiAPIPatch(o["stub-type"]) {
			return fmt.Errorf("error reading stub_type: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenRouterospfAreaType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("error reading type: %v", err)
		}
	}

	if err = d.Set("virtual_link", dataSourceFlattenRouterospfAreaVirtualLink(o["virtual-link"], d, "virtual_link")); err != nil {
		if !fortiAPIPatch(o["virtual-link"]) {
			return fmt.Errorf("error reading virtual_link: %v", err)
		}
	}

	return nil
}
