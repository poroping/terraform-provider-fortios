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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemFortimanager() *schema.Resource {
	return &schema.Resource{
		Read:               dataSourceSystemFortimanagerRead,
		DeprecationMessage: "This resource has been depecated. Please use the dedicated FortiManager provider. This resource will be removed in v2.0.0",
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdom": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipsec": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"central_management": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"central_mgmt_auto_backup": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"central_mgmt_schedule_config_restore": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"central_mgmt_schedule_script_restore": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemFortimanagerRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "SystemFortimanager"

	o, err := c.ReadSystemFortimanager(mkey, vdomparam, make(map[string][]string), 0)
	if err != nil {
		return fmt.Errorf("error describing SystemFortimanager: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemFortimanager(d, o)
	if err != nil {
		return fmt.Errorf("error describing SystemFortimanager from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemFortimanagerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortimanagerVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortimanagerIpsec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortimanagerCentralManagement(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortimanagerCentralMgmtAutoBackup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortimanagerCentralMgmtScheduleConfigRestore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortimanagerCentralMgmtScheduleScriptRestore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemFortimanager(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ip", dataSourceFlattenSystemFortimanagerIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("error reading ip: %v", err)
		}
	}

	if err = d.Set("vdom", dataSourceFlattenSystemFortimanagerVdom(o["vdom"], d, "vdom")); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("error reading vdom: %v", err)
		}
	}

	if err = d.Set("ipsec", dataSourceFlattenSystemFortimanagerIpsec(o["ipsec"], d, "ipsec")); err != nil {
		if !fortiAPIPatch(o["ipsec"]) {
			return fmt.Errorf("error reading ipsec: %v", err)
		}
	}

	if err = d.Set("central_management", dataSourceFlattenSystemFortimanagerCentralManagement(o["central-management"], d, "central_management")); err != nil {
		if !fortiAPIPatch(o["central-management"]) {
			return fmt.Errorf("error reading central_management: %v", err)
		}
	}

	if err = d.Set("central_mgmt_auto_backup", dataSourceFlattenSystemFortimanagerCentralMgmtAutoBackup(o["central-mgmt-auto-backup"], d, "central_mgmt_auto_backup")); err != nil {
		if !fortiAPIPatch(o["central-mgmt-auto-backup"]) {
			return fmt.Errorf("error reading central_mgmt_auto_backup: %v", err)
		}
	}

	if err = d.Set("central_mgmt_schedule_config_restore", dataSourceFlattenSystemFortimanagerCentralMgmtScheduleConfigRestore(o["central-mgmt-schedule-config-restore"], d, "central_mgmt_schedule_config_restore")); err != nil {
		if !fortiAPIPatch(o["central-mgmt-schedule-config-restore"]) {
			return fmt.Errorf("error reading central_mgmt_schedule_config_restore: %v", err)
		}
	}

	if err = d.Set("central_mgmt_schedule_script_restore", dataSourceFlattenSystemFortimanagerCentralMgmtScheduleScriptRestore(o["central-mgmt-schedule-script-restore"], d, "central_mgmt_schedule_script_restore")); err != nil {
		if !fortiAPIPatch(o["central-mgmt-schedule-script-restore"]) {
			return fmt.Errorf("error reading central_mgmt_schedule_script_restore: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemFortimanagerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
