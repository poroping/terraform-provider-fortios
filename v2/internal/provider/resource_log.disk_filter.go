// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure filters for local disk logging. Use these filters to determine the log messages to record according to severity and type.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceLogDiskFilter() *schema.Resource {
	return &schema.Resource{
		Description: "Configure filters for local disk logging. Use these filters to determine the log messages to record according to severity and type.",

		CreateContext: resourceLogDiskFilterCreate,
		ReadContext:   resourceLogDiskFilterRead,
		UpdateContext: resourceLogDiskFilterUpdate,
		DeleteContext: resourceLogDiskFilterDelete,

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
			"admin": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable admin login/logout logging.",
				Optional:    true,
				Computed:    true,
			},
			"anomaly": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable anomaly logging.",
				Optional:    true,
				Computed:    true,
			},
			"auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable firewall authentication logging.",
				Optional:    true,
				Computed:    true,
			},
			"cpu_memory_usage": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable CPU & memory usage logging every 5 minutes.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable DHCP service messages logging.",
				Optional:    true,
				Computed:    true,
			},
			"dlp_archive": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable DLP archive logging.",
				Optional:    true,
				Computed:    true,
			},
			"event": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable event logging.",
				Optional:    true,
				Computed:    true,
			},
			"filter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),

				Description: "Disk log filter.",
				Optional:    true,
				Computed:    true,
			},
			"filter_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"include", "exclude"}, false),

				Description: "Include/exclude logs that match the filter.",
				Optional:    true,
				Computed:    true,
			},
			"forward_traffic": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable forward traffic logging.",
				Optional:    true,
				Computed:    true,
			},
			"free_style": {
				Type:        schema.TypeList,
				Description: "Free style filters.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"traffic", "event", "virus", "webfilter", "attack", "spam", "anomaly", "voip", "dlp", "app-ctrl", "waf", "gtp", "dns", "ssh", "ssl", "file-filter", "icap"}, false),

							Description: "Log category.",
							Optional:    true,
							Computed:    true,
						},
						"filter": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1023),

							Description: "Free style filter string.",
							Optional:    true,
							Computed:    true,
						},
						"filter_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"include", "exclude"}, false),

							Description: "Include/exclude logs that match the filter.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Entry ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"gtp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable GTP messages logging.",
				Optional:    true,
				Computed:    true,
			},
			"ha": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable HA logging.",
				Optional:    true,
				Computed:    true,
			},
			"ipsec": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPsec negotiation messages logging.",
				Optional:    true,
				Computed:    true,
			},
			"ldb_monitor": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable VIP real server health monitoring logging.",
				Optional:    true,
				Computed:    true,
			},
			"local_traffic": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable local in or out traffic logging.",
				Optional:    true,
				Computed:    true,
			},
			"multicast_traffic": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable multicast traffic logging.",
				Optional:    true,
				Computed:    true,
			},
			"pattern": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable pattern update logging.",
				Optional:    true,
				Computed:    true,
			},
			"ppp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable L2TP/PPTP/PPPoE logging.",
				Optional:    true,
				Computed:    true,
			},
			"radius": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable RADIUS messages logging.",
				Optional:    true,
				Computed:    true,
			},
			"severity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"emergency", "alert", "critical", "error", "warning", "notification", "information", "debug"}, false),

				Description: "Log to disk every message above and including this severity level.",
				Optional:    true,
				Computed:    true,
			},
			"sniffer_traffic": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable sniffer traffic logging.",
				Optional:    true,
				Computed:    true,
			},
			"sslvpn_log_adm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SSL administrator login logging.",
				Optional:    true,
				Computed:    true,
			},
			"sslvpn_log_auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SSL user authentication logging.",
				Optional:    true,
				Computed:    true,
			},
			"sslvpn_log_session": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SSL session logging.",
				Optional:    true,
				Computed:    true,
			},
			"system": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable system activity logging.",
				Optional:    true,
				Computed:    true,
			},
			"vip_ssl": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable VIP SSL logging.",
				Optional:    true,
				Computed:    true,
			},
			"voip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable VoIP logging.",
				Optional:    true,
				Computed:    true,
			},
			"wan_opt": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable WAN optimization event logging.",
				Optional:    true,
				Computed:    true,
			},
			"wireless_activity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable wireless activity event logging.",
				Optional:    true,
				Computed:    true,
			},
			"ztna_traffic": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable ztna traffic logging.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceLogDiskFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogDiskFilter(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateLogDiskFilter(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogDiskFilter")
	}

	return resourceLogDiskFilterRead(ctx, d, meta)
}

func resourceLogDiskFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogDiskFilter(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateLogDiskFilter(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating LogDiskFilter resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogDiskFilter")
	}

	return resourceLogDiskFilterRead(ctx, d, meta)
}

func resourceLogDiskFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectLogDiskFilter(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateLogDiskFilter(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting LogDiskFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogDiskFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadLogDiskFilter(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogDiskFilter resource: %v", err)
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

	diags := refreshObjectLogDiskFilter(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenLogDiskFilterFreeStyle(d *schema.ResourceData, v *[]models.LogDiskFilterFreeStyle, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Category; tmp != nil {
				v["category"] = *tmp
			}

			if tmp := cfg.Filter; tmp != nil {
				v["filter"] = *tmp
			}

			if tmp := cfg.FilterType; tmp != nil {
				v["filter_type"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectLogDiskFilter(d *schema.ResourceData, o *models.LogDiskFilter, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Admin != nil {
		v := *o.Admin

		if err = d.Set("admin", v); err != nil {
			return diag.Errorf("error reading admin: %v", err)
		}
	}

	if o.Anomaly != nil {
		v := *o.Anomaly

		if err = d.Set("anomaly", v); err != nil {
			return diag.Errorf("error reading anomaly: %v", err)
		}
	}

	if o.Auth != nil {
		v := *o.Auth

		if err = d.Set("auth", v); err != nil {
			return diag.Errorf("error reading auth: %v", err)
		}
	}

	if o.CpuMemoryUsage != nil {
		v := *o.CpuMemoryUsage

		if err = d.Set("cpu_memory_usage", v); err != nil {
			return diag.Errorf("error reading cpu_memory_usage: %v", err)
		}
	}

	if o.Dhcp != nil {
		v := *o.Dhcp

		if err = d.Set("dhcp", v); err != nil {
			return diag.Errorf("error reading dhcp: %v", err)
		}
	}

	if o.DlpArchive != nil {
		v := *o.DlpArchive

		if err = d.Set("dlp_archive", v); err != nil {
			return diag.Errorf("error reading dlp_archive: %v", err)
		}
	}

	if o.Event != nil {
		v := *o.Event

		if err = d.Set("event", v); err != nil {
			return diag.Errorf("error reading event: %v", err)
		}
	}

	if o.Filter != nil {
		v := *o.Filter

		if err = d.Set("filter", v); err != nil {
			return diag.Errorf("error reading filter: %v", err)
		}
	}

	if o.FilterType != nil {
		v := *o.FilterType

		if err = d.Set("filter_type", v); err != nil {
			return diag.Errorf("error reading filter_type: %v", err)
		}
	}

	if o.ForwardTraffic != nil {
		v := *o.ForwardTraffic

		if err = d.Set("forward_traffic", v); err != nil {
			return diag.Errorf("error reading forward_traffic: %v", err)
		}
	}

	if o.FreeStyle != nil {
		if err = d.Set("free_style", flattenLogDiskFilterFreeStyle(d, o.FreeStyle, "free_style", sort)); err != nil {
			return diag.Errorf("error reading free_style: %v", err)
		}
	}

	if o.Gtp != nil {
		v := *o.Gtp

		if err = d.Set("gtp", v); err != nil {
			return diag.Errorf("error reading gtp: %v", err)
		}
	}

	if o.Ha != nil {
		v := *o.Ha

		if err = d.Set("ha", v); err != nil {
			return diag.Errorf("error reading ha: %v", err)
		}
	}

	if o.Ipsec != nil {
		v := *o.Ipsec

		if err = d.Set("ipsec", v); err != nil {
			return diag.Errorf("error reading ipsec: %v", err)
		}
	}

	if o.LdbMonitor != nil {
		v := *o.LdbMonitor

		if err = d.Set("ldb_monitor", v); err != nil {
			return diag.Errorf("error reading ldb_monitor: %v", err)
		}
	}

	if o.LocalTraffic != nil {
		v := *o.LocalTraffic

		if err = d.Set("local_traffic", v); err != nil {
			return diag.Errorf("error reading local_traffic: %v", err)
		}
	}

	if o.MulticastTraffic != nil {
		v := *o.MulticastTraffic

		if err = d.Set("multicast_traffic", v); err != nil {
			return diag.Errorf("error reading multicast_traffic: %v", err)
		}
	}

	if o.Pattern != nil {
		v := *o.Pattern

		if err = d.Set("pattern", v); err != nil {
			return diag.Errorf("error reading pattern: %v", err)
		}
	}

	if o.Ppp != nil {
		v := *o.Ppp

		if err = d.Set("ppp", v); err != nil {
			return diag.Errorf("error reading ppp: %v", err)
		}
	}

	if o.Radius != nil {
		v := *o.Radius

		if err = d.Set("radius", v); err != nil {
			return diag.Errorf("error reading radius: %v", err)
		}
	}

	if o.Severity != nil {
		v := *o.Severity

		if err = d.Set("severity", v); err != nil {
			return diag.Errorf("error reading severity: %v", err)
		}
	}

	if o.SnifferTraffic != nil {
		v := *o.SnifferTraffic

		if err = d.Set("sniffer_traffic", v); err != nil {
			return diag.Errorf("error reading sniffer_traffic: %v", err)
		}
	}

	if o.SslvpnLogAdm != nil {
		v := *o.SslvpnLogAdm

		if err = d.Set("sslvpn_log_adm", v); err != nil {
			return diag.Errorf("error reading sslvpn_log_adm: %v", err)
		}
	}

	if o.SslvpnLogAuth != nil {
		v := *o.SslvpnLogAuth

		if err = d.Set("sslvpn_log_auth", v); err != nil {
			return diag.Errorf("error reading sslvpn_log_auth: %v", err)
		}
	}

	if o.SslvpnLogSession != nil {
		v := *o.SslvpnLogSession

		if err = d.Set("sslvpn_log_session", v); err != nil {
			return diag.Errorf("error reading sslvpn_log_session: %v", err)
		}
	}

	if o.System != nil {
		v := *o.System

		if err = d.Set("system", v); err != nil {
			return diag.Errorf("error reading system: %v", err)
		}
	}

	if o.VipSsl != nil {
		v := *o.VipSsl

		if err = d.Set("vip_ssl", v); err != nil {
			return diag.Errorf("error reading vip_ssl: %v", err)
		}
	}

	if o.Voip != nil {
		v := *o.Voip

		if err = d.Set("voip", v); err != nil {
			return diag.Errorf("error reading voip: %v", err)
		}
	}

	if o.WanOpt != nil {
		v := *o.WanOpt

		if err = d.Set("wan_opt", v); err != nil {
			return diag.Errorf("error reading wan_opt: %v", err)
		}
	}

	if o.WirelessActivity != nil {
		v := *o.WirelessActivity

		if err = d.Set("wireless_activity", v); err != nil {
			return diag.Errorf("error reading wireless_activity: %v", err)
		}
	}

	if o.ZtnaTraffic != nil {
		v := *o.ZtnaTraffic

		if err = d.Set("ztna_traffic", v); err != nil {
			return diag.Errorf("error reading ztna_traffic: %v", err)
		}
	}

	return nil
}

func expandLogDiskFilterFreeStyle(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.LogDiskFilterFreeStyle, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.LogDiskFilterFreeStyle

	for i := range l {
		tmp := models.LogDiskFilterFreeStyle{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.category", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Category = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.filter", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Filter = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.filter_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FilterType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectLogDiskFilter(d *schema.ResourceData, sv string) (*models.LogDiskFilter, diag.Diagnostics) {
	obj := models.LogDiskFilter{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("admin"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("admin", sv)
				diags = append(diags, e)
			}
			obj.Admin = &v2
		}
	}
	if v1, ok := d.GetOk("anomaly"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("anomaly", sv)
				diags = append(diags, e)
			}
			obj.Anomaly = &v2
		}
	}
	if v1, ok := d.GetOk("auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("auth", sv)
				diags = append(diags, e)
			}
			obj.Auth = &v2
		}
	}
	if v1, ok := d.GetOk("cpu_memory_usage"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("cpu_memory_usage", sv)
				diags = append(diags, e)
			}
			obj.CpuMemoryUsage = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("dhcp", sv)
				diags = append(diags, e)
			}
			obj.Dhcp = &v2
		}
	}
	if v1, ok := d.GetOk("dlp_archive"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("dlp_archive", sv)
				diags = append(diags, e)
			}
			obj.DlpArchive = &v2
		}
	}
	if v1, ok := d.GetOk("event"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("event", sv)
				diags = append(diags, e)
			}
			obj.Event = &v2
		}
	}
	if v1, ok := d.GetOk("filter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("filter", sv)
				diags = append(diags, e)
			}
			obj.Filter = &v2
		}
	}
	if v1, ok := d.GetOk("filter_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("filter_type", sv)
				diags = append(diags, e)
			}
			obj.FilterType = &v2
		}
	}
	if v1, ok := d.GetOk("forward_traffic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("forward_traffic", sv)
				diags = append(diags, e)
			}
			obj.ForwardTraffic = &v2
		}
	}
	if v, ok := d.GetOk("free_style"); ok {
		if !utils.CheckVer(sv, "v7.0.0", "") {
			e := utils.AttributeVersionWarning("free_style", sv)
			diags = append(diags, e)
		}
		t, err := expandLogDiskFilterFreeStyle(d, v, "free_style", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.FreeStyle = t
		}
	} else if d.HasChange("free_style") {
		old, new := d.GetChange("free_style")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.FreeStyle = &[]models.LogDiskFilterFreeStyle{}
		}
	}
	if v1, ok := d.GetOk("gtp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("gtp", sv)
				diags = append(diags, e)
			}
			obj.Gtp = &v2
		}
	}
	if v1, ok := d.GetOk("ha"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("ha", sv)
				diags = append(diags, e)
			}
			obj.Ha = &v2
		}
	}
	if v1, ok := d.GetOk("ipsec"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("ipsec", sv)
				diags = append(diags, e)
			}
			obj.Ipsec = &v2
		}
	}
	if v1, ok := d.GetOk("ldb_monitor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("ldb_monitor", sv)
				diags = append(diags, e)
			}
			obj.LdbMonitor = &v2
		}
	}
	if v1, ok := d.GetOk("local_traffic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_traffic", sv)
				diags = append(diags, e)
			}
			obj.LocalTraffic = &v2
		}
	}
	if v1, ok := d.GetOk("multicast_traffic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("multicast_traffic", sv)
				diags = append(diags, e)
			}
			obj.MulticastTraffic = &v2
		}
	}
	if v1, ok := d.GetOk("pattern"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("pattern", sv)
				diags = append(diags, e)
			}
			obj.Pattern = &v2
		}
	}
	if v1, ok := d.GetOk("ppp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("ppp", sv)
				diags = append(diags, e)
			}
			obj.Ppp = &v2
		}
	}
	if v1, ok := d.GetOk("radius"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("radius", sv)
				diags = append(diags, e)
			}
			obj.Radius = &v2
		}
	}
	if v1, ok := d.GetOk("severity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("severity", sv)
				diags = append(diags, e)
			}
			obj.Severity = &v2
		}
	}
	if v1, ok := d.GetOk("sniffer_traffic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sniffer_traffic", sv)
				diags = append(diags, e)
			}
			obj.SnifferTraffic = &v2
		}
	}
	if v1, ok := d.GetOk("sslvpn_log_adm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("sslvpn_log_adm", sv)
				diags = append(diags, e)
			}
			obj.SslvpnLogAdm = &v2
		}
	}
	if v1, ok := d.GetOk("sslvpn_log_auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("sslvpn_log_auth", sv)
				diags = append(diags, e)
			}
			obj.SslvpnLogAuth = &v2
		}
	}
	if v1, ok := d.GetOk("sslvpn_log_session"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("sslvpn_log_session", sv)
				diags = append(diags, e)
			}
			obj.SslvpnLogSession = &v2
		}
	}
	if v1, ok := d.GetOk("system"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("system", sv)
				diags = append(diags, e)
			}
			obj.System = &v2
		}
	}
	if v1, ok := d.GetOk("vip_ssl"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("vip_ssl", sv)
				diags = append(diags, e)
			}
			obj.VipSsl = &v2
		}
	}
	if v1, ok := d.GetOk("voip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("voip", sv)
				diags = append(diags, e)
			}
			obj.Voip = &v2
		}
	}
	if v1, ok := d.GetOk("wan_opt"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("wan_opt", sv)
				diags = append(diags, e)
			}
			obj.WanOpt = &v2
		}
	}
	if v1, ok := d.GetOk("wireless_activity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("wireless_activity", sv)
				diags = append(diags, e)
			}
			obj.WirelessActivity = &v2
		}
	}
	if v1, ok := d.GetOk("ztna_traffic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.4", "") {
				e := utils.AttributeVersionWarning("ztna_traffic", sv)
				diags = append(diags, e)
			}
			obj.ZtnaTraffic = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectLogDiskFilter(d *schema.ResourceData, sv string) (*models.LogDiskFilter, diag.Diagnostics) {
	obj := models.LogDiskFilter{}
	diags := diag.Diagnostics{}

	obj.FreeStyle = &[]models.LogDiskFilterFreeStyle{}

	return &obj, diags
}
