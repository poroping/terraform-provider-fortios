// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSystemIpv6NeighborCacheList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemIpv6NeighborCacheListRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"filter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosidlist": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func dataSourceSystemIpv6NeighborCacheListRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	filter := d.Get("filter").(string)
	if filter != "" {
		filter = escapeFilter(filter)
	}

	o, err := c.GenericGroupRead("/api/v2/cmdb/system/ipv6-neighbor-cache", filter, vdomparam, 0)
	if err != nil {
		return fmt.Errorf("error describing SystemIpv6NeighborCache: %v", err)
	}

	var tmps []int
	if o != nil {
		if len(o) == 0 || o[0] == nil {
			return nil
		}

		for _, r := range o {
			i := r.(map[string]interface{})

			if _, ok := i["id"]; ok {
				tmps = append(tmps, fortiIntValue(i["id"]))
			}
		}
	}
	d.Set("fosidlist", tmps)

	d.SetId("DataSourceSystemIpv6NeighborCacheList" + filter)

	return nil
}
