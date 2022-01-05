// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure access profiles for system administrators.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemAccprofile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure access profiles for system administrators.",

		ReadContext: dataSourceSystemAccprofileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"admintimeout": {
				Type:        schema.TypeInt,
				Description: "Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).",
				Computed:    true,
			},
			"admintimeout_override": {
				Type:        schema.TypeString,
				Description: "Enable/disable overriding the global administrator idle timeout.",
				Computed:    true,
			},
			"authgrp": {
				Type:        schema.TypeString,
				Description: "Administrator access to Users and Devices.",
				Computed:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"ftviewgrp": {
				Type:        schema.TypeString,
				Description: "FortiView.",
				Computed:    true,
			},
			"fwgrp": {
				Type:        schema.TypeString,
				Description: "Administrator access to the Firewall configuration.",
				Computed:    true,
			},
			"fwgrp_permission": {
				Type:        schema.TypeList,
				Description: "Custom firewall permission.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Description: "Address Configuration.",
							Computed:    true,
						},
						"others": {
							Type:        schema.TypeString,
							Description: "Other Firewall Configuration.",
							Computed:    true,
						},
						"policy": {
							Type:        schema.TypeString,
							Description: "Policy Configuration.",
							Computed:    true,
						},
						"schedule": {
							Type:        schema.TypeString,
							Description: "Schedule Configuration.",
							Computed:    true,
						},
						"service": {
							Type:        schema.TypeString,
							Description: "Service Configuration.",
							Computed:    true,
						},
					},
				},
			},
			"loggrp": {
				Type:        schema.TypeString,
				Description: "Administrator access to Logging and Reporting including viewing log messages.",
				Computed:    true,
			},
			"loggrp_permission": {
				Type:        schema.TypeList,
				Description: "Custom Log & Report permission.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config": {
							Type:        schema.TypeString,
							Description: "Log & Report configuration.",
							Computed:    true,
						},
						"data_access": {
							Type:        schema.TypeString,
							Description: "Log & Report Data Access.",
							Computed:    true,
						},
						"report_access": {
							Type:        schema.TypeString,
							Description: "Log & Report Report Access.",
							Computed:    true,
						},
						"threat_weight": {
							Type:        schema.TypeString,
							Description: "Log & Report Threat Weight.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Profile name.",
				Required:    true,
			},
			"netgrp": {
				Type:        schema.TypeString,
				Description: "Network Configuration.",
				Computed:    true,
			},
			"netgrp_permission": {
				Type:        schema.TypeList,
				Description: "Custom network permission.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cfg": {
							Type:        schema.TypeString,
							Description: "Network Configuration.",
							Computed:    true,
						},
						"packet_capture": {
							Type:        schema.TypeString,
							Description: "Packet Capture Configuration.",
							Computed:    true,
						},
						"route_cfg": {
							Type:        schema.TypeString,
							Description: "Router Configuration.",
							Computed:    true,
						},
					},
				},
			},
			"scope": {
				Type:        schema.TypeString,
				Description: "Scope of admin access: global or specific VDOM(s).",
				Computed:    true,
			},
			"secfabgrp": {
				Type:        schema.TypeString,
				Description: "Security Fabric.",
				Computed:    true,
			},
			"sysgrp": {
				Type:        schema.TypeString,
				Description: "System Configuration.",
				Computed:    true,
			},
			"sysgrp_permission": {
				Type:        schema.TypeList,
				Description: "Custom system permission.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin": {
							Type:        schema.TypeString,
							Description: "Administrator Users.",
							Computed:    true,
						},
						"cfg": {
							Type:        schema.TypeString,
							Description: "System Configuration.",
							Computed:    true,
						},
						"mnt": {
							Type:        schema.TypeString,
							Description: "Maintenance.",
							Computed:    true,
						},
						"upd": {
							Type:        schema.TypeString,
							Description: "FortiGuard Updates.",
							Computed:    true,
						},
					},
				},
			},
			"system_diagnostics": {
				Type:        schema.TypeString,
				Description: "Enable/disable permission to run system diagnostic commands.",
				Computed:    true,
			},
			"utmgrp": {
				Type:        schema.TypeString,
				Description: "Administrator access to Security Profiles.",
				Computed:    true,
			},
			"utmgrp_permission": {
				Type:        schema.TypeList,
				Description: "Custom Security Profile permissions.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"antivirus": {
							Type:        schema.TypeString,
							Description: "Antivirus profiles and settings.",
							Computed:    true,
						},
						"application_control": {
							Type:        schema.TypeString,
							Description: "Application Control profiles and settings.",
							Computed:    true,
						},
						"data_loss_prevention": {
							Type:        schema.TypeString,
							Description: "DLP profiles and settings.",
							Computed:    true,
						},
						"dnsfilter": {
							Type:        schema.TypeString,
							Description: "DNS Filter profiles and settings.",
							Computed:    true,
						},
						"emailfilter": {
							Type:        schema.TypeString,
							Description: "Email Filter and settings.",
							Computed:    true,
						},
						"endpoint_control": {
							Type:        schema.TypeString,
							Description: "FortiClient Profiles.",
							Computed:    true,
						},
						"file_filter": {
							Type:        schema.TypeString,
							Description: "File-filter profiles and settings.",
							Computed:    true,
						},
						"icap": {
							Type:        schema.TypeString,
							Description: "ICAP profiles and settings.",
							Computed:    true,
						},
						"ips": {
							Type:        schema.TypeString,
							Description: "IPS profiles and settings.",
							Computed:    true,
						},
						"voip": {
							Type:        schema.TypeString,
							Description: "VoIP profiles and settings.",
							Computed:    true,
						},
						"waf": {
							Type:        schema.TypeString,
							Description: "Web Application Firewall profiles and settings.",
							Computed:    true,
						},
						"webfilter": {
							Type:        schema.TypeString,
							Description: "Web Filter profiles and settings.",
							Computed:    true,
						},
					},
				},
			},
			"vpngrp": {
				Type:        schema.TypeString,
				Description: "Administrator access to IPsec, SSL, PPTP, and L2TP VPN.",
				Computed:    true,
			},
			"wanoptgrp": {
				Type:        schema.TypeString,
				Description: "Administrator access to WAN Opt & Cache.",
				Computed:    true,
			},
			"wifi": {
				Type:        schema.TypeString,
				Description: "Administrator access to the WiFi controller and Switch controller.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemAccprofileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemAccprofile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemAccprofile dataSource: %v", err)
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

	diags := refreshObjectSystemAccprofile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
