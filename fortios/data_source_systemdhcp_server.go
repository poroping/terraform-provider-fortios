// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure DHCP servers.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemDhcpServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemDhcpServerRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"fosid": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lease_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"mac_acl_default_action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"forticlient_on_net_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_service": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_server1": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_server2": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_server3": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_server4": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"wifi_ac_service": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"wifi_ac1": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"wifi_ac2": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"wifi_ac3": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ntp_service": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ntp_server1": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ntp_server2": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ntp_server3": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"wins_server1": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"wins_server2": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_gateway": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"next_server": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_range": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"start_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"end_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"timezone_option": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"timezone": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tftp_server": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tftp_server": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"filename": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"options": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"code": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"server_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"conflicted_ip_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ipsec_lease_hold": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"auto_configuration": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_settings_from_fortiipam": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_managed_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ddns_update": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ddns_update_override": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ddns_server_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ddns_zone": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ddns_auth": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ddns_keyname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ddns_key": {
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"ddns_ttl": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vci_match": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vci_string": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vci_string": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"exclude_range": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"start_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"end_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"reserved_address": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mac": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"action": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"circuit_id_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"circuit_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"remote_id_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"remote_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemDhcpServerRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""

	t := d.Get("fosid")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("error describing SystemDhcpServer: type error")
	}

	o, err := c.ReadSystemDhcpServer(mkey, vdomparam, make(map[string][]string), 0)
	if err != nil {
		return fmt.Errorf("error describing SystemDhcpServer: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemDhcpServer(d, o)
	if err != nil {
		return fmt.Errorf("error describing SystemDhcpServer from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemDhcpServerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerLeaseTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerMacAclDefaultAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerForticlientOnNetStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerDnsService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerDnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerDnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerDnsServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerDnsServer4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerWifiAcService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerWifiAc1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerWifiAc2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerWifiAc3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerNtpService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerNtpServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerNtpServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerNtpServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerWinsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerWinsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerDefaultGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerNextServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerNetmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerIpRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemDhcpServerIpRangeId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			tmp["start_ip"] = dataSourceFlattenSystemDhcpServerIpRangeStartIp(i["start-ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			tmp["end_ip"] = dataSourceFlattenSystemDhcpServerIpRangeEndIp(i["end-ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemDhcpServerIpRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerIpRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerIpRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerTimezoneOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerTimezone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerTftpServer(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tftp_server"
		if _, ok := i["tftp-server"]; ok {
			tmp["tftp_server"] = dataSourceFlattenSystemDhcpServerTftpServerTftpServer(i["tftp-server"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemDhcpServerTftpServerTftpServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerFilename(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerOptions(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemDhcpServerOptionsId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "code"
		if _, ok := i["code"]; ok {
			tmp["code"] = dataSourceFlattenSystemDhcpServerOptionsCode(i["code"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = dataSourceFlattenSystemDhcpServerOptionsType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			tmp["value"] = dataSourceFlattenSystemDhcpServerOptionsValue(i["value"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = dataSourceFlattenSystemDhcpServerOptionsIp(i["ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemDhcpServerOptionsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerOptionsCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerOptionsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerOptionsValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerOptionsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerIpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerConflictedIpTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerIpsecLeaseHold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerAutoConfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerDhcpSettingsFromFortiipam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerAutoManagedStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerDdnsUpdate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerDdnsUpdateOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerDdnsServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerDdnsZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerDdnsAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerDdnsKeyname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerDdnsKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerDdnsTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerVciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerVciString(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			tmp["vci_string"] = dataSourceFlattenSystemDhcpServerVciStringVciString(i["vci-string"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemDhcpServerVciStringVciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerExcludeRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemDhcpServerExcludeRangeId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			tmp["start_ip"] = dataSourceFlattenSystemDhcpServerExcludeRangeStartIp(i["start-ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			tmp["end_ip"] = dataSourceFlattenSystemDhcpServerExcludeRangeEndIp(i["end-ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemDhcpServerExcludeRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerExcludeRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerExcludeRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerReservedAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemDhcpServerReservedAddressId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = dataSourceFlattenSystemDhcpServerReservedAddressType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = dataSourceFlattenSystemDhcpServerReservedAddressIp(i["ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			tmp["mac"] = dataSourceFlattenSystemDhcpServerReservedAddressMac(i["mac"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = dataSourceFlattenSystemDhcpServerReservedAddressAction(i["action"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id_type"
		if _, ok := i["circuit-id-type"]; ok {
			tmp["circuit_id_type"] = dataSourceFlattenSystemDhcpServerReservedAddressCircuitIdType(i["circuit-id-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := i["circuit-id"]; ok {
			tmp["circuit_id"] = dataSourceFlattenSystemDhcpServerReservedAddressCircuitId(i["circuit-id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id_type"
		if _, ok := i["remote-id-type"]; ok {
			tmp["remote_id_type"] = dataSourceFlattenSystemDhcpServerReservedAddressRemoteIdType(i["remote-id-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := i["remote-id"]; ok {
			tmp["remote_id"] = dataSourceFlattenSystemDhcpServerReservedAddressRemoteId(i["remote-id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			tmp["description"] = dataSourceFlattenSystemDhcpServerReservedAddressDescription(i["description"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemDhcpServerReservedAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerReservedAddressType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerReservedAddressIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerReservedAddressMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerReservedAddressAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerReservedAddressCircuitIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerReservedAddressCircuitId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerReservedAddressRemoteIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerReservedAddressRemoteId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDhcpServerReservedAddressDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemDhcpServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", dataSourceFlattenSystemDhcpServerId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("error reading fosid: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenSystemDhcpServerStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("error reading status: %v", err)
		}
	}

	if err = d.Set("lease_time", dataSourceFlattenSystemDhcpServerLeaseTime(o["lease-time"], d, "lease_time")); err != nil {
		if !fortiAPIPatch(o["lease-time"]) {
			return fmt.Errorf("error reading lease_time: %v", err)
		}
	}

	if err = d.Set("mac_acl_default_action", dataSourceFlattenSystemDhcpServerMacAclDefaultAction(o["mac-acl-default-action"], d, "mac_acl_default_action")); err != nil {
		if !fortiAPIPatch(o["mac-acl-default-action"]) {
			return fmt.Errorf("error reading mac_acl_default_action: %v", err)
		}
	}

	if err = d.Set("forticlient_on_net_status", dataSourceFlattenSystemDhcpServerForticlientOnNetStatus(o["forticlient-on-net-status"], d, "forticlient_on_net_status")); err != nil {
		if !fortiAPIPatch(o["forticlient-on-net-status"]) {
			return fmt.Errorf("error reading forticlient_on_net_status: %v", err)
		}
	}

	if err = d.Set("dns_service", dataSourceFlattenSystemDhcpServerDnsService(o["dns-service"], d, "dns_service")); err != nil {
		if !fortiAPIPatch(o["dns-service"]) {
			return fmt.Errorf("error reading dns_service: %v", err)
		}
	}

	if err = d.Set("dns_server1", dataSourceFlattenSystemDhcpServerDnsServer1(o["dns-server1"], d, "dns_server1")); err != nil {
		if !fortiAPIPatch(o["dns-server1"]) {
			return fmt.Errorf("error reading dns_server1: %v", err)
		}
	}

	if err = d.Set("dns_server2", dataSourceFlattenSystemDhcpServerDnsServer2(o["dns-server2"], d, "dns_server2")); err != nil {
		if !fortiAPIPatch(o["dns-server2"]) {
			return fmt.Errorf("error reading dns_server2: %v", err)
		}
	}

	if err = d.Set("dns_server3", dataSourceFlattenSystemDhcpServerDnsServer3(o["dns-server3"], d, "dns_server3")); err != nil {
		if !fortiAPIPatch(o["dns-server3"]) {
			return fmt.Errorf("error reading dns_server3: %v", err)
		}
	}

	if err = d.Set("dns_server4", dataSourceFlattenSystemDhcpServerDnsServer4(o["dns-server4"], d, "dns_server4")); err != nil {
		if !fortiAPIPatch(o["dns-server4"]) {
			return fmt.Errorf("error reading dns_server4: %v", err)
		}
	}

	if err = d.Set("wifi_ac_service", dataSourceFlattenSystemDhcpServerWifiAcService(o["wifi-ac-service"], d, "wifi_ac_service")); err != nil {
		if !fortiAPIPatch(o["wifi-ac-service"]) {
			return fmt.Errorf("error reading wifi_ac_service: %v", err)
		}
	}

	if err = d.Set("wifi_ac1", dataSourceFlattenSystemDhcpServerWifiAc1(o["wifi-ac1"], d, "wifi_ac1")); err != nil {
		if !fortiAPIPatch(o["wifi-ac1"]) {
			return fmt.Errorf("error reading wifi_ac1: %v", err)
		}
	}

	if err = d.Set("wifi_ac2", dataSourceFlattenSystemDhcpServerWifiAc2(o["wifi-ac2"], d, "wifi_ac2")); err != nil {
		if !fortiAPIPatch(o["wifi-ac2"]) {
			return fmt.Errorf("error reading wifi_ac2: %v", err)
		}
	}

	if err = d.Set("wifi_ac3", dataSourceFlattenSystemDhcpServerWifiAc3(o["wifi-ac3"], d, "wifi_ac3")); err != nil {
		if !fortiAPIPatch(o["wifi-ac3"]) {
			return fmt.Errorf("error reading wifi_ac3: %v", err)
		}
	}

	if err = d.Set("ntp_service", dataSourceFlattenSystemDhcpServerNtpService(o["ntp-service"], d, "ntp_service")); err != nil {
		if !fortiAPIPatch(o["ntp-service"]) {
			return fmt.Errorf("error reading ntp_service: %v", err)
		}
	}

	if err = d.Set("ntp_server1", dataSourceFlattenSystemDhcpServerNtpServer1(o["ntp-server1"], d, "ntp_server1")); err != nil {
		if !fortiAPIPatch(o["ntp-server1"]) {
			return fmt.Errorf("error reading ntp_server1: %v", err)
		}
	}

	if err = d.Set("ntp_server2", dataSourceFlattenSystemDhcpServerNtpServer2(o["ntp-server2"], d, "ntp_server2")); err != nil {
		if !fortiAPIPatch(o["ntp-server2"]) {
			return fmt.Errorf("error reading ntp_server2: %v", err)
		}
	}

	if err = d.Set("ntp_server3", dataSourceFlattenSystemDhcpServerNtpServer3(o["ntp-server3"], d, "ntp_server3")); err != nil {
		if !fortiAPIPatch(o["ntp-server3"]) {
			return fmt.Errorf("error reading ntp_server3: %v", err)
		}
	}

	if err = d.Set("domain", dataSourceFlattenSystemDhcpServerDomain(o["domain"], d, "domain")); err != nil {
		if !fortiAPIPatch(o["domain"]) {
			return fmt.Errorf("error reading domain: %v", err)
		}
	}

	if err = d.Set("wins_server1", dataSourceFlattenSystemDhcpServerWinsServer1(o["wins-server1"], d, "wins_server1")); err != nil {
		if !fortiAPIPatch(o["wins-server1"]) {
			return fmt.Errorf("error reading wins_server1: %v", err)
		}
	}

	if err = d.Set("wins_server2", dataSourceFlattenSystemDhcpServerWinsServer2(o["wins-server2"], d, "wins_server2")); err != nil {
		if !fortiAPIPatch(o["wins-server2"]) {
			return fmt.Errorf("error reading wins_server2: %v", err)
		}
	}

	if err = d.Set("default_gateway", dataSourceFlattenSystemDhcpServerDefaultGateway(o["default-gateway"], d, "default_gateway")); err != nil {
		if !fortiAPIPatch(o["default-gateway"]) {
			return fmt.Errorf("error reading default_gateway: %v", err)
		}
	}

	if err = d.Set("next_server", dataSourceFlattenSystemDhcpServerNextServer(o["next-server"], d, "next_server")); err != nil {
		if !fortiAPIPatch(o["next-server"]) {
			return fmt.Errorf("error reading next_server: %v", err)
		}
	}

	if err = d.Set("netmask", dataSourceFlattenSystemDhcpServerNetmask(o["netmask"], d, "netmask")); err != nil {
		if !fortiAPIPatch(o["netmask"]) {
			return fmt.Errorf("error reading netmask: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenSystemDhcpServerInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("error reading interface: %v", err)
		}
	}

	if err = d.Set("ip_range", dataSourceFlattenSystemDhcpServerIpRange(o["ip-range"], d, "ip_range")); err != nil {
		if !fortiAPIPatch(o["ip-range"]) {
			return fmt.Errorf("error reading ip_range: %v", err)
		}
	}

	if err = d.Set("timezone_option", dataSourceFlattenSystemDhcpServerTimezoneOption(o["timezone-option"], d, "timezone_option")); err != nil {
		if !fortiAPIPatch(o["timezone-option"]) {
			return fmt.Errorf("error reading timezone_option: %v", err)
		}
	}

	if err = d.Set("timezone", dataSourceFlattenSystemDhcpServerTimezone(o["timezone"], d, "timezone")); err != nil {
		if !fortiAPIPatch(o["timezone"]) {
			return fmt.Errorf("error reading timezone: %v", err)
		}
	}

	if err = d.Set("tftp_server", dataSourceFlattenSystemDhcpServerTftpServer(o["tftp-server"], d, "tftp_server")); err != nil {
		if !fortiAPIPatch(o["tftp-server"]) {
			return fmt.Errorf("error reading tftp_server: %v", err)
		}
	}

	if err = d.Set("filename", dataSourceFlattenSystemDhcpServerFilename(o["filename"], d, "filename")); err != nil {
		if !fortiAPIPatch(o["filename"]) {
			return fmt.Errorf("error reading filename: %v", err)
		}
	}

	if err = d.Set("options", dataSourceFlattenSystemDhcpServerOptions(o["options"], d, "options")); err != nil {
		if !fortiAPIPatch(o["options"]) {
			return fmt.Errorf("error reading options: %v", err)
		}
	}

	if err = d.Set("server_type", dataSourceFlattenSystemDhcpServerServerType(o["server-type"], d, "server_type")); err != nil {
		if !fortiAPIPatch(o["server-type"]) {
			return fmt.Errorf("error reading server_type: %v", err)
		}
	}

	if err = d.Set("ip_mode", dataSourceFlattenSystemDhcpServerIpMode(o["ip-mode"], d, "ip_mode")); err != nil {
		if !fortiAPIPatch(o["ip-mode"]) {
			return fmt.Errorf("error reading ip_mode: %v", err)
		}
	}

	if err = d.Set("conflicted_ip_timeout", dataSourceFlattenSystemDhcpServerConflictedIpTimeout(o["conflicted-ip-timeout"], d, "conflicted_ip_timeout")); err != nil {
		if !fortiAPIPatch(o["conflicted-ip-timeout"]) {
			return fmt.Errorf("error reading conflicted_ip_timeout: %v", err)
		}
	}

	if err = d.Set("ipsec_lease_hold", dataSourceFlattenSystemDhcpServerIpsecLeaseHold(o["ipsec-lease-hold"], d, "ipsec_lease_hold")); err != nil {
		if !fortiAPIPatch(o["ipsec-lease-hold"]) {
			return fmt.Errorf("error reading ipsec_lease_hold: %v", err)
		}
	}

	if err = d.Set("auto_configuration", dataSourceFlattenSystemDhcpServerAutoConfiguration(o["auto-configuration"], d, "auto_configuration")); err != nil {
		if !fortiAPIPatch(o["auto-configuration"]) {
			return fmt.Errorf("error reading auto_configuration: %v", err)
		}
	}

	if err = d.Set("dhcp_settings_from_fortiipam", dataSourceFlattenSystemDhcpServerDhcpSettingsFromFortiipam(o["dhcp-settings-from-fortiipam"], d, "dhcp_settings_from_fortiipam")); err != nil {
		if !fortiAPIPatch(o["dhcp-settings-from-fortiipam"]) {
			return fmt.Errorf("error reading dhcp_settings_from_fortiipam: %v", err)
		}
	}

	if err = d.Set("auto_managed_status", dataSourceFlattenSystemDhcpServerAutoManagedStatus(o["auto-managed-status"], d, "auto_managed_status")); err != nil {
		if !fortiAPIPatch(o["auto-managed-status"]) {
			return fmt.Errorf("error reading auto_managed_status: %v", err)
		}
	}

	if err = d.Set("ddns_update", dataSourceFlattenSystemDhcpServerDdnsUpdate(o["ddns-update"], d, "ddns_update")); err != nil {
		if !fortiAPIPatch(o["ddns-update"]) {
			return fmt.Errorf("error reading ddns_update: %v", err)
		}
	}

	if err = d.Set("ddns_update_override", dataSourceFlattenSystemDhcpServerDdnsUpdateOverride(o["ddns-update-override"], d, "ddns_update_override")); err != nil {
		if !fortiAPIPatch(o["ddns-update-override"]) {
			return fmt.Errorf("error reading ddns_update_override: %v", err)
		}
	}

	if err = d.Set("ddns_server_ip", dataSourceFlattenSystemDhcpServerDdnsServerIp(o["ddns-server-ip"], d, "ddns_server_ip")); err != nil {
		if !fortiAPIPatch(o["ddns-server-ip"]) {
			return fmt.Errorf("error reading ddns_server_ip: %v", err)
		}
	}

	if err = d.Set("ddns_zone", dataSourceFlattenSystemDhcpServerDdnsZone(o["ddns-zone"], d, "ddns_zone")); err != nil {
		if !fortiAPIPatch(o["ddns-zone"]) {
			return fmt.Errorf("error reading ddns_zone: %v", err)
		}
	}

	if err = d.Set("ddns_auth", dataSourceFlattenSystemDhcpServerDdnsAuth(o["ddns-auth"], d, "ddns_auth")); err != nil {
		if !fortiAPIPatch(o["ddns-auth"]) {
			return fmt.Errorf("error reading ddns_auth: %v", err)
		}
	}

	if err = d.Set("ddns_keyname", dataSourceFlattenSystemDhcpServerDdnsKeyname(o["ddns-keyname"], d, "ddns_keyname")); err != nil {
		if !fortiAPIPatch(o["ddns-keyname"]) {
			return fmt.Errorf("error reading ddns_keyname: %v", err)
		}
	}

	if err = d.Set("ddns_ttl", dataSourceFlattenSystemDhcpServerDdnsTtl(o["ddns-ttl"], d, "ddns_ttl")); err != nil {
		if !fortiAPIPatch(o["ddns-ttl"]) {
			return fmt.Errorf("error reading ddns_ttl: %v", err)
		}
	}

	if err = d.Set("vci_match", dataSourceFlattenSystemDhcpServerVciMatch(o["vci-match"], d, "vci_match")); err != nil {
		if !fortiAPIPatch(o["vci-match"]) {
			return fmt.Errorf("error reading vci_match: %v", err)
		}
	}

	if err = d.Set("vci_string", dataSourceFlattenSystemDhcpServerVciString(o["vci-string"], d, "vci_string")); err != nil {
		if !fortiAPIPatch(o["vci-string"]) {
			return fmt.Errorf("error reading vci_string: %v", err)
		}
	}

	if err = d.Set("exclude_range", dataSourceFlattenSystemDhcpServerExcludeRange(o["exclude-range"], d, "exclude_range")); err != nil {
		if !fortiAPIPatch(o["exclude-range"]) {
			return fmt.Errorf("error reading exclude_range: %v", err)
		}
	}

	if err = d.Set("reserved_address", dataSourceFlattenSystemDhcpServerReservedAddress(o["reserved-address"], d, "reserved_address")); err != nil {
		if !fortiAPIPatch(o["reserved-address"]) {
			return fmt.Errorf("error reading reserved_address: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemDhcpServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
