// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure access profiles for system administrators.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceSystemAccprofile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure access profiles for system administrators.",

		CreateContext: resourceSystemAccprofileCreate,
		ReadContext:   resourceSystemAccprofileRead,
		UpdateContext: resourceSystemAccprofileUpdate,
		DeleteContext: resourceSystemAccprofileDelete,

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
			"admintimeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 480),

				Description: "Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).",
				Optional:    true,
				Computed:    true,
			},
			"admintimeout_override": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable overriding the global administrator idle timeout.",
				Optional:    true,
				Computed:    true,
			},
			"authgrp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

				Description: "Administrator access to Users and Devices.",
				Optional:    true,
				Computed:    true,
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"ftviewgrp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

				Description: "FortiView.",
				Optional:    true,
				Computed:    true,
			},
			"fwgrp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write", "custom"}, false),

				Description: "Administrator access to the Firewall configuration.",
				Optional:    true,
				Computed:    true,
			},
			"fwgrp_permission": {
				Type:        schema.TypeList,
				Description: "Custom firewall permission.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "Address Configuration.",
							Optional:    true,
							Computed:    true,
						},
						"others": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "Other Firewall Configuration.",
							Optional:    true,
							Computed:    true,
						},
						"policy": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "Policy Configuration.",
							Optional:    true,
							Computed:    true,
						},
						"schedule": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "Schedule Configuration.",
							Optional:    true,
							Computed:    true,
						},
						"service": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "Service Configuration.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"loggrp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write", "custom"}, false),

				Description: "Administrator access to Logging and Reporting including viewing log messages.",
				Optional:    true,
				Computed:    true,
			},
			"loggrp_permission": {
				Type:        schema.TypeList,
				Description: "Custom Log & Report permission.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "Log & Report configuration.",
							Optional:    true,
							Computed:    true,
						},
						"data_access": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "Log & Report Data Access.",
							Optional:    true,
							Computed:    true,
						},
						"report_access": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "Log & Report Report Access.",
							Optional:    true,
							Computed:    true,
						},
						"threat_weight": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "Log & Report Threat Weight.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Profile name.",
				ForceNew:    true,
				Required:    true,
			},
			"netgrp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write", "custom"}, false),

				Description: "Network Configuration.",
				Optional:    true,
				Computed:    true,
			},
			"netgrp_permission": {
				Type:        schema.TypeList,
				Description: "Custom network permission.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cfg": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "Network Configuration.",
							Optional:    true,
							Computed:    true,
						},
						"packet_capture": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "Packet Capture Configuration.",
							Optional:    true,
							Computed:    true,
						},
						"route_cfg": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "Router Configuration.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"scope": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"vdom", "global"}, false),

				Description: "Scope of admin access: global or specific VDOM(s).",
				Optional:    true,
				Computed:    true,
			},
			"secfabgrp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

				Description: "Security Fabric.",
				Optional:    true,
				Computed:    true,
			},
			"sysgrp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write", "custom"}, false),

				Description: "System Configuration.",
				Optional:    true,
				Computed:    true,
			},
			"sysgrp_permission": {
				Type:        schema.TypeList,
				Description: "Custom system permission.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "Administrator Users.",
							Optional:    true,
							Computed:    true,
						},
						"cfg": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "System Configuration.",
							Optional:    true,
							Computed:    true,
						},
						"mnt": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "Maintenance.",
							Optional:    true,
							Computed:    true,
						},
						"upd": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "FortiGuard Updates.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"system_diagnostics": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable permission to run system diagnostic commands.",
				Optional:    true,
				Computed:    true,
			},
			"utmgrp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write", "custom"}, false),

				Description: "Administrator access to Security Profiles.",
				Optional:    true,
				Computed:    true,
			},
			"utmgrp_permission": {
				Type:        schema.TypeList,
				Description: "Custom Security Profile permissions.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"antivirus": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "Antivirus profiles and settings.",
							Optional:    true,
							Computed:    true,
						},
						"application_control": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "Application Control profiles and settings.",
							Optional:    true,
							Computed:    true,
						},
						"data_loss_prevention": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "DLP profiles and settings.",
							Optional:    true,
							Computed:    true,
						},
						"dnsfilter": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "DNS Filter profiles and settings.",
							Optional:    true,
							Computed:    true,
						},
						"emailfilter": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "Email Filter and settings.",
							Optional:    true,
							Computed:    true,
						},
						"endpoint_control": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "FortiClient Profiles.",
							Optional:    true,
							Computed:    true,
						},
						"file_filter": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "File-filter profiles and settings.",
							Optional:    true,
							Computed:    true,
						},
						"icap": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "ICAP profiles and settings.",
							Optional:    true,
							Computed:    true,
						},
						"ips": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "IPS profiles and settings.",
							Optional:    true,
							Computed:    true,
						},
						"voip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "VoIP profiles and settings.",
							Optional:    true,
							Computed:    true,
						},
						"waf": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "Web Application Firewall profiles and settings.",
							Optional:    true,
							Computed:    true,
						},
						"webfilter": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

							Description: "Web Filter profiles and settings.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"vpngrp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

				Description: "Administrator access to IPsec, SSL, PPTP, and L2TP VPN.",
				Optional:    true,
				Computed:    true,
			},
			"wanoptgrp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

				Description: "Administrator access to WAN Opt & Cache.",
				Optional:    true,
				Computed:    true,
			},
			"wifi": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "read", "read-write"}, false),

				Description: "Administrator access to the WiFi controller and Switch controller.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemAccprofileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemAccprofile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemAccprofile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemAccprofile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemAccprofile")
	}

	return resourceSystemAccprofileRead(ctx, d, meta)
}

func resourceSystemAccprofileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemAccprofile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemAccprofile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemAccprofile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemAccprofile")
	}

	return resourceSystemAccprofileRead(ctx, d, meta)
}

func resourceSystemAccprofileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemAccprofile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemAccprofile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAccprofileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemAccprofile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemAccprofile resource: %v", err)
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

	diags := refreshObjectSystemAccprofile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemAccprofileFwgrpPermission(v *[]models.SystemAccprofileFwgrpPermission, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Address; tmp != nil {
				v["address"] = *tmp
			}

			if tmp := cfg.Others; tmp != nil {
				v["others"] = *tmp
			}

			if tmp := cfg.Policy; tmp != nil {
				v["policy"] = *tmp
			}

			if tmp := cfg.Schedule; tmp != nil {
				v["schedule"] = *tmp
			}

			if tmp := cfg.Service; tmp != nil {
				v["service"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemAccprofileLoggrpPermission(v *[]models.SystemAccprofileLoggrpPermission, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Config; tmp != nil {
				v["config"] = *tmp
			}

			if tmp := cfg.DataAccess; tmp != nil {
				v["data_access"] = *tmp
			}

			if tmp := cfg.ReportAccess; tmp != nil {
				v["report_access"] = *tmp
			}

			if tmp := cfg.ThreatWeight; tmp != nil {
				v["threat_weight"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemAccprofileNetgrpPermission(v *[]models.SystemAccprofileNetgrpPermission, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Cfg; tmp != nil {
				v["cfg"] = *tmp
			}

			if tmp := cfg.PacketCapture; tmp != nil {
				v["packet_capture"] = *tmp
			}

			if tmp := cfg.RouteCfg; tmp != nil {
				v["route_cfg"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemAccprofileSysgrpPermission(v *[]models.SystemAccprofileSysgrpPermission, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Admin; tmp != nil {
				v["admin"] = *tmp
			}

			if tmp := cfg.Cfg; tmp != nil {
				v["cfg"] = *tmp
			}

			if tmp := cfg.Mnt; tmp != nil {
				v["mnt"] = *tmp
			}

			if tmp := cfg.Upd; tmp != nil {
				v["upd"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSystemAccprofileUtmgrpPermission(v *[]models.SystemAccprofileUtmgrpPermission, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Antivirus; tmp != nil {
				v["antivirus"] = *tmp
			}

			if tmp := cfg.ApplicationControl; tmp != nil {
				v["application_control"] = *tmp
			}

			if tmp := cfg.DataLossPrevention; tmp != nil {
				v["data_loss_prevention"] = *tmp
			}

			if tmp := cfg.Dnsfilter; tmp != nil {
				v["dnsfilter"] = *tmp
			}

			if tmp := cfg.Emailfilter; tmp != nil {
				v["emailfilter"] = *tmp
			}

			if tmp := cfg.EndpointControl; tmp != nil {
				v["endpoint_control"] = *tmp
			}

			if tmp := cfg.FileFilter; tmp != nil {
				v["file_filter"] = *tmp
			}

			if tmp := cfg.Icap; tmp != nil {
				v["icap"] = *tmp
			}

			if tmp := cfg.Ips; tmp != nil {
				v["ips"] = *tmp
			}

			if tmp := cfg.Voip; tmp != nil {
				v["voip"] = *tmp
			}

			if tmp := cfg.Waf; tmp != nil {
				v["waf"] = *tmp
			}

			if tmp := cfg.Webfilter; tmp != nil {
				v["webfilter"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func refreshObjectSystemAccprofile(d *schema.ResourceData, o *models.SystemAccprofile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Admintimeout != nil {
		v := *o.Admintimeout

		if err = d.Set("admintimeout", v); err != nil {
			return diag.Errorf("error reading admintimeout: %v", err)
		}
	}

	if o.AdmintimeoutOverride != nil {
		v := *o.AdmintimeoutOverride

		if err = d.Set("admintimeout_override", v); err != nil {
			return diag.Errorf("error reading admintimeout_override: %v", err)
		}
	}

	if o.Authgrp != nil {
		v := *o.Authgrp

		if err = d.Set("authgrp", v); err != nil {
			return diag.Errorf("error reading authgrp: %v", err)
		}
	}

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.Ftviewgrp != nil {
		v := *o.Ftviewgrp

		if err = d.Set("ftviewgrp", v); err != nil {
			return diag.Errorf("error reading ftviewgrp: %v", err)
		}
	}

	if o.Fwgrp != nil {
		v := *o.Fwgrp

		if err = d.Set("fwgrp", v); err != nil {
			return diag.Errorf("error reading fwgrp: %v", err)
		}
	}

	if o.FwgrpPermission != nil {
		if err = d.Set("fwgrp_permission", flattenSystemAccprofileFwgrpPermission(o.FwgrpPermission, sort)); err != nil {
			return diag.Errorf("error reading fwgrp_permission: %v", err)
		}
	}

	if o.Loggrp != nil {
		v := *o.Loggrp

		if err = d.Set("loggrp", v); err != nil {
			return diag.Errorf("error reading loggrp: %v", err)
		}
	}

	if o.LoggrpPermission != nil {
		if err = d.Set("loggrp_permission", flattenSystemAccprofileLoggrpPermission(o.LoggrpPermission, sort)); err != nil {
			return diag.Errorf("error reading loggrp_permission: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Netgrp != nil {
		v := *o.Netgrp

		if err = d.Set("netgrp", v); err != nil {
			return diag.Errorf("error reading netgrp: %v", err)
		}
	}

	if o.NetgrpPermission != nil {
		if err = d.Set("netgrp_permission", flattenSystemAccprofileNetgrpPermission(o.NetgrpPermission, sort)); err != nil {
			return diag.Errorf("error reading netgrp_permission: %v", err)
		}
	}

	if o.Scope != nil {
		v := *o.Scope

		if err = d.Set("scope", v); err != nil {
			return diag.Errorf("error reading scope: %v", err)
		}
	}

	if o.Secfabgrp != nil {
		v := *o.Secfabgrp

		if err = d.Set("secfabgrp", v); err != nil {
			return diag.Errorf("error reading secfabgrp: %v", err)
		}
	}

	if o.Sysgrp != nil {
		v := *o.Sysgrp

		if err = d.Set("sysgrp", v); err != nil {
			return diag.Errorf("error reading sysgrp: %v", err)
		}
	}

	if o.SysgrpPermission != nil {
		if err = d.Set("sysgrp_permission", flattenSystemAccprofileSysgrpPermission(o.SysgrpPermission, sort)); err != nil {
			return diag.Errorf("error reading sysgrp_permission: %v", err)
		}
	}

	if o.SystemDiagnostics != nil {
		v := *o.SystemDiagnostics

		if err = d.Set("system_diagnostics", v); err != nil {
			return diag.Errorf("error reading system_diagnostics: %v", err)
		}
	}

	if o.Utmgrp != nil {
		v := *o.Utmgrp

		if err = d.Set("utmgrp", v); err != nil {
			return diag.Errorf("error reading utmgrp: %v", err)
		}
	}

	if o.UtmgrpPermission != nil {
		if err = d.Set("utmgrp_permission", flattenSystemAccprofileUtmgrpPermission(o.UtmgrpPermission, sort)); err != nil {
			return diag.Errorf("error reading utmgrp_permission: %v", err)
		}
	}

	if o.Vpngrp != nil {
		v := *o.Vpngrp

		if err = d.Set("vpngrp", v); err != nil {
			return diag.Errorf("error reading vpngrp: %v", err)
		}
	}

	if o.Wanoptgrp != nil {
		v := *o.Wanoptgrp

		if err = d.Set("wanoptgrp", v); err != nil {
			return diag.Errorf("error reading wanoptgrp: %v", err)
		}
	}

	if o.Wifi != nil {
		v := *o.Wifi

		if err = d.Set("wifi", v); err != nil {
			return diag.Errorf("error reading wifi: %v", err)
		}
	}

	return nil
}

func expandSystemAccprofileFwgrpPermission(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemAccprofileFwgrpPermission, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemAccprofileFwgrpPermission

	for i := range l {
		tmp := models.SystemAccprofileFwgrpPermission{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Address = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.others", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Others = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.policy", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Policy = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.schedule", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Schedule = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.service", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Service = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemAccprofileLoggrpPermission(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemAccprofileLoggrpPermission, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemAccprofileLoggrpPermission

	for i := range l {
		tmp := models.SystemAccprofileLoggrpPermission{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.config", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Config = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.data_access", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DataAccess = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.report_access", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ReportAccess = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.threat_weight", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ThreatWeight = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemAccprofileNetgrpPermission(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemAccprofileNetgrpPermission, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemAccprofileNetgrpPermission

	for i := range l {
		tmp := models.SystemAccprofileNetgrpPermission{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.cfg", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Cfg = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.packet_capture", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PacketCapture = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_cfg", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteCfg = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemAccprofileSysgrpPermission(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemAccprofileSysgrpPermission, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemAccprofileSysgrpPermission

	for i := range l {
		tmp := models.SystemAccprofileSysgrpPermission{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.admin", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Admin = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cfg", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Cfg = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mnt", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mnt = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.upd", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Upd = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemAccprofileUtmgrpPermission(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemAccprofileUtmgrpPermission, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemAccprofileUtmgrpPermission

	for i := range l {
		tmp := models.SystemAccprofileUtmgrpPermission{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.antivirus", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Antivirus = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.application_control", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ApplicationControl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.data_loss_prevention", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DataLossPrevention = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dnsfilter", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dnsfilter = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.emailfilter", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Emailfilter = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.endpoint_control", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EndpointControl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.file_filter", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FileFilter = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.icap", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Icap = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ips", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ips = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.voip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Voip = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.waf", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Waf = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.webfilter", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Webfilter = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemAccprofile(d *schema.ResourceData, sv string) (*models.SystemAccprofile, diag.Diagnostics) {
	obj := models.SystemAccprofile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("admintimeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admintimeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Admintimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("admintimeout_override"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("admintimeout_override", sv)
				diags = append(diags, e)
			}
			obj.AdmintimeoutOverride = &v2
		}
	}
	if v1, ok := d.GetOk("authgrp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authgrp", sv)
				diags = append(diags, e)
			}
			obj.Authgrp = &v2
		}
	}
	if v1, ok := d.GetOk("comments"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comments", sv)
				diags = append(diags, e)
			}
			obj.Comments = &v2
		}
	}
	if v1, ok := d.GetOk("ftviewgrp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ftviewgrp", sv)
				diags = append(diags, e)
			}
			obj.Ftviewgrp = &v2
		}
	}
	if v1, ok := d.GetOk("fwgrp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fwgrp", sv)
				diags = append(diags, e)
			}
			obj.Fwgrp = &v2
		}
	}
	if v, ok := d.GetOk("fwgrp_permission"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("fwgrp_permission", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemAccprofileFwgrpPermission(d, v, "fwgrp_permission", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.FwgrpPermission = t
		}
	} else if d.HasChange("fwgrp_permission") {
		old, new := d.GetChange("fwgrp_permission")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.FwgrpPermission = &[]models.SystemAccprofileFwgrpPermission{}
		}
	}
	if v1, ok := d.GetOk("loggrp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("loggrp", sv)
				diags = append(diags, e)
			}
			obj.Loggrp = &v2
		}
	}
	if v, ok := d.GetOk("loggrp_permission"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("loggrp_permission", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemAccprofileLoggrpPermission(d, v, "loggrp_permission", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.LoggrpPermission = t
		}
	} else if d.HasChange("loggrp_permission") {
		old, new := d.GetChange("loggrp_permission")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.LoggrpPermission = &[]models.SystemAccprofileLoggrpPermission{}
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
	if v1, ok := d.GetOk("netgrp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("netgrp", sv)
				diags = append(diags, e)
			}
			obj.Netgrp = &v2
		}
	}
	if v, ok := d.GetOk("netgrp_permission"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("netgrp_permission", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemAccprofileNetgrpPermission(d, v, "netgrp_permission", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.NetgrpPermission = t
		}
	} else if d.HasChange("netgrp_permission") {
		old, new := d.GetChange("netgrp_permission")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.NetgrpPermission = &[]models.SystemAccprofileNetgrpPermission{}
		}
	}
	if v1, ok := d.GetOk("scope"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("scope", sv)
				diags = append(diags, e)
			}
			obj.Scope = &v2
		}
	}
	if v1, ok := d.GetOk("secfabgrp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("secfabgrp", sv)
				diags = append(diags, e)
			}
			obj.Secfabgrp = &v2
		}
	}
	if v1, ok := d.GetOk("sysgrp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sysgrp", sv)
				diags = append(diags, e)
			}
			obj.Sysgrp = &v2
		}
	}
	if v, ok := d.GetOk("sysgrp_permission"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("sysgrp_permission", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemAccprofileSysgrpPermission(d, v, "sysgrp_permission", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SysgrpPermission = t
		}
	} else if d.HasChange("sysgrp_permission") {
		old, new := d.GetChange("sysgrp_permission")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SysgrpPermission = &[]models.SystemAccprofileSysgrpPermission{}
		}
	}
	if v1, ok := d.GetOk("system_diagnostics"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("system_diagnostics", sv)
				diags = append(diags, e)
			}
			obj.SystemDiagnostics = &v2
		}
	}
	if v1, ok := d.GetOk("utmgrp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("utmgrp", sv)
				diags = append(diags, e)
			}
			obj.Utmgrp = &v2
		}
	}
	if v, ok := d.GetOk("utmgrp_permission"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("utmgrp_permission", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemAccprofileUtmgrpPermission(d, v, "utmgrp_permission", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.UtmgrpPermission = t
		}
	} else if d.HasChange("utmgrp_permission") {
		old, new := d.GetChange("utmgrp_permission")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.UtmgrpPermission = &[]models.SystemAccprofileUtmgrpPermission{}
		}
	}
	if v1, ok := d.GetOk("vpngrp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vpngrp", sv)
				diags = append(diags, e)
			}
			obj.Vpngrp = &v2
		}
	}
	if v1, ok := d.GetOk("wanoptgrp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wanoptgrp", sv)
				diags = append(diags, e)
			}
			obj.Wanoptgrp = &v2
		}
	}
	if v1, ok := d.GetOk("wifi"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wifi", sv)
				diags = append(diags, e)
			}
			obj.Wifi = &v2
		}
	}
	return &obj, diags
}
