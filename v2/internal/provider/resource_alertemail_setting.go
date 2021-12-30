// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure alert email settings.

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

func resourceAlertemailSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Configure alert email settings.",

		CreateContext: resourceAlertemailSettingCreate,
		ReadContext:   resourceAlertemailSettingRead,
		UpdateContext: resourceAlertemailSettingUpdate,
		DeleteContext: resourceAlertemailSettingDelete,

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
			"fds_license_expiring_days": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),

				Description: "Number of days to send alert email prior to FortiGuard license expiration (1 - 100 days, default = 100).",
				Optional:    true,
				Computed:    true,
			},
			"fds_license_expiring_warning": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiGuard license expiration warnings in alert email.",
				Optional:    true,
				Computed:    true,
			},
			"fds_update_logs": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiGuard update logs in alert email.",
				Optional:    true,
				Computed:    true,
			},
			"fips_cc_errors": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FIPS and Common Criteria error logs in alert email.",
				Optional:    true,
				Computed:    true,
			},
			"fsso_disconnect_logs": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging of FSSO collector agent disconnect.",
				Optional:    true,
				Computed:    true,
			},
			"ha_logs": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable HA logs in alert email.",
				Optional:    true,
				Computed:    true,
			},
			"ips_logs": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPS logs in alert email.",
				Optional:    true,
				Computed:    true,
			},
			"ipsec_errors_logs": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPsec error logs in alert email.",
				Optional:    true,
				Computed:    true,
			},
			"ppp_errors_logs": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable PPP error logs in alert email.",
				Optional:    true,
				Computed:    true,
			},
			"admin_login_logs": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable administrator login/logout logs in alert email.",
				Optional:    true,
				Computed:    true,
			},
			"alert_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 99999),

				Description: "Alert alert interval in minutes.",
				Optional:    true,
				Computed:    true,
			},
			"amc_interface_bypass_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Fortinet Advanced Mezzanine Card (AMC) interface bypass mode logs in alert email.",
				Optional:    true,
				Computed:    true,
			},
			"antivirus_logs": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable antivirus logs in alert email.",
				Optional:    true,
				Computed:    true,
			},
			"configuration_changes_logs": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable configuration change logs in alert email.",
				Optional:    true,
				Computed:    true,
			},
			"critical_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 99999),

				Description: "Critical alert interval in minutes.",
				Optional:    true,
				Computed:    true,
			},
			"debug_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 99999),

				Description: "Debug alert interval in minutes.",
				Optional:    true,
				Computed:    true,
			},
			"email_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 99999),

				Description: "Interval between sending alert emails (1 - 99999 min, default = 5).",
				Optional:    true,
				Computed:    true,
			},
			"emergency_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 99999),

				Description: "Emergency alert interval in minutes.",
				Optional:    true,
				Computed:    true,
			},
			"error_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 99999),

				Description: "Error alert interval in minutes.",
				Optional:    true,
				Computed:    true,
			},
			"filter_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"category", "threshold"}, false),

				Description: "How to filter log messages that are sent to alert emails.",
				Optional:    true,
				Computed:    true,
			},
			"firewall_authentication_failure_logs": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable firewall authentication failure logs in alert email.",
				Optional:    true,
				Computed:    true,
			},
			"fortiguard_log_quota_warning": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiCloud log quota warnings in alert email.",
				Optional:    true,
				Computed:    true,
			},
			"information_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 99999),

				Description: "Information alert interval in minutes.",
				Optional:    true,
				Computed:    true,
			},
			"local_disk_usage": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 99),

				Description: "Disk usage percentage at which to send alert email (1 - 99 percent, default = 75).",
				Optional:    true,
				Computed:    true,
			},
			"log_disk_usage_warning": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable disk usage warnings in alert email.",
				Optional:    true,
				Computed:    true,
			},
			"mailto1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Email address to send alert email to (usually a system administrator) (max. 63 characters).",
				Optional:    true,
				Computed:    true,
			},
			"mailto2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Optional second email address to send alert email to (max. 63 characters).",
				Optional:    true,
				Computed:    true,
			},
			"mailto3": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Optional third email address to send alert email to (max. 63 characters).",
				Optional:    true,
				Computed:    true,
			},
			"notification_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 99999),

				Description: "Notification alert interval in minutes.",
				Optional:    true,
				Computed:    true,
			},
			"severity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"emergency", "alert", "critical", "error", "warning", "notification", "information", "debug"}, false),

				Description: "Lowest severity level to log.",
				Optional:    true,
				Computed:    true,
			},
			"ssh_logs": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SSH logs in alert email.",
				Optional:    true,
				Computed:    true,
			},
			"sslvpn_authentication_errors_logs": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SSL-VPN authentication error logs in alert email.",
				Optional:    true,
				Computed:    true,
			},
			"username": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Name that appears in the From: field of alert emails (max. 63 characters).",
				Optional:    true,
				Computed:    true,
			},
			"violation_traffic_logs": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable violation traffic logs in alert email.",
				Optional:    true,
				Computed:    true,
			},
			"warning_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 99999),

				Description: "Warning alert interval in minutes.",
				Optional:    true,
				Computed:    true,
			},
			"webfilter_logs": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable web filter logs in alert email.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceAlertemailSettingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectAlertemailSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateAlertemailSetting(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("AlertemailSetting")
	}

	return resourceAlertemailSettingRead(ctx, d, meta)
}

func resourceAlertemailSettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectAlertemailSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateAlertemailSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating AlertemailSetting resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("AlertemailSetting")
	}

	return resourceAlertemailSettingRead(ctx, d, meta)
}

func resourceAlertemailSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteAlertemailSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting AlertemailSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAlertemailSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadAlertemailSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading AlertemailSetting resource: %v", err)
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

	diags := refreshObjectAlertemailSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectAlertemailSetting(d *schema.ResourceData, o *models.AlertemailSetting, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.FDSLicenseExpiringDays != nil {
		v := *o.FDSLicenseExpiringDays

		if err = d.Set("fds_license_expiring_days", v); err != nil {
			return diag.Errorf("error reading fds_license_expiring_days: %v", err)
		}
	}

	if o.FDSLicenseExpiringWarning != nil {
		v := *o.FDSLicenseExpiringWarning

		if err = d.Set("fds_license_expiring_warning", v); err != nil {
			return diag.Errorf("error reading fds_license_expiring_warning: %v", err)
		}
	}

	if o.FDSUpdateLogs != nil {
		v := *o.FDSUpdateLogs

		if err = d.Set("fds_update_logs", v); err != nil {
			return diag.Errorf("error reading fds_update_logs: %v", err)
		}
	}

	if o.FIPSCCErrors != nil {
		v := *o.FIPSCCErrors

		if err = d.Set("fips_cc_errors", v); err != nil {
			return diag.Errorf("error reading fips_cc_errors: %v", err)
		}
	}

	if o.FSSODisconnectLogs != nil {
		v := *o.FSSODisconnectLogs

		if err = d.Set("fsso_disconnect_logs", v); err != nil {
			return diag.Errorf("error reading fsso_disconnect_logs: %v", err)
		}
	}

	if o.HALogs != nil {
		v := *o.HALogs

		if err = d.Set("ha_logs", v); err != nil {
			return diag.Errorf("error reading ha_logs: %v", err)
		}
	}

	if o.IPSLogs != nil {
		v := *o.IPSLogs

		if err = d.Set("ips_logs", v); err != nil {
			return diag.Errorf("error reading ips_logs: %v", err)
		}
	}

	if o.IPsecErrorsLogs != nil {
		v := *o.IPsecErrorsLogs

		if err = d.Set("ipsec_errors_logs", v); err != nil {
			return diag.Errorf("error reading ipsec_errors_logs: %v", err)
		}
	}

	if o.PPPErrorsLogs != nil {
		v := *o.PPPErrorsLogs

		if err = d.Set("ppp_errors_logs", v); err != nil {
			return diag.Errorf("error reading ppp_errors_logs: %v", err)
		}
	}

	if o.AdminLoginLogs != nil {
		v := *o.AdminLoginLogs

		if err = d.Set("admin_login_logs", v); err != nil {
			return diag.Errorf("error reading admin_login_logs: %v", err)
		}
	}

	if o.AlertInterval != nil {
		v := *o.AlertInterval

		if err = d.Set("alert_interval", v); err != nil {
			return diag.Errorf("error reading alert_interval: %v", err)
		}
	}

	if o.AmcInterfaceBypassMode != nil {
		v := *o.AmcInterfaceBypassMode

		if err = d.Set("amc_interface_bypass_mode", v); err != nil {
			return diag.Errorf("error reading amc_interface_bypass_mode: %v", err)
		}
	}

	if o.AntivirusLogs != nil {
		v := *o.AntivirusLogs

		if err = d.Set("antivirus_logs", v); err != nil {
			return diag.Errorf("error reading antivirus_logs: %v", err)
		}
	}

	if o.ConfigurationChangesLogs != nil {
		v := *o.ConfigurationChangesLogs

		if err = d.Set("configuration_changes_logs", v); err != nil {
			return diag.Errorf("error reading configuration_changes_logs: %v", err)
		}
	}

	if o.CriticalInterval != nil {
		v := *o.CriticalInterval

		if err = d.Set("critical_interval", v); err != nil {
			return diag.Errorf("error reading critical_interval: %v", err)
		}
	}

	if o.DebugInterval != nil {
		v := *o.DebugInterval

		if err = d.Set("debug_interval", v); err != nil {
			return diag.Errorf("error reading debug_interval: %v", err)
		}
	}

	if o.EmailInterval != nil {
		v := *o.EmailInterval

		if err = d.Set("email_interval", v); err != nil {
			return diag.Errorf("error reading email_interval: %v", err)
		}
	}

	if o.EmergencyInterval != nil {
		v := *o.EmergencyInterval

		if err = d.Set("emergency_interval", v); err != nil {
			return diag.Errorf("error reading emergency_interval: %v", err)
		}
	}

	if o.ErrorInterval != nil {
		v := *o.ErrorInterval

		if err = d.Set("error_interval", v); err != nil {
			return diag.Errorf("error reading error_interval: %v", err)
		}
	}

	if o.FilterMode != nil {
		v := *o.FilterMode

		if err = d.Set("filter_mode", v); err != nil {
			return diag.Errorf("error reading filter_mode: %v", err)
		}
	}

	if o.FirewallAuthenticationFailureLogs != nil {
		v := *o.FirewallAuthenticationFailureLogs

		if err = d.Set("firewall_authentication_failure_logs", v); err != nil {
			return diag.Errorf("error reading firewall_authentication_failure_logs: %v", err)
		}
	}

	if o.FortiguardLogQuotaWarning != nil {
		v := *o.FortiguardLogQuotaWarning

		if err = d.Set("fortiguard_log_quota_warning", v); err != nil {
			return diag.Errorf("error reading fortiguard_log_quota_warning: %v", err)
		}
	}

	if o.InformationInterval != nil {
		v := *o.InformationInterval

		if err = d.Set("information_interval", v); err != nil {
			return diag.Errorf("error reading information_interval: %v", err)
		}
	}

	if o.LocalDiskUsage != nil {
		v := *o.LocalDiskUsage

		if err = d.Set("local_disk_usage", v); err != nil {
			return diag.Errorf("error reading local_disk_usage: %v", err)
		}
	}

	if o.LogDiskUsageWarning != nil {
		v := *o.LogDiskUsageWarning

		if err = d.Set("log_disk_usage_warning", v); err != nil {
			return diag.Errorf("error reading log_disk_usage_warning: %v", err)
		}
	}

	if o.Mailto1 != nil {
		v := *o.Mailto1

		if err = d.Set("mailto1", v); err != nil {
			return diag.Errorf("error reading mailto1: %v", err)
		}
	}

	if o.Mailto2 != nil {
		v := *o.Mailto2

		if err = d.Set("mailto2", v); err != nil {
			return diag.Errorf("error reading mailto2: %v", err)
		}
	}

	if o.Mailto3 != nil {
		v := *o.Mailto3

		if err = d.Set("mailto3", v); err != nil {
			return diag.Errorf("error reading mailto3: %v", err)
		}
	}

	if o.NotificationInterval != nil {
		v := *o.NotificationInterval

		if err = d.Set("notification_interval", v); err != nil {
			return diag.Errorf("error reading notification_interval: %v", err)
		}
	}

	if o.Severity != nil {
		v := *o.Severity

		if err = d.Set("severity", v); err != nil {
			return diag.Errorf("error reading severity: %v", err)
		}
	}

	if o.SshLogs != nil {
		v := *o.SshLogs

		if err = d.Set("ssh_logs", v); err != nil {
			return diag.Errorf("error reading ssh_logs: %v", err)
		}
	}

	if o.SslvpnAuthenticationErrorsLogs != nil {
		v := *o.SslvpnAuthenticationErrorsLogs

		if err = d.Set("sslvpn_authentication_errors_logs", v); err != nil {
			return diag.Errorf("error reading sslvpn_authentication_errors_logs: %v", err)
		}
	}

	if o.Username != nil {
		v := *o.Username

		if err = d.Set("username", v); err != nil {
			return diag.Errorf("error reading username: %v", err)
		}
	}

	if o.ViolationTrafficLogs != nil {
		v := *o.ViolationTrafficLogs

		if err = d.Set("violation_traffic_logs", v); err != nil {
			return diag.Errorf("error reading violation_traffic_logs: %v", err)
		}
	}

	if o.WarningInterval != nil {
		v := *o.WarningInterval

		if err = d.Set("warning_interval", v); err != nil {
			return diag.Errorf("error reading warning_interval: %v", err)
		}
	}

	if o.WebfilterLogs != nil {
		v := *o.WebfilterLogs

		if err = d.Set("webfilter_logs", v); err != nil {
			return diag.Errorf("error reading webfilter_logs: %v", err)
		}
	}

	return nil
}

