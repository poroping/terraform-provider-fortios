// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: BGP neighbor range table.

package fortios

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceRouterbgpNeighborRange() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterbgpNeighborRangeRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "Neighbor range ID.",
				Computed:    true,
			},
			"max_neighbor_num": {
				Type:        schema.TypeInt,
				Description: "Maximum number of neighbors.",
				Computed:    true,
			},
			"neighbor_group": {
				Type:        schema.TypeString,
				Description: "Neighbor group name.",
				Computed:    true,
			},
			"prefix": {
				Type:        schema.TypeString,
				Description: "Neighbor range prefix.",
				Computed:    true,
			},
		},
	}
}

func dataSourceRouterbgpNeighborRangeRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("error describing RouterbgpNeighborRange: type error")
	}

	o, err := c.ReadRouterbgpNeighborRange(mkey, vdomparam, urlparams, 0)
	if err != nil {
		return fmt.Errorf("error describing RouterbgpNeighborRange: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterbgpNeighborRange(d, o)
	if err != nil {
		return fmt.Errorf("error describing RouterbgpNeighborRange from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterbgpNeighborRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRangeMaxNeighborNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRangeNeighborGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRangePrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}
	return v
}

func dataSourceRefreshObjectRouterbgpNeighborRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", dataSourceFlattenRouterbgpNeighborRangeId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("error reading fosid: %v", err)
		}
	}

	if err = d.Set("max_neighbor_num", dataSourceFlattenRouterbgpNeighborRangeMaxNeighborNum(o["max-neighbor-num"], d, "max_neighbor_num")); err != nil {
		if !fortiAPIPatch(o["max-neighbor-num"]) {
			return fmt.Errorf("error reading max_neighbor_num: %v", err)
		}
	}

	if err = d.Set("neighbor_group", dataSourceFlattenRouterbgpNeighborRangeNeighborGroup(o["neighbor-group"], d, "neighbor_group")); err != nil {
		if !fortiAPIPatch(o["neighbor-group"]) {
			return fmt.Errorf("error reading neighbor_group: %v", err)
		}
	}

	if err = d.Set("prefix", dataSourceFlattenRouterbgpNeighborRangePrefix(o["prefix"], d, "prefix")); err != nil {
		if !fortiAPIPatch(o["prefix"]) {
			return fmt.Errorf("error reading prefix: %v", err)
		}
	}

	return nil
}
