// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: Configure RIP.

package fortios

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceRouterRip() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterRipRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"default_information_originate": {
				Type:        schema.TypeString,
				Description: "Enable/disable generation of default route.",
				Computed:    true,
			},
			"default_metric": {
				Type:        schema.TypeInt,
				Description: "Default metric.",
				Computed:    true,
			},
			"distance": {
				Type:        schema.TypeList,
				Description: "distance",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_list": {
							Type:        schema.TypeString,
							Description: "Access list for route destination.",
							Computed:    true,
						},
						"distance": {
							Type:        schema.TypeInt,
							Description: "Distance (1 - 255).",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Distance ID.",
							Computed:    true,
						},
						"prefix": {
							Type:        schema.TypeString,
							Description: "Distance prefix.",
							Computed:    true,
						},
					},
				},
			},
			"distribute_list": {
				Type:        schema.TypeList,
				Description: "Distribute list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"direction": {
							Type:        schema.TypeString,
							Description: "Distribute list direction.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Distribute list ID.",
							Computed:    true,
						},
						"interface": {
							Type:        schema.TypeString,
							Description: "Distribute list interface name.",
							Computed:    true,
						},
						"listname": {
							Type:        schema.TypeString,
							Description: "Distribute access/prefix list name.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "status",
							Computed:    true,
						},
					},
				},
			},
			"garbage_timer": {
				Type:        schema.TypeInt,
				Description: "Garbage timer in seconds.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeList,
				Description: "RIP interface configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_keychain": {
							Type:        schema.TypeString,
							Description: "Authentication key-chain name.",
							Computed:    true,
						},
						"auth_mode": {
							Type:        schema.TypeString,
							Description: "Authentication mode.",
							Computed:    true,
						},
						"auth_string": {
							Type:        schema.TypeString,
							Description: "Authentication string/password.",
							Computed:    true,
							Sensitive:   true,
						},
						"flags": {
							Type:        schema.TypeInt,
							Description: "flags",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
						"receive_version": {
							Type:        schema.TypeString,
							Description: "Receive version.",
							Computed:    true,
						},
						"send_version": {
							Type:        schema.TypeString,
							Description: "Send version.",
							Computed:    true,
						},
						"send_version2_broadcast": {
							Type:        schema.TypeString,
							Description: "Enable/disable broadcast version 1 compatible packets.",
							Computed:    true,
						},
						"split_horizon": {
							Type:        schema.TypeString,
							Description: "Enable/disable split horizon.",
							Computed:    true,
						},
						"split_horizon_status": {
							Type:        schema.TypeString,
							Description: "Enable/disable split horizon.",
							Computed:    true,
						},
					},
				},
			},
			"max_out_metric": {
				Type:        schema.TypeInt,
				Description: "Maximum metric allowed to output(0 means 'not set').",
				Computed:    true,
			},
			"neighbor": {
				Type:        schema.TypeList,
				Description: "neighbor",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Neighbor entry ID.",
							Computed:    true,
						},
						"ip": {
							Type:        schema.TypeString,
							Description: "IP address.",
							Computed:    true,
						},
					},
				},
			},
			"network": {
				Type:        schema.TypeList,
				Description: "network",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Network entry ID.",
							Computed:    true,
						},
						"prefix": {
							Type:        schema.TypeString,
							Description: "Network prefix.",
							Computed:    true,
						},
					},
				},
			},
			"offset_list": {
				Type:        schema.TypeList,
				Description: "Offset list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_list": {
							Type:        schema.TypeString,
							Description: "Access list name.",
							Computed:    true,
						},
						"direction": {
							Type:        schema.TypeString,
							Description: "Offset list direction.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Offset-list ID.",
							Computed:    true,
						},
						"interface": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
						"offset": {
							Type:        schema.TypeInt,
							Description: "offset",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "status",
							Computed:    true,
						},
					},
				},
			},
			"passive_interface": {
				Type:        schema.TypeList,
				Description: "Passive interface configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Passive interface name.",
							Computed:    true,
						},
					},
				},
			},
			"recv_buffer_size": {
				Type:        schema.TypeInt,
				Description: "Receiving buffer size.",
				Computed:    true,
			},
			"redistribute": {
				Type:        schema.TypeList,
				Description: "Redistribute configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metric": {
							Type:        schema.TypeInt,
							Description: "Redistribute metric setting.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Redistribute name.",
							Computed:    true,
						},
						"routemap": {
							Type:        schema.TypeString,
							Description: "Route map name.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "status",
							Computed:    true,
						},
					},
				},
			},
			"timeout_timer": {
				Type:        schema.TypeInt,
				Description: "Timeout timer in seconds.",
				Computed:    true,
			},
			"update_timer": {
				Type:        schema.TypeInt,
				Description: "Update timer in seconds.",
				Computed:    true,
			},
			"version": {
				Type:        schema.TypeString,
				Description: "RIP version.",
				Computed:    true,
			},
		},
	}
}

func dataSourceRouterRipRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""
	mkey = "RouterRip"

	o, err := c.ReadRouterRip(mkey, vdomparam, urlparams, 0)
	if err != nil {
		return fmt.Errorf("error describing RouterRip: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterRip(d, o)
	if err != nil {
		return fmt.Errorf("error describing RouterRip from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterRipDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDefaultMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistance(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list"
		if _, ok := i["access-list"]; ok {

			tmp["access_list"] = dataSourceFlattenRouterRipDistanceAccessList(i["access-list"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := i["distance"]; ok {

			tmp["distance"] = dataSourceFlattenRouterRipDistanceDistance(i["distance"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = dataSourceFlattenRouterRipDistanceId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {

			tmp["prefix"] = dataSourceFlattenRouterRipDistancePrefix(i["prefix"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipDistanceAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistanceDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistanceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistancePrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}
	return v
}

func dataSourceFlattenRouterRipDistributeList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {

			tmp["direction"] = dataSourceFlattenRouterRipDistributeListDirection(i["direction"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = dataSourceFlattenRouterRipDistributeListId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {

			tmp["interface"] = dataSourceFlattenRouterRipDistributeListInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listname"
		if _, ok := i["listname"]; ok {

			tmp["listname"] = dataSourceFlattenRouterRipDistributeListListname(i["listname"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = dataSourceFlattenRouterRipDistributeListStatus(i["status"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipDistributeListDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistributeListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistributeListInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistributeListListname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistributeListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipGarbageTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_keychain"
		if _, ok := i["auth-keychain"]; ok {

			tmp["auth_keychain"] = dataSourceFlattenRouterRipInterfaceAuthKeychain(i["auth-keychain"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode"
		if _, ok := i["auth-mode"]; ok {

			tmp["auth_mode"] = dataSourceFlattenRouterRipInterfaceAuthMode(i["auth-mode"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := i["flags"]; ok {

			tmp["flags"] = dataSourceFlattenRouterRipInterfaceFlags(i["flags"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = dataSourceFlattenRouterRipInterfaceName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "receive_version"
		if _, ok := i["receive-version"]; ok {

			tmp["receive_version"] = dataSourceFlattenRouterRipInterfaceReceiveVersion(i["receive-version"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_version"
		if _, ok := i["send-version"]; ok {

			tmp["send_version"] = dataSourceFlattenRouterRipInterfaceSendVersion(i["send-version"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_version2_broadcast"
		if _, ok := i["send-version2-broadcast"]; ok {

			tmp["send_version2_broadcast"] = dataSourceFlattenRouterRipInterfaceSendVersion2Broadcast(i["send-version2-broadcast"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon"
		if _, ok := i["split-horizon"]; ok {

			tmp["split_horizon"] = dataSourceFlattenRouterRipInterfaceSplitHorizon(i["split-horizon"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon_status"
		if _, ok := i["split-horizon-status"]; ok {

			tmp["split_horizon_status"] = dataSourceFlattenRouterRipInterfaceSplitHorizonStatus(i["split-horizon-status"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipInterfaceAuthKeychain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceAuthMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceReceiveVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceSendVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceSendVersion2Broadcast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceSplitHorizon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceSplitHorizonStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipMaxOutMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipNeighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["id"] = dataSourceFlattenRouterRipNeighborId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = dataSourceFlattenRouterRipNeighborIp(i["ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipNeighborId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipNeighborIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipNetwork(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["id"] = dataSourceFlattenRouterRipNetworkId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {

			tmp["prefix"] = dataSourceFlattenRouterRipNetworkPrefix(i["prefix"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipNetworkId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipNetworkPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}
	return v
}

func dataSourceFlattenRouterRipOffsetList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list"
		if _, ok := i["access-list"]; ok {

			tmp["access_list"] = dataSourceFlattenRouterRipOffsetListAccessList(i["access-list"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {

			tmp["direction"] = dataSourceFlattenRouterRipOffsetListDirection(i["direction"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = dataSourceFlattenRouterRipOffsetListId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {

			tmp["interface"] = dataSourceFlattenRouterRipOffsetListInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "offset"
		if _, ok := i["offset"]; ok {

			tmp["offset"] = dataSourceFlattenRouterRipOffsetListOffset(i["offset"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = dataSourceFlattenRouterRipOffsetListStatus(i["status"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipOffsetListAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipOffsetListDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipOffsetListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipOffsetListInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipOffsetListOffset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipOffsetListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipPassiveInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenRouterRipPassiveInterfaceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipPassiveInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipRecvBufferSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipRedistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := i["metric"]; ok {

			tmp["metric"] = dataSourceFlattenRouterRipRedistributeMetric(i["metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = dataSourceFlattenRouterRipRedistributeName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := i["routemap"]; ok {

			tmp["routemap"] = dataSourceFlattenRouterRipRedistributeRoutemap(i["routemap"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = dataSourceFlattenRouterRipRedistributeStatus(i["status"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipRedistributeMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipRedistributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipRedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipRedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipTimeoutTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipUpdateTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterRip(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("default_information_originate", dataSourceFlattenRouterRipDefaultInformationOriginate(o["default-information-originate"], d, "default_information_originate")); err != nil {
		if !fortiAPIPatch(o["default-information-originate"]) {
			return fmt.Errorf("error reading default_information_originate: %v", err)
		}
	}

	if err = d.Set("default_metric", dataSourceFlattenRouterRipDefaultMetric(o["default-metric"], d, "default_metric")); err != nil {
		if !fortiAPIPatch(o["default-metric"]) {
			return fmt.Errorf("error reading default_metric: %v", err)
		}
	}

	if err = d.Set("distance", dataSourceFlattenRouterRipDistance(o["distance"], d, "distance")); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("error reading distance: %v", err)
		}
	}

	if err = d.Set("distribute_list", dataSourceFlattenRouterRipDistributeList(o["distribute-list"], d, "distribute_list")); err != nil {
		if !fortiAPIPatch(o["distribute-list"]) {
			return fmt.Errorf("error reading distribute_list: %v", err)
		}
	}

	if err = d.Set("garbage_timer", dataSourceFlattenRouterRipGarbageTimer(o["garbage-timer"], d, "garbage_timer")); err != nil {
		if !fortiAPIPatch(o["garbage-timer"]) {
			return fmt.Errorf("error reading garbage_timer: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenRouterRipInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("error reading interface: %v", err)
		}
	}

	if err = d.Set("max_out_metric", dataSourceFlattenRouterRipMaxOutMetric(o["max-out-metric"], d, "max_out_metric")); err != nil {
		if !fortiAPIPatch(o["max-out-metric"]) {
			return fmt.Errorf("error reading max_out_metric: %v", err)
		}
	}

	if err = d.Set("neighbor", dataSourceFlattenRouterRipNeighbor(o["neighbor"], d, "neighbor")); err != nil {
		if !fortiAPIPatch(o["neighbor"]) {
			return fmt.Errorf("error reading neighbor: %v", err)
		}
	}

	if err = d.Set("network", dataSourceFlattenRouterRipNetwork(o["network"], d, "network")); err != nil {
		if !fortiAPIPatch(o["network"]) {
			return fmt.Errorf("error reading network: %v", err)
		}
	}

	if err = d.Set("offset_list", dataSourceFlattenRouterRipOffsetList(o["offset-list"], d, "offset_list")); err != nil {
		if !fortiAPIPatch(o["offset-list"]) {
			return fmt.Errorf("error reading offset_list: %v", err)
		}
	}

	if err = d.Set("passive_interface", dataSourceFlattenRouterRipPassiveInterface(o["passive-interface"], d, "passive_interface")); err != nil {
		if !fortiAPIPatch(o["passive-interface"]) {
			return fmt.Errorf("error reading passive_interface: %v", err)
		}
	}

	if err = d.Set("recv_buffer_size", dataSourceFlattenRouterRipRecvBufferSize(o["recv-buffer-size"], d, "recv_buffer_size")); err != nil {
		if !fortiAPIPatch(o["recv-buffer-size"]) {
			return fmt.Errorf("error reading recv_buffer_size: %v", err)
		}
	}

	if err = d.Set("redistribute", dataSourceFlattenRouterRipRedistribute(o["redistribute"], d, "redistribute")); err != nil {
		if !fortiAPIPatch(o["redistribute"]) {
			return fmt.Errorf("error reading redistribute: %v", err)
		}
	}

	if err = d.Set("timeout_timer", dataSourceFlattenRouterRipTimeoutTimer(o["timeout-timer"], d, "timeout_timer")); err != nil {
		if !fortiAPIPatch(o["timeout-timer"]) {
			return fmt.Errorf("error reading timeout_timer: %v", err)
		}
	}

	if err = d.Set("update_timer", dataSourceFlattenRouterRipUpdateTimer(o["update-timer"], d, "update_timer")); err != nil {
		if !fortiAPIPatch(o["update-timer"]) {
			return fmt.Errorf("error reading update_timer: %v", err)
		}
	}

	if err = d.Set("version", dataSourceFlattenRouterRipVersion(o["version"], d, "version")); err != nil {
		if !fortiAPIPatch(o["version"]) {
			return fmt.Errorf("error reading version: %v", err)
		}
	}

	return nil
}
