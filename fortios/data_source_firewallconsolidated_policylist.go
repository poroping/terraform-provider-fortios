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

func dataSourceFirewallConsolidatedPolicyList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallConsolidatedPolicyListRead,

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
			"policyidlist": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func dataSourceFirewallConsolidatedPolicyListRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.GenericGroupRead("/api/v2/cmdb/firewall.consolidated/policy", filter, vdomparam, 0)
	if err != nil {
		return fmt.Errorf("error describing FirewallConsolidatedPolicy: %v", err)
	}

	var tmps []int
	if o != nil {
		if len(o) == 0 || o[0] == nil {
			return nil
		}

		for _, r := range o {
			i := r.(map[string]interface{})

			if _, ok := i["policyid"]; ok {
				tmps = append(tmps, fortiIntValue(i["policyid"]))
			}
		}
	}
	d.Set("policyidlist", tmps)

	d.SetId("DataSourceFirewallConsolidatedPolicyList" + filter)

	return nil
}
