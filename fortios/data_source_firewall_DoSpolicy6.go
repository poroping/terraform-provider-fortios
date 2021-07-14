// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6 DoS policies.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceFirewallDosPolicy6() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallDosPolicy6Read,
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"policyid": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"srcaddr": {
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
			"dstaddr": {
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
			"service": {
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
			"anomaly": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"log": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"action": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"quarantine": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"quarantine_expiry": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"quarantine_log": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"threshold": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"thresholddefault": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceFirewallDosPolicy6Read(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""

	t := d.Get("policyid")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("error describing FirewallDosPolicy6: type error")
	}

	o, err := c.ReadFirewallDosPolicy6(mkey, vdomparam, make(map[string][]string), 0)
	if err != nil {
		return fmt.Errorf("error describing FirewallDosPolicy6: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectFirewallDosPolicy6(d, o)
	if err != nil {
		return fmt.Errorf("error describing FirewallDosPolicy6 from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenFirewallDosPolicy6Policyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallDosPolicy6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallDosPolicy6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallDosPolicy6Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallDosPolicy6Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallDosPolicy6Srcaddr(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallDosPolicy6SrcaddrName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallDosPolicy6SrcaddrName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallDosPolicy6Dstaddr(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallDosPolicy6DstaddrName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallDosPolicy6DstaddrName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallDosPolicy6Service(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallDosPolicy6ServiceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallDosPolicy6ServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallDosPolicy6Anomaly(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallDosPolicy6AnomalyName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = dataSourceFlattenFirewallDosPolicy6AnomalyStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			tmp["log"] = dataSourceFlattenFirewallDosPolicy6AnomalyLog(i["log"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = dataSourceFlattenFirewallDosPolicy6AnomalyAction(i["action"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := i["quarantine"]; ok {
			tmp["quarantine"] = dataSourceFlattenFirewallDosPolicy6AnomalyQuarantine(i["quarantine"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := i["quarantine-expiry"]; ok {
			tmp["quarantine_expiry"] = dataSourceFlattenFirewallDosPolicy6AnomalyQuarantineExpiry(i["quarantine-expiry"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := i["quarantine-log"]; ok {
			tmp["quarantine_log"] = dataSourceFlattenFirewallDosPolicy6AnomalyQuarantineLog(i["quarantine-log"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold"
		if _, ok := i["threshold"]; ok {
			tmp["threshold"] = dataSourceFlattenFirewallDosPolicy6AnomalyThreshold(i["threshold"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "thresholddefault"
		if _, ok := i["threshold(default)"]; ok {
			tmp["thresholddefault"] = dataSourceFlattenFirewallDosPolicy6AnomalyThresholdDefault(i["threshold(default)"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallDosPolicy6AnomalyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallDosPolicy6AnomalyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallDosPolicy6AnomalyLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallDosPolicy6AnomalyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallDosPolicy6AnomalyQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallDosPolicy6AnomalyQuarantineExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallDosPolicy6AnomalyQuarantineLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallDosPolicy6AnomalyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallDosPolicy6AnomalyThresholdDefault(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectFirewallDosPolicy6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("policyid", dataSourceFlattenFirewallDosPolicy6Policyid(o["policyid"], d, "policyid")); err != nil {
		if !fortiAPIPatch(o["policyid"]) {
			return fmt.Errorf("error reading policyid: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenFirewallDosPolicy6Status(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("error reading status: %v", err)
		}
	}

	if err = d.Set("name", dataSourceFlattenFirewallDosPolicy6Name(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("comments", dataSourceFlattenFirewallDosPolicy6Comments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("error reading comments: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenFirewallDosPolicy6Interface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("error reading interface: %v", err)
		}
	}

	if err = d.Set("srcaddr", dataSourceFlattenFirewallDosPolicy6Srcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if !fortiAPIPatch(o["srcaddr"]) {
			return fmt.Errorf("error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("dstaddr", dataSourceFlattenFirewallDosPolicy6Dstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if !fortiAPIPatch(o["dstaddr"]) {
			return fmt.Errorf("error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("service", dataSourceFlattenFirewallDosPolicy6Service(o["service"], d, "service")); err != nil {
		if !fortiAPIPatch(o["service"]) {
			return fmt.Errorf("error reading service: %v", err)
		}
	}

	if err = d.Set("anomaly", dataSourceFlattenFirewallDosPolicy6Anomaly(o["anomaly"], d, "anomaly")); err != nil {
		if !fortiAPIPatch(o["anomaly"]) {
			return fmt.Errorf("error reading anomaly: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenFirewallDosPolicy6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
