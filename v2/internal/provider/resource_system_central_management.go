// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure central management.

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

func resourceSystemCentralManagement() *schema.Resource {
	return &schema.Resource{
		Description: "Configure central management.",

		CreateContext: resourceSystemCentralManagementCreate,
		ReadContext:   resourceSystemCentralManagementRead,
		UpdateContext: resourceSystemCentralManagementUpdate,
		DeleteContext: resourceSystemCentralManagementDelete,

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
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"allow_monitor": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable allowing the central management server to remotely monitor this FortiGate unit.",
				Optional:    true,
				Computed:    true,
			},
			"allow_push_configuration": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable allowing the central management server to push configuration changes to this FortiGate.",
				Optional:    true,
				Computed:    true,
			},
			"allow_push_firmware": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable allowing the central management server to push firmware updates to this FortiGate.",
				Optional:    true,
				Computed:    true,
			},
			"allow_remote_firmware_upgrade": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable remotely upgrading the firmware on this FortiGate from the central management server.",
				Optional:    true,
				Computed:    true,
			},
			"ca_cert": {
				Type: schema.TypeString,

				Description: "CA certificate to be used by FGFM protocol.",
				Optional:    true,
				Computed:    true,
			},
			"enc_algorithm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"default", "high", "low"}, false),

				Description: "Encryption strength for communications between the FortiGate and central management.",
				Optional:    true,
				Computed:    true,
			},
			"fmg": {
				Type: schema.TypeString,

				Description: "IP address or FQDN of the FortiManager.",
				Optional:    true,
				Computed:    true,
			},
			"fmg_source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IPv4 source address that this FortiGate uses when communicating with FortiManager.",
				Optional:    true,
				Computed:    true,
			},
			"fmg_source_ip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 source address that this FortiGate uses when communicating with FortiManager.",
				Optional:         true,
				Computed:         true,
			},
			"fmg_update_port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"8890", "443"}, false),

				Description: "Port used to communicate with FortiManager that is acting as a FortiGuard update server.",
				Optional:    true,
				Computed:    true,
			},
			"include_default_servers": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable inclusion of public FortiGuard servers in the override server list.",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Specify outgoing interface to reach server.",
				Optional:    true,
				Computed:    true,
			},
			"interface_select_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "sdwan", "specify"}, false),

				Description: "Specify how to select outgoing interface to reach server.",
				Optional:    true,
				Computed:    true,
			},
			"local_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Certificate to be used by FGFM protocol.",
				Optional:    true,
				Computed:    true,
			},
			"mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"normal", "backup"}, false),

				Description: "Central management mode.",
				Optional:    true,
				Computed:    true,
			},
			"schedule_config_restore": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable allowing the central management server to restore the configuration of this FortiGate.",
				Optional:    true,
				Computed:    true,
			},
			"schedule_script_restore": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable allowing the central management server to restore the scripts stored on this FortiGate.",
				Optional:    true,
				Computed:    true,
			},
			"serial_number": {
				Type: schema.TypeString,

				Description: "Serial number.",
				Optional:    true,
				Computed:    true,
			},
			"server_list": {
				Type:        schema.TypeList,
				Description: "Additional severs that the FortiGate can use for updates (for AV, IPS, updates) and ratings (for web filter and antispam ratings) servers.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ipv4", "ipv6", "fqdn"}, false),

							Description: "Indicate whether the FortiGate communicates with the override server using an IPv4 address, an IPv6 address or a FQDN.",
							Optional:    true,
							Computed:    true,
						},
						"fqdn": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "FQDN address of override server.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"server_address": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "IPv4 address of override server.",
							Optional:    true,
							Computed:    true,
						},
						"server_address6": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "IPv6 address of override server.",
							Optional:         true,
							Computed:         true,
						},
						"server_type": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "FortiGuard service type.",
							Optional:         true,
							Computed:         true,
						},
					},
				},
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"fortimanager", "fortiguard", "none"}, false),

				Description: "Central management type.",
				Optional:    true,
				Computed:    true,
			},
			"vdom": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "Virtual domain (VDOM) name to use when communicating with FortiManager.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemCentralManagementCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemCentralManagement(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemCentralManagement(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemCentralManagement")
	}

	return resourceSystemCentralManagementRead(ctx, d, meta)
}

func resourceSystemCentralManagementUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemCentralManagement(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemCentralManagement(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemCentralManagement resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemCentralManagement")
	}

	return resourceSystemCentralManagementRead(ctx, d, meta)
}

func resourceSystemCentralManagementDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemCentralManagement(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemCentralManagement(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemCentralManagement resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCentralManagementRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemCentralManagement(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemCentralManagement resource: %v", err)
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

	diags := refreshObjectSystemCentralManagement(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemCentralManagementServerList(d *schema.ResourceData, v *[]models.SystemCentralManagementServerList, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AddrType; tmp != nil {
				v["addr_type"] = *tmp
			}

			if tmp := cfg.Fqdn; tmp != nil {
				v["fqdn"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.ServerAddress; tmp != nil {
				v["server_address"] = *tmp
			}

			if tmp := cfg.ServerAddress6; tmp != nil {
				v["server_address6"] = *tmp
			}

			if tmp := cfg.ServerType; tmp != nil {
				v["server_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectSystemCentralManagement(d *schema.ResourceData, o *models.SystemCentralManagement, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AllowMonitor != nil {
		v := *o.AllowMonitor

		if err = d.Set("allow_monitor", v); err != nil {
			return diag.Errorf("error reading allow_monitor: %v", err)
		}
	}

	if o.AllowPushConfiguration != nil {
		v := *o.AllowPushConfiguration

		if err = d.Set("allow_push_configuration", v); err != nil {
			return diag.Errorf("error reading allow_push_configuration: %v", err)
		}
	}

	if o.AllowPushFirmware != nil {
		v := *o.AllowPushFirmware

		if err = d.Set("allow_push_firmware", v); err != nil {
			return diag.Errorf("error reading allow_push_firmware: %v", err)
		}
	}

	if o.AllowRemoteFirmwareUpgrade != nil {
		v := *o.AllowRemoteFirmwareUpgrade

		if err = d.Set("allow_remote_firmware_upgrade", v); err != nil {
			return diag.Errorf("error reading allow_remote_firmware_upgrade: %v", err)
		}
	}

	if o.CaCert != nil {
		v := *o.CaCert

		if err = d.Set("ca_cert", v); err != nil {
			return diag.Errorf("error reading ca_cert: %v", err)
		}
	}

	if o.EncAlgorithm != nil {
		v := *o.EncAlgorithm

		if err = d.Set("enc_algorithm", v); err != nil {
			return diag.Errorf("error reading enc_algorithm: %v", err)
		}
	}

	if o.Fmg != nil {
		v := *o.Fmg

		if err = d.Set("fmg", v); err != nil {
			return diag.Errorf("error reading fmg: %v", err)
		}
	}

	if o.FmgSourceIp != nil {
		v := *o.FmgSourceIp

		if err = d.Set("fmg_source_ip", v); err != nil {
			return diag.Errorf("error reading fmg_source_ip: %v", err)
		}
	}

	if o.FmgSourceIp6 != nil {
		v := *o.FmgSourceIp6

		if err = d.Set("fmg_source_ip6", v); err != nil {
			return diag.Errorf("error reading fmg_source_ip6: %v", err)
		}
	}

	if o.FmgUpdatePort != nil {
		v := *o.FmgUpdatePort

		if err = d.Set("fmg_update_port", v); err != nil {
			return diag.Errorf("error reading fmg_update_port: %v", err)
		}
	}

	if o.IncludeDefaultServers != nil {
		v := *o.IncludeDefaultServers

		if err = d.Set("include_default_servers", v); err != nil {
			return diag.Errorf("error reading include_default_servers: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.InterfaceSelectMethod != nil {
		v := *o.InterfaceSelectMethod

		if err = d.Set("interface_select_method", v); err != nil {
			return diag.Errorf("error reading interface_select_method: %v", err)
		}
	}

	if o.LocalCert != nil {
		v := *o.LocalCert

		if err = d.Set("local_cert", v); err != nil {
			return diag.Errorf("error reading local_cert: %v", err)
		}
	}

	if o.Mode != nil {
		v := *o.Mode

		if err = d.Set("mode", v); err != nil {
			return diag.Errorf("error reading mode: %v", err)
		}
	}

	if o.ScheduleConfigRestore != nil {
		v := *o.ScheduleConfigRestore

		if err = d.Set("schedule_config_restore", v); err != nil {
			return diag.Errorf("error reading schedule_config_restore: %v", err)
		}
	}

	if o.ScheduleScriptRestore != nil {
		v := *o.ScheduleScriptRestore

		if err = d.Set("schedule_script_restore", v); err != nil {
			return diag.Errorf("error reading schedule_script_restore: %v", err)
		}
	}

	if o.SerialNumber != nil {
		v := *o.SerialNumber

		if err = d.Set("serial_number", v); err != nil {
			return diag.Errorf("error reading serial_number: %v", err)
		}
	}

	if o.ServerList != nil {
		if err = d.Set("server_list", flattenSystemCentralManagementServerList(d, o.ServerList, "server_list", sort)); err != nil {
			return diag.Errorf("error reading server_list: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
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

func expandSystemCentralManagementServerList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemCentralManagementServerList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemCentralManagementServerList

	for i := range l {
		tmp := models.SystemCentralManagementServerList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.addr_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AddrType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fqdn", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Fqdn = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.server_address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ServerAddress = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.server_address6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ServerAddress6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.server_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ServerType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemCentralManagement(d *schema.ResourceData, sv string) (*models.SystemCentralManagement, diag.Diagnostics) {
	obj := models.SystemCentralManagement{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("allow_monitor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("allow_monitor", sv)
				diags = append(diags, e)
			}
			obj.AllowMonitor = &v2
		}
	}
	if v1, ok := d.GetOk("allow_push_configuration"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("allow_push_configuration", sv)
				diags = append(diags, e)
			}
			obj.AllowPushConfiguration = &v2
		}
	}
	if v1, ok := d.GetOk("allow_push_firmware"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("allow_push_firmware", sv)
				diags = append(diags, e)
			}
			obj.AllowPushFirmware = &v2
		}
	}
	if v1, ok := d.GetOk("allow_remote_firmware_upgrade"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("allow_remote_firmware_upgrade", sv)
				diags = append(diags, e)
			}
			obj.AllowRemoteFirmwareUpgrade = &v2
		}
	}
	if v1, ok := d.GetOk("ca_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ca_cert", sv)
				diags = append(diags, e)
			}
			obj.CaCert = &v2
		}
	}
	if v1, ok := d.GetOk("enc_algorithm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("enc_algorithm", sv)
				diags = append(diags, e)
			}
			obj.EncAlgorithm = &v2
		}
	}
	if v1, ok := d.GetOk("fmg"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fmg", sv)
				diags = append(diags, e)
			}
			obj.Fmg = &v2
		}
	}
	if v1, ok := d.GetOk("fmg_source_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fmg_source_ip", sv)
				diags = append(diags, e)
			}
			obj.FmgSourceIp = &v2
		}
	}
	if v1, ok := d.GetOk("fmg_source_ip6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fmg_source_ip6", sv)
				diags = append(diags, e)
			}
			obj.FmgSourceIp6 = &v2
		}
	}
	if v1, ok := d.GetOk("fmg_update_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fmg_update_port", sv)
				diags = append(diags, e)
			}
			obj.FmgUpdatePort = &v2
		}
	}
	if v1, ok := d.GetOk("include_default_servers"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("include_default_servers", sv)
				diags = append(diags, e)
			}
			obj.IncludeDefaultServers = &v2
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("interface_select_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("interface_select_method", sv)
				diags = append(diags, e)
			}
			obj.InterfaceSelectMethod = &v2
		}
	}
	if v1, ok := d.GetOk("local_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_cert", sv)
				diags = append(diags, e)
			}
			obj.LocalCert = &v2
		}
	}
	if v1, ok := d.GetOk("mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mode", sv)
				diags = append(diags, e)
			}
			obj.Mode = &v2
		}
	}
	if v1, ok := d.GetOk("schedule_config_restore"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("schedule_config_restore", sv)
				diags = append(diags, e)
			}
			obj.ScheduleConfigRestore = &v2
		}
	}
	if v1, ok := d.GetOk("schedule_script_restore"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("schedule_script_restore", sv)
				diags = append(diags, e)
			}
			obj.ScheduleScriptRestore = &v2
		}
	}
	if v1, ok := d.GetOk("serial_number"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("serial_number", sv)
				diags = append(diags, e)
			}
			obj.SerialNumber = &v2
		}
	}
	if v, ok := d.GetOk("server_list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("server_list", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemCentralManagementServerList(d, v, "server_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ServerList = t
		}
	} else if d.HasChange("server_list") {
		old, new := d.GetChange("server_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ServerList = &[]models.SystemCentralManagementServerList{}
		}
	}
	if v1, ok := d.GetOk("type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("type", sv)
				diags = append(diags, e)
			}
			obj.Type = &v2
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
func getEmptyObjectSystemCentralManagement(d *schema.ResourceData, sv string) (*models.SystemCentralManagement, diag.Diagnostics) {
	obj := models.SystemCentralManagement{}
	diags := diag.Diagnostics{}

	obj.ServerList = &[]models.SystemCentralManagementServerList{}

	return &obj, diags
}
