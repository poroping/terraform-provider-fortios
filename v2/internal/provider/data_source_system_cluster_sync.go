// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiGate Session Life Support Protocol (FGSP) session synchronization.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemClusterSync() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiGate Session Life Support Protocol (FGSP) session synchronization.",

		ReadContext: dataSourceSystemClusterSyncRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"down_intfs_before_sess_sync": {
				Type:        schema.TypeList,
				Description: "List of interfaces to be turned down before session synchronization is complete.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
					},
				},
			},
			"hb_interval": {
				Type:        schema.TypeInt,
				Description: "Heartbeat interval (1 - 10 sec).",
				Computed:    true,
			},
			"hb_lost_threshold": {
				Type:        schema.TypeInt,
				Description: "Lost heartbeat threshold (1 - 10).",
				Computed:    true,
			},
			"ike_heartbeat_interval": {
				Type:        schema.TypeInt,
				Description: "IKE heartbeat interval (1 - 60 secs).",
				Computed:    true,
			},
			"ike_monitor": {
				Type:        schema.TypeString,
				Description: "Enable/disable IKE HA monitor.",
				Computed:    true,
			},
			"ike_monitor_interval": {
				Type:        schema.TypeInt,
				Description: "IKE HA monitor interval (10 - 300 secs).",
				Computed:    true,
			},
			"ike_seqjump_speed": {
				Type:        schema.TypeInt,
				Description: "ESP jump ahead factor (1G - 10G pps equivalent).",
				Computed:    true,
			},
			"ipsec_tunnel_sync": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPsec tunnel synchronization.",
				Computed:    true,
			},
			"peerip": {
				Type:        schema.TypeString,
				Description: "IP address of the interface on the peer unit that is used for the session synchronization link.",
				Computed:    true,
			},
			"peervd": {
				Type:        schema.TypeString,
				Description: "VDOM that contains the session synchronization link interface on the peer unit. Usually both peers would have the same peervd.",
				Computed:    true,
			},
			"secondary_add_ipsec_routes": {
				Type:        schema.TypeString,
				Description: "Enable/disable IKE route announcement on the backup unit.",
				Computed:    true,
			},
			"session_sync_filter": {
				Type:        schema.TypeList,
				Description: "Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"custom_service": {
							Type:        schema.TypeList,
							Description: "Only sessions using these custom services are synchronized. Use source and destination port ranges to define these custom services.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dst_port_range": {
										Type:        schema.TypeString,
										Description: "Custom service destination port range.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "Custom service ID.",
										Computed:    true,
									},
									"src_port_range": {
										Type:        schema.TypeString,
										Description: "Custom service source port range.",
										Computed:    true,
									},
								},
							},
						},
						"dstaddr": {
							Type:        schema.TypeString,
							Description: "Only sessions to this IPv4 address are synchronized. You can only enter one address. To synchronize sessions for multiple destination addresses, add multiple filters.",
							Computed:    true,
						},
						"dstaddr6": {
							Type:        schema.TypeString,
							Description: "Only sessions to this IPv6 address are synchronized. You can only enter one address. To synchronize sessions for multiple destination addresses, add multiple filters.",
							Computed:    true,
						},
						"dstintf": {
							Type:        schema.TypeString,
							Description: "Only sessions to this interface are synchronized. You can only enter one interface name. To synchronize sessions to multiple destination interfaces, add multiple filters.",
							Computed:    true,
						},
						"srcaddr": {
							Type:        schema.TypeString,
							Description: "Only sessions from this IPv4 address are synchronized. You can only enter one address. To synchronize sessions from multiple source addresses, add multiple filters.",
							Computed:    true,
						},
						"srcaddr6": {
							Type:        schema.TypeString,
							Description: "Only sessions from this IPv6 address are synchronized. You can only enter one address. To synchronize sessions from multiple source addresses, add multiple filters.",
							Computed:    true,
						},
						"srcintf": {
							Type:        schema.TypeString,
							Description: "Only sessions from this interface are synchronized. You can only enter one interface name. To synchronize sessions for multiple source interfaces, add multiple filters.",
							Computed:    true,
						},
					},
				},
			},
			"slave_add_ike_routes": {
				Type:        schema.TypeString,
				Description: "Enable/disable IKE route announcement on the backup unit.",
				Computed:    true,
			},
			"sync_id": {
				Type:        schema.TypeInt,
				Description: "Sync ID.",
				Required:    true,
			},
			"syncvd": {
				Type:        schema.TypeList,
				Description: "Sessions from these VDOMs are synchronized using this session synchronization configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "VDOM name.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemClusterSyncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("sync_id")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadSystemClusterSync(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemClusterSync dataSource: %v", err)
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

	diags := refreshObjectSystemClusterSync(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
