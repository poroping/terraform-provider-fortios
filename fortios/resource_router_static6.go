// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 | v6.4.2 | v6.2.7 merged schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: Configure IPv6 static routing tables.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterStatic6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterStatic6Create,
		Read:   resourceRouterStatic6Read,
		Update: resourceRouterStatic6Update,
		Delete: resourceRouterStatic6Delete,

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
			"dynamic_sort_table": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"bfd": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable Bidirectional Forwarding Detection (BFD).",
				Optional:    true,
				Computed:    true,
			},
			"blackhole": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable black hole.",
				Optional:    true,
				Computed:    true,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Optional comments.",
				Optional:    true,
				Computed:    true,
			},
			"device": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Gateway out interface or tunnel.",
				Optional:    true,
				Computed:    true,
			},
			"devindex": {
				Type: schema.TypeInt,

				Description: "Device index (0 - 4294967295).",
				Optional:    true,
				Computed:    true,
			},
			"distance": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Administrative distance (1 - 255).",
				Optional:    true,
				Computed:    true,
			},
			"dst": {
				Type:             schema.TypeString,
				ValidateFunc:     fortiValidateIPv6Network,
				DiffSuppressFunc: diffCidrEqual,
				Description:      "Destination IPv6 prefix.",
				Optional:         true,
				Computed:         true,
			},
			"dynamic_gateway": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable use of dynamic gateway retrieved from Router Advertisement (RA).",
				Optional:    true,
				Computed:    true,
			},
			"gateway": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: diffIPEqual,
				Description:      "IPv6 address of the gateway.",
				Optional:         true,
				Computed:         true,
			},
			"link_monitor_exempt": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable withdrawal of this static route when link monitor or health check is down.",
				Optional:    true,
				Computed:    true,
			},
			"priority": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Administrative priority (1 - 65535).",
				Optional:    true,
				Computed:    true,
			},
			"sdwan": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable egress through the SD-WAN.",
				Optional:    true,
				Computed:    true,
			},
			"sdwan_zone": {
				Type:        schema.TypeList,
				Description: "Choose SD-WAN Zone.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "SD-WAN zone name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"seq_num": {
				Type: schema.TypeInt,

				Description: "Sequence number.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable this static route.",
				Optional:    true,
				Computed:    true,
			},
			"virtual_wan_link": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable egress through the virtual-wan-link.",
				Optional:    true,
				Computed:    true,
			},
			"vrf": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 31),

				Description: "Virtual Routing Forwarding ID.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceRouterStatic6Create(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterStatic6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating RouterStatic6 resource while getting object: %v", err)
	}

	o, err := c.CreateRouterStatic6(obj, vdomparam, urlparams, batchid)

	if err != nil {
		return fmt.Errorf("error creating RouterStatic6 resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterStatic6")
	}

	return resourceRouterStatic6Read(d, m)
}

func resourceRouterStatic6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterStatic6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating RouterStatic6 resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterStatic6(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating RouterStatic6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterStatic6")
	}

	return resourceRouterStatic6Read(d, m)
}

