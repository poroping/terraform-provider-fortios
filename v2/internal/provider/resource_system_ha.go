// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure HA.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceSystemHa() *schema.Resource {
	return &schema.Resource{
		Description: "Configure HA.",

		CreateContext: resourceSystemHaCreate,
		ReadContext:   resourceSystemHaRead,
		UpdateContext: resourceSystemHaUpdate,
		DeleteContext: resourceSystemHaDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"arps": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 60),

				Description: "Number of gratuitous ARPs (1 - 60). Lower to reduce traffic. Higher to reduce failover time.",
				Optional:    true,
				Computed:    true,
			},
			"arps_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 20),

				Description: "Time between gratuitous ARPs  (1 - 20 sec). Lower to reduce failover time. Higher to reduce traffic.",
				Optional:    true,
				Computed:    true,
			},
			"authentication": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable heartbeat message authentication.",
				Optional:    true,
				Computed:    true,
			},
			"cpu_threshold": {
				Type: schema.TypeString,

				Description: "Dynamic weighted load balancing CPU usage weight and high and low thresholds.",
				Optional:    true,
				Computed:    true,
			},
			"encryption": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable heartbeat message encryption.",
				Optional:    true,
				Computed:    true,
			},
			"failover_hold_time": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 300),

				Description: "Time to wait before failover (0 - 300 sec, default = 0), to avoid flip.",
				Optional:    true,
				Computed:    true,
			},
			"ftp_proxy_threshold": {
				Type: schema.TypeString,

				Description: "Dynamic weighted load balancing weight and high and low number of FTP proxy sessions.",
				Optional:    true,
				Computed:    true,
			},
			"gratuitous_arps": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable gratuitous ARPs. Disable if link-failed-signal enabled.",
				Optional:    true,
				Computed:    true,
			},
			"group_id": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1023),

				Description: "HA group ID  (0 - 1023). Must be the same for all members.",
				Optional:    true,
				Computed:    true,
			},
			"group_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),

				Description: "Cluster group name. Must be the same for all members.",
				Optional:    true,
				Computed:    true,
			},
			"ha_direct": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable using ha-mgmt interface for syslog, SNMP, remote authentication (RADIUS), FortiAnalyzer, FortiSandbox, sFlow, and Netflow.",
				Optional:    true,
				Computed:    true,
			},
			"ha_eth_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 4),

				Description: "HA heartbeat packet Ethertype (4-digit hex).",
				Optional:    true,
				Computed:    true,
			},
			"ha_mgmt_interfaces": {
				Type:        schema.TypeList,
				Description: "Reserve interfaces to manage individual cluster units.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst": {
							Type:         schema.TypeString,
							ValidateFunc: validators.FortiValidateIPv4Classnet,

							Description: "Default route destination for reserved HA management interface.",
							Optional:    true,
							Computed:    true,
						},
						"gateway": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Default route gateway for reserved HA management interface.",
							Optional:    true,
							Computed:    true,
						},
						"gateway6": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "Default IPv6 gateway for reserved HA management interface.",
							Optional:         true,
							Computed:         true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Table ID.",
							Optional:    true,
							Computed:    true,
						},
						"interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Interface to reserve for HA management.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ha_mgmt_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to reserve interfaces to manage individual cluster units.",
				Optional:    true,
				Computed:    true,
			},
			"ha_uptime_diff_margin": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Normally you would only reduce this value for failover testing.",
				Optional:    true,
				Computed:    true,
			},
			"hb_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 20),

				Description: "Time between sending heartbeat packets (1 - 20). Increase to reduce false positives.",
				Optional:    true,
				Computed:    true,
			},
			"hb_interval_in_milliseconds": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"100ms", "10ms"}, false),

				Description: "Number of milliseconds for each heartbeat interval: 100ms or 10ms.",
				Optional:    true,
				Computed:    true,
			},
			"hb_lost_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 60),

				Description: "Number of lost heartbeats to signal a failure (1 - 60). Increase to reduce false positives.",
				Optional:    true,
				Computed:    true,
			},
			"hbdev": {
				Type: schema.TypeString,

				Description: "Heartbeat interfaces. Must be the same for all members.",
				Optional:    true,
				Computed:    true,
			},
			"hc_eth_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 4),

				Description: "Transparent mode HA heartbeat packet Ethertype (4-digit hex).",
				Optional:    true,
				Computed:    true,
			},
			"hello_holddown": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 300),

				Description: "Time to wait before changing from hello to work state (5 - 300 sec).",
				Optional:    true,
				Computed:    true,
			},
			"http_proxy_threshold": {
				Type: schema.TypeString,

				Description: "Dynamic weighted load balancing weight and high and low number of HTTP proxy sessions.",
				Optional:    true,
				Computed:    true,
			},
			"imap_proxy_threshold": {
				Type: schema.TypeString,

				Description: "Dynamic weighted load balancing weight and high and low number of IMAP proxy sessions.",
				Optional:    true,
				Computed:    true,
			},
			"inter_cluster_session_sync": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable synchronization of sessions among HA clusters.",
				Optional:    true,
				Computed:    true,
			},
			"key": {
				Type: schema.TypeString,

				Description: "key",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"l2ep_eth_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 4),

				Description: "Telnet session HA heartbeat packet Ethertype (4-digit hex).",
				Optional:    true,
				Computed:    true,
			},
			"link_failed_signal": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to shut down all interfaces for 1 sec after a failover. Use if gratuitous ARPs do not update network.",
				Optional:    true,
				Computed:    true,
			},
			"load_balance_all": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to load balance TCP sessions. Disable to load balance proxy sessions only.",
				Optional:    true,
				Computed:    true,
			},
			"logical_sn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable usage of the logical serial number.",
				Optional:    true,
				Computed:    true,
			},
			"memory_based_failover": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable memory based failover.",
				Optional:    true,
				Computed:    true,
			},
			"memory_compatible_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable memory compatible mode.",
				Optional:    true,
				Computed:    true,
			},
			"memory_failover_flip_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(6, 2147483647),

				Description: "Time to wait between subsequent memory based failovers in minutes (6 - 2147483647, default = 6).",
				Optional:    true,
				Computed:    true,
			},
			"memory_failover_monitor_period": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 300),

				Description: "Duration of high memory usage before memory based failover is triggered in seconds (1 - 300, default = 60).",
				Optional:    true,
				Computed:    true,
			},
			"memory_failover_sample_rate": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 60),

				Description: "Rate at which memory usage is sampled in order to measure memory usage in seconds (1 - 60, default = 1).",
				Optional:    true,
				Computed:    true,
			},
			"memory_failover_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 95),

				Description: "Memory usage threshold to trigger memory based failover (0 means using conserve mode threshold in system.global).",
				Optional:    true,
				Computed:    true,
			},
			"memory_threshold": {
				Type: schema.TypeString,

				Description: "Dynamic weighted load balancing memory usage weight and high and low thresholds.",
				Optional:    true,
				Computed:    true,
			},
			"mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"standalone", "a-a", "a-p"}, false),

				Description: "HA mode. Must be the same for all members. FGSP requires standalone.",
				Optional:    true,
				Computed:    true,
			},
			"monitor": {
				Type: schema.TypeString,

				Description: "Interfaces to check for port monitoring (or link failure).",
				Optional:    true,
				Computed:    true,
			},
			"multicast_ttl": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 3600),

				Description: "HA multicast TTL on primary (5 - 3600 sec).",
				Optional:    true,
				Computed:    true,
			},
			"nntp_proxy_threshold": {
				Type: schema.TypeString,

				Description: "Dynamic weighted load balancing weight and high and low number of NNTP proxy sessions.",
				Optional:    true,
				Computed:    true,
			},
			"override": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable and increase the priority of the unit that should always be primary.",
				Optional:    true,
				Computed:    true,
			},
			"override_wait_time": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3600),

				Description: "Delay negotiating if override is enabled (0 - 3600 sec). Reduces how often the cluster negotiates.",
				Optional:    true,
				Computed:    true,
			},
			"password": {
				Type: schema.TypeString,

				Description: "Cluster password. Must be the same for all members.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"pingserver_failover_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 50),

				Description: "Remote IP monitoring failover threshold (0 - 50).",
				Optional:    true,
				Computed:    true,
			},
			"pingserver_flip_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(6, 2147483647),

				Description: "Time to wait in minutes before renegotiating after a remote IP monitoring failover.",
				Optional:    true,
				Computed:    true,
			},
			"pingserver_monitor_interface": {
				Type: schema.TypeString,

				Description: "Interfaces to check for remote IP monitoring.",
				Optional:    true,
				Computed:    true,
			},
			"pingserver_secondary_force_reset": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to force the cluster to negotiate after a remote IP monitoring failover.",
				Optional:    true,
				Computed:    true,
			},
			"pingserver_slave_force_reset": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to force the cluster to negotiate after a remote IP monitoring failover.",
				Optional:    true,
				Computed:    true,
			},
			"pop3_proxy_threshold": {
				Type: schema.TypeString,

				Description: "Dynamic weighted load balancing weight and high and low number of POP3 proxy sessions.",
				Optional:    true,
				Computed:    true,
			},
			"priority": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Increase the priority to select the primary unit (0 - 255).",
				Optional:    true,
				Computed:    true,
			},
			"route_hold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3600),

				Description: "Time to wait between routing table updates to the cluster (0 - 3600 sec).",
				Optional:    true,
				Computed:    true,
			},
			"route_ttl": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 3600),

				Description: "TTL for primary unit routes (5 - 3600 sec). Increase to maintain active routes during failover.",
				Optional:    true,
				Computed:    true,
			},
			"route_wait": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3600),

				Description: "Time to wait before sending new routes to the cluster (0 - 3600 sec).",
				Optional:    true,
				Computed:    true,
			},
			"schedule": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "hub", "leastconnection", "round-robin", "weight-round-robin", "random", "ip", "ipport"}, false),

				Description: "Type of A-A load balancing. Use none if you have external load balancers.",
				Optional:    true,
				Computed:    true,
			},
			"secondary_vcluster": {
				Type:        schema.TypeList,
				Description: "Configure virtual cluster 2.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitor": {
							Type: schema.TypeString,

							Description: "Interfaces to check for port monitoring (or link failure).",
							Optional:    true,
							Computed:    true,
						},
						"override": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable and increase the priority of the unit that should always be primary.",
							Optional:    true,
							Computed:    true,
						},
						"override_wait_time": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 3600),

							Description: "Delay negotiating if override is enabled (0 - 3600 sec). Reduces how often the cluster negotiates.",
							Optional:    true,
							Computed:    true,
						},
						"pingserver_failover_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50),

							Description: "Remote IP monitoring failover threshold (0 - 50).",
							Optional:    true,
							Computed:    true,
						},
						"pingserver_monitor_interface": {
							Type: schema.TypeString,

							Description: "Interfaces to check for remote IP monitoring.",
							Optional:    true,
							Computed:    true,
						},
						"pingserver_secondary_force_reset": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to force the cluster to negotiate after a remote IP monitoring failover.",
							Optional:    true,
							Computed:    true,
						},
						"pingserver_slave_force_reset": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to force the cluster to negotiate after a remote IP monitoring failover.",
							Optional:    true,
							Computed:    true,
						},
						"priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Increase the priority to select the primary unit (0 - 255).",
							Optional:    true,
							Computed:    true,
						},
						"vcluster_id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Cluster ID.",
							Optional:    true,
							Computed:    true,
						},
						"vdom": {
							Type: schema.TypeString,

							Description: "VDOMs in virtual cluster 2.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"session_pickup": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable session pickup. Enabling it can reduce session down time when fail over happens.",
				Optional:    true,
				Computed:    true,
			},
			"session_pickup_connectionless": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable UDP and ICMP session sync.",
				Optional:    true,
				Computed:    true,
			},
			"session_pickup_delay": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to sync sessions longer than 30 sec. Only longer lived sessions need to be synced.",
				Optional:    true,
				Computed:    true,
			},
			"session_pickup_expectation": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable session helper expectation session sync for FGSP.",
				Optional:    true,
				Computed:    true,
			},
			"session_pickup_nat": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable NAT session sync for FGSP.",
				Optional:    true,
				Computed:    true,
			},
			"session_sync_dev": {
				Type: schema.TypeString,

				Description: "Offload session-sync process to kernel and sync sessions using connected interface(s) directly.",
				Optional:    true,
				Computed:    true,
			},
			"smtp_proxy_threshold": {
				Type: schema.TypeString,

				Description: "Dynamic weighted load balancing weight and high and low number of SMTP proxy sessions.",
				Optional:    true,
				Computed:    true,
			},
			"ssd_failover": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable automatic HA failover on SSD disk failure.",
				Optional:    true,
				Computed:    true,
			},
			"standalone_config_sync": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FGSP configuration synchronization.",
				Optional:    true,
				Computed:    true,
			},
			"standalone_mgmt_vdom": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable standalone management VDOM.",
				Optional:    true,
				Computed:    true,
			},
			"sync_config": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable configuration synchronization.",
				Optional:    true,
				Computed:    true,
			},
			"sync_packet_balance": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable HA packet distribution to multiple CPUs.",
				Optional:    true,
				Computed:    true,
			},
			"unicast_gateway": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Default route gateway for unicast interface.",
				Optional:    true,
				Computed:    true,
			},
			"unicast_hb": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable unicast heartbeat.",
				Optional:    true,
				Computed:    true,
			},
			"unicast_hb_netmask": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4Netmask,

				Description: "Unicast heartbeat netmask.",
				Optional:    true,
				Computed:    true,
			},
			"unicast_hb_peerip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Unicast heartbeat peer IP.",
				Optional:    true,
				Computed:    true,
			},
			"unicast_peers": {
				Type:        schema.TypeList,
				Description: "Number of unicast peers.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Table ID.",
							Optional:    true,
							Computed:    true,
						},
						"peer_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Unicast peer IP.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"unicast_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable unicast connection.",
				Optional:    true,
				Computed:    true,
			},
			"uninterruptible_primary_wait": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 300),

				Description: "Number of minutes the primary HA unit waits before the secondary HA unit is considered upgraded and the system is started before starting its own upgrade (1 - 300, default = 30).",
				Optional:    true,
				Computed:    true,
			},
			"uninterruptible_upgrade": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to upgrade a cluster without blocking network traffic.",
				Optional:    true,
				Computed:    true,
			},
			"vcluster_id": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Cluster ID.",
				Optional:    true,
				Computed:    true,
			},
			"vcluster2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable virtual cluster 2 for virtual clustering.",
				Optional:    true,
				Computed:    true,
			},
			"vdom": {
				Type: schema.TypeString,

				Description: "VDOMs in virtual cluster 1.",
				Optional:    true,
				Computed:    true,
			},
			"weight": {
				Type: schema.TypeString,

				Description: "Weight-round-robin weight for each cluster unit. Syntax <priority> <weight>.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemHaCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	var diags diag.Diagnostics
	var err error
	// c.Retries = 1
	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	allow_append := false
	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}
	urlparams.AllowAppend = &allow_append

	obj, diags := getObjectSystemHa(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemHa(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemHa")
	}

	return resourceSystemHaRead(ctx, d, meta)
}

func resourceSystemHaUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemHa(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemHa(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemHa resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemHa")
	}

	return resourceSystemHaRead(ctx, d, meta)
}

func resourceSystemHaDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemHa(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemHa(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemHa resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemHaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemHa(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemHa resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	sort := false
	if v, ok := d.GetOk("dynamic_sort_subtable"); ok {
		if b, ok := v.(bool); ok {
			sort = b
		}
	}

	diags := refreshObjectSystemHa(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemHaHaMgmtInterfaces(v *[]models.SystemHaHaMgmtInterfaces, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Dst; tmp != nil {
				v["dst"] = *tmp
			}

			if tmp := cfg.Gateway; tmp != nil {
				v["gateway"] = *tmp
			}

			if tmp := cfg.Gateway6; tmp != nil {
				v["gateway6"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemHaSecondaryVcluster(v *[]models.SystemHaSecondaryVcluster, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Monitor; tmp != nil {
				v["monitor"] = *tmp
			}

			if tmp := cfg.Override; tmp != nil {
				v["override"] = *tmp
			}

			if tmp := cfg.OverrideWaitTime; tmp != nil {
				v["override_wait_time"] = *tmp
			}

			if tmp := cfg.PingserverFailoverThreshold; tmp != nil {
				v["pingserver_failover_threshold"] = *tmp
			}

			if tmp := cfg.PingserverMonitorInterface; tmp != nil {
				v["pingserver_monitor_interface"] = *tmp
			}

			if tmp := cfg.PingserverSecondaryForceReset; tmp != nil {
				v["pingserver_secondary_force_reset"] = *tmp
			}

			if tmp := cfg.PingserverSlaveForceReset; tmp != nil {
				v["pingserver_slave_force_reset"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			if tmp := cfg.VclusterId; tmp != nil {
				v["vcluster_id"] = *tmp
			}

			if tmp := cfg.Vdom; tmp != nil {
				v["vdom"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemHaUnicastPeers(v *[]models.SystemHaUnicastPeers, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.PeerIp; tmp != nil {
				v["peer_ip"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectSystemHa(d *schema.ResourceData, o *models.SystemHa, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Arps != nil {
		v := *o.Arps

		if err = d.Set("arps", v); err != nil {
			return diag.Errorf("error reading arps: %v", err)
		}
	}

	if o.ArpsInterval != nil {
		v := *o.ArpsInterval

		if err = d.Set("arps_interval", v); err != nil {
			return diag.Errorf("error reading arps_interval: %v", err)
		}
	}

	if o.Authentication != nil {
		v := *o.Authentication

		if err = d.Set("authentication", v); err != nil {
			return diag.Errorf("error reading authentication: %v", err)
		}
	}

	if o.CpuThreshold != nil {
		v := *o.CpuThreshold

		if err = d.Set("cpu_threshold", v); err != nil {
			return diag.Errorf("error reading cpu_threshold: %v", err)
		}
	}

	if o.Encryption != nil {
		v := *o.Encryption

		if err = d.Set("encryption", v); err != nil {
			return diag.Errorf("error reading encryption: %v", err)
		}
	}

	if o.FailoverHoldTime != nil {
		v := *o.FailoverHoldTime

		if err = d.Set("failover_hold_time", v); err != nil {
			return diag.Errorf("error reading failover_hold_time: %v", err)
		}
	}

	if o.FtpProxyThreshold != nil {
		v := *o.FtpProxyThreshold

		if err = d.Set("ftp_proxy_threshold", v); err != nil {
			return diag.Errorf("error reading ftp_proxy_threshold: %v", err)
		}
	}

	if o.GratuitousArps != nil {
		v := *o.GratuitousArps

		if err = d.Set("gratuitous_arps", v); err != nil {
			return diag.Errorf("error reading gratuitous_arps: %v", err)
		}
	}

	if o.GroupId != nil {
		v := *o.GroupId

		if err = d.Set("group_id", v); err != nil {
			return diag.Errorf("error reading group_id: %v", err)
		}
	}

	if o.GroupName != nil {
		v := *o.GroupName

		if err = d.Set("group_name", v); err != nil {
			return diag.Errorf("error reading group_name: %v", err)
		}
	}

	if o.HaDirect != nil {
		v := *o.HaDirect

		if err = d.Set("ha_direct", v); err != nil {
			return diag.Errorf("error reading ha_direct: %v", err)
		}
	}

	if o.HaEthType != nil {
		v := *o.HaEthType

		if err = d.Set("ha_eth_type", v); err != nil {
			return diag.Errorf("error reading ha_eth_type: %v", err)
		}
	}

	if o.HaMgmtInterfaces != nil {
		if err = d.Set("ha_mgmt_interfaces", flattenSystemHaHaMgmtInterfaces(o.HaMgmtInterfaces, sort)); err != nil {
			return diag.Errorf("error reading ha_mgmt_interfaces: %v", err)
		}
	}

	if o.HaMgmtStatus != nil {
		v := *o.HaMgmtStatus

		if err = d.Set("ha_mgmt_status", v); err != nil {
			return diag.Errorf("error reading ha_mgmt_status: %v", err)
		}
	}

	if o.HaUptimeDiffMargin != nil {
		v := *o.HaUptimeDiffMargin

		if err = d.Set("ha_uptime_diff_margin", v); err != nil {
			return diag.Errorf("error reading ha_uptime_diff_margin: %v", err)
		}
	}

	if o.HbInterval != nil {
		v := *o.HbInterval

		if err = d.Set("hb_interval", v); err != nil {
			return diag.Errorf("error reading hb_interval: %v", err)
		}
	}

	if o.HbIntervalInMilliseconds != nil {
		v := *o.HbIntervalInMilliseconds

		if err = d.Set("hb_interval_in_milliseconds", v); err != nil {
			return diag.Errorf("error reading hb_interval_in_milliseconds: %v", err)
		}
	}

	if o.HbLostThreshold != nil {
		v := *o.HbLostThreshold

		if err = d.Set("hb_lost_threshold", v); err != nil {
			return diag.Errorf("error reading hb_lost_threshold: %v", err)
		}
	}

	if o.Hbdev != nil {
		v := *o.Hbdev

		if err = d.Set("hbdev", v); err != nil {
			return diag.Errorf("error reading hbdev: %v", err)
		}
	}

	if o.HcEthType != nil {
		v := *o.HcEthType

		if err = d.Set("hc_eth_type", v); err != nil {
			return diag.Errorf("error reading hc_eth_type: %v", err)
		}
	}

	if o.HelloHolddown != nil {
		v := *o.HelloHolddown

		if err = d.Set("hello_holddown", v); err != nil {
			return diag.Errorf("error reading hello_holddown: %v", err)
		}
	}

	if o.HttpProxyThreshold != nil {
		v := *o.HttpProxyThreshold

		if err = d.Set("http_proxy_threshold", v); err != nil {
			return diag.Errorf("error reading http_proxy_threshold: %v", err)
		}
	}

	if o.ImapProxyThreshold != nil {
		v := *o.ImapProxyThreshold

		if err = d.Set("imap_proxy_threshold", v); err != nil {
			return diag.Errorf("error reading imap_proxy_threshold: %v", err)
		}
	}

	if o.InterClusterSessionSync != nil {
		v := *o.InterClusterSessionSync

		if err = d.Set("inter_cluster_session_sync", v); err != nil {
			return diag.Errorf("error reading inter_cluster_session_sync: %v", err)
		}
	}

	if o.Key != nil {
		v := *o.Key

		if err = d.Set("key", v); err != nil {
			return diag.Errorf("error reading key: %v", err)
		}
	}

	if o.L2epEthType != nil {
		v := *o.L2epEthType

		if err = d.Set("l2ep_eth_type", v); err != nil {
			return diag.Errorf("error reading l2ep_eth_type: %v", err)
		}
	}

	if o.LinkFailedSignal != nil {
		v := *o.LinkFailedSignal

		if err = d.Set("link_failed_signal", v); err != nil {
			return diag.Errorf("error reading link_failed_signal: %v", err)
		}
	}

	if o.LoadBalanceAll != nil {
		v := *o.LoadBalanceAll

		if err = d.Set("load_balance_all", v); err != nil {
			return diag.Errorf("error reading load_balance_all: %v", err)
		}
	}

	if o.LogicalSn != nil {
		v := *o.LogicalSn

		if err = d.Set("logical_sn", v); err != nil {
			return diag.Errorf("error reading logical_sn: %v", err)
		}
	}

	if o.MemoryBasedFailover != nil {
		v := *o.MemoryBasedFailover

		if err = d.Set("memory_based_failover", v); err != nil {
			return diag.Errorf("error reading memory_based_failover: %v", err)
		}
	}

	if o.MemoryCompatibleMode != nil {
		v := *o.MemoryCompatibleMode

		if err = d.Set("memory_compatible_mode", v); err != nil {
			return diag.Errorf("error reading memory_compatible_mode: %v", err)
		}
	}

	if o.MemoryFailoverFlipTimeout != nil {
		v := *o.MemoryFailoverFlipTimeout

		if err = d.Set("memory_failover_flip_timeout", v); err != nil {
			return diag.Errorf("error reading memory_failover_flip_timeout: %v", err)
		}
	}

	if o.MemoryFailoverMonitorPeriod != nil {
		v := *o.MemoryFailoverMonitorPeriod

		if err = d.Set("memory_failover_monitor_period", v); err != nil {
			return diag.Errorf("error reading memory_failover_monitor_period: %v", err)
		}
	}

	if o.MemoryFailoverSampleRate != nil {
		v := *o.MemoryFailoverSampleRate

		if err = d.Set("memory_failover_sample_rate", v); err != nil {
			return diag.Errorf("error reading memory_failover_sample_rate: %v", err)
		}
	}

	if o.MemoryFailoverThreshold != nil {
		v := *o.MemoryFailoverThreshold

		if err = d.Set("memory_failover_threshold", v); err != nil {
			return diag.Errorf("error reading memory_failover_threshold: %v", err)
		}
	}

	if o.MemoryThreshold != nil {
		v := *o.MemoryThreshold

		if err = d.Set("memory_threshold", v); err != nil {
			return diag.Errorf("error reading memory_threshold: %v", err)
		}
	}

	if o.Mode != nil {
		v := *o.Mode

		if err = d.Set("mode", v); err != nil {
			return diag.Errorf("error reading mode: %v", err)
		}
	}

	if o.Monitor != nil {
		v := *o.Monitor

		if err = d.Set("monitor", v); err != nil {
			return diag.Errorf("error reading monitor: %v", err)
		}
	}

	if o.MulticastTtl != nil {
		v := *o.MulticastTtl

		if err = d.Set("multicast_ttl", v); err != nil {
			return diag.Errorf("error reading multicast_ttl: %v", err)
		}
	}

	if o.NntpProxyThreshold != nil {
		v := *o.NntpProxyThreshold

		if err = d.Set("nntp_proxy_threshold", v); err != nil {
			return diag.Errorf("error reading nntp_proxy_threshold: %v", err)
		}
	}

	if o.Override != nil {
		v := *o.Override

		if err = d.Set("override", v); err != nil {
			return diag.Errorf("error reading override: %v", err)
		}
	}

	if o.OverrideWaitTime != nil {
		v := *o.OverrideWaitTime

		if err = d.Set("override_wait_time", v); err != nil {
			return diag.Errorf("error reading override_wait_time: %v", err)
		}
	}

	if o.Password != nil {
		v := *o.Password

		if err = d.Set("password", v); err != nil {
			return diag.Errorf("error reading password: %v", err)
		}
	}

	if o.PingserverFailoverThreshold != nil {
		v := *o.PingserverFailoverThreshold

		if err = d.Set("pingserver_failover_threshold", v); err != nil {
			return diag.Errorf("error reading pingserver_failover_threshold: %v", err)
		}
	}

	if o.PingserverFlipTimeout != nil {
		v := *o.PingserverFlipTimeout

		if err = d.Set("pingserver_flip_timeout", v); err != nil {
			return diag.Errorf("error reading pingserver_flip_timeout: %v", err)
		}
	}

	if o.PingserverMonitorInterface != nil {
		v := *o.PingserverMonitorInterface

		if err = d.Set("pingserver_monitor_interface", v); err != nil {
			return diag.Errorf("error reading pingserver_monitor_interface: %v", err)
		}
	}

	if o.PingserverSecondaryForceReset != nil {
		v := *o.PingserverSecondaryForceReset

		if err = d.Set("pingserver_secondary_force_reset", v); err != nil {
			return diag.Errorf("error reading pingserver_secondary_force_reset: %v", err)
		}
	}

	if o.PingserverSlaveForceReset != nil {
		v := *o.PingserverSlaveForceReset

		if err = d.Set("pingserver_slave_force_reset", v); err != nil {
			return diag.Errorf("error reading pingserver_slave_force_reset: %v", err)
		}
	}

	if o.Pop3ProxyThreshold != nil {
		v := *o.Pop3ProxyThreshold

		if err = d.Set("pop3_proxy_threshold", v); err != nil {
			return diag.Errorf("error reading pop3_proxy_threshold: %v", err)
		}
	}

	if o.Priority != nil {
		v := *o.Priority

		if err = d.Set("priority", v); err != nil {
			return diag.Errorf("error reading priority: %v", err)
		}
	}

	if o.RouteHold != nil {
		v := *o.RouteHold

		if err = d.Set("route_hold", v); err != nil {
			return diag.Errorf("error reading route_hold: %v", err)
		}
	}

	if o.RouteTtl != nil {
		v := *o.RouteTtl

		if err = d.Set("route_ttl", v); err != nil {
			return diag.Errorf("error reading route_ttl: %v", err)
		}
	}

	if o.RouteWait != nil {
		v := *o.RouteWait

		if err = d.Set("route_wait", v); err != nil {
			return diag.Errorf("error reading route_wait: %v", err)
		}
	}

	if o.Schedule != nil {
		v := *o.Schedule

		if err = d.Set("schedule", v); err != nil {
			return diag.Errorf("error reading schedule: %v", err)
		}
	}

	if o.SecondaryVcluster != nil {
		if err = d.Set("secondary_vcluster", flattenSystemHaSecondaryVcluster(o.SecondaryVcluster, sort)); err != nil {
			return diag.Errorf("error reading secondary_vcluster: %v", err)
		}
	}

	if o.SessionPickup != nil {
		v := *o.SessionPickup

		if err = d.Set("session_pickup", v); err != nil {
			return diag.Errorf("error reading session_pickup: %v", err)
		}
	}

	if o.SessionPickupConnectionless != nil {
		v := *o.SessionPickupConnectionless

		if err = d.Set("session_pickup_connectionless", v); err != nil {
			return diag.Errorf("error reading session_pickup_connectionless: %v", err)
		}
	}

	if o.SessionPickupDelay != nil {
		v := *o.SessionPickupDelay

		if err = d.Set("session_pickup_delay", v); err != nil {
			return diag.Errorf("error reading session_pickup_delay: %v", err)
		}
	}

	if o.SessionPickupExpectation != nil {
		v := *o.SessionPickupExpectation

		if err = d.Set("session_pickup_expectation", v); err != nil {
			return diag.Errorf("error reading session_pickup_expectation: %v", err)
		}
	}

	if o.SessionPickupNat != nil {
		v := *o.SessionPickupNat

		if err = d.Set("session_pickup_nat", v); err != nil {
			return diag.Errorf("error reading session_pickup_nat: %v", err)
		}
	}

	if o.SessionSyncDev != nil {
		v := *o.SessionSyncDev

		if err = d.Set("session_sync_dev", v); err != nil {
			return diag.Errorf("error reading session_sync_dev: %v", err)
		}
	}

	if o.SmtpProxyThreshold != nil {
		v := *o.SmtpProxyThreshold

		if err = d.Set("smtp_proxy_threshold", v); err != nil {
			return diag.Errorf("error reading smtp_proxy_threshold: %v", err)
		}
	}

	if o.SsdFailover != nil {
		v := *o.SsdFailover

		if err = d.Set("ssd_failover", v); err != nil {
			return diag.Errorf("error reading ssd_failover: %v", err)
		}
	}

	if o.StandaloneConfigSync != nil {
		v := *o.StandaloneConfigSync

		if err = d.Set("standalone_config_sync", v); err != nil {
			return diag.Errorf("error reading standalone_config_sync: %v", err)
		}
	}

	if o.StandaloneMgmtVdom != nil {
		v := *o.StandaloneMgmtVdom

		if err = d.Set("standalone_mgmt_vdom", v); err != nil {
			return diag.Errorf("error reading standalone_mgmt_vdom: %v", err)
		}
	}

	if o.SyncConfig != nil {
		v := *o.SyncConfig

		if err = d.Set("sync_config", v); err != nil {
			return diag.Errorf("error reading sync_config: %v", err)
		}
	}

	if o.SyncPacketBalance != nil {
		v := *o.SyncPacketBalance

		if err = d.Set("sync_packet_balance", v); err != nil {
			return diag.Errorf("error reading sync_packet_balance: %v", err)
		}
	}

	if o.UnicastGateway != nil {
		v := *o.UnicastGateway

		if err = d.Set("unicast_gateway", v); err != nil {
			return diag.Errorf("error reading unicast_gateway: %v", err)
		}
	}

	if o.UnicastHb != nil {
		v := *o.UnicastHb

		if err = d.Set("unicast_hb", v); err != nil {
			return diag.Errorf("error reading unicast_hb: %v", err)
		}
	}

	if o.UnicastHbNetmask != nil {
		v := *o.UnicastHbNetmask

		if err = d.Set("unicast_hb_netmask", v); err != nil {
			return diag.Errorf("error reading unicast_hb_netmask: %v", err)
		}
	}

	if o.UnicastHbPeerip != nil {
		v := *o.UnicastHbPeerip

		if err = d.Set("unicast_hb_peerip", v); err != nil {
			return diag.Errorf("error reading unicast_hb_peerip: %v", err)
		}
	}

	if o.UnicastPeers != nil {
		if err = d.Set("unicast_peers", flattenSystemHaUnicastPeers(o.UnicastPeers, sort)); err != nil {
			return diag.Errorf("error reading unicast_peers: %v", err)
		}
	}

	if o.UnicastStatus != nil {
		v := *o.UnicastStatus

		if err = d.Set("unicast_status", v); err != nil {
			return diag.Errorf("error reading unicast_status: %v", err)
		}
	}

	if o.UninterruptiblePrimaryWait != nil {
		v := *o.UninterruptiblePrimaryWait

		if err = d.Set("uninterruptible_primary_wait", v); err != nil {
			return diag.Errorf("error reading uninterruptible_primary_wait: %v", err)
		}
	}

	if o.UninterruptibleUpgrade != nil {
		v := *o.UninterruptibleUpgrade

		if err = d.Set("uninterruptible_upgrade", v); err != nil {
			return diag.Errorf("error reading uninterruptible_upgrade: %v", err)
		}
	}

	if o.VclusterId != nil {
		v := *o.VclusterId

		if err = d.Set("vcluster_id", v); err != nil {
			return diag.Errorf("error reading vcluster_id: %v", err)
		}
	}

	if o.Vcluster2 != nil {
		v := *o.Vcluster2

		if err = d.Set("vcluster2", v); err != nil {
			return diag.Errorf("error reading vcluster2: %v", err)
		}
	}

	if o.Vdom != nil {
		v := *o.Vdom

		if err = d.Set("vdom", v); err != nil {
			return diag.Errorf("error reading vdom: %v", err)
		}
	}

	if o.Weight != nil {
		v := *o.Weight

		if err = d.Set("weight", v); err != nil {
			return diag.Errorf("error reading weight: %v", err)
		}
	}

	return nil
}

func expandSystemHaHaMgmtInterfaces(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemHaHaMgmtInterfaces, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemHaHaMgmtInterfaces

	for i := range l {
		tmp := models.SystemHaHaMgmtInterfaces{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dst", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dst = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.gateway", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Gateway = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.gateway6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Gateway6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemHaSecondaryVcluster(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemHaSecondaryVcluster, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemHaSecondaryVcluster

	for i := range l {
		tmp := models.SystemHaSecondaryVcluster{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.monitor", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Monitor = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Override = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_wait_time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.OverrideWaitTime = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pingserver_failover_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.PingserverFailoverThreshold = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pingserver_monitor_interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PingserverMonitorInterface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pingserver_secondary_force_reset", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PingserverSecondaryForceReset = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pingserver_slave_force_reset", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PingserverSlaveForceReset = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Priority = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vcluster_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.VclusterId = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vdom", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Vdom = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemHaUnicastPeers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemHaUnicastPeers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemHaUnicastPeers

	for i := range l {
		tmp := models.SystemHaUnicastPeers{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.peer_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PeerIp = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemHa(d *schema.ResourceData, sv string) (*models.SystemHa, diag.Diagnostics) {
	obj := models.SystemHa{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("arps"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("arps", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Arps = &tmp
		}
	}
	if v1, ok := d.GetOk("arps_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("arps_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ArpsInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("authentication"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authentication", sv)
				diags = append(diags, e)
			}
			obj.Authentication = &v2
		}
	}
	if v1, ok := d.GetOk("cpu_threshold"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cpu_threshold", sv)
				diags = append(diags, e)
			}
			obj.CpuThreshold = &v2
		}
	}
	if v1, ok := d.GetOk("encryption"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("encryption", sv)
				diags = append(diags, e)
			}
			obj.Encryption = &v2
		}
	}
	if v1, ok := d.GetOk("failover_hold_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("failover_hold_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FailoverHoldTime = &tmp
		}
	}
	if v1, ok := d.GetOk("ftp_proxy_threshold"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ftp_proxy_threshold", sv)
				diags = append(diags, e)
			}
			obj.FtpProxyThreshold = &v2
		}
	}
	if v1, ok := d.GetOk("gratuitous_arps"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gratuitous_arps", sv)
				diags = append(diags, e)
			}
			obj.GratuitousArps = &v2
		}
	}
	if v1, ok := d.GetOk("group_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("group_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.GroupId = &tmp
		}
	}
	if v1, ok := d.GetOk("group_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("group_name", sv)
				diags = append(diags, e)
			}
			obj.GroupName = &v2
		}
	}
	if v1, ok := d.GetOk("ha_direct"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ha_direct", sv)
				diags = append(diags, e)
			}
			obj.HaDirect = &v2
		}
	}
	if v1, ok := d.GetOk("ha_eth_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ha_eth_type", sv)
				diags = append(diags, e)
			}
			obj.HaEthType = &v2
		}
	}
	if v, ok := d.GetOk("ha_mgmt_interfaces"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ha_mgmt_interfaces", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemHaHaMgmtInterfaces(d, v, "ha_mgmt_interfaces", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.HaMgmtInterfaces = t
		}
	} else if d.HasChange("ha_mgmt_interfaces") {
		old, new := d.GetChange("ha_mgmt_interfaces")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.HaMgmtInterfaces = &[]models.SystemHaHaMgmtInterfaces{}
		}
	}
	if v1, ok := d.GetOk("ha_mgmt_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ha_mgmt_status", sv)
				diags = append(diags, e)
			}
			obj.HaMgmtStatus = &v2
		}
	}
	if v1, ok := d.GetOk("ha_uptime_diff_margin"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ha_uptime_diff_margin", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HaUptimeDiffMargin = &tmp
		}
	}
	if v1, ok := d.GetOk("hb_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hb_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HbInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("hb_interval_in_milliseconds"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("hb_interval_in_milliseconds", sv)
				diags = append(diags, e)
			}
			obj.HbIntervalInMilliseconds = &v2
		}
	}
	if v1, ok := d.GetOk("hb_lost_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hb_lost_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HbLostThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("hbdev"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hbdev", sv)
				diags = append(diags, e)
			}
			obj.Hbdev = &v2
		}
	}
	if v1, ok := d.GetOk("hc_eth_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hc_eth_type", sv)
				diags = append(diags, e)
			}
			obj.HcEthType = &v2
		}
	}
	if v1, ok := d.GetOk("hello_holddown"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hello_holddown", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HelloHolddown = &tmp
		}
	}
	if v1, ok := d.GetOk("http_proxy_threshold"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_proxy_threshold", sv)
				diags = append(diags, e)
			}
			obj.HttpProxyThreshold = &v2
		}
	}
	if v1, ok := d.GetOk("imap_proxy_threshold"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("imap_proxy_threshold", sv)
				diags = append(diags, e)
			}
			obj.ImapProxyThreshold = &v2
		}
	}
	if v1, ok := d.GetOk("inter_cluster_session_sync"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("inter_cluster_session_sync", sv)
				diags = append(diags, e)
			}
			obj.InterClusterSessionSync = &v2
		}
	}
	if v1, ok := d.GetOk("key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("key", sv)
				diags = append(diags, e)
			}
			obj.Key = &v2
		}
	}
	if v1, ok := d.GetOk("l2ep_eth_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("l2ep_eth_type", sv)
				diags = append(diags, e)
			}
			obj.L2epEthType = &v2
		}
	}
	if v1, ok := d.GetOk("link_failed_signal"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("link_failed_signal", sv)
				diags = append(diags, e)
			}
			obj.LinkFailedSignal = &v2
		}
	}
	if v1, ok := d.GetOk("load_balance_all"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("load_balance_all", sv)
				diags = append(diags, e)
			}
			obj.LoadBalanceAll = &v2
		}
	}
	if v1, ok := d.GetOk("logical_sn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("logical_sn", sv)
				diags = append(diags, e)
			}
			obj.LogicalSn = &v2
		}
	}
	if v1, ok := d.GetOk("memory_based_failover"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("memory_based_failover", sv)
				diags = append(diags, e)
			}
			obj.MemoryBasedFailover = &v2
		}
	}
	if v1, ok := d.GetOk("memory_compatible_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("memory_compatible_mode", sv)
				diags = append(diags, e)
			}
			obj.MemoryCompatibleMode = &v2
		}
	}
	if v1, ok := d.GetOk("memory_failover_flip_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("memory_failover_flip_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MemoryFailoverFlipTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("memory_failover_monitor_period"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("memory_failover_monitor_period", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MemoryFailoverMonitorPeriod = &tmp
		}
	}
	if v1, ok := d.GetOk("memory_failover_sample_rate"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("memory_failover_sample_rate", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MemoryFailoverSampleRate = &tmp
		}
	}
	if v1, ok := d.GetOk("memory_failover_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("memory_failover_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MemoryFailoverThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("memory_threshold"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("memory_threshold", sv)
				diags = append(diags, e)
			}
			obj.MemoryThreshold = &v2
		}
	}
	if v1, ok := d.GetOk("mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mode", sv)
				diags = append(diags, e)
			}
			obj.Mode = &v2
		}
	}
	if v1, ok := d.GetOk("monitor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("monitor", sv)
				diags = append(diags, e)
			}
			obj.Monitor = &v2
		}
	}
	if v1, ok := d.GetOk("multicast_ttl"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("multicast_ttl", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MulticastTtl = &tmp
		}
	}
	if v1, ok := d.GetOk("nntp_proxy_threshold"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("nntp_proxy_threshold", sv)
				diags = append(diags, e)
			}
			obj.NntpProxyThreshold = &v2
		}
	}
	if v1, ok := d.GetOk("override"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("override", sv)
				diags = append(diags, e)
			}
			obj.Override = &v2
		}
	}
	if v1, ok := d.GetOk("override_wait_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("override_wait_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.OverrideWaitTime = &tmp
		}
	}
	if v1, ok := d.GetOk("password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("password", sv)
				diags = append(diags, e)
			}
			obj.Password = &v2
		}
	}
	if v1, ok := d.GetOk("pingserver_failover_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pingserver_failover_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PingserverFailoverThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("pingserver_flip_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pingserver_flip_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PingserverFlipTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("pingserver_monitor_interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pingserver_monitor_interface", sv)
				diags = append(diags, e)
			}
			obj.PingserverMonitorInterface = &v2
		}
	}
	if v1, ok := d.GetOk("pingserver_secondary_force_reset"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("pingserver_secondary_force_reset", sv)
				diags = append(diags, e)
			}
			obj.PingserverSecondaryForceReset = &v2
		}
	}
	if v1, ok := d.GetOk("pingserver_slave_force_reset"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.3") {
				e := utils.AttributeVersionWarning("pingserver_slave_force_reset", sv)
				diags = append(diags, e)
			}
			obj.PingserverSlaveForceReset = &v2
		}
	}
	if v1, ok := d.GetOk("pop3_proxy_threshold"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pop3_proxy_threshold", sv)
				diags = append(diags, e)
			}
			obj.Pop3ProxyThreshold = &v2
		}
	}
	if v1, ok := d.GetOk("priority"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("priority", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Priority = &tmp
		}
	}
	if v1, ok := d.GetOk("route_hold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("route_hold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RouteHold = &tmp
		}
	}
	if v1, ok := d.GetOk("route_ttl"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("route_ttl", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RouteTtl = &tmp
		}
	}
	if v1, ok := d.GetOk("route_wait"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("route_wait", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RouteWait = &tmp
		}
	}
	if v1, ok := d.GetOk("schedule"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("schedule", sv)
				diags = append(diags, e)
			}
			obj.Schedule = &v2
		}
	}
	if v, ok := d.GetOk("secondary_vcluster"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("secondary_vcluster", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemHaSecondaryVcluster(d, v, "secondary_vcluster", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SecondaryVcluster = t
		}
	} else if d.HasChange("secondary_vcluster") {
		old, new := d.GetChange("secondary_vcluster")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SecondaryVcluster = &[]models.SystemHaSecondaryVcluster{}
		}
	}
	if v1, ok := d.GetOk("session_pickup"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("session_pickup", sv)
				diags = append(diags, e)
			}
			obj.SessionPickup = &v2
		}
	}
	if v1, ok := d.GetOk("session_pickup_connectionless"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("session_pickup_connectionless", sv)
				diags = append(diags, e)
			}
			obj.SessionPickupConnectionless = &v2
		}
	}
	if v1, ok := d.GetOk("session_pickup_delay"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("session_pickup_delay", sv)
				diags = append(diags, e)
			}
			obj.SessionPickupDelay = &v2
		}
	}
	if v1, ok := d.GetOk("session_pickup_expectation"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("session_pickup_expectation", sv)
				diags = append(diags, e)
			}
			obj.SessionPickupExpectation = &v2
		}
	}
	if v1, ok := d.GetOk("session_pickup_nat"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("session_pickup_nat", sv)
				diags = append(diags, e)
			}
			obj.SessionPickupNat = &v2
		}
	}
	if v1, ok := d.GetOk("session_sync_dev"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("session_sync_dev", sv)
				diags = append(diags, e)
			}
			obj.SessionSyncDev = &v2
		}
	}
	if v1, ok := d.GetOk("smtp_proxy_threshold"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("smtp_proxy_threshold", sv)
				diags = append(diags, e)
			}
			obj.SmtpProxyThreshold = &v2
		}
	}
	if v1, ok := d.GetOk("ssd_failover"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssd_failover", sv)
				diags = append(diags, e)
			}
			obj.SsdFailover = &v2
		}
	}
	if v1, ok := d.GetOk("standalone_config_sync"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("standalone_config_sync", sv)
				diags = append(diags, e)
			}
			obj.StandaloneConfigSync = &v2
		}
	}
	if v1, ok := d.GetOk("standalone_mgmt_vdom"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("standalone_mgmt_vdom", sv)
				diags = append(diags, e)
			}
			obj.StandaloneMgmtVdom = &v2
		}
	}
	if v1, ok := d.GetOk("sync_config"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sync_config", sv)
				diags = append(diags, e)
			}
			obj.SyncConfig = &v2
		}
	}
	if v1, ok := d.GetOk("sync_packet_balance"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sync_packet_balance", sv)
				diags = append(diags, e)
			}
			obj.SyncPacketBalance = &v2
		}
	}
	if v1, ok := d.GetOk("unicast_gateway"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("unicast_gateway", sv)
				diags = append(diags, e)
			}
			obj.UnicastGateway = &v2
		}
	}
	if v1, ok := d.GetOk("unicast_hb"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("unicast_hb", sv)
				diags = append(diags, e)
			}
			obj.UnicastHb = &v2
		}
	}
	if v1, ok := d.GetOk("unicast_hb_netmask"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("unicast_hb_netmask", sv)
				diags = append(diags, e)
			}
			obj.UnicastHbNetmask = &v2
		}
	}
	if v1, ok := d.GetOk("unicast_hb_peerip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("unicast_hb_peerip", sv)
				diags = append(diags, e)
			}
			obj.UnicastHbPeerip = &v2
		}
	}
	if v, ok := d.GetOk("unicast_peers"); ok {
		if !utils.CheckVer(sv, "v7.0.0", "") {
			e := utils.AttributeVersionWarning("unicast_peers", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemHaUnicastPeers(d, v, "unicast_peers", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.UnicastPeers = t
		}
	} else if d.HasChange("unicast_peers") {
		old, new := d.GetChange("unicast_peers")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.UnicastPeers = &[]models.SystemHaUnicastPeers{}
		}
	}
	if v1, ok := d.GetOk("unicast_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("unicast_status", sv)
				diags = append(diags, e)
			}
			obj.UnicastStatus = &v2
		}
	}
	if v1, ok := d.GetOk("uninterruptible_primary_wait"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("uninterruptible_primary_wait", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.UninterruptiblePrimaryWait = &tmp
		}
	}
	if v1, ok := d.GetOk("uninterruptible_upgrade"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("uninterruptible_upgrade", sv)
				diags = append(diags, e)
			}
			obj.UninterruptibleUpgrade = &v2
		}
	}
	if v1, ok := d.GetOk("vcluster_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vcluster_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.VclusterId = &tmp
		}
	}
	if v1, ok := d.GetOk("vcluster2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vcluster2", sv)
				diags = append(diags, e)
			}
			obj.Vcluster2 = &v2
		}
	}
	if v1, ok := d.GetOk("vdom"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vdom", sv)
				diags = append(diags, e)
			}
			obj.Vdom = &v2
		}
	}
	if v1, ok := d.GetOk("weight"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("weight", sv)
				diags = append(diags, e)
			}
			obj.Weight = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemHa(d *schema.ResourceData, sv string) (*models.SystemHa, diag.Diagnostics) {
	obj := models.SystemHa{}
	diags := diag.Diagnostics{}

	obj.HaMgmtInterfaces = &[]models.SystemHaHaMgmtInterfaces{}
	obj.SecondaryVcluster = &[]models.SystemHaSecondaryVcluster{}
	obj.UnicastPeers = &[]models.SystemHaUnicastPeers{}

	return &obj, diags
}
