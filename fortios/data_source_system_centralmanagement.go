// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure central management.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemCentralManagement() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemCentralManagementRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"schedule_config_restore": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"schedule_script_restore": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_push_configuration": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_push_firmware": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_remote_firmware_upgrade": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_monitor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fmg": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fmg_source_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fmg_source_ip6": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"local_cert": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ca_cert": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdom": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"server_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"addr_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"server_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"server_address6": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fqdn": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"fmg_update_port": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"include_default_servers": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enc_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface_select_method": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemCentralManagementRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "SystemCentralManagement"

	o, err := c.ReadSystemCentralManagement(mkey, vdomparam, make(map[string][]string), 0)
	if err != nil {
		return fmt.Errorf("error describing SystemCentralManagement: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemCentralManagement(d, o)
	if err != nil {
		return fmt.Errorf("error describing SystemCentralManagement from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemCentralManagementMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementScheduleConfigRestore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementScheduleScriptRestore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementAllowPushConfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementAllowPushFirmware(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementAllowRemoteFirmwareUpgrade(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementAllowMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementSerialNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementFmg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementFmgSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementFmgSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementLocalCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementCaCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementServerList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenSystemCentralManagementServerListId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_type"
		if _, ok := i["server-type"]; ok {
			tmp["server_type"] = dataSourceFlattenSystemCentralManagementServerListServerType(i["server-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := i["addr-type"]; ok {
			tmp["addr_type"] = dataSourceFlattenSystemCentralManagementServerListAddrType(i["addr-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_address"
		if _, ok := i["server-address"]; ok {
			tmp["server_address"] = dataSourceFlattenSystemCentralManagementServerListServerAddress(i["server-address"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_address6"
		if _, ok := i["server-address6"]; ok {
			tmp["server_address6"] = dataSourceFlattenSystemCentralManagementServerListServerAddress6(i["server-address6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fqdn"
		if _, ok := i["fqdn"]; ok {
			tmp["fqdn"] = dataSourceFlattenSystemCentralManagementServerListFqdn(i["fqdn"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemCentralManagementServerListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementServerListServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementServerListAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementServerListServerAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementServerListServerAddress6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementServerListFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementFmgUpdatePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementIncludeDefaultServers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementEncAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemCentralManagement(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("mode", dataSourceFlattenSystemCentralManagementMode(o["mode"], d, "mode")); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("error reading mode: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenSystemCentralManagementType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("error reading type: %v", err)
		}
	}

	if err = d.Set("schedule_config_restore", dataSourceFlattenSystemCentralManagementScheduleConfigRestore(o["schedule-config-restore"], d, "schedule_config_restore")); err != nil {
		if !fortiAPIPatch(o["schedule-config-restore"]) {
			return fmt.Errorf("error reading schedule_config_restore: %v", err)
		}
	}

	if err = d.Set("schedule_script_restore", dataSourceFlattenSystemCentralManagementScheduleScriptRestore(o["schedule-script-restore"], d, "schedule_script_restore")); err != nil {
		if !fortiAPIPatch(o["schedule-script-restore"]) {
			return fmt.Errorf("error reading schedule_script_restore: %v", err)
		}
	}

	if err = d.Set("allow_push_configuration", dataSourceFlattenSystemCentralManagementAllowPushConfiguration(o["allow-push-configuration"], d, "allow_push_configuration")); err != nil {
		if !fortiAPIPatch(o["allow-push-configuration"]) {
			return fmt.Errorf("error reading allow_push_configuration: %v", err)
		}
	}

	if err = d.Set("allow_push_firmware", dataSourceFlattenSystemCentralManagementAllowPushFirmware(o["allow-push-firmware"], d, "allow_push_firmware")); err != nil {
		if !fortiAPIPatch(o["allow-push-firmware"]) {
			return fmt.Errorf("error reading allow_push_firmware: %v", err)
		}
	}

	if err = d.Set("allow_remote_firmware_upgrade", dataSourceFlattenSystemCentralManagementAllowRemoteFirmwareUpgrade(o["allow-remote-firmware-upgrade"], d, "allow_remote_firmware_upgrade")); err != nil {
		if !fortiAPIPatch(o["allow-remote-firmware-upgrade"]) {
			return fmt.Errorf("error reading allow_remote_firmware_upgrade: %v", err)
		}
	}

	if err = d.Set("allow_monitor", dataSourceFlattenSystemCentralManagementAllowMonitor(o["allow-monitor"], d, "allow_monitor")); err != nil {
		if !fortiAPIPatch(o["allow-monitor"]) {
			return fmt.Errorf("error reading allow_monitor: %v", err)
		}
	}

	if err = d.Set("serial_number", dataSourceFlattenSystemCentralManagementSerialNumber(o["serial-number"], d, "serial_number")); err != nil {
		if !fortiAPIPatch(o["serial-number"]) {
			return fmt.Errorf("error reading serial_number: %v", err)
		}
	}

	if err = d.Set("fmg", dataSourceFlattenSystemCentralManagementFmg(o["fmg"], d, "fmg")); err != nil {
		if !fortiAPIPatch(o["fmg"]) {
			return fmt.Errorf("error reading fmg: %v", err)
		}
	}

	if err = d.Set("fmg_source_ip", dataSourceFlattenSystemCentralManagementFmgSourceIp(o["fmg-source-ip"], d, "fmg_source_ip")); err != nil {
		if !fortiAPIPatch(o["fmg-source-ip"]) {
			return fmt.Errorf("error reading fmg_source_ip: %v", err)
		}
	}

	if err = d.Set("fmg_source_ip6", dataSourceFlattenSystemCentralManagementFmgSourceIp6(o["fmg-source-ip6"], d, "fmg_source_ip6")); err != nil {
		if !fortiAPIPatch(o["fmg-source-ip6"]) {
			return fmt.Errorf("error reading fmg_source_ip6: %v", err)
		}
	}

	if err = d.Set("local_cert", dataSourceFlattenSystemCentralManagementLocalCert(o["local-cert"], d, "local_cert")); err != nil {
		if !fortiAPIPatch(o["local-cert"]) {
			return fmt.Errorf("error reading local_cert: %v", err)
		}
	}

	if err = d.Set("ca_cert", dataSourceFlattenSystemCentralManagementCaCert(o["ca-cert"], d, "ca_cert")); err != nil {
		if !fortiAPIPatch(o["ca-cert"]) {
			return fmt.Errorf("error reading ca_cert: %v", err)
		}
	}

	if err = d.Set("vdom", dataSourceFlattenSystemCentralManagementVdom(o["vdom"], d, "vdom")); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("error reading vdom: %v", err)
		}
	}

	if err = d.Set("server_list", dataSourceFlattenSystemCentralManagementServerList(o["server-list"], d, "server_list")); err != nil {
		if !fortiAPIPatch(o["server-list"]) {
			return fmt.Errorf("error reading server_list: %v", err)
		}
	}

	if err = d.Set("fmg_update_port", dataSourceFlattenSystemCentralManagementFmgUpdatePort(o["fmg-update-port"], d, "fmg_update_port")); err != nil {
		if !fortiAPIPatch(o["fmg-update-port"]) {
			return fmt.Errorf("error reading fmg_update_port: %v", err)
		}
	}

	if err = d.Set("include_default_servers", dataSourceFlattenSystemCentralManagementIncludeDefaultServers(o["include-default-servers"], d, "include_default_servers")); err != nil {
		if !fortiAPIPatch(o["include-default-servers"]) {
			return fmt.Errorf("error reading include_default_servers: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", dataSourceFlattenSystemCentralManagementEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm")); err != nil {
		if !fortiAPIPatch(o["enc-algorithm"]) {
			return fmt.Errorf("error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("interface_select_method", dataSourceFlattenSystemCentralManagementInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenSystemCentralManagementInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("error reading interface: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemCentralManagementFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
