// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure logging to FortiCloud.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceLogFortiguardSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Configure logging to FortiCloud.",

		CreateContext: resourceLogFortiguardSettingCreate,
		ReadContext:   resourceLogFortiguardSettingRead,
		UpdateContext: resourceLogFortiguardSettingUpdate,
		DeleteContext: resourceLogFortiguardSettingDelete,

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
			"access_config": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiCloud access to configuration and data.",
				Optional:    true,
				Computed:    true,
			},
			"conn_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),

				Description: "FortiGate Cloud connection timeout in seconds.",
				Optional:    true,
				Computed:    true,
			},
			"enc_algorithm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"high-medium", "high", "low"}, false),

				Description: "Configure the level of SSL protection for secure communication with FortiCloud.",
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
			"max_log_rate": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100000),

				Description: "FortiCloud maximum log rate in MBps (0 = unlimited).",
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
			"source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Source IP address used to connect FortiCloud.",
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

				Description: "Enable/disable logging to FortiCloud.",
				Optional:    true,
				Computed:    true,
			},
			"upload_day": {
				Type: schema.TypeString,

				Description: "Day of week to roll logs.",
				Optional:    true,
				Computed:    true,
			},
			"upload_interval": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"daily", "weekly", "monthly"}, false),

				Description: "Frequency of uploading log files to FortiCloud.",
				Optional:    true,
				Computed:    true,
			},
			"upload_option": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"store-and-upload", "realtime", "1-minute", "5-minute"}, false),

				Description: "Configure how log messages are sent to FortiCloud.",
				Optional:    true,
				Computed:    true,
			},
			"upload_time": {
				Type: schema.TypeString,

				Description: "Time of day to roll logs (hh:mm).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceLogFortiguardSettingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogFortiguardSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateLogFortiguardSetting(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogFortiguardSetting")
	}

	return resourceLogFortiguardSettingRead(ctx, d, meta)
}

func resourceLogFortiguardSettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogFortiguardSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateLogFortiguardSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating LogFortiguardSetting resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogFortiguardSetting")
	}

	return resourceLogFortiguardSettingRead(ctx, d, meta)
}

func resourceLogFortiguardSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectLogFortiguardSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateLogFortiguardSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting LogFortiguardSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogFortiguardSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadLogFortiguardSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogFortiguardSetting resource: %v", err)
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

	diags := refreshObjectLogFortiguardSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectLogFortiguardSetting(d *schema.ResourceData, o *models.LogFortiguardSetting, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AccessConfig != nil {
		v := *o.AccessConfig

		if err = d.Set("access_config", v); err != nil {
			return diag.Errorf("error reading access_config: %v", err)
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

	if o.MaxLogRate != nil {
		v := *o.MaxLogRate

		if err = d.Set("max_log_rate", v); err != nil {
			return diag.Errorf("error reading max_log_rate: %v", err)
		}
	}

	if o.Priority != nil {
		v := *o.Priority

		if err = d.Set("priority", v); err != nil {
			return diag.Errorf("error reading priority: %v", err)
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

func getObjectLogFortiguardSetting(d *schema.ResourceData, sv string) (*models.LogFortiguardSetting, diag.Diagnostics) {
	obj := models.LogFortiguardSetting{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("access_config"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("access_config", sv)
				diags = append(diags, e)
			}
			obj.AccessConfig = &v2
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
	if v1, ok := d.GetOk("priority"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("priority", sv)
				diags = append(diags, e)
			}
			obj.Priority = &v2
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
func getEmptyObjectLogFortiguardSetting(d *schema.ResourceData, sv string) (*models.LogFortiguardSetting, diag.Diagnostics) {
	obj := models.LogFortiguardSetting{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
