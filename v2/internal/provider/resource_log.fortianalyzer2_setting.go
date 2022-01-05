// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Global FortiAnalyzer settings.

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

func resourceLogFortianalyzer2Setting() *schema.Resource {
	return &schema.Resource{
		Description: "Global FortiAnalyzer settings.",

		CreateContext: resourceLogFortianalyzer2SettingCreate,
		ReadContext:   resourceLogFortianalyzer2SettingRead,
		UpdateContext: resourceLogFortianalyzer2SettingUpdate,
		DeleteContext: resourceLogFortianalyzer2SettingDelete,

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
			"access_config": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiAnalyzer access to configuration and data.",
				Optional:    true,
				Computed:    true,
			},
			"certificate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Certificate used to communicate with FortiAnalyzer.",
				Optional:    true,
				Computed:    true,
			},
			"certificate_verification": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable identity verification of FortiAnalyzer by use of certificate.",
				Optional:    true,
				Computed:    true,
			},
			"conn_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),

				Description: "FortiAnalyzer connection time-out in seconds (for status and log buffer).",
				Optional:    true,
				Computed:    true,
			},
			"enc_algorithm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"high-medium", "high", "low"}, false),

				Description: "Configure the level of SSL protection for secure communication with FortiAnalyzer.",
				Optional:    true,
				Computed:    true,
			},
			"hmac_algorithm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"sha256", "sha1"}, false),

				Description: "FortiAnalyzer IPsec tunnel HMAC algorithm.",
				Optional:    true,
				Computed:    true,
			},
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
			"ips_archive": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPS packet archive logging.",
				Optional:    true,
				Computed:    true,
			},
			"max_log_rate": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100000),

				Description: "FortiAnalyzer maximum log rate in MBps (0 = unlimited).",
				Optional:    true,
				Computed:    true,
			},
			"monitor_failure_retry_period": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 86400),

				Description: "Time between FortiAnalyzer connection retries in seconds (for status and log buffer).",
				Optional:    true,
				Computed:    true,
			},
			"monitor_keepalive_period": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 120),

				Description: "Time between OFTP keepalives in seconds (for status and log buffer).",
				Optional:    true,
				Computed:    true,
			},
			"preshared_key": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Preshared-key used for auto-authorization on FortiAnalyzer.",
				Optional:    true,
				Computed:    true,
			},
			"priority": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"default", "low"}, false),

				Description: "Set log transmission priority.",
				Optional:    true,
				Computed:    true,
			},
			"reliable": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable reliable logging to FortiAnalyzer.",
				Optional:    true,
				Computed:    true,
			},
			"serial": {
				Type:        schema.TypeList,
				Description: "Serial numbers of the FortiAnalyzer.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Serial Number.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "The remote FortiAnalyzer.",
				Optional:    true,
				Computed:    true,
			},
			"source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_min_proto_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"default", "SSLv3", "TLSv1", "TLSv1-1", "TLSv1-2"}, false),

				Description: "Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging to FortiAnalyzer.",
				Optional:    true,
				Computed:    true,
			},
			"upload_day": {
				Type: schema.TypeString,

				Description: "Day of week (month) to upload logs.",
				Optional:    true,
				Computed:    true,
			},
			"upload_interval": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"daily", "weekly", "monthly"}, false),

				Description: "Frequency to upload log files to FortiAnalyzer.",
				Optional:    true,
				Computed:    true,
			},
			"upload_option": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"store-and-upload", "realtime", "1-minute", "5-minute"}, false),

				Description: "Enable/disable logging to hard disk and then uploading to FortiAnalyzer.",
				Optional:    true,
				Computed:    true,
			},
			"upload_time": {
				Type: schema.TypeString,

				Description: "Time to upload logs (hh:mm).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceLogFortianalyzer2SettingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogFortianalyzer2Setting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateLogFortianalyzer2Setting(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogFortianalyzer2Setting")
	}

	return resourceLogFortianalyzer2SettingRead(ctx, d, meta)
}

func resourceLogFortianalyzer2SettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogFortianalyzer2Setting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateLogFortianalyzer2Setting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating LogFortianalyzer2Setting resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogFortianalyzer2Setting")
	}

	return resourceLogFortianalyzer2SettingRead(ctx, d, meta)
}

