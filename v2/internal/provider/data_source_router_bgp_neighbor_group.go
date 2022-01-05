// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: BGP neighbor group table.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceRouterBgpNeighborGroup() *schema.Resource {
	return &schema.Resource{
		Description: "BGP neighbor group table.",

		ReadContext: dataSourceRouterBgpNeighborGroupRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
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

func dataSourceRouterBgpNeighborGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

	c := meta.(*apiClient).Client
	// c.Retries = 1

	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	o, err := c.Cmdb.ReadRouterBgpNeighborGroup(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterBgpNeighborGroup dataSource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] dataSource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	sort := false
	if v, ok := d.GetOk("dynamic_sort_subtable"); ok {
		if b, ok := v.(bool); ok {
			sort = b
		}
	}

	diags := refreshObjectRouterBgpNeighborGroup(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
