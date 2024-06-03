// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes.

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

func resourceSystemStandaloneCluster() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes.",

		CreateContext: resourceSystemStandaloneClusterCreate,
		ReadContext:   resourceSystemStandaloneClusterRead,
		UpdateContext: resourceSystemStandaloneClusterUpdate,
		DeleteContext: resourceSystemStandaloneClusterDelete,

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
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"cluster_peer": {
				Type:        schema.TypeList,
				Description: "Configure FortiGate Session Life Support Protocol (FGSP) session synchronization.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"down_intfs_before_sess_sync": {
							Type:        schema.TypeList,
							Description: "List of interfaces to be turned down before session synchronization is complete.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Interface name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"hb_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 20),

							Description: "Heartbeat interval (1 - 20 (100*ms). Increase to reduce false positives.",
							Optional:    true,
							Computed:    true,
						},
						"hb_lost_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 60),

							Description: "Lost heartbeat threshold (1 - 60). Increase to reduce false positives.",
							Optional:    true,
							Computed:    true,
						},
						"ipsec_tunnel_sync": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable IPsec tunnel synchronization.",
							Optional:    true,
							Computed:    true,
						},
						"peerip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "IP address of the interface on the peer unit that is used for the session synchronization link.",
							Optional:    true,
							Computed:    true,
						},
						"peervd": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),

							Description: "VDOM that contains the session synchronization link interface on the peer unit. Usually both peers would have the same peervd.",
							Optional:    true,
							Computed:    true,
						},
						"secondary_add_ipsec_routes": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable IKE route announcement on the backup unit.",
							Optional:    true,
							Computed:    true,
						},
						"session_sync_filter": {
							Type:        schema.TypeList,
							Description: "Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize.",
							Optional:    true, MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"custom_service": {
										Type:        schema.TypeList,
										Description: "Only sessions using these custom services are synchronized. Use source and destination port ranges to define these custom services.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dst_port_range": {
													Type: schema.TypeString,

													Description: "Custom service destination port range.",
													Optional:    true,
													Computed:    true,
												},
												"id": {
													Type: schema.TypeInt,

													Description: "Custom service ID.",
													Optional:    true,
													Computed:    true,
												},
												"src_port_range": {
													Type: schema.TypeString,

													Description: "Custom service source port range.",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"dstaddr": {
										Type:         schema.TypeString,
										ValidateFunc: validators.FortiValidateIPv4ClassnetAny,

										Description: "Only sessions to this IPv4 address are synchronized.",
										Optional:    true,
										Computed:    true,
									},
									"dstaddr6": {
										Type:             schema.TypeString,
										ValidateFunc:     validators.FortiValidateIPv6Network,
										DiffSuppressFunc: suppressors.DiffCidrEqual,
										Description:      "Only sessions to this IPv6 address are synchronized.",
										Optional:         true,
										Computed:         true,
									},
									"dstintf": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 15),

										Description: "Only sessions to this interface are synchronized.",
										Optional:    true,
										Computed:    true,
									},
									"srcaddr": {
										Type:         schema.TypeString,
										ValidateFunc: validators.FortiValidateIPv4ClassnetAny,

										Description: "Only sessions from this IPv4 address are synchronized.",
										Optional:    true,
										Computed:    true,
									},
									"srcaddr6": {
										Type:             schema.TypeString,
										ValidateFunc:     validators.FortiValidateIPv6Network,
										DiffSuppressFunc: suppressors.DiffCidrEqual,
										Description:      "Only sessions from this IPv6 address are synchronized.",
										Optional:         true,
										Computed:         true,
									},
									"srcintf": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 15),

										Description: "Only sessions from this interface are synchronized.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"sync_id": {
							Type: schema.TypeInt,

							Description: "Sync ID.",
							Optional:    true,
							Computed:    true,
						},
						"syncvd": {
							Type:        schema.TypeList,
							Description: "Sessions from these VDOMs are synchronized using this session synchronization configuration.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "VDOM name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"encryption": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable encryption when synchronizing sessions.",
				Optional:    true,
				Computed:    true,
			},
			"group_member_id": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 15),

				Description: "Cluster member ID (0 - 15).",
				Optional:    true,
				Computed:    true,
			},
			"layer2_connection": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"available", "unavailable"}, false),

				Description: "Indicate whether layer 2 connections are present among FGSP members.",
				Optional:    true,
				Computed:    true,
			},
			"psksecret": {
				Type: schema.TypeString,

				Description: "Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"session_sync_dev": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Offload session-sync process to kernel and sync sessions using connected interface(s) directly.",
				Optional:         true,
				Computed:         true,
			},
			"standalone_group_id": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Cluster group ID (0 - 255). Must be the same for all members.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemStandaloneClusterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemStandaloneCluster(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemStandaloneCluster(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemStandaloneCluster")
	}

	return resourceSystemStandaloneClusterRead(ctx, d, meta)
}

func resourceSystemStandaloneClusterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemStandaloneCluster(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemStandaloneCluster(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemStandaloneCluster resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemStandaloneCluster")
	}

	return resourceSystemStandaloneClusterRead(ctx, d, meta)
}

func resourceSystemStandaloneClusterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemStandaloneCluster(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemStandaloneCluster(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemStandaloneCluster resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemStandaloneClusterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	ptp := true
	urlparams.PlainTextPassword = &ptp

	o, err := c.Cmdb.ReadSystemStandaloneCluster(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemStandaloneCluster resource: %v", err)
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

	diags := refreshObjectSystemStandaloneCluster(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemStandaloneClusterClusterPeer(d *schema.ResourceData, v *[]models.SystemStandaloneClusterClusterPeer, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.DownIntfsBeforeSessSync; tmp != nil {
				v["down_intfs_before_sess_sync"] = flattenSystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "down_intfs_before_sess_sync"), sort)
			}

			if tmp := cfg.HbInterval; tmp != nil {
				v["hb_interval"] = *tmp
			}

			if tmp := cfg.HbLostThreshold; tmp != nil {
				v["hb_lost_threshold"] = *tmp
			}

			if tmp := cfg.IpsecTunnelSync; tmp != nil {
				v["ipsec_tunnel_sync"] = *tmp
			}

			if tmp := cfg.Peerip; tmp != nil {
				v["peerip"] = *tmp
			}

			if tmp := cfg.Peervd; tmp != nil {
				v["peervd"] = *tmp
			}

			if tmp := cfg.SecondaryAddIpsecRoutes; tmp != nil {
				v["secondary_add_ipsec_routes"] = *tmp
			}

			if tmp := cfg.SessionSyncFilter; tmp != nil {
				v["session_sync_filter"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilter(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "session_sync_filter"), sort)
			}

			if tmp := cfg.SyncId; tmp != nil {
				v["sync_id"] = *tmp
			}

			if tmp := cfg.Syncvd; tmp != nil {
				v["syncvd"] = flattenSystemStandaloneClusterClusterPeerSyncvd(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "syncvd"), sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "sync_id")
	}

	return flat
}

func flattenSystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync(d *schema.ResourceData, v *[]models.SystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilter(d *schema.ResourceData, v *models.SystemStandaloneClusterClusterPeerSessionSyncFilter, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemStandaloneClusterClusterPeerSessionSyncFilter{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.CustomService; tmp != nil {
				v["custom_service"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "custom_service"), sort)
			}

			if tmp := cfg.Dstaddr; tmp != nil {
				if tmp = utils.Ipv4Read(d, fmt.Sprintf("%s.%d.dstaddr", prefix, i), *tmp); tmp != nil {
					v["dstaddr"] = *tmp
				}
			}

			if tmp := cfg.Dstaddr6; tmp != nil {
				v["dstaddr6"] = *tmp
			}

			if tmp := cfg.Dstintf; tmp != nil {
				v["dstintf"] = *tmp
			}

			if tmp := cfg.Srcaddr; tmp != nil {
				if tmp = utils.Ipv4Read(d, fmt.Sprintf("%s.%d.srcaddr", prefix, i), *tmp); tmp != nil {
					v["srcaddr"] = *tmp
				}
			}

			if tmp := cfg.Srcaddr6; tmp != nil {
				v["srcaddr6"] = *tmp
			}

			if tmp := cfg.Srcintf; tmp != nil {
				v["srcintf"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService(d *schema.ResourceData, v *[]models.SystemStandaloneClusterClusterPeerSessionSyncFilterCustomService, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.DstPortRange; tmp != nil {
				v["dst_port_range"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.SrcPortRange; tmp != nil {
				v["src_port_range"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemStandaloneClusterClusterPeerSyncvd(d *schema.ResourceData, v *[]models.SystemStandaloneClusterClusterPeerSyncvd, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectSystemStandaloneCluster(d *schema.ResourceData, o *models.SystemStandaloneCluster, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.ClusterPeer != nil {
		if err = d.Set("cluster_peer", flattenSystemStandaloneClusterClusterPeer(d, o.ClusterPeer, "cluster_peer", sort)); err != nil {
			return diag.Errorf("error reading cluster_peer: %v", err)
		}
	}

	if o.Encryption != nil {
		v := *o.Encryption

		if err = d.Set("encryption", v); err != nil {
			return diag.Errorf("error reading encryption: %v", err)
		}
	}

	if o.GroupMemberId != nil {
		v := *o.GroupMemberId

		if err = d.Set("group_member_id", v); err != nil {
			return diag.Errorf("error reading group_member_id: %v", err)
		}
	}

	if o.Layer2Connection != nil {
		v := *o.Layer2Connection

		if err = d.Set("layer2_connection", v); err != nil {
			return diag.Errorf("error reading layer2_connection: %v", err)
		}
	}

	if o.Psksecret != nil {
		v := *o.Psksecret

		if v == "ENC XXXX" {
		} else if err = d.Set("psksecret", v); err != nil {
			return diag.Errorf("error reading psksecret: %v", err)
		}
	}

	if o.SessionSyncDev != nil {
		v := *o.SessionSyncDev

		if err = d.Set("session_sync_dev", v); err != nil {
			return diag.Errorf("error reading session_sync_dev: %v", err)
		}
	}

	if o.StandaloneGroupId != nil {
		v := *o.StandaloneGroupId

		if err = d.Set("standalone_group_id", v); err != nil {
			return diag.Errorf("error reading standalone_group_id: %v", err)
		}
	}

	return nil
}

func expandSystemStandaloneClusterClusterPeer(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemStandaloneClusterClusterPeer, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemStandaloneClusterClusterPeer

	for i := range l {
		tmp := models.SystemStandaloneClusterClusterPeer{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.down_intfs_before_sess_sync", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync
			// 	}
			tmp.DownIntfsBeforeSessSync = v2

		}

		pre_append = fmt.Sprintf("%s.%d.hb_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.HbInterval = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.hb_lost_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.HbLostThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipsec_tunnel_sync", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IpsecTunnelSync = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.peerip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Peerip = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.peervd", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Peervd = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.secondary_add_ipsec_routes", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SecondaryAddIpsecRoutes = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.session_sync_filter", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemStandaloneClusterClusterPeerSessionSyncFilter(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemStandaloneClusterClusterPeerSessionSyncFilter
			// 	}
			tmp.SessionSyncFilter = v2

		}

		pre_append = fmt.Sprintf("%s.%d.sync_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SyncId = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.syncvd", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemStandaloneClusterClusterPeerSyncvd(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemStandaloneClusterClusterPeerSyncvd
			// 	}
			tmp.Syncvd = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync

	for i := range l {
		tmp := models.SystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemStandaloneClusterClusterPeerSessionSyncFilter, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemStandaloneClusterClusterPeerSessionSyncFilter

	for i := range l {
		tmp := models.SystemStandaloneClusterClusterPeerSessionSyncFilter{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.custom_service", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemStandaloneClusterClusterPeerSessionSyncFilterCustomService
			// 	}
			tmp.CustomService = v2

		}

		pre_append = fmt.Sprintf("%s.%d.dstaddr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dstaddr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dstaddr6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dstaddr6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dstintf", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dstintf = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.srcaddr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Srcaddr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.srcaddr6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Srcaddr6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.srcintf", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Srcintf = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemStandaloneClusterClusterPeerSessionSyncFilterCustomService, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemStandaloneClusterClusterPeerSessionSyncFilterCustomService

	for i := range l {
		tmp := models.SystemStandaloneClusterClusterPeerSessionSyncFilterCustomService{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dst_port_range", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DstPortRange = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.src_port_range", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SrcPortRange = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemStandaloneClusterClusterPeerSyncvd(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemStandaloneClusterClusterPeerSyncvd, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemStandaloneClusterClusterPeerSyncvd

	for i := range l {
		tmp := models.SystemStandaloneClusterClusterPeerSyncvd{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemStandaloneCluster(d *schema.ResourceData, sv string) (*models.SystemStandaloneCluster, diag.Diagnostics) {
	obj := models.SystemStandaloneCluster{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("cluster_peer"); ok {
		if !utils.CheckVer(sv, "v7.2.1", "") {
			e := utils.AttributeVersionWarning("cluster_peer", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemStandaloneClusterClusterPeer(d, v, "cluster_peer", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ClusterPeer = t
		}
	} else if d.HasChange("cluster_peer") {
		old, new := d.GetChange("cluster_peer")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ClusterPeer = &[]models.SystemStandaloneClusterClusterPeer{}
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
	if v1, ok := d.GetOk("group_member_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("group_member_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.GroupMemberId = &tmp
		}
	}
	if v1, ok := d.GetOk("layer2_connection"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("layer2_connection", sv)
				diags = append(diags, e)
			}
			obj.Layer2Connection = &v2
		}
	}
	if v1, ok := d.GetOk("psksecret"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("psksecret", sv)
				diags = append(diags, e)
			}
			obj.Psksecret = &v2
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
	if v1, ok := d.GetOk("standalone_group_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("standalone_group_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.StandaloneGroupId = &tmp
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemStandaloneCluster(d *schema.ResourceData, sv string) (*models.SystemStandaloneCluster, diag.Diagnostics) {
	obj := models.SystemStandaloneCluster{}
	diags := diag.Diagnostics{}

	obj.ClusterPeer = &[]models.SystemStandaloneClusterClusterPeer{}

	return &obj, diags
}
