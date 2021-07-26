// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemsdwanNeighbor() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemsdwanNeighborCreate,
		Read:   resourceSystemsdwanNeighborRead,
		Update: resourceSystemsdwanNeighborUpdate,
		Delete: resourceSystemsdwanNeighborDelete,

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
			"health_check": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "SD-WAN health-check name.",
				Optional:     true,
				Computed:     true,
			},
			"ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 45),
				Description:  "IP/IPv6 address of neighbor.",
				ForceNew:     true,
				Required:     true,
			},
			"member": {
				Type: schema.TypeInt,

				Description: "Member sequence number.",
				Optional:    true,
				Computed:    true,
			},
			"mode": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"sla", "speedtest"}),
				Description:  "What metric to select the neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"role": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"standalone", "primary", "secondary"}),
				Description:  "Role of neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"sla_id": {
				Type: schema.TypeInt,

				Description: "SLA ID.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemsdwanNeighborCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemsdwanNeighbor(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating SystemsdwanNeighbor resource while getting object: %v", err)
	}

	o, err := c.CreateSystemsdwanNeighbor(obj, vdomparam, urlparams, batchid)

	if err != nil {
		return fmt.Errorf("error creating SystemsdwanNeighbor resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemsdwanNeighbor")
	}

	return resourceSystemsdwanNeighborRead(d, m)
}

func resourceSystemsdwanNeighborUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemsdwanNeighbor(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating SystemsdwanNeighbor resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemsdwanNeighbor(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating SystemsdwanNeighbor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemsdwanNeighbor")
	}

	return resourceSystemsdwanNeighborRead(d, m)
}

func resourceSystemsdwanNeighborDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteSystemsdwanNeighbor(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting SystemsdwanNeighbor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemsdwanNeighborRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemsdwanNeighbor(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading SystemsdwanNeighbor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemsdwanNeighbor(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading SystemsdwanNeighbor resource from API: %v", err)
	}
	return nil
}

func flattenSystemsdwanNeighborHealthCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanNeighborIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanNeighborMember(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanNeighborMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanNeighborRole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemsdwanNeighborSlaId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemsdwanNeighbor(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("health_check", flattenSystemsdwanNeighborHealthCheck(o["health-check"], d, "health_check", sv)); err != nil {
		if !fortiAPIPatch(o["health-check"]) {
			return fmt.Errorf("error reading health_check: %v", err)
		}
	}

	if err = d.Set("ip", flattenSystemsdwanNeighborIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("error reading ip: %v", err)
		}
	}

	if err = d.Set("member", flattenSystemsdwanNeighborMember(o["member"], d, "member", sv)); err != nil {
		if !fortiAPIPatch(o["member"]) {
			return fmt.Errorf("error reading member: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemsdwanNeighborMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("error reading mode: %v", err)
		}
	}

	if err = d.Set("role", flattenSystemsdwanNeighborRole(o["role"], d, "role", sv)); err != nil {
		if !fortiAPIPatch(o["role"]) {
			return fmt.Errorf("error reading role: %v", err)
		}
	}

	if err = d.Set("sla_id", flattenSystemsdwanNeighborSlaId(o["sla-id"], d, "sla_id", sv)); err != nil {
		if !fortiAPIPatch(o["sla-id"]) {
			return fmt.Errorf("error reading sla_id: %v", err)
		}
	}

	return nil
}

func expandSystemsdwanNeighborHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanNeighborIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanNeighborMember(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanNeighborMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanNeighborRole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemsdwanNeighborSlaId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemsdwanNeighbor(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("health_check"); ok {
		t, err := expandSystemsdwanNeighborHealthCheck(d, v, "health_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandSystemsdwanNeighborIp(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok {
		t, err := expandSystemsdwanNeighborMember(d, v, "member", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandSystemsdwanNeighborMode(d, v, "mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok {
		t, err := expandSystemsdwanNeighborRole(d, v, "role", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("sla_id"); ok {
		t, err := expandSystemsdwanNeighborSlaId(d, v, "sla_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-id"] = t
		}
	}

	return &obj, nil
}
