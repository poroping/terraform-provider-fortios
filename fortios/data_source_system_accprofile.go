// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure access profiles for system administrators.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemAccprofile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemAccprofileRead,
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
			"scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secfabgrp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ftviewgrp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"authgrp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sysgrp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"netgrp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"loggrp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fwgrp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vpngrp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"utmgrp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"wanoptgrp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"wifi": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"netgrp_permission": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cfg": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"packet_capture": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_cfg": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"sysgrp_permission": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"upd": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cfg": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mnt": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"fwgrp_permission": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"service": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"schedule": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"loggrp_permission": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"data_access": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"report_access": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"threat_weight": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"utmgrp_permission": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"antivirus": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ips": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"webfilter": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"emailfilter": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"spamfilter": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"data_loss_prevention": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"file_filter": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"application_control": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"icap": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"voip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"waf": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dnsfilter": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"endpoint_control": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"admintimeout_override": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"admintimeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"system_diagnostics": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemAccprofileRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("error describing SystemAccprofile: type error")
	}

	o, err := c.ReadSystemAccprofile(mkey, vdomparam, make(map[string][]string), 0)
	if err != nil {
		return fmt.Errorf("error describing SystemAccprofile: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemAccprofile(d, o)
	if err != nil {
		return fmt.Errorf("error describing SystemAccprofile from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemAccprofileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileScope(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileSecfabgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileFtviewgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileAuthgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileSysgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileNetgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileLoggrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileFwgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileVpngrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileWanoptgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileWifi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileNetgrpPermission(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cfg"
	if _, ok := i["cfg"]; ok {
		result["cfg"] = dataSourceFlattenSystemAccprofileNetgrpPermissionCfg(i["cfg"], d, pre_append)
	}

	pre_append = pre + ".0." + "packet_capture"
	if _, ok := i["packet-capture"]; ok {
		result["packet_capture"] = dataSourceFlattenSystemAccprofileNetgrpPermissionPacketCapture(i["packet-capture"], d, pre_append)
	}

	pre_append = pre + ".0." + "route_cfg"
	if _, ok := i["route-cfg"]; ok {
		result["route_cfg"] = dataSourceFlattenSystemAccprofileNetgrpPermissionRouteCfg(i["route-cfg"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenSystemAccprofileNetgrpPermissionCfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileNetgrpPermissionPacketCapture(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileNetgrpPermissionRouteCfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileSysgrpPermission(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "admin"
	if _, ok := i["admin"]; ok {
		result["admin"] = dataSourceFlattenSystemAccprofileSysgrpPermissionAdmin(i["admin"], d, pre_append)
	}

	pre_append = pre + ".0." + "upd"
	if _, ok := i["upd"]; ok {
		result["upd"] = dataSourceFlattenSystemAccprofileSysgrpPermissionUpd(i["upd"], d, pre_append)
	}

	pre_append = pre + ".0." + "cfg"
	if _, ok := i["cfg"]; ok {
		result["cfg"] = dataSourceFlattenSystemAccprofileSysgrpPermissionCfg(i["cfg"], d, pre_append)
	}

	pre_append = pre + ".0." + "mnt"
	if _, ok := i["mnt"]; ok {
		result["mnt"] = dataSourceFlattenSystemAccprofileSysgrpPermissionMnt(i["mnt"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenSystemAccprofileSysgrpPermissionAdmin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileSysgrpPermissionUpd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileSysgrpPermissionCfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileSysgrpPermissionMnt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileFwgrpPermission(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "policy"
	if _, ok := i["policy"]; ok {
		result["policy"] = dataSourceFlattenSystemAccprofileFwgrpPermissionPolicy(i["policy"], d, pre_append)
	}

	pre_append = pre + ".0." + "address"
	if _, ok := i["address"]; ok {
		result["address"] = dataSourceFlattenSystemAccprofileFwgrpPermissionAddress(i["address"], d, pre_append)
	}

	pre_append = pre + ".0." + "service"
	if _, ok := i["service"]; ok {
		result["service"] = dataSourceFlattenSystemAccprofileFwgrpPermissionService(i["service"], d, pre_append)
	}

	pre_append = pre + ".0." + "schedule"
	if _, ok := i["schedule"]; ok {
		result["schedule"] = dataSourceFlattenSystemAccprofileFwgrpPermissionSchedule(i["schedule"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenSystemAccprofileFwgrpPermissionPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileFwgrpPermissionAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileFwgrpPermissionService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileFwgrpPermissionSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileLoggrpPermission(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "config"
	if _, ok := i["config"]; ok {
		result["config"] = dataSourceFlattenSystemAccprofileLoggrpPermissionConfig(i["config"], d, pre_append)
	}

	pre_append = pre + ".0." + "data_access"
	if _, ok := i["data-access"]; ok {
		result["data_access"] = dataSourceFlattenSystemAccprofileLoggrpPermissionDataAccess(i["data-access"], d, pre_append)
	}

	pre_append = pre + ".0." + "report_access"
	if _, ok := i["report-access"]; ok {
		result["report_access"] = dataSourceFlattenSystemAccprofileLoggrpPermissionReportAccess(i["report-access"], d, pre_append)
	}

	pre_append = pre + ".0." + "threat_weight"
	if _, ok := i["threat-weight"]; ok {
		result["threat_weight"] = dataSourceFlattenSystemAccprofileLoggrpPermissionThreatWeight(i["threat-weight"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenSystemAccprofileLoggrpPermissionConfig(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileLoggrpPermissionDataAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileLoggrpPermissionReportAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileLoggrpPermissionThreatWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermission(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "antivirus"
	if _, ok := i["antivirus"]; ok {
		result["antivirus"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionAntivirus(i["antivirus"], d, pre_append)
	}

	pre_append = pre + ".0." + "ips"
	if _, ok := i["ips"]; ok {
		result["ips"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionIps(i["ips"], d, pre_append)
	}

	pre_append = pre + ".0." + "webfilter"
	if _, ok := i["webfilter"]; ok {
		result["webfilter"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionWebfilter(i["webfilter"], d, pre_append)
	}

	pre_append = pre + ".0." + "emailfilter"
	if _, ok := i["emailfilter"]; ok {
		result["emailfilter"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionEmailfilter(i["emailfilter"], d, pre_append)
	}

	pre_append = pre + ".0." + "spamfilter"
	if _, ok := i["spamfilter"]; ok {
		result["spamfilter"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionSpamfilter(i["spamfilter"], d, pre_append)
	}

	pre_append = pre + ".0." + "data_loss_prevention"
	if _, ok := i["data-loss-prevention"]; ok {
		result["data_loss_prevention"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionDataLossPrevention(i["data-loss-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "file_filter"
	if _, ok := i["file-filter"]; ok {
		result["file_filter"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionFileFilter(i["file-filter"], d, pre_append)
	}

	pre_append = pre + ".0." + "application_control"
	if _, ok := i["application-control"]; ok {
		result["application_control"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionApplicationControl(i["application-control"], d, pre_append)
	}

	pre_append = pre + ".0." + "icap"
	if _, ok := i["icap"]; ok {
		result["icap"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionIcap(i["icap"], d, pre_append)
	}

	pre_append = pre + ".0." + "voip"
	if _, ok := i["voip"]; ok {
		result["voip"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionVoip(i["voip"], d, pre_append)
	}

	pre_append = pre + ".0." + "waf"
	if _, ok := i["waf"]; ok {
		result["waf"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionWaf(i["waf"], d, pre_append)
	}

	pre_append = pre + ".0." + "dnsfilter"
	if _, ok := i["dnsfilter"]; ok {
		result["dnsfilter"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionDnsfilter(i["dnsfilter"], d, pre_append)
	}

	pre_append = pre + ".0." + "endpoint_control"
	if _, ok := i["endpoint-control"]; ok {
		result["endpoint_control"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionEndpointControl(i["endpoint-control"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionAntivirus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionIps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionWebfilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionEmailfilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionSpamfilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionDataLossPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionFileFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionApplicationControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionIcap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionVoip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionWaf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionDnsfilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionEndpointControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileAdmintimeoutOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileAdmintimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileSystemDiagnostics(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemAccprofile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenSystemAccprofileName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("scope", dataSourceFlattenSystemAccprofileScope(o["scope"], d, "scope")); err != nil {
		if !fortiAPIPatch(o["scope"]) {
			return fmt.Errorf("error reading scope: %v", err)
		}
	}

	if err = d.Set("comments", dataSourceFlattenSystemAccprofileComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("error reading comments: %v", err)
		}
	}

	if err = d.Set("secfabgrp", dataSourceFlattenSystemAccprofileSecfabgrp(o["secfabgrp"], d, "secfabgrp")); err != nil {
		if !fortiAPIPatch(o["secfabgrp"]) {
			return fmt.Errorf("error reading secfabgrp: %v", err)
		}
	}

	if err = d.Set("ftviewgrp", dataSourceFlattenSystemAccprofileFtviewgrp(o["ftviewgrp"], d, "ftviewgrp")); err != nil {
		if !fortiAPIPatch(o["ftviewgrp"]) {
			return fmt.Errorf("error reading ftviewgrp: %v", err)
		}
	}

	if err = d.Set("authgrp", dataSourceFlattenSystemAccprofileAuthgrp(o["authgrp"], d, "authgrp")); err != nil {
		if !fortiAPIPatch(o["authgrp"]) {
			return fmt.Errorf("error reading authgrp: %v", err)
		}
	}

	if err = d.Set("sysgrp", dataSourceFlattenSystemAccprofileSysgrp(o["sysgrp"], d, "sysgrp")); err != nil {
		if !fortiAPIPatch(o["sysgrp"]) {
			return fmt.Errorf("error reading sysgrp: %v", err)
		}
	}

	if err = d.Set("netgrp", dataSourceFlattenSystemAccprofileNetgrp(o["netgrp"], d, "netgrp")); err != nil {
		if !fortiAPIPatch(o["netgrp"]) {
			return fmt.Errorf("error reading netgrp: %v", err)
		}
	}

	if err = d.Set("loggrp", dataSourceFlattenSystemAccprofileLoggrp(o["loggrp"], d, "loggrp")); err != nil {
		if !fortiAPIPatch(o["loggrp"]) {
			return fmt.Errorf("error reading loggrp: %v", err)
		}
	}

	if err = d.Set("fwgrp", dataSourceFlattenSystemAccprofileFwgrp(o["fwgrp"], d, "fwgrp")); err != nil {
		if !fortiAPIPatch(o["fwgrp"]) {
			return fmt.Errorf("error reading fwgrp: %v", err)
		}
	}

	if err = d.Set("vpngrp", dataSourceFlattenSystemAccprofileVpngrp(o["vpngrp"], d, "vpngrp")); err != nil {
		if !fortiAPIPatch(o["vpngrp"]) {
			return fmt.Errorf("error reading vpngrp: %v", err)
		}
	}

	if err = d.Set("utmgrp", dataSourceFlattenSystemAccprofileUtmgrp(o["utmgrp"], d, "utmgrp")); err != nil {
		if !fortiAPIPatch(o["utmgrp"]) {
			return fmt.Errorf("error reading utmgrp: %v", err)
		}
	}

	if err = d.Set("wanoptgrp", dataSourceFlattenSystemAccprofileWanoptgrp(o["wanoptgrp"], d, "wanoptgrp")); err != nil {
		if !fortiAPIPatch(o["wanoptgrp"]) {
			return fmt.Errorf("error reading wanoptgrp: %v", err)
		}
	}

	if err = d.Set("wifi", dataSourceFlattenSystemAccprofileWifi(o["wifi"], d, "wifi")); err != nil {
		if !fortiAPIPatch(o["wifi"]) {
			return fmt.Errorf("error reading wifi: %v", err)
		}
	}

	if err = d.Set("netgrp_permission", dataSourceFlattenSystemAccprofileNetgrpPermission(o["netgrp-permission"], d, "netgrp_permission")); err != nil {
		if !fortiAPIPatch(o["netgrp-permission"]) {
			return fmt.Errorf("error reading netgrp_permission: %v", err)
		}
	}

	if err = d.Set("sysgrp_permission", dataSourceFlattenSystemAccprofileSysgrpPermission(o["sysgrp-permission"], d, "sysgrp_permission")); err != nil {
		if !fortiAPIPatch(o["sysgrp-permission"]) {
			return fmt.Errorf("error reading sysgrp_permission: %v", err)
		}
	}

	if err = d.Set("fwgrp_permission", dataSourceFlattenSystemAccprofileFwgrpPermission(o["fwgrp-permission"], d, "fwgrp_permission")); err != nil {
		if !fortiAPIPatch(o["fwgrp-permission"]) {
			return fmt.Errorf("error reading fwgrp_permission: %v", err)
		}
	}

	if err = d.Set("loggrp_permission", dataSourceFlattenSystemAccprofileLoggrpPermission(o["loggrp-permission"], d, "loggrp_permission")); err != nil {
		if !fortiAPIPatch(o["loggrp-permission"]) {
			return fmt.Errorf("error reading loggrp_permission: %v", err)
		}
	}

	if err = d.Set("utmgrp_permission", dataSourceFlattenSystemAccprofileUtmgrpPermission(o["utmgrp-permission"], d, "utmgrp_permission")); err != nil {
		if !fortiAPIPatch(o["utmgrp-permission"]) {
			return fmt.Errorf("error reading utmgrp_permission: %v", err)
		}
	}

	if err = d.Set("admintimeout_override", dataSourceFlattenSystemAccprofileAdmintimeoutOverride(o["admintimeout-override"], d, "admintimeout_override")); err != nil {
		if !fortiAPIPatch(o["admintimeout-override"]) {
			return fmt.Errorf("error reading admintimeout_override: %v", err)
		}
	}

	if err = d.Set("admintimeout", dataSourceFlattenSystemAccprofileAdmintimeout(o["admintimeout"], d, "admintimeout")); err != nil {
		if !fortiAPIPatch(o["admintimeout"]) {
			return fmt.Errorf("error reading admintimeout: %v", err)
		}
	}

	if err = d.Set("system_diagnostics", dataSourceFlattenSystemAccprofileSystemDiagnostics(o["system-diagnostics"], d, "system_diagnostics")); err != nil {
		if !fortiAPIPatch(o["system-diagnostics"]) {
			return fmt.Errorf("error reading system_diagnostics: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemAccprofileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
