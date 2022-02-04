// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure alarm.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemAlarm() *schema.Resource {
	return &schema.Resource{
		Description: "Configure alarm.",

		ReadContext: dataSourceSystemAlarmRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"audible": {
				Type:        schema.TypeString,
				Description: "Enable/disable audible alarm.",
				Computed:    true,
			},
			"groups": {
				Type:        schema.TypeList,
				Description: "Alarm groups.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin_auth_failure_threshold": {
							Type:        schema.TypeInt,
							Description: "Admin authentication failure threshold.",
							Computed:    true,
						},
						"admin_auth_lockout_threshold": {
							Type:        schema.TypeInt,
							Description: "Admin authentication lockout threshold.",
							Computed:    true,
						},
						"decryption_failure_threshold": {
							Type:        schema.TypeInt,
							Description: "Decryption failure threshold.",
							Computed:    true,
						},
						"encryption_failure_threshold": {
							Type:        schema.TypeInt,
							Description: "Encryption failure threshold.",
							Computed:    true,
						},
						"fw_policy_id": {
							Type:        schema.TypeInt,
							Description: "Firewall policy ID.",
							Computed:    true,
						},
						"fw_policy_id_threshold": {
							Type:        schema.TypeInt,
							Description: "Firewall policy ID threshold.",
							Computed:    true,
						},
						"fw_policy_violations": {
							Type:        schema.TypeList,
							Description: "Firewall policy violations.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dst_ip": {
										Type:        schema.TypeString,
										Description: "Destination IP (0=all).",
										Computed:    true,
									},
									"dst_port": {
										Type:        schema.TypeInt,
										Description: "Destination port (0=all).",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "Firewall policy violations ID.",
										Computed:    true,
									},
									"src_ip": {
										Type:        schema.TypeString,
										Description: "Source IP (0=all).",
										Computed:    true,
									},
									"src_port": {
										Type:        schema.TypeInt,
										Description: "Source port (0=all).",
										Computed:    true,
									},
									"threshold": {
										Type:        schema.TypeInt,
										Description: "Firewall policy violation threshold.",
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Group ID.",
							Computed:    true,
						},
						"log_full_warning_threshold": {
							Type:        schema.TypeInt,
							Description: "Log full warning threshold.",
							Computed:    true,
						},
						"period": {
							Type:        schema.TypeInt,
							Description: "Time period in seconds (0 = from start up).",
							Computed:    true,
						},
						"replay_attempt_threshold": {
							Type:        schema.TypeInt,
							Description: "Replay attempt threshold.",
							Computed:    true,
						},
						"self_test_failure_threshold": {
							Type:        schema.TypeInt,
							Description: "Self-test failure threshold.",
							Computed:    true,
						},
						"user_auth_failure_threshold": {
							Type:        schema.TypeInt,
							Description: "User authentication failure threshold.",
							Computed:    true,
						},
						"user_auth_lockout_threshold": {
							Type:        schema.TypeInt,
							Description: "User authentication lockout threshold.",
							Computed:    true,
						},
					},
				},
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable alarm.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemAlarmRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemAlarm"

	o, err := c.Cmdb.ReadSystemAlarm(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemAlarm dataSource: %v", err)
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

	diags := refreshObjectSystemAlarm(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
