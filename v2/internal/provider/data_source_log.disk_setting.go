// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Settings for local disk logging.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceLogDiskSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Settings for local disk logging.",

		ReadContext: dataSourceLogDiskSettingRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"diskfull": {
				Type:        schema.TypeString,
				Description: "Action to take when disk is full. The system can overwrite the oldest log messages or stop logging when the disk is full (default = overwrite).",
				Computed:    true,
			},
			"dlp_archive_quota": {
				Type:        schema.TypeInt,
				Description: "DLP archive quota (MB).",
				Computed:    true,
			},
			"full_final_warning_threshold": {
				Type:        schema.TypeInt,
				Description: "Log full final warning threshold as a percent (3 - 100, default = 95).",
				Computed:    true,
			},
			"full_first_warning_threshold": {
				Type:        schema.TypeInt,
				Description: "Log full first warning threshold as a percent (1 - 98, default = 75).",
				Computed:    true,
			},
			"full_second_warning_threshold": {
				Type:        schema.TypeInt,
				Description: "Log full second warning threshold as a percent (2 - 99, default = 90).",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Specify outgoing interface to reach server.",
				Computed:    true,
			},
			"interface_select_method": {
				Type:        schema.TypeString,
				Description: "Specify how to select outgoing interface to reach server.",
				Computed:    true,
			},
			"ips_archive": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPS packet archiving to the local disk.",
				Computed:    true,
			},
			"log_quota": {
				Type:        schema.TypeInt,
				Description: "Disk log quota (MB).",
				Computed:    true,
			},
			"max_log_file_size": {
				Type:        schema.TypeInt,
				Description: "Maximum log file size before rolling (1 - 100 Mbytes).",
				Computed:    true,
			},
			"max_policy_packet_capture_size": {
				Type:        schema.TypeInt,
				Description: "Maximum size of policy sniffer in MB (0 means unlimited).",
				Computed:    true,
			},
			"maximum_log_age": {
				Type:        schema.TypeInt,
				Description: "Delete log files older than (days).",
				Computed:    true,
			},
			"report_quota": {
				Type:        schema.TypeInt,
				Description: "Report db quota (MB).",
				Computed:    true,
			},
			"roll_day": {
				Type:        schema.TypeString,
				Description: "Day of week on which to roll log file.",
				Computed:    true,
			},
			"roll_schedule": {
				Type:        schema.TypeString,
				Description: "Frequency to check log file for rolling.",
				Computed:    true,
			},
			"roll_time": {
				Type:        schema.TypeString,
				Description: "Time of day to roll the log file (hh:mm).",
				Computed:    true,
			},
			"source_ip": {
				Type:        schema.TypeString,
				Description: "Source IP address to use for uploading disk log files.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable local disk logging.",
				Computed:    true,
			},
			"upload": {
				Type:        schema.TypeString,
				Description: "Enable/disable uploading log files when they are rolled.",
				Computed:    true,
			},
			"upload_delete_files": {
				Type:        schema.TypeString,
				Description: "Delete log files after uploading (default = enable).",
				Computed:    true,
			},
			"upload_destination": {
				Type:        schema.TypeString,
				Description: "The type of server to upload log files to. Only FTP is currently supported.",
				Computed:    true,
			},
			"upload_ssl_conn": {
				Type:        schema.TypeString,
				Description: "Enable/disable encrypted FTPS communication to upload log files.",
				Computed:    true,
			},
			"uploaddir": {
				Type:        schema.TypeString,
				Description: "The remote directory on the FTP server to upload log files to.",
				Computed:    true,
			},
			"uploadip": {
				Type:        schema.TypeString,
				Description: "IP address of the FTP server to upload log files to.",
				Computed:    true,
			},
			"uploadpass": {
				Type:        schema.TypeString,
				Description: "Password required to log into the FTP server to upload disk log files.",
				Computed:    true,
				Sensitive:   true,
			},
			"uploadport": {
				Type:        schema.TypeInt,
				Description: "TCP port to use for communicating with the FTP server (default = 21).",
				Computed:    true,
			},
			"uploadsched": {
				Type:        schema.TypeString,
				Description: "Set the schedule for uploading log files to the FTP server (default = disable = upload when rolling).",
				Computed:    true,
			},
			"uploadtime": {
				Type:        schema.TypeString,
				Description: "Time of day at which log files are uploaded if uploadsched is enabled (hh:mm or hh).",
				Computed:    true,
			},
			"uploadtype": {
				Type:        schema.TypeString,
				Description: "Types of log files to upload. Separate multiple entries with a space.",
				Computed:    true,
			},
			"uploaduser": {
				Type:        schema.TypeString,
				Description: "Username required to log into the FTP server to upload disk log files.",
				Computed:    true,
			},
		},
	}
}

func dataSourceLogDiskSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "LogDiskSetting"

	o, err := c.Cmdb.ReadLogDiskSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogDiskSetting dataSource: %v", err)
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

	diags := refreshObjectLogDiskSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
