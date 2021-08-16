// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 | v6.4.2 | v6.2.7 merged schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: Configure IPv6 static routing tables.

package fortios

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceRouterStatic6() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterStatic6Read,
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"bfd": {
				Type:        schema.TypeString,
				Description: "Enable/disable Bidirectional Forwarding Detection (BFD).",
				Computed:    true,
			},
			"blackhole": {
				Type:        schema.TypeString,
				Description: "Enable/disable black hole.",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Optional comments.",
				Computed:    true,
			},
			"device": {
				Type:        schema.TypeString,
				Description: "Gateway out interface or tunnel.",
				Computed:    true,
			},
			"devindex": {
				Type:        schema.TypeInt,
				Description: "Device index (0 - 4294967295).",
				Computed:    true,
			},
			"distance": {
				Type:        schema.TypeInt,
				Description: "Administrative distance (1 - 255).",
				Computed:    true,
			},
			"dst": {
				Type:        schema.TypeString,
				Description: "Destination IPv6 prefix.",
				Computed:    true,
			},
			"dynamic_gateway": {
				Type:        schema.TypeString,
				Description: "Enable use of dynamic gateway retrieved from Router Advertisement (RA).",
				Computed:    true,
			},
			"gateway": {
				Type:        schema.TypeString,
				Description: "IPv6 address of the gateway.",
				Computed:    true,
			},
			"link_monitor_exempt": {
				Type:        schema.TypeString,
				Description: "Enable/disable withdrawal of this static route when link monitor or health check is down.",
				Computed:    true,
			},
			"priority": {
				Type:        schema.TypeInt,
				Description: "Administrative priority (1 - 65535).",
				Computed:    true,
			},
			"sdwan": {
				Type:        schema.TypeString,
				Description: "Enable/disable egress through the SD-WAN.",
				Computed:    true,
			},
			"sdwan_zone": {
				Type:        schema.TypeList,
				Description: "Choose SD-WAN Zone.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "SD-WAN zone name.",
							Computed:    true,
						},
					},
				},
			},
			"seq_num": {
				Type:        schema.TypeInt,
				Description: "Sequence number.",
				Required:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable this static route.",
				Computed:    true,
			},
			"virtual_wan_link": {
				Type:        schema.TypeString,
				Description: "Enable/disable egress through the virtual-wan-link.",
				Computed:    true,
			},
			"vrf": {
				Type:        schema.TypeInt,
				Description: "Virtual Routing Forwarding ID.",
				Computed:    true,
			},
		},
	}
}

func dataSourceRouterStatic6Read(d *schema.ResourceData, m interface{}) error {
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
	key := "seq_num"
	t := d.Get(key)
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("error describing RouterStatic6: type error")
	}

	o, err := c.ReadRouterStatic6(mkey, vdomparam, urlparams, 0)
	if err != nil {
		return fmt.Errorf("error describing RouterStatic6: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterStatic6(d, o)
	if err != nil {
		return fmt.Errorf("error describing RouterStatic6 from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterStatic6Bfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6Blackhole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6Comment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6Device(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6Devindex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6Distance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6Dst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6DynamicGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6Gateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6LinkMonitorExempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6Priority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6Sdwan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6SdwanZone(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenRouterStatic6SdwanZoneName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterStatic6SdwanZoneName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6SeqNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6VirtualWanLink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6Vrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterStatic6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("bfd", dataSourceFlattenRouterStatic6Bfd(o["bfd"], d, "bfd")); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("error reading bfd: %v", err)
		}
	}

	if err = d.Set("blackhole", dataSourceFlattenRouterStatic6Blackhole(o["blackhole"], d, "blackhole")); err != nil {
		if !fortiAPIPatch(o["blackhole"]) {
			return fmt.Errorf("error reading blackhole: %v", err)
		}
	}

	if err = d.Set("comment", dataSourceFlattenRouterStatic6Comment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("error reading comment: %v", err)
		}
	}

	if err = d.Set("device", dataSourceFlattenRouterStatic6Device(o["device"], d, "device")); err != nil {
		if !fortiAPIPatch(o["device"]) {
			return fmt.Errorf("error reading device: %v", err)
		}
	}

	if err = d.Set("devindex", dataSourceFlattenRouterStatic6Devindex(o["devindex"], d, "devindex")); err != nil {
		if !fortiAPIPatch(o["devindex"]) {
			return fmt.Errorf("error reading devindex: %v", err)
		}
	}

	if err = d.Set("distance", dataSourceFlattenRouterStatic6Distance(o["distance"], d, "distance")); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("error reading distance: %v", err)
		}
	}

	if err = d.Set("dst", dataSourceFlattenRouterStatic6Dst(o["dst"], d, "dst")); err != nil {
		if !fortiAPIPatch(o["dst"]) {
			return fmt.Errorf("error reading dst: %v", err)
		}
	}

	if err = d.Set("dynamic_gateway", dataSourceFlattenRouterStatic6DynamicGateway(o["dynamic-gateway"], d, "dynamic_gateway")); err != nil {
		if !fortiAPIPatch(o["dynamic-gateway"]) {
			return fmt.Errorf("error reading dynamic_gateway: %v", err)
		}
	}

	if err = d.Set("gateway", dataSourceFlattenRouterStatic6Gateway(o["gateway"], d, "gateway")); err != nil {
		if !fortiAPIPatch(o["gateway"]) {
			return fmt.Errorf("error reading gateway: %v", err)
		}
	}

	if err = d.Set("link_monitor_exempt", dataSourceFlattenRouterStatic6LinkMonitorExempt(o["link-monitor-exempt"], d, "link_monitor_exempt")); err != nil {
		if !fortiAPIPatch(o["link-monitor-exempt"]) {
			return fmt.Errorf("error reading link_monitor_exempt: %v", err)
		}
	}

	if err = d.Set("priority", dataSourceFlattenRouterStatic6Priority(o["priority"], d, "priority")); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("error reading priority: %v", err)
		}
	}

	if err = d.Set("sdwan", dataSourceFlattenRouterStatic6Sdwan(o["sdwan"], d, "sdwan")); err != nil {
		if !fortiAPIPatch(o["sdwan"]) {
			return fmt.Errorf("error reading sdwan: %v", err)
		}
	}

	if err = d.Set("sdwan_zone", dataSourceFlattenRouterStatic6SdwanZone(o["sdwan-zone"], d, "sdwan_zone")); err != nil {
		if !fortiAPIPatch(o["sdwan-zone"]) {
			return fmt.Errorf("error reading sdwan_zone: %v", err)
		}
	}

	if err = d.Set("seq_num", dataSourceFlattenRouterStatic6SeqNum(o["seq-num"], d, "seq_num")); err != nil {
		if !fortiAPIPatch(o["seq-num"]) {
			return fmt.Errorf("error reading seq_num: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenRouterStatic6Status(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("error reading status: %v", err)
		}
	}

	if err = d.Set("virtual_wan_link", dataSourceFlattenRouterStatic6VirtualWanLink(o["virtual-wan-link"], d, "virtual_wan_link")); err != nil {
		if !fortiAPIPatch(o["virtual-wan-link"]) {
			return fmt.Errorf("error reading virtual_wan_link: %v", err)
		}
	}

	if err = d.Set("vrf", dataSourceFlattenRouterStatic6Vrf(o["vrf"], d, "vrf")); err != nil {
		if !fortiAPIPatch(o["vrf"]) {
			return fmt.Errorf("error reading vrf: %v", err)
		}
	}

	return nil
}
