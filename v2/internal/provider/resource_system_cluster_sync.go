// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiGate Session Life Support Protocol (FGSP) session synchronization.

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

func resourceSystemClusterSync() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiGate Session Life Support Protocol (FGSP) session synchronization.",

		CreateContext: resourceSystemClusterSyncCreate,
		ReadContext:   resourceSystemClusterSyncRead,
		UpdateContext: resourceSystemClusterSyncUpdate,
		DeleteContext: resourceSystemClusterSyncDelete,

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
			"allow_append": {
				Type:        schema.TypeBool,
				Description: "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
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
				ValidateFunc: validation.IntBetween(1, 10),

				Description: "Heartbeat interval (1 - 10 sec).",
				Optional:    true,
				Computed:    true,
			},
			"hb_lost_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),

				Description: "Lost heartbeat threshold (1 - 10).",
				Optional:    true,
				Computed:    true,
			},
			"ike_heartbeat_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 60),

				Description: "IKE heartbeat interval (1 - 60 secs).",
				Optional:    true,
				Computed:    true,
			},
			"ike_monitor": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IKE HA monitor.",
				Optional:    true,
				Computed:    true,
			},
			"ike_monitor_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 300),

				Description: "IKE HA monitor interval (10 - 300 secs).",
				Optional:    true,
				Computed:    true,
			},
			"ike_seqjump_speed": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),

				Description: "ESP jump ahead factor (1G - 10G pps equivalent).",
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
			"slave_add_ike_routes": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IKE route announcement on the backup unit.",
				Optional:    true,
				Computed:    true,
			},
			"sync_id": {
				Type: schema.TypeInt,

				Description: "Sync ID.",
				ForceNew:    true,
				Required:    true,
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
	}
}

func resourceSystemClusterSyncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := ""
	key := "sync_id"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating SystemClusterSync resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemClusterSync(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemClusterSync(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemClusterSync")
	}

	return resourceSystemClusterSyncRead(ctx, d, meta)
}

func resourceSystemClusterSyncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemClusterSync(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemClusterSync(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemClusterSync resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemClusterSync")
	}

	return resourceSystemClusterSyncRead(ctx, d, meta)
}

func resourceSystemClusterSyncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemClusterSync(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemClusterSync resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemClusterSyncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemClusterSync(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemClusterSync resource: %v", err)
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

	diags := refreshObjectSystemClusterSync(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemClusterSyncDownIntfsBeforeSessSync(d *schema.ResourceData, v *[]models.SystemClusterSyncDownIntfsBeforeSessSync, prefix string, sort bool) interface{} {
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

func flattenSystemClusterSyncSessionSyncFilter(d *schema.ResourceData, v *models.SystemClusterSyncSessionSyncFilter, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SystemClusterSyncSessionSyncFilter{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.CustomService; tmp != nil {
				v["custom_service"] = flattenSystemClusterSyncSessionSyncFilterCustomService(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "custom_service"), sort)
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

func flattenSystemClusterSyncSessionSyncFilterCustomService(d *schema.ResourceData, v *[]models.SystemClusterSyncSessionSyncFilterCustomService, prefix string, sort bool) interface{} {
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

func flattenSystemClusterSyncSyncvd(d *schema.ResourceData, v *[]models.SystemClusterSyncSyncvd, prefix string, sort bool) interface{} {
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

func refreshObjectSystemClusterSync(d *schema.ResourceData, o *models.SystemClusterSync, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.DownIntfsBeforeSessSync != nil {
		if err = d.Set("down_intfs_before_sess_sync", flattenSystemClusterSyncDownIntfsBeforeSessSync(d, o.DownIntfsBeforeSessSync, "down_intfs_before_sess_sync", sort)); err != nil {
			return diag.Errorf("error reading down_intfs_before_sess_sync: %v", err)
		}
	}

	if o.HbInterval != nil {
		v := *o.HbInterval

		if err = d.Set("hb_interval", v); err != nil {
			return diag.Errorf("error reading hb_interval: %v", err)
		}
	}

	if o.HbLostThreshold != nil {
		v := *o.HbLostThreshold

		if err = d.Set("hb_lost_threshold", v); err != nil {
			return diag.Errorf("error reading hb_lost_threshold: %v", err)
		}
	}

	if o.IkeHeartbeatInterval != nil {
		v := *o.IkeHeartbeatInterval

		if err = d.Set("ike_heartbeat_interval", v); err != nil {
			return diag.Errorf("error reading ike_heartbeat_interval: %v", err)
		}
	}

	if o.IkeMonitor != nil {
		v := *o.IkeMonitor

		if err = d.Set("ike_monitor", v); err != nil {
			return diag.Errorf("error reading ike_monitor: %v", err)
		}
	}

	if o.IkeMonitorInterval != nil {
		v := *o.IkeMonitorInterval

		if err = d.Set("ike_monitor_interval", v); err != nil {
			return diag.Errorf("error reading ike_monitor_interval: %v", err)
		}
	}

	if o.IkeSeqjumpSpeed != nil {
		v := *o.IkeSeqjumpSpeed

		if err = d.Set("ike_seqjump_speed", v); err != nil {
			return diag.Errorf("error reading ike_seqjump_speed: %v", err)
		}
	}

	if o.IpsecTunnelSync != nil {
		v := *o.IpsecTunnelSync

		if err = d.Set("ipsec_tunnel_sync", v); err != nil {
			return diag.Errorf("error reading ipsec_tunnel_sync: %v", err)
		}
	}

	if o.Peerip != nil {
		v := *o.Peerip

		if err = d.Set("peerip", v); err != nil {
			return diag.Errorf("error reading peerip: %v", err)
		}
	}

	if o.Peervd != nil {
		v := *o.Peervd

		if err = d.Set("peervd", v); err != nil {
			return diag.Errorf("error reading peervd: %v", err)
		}
	}

	if o.SecondaryAddIpsecRoutes != nil {
		v := *o.SecondaryAddIpsecRoutes

		if err = d.Set("secondary_add_ipsec_routes", v); err != nil {
			return diag.Errorf("error reading secondary_add_ipsec_routes: %v", err)
		}
	}

	if _, ok := d.GetOk("session_sync_filter"); ok {
		if o.SessionSyncFilter != nil {
			if err = d.Set("session_sync_filter", flattenSystemClusterSyncSessionSyncFilter(d, o.SessionSyncFilter, "session_sync_filter", sort)); err != nil {
				return diag.Errorf("error reading session_sync_filter: %v", err)
			}
		}
	}

	if o.SlaveAddIkeRoutes != nil {
		v := *o.SlaveAddIkeRoutes

		if err = d.Set("slave_add_ike_routes", v); err != nil {
			return diag.Errorf("error reading slave_add_ike_routes: %v", err)
		}
	}

	if o.SyncId != nil {
		v := *o.SyncId

		if err = d.Set("sync_id", v); err != nil {
			return diag.Errorf("error reading sync_id: %v", err)
		}
	}

	if o.Syncvd != nil {
		if err = d.Set("syncvd", flattenSystemClusterSyncSyncvd(d, o.Syncvd, "syncvd", sort)); err != nil {
			return diag.Errorf("error reading syncvd: %v", err)
		}
	}

	return nil
}

func expandSystemClusterSyncDownIntfsBeforeSessSync(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemClusterSyncDownIntfsBeforeSessSync, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemClusterSyncDownIntfsBeforeSessSync

	for i := range l {
		tmp := models.SystemClusterSyncDownIntfsBeforeSessSync{}
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

func expandSystemClusterSyncSessionSyncFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SystemClusterSyncSessionSyncFilter, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemClusterSyncSessionSyncFilter

	for i := range l {
		tmp := models.SystemClusterSyncSessionSyncFilter{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.custom_service", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemClusterSyncSessionSyncFilterCustomService(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemClusterSyncSessionSyncFilterCustomService
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

func expandSystemClusterSyncSessionSyncFilterCustomService(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemClusterSyncSessionSyncFilterCustomService, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemClusterSyncSessionSyncFilterCustomService

	for i := range l {
		tmp := models.SystemClusterSyncSessionSyncFilterCustomService{}
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

func expandSystemClusterSyncSyncvd(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemClusterSyncSyncvd, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemClusterSyncSyncvd

	for i := range l {
		tmp := models.SystemClusterSyncSyncvd{}
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

func getObjectSystemClusterSync(d *schema.ResourceData, sv string) (*models.SystemClusterSync, diag.Diagnostics) {
	obj := models.SystemClusterSync{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("down_intfs_before_sess_sync"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("down_intfs_before_sess_sync", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemClusterSyncDownIntfsBeforeSessSync(d, v, "down_intfs_before_sess_sync", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DownIntfsBeforeSessSync = t
		}
	} else if d.HasChange("down_intfs_before_sess_sync") {
		old, new := d.GetChange("down_intfs_before_sess_sync")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DownIntfsBeforeSessSync = &[]models.SystemClusterSyncDownIntfsBeforeSessSync{}
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
	if v1, ok := d.GetOk("ike_heartbeat_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("ike_heartbeat_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IkeHeartbeatInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("ike_monitor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("ike_monitor", sv)
				diags = append(diags, e)
			}
			obj.IkeMonitor = &v2
		}
	}
	if v1, ok := d.GetOk("ike_monitor_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("ike_monitor_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IkeMonitorInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("ike_seqjump_speed"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "v7.0.1") {
				e := utils.AttributeVersionWarning("ike_seqjump_speed", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IkeSeqjumpSpeed = &tmp
		}
	}
	if v1, ok := d.GetOk("ipsec_tunnel_sync"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipsec_tunnel_sync", sv)
				diags = append(diags, e)
			}
			obj.IpsecTunnelSync = &v2
		}
	}
	if v1, ok := d.GetOk("peerip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("peerip", sv)
				diags = append(diags, e)
			}
			obj.Peerip = &v2
		}
	}
	if v1, ok := d.GetOk("peervd"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("peervd", sv)
				diags = append(diags, e)
			}
			obj.Peervd = &v2
		}
	}
	if v1, ok := d.GetOk("secondary_add_ipsec_routes"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("secondary_add_ipsec_routes", sv)
				diags = append(diags, e)
			}
			obj.SecondaryAddIpsecRoutes = &v2
		}
	}
	if v, ok := d.GetOk("session_sync_filter"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("session_sync_filter", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemClusterSyncSessionSyncFilter(d, v, "session_sync_filter", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SessionSyncFilter = t
		}
	} else if d.HasChange("session_sync_filter") {
		old, new := d.GetChange("session_sync_filter")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SessionSyncFilter = &models.SystemClusterSyncSessionSyncFilter{}
		}
	}
	if v1, ok := d.GetOk("slave_add_ike_routes"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.1") {
				e := utils.AttributeVersionWarning("slave_add_ike_routes", sv)
				diags = append(diags, e)
			}
			obj.SlaveAddIkeRoutes = &v2
		}
	}
	if v1, ok := d.GetOk("sync_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sync_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SyncId = &tmp
		}
	}
	if v, ok := d.GetOk("syncvd"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("syncvd", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemClusterSyncSyncvd(d, v, "syncvd", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Syncvd = t
		}
	} else if d.HasChange("syncvd") {
		old, new := d.GetChange("syncvd")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Syncvd = &[]models.SystemClusterSyncSyncvd{}
		}
	}
	return &obj, diags
}
