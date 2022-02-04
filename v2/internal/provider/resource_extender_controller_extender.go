// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Extender controller configuration.

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

func resourceExtenderControllerExtender() *schema.Resource {
	return &schema.Resource{
		Description: "Extender controller configuration.",

		CreateContext: resourceExtenderControllerExtenderCreate,
		ReadContext:   resourceExtenderControllerExtenderRead,
		UpdateContext: resourceExtenderControllerExtenderUpdate,
		DeleteContext: resourceExtenderControllerExtenderDelete,

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
			"aaa_shared_secret": {
				Type: schema.TypeString,

				Description: "AAA shared secret.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"access_point_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Access point name(APN).",
				Optional:    true,
				Computed:    true,
			},
			"admin": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "discovered", "enable"}, false),

				Description: "FortiExtender Administration (enable or disable).",
				Optional:    true,
				Computed:    true,
			},
			"allowaccess": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Control management access to the managed extender. Separate entries with a space.",
				Optional:         true,
				Computed:         true,
			},
			"at_dial_script": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Initialization AT commands specific to the MODEM.",
				Optional:    true,
				Computed:    true,
			},
			"authorized": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "FortiExtender Administration (enable or disable).",
				Optional:    true,
				Computed:    true,
			},
			"bandwidth_limit": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16776000),

				Description: "FortiExtender LAN extension bandwidth limit (Mbps).",
				Optional:    true,
				Computed:    true,
			},
			"billing_start_day": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 28),

				Description: "Billing start day.",
				Optional:    true,
				Computed:    true,
			},
			"cdma_aaa_spi": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "CDMA AAA SPI.",
				Optional:    true,
				Computed:    true,
			},
			"cdma_ha_spi": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "CDMA HA SPI.",
				Optional:    true,
				Computed:    true,
			},
			"cdma_nai": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "NAI for CDMA MODEMS.",
				Optional:    true,
				Computed:    true,
			},
			"conn_status": {
				Type: schema.TypeInt,

				Description: "Connection status.",
				Optional:    true,
				Computed:    true,
			},
			"controller_report": {
				Type:        schema.TypeList,
				Description: "FortiExtender controller report configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interval": {
							Type: schema.TypeInt,

							Description: "Controller report interval.",
							Optional:    true,
							Computed:    true,
						},
						"signal_threshold": {
							Type: schema.TypeInt,

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
			"description": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Description.",
				Optional:    true,
				Computed:    true,
			},
			"device_id": {
				Type: schema.TypeInt,

				Description: "Device ID.",
				Optional:    true,
				Computed:    true,
			},
			"dial_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"dial-on-demand", "always-connect"}, false),

				Description: "Dial mode (dial-on-demand or always-connect).",
				Optional:    true,
				Computed:    true,
			},
			"dial_status": {
				Type: schema.TypeInt,

				Description: "Dial status.",
				Optional:    true,
				Computed:    true,
			},
			"enforce_bandwidth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable enforcement of bandwidth on LAN extension interface.",
				Optional:    true,
				Computed:    true,
			},
			"ext_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "FortiExtender name.",
				Optional:    true,
				Computed:    true,
			},
			"extension_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"lan-extension"}, false),

				Description: "Extension type for this FortiExtender.",
				Optional:    true,
				Computed:    true,
			},
			"ha_shared_secret": {
				Type: schema.TypeString,

				Description: "HA shared secret.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"fosid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),

				Description: "FortiExtender serial number.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"ifname": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "FortiExtender interface name.",
				Optional:    true,
				Computed:    true,
			},
			"initiated_update": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Allow/disallow network initiated updates to the MODEM.",
				Optional:    true,
				Computed:    true,
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
			"mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"standalone", "redundant"}, false),

				Description: "FortiExtender mode.",
				Optional:    true,
				Computed:    true,
			},
			"modem_passwd": {
				Type: schema.TypeString,

				Description: "MODEM password.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"modem_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"cdma", "gsm/lte", "wimax"}, false),

				Description: "MODEM type (CDMA, GSM/LTE or WIMAX).",
				Optional:    true,
				Computed:    true,
			},
			"modem1": {
				Type:        schema.TypeList,
				Description: "Configuration options for modem 1.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_switch": {
							Type:        schema.TypeList,
							Description: "FortiExtender auto switch configuration.",
							Optional:    true,
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
										Type: schema.TypeInt,

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
						"ifname": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "FortiExtender interface name.",
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
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_switch": {
							Type:        schema.TypeList,
							Description: "FortiExtender auto switch configuration.",
							Optional:    true,
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
										Type: schema.TypeInt,

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
						"ifname": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "FortiExtender interface name.",
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
			"multi_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "auto-3g", "force-lte", "force-3g", "force-2g"}, false),

				Description: "MODEM mode of operation(3G,LTE,etc).",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),

				Description: "FortiExtender entry name.",
				ForceNew:    true,
				Required:    true,
			},
			"override_allowaccess": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to override the extender profile management access configuration.",
				Optional:    true,
				Computed:    true,
			},
			"override_enforce_bandwidth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to override the extender profile enforce-bandwidth setting.",
				Optional:    true,
				Computed:    true,
			},
			"override_login_password_change": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to override the extender profile login-password (administrator password) setting.",
				Optional:    true,
				Computed:    true,
			},
			"ppp_auth_protocol": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "pap", "chap"}, false),

				Description: "PPP authentication protocol (PAP,CHAP or auto).",
				Optional:    true,
				Computed:    true,
			},
			"ppp_echo_request": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable PPP echo request.",
				Optional:    true,
				Computed:    true,
			},
			"ppp_password": {
				Type: schema.TypeString,

				Description: "PPP password.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"ppp_username": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "PPP username.",
				Optional:    true,
				Computed:    true,
			},
			"primary_ha": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "Primary HA.",
				Optional:    true,
				Computed:    true,
			},
			"profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "FortiExtender profile configuration.",
				Optional:    true,
				Computed:    true,
			},
			"quota_limit_mb": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10485760),

				Description: "Monthly quota limit (MB).",
				Optional:    true,
				Computed:    true,
			},
			"redial": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, false),

				Description: "Number of redials allowed based on failed attempts.",
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
			"roaming": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable MODEM roaming.",
				Optional:    true,
				Computed:    true,
			},
			"role": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "primary", "secondary"}, false),

				Description: "FortiExtender work role(Primary, Secondary, None).",
				Optional:    true,
				Computed:    true,
			},
			"secondary_ha": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "Secondary HA.",
				Optional:    true,
				Computed:    true,
			},
			"sim_pin": {
				Type: schema.TypeString,

				Description: "SIM PIN.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"vdom": {
				Type: schema.TypeInt,

				Description: "VDOM.",
				Optional:    true,
				Computed:    true,
			},
			"wan_extension": {
				Type:        schema.TypeList,
				Description: "FortiExtender wan extension configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"modem1_extension": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),

							Description: "FortiExtender interface name.",
							Optional:    true,
							Computed:    true,
						},
						"modem2_extension": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),

							Description: "FortiExtender interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"wimax_auth_protocol": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"tls", "ttls"}, false),

				Description: "WiMax authentication protocol(TLS or TTLS).",
				Optional:    true,
				Computed:    true,
			},
			"wimax_carrier": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "WiMax carrier.",
				Optional:    true,
				Computed:    true,
			},
			"wimax_realm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "WiMax realm.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceExtenderControllerExtenderCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating ExtenderControllerExtender resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectExtenderControllerExtender(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateExtenderControllerExtender(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ExtenderControllerExtender")
	}

	return resourceExtenderControllerExtenderRead(ctx, d, meta)
}

func resourceExtenderControllerExtenderUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectExtenderControllerExtender(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateExtenderControllerExtender(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating ExtenderControllerExtender resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ExtenderControllerExtender")
	}

	return resourceExtenderControllerExtenderRead(ctx, d, meta)
}

func resourceExtenderControllerExtenderDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteExtenderControllerExtender(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting ExtenderControllerExtender resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtenderControllerExtenderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadExtenderControllerExtender(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading ExtenderControllerExtender resource: %v", err)
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

	diags := refreshObjectExtenderControllerExtender(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenExtenderControllerExtenderControllerReport(v *[]models.ExtenderControllerExtenderControllerReport, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenExtenderControllerExtenderModem1(v *[]models.ExtenderControllerExtenderModem1, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AutoSwitch; tmp != nil {
				v["auto_switch"] = flattenExtenderControllerExtenderModem1AutoSwitch(tmp, sort)
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

			if tmp := cfg.Ifname; tmp != nil {
				v["ifname"] = *tmp
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

func flattenExtenderControllerExtenderModem1AutoSwitch(v *[]models.ExtenderControllerExtenderModem1AutoSwitch, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenExtenderControllerExtenderModem2(v *[]models.ExtenderControllerExtenderModem2, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AutoSwitch; tmp != nil {
				v["auto_switch"] = flattenExtenderControllerExtenderModem2AutoSwitch(tmp, sort)
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

			if tmp := cfg.Ifname; tmp != nil {
				v["ifname"] = *tmp
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

func flattenExtenderControllerExtenderModem2AutoSwitch(v *[]models.ExtenderControllerExtenderModem2AutoSwitch, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenExtenderControllerExtenderWanExtension(v *[]models.ExtenderControllerExtenderWanExtension, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Modem1Extension; tmp != nil {
				v["modem1_extension"] = *tmp
			}

			if tmp := cfg.Modem2Extension; tmp != nil {
				v["modem2_extension"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func refreshObjectExtenderControllerExtender(d *schema.ResourceData, o *models.ExtenderControllerExtender, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AaaSharedSecret != nil {
		v := *o.AaaSharedSecret

		if v == "ENC XXXX" {
		} else if err = d.Set("aaa_shared_secret", v); err != nil {
			return diag.Errorf("error reading aaa_shared_secret: %v", err)
		}
	}

	if o.AccessPointName != nil {
		v := *o.AccessPointName

		if err = d.Set("access_point_name", v); err != nil {
			return diag.Errorf("error reading access_point_name: %v", err)
		}
	}

	if o.Admin != nil {
		v := *o.Admin

		if err = d.Set("admin", v); err != nil {
			return diag.Errorf("error reading admin: %v", err)
		}
	}

	if o.Allowaccess != nil {
		v := *o.Allowaccess

		if err = d.Set("allowaccess", v); err != nil {
			return diag.Errorf("error reading allowaccess: %v", err)
		}
	}

	if o.AtDialScript != nil {
		v := *o.AtDialScript

		if err = d.Set("at_dial_script", v); err != nil {
			return diag.Errorf("error reading at_dial_script: %v", err)
		}
	}

	if o.Authorized != nil {
		v := *o.Authorized

		if err = d.Set("authorized", v); err != nil {
			return diag.Errorf("error reading authorized: %v", err)
		}
	}

	if o.BandwidthLimit != nil {
		v := *o.BandwidthLimit

		if err = d.Set("bandwidth_limit", v); err != nil {
			return diag.Errorf("error reading bandwidth_limit: %v", err)
		}
	}

	if o.BillingStartDay != nil {
		v := *o.BillingStartDay

		if err = d.Set("billing_start_day", v); err != nil {
			return diag.Errorf("error reading billing_start_day: %v", err)
		}
	}

	if o.CdmaAaaSpi != nil {
		v := *o.CdmaAaaSpi

		if err = d.Set("cdma_aaa_spi", v); err != nil {
			return diag.Errorf("error reading cdma_aaa_spi: %v", err)
		}
	}

	if o.CdmaHaSpi != nil {
		v := *o.CdmaHaSpi

		if err = d.Set("cdma_ha_spi", v); err != nil {
			return diag.Errorf("error reading cdma_ha_spi: %v", err)
		}
	}

	if o.CdmaNai != nil {
		v := *o.CdmaNai

		if err = d.Set("cdma_nai", v); err != nil {
			return diag.Errorf("error reading cdma_nai: %v", err)
		}
	}

	if o.ConnStatus != nil {
		v := *o.ConnStatus

		if err = d.Set("conn_status", v); err != nil {
			return diag.Errorf("error reading conn_status: %v", err)
		}
	}

	if o.ControllerReport != nil {
		if err = d.Set("controller_report", flattenExtenderControllerExtenderControllerReport(o.ControllerReport, sort)); err != nil {
			return diag.Errorf("error reading controller_report: %v", err)
		}
	}

	if o.Description != nil {
		v := *o.Description

		if err = d.Set("description", v); err != nil {
			return diag.Errorf("error reading description: %v", err)
		}
	}

	if o.DeviceId != nil {
		v := *o.DeviceId

		if err = d.Set("device_id", v); err != nil {
			return diag.Errorf("error reading device_id: %v", err)
		}
	}

	if o.DialMode != nil {
		v := *o.DialMode

		if err = d.Set("dial_mode", v); err != nil {
			return diag.Errorf("error reading dial_mode: %v", err)
		}
	}

	if o.DialStatus != nil {
		v := *o.DialStatus

		if err = d.Set("dial_status", v); err != nil {
			return diag.Errorf("error reading dial_status: %v", err)
		}
	}

	if o.EnforceBandwidth != nil {
		v := *o.EnforceBandwidth

		if err = d.Set("enforce_bandwidth", v); err != nil {
			return diag.Errorf("error reading enforce_bandwidth: %v", err)
		}
	}

	if o.ExtName != nil {
		v := *o.ExtName

		if err = d.Set("ext_name", v); err != nil {
			return diag.Errorf("error reading ext_name: %v", err)
		}
	}

	if o.ExtensionType != nil {
		v := *o.ExtensionType

		if err = d.Set("extension_type", v); err != nil {
			return diag.Errorf("error reading extension_type: %v", err)
		}
	}

	if o.HaSharedSecret != nil {
		v := *o.HaSharedSecret

		if v == "ENC XXXX" {
		} else if err = d.Set("ha_shared_secret", v); err != nil {
			return diag.Errorf("error reading ha_shared_secret: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.Ifname != nil {
		v := *o.Ifname

		if err = d.Set("ifname", v); err != nil {
			return diag.Errorf("error reading ifname: %v", err)
		}
	}

	if o.InitiatedUpdate != nil {
		v := *o.InitiatedUpdate

		if err = d.Set("initiated_update", v); err != nil {
			return diag.Errorf("error reading initiated_update: %v", err)
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

	if o.Mode != nil {
		v := *o.Mode

		if err = d.Set("mode", v); err != nil {
			return diag.Errorf("error reading mode: %v", err)
		}
	}

	if o.ModemPasswd != nil {
		v := *o.ModemPasswd

		if v == "ENC XXXX" {
		} else if err = d.Set("modem_passwd", v); err != nil {
			return diag.Errorf("error reading modem_passwd: %v", err)
		}
	}

	if o.ModemType != nil {
		v := *o.ModemType

		if err = d.Set("modem_type", v); err != nil {
			return diag.Errorf("error reading modem_type: %v", err)
		}
	}

	if o.Modem1 != nil {
		if err = d.Set("modem1", flattenExtenderControllerExtenderModem1(o.Modem1, sort)); err != nil {
			return diag.Errorf("error reading modem1: %v", err)
		}
	}

	if o.Modem2 != nil {
		if err = d.Set("modem2", flattenExtenderControllerExtenderModem2(o.Modem2, sort)); err != nil {
			return diag.Errorf("error reading modem2: %v", err)
		}
	}

	if o.MultiMode != nil {
		v := *o.MultiMode

		if err = d.Set("multi_mode", v); err != nil {
			return diag.Errorf("error reading multi_mode: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.OverrideAllowaccess != nil {
		v := *o.OverrideAllowaccess

		if err = d.Set("override_allowaccess", v); err != nil {
			return diag.Errorf("error reading override_allowaccess: %v", err)
		}
	}

	if o.OverrideEnforceBandwidth != nil {
		v := *o.OverrideEnforceBandwidth

		if err = d.Set("override_enforce_bandwidth", v); err != nil {
			return diag.Errorf("error reading override_enforce_bandwidth: %v", err)
		}
	}

	if o.OverrideLoginPasswordChange != nil {
		v := *o.OverrideLoginPasswordChange

		if err = d.Set("override_login_password_change", v); err != nil {
			return diag.Errorf("error reading override_login_password_change: %v", err)
		}
	}

	if o.PppAuthProtocol != nil {
		v := *o.PppAuthProtocol

		if err = d.Set("ppp_auth_protocol", v); err != nil {
			return diag.Errorf("error reading ppp_auth_protocol: %v", err)
		}
	}

	if o.PppEchoRequest != nil {
		v := *o.PppEchoRequest

		if err = d.Set("ppp_echo_request", v); err != nil {
			return diag.Errorf("error reading ppp_echo_request: %v", err)
		}
	}

	if o.PppPassword != nil {
		v := *o.PppPassword

		if v == "ENC XXXX" {
		} else if err = d.Set("ppp_password", v); err != nil {
			return diag.Errorf("error reading ppp_password: %v", err)
		}
	}

	if o.PppUsername != nil {
		v := *o.PppUsername

		if err = d.Set("ppp_username", v); err != nil {
			return diag.Errorf("error reading ppp_username: %v", err)
		}
	}

	if o.PrimaryHa != nil {
		v := *o.PrimaryHa

		if err = d.Set("primary_ha", v); err != nil {
			return diag.Errorf("error reading primary_ha: %v", err)
		}
	}

	if o.Profile != nil {
		v := *o.Profile

		if err = d.Set("profile", v); err != nil {
			return diag.Errorf("error reading profile: %v", err)
		}
	}

	if o.QuotaLimitMb != nil {
		v := *o.QuotaLimitMb

		if err = d.Set("quota_limit_mb", v); err != nil {
			return diag.Errorf("error reading quota_limit_mb: %v", err)
		}
	}

	if o.Redial != nil {
		v := *o.Redial

		if err = d.Set("redial", v); err != nil {
			return diag.Errorf("error reading redial: %v", err)
		}
	}

	if o.RedundantIntf != nil {
		v := *o.RedundantIntf

		if err = d.Set("redundant_intf", v); err != nil {
			return diag.Errorf("error reading redundant_intf: %v", err)
		}
	}

	if o.Roaming != nil {
		v := *o.Roaming

		if err = d.Set("roaming", v); err != nil {
			return diag.Errorf("error reading roaming: %v", err)
		}
	}

	if o.Role != nil {
		v := *o.Role

		if err = d.Set("role", v); err != nil {
			return diag.Errorf("error reading role: %v", err)
		}
	}

	if o.SecondaryHa != nil {
		v := *o.SecondaryHa

		if err = d.Set("secondary_ha", v); err != nil {
			return diag.Errorf("error reading secondary_ha: %v", err)
		}
	}

	if o.SimPin != nil {
		v := *o.SimPin

		if v == "ENC XXXX" {
		} else if err = d.Set("sim_pin", v); err != nil {
			return diag.Errorf("error reading sim_pin: %v", err)
		}
	}

	if o.Vdom != nil {
		v := *o.Vdom

		if err = d.Set("vdom", v); err != nil {
			return diag.Errorf("error reading vdom: %v", err)
		}
	}

	if o.WanExtension != nil {
		if err = d.Set("wan_extension", flattenExtenderControllerExtenderWanExtension(o.WanExtension, sort)); err != nil {
			return diag.Errorf("error reading wan_extension: %v", err)
		}
	}

	if o.WimaxAuthProtocol != nil {
		v := *o.WimaxAuthProtocol

		if err = d.Set("wimax_auth_protocol", v); err != nil {
			return diag.Errorf("error reading wimax_auth_protocol: %v", err)
		}
	}

	if o.WimaxCarrier != nil {
		v := *o.WimaxCarrier

		if err = d.Set("wimax_carrier", v); err != nil {
			return diag.Errorf("error reading wimax_carrier: %v", err)
		}
	}

	if o.WimaxRealm != nil {
		v := *o.WimaxRealm

		if err = d.Set("wimax_realm", v); err != nil {
			return diag.Errorf("error reading wimax_realm: %v", err)
		}
	}

	return nil
}

func expandExtenderControllerExtenderControllerReport(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ExtenderControllerExtenderControllerReport, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ExtenderControllerExtenderControllerReport

	for i := range l {
		tmp := models.ExtenderControllerExtenderControllerReport{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Interval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.signal_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.SignalThreshold = &v2
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

func expandExtenderControllerExtenderModem1(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ExtenderControllerExtenderModem1, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ExtenderControllerExtenderModem1

	for i := range l {
		tmp := models.ExtenderControllerExtenderModem1{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.auto_switch", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandExtenderControllerExtenderModem1AutoSwitch(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ExtenderControllerExtenderModem1AutoSwitch
			// 	}
			tmp.AutoSwitch = v2

		}

		pre_append = fmt.Sprintf("%s.%d.conn_status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.ConnStatus = &v2
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

		pre_append = fmt.Sprintf("%s.%d.ifname", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ifname = &v2
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
	return &result, nil
}

func expandExtenderControllerExtenderModem1AutoSwitch(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ExtenderControllerExtenderModem1AutoSwitch, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ExtenderControllerExtenderModem1AutoSwitch

	for i := range l {
		tmp := models.ExtenderControllerExtenderModem1AutoSwitch{}
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
			if v2, ok := v1.(int64); ok {
				tmp.DisconnectPeriod = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.disconnect_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.DisconnectThreshold = &v2
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
			if v2, ok := v1.(int64); ok {
				tmp.SwitchBackTimer = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandExtenderControllerExtenderModem2(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ExtenderControllerExtenderModem2, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ExtenderControllerExtenderModem2

	for i := range l {
		tmp := models.ExtenderControllerExtenderModem2{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.auto_switch", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandExtenderControllerExtenderModem2AutoSwitch(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ExtenderControllerExtenderModem2AutoSwitch
			// 	}
			tmp.AutoSwitch = v2

		}

		pre_append = fmt.Sprintf("%s.%d.conn_status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.ConnStatus = &v2
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

		pre_append = fmt.Sprintf("%s.%d.ifname", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ifname = &v2
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
	return &result, nil
}

func expandExtenderControllerExtenderModem2AutoSwitch(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ExtenderControllerExtenderModem2AutoSwitch, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ExtenderControllerExtenderModem2AutoSwitch

	for i := range l {
		tmp := models.ExtenderControllerExtenderModem2AutoSwitch{}
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
			if v2, ok := v1.(int64); ok {
				tmp.DisconnectPeriod = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.disconnect_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.DisconnectThreshold = &v2
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
			if v2, ok := v1.(int64); ok {
				tmp.SwitchBackTimer = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandExtenderControllerExtenderWanExtension(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ExtenderControllerExtenderWanExtension, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ExtenderControllerExtenderWanExtension

	for i := range l {
		tmp := models.ExtenderControllerExtenderWanExtension{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.modem1_extension", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Modem1Extension = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.modem2_extension", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Modem2Extension = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectExtenderControllerExtender(d *schema.ResourceData, sv string) (*models.ExtenderControllerExtender, diag.Diagnostics) {
	obj := models.ExtenderControllerExtender{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("aaa_shared_secret"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("aaa_shared_secret", sv)
				diags = append(diags, e)
			}
			obj.AaaSharedSecret = &v2
		}
	}
	if v1, ok := d.GetOk("access_point_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("access_point_name", sv)
				diags = append(diags, e)
			}
			obj.AccessPointName = &v2
		}
	}
	if v1, ok := d.GetOk("admin"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("admin", sv)
				diags = append(diags, e)
			}
			obj.Admin = &v2
		}
	}
	if v1, ok := d.GetOk("allowaccess"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("allowaccess", sv)
				diags = append(diags, e)
			}
			obj.Allowaccess = &v2
		}
	}
	if v1, ok := d.GetOk("at_dial_script"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("at_dial_script", sv)
				diags = append(diags, e)
			}
			obj.AtDialScript = &v2
		}
	}
	if v1, ok := d.GetOk("authorized"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("authorized", sv)
				diags = append(diags, e)
			}
			obj.Authorized = &v2
		}
	}
	if v1, ok := d.GetOk("bandwidth_limit"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("bandwidth_limit", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.BandwidthLimit = &tmp
		}
	}
	if v1, ok := d.GetOk("billing_start_day"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("billing_start_day", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.BillingStartDay = &tmp
		}
	}
	if v1, ok := d.GetOk("cdma_aaa_spi"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("cdma_aaa_spi", sv)
				diags = append(diags, e)
			}
			obj.CdmaAaaSpi = &v2
		}
	}
	if v1, ok := d.GetOk("cdma_ha_spi"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("cdma_ha_spi", sv)
				diags = append(diags, e)
			}
			obj.CdmaHaSpi = &v2
		}
	}
	if v1, ok := d.GetOk("cdma_nai"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("cdma_nai", sv)
				diags = append(diags, e)
			}
			obj.CdmaNai = &v2
		}
	}
	if v1, ok := d.GetOk("conn_status"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("conn_status", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ConnStatus = &tmp
		}
	}
	if v, ok := d.GetOk("controller_report"); ok {
		if !utils.CheckVer(sv, "v6.4.2", "v7.0.2") {
			e := utils.AttributeVersionWarning("controller_report", sv)
			diags = append(diags, e)
		}
		t, err := expandExtenderControllerExtenderControllerReport(d, v, "controller_report", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ControllerReport = t
		}
	} else if d.HasChange("controller_report") {
		old, new := d.GetChange("controller_report")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ControllerReport = &[]models.ExtenderControllerExtenderControllerReport{}
		}
	}
	if v1, ok := d.GetOk("description"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("description", sv)
				diags = append(diags, e)
			}
			obj.Description = &v2
		}
	}
	if v1, ok := d.GetOk("device_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("device_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DeviceId = &tmp
		}
	}
	if v1, ok := d.GetOk("dial_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("dial_mode", sv)
				diags = append(diags, e)
			}
			obj.DialMode = &v2
		}
	}
	if v1, ok := d.GetOk("dial_status"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("dial_status", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DialStatus = &tmp
		}
	}
	if v1, ok := d.GetOk("enforce_bandwidth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("enforce_bandwidth", sv)
				diags = append(diags, e)
			}
			obj.EnforceBandwidth = &v2
		}
	}
	if v1, ok := d.GetOk("ext_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ext_name", sv)
				diags = append(diags, e)
			}
			obj.ExtName = &v2
		}
	}
	if v1, ok := d.GetOk("extension_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("extension_type", sv)
				diags = append(diags, e)
			}
			obj.ExtensionType = &v2
		}
	}
	if v1, ok := d.GetOk("ha_shared_secret"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("ha_shared_secret", sv)
				diags = append(diags, e)
			}
			obj.HaSharedSecret = &v2
		}
	}
	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			obj.Id = &v2
		}
	}
	if v1, ok := d.GetOk("ifname"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("ifname", sv)
				diags = append(diags, e)
			}
			obj.Ifname = &v2
		}
	}
	if v1, ok := d.GetOk("initiated_update"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("initiated_update", sv)
				diags = append(diags, e)
			}
			obj.InitiatedUpdate = &v2
		}
	}
	if v1, ok := d.GetOk("login_password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("login_password", sv)
				diags = append(diags, e)
			}
			obj.LoginPassword = &v2
		}
	}
	if v1, ok := d.GetOk("login_password_change"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("login_password_change", sv)
				diags = append(diags, e)
			}
			obj.LoginPasswordChange = &v2
		}
	}
	if v1, ok := d.GetOk("mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("mode", sv)
				diags = append(diags, e)
			}
			obj.Mode = &v2
		}
	}
	if v1, ok := d.GetOk("modem_passwd"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("modem_passwd", sv)
				diags = append(diags, e)
			}
			obj.ModemPasswd = &v2
		}
	}
	if v1, ok := d.GetOk("modem_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("modem_type", sv)
				diags = append(diags, e)
			}
			obj.ModemType = &v2
		}
	}
	if v, ok := d.GetOk("modem1"); ok {
		if !utils.CheckVer(sv, "v6.4.2", "v7.0.2") {
			e := utils.AttributeVersionWarning("modem1", sv)
			diags = append(diags, e)
		}
		t, err := expandExtenderControllerExtenderModem1(d, v, "modem1", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Modem1 = t
		}
	} else if d.HasChange("modem1") {
		old, new := d.GetChange("modem1")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Modem1 = &[]models.ExtenderControllerExtenderModem1{}
		}
	}
	if v, ok := d.GetOk("modem2"); ok {
		if !utils.CheckVer(sv, "v6.4.2", "v7.0.2") {
			e := utils.AttributeVersionWarning("modem2", sv)
			diags = append(diags, e)
		}
		t, err := expandExtenderControllerExtenderModem2(d, v, "modem2", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Modem2 = t
		}
	} else if d.HasChange("modem2") {
		old, new := d.GetChange("modem2")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Modem2 = &[]models.ExtenderControllerExtenderModem2{}
		}
	}
	if v1, ok := d.GetOk("multi_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("multi_mode", sv)
				diags = append(diags, e)
			}
			obj.MultiMode = &v2
		}
	}
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v1, ok := d.GetOk("override_allowaccess"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("override_allowaccess", sv)
				diags = append(diags, e)
			}
			obj.OverrideAllowaccess = &v2
		}
	}
	if v1, ok := d.GetOk("override_enforce_bandwidth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("override_enforce_bandwidth", sv)
				diags = append(diags, e)
			}
			obj.OverrideEnforceBandwidth = &v2
		}
	}
	if v1, ok := d.GetOk("override_login_password_change"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("override_login_password_change", sv)
				diags = append(diags, e)
			}
			obj.OverrideLoginPasswordChange = &v2
		}
	}
	if v1, ok := d.GetOk("ppp_auth_protocol"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("ppp_auth_protocol", sv)
				diags = append(diags, e)
			}
			obj.PppAuthProtocol = &v2
		}
	}
	if v1, ok := d.GetOk("ppp_echo_request"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("ppp_echo_request", sv)
				diags = append(diags, e)
			}
			obj.PppEchoRequest = &v2
		}
	}
	if v1, ok := d.GetOk("ppp_password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("ppp_password", sv)
				diags = append(diags, e)
			}
			obj.PppPassword = &v2
		}
	}
	if v1, ok := d.GetOk("ppp_username"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("ppp_username", sv)
				diags = append(diags, e)
			}
			obj.PppUsername = &v2
		}
	}
	if v1, ok := d.GetOk("primary_ha"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("primary_ha", sv)
				diags = append(diags, e)
			}
			obj.PrimaryHa = &v2
		}
	}
	if v1, ok := d.GetOk("profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("profile", sv)
				diags = append(diags, e)
			}
			obj.Profile = &v2
		}
	}
	if v1, ok := d.GetOk("quota_limit_mb"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("quota_limit_mb", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.QuotaLimitMb = &tmp
		}
	}
	if v1, ok := d.GetOk("redial"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("redial", sv)
				diags = append(diags, e)
			}
			obj.Redial = &v2
		}
	}
	if v1, ok := d.GetOk("redundant_intf"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("redundant_intf", sv)
				diags = append(diags, e)
			}
			obj.RedundantIntf = &v2
		}
	}
	if v1, ok := d.GetOk("roaming"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("roaming", sv)
				diags = append(diags, e)
			}
			obj.Roaming = &v2
		}
	}
	if v1, ok := d.GetOk("role"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("role", sv)
				diags = append(diags, e)
			}
			obj.Role = &v2
		}
	}
	if v1, ok := d.GetOk("secondary_ha"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("secondary_ha", sv)
				diags = append(diags, e)
			}
			obj.SecondaryHa = &v2
		}
	}
	if v1, ok := d.GetOk("sim_pin"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("sim_pin", sv)
				diags = append(diags, e)
			}
			obj.SimPin = &v2
		}
	}
	if v1, ok := d.GetOk("vdom"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vdom", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Vdom = &tmp
		}
	}
	if v, ok := d.GetOk("wan_extension"); ok {
		if !utils.CheckVer(sv, "v7.0.2", "") {
			e := utils.AttributeVersionWarning("wan_extension", sv)
			diags = append(diags, e)
		}
		t, err := expandExtenderControllerExtenderWanExtension(d, v, "wan_extension", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.WanExtension = t
		}
	} else if d.HasChange("wan_extension") {
		old, new := d.GetChange("wan_extension")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.WanExtension = &[]models.ExtenderControllerExtenderWanExtension{}
		}
	}
	if v1, ok := d.GetOk("wimax_auth_protocol"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("wimax_auth_protocol", sv)
				diags = append(diags, e)
			}
			obj.WimaxAuthProtocol = &v2
		}
	}
	if v1, ok := d.GetOk("wimax_carrier"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("wimax_carrier", sv)
				diags = append(diags, e)
			}
			obj.WimaxCarrier = &v2
		}
	}
	if v1, ok := d.GetOk("wimax_realm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("wimax_realm", sv)
				diags = append(diags, e)
			}
			obj.WimaxRealm = &v2
		}
	}
	return &obj, diags
}
