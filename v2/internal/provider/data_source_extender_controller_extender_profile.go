// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: FortiExtender extender profile configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceExtenderControllerExtenderProfile() *schema.Resource {
	return &schema.Resource{
		Description: "FortiExtender extender profile configuration.",

		ReadContext: dataSourceExtenderControllerExtenderProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"allowaccess": {
				Type:        schema.TypeString,
				Description: "Control management access to the managed extender. Separate entries with a space.",
				Computed:    true,
			},
			"bandwidth_limit": {
				Type:        schema.TypeInt,
				Description: "FortiExtender LAN extension bandwidth limit (Mbps).",
				Computed:    true,
			},
			"cellular": {
				Type:        schema.TypeList,
				Description: "FortiExtender cellular configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"dataplan": {
							Type:        schema.TypeList,
							Description: "Dataplan names.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Dataplan name.",
										Computed:    true,
									},
								},
							},
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
						"sms_notification": {
							Type:        schema.TypeList,
							Description: "FortiExtender cellular SMS notification configuration.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alert": {
										Type:        schema.TypeList,
										Description: "SMS alert list.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"data_exhausted": {
													Type:        schema.TypeString,
													Description: "Display string when data exhausted.",
													Computed:    true,
												},
												"fgt_backup_mode_switch": {
													Type:        schema.TypeString,
													Description: "Display string when FortiGate backup mode switched.",
													Computed:    true,
												},
												"low_signal_strength": {
													Type:        schema.TypeString,
													Description: "Display string when signal strength is low.",
													Computed:    true,
												},
												"mode_switch": {
													Type:        schema.TypeString,
													Description: "Display string when mode is switched.",
													Computed:    true,
												},
												"os_image_fallback": {
													Type:        schema.TypeString,
													Description: "Display string when falling back to a previous OS image.",
													Computed:    true,
												},
												"session_disconnect": {
													Type:        schema.TypeString,
													Description: "Display string when session disconnected.",
													Computed:    true,
												},
												"system_reboot": {
													Type:        schema.TypeString,
													Description: "Display string when system rebooted.",
													Computed:    true,
												},
											},
										},
									},
									"receiver": {
										Type:        schema.TypeList,
										Description: "SMS notification receiver list.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"alert": {
													Type:        schema.TypeString,
													Description: "Alert multi-options.",
													Computed:    true,
												},
												"name": {
													Type:        schema.TypeString,
													Description: "FortiExtender SMS notification receiver name.",
													Computed:    true,
												},
												"phone_number": {
													Type:        schema.TypeString,
													Description: "Receiver phone number. Format: [+][country code][area code][local phone number]. For example, +16501234567.",
													Computed:    true,
												},
												"status": {
													Type:        schema.TypeString,
													Description: "SMS notification receiver status.",
													Computed:    true,
												},
											},
										},
									},
									"status": {
										Type:        schema.TypeString,
										Description: "FortiExtender SMS notification status.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"enforce_bandwidth": {
				Type:        schema.TypeString,
				Description: "Enable/disable enforcement of bandwidth on LAN extension interface.",
				Computed:    true,
			},
			"extension": {
				Type:        schema.TypeString,
				Description: "Extension option.",
				Computed:    true,
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "ID.",
				Computed:    true,
			},
			"lan_extension": {
				Type:        schema.TypeList,
				Description: "FortiExtender lan extension configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backhaul": {
							Type:        schema.TypeList,
							Description: "LAN extension backhaul tunnel configuration.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "FortiExtender LAN extension backhaul name.",
										Computed:    true,
									},
									"port": {
										Type:        schema.TypeString,
										Description: "FortiExtender uplink port.",
										Computed:    true,
									},
									"role": {
										Type:        schema.TypeString,
										Description: "FortiExtender uplink port.",
										Computed:    true,
									},
									"weight": {
										Type:        schema.TypeInt,
										Description: "WRR weight parameter.",
										Computed:    true,
									},
								},
							},
						},
						"backhaul_interface": {
							Type:        schema.TypeString,
							Description: "IPsec phase1 interface.",
							Computed:    true,
						},
						"backhaul_ip": {
							Type:        schema.TypeString,
							Description: "IPsec phase1 IPv4/FQDN. Used to specify the external IP/FQDN when the FortiGate unit is behind a NAT device.",
							Computed:    true,
						},
						"ipsec_tunnel": {
							Type:        schema.TypeString,
							Description: "IPsec tunnel name.",
							Computed:    true,
						},
						"link_loadbalance": {
							Type:        schema.TypeString,
							Description: "LAN extension link load balance strategy.",
							Computed:    true,
						},
					},
				},
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
			"model": {
				Type:        schema.TypeString,
				Description: "Model.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "FortiExtender profile name.",
				Required:    true,
			},
		},
	}
}

func dataSourceExtenderControllerExtenderProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadExtenderControllerExtenderProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading ExtenderControllerExtenderProfile dataSource: %v", err)
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

	diags := refreshObjectExtenderControllerExtenderProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
