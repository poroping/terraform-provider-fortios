// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Extender controller configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceExtenderControllerExtender() *schema.Resource {
	return &schema.Resource{
		Description: "Extender controller configuration.",

		ReadContext: dataSourceExtenderControllerExtenderRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"aaa_shared_secret": {
				Type:        schema.TypeString,
				Description: "AAA shared secret.",
				Computed:    true,
				Sensitive:   true,
			},
			"access_point_name": {
				Type:        schema.TypeString,
				Description: "Access point name(APN).",
				Computed:    true,
			},
			"admin": {
				Type:        schema.TypeString,
				Description: "FortiExtender Administration (enable or disable).",
				Computed:    true,
			},
			"allowaccess": {
				Type:        schema.TypeString,
				Description: "Control management access to the managed extender. Separate entries with a space.",
				Computed:    true,
			},
			"at_dial_script": {
				Type:        schema.TypeString,
				Description: "Initialization AT commands specific to the MODEM.",
				Computed:    true,
			},
			"authorized": {
				Type:        schema.TypeString,
				Description: "FortiExtender Administration (enable or disable).",
				Computed:    true,
			},
			"bandwidth_limit": {
				Type:        schema.TypeInt,
				Description: "FortiExtender LAN extension bandwidth limit (Mbps).",
				Computed:    true,
			},
			"billing_start_day": {
				Type:        schema.TypeInt,
				Description: "Billing start day.",
				Computed:    true,
			},
			"cdma_aaa_spi": {
				Type:        schema.TypeString,
				Description: "CDMA AAA SPI.",
				Computed:    true,
			},
			"cdma_ha_spi": {
				Type:        schema.TypeString,
				Description: "CDMA HA SPI.",
				Computed:    true,
			},
			"cdma_nai": {
				Type:        schema.TypeString,
				Description: "NAI for CDMA MODEMS.",
				Computed:    true,
			},
			"conn_status": {
				Type:        schema.TypeInt,
				Description: "Connection status.",
				Computed:    true,
			},
			"controller_report": {
				Type:        schema.TypeList,
				Description: "FortiExtender controller report configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interval": {
							Type:        schema.TypeInt,
							Description: "Controller report interval.",
							Computed:    true,
						},
						"signal_threshold": {
							Type:        schema.TypeInt,
							Description: "Controller report signal threshold.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "FortiExtender controller report status.",
							Computed:    true,
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description.",
				Computed:    true,
			},
			"device_id": {
				Type:        schema.TypeInt,
				Description: "Device ID.",
				Computed:    true,
			},
			"dial_mode": {
				Type:        schema.TypeString,
				Description: "Dial mode (dial-on-demand or always-connect).",
				Computed:    true,
			},
			"dial_status": {
				Type:        schema.TypeInt,
				Description: "Dial status.",
				Computed:    true,
			},
			"enforce_bandwidth": {
				Type:        schema.TypeString,
				Description: "Enable/disable enforcement of bandwidth on LAN extension interface.",
				Computed:    true,
			},
			"ext_name": {
				Type:        schema.TypeString,
				Description: "FortiExtender name.",
				Computed:    true,
			},
			"extension_type": {
				Type:        schema.TypeString,
				Description: "Extension type for this FortiExtender.",
				Computed:    true,
			},
			"ha_shared_secret": {
				Type:        schema.TypeString,
				Description: "HA shared secret.",
				Computed:    true,
				Sensitive:   true,
			},
			"fosid": {
				Type:        schema.TypeString,
				Description: "FortiExtender serial number.",
				Computed:    true,
			},
			"ifname": {
				Type:        schema.TypeString,
				Description: "FortiExtender interface name.",
				Computed:    true,
			},
			"initiated_update": {
				Type:        schema.TypeString,
				Description: "Allow/disallow network initiated updates to the MODEM.",
				Computed:    true,
			},
			"login_password": {
				Type:        schema.TypeString,
				Description: "Set the managed extender's administrator password.",
				Computed:    true,
				Sensitive:   true,
			},
			"login_password_change": {
				Type:        schema.TypeString,
				Description: "Change or reset the administrator password of a managed extender (yes, default, or no, default = no).",
				Computed:    true,
			},
			"mode": {
				Type:        schema.TypeString,
				Description: "FortiExtender mode.",
				Computed:    true,
			},
			"modem_passwd": {
				Type:        schema.TypeString,
				Description: "MODEM password.",
				Computed:    true,
				Sensitive:   true,
			},
			"modem_type": {
				Type:        schema.TypeString,
				Description: "MODEM type (CDMA, GSM/LTE or WIMAX).",
				Computed:    true,
			},
			"modem1": {
				Type:        schema.TypeList,
				Description: "Configuration options for modem 1.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_switch": {
							Type:        schema.TypeList,
							Description: "FortiExtender auto switch configuration.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dataplan": {
										Type:        schema.TypeString,
										Description: "Automatically switch based on data usage.",
										Computed:    true,
									},
									"disconnect": {
										Type:        schema.TypeString,
										Description: "Auto switch by disconnect.",
										Computed:    true,
									},
									"disconnect_period": {
										Type:        schema.TypeInt,
										Description: "Automatically switch based on disconnect period.",
										Computed:    true,
									},
									"disconnect_threshold": {
										Type:        schema.TypeInt,
										Description: "Automatically switch based on disconnect threshold.",
										Computed:    true,
									},
									"signal": {
										Type:        schema.TypeString,
										Description: "Automatically switch based on signal strength.",
										Computed:    true,
									},
									"switch_back": {
										Type:        schema.TypeString,
										Description: "Auto switch with switch back multi-options.",
										Computed:    true,
									},
									"switch_back_time": {
										Type:        schema.TypeString,
										Description: "Automatically switch over to preferred SIM/carrier at a specified time in UTC (HH:MM).",
										Computed:    true,
									},
									"switch_back_timer": {
										Type:        schema.TypeInt,
										Description: "Automatically switch over to preferred SIM/carrier after the given time (3600 - 2147483647 sec).",
										Computed:    true,
									},
								},
							},
						},
						"conn_status": {
							Type:        schema.TypeInt,
							Description: "Connection status.",
							Computed:    true,
						},
						"default_sim": {
							Type:        schema.TypeString,
							Description: "Default SIM selection.",
							Computed:    true,
						},
						"gps": {
							Type:        schema.TypeString,
							Description: "FortiExtender GPS enable/disable.",
							Computed:    true,
						},
						"ifname": {
							Type:        schema.TypeString,
							Description: "FortiExtender interface name.",
							Computed:    true,
						},
						"preferred_carrier": {
							Type:        schema.TypeString,
							Description: "Preferred carrier.",
							Computed:    true,
						},
						"redundant_intf": {
							Type:        schema.TypeString,
							Description: "Redundant interface.",
							Computed:    true,
						},
						"redundant_mode": {
							Type:        schema.TypeString,
							Description: "FortiExtender mode.",
							Computed:    true,
						},
						"sim1_pin": {
							Type:        schema.TypeString,
							Description: "SIM #1 PIN status.",
							Computed:    true,
						},
						"sim1_pin_code": {
							Type:        schema.TypeString,
							Description: "SIM #1 PIN password.",
							Computed:    true,
							Sensitive:   true,
						},
						"sim2_pin": {
							Type:        schema.TypeString,
							Description: "SIM #2 PIN status.",
							Computed:    true,
						},
						"sim2_pin_code": {
							Type:        schema.TypeString,
							Description: "SIM #2 PIN password.",
							Computed:    true,
							Sensitive:   true,
						},
					},
				},
			},
			"modem2": {
				Type:        schema.TypeList,
				Description: "Configuration options for modem 2.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_switch": {
							Type:        schema.TypeList,
							Description: "FortiExtender auto switch configuration.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dataplan": {
										Type:        schema.TypeString,
										Description: "Automatically switch based on data usage.",
										Computed:    true,
									},
									"disconnect": {
										Type:        schema.TypeString,
										Description: "Auto switch by disconnect.",
										Computed:    true,
									},
									"disconnect_period": {
										Type:        schema.TypeInt,
										Description: "Automatically switch based on disconnect period.",
										Computed:    true,
									},
									"disconnect_threshold": {
										Type:        schema.TypeInt,
										Description: "Automatically switch based on disconnect threshold.",
										Computed:    true,
									},
									"signal": {
										Type:        schema.TypeString,
										Description: "Automatically switch based on signal strength.",
										Computed:    true,
									},
									"switch_back": {
										Type:        schema.TypeString,
										Description: "Auto switch with switch back multi-options.",
										Computed:    true,
									},
									"switch_back_time": {
										Type:        schema.TypeString,
										Description: "Automatically switch over to preferred SIM/carrier at a specified time in UTC (HH:MM).",
										Computed:    true,
									},
									"switch_back_timer": {
										Type:        schema.TypeInt,
										Description: "Automatically switch over to preferred SIM/carrier after the given time (3600 - 2147483647 sec).",
										Computed:    true,
									},
								},
							},
						},
						"conn_status": {
							Type:        schema.TypeInt,
							Description: "Connection status.",
							Computed:    true,
						},
						"default_sim": {
							Type:        schema.TypeString,
							Description: "Default SIM selection.",
							Computed:    true,
						},
						"gps": {
							Type:        schema.TypeString,
							Description: "FortiExtender GPS enable/disable.",
							Computed:    true,
						},
						"ifname": {
							Type:        schema.TypeString,
							Description: "FortiExtender interface name.",
							Computed:    true,
						},
						"preferred_carrier": {
							Type:        schema.TypeString,
							Description: "Preferred carrier.",
							Computed:    true,
						},
						"redundant_intf": {
							Type:        schema.TypeString,
							Description: "Redundant interface.",
							Computed:    true,
						},
						"redundant_mode": {
							Type:        schema.TypeString,
							Description: "FortiExtender mode.",
							Computed:    true,
						},
						"sim1_pin": {
							Type:        schema.TypeString,
							Description: "SIM #1 PIN status.",
							Computed:    true,
						},
						"sim1_pin_code": {
							Type:        schema.TypeString,
							Description: "SIM #1 PIN password.",
							Computed:    true,
							Sensitive:   true,
						},
						"sim2_pin": {
							Type:        schema.TypeString,
							Description: "SIM #2 PIN status.",
							Computed:    true,
						},
						"sim2_pin_code": {
							Type:        schema.TypeString,
							Description: "SIM #2 PIN password.",
							Computed:    true,
							Sensitive:   true,
						},
					},
				},
			},
			"multi_mode": {
				Type:        schema.TypeString,
				Description: "MODEM mode of operation(3G,LTE,etc).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "FortiExtender entry name.",
				Required:    true,
			},
			"override_allowaccess": {
				Type:        schema.TypeString,
				Description: "Enable to override the extender profile management access configuration.",
				Computed:    true,
			},
			"override_enforce_bandwidth": {
				Type:        schema.TypeString,
				Description: "Enable to override the extender profile enforce-bandwidth setting.",
				Computed:    true,
			},
			"override_login_password_change": {
				Type:        schema.TypeString,
				Description: "Enable to override the extender profile login-password (administrator password) setting.",
				Computed:    true,
			},
			"ppp_auth_protocol": {
				Type:        schema.TypeString,
				Description: "PPP authentication protocol (PAP,CHAP or auto).",
				Computed:    true,
			},
			"ppp_echo_request": {
				Type:        schema.TypeString,
				Description: "Enable/disable PPP echo request.",
				Computed:    true,
			},
			"ppp_password": {
				Type:        schema.TypeString,
				Description: "PPP password.",
				Computed:    true,
				Sensitive:   true,
			},
			"ppp_username": {
				Type:        schema.TypeString,
				Description: "PPP username.",
				Computed:    true,
			},
			"primary_ha": {
				Type:        schema.TypeString,
				Description: "Primary HA.",
				Computed:    true,
			},
			"profile": {
				Type:        schema.TypeString,
				Description: "FortiExtender profile configuration.",
				Computed:    true,
			},
			"quota_limit_mb": {
				Type:        schema.TypeInt,
				Description: "Monthly quota limit (MB).",
				Computed:    true,
			},
			"redial": {
				Type:        schema.TypeString,
				Description: "Number of redials allowed based on failed attempts.",
				Computed:    true,
			},
			"redundant_intf": {
				Type:        schema.TypeString,
				Description: "Redundant interface.",
				Computed:    true,
			},
			"roaming": {
				Type:        schema.TypeString,
				Description: "Enable/disable MODEM roaming.",
				Computed:    true,
			},
			"role": {
				Type:        schema.TypeString,
				Description: "FortiExtender work role(Primary, Secondary, None).",
				Computed:    true,
			},
			"secondary_ha": {
				Type:        schema.TypeString,
				Description: "Secondary HA.",
				Computed:    true,
			},
			"sim_pin": {
				Type:        schema.TypeString,
				Description: "SIM PIN.",
				Computed:    true,
				Sensitive:   true,
			},
			"vdom": {
				Type:        schema.TypeInt,
				Description: "VDOM.",
				Computed:    true,
			},
			"wan_extension": {
				Type:        schema.TypeList,
				Description: "FortiExtender wan extension configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"modem1_extension": {
							Type:        schema.TypeString,
							Description: "FortiExtender interface name.",
							Computed:    true,
						},
						"modem2_extension": {
							Type:        schema.TypeString,
							Description: "FortiExtender interface name.",
							Computed:    true,
						},
					},
				},
			},
			"wimax_auth_protocol": {
				Type:        schema.TypeString,
				Description: "WiMax authentication protocol(TLS or TTLS).",
				Computed:    true,
			},
			"wimax_carrier": {
				Type:        schema.TypeString,
				Description: "WiMax carrier.",
				Computed:    true,
			},
			"wimax_realm": {
				Type:        schema.TypeString,
				Description: "WiMax realm.",
				Computed:    true,
			},
		},
	}
}

func dataSourceExtenderControllerExtenderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("name")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadExtenderControllerExtender(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading ExtenderControllerExtender dataSource: %v", err)
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

	diags := refreshObjectExtenderControllerExtender(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
