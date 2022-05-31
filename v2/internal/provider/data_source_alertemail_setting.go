// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure alert email settings.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceAlertemailSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Configure alert email settings.",

		ReadContext: dataSourceAlertemailSettingRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"fds_license_expiring_days": {
				Type:        schema.TypeInt,
				Description: "Number of days to send alert email prior to FortiGuard license expiration (1 - 100 days, default = 100).",
				Computed:    true,
			},
			"fds_license_expiring_warning": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiGuard license expiration warnings in alert email.",
				Computed:    true,
			},
			"fds_update_logs": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiGuard update logs in alert email.",
				Computed:    true,
			},
			"fips_cc_errors": {
				Type:        schema.TypeString,
				Description: "Enable/disable FIPS and Common Criteria error logs in alert email.",
				Computed:    true,
			},
			"fsso_disconnect_logs": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging of FSSO collector agent disconnect.",
				Computed:    true,
			},
			"ha_logs": {
				Type:        schema.TypeString,
				Description: "Enable/disable HA logs in alert email.",
				Computed:    true,
			},
			"ips_logs": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPS logs in alert email.",
				Computed:    true,
			},
			"ipsec_errors_logs": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPsec error logs in alert email.",
				Computed:    true,
			},
			"ppp_errors_logs": {
				Type:        schema.TypeString,
				Description: "Enable/disable PPP error logs in alert email.",
				Computed:    true,
			},
			"admin_login_logs": {
				Type:        schema.TypeString,
				Description: "Enable/disable administrator login/logout logs in alert email.",
				Computed:    true,
			},
			"alert_interval": {
				Type:        schema.TypeInt,
				Description: "Alert alert interval in minutes.",
				Computed:    true,
			},
			"amc_interface_bypass_mode": {
				Type:        schema.TypeString,
				Description: "Enable/disable Fortinet Advanced Mezzanine Card (AMC) interface bypass mode logs in alert email.",
				Computed:    true,
			},
			"antivirus_logs": {
				Type:        schema.TypeString,
				Description: "Enable/disable antivirus logs in alert email.",
				Computed:    true,
			},
			"configuration_changes_logs": {
				Type:        schema.TypeString,
				Description: "Enable/disable configuration change logs in alert email.",
				Computed:    true,
			},
			"critical_interval": {
				Type:        schema.TypeInt,
				Description: "Critical alert interval in minutes.",
				Computed:    true,
			},
			"debug_interval": {
				Type:        schema.TypeInt,
				Description: "Debug alert interval in minutes.",
				Computed:    true,
			},
			"email_interval": {
				Type:        schema.TypeInt,
				Description: "Interval between sending alert emails (1 - 99999 min, default = 5).",
				Computed:    true,
			},
			"emergency_interval": {
				Type:        schema.TypeInt,
				Description: "Emergency alert interval in minutes.",
				Computed:    true,
			},
			"error_interval": {
				Type:        schema.TypeInt,
				Description: "Error alert interval in minutes.",
				Computed:    true,
			},
			"filter_mode": {
				Type:        schema.TypeString,
				Description: "How to filter log messages that are sent to alert emails.",
				Computed:    true,
			},
			"firewall_authentication_failure_logs": {
				Type:        schema.TypeString,
				Description: "Enable/disable firewall authentication failure logs in alert email.",
				Computed:    true,
			},
			"fortiguard_log_quota_warning": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiCloud log quota warnings in alert email.",
				Computed:    true,
			},
			"information_interval": {
				Type:        schema.TypeInt,
				Description: "Information alert interval in minutes.",
				Computed:    true,
			},
			"local_disk_usage": {
				Type:        schema.TypeInt,
				Description: "Disk usage percentage at which to send alert email (1 - 99 percent, default = 75).",
				Computed:    true,
			},
			"log_disk_usage_warning": {
				Type:        schema.TypeString,
				Description: "Enable/disable disk usage warnings in alert email.",
				Computed:    true,
			},
			"mailto1": {
				Type:        schema.TypeString,
				Description: "Email address to send alert email to (usually a system administrator) (max. 63 characters).",
				Computed:    true,
			},
			"mailto2": {
				Type:        schema.TypeString,
				Description: "Optional second email address to send alert email to (max. 63 characters).",
				Computed:    true,
			},
			"mailto3": {
				Type:        schema.TypeString,
				Description: "Optional third email address to send alert email to (max. 63 characters).",
				Computed:    true,
			},
			"notification_interval": {
				Type:        schema.TypeInt,
				Description: "Notification alert interval in minutes.",
				Computed:    true,
			},
			"severity": {
				Type:        schema.TypeString,
				Description: "Lowest severity level to log.",
				Computed:    true,
			},
			"ssh_logs": {
				Type:        schema.TypeString,
				Description: "Enable/disable SSH logs in alert email.",
				Computed:    true,
			},
			"sslvpn_authentication_errors_logs": {
				Type:        schema.TypeString,
				Description: "Enable/disable SSL-VPN authentication error logs in alert email.",
				Computed:    true,
			},
			"username": {
				Type:        schema.TypeString,
				Description: "Name that appears in the From: field of alert emails (max. 63 characters).",
				Computed:    true,
			},
			"violation_traffic_logs": {
				Type:        schema.TypeString,
				Description: "Enable/disable violation traffic logs in alert email.",
				Computed:    true,
			},
			"warning_interval": {
				Type:        schema.TypeInt,
				Description: "Warning alert interval in minutes.",
				Computed:    true,
			},
			"webfilter_logs": {
				Type:        schema.TypeString,
				Description: "Enable/disable web filter logs in alert email.",
				Computed:    true,
			},
		},
	}
}

func dataSourceAlertemailSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "AlertemailSetting"

	o, err := c.Cmdb.ReadAlertemailSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading AlertemailSetting dataSource: %v", err)
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

	diags := refreshObjectAlertemailSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