func getObjectAlertemailSetting(d *schema.ResourceData, sv string) (*models.AlertemailSetting, diag.Diagnostics) {
	obj := models.AlertemailSetting{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("fds_license_expiring_days"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fds_license_expiring_days", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FDSLicenseExpiringDays = &tmp
		}
	}
	if v1, ok := d.GetOk("fds_license_expiring_warning"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fds_license_expiring_warning", sv)
				diags = append(diags, e)
			}
			obj.FDSLicenseExpiringWarning = &v2
		}
	}
	if v1, ok := d.GetOk("fds_update_logs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fds_update_logs", sv)
				diags = append(diags, e)
			}
			obj.FDSUpdateLogs = &v2
		}
	}
	if v1, ok := d.GetOk("fips_cc_errors"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fips_cc_errors", sv)
				diags = append(diags, e)
			}
			obj.FIPSCCErrors = &v2
		}
	}
	if v1, ok := d.GetOk("fsso_disconnect_logs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fsso_disconnect_logs", sv)
				diags = append(diags, e)
			}
			obj.FSSODisconnectLogs = &v2
		}
	}
	if v1, ok := d.GetOk("ha_logs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ha_logs", sv)
				diags = append(diags, e)
			}
			obj.HALogs = &v2
		}
	}
	if v1, ok := d.GetOk("ips_logs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ips_logs", sv)
				diags = append(diags, e)
			}
			obj.IPSLogs = &v2
		}
	}
	if v1, ok := d.GetOk("ipsec_errors_logs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipsec_errors_logs", sv)
				diags = append(diags, e)
			}
			obj.IPsecErrorsLogs = &v2
		}
	}
	if v1, ok := d.GetOk("ppp_errors_logs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ppp_errors_logs", sv)
				diags = append(diags, e)
			}
			obj.PPPErrorsLogs = &v2
		}
	}
	if v1, ok := d.GetOk("admin_login_logs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admin_login_logs", sv)
				diags = append(diags, e)
			}
			obj.AdminLoginLogs = &v2
		}
	}
	if v1, ok := d.GetOk("alert_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("alert_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AlertInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("amc_interface_bypass_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("amc_interface_bypass_mode", sv)
				diags = append(diags, e)
			}
			obj.AmcInterfaceBypassMode = &v2
		}
	}
	if v1, ok := d.GetOk("antivirus_logs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("antivirus_logs", sv)
				diags = append(diags, e)
			}
			obj.AntivirusLogs = &v2
		}
	}
	if v1, ok := d.GetOk("configuration_changes_logs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("configuration_changes_logs", sv)
				diags = append(diags, e)
			}
			obj.ConfigurationChangesLogs = &v2
		}
	}
	if v1, ok := d.GetOk("critical_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("critical_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CriticalInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("debug_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("debug_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DebugInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("email_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("email_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EmailInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("emergency_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("emergency_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EmergencyInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("error_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("error_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ErrorInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("filter_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("filter_mode", sv)
				diags = append(diags, e)
			}
			obj.FilterMode = &v2
		}
	}
	if v1, ok := d.GetOk("firewall_authentication_failure_logs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("firewall_authentication_failure_logs", sv)
				diags = append(diags, e)
			}
			obj.FirewallAuthenticationFailureLogs = &v2
		}
	}
	if v1, ok := d.GetOk("fortiguard_log_quota_warning"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fortiguard_log_quota_warning", sv)
				diags = append(diags, e)
			}
			obj.FortiguardLogQuotaWarning = &v2
		}
	}
	if v1, ok := d.GetOk("information_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("information_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.InformationInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("local_disk_usage"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_disk_usage", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LocalDiskUsage = &tmp
		}
	}
	if v1, ok := d.GetOk("log_disk_usage_warning"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("log_disk_usage_warning", sv)
				diags = append(diags, e)
			}
			obj.LogDiskUsageWarning = &v2
		}
	}
	if v1, ok := d.GetOk("mailto1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mailto1", sv)
				diags = append(diags, e)
			}
			obj.Mailto1 = &v2
		}
	}
	if v1, ok := d.GetOk("mailto2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mailto2", sv)
				diags = append(diags, e)
			}
			obj.Mailto2 = &v2
		}
	}
	if v1, ok := d.GetOk("mailto3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mailto3", sv)
				diags = append(diags, e)
			}
			obj.Mailto3 = &v2
		}
	}
	if v1, ok := d.GetOk("notification_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("notification_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.NotificationInterval = &tmp
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
	if v1, ok := d.GetOk("ssh_logs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssh_logs", sv)
				diags = append(diags, e)
			}
			obj.SshLogs = &v2
		}
	}
	if v1, ok := d.GetOk("sslvpn_authentication_errors_logs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sslvpn_authentication_errors_logs", sv)
				diags = append(diags, e)
			}
			obj.SslvpnAuthenticationErrorsLogs = &v2
		}
	}
	if v1, ok := d.GetOk("username"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("username", sv)
				diags = append(diags, e)
			}
			obj.Username = &v2
		}
	}
	if v1, ok := d.GetOk("violation_traffic_logs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("violation_traffic_logs", sv)
				diags = append(diags, e)
			}
			obj.ViolationTrafficLogs = &v2
		}
	}
	if v1, ok := d.GetOk("warning_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("warning_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WarningInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("webfilter_logs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("webfilter_logs", sv)
				diags = append(diags, e)
			}
			obj.WebfilterLogs = &v2
		}
	}
	return &obj, diags
}
