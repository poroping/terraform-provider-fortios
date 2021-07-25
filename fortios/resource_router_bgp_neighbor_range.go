// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: BGP neighbor range table.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterbgpNeighborRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterbgpNeighborRangeCreate,
		Read:   resourceRouterbgpNeighborRangeRead,
		Update: resourceRouterbgpNeighborRangeUpdate,
		Delete: resourceRouterbgpNeighborRangeDelete,

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
			"fosid": {
				Type: schema.TypeInt,

				Description: "Neighbor range ID.",
				Optional:    true,
				Computed:    true,
			},
			"max_neighbor_num": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1000),
				Description:  "Maximum number of neighbors.",
				Optional:     true,
				Computed:     true,
			},
			"neighbor_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Description:  "Neighbor group name.",
				Optional:     true,
				Computed:     true,
			},
			"prefix": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateIPv4Classnet,
				Description:  "Neighbor range prefix.",
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceRouterbgpNeighborRangeCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterbgpNeighborRange(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating RouterbgpNeighborRange resource while getting object: %v", err)
	}

	o, err := c.CreateRouterbgpNeighborRange(obj, vdomparam, urlparams, batchid)

	if err != nil {
		return fmt.Errorf("error creating RouterbgpNeighborRange resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterbgpNeighborRange")
	}

	return resourceRouterbgpNeighborRangeRead(d, m)
}

func resourceRouterbgpNeighborRangeUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterbgpNeighborRange(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating RouterbgpNeighborRange resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterbgpNeighborRange(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating RouterbgpNeighborRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterbgpNeighborRange")
	}

	return resourceRouterbgpNeighborRangeRead(d, m)
}

func resourceRouterbgpNeighborRangeDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteRouterbgpNeighborRange(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting RouterbgpNeighborRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterbgpNeighborRangeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterbgpNeighborRange(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading RouterbgpNeighborRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterbgpNeighborRange(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading RouterbgpNeighborRange resource from API: %v", err)
	}
	return nil
}

func flattenRouterbgpNeighborRangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRangeMaxNeighborNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRangeNeighborGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRangePrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}
	return v
}

func refreshObjectRouterbgpNeighborRange(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenRouterbgpNeighborRangeId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("error reading fosid: %v", err)
		}
	}

	if err = d.Set("max_neighbor_num", flattenRouterbgpNeighborRangeMaxNeighborNum(o["max-neighbor-num"], d, "max_neighbor_num", sv)); err != nil {
		if !fortiAPIPatch(o["max-neighbor-num"]) {
			return fmt.Errorf("error reading max_neighbor_num: %v", err)
		}
	}

	if err = d.Set("neighbor_group", flattenRouterbgpNeighborRangeNeighborGroup(o["neighbor-group"], d, "neighbor_group", sv)); err != nil {
		if !fortiAPIPatch(o["neighbor-group"]) {
			return fmt.Errorf("error reading neighbor_group: %v", err)
		}
	}

	if err = d.Set("prefix", flattenRouterbgpNeighborRangePrefix(o["prefix"], d, "prefix", sv)); err != nil {
		if !fortiAPIPatch(o["prefix"]) {
			return fmt.Errorf("error reading prefix: %v", err)
		}
	}

	return nil
}

func expandRouterbgpNeighborRangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRangeMaxNeighborNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRangeNeighborGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRangePrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterbgpNeighborRange(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandRouterbgpNeighborRangeId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("max_neighbor_num"); ok {
		t, err := expandRouterbgpNeighborRangeMaxNeighborNum(d, v, "max_neighbor_num", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-neighbor-num"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_group"); ok {
		t, err := expandRouterbgpNeighborRangeNeighborGroup(d, v, "neighbor_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-group"] = t
		}
	}

	if v, ok := d.GetOk("prefix"); ok {
		t, err := expandRouterbgpNeighborRangePrefix(d, v, "prefix", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	return &obj, nil
}
