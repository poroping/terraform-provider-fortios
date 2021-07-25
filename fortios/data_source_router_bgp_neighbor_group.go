// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: BGP neighbor group table.

package fortios

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceRouterbgpNeighborGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterbgpNeighborGroupRead,
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
			"name": {
				Type:        schema.TypeString,
				Description: "Neighbor group name.",
				Required:    true,
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

func dataSourceRouterbgpNeighborGroupRead(d *schema.ResourceData, m interface{}) error {
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
	key := "name"
	t := d.Get(key)
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("error describing RouterbgpNeighborGroup: type error")
	}

	o, err := c.ReadRouterbgpNeighborGroup(mkey, vdomparam, urlparams, 0)
	if err != nil {
		return fmt.Errorf("error describing RouterbgpNeighborGroup: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterbgpNeighborGroup(d, o)
	if err != nil {
		return fmt.Errorf("error describing RouterbgpNeighborGroup from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterbgpNeighborGroupActivate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupActivate6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupAdditionalPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupAdditionalPath6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupAdvAdditionalPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupAdvAdditionalPath6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupAdvertisementInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupAllowasIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupAllowasInEnable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupAllowasInEnable6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupAllowasIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupAsOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupAsOverride6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupAttributeUnchanged(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupAttributeUnchanged6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupCapabilityDefaultOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupCapabilityDefaultOriginate6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupCapabilityDynamic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupCapabilityGracefulRestart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupCapabilityGracefulRestart6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupCapabilityOrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupCapabilityOrf6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupCapabilityRouteRefresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupConnectTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupDefaultOriginateRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupDefaultOriginateRoutemap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupDistributeListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupDistributeListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupDistributeListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupDistributeListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupDontCapabilityNegotiate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupEbgpEnforceMultihop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupEbgpMultihopTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupFilterListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupFilterListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupFilterListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupFilterListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupHoldtimeTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupKeepAliveTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupLinkDownFailover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupLocalAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupLocalAsNoPrepend(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupLocalAsReplaceAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupMaximumPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupMaximumPrefixThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupMaximumPrefixThreshold6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupMaximumPrefixWarningOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupMaximumPrefixWarningOnly6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupMaximumPrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupNextHopSelf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupNextHopSelfRr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupNextHopSelfRr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupNextHopSelf6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupOverrideCapability(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupPassive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupPrefixListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupPrefixListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupPrefixListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupPrefixListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupRemoteAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupRemovePrivateAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupRemovePrivateAs6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupRestartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupRetainStaleTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupRouteMapIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupRouteMapIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupRouteMapOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupRouteMapOutPreferable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupRouteMapOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupRouteMapOut6Preferable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupRouteReflectorClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupRouteReflectorClient6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupRouteServerClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupRouteServerClient6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupSendCommunity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupSendCommunity6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupShutdown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupSoftReconfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupSoftReconfiguration6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupStaleRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupStrictCapabilityMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupUnsuppressMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupUnsuppressMap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupUpdateSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborGroupWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterbgpNeighborGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("activate", dataSourceFlattenRouterbgpNeighborGroupActivate(o["activate"], d, "activate")); err != nil {
		if !fortiAPIPatch(o["activate"]) {
			return fmt.Errorf("error reading activate: %v", err)
		}
	}

	if err = d.Set("activate6", dataSourceFlattenRouterbgpNeighborGroupActivate6(o["activate6"], d, "activate6")); err != nil {
		if !fortiAPIPatch(o["activate6"]) {
			return fmt.Errorf("error reading activate6: %v", err)
		}
	}

	if err = d.Set("additional_path", dataSourceFlattenRouterbgpNeighborGroupAdditionalPath(o["additional-path"], d, "additional_path")); err != nil {
		if !fortiAPIPatch(o["additional-path"]) {
			return fmt.Errorf("error reading additional_path: %v", err)
		}
	}

	if err = d.Set("additional_path6", dataSourceFlattenRouterbgpNeighborGroupAdditionalPath6(o["additional-path6"], d, "additional_path6")); err != nil {
		if !fortiAPIPatch(o["additional-path6"]) {
			return fmt.Errorf("error reading additional_path6: %v", err)
		}
	}

	if err = d.Set("adv_additional_path", dataSourceFlattenRouterbgpNeighborGroupAdvAdditionalPath(o["adv-additional-path"], d, "adv_additional_path")); err != nil {
		if !fortiAPIPatch(o["adv-additional-path"]) {
			return fmt.Errorf("error reading adv_additional_path: %v", err)
		}
	}

	if err = d.Set("adv_additional_path6", dataSourceFlattenRouterbgpNeighborGroupAdvAdditionalPath6(o["adv-additional-path6"], d, "adv_additional_path6")); err != nil {
		if !fortiAPIPatch(o["adv-additional-path6"]) {
			return fmt.Errorf("error reading adv_additional_path6: %v", err)
		}
	}

	if err = d.Set("advertisement_interval", dataSourceFlattenRouterbgpNeighborGroupAdvertisementInterval(o["advertisement-interval"], d, "advertisement_interval")); err != nil {
		if !fortiAPIPatch(o["advertisement-interval"]) {
			return fmt.Errorf("error reading advertisement_interval: %v", err)
		}
	}

	if err = d.Set("allowas_in", dataSourceFlattenRouterbgpNeighborGroupAllowasIn(o["allowas-in"], d, "allowas_in")); err != nil {
		if !fortiAPIPatch(o["allowas-in"]) {
			return fmt.Errorf("error reading allowas_in: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable", dataSourceFlattenRouterbgpNeighborGroupAllowasInEnable(o["allowas-in-enable"], d, "allowas_in_enable")); err != nil {
		if !fortiAPIPatch(o["allowas-in-enable"]) {
			return fmt.Errorf("error reading allowas_in_enable: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable6", dataSourceFlattenRouterbgpNeighborGroupAllowasInEnable6(o["allowas-in-enable6"], d, "allowas_in_enable6")); err != nil {
		if !fortiAPIPatch(o["allowas-in-enable6"]) {
			return fmt.Errorf("error reading allowas_in_enable6: %v", err)
		}
	}

	if err = d.Set("allowas_in6", dataSourceFlattenRouterbgpNeighborGroupAllowasIn6(o["allowas-in6"], d, "allowas_in6")); err != nil {
		if !fortiAPIPatch(o["allowas-in6"]) {
			return fmt.Errorf("error reading allowas_in6: %v", err)
		}
	}

	if err = d.Set("as_override", dataSourceFlattenRouterbgpNeighborGroupAsOverride(o["as-override"], d, "as_override")); err != nil {
		if !fortiAPIPatch(o["as-override"]) {
			return fmt.Errorf("error reading as_override: %v", err)
		}
	}

	if err = d.Set("as_override6", dataSourceFlattenRouterbgpNeighborGroupAsOverride6(o["as-override6"], d, "as_override6")); err != nil {
		if !fortiAPIPatch(o["as-override6"]) {
			return fmt.Errorf("error reading as_override6: %v", err)
		}
	}

	if err = d.Set("attribute_unchanged", dataSourceFlattenRouterbgpNeighborGroupAttributeUnchanged(o["attribute-unchanged"], d, "attribute_unchanged")); err != nil {
		if !fortiAPIPatch(o["attribute-unchanged"]) {
			return fmt.Errorf("error reading attribute_unchanged: %v", err)
		}
	}

	if err = d.Set("attribute_unchanged6", dataSourceFlattenRouterbgpNeighborGroupAttributeUnchanged6(o["attribute-unchanged6"], d, "attribute_unchanged6")); err != nil {
		if !fortiAPIPatch(o["attribute-unchanged6"]) {
			return fmt.Errorf("error reading attribute_unchanged6: %v", err)
		}
	}

	if err = d.Set("bfd", dataSourceFlattenRouterbgpNeighborGroupBfd(o["bfd"], d, "bfd")); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("error reading bfd: %v", err)
		}
	}

	if err = d.Set("capability_default_originate", dataSourceFlattenRouterbgpNeighborGroupCapabilityDefaultOriginate(o["capability-default-originate"], d, "capability_default_originate")); err != nil {
		if !fortiAPIPatch(o["capability-default-originate"]) {
			return fmt.Errorf("error reading capability_default_originate: %v", err)
		}
	}

	if err = d.Set("capability_default_originate6", dataSourceFlattenRouterbgpNeighborGroupCapabilityDefaultOriginate6(o["capability-default-originate6"], d, "capability_default_originate6")); err != nil {
		if !fortiAPIPatch(o["capability-default-originate6"]) {
			return fmt.Errorf("error reading capability_default_originate6: %v", err)
		}
	}

	if err = d.Set("capability_dynamic", dataSourceFlattenRouterbgpNeighborGroupCapabilityDynamic(o["capability-dynamic"], d, "capability_dynamic")); err != nil {
		if !fortiAPIPatch(o["capability-dynamic"]) {
			return fmt.Errorf("error reading capability_dynamic: %v", err)
		}
	}

	if err = d.Set("capability_graceful_restart", dataSourceFlattenRouterbgpNeighborGroupCapabilityGracefulRestart(o["capability-graceful-restart"], d, "capability_graceful_restart")); err != nil {
		if !fortiAPIPatch(o["capability-graceful-restart"]) {
			return fmt.Errorf("error reading capability_graceful_restart: %v", err)
		}
	}

	if err = d.Set("capability_graceful_restart6", dataSourceFlattenRouterbgpNeighborGroupCapabilityGracefulRestart6(o["capability-graceful-restart6"], d, "capability_graceful_restart6")); err != nil {
		if !fortiAPIPatch(o["capability-graceful-restart6"]) {
			return fmt.Errorf("error reading capability_graceful_restart6: %v", err)
		}
	}

	if err = d.Set("capability_orf", dataSourceFlattenRouterbgpNeighborGroupCapabilityOrf(o["capability-orf"], d, "capability_orf")); err != nil {
		if !fortiAPIPatch(o["capability-orf"]) {
			return fmt.Errorf("error reading capability_orf: %v", err)
		}
	}

	if err = d.Set("capability_orf6", dataSourceFlattenRouterbgpNeighborGroupCapabilityOrf6(o["capability-orf6"], d, "capability_orf6")); err != nil {
		if !fortiAPIPatch(o["capability-orf6"]) {
			return fmt.Errorf("error reading capability_orf6: %v", err)
		}
	}

	if err = d.Set("capability_route_refresh", dataSourceFlattenRouterbgpNeighborGroupCapabilityRouteRefresh(o["capability-route-refresh"], d, "capability_route_refresh")); err != nil {
		if !fortiAPIPatch(o["capability-route-refresh"]) {
			return fmt.Errorf("error reading capability_route_refresh: %v", err)
		}
	}

	if err = d.Set("connect_timer", dataSourceFlattenRouterbgpNeighborGroupConnectTimer(o["connect-timer"], d, "connect_timer")); err != nil {
		if !fortiAPIPatch(o["connect-timer"]) {
			return fmt.Errorf("error reading connect_timer: %v", err)
		}
	}

	if err = d.Set("default_originate_routemap", dataSourceFlattenRouterbgpNeighborGroupDefaultOriginateRoutemap(o["default-originate-routemap"], d, "default_originate_routemap")); err != nil {
		if !fortiAPIPatch(o["default-originate-routemap"]) {
			return fmt.Errorf("error reading default_originate_routemap: %v", err)
		}
	}

	if err = d.Set("default_originate_routemap6", dataSourceFlattenRouterbgpNeighborGroupDefaultOriginateRoutemap6(o["default-originate-routemap6"], d, "default_originate_routemap6")); err != nil {
		if !fortiAPIPatch(o["default-originate-routemap6"]) {
			return fmt.Errorf("error reading default_originate_routemap6: %v", err)
		}
	}

	if err = d.Set("description", dataSourceFlattenRouterbgpNeighborGroupDescription(o["description"], d, "description")); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("error reading description: %v", err)
		}
	}

	if err = d.Set("distribute_list_in", dataSourceFlattenRouterbgpNeighborGroupDistributeListIn(o["distribute-list-in"], d, "distribute_list_in")); err != nil {
		if !fortiAPIPatch(o["distribute-list-in"]) {
			return fmt.Errorf("error reading distribute_list_in: %v", err)
		}
	}

	if err = d.Set("distribute_list_in6", dataSourceFlattenRouterbgpNeighborGroupDistributeListIn6(o["distribute-list-in6"], d, "distribute_list_in6")); err != nil {
		if !fortiAPIPatch(o["distribute-list-in6"]) {
			return fmt.Errorf("error reading distribute_list_in6: %v", err)
		}
	}

	if err = d.Set("distribute_list_out", dataSourceFlattenRouterbgpNeighborGroupDistributeListOut(o["distribute-list-out"], d, "distribute_list_out")); err != nil {
		if !fortiAPIPatch(o["distribute-list-out"]) {
			return fmt.Errorf("error reading distribute_list_out: %v", err)
		}
	}

	if err = d.Set("distribute_list_out6", dataSourceFlattenRouterbgpNeighborGroupDistributeListOut6(o["distribute-list-out6"], d, "distribute_list_out6")); err != nil {
		if !fortiAPIPatch(o["distribute-list-out6"]) {
			return fmt.Errorf("error reading distribute_list_out6: %v", err)
		}
	}

	if err = d.Set("dont_capability_negotiate", dataSourceFlattenRouterbgpNeighborGroupDontCapabilityNegotiate(o["dont-capability-negotiate"], d, "dont_capability_negotiate")); err != nil {
		if !fortiAPIPatch(o["dont-capability-negotiate"]) {
			return fmt.Errorf("error reading dont_capability_negotiate: %v", err)
		}
	}

	if err = d.Set("ebgp_enforce_multihop", dataSourceFlattenRouterbgpNeighborGroupEbgpEnforceMultihop(o["ebgp-enforce-multihop"], d, "ebgp_enforce_multihop")); err != nil {
		if !fortiAPIPatch(o["ebgp-enforce-multihop"]) {
			return fmt.Errorf("error reading ebgp_enforce_multihop: %v", err)
		}
	}

	if err = d.Set("ebgp_multihop_ttl", dataSourceFlattenRouterbgpNeighborGroupEbgpMultihopTtl(o["ebgp-multihop-ttl"], d, "ebgp_multihop_ttl")); err != nil {
		if !fortiAPIPatch(o["ebgp-multihop-ttl"]) {
			return fmt.Errorf("error reading ebgp_multihop_ttl: %v", err)
		}
	}

	if err = d.Set("filter_list_in", dataSourceFlattenRouterbgpNeighborGroupFilterListIn(o["filter-list-in"], d, "filter_list_in")); err != nil {
		if !fortiAPIPatch(o["filter-list-in"]) {
			return fmt.Errorf("error reading filter_list_in: %v", err)
		}
	}

	if err = d.Set("filter_list_in6", dataSourceFlattenRouterbgpNeighborGroupFilterListIn6(o["filter-list-in6"], d, "filter_list_in6")); err != nil {
		if !fortiAPIPatch(o["filter-list-in6"]) {
			return fmt.Errorf("error reading filter_list_in6: %v", err)
		}
	}

	if err = d.Set("filter_list_out", dataSourceFlattenRouterbgpNeighborGroupFilterListOut(o["filter-list-out"], d, "filter_list_out")); err != nil {
		if !fortiAPIPatch(o["filter-list-out"]) {
			return fmt.Errorf("error reading filter_list_out: %v", err)
		}
	}

	if err = d.Set("filter_list_out6", dataSourceFlattenRouterbgpNeighborGroupFilterListOut6(o["filter-list-out6"], d, "filter_list_out6")); err != nil {
		if !fortiAPIPatch(o["filter-list-out6"]) {
			return fmt.Errorf("error reading filter_list_out6: %v", err)
		}
	}

	if err = d.Set("holdtime_timer", dataSourceFlattenRouterbgpNeighborGroupHoldtimeTimer(o["holdtime-timer"], d, "holdtime_timer")); err != nil {
		if !fortiAPIPatch(o["holdtime-timer"]) {
			return fmt.Errorf("error reading holdtime_timer: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenRouterbgpNeighborGroupInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("error reading interface: %v", err)
		}
	}

	if err = d.Set("keep_alive_timer", dataSourceFlattenRouterbgpNeighborGroupKeepAliveTimer(o["keep-alive-timer"], d, "keep_alive_timer")); err != nil {
		if !fortiAPIPatch(o["keep-alive-timer"]) {
			return fmt.Errorf("error reading keep_alive_timer: %v", err)
		}
	}

	if err = d.Set("link_down_failover", dataSourceFlattenRouterbgpNeighborGroupLinkDownFailover(o["link-down-failover"], d, "link_down_failover")); err != nil {
		if !fortiAPIPatch(o["link-down-failover"]) {
			return fmt.Errorf("error reading link_down_failover: %v", err)
		}
	}

	if err = d.Set("local_as", dataSourceFlattenRouterbgpNeighborGroupLocalAs(o["local-as"], d, "local_as")); err != nil {
		if !fortiAPIPatch(o["local-as"]) {
			return fmt.Errorf("error reading local_as: %v", err)
		}
	}

	if err = d.Set("local_as_no_prepend", dataSourceFlattenRouterbgpNeighborGroupLocalAsNoPrepend(o["local-as-no-prepend"], d, "local_as_no_prepend")); err != nil {
		if !fortiAPIPatch(o["local-as-no-prepend"]) {
			return fmt.Errorf("error reading local_as_no_prepend: %v", err)
		}
	}

	if err = d.Set("local_as_replace_as", dataSourceFlattenRouterbgpNeighborGroupLocalAsReplaceAs(o["local-as-replace-as"], d, "local_as_replace_as")); err != nil {
		if !fortiAPIPatch(o["local-as-replace-as"]) {
			return fmt.Errorf("error reading local_as_replace_as: %v", err)
		}
	}

	if err = d.Set("maximum_prefix", dataSourceFlattenRouterbgpNeighborGroupMaximumPrefix(o["maximum-prefix"], d, "maximum_prefix")); err != nil {
		if !fortiAPIPatch(o["maximum-prefix"]) {
			return fmt.Errorf("error reading maximum_prefix: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold", dataSourceFlattenRouterbgpNeighborGroupMaximumPrefixThreshold(o["maximum-prefix-threshold"], d, "maximum_prefix_threshold")); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-threshold"]) {
			return fmt.Errorf("error reading maximum_prefix_threshold: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold6", dataSourceFlattenRouterbgpNeighborGroupMaximumPrefixThreshold6(o["maximum-prefix-threshold6"], d, "maximum_prefix_threshold6")); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-threshold6"]) {
			return fmt.Errorf("error reading maximum_prefix_threshold6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only", dataSourceFlattenRouterbgpNeighborGroupMaximumPrefixWarningOnly(o["maximum-prefix-warning-only"], d, "maximum_prefix_warning_only")); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-warning-only"]) {
			return fmt.Errorf("error reading maximum_prefix_warning_only: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only6", dataSourceFlattenRouterbgpNeighborGroupMaximumPrefixWarningOnly6(o["maximum-prefix-warning-only6"], d, "maximum_prefix_warning_only6")); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-warning-only6"]) {
			return fmt.Errorf("error reading maximum_prefix_warning_only6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix6", dataSourceFlattenRouterbgpNeighborGroupMaximumPrefix6(o["maximum-prefix6"], d, "maximum_prefix6")); err != nil {
		if !fortiAPIPatch(o["maximum-prefix6"]) {
			return fmt.Errorf("error reading maximum_prefix6: %v", err)
		}
	}

	if err = d.Set("name", dataSourceFlattenRouterbgpNeighborGroupName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("next_hop_self", dataSourceFlattenRouterbgpNeighborGroupNextHopSelf(o["next-hop-self"], d, "next_hop_self")); err != nil {
		if !fortiAPIPatch(o["next-hop-self"]) {
			return fmt.Errorf("error reading next_hop_self: %v", err)
		}
	}

	if err = d.Set("next_hop_self_rr", dataSourceFlattenRouterbgpNeighborGroupNextHopSelfRr(o["next-hop-self-rr"], d, "next_hop_self_rr")); err != nil {
		if !fortiAPIPatch(o["next-hop-self-rr"]) {
			return fmt.Errorf("error reading next_hop_self_rr: %v", err)
		}
	}

	if err = d.Set("next_hop_self_rr6", dataSourceFlattenRouterbgpNeighborGroupNextHopSelfRr6(o["next-hop-self-rr6"], d, "next_hop_self_rr6")); err != nil {
		if !fortiAPIPatch(o["next-hop-self-rr6"]) {
			return fmt.Errorf("error reading next_hop_self_rr6: %v", err)
		}
	}

	if err = d.Set("next_hop_self6", dataSourceFlattenRouterbgpNeighborGroupNextHopSelf6(o["next-hop-self6"], d, "next_hop_self6")); err != nil {
		if !fortiAPIPatch(o["next-hop-self6"]) {
			return fmt.Errorf("error reading next_hop_self6: %v", err)
		}
	}

	if err = d.Set("override_capability", dataSourceFlattenRouterbgpNeighborGroupOverrideCapability(o["override-capability"], d, "override_capability")); err != nil {
		if !fortiAPIPatch(o["override-capability"]) {
			return fmt.Errorf("error reading override_capability: %v", err)
		}
	}

	if err = d.Set("passive", dataSourceFlattenRouterbgpNeighborGroupPassive(o["passive"], d, "passive")); err != nil {
		if !fortiAPIPatch(o["passive"]) {
			return fmt.Errorf("error reading passive: %v", err)
		}
	}

	if err = d.Set("prefix_list_in", dataSourceFlattenRouterbgpNeighborGroupPrefixListIn(o["prefix-list-in"], d, "prefix_list_in")); err != nil {
		if !fortiAPIPatch(o["prefix-list-in"]) {
			return fmt.Errorf("error reading prefix_list_in: %v", err)
		}
	}

	if err = d.Set("prefix_list_in6", dataSourceFlattenRouterbgpNeighborGroupPrefixListIn6(o["prefix-list-in6"], d, "prefix_list_in6")); err != nil {
		if !fortiAPIPatch(o["prefix-list-in6"]) {
			return fmt.Errorf("error reading prefix_list_in6: %v", err)
		}
	}

	if err = d.Set("prefix_list_out", dataSourceFlattenRouterbgpNeighborGroupPrefixListOut(o["prefix-list-out"], d, "prefix_list_out")); err != nil {
		if !fortiAPIPatch(o["prefix-list-out"]) {
			return fmt.Errorf("error reading prefix_list_out: %v", err)
		}
	}

	if err = d.Set("prefix_list_out6", dataSourceFlattenRouterbgpNeighborGroupPrefixListOut6(o["prefix-list-out6"], d, "prefix_list_out6")); err != nil {
		if !fortiAPIPatch(o["prefix-list-out6"]) {
			return fmt.Errorf("error reading prefix_list_out6: %v", err)
		}
	}

	if err = d.Set("remote_as", dataSourceFlattenRouterbgpNeighborGroupRemoteAs(o["remote-as"], d, "remote_as")); err != nil {
		if !fortiAPIPatch(o["remote-as"]) {
			return fmt.Errorf("error reading remote_as: %v", err)
		}
	}

	if err = d.Set("remove_private_as", dataSourceFlattenRouterbgpNeighborGroupRemovePrivateAs(o["remove-private-as"], d, "remove_private_as")); err != nil {
		if !fortiAPIPatch(o["remove-private-as"]) {
			return fmt.Errorf("error reading remove_private_as: %v", err)
		}
	}

	if err = d.Set("remove_private_as6", dataSourceFlattenRouterbgpNeighborGroupRemovePrivateAs6(o["remove-private-as6"], d, "remove_private_as6")); err != nil {
		if !fortiAPIPatch(o["remove-private-as6"]) {
			return fmt.Errorf("error reading remove_private_as6: %v", err)
		}
	}

	if err = d.Set("restart_time", dataSourceFlattenRouterbgpNeighborGroupRestartTime(o["restart-time"], d, "restart_time")); err != nil {
		if !fortiAPIPatch(o["restart-time"]) {
			return fmt.Errorf("error reading restart_time: %v", err)
		}
	}

	if err = d.Set("retain_stale_time", dataSourceFlattenRouterbgpNeighborGroupRetainStaleTime(o["retain-stale-time"], d, "retain_stale_time")); err != nil {
		if !fortiAPIPatch(o["retain-stale-time"]) {
			return fmt.Errorf("error reading retain_stale_time: %v", err)
		}
	}

	if err = d.Set("route_map_in", dataSourceFlattenRouterbgpNeighborGroupRouteMapIn(o["route-map-in"], d, "route_map_in")); err != nil {
		if !fortiAPIPatch(o["route-map-in"]) {
			return fmt.Errorf("error reading route_map_in: %v", err)
		}
	}

	if err = d.Set("route_map_in6", dataSourceFlattenRouterbgpNeighborGroupRouteMapIn6(o["route-map-in6"], d, "route_map_in6")); err != nil {
		if !fortiAPIPatch(o["route-map-in6"]) {
			return fmt.Errorf("error reading route_map_in6: %v", err)
		}
	}

	if err = d.Set("route_map_out", dataSourceFlattenRouterbgpNeighborGroupRouteMapOut(o["route-map-out"], d, "route_map_out")); err != nil {
		if !fortiAPIPatch(o["route-map-out"]) {
			return fmt.Errorf("error reading route_map_out: %v", err)
		}
	}

	if err = d.Set("route_map_out_preferable", dataSourceFlattenRouterbgpNeighborGroupRouteMapOutPreferable(o["route-map-out-preferable"], d, "route_map_out_preferable")); err != nil {
		if !fortiAPIPatch(o["route-map-out-preferable"]) {
			return fmt.Errorf("error reading route_map_out_preferable: %v", err)
		}
	}

	if err = d.Set("route_map_out6", dataSourceFlattenRouterbgpNeighborGroupRouteMapOut6(o["route-map-out6"], d, "route_map_out6")); err != nil {
		if !fortiAPIPatch(o["route-map-out6"]) {
			return fmt.Errorf("error reading route_map_out6: %v", err)
		}
	}

	if err = d.Set("route_map_out6_preferable", dataSourceFlattenRouterbgpNeighborGroupRouteMapOut6Preferable(o["route-map-out6-preferable"], d, "route_map_out6_preferable")); err != nil {
		if !fortiAPIPatch(o["route-map-out6-preferable"]) {
			return fmt.Errorf("error reading route_map_out6_preferable: %v", err)
		}
	}

	if err = d.Set("route_reflector_client", dataSourceFlattenRouterbgpNeighborGroupRouteReflectorClient(o["route-reflector-client"], d, "route_reflector_client")); err != nil {
		if !fortiAPIPatch(o["route-reflector-client"]) {
			return fmt.Errorf("error reading route_reflector_client: %v", err)
		}
	}

	if err = d.Set("route_reflector_client6", dataSourceFlattenRouterbgpNeighborGroupRouteReflectorClient6(o["route-reflector-client6"], d, "route_reflector_client6")); err != nil {
		if !fortiAPIPatch(o["route-reflector-client6"]) {
			return fmt.Errorf("error reading route_reflector_client6: %v", err)
		}
	}

	if err = d.Set("route_server_client", dataSourceFlattenRouterbgpNeighborGroupRouteServerClient(o["route-server-client"], d, "route_server_client")); err != nil {
		if !fortiAPIPatch(o["route-server-client"]) {
			return fmt.Errorf("error reading route_server_client: %v", err)
		}
	}

	if err = d.Set("route_server_client6", dataSourceFlattenRouterbgpNeighborGroupRouteServerClient6(o["route-server-client6"], d, "route_server_client6")); err != nil {
		if !fortiAPIPatch(o["route-server-client6"]) {
			return fmt.Errorf("error reading route_server_client6: %v", err)
		}
	}

	if err = d.Set("send_community", dataSourceFlattenRouterbgpNeighborGroupSendCommunity(o["send-community"], d, "send_community")); err != nil {
		if !fortiAPIPatch(o["send-community"]) {
			return fmt.Errorf("error reading send_community: %v", err)
		}
	}

	if err = d.Set("send_community6", dataSourceFlattenRouterbgpNeighborGroupSendCommunity6(o["send-community6"], d, "send_community6")); err != nil {
		if !fortiAPIPatch(o["send-community6"]) {
			return fmt.Errorf("error reading send_community6: %v", err)
		}
	}

	if err = d.Set("shutdown", dataSourceFlattenRouterbgpNeighborGroupShutdown(o["shutdown"], d, "shutdown")); err != nil {
		if !fortiAPIPatch(o["shutdown"]) {
			return fmt.Errorf("error reading shutdown: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration", dataSourceFlattenRouterbgpNeighborGroupSoftReconfiguration(o["soft-reconfiguration"], d, "soft_reconfiguration")); err != nil {
		if !fortiAPIPatch(o["soft-reconfiguration"]) {
			return fmt.Errorf("error reading soft_reconfiguration: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration6", dataSourceFlattenRouterbgpNeighborGroupSoftReconfiguration6(o["soft-reconfiguration6"], d, "soft_reconfiguration6")); err != nil {
		if !fortiAPIPatch(o["soft-reconfiguration6"]) {
			return fmt.Errorf("error reading soft_reconfiguration6: %v", err)
		}
	}

	if err = d.Set("stale_route", dataSourceFlattenRouterbgpNeighborGroupStaleRoute(o["stale-route"], d, "stale_route")); err != nil {
		if !fortiAPIPatch(o["stale-route"]) {
			return fmt.Errorf("error reading stale_route: %v", err)
		}
	}

	if err = d.Set("strict_capability_match", dataSourceFlattenRouterbgpNeighborGroupStrictCapabilityMatch(o["strict-capability-match"], d, "strict_capability_match")); err != nil {
		if !fortiAPIPatch(o["strict-capability-match"]) {
			return fmt.Errorf("error reading strict_capability_match: %v", err)
		}
	}

	if err = d.Set("unsuppress_map", dataSourceFlattenRouterbgpNeighborGroupUnsuppressMap(o["unsuppress-map"], d, "unsuppress_map")); err != nil {
		if !fortiAPIPatch(o["unsuppress-map"]) {
			return fmt.Errorf("error reading unsuppress_map: %v", err)
		}
	}

	if err = d.Set("unsuppress_map6", dataSourceFlattenRouterbgpNeighborGroupUnsuppressMap6(o["unsuppress-map6"], d, "unsuppress_map6")); err != nil {
		if !fortiAPIPatch(o["unsuppress-map6"]) {
			return fmt.Errorf("error reading unsuppress_map6: %v", err)
		}
	}

	if err = d.Set("update_source", dataSourceFlattenRouterbgpNeighborGroupUpdateSource(o["update-source"], d, "update_source")); err != nil {
		if !fortiAPIPatch(o["update-source"]) {
			return fmt.Errorf("error reading update_source: %v", err)
		}
	}

	if err = d.Set("weight", dataSourceFlattenRouterbgpNeighborGroupWeight(o["weight"], d, "weight")); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("error reading weight: %v", err)
		}
	}

	return nil
}
