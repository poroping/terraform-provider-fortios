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

func dataSourceFirewallShaperTrafficShaperList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallShaperTrafficShaperListRead,

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
			"namelist": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceFirewallShaperTrafficShaperListRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.GenericGroupRead("/api/v2/cmdb/firewall.shaper/traffic-shaper", filter, vdomparam, 0)
	if err != nil {
		return fmt.Errorf("error describing FirewallShaperTrafficShaper: %v", err)
	}

	var tmps []string
	if o != nil {
		if len(o) == 0 || o[0] == nil {
			return nil
		}

		for _, r := range o {
			i := r.(map[string]interface{})

			if _, ok := i["name"]; ok {
				tmps = append(tmps, fortiStringValue(i["name"]))
			}
		}
	}
	d.Set("namelist", tmps)

	d.SetId("DataSourceFirewallShaperTrafficShaperList" + filter)

	return nil
}
