// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: BGP neighbor table.

package fortios

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceRouterbgpNeighbor() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterbgpNeighborRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"activate": {
				Type:        schema.TypeString,
				Description: "Enable/disable address family IPv4 for this neighbor.",
				Computed:    true,
			},
			"activate6": {
				Type:        schema.TypeString,
				Description: "Enable/disable address family IPv6 for this neighbor.",
				Computed:    true,
			},
			"additional_path": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv4 additional-path capability.",
				Computed:    true,
			},
			"additional_path6": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv6 additional-path capability.",
				Computed:    true,
			},
			"adv_additional_path": {
				Type:        schema.TypeInt,
				Description: "Number of IPv4 additional paths that can be advertised to this neighbor.",
				Computed:    true,
			},
			"adv_additional_path6": {
				Type:        schema.TypeInt,
				Description: "Number of IPv6 additional paths that can be advertised to this neighbor.",
				Computed:    true,
			},
			"advertisement_interval": {
				Type:        schema.TypeInt,
				Description: "Minimum interval (sec) between sending updates.",
				Computed:    true,
			},
			"allowas_in": {
				Type:        schema.TypeInt,
				Description: "IPv4 The maximum number of occurrence of my AS number allowed.",
				Computed:    true,
			},
			"allowas_in_enable": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv4 Enable to allow my AS in AS path.",
				Computed:    true,
			},
			"allowas_in_enable6": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv6 Enable to allow my AS in AS path.",
				Computed:    true,
			},
			"allowas_in6": {
				Type:        schema.TypeInt,
				Description: "IPv6 The maximum number of occurrence of my AS number allowed.",
				Computed:    true,
			},
			"as_override": {
				Type:        schema.TypeString,
				Description: "Enable/disable replace peer AS with own AS for IPv4.",
				Computed:    true,
			},
			"as_override6": {
				Type:        schema.TypeString,
				Description: "Enable/disable replace peer AS with own AS for IPv6.",
				Computed:    true,
			},
			"attribute_unchanged": {
				Type:        schema.TypeString,
				Description: "IPv4 List of attributes that should be unchanged.",
				Computed:    true,
			},
			"attribute_unchanged6": {
				Type:        schema.TypeString,
				Description: "IPv6 List of attributes that should be unchanged.",
				Computed:    true,
			},
			"bfd": {
				Type:        schema.TypeString,
				Description: "Enable/disable BFD for this neighbor.",
				Computed:    true,
			},
			"capability_default_originate": {
				Type:        schema.TypeString,
				Description: "Enable/disable advertise default IPv4 route to this neighbor.",
				Computed:    true,
			},
			"capability_default_originate6": {
				Type:        schema.TypeString,
				Description: "Enable/disable advertise default IPv6 route to this neighbor.",
				Computed:    true,
			},
			"capability_dynamic": {
				Type:        schema.TypeString,
				Description: "Enable/disable advertise dynamic capability to this neighbor.",
				Computed:    true,
			},
			"capability_graceful_restart": {
				Type:        schema.TypeString,
				Description: "Enable/disable advertise IPv4 graceful restart capability to this neighbor.",
				Computed:    true,
			},
			"capability_graceful_restart6": {
				Type:        schema.TypeString,
				Description: "Enable/disable advertise IPv6 graceful restart capability to this neighbor.",
				Computed:    true,
			},
			"capability_orf": {
				Type:        schema.TypeString,
				Description: "Accept/Send IPv4 ORF lists to/from this neighbor.",
				Computed:    true,
			},
			"capability_orf6": {
				Type:        schema.TypeString,
				Description: "Accept/Send IPv6 ORF lists to/from this neighbor.",
				Computed:    true,
			},
			"capability_route_refresh": {
				Type:        schema.TypeString,
				Description: "Enable/disable advertise route refresh capability to this neighbor.",
				Computed:    true,
			},
			"conditional_advertise": {
				Type:        schema.TypeList,
				Description: "Conditional advertisement.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advertise_routemap": {
							Type:        schema.TypeString,
							Description: "Name of advertising route map.",
							Computed:    true,
						},
						"condition_routemap": {
							Type:        schema.TypeString,
							Description: "Name of condition route map.",
							Computed:    true,
						},
						"condition_type": {
							Type:        schema.TypeString,
							Description: "Type of condition.",
							Computed:    true,
						},
					},
				},
			},
			"conditional_advertise6": {
				Type:        schema.TypeList,
				Description: "IPv6 conditional advertisement.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advertise_routemap": {
							Type:        schema.TypeString,
							Description: "Name of advertising route map.",
							Computed:    true,
						},
						"condition_routemap": {
							Type:        schema.TypeString,
							Description: "Name of condition route map.",
							Computed:    true,
						},
						"condition_type": {
							Type:        schema.TypeString,
							Description: "Type of condition.",
							Computed:    true,
						},
					},
				},
			},
			"connect_timer": {
				Type:        schema.TypeInt,
				Description: "Interval (sec) for connect timer.",
				Computed:    true,
			},
			"default_originate_routemap": {
				Type:        schema.TypeString,
				Description: "Route map to specify criteria to originate IPv4 default.",
				Computed:    true,
			},
			"default_originate_routemap6": {
				Type:        schema.TypeString,
				Description: "Route map to specify criteria to originate IPv6 default.",
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description.",
				Computed:    true,
			},
			"distribute_list_in": {
				Type:        schema.TypeString,
				Description: "Filter for IPv4 updates from this neighbor.",
				Computed:    true,
			},
			"distribute_list_in6": {
				Type:        schema.TypeString,
				Description: "Filter for IPv6 updates from this neighbor.",
				Computed:    true,
			},
			"distribute_list_out": {
				Type:        schema.TypeString,
				Description: "Filter for IPv4 updates to this neighbor.",
				Computed:    true,
			},
			"distribute_list_out6": {
				Type:        schema.TypeString,
				Description: "Filter for IPv6 updates to this neighbor.",
				Computed:    true,
			},
			"dont_capability_negotiate": {
				Type:        schema.TypeString,
				Description: "Don't negotiate capabilities with this neighbor",
				Computed:    true,
			},
			"ebgp_enforce_multihop": {
				Type:        schema.TypeString,
				Description: "Enable/disable allow multi-hop EBGP neighbors.",
				Computed:    true,
			},
			"ebgp_multihop_ttl": {
				Type:        schema.TypeInt,
				Description: "EBGP multihop TTL for this peer.",
				Computed:    true,
			},
			"filter_list_in": {
				Type:        schema.TypeString,
				Description: "BGP filter for IPv4 inbound routes.",
				Computed:    true,
			},
			"filter_list_in6": {
				Type:        schema.TypeString,
				Description: "BGP filter for IPv6 inbound routes.",
				Computed:    true,
			},
			"filter_list_out": {
				Type:        schema.TypeString,
				Description: "BGP filter for IPv4 outbound routes.",
				Computed:    true,
			},
			"filter_list_out6": {
				Type:        schema.TypeString,
				Description: "BGP filter for IPv6 outbound routes.",
				Computed:    true,
			},
			"holdtime_timer": {
				Type:        schema.TypeInt,
				Description: "Interval (sec) before peer considered dead.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Specify outgoing interface for peer connection. For IPv6 peer, the interface should have link-local address.",
				Computed:    true,
			},
			"ip": {
				Type:        schema.TypeString,
				Description: "IP/IPv6 address of neighbor.",
				Required:    true,
			},
			"keep_alive_timer": {
				Type:        schema.TypeInt,
				Description: "Keep alive timer interval (sec).",
				Computed:    true,
			},
			"link_down_failover": {
				Type:        schema.TypeString,
				Description: "Enable/disable failover upon link down.",
				Computed:    true,
			},
			"local_as": {
				Type:        schema.TypeInt,
				Description: "Local AS number of neighbor.",
				Computed:    true,
			},
			"local_as_no_prepend": {
				Type:        schema.TypeString,
				Description: "Do not prepend local-as to incoming updates.",
				Computed:    true,
			},
			"local_as_replace_as": {
				Type:        schema.TypeString,
				Description: "Replace real AS with local-as in outgoing updates.",
				Computed:    true,
			},
			"maximum_prefix": {
				Type:        schema.TypeInt,
				Description: "Maximum number of IPv4 prefixes to accept from this peer.",
				Computed:    true,
			},
			"maximum_prefix_threshold": {
				Type:        schema.TypeInt,
				Description: "Maximum IPv4 prefix threshold value (1 - 100 percent).",
				Computed:    true,
			},
			"maximum_prefix_threshold6": {
				Type:        schema.TypeInt,
				Description: "Maximum IPv6 prefix threshold value (1 - 100 percent).",
				Computed:    true,
			},
			"maximum_prefix_warning_only": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv4 Only give warning message when limit is exceeded.",
				Computed:    true,
			},
			"maximum_prefix_warning_only6": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv6 Only give warning message when limit is exceeded.",
				Computed:    true,
			},
			"maximum_prefix6": {
				Type:        schema.TypeInt,
				Description: "Maximum number of IPv6 prefixes to accept from this peer.",
				Computed:    true,
			},
			"next_hop_self": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv4 next-hop calculation for this neighbor.",
				Computed:    true,
			},
			"next_hop_self_rr": {
				Type:        schema.TypeString,
				Description: "Enable/disable setting nexthop's address to interface's IPv4 address for route-reflector routes.",
				Computed:    true,
			},
			"next_hop_self_rr6": {
				Type:        schema.TypeString,
				Description: "Enable/disable setting nexthop's address to interface's IPv6 address for route-reflector routes.",
				Computed:    true,
			},
			"next_hop_self6": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv6 next-hop calculation for this neighbor.",
				Computed:    true,
			},
			"override_capability": {
				Type:        schema.TypeString,
				Description: "Enable/disable override result of capability negotiation.",
				Computed:    true,
			},
			"passive": {
				Type:        schema.TypeString,
				Description: "Enable/disable sending of open messages to this neighbor.",
				Computed:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "Password used in MD5 authentication.",
				Computed:    true,
				Sensitive:   true,
			},
			"prefix_list_in": {
				Type:        schema.TypeString,
				Description: "IPv4 Inbound filter for updates from this neighbor.",
				Computed:    true,
			},
			"prefix_list_in6": {
				Type:        schema.TypeString,
				Description: "IPv6 Inbound filter for updates from this neighbor.",
				Computed:    true,
			},
			"prefix_list_out": {
				Type:        schema.TypeString,
				Description: "IPv4 Outbound filter for updates to this neighbor.",
				Computed:    true,
			},
			"prefix_list_out6": {
				Type:        schema.TypeString,
				Description: "IPv6 Outbound filter for updates to this neighbor.",
				Computed:    true,
			},
			"remote_as": {
				Type:        schema.TypeInt,
				Description: "AS number of neighbor.",
				Computed:    true,
			},
			"remove_private_as": {
				Type:        schema.TypeString,
				Description: "Enable/disable remove private AS number from IPv4 outbound updates.",
				Computed:    true,
			},
			"remove_private_as6": {
				Type:        schema.TypeString,
				Description: "Enable/disable remove private AS number from IPv6 outbound updates.",
				Computed:    true,
			},
			"restart_time": {
				Type:        schema.TypeInt,
				Description: "Graceful restart delay time (sec, 0 = global default).",
				Computed:    true,
			},
			"retain_stale_time": {
				Type:        schema.TypeInt,
				Description: "Time to retain stale routes.",
				Computed:    true,
			},
			"route_map_in": {
				Type:        schema.TypeString,
				Description: "IPv4 Inbound route map filter.",
				Computed:    true,
			},
			"route_map_in6": {
				Type:        schema.TypeString,
				Description: "IPv6 Inbound route map filter.",
				Computed:    true,
			},
			"route_map_out": {
				Type:        schema.TypeString,
				Description: "IPv4 outbound route map filter.",
				Computed:    true,
			},
			"route_map_out_preferable": {
				Type:        schema.TypeString,
				Description: "IPv4 outbound route map filter if the peer is preferred.",
				Computed:    true,
			},
			"route_map_out6": {
				Type:        schema.TypeString,
				Description: "IPv6 Outbound route map filter.",
				Computed:    true,
			},
			"route_map_out6_preferable": {
				Type:        schema.TypeString,
				Description: "IPv6 outbound route map filter if the peer is preferred.",
				Computed:    true,
			},
			"route_reflector_client": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv4 AS route reflector client.",
				Computed:    true,
			},
			"route_reflector_client6": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv6 AS route reflector client.",
				Computed:    true,
			},
			"route_server_client": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv4 AS route server client.",
				Computed:    true,
			},
			"route_server_client6": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv6 AS route server client.",
				Computed:    true,
			},
			"send_community": {
				Type:        schema.TypeString,
				Description: "IPv4 Send community attribute to neighbor.",
				Computed:    true,
			},
			"send_community6": {
				Type:        schema.TypeString,
				Description: "IPv6 Send community attribute to neighbor.",
				Computed:    true,
			},
			"shutdown": {
				Type:        schema.TypeString,
				Description: "Enable/disable shutdown this neighbor.",
				Computed:    true,
			},
			"soft_reconfiguration": {
				Type:        schema.TypeString,
				Description: "Enable/disable allow IPv4 inbound soft reconfiguration.",
				Computed:    true,
			},
			"soft_reconfiguration6": {
				Type:        schema.TypeString,
				Description: "Enable/disable allow IPv6 inbound soft reconfiguration.",
				Computed:    true,
			},
			"stale_route": {
				Type:        schema.TypeString,
				Description: "Enable/disable stale route after neighbor down.",
				Computed:    true,
			},
			"strict_capability_match": {
				Type:        schema.TypeString,
				Description: "Enable/disable strict capability matching.",
				Computed:    true,
			},
			"unsuppress_map": {
				Type:        schema.TypeString,
				Description: "IPv4 Route map to selectively unsuppress suppressed routes.",
				Computed:    true,
			},
			"unsuppress_map6": {
				Type:        schema.TypeString,
				Description: "IPv6 Route map to selectively unsuppress suppressed routes.",
				Computed:    true,
			},
			"update_source": {
				Type:        schema.TypeString,
				Description: "Interface to use as source IP/IPv6 address of TCP connections.",
				Computed:    true,
			},
			"weight": {
				Type:        schema.TypeInt,
				Description: "Neighbor weight.",
				Computed:    true,
			},
		},
	}
}

