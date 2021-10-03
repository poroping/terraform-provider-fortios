// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: OSPF area configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterospfArea() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterospfAreaCreate,
		Read:   resourceRouterospfAreaRead,
		Update: resourceRouterospfAreaUpdate,
		Delete: resourceRouterospfAreaDelete,

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
			"authentication": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"none", "text", "message-digest"}),

				Description: "Authentication type.",
				Optional:    true,
				Computed:    true,
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"default_cost": {
				Type: schema.TypeInt,

				Description: "Summary default cost of stub or NSSA area.",
				Optional:    true,
				Computed:    true,
			},
			"filter_list": {
				Type:        schema.TypeList,
				Description: "OSPF area filter-list configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"direction": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"in", "out"}),

							Description: "Direction.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Filter list entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"list": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Access-list or prefix-list name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"fosid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Area entry IP address.",
				Optional:    true,
				Computed:    true,
			},
			"nssa_default_information_originate": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"enable", "always", "disable"}),

				Description: "Redistribute, advertise, or do not originate Type-7 default route into NSSA area.",
				Optional:    true,
				Computed:    true,
			},
			"nssa_default_information_originate_metric": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16777214),

				Description: "OSPF default metric.",
				Optional:    true,
				Computed:    true,
			},
			"nssa_default_information_originate_metric_type": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"1", "2"}),

				Description: "OSPF metric type for default routes.",
				Optional:    true,
				Computed:    true,
			},
			"nssa_redistribution": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable redistribute into NSSA area.",
				Optional:    true,
				Computed:    true,
			},
			"nssa_translator_role": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"candidate", "never", "always"}),

				Description: "NSSA translator role type.",
				Optional:    true,
				Computed:    true,
			},
			"range": {
				Type:        schema.TypeList,
				Description: "OSPF area range configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advertise": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable advertise status.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Range entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"prefix": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateIPv4ClassnetAny,

							Description: "Prefix.",
							Optional:    true,
							Computed:    true,
						},
						"substitute": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateIPv4ClassnetAny,

							Description: "Substitute prefix.",
							Optional:    true,
							Computed:    true,
						},
						"substitute_status": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable substitute status.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"shortcut": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"disable", "enable", "default"}),

				Description: "Enable/disable shortcut option.",
				Optional:    true,
				Computed:    true,
			},
			"stub_type": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"no-summary", "summary"}),

				Description: "Stub summary setting.",
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"regular", "nssa", "stub"}),

				Description: "Area type setting.",
				Optional:    true,
				Computed:    true,
			},
			"virtual_link": {
				Type:        schema.TypeList,
				Description: "OSPF virtual link configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"none", "text", "message-digest"}),

							Description: "Authentication type.",
							Optional:    true,
							Computed:    true,
						},
						"authentication_key": {
							Type: schema.TypeString,

							Description: "Authentication key.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"dead_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Dead interval.",
							Optional:    true,
							Computed:    true,
						},
						"hello_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Hello interval.",
							Optional:    true,
							Computed:    true,
						},
						"keychain": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Message-digest key-chain name.",
							Optional:    true,
							Computed:    true,
						},
						"md5_keys": {
							Type:        schema.TypeList,
							Description: "MD5 key.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 255),

										Description: "Key ID (1 - 255).",
										Optional:    true,
										Computed:    true,
									},
									"key_string": {
										Type: schema.TypeString,

										Description: "Password for the key.",
										Optional:    true,
										Computed:    true,
										Sensitive:   true,
									},
								},
							},
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Virtual link entry name.",
							Optional:    true,
							Computed:    true,
						},
						"peer": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Peer IP.",
							Optional:    true,
							Computed:    true,
						},
						"retransmit_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Retransmit interval.",
							Optional:    true,
							Computed:    true,
						},
						"transmit_delay": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Transmit delay.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceRouterospfAreaCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterospfArea(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating RouterospfArea resource while getting object: %v", err)
	}

	if mkey == "" && allow_append {
		return fmt.Errorf("error creating RouterospfArea resource: %q must be set if \"allow_append\" is true", key)
	}

	o := make(map[string]interface{})
	if mkey != "" && allow_append {
		o, err = c.UpdateRouterospfArea(obj, mkey, vdomparam, urlparams, batchid)
	} else {
		o, err = c.CreateRouterospfArea(obj, vdomparam, urlparams, batchid)
	}

	if err != nil {
		return fmt.Errorf("error creating RouterospfArea resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterospfArea")
	}

	return resourceRouterospfAreaRead(d, m)
}

func resourceRouterospfAreaUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterospfArea(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating RouterospfArea resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterospfArea(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating RouterospfArea resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterospfArea")
	}

	return resourceRouterospfAreaRead(d, m)
}

func resourceRouterospfAreaDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteRouterospfArea(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting RouterospfArea resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterospfAreaRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterospfArea(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading RouterospfArea resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterospfArea(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading RouterospfArea resource from API: %v", err)
	}
	return nil
}

func flattenRouterospfAreaAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaDefaultCost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaFilterList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["direction"] = flattenRouterospfAreaFilterListDirection(i["direction"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenRouterospfAreaFilterListId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "list"
		if _, ok := i["list"]; ok {

			tmp["list"] = flattenRouterospfAreaFilterListList(i["list"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterospfAreaFilterListDirection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaFilterListId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaFilterListList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaNssaDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaNssaDefaultInformationOriginateMetric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaNssaDefaultInformationOriginateMetricType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaNssaRedistribution(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaNssaTranslatorRole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaRange(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["advertise"] = flattenRouterospfAreaRangeAdvertise(i["advertise"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenRouterospfAreaRangeId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {

			tmp["prefix"] = flattenRouterospfAreaRangePrefix(i["prefix"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute"
		if _, ok := i["substitute"]; ok {

			tmp["substitute"] = flattenRouterospfAreaRangeSubstitute(i["substitute"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute_status"
		if _, ok := i["substitute-status"]; ok {

			tmp["substitute_status"] = flattenRouterospfAreaRangeSubstituteStatus(i["substitute-status"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterospfAreaRangeAdvertise(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaRangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaRangePrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}
	return v
}

func flattenRouterospfAreaRangeSubstitute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}
	return v
}

func flattenRouterospfAreaRangeSubstituteStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaShortcut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaStubType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaVirtualLink(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["authentication"] = flattenRouterospfAreaVirtualLinkAuthentication(i["authentication"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication_key"
		if s, ok := i["authentication-key"].(string); ok && s != "ENC XXXX" {
			tmp["authentication_key"] = flattenRouterospfAreaVirtualLinkAuthenticationKey(i["authentication-key"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := i["dead-interval"]; ok {

			tmp["dead_interval"] = flattenRouterospfAreaVirtualLinkDeadInterval(i["dead-interval"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {

			tmp["hello_interval"] = flattenRouterospfAreaVirtualLinkHelloInterval(i["hello-interval"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keychain"
		if _, ok := i["keychain"]; ok {

			tmp["keychain"] = flattenRouterospfAreaVirtualLinkKeychain(i["keychain"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keys"
		if _, ok := i["md5-keys"]; ok {

			tmp["md5_keys"] = flattenRouterospfAreaVirtualLinkMd5Keys(i["md5-keys"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenRouterospfAreaVirtualLinkName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := i["peer"]; ok {

			tmp["peer"] = flattenRouterospfAreaVirtualLinkPeer(i["peer"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := i["retransmit-interval"]; ok {

			tmp["retransmit_interval"] = flattenRouterospfAreaVirtualLinkRetransmitInterval(i["retransmit-interval"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := i["transmit-delay"]; ok {

			tmp["transmit_delay"] = flattenRouterospfAreaVirtualLinkTransmitDelay(i["transmit-delay"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterospfAreaVirtualLinkAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaVirtualLinkAuthenticationKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaVirtualLinkDeadInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaVirtualLinkHelloInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaVirtualLinkKeychain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaVirtualLinkMd5Keys(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenRouterospfAreaVirtualLinkMd5KeysId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_string"
		if s, ok := i["key-string"].(string); ok && s != "ENC XXXX" {
			tmp["key_string"] = flattenRouterospfAreaVirtualLinkMd5KeysKeyString(i["key-string"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterospfAreaVirtualLinkMd5KeysId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaVirtualLinkMd5KeysKeyString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaVirtualLinkName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaVirtualLinkPeer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaVirtualLinkRetransmitInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfAreaVirtualLinkTransmitDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterospfArea(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("authentication", flattenRouterospfAreaAuthentication(o["authentication"], d, "authentication", sv)); err != nil {
		if !fortiAPIPatch(o["authentication"]) {
			return fmt.Errorf("error reading authentication: %v", err)
		}
	}

	if err = d.Set("comments", flattenRouterospfAreaComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("error reading comments: %v", err)
		}
	}

	if err = d.Set("default_cost", flattenRouterospfAreaDefaultCost(o["default-cost"], d, "default_cost", sv)); err != nil {
		if !fortiAPIPatch(o["default-cost"]) {
			return fmt.Errorf("error reading default_cost: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("filter_list", flattenRouterospfAreaFilterList(o["filter-list"], d, "filter_list", sv)); err != nil {
			if !fortiAPIPatch(o["filter-list"]) {
				return fmt.Errorf("error reading filter_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("filter_list"); ok {
			if err = d.Set("filter_list", flattenRouterospfAreaFilterList(o["filter-list"], d, "filter_list", sv)); err != nil {
				if !fortiAPIPatch(o["filter-list"]) {
					return fmt.Errorf("error reading filter_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenRouterospfAreaId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("error reading fosid: %v", err)
		}
	}

	if err = d.Set("nssa_default_information_originate", flattenRouterospfAreaNssaDefaultInformationOriginate(o["nssa-default-information-originate"], d, "nssa_default_information_originate", sv)); err != nil {
		if !fortiAPIPatch(o["nssa-default-information-originate"]) {
			return fmt.Errorf("error reading nssa_default_information_originate: %v", err)
		}
	}

	if err = d.Set("nssa_default_information_originate_metric", flattenRouterospfAreaNssaDefaultInformationOriginateMetric(o["nssa-default-information-originate-metric"], d, "nssa_default_information_originate_metric", sv)); err != nil {
		if !fortiAPIPatch(o["nssa-default-information-originate-metric"]) {
			return fmt.Errorf("error reading nssa_default_information_originate_metric: %v", err)
		}
	}

	if err = d.Set("nssa_default_information_originate_metric_type", flattenRouterospfAreaNssaDefaultInformationOriginateMetricType(o["nssa-default-information-originate-metric-type"], d, "nssa_default_information_originate_metric_type", sv)); err != nil {
		if !fortiAPIPatch(o["nssa-default-information-originate-metric-type"]) {
			return fmt.Errorf("error reading nssa_default_information_originate_metric_type: %v", err)
		}
	}

	if err = d.Set("nssa_redistribution", flattenRouterospfAreaNssaRedistribution(o["nssa-redistribution"], d, "nssa_redistribution", sv)); err != nil {
		if !fortiAPIPatch(o["nssa-redistribution"]) {
			return fmt.Errorf("error reading nssa_redistribution: %v", err)
		}
	}

	if err = d.Set("nssa_translator_role", flattenRouterospfAreaNssaTranslatorRole(o["nssa-translator-role"], d, "nssa_translator_role", sv)); err != nil {
		if !fortiAPIPatch(o["nssa-translator-role"]) {
			return fmt.Errorf("error reading nssa_translator_role: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("range", flattenRouterospfAreaRange(o["range"], d, "range", sv)); err != nil {
			if !fortiAPIPatch(o["range"]) {
				return fmt.Errorf("error reading range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("range"); ok {
			if err = d.Set("range", flattenRouterospfAreaRange(o["range"], d, "range", sv)); err != nil {
				if !fortiAPIPatch(o["range"]) {
					return fmt.Errorf("error reading range: %v", err)
				}
			}
		}
	}

	if err = d.Set("shortcut", flattenRouterospfAreaShortcut(o["shortcut"], d, "shortcut", sv)); err != nil {
		if !fortiAPIPatch(o["shortcut"]) {
			return fmt.Errorf("error reading shortcut: %v", err)
		}
	}

	if err = d.Set("stub_type", flattenRouterospfAreaStubType(o["stub-type"], d, "stub_type", sv)); err != nil {
		if !fortiAPIPatch(o["stub-type"]) {
			return fmt.Errorf("error reading stub_type: %v", err)
		}
	}

	if err = d.Set("type", flattenRouterospfAreaType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("error reading type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("virtual_link", flattenRouterospfAreaVirtualLink(o["virtual-link"], d, "virtual_link", sv)); err != nil {
			if !fortiAPIPatch(o["virtual-link"]) {
				return fmt.Errorf("error reading virtual_link: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("virtual_link"); ok {
			if err = d.Set("virtual_link", flattenRouterospfAreaVirtualLink(o["virtual-link"], d, "virtual_link", sv)); err != nil {
				if !fortiAPIPatch(o["virtual-link"]) {
					return fmt.Errorf("error reading virtual_link: %v", err)
				}
			}
		}
	}

	return nil
}

func expandRouterospfAreaAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaDefaultCost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaFilterList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["direction"], _ = expandRouterospfAreaFilterListDirection(d, i["direction"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandRouterospfAreaFilterListId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "list"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["list"], _ = expandRouterospfAreaFilterListList(d, i["list"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterospfAreaFilterListDirection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaFilterListId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaFilterListList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaNssaDefaultInformationOriginate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaNssaDefaultInformationOriginateMetric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaNssaDefaultInformationOriginateMetricType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaNssaRedistribution(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaNssaTranslatorRole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["advertise"], _ = expandRouterospfAreaRangeAdvertise(d, i["advertise"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandRouterospfAreaRangeId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix"], _ = expandRouterospfAreaRangePrefix(d, i["prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["substitute"], _ = expandRouterospfAreaRangeSubstitute(d, i["substitute"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute_status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["substitute-status"], _ = expandRouterospfAreaRangeSubstituteStatus(d, i["substitute_status"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterospfAreaRangeAdvertise(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaRangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaRangePrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaRangeSubstitute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaRangeSubstituteStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaShortcut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaStubType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaVirtualLink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["authentication"], _ = expandRouterospfAreaVirtualLinkAuthentication(d, i["authentication"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication_key"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["authentication-key"], _ = expandRouterospfAreaVirtualLinkAuthenticationKey(d, i["authentication_key"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dead-interval"], _ = expandRouterospfAreaVirtualLinkDeadInterval(d, i["dead_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["hello-interval"], _ = expandRouterospfAreaVirtualLinkHelloInterval(d, i["hello_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keychain"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["keychain"], _ = expandRouterospfAreaVirtualLinkKeychain(d, i["keychain"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keys"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["md5-keys"], _ = expandRouterospfAreaVirtualLinkMd5Keys(d, i["md5_keys"], pre_append, sv)
		} else {
			tmp["md5-keys"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandRouterospfAreaVirtualLinkName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["peer"], _ = expandRouterospfAreaVirtualLinkPeer(d, i["peer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["retransmit-interval"], _ = expandRouterospfAreaVirtualLinkRetransmitInterval(d, i["retransmit_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["transmit-delay"], _ = expandRouterospfAreaVirtualLinkTransmitDelay(d, i["transmit_delay"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterospfAreaVirtualLinkAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaVirtualLinkAuthenticationKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaVirtualLinkDeadInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaVirtualLinkHelloInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaVirtualLinkKeychain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaVirtualLinkMd5Keys(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandRouterospfAreaVirtualLinkMd5KeysId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_string"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["key-string"], _ = expandRouterospfAreaVirtualLinkMd5KeysKeyString(d, i["key_string"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterospfAreaVirtualLinkMd5KeysId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaVirtualLinkMd5KeysKeyString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaVirtualLinkName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaVirtualLinkPeer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaVirtualLinkRetransmitInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfAreaVirtualLinkTransmitDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterospfArea(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("authentication"); ok {
		t, err := expandRouterospfAreaAuthentication(d, v, "authentication", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandRouterospfAreaComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("default_cost"); ok {
		t, err := expandRouterospfAreaDefaultCost(d, v, "default_cost", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-cost"] = t
		}
	}

	if v, ok := d.GetOk("filter_list"); ok {
		t, err := expandRouterospfAreaFilterList(d, v, "filter_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list"] = t
		}
	} else if d.HasChange("filter_list") {
		old, new := d.GetChange("filter_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["filter-list"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandRouterospfAreaId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("nssa_default_information_originate"); ok {
		t, err := expandRouterospfAreaNssaDefaultInformationOriginate(d, v, "nssa_default_information_originate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nssa-default-information-originate"] = t
		}
	}

	if v, ok := d.GetOk("nssa_default_information_originate_metric"); ok {
		t, err := expandRouterospfAreaNssaDefaultInformationOriginateMetric(d, v, "nssa_default_information_originate_metric", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nssa-default-information-originate-metric"] = t
		}
	}

	if v, ok := d.GetOk("nssa_default_information_originate_metric_type"); ok {
		t, err := expandRouterospfAreaNssaDefaultInformationOriginateMetricType(d, v, "nssa_default_information_originate_metric_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nssa-default-information-originate-metric-type"] = t
		}
	}

	if v, ok := d.GetOk("nssa_redistribution"); ok {
		t, err := expandRouterospfAreaNssaRedistribution(d, v, "nssa_redistribution", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nssa-redistribution"] = t
		}
	}

	if v, ok := d.GetOk("nssa_translator_role"); ok {
		t, err := expandRouterospfAreaNssaTranslatorRole(d, v, "nssa_translator_role", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nssa-translator-role"] = t
		}
	}

	if v, ok := d.GetOk("range"); ok {
		t, err := expandRouterospfAreaRange(d, v, "range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	} else if d.HasChange("range") {
		old, new := d.GetChange("range")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["range"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("shortcut"); ok {
		t, err := expandRouterospfAreaShortcut(d, v, "shortcut", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shortcut"] = t
		}
	}

	if v, ok := d.GetOk("stub_type"); ok {
		t, err := expandRouterospfAreaStubType(d, v, "stub_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stub-type"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandRouterospfAreaType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("virtual_link"); ok {
		t, err := expandRouterospfAreaVirtualLink(d, v, "virtual_link", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-link"] = t
		}
	} else if d.HasChange("virtual_link") {
		old, new := d.GetChange("virtual_link")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["virtual-link"] = make([]struct{}, 0)
		}
	}

	return &obj, nil
}
