// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: Configure RIPng.

package fortios

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceRouterRipng() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterRipngRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"aggregate_address": {
				Type:        schema.TypeList,
				Description: "Aggregate address.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Aggregate address entry ID.",
							Computed:    true,
						},
						"prefix6": {
							Type:        schema.TypeString,
							Description: "Aggregate address prefix.",
							Computed:    true,
						},
					},
				},
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
						"access_list6": {
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
						"prefix6": {
							Type:        schema.TypeString,
							Description: "Distance prefix6.",
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
				Description: "Garbage timer.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeList,
				Description: "RIPng interface configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"flags": {
							Type:        schema.TypeInt,
							Description: "Flags.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Interface name.",
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
						"interface": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
						"ip6": {
							Type:        schema.TypeString,
							Description: "IPv6 link-local address.",
							Computed:    true,
						},
					},
				},
			},
			"network": {
				Type:        schema.TypeList,
				Description: "Network.",
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
							Description: "Network IPv6 link-local prefix.",
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
						"access_list6": {
							Type:        schema.TypeString,
							Description: "IPv6 access list name.",
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
				Description: "Timeout timer.",
				Computed:    true,
			},
			"update_timer": {
				Type:        schema.TypeInt,
				Description: "Update timer.",
				Computed:    true,
			},
		},
	}
}

