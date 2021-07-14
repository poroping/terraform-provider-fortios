// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Config global Wildcard FQDN address groups.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceFirewallWildcardFqdnGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallWildcardFqdnGroupRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"member": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"color": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"visibility": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceFirewallWildcardFqdnGroupRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("error describing FirewallWildcardFqdnGroup: type error")
	}

	o, err := c.ReadFirewallWildcardFqdnGroup(mkey, vdomparam, make(map[string][]string), 0)
	if err != nil {
		return fmt.Errorf("error describing FirewallWildcardFqdnGroup: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectFirewallWildcardFqdnGroup(d, o)
	if err != nil {
		return fmt.Errorf("error describing FirewallWildcardFqdnGroup from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenFirewallWildcardFqdnGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallWildcardFqdnGroupUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallWildcardFqdnGroupMember(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallWildcardFqdnGroupMemberName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallWildcardFqdnGroupMemberName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallWildcardFqdnGroupColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallWildcardFqdnGroupComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallWildcardFqdnGroupVisibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectFirewallWildcardFqdnGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenFirewallWildcardFqdnGroupName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("uuid", dataSourceFlattenFirewallWildcardFqdnGroupUuid(o["uuid"], d, "uuid")); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("error reading uuid: %v", err)
		}
	}

	if err = d.Set("member", dataSourceFlattenFirewallWildcardFqdnGroupMember(o["member"], d, "member")); err != nil {
		if !fortiAPIPatch(o["member"]) {
			return fmt.Errorf("error reading member: %v", err)
		}
	}

	if err = d.Set("color", dataSourceFlattenFirewallWildcardFqdnGroupColor(o["color"], d, "color")); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("error reading color: %v", err)
		}
	}

	if err = d.Set("comment", dataSourceFlattenFirewallWildcardFqdnGroupComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("error reading comment: %v", err)
		}
	}

	if err = d.Set("visibility", dataSourceFlattenFirewallWildcardFqdnGroupVisibility(o["visibility"], d, "visibility")); err != nil {
		if !fortiAPIPatch(o["visibility"]) {
			return fmt.Errorf("error reading visibility: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenFirewallWildcardFqdnGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
