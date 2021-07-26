// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: FortiGate interfaces added to the SD-WAN.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemsdwanMembers() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemsdwanMembersCreate,
		Read:   resourceSystemsdwanMembersRead,
		Update: resourceSystemsdwanMembersUpdate,
		Delete: resourceSystemsdwanMembersDelete,

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
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Description:  "Comments.",
				Optional:     true,
				Computed:     true,
			},
			"cost": {
				Type: schema.TypeInt,

				Description: "Cost of this interface for services in SLA mode (0 - 4294967295, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"gateway": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,
				Description:  "The default gateway for this interface. Usually the default gateway of the Internet service provider that this interface is connected to.",
				Optional:     true,
				Computed:     true,
			},
			"gateway6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv6Address,
				Description:  "IPv6 gateway.",
				Optional:     true,
				Computed:     true,
			},
			"ingress_spillover_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),
				Description:  "Ingress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.",
				Optional:     true,
				Computed:     true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Description:  "Interface name.",
				Optional:     true,
				Computed:     true,
			},
			"priority": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Description:  "Priority of the interface for IPv4 (0 - 65535, default = 0). Used for SD-WAN rules or priority rules.",
				Optional:     true,
				Computed:     true,
			},
			"priority6": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Description:  "Priority of the interface for IPv6 (1 - 65535, default = 1024). Used for SD-WAN rules or priority rules.",
				Optional:     true,
				Computed:     true,
			},
			"seq_num": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 512),
				Description:  "Sequence number(1-512).",
				ForceNew:     true,
				Required:     true,
			},
			"source": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,
				Description:  "Source IP address used in the health-check packet to the server.",
				Optional:     true,
				Computed:     true,
			},
			"source6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv6Address,
				Description:  "Source IPv6 address used in the health-check packet to the server.",
				Optional:     true,
				Computed:     true,
			},
			"spillover_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),
				Description:  "Egress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.",
				Optional:     true,
				Computed:     true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable this interface in the SD-WAN.",
				Optional:     true,
				Computed:     true,
			},
			"volume_ratio": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Description:  "Measured volume ratio (this value / sum of all values = percentage of link volume, 1 - 255).",
				Optional:     true,
				Computed:     true,
			},
			"weight": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Description:  "Weight of this interface for weighted load balancing. (1 - 255) More traffic is directed to interfaces with higher weights.",
				Optional:     true,
				Computed:     true,
			},
			"zone": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "Zone name.",
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSystemsdwanMembersCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemsdwanMembers(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating SystemsdwanMembers resource while getting object: %v", err)
	}

	o, err := c.CreateSystemsdwanMembers(obj, vdomparam, urlparams, batchid)

	if err != nil {
		return fmt.Errorf("error creating SystemsdwanMembers resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemsdwanMembers")
	}

	return resourceSystemsdwanMembersRead(d, m)
}

func resourceSystemsdwanMembersUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemsdwanMembers(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating SystemsdwanMembers resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemsdwanMembers(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating SystemsdwanMembers resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemsdwanMembers")
	}

	return resourceSystemsdwanMembersRead(d, m)
}

func resourceSystemsdwanMembersDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteSystemsdwanMembers(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting SystemsdwanMembers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemsdwanMembersRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemsdwanMembers(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading SystemsdwanMembers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemsdwanMembers(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading SystemsdwanMembers resource from API: %v", err)
	}
	return nil
}

func flattenSystemsdwanMembersComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanMembersCost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanMembersGateway(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanMembersGateway6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanMembersIngressSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanMembersInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanMembersPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanMembersPriority6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanMembersSeqNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanMembersSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanMembersSource6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanMembersSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanMembersStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanMembersVolumeRatio(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanMembersWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanMembersZone(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemsdwanMembers(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("comment", flattenSystemsdwanMembersComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("error reading comment: %v", err)
		}
	}

	if err = d.Set("cost", flattenSystemsdwanMembersCost(o["cost"], d, "cost", sv)); err != nil {
		if !fortiAPIPatch(o["cost"]) {
			return fmt.Errorf("error reading cost: %v", err)
		}
	}

	if err = d.Set("gateway", flattenSystemsdwanMembersGateway(o["gateway"], d, "gateway", sv)); err != nil {
		if !fortiAPIPatch(o["gateway"]) {
			return fmt.Errorf("error reading gateway: %v", err)
		}
	}

	if err = d.Set("gateway6", flattenSystemsdwanMembersGateway6(o["gateway6"], d, "gateway6", sv)); err != nil {
		if !fortiAPIPatch(o["gateway6"]) {
			return fmt.Errorf("error reading gateway6: %v", err)
		}
	}

	if err = d.Set("ingress_spillover_threshold", flattenSystemsdwanMembersIngressSpilloverThreshold(o["ingress-spillover-threshold"], d, "ingress_spillover_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["ingress-spillover-threshold"]) {
			return fmt.Errorf("error reading ingress_spillover_threshold: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemsdwanMembersInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("error reading interface: %v", err)
		}
	}

	if err = d.Set("priority", flattenSystemsdwanMembersPriority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("error reading priority: %v", err)
		}
	}

	if err = d.Set("priority6", flattenSystemsdwanMembersPriority6(o["priority6"], d, "priority6", sv)); err != nil {
		if !fortiAPIPatch(o["priority6"]) {
			return fmt.Errorf("error reading priority6: %v", err)
		}
	}

	if err = d.Set("seq_num", flattenSystemsdwanMembersSeqNum(o["seq-num"], d, "seq_num", sv)); err != nil {
		if !fortiAPIPatch(o["seq-num"]) {
			return fmt.Errorf("error reading seq_num: %v", err)
		}
	}

	if err = d.Set("source", flattenSystemsdwanMembersSource(o["source"], d, "source", sv)); err != nil {
		if !fortiAPIPatch(o["source"]) {
			return fmt.Errorf("error reading source: %v", err)
		}
	}

	if err = d.Set("source6", flattenSystemsdwanMembersSource6(o["source6"], d, "source6", sv)); err != nil {
		if !fortiAPIPatch(o["source6"]) {
			return fmt.Errorf("error reading source6: %v", err)
		}
	}

	if err = d.Set("spillover_threshold", flattenSystemsdwanMembersSpilloverThreshold(o["spillover-threshold"], d, "spillover_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["spillover-threshold"]) {
			return fmt.Errorf("error reading spillover_threshold: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemsdwanMembersStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("error reading status: %v", err)
		}
	}

	if err = d.Set("volume_ratio", flattenSystemsdwanMembersVolumeRatio(o["volume-ratio"], d, "volume_ratio", sv)); err != nil {
		if !fortiAPIPatch(o["volume-ratio"]) {
			return fmt.Errorf("error reading volume_ratio: %v", err)
		}
	}

	if err = d.Set("weight", flattenSystemsdwanMembersWeight(o["weight"], d, "weight", sv)); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("error reading weight: %v", err)
		}
	}

	if err = d.Set("zone", flattenSystemsdwanMembersZone(o["zone"], d, "zone", sv)); err != nil {
		if !fortiAPIPatch(o["zone"]) {
			return fmt.Errorf("error reading zone: %v", err)
		}
	}

	return nil
}

func expandSystemsdwanMembersComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanMembersCost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanMembersGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanMembersGateway6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanMembersIngressSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanMembersInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanMembersPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanMembersPriority6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanMembersSeqNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanMembersSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanMembersSource6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanMembersSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanMembersStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanMembersVolumeRatio(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanMembersWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanMembersZone(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemsdwanMembers(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandSystemsdwanMembersComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("cost"); ok {
		t, err := expandSystemsdwanMembersCost(d, v, "cost", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cost"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok {
		t, err := expandSystemsdwanMembersGateway(d, v, "gateway", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("gateway6"); ok {
		t, err := expandSystemsdwanMembersGateway6(d, v, "gateway6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway6"] = t
		}
	}

	if v, ok := d.GetOk("ingress_spillover_threshold"); ok {
		t, err := expandSystemsdwanMembersIngressSpilloverThreshold(d, v, "ingress_spillover_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ingress-spillover-threshold"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandSystemsdwanMembersInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok {
		t, err := expandSystemsdwanMembersPriority(d, v, "priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("priority6"); ok {
		t, err := expandSystemsdwanMembersPriority6(d, v, "priority6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority6"] = t
		}
	}

	if v, ok := d.GetOk("seq_num"); ok {
		t, err := expandSystemsdwanMembersSeqNum(d, v, "seq_num", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seq-num"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok {
		t, err := expandSystemsdwanMembersSource(d, v, "source", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("source6"); ok {
		t, err := expandSystemsdwanMembersSource6(d, v, "source6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source6"] = t
		}
	}

	if v, ok := d.GetOk("spillover_threshold"); ok {
		t, err := expandSystemsdwanMembersSpilloverThreshold(d, v, "spillover_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spillover-threshold"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemsdwanMembersStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("volume_ratio"); ok {
		t, err := expandSystemsdwanMembersVolumeRatio(d, v, "volume_ratio", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["volume-ratio"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok {
		t, err := expandSystemsdwanMembersWeight(d, v, "weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	if v, ok := d.GetOk("zone"); ok {
		t, err := expandSystemsdwanMembersZone(d, v, "zone", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["zone"] = t
		}
	}

	return &obj, nil
}