func dataSourceRouterRipngRead(d *schema.ResourceData, m interface{}) error {
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
	mkey = "RouterRipng"

	o, err := c.ReadRouterRipng(mkey, vdomparam, urlparams, 0)
	if err != nil {
		return fmt.Errorf("error describing RouterRipng: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterRipng(d, o)
	if err != nil {
		return fmt.Errorf("error describing RouterRipng from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterRipngAggregateAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["id"] = dataSourceFlattenRouterRipngAggregateAddressId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {

			tmp["prefix6"] = dataSourceFlattenRouterRipngAggregateAddressPrefix6(i["prefix6"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipngAggregateAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngAggregateAddressPrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngDefaultMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngDistance(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list6"
		if _, ok := i["access-list6"]; ok {

			tmp["access_list6"] = dataSourceFlattenRouterRipngDistanceAccessList6(i["access-list6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := i["distance"]; ok {

			tmp["distance"] = dataSourceFlattenRouterRipngDistanceDistance(i["distance"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = dataSourceFlattenRouterRipngDistanceId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {

			tmp["prefix6"] = dataSourceFlattenRouterRipngDistancePrefix6(i["prefix6"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipngDistanceAccessList6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngDistanceDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngDistanceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngDistancePrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngDistributeList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["direction"] = dataSourceFlattenRouterRipngDistributeListDirection(i["direction"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = dataSourceFlattenRouterRipngDistributeListId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {

			tmp["interface"] = dataSourceFlattenRouterRipngDistributeListInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listname"
		if _, ok := i["listname"]; ok {

			tmp["listname"] = dataSourceFlattenRouterRipngDistributeListListname(i["listname"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = dataSourceFlattenRouterRipngDistributeListStatus(i["status"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipngDistributeListDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngDistributeListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngDistributeListInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngDistributeListListname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngDistributeListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngGarbageTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := i["flags"]; ok {

			tmp["flags"] = dataSourceFlattenRouterRipngInterfaceFlags(i["flags"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = dataSourceFlattenRouterRipngInterfaceName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon"
		if _, ok := i["split-horizon"]; ok {

			tmp["split_horizon"] = dataSourceFlattenRouterRipngInterfaceSplitHorizon(i["split-horizon"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon_status"
		if _, ok := i["split-horizon-status"]; ok {

			tmp["split_horizon_status"] = dataSourceFlattenRouterRipngInterfaceSplitHorizonStatus(i["split-horizon-status"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipngInterfaceFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngInterfaceSplitHorizon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngInterfaceSplitHorizonStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngMaxOutMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngNeighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["id"] = dataSourceFlattenRouterRipngNeighborId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {

			tmp["interface"] = dataSourceFlattenRouterRipngNeighborInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {

			tmp["ip6"] = dataSourceFlattenRouterRipngNeighborIp6(i["ip6"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipngNeighborId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngNeighborInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngNeighborIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngNetwork(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["id"] = dataSourceFlattenRouterRipngNetworkId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {

			tmp["prefix"] = dataSourceFlattenRouterRipngNetworkPrefix(i["prefix"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipngNetworkId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngNetworkPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngOffsetList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list6"
		if _, ok := i["access-list6"]; ok {

			tmp["access_list6"] = dataSourceFlattenRouterRipngOffsetListAccessList6(i["access-list6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {

			tmp["direction"] = dataSourceFlattenRouterRipngOffsetListDirection(i["direction"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = dataSourceFlattenRouterRipngOffsetListId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {

			tmp["interface"] = dataSourceFlattenRouterRipngOffsetListInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "offset"
		if _, ok := i["offset"]; ok {

			tmp["offset"] = dataSourceFlattenRouterRipngOffsetListOffset(i["offset"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = dataSourceFlattenRouterRipngOffsetListStatus(i["status"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipngOffsetListAccessList6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngOffsetListDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngOffsetListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngOffsetListInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngOffsetListOffset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngOffsetListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngPassiveInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["name"] = dataSourceFlattenRouterRipngPassiveInterfaceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipngPassiveInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngRedistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

			tmp["metric"] = dataSourceFlattenRouterRipngRedistributeMetric(i["metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = dataSourceFlattenRouterRipngRedistributeName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := i["routemap"]; ok {

			tmp["routemap"] = dataSourceFlattenRouterRipngRedistributeRoutemap(i["routemap"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = dataSourceFlattenRouterRipngRedistributeStatus(i["status"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipngRedistributeMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngRedistributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngRedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngRedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngTimeoutTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipngUpdateTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterRipng(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("aggregate_address", dataSourceFlattenRouterRipngAggregateAddress(o["aggregate-address"], d, "aggregate_address")); err != nil {
		if !fortiAPIPatch(o["aggregate-address"]) {
			return fmt.Errorf("error reading aggregate_address: %v", err)
		}
	}

	if err = d.Set("default_information_originate", dataSourceFlattenRouterRipngDefaultInformationOriginate(o["default-information-originate"], d, "default_information_originate")); err != nil {
		if !fortiAPIPatch(o["default-information-originate"]) {
			return fmt.Errorf("error reading default_information_originate: %v", err)
		}
	}

	if err = d.Set("default_metric", dataSourceFlattenRouterRipngDefaultMetric(o["default-metric"], d, "default_metric")); err != nil {
		if !fortiAPIPatch(o["default-metric"]) {
			return fmt.Errorf("error reading default_metric: %v", err)
		}
	}

	if err = d.Set("distance", dataSourceFlattenRouterRipngDistance(o["distance"], d, "distance")); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("error reading distance: %v", err)
		}
	}

	if err = d.Set("distribute_list", dataSourceFlattenRouterRipngDistributeList(o["distribute-list"], d, "distribute_list")); err != nil {
		if !fortiAPIPatch(o["distribute-list"]) {
			return fmt.Errorf("error reading distribute_list: %v", err)
		}
	}

	if err = d.Set("garbage_timer", dataSourceFlattenRouterRipngGarbageTimer(o["garbage-timer"], d, "garbage_timer")); err != nil {
		if !fortiAPIPatch(o["garbage-timer"]) {
			return fmt.Errorf("error reading garbage_timer: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenRouterRipngInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("error reading interface: %v", err)
		}
	}

	if err = d.Set("max_out_metric", dataSourceFlattenRouterRipngMaxOutMetric(o["max-out-metric"], d, "max_out_metric")); err != nil {
		if !fortiAPIPatch(o["max-out-metric"]) {
			return fmt.Errorf("error reading max_out_metric: %v", err)
		}
	}

	if err = d.Set("neighbor", dataSourceFlattenRouterRipngNeighbor(o["neighbor"], d, "neighbor")); err != nil {
		if !fortiAPIPatch(o["neighbor"]) {
			return fmt.Errorf("error reading neighbor: %v", err)
		}
	}

	if err = d.Set("network", dataSourceFlattenRouterRipngNetwork(o["network"], d, "network")); err != nil {
		if !fortiAPIPatch(o["network"]) {
			return fmt.Errorf("error reading network: %v", err)
		}
	}

	if err = d.Set("offset_list", dataSourceFlattenRouterRipngOffsetList(o["offset-list"], d, "offset_list")); err != nil {
		if !fortiAPIPatch(o["offset-list"]) {
			return fmt.Errorf("error reading offset_list: %v", err)
		}
	}

	if err = d.Set("passive_interface", dataSourceFlattenRouterRipngPassiveInterface(o["passive-interface"], d, "passive_interface")); err != nil {
		if !fortiAPIPatch(o["passive-interface"]) {
			return fmt.Errorf("error reading passive_interface: %v", err)
		}
	}

	if err = d.Set("redistribute", dataSourceFlattenRouterRipngRedistribute(o["redistribute"], d, "redistribute")); err != nil {
		if !fortiAPIPatch(o["redistribute"]) {
			return fmt.Errorf("error reading redistribute: %v", err)
		}
	}

	if err = d.Set("timeout_timer", dataSourceFlattenRouterRipngTimeoutTimer(o["timeout-timer"], d, "timeout_timer")); err != nil {
		if !fortiAPIPatch(o["timeout-timer"]) {
			return fmt.Errorf("error reading timeout_timer: %v", err)
		}
	}

	if err = d.Set("update_timer", dataSourceFlattenRouterRipngUpdateTimer(o["update-timer"], d, "update_timer")); err != nil {
		if !fortiAPIPatch(o["update-timer"]) {
			return fmt.Errorf("error reading update_timer: %v", err)
		}
	}

	return nil
}
