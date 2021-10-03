// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 | v6.4.6 merged schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: OSPF interface configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterospfOspfInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterospfOspfInterfaceCreate,
		Read:   resourceRouterospfOspfInterfaceRead,
		Update: resourceRouterospfOspfInterfaceUpdate,
		Delete: resourceRouterospfOspfInterfaceDelete,

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
				Type:        schema.TypeBool,
				Description: "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"authentication": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"none", "text", "message-digest", "md5"}),

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
			"bfd": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"global", "enable", "disable"}),

				Description: "Bidirectional Forwarding Detection (BFD).",
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
			"cost": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Cost of the interface, value range from 0 to 65535, 0 means auto-cost.",
				Optional:    true,
				Computed:    true,
			},
			"database_filter_out": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable control of flooding out LSAs.",
				Optional:    true,
				Computed:    true,
			},
			"dead_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Dead interval.",
				Optional:    true,
				Computed:    true,
			},
			"hello_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Hello interval.",
				Optional:    true,
				Computed:    true,
			},
			"hello_multiplier": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(3, 10),

				Description: "Number of hello packets within dead interval.",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Configuration interface name.",
				Optional:    true,
				Computed:    true,
			},
			"ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address.",
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
						},
					},
				},
			},
			"mtu": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(576, 65535),

				Description: "MTU for database description packets.",
				Optional:    true,
				Computed:    true,
			},
			"mtu_ignore": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable ignore MTU.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Interface entry name.",
				ForceNew:    true,
				Required:    true,
			},
			"network_type": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"broadcast", "non-broadcast", "point-to-point", "point-to-multipoint", "point-to-multipoint-non-broadcast"}),

				Description: "Network type.",
				Optional:    true,
				Computed:    true,
			},
			"prefix_length": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),

				Description: "Prefix length.",
				Optional:    true,
				Computed:    true,
			},
			"priority": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Priority.",
				Optional:    true,
				Computed:    true,
			},
			"resync_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),

				Description: "Graceful restart neighbor resynchronization timeout.",
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
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable status.",
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
	}
}

func resourceRouterospfOspfInterfaceCreate(d *schema.ResourceData, m interface{}) error {
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

	key := "name"
	mkey := ""
	if v, ok := d.GetOk(key); ok {
		if s, ok := v.(string); ok {
			mkey = s
		}
	}

	obj, err := getObjectRouterospfOspfInterface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating RouterospfOspfInterface resource while getting object: %v", err)
	}

	if mkey == "" && allow_append {
		return fmt.Errorf("error creating RouterospfOspfInterface resource: %q must be set if \"allow_append\" is true", key)
	}

	o := make(map[string]interface{})
	if mkey != "" && allow_append {
		o, err = c.UpdateRouterospfOspfInterface(obj, mkey, vdomparam, urlparams, batchid)
	} else {
		o, err = c.CreateRouterospfOspfInterface(obj, vdomparam, urlparams, batchid)
	}

	if err != nil {
		return fmt.Errorf("error creating RouterospfOspfInterface resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterospfOspfInterface")
	}

	return resourceRouterospfOspfInterfaceRead(d, m)
}

func resourceRouterospfOspfInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterospfOspfInterface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating RouterospfOspfInterface resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterospfOspfInterface(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating RouterospfOspfInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterospfOspfInterface")
	}

	return resourceRouterospfOspfInterfaceRead(d, m)
}

func resourceRouterospfOspfInterfaceDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteRouterospfOspfInterface(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting RouterospfOspfInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterospfOspfInterfaceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterospfOspfInterface(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading RouterospfOspfInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterospfOspfInterface(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading RouterospfOspfInterface resource from API: %v", err)
	}
	return nil
}

func flattenRouterospfOspfInterfaceAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceAuthenticationKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceBfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceCost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceDatabaseFilterOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceDeadInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceHelloMultiplier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceKeychain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceMd5Keys(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenRouterospfOspfInterfaceMd5KeysId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_string"
		if s, ok := i["key-string"].(string); ok && s != "ENC XXXX" {
			tmp["key_string"] = flattenRouterospfOspfInterfaceMd5KeysKeyString(i["key-string"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterospfOspfInterfaceMd5KeysId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceMd5KeysKeyString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceMtu(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceMtuIgnore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceNetworkType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfacePrefixLength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfacePriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceResyncTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceRetransmitInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceTransmitDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterospfOspfInterface(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("authentication", flattenRouterospfOspfInterfaceAuthentication(o["authentication"], d, "authentication", sv)); err != nil {
		if !fortiAPIPatch(o["authentication"]) {
			return fmt.Errorf("error reading authentication: %v", err)
		}
	}

	if s, ok := o["authentication-key"].(string); ok && s != "ENC XXXX" {
		if err = d.Set("authentication_key", flattenRouterospfOspfInterfaceAuthenticationKey(o["authentication-key"], d, "authentication_key", sv)); err != nil {
			if !fortiAPIPatch(o["authentication-key"]) {
				return fmt.Errorf("error reading authentication_key: %v", err)
			}
		}
	}

	if err = d.Set("bfd", flattenRouterospfOspfInterfaceBfd(o["bfd"], d, "bfd", sv)); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("error reading bfd: %v", err)
		}
	}

	if err = d.Set("comments", flattenRouterospfOspfInterfaceComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("error reading comments: %v", err)
		}
	}

	if err = d.Set("cost", flattenRouterospfOspfInterfaceCost(o["cost"], d, "cost", sv)); err != nil {
		if !fortiAPIPatch(o["cost"]) {
			return fmt.Errorf("error reading cost: %v", err)
		}
	}

	if err = d.Set("database_filter_out", flattenRouterospfOspfInterfaceDatabaseFilterOut(o["database-filter-out"], d, "database_filter_out", sv)); err != nil {
		if !fortiAPIPatch(o["database-filter-out"]) {
			return fmt.Errorf("error reading database_filter_out: %v", err)
		}
	}

	if err = d.Set("dead_interval", flattenRouterospfOspfInterfaceDeadInterval(o["dead-interval"], d, "dead_interval", sv)); err != nil {
		if !fortiAPIPatch(o["dead-interval"]) {
			return fmt.Errorf("error reading dead_interval: %v", err)
		}
	}

	if err = d.Set("hello_interval", flattenRouterospfOspfInterfaceHelloInterval(o["hello-interval"], d, "hello_interval", sv)); err != nil {
		if !fortiAPIPatch(o["hello-interval"]) {
			return fmt.Errorf("error reading hello_interval: %v", err)
		}
	}

	if err = d.Set("hello_multiplier", flattenRouterospfOspfInterfaceHelloMultiplier(o["hello-multiplier"], d, "hello_multiplier", sv)); err != nil {
		if !fortiAPIPatch(o["hello-multiplier"]) {
			return fmt.Errorf("error reading hello_multiplier: %v", err)
		}
	}

	if err = d.Set("interface", flattenRouterospfOspfInterfaceInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("error reading interface: %v", err)
		}
	}

	if err = d.Set("ip", flattenRouterospfOspfInterfaceIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("error reading ip: %v", err)
		}
	}

	if err = d.Set("keychain", flattenRouterospfOspfInterfaceKeychain(o["keychain"], d, "keychain", sv)); err != nil {
		if !fortiAPIPatch(o["keychain"]) {
			return fmt.Errorf("error reading keychain: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("md5_keys", flattenRouterospfOspfInterfaceMd5Keys(o["md5-keys"], d, "md5_keys", sv)); err != nil {
			if !fortiAPIPatch(o["md5-keys"]) {
				return fmt.Errorf("error reading md5_keys: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("md5_keys"); ok {
			if err = d.Set("md5_keys", flattenRouterospfOspfInterfaceMd5Keys(o["md5-keys"], d, "md5_keys", sv)); err != nil {
				if !fortiAPIPatch(o["md5-keys"]) {
					return fmt.Errorf("error reading md5_keys: %v", err)
				}
			}
		}
	}

	if err = d.Set("mtu", flattenRouterospfOspfInterfaceMtu(o["mtu"], d, "mtu", sv)); err != nil {
		if !fortiAPIPatch(o["mtu"]) {
			return fmt.Errorf("error reading mtu: %v", err)
		}
	}

	if err = d.Set("mtu_ignore", flattenRouterospfOspfInterfaceMtuIgnore(o["mtu-ignore"], d, "mtu_ignore", sv)); err != nil {
		if !fortiAPIPatch(o["mtu-ignore"]) {
			return fmt.Errorf("error reading mtu_ignore: %v", err)
		}
	}

	if err = d.Set("name", flattenRouterospfOspfInterfaceName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("network_type", flattenRouterospfOspfInterfaceNetworkType(o["network-type"], d, "network_type", sv)); err != nil {
		if !fortiAPIPatch(o["network-type"]) {
			return fmt.Errorf("error reading network_type: %v", err)
		}
	}

	if err = d.Set("prefix_length", flattenRouterospfOspfInterfacePrefixLength(o["prefix-length"], d, "prefix_length", sv)); err != nil {
		if !fortiAPIPatch(o["prefix-length"]) {
			return fmt.Errorf("error reading prefix_length: %v", err)
		}
	}

	if err = d.Set("priority", flattenRouterospfOspfInterfacePriority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("error reading priority: %v", err)
		}
	}

	if err = d.Set("resync_timeout", flattenRouterospfOspfInterfaceResyncTimeout(o["resync-timeout"], d, "resync_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["resync-timeout"]) {
			return fmt.Errorf("error reading resync_timeout: %v", err)
		}
	}

	if err = d.Set("retransmit_interval", flattenRouterospfOspfInterfaceRetransmitInterval(o["retransmit-interval"], d, "retransmit_interval", sv)); err != nil {
		if !fortiAPIPatch(o["retransmit-interval"]) {
			return fmt.Errorf("error reading retransmit_interval: %v", err)
		}
	}

	if err = d.Set("status", flattenRouterospfOspfInterfaceStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("error reading status: %v", err)
		}
	}

	if err = d.Set("transmit_delay", flattenRouterospfOspfInterfaceTransmitDelay(o["transmit-delay"], d, "transmit_delay", sv)); err != nil {
		if !fortiAPIPatch(o["transmit-delay"]) {
			return fmt.Errorf("error reading transmit_delay: %v", err)
		}
	}

	return nil
}

func expandRouterospfOspfInterfaceAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceAuthenticationKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceBfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceCost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceDatabaseFilterOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceDeadInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceHelloInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceHelloMultiplier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceKeychain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceMd5Keys(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandRouterospfOspfInterfaceMd5KeysId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_string"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["key-string"], _ = expandRouterospfOspfInterfaceMd5KeysKeyString(d, i["key_string"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterospfOspfInterfaceMd5KeysId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceMd5KeysKeyString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceMtu(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceMtuIgnore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceNetworkType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfacePrefixLength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfacePriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceResyncTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceRetransmitInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceTransmitDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterospfOspfInterface(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("authentication"); ok {
		t, err := expandRouterospfOspfInterfaceAuthentication(d, v, "authentication", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("authentication_key"); ok {
		t, err := expandRouterospfOspfInterfaceAuthenticationKey(d, v, "authentication_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication-key"] = t
		}
	}

	if v, ok := d.GetOk("bfd"); ok {
		t, err := expandRouterospfOspfInterfaceBfd(d, v, "bfd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandRouterospfOspfInterfaceComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("cost"); ok {
		t, err := expandRouterospfOspfInterfaceCost(d, v, "cost", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cost"] = t
		}
	}

	if v, ok := d.GetOk("database_filter_out"); ok {
		t, err := expandRouterospfOspfInterfaceDatabaseFilterOut(d, v, "database_filter_out", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database-filter-out"] = t
		}
	}

	if v, ok := d.GetOk("dead_interval"); ok {
		t, err := expandRouterospfOspfInterfaceDeadInterval(d, v, "dead_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dead-interval"] = t
		}
	}

	if v, ok := d.GetOk("hello_interval"); ok {
		t, err := expandRouterospfOspfInterfaceHelloInterval(d, v, "hello_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-interval"] = t
		}
	}

	if v, ok := d.GetOk("hello_multiplier"); ok {
		t, err := expandRouterospfOspfInterfaceHelloMultiplier(d, v, "hello_multiplier", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-multiplier"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandRouterospfOspfInterfaceInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandRouterospfOspfInterfaceIp(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("keychain"); ok {
		t, err := expandRouterospfOspfInterfaceKeychain(d, v, "keychain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keychain"] = t
		}
	}

	if v, ok := d.GetOk("md5_keys"); ok {
		t, err := expandRouterospfOspfInterfaceMd5Keys(d, v, "md5_keys", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["md5-keys"] = t
		}
	} else if d.HasChange("md5_keys") {
		old, new := d.GetChange("md5_keys")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["md5-keys"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("mtu"); ok {
		t, err := expandRouterospfOspfInterfaceMtu(d, v, "mtu", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mtu"] = t
		}
	}

	if v, ok := d.GetOk("mtu_ignore"); ok {
		t, err := expandRouterospfOspfInterfaceMtuIgnore(d, v, "mtu_ignore", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mtu-ignore"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandRouterospfOspfInterfaceName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("network_type"); ok {
		t, err := expandRouterospfOspfInterfaceNetworkType(d, v, "network_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-type"] = t
		}
	}

	if v, ok := d.GetOk("prefix_length"); ok {
		t, err := expandRouterospfOspfInterfacePrefixLength(d, v, "prefix_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-length"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok {
		t, err := expandRouterospfOspfInterfacePriority(d, v, "priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("resync_timeout"); ok {
		t, err := expandRouterospfOspfInterfaceResyncTimeout(d, v, "resync_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resync-timeout"] = t
		}
	}

	if v, ok := d.GetOk("retransmit_interval"); ok {
		t, err := expandRouterospfOspfInterfaceRetransmitInterval(d, v, "retransmit_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retransmit-interval"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandRouterospfOspfInterfaceStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("transmit_delay"); ok {
		t, err := expandRouterospfOspfInterfaceTransmitDelay(d, v, "transmit_delay", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transmit-delay"] = t
		}
	}

	return &obj, nil
}
