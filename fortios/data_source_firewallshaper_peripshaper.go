// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure per-IP traffic shaper.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceFirewallShaperPerIpShaper() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallShaperPerIpShaperRead,
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
			"max_bandwidth": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"bandwidth_unit": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_concurrent_session": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_concurrent_tcp_session": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_concurrent_udp_session": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"diffserv_forward": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"diffserv_reverse": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"diffservcode_forward": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"diffservcode_rev": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceFirewallShaperPerIpShaperRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("error describing FirewallShaperPerIpShaper: type error")
	}

	o, err := c.ReadFirewallShaperPerIpShaper(mkey, vdomparam, make(map[string][]string), 0)
	if err != nil {
		return fmt.Errorf("error describing FirewallShaperPerIpShaper: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectFirewallShaperPerIpShaper(d, o)
	if err != nil {
		return fmt.Errorf("error describing FirewallShaperPerIpShaper from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenFirewallShaperPerIpShaperName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallShaperPerIpShaperMaxBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallShaperPerIpShaperBandwidthUnit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallShaperPerIpShaperMaxConcurrentSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallShaperPerIpShaperMaxConcurrentTcpSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallShaperPerIpShaperMaxConcurrentUdpSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallShaperPerIpShaperDiffservForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallShaperPerIpShaperDiffservReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallShaperPerIpShaperDiffservcodeForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallShaperPerIpShaperDiffservcodeRev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectFirewallShaperPerIpShaper(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenFirewallShaperPerIpShaperName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("max_bandwidth", dataSourceFlattenFirewallShaperPerIpShaperMaxBandwidth(o["max-bandwidth"], d, "max_bandwidth")); err != nil {
		if !fortiAPIPatch(o["max-bandwidth"]) {
			return fmt.Errorf("error reading max_bandwidth: %v", err)
		}
	}

	if err = d.Set("bandwidth_unit", dataSourceFlattenFirewallShaperPerIpShaperBandwidthUnit(o["bandwidth-unit"], d, "bandwidth_unit")); err != nil {
		if !fortiAPIPatch(o["bandwidth-unit"]) {
			return fmt.Errorf("error reading bandwidth_unit: %v", err)
		}
	}

	if err = d.Set("max_concurrent_session", dataSourceFlattenFirewallShaperPerIpShaperMaxConcurrentSession(o["max-concurrent-session"], d, "max_concurrent_session")); err != nil {
		if !fortiAPIPatch(o["max-concurrent-session"]) {
			return fmt.Errorf("error reading max_concurrent_session: %v", err)
		}
	}

	if err = d.Set("max_concurrent_tcp_session", dataSourceFlattenFirewallShaperPerIpShaperMaxConcurrentTcpSession(o["max-concurrent-tcp-session"], d, "max_concurrent_tcp_session")); err != nil {
		if !fortiAPIPatch(o["max-concurrent-tcp-session"]) {
			return fmt.Errorf("error reading max_concurrent_tcp_session: %v", err)
		}
	}

	if err = d.Set("max_concurrent_udp_session", dataSourceFlattenFirewallShaperPerIpShaperMaxConcurrentUdpSession(o["max-concurrent-udp-session"], d, "max_concurrent_udp_session")); err != nil {
		if !fortiAPIPatch(o["max-concurrent-udp-session"]) {
			return fmt.Errorf("error reading max_concurrent_udp_session: %v", err)
		}
	}

	if err = d.Set("diffserv_forward", dataSourceFlattenFirewallShaperPerIpShaperDiffservForward(o["diffserv-forward"], d, "diffserv_forward")); err != nil {
		if !fortiAPIPatch(o["diffserv-forward"]) {
			return fmt.Errorf("error reading diffserv_forward: %v", err)
		}
	}

	if err = d.Set("diffserv_reverse", dataSourceFlattenFirewallShaperPerIpShaperDiffservReverse(o["diffserv-reverse"], d, "diffserv_reverse")); err != nil {
		if !fortiAPIPatch(o["diffserv-reverse"]) {
			return fmt.Errorf("error reading diffserv_reverse: %v", err)
		}
	}

	if err = d.Set("diffservcode_forward", dataSourceFlattenFirewallShaperPerIpShaperDiffservcodeForward(o["diffservcode-forward"], d, "diffservcode_forward")); err != nil {
		if !fortiAPIPatch(o["diffservcode-forward"]) {
			return fmt.Errorf("error reading diffservcode_forward: %v", err)
		}
	}

	if err = d.Set("diffservcode_rev", dataSourceFlattenFirewallShaperPerIpShaperDiffservcodeRev(o["diffservcode-rev"], d, "diffservcode_rev")); err != nil {
		if !fortiAPIPatch(o["diffservcode-rev"]) {
			return fmt.Errorf("error reading diffservcode_rev: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenFirewallShaperPerIpShaperFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