func resourceLogFortianalyzer2SettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectLogFortianalyzer2Setting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateLogFortianalyzer2Setting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting LogFortianalyzer2Setting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogFortianalyzer2SettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadLogFortianalyzer2Setting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogFortianalyzer2Setting resource: %v", err)
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

	diags := refreshObjectLogFortianalyzer2Setting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenLogFortianalyzer2SettingSerial(v *[]models.LogFortianalyzer2SettingSerial, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func refreshObjectLogFortianalyzer2Setting(d *schema.ResourceData, o *models.LogFortianalyzer2Setting, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AccessConfig != nil {
		v := *o.AccessConfig

		if err = d.Set("access_config", v); err != nil {
			return diag.Errorf("error reading access_config: %v", err)
		}
	}

	if o.Certificate != nil {
		v := *o.Certificate

		if err = d.Set("certificate", v); err != nil {
			return diag.Errorf("error reading certificate: %v", err)
		}
	}

	if o.CertificateVerification != nil {
		v := *o.CertificateVerification

		if err = d.Set("certificate_verification", v); err != nil {
			return diag.Errorf("error reading certificate_verification: %v", err)
		}
	}

	if o.ConnTimeout != nil {
		v := *o.ConnTimeout

		if err = d.Set("conn_timeout", v); err != nil {
			return diag.Errorf("error reading conn_timeout: %v", err)
		}
	}

	if o.EncAlgorithm != nil {
		v := *o.EncAlgorithm

		if err = d.Set("enc_algorithm", v); err != nil {
			return diag.Errorf("error reading enc_algorithm: %v", err)
		}
	}

	if o.HmacAlgorithm != nil {
		v := *o.HmacAlgorithm

		if err = d.Set("hmac_algorithm", v); err != nil {
			return diag.Errorf("error reading hmac_algorithm: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.InterfaceSelectMethod != nil {
		v := *o.InterfaceSelectMethod

		if err = d.Set("interface_select_method", v); err != nil {
			return diag.Errorf("error reading interface_select_method: %v", err)
		}
	}

	if o.IpsArchive != nil {
		v := *o.IpsArchive

		if err = d.Set("ips_archive", v); err != nil {
			return diag.Errorf("error reading ips_archive: %v", err)
		}
	}

	if o.MaxLogRate != nil {
		v := *o.MaxLogRate

		if err = d.Set("max_log_rate", v); err != nil {
			return diag.Errorf("error reading max_log_rate: %v", err)
		}
	}

	if o.MonitorFailureRetryPeriod != nil {
		v := *o.MonitorFailureRetryPeriod

		if err = d.Set("monitor_failure_retry_period", v); err != nil {
			return diag.Errorf("error reading monitor_failure_retry_period: %v", err)
		}
	}

	if o.MonitorKeepalivePeriod != nil {
		v := *o.MonitorKeepalivePeriod

		if err = d.Set("monitor_keepalive_period", v); err != nil {
			return diag.Errorf("error reading monitor_keepalive_period: %v", err)
		}
	}

	if o.PresharedKey != nil {
		v := *o.PresharedKey

		if err = d.Set("preshared_key", v); err != nil {
			return diag.Errorf("error reading preshared_key: %v", err)
		}
	}

	if o.Priority != nil {
		v := *o.Priority

		if err = d.Set("priority", v); err != nil {
			return diag.Errorf("error reading priority: %v", err)
		}
	}

	if o.Reliable != nil {
		v := *o.Reliable

		if err = d.Set("reliable", v); err != nil {
			return diag.Errorf("error reading reliable: %v", err)
		}
	}

	if o.Serial != nil {
		if err = d.Set("serial", flattenLogFortianalyzer2SettingSerial(o.Serial, sort)); err != nil {
			return diag.Errorf("error reading serial: %v", err)
		}
	}

	if o.Server != nil {
		v := *o.Server

		if err = d.Set("server", v); err != nil {
			return diag.Errorf("error reading server: %v", err)
		}
	}

	if o.SourceIp != nil {
		v := *o.SourceIp

		if err = d.Set("source_ip", v); err != nil {
			return diag.Errorf("error reading source_ip: %v", err)
		}
	}

	if o.SslMinProtoVersion != nil {
		v := *o.SslMinProtoVersion

		if err = d.Set("ssl_min_proto_version", v); err != nil {
			return diag.Errorf("error reading ssl_min_proto_version: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.UploadDay != nil {
		v := *o.UploadDay

		if err = d.Set("upload_day", v); err != nil {
			return diag.Errorf("error reading upload_day: %v", err)
		}
	}

	if o.UploadInterval != nil {
		v := *o.UploadInterval

		if err = d.Set("upload_interval", v); err != nil {
			return diag.Errorf("error reading upload_interval: %v", err)
		}
	}

	if o.UploadOption != nil {
		v := *o.UploadOption

		if err = d.Set("upload_option", v); err != nil {
			return diag.Errorf("error reading upload_option: %v", err)
		}
	}

	if o.UploadTime != nil {
		v := *o.UploadTime

		if err = d.Set("upload_time", v); err != nil {
			return diag.Errorf("error reading upload_time: %v", err)
		}
	}

	return nil
}

func expandLogFortianalyzer2SettingSerial(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.LogFortianalyzer2SettingSerial, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.LogFortianalyzer2SettingSerial

	for i := range l {
		tmp := models.LogFortianalyzer2SettingSerial{}
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

func getObjectLogFortianalyzer2Setting(d *schema.ResourceData, sv string) (*models.LogFortianalyzer2Setting, diag.Diagnostics) {
	obj := models.LogFortianalyzer2Setting{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("access_config"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("access_config", sv)
				diags = append(diags, e)
			}
			obj.AccessConfig = &v2
		}
	}
	if v1, ok := d.GetOk("certificate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("certificate", sv)
				diags = append(diags, e)
			}
			obj.Certificate = &v2
		}
	}
	if v1, ok := d.GetOk("certificate_verification"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("certificate_verification", sv)
				diags = append(diags, e)
			}
			obj.CertificateVerification = &v2
		}
	}
	if v1, ok := d.GetOk("conn_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("conn_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ConnTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("enc_algorithm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("enc_algorithm", sv)
				diags = append(diags, e)
			}
			obj.EncAlgorithm = &v2
		}
	}
	if v1, ok := d.GetOk("hmac_algorithm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hmac_algorithm", sv)
				diags = append(diags, e)
			}
			obj.HmacAlgorithm = &v2
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("interface_select_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("interface_select_method", sv)
				diags = append(diags, e)
			}
			obj.InterfaceSelectMethod = &v2
		}
	}
	if v1, ok := d.GetOk("ips_archive"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ips_archive", sv)
				diags = append(diags, e)
			}
			obj.IpsArchive = &v2
		}
	}
	if v1, ok := d.GetOk("max_log_rate"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_log_rate", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxLogRate = &tmp
		}
	}
	if v1, ok := d.GetOk("monitor_failure_retry_period"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("monitor_failure_retry_period", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MonitorFailureRetryPeriod = &tmp
		}
	}
	if v1, ok := d.GetOk("monitor_keepalive_period"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("monitor_keepalive_period", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MonitorKeepalivePeriod = &tmp
		}
	}
	if v1, ok := d.GetOk("preshared_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("preshared_key", sv)
				diags = append(diags, e)
			}
			obj.PresharedKey = &v2
		}
	}
	if v1, ok := d.GetOk("priority"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("priority", sv)
				diags = append(diags, e)
			}
			obj.Priority = &v2
		}
	}
	if v1, ok := d.GetOk("reliable"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("reliable", sv)
				diags = append(diags, e)
			}
			obj.Reliable = &v2
		}
	}
	if v, ok := d.GetOk("serial"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("serial", sv)
			diags = append(diags, e)
		}
		t, err := expandLogFortianalyzer2SettingSerial(d, v, "serial", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Serial = t
		}
	} else if d.HasChange("serial") {
		old, new := d.GetChange("serial")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Serial = &[]models.LogFortianalyzer2SettingSerial{}
		}
	}
	if v1, ok := d.GetOk("server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server", sv)
				diags = append(diags, e)
			}
			obj.Server = &v2
		}
	}
	if v1, ok := d.GetOk("source_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_ip", sv)
				diags = append(diags, e)
			}
			obj.SourceIp = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_min_proto_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_min_proto_version", sv)
				diags = append(diags, e)
			}
			obj.SslMinProtoVersion = &v2
		}
	}
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	if v1, ok := d.GetOk("upload_day"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("upload_day", sv)
				diags = append(diags, e)
			}
			obj.UploadDay = &v2
		}
	}
	if v1, ok := d.GetOk("upload_interval"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("upload_interval", sv)
				diags = append(diags, e)
			}
			obj.UploadInterval = &v2
		}
	}
	if v1, ok := d.GetOk("upload_option"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("upload_option", sv)
				diags = append(diags, e)
			}
			obj.UploadOption = &v2
		}
	}
	if v1, ok := d.GetOk("upload_time"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("upload_time", sv)
				diags = append(diags, e)
			}
			obj.UploadTime = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectLogFortianalyzer2Setting(d *schema.ResourceData, sv string) (*models.LogFortianalyzer2Setting, diag.Diagnostics) {
	obj := models.LogFortianalyzer2Setting{}
	diags := diag.Diagnostics{}

	obj.Serial = &[]models.LogFortianalyzer2SettingSerial{}

	return &obj, diags
}
