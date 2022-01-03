// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Settings for local disk logging.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceLogDiskSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Settings for local disk logging.",

		CreateContext: resourceLogDiskSettingCreate,
		ReadContext:   resourceLogDiskSettingRead,
		UpdateContext: resourceLogDiskSettingUpdate,
		DeleteContext: resourceLogDiskSettingDelete,

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
			"diskfull": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"overwrite", "nolog"}, false),

				Description: "Action to take when disk is full. The system can overwrite the oldest log messages or stop logging when the disk is full (default = overwrite).",
				Optional:    true,
				Computed:    true,
			},
			"dlp_archive_quota": {
				Type: schema.TypeInt,

				Description: "DLP archive quota (MB).",
				Optional:    true,
				Computed:    true,
			},
			"full_final_warning_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(3, 100),

				Description: "Log full final warning threshold as a percent (3 - 100, default = 95).",
				Optional:    true,
				Computed:    true,
			},
			"full_first_warning_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 98),

				Description: "Log full first warning threshold as a percent (1 - 98, default = 75).",
				Optional:    true,
				Computed:    true,
			},
			"full_second_warning_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 99),

				Description: "Log full second warning threshold as a percent (2 - 99, default = 90).",
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

				Description: "Enable/disable IPS packet archiving to the local disk.",
				Optional:    true,
				Computed:    true,
			},
			"log_quota": {
				Type: schema.TypeInt,

				Description: "Disk log quota (MB).",
				Optional:    true,
				Computed:    true,
			},
			"max_log_file_size": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),

				Description: "Maximum log file size before rolling (1 - 100 Mbytes).",
				Optional:    true,
				Computed:    true,
			},
			"max_policy_packet_capture_size": {
				Type: schema.TypeInt,

				Description: "Maximum size of policy sniffer in MB (0 means unlimited).",
				Optional:    true,
				Computed:    true,
			},
			"maximum_log_age": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3650),

				Description: "Delete log files older than (days).",
				Optional:    true,
				Computed:    true,
			},
			"report_quota": {
				Type: schema.TypeInt,

				Description: "Report db quota (MB).",
				Optional:    true,
				Computed:    true,
			},
			"roll_day": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Day of week on which to roll log file.",
				Optional:         true,
				Computed:         true,
			},
			"roll_schedule": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"daily", "weekly"}, false),

				Description: "Frequency to check log file for rolling.",
				Optional:    true,
				Computed:    true,
			},
			"roll_time": {
				Type: schema.TypeString,

				Description: "Time of day to roll the log file (hh:mm).",
				Optional:    true,
				Computed:    true,
			},
			"source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Source IP address to use for uploading disk log files.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable local disk logging.",
				Optional:    true,
				Computed:    true,
			},
			"upload": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable uploading log files when they are rolled.",
				Optional:    true,
				Computed:    true,
			},
			"upload_delete_files": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Delete log files after uploading (default = enable).",
				Optional:    true,
				Computed:    true,
			},
			"upload_destination": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ftp-server"}, false),

				Description: "The type of server to upload log files to. Only FTP is currently supported.",
				Optional:    true,
				Computed:    true,
			},
			"upload_ssl_conn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"default", "high", "low", "disable"}, false),

				Description: "Enable/disable encrypted FTPS communication to upload log files.",
				Optional:    true,
				Computed:    true,
			},
			"uploaddir": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "The remote directory on the FTP server to upload log files to.",
				Optional:    true,
				Computed:    true,
			},
			"uploadip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address of the FTP server to upload log files to.",
				Optional:    true,
				Computed:    true,
			},
			"uploadpass": {
				Type: schema.TypeString,

				Description: "Password required to log into the FTP server to upload disk log files.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"uploadport": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "TCP port to use for communicating with the FTP server (default = 21).",
				Optional:    true,
				Computed:    true,
			},
			"uploadsched": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Set the schedule for uploading log files to the FTP server (default = disable = upload when rolling).",
				Optional:    true,
				Computed:    true,
			},
			"uploadtime": {
				Type: schema.TypeString,

				Description: "Time of day at which log files are uploaded if uploadsched is enabled (hh:mm or hh).",
				Optional:    true,
				Computed:    true,
			},
			"uploadtype": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Types of log files to upload. Separate multiple entries with a space.",
				Optional:         true,
				Computed:         true,
			},
			"uploaduser": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Username required to log into the FTP server to upload disk log files.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceLogDiskSettingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogDiskSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateLogDiskSetting(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogDiskSetting")
	}

	return resourceLogDiskSettingRead(ctx, d, meta)
}

func resourceLogDiskSettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogDiskSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateLogDiskSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating LogDiskSetting resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogDiskSetting")
	}

	return resourceLogDiskSettingRead(ctx, d, meta)
}

func resourceLogDiskSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectLogDiskSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateLogDiskSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting LogDiskSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogDiskSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadLogDiskSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogDiskSetting resource: %v", err)
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

	diags := refreshObjectLogDiskSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectLogDiskSetting(d *schema.ResourceData, o *models.LogDiskSetting, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Diskfull != nil {
		v := *o.Diskfull

		if err = d.Set("diskfull", v); err != nil {
			return diag.Errorf("error reading diskfull: %v", err)
		}
	}

	if o.DlpArchiveQuota != nil {
		v := *o.DlpArchiveQuota

		if err = d.Set("dlp_archive_quota", v); err != nil {
			return diag.Errorf("error reading dlp_archive_quota: %v", err)
		}
	}

	if o.FullFinalWarningThreshold != nil {
		v := *o.FullFinalWarningThreshold

		if err = d.Set("full_final_warning_threshold", v); err != nil {
			return diag.Errorf("error reading full_final_warning_threshold: %v", err)
		}
	}

	if o.FullFirstWarningThreshold != nil {
		v := *o.FullFirstWarningThreshold

		if err = d.Set("full_first_warning_threshold", v); err != nil {
			return diag.Errorf("error reading full_first_warning_threshold: %v", err)
		}
	}

	if o.FullSecondWarningThreshold != nil {
		v := *o.FullSecondWarningThreshold

		if err = d.Set("full_second_warning_threshold", v); err != nil {
			return diag.Errorf("error reading full_second_warning_threshold: %v", err)
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

	if o.LogQuota != nil {
		v := *o.LogQuota

		if err = d.Set("log_quota", v); err != nil {
			return diag.Errorf("error reading log_quota: %v", err)
		}
	}

	if o.MaxLogFileSize != nil {
		v := *o.MaxLogFileSize

		if err = d.Set("max_log_file_size", v); err != nil {
			return diag.Errorf("error reading max_log_file_size: %v", err)
		}
	}

	if o.MaxPolicyPacketCaptureSize != nil {
		v := *o.MaxPolicyPacketCaptureSize

		if err = d.Set("max_policy_packet_capture_size", v); err != nil {
			return diag.Errorf("error reading max_policy_packet_capture_size: %v", err)
		}
	}

	if o.MaximumLogAge != nil {
		v := *o.MaximumLogAge

		if err = d.Set("maximum_log_age", v); err != nil {
			return diag.Errorf("error reading maximum_log_age: %v", err)
		}
	}

	if o.ReportQuota != nil {
		v := *o.ReportQuota

		if err = d.Set("report_quota", v); err != nil {
			return diag.Errorf("error reading report_quota: %v", err)
		}
	}

	if o.RollDay != nil {
		v := *o.RollDay

		if err = d.Set("roll_day", v); err != nil {
			return diag.Errorf("error reading roll_day: %v", err)
		}
	}

	if o.RollSchedule != nil {
		v := *o.RollSchedule

		if err = d.Set("roll_schedule", v); err != nil {
			return diag.Errorf("error reading roll_schedule: %v", err)
		}
	}

	if o.RollTime != nil {
		v := *o.RollTime

		if err = d.Set("roll_time", v); err != nil {
			return diag.Errorf("error reading roll_time: %v", err)
		}
	}

	if o.SourceIp != nil {
		v := *o.SourceIp

		if err = d.Set("source_ip", v); err != nil {
			return diag.Errorf("error reading source_ip: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Upload != nil {
		v := *o.Upload

		if err = d.Set("upload", v); err != nil {
			return diag.Errorf("error reading upload: %v", err)
		}
	}

	if o.UploadDeleteFiles != nil {
		v := *o.UploadDeleteFiles

		if err = d.Set("upload_delete_files", v); err != nil {
			return diag.Errorf("error reading upload_delete_files: %v", err)
		}
	}

	if o.UploadDestination != nil {
		v := *o.UploadDestination

		if err = d.Set("upload_destination", v); err != nil {
			return diag.Errorf("error reading upload_destination: %v", err)
		}
	}

	if o.UploadSslConn != nil {
		v := *o.UploadSslConn

		if err = d.Set("upload_ssl_conn", v); err != nil {
			return diag.Errorf("error reading upload_ssl_conn: %v", err)
		}
	}

	if o.Uploaddir != nil {
		v := *o.Uploaddir

		if err = d.Set("uploaddir", v); err != nil {
			return diag.Errorf("error reading uploaddir: %v", err)
		}
	}

	if o.Uploadip != nil {
		v := *o.Uploadip

		if err = d.Set("uploadip", v); err != nil {
			return diag.Errorf("error reading uploadip: %v", err)
		}
	}

	if o.Uploadpass != nil {
		v := *o.Uploadpass

		if err = d.Set("uploadpass", v); err != nil {
			return diag.Errorf("error reading uploadpass: %v", err)
		}
	}

	if o.Uploadport != nil {
		v := *o.Uploadport

		if err = d.Set("uploadport", v); err != nil {
			return diag.Errorf("error reading uploadport: %v", err)
		}
	}

	if o.Uploadsched != nil {
		v := *o.Uploadsched

		if err = d.Set("uploadsched", v); err != nil {
			return diag.Errorf("error reading uploadsched: %v", err)
		}
	}

	if o.Uploadtime != nil {
		v := *o.Uploadtime

		if err = d.Set("uploadtime", v); err != nil {
			return diag.Errorf("error reading uploadtime: %v", err)
		}
	}

	if o.Uploadtype != nil {
		v := *o.Uploadtype

		if err = d.Set("uploadtype", v); err != nil {
			return diag.Errorf("error reading uploadtype: %v", err)
		}
	}

	if o.Uploaduser != nil {
		v := *o.Uploaduser

		if err = d.Set("uploaduser", v); err != nil {
			return diag.Errorf("error reading uploaduser: %v", err)
		}
	}

	return nil
}

func getObjectLogDiskSetting(d *schema.ResourceData, sv string) (*models.LogDiskSetting, diag.Diagnostics) {
	obj := models.LogDiskSetting{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("diskfull"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("diskfull", sv)
				diags = append(diags, e)
			}
			obj.Diskfull = &v2
		}
	}
	if v1, ok := d.GetOk("dlp_archive_quota"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dlp_archive_quota", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DlpArchiveQuota = &tmp
		}
	}
	if v1, ok := d.GetOk("full_final_warning_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("full_final_warning_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FullFinalWarningThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("full_first_warning_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("full_first_warning_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FullFirstWarningThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("full_second_warning_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("full_second_warning_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FullSecondWarningThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("interface_select_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
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
	if v1, ok := d.GetOk("log_quota"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("log_quota", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LogQuota = &tmp
		}
	}
	if v1, ok := d.GetOk("max_log_file_size"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_log_file_size", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxLogFileSize = &tmp
		}
	}
	if v1, ok := d.GetOk("max_policy_packet_capture_size"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_policy_packet_capture_size", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxPolicyPacketCaptureSize = &tmp
		}
	}
	if v1, ok := d.GetOk("maximum_log_age"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("maximum_log_age", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaximumLogAge = &tmp
		}
	}
	if v1, ok := d.GetOk("report_quota"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("report_quota", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ReportQuota = &tmp
		}
	}
	if v1, ok := d.GetOk("roll_day"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("roll_day", sv)
				diags = append(diags, e)
			}
			obj.RollDay = &v2
		}
	}
	if v1, ok := d.GetOk("roll_schedule"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("roll_schedule", sv)
				diags = append(diags, e)
			}
			obj.RollSchedule = &v2
		}
	}
	if v1, ok := d.GetOk("roll_time"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("roll_time", sv)
				diags = append(diags, e)
			}
			obj.RollTime = &v2
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
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	if v1, ok := d.GetOk("upload"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("upload", sv)
				diags = append(diags, e)
			}
			obj.Upload = &v2
		}
	}
	if v1, ok := d.GetOk("upload_delete_files"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("upload_delete_files", sv)
				diags = append(diags, e)
			}
			obj.UploadDeleteFiles = &v2
		}
	}
	if v1, ok := d.GetOk("upload_destination"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("upload_destination", sv)
				diags = append(diags, e)
			}
			obj.UploadDestination = &v2
		}
	}
	if v1, ok := d.GetOk("upload_ssl_conn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("upload_ssl_conn", sv)
				diags = append(diags, e)
			}
			obj.UploadSslConn = &v2
		}
	}
	if v1, ok := d.GetOk("uploaddir"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("uploaddir", sv)
				diags = append(diags, e)
			}
			obj.Uploaddir = &v2
		}
	}
	if v1, ok := d.GetOk("uploadip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("uploadip", sv)
				diags = append(diags, e)
			}
			obj.Uploadip = &v2
		}
	}
	if v1, ok := d.GetOk("uploadpass"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("uploadpass", sv)
				diags = append(diags, e)
			}
			obj.Uploadpass = &v2
		}
	}
	if v1, ok := d.GetOk("uploadport"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("uploadport", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Uploadport = &tmp
		}
	}
	if v1, ok := d.GetOk("uploadsched"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("uploadsched", sv)
				diags = append(diags, e)
			}
			obj.Uploadsched = &v2
		}
	}
	if v1, ok := d.GetOk("uploadtime"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("uploadtime", sv)
				diags = append(diags, e)
			}
			obj.Uploadtime = &v2
		}
	}
	if v1, ok := d.GetOk("uploadtype"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("uploadtype", sv)
				diags = append(diags, e)
			}
			obj.Uploadtype = &v2
		}
	}
	if v1, ok := d.GetOk("uploaduser"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("uploaduser", sv)
				diags = append(diags, e)
			}
			obj.Uploaduser = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectLogDiskSetting(d *schema.ResourceData, sv string) (*models.LogDiskSetting, diag.Diagnostics) {
	obj := models.LogDiskSetting{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
