// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure HA.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemHa() *schema.Resource {
	return &schema.Resource{
		Description: "Configure HA.",

		ReadContext: dataSourceSystemHaRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"arps": {
				Type:        schema.TypeInt,
				Description: "Number of gratuitous ARPs (1 - 60). Lower to reduce traffic. Higher to reduce failover time.",
				Computed:    true,
			},
			"arps_interval": {
				Type:        schema.TypeInt,
				Description: "Time between gratuitous ARPs  (1 - 20 sec). Lower to reduce failover time. Higher to reduce traffic.",
				Computed:    true,
			},
			"authentication": {
				Type:        schema.TypeString,
				Description: "Enable/disable heartbeat message authentication.",
				Computed:    true,
			},
			"cpu_threshold": {
				Type:        schema.TypeString,
				Description: "Dynamic weighted load balancing CPU usage weight and high and low thresholds.",
				Computed:    true,
			},
			"encryption": {
				Type:        schema.TypeString,
				Description: "Enable/disable heartbeat message encryption.",
				Computed:    true,
			},
			"failover_hold_time": {
				Type:        schema.TypeInt,
				Description: "Time to wait before failover (0 - 300 sec, default = 0), to avoid flip.",
				Computed:    true,
			},
			"ftp_proxy_threshold": {
				Type:        schema.TypeString,
				Description: "Dynamic weighted load balancing weight and high and low number of FTP proxy sessions.",
				Computed:    true,
			},
			"gratuitous_arps": {
				Type:        schema.TypeString,
				Description: "Enable/disable gratuitous ARPs. Disable if link-failed-signal enabled.",
				Computed:    true,
			},
			"group_id": {
				Type:        schema.TypeInt,
				Description: "HA group ID  (0 - 1023;  or 0 - 7 when vcluster is enabled). Must be the same for all members.",
				Computed:    true,
			},
			"group_name": {
				Type:        schema.TypeString,
				Description: "Cluster group name. Must be the same for all members.",
				Computed:    true,
			},
			"ha_direct": {
				Type:        schema.TypeString,
				Description: "Enable/disable using ha-mgmt interface for syslog, SNMP, remote authentication (RADIUS), FortiAnalyzer, FortiSandbox, sFlow, and Netflow.",
				Computed:    true,
			},
			"ha_eth_type": {
				Type:        schema.TypeString,
				Description: "HA heartbeat packet Ethertype (4-digit hex).",
				Computed:    true,
			},
			"ha_mgmt_interfaces": {
				Type:        schema.TypeList,
				Description: "Reserve interfaces to manage individual cluster units.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst": {
							Type:        schema.TypeString,
							Description: "Default route destination for reserved HA management interface.",
							Computed:    true,
						},
						"gateway": {
							Type:        schema.TypeString,
							Description: "Default route gateway for reserved HA management interface.",
							Computed:    true,
						},
						"gateway6": {
							Type:        schema.TypeString,
							Description: "Default IPv6 gateway for reserved HA management interface.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Table ID.",
							Computed:    true,
						},
						"interface": {
							Type:        schema.TypeString,
							Description: "Interface to reserve for HA management.",
							Computed:    true,
						},
					},
				},
			},
			"ha_mgmt_status": {
				Type:        schema.TypeString,
				Description: "Enable to reserve interfaces to manage individual cluster units.",
				Computed:    true,
			},
			"ha_uptime_diff_margin": {
				Type:        schema.TypeInt,
				Description: "Normally you would only reduce this value for failover testing.",
				Computed:    true,
			},
			"hb_interval": {
				Type:        schema.TypeInt,
				Description: "Time between sending heartbeat packets (1 - 20). Increase to reduce false positives.",
				Computed:    true,
			},
			"hb_interval_in_milliseconds": {
				Type:        schema.TypeString,
				Description: "Number of milliseconds for each heartbeat interval: 100ms or 10ms.",
				Computed:    true,
			},
			"hb_lost_threshold": {
				Type:        schema.TypeInt,
				Description: "Number of lost heartbeats to signal a failure (1 - 60). Increase to reduce false positives.",
				Computed:    true,
			},
			"hbdev": {
				Type:        schema.TypeString,
				Description: "Heartbeat interfaces. Must be the same for all members.",
				Computed:    true,
			},
			"hc_eth_type": {
				Type:        schema.TypeString,
				Description: "Transparent mode HA heartbeat packet Ethertype (4-digit hex).",
				Computed:    true,
			},
			"hello_holddown": {
				Type:        schema.TypeInt,
				Description: "Time to wait before changing from hello to work state (5 - 300 sec).",
				Computed:    true,
			},
			"http_proxy_threshold": {
				Type:        schema.TypeString,
				Description: "Dynamic weighted load balancing weight and high and low number of HTTP proxy sessions.",
				Computed:    true,
			},
			"imap_proxy_threshold": {
				Type:        schema.TypeString,
				Description: "Dynamic weighted load balancing weight and high and low number of IMAP proxy sessions.",
				Computed:    true,
			},
			"inter_cluster_session_sync": {
				Type:        schema.TypeString,
				Description: "Enable/disable synchronization of sessions among HA clusters.",
				Computed:    true,
			},
			"key": {
				Type:        schema.TypeString,
				Description: "Key.",
				Computed:    true,
				Sensitive:   true,
			},
			"l2ep_eth_type": {
				Type:        schema.TypeString,
				Description: "Telnet session HA heartbeat packet Ethertype (4-digit hex).",
				Computed:    true,
			},
			"link_failed_signal": {
				Type:        schema.TypeString,
				Description: "Enable to shut down all interfaces for 1 sec after a failover. Use if gratuitous ARPs do not update network.",
				Computed:    true,
			},
			"load_balance_all": {
				Type:        schema.TypeString,
				Description: "Enable to load balance TCP sessions. Disable to load balance proxy sessions only.",
				Computed:    true,
			},
			"logical_sn": {
				Type:        schema.TypeString,
				Description: "Enable/disable usage of the logical serial number.",
				Computed:    true,
			},
			"memory_based_failover": {
				Type:        schema.TypeString,
				Description: "Enable/disable memory based failover.",
				Computed:    true,
			},
			"memory_compatible_mode": {
				Type:        schema.TypeString,
				Description: "Enable/disable memory compatible mode.",
				Computed:    true,
			},
			"memory_failover_flip_timeout": {
				Type:        schema.TypeInt,
				Description: "Time to wait between subsequent memory based failovers in minutes (6 - 2147483647, default = 6).",
				Computed:    true,
			},
			"memory_failover_monitor_period": {
				Type:        schema.TypeInt,
				Description: "Duration of high memory usage before memory based failover is triggered in seconds (1 - 300, default = 60).",
				Computed:    true,
			},
			"memory_failover_sample_rate": {
				Type:        schema.TypeInt,
				Description: "Rate at which memory usage is sampled in order to measure memory usage in seconds (1 - 60, default = 1).",
				Computed:    true,
			},
			"memory_failover_threshold": {
				Type:        schema.TypeInt,
				Description: "Memory usage threshold to trigger memory based failover (0 means using conserve mode threshold in system.global).",
				Computed:    true,
			},
			"memory_threshold": {
				Type:        schema.TypeString,
				Description: "Dynamic weighted load balancing memory usage weight and high and low thresholds.",
				Computed:    true,
			},
			"mode": {
				Type:        schema.TypeString,
				Description: "HA mode. Must be the same for all members. FGSP requires standalone.",
				Computed:    true,
			},
			"monitor": {
				Type:        schema.TypeString,
				Description: "Interfaces to check for port monitoring (or link failure).",
				Computed:    true,
			},
			"multicast_ttl": {
				Type:        schema.TypeInt,
				Description: "HA multicast TTL on primary (5 - 3600 sec).",
				Computed:    true,
			},
			"nntp_proxy_threshold": {
				Type:        schema.TypeString,
				Description: "Dynamic weighted load balancing weight and high and low number of NNTP proxy sessions.",
				Computed:    true,
			},
			"override": {
				Type:        schema.TypeString,
				Description: "Enable and increase the priority of the unit that should always be primary (master).",
				Computed:    true,
			},
			"override_wait_time": {
				Type:        schema.TypeInt,
				Description: "Delay negotiating if override is enabled (0 - 3600 sec). Reduces how often the cluster negotiates.",
				Computed:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "Cluster password. Must be the same for all members.",
				Computed:    true,
				Sensitive:   true,
			},
			"pingserver_failover_threshold": {
				Type:        schema.TypeInt,
				Description: "Remote IP monitoring failover threshold (0 - 50).",
				Computed:    true,
			},
			"pingserver_flip_timeout": {
				Type:        schema.TypeInt,
				Description: "Time to wait in minutes before renegotiating after a remote IP monitoring failover.",
				Computed:    true,
			},
			"pingserver_monitor_interface": {
				Type:        schema.TypeString,
				Description: "Interfaces to check for remote IP monitoring.",
				Computed:    true,
			},
			"pingserver_secondary_force_reset": {
				Type:        schema.TypeString,
				Description: "Enable to force the cluster to negotiate after a remote IP monitoring failover.",
				Computed:    true,
			},
			"pingserver_slave_force_reset": {
				Type:        schema.TypeString,
				Description: "Enable to force the cluster to negotiate after a remote IP monitoring failover.",
				Computed:    true,
			},
			"pop3_proxy_threshold": {
				Type:        schema.TypeString,
				Description: "Dynamic weighted load balancing weight and high and low number of POP3 proxy sessions.",
				Computed:    true,
			},
			"priority": {
				Type:        schema.TypeInt,
				Description: "Increase the priority to select the primary unit (0 - 255).",
				Computed:    true,
			},
			"route_hold": {
				Type:        schema.TypeInt,
				Description: "Time to wait between routing table updates to the cluster (0 - 3600 sec).",
				Computed:    true,
			},
			"route_ttl": {
				Type:        schema.TypeInt,
				Description: "TTL for primary unit routes (5 - 3600 sec). Increase to maintain active routes during failover.",
				Computed:    true,
			},
			"route_wait": {
				Type:        schema.TypeInt,
				Description: "Time to wait before sending new routes to the cluster (0 - 3600 sec).",
				Computed:    true,
			},
			"schedule": {
				Type:        schema.TypeString,
				Description: "Type of A-A load balancing. Use none if you have external load balancers.",
				Computed:    true,
			},
			"secondary_vcluster": {
				Type:        schema.TypeList,
				Description: "Configure virtual cluster 2.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitor": {
							Type:        schema.TypeString,
							Description: "Interfaces to check for port monitoring (or link failure).",
							Computed:    true,
						},
						"override": {
							Type:        schema.TypeString,
							Description: "Enable and increase the priority of the unit that should always be primary.",
							Computed:    true,
						},
						"override_wait_time": {
							Type:        schema.TypeInt,
							Description: "Delay negotiating if override is enabled (0 - 3600 sec). Reduces how often the cluster negotiates.",
							Computed:    true,
						},
						"pingserver_failover_threshold": {
							Type:        schema.TypeInt,
							Description: "Remote IP monitoring failover threshold (0 - 50).",
							Computed:    true,
						},
						"pingserver_monitor_interface": {
							Type:        schema.TypeString,
							Description: "Interfaces to check for remote IP monitoring.",
							Computed:    true,
						},
						"pingserver_secondary_force_reset": {
							Type:        schema.TypeString,
							Description: "Enable to force the cluster to negotiate after a remote IP monitoring failover.",
							Computed:    true,
						},
						"pingserver_slave_force_reset": {
							Type:        schema.TypeString,
							Description: "Enable to force the cluster to negotiate after a remote IP monitoring failover.",
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "Increase the priority to select the primary unit (0 - 255).",
							Computed:    true,
						},
						"vcluster_id": {
							Type:        schema.TypeInt,
							Description: "Cluster ID.",
							Computed:    true,
						},
						"vdom": {
							Type:        schema.TypeString,
							Description: "VDOMs in virtual cluster 2.",
							Computed:    true,
						},
					},
				},
			},
			"session_pickup": {
				Type:        schema.TypeString,
				Description: "Enable/disable session pickup. Enabling it can reduce session down time when fail over happens.",
				Computed:    true,
			},
			"session_pickup_connectionless": {
				Type:        schema.TypeString,
				Description: "Enable/disable UDP and ICMP session sync.",
				Computed:    true,
			},
			"session_pickup_delay": {
				Type:        schema.TypeString,
				Description: "Enable to sync sessions longer than 30 sec. Only longer lived sessions need to be synced.",
				Computed:    true,
			},
			"session_pickup_expectation": {
				Type:        schema.TypeString,
				Description: "Enable/disable session helper expectation session sync for FGSP.",
				Computed:    true,
			},
			"session_pickup_nat": {
				Type:        schema.TypeString,
				Description: "Enable/disable NAT session sync for FGSP.",
				Computed:    true,
			},
			"session_sync_dev": {
				Type:        schema.TypeString,
				Description: "Offload session-sync process to kernel and sync sessions using connected interface(s) directly.",
				Computed:    true,
			},
			"smtp_proxy_threshold": {
				Type:        schema.TypeString,
				Description: "Dynamic weighted load balancing weight and high and low number of SMTP proxy sessions.",
				Computed:    true,
			},
			"ssd_failover": {
				Type:        schema.TypeString,
				Description: "Enable/disable automatic HA failover on SSD disk failure.",
				Computed:    true,
			},
			"standalone_config_sync": {
				Type:        schema.TypeString,
				Description: "Enable/disable FGSP configuration synchronization.",
				Computed:    true,
			},
			"standalone_mgmt_vdom": {
				Type:        schema.TypeString,
				Description: "Enable/disable standalone management VDOM.",
				Computed:    true,
			},
			"sync_config": {
				Type:        schema.TypeString,
				Description: "Enable/disable configuration synchronization.",
				Computed:    true,
			},
			"sync_packet_balance": {
				Type:        schema.TypeString,
				Description: "Enable/disable HA packet distribution to multiple CPUs.",
				Computed:    true,
			},
			"unicast_gateway": {
				Type:        schema.TypeString,
				Description: "Default route gateway for unicast interface.",
				Computed:    true,
			},
			"unicast_hb": {
				Type:        schema.TypeString,
				Description: "Enable/disable unicast heartbeat.",
				Computed:    true,
			},
			"unicast_hb_netmask": {
				Type:        schema.TypeString,
				Description: "Unicast heartbeat netmask.",
				Computed:    true,
			},
			"unicast_hb_peerip": {
				Type:        schema.TypeString,
				Description: "Unicast heartbeat peer IP.",
				Computed:    true,
			},
			"unicast_peers": {
				Type:        schema.TypeList,
				Description: "Number of unicast peers.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Table ID.",
							Computed:    true,
						},
						"peer_ip": {
							Type:        schema.TypeString,
							Description: "Unicast peer IP.",
							Computed:    true,
						},
					},
				},
			},
			"unicast_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable unicast connection.",
				Computed:    true,
			},
			"uninterruptible_primary_wait": {
				Type:        schema.TypeInt,
				Description: "Number of minutes the primary HA unit waits before the secondary HA unit is considered upgraded and the system is started before starting its own upgrade (15 - 300, default = 30).",
				Computed:    true,
			},
			"uninterruptible_upgrade": {
				Type:        schema.TypeString,
				Description: "Enable to upgrade a cluster without blocking network traffic.",
				Computed:    true,
			},
			"vcluster": {
				Type:        schema.TypeList,
				Description: "Virtual cluster table.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitor": {
							Type:        schema.TypeString,
							Description: "Interfaces to check for port monitoring (or link failure).",
							Computed:    true,
						},
						"override": {
							Type:        schema.TypeString,
							Description: "Enable and increase the priority of the unit that should always be primary (master).",
							Computed:    true,
						},
						"override_wait_time": {
							Type:        schema.TypeInt,
							Description: "Delay negotiating if override is enabled (0 - 3600 sec). Reduces how often the cluster negotiates.",
							Computed:    true,
						},
						"pingserver_failover_threshold": {
							Type:        schema.TypeInt,
							Description: "Remote IP monitoring failover threshold (0 - 50).",
							Computed:    true,
						},
						"pingserver_monitor_interface": {
							Type:        schema.TypeString,
							Description: "Interfaces to check for remote IP monitoring.",
							Computed:    true,
						},
						"pingserver_slave_force_reset": {
							Type:        schema.TypeString,
							Description: "Enable to force the cluster to negotiate after a remote IP monitoring failover.",
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "Increase the priority to select the primary unit (0 - 255).",
							Computed:    true,
						},
						"vcluster_id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"vdom": {
							Type:        schema.TypeList,
							Description: "Virtual domain(s) in the virtual cluster.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Virtual domain name.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"vcluster_id": {
				Type:        schema.TypeInt,
				Description: "Cluster ID.",
				Computed:    true,
			},
			"vcluster_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable virtual cluster for virtual clustering.",
				Computed:    true,
			},
			"vcluster2": {
				Type:        schema.TypeString,
				Description: "Enable/disable virtual cluster 2 for virtual clustering.",
				Computed:    true,
			},
			"vdom": {
				Type:        schema.TypeString,
				Description: "VDOMs in virtual cluster 1.",
				Computed:    true,
			},
			"weight": {
				Type:        schema.TypeString,
				Description: "Weight-round-robin weight for each cluster unit. Syntax <priority> <weight>.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemHaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemHa"

	o, err := c.Cmdb.ReadSystemHa(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemHa dataSource: %v", err)
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

	diags := refreshObjectSystemHa(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
