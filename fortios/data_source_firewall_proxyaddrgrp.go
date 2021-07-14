// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Web proxy address group configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceFirewallProxyAddrgrp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallProxyAddrgrpRead,
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
			"type": {
				Type:     schema.TypeString,
				Computed: true,
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
			"tagging": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"category": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tags": {
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
					},
				},
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

func dataSourceFirewallProxyAddrgrpRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("error describing FirewallProxyAddrgrp: type error")
	}

	o, err := c.ReadFirewallProxyAddrgrp(mkey, vdomparam, make(map[string][]string), 0)
	if err != nil {
		return fmt.Errorf("error describing FirewallProxyAddrgrp: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectFirewallProxyAddrgrp(d, o)
	if err != nil {
		return fmt.Errorf("error describing FirewallProxyAddrgrp from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenFirewallProxyAddrgrpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddrgrpType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddrgrpUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddrgrpMember(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallProxyAddrgrpMemberName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallProxyAddrgrpMemberName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddrgrpColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddrgrpTagging(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallProxyAddrgrpTaggingName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			tmp["category"] = dataSourceFlattenFirewallProxyAddrgrpTaggingCategory(i["category"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {
			tmp["tags"] = dataSourceFlattenFirewallProxyAddrgrpTaggingTags(i["tags"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallProxyAddrgrpTaggingName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddrgrpTaggingCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddrgrpTaggingTags(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallProxyAddrgrpTaggingTagsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallProxyAddrgrpTaggingTagsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddrgrpComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddrgrpVisibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectFirewallProxyAddrgrp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenFirewallProxyAddrgrpName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenFirewallProxyAddrgrpType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("error reading type: %v", err)
		}
	}

	if err = d.Set("uuid", dataSourceFlattenFirewallProxyAddrgrpUuid(o["uuid"], d, "uuid")); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("error reading uuid: %v", err)
		}
	}

	if err = d.Set("member", dataSourceFlattenFirewallProxyAddrgrpMember(o["member"], d, "member")); err != nil {
		if !fortiAPIPatch(o["member"]) {
			return fmt.Errorf("error reading member: %v", err)
		}
	}

	if err = d.Set("color", dataSourceFlattenFirewallProxyAddrgrpColor(o["color"], d, "color")); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("error reading color: %v", err)
		}
	}

	if err = d.Set("tagging", dataSourceFlattenFirewallProxyAddrgrpTagging(o["tagging"], d, "tagging")); err != nil {
		if !fortiAPIPatch(o["tagging"]) {
			return fmt.Errorf("error reading tagging: %v", err)
		}
	}

	if err = d.Set("comment", dataSourceFlattenFirewallProxyAddrgrpComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("error reading comment: %v", err)
		}
	}

	if err = d.Set("visibility", dataSourceFlattenFirewallProxyAddrgrpVisibility(o["visibility"], d, "visibility")); err != nil {
		if !fortiAPIPatch(o["visibility"]) {
			return fmt.Errorf("error reading visibility: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenFirewallProxyAddrgrpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
