// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status.

package fortios

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSystemsdwanNeighbor() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemsdwanNeighborRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"health_check": {
				Type:        schema.TypeString,
				Description: "SD-WAN health-check name.",
				Computed:    true,
			},
			"ip": {
				Type:        schema.TypeString,
				Description: "IP/IPv6 address of neighbor.",
				Required:    true,
			},
			"member": {
				Type:        schema.TypeInt,
				Description: "Member sequence number.",
				Computed:    true,
			},
			"mode": {
				Type:        schema.TypeString,
				Description: "What metric to select the neighbor.",
				Computed:    true,
			},
			"role": {
				Type:        schema.TypeString,
				Description: "Role of neighbor.",
				Computed:    true,
			},
			"sla_id": {
				Type:        schema.TypeInt,
				Description: "SLA ID.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemsdwanNeighborRead(d *schema.ResourceData, m interface{}) error {
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
	key := "ip"
	t := d.Get(key)
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("error describing SystemsdwanNeighbor: type error")
	}

	o, err := c.ReadSystemsdwanNeighbor(mkey, vdomparam, urlparams, 0)
	if err != nil {
		return fmt.Errorf("error describing SystemsdwanNeighbor: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemsdwanNeighbor(d, o)
	if err != nil {
		return fmt.Errorf("error describing SystemsdwanNeighbor from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemsdwanNeighborHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanNeighborIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanNeighborMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanNeighborMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanNeighborRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemsdwanNeighborSlaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemsdwanNeighbor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("health_check", dataSourceFlattenSystemsdwanNeighborHealthCheck(o["health-check"], d, "health_check")); err != nil {
		if !fortiAPIPatch(o["health-check"]) {
			return fmt.Errorf("error reading health_check: %v", err)
		}
	}

	if err = d.Set("ip", dataSourceFlattenSystemsdwanNeighborIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("error reading ip: %v", err)
		}
	}

	if err = d.Set("member", dataSourceFlattenSystemsdwanNeighborMember(o["member"], d, "member")); err != nil {
		if !fortiAPIPatch(o["member"]) {
			return fmt.Errorf("error reading member: %v", err)
		}
	}

	if err = d.Set("mode", dataSourceFlattenSystemsdwanNeighborMode(o["mode"], d, "mode")); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("error reading mode: %v", err)
		}
	}

	if err = d.Set("role", dataSourceFlattenSystemsdwanNeighborRole(o["role"], d, "role")); err != nil {
		if !fortiAPIPatch(o["role"]) {
			return fmt.Errorf("error reading role: %v", err)
		}
	}

	if err = d.Set("sla_id", dataSourceFlattenSystemsdwanNeighborSlaId(o["sla-id"], d, "sla_id")); err != nil {
		if !fortiAPIPatch(o["sla-id"]) {
			return fmt.Errorf("error reading sla_id: %v", err)
		}
	}

	return nil
}
