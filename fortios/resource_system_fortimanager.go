// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiManager.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemFortimanager() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFortimanagerUpdate,
		Read:   resourceSystemFortimanagerRead,
		Update: resourceSystemFortimanagerUpdate,
		Delete: resourceSystemFortimanagerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vdom": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"ipsec": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"central_management": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"central_mgmt_auto_backup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"central_mgmt_schedule_config_restore": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"central_mgmt_schedule_script_restore": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"batchid": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func resourceSystemFortimanagerUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemFortimanager(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating SystemFortimanager resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemFortimanager(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating SystemFortimanager resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemFortimanager")
	}

	return resourceSystemFortimanagerRead(d, m)
}

func resourceSystemFortimanagerDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteSystemFortimanager(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting SystemFortimanager resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFortimanagerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemFortimanager(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading SystemFortimanager resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFortimanager(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading SystemFortimanager resource from API: %v", err)
	}
	return nil
}

func flattenSystemFortimanagerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortimanagerVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortimanagerIpsec(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortimanagerCentralManagement(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortimanagerCentralMgmtAutoBackup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortimanagerCentralMgmtScheduleConfigRestore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortimanagerCentralMgmtScheduleScriptRestore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemFortimanager(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ip", flattenSystemFortimanagerIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("error reading ip: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemFortimanagerVdom(o["vdom"], d, "vdom", sv)); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("error reading vdom: %v", err)
		}
	}

	if err = d.Set("ipsec", flattenSystemFortimanagerIpsec(o["ipsec"], d, "ipsec", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec"]) {
			return fmt.Errorf("error reading ipsec: %v", err)
		}
	}

	if err = d.Set("central_management", flattenSystemFortimanagerCentralManagement(o["central-management"], d, "central_management", sv)); err != nil {
		if !fortiAPIPatch(o["central-management"]) {
			return fmt.Errorf("error reading central_management: %v", err)
		}
	}

	if err = d.Set("central_mgmt_auto_backup", flattenSystemFortimanagerCentralMgmtAutoBackup(o["central-mgmt-auto-backup"], d, "central_mgmt_auto_backup", sv)); err != nil {
		if !fortiAPIPatch(o["central-mgmt-auto-backup"]) {
			return fmt.Errorf("error reading central_mgmt_auto_backup: %v", err)
		}
	}

	if err = d.Set("central_mgmt_schedule_config_restore", flattenSystemFortimanagerCentralMgmtScheduleConfigRestore(o["central-mgmt-schedule-config-restore"], d, "central_mgmt_schedule_config_restore", sv)); err != nil {
		if !fortiAPIPatch(o["central-mgmt-schedule-config-restore"]) {
			return fmt.Errorf("error reading central_mgmt_schedule_config_restore: %v", err)
		}
	}

	if err = d.Set("central_mgmt_schedule_script_restore", flattenSystemFortimanagerCentralMgmtScheduleScriptRestore(o["central-mgmt-schedule-script-restore"], d, "central_mgmt_schedule_script_restore", sv)); err != nil {
		if !fortiAPIPatch(o["central-mgmt-schedule-script-restore"]) {
			return fmt.Errorf("error reading central_mgmt_schedule_script_restore: %v", err)
		}
	}

	return nil
}

func flattenSystemFortimanagerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemFortimanagerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortimanagerVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortimanagerIpsec(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortimanagerCentralManagement(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortimanagerCentralMgmtAutoBackup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortimanagerCentralMgmtScheduleConfigRestore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortimanagerCentralMgmtScheduleScriptRestore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFortimanager(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ip"); ok {

		t, err := expandSystemFortimanagerIp(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok {

		t, err := expandSystemFortimanagerVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("ipsec"); ok {

		t, err := expandSystemFortimanagerIpsec(d, v, "ipsec", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec"] = t
		}
	}

	if v, ok := d.GetOk("central_management"); ok {

		t, err := expandSystemFortimanagerCentralManagement(d, v, "central_management", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["central-management"] = t
		}
	}

	if v, ok := d.GetOk("central_mgmt_auto_backup"); ok {

		t, err := expandSystemFortimanagerCentralMgmtAutoBackup(d, v, "central_mgmt_auto_backup", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["central-mgmt-auto-backup"] = t
		}
	}

	if v, ok := d.GetOk("central_mgmt_schedule_config_restore"); ok {

		t, err := expandSystemFortimanagerCentralMgmtScheduleConfigRestore(d, v, "central_mgmt_schedule_config_restore", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["central-mgmt-schedule-config-restore"] = t
		}
	}

	if v, ok := d.GetOk("central_mgmt_schedule_script_restore"); ok {

		t, err := expandSystemFortimanagerCentralMgmtScheduleScriptRestore(d, v, "central_mgmt_schedule_script_restore", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["central-mgmt-schedule-script-restore"] = t
		}
	}

	return &obj, nil
}
