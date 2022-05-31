// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPS global parameter.

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
)

func resourceIpsGlobal() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPS global parameter.",

		CreateContext: resourceIpsGlobalCreate,
		ReadContext:   resourceIpsGlobalRead,
		UpdateContext: resourceIpsGlobalUpdate,
		DeleteContext: resourceIpsGlobalDelete,

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
			"anomaly_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"periodical", "continuous"}, false),

				Description: "Global blocking mode for rate-based anomalies.",
				Optional:    true,
				Computed:    true,
			},
			"cp_accel_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "basic", "advanced"}, false),

				Description: "IPS Pattern matching acceleration/offloading to CPx processors.",
				Optional:    true,
				Computed:    true,
			},
			"database": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"regular", "extended"}, false),

				Description: "Regular or extended IPS database. Regular protects against the latest common and in-the-wild attacks. Extended includes protection from legacy attacks.",
				Optional:    true,
				Computed:    true,
			},
			"deep_app_insp_db_limit": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2147483647),

				Description: "Limit on number of entries in deep application inspection database (1 - 2147483647, use recommended setting = 0).",
				Optional:    true,
				Computed:    true,
			},
			"deep_app_insp_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2147483647),

				Description: "Timeout for Deep application inspection (1 - 2147483647 sec., 0 = use recommended setting).",
				Optional:    true,
				Computed:    true,
			},
			"engine_count": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Number of IPS engines running. If set to the default value of 0, FortiOS sets the number to optimize performance depending on the number of CPU cores.",
				Optional:    true,
				Computed:    true,
			},
			"exclude_signatures": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "industrial"}, false),

				Description: "Excluded signatures.",
				Optional:    true,
				Computed:    true,
			},
			"fail_open": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to allow traffic if the IPS buffer is full. Default is disable and IPS traffic is blocked when the IPS buffer is full.",
				Optional:    true,
				Computed:    true,
			},
			"intelligent_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPS adaptive scanning (intelligent mode). Intelligent mode optimizes the scanning method for the type of traffic.",
				Optional:    true,
				Computed:    true,
			},
			"ips_reserve_cpu": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable IPS daemon's use of CPUs other than CPU 0",
				Optional:    true,
				Computed:    true,
			},
			"ngfw_max_scan_range": {
				Type: schema.TypeInt,

				Description: "NGFW policy-mode app detection threshold.",
				Optional:    true,
				Computed:    true,
			},
			"np_accel_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "basic"}, false),

				Description: "Acceleration mode for IPS processing by NPx processors.",
				Optional:    true,
				Computed:    true,
			},
			"packet_log_queue_depth": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(128, 4096),

				Description: "Packet/pcap log queue depth per IPS engine.",
				Optional:    true,
				Computed:    true,
			},
			"session_limit_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"accurate", "heuristic"}, false),

				Description: "Method of counting concurrent sessions used by session limit anomalies. Choose between greater accuracy (accurate) or improved performance (heuristics).",
				Optional:    true,
				Computed:    true,
			},
			"skype_client_public_ipaddr": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Public IP addresses of your network that receive Skype sessions. Helps identify Skype sessions. Separate IP addresses with commas.",
				Optional:    true,
				Computed:    true,
			},
			"socket_size": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),

				Description: "IPS socket buffer size. Max and default value depend on available memory. Can be changed to tune performance.",
				Optional:    true,
				Computed:    true,
			},
			"sync_session_ttl": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of kernel session TTL for IPS sessions.",
				Optional:    true,
				Computed:    true,
			},
			"tls_active_probe": {
				Type:        schema.TypeList,
				Description: "TLS active probe configuration.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Specify outgoing interface to reach server.",
							Optional:    true,
							Computed:    true,
						},
						"interface_select_method": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"auto", "sdwan", "specify"}, false),

							Description: "Specify how to select outgoing interface to reach server.",
							Optional:    true,
							Computed:    true,
						},
						"source_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Source IP address used for TLS active probe.",
							Optional:    true,
							Computed:    true,
						},
						"source_ip6": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "Source IPv6 address used for TLS active probe.",
							Optional:         true,
							Computed:         true,
						},
						"vdom": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),

							Description: "Virtual domain name for TLS active probe.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"traffic_submit": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable submitting attack data found by this FortiGate to FortiGuard.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceIpsGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectIpsGlobal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateIpsGlobal(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("IpsGlobal")
	}

	return resourceIpsGlobalRead(ctx, d, meta)
}

func resourceIpsGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectIpsGlobal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateIpsGlobal(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating IpsGlobal resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("IpsGlobal")
	}

	return resourceIpsGlobalRead(ctx, d, meta)
}

func resourceIpsGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectIpsGlobal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateIpsGlobal(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting IpsGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadIpsGlobal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading IpsGlobal resource: %v", err)
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

	diags := refreshObjectIpsGlobal(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenIpsGlobalTlsActiveProbe(d *schema.ResourceData, v *models.IpsGlobalTlsActiveProbe, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.IpsGlobalTlsActiveProbe{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			if tmp := cfg.InterfaceSelectMethod; tmp != nil {
				v["interface_select_method"] = *tmp
			}

			if tmp := cfg.SourceIp; tmp != nil {
				v["source_ip"] = *tmp
			}

			if tmp := cfg.SourceIp6; tmp != nil {
				v["source_ip6"] = *tmp
			}

			if tmp := cfg.Vdom; tmp != nil {
				v["vdom"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func refreshObjectIpsGlobal(d *schema.ResourceData, o *models.IpsGlobal, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AnomalyMode != nil {
		v := *o.AnomalyMode

		if err = d.Set("anomaly_mode", v); err != nil {
			return diag.Errorf("error reading anomaly_mode: %v", err)
		}
	}

	if o.CpAccelMode != nil {
		v := *o.CpAccelMode

		if err = d.Set("cp_accel_mode", v); err != nil {
			return diag.Errorf("error reading cp_accel_mode: %v", err)
		}
	}

	if o.Database != nil {
		v := *o.Database

		if err = d.Set("database", v); err != nil {
			return diag.Errorf("error reading database: %v", err)
		}
	}

	if o.DeepAppInspDbLimit != nil {
		v := *o.DeepAppInspDbLimit

		if err = d.Set("deep_app_insp_db_limit", v); err != nil {
			return diag.Errorf("error reading deep_app_insp_db_limit: %v", err)
		}
	}

	if o.DeepAppInspTimeout != nil {
		v := *o.DeepAppInspTimeout

		if err = d.Set("deep_app_insp_timeout", v); err != nil {
			return diag.Errorf("error reading deep_app_insp_timeout: %v", err)
		}
	}

	if o.EngineCount != nil {
		v := *o.EngineCount

		if err = d.Set("engine_count", v); err != nil {
			return diag.Errorf("error reading engine_count: %v", err)
		}
	}

	if o.ExcludeSignatures != nil {
		v := *o.ExcludeSignatures

		if err = d.Set("exclude_signatures", v); err != nil {
			return diag.Errorf("error reading exclude_signatures: %v", err)
		}
	}

	if o.FailOpen != nil {
		v := *o.FailOpen

		if err = d.Set("fail_open", v); err != nil {
			return diag.Errorf("error reading fail_open: %v", err)
		}
	}

	if o.IntelligentMode != nil {
		v := *o.IntelligentMode

		if err = d.Set("intelligent_mode", v); err != nil {
			return diag.Errorf("error reading intelligent_mode: %v", err)
		}
	}

	if o.IpsReserveCpu != nil {
		v := *o.IpsReserveCpu

		if err = d.Set("ips_reserve_cpu", v); err != nil {
			return diag.Errorf("error reading ips_reserve_cpu: %v", err)
		}
	}

	if o.NgfwMaxScanRange != nil {
		v := *o.NgfwMaxScanRange

		if err = d.Set("ngfw_max_scan_range", v); err != nil {
			return diag.Errorf("error reading ngfw_max_scan_range: %v", err)
		}
	}

	if o.NpAccelMode != nil {
		v := *o.NpAccelMode

		if err = d.Set("np_accel_mode", v); err != nil {
			return diag.Errorf("error reading np_accel_mode: %v", err)
		}
	}

	if o.PacketLogQueueDepth != nil {
		v := *o.PacketLogQueueDepth

		if err = d.Set("packet_log_queue_depth", v); err != nil {
			return diag.Errorf("error reading packet_log_queue_depth: %v", err)
		}
	}

	if o.SessionLimitMode != nil {
		v := *o.SessionLimitMode

		if err = d.Set("session_limit_mode", v); err != nil {
			return diag.Errorf("error reading session_limit_mode: %v", err)
		}
	}

	if o.SkypeClientPublicIpaddr != nil {
		v := *o.SkypeClientPublicIpaddr

		if err = d.Set("skype_client_public_ipaddr", v); err != nil {
			return diag.Errorf("error reading skype_client_public_ipaddr: %v", err)
		}
	}

	if o.SocketSize != nil {
		v := *o.SocketSize

		if err = d.Set("socket_size", v); err != nil {
			return diag.Errorf("error reading socket_size: %v", err)
		}
	}

	if o.SyncSessionTtl != nil {
		v := *o.SyncSessionTtl

		if err = d.Set("sync_session_ttl", v); err != nil {
			return diag.Errorf("error reading sync_session_ttl: %v", err)
		}
	}

	if _, ok := d.GetOk("tls_active_probe"); ok {
		if o.TlsActiveProbe != nil {
			if err = d.Set("tls_active_probe", flattenIpsGlobalTlsActiveProbe(d, o.TlsActiveProbe, "tls_active_probe", sort)); err != nil {
				return diag.Errorf("error reading tls_active_probe: %v", err)
			}
		}
	}

	if o.TrafficSubmit != nil {
		v := *o.TrafficSubmit

		if err = d.Set("traffic_submit", v); err != nil {
			return diag.Errorf("error reading traffic_submit: %v", err)
		}
	}

	return nil
}

func expandIpsGlobalTlsActiveProbe(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.IpsGlobalTlsActiveProbe, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.IpsGlobalTlsActiveProbe

	for i := range l {
		tmp := models.IpsGlobalTlsActiveProbe{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface_select_method", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InterfaceSelectMethod = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.source_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SourceIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.source_ip6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SourceIp6 = &v2
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
	return &result[0], nil
}

func getObjectIpsGlobal(d *schema.ResourceData, sv string) (*models.IpsGlobal, diag.Diagnostics) {
	obj := models.IpsGlobal{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("anomaly_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("anomaly_mode", sv)
				diags = append(diags, e)
			}
			obj.AnomalyMode = &v2
		}
	}
	if v1, ok := d.GetOk("cp_accel_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("cp_accel_mode", sv)
				diags = append(diags, e)
			}
			obj.CpAccelMode = &v2
		}
	}
	if v1, ok := d.GetOk("database"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("database", sv)
				diags = append(diags, e)
			}
			obj.Database = &v2
		}
	}
	if v1, ok := d.GetOk("deep_app_insp_db_limit"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("deep_app_insp_db_limit", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DeepAppInspDbLimit = &tmp
		}
	}
	if v1, ok := d.GetOk("deep_app_insp_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("deep_app_insp_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DeepAppInspTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("engine_count"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("engine_count", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EngineCount = &tmp
		}
	}
	if v1, ok := d.GetOk("exclude_signatures"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("exclude_signatures", sv)
				diags = append(diags, e)
			}
			obj.ExcludeSignatures = &v2
		}
	}
	if v1, ok := d.GetOk("fail_open"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fail_open", sv)
				diags = append(diags, e)
			}
			obj.FailOpen = &v2
		}
	}
	if v1, ok := d.GetOk("intelligent_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.3") {
				e := utils.AttributeVersionWarning("intelligent_mode", sv)
				diags = append(diags, e)
			}
			obj.IntelligentMode = &v2
		}
	}
	if v1, ok := d.GetOk("ips_reserve_cpu"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("ips_reserve_cpu", sv)
				diags = append(diags, e)
			}
			obj.IpsReserveCpu = &v2
		}
	}
	if v1, ok := d.GetOk("ngfw_max_scan_range"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("ngfw_max_scan_range", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.NgfwMaxScanRange = &tmp
		}
	}
	if v1, ok := d.GetOk("np_accel_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("np_accel_mode", sv)
				diags = append(diags, e)
			}
			obj.NpAccelMode = &v2
		}
	}
	if v1, ok := d.GetOk("packet_log_queue_depth"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("packet_log_queue_depth", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PacketLogQueueDepth = &tmp
		}
	}
	if v1, ok := d.GetOk("session_limit_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("session_limit_mode", sv)
				diags = append(diags, e)
			}
			obj.SessionLimitMode = &v2
		}
	}
	if v1, ok := d.GetOk("skype_client_public_ipaddr"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("skype_client_public_ipaddr", sv)
				diags = append(diags, e)
			}
			obj.SkypeClientPublicIpaddr = &v2
		}
	}
	if v1, ok := d.GetOk("socket_size"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("socket_size", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SocketSize = &tmp
		}
	}
	if v1, ok := d.GetOk("sync_session_ttl"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sync_session_ttl", sv)
				diags = append(diags, e)
			}
			obj.SyncSessionTtl = &v2
		}
	}
	if v, ok := d.GetOk("tls_active_probe"); ok {
		if !utils.CheckVer(sv, "", "v6.4.0") {
			e := utils.AttributeVersionWarning("tls_active_probe", sv)
			diags = append(diags, e)
		}
		t, err := expandIpsGlobalTlsActiveProbe(d, v, "tls_active_probe", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.TlsActiveProbe = t
		}
	} else if d.HasChange("tls_active_probe") {
		old, new := d.GetChange("tls_active_probe")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.TlsActiveProbe = &models.IpsGlobalTlsActiveProbe{}
		}
	}
	if v1, ok := d.GetOk("traffic_submit"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("traffic_submit", sv)
				diags = append(diags, e)
			}
			obj.TrafficSubmit = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectIpsGlobal(d *schema.ResourceData, sv string) (*models.IpsGlobal, diag.Diagnostics) {
	obj := models.IpsGlobal{}
	diags := diag.Diagnostics{}

	obj.TlsActiveProbe = &models.IpsGlobalTlsActiveProbe{}

	return &obj, diags
}
