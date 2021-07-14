// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure global DPDK options.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceDpdkGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceDpdkGlobalUpdate,
		Read:   resourceDpdkGlobalRead,
		Update: resourceDpdkGlobalUpdate,
		Delete: resourceDpdkGlobalDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"multiqueue": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sleep_on_idle": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"elasticbuffer": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"per_session_accounting": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hugepage_percentage": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 50),
				Optional:     true,
				Computed:     true,
			},
			"mbufpool_percentage": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 45),
				Optional:     true,
				Computed:     true,
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"batchid": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func resourceDpdkGlobalUpdate(d *schema.ResourceData, m interface{}) error {
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

	urlparams := make(map[string][]string)

	obj, err := getObjectDpdkGlobal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating DpdkGlobal resource while getting object: %v", err)
	}

	o, err := c.UpdateDpdkGlobal(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating DpdkGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("DpdkGlobal")
	}

	return resourceDpdkGlobalRead(d, m)
}

func resourceDpdkGlobalDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteDpdkGlobal(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting DpdkGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDpdkGlobalRead(d *schema.ResourceData, m interface{}) error {
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

	urlparams := make(map[string][]string)

	o, err := c.ReadDpdkGlobal(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading DpdkGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDpdkGlobal(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading DpdkGlobal resource from API: %v", err)
	}
	return nil
}

func flattenDpdkGlobalStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDpdkGlobalInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if _, ok := i["interface-name"]; ok {

			tmp["interface_name"] = flattenDpdkGlobalInterfaceInterfaceName(i["interface-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "interface_name", d)
	return result
}

func flattenDpdkGlobalInterfaceInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDpdkGlobalMultiqueue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDpdkGlobalSleepOnIdle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDpdkGlobalElasticbuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDpdkGlobalPerSessionAccounting(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDpdkGlobalHugepagePercentage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDpdkGlobalMbufpoolPercentage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectDpdkGlobal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenDpdkGlobalStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("error reading status: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("interface", flattenDpdkGlobalInterface(o["interface"], d, "interface", sv)); err != nil {
			if !fortiAPIPatch(o["interface"]) {
				return fmt.Errorf("error reading interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("interface"); ok {
			if err = d.Set("interface", flattenDpdkGlobalInterface(o["interface"], d, "interface", sv)); err != nil {
				if !fortiAPIPatch(o["interface"]) {
					return fmt.Errorf("error reading interface: %v", err)
				}
			}
		}
	}

	if err = d.Set("multiqueue", flattenDpdkGlobalMultiqueue(o["multiqueue"], d, "multiqueue", sv)); err != nil {
		if !fortiAPIPatch(o["multiqueue"]) {
			return fmt.Errorf("error reading multiqueue: %v", err)
		}
	}

	if err = d.Set("sleep_on_idle", flattenDpdkGlobalSleepOnIdle(o["sleep-on-idle"], d, "sleep_on_idle", sv)); err != nil {
		if !fortiAPIPatch(o["sleep-on-idle"]) {
			return fmt.Errorf("error reading sleep_on_idle: %v", err)
		}
	}

	if err = d.Set("elasticbuffer", flattenDpdkGlobalElasticbuffer(o["elasticbuffer"], d, "elasticbuffer", sv)); err != nil {
		if !fortiAPIPatch(o["elasticbuffer"]) {
			return fmt.Errorf("error reading elasticbuffer: %v", err)
		}
	}

	if err = d.Set("per_session_accounting", flattenDpdkGlobalPerSessionAccounting(o["per-session-accounting"], d, "per_session_accounting", sv)); err != nil {
		if !fortiAPIPatch(o["per-session-accounting"]) {
			return fmt.Errorf("error reading per_session_accounting: %v", err)
		}
	}

	if err = d.Set("hugepage_percentage", flattenDpdkGlobalHugepagePercentage(o["hugepage-percentage"], d, "hugepage_percentage", sv)); err != nil {
		if !fortiAPIPatch(o["hugepage-percentage"]) {
			return fmt.Errorf("error reading hugepage_percentage: %v", err)
		}
	}

	if err = d.Set("mbufpool_percentage", flattenDpdkGlobalMbufpoolPercentage(o["mbufpool-percentage"], d, "mbufpool_percentage", sv)); err != nil {
		if !fortiAPIPatch(o["mbufpool-percentage"]) {
			return fmt.Errorf("error reading mbufpool_percentage: %v", err)
		}
	}

	return nil
}

func flattenDpdkGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandDpdkGlobalStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDpdkGlobalInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["interface-name"], _ = expandDpdkGlobalInterfaceInterfaceName(d, i["interface_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandDpdkGlobalInterfaceInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDpdkGlobalMultiqueue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDpdkGlobalSleepOnIdle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDpdkGlobalElasticbuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDpdkGlobalPerSessionAccounting(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDpdkGlobalHugepagePercentage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDpdkGlobalMbufpoolPercentage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectDpdkGlobal(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {

		t, err := expandDpdkGlobalStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {

		t, err := expandDpdkGlobalInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("multiqueue"); ok {

		t, err := expandDpdkGlobalMultiqueue(d, v, "multiqueue", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multiqueue"] = t
		}
	}

	if v, ok := d.GetOk("sleep_on_idle"); ok {

		t, err := expandDpdkGlobalSleepOnIdle(d, v, "sleep_on_idle", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sleep-on-idle"] = t
		}
	}

	if v, ok := d.GetOk("elasticbuffer"); ok {

		t, err := expandDpdkGlobalElasticbuffer(d, v, "elasticbuffer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["elasticbuffer"] = t
		}
	}

	if v, ok := d.GetOk("per_session_accounting"); ok {

		t, err := expandDpdkGlobalPerSessionAccounting(d, v, "per_session_accounting", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-session-accounting"] = t
		}
	}

	if v, ok := d.GetOk("hugepage_percentage"); ok {

		t, err := expandDpdkGlobalHugepagePercentage(d, v, "hugepage_percentage", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hugepage-percentage"] = t
		}
	}

	if v, ok := d.GetOk("mbufpool_percentage"); ok {

		t, err := expandDpdkGlobalMbufpoolPercentage(d, v, "mbufpool_percentage", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbufpool-percentage"] = t
		}
	}

	return &obj, nil
}
