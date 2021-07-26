// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: FortiGate interfaces added to the SD-WAN.

package fortios

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSystemsdwanMembers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemsdwanMembersRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comments.",
				Computed:    true,
			},
			"cost": {
				Type:        schema.TypeInt,
				Description: "Cost of this interface for services in SLA mode (0 - 4294967295, default = 0).",
				Computed:    true,
			},
			"gateway": {
				Type:        schema.TypeString,
				Description: "The default gateway for this interface. Usually the default gateway of the Internet service provider that this interface is connected to.",
				Computed:    true,
			},
			"gateway6": {
				Type:        schema.TypeString,
				Description: "IPv6 gateway.",
				Computed:    true,
			},
			"ingress_spillover_threshold": {
				Type:        schema.TypeInt,
				Description: "Ingress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Interface name.",
				Computed:    true,
			},
			"priority": {
				Type:        schema.TypeInt,
				Description: "Priority of the interface for IPv4 (0 - 65535, default = 0). Used for SD-WAN rules or priority rules.",
				Computed:    true,
			},
			"priority6": {
				Type:        schema.TypeInt,
				Description: "Priority of the interface for IPv6 (1 - 65535, default = 1024). Used for SD-WAN rules or priority rules.",
				Computed:    true,
			},
			"seq_num": {
				Type:        schema.TypeInt,
				Description: "Sequence number(1-512).",
				Required:    true,
			},
			"source": {
				Type:        schema.TypeString,
				Description: "Source IP address used in the health-check packet to the server.",
				Computed:    true,
			},
			"source6": {
				Type:        schema.TypeString,
				Description: "Source IPv6 address used in the health-check packet to the server.",
				Computed:    true,
			},
			"spillover_threshold": {
				Type:        schema.TypeInt,
				Description: "Egress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable this interface in the SD-WAN.",
				Computed:    true,
			},
			"volume_ratio": {
				Type:        schema.TypeInt,
				Description: "Measured volume ratio (this value / sum of all values = percentage of link volume, 1 - 255).",
				Computed:    true,
			},
			"weight": {
				Type:        schema.TypeInt,
				Description: "Weight of this interface for weighted load balancing. (1 - 255) More traffic is directed to interfaces with higher weights.",
				Computed:    true,
			},
			"zone": {
				Type:        schema.TypeString,
				Description: "Zone name.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemsdwanMembersRead(d *schema.ResourceData, m interface{}) error {
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
	key := "seq-num"
	t := d.Get(key)
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("error describing SystemsdwanMembers: type error")
	}

	o, err := c.ReadSystemsdwanMembers(mkey, vdomparam, urlparams, 0)
	if err != nil {
		return fmt.Errorf("error describing SystemsdwanMembers: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemsdwanMembers(d, o)
	if err != nil {
		return fmt.Errorf("error describing SystemsdwanMembers from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemsdwanMembersComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanMembersCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanMembersGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanMembersGateway6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanMembersIngressSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanMembersInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanMembersPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanMembersPriority6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanMembersSeqNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanMembersSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanMembersSource6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanMembersSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanMembersStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanMembersVolumeRatio(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanMembersWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanMembersZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemsdwanMembers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", dataSourceFlattenSystemsdwanMembersComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("error reading comment: %v", err)
		}
	}

	if err = d.Set("cost", dataSourceFlattenSystemsdwanMembersCost(o["cost"], d, "cost")); err != nil {
		if !fortiAPIPatch(o["cost"]) {
			return fmt.Errorf("error reading cost: %v", err)
		}
	}

	if err = d.Set("gateway", dataSourceFlattenSystemsdwanMembersGateway(o["gateway"], d, "gateway")); err != nil {
		if !fortiAPIPatch(o["gateway"]) {
			return fmt.Errorf("error reading gateway: %v", err)
		}
	}

	if err = d.Set("gateway6", dataSourceFlattenSystemsdwanMembersGateway6(o["gateway6"], d, "gateway6")); err != nil {
		if !fortiAPIPatch(o["gateway6"]) {
			return fmt.Errorf("error reading gateway6: %v", err)
		}
	}

	if err = d.Set("ingress_spillover_threshold", dataSourceFlattenSystemsdwanMembersIngressSpilloverThreshold(o["ingress-spillover-threshold"], d, "ingress_spillover_threshold")); err != nil {
		if !fortiAPIPatch(o["ingress-spillover-threshold"]) {
			return fmt.Errorf("error reading ingress_spillover_threshold: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenSystemsdwanMembersInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("error reading interface: %v", err)
		}
	}

	if err = d.Set("priority", dataSourceFlattenSystemsdwanMembersPriority(o["priority"], d, "priority")); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("error reading priority: %v", err)
		}
	}

	if err = d.Set("priority6", dataSourceFlattenSystemsdwanMembersPriority6(o["priority6"], d, "priority6")); err != nil {
		if !fortiAPIPatch(o["priority6"]) {
			return fmt.Errorf("error reading priority6: %v", err)
		}
	}

	if err = d.Set("seq_num", dataSourceFlattenSystemsdwanMembersSeqNum(o["seq-num"], d, "seq_num")); err != nil {
		if !fortiAPIPatch(o["seq-num"]) {
			return fmt.Errorf("error reading seq_num: %v", err)
		}
	}

	if err = d.Set("source", dataSourceFlattenSystemsdwanMembersSource(o["source"], d, "source")); err != nil {
		if !fortiAPIPatch(o["source"]) {
			return fmt.Errorf("error reading source: %v", err)
		}
	}

	if err = d.Set("source6", dataSourceFlattenSystemsdwanMembersSource6(o["source6"], d, "source6")); err != nil {
		if !fortiAPIPatch(o["source6"]) {
			return fmt.Errorf("error reading source6: %v", err)
		}
	}

	if err = d.Set("spillover_threshold", dataSourceFlattenSystemsdwanMembersSpilloverThreshold(o["spillover-threshold"], d, "spillover_threshold")); err != nil {
		if !fortiAPIPatch(o["spillover-threshold"]) {
			return fmt.Errorf("error reading spillover_threshold: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenSystemsdwanMembersStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("error reading status: %v", err)
		}
	}

	if err = d.Set("volume_ratio", dataSourceFlattenSystemsdwanMembersVolumeRatio(o["volume-ratio"], d, "volume_ratio")); err != nil {
		if !fortiAPIPatch(o["volume-ratio"]) {
			return fmt.Errorf("error reading volume_ratio: %v", err)
		}
	}

	if err = d.Set("weight", dataSourceFlattenSystemsdwanMembersWeight(o["weight"], d, "weight")); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("error reading weight: %v", err)
		}
	}

	if err = d.Set("zone", dataSourceFlattenSystemsdwanMembersZone(o["zone"], d, "zone")); err != nil {
		if !fortiAPIPatch(o["zone"]) {
			return fmt.Errorf("error reading zone: %v", err)
		}
	}

	return nil
}
