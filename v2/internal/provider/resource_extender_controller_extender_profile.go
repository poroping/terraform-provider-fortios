// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: FortiExtender extender profile configuration.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceExtenderControllerExtenderProfile() *schema.Resource {
	return &schema.Resource{
		Description: "FortiExtender extender profile configuration.",

		CreateContext: resourceExtenderControllerExtenderProfileCreate,
		ReadContext:   resourceExtenderControllerExtenderProfileRead,
		UpdateContext: resourceExtenderControllerExtenderProfileUpdate,
		DeleteContext: resourceExtenderControllerExtenderProfileDelete,

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
			"allow_append": {
				Type:        schema.TypeBool,
				Description: "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"allowaccess": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Control management access to the managed extender. Separate entries with a space.",
				Optional:         true,
				Computed:         true,
			},
			"bandwidth_limit": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16776000),

				Description: "FortiExtender LAN extension bandwidth limit (Mbps).",
				Optional:    true,
				Computed:    true,
			},
			"cellular": {
				Type:        schema.TypeList,
				Description: "FortiExtender cellular configuration.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"controller_report": {
							Type:        schema.TypeList,
							Description: "FortiExtender controller report configuration.",
							Optional:    true, MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interval": {
										Type: schema.TypeInt,

										Description: "Controller report interval.",
										Optional:    true,
										Computed:    true,
									},
									"signal_threshold": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(10, 50),

										Description: "Controller report signal threshold.",
										Optional:    true,
										Computed:    true,
									},
									"status": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

										Description: "FortiExtender controller report status.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"dataplan": {
							Type:        schema.TypeList,
							Description: "Dataplan names.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Dataplan name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"modem1": {
							Type:        schema.TypeList,
							Description: "Configuration options for modem 1.",
							Optional:    true, MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auto_switch": {
										Type:        schema.TypeList,
										Description: "FortiExtender auto switch configuration.",
										Optional:    true, MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dataplan": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

													Description: "Automatically switch based on data usage.",
													Optional:    true,
													Computed:    true,
												},
												"disconnect": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

													Description: "Auto switch by disconnect.",
													Optional:    true,
													Computed:    true,
												},
												"disconnect_period": {
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntBetween(600, 18000),

													Description: "Automatically switch based on disconnect period.",
													Optional:    true,
													Computed:    true,
												},
												"disconnect_threshold": {
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntBetween(1, 100),

													Description: "Automatically switch based on disconnect threshold.",
													Optional:    true,
													Computed:    true,
												},
												"signal": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

													Description: "Automatically switch based on signal strength.",
													Optional:    true,
													Computed:    true,
												},
												"switch_back": {
													Type: schema.TypeString,

													DiffSuppressFunc: suppressors.DiffFakeListEqual,
													Description:      "Auto switch with switch back multi-options.",
													Optional:         true,
													Computed:         true,
												},
												"switch_back_time": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 31),

													Description: "Automatically switch over to preferred SIM/carrier at a specified time in UTC (HH:MM).",
													Optional:    true,
													Computed:    true,
												},
												"switch_back_timer": {
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntBetween(3600, 2147483647),

													Description: "Automatically switch over to preferred SIM/carrier after the given time (3600 - 2147483647 sec).",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"conn_status": {
										Type: schema.TypeInt,

										Description: "Connection status.",
										Optional:    true,
										Computed:    true,
									},
									"default_sim": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"sim1", "sim2", "carrier", "cost"}, false),

										Description: "Default SIM selection.",
										Optional:    true,
										Computed:    true,
									},
									"gps": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

										Description: "FortiExtender GPS enable/disable.",
										Optional:    true,
										Computed:    true,
									},
									"preferred_carrier": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),

										Description: "Preferred carrier.",
										Optional:    true,
										Computed:    true,
									},
									"redundant_intf": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 15),

										Description: "Redundant interface.",
										Optional:    true,
										Computed:    true,
									},
									"redundant_mode": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

										Description: "FortiExtender mode.",
										Optional:    true,
										Computed:    true,
									},
									"sim1_pin": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

										Description: "SIM #1 PIN status.",
										Optional:    true,
										Computed:    true,
									},
									"sim1_pin_code": {
										Type: schema.TypeString,

										Description: "SIM #1 PIN password.",
										Optional:    true,
										Computed:    true,
										Sensitive:   true,
									},
									"sim2_pin": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

										Description: "SIM #2 PIN status.",
										Optional:    true,
										Computed:    true,
									},
									"sim2_pin_code": {
										Type: schema.TypeString,

										Description: "SIM #2 PIN password.",
										Optional:    true,
										Computed:    true,
										Sensitive:   true,
									},
								},
							},
						},
						"modem2": {
							Type:        schema.TypeList,
							Description: "Configuration options for modem 2.",
							Optional:    true, MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auto_switch": {
										Type:        schema.TypeList,
										Description: "FortiExtender auto switch configuration.",
										Optional:    true, MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dataplan": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

													Description: "Automatically switch based on data usage.",
													Optional:    true,
													Computed:    true,
												},
												"disconnect": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

													Description: "Auto switch by disconnect.",
													Optional:    true,
													Computed:    true,
												},
												"disconnect_period": {
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntBetween(600, 18000),

													Description: "Automatically switch based on disconnect period.",
													Optional:    true,
													Computed:    true,
												},
												"disconnect_threshold": {
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntBetween(1, 100),

													Description: "Automatically switch based on disconnect threshold.",
													Optional:    true,
													Computed:    true,
												},
												"signal": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

													Description: "Automatically switch based on signal strength.",
													Optional:    true,
													Computed:    true,
												},
												"switch_back": {
													Type: schema.TypeString,

													DiffSuppressFunc: suppressors.DiffFakeListEqual,
													Description:      "Auto switch with switch back multi-options.",
													Optional:         true,
													Computed:         true,
												},
												"switch_back_time": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 31),

													Description: "Automatically switch over to preferred SIM/carrier at a specified time in UTC (HH:MM).",
													Optional:    true,
													Computed:    true,
												},
												"switch_back_timer": {
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntBetween(3600, 2147483647),

													Description: "Automatically switch over to preferred SIM/carrier after the given time (3600 - 2147483647 sec).",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"conn_status": {
										Type: schema.TypeInt,

										Description: "Connection status.",
										Optional:    true,
										Computed:    true,
									},
									"default_sim": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"sim1", "sim2", "carrier", "cost"}, false),

										Description: "Default SIM selection.",
										Optional:    true,
										Computed:    true,
									},
									"gps": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

										Description: "FortiExtender GPS enable/disable.",
										Optional:    true,
										Computed:    true,
									},
									"preferred_carrier": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),

										Description: "Preferred carrier.",
										Optional:    true,
										Computed:    true,
									},
									"redundant_intf": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 15),

										Description: "Redundant interface.",
										Optional:    true,
										Computed:    true,
									},
									"redundant_mode": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

										Description: "FortiExtender mode.",
										Optional:    true,
										Computed:    true,
									},
									"sim1_pin": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

										Description: "SIM #1 PIN status.",
										Optional:    true,
										Computed:    true,
									},
									"sim1_pin_code": {
										Type: schema.TypeString,

										Description: "SIM #1 PIN password.",
										Optional:    true,
										Computed:    true,
										Sensitive:   true,
									},
									"sim2_pin": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

										Description: "SIM #2 PIN status.",
										Optional:    true,
										Computed:    true,
									},
									"sim2_pin_code": {
										Type: schema.TypeString,

										Description: "SIM #2 PIN password.",
										Optional:    true,
										Computed:    true,
										Sensitive:   true,
									},
								},
							},
						},
						"sms_notification": {
							Type:        schema.TypeList,
							Description: "FortiExtender cellular SMS notification configuration.",
							Optional:    true, MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alert": {
										Type:        schema.TypeList,
										Description: "SMS alert list.",
										Optional:    true, MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"data_exhausted": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 63),

													Description: "Display string when data exhausted.",
													Optional:    true,
													Computed:    true,
												},
												"fgt_backup_mode_switch": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 63),

													Description: "Display string when FortiGate backup mode switched.",
													Optional:    true,
													Computed:    true,
												},
												"low_signal_strength": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 63),

													Description: "Display string when signal strength is low.",
													Optional:    true,
													Computed:    true,
												},
												"mode_switch": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 63),

													Description: "Display string when mode is switched.",
													Optional:    true,
													Computed:    true,
												},
												"os_image_fallback": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 63),

													Description: "Display string when falling back to a previous OS image.",
													Optional:    true,
													Computed:    true,
												},
												"session_disconnect": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 63),

													Description: "Display string when session disconnected.",
													Optional:    true,
													Computed:    true,
												},
												"system_reboot": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 63),

													Description: "Display string when system rebooted.",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"receiver": {
										Type:        schema.TypeList,
										Description: "SMS notification receiver list.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"alert": {
													Type: schema.TypeString,

													DiffSuppressFunc: suppressors.DiffFakeListEqual,
													Description:      "Alert multi-options.",
													Optional:         true,
													Computed:         true,
												},
												"name": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 31),

													Description: "FortiExtender SMS notification receiver name.",
													Optional:    true,
													Computed:    true,
												},
												"phone_number": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 31),

													Description: "Receiver phone number. Format: [+][country code][area code][local phone number]. For example, +16501234567.",
													Optional:    true,
													Computed:    true,
												},
												"status": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

													Description: "SMS notification receiver status.",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"status": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

										Description: "FortiExtender SMS notification status.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"enforce_bandwidth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable enforcement of bandwidth on LAN extension interface.",
				Optional:    true,
				Computed:    true,
			},
			"extension": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"lan-extension"}, false),

				Description: "Extension option.",
				Optional:    true,
				Computed:    true,
			},
			"fosid": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 102400000),

				Description: "ID.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"lan_extension": {
				Type:        schema.TypeList,
				Description: "FortiExtender lan extension configuration.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backhaul": {
							Type:        schema.TypeList,
							Description: "LAN extension backhaul tunnel configuration.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),

										Description: "FortiExtender LAN extension backhaul name.",
										Optional:    true,
										Computed:    true,
									},
									"port": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"wan", "lte1", "lte2", "port1", "port2", "port3", "port4", "port5", "sfp"}, false),

										Description: "FortiExtender uplink port.",
										Optional:    true,
										Computed:    true,
									},
									"role": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"primary", "secondary"}, false),

										Description: "FortiExtender uplink port.",
										Optional:    true,
										Computed:    true,
									},
									"weight": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 256),

										Description: "WRR weight parameter.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"backhaul_interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "IPsec phase1 interface.",
							Optional:    true,
							Computed:    true,
						},
						"backhaul_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "IPsec phase1 IPv4/FQDN. Used to specify the external IP/FQDN when the FortiGate unit is behind a NAT device.",
							Optional:    true,
							Computed:    true,
						},
						"ipsec_tunnel": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "IPsec tunnel name.",
							Optional:    true,
							Computed:    true,
						},
						"link_loadbalance": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"activebackup", "loadbalance"}, false),

							Description: "LAN extension link load balance strategy.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"login_password": {
				Type: schema.TypeString,

				Description: "Set the managed extender's administrator password.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"login_password_change": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"yes", "default", "no"}, false),

				Description: "Change or reset the administrator password of a managed extender (yes, default, or no, default = no).",
				Optional:    true,
				Computed:    true,
			},
			"model": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"FX201E", "FX211E", "FX200F", "FXA11F", "FXE11F", "FXA21F", "FXE21F", "FXA22F", "FXE22F", "FX212F", "FX311F", "FX312F", "FX511F", "FVG21F", "FVA21F", "FVG22F", "FVA22F", "FX04DA"}, false),

				Description: "Model.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "FortiExtender profile name.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceExtenderControllerExtenderProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := ""
	key := "name"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating ExtenderControllerExtenderProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectExtenderControllerExtenderProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateExtenderControllerExtenderProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ExtenderControllerExtenderProfile")
	}

	return resourceExtenderControllerExtenderProfileRead(ctx, d, meta)
}

func resourceExtenderControllerExtenderProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectExtenderControllerExtenderProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateExtenderControllerExtenderProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating ExtenderControllerExtenderProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ExtenderControllerExtenderProfile")
	}

	return resourceExtenderControllerExtenderProfileRead(ctx, d, meta)
}

func resourceExtenderControllerExtenderProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteExtenderControllerExtenderProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting ExtenderControllerExtenderProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtenderControllerExtenderProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	ptp := true
	urlparams.PlainTextPassword = &ptp

	o, err := c.Cmdb.ReadExtenderControllerExtenderProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading ExtenderControllerExtenderProfile resource: %v", err)
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

	diags := refreshObjectExtenderControllerExtenderProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenExtenderControllerExtenderProfileCellular(d *schema.ResourceData, v *models.ExtenderControllerExtenderProfileCellular, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.ExtenderControllerExtenderProfileCellular{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ControllerReport; tmp != nil {
				v["controller_report"] = flattenExtenderControllerExtenderProfileCellularControllerReport(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "controller_report"), sort)
			}

			if tmp := cfg.Dataplan; tmp != nil {
				v["dataplan"] = flattenExtenderControllerExtenderProfileCellularDataplan(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "dataplan"), sort)
			}

			if tmp := cfg.Modem1; tmp != nil {
				v["modem1"] = flattenExtenderControllerExtenderProfileCellularModem1(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "modem1"), sort)
			}

			if tmp := cfg.Modem2; tmp != nil {
				v["modem2"] = flattenExtenderControllerExtenderProfileCellularModem2(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "modem2"), sort)
			}

			if tmp := cfg.SmsNotification; tmp != nil {
				v["sms_notification"] = flattenExtenderControllerExtenderProfileCellularSmsNotification(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "sms_notification"), sort)
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenExtenderControllerExtenderProfileCellularControllerReport(d *schema.ResourceData, v *models.ExtenderControllerExtenderProfileCellularControllerReport, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.ExtenderControllerExtenderProfileCellularControllerReport{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Interval; tmp != nil {
				v["interval"] = *tmp
			}

			if tmp := cfg.SignalThreshold; tmp != nil {
				v["signal_threshold"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenExtenderControllerExtenderProfileCellularDataplan(d *schema.ResourceData, v *[]models.ExtenderControllerExtenderProfileCellularDataplan, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenExtenderControllerExtenderProfileCellularModem1(d *schema.ResourceData, v *models.ExtenderControllerExtenderProfileCellularModem1, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.ExtenderControllerExtenderProfileCellularModem1{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AutoSwitch; tmp != nil {
				v["auto_switch"] = flattenExtenderControllerExtenderProfileCellularModem1AutoSwitch(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "auto_switch"), sort)
			}

			if tmp := cfg.ConnStatus; tmp != nil {
				v["conn_status"] = *tmp
			}

			if tmp := cfg.DefaultSim; tmp != nil {
				v["default_sim"] = *tmp
			}

			if tmp := cfg.Gps; tmp != nil {
				v["gps"] = *tmp
			}

			if tmp := cfg.PreferredCarrier; tmp != nil {
				v["preferred_carrier"] = *tmp
			}

			if tmp := cfg.RedundantIntf; tmp != nil {
				v["redundant_intf"] = *tmp
			}

			if tmp := cfg.RedundantMode; tmp != nil {
				v["redundant_mode"] = *tmp
			}

			if tmp := cfg.Sim1Pin; tmp != nil {
				v["sim1_pin"] = *tmp
			}

			if tmp := cfg.Sim1PinCode; tmp != nil {
				v["sim1_pin_code"] = *tmp
			}

			if tmp := cfg.Sim2Pin; tmp != nil {
				v["sim2_pin"] = *tmp
			}

			if tmp := cfg.Sim2PinCode; tmp != nil {
				v["sim2_pin_code"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenExtenderControllerExtenderProfileCellularModem1AutoSwitch(d *schema.ResourceData, v *models.ExtenderControllerExtenderProfileCellularModem1AutoSwitch, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.ExtenderControllerExtenderProfileCellularModem1AutoSwitch{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Dataplan; tmp != nil {
				v["dataplan"] = *tmp
			}

			if tmp := cfg.Disconnect; tmp != nil {
				v["disconnect"] = *tmp
			}

			if tmp := cfg.DisconnectPeriod; tmp != nil {
				v["disconnect_period"] = *tmp
			}

			if tmp := cfg.DisconnectThreshold; tmp != nil {
				v["disconnect_threshold"] = *tmp
			}

			if tmp := cfg.Signal; tmp != nil {
				v["signal"] = *tmp
			}

			if tmp := cfg.SwitchBack; tmp != nil {
				v["switch_back"] = *tmp
			}

			if tmp := cfg.SwitchBackTime; tmp != nil {
				v["switch_back_time"] = *tmp
			}

			if tmp := cfg.SwitchBackTimer; tmp != nil {
				v["switch_back_timer"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenExtenderControllerExtenderProfileCellularModem2(d *schema.ResourceData, v *models.ExtenderControllerExtenderProfileCellularModem2, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.ExtenderControllerExtenderProfileCellularModem2{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AutoSwitch; tmp != nil {
				v["auto_switch"] = flattenExtenderControllerExtenderProfileCellularModem2AutoSwitch(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "auto_switch"), sort)
			}

			if tmp := cfg.ConnStatus; tmp != nil {
				v["conn_status"] = *tmp
			}

			if tmp := cfg.DefaultSim; tmp != nil {
				v["default_sim"] = *tmp
			}

			if tmp := cfg.Gps; tmp != nil {
				v["gps"] = *tmp
			}

			if tmp := cfg.PreferredCarrier; tmp != nil {
				v["preferred_carrier"] = *tmp
			}

			if tmp := cfg.RedundantIntf; tmp != nil {
				v["redundant_intf"] = *tmp
			}

			if tmp := cfg.RedundantMode; tmp != nil {
				v["redundant_mode"] = *tmp
			}

			if tmp := cfg.Sim1Pin; tmp != nil {
				v["sim1_pin"] = *tmp
			}

			if tmp := cfg.Sim1PinCode; tmp != nil {
				v["sim1_pin_code"] = *tmp
			}

			if tmp := cfg.Sim2Pin; tmp != nil {
				v["sim2_pin"] = *tmp
			}

			if tmp := cfg.Sim2PinCode; tmp != nil {
				v["sim2_pin_code"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenExtenderControllerExtenderProfileCellularModem2AutoSwitch(d *schema.ResourceData, v *models.ExtenderControllerExtenderProfileCellularModem2AutoSwitch, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.ExtenderControllerExtenderProfileCellularModem2AutoSwitch{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Dataplan; tmp != nil {
				v["dataplan"] = *tmp
			}

			if tmp := cfg.Disconnect; tmp != nil {
				v["disconnect"] = *tmp
			}

			if tmp := cfg.DisconnectPeriod; tmp != nil {
				v["disconnect_period"] = *tmp
			}

			if tmp := cfg.DisconnectThreshold; tmp != nil {
				v["disconnect_threshold"] = *tmp
			}

			if tmp := cfg.Signal; tmp != nil {
				v["signal"] = *tmp
			}

			if tmp := cfg.SwitchBack; tmp != nil {
				v["switch_back"] = *tmp
			}

			if tmp := cfg.SwitchBackTime; tmp != nil {
				v["switch_back_time"] = *tmp
			}

			if tmp := cfg.SwitchBackTimer; tmp != nil {
				v["switch_back_timer"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenExtenderControllerExtenderProfileCellularSmsNotification(d *schema.ResourceData, v *models.ExtenderControllerExtenderProfileCellularSmsNotification, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.ExtenderControllerExtenderProfileCellularSmsNotification{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Alert; tmp != nil {
				v["alert"] = flattenExtenderControllerExtenderProfileCellularSmsNotificationAlert(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "alert"), sort)
			}

			if tmp := cfg.Receiver; tmp != nil {
				v["receiver"] = flattenExtenderControllerExtenderProfileCellularSmsNotificationReceiver(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "receiver"), sort)
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenExtenderControllerExtenderProfileCellularSmsNotificationAlert(d *schema.ResourceData, v *models.ExtenderControllerExtenderProfileCellularSmsNotificationAlert, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.ExtenderControllerExtenderProfileCellularSmsNotificationAlert{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.DataExhausted; tmp != nil {
				v["data_exhausted"] = *tmp
			}

			if tmp := cfg.FgtBackupModeSwitch; tmp != nil {
				v["fgt_backup_mode_switch"] = *tmp
			}

			if tmp := cfg.LowSignalStrength; tmp != nil {
				v["low_signal_strength"] = *tmp
			}

			if tmp := cfg.ModeSwitch; tmp != nil {
				v["mode_switch"] = *tmp
			}

			if tmp := cfg.OsImageFallback; tmp != nil {
				v["os_image_fallback"] = *tmp
			}

			if tmp := cfg.SessionDisconnect; tmp != nil {
				v["session_disconnect"] = *tmp
			}

			if tmp := cfg.SystemReboot; tmp != nil {
				v["system_reboot"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenExtenderControllerExtenderProfileCellularSmsNotificationReceiver(d *schema.ResourceData, v *[]models.ExtenderControllerExtenderProfileCellularSmsNotificationReceiver, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Alert; tmp != nil {
				v["alert"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.PhoneNumber; tmp != nil {
				v["phone_number"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenExtenderControllerExtenderProfileLanExtension(d *schema.ResourceData, v *models.ExtenderControllerExtenderProfileLanExtension, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.ExtenderControllerExtenderProfileLanExtension{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Backhaul; tmp != nil {
				v["backhaul"] = flattenExtenderControllerExtenderProfileLanExtensionBackhaul(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "backhaul"), sort)
			}

			if tmp := cfg.BackhaulInterface; tmp != nil {
				v["backhaul_interface"] = *tmp
			}

			if tmp := cfg.BackhaulIp; tmp != nil {
				v["backhaul_ip"] = *tmp
			}

			if tmp := cfg.IpsecTunnel; tmp != nil {
				v["ipsec_tunnel"] = *tmp
			}

			if tmp := cfg.LinkLoadbalance; tmp != nil {
				v["link_loadbalance"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenExtenderControllerExtenderProfileLanExtensionBackhaul(d *schema.ResourceData, v *[]models.ExtenderControllerExtenderProfileLanExtensionBackhaul, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Port; tmp != nil {
				v["port"] = *tmp
			}

			if tmp := cfg.Role; tmp != nil {
				v["role"] = *tmp
			}

			if tmp := cfg.Weight; tmp != nil {
				v["weight"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectExtenderControllerExtenderProfile(d *schema.ResourceData, o *models.ExtenderControllerExtenderProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Allowaccess != nil {
		v := *o.Allowaccess

		if err = d.Set("allowaccess", v); err != nil {
			return diag.Errorf("error reading allowaccess: %v", err)
		}
	}

	if o.BandwidthLimit != nil {
		v := *o.BandwidthLimit

		if err = d.Set("bandwidth_limit", v); err != nil {
			return diag.Errorf("error reading bandwidth_limit: %v", err)
		}
	}

	if _, ok := d.GetOk("cellular"); ok {
		if o.Cellular != nil {
			if err = d.Set("cellular", flattenExtenderControllerExtenderProfileCellular(d, o.Cellular, "cellular", sort)); err != nil {
				return diag.Errorf("error reading cellular: %v", err)
			}
		}
	}

	if o.EnforceBandwidth != nil {
		v := *o.EnforceBandwidth

		if err = d.Set("enforce_bandwidth", v); err != nil {
			return diag.Errorf("error reading enforce_bandwidth: %v", err)
		}
	}

	if o.Extension != nil {
		v := *o.Extension

		if err = d.Set("extension", v); err != nil {
			return diag.Errorf("error reading extension: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if _, ok := d.GetOk("lan_extension"); ok {
		if o.LanExtension != nil {
			if err = d.Set("lan_extension", flattenExtenderControllerExtenderProfileLanExtension(d, o.LanExtension, "lan_extension", sort)); err != nil {
				return diag.Errorf("error reading lan_extension: %v", err)
			}
		}
	}

	if o.LoginPassword != nil {
		v := *o.LoginPassword

		if v == "ENC XXXX" {
		} else if err = d.Set("login_password", v); err != nil {
			return diag.Errorf("error reading login_password: %v", err)
		}
	}

	if o.LoginPasswordChange != nil {
		v := *o.LoginPasswordChange

		if err = d.Set("login_password_change", v); err != nil {
			return diag.Errorf("error reading login_password_change: %v", err)
		}
	}

	if o.Model != nil {
		v := *o.Model

		if err = d.Set("model", v); err != nil {
			return diag.Errorf("error reading model: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	return nil
}

func expandExtenderControllerExtenderProfileCellular(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.ExtenderControllerExtenderProfileCellular, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ExtenderControllerExtenderProfileCellular

	for i := range l {
		tmp := models.ExtenderControllerExtenderProfileCellular{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.controller_report", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandExtenderControllerExtenderProfileCellularControllerReport(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ExtenderControllerExtenderProfileCellularControllerReport
			// 	}
			tmp.ControllerReport = v2

		}

		pre_append = fmt.Sprintf("%s.%d.dataplan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandExtenderControllerExtenderProfileCellularDataplan(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ExtenderControllerExtenderProfileCellularDataplan
			// 	}
			tmp.Dataplan = v2

		}

		pre_append = fmt.Sprintf("%s.%d.modem1", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandExtenderControllerExtenderProfileCellularModem1(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ExtenderControllerExtenderProfileCellularModem1
			// 	}
			tmp.Modem1 = v2

		}

		pre_append = fmt.Sprintf("%s.%d.modem2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandExtenderControllerExtenderProfileCellularModem2(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ExtenderControllerExtenderProfileCellularModem2
			// 	}
			tmp.Modem2 = v2

		}

		pre_append = fmt.Sprintf("%s.%d.sms_notification", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandExtenderControllerExtenderProfileCellularSmsNotification(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ExtenderControllerExtenderProfileCellularSmsNotification
			// 	}
			tmp.SmsNotification = v2

		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandExtenderControllerExtenderProfileCellularControllerReport(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.ExtenderControllerExtenderProfileCellularControllerReport, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ExtenderControllerExtenderProfileCellularControllerReport

	for i := range l {
		tmp := models.ExtenderControllerExtenderProfileCellularControllerReport{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Interval = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.signal_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SignalThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandExtenderControllerExtenderProfileCellularDataplan(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ExtenderControllerExtenderProfileCellularDataplan, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ExtenderControllerExtenderProfileCellularDataplan

	for i := range l {
		tmp := models.ExtenderControllerExtenderProfileCellularDataplan{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandExtenderControllerExtenderProfileCellularModem1(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.ExtenderControllerExtenderProfileCellularModem1, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ExtenderControllerExtenderProfileCellularModem1

	for i := range l {
		tmp := models.ExtenderControllerExtenderProfileCellularModem1{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.auto_switch", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandExtenderControllerExtenderProfileCellularModem1AutoSwitch(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ExtenderControllerExtenderProfileCellularModem1AutoSwitch
			// 	}
			tmp.AutoSwitch = v2

		}

		pre_append = fmt.Sprintf("%s.%d.conn_status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ConnStatus = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.default_sim", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DefaultSim = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.gps", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Gps = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.preferred_carrier", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PreferredCarrier = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.redundant_intf", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RedundantIntf = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.redundant_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RedundantMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sim1_pin", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Sim1Pin = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sim1_pin_code", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Sim1PinCode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sim2_pin", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Sim2Pin = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sim2_pin_code", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Sim2PinCode = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandExtenderControllerExtenderProfileCellularModem1AutoSwitch(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.ExtenderControllerExtenderProfileCellularModem1AutoSwitch, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ExtenderControllerExtenderProfileCellularModem1AutoSwitch

	for i := range l {
		tmp := models.ExtenderControllerExtenderProfileCellularModem1AutoSwitch{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dataplan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dataplan = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.disconnect", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Disconnect = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.disconnect_period", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.DisconnectPeriod = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.disconnect_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.DisconnectThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.signal", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Signal = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.switch_back", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SwitchBack = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.switch_back_time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SwitchBackTime = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.switch_back_timer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SwitchBackTimer = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandExtenderControllerExtenderProfileCellularModem2(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.ExtenderControllerExtenderProfileCellularModem2, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ExtenderControllerExtenderProfileCellularModem2

	for i := range l {
		tmp := models.ExtenderControllerExtenderProfileCellularModem2{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.auto_switch", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandExtenderControllerExtenderProfileCellularModem2AutoSwitch(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ExtenderControllerExtenderProfileCellularModem2AutoSwitch
			// 	}
			tmp.AutoSwitch = v2

		}

		pre_append = fmt.Sprintf("%s.%d.conn_status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ConnStatus = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.default_sim", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DefaultSim = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.gps", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Gps = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.preferred_carrier", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PreferredCarrier = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.redundant_intf", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RedundantIntf = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.redundant_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RedundantMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sim1_pin", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Sim1Pin = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sim1_pin_code", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Sim1PinCode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sim2_pin", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Sim2Pin = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sim2_pin_code", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Sim2PinCode = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandExtenderControllerExtenderProfileCellularModem2AutoSwitch(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.ExtenderControllerExtenderProfileCellularModem2AutoSwitch, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ExtenderControllerExtenderProfileCellularModem2AutoSwitch

	for i := range l {
		tmp := models.ExtenderControllerExtenderProfileCellularModem2AutoSwitch{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dataplan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dataplan = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.disconnect", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Disconnect = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.disconnect_period", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.DisconnectPeriod = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.disconnect_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.DisconnectThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.signal", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Signal = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.switch_back", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SwitchBack = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.switch_back_time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SwitchBackTime = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.switch_back_timer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SwitchBackTimer = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandExtenderControllerExtenderProfileCellularSmsNotification(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.ExtenderControllerExtenderProfileCellularSmsNotification, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ExtenderControllerExtenderProfileCellularSmsNotification

	for i := range l {
		tmp := models.ExtenderControllerExtenderProfileCellularSmsNotification{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.alert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandExtenderControllerExtenderProfileCellularSmsNotificationAlert(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ExtenderControllerExtenderProfileCellularSmsNotificationAlert
			// 	}
			tmp.Alert = v2

		}

		pre_append = fmt.Sprintf("%s.%d.receiver", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandExtenderControllerExtenderProfileCellularSmsNotificationReceiver(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ExtenderControllerExtenderProfileCellularSmsNotificationReceiver
			// 	}
			tmp.Receiver = v2

		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandExtenderControllerExtenderProfileCellularSmsNotificationAlert(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.ExtenderControllerExtenderProfileCellularSmsNotificationAlert, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ExtenderControllerExtenderProfileCellularSmsNotificationAlert

	for i := range l {
		tmp := models.ExtenderControllerExtenderProfileCellularSmsNotificationAlert{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.data_exhausted", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DataExhausted = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fgt_backup_mode_switch", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FgtBackupModeSwitch = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.low_signal_strength", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LowSignalStrength = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mode_switch", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ModeSwitch = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.os_image_fallback", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OsImageFallback = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.session_disconnect", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SessionDisconnect = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.system_reboot", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SystemReboot = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandExtenderControllerExtenderProfileCellularSmsNotificationReceiver(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ExtenderControllerExtenderProfileCellularSmsNotificationReceiver, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ExtenderControllerExtenderProfileCellularSmsNotificationReceiver

	for i := range l {
		tmp := models.ExtenderControllerExtenderProfileCellularSmsNotificationReceiver{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.alert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Alert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.phone_number", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PhoneNumber = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandExtenderControllerExtenderProfileLanExtension(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.ExtenderControllerExtenderProfileLanExtension, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ExtenderControllerExtenderProfileLanExtension

	for i := range l {
		tmp := models.ExtenderControllerExtenderProfileLanExtension{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.backhaul", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandExtenderControllerExtenderProfileLanExtensionBackhaul(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ExtenderControllerExtenderProfileLanExtensionBackhaul
			// 	}
			tmp.Backhaul = v2

		}

		pre_append = fmt.Sprintf("%s.%d.backhaul_interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BackhaulInterface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.backhaul_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BackhaulIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipsec_tunnel", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IpsecTunnel = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.link_loadbalance", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LinkLoadbalance = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandExtenderControllerExtenderProfileLanExtensionBackhaul(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ExtenderControllerExtenderProfileLanExtensionBackhaul, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ExtenderControllerExtenderProfileLanExtensionBackhaul

	for i := range l {
		tmp := models.ExtenderControllerExtenderProfileLanExtensionBackhaul{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Port = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.role", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Role = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.weight", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Weight = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectExtenderControllerExtenderProfile(d *schema.ResourceData, sv string) (*models.ExtenderControllerExtenderProfile, diag.Diagnostics) {
	obj := models.ExtenderControllerExtenderProfile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("allowaccess"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("allowaccess", sv)
				diags = append(diags, e)
			}
			obj.Allowaccess = &v2
		}
	}
	if v1, ok := d.GetOk("bandwidth_limit"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bandwidth_limit", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.BandwidthLimit = &tmp
		}
	}
	if v, ok := d.GetOk("cellular"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("cellular", sv)
			diags = append(diags, e)
		}
		t, err := expandExtenderControllerExtenderProfileCellular(d, v, "cellular", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Cellular = t
		}
	} else if d.HasChange("cellular") {
		old, new := d.GetChange("cellular")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Cellular = &models.ExtenderControllerExtenderProfileCellular{}
		}
	}
	if v1, ok := d.GetOk("enforce_bandwidth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("enforce_bandwidth", sv)
				diags = append(diags, e)
			}
			obj.EnforceBandwidth = &v2
		}
	}
	if v1, ok := d.GetOk("extension"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("extension", sv)
				diags = append(diags, e)
			}
			obj.Extension = &v2
		}
	}
	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Id = &tmp
		}
	}
	if v, ok := d.GetOk("lan_extension"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("lan_extension", sv)
			diags = append(diags, e)
		}
		t, err := expandExtenderControllerExtenderProfileLanExtension(d, v, "lan_extension", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.LanExtension = t
		}
	} else if d.HasChange("lan_extension") {
		old, new := d.GetChange("lan_extension")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.LanExtension = &models.ExtenderControllerExtenderProfileLanExtension{}
		}
	}
	if v1, ok := d.GetOk("login_password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("login_password", sv)
				diags = append(diags, e)
			}
			obj.LoginPassword = &v2
		}
	}
	if v1, ok := d.GetOk("login_password_change"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("login_password_change", sv)
				diags = append(diags, e)
			}
			obj.LoginPasswordChange = &v2
		}
	}
	if v1, ok := d.GetOk("model"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("model", sv)
				diags = append(diags, e)
			}
			obj.Model = &v2
		}
	}
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	return &obj, diags
}
