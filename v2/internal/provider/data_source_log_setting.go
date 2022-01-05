// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure general log settings.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceLogSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Configure general log settings.",

		ReadContext: dataSourceLogSettingRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"anonymization_hash": {
				Type:        schema.TypeString,
				Description: "User name anonymization hash salt.",
				Computed:    true,
			},
			"brief_traffic_format": {
				Type:        schema.TypeString,
				Description: "Enable/disable brief format traffic logging.",
				Computed:    true,
			},
			"custom_log_fields": {
				Type:        schema.TypeList,
				Description: "Custom fields to append to all log messages.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field_id": {
							Type:        schema.TypeString,
							Description: "Custom log field.",
							Computed:    true,
						},
					},
				},
			},
			"daemon_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable daemon logging.",
				Computed:    true,
			},
			"expolicy_implicit_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable explicit proxy firewall implicit policy logging.",
				Computed:    true,
			},
			"faz_override": {
				Type:        schema.TypeString,
				Description: "Enable/disable override FortiAnalyzer settings.",
				Computed:    true,
			},
			"fortiview_weekly_data": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiView weekly data.",
				Computed:    true,
			},
			"fwpolicy_implicit_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable implicit firewall policy logging.",
				Computed:    true,
			},
			"fwpolicy6_implicit_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable implicit firewall policy6 logging.",
				Computed:    true,
			},
			"local_in_allow": {
				Type:        schema.TypeString,
				Description: "Enable/disable local-in-allow logging.",
				Computed:    true,
			},
			"local_in_deny_broadcast": {
				Type:        schema.TypeString,
				Description: "Enable/disable local-in-deny-broadcast logging.",
				Computed:    true,
			},
			"local_in_deny_unicast": {
				Type:        schema.TypeString,
				Description: "Enable/disable local-in-deny-unicast logging.",
				Computed:    true,
			},
			"local_out": {
				Type:        schema.TypeString,
				Description: "Enable/disable local-out logging.",
				Computed:    true,
			},
			"log_invalid_packet": {
				Type:        schema.TypeString,
				Description: "Enable/disable invalid packet traffic logging.",
				Computed:    true,
			},
			"log_policy_comment": {
				Type:        schema.TypeString,
				Description: "Enable/disable inserting policy comments into traffic logs.",
				Computed:    true,
			},
			"log_policy_name": {
				Type:        schema.TypeString,
				Description: "Enable/disable inserting policy name into traffic logs.",
				Computed:    true,
			},
			"log_user_in_upper": {
				Type:        schema.TypeString,
				Description: "Enable/disable logs with user-in-upper.",
				Computed:    true,
			},
			"neighbor_event": {
				Type:        schema.TypeString,
				Description: "Enable/disable neighbor event logging.",
				Computed:    true,
			},
			"resolve_ip": {
				Type:        schema.TypeString,
				Description: "Enable/disable adding resolved domain names to traffic logs if possible.",
				Computed:    true,
			},
			"resolve_port": {
				Type:        schema.TypeString,
				Description: "Enable/disable adding resolved service names to traffic logs.",
				Computed:    true,
			},
			"syslog_override": {
				Type:        schema.TypeString,
				Description: "Enable/disable override Syslog settings.",
				Computed:    true,
			},
			"user_anonymize": {
				Type:        schema.TypeString,
				Description: "Enable/disable anonymizing user names in log messages.",
				Computed:    true,
			},
		},
	}
}

func dataSourceLogSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "LogSetting"

	o, err := c.Cmdb.ReadLogSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogSetting dataSource: %v", err)
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

	diags := refreshObjectLogSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