func dataSourceRouterbgpNeighborRead(d *schema.ResourceData, m interface{}) error {
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
	key := "ip"
	t := d.Get(key)
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("error describing RouterbgpNeighbor: type error")
	}

	o, err := c.ReadRouterbgpNeighbor(mkey, vdomparam, urlparams, 0)
	if err != nil {
		return fmt.Errorf("error describing RouterbgpNeighbor: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterbgpNeighbor(d, o)
	if err != nil {
		return fmt.Errorf("error describing RouterbgpNeighbor from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterbgpNeighborActivate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborActivate6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAdditionalPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAdditionalPath6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAdvAdditionalPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAdvAdditionalPath6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAdvertisementInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAllowasIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAllowasInEnable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAllowasInEnable6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAllowasIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAsOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAsOverride6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAttributeUnchanged(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAttributeUnchanged6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborCapabilityDefaultOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborCapabilityDefaultOriginate6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborCapabilityDynamic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborCapabilityGracefulRestart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborCapabilityGracefulRestart6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborCapabilityOrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborCapabilityOrf6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborCapabilityRouteRefresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborConditionalAdvertise(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise_routemap"
		if _, ok := i["advertise-routemap"]; ok {

			tmp["advertise_routemap"] = dataSourceFlattenRouterbgpNeighborConditionalAdvertiseAdvertiseRoutemap(i["advertise-routemap"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if _, ok := i["condition-routemap"]; ok {

			tmp["condition_routemap"] = dataSourceFlattenRouterbgpNeighborConditionalAdvertiseConditionRoutemap(i["condition-routemap"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if _, ok := i["condition-type"]; ok {

			tmp["condition_type"] = dataSourceFlattenRouterbgpNeighborConditionalAdvertiseConditionType(i["condition-type"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterbgpNeighborConditionalAdvertiseAdvertiseRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborConditionalAdvertiseConditionRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborConditionalAdvertiseConditionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborConditionalAdvertise6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise_routemap"
		if _, ok := i["advertise-routemap"]; ok {

			tmp["advertise_routemap"] = dataSourceFlattenRouterbgpNeighborConditionalAdvertise6AdvertiseRoutemap(i["advertise-routemap"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if _, ok := i["condition-routemap"]; ok {

			tmp["condition_routemap"] = dataSourceFlattenRouterbgpNeighborConditionalAdvertise6ConditionRoutemap(i["condition-routemap"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if _, ok := i["condition-type"]; ok {

			tmp["condition_type"] = dataSourceFlattenRouterbgpNeighborConditionalAdvertise6ConditionType(i["condition-type"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterbgpNeighborConditionalAdvertise6AdvertiseRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborConditionalAdvertise6ConditionRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborConditionalAdvertise6ConditionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborConnectTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborDefaultOriginateRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborDefaultOriginateRoutemap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborDistributeListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborDistributeListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborDistributeListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborDistributeListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborDontCapabilityNegotiate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborEbgpEnforceMultihop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborEbgpMultihopTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborFilterListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborFilterListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborFilterListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborFilterListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborHoldtimeTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborKeepAliveTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborLinkDownFailover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborLocalAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborLocalAsNoPrepend(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborLocalAsReplaceAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborMaximumPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborMaximumPrefixThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborMaximumPrefixThreshold6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborMaximumPrefixWarningOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborMaximumPrefixWarningOnly6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborMaximumPrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborNextHopSelf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborNextHopSelfRr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborNextHopSelfRr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborNextHopSelf6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborOverrideCapability(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborPassive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborPrefixListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborPrefixListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborPrefixListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborPrefixListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRemoteAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRemovePrivateAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRemovePrivateAs6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRestartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRetainStaleTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRouteMapIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRouteMapIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRouteMapOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRouteMapOutPreferable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRouteMapOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRouteMapOut6Preferable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRouteReflectorClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRouteReflectorClient6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRouteServerClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRouteServerClient6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborSendCommunity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborSendCommunity6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborShutdown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborSoftReconfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborSoftReconfiguration6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborStaleRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborStrictCapabilityMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborUnsuppressMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborUnsuppressMap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborUpdateSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterbgpNeighbor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("activate", dataSourceFlattenRouterbgpNeighborActivate(o["activate"], d, "activate")); err != nil {
		if !fortiAPIPatch(o["activate"]) {
			return fmt.Errorf("error reading activate: %v", err)
		}
	}

	if err = d.Set("activate6", dataSourceFlattenRouterbgpNeighborActivate6(o["activate6"], d, "activate6")); err != nil {
		if !fortiAPIPatch(o["activate6"]) {
			return fmt.Errorf("error reading activate6: %v", err)
		}
	}

	if err = d.Set("additional_path", dataSourceFlattenRouterbgpNeighborAdditionalPath(o["additional-path"], d, "additional_path")); err != nil {
		if !fortiAPIPatch(o["additional-path"]) {
			return fmt.Errorf("error reading additional_path: %v", err)
		}
	}

	if err = d.Set("additional_path6", dataSourceFlattenRouterbgpNeighborAdditionalPath6(o["additional-path6"], d, "additional_path6")); err != nil {
		if !fortiAPIPatch(o["additional-path6"]) {
			return fmt.Errorf("error reading additional_path6: %v", err)
		}
	}

	if err = d.Set("adv_additional_path", dataSourceFlattenRouterbgpNeighborAdvAdditionalPath(o["adv-additional-path"], d, "adv_additional_path")); err != nil {
		if !fortiAPIPatch(o["adv-additional-path"]) {
			return fmt.Errorf("error reading adv_additional_path: %v", err)
		}
	}

	if err = d.Set("adv_additional_path6", dataSourceFlattenRouterbgpNeighborAdvAdditionalPath6(o["adv-additional-path6"], d, "adv_additional_path6")); err != nil {
		if !fortiAPIPatch(o["adv-additional-path6"]) {
			return fmt.Errorf("error reading adv_additional_path6: %v", err)
		}
	}

	if err = d.Set("advertisement_interval", dataSourceFlattenRouterbgpNeighborAdvertisementInterval(o["advertisement-interval"], d, "advertisement_interval")); err != nil {
		if !fortiAPIPatch(o["advertisement-interval"]) {
			return fmt.Errorf("error reading advertisement_interval: %v", err)
		}
	}

	if err = d.Set("allowas_in", dataSourceFlattenRouterbgpNeighborAllowasIn(o["allowas-in"], d, "allowas_in")); err != nil {
		if !fortiAPIPatch(o["allowas-in"]) {
			return fmt.Errorf("error reading allowas_in: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable", dataSourceFlattenRouterbgpNeighborAllowasInEnable(o["allowas-in-enable"], d, "allowas_in_enable")); err != nil {
		if !fortiAPIPatch(o["allowas-in-enable"]) {
			return fmt.Errorf("error reading allowas_in_enable: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable6", dataSourceFlattenRouterbgpNeighborAllowasInEnable6(o["allowas-in-enable6"], d, "allowas_in_enable6")); err != nil {
		if !fortiAPIPatch(o["allowas-in-enable6"]) {
			return fmt.Errorf("error reading allowas_in_enable6: %v", err)
		}
	}

	if err = d.Set("allowas_in6", dataSourceFlattenRouterbgpNeighborAllowasIn6(o["allowas-in6"], d, "allowas_in6")); err != nil {
		if !fortiAPIPatch(o["allowas-in6"]) {
			return fmt.Errorf("error reading allowas_in6: %v", err)
		}
	}

	if err = d.Set("as_override", dataSourceFlattenRouterbgpNeighborAsOverride(o["as-override"], d, "as_override")); err != nil {
		if !fortiAPIPatch(o["as-override"]) {
			return fmt.Errorf("error reading as_override: %v", err)
		}
	}

	if err = d.Set("as_override6", dataSourceFlattenRouterbgpNeighborAsOverride6(o["as-override6"], d, "as_override6")); err != nil {
		if !fortiAPIPatch(o["as-override6"]) {
			return fmt.Errorf("error reading as_override6: %v", err)
		}
	}

	if err = d.Set("attribute_unchanged", dataSourceFlattenRouterbgpNeighborAttributeUnchanged(o["attribute-unchanged"], d, "attribute_unchanged")); err != nil {
		if !fortiAPIPatch(o["attribute-unchanged"]) {
			return fmt.Errorf("error reading attribute_unchanged: %v", err)
		}
	}

	if err = d.Set("attribute_unchanged6", dataSourceFlattenRouterbgpNeighborAttributeUnchanged6(o["attribute-unchanged6"], d, "attribute_unchanged6")); err != nil {
		if !fortiAPIPatch(o["attribute-unchanged6"]) {
			return fmt.Errorf("error reading attribute_unchanged6: %v", err)
		}
	}

	if err = d.Set("bfd", dataSourceFlattenRouterbgpNeighborBfd(o["bfd"], d, "bfd")); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("error reading bfd: %v", err)
		}
	}

	if err = d.Set("capability_default_originate", dataSourceFlattenRouterbgpNeighborCapabilityDefaultOriginate(o["capability-default-originate"], d, "capability_default_originate")); err != nil {
		if !fortiAPIPatch(o["capability-default-originate"]) {
			return fmt.Errorf("error reading capability_default_originate: %v", err)
		}
	}

	if err = d.Set("capability_default_originate6", dataSourceFlattenRouterbgpNeighborCapabilityDefaultOriginate6(o["capability-default-originate6"], d, "capability_default_originate6")); err != nil {
		if !fortiAPIPatch(o["capability-default-originate6"]) {
			return fmt.Errorf("error reading capability_default_originate6: %v", err)
		}
	}

	if err = d.Set("capability_dynamic", dataSourceFlattenRouterbgpNeighborCapabilityDynamic(o["capability-dynamic"], d, "capability_dynamic")); err != nil {
		if !fortiAPIPatch(o["capability-dynamic"]) {
			return fmt.Errorf("error reading capability_dynamic: %v", err)
		}
	}

	if err = d.Set("capability_graceful_restart", dataSourceFlattenRouterbgpNeighborCapabilityGracefulRestart(o["capability-graceful-restart"], d, "capability_graceful_restart")); err != nil {
		if !fortiAPIPatch(o["capability-graceful-restart"]) {
			return fmt.Errorf("error reading capability_graceful_restart: %v", err)
		}
	}

	if err = d.Set("capability_graceful_restart6", dataSourceFlattenRouterbgpNeighborCapabilityGracefulRestart6(o["capability-graceful-restart6"], d, "capability_graceful_restart6")); err != nil {
		if !fortiAPIPatch(o["capability-graceful-restart6"]) {
			return fmt.Errorf("error reading capability_graceful_restart6: %v", err)
		}
	}

	if err = d.Set("capability_orf", dataSourceFlattenRouterbgpNeighborCapabilityOrf(o["capability-orf"], d, "capability_orf")); err != nil {
		if !fortiAPIPatch(o["capability-orf"]) {
			return fmt.Errorf("error reading capability_orf: %v", err)
		}
	}

	if err = d.Set("capability_orf6", dataSourceFlattenRouterbgpNeighborCapabilityOrf6(o["capability-orf6"], d, "capability_orf6")); err != nil {
		if !fortiAPIPatch(o["capability-orf6"]) {
			return fmt.Errorf("error reading capability_orf6: %v", err)
		}
	}

	if err = d.Set("capability_route_refresh", dataSourceFlattenRouterbgpNeighborCapabilityRouteRefresh(o["capability-route-refresh"], d, "capability_route_refresh")); err != nil {
		if !fortiAPIPatch(o["capability-route-refresh"]) {
			return fmt.Errorf("error reading capability_route_refresh: %v", err)
		}
	}

	if err = d.Set("conditional_advertise", dataSourceFlattenRouterbgpNeighborConditionalAdvertise(o["conditional-advertise"], d, "conditional_advertise")); err != nil {
		if !fortiAPIPatch(o["conditional-advertise"]) {
			return fmt.Errorf("error reading conditional_advertise: %v", err)
		}
	}

	if err = d.Set("conditional_advertise6", dataSourceFlattenRouterbgpNeighborConditionalAdvertise6(o["conditional-advertise6"], d, "conditional_advertise6")); err != nil {
		if !fortiAPIPatch(o["conditional-advertise6"]) {
			return fmt.Errorf("error reading conditional_advertise6: %v", err)
		}
	}

	if err = d.Set("connect_timer", dataSourceFlattenRouterbgpNeighborConnectTimer(o["connect-timer"], d, "connect_timer")); err != nil {
		if !fortiAPIPatch(o["connect-timer"]) {
			return fmt.Errorf("error reading connect_timer: %v", err)
		}
	}

	if err = d.Set("default_originate_routemap", dataSourceFlattenRouterbgpNeighborDefaultOriginateRoutemap(o["default-originate-routemap"], d, "default_originate_routemap")); err != nil {
		if !fortiAPIPatch(o["default-originate-routemap"]) {
			return fmt.Errorf("error reading default_originate_routemap: %v", err)
		}
	}

	if err = d.Set("default_originate_routemap6", dataSourceFlattenRouterbgpNeighborDefaultOriginateRoutemap6(o["default-originate-routemap6"], d, "default_originate_routemap6")); err != nil {
		if !fortiAPIPatch(o["default-originate-routemap6"]) {
			return fmt.Errorf("error reading default_originate_routemap6: %v", err)
		}
	}

	if err = d.Set("description", dataSourceFlattenRouterbgpNeighborDescription(o["description"], d, "description")); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("error reading description: %v", err)
		}
	}

	if err = d.Set("distribute_list_in", dataSourceFlattenRouterbgpNeighborDistributeListIn(o["distribute-list-in"], d, "distribute_list_in")); err != nil {
		if !fortiAPIPatch(o["distribute-list-in"]) {
			return fmt.Errorf("error reading distribute_list_in: %v", err)
		}
	}

	if err = d.Set("distribute_list_in6", dataSourceFlattenRouterbgpNeighborDistributeListIn6(o["distribute-list-in6"], d, "distribute_list_in6")); err != nil {
		if !fortiAPIPatch(o["distribute-list-in6"]) {
			return fmt.Errorf("error reading distribute_list_in6: %v", err)
		}
	}

	if err = d.Set("distribute_list_out", dataSourceFlattenRouterbgpNeighborDistributeListOut(o["distribute-list-out"], d, "distribute_list_out")); err != nil {
		if !fortiAPIPatch(o["distribute-list-out"]) {
			return fmt.Errorf("error reading distribute_list_out: %v", err)
		}
	}

	if err = d.Set("distribute_list_out6", dataSourceFlattenRouterbgpNeighborDistributeListOut6(o["distribute-list-out6"], d, "distribute_list_out6")); err != nil {
		if !fortiAPIPatch(o["distribute-list-out6"]) {
			return fmt.Errorf("error reading distribute_list_out6: %v", err)
		}
	}

	if err = d.Set("dont_capability_negotiate", dataSourceFlattenRouterbgpNeighborDontCapabilityNegotiate(o["dont-capability-negotiate"], d, "dont_capability_negotiate")); err != nil {
		if !fortiAPIPatch(o["dont-capability-negotiate"]) {
			return fmt.Errorf("error reading dont_capability_negotiate: %v", err)
		}
	}

	if err = d.Set("ebgp_enforce_multihop", dataSourceFlattenRouterbgpNeighborEbgpEnforceMultihop(o["ebgp-enforce-multihop"], d, "ebgp_enforce_multihop")); err != nil {
		if !fortiAPIPatch(o["ebgp-enforce-multihop"]) {
			return fmt.Errorf("error reading ebgp_enforce_multihop: %v", err)
		}
	}

	if err = d.Set("ebgp_multihop_ttl", dataSourceFlattenRouterbgpNeighborEbgpMultihopTtl(o["ebgp-multihop-ttl"], d, "ebgp_multihop_ttl")); err != nil {
		if !fortiAPIPatch(o["ebgp-multihop-ttl"]) {
			return fmt.Errorf("error reading ebgp_multihop_ttl: %v", err)
		}
	}

	if err = d.Set("filter_list_in", dataSourceFlattenRouterbgpNeighborFilterListIn(o["filter-list-in"], d, "filter_list_in")); err != nil {
		if !fortiAPIPatch(o["filter-list-in"]) {
			return fmt.Errorf("error reading filter_list_in: %v", err)
		}
	}

	if err = d.Set("filter_list_in6", dataSourceFlattenRouterbgpNeighborFilterListIn6(o["filter-list-in6"], d, "filter_list_in6")); err != nil {
		if !fortiAPIPatch(o["filter-list-in6"]) {
			return fmt.Errorf("error reading filter_list_in6: %v", err)
		}
	}

	if err = d.Set("filter_list_out", dataSourceFlattenRouterbgpNeighborFilterListOut(o["filter-list-out"], d, "filter_list_out")); err != nil {
		if !fortiAPIPatch(o["filter-list-out"]) {
			return fmt.Errorf("error reading filter_list_out: %v", err)
		}
	}

	if err = d.Set("filter_list_out6", dataSourceFlattenRouterbgpNeighborFilterListOut6(o["filter-list-out6"], d, "filter_list_out6")); err != nil {
		if !fortiAPIPatch(o["filter-list-out6"]) {
			return fmt.Errorf("error reading filter_list_out6: %v", err)
		}
	}

	if err = d.Set("holdtime_timer", dataSourceFlattenRouterbgpNeighborHoldtimeTimer(o["holdtime-timer"], d, "holdtime_timer")); err != nil {
		if !fortiAPIPatch(o["holdtime-timer"]) {
			return fmt.Errorf("error reading holdtime_timer: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenRouterbgpNeighborInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("error reading interface: %v", err)
		}
	}

	if err = d.Set("ip", dataSourceFlattenRouterbgpNeighborIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("error reading ip: %v", err)
		}
	}

	if err = d.Set("keep_alive_timer", dataSourceFlattenRouterbgpNeighborKeepAliveTimer(o["keep-alive-timer"], d, "keep_alive_timer")); err != nil {
		if !fortiAPIPatch(o["keep-alive-timer"]) {
			return fmt.Errorf("error reading keep_alive_timer: %v", err)
		}
	}

	if err = d.Set("link_down_failover", dataSourceFlattenRouterbgpNeighborLinkDownFailover(o["link-down-failover"], d, "link_down_failover")); err != nil {
		if !fortiAPIPatch(o["link-down-failover"]) {
			return fmt.Errorf("error reading link_down_failover: %v", err)
		}
	}

	if err = d.Set("local_as", dataSourceFlattenRouterbgpNeighborLocalAs(o["local-as"], d, "local_as")); err != nil {
		if !fortiAPIPatch(o["local-as"]) {
			return fmt.Errorf("error reading local_as: %v", err)
		}
	}

	if err = d.Set("local_as_no_prepend", dataSourceFlattenRouterbgpNeighborLocalAsNoPrepend(o["local-as-no-prepend"], d, "local_as_no_prepend")); err != nil {
		if !fortiAPIPatch(o["local-as-no-prepend"]) {
			return fmt.Errorf("error reading local_as_no_prepend: %v", err)
		}
	}

	if err = d.Set("local_as_replace_as", dataSourceFlattenRouterbgpNeighborLocalAsReplaceAs(o["local-as-replace-as"], d, "local_as_replace_as")); err != nil {
		if !fortiAPIPatch(o["local-as-replace-as"]) {
			return fmt.Errorf("error reading local_as_replace_as: %v", err)
		}
	}

	if err = d.Set("maximum_prefix", dataSourceFlattenRouterbgpNeighborMaximumPrefix(o["maximum-prefix"], d, "maximum_prefix")); err != nil {
		if !fortiAPIPatch(o["maximum-prefix"]) {
			return fmt.Errorf("error reading maximum_prefix: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold", dataSourceFlattenRouterbgpNeighborMaximumPrefixThreshold(o["maximum-prefix-threshold"], d, "maximum_prefix_threshold")); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-threshold"]) {
			return fmt.Errorf("error reading maximum_prefix_threshold: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold6", dataSourceFlattenRouterbgpNeighborMaximumPrefixThreshold6(o["maximum-prefix-threshold6"], d, "maximum_prefix_threshold6")); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-threshold6"]) {
			return fmt.Errorf("error reading maximum_prefix_threshold6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only", dataSourceFlattenRouterbgpNeighborMaximumPrefixWarningOnly(o["maximum-prefix-warning-only"], d, "maximum_prefix_warning_only")); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-warning-only"]) {
			return fmt.Errorf("error reading maximum_prefix_warning_only: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only6", dataSourceFlattenRouterbgpNeighborMaximumPrefixWarningOnly6(o["maximum-prefix-warning-only6"], d, "maximum_prefix_warning_only6")); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-warning-only6"]) {
			return fmt.Errorf("error reading maximum_prefix_warning_only6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix6", dataSourceFlattenRouterbgpNeighborMaximumPrefix6(o["maximum-prefix6"], d, "maximum_prefix6")); err != nil {
		if !fortiAPIPatch(o["maximum-prefix6"]) {
			return fmt.Errorf("error reading maximum_prefix6: %v", err)
		}
	}

	if err = d.Set("next_hop_self", dataSourceFlattenRouterbgpNeighborNextHopSelf(o["next-hop-self"], d, "next_hop_self")); err != nil {
		if !fortiAPIPatch(o["next-hop-self"]) {
			return fmt.Errorf("error reading next_hop_self: %v", err)
		}
	}

	if err = d.Set("next_hop_self_rr", dataSourceFlattenRouterbgpNeighborNextHopSelfRr(o["next-hop-self-rr"], d, "next_hop_self_rr")); err != nil {
		if !fortiAPIPatch(o["next-hop-self-rr"]) {
			return fmt.Errorf("error reading next_hop_self_rr: %v", err)
		}
	}

	if err = d.Set("next_hop_self_rr6", dataSourceFlattenRouterbgpNeighborNextHopSelfRr6(o["next-hop-self-rr6"], d, "next_hop_self_rr6")); err != nil {
		if !fortiAPIPatch(o["next-hop-self-rr6"]) {
			return fmt.Errorf("error reading next_hop_self_rr6: %v", err)
		}
	}

	if err = d.Set("next_hop_self6", dataSourceFlattenRouterbgpNeighborNextHopSelf6(o["next-hop-self6"], d, "next_hop_self6")); err != nil {
		if !fortiAPIPatch(o["next-hop-self6"]) {
			return fmt.Errorf("error reading next_hop_self6: %v", err)
		}
	}

	if err = d.Set("override_capability", dataSourceFlattenRouterbgpNeighborOverrideCapability(o["override-capability"], d, "override_capability")); err != nil {
		if !fortiAPIPatch(o["override-capability"]) {
			return fmt.Errorf("error reading override_capability: %v", err)
		}
	}

	if err = d.Set("passive", dataSourceFlattenRouterbgpNeighborPassive(o["passive"], d, "passive")); err != nil {
		if !fortiAPIPatch(o["passive"]) {
			return fmt.Errorf("error reading passive: %v", err)
		}
	}

	if err = d.Set("prefix_list_in", dataSourceFlattenRouterbgpNeighborPrefixListIn(o["prefix-list-in"], d, "prefix_list_in")); err != nil {
		if !fortiAPIPatch(o["prefix-list-in"]) {
			return fmt.Errorf("error reading prefix_list_in: %v", err)
		}
	}

	if err = d.Set("prefix_list_in6", dataSourceFlattenRouterbgpNeighborPrefixListIn6(o["prefix-list-in6"], d, "prefix_list_in6")); err != nil {
		if !fortiAPIPatch(o["prefix-list-in6"]) {
			return fmt.Errorf("error reading prefix_list_in6: %v", err)
		}
	}

	if err = d.Set("prefix_list_out", dataSourceFlattenRouterbgpNeighborPrefixListOut(o["prefix-list-out"], d, "prefix_list_out")); err != nil {
		if !fortiAPIPatch(o["prefix-list-out"]) {
			return fmt.Errorf("error reading prefix_list_out: %v", err)
		}
	}

	if err = d.Set("prefix_list_out6", dataSourceFlattenRouterbgpNeighborPrefixListOut6(o["prefix-list-out6"], d, "prefix_list_out6")); err != nil {
		if !fortiAPIPatch(o["prefix-list-out6"]) {
			return fmt.Errorf("error reading prefix_list_out6: %v", err)
		}
	}

	if err = d.Set("remote_as", dataSourceFlattenRouterbgpNeighborRemoteAs(o["remote-as"], d, "remote_as")); err != nil {
		if !fortiAPIPatch(o["remote-as"]) {
			return fmt.Errorf("error reading remote_as: %v", err)
		}
	}

	if err = d.Set("remove_private_as", dataSourceFlattenRouterbgpNeighborRemovePrivateAs(o["remove-private-as"], d, "remove_private_as")); err != nil {
		if !fortiAPIPatch(o["remove-private-as"]) {
			return fmt.Errorf("error reading remove_private_as: %v", err)
		}
	}

	if err = d.Set("remove_private_as6", dataSourceFlattenRouterbgpNeighborRemovePrivateAs6(o["remove-private-as6"], d, "remove_private_as6")); err != nil {
		if !fortiAPIPatch(o["remove-private-as6"]) {
			return fmt.Errorf("error reading remove_private_as6: %v", err)
		}
	}

	if err = d.Set("restart_time", dataSourceFlattenRouterbgpNeighborRestartTime(o["restart-time"], d, "restart_time")); err != nil {
		if !fortiAPIPatch(o["restart-time"]) {
			return fmt.Errorf("error reading restart_time: %v", err)
		}
	}

	if err = d.Set("retain_stale_time", dataSourceFlattenRouterbgpNeighborRetainStaleTime(o["retain-stale-time"], d, "retain_stale_time")); err != nil {
		if !fortiAPIPatch(o["retain-stale-time"]) {
			return fmt.Errorf("error reading retain_stale_time: %v", err)
		}
	}

	if err = d.Set("route_map_in", dataSourceFlattenRouterbgpNeighborRouteMapIn(o["route-map-in"], d, "route_map_in")); err != nil {
		if !fortiAPIPatch(o["route-map-in"]) {
			return fmt.Errorf("error reading route_map_in: %v", err)
		}
	}

	if err = d.Set("route_map_in6", dataSourceFlattenRouterbgpNeighborRouteMapIn6(o["route-map-in6"], d, "route_map_in6")); err != nil {
		if !fortiAPIPatch(o["route-map-in6"]) {
			return fmt.Errorf("error reading route_map_in6: %v", err)
		}
	}

	if err = d.Set("route_map_out", dataSourceFlattenRouterbgpNeighborRouteMapOut(o["route-map-out"], d, "route_map_out")); err != nil {
		if !fortiAPIPatch(o["route-map-out"]) {
			return fmt.Errorf("error reading route_map_out: %v", err)
		}
	}

	if err = d.Set("route_map_out_preferable", dataSourceFlattenRouterbgpNeighborRouteMapOutPreferable(o["route-map-out-preferable"], d, "route_map_out_preferable")); err != nil {
		if !fortiAPIPatch(o["route-map-out-preferable"]) {
			return fmt.Errorf("error reading route_map_out_preferable: %v", err)
		}
	}

	if err = d.Set("route_map_out6", dataSourceFlattenRouterbgpNeighborRouteMapOut6(o["route-map-out6"], d, "route_map_out6")); err != nil {
		if !fortiAPIPatch(o["route-map-out6"]) {
			return fmt.Errorf("error reading route_map_out6: %v", err)
		}
	}

	if err = d.Set("route_map_out6_preferable", dataSourceFlattenRouterbgpNeighborRouteMapOut6Preferable(o["route-map-out6-preferable"], d, "route_map_out6_preferable")); err != nil {
		if !fortiAPIPatch(o["route-map-out6-preferable"]) {
			return fmt.Errorf("error reading route_map_out6_preferable: %v", err)
		}
	}

	if err = d.Set("route_reflector_client", dataSourceFlattenRouterbgpNeighborRouteReflectorClient(o["route-reflector-client"], d, "route_reflector_client")); err != nil {
		if !fortiAPIPatch(o["route-reflector-client"]) {
			return fmt.Errorf("error reading route_reflector_client: %v", err)
		}
	}

	if err = d.Set("route_reflector_client6", dataSourceFlattenRouterbgpNeighborRouteReflectorClient6(o["route-reflector-client6"], d, "route_reflector_client6")); err != nil {
		if !fortiAPIPatch(o["route-reflector-client6"]) {
			return fmt.Errorf("error reading route_reflector_client6: %v", err)
		}
	}

	if err = d.Set("route_server_client", dataSourceFlattenRouterbgpNeighborRouteServerClient(o["route-server-client"], d, "route_server_client")); err != nil {
		if !fortiAPIPatch(o["route-server-client"]) {
			return fmt.Errorf("error reading route_server_client: %v", err)
		}
	}

	if err = d.Set("route_server_client6", dataSourceFlattenRouterbgpNeighborRouteServerClient6(o["route-server-client6"], d, "route_server_client6")); err != nil {
		if !fortiAPIPatch(o["route-server-client6"]) {
			return fmt.Errorf("error reading route_server_client6: %v", err)
		}
	}

	if err = d.Set("send_community", dataSourceFlattenRouterbgpNeighborSendCommunity(o["send-community"], d, "send_community")); err != nil {
		if !fortiAPIPatch(o["send-community"]) {
			return fmt.Errorf("error reading send_community: %v", err)
		}
	}

	if err = d.Set("send_community6", dataSourceFlattenRouterbgpNeighborSendCommunity6(o["send-community6"], d, "send_community6")); err != nil {
		if !fortiAPIPatch(o["send-community6"]) {
			return fmt.Errorf("error reading send_community6: %v", err)
		}
	}

	if err = d.Set("shutdown", dataSourceFlattenRouterbgpNeighborShutdown(o["shutdown"], d, "shutdown")); err != nil {
		if !fortiAPIPatch(o["shutdown"]) {
			return fmt.Errorf("error reading shutdown: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration", dataSourceFlattenRouterbgpNeighborSoftReconfiguration(o["soft-reconfiguration"], d, "soft_reconfiguration")); err != nil {
		if !fortiAPIPatch(o["soft-reconfiguration"]) {
			return fmt.Errorf("error reading soft_reconfiguration: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration6", dataSourceFlattenRouterbgpNeighborSoftReconfiguration6(o["soft-reconfiguration6"], d, "soft_reconfiguration6")); err != nil {
		if !fortiAPIPatch(o["soft-reconfiguration6"]) {
			return fmt.Errorf("error reading soft_reconfiguration6: %v", err)
		}
	}

	if err = d.Set("stale_route", dataSourceFlattenRouterbgpNeighborStaleRoute(o["stale-route"], d, "stale_route")); err != nil {
		if !fortiAPIPatch(o["stale-route"]) {
			return fmt.Errorf("error reading stale_route: %v", err)
		}
	}

	if err = d.Set("strict_capability_match", dataSourceFlattenRouterbgpNeighborStrictCapabilityMatch(o["strict-capability-match"], d, "strict_capability_match")); err != nil {
		if !fortiAPIPatch(o["strict-capability-match"]) {
			return fmt.Errorf("error reading strict_capability_match: %v", err)
		}
	}

	if err = d.Set("unsuppress_map", dataSourceFlattenRouterbgpNeighborUnsuppressMap(o["unsuppress-map"], d, "unsuppress_map")); err != nil {
		if !fortiAPIPatch(o["unsuppress-map"]) {
			return fmt.Errorf("error reading unsuppress_map: %v", err)
		}
	}

	if err = d.Set("unsuppress_map6", dataSourceFlattenRouterbgpNeighborUnsuppressMap6(o["unsuppress-map6"], d, "unsuppress_map6")); err != nil {
		if !fortiAPIPatch(o["unsuppress-map6"]) {
			return fmt.Errorf("error reading unsuppress_map6: %v", err)
		}
	}

	if err = d.Set("update_source", dataSourceFlattenRouterbgpNeighborUpdateSource(o["update-source"], d, "update_source")); err != nil {
		if !fortiAPIPatch(o["update-source"]) {
			return fmt.Errorf("error reading update_source: %v", err)
		}
	}

	if err = d.Set("weight", dataSourceFlattenRouterbgpNeighborWeight(o["weight"], d, "weight")); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("error reading weight: %v", err)
		}
	}

	return nil
}
