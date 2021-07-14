// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Type of Service (ToS) based priority table to set network traffic priorities.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemTosBasedPriority() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemTosBasedPriorityRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"fosid": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"tos": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemTosBasedPriorityRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""

	t := d.Get("fosid")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("error describing SystemTosBasedPriority: type error")
	}

	o, err := c.ReadSystemTosBasedPriority(mkey, vdomparam, make(map[string][]string), 0)
	if err != nil {
		return fmt.Errorf("error describing SystemTosBasedPriority: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemTosBasedPriority(d, o)
	if err != nil {
		return fmt.Errorf("error describing SystemTosBasedPriority from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemTosBasedPriorityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemTosBasedPriorityTos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemTosBasedPriorityPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemTosBasedPriority(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", dataSourceFlattenSystemTosBasedPriorityId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("error reading fosid: %v", err)
		}
	}

	if err = d.Set("tos", dataSourceFlattenSystemTosBasedPriorityTos(o["tos"], d, "tos")); err != nil {
		if !fortiAPIPatch(o["tos"]) {
			return fmt.Errorf("error reading tos: %v", err)
		}
	}

	if err = d.Set("priority", dataSourceFlattenSystemTosBasedPriorityPriority(o["priority"], d, "priority")); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("error reading priority: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemTosBasedPriorityFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