func resourceRouterStatic6Delete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteRouterStatic6(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting RouterStatic6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterStatic6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterStatic6(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading RouterStatic6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterStatic6(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading RouterStatic6 resource from API: %v", err)
	}
	return nil
}

func flattenRouterStatic6Bfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStatic6Blackhole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStatic6Comment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStatic6Device(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStatic6Devindex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStatic6Distance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStatic6Dst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStatic6DynamicGateway(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStatic6Gateway(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStatic6LinkMonitorExempt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStatic6Priority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStatic6Sdwan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStatic6SdwanZone(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenRouterStatic6SdwanZoneName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterStatic6SdwanZoneName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStatic6SeqNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStatic6Status(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStatic6VirtualWanLink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStatic6Vrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterStatic6(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("bfd", flattenRouterStatic6Bfd(o["bfd"], d, "bfd", sv)); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("error reading bfd: %v", err)
		}
	}

	if err = d.Set("blackhole", flattenRouterStatic6Blackhole(o["blackhole"], d, "blackhole", sv)); err != nil {
		if !fortiAPIPatch(o["blackhole"]) {
			return fmt.Errorf("error reading blackhole: %v", err)
		}
	}

	if err = d.Set("comment", flattenRouterStatic6Comment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("error reading comment: %v", err)
		}
	}

	if err = d.Set("device", flattenRouterStatic6Device(o["device"], d, "device", sv)); err != nil {
		if !fortiAPIPatch(o["device"]) {
			return fmt.Errorf("error reading device: %v", err)
		}
	}

	if err = d.Set("devindex", flattenRouterStatic6Devindex(o["devindex"], d, "devindex", sv)); err != nil {
		if !fortiAPIPatch(o["devindex"]) {
			return fmt.Errorf("error reading devindex: %v", err)
		}
	}

	if err = d.Set("distance", flattenRouterStatic6Distance(o["distance"], d, "distance", sv)); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("error reading distance: %v", err)
		}
	}

	if err = d.Set("dst", flattenRouterStatic6Dst(o["dst"], d, "dst", sv)); err != nil {
		if !fortiAPIPatch(o["dst"]) {
			return fmt.Errorf("error reading dst: %v", err)
		}
	}

	if err = d.Set("dynamic_gateway", flattenRouterStatic6DynamicGateway(o["dynamic-gateway"], d, "dynamic_gateway", sv)); err != nil {
		if !fortiAPIPatch(o["dynamic-gateway"]) {
			return fmt.Errorf("error reading dynamic_gateway: %v", err)
		}
	}

	if err = d.Set("gateway", flattenRouterStatic6Gateway(o["gateway"], d, "gateway", sv)); err != nil {
		if !fortiAPIPatch(o["gateway"]) {
			return fmt.Errorf("error reading gateway: %v", err)
		}
	}

	if err = d.Set("link_monitor_exempt", flattenRouterStatic6LinkMonitorExempt(o["link-monitor-exempt"], d, "link_monitor_exempt", sv)); err != nil {
		if !fortiAPIPatch(o["link-monitor-exempt"]) {
			return fmt.Errorf("error reading link_monitor_exempt: %v", err)
		}
	}

	if err = d.Set("priority", flattenRouterStatic6Priority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("error reading priority: %v", err)
		}
	}

	if err = d.Set("sdwan", flattenRouterStatic6Sdwan(o["sdwan"], d, "sdwan", sv)); err != nil {
		if !fortiAPIPatch(o["sdwan"]) {
			return fmt.Errorf("error reading sdwan: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("sdwan_zone", flattenRouterStatic6SdwanZone(o["sdwan-zone"], d, "sdwan_zone", sv)); err != nil {
			if !fortiAPIPatch(o["sdwan-zone"]) {
				return fmt.Errorf("error reading sdwan_zone: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sdwan_zone"); ok {
			if err = d.Set("sdwan_zone", flattenRouterStatic6SdwanZone(o["sdwan-zone"], d, "sdwan_zone", sv)); err != nil {
				if !fortiAPIPatch(o["sdwan-zone"]) {
					return fmt.Errorf("error reading sdwan_zone: %v", err)
				}
			}
		}
	}

	if err = d.Set("seq_num", flattenRouterStatic6SeqNum(o["seq-num"], d, "seq_num", sv)); err != nil {
		if !fortiAPIPatch(o["seq-num"]) {
			return fmt.Errorf("error reading seq_num: %v", err)
		}
	}

	if err = d.Set("status", flattenRouterStatic6Status(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("error reading status: %v", err)
		}
	}

	if err = d.Set("virtual_wan_link", flattenRouterStatic6VirtualWanLink(o["virtual-wan-link"], d, "virtual_wan_link", sv)); err != nil {
		if !fortiAPIPatch(o["virtual-wan-link"]) {
			return fmt.Errorf("error reading virtual_wan_link: %v", err)
		}
	}

	if err = d.Set("vrf", flattenRouterStatic6Vrf(o["vrf"], d, "vrf", sv)); err != nil {
		if !fortiAPIPatch(o["vrf"]) {
			return fmt.Errorf("error reading vrf: %v", err)
		}
	}

	return nil
}

func expandRouterStatic6Bfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Blackhole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Comment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Device(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Devindex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Distance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Dst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6DynamicGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Gateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6LinkMonitorExempt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Priority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Sdwan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6SdwanZone(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandRouterStatic6SdwanZoneName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterStatic6SdwanZoneName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6SeqNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Status(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6VirtualWanLink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Vrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterStatic6(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bfd"); ok {
		t, err := expandRouterStatic6Bfd(d, v, "bfd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("blackhole"); ok {
		t, err := expandRouterStatic6Blackhole(d, v, "blackhole", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["blackhole"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandRouterStatic6Comment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("device"); ok {
		t, err := expandRouterStatic6Device(d, v, "device", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("devindex"); ok {
		t, err := expandRouterStatic6Devindex(d, v, "devindex", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["devindex"] = t
		}
	}

	if v, ok := d.GetOk("distance"); ok {
		t, err := expandRouterStatic6Distance(d, v, "distance", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok {
		t, err := expandRouterStatic6Dst(d, v, "dst", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_gateway"); ok {
		t, err := expandRouterStatic6DynamicGateway(d, v, "dynamic_gateway", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-gateway"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok {
		t, err := expandRouterStatic6Gateway(d, v, "gateway", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("link_monitor_exempt"); ok {
		t, err := expandRouterStatic6LinkMonitorExempt(d, v, "link_monitor_exempt", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-monitor-exempt"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok {
		t, err := expandRouterStatic6Priority(d, v, "priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("sdwan"); ok {
		t, err := expandRouterStatic6Sdwan(d, v, "sdwan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdwan"] = t
		}
	}

	if v, ok := d.GetOk("sdwan_zone"); ok {
		t, err := expandRouterStatic6SdwanZone(d, v, "sdwan_zone", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdwan-zone"] = t
		}
	}

	if v, ok := d.GetOk("seq_num"); ok {
		t, err := expandRouterStatic6SeqNum(d, v, "seq_num", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seq-num"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandRouterStatic6Status(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("virtual_wan_link"); ok {
		t, err := expandRouterStatic6VirtualWanLink(d, v, "virtual_wan_link", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-wan-link"] = t
		}
	}

	if v, ok := d.GetOk("vrf"); ok {
		t, err := expandRouterStatic6Vrf(d, v, "vrf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf"] = t
		}
	}

	return &obj, nil
}
