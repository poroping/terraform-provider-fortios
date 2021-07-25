// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: BGP neighbor group table.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterbgpNeighborGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterbgpNeighborGroupCreate,
		Read:   resourceRouterbgpNeighborGroupRead,
		Update: resourceRouterbgpNeighborGroupUpdate,
		Delete: resourceRouterbgpNeighborGroupDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"batchid": {
				Type:        schema.TypeInt,
				Description: "Associate with batch. From 6.4.x+. Currently a WIP and broken.",
				Optional:    true,
				Default:     0,
			},
			"allow_append": {
				Type:        schema.TypeBool,
				Description: "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"activate": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable address family IPv4 for this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"activate6": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable address family IPv6 for this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"additional_path": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"send", "receive", "both", "disable"}),
				Description:  "Enable/disable IPv4 additional-path capability.",
				Optional:     true,
				Computed:     true,
			},
			"additional_path6": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"send", "receive", "both", "disable"}),
				Description:  "Enable/disable IPv6 additional-path capability.",
				Optional:     true,
				Computed:     true,
			},
			"adv_additional_path": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 255),
				Description:  "Number of IPv4 additional paths that can be advertised to this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"adv_additional_path6": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 255),
				Description:  "Number of IPv6 additional paths that can be advertised to this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"advertisement_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 600),
				Description:  "Minimum interval (sec) between sending updates.",
				Optional:     true,
				Computed:     true,
			},
			"allowas_in": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),
				Description:  "IPv4 The maximum number of occurrence of my AS number allowed.",
				Optional:     true,
				Computed:     true,
			},
			"allowas_in_enable": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable IPv4 Enable to allow my AS in AS path.",
				Optional:     true,
				Computed:     true,
			},
			"allowas_in_enable6": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable IPv6 Enable to allow my AS in AS path.",
				Optional:     true,
				Computed:     true,
			},
			"allowas_in6": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),
				Description:  "IPv6 The maximum number of occurrence of my AS number allowed.",
				Optional:     true,
				Computed:     true,
			},
			"as_override": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable replace peer AS with own AS for IPv4.",
				Optional:     true,
				Computed:     true,
			},
			"as_override6": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable replace peer AS with own AS for IPv6.",
				Optional:     true,
				Computed:     true,
			},
			"attribute_unchanged": {
				Type: schema.TypeString,

				Description: "IPv4 List of attributes that should be unchanged.",
				Optional:    true,
				Computed:    true,
			},
			"attribute_unchanged6": {
				Type: schema.TypeString,

				Description: "IPv6 List of attributes that should be unchanged.",
				Optional:    true,
				Computed:    true,
			},
			"bfd": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable BFD for this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"capability_default_originate": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable advertise default IPv4 route to this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"capability_default_originate6": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable advertise default IPv6 route to this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"capability_dynamic": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable advertise dynamic capability to this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"capability_graceful_restart": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable advertise IPv4 graceful restart capability to this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"capability_graceful_restart6": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable advertise IPv6 graceful restart capability to this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"capability_orf": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"none", "receive", "send", "both"}),
				Description:  "Accept/Send IPv4 ORF lists to/from this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"capability_orf6": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"none", "receive", "send", "both"}),
				Description:  "Accept/Send IPv6 ORF lists to/from this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"capability_route_refresh": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable advertise route refresh capability to this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"connect_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Description:  "Interval (sec) for connect timer.",
				Optional:     true,
				Computed:     true,
			},
			"default_originate_routemap": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "Route map to specify criteria to originate IPv4 default.",
				Optional:     true,
				Computed:     true,
			},
			"default_originate_routemap6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "Route map to specify criteria to originate IPv6 default.",
				Optional:     true,
				Computed:     true,
			},
			"description": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Description:  "Description.",
				Optional:     true,
				Computed:     true,
			},
			"distribute_list_in": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "Filter for IPv4 updates from this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"distribute_list_in6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "Filter for IPv6 updates from this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"distribute_list_out": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "Filter for IPv4 updates to this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"distribute_list_out6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "Filter for IPv6 updates to this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"dont_capability_negotiate": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Don't negotiate capabilities with this neighbor",
				Optional:     true,
				Computed:     true,
			},
			"ebgp_enforce_multihop": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable allow multi-hop EBGP neighbors.",
				Optional:     true,
				Computed:     true,
			},
			"ebgp_multihop_ttl": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Description:  "EBGP multihop TTL for this peer.",
				Optional:     true,
				Computed:     true,
			},
			"filter_list_in": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "BGP filter for IPv4 inbound routes.",
				Optional:     true,
				Computed:     true,
			},
			"filter_list_in6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "BGP filter for IPv6 inbound routes.",
				Optional:     true,
				Computed:     true,
			},
			"filter_list_out": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "BGP filter for IPv4 outbound routes.",
				Optional:     true,
				Computed:     true,
			},
			"filter_list_out6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "BGP filter for IPv6 outbound routes.",
				Optional:     true,
				Computed:     true,
			},
			"holdtime_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(3, 65535),
				Description:  "Interval (sec) before peer considered dead.",
				Optional:     true,
				Computed:     true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Description:  "Specify outgoing interface for peer connection. For IPv6 peer, the interface should have link-local address.",
				Optional:     true,
				Computed:     true,
			},
			"keep_alive_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Description:  "Keep alive timer interval (sec).",
				Optional:     true,
				Computed:     true,
			},
			"link_down_failover": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable failover upon link down.",
				Optional:     true,
				Computed:     true,
			},
			"local_as": {
				Type: schema.TypeInt,

				Description: "Local AS number of neighbor.",
				Optional:    true,
				Computed:    true,
			},
			"local_as_no_prepend": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Do not prepend local-as to incoming updates.",
				Optional:     true,
				Computed:     true,
			},
			"local_as_replace_as": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Replace real AS with local-as in outgoing updates.",
				Optional:     true,
				Computed:     true,
			},
			"maximum_prefix": {
				Type: schema.TypeInt,

				Description: "Maximum number of IPv4 prefixes to accept from this peer.",
				Optional:    true,
				Computed:    true,
			},
			"maximum_prefix_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),
				Description:  "Maximum IPv4 prefix threshold value (1 - 100 percent).",
				Optional:     true,
				Computed:     true,
			},
			"maximum_prefix_threshold6": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),
				Description:  "Maximum IPv6 prefix threshold value (1 - 100 percent).",
				Optional:     true,
				Computed:     true,
			},
			"maximum_prefix_warning_only": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable IPv4 Only give warning message when limit is exceeded.",
				Optional:     true,
				Computed:     true,
			},
			"maximum_prefix_warning_only6": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable IPv6 Only give warning message when limit is exceeded.",
				Optional:     true,
				Computed:     true,
			},
			"maximum_prefix6": {
				Type: schema.TypeInt,

				Description: "Maximum number of IPv6 prefixes to accept from this peer.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 45),
				Description:  "Neighbor group name.",
				ForceNew:     true,
				Required:     true,
			},
			"next_hop_self": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable IPv4 next-hop calculation for this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"next_hop_self_rr": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable setting nexthop's address to interface's IPv4 address for route-reflector routes.",
				Optional:     true,
				Computed:     true,
			},
			"next_hop_self_rr6": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable setting nexthop's address to interface's IPv6 address for route-reflector routes.",
				Optional:     true,
				Computed:     true,
			},
			"next_hop_self6": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable IPv6 next-hop calculation for this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"override_capability": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable override result of capability negotiation.",
				Optional:     true,
				Computed:     true,
			},
			"passive": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable sending of open messages to this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"prefix_list_in": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "IPv4 Inbound filter for updates from this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"prefix_list_in6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "IPv6 Inbound filter for updates from this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"prefix_list_out": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "IPv4 Outbound filter for updates to this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"prefix_list_out6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "IPv6 Outbound filter for updates to this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"remote_as": {
				Type: schema.TypeInt,

				Description: "AS number of neighbor.",
				Optional:    true,
				Computed:    true,
			},
			"remove_private_as": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable remove private AS number from IPv4 outbound updates.",
				Optional:     true,
				Computed:     true,
			},
			"remove_private_as6": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable remove private AS number from IPv6 outbound updates.",
				Optional:     true,
				Computed:     true,
			},
			"restart_time": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3600),
				Description:  "Graceful restart delay time (sec, 0 = global default).",
				Optional:     true,
				Computed:     true,
			},
			"retain_stale_time": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Description:  "Time to retain stale routes.",
				Optional:     true,
				Computed:     true,
			},
			"route_map_in": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "IPv4 Inbound route map filter.",
				Optional:     true,
				Computed:     true,
			},
			"route_map_in6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "IPv6 Inbound route map filter.",
				Optional:     true,
				Computed:     true,
			},
			"route_map_out": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "IPv4 outbound route map filter.",
				Optional:     true,
				Computed:     true,
			},
			"route_map_out_preferable": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "IPv4 outbound route map filter if the peer is preferred.",
				Optional:     true,
				Computed:     true,
			},
			"route_map_out6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "IPv6 Outbound route map filter.",
				Optional:     true,
				Computed:     true,
			},
			"route_map_out6_preferable": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "IPv6 outbound route map filter if the peer is preferred.",
				Optional:     true,
				Computed:     true,
			},
			"route_reflector_client": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable IPv4 AS route reflector client.",
				Optional:     true,
				Computed:     true,
			},
			"route_reflector_client6": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable IPv6 AS route reflector client.",
				Optional:     true,
				Computed:     true,
			},
			"route_server_client": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable IPv4 AS route server client.",
				Optional:     true,
				Computed:     true,
			},
			"route_server_client6": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable IPv6 AS route server client.",
				Optional:     true,
				Computed:     true,
			},
			"send_community": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"standard", "extended", "both", "disable"}),
				Description:  "IPv4 Send community attribute to neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"send_community6": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnum([]string{"standard", "extended", "both", "disable"}),
				Description:  "IPv6 Send community attribute to neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"shutdown": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable shutdown this neighbor.",
				Optional:     true,
				Computed:     true,
			},
			"soft_reconfiguration": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable allow IPv4 inbound soft reconfiguration.",
				Optional:     true,
				Computed:     true,
			},
			"soft_reconfiguration6": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable allow IPv6 inbound soft reconfiguration.",
				Optional:     true,
				Computed:     true,
			},
			"stale_route": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable stale route after neighbor down.",
				Optional:     true,
				Computed:     true,
			},
			"strict_capability_match": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),
				Description:  "Enable/disable strict capability matching.",
				Optional:     true,
				Computed:     true,
			},
			"unsuppress_map": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "IPv4 Route map to selectively unsuppress suppressed routes.",
				Optional:     true,
				Computed:     true,
			},
			"unsuppress_map6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Description:  "IPv6 Route map to selectively unsuppress suppressed routes.",
				Optional:     true,
				Computed:     true,
			},
			"update_source": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Description:  "Interface to use as source IP/IPv6 address of TCP connections.",
				Optional:     true,
				Computed:     true,
			},
			"weight": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Description:  "Neighbor weight.",
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceRouterbgpNeighborGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

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

	allow_append := false

	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}

	urlparams["allow_append"] = []string{strconv.FormatBool(allow_append)}

	key := "name"
	mkey := ""
	if v, ok := d.GetOk(key); ok {
		if s, ok := v.(string); ok {
			mkey = s
		}
	}

	obj, err := getObjectRouterbgpNeighborGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating RouterbgpNeighborGroup resource while getting object: %v", err)
	}

	if mkey == "" && allow_append {
		return fmt.Errorf("error creating RouterbgpNeighborGroup resource: %q must be set if \"allow_append\" is true", key)
	}

	o := make(map[string]interface{})
	if mkey != "" && allow_append {
		o, err = c.UpdateRouterbgpNeighborGroup(obj, mkey, vdomparam, urlparams, batchid)
	} else {
		o, err = c.CreateRouterbgpNeighborGroup(obj, vdomparam, urlparams, batchid)
	}

	if err != nil {
		return fmt.Errorf("error creating RouterbgpNeighborGroup resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterbgpNeighborGroup")
	}

	return resourceRouterbgpNeighborGroupRead(d, m)
}

func resourceRouterbgpNeighborGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

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

	obj, err := getObjectRouterbgpNeighborGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating RouterbgpNeighborGroup resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterbgpNeighborGroup(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating RouterbgpNeighborGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterbgpNeighborGroup")
	}

	return resourceRouterbgpNeighborGroupRead(d, m)
}

func resourceRouterbgpNeighborGroupDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteRouterbgpNeighborGroup(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting RouterbgpNeighborGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterbgpNeighborGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

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

	o, err := c.ReadRouterbgpNeighborGroup(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading RouterbgpNeighborGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterbgpNeighborGroup(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading RouterbgpNeighborGroup resource from API: %v", err)
	}
	return nil
}

func flattenRouterbgpNeighborGroupActivate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupActivate6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupAdditionalPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupAdditionalPath6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupAdvAdditionalPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupAdvAdditionalPath6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupAdvertisementInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupAllowasIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupAllowasInEnable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupAllowasInEnable6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupAllowasIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupAsOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupAsOverride6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupAttributeUnchanged(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupAttributeUnchanged6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupBfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupCapabilityDefaultOriginate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupCapabilityDefaultOriginate6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupCapabilityDynamic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupCapabilityGracefulRestart(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupCapabilityGracefulRestart6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupCapabilityOrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupCapabilityOrf6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupCapabilityRouteRefresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupConnectTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupDefaultOriginateRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupDefaultOriginateRoutemap6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupDistributeListIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupDistributeListIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupDistributeListOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupDistributeListOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupDontCapabilityNegotiate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupEbgpEnforceMultihop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupEbgpMultihopTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupFilterListIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupFilterListIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupFilterListOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupFilterListOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupHoldtimeTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupKeepAliveTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupLinkDownFailover(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupLocalAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupLocalAsNoPrepend(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupLocalAsReplaceAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupMaximumPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupMaximumPrefixThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupMaximumPrefixThreshold6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupMaximumPrefixWarningOnly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupMaximumPrefixWarningOnly6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupMaximumPrefix6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupNextHopSelf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupNextHopSelfRr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupNextHopSelfRr6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupNextHopSelf6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupOverrideCapability(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupPassive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupPrefixListIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupPrefixListIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupPrefixListOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupPrefixListOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupRemoteAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupRemovePrivateAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupRemovePrivateAs6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupRestartTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupRetainStaleTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupRouteMapIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupRouteMapIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupRouteMapOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupRouteMapOutPreferable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupRouteMapOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupRouteMapOut6Preferable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupRouteReflectorClient(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupRouteReflectorClient6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupRouteServerClient(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupRouteServerClient6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupSendCommunity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupSendCommunity6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupShutdown(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupSoftReconfiguration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupSoftReconfiguration6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupStaleRoute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupStrictCapabilityMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupUnsuppressMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupUnsuppressMap6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupUpdateSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborGroupWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterbgpNeighborGroup(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("activate", flattenRouterbgpNeighborGroupActivate(o["activate"], d, "activate", sv)); err != nil {
		if !fortiAPIPatch(o["activate"]) {
			return fmt.Errorf("error reading activate: %v", err)
		}
	}

	if err = d.Set("activate6", flattenRouterbgpNeighborGroupActivate6(o["activate6"], d, "activate6", sv)); err != nil {
		if !fortiAPIPatch(o["activate6"]) {
			return fmt.Errorf("error reading activate6: %v", err)
		}
	}

	if err = d.Set("additional_path", flattenRouterbgpNeighborGroupAdditionalPath(o["additional-path"], d, "additional_path", sv)); err != nil {
		if !fortiAPIPatch(o["additional-path"]) {
			return fmt.Errorf("error reading additional_path: %v", err)
		}
	}

	if err = d.Set("additional_path6", flattenRouterbgpNeighborGroupAdditionalPath6(o["additional-path6"], d, "additional_path6", sv)); err != nil {
		if !fortiAPIPatch(o["additional-path6"]) {
			return fmt.Errorf("error reading additional_path6: %v", err)
		}
	}

	if err = d.Set("adv_additional_path", flattenRouterbgpNeighborGroupAdvAdditionalPath(o["adv-additional-path"], d, "adv_additional_path", sv)); err != nil {
		if !fortiAPIPatch(o["adv-additional-path"]) {
			return fmt.Errorf("error reading adv_additional_path: %v", err)
		}
	}

	if err = d.Set("adv_additional_path6", flattenRouterbgpNeighborGroupAdvAdditionalPath6(o["adv-additional-path6"], d, "adv_additional_path6", sv)); err != nil {
		if !fortiAPIPatch(o["adv-additional-path6"]) {
			return fmt.Errorf("error reading adv_additional_path6: %v", err)
		}
	}

	if err = d.Set("advertisement_interval", flattenRouterbgpNeighborGroupAdvertisementInterval(o["advertisement-interval"], d, "advertisement_interval", sv)); err != nil {
		if !fortiAPIPatch(o["advertisement-interval"]) {
			return fmt.Errorf("error reading advertisement_interval: %v", err)
		}
	}

	if err = d.Set("allowas_in", flattenRouterbgpNeighborGroupAllowasIn(o["allowas-in"], d, "allowas_in", sv)); err != nil {
		if !fortiAPIPatch(o["allowas-in"]) {
			return fmt.Errorf("error reading allowas_in: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable", flattenRouterbgpNeighborGroupAllowasInEnable(o["allowas-in-enable"], d, "allowas_in_enable", sv)); err != nil {
		if !fortiAPIPatch(o["allowas-in-enable"]) {
			return fmt.Errorf("error reading allowas_in_enable: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable6", flattenRouterbgpNeighborGroupAllowasInEnable6(o["allowas-in-enable6"], d, "allowas_in_enable6", sv)); err != nil {
		if !fortiAPIPatch(o["allowas-in-enable6"]) {
			return fmt.Errorf("error reading allowas_in_enable6: %v", err)
		}
	}

	if err = d.Set("allowas_in6", flattenRouterbgpNeighborGroupAllowasIn6(o["allowas-in6"], d, "allowas_in6", sv)); err != nil {
		if !fortiAPIPatch(o["allowas-in6"]) {
			return fmt.Errorf("error reading allowas_in6: %v", err)
		}
	}

	if err = d.Set("as_override", flattenRouterbgpNeighborGroupAsOverride(o["as-override"], d, "as_override", sv)); err != nil {
		if !fortiAPIPatch(o["as-override"]) {
			return fmt.Errorf("error reading as_override: %v", err)
		}
	}

	if err = d.Set("as_override6", flattenRouterbgpNeighborGroupAsOverride6(o["as-override6"], d, "as_override6", sv)); err != nil {
		if !fortiAPIPatch(o["as-override6"]) {
			return fmt.Errorf("error reading as_override6: %v", err)
		}
	}

	if err = d.Set("attribute_unchanged", flattenRouterbgpNeighborGroupAttributeUnchanged(o["attribute-unchanged"], d, "attribute_unchanged", sv)); err != nil {
		if !fortiAPIPatch(o["attribute-unchanged"]) {
			return fmt.Errorf("error reading attribute_unchanged: %v", err)
		}
	}

	if err = d.Set("attribute_unchanged6", flattenRouterbgpNeighborGroupAttributeUnchanged6(o["attribute-unchanged6"], d, "attribute_unchanged6", sv)); err != nil {
		if !fortiAPIPatch(o["attribute-unchanged6"]) {
			return fmt.Errorf("error reading attribute_unchanged6: %v", err)
		}
	}

	if err = d.Set("bfd", flattenRouterbgpNeighborGroupBfd(o["bfd"], d, "bfd", sv)); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("error reading bfd: %v", err)
		}
	}

	if err = d.Set("capability_default_originate", flattenRouterbgpNeighborGroupCapabilityDefaultOriginate(o["capability-default-originate"], d, "capability_default_originate", sv)); err != nil {
		if !fortiAPIPatch(o["capability-default-originate"]) {
			return fmt.Errorf("error reading capability_default_originate: %v", err)
		}
	}

	if err = d.Set("capability_default_originate6", flattenRouterbgpNeighborGroupCapabilityDefaultOriginate6(o["capability-default-originate6"], d, "capability_default_originate6", sv)); err != nil {
		if !fortiAPIPatch(o["capability-default-originate6"]) {
			return fmt.Errorf("error reading capability_default_originate6: %v", err)
		}
	}

	if err = d.Set("capability_dynamic", flattenRouterbgpNeighborGroupCapabilityDynamic(o["capability-dynamic"], d, "capability_dynamic", sv)); err != nil {
		if !fortiAPIPatch(o["capability-dynamic"]) {
			return fmt.Errorf("error reading capability_dynamic: %v", err)
		}
	}

	if err = d.Set("capability_graceful_restart", flattenRouterbgpNeighborGroupCapabilityGracefulRestart(o["capability-graceful-restart"], d, "capability_graceful_restart", sv)); err != nil {
		if !fortiAPIPatch(o["capability-graceful-restart"]) {
			return fmt.Errorf("error reading capability_graceful_restart: %v", err)
		}
	}

	if err = d.Set("capability_graceful_restart6", flattenRouterbgpNeighborGroupCapabilityGracefulRestart6(o["capability-graceful-restart6"], d, "capability_graceful_restart6", sv)); err != nil {
		if !fortiAPIPatch(o["capability-graceful-restart6"]) {
			return fmt.Errorf("error reading capability_graceful_restart6: %v", err)
		}
	}

	if err = d.Set("capability_orf", flattenRouterbgpNeighborGroupCapabilityOrf(o["capability-orf"], d, "capability_orf", sv)); err != nil {
		if !fortiAPIPatch(o["capability-orf"]) {
			return fmt.Errorf("error reading capability_orf: %v", err)
		}
	}

	if err = d.Set("capability_orf6", flattenRouterbgpNeighborGroupCapabilityOrf6(o["capability-orf6"], d, "capability_orf6", sv)); err != nil {
		if !fortiAPIPatch(o["capability-orf6"]) {
			return fmt.Errorf("error reading capability_orf6: %v", err)
		}
	}

	if err = d.Set("capability_route_refresh", flattenRouterbgpNeighborGroupCapabilityRouteRefresh(o["capability-route-refresh"], d, "capability_route_refresh", sv)); err != nil {
		if !fortiAPIPatch(o["capability-route-refresh"]) {
			return fmt.Errorf("error reading capability_route_refresh: %v", err)
		}
	}

	if err = d.Set("connect_timer", flattenRouterbgpNeighborGroupConnectTimer(o["connect-timer"], d, "connect_timer", sv)); err != nil {
		if !fortiAPIPatch(o["connect-timer"]) {
			return fmt.Errorf("error reading connect_timer: %v", err)
		}
	}

	if err = d.Set("default_originate_routemap", flattenRouterbgpNeighborGroupDefaultOriginateRoutemap(o["default-originate-routemap"], d, "default_originate_routemap", sv)); err != nil {
		if !fortiAPIPatch(o["default-originate-routemap"]) {
			return fmt.Errorf("error reading default_originate_routemap: %v", err)
		}
	}

	if err = d.Set("default_originate_routemap6", flattenRouterbgpNeighborGroupDefaultOriginateRoutemap6(o["default-originate-routemap6"], d, "default_originate_routemap6", sv)); err != nil {
		if !fortiAPIPatch(o["default-originate-routemap6"]) {
			return fmt.Errorf("error reading default_originate_routemap6: %v", err)
		}
	}

	if err = d.Set("description", flattenRouterbgpNeighborGroupDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("error reading description: %v", err)
		}
	}

	if err = d.Set("distribute_list_in", flattenRouterbgpNeighborGroupDistributeListIn(o["distribute-list-in"], d, "distribute_list_in", sv)); err != nil {
		if !fortiAPIPatch(o["distribute-list-in"]) {
			return fmt.Errorf("error reading distribute_list_in: %v", err)
		}
	}

	if err = d.Set("distribute_list_in6", flattenRouterbgpNeighborGroupDistributeListIn6(o["distribute-list-in6"], d, "distribute_list_in6", sv)); err != nil {
		if !fortiAPIPatch(o["distribute-list-in6"]) {
			return fmt.Errorf("error reading distribute_list_in6: %v", err)
		}
	}

	if err = d.Set("distribute_list_out", flattenRouterbgpNeighborGroupDistributeListOut(o["distribute-list-out"], d, "distribute_list_out", sv)); err != nil {
		if !fortiAPIPatch(o["distribute-list-out"]) {
			return fmt.Errorf("error reading distribute_list_out: %v", err)
		}
	}

	if err = d.Set("distribute_list_out6", flattenRouterbgpNeighborGroupDistributeListOut6(o["distribute-list-out6"], d, "distribute_list_out6", sv)); err != nil {
		if !fortiAPIPatch(o["distribute-list-out6"]) {
			return fmt.Errorf("error reading distribute_list_out6: %v", err)
		}
	}

	if err = d.Set("dont_capability_negotiate", flattenRouterbgpNeighborGroupDontCapabilityNegotiate(o["dont-capability-negotiate"], d, "dont_capability_negotiate", sv)); err != nil {
		if !fortiAPIPatch(o["dont-capability-negotiate"]) {
			return fmt.Errorf("error reading dont_capability_negotiate: %v", err)
		}
	}

	if err = d.Set("ebgp_enforce_multihop", flattenRouterbgpNeighborGroupEbgpEnforceMultihop(o["ebgp-enforce-multihop"], d, "ebgp_enforce_multihop", sv)); err != nil {
		if !fortiAPIPatch(o["ebgp-enforce-multihop"]) {
			return fmt.Errorf("error reading ebgp_enforce_multihop: %v", err)
		}
	}

	if err = d.Set("ebgp_multihop_ttl", flattenRouterbgpNeighborGroupEbgpMultihopTtl(o["ebgp-multihop-ttl"], d, "ebgp_multihop_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["ebgp-multihop-ttl"]) {
			return fmt.Errorf("error reading ebgp_multihop_ttl: %v", err)
		}
	}

	if err = d.Set("filter_list_in", flattenRouterbgpNeighborGroupFilterListIn(o["filter-list-in"], d, "filter_list_in", sv)); err != nil {
		if !fortiAPIPatch(o["filter-list-in"]) {
			return fmt.Errorf("error reading filter_list_in: %v", err)
		}
	}

	if err = d.Set("filter_list_in6", flattenRouterbgpNeighborGroupFilterListIn6(o["filter-list-in6"], d, "filter_list_in6", sv)); err != nil {
		if !fortiAPIPatch(o["filter-list-in6"]) {
			return fmt.Errorf("error reading filter_list_in6: %v", err)
		}
	}

	if err = d.Set("filter_list_out", flattenRouterbgpNeighborGroupFilterListOut(o["filter-list-out"], d, "filter_list_out", sv)); err != nil {
		if !fortiAPIPatch(o["filter-list-out"]) {
			return fmt.Errorf("error reading filter_list_out: %v", err)
		}
	}

	if err = d.Set("filter_list_out6", flattenRouterbgpNeighborGroupFilterListOut6(o["filter-list-out6"], d, "filter_list_out6", sv)); err != nil {
		if !fortiAPIPatch(o["filter-list-out6"]) {
			return fmt.Errorf("error reading filter_list_out6: %v", err)
		}
	}

	if err = d.Set("holdtime_timer", flattenRouterbgpNeighborGroupHoldtimeTimer(o["holdtime-timer"], d, "holdtime_timer", sv)); err != nil {
		if !fortiAPIPatch(o["holdtime-timer"]) {
			return fmt.Errorf("error reading holdtime_timer: %v", err)
		}
	}

	if err = d.Set("interface", flattenRouterbgpNeighborGroupInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("error reading interface: %v", err)
		}
	}

	if err = d.Set("keep_alive_timer", flattenRouterbgpNeighborGroupKeepAliveTimer(o["keep-alive-timer"], d, "keep_alive_timer", sv)); err != nil {
		if !fortiAPIPatch(o["keep-alive-timer"]) {
			return fmt.Errorf("error reading keep_alive_timer: %v", err)
		}
	}

	if err = d.Set("link_down_failover", flattenRouterbgpNeighborGroupLinkDownFailover(o["link-down-failover"], d, "link_down_failover", sv)); err != nil {
		if !fortiAPIPatch(o["link-down-failover"]) {
			return fmt.Errorf("error reading link_down_failover: %v", err)
		}
	}

	if err = d.Set("local_as", flattenRouterbgpNeighborGroupLocalAs(o["local-as"], d, "local_as", sv)); err != nil {
		if !fortiAPIPatch(o["local-as"]) {
			return fmt.Errorf("error reading local_as: %v", err)
		}
	}

	if err = d.Set("local_as_no_prepend", flattenRouterbgpNeighborGroupLocalAsNoPrepend(o["local-as-no-prepend"], d, "local_as_no_prepend", sv)); err != nil {
		if !fortiAPIPatch(o["local-as-no-prepend"]) {
			return fmt.Errorf("error reading local_as_no_prepend: %v", err)
		}
	}

	if err = d.Set("local_as_replace_as", flattenRouterbgpNeighborGroupLocalAsReplaceAs(o["local-as-replace-as"], d, "local_as_replace_as", sv)); err != nil {
		if !fortiAPIPatch(o["local-as-replace-as"]) {
			return fmt.Errorf("error reading local_as_replace_as: %v", err)
		}
	}

	if err = d.Set("maximum_prefix", flattenRouterbgpNeighborGroupMaximumPrefix(o["maximum-prefix"], d, "maximum_prefix", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-prefix"]) {
			return fmt.Errorf("error reading maximum_prefix: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold", flattenRouterbgpNeighborGroupMaximumPrefixThreshold(o["maximum-prefix-threshold"], d, "maximum_prefix_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-threshold"]) {
			return fmt.Errorf("error reading maximum_prefix_threshold: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold6", flattenRouterbgpNeighborGroupMaximumPrefixThreshold6(o["maximum-prefix-threshold6"], d, "maximum_prefix_threshold6", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-threshold6"]) {
			return fmt.Errorf("error reading maximum_prefix_threshold6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only", flattenRouterbgpNeighborGroupMaximumPrefixWarningOnly(o["maximum-prefix-warning-only"], d, "maximum_prefix_warning_only", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-warning-only"]) {
			return fmt.Errorf("error reading maximum_prefix_warning_only: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only6", flattenRouterbgpNeighborGroupMaximumPrefixWarningOnly6(o["maximum-prefix-warning-only6"], d, "maximum_prefix_warning_only6", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-warning-only6"]) {
			return fmt.Errorf("error reading maximum_prefix_warning_only6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix6", flattenRouterbgpNeighborGroupMaximumPrefix6(o["maximum-prefix6"], d, "maximum_prefix6", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-prefix6"]) {
			return fmt.Errorf("error reading maximum_prefix6: %v", err)
		}
	}

	if err = d.Set("name", flattenRouterbgpNeighborGroupName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("next_hop_self", flattenRouterbgpNeighborGroupNextHopSelf(o["next-hop-self"], d, "next_hop_self", sv)); err != nil {
		if !fortiAPIPatch(o["next-hop-self"]) {
			return fmt.Errorf("error reading next_hop_self: %v", err)
		}
	}

	if err = d.Set("next_hop_self_rr", flattenRouterbgpNeighborGroupNextHopSelfRr(o["next-hop-self-rr"], d, "next_hop_self_rr", sv)); err != nil {
		if !fortiAPIPatch(o["next-hop-self-rr"]) {
			return fmt.Errorf("error reading next_hop_self_rr: %v", err)
		}
	}

	if err = d.Set("next_hop_self_rr6", flattenRouterbgpNeighborGroupNextHopSelfRr6(o["next-hop-self-rr6"], d, "next_hop_self_rr6", sv)); err != nil {
		if !fortiAPIPatch(o["next-hop-self-rr6"]) {
			return fmt.Errorf("error reading next_hop_self_rr6: %v", err)
		}
	}

	if err = d.Set("next_hop_self6", flattenRouterbgpNeighborGroupNextHopSelf6(o["next-hop-self6"], d, "next_hop_self6", sv)); err != nil {
		if !fortiAPIPatch(o["next-hop-self6"]) {
			return fmt.Errorf("error reading next_hop_self6: %v", err)
		}
	}

	if err = d.Set("override_capability", flattenRouterbgpNeighborGroupOverrideCapability(o["override-capability"], d, "override_capability", sv)); err != nil {
		if !fortiAPIPatch(o["override-capability"]) {
			return fmt.Errorf("error reading override_capability: %v", err)
		}
	}

	if err = d.Set("passive", flattenRouterbgpNeighborGroupPassive(o["passive"], d, "passive", sv)); err != nil {
		if !fortiAPIPatch(o["passive"]) {
			return fmt.Errorf("error reading passive: %v", err)
		}
	}

	if err = d.Set("prefix_list_in", flattenRouterbgpNeighborGroupPrefixListIn(o["prefix-list-in"], d, "prefix_list_in", sv)); err != nil {
		if !fortiAPIPatch(o["prefix-list-in"]) {
			return fmt.Errorf("error reading prefix_list_in: %v", err)
		}
	}

	if err = d.Set("prefix_list_in6", flattenRouterbgpNeighborGroupPrefixListIn6(o["prefix-list-in6"], d, "prefix_list_in6", sv)); err != nil {
		if !fortiAPIPatch(o["prefix-list-in6"]) {
			return fmt.Errorf("error reading prefix_list_in6: %v", err)
		}
	}

	if err = d.Set("prefix_list_out", flattenRouterbgpNeighborGroupPrefixListOut(o["prefix-list-out"], d, "prefix_list_out", sv)); err != nil {
		if !fortiAPIPatch(o["prefix-list-out"]) {
			return fmt.Errorf("error reading prefix_list_out: %v", err)
		}
	}

	if err = d.Set("prefix_list_out6", flattenRouterbgpNeighborGroupPrefixListOut6(o["prefix-list-out6"], d, "prefix_list_out6", sv)); err != nil {
		if !fortiAPIPatch(o["prefix-list-out6"]) {
			return fmt.Errorf("error reading prefix_list_out6: %v", err)
		}
	}

	if err = d.Set("remote_as", flattenRouterbgpNeighborGroupRemoteAs(o["remote-as"], d, "remote_as", sv)); err != nil {
		if !fortiAPIPatch(o["remote-as"]) {
			return fmt.Errorf("error reading remote_as: %v", err)
		}
	}

	if err = d.Set("remove_private_as", flattenRouterbgpNeighborGroupRemovePrivateAs(o["remove-private-as"], d, "remove_private_as", sv)); err != nil {
		if !fortiAPIPatch(o["remove-private-as"]) {
			return fmt.Errorf("error reading remove_private_as: %v", err)
		}
	}

	if err = d.Set("remove_private_as6", flattenRouterbgpNeighborGroupRemovePrivateAs6(o["remove-private-as6"], d, "remove_private_as6", sv)); err != nil {
		if !fortiAPIPatch(o["remove-private-as6"]) {
			return fmt.Errorf("error reading remove_private_as6: %v", err)
		}
	}

	if err = d.Set("restart_time", flattenRouterbgpNeighborGroupRestartTime(o["restart-time"], d, "restart_time", sv)); err != nil {
		if !fortiAPIPatch(o["restart-time"]) {
			return fmt.Errorf("error reading restart_time: %v", err)
		}
	}

	if err = d.Set("retain_stale_time", flattenRouterbgpNeighborGroupRetainStaleTime(o["retain-stale-time"], d, "retain_stale_time", sv)); err != nil {
		if !fortiAPIPatch(o["retain-stale-time"]) {
			return fmt.Errorf("error reading retain_stale_time: %v", err)
		}
	}

	if err = d.Set("route_map_in", flattenRouterbgpNeighborGroupRouteMapIn(o["route-map-in"], d, "route_map_in", sv)); err != nil {
		if !fortiAPIPatch(o["route-map-in"]) {
			return fmt.Errorf("error reading route_map_in: %v", err)
		}
	}

	if err = d.Set("route_map_in6", flattenRouterbgpNeighborGroupRouteMapIn6(o["route-map-in6"], d, "route_map_in6", sv)); err != nil {
		if !fortiAPIPatch(o["route-map-in6"]) {
			return fmt.Errorf("error reading route_map_in6: %v", err)
		}
	}

	if err = d.Set("route_map_out", flattenRouterbgpNeighborGroupRouteMapOut(o["route-map-out"], d, "route_map_out", sv)); err != nil {
		if !fortiAPIPatch(o["route-map-out"]) {
			return fmt.Errorf("error reading route_map_out: %v", err)
		}
	}

	if err = d.Set("route_map_out_preferable", flattenRouterbgpNeighborGroupRouteMapOutPreferable(o["route-map-out-preferable"], d, "route_map_out_preferable", sv)); err != nil {
		if !fortiAPIPatch(o["route-map-out-preferable"]) {
			return fmt.Errorf("error reading route_map_out_preferable: %v", err)
		}
	}

	if err = d.Set("route_map_out6", flattenRouterbgpNeighborGroupRouteMapOut6(o["route-map-out6"], d, "route_map_out6", sv)); err != nil {
		if !fortiAPIPatch(o["route-map-out6"]) {
			return fmt.Errorf("error reading route_map_out6: %v", err)
		}
	}

	if err = d.Set("route_map_out6_preferable", flattenRouterbgpNeighborGroupRouteMapOut6Preferable(o["route-map-out6-preferable"], d, "route_map_out6_preferable", sv)); err != nil {
		if !fortiAPIPatch(o["route-map-out6-preferable"]) {
			return fmt.Errorf("error reading route_map_out6_preferable: %v", err)
		}
	}

	if err = d.Set("route_reflector_client", flattenRouterbgpNeighborGroupRouteReflectorClient(o["route-reflector-client"], d, "route_reflector_client", sv)); err != nil {
		if !fortiAPIPatch(o["route-reflector-client"]) {
			return fmt.Errorf("error reading route_reflector_client: %v", err)
		}
	}

	if err = d.Set("route_reflector_client6", flattenRouterbgpNeighborGroupRouteReflectorClient6(o["route-reflector-client6"], d, "route_reflector_client6", sv)); err != nil {
		if !fortiAPIPatch(o["route-reflector-client6"]) {
			return fmt.Errorf("error reading route_reflector_client6: %v", err)
		}
	}

	if err = d.Set("route_server_client", flattenRouterbgpNeighborGroupRouteServerClient(o["route-server-client"], d, "route_server_client", sv)); err != nil {
		if !fortiAPIPatch(o["route-server-client"]) {
			return fmt.Errorf("error reading route_server_client: %v", err)
		}
	}

	if err = d.Set("route_server_client6", flattenRouterbgpNeighborGroupRouteServerClient6(o["route-server-client6"], d, "route_server_client6", sv)); err != nil {
		if !fortiAPIPatch(o["route-server-client6"]) {
			return fmt.Errorf("error reading route_server_client6: %v", err)
		}
	}

	if err = d.Set("send_community", flattenRouterbgpNeighborGroupSendCommunity(o["send-community"], d, "send_community", sv)); err != nil {
		if !fortiAPIPatch(o["send-community"]) {
			return fmt.Errorf("error reading send_community: %v", err)
		}
	}

	if err = d.Set("send_community6", flattenRouterbgpNeighborGroupSendCommunity6(o["send-community6"], d, "send_community6", sv)); err != nil {
		if !fortiAPIPatch(o["send-community6"]) {
			return fmt.Errorf("error reading send_community6: %v", err)
		}
	}

	if err = d.Set("shutdown", flattenRouterbgpNeighborGroupShutdown(o["shutdown"], d, "shutdown", sv)); err != nil {
		if !fortiAPIPatch(o["shutdown"]) {
			return fmt.Errorf("error reading shutdown: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration", flattenRouterbgpNeighborGroupSoftReconfiguration(o["soft-reconfiguration"], d, "soft_reconfiguration", sv)); err != nil {
		if !fortiAPIPatch(o["soft-reconfiguration"]) {
			return fmt.Errorf("error reading soft_reconfiguration: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration6", flattenRouterbgpNeighborGroupSoftReconfiguration6(o["soft-reconfiguration6"], d, "soft_reconfiguration6", sv)); err != nil {
		if !fortiAPIPatch(o["soft-reconfiguration6"]) {
			return fmt.Errorf("error reading soft_reconfiguration6: %v", err)
		}
	}

	if err = d.Set("stale_route", flattenRouterbgpNeighborGroupStaleRoute(o["stale-route"], d, "stale_route", sv)); err != nil {
		if !fortiAPIPatch(o["stale-route"]) {
			return fmt.Errorf("error reading stale_route: %v", err)
		}
	}

	if err = d.Set("strict_capability_match", flattenRouterbgpNeighborGroupStrictCapabilityMatch(o["strict-capability-match"], d, "strict_capability_match", sv)); err != nil {
		if !fortiAPIPatch(o["strict-capability-match"]) {
			return fmt.Errorf("error reading strict_capability_match: %v", err)
		}
	}

	if err = d.Set("unsuppress_map", flattenRouterbgpNeighborGroupUnsuppressMap(o["unsuppress-map"], d, "unsuppress_map", sv)); err != nil {
		if !fortiAPIPatch(o["unsuppress-map"]) {
			return fmt.Errorf("error reading unsuppress_map: %v", err)
		}
	}

	if err = d.Set("unsuppress_map6", flattenRouterbgpNeighborGroupUnsuppressMap6(o["unsuppress-map6"], d, "unsuppress_map6", sv)); err != nil {
		if !fortiAPIPatch(o["unsuppress-map6"]) {
			return fmt.Errorf("error reading unsuppress_map6: %v", err)
		}
	}

	if err = d.Set("update_source", flattenRouterbgpNeighborGroupUpdateSource(o["update-source"], d, "update_source", sv)); err != nil {
		if !fortiAPIPatch(o["update-source"]) {
			return fmt.Errorf("error reading update_source: %v", err)
		}
	}

	if err = d.Set("weight", flattenRouterbgpNeighborGroupWeight(o["weight"], d, "weight", sv)); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("error reading weight: %v", err)
		}
	}

	return nil
}

func expandRouterbgpNeighborGroupActivate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupActivate6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupAdditionalPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupAdditionalPath6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupAdvAdditionalPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupAdvAdditionalPath6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupAdvertisementInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupAllowasIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupAllowasInEnable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupAllowasInEnable6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupAllowasIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupAsOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupAsOverride6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupAttributeUnchanged(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupAttributeUnchanged6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupBfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupCapabilityDefaultOriginate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupCapabilityDefaultOriginate6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupCapabilityDynamic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupCapabilityGracefulRestart(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupCapabilityGracefulRestart6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupCapabilityOrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupCapabilityOrf6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupCapabilityRouteRefresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupConnectTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupDefaultOriginateRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupDefaultOriginateRoutemap6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupDistributeListIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupDistributeListIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupDistributeListOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupDistributeListOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupDontCapabilityNegotiate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupEbgpEnforceMultihop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupEbgpMultihopTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupFilterListIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupFilterListIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupFilterListOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupFilterListOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupHoldtimeTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupKeepAliveTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupLinkDownFailover(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupLocalAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupLocalAsNoPrepend(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupLocalAsReplaceAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupMaximumPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupMaximumPrefixThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupMaximumPrefixThreshold6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupMaximumPrefixWarningOnly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupMaximumPrefixWarningOnly6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupMaximumPrefix6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupNextHopSelf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupNextHopSelfRr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupNextHopSelfRr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupNextHopSelf6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupOverrideCapability(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupPassive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupPrefixListIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupPrefixListIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupPrefixListOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupPrefixListOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupRemoteAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupRemovePrivateAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupRemovePrivateAs6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupRestartTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupRetainStaleTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupRouteMapIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupRouteMapIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupRouteMapOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupRouteMapOutPreferable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupRouteMapOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupRouteMapOut6Preferable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupRouteReflectorClient(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupRouteReflectorClient6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupRouteServerClient(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupRouteServerClient6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupSendCommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupSendCommunity6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupShutdown(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupSoftReconfiguration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupSoftReconfiguration6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupStaleRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupStrictCapabilityMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupUnsuppressMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupUnsuppressMap6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupUpdateSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborGroupWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterbgpNeighborGroup(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("activate"); ok {
		t, err := expandRouterbgpNeighborGroupActivate(d, v, "activate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["activate"] = t
		}
	}

	if v, ok := d.GetOk("activate6"); ok {
		t, err := expandRouterbgpNeighborGroupActivate6(d, v, "activate6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["activate6"] = t
		}
	}

	if v, ok := d.GetOk("additional_path"); ok {
		t, err := expandRouterbgpNeighborGroupAdditionalPath(d, v, "additional_path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path"] = t
		}
	}

	if v, ok := d.GetOk("additional_path6"); ok {
		t, err := expandRouterbgpNeighborGroupAdditionalPath6(d, v, "additional_path6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path6"] = t
		}
	}

	if v, ok := d.GetOk("adv_additional_path"); ok {
		t, err := expandRouterbgpNeighborGroupAdvAdditionalPath(d, v, "adv_additional_path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adv-additional-path"] = t
		}
	}

	if v, ok := d.GetOk("adv_additional_path6"); ok {
		t, err := expandRouterbgpNeighborGroupAdvAdditionalPath6(d, v, "adv_additional_path6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adv-additional-path6"] = t
		}
	}

	if v, ok := d.GetOk("advertisement_interval"); ok {
		t, err := expandRouterbgpNeighborGroupAdvertisementInterval(d, v, "advertisement_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["advertisement-interval"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in"); ok {
		t, err := expandRouterbgpNeighborGroupAllowasIn(d, v, "allowas_in", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_enable"); ok {
		t, err := expandRouterbgpNeighborGroupAllowasInEnable(d, v, "allowas_in_enable", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-enable"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_enable6"); ok {
		t, err := expandRouterbgpNeighborGroupAllowasInEnable6(d, v, "allowas_in_enable6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-enable6"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in6"); ok {
		t, err := expandRouterbgpNeighborGroupAllowasIn6(d, v, "allowas_in6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in6"] = t
		}
	}

	if v, ok := d.GetOk("as_override"); ok {
		t, err := expandRouterbgpNeighborGroupAsOverride(d, v, "as_override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["as-override"] = t
		}
	}

	if v, ok := d.GetOk("as_override6"); ok {
		t, err := expandRouterbgpNeighborGroupAsOverride6(d, v, "as_override6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["as-override6"] = t
		}
	}

	if v, ok := d.GetOk("attribute_unchanged"); ok {
		t, err := expandRouterbgpNeighborGroupAttributeUnchanged(d, v, "attribute_unchanged", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute-unchanged"] = t
		}
	}

	if v, ok := d.GetOk("attribute_unchanged6"); ok {
		t, err := expandRouterbgpNeighborGroupAttributeUnchanged6(d, v, "attribute_unchanged6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute-unchanged6"] = t
		}
	}

	if v, ok := d.GetOk("bfd"); ok {
		t, err := expandRouterbgpNeighborGroupBfd(d, v, "bfd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("capability_default_originate"); ok {
		t, err := expandRouterbgpNeighborGroupCapabilityDefaultOriginate(d, v, "capability_default_originate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-default-originate"] = t
		}
	}

	if v, ok := d.GetOk("capability_default_originate6"); ok {
		t, err := expandRouterbgpNeighborGroupCapabilityDefaultOriginate6(d, v, "capability_default_originate6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-default-originate6"] = t
		}
	}

	if v, ok := d.GetOk("capability_dynamic"); ok {
		t, err := expandRouterbgpNeighborGroupCapabilityDynamic(d, v, "capability_dynamic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-dynamic"] = t
		}
	}

	if v, ok := d.GetOk("capability_graceful_restart"); ok {
		t, err := expandRouterbgpNeighborGroupCapabilityGracefulRestart(d, v, "capability_graceful_restart", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-graceful-restart"] = t
		}
	}

	if v, ok := d.GetOk("capability_graceful_restart6"); ok {
		t, err := expandRouterbgpNeighborGroupCapabilityGracefulRestart6(d, v, "capability_graceful_restart6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-graceful-restart6"] = t
		}
	}

	if v, ok := d.GetOk("capability_orf"); ok {
		t, err := expandRouterbgpNeighborGroupCapabilityOrf(d, v, "capability_orf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-orf"] = t
		}
	}

	if v, ok := d.GetOk("capability_orf6"); ok {
		t, err := expandRouterbgpNeighborGroupCapabilityOrf6(d, v, "capability_orf6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-orf6"] = t
		}
	}

	if v, ok := d.GetOk("capability_route_refresh"); ok {
		t, err := expandRouterbgpNeighborGroupCapabilityRouteRefresh(d, v, "capability_route_refresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-route-refresh"] = t
		}
	}

	if v, ok := d.GetOk("connect_timer"); ok {
		t, err := expandRouterbgpNeighborGroupConnectTimer(d, v, "connect_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connect-timer"] = t
		}
	}

	if v, ok := d.GetOk("default_originate_routemap"); ok {
		t, err := expandRouterbgpNeighborGroupDefaultOriginateRoutemap(d, v, "default_originate_routemap", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-originate-routemap"] = t
		}
	}

	if v, ok := d.GetOk("default_originate_routemap6"); ok {
		t, err := expandRouterbgpNeighborGroupDefaultOriginateRoutemap6(d, v, "default_originate_routemap6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-originate-routemap6"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandRouterbgpNeighborGroupDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_in"); ok {
		t, err := expandRouterbgpNeighborGroupDistributeListIn(d, v, "distribute_list_in", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-in"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_in6"); ok {
		t, err := expandRouterbgpNeighborGroupDistributeListIn6(d, v, "distribute_list_in6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-in6"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_out"); ok {
		t, err := expandRouterbgpNeighborGroupDistributeListOut(d, v, "distribute_list_out", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-out"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_out6"); ok {
		t, err := expandRouterbgpNeighborGroupDistributeListOut6(d, v, "distribute_list_out6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-out6"] = t
		}
	}

	if v, ok := d.GetOk("dont_capability_negotiate"); ok {
		t, err := expandRouterbgpNeighborGroupDontCapabilityNegotiate(d, v, "dont_capability_negotiate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dont-capability-negotiate"] = t
		}
	}

	if v, ok := d.GetOk("ebgp_enforce_multihop"); ok {
		t, err := expandRouterbgpNeighborGroupEbgpEnforceMultihop(d, v, "ebgp_enforce_multihop", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ebgp-enforce-multihop"] = t
		}
	}

	if v, ok := d.GetOk("ebgp_multihop_ttl"); ok {
		t, err := expandRouterbgpNeighborGroupEbgpMultihopTtl(d, v, "ebgp_multihop_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ebgp-multihop-ttl"] = t
		}
	}

	if v, ok := d.GetOk("filter_list_in"); ok {
		t, err := expandRouterbgpNeighborGroupFilterListIn(d, v, "filter_list_in", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-in"] = t
		}
	}

	if v, ok := d.GetOk("filter_list_in6"); ok {
		t, err := expandRouterbgpNeighborGroupFilterListIn6(d, v, "filter_list_in6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-in6"] = t
		}
	}

	if v, ok := d.GetOk("filter_list_out"); ok {
		t, err := expandRouterbgpNeighborGroupFilterListOut(d, v, "filter_list_out", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-out"] = t
		}
	}

	if v, ok := d.GetOk("filter_list_out6"); ok {
		t, err := expandRouterbgpNeighborGroupFilterListOut6(d, v, "filter_list_out6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-out6"] = t
		}
	}

	if v, ok := d.GetOk("holdtime_timer"); ok {
		t, err := expandRouterbgpNeighborGroupHoldtimeTimer(d, v, "holdtime_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["holdtime-timer"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandRouterbgpNeighborGroupInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("keep_alive_timer"); ok {
		t, err := expandRouterbgpNeighborGroupKeepAliveTimer(d, v, "keep_alive_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keep-alive-timer"] = t
		}
	}

	if v, ok := d.GetOk("link_down_failover"); ok {
		t, err := expandRouterbgpNeighborGroupLinkDownFailover(d, v, "link_down_failover", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-down-failover"] = t
		}
	}

	if v, ok := d.GetOk("local_as"); ok {
		t, err := expandRouterbgpNeighborGroupLocalAs(d, v, "local_as", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-as"] = t
		}
	}

	if v, ok := d.GetOk("local_as_no_prepend"); ok {
		t, err := expandRouterbgpNeighborGroupLocalAsNoPrepend(d, v, "local_as_no_prepend", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-as-no-prepend"] = t
		}
	}

	if v, ok := d.GetOk("local_as_replace_as"); ok {
		t, err := expandRouterbgpNeighborGroupLocalAsReplaceAs(d, v, "local_as_replace_as", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-as-replace-as"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix"); ok {
		t, err := expandRouterbgpNeighborGroupMaximumPrefix(d, v, "maximum_prefix", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_threshold"); ok {
		t, err := expandRouterbgpNeighborGroupMaximumPrefixThreshold(d, v, "maximum_prefix_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-threshold"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_threshold6"); ok {
		t, err := expandRouterbgpNeighborGroupMaximumPrefixThreshold6(d, v, "maximum_prefix_threshold6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-threshold6"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_warning_only"); ok {
		t, err := expandRouterbgpNeighborGroupMaximumPrefixWarningOnly(d, v, "maximum_prefix_warning_only", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-warning-only"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_warning_only6"); ok {
		t, err := expandRouterbgpNeighborGroupMaximumPrefixWarningOnly6(d, v, "maximum_prefix_warning_only6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-warning-only6"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix6"); ok {
		t, err := expandRouterbgpNeighborGroupMaximumPrefix6(d, v, "maximum_prefix6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix6"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandRouterbgpNeighborGroupName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("next_hop_self"); ok {
		t, err := expandRouterbgpNeighborGroupNextHopSelf(d, v, "next_hop_self", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop-self"] = t
		}
	}

	if v, ok := d.GetOk("next_hop_self_rr"); ok {
		t, err := expandRouterbgpNeighborGroupNextHopSelfRr(d, v, "next_hop_self_rr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop-self-rr"] = t
		}
	}

	if v, ok := d.GetOk("next_hop_self_rr6"); ok {
		t, err := expandRouterbgpNeighborGroupNextHopSelfRr6(d, v, "next_hop_self_rr6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop-self-rr6"] = t
		}
	}

	if v, ok := d.GetOk("next_hop_self6"); ok {
		t, err := expandRouterbgpNeighborGroupNextHopSelf6(d, v, "next_hop_self6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop-self6"] = t
		}
	}

	if v, ok := d.GetOk("override_capability"); ok {
		t, err := expandRouterbgpNeighborGroupOverrideCapability(d, v, "override_capability", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-capability"] = t
		}
	}

	if v, ok := d.GetOk("passive"); ok {
		t, err := expandRouterbgpNeighborGroupPassive(d, v, "passive", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passive"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_in"); ok {
		t, err := expandRouterbgpNeighborGroupPrefixListIn(d, v, "prefix_list_in", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-in"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_in6"); ok {
		t, err := expandRouterbgpNeighborGroupPrefixListIn6(d, v, "prefix_list_in6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-in6"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_out"); ok {
		t, err := expandRouterbgpNeighborGroupPrefixListOut(d, v, "prefix_list_out", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-out"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_out6"); ok {
		t, err := expandRouterbgpNeighborGroupPrefixListOut6(d, v, "prefix_list_out6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-out6"] = t
		}
	}

	if v, ok := d.GetOk("remote_as"); ok {
		t, err := expandRouterbgpNeighborGroupRemoteAs(d, v, "remote_as", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-as"] = t
		}
	}

	if v, ok := d.GetOk("remove_private_as"); ok {
		t, err := expandRouterbgpNeighborGroupRemovePrivateAs(d, v, "remove_private_as", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-private-as"] = t
		}
	}

	if v, ok := d.GetOk("remove_private_as6"); ok {
		t, err := expandRouterbgpNeighborGroupRemovePrivateAs6(d, v, "remove_private_as6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-private-as6"] = t
		}
	}

	if v, ok := d.GetOk("restart_time"); ok {
		t, err := expandRouterbgpNeighborGroupRestartTime(d, v, "restart_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restart-time"] = t
		}
	}

	if v, ok := d.GetOk("retain_stale_time"); ok {
		t, err := expandRouterbgpNeighborGroupRetainStaleTime(d, v, "retain_stale_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retain-stale-time"] = t
		}
	}

	if v, ok := d.GetOk("route_map_in"); ok {
		t, err := expandRouterbgpNeighborGroupRouteMapIn(d, v, "route_map_in", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-in"] = t
		}
	}

	if v, ok := d.GetOk("route_map_in6"); ok {
		t, err := expandRouterbgpNeighborGroupRouteMapIn6(d, v, "route_map_in6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-in6"] = t
		}
	}

	if v, ok := d.GetOk("route_map_out"); ok {
		t, err := expandRouterbgpNeighborGroupRouteMapOut(d, v, "route_map_out", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out"] = t
		}
	}

	if v, ok := d.GetOk("route_map_out_preferable"); ok {
		t, err := expandRouterbgpNeighborGroupRouteMapOutPreferable(d, v, "route_map_out_preferable", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out-preferable"] = t
		}
	}

	if v, ok := d.GetOk("route_map_out6"); ok {
		t, err := expandRouterbgpNeighborGroupRouteMapOut6(d, v, "route_map_out6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out6"] = t
		}
	}

	if v, ok := d.GetOk("route_map_out6_preferable"); ok {
		t, err := expandRouterbgpNeighborGroupRouteMapOut6Preferable(d, v, "route_map_out6_preferable", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out6-preferable"] = t
		}
	}

	if v, ok := d.GetOk("route_reflector_client"); ok {
		t, err := expandRouterbgpNeighborGroupRouteReflectorClient(d, v, "route_reflector_client", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-reflector-client"] = t
		}
	}

	if v, ok := d.GetOk("route_reflector_client6"); ok {
		t, err := expandRouterbgpNeighborGroupRouteReflectorClient6(d, v, "route_reflector_client6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-reflector-client6"] = t
		}
	}

	if v, ok := d.GetOk("route_server_client"); ok {
		t, err := expandRouterbgpNeighborGroupRouteServerClient(d, v, "route_server_client", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-server-client"] = t
		}
	}

	if v, ok := d.GetOk("route_server_client6"); ok {
		t, err := expandRouterbgpNeighborGroupRouteServerClient6(d, v, "route_server_client6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-server-client6"] = t
		}
	}

	if v, ok := d.GetOk("send_community"); ok {
		t, err := expandRouterbgpNeighborGroupSendCommunity(d, v, "send_community", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-community"] = t
		}
	}

	if v, ok := d.GetOk("send_community6"); ok {
		t, err := expandRouterbgpNeighborGroupSendCommunity6(d, v, "send_community6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-community6"] = t
		}
	}

	if v, ok := d.GetOk("shutdown"); ok {
		t, err := expandRouterbgpNeighborGroupShutdown(d, v, "shutdown", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shutdown"] = t
		}
	}

	if v, ok := d.GetOk("soft_reconfiguration"); ok {
		t, err := expandRouterbgpNeighborGroupSoftReconfiguration(d, v, "soft_reconfiguration", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["soft-reconfiguration"] = t
		}
	}

	if v, ok := d.GetOk("soft_reconfiguration6"); ok {
		t, err := expandRouterbgpNeighborGroupSoftReconfiguration6(d, v, "soft_reconfiguration6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["soft-reconfiguration6"] = t
		}
	}

	if v, ok := d.GetOk("stale_route"); ok {
		t, err := expandRouterbgpNeighborGroupStaleRoute(d, v, "stale_route", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stale-route"] = t
		}
	}

	if v, ok := d.GetOk("strict_capability_match"); ok {
		t, err := expandRouterbgpNeighborGroupStrictCapabilityMatch(d, v, "strict_capability_match", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strict-capability-match"] = t
		}
	}

	if v, ok := d.GetOk("unsuppress_map"); ok {
		t, err := expandRouterbgpNeighborGroupUnsuppressMap(d, v, "unsuppress_map", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unsuppress-map"] = t
		}
	}

	if v, ok := d.GetOk("unsuppress_map6"); ok {
		t, err := expandRouterbgpNeighborGroupUnsuppressMap6(d, v, "unsuppress_map6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unsuppress-map6"] = t
		}
	}

	if v, ok := d.GetOk("update_source"); ok {
		t, err := expandRouterbgpNeighborGroupUpdateSource(d, v, "update_source", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-source"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok {
		t, err := expandRouterbgpNeighborGroupWeight(d, v, "weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
