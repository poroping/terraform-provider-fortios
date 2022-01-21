// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiManager.

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

func resourceSystemFortimanager() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiManager.",

		CreateContext: resourceSystemFortimanagerCreate,
		ReadContext:   resourceSystemFortimanagerRead,
		UpdateContext: resourceSystemFortimanagerUpdate,
		DeleteContext: resourceSystemFortimanagerDelete,

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
			"central_management": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiManager central management.",
				Optional:    true,
				Computed:    true,
			},
			"central_mgmt_auto_backup": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable central management auto backup.",
				Optional:    true,
				Computed:    true,
			},
			"central_mgmt_schedule_config_restore": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable central management schedule config restore.",
				Optional:    true,
				Computed:    true,
			},
			"central_mgmt_schedule_script_restore": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable central management schedule script restore.",
				Optional:    true,
				Computed:    true,
			},
			"ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address.",
				Optional:    true,
				Computed:    true,
			},
			"ipsec": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiManager IPsec tunnel.",
				Optional:    true,
				Computed:    true,
			},
			"vdom": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "Virtual domain name.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemFortimanagerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemFortimanager(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemFortimanager(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemFortimanager")
	}

	return resourceSystemFortimanagerRead(ctx, d, meta)
}

func resourceSystemFortimanagerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemFortimanager(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemFortimanager(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemFortimanager resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemFortimanager")
	}

	return resourceSystemFortimanagerRead(ctx, d, meta)
}

func resourceSystemFortimanagerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemFortimanager(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemFortimanager(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemFortimanager resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFortimanagerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemFortimanager(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemFortimanager resource: %v", err)
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

	diags := refreshObjectSystemFortimanager(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemFortimanager(d *schema.ResourceData, o *models.SystemFortimanager, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.CentralManagement != nil {
		v := *o.CentralManagement

		if err = d.Set("central_management", v); err != nil {
			return diag.Errorf("error reading central_management: %v", err)
		}
	}

	if o.CentralMgmtAutoBackup != nil {
		v := *o.CentralMgmtAutoBackup

		if err = d.Set("central_mgmt_auto_backup", v); err != nil {
			return diag.Errorf("error reading central_mgmt_auto_backup: %v", err)
		}
	}

	if o.CentralMgmtScheduleConfigRestore != nil {
		v := *o.CentralMgmtScheduleConfigRestore

		if err = d.Set("central_mgmt_schedule_config_restore", v); err != nil {
			return diag.Errorf("error reading central_mgmt_schedule_config_restore: %v", err)
		}
	}

	if o.CentralMgmtScheduleScriptRestore != nil {
		v := *o.CentralMgmtScheduleScriptRestore

		if err = d.Set("central_mgmt_schedule_script_restore", v); err != nil {
			return diag.Errorf("error reading central_mgmt_schedule_script_restore: %v", err)
		}
	}

	if o.Ip != nil {
		v := *o.Ip

		if err = d.Set("ip", v); err != nil {
			return diag.Errorf("error reading ip: %v", err)
		}
	}

	if o.Ipsec != nil {
		v := *o.Ipsec

		if err = d.Set("ipsec", v); err != nil {
			return diag.Errorf("error reading ipsec: %v", err)
		}
	}

	if o.Vdom != nil {
		v := *o.Vdom

		if err = d.Set("vdom", v); err != nil {
			return diag.Errorf("error reading vdom: %v", err)
		}
	}

	return nil
}

func getObjectSystemFortimanager(d *schema.ResourceData, sv string) (*models.SystemFortimanager, diag.Diagnostics) {
	obj := models.SystemFortimanager{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("central_management"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("central_management", sv)
				diags = append(diags, e)
			}
			obj.CentralManagement = &v2
		}
	}
	if v1, ok := d.GetOk("central_mgmt_auto_backup"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("central_mgmt_auto_backup", sv)
				diags = append(diags, e)
			}
			obj.CentralMgmtAutoBackup = &v2
		}
	}
	if v1, ok := d.GetOk("central_mgmt_schedule_config_restore"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("central_mgmt_schedule_config_restore", sv)
				diags = append(diags, e)
			}
			obj.CentralMgmtScheduleConfigRestore = &v2
		}
	}
	if v1, ok := d.GetOk("central_mgmt_schedule_script_restore"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("central_mgmt_schedule_script_restore", sv)
				diags = append(diags, e)
			}
			obj.CentralMgmtScheduleScriptRestore = &v2
		}
	}
	if v1, ok := d.GetOk("ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip", sv)
				diags = append(diags, e)
			}
			obj.Ip = &v2
		}
	}
	if v1, ok := d.GetOk("ipsec"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipsec", sv)
				diags = append(diags, e)
			}
			obj.Ipsec = &v2
		}
	}
	if v1, ok := d.GetOk("vdom"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vdom", sv)
				diags = append(diags, e)
			}
			obj.Vdom = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemFortimanager(d *schema.ResourceData, sv string) (*models.SystemFortimanager, diag.Diagnostics) {
	obj := models.SystemFortimanager{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
