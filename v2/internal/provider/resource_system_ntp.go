// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure system NTP information.

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

func resourceSystemNtp() *schema.Resource {
	return &schema.Resource{
		Description: "Configure system NTP information.",

		CreateContext: resourceSystemNtpCreate,
		ReadContext:   resourceSystemNtpRead,
		UpdateContext: resourceSystemNtpUpdate,
		DeleteContext: resourceSystemNtpDelete,

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
			"authentication": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable authentication.",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeList,
				Description: "FortiGate interface(s) with NTP server mode enabled. Devices on your network can contact these interfaces for NTP services.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"key": {
				Type: schema.TypeString,

				Description: "Key for authentication.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"key_id": {
				Type: schema.TypeInt,

				Description: "Key ID for authentication.",
				Optional:    true,
				Computed:    true,
			},
			"key_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"MD5", "SHA1"}, false),

				Description: "Key type for authentication (MD5, SHA1).",
				Optional:    true,
				Computed:    true,
			},
			"ntpserver": {
				Type:        schema.TypeList,
				Description: "Configure the FortiGate to connect to any available third-party NTP server.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable MD5(NTPv3)/SHA1(NTPv4) authentication.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "NTP server ID.",
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
						"key": {
							Type: schema.TypeString,

							Description: "Key for MD5(NTPv3)/SHA1(NTPv4) authentication.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"key_id": {
							Type: schema.TypeInt,

							Description: "Key ID for authentication.",
							Optional:    true,
							Computed:    true,
						},
						"ntpv3": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable to use NTPv3 instead of NTPv4.",
							Optional:    true,
							Computed:    true,
						},
						"server": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "IP address or hostname of the NTP Server.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ntpsync": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable setting the FortiGate system time by synchronizing with an NTP Server.",
				Optional:    true,
				Computed:    true,
			},
			"server_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiGate NTP Server Mode. Your FortiGate becomes an NTP server for other devices on your network. The FortiGate relays NTP requests to its configured NTP server.",
				Optional:    true,
				Computed:    true,
			},
			"source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Source IP address for communication to the NTP server.",
				Optional:    true,
				Computed:    true,
			},
			"source_ip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Source IPv6 address for communication to the NTP server.",
				Optional:         true,
				Computed:         true,
			},
			"syncinterval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1440),

				Description: "NTP synchronization interval (1 - 1440 min).",
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"fortiguard", "custom"}, false),

				Description: "Use the FortiGuard NTP server or any other available NTP Server.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemNtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemNtp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemNtp(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemNtp")
	}

	return resourceSystemNtpRead(ctx, d, meta)
}

func resourceSystemNtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemNtp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemNtp(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemNtp resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemNtp")
	}

	return resourceSystemNtpRead(ctx, d, meta)
}

func resourceSystemNtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemNtp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemNtp(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemNtp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemNtp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemNtp resource: %v", err)
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

	diags := refreshObjectSystemNtp(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemNtpInterface(d *schema.ResourceData, v *[]models.SystemNtpInterface, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.InterfaceName; tmp != nil {
				v["interface_name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "interface_name")
	}

	return flat
}

func flattenSystemNtpNtpserver(d *schema.ResourceData, v *[]models.SystemNtpNtpserver, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Authentication; tmp != nil {
				v["authentication"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			if tmp := cfg.InterfaceSelectMethod; tmp != nil {
				v["interface_select_method"] = *tmp
			}

			if tmp := cfg.Key; tmp != nil {
				v["key"] = *tmp
			}

			if tmp := cfg.KeyId; tmp != nil {
				v["key_id"] = *tmp
			}

			if tmp := cfg.Ntpv3; tmp != nil {
				v["ntpv3"] = *tmp
			}

			if tmp := cfg.Server; tmp != nil {
				v["server"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectSystemNtp(d *schema.ResourceData, o *models.SystemNtp, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Authentication != nil {
		v := *o.Authentication

		if err = d.Set("authentication", v); err != nil {
			return diag.Errorf("error reading authentication: %v", err)
		}
	}

	if o.Interface != nil {
		if err = d.Set("interface", flattenSystemNtpInterface(d, o.Interface, "interface", sort)); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.Key != nil {
		v := *o.Key

		if v == "ENC XXXX" {
		} else if err = d.Set("key", v); err != nil {
			return diag.Errorf("error reading key: %v", err)
		}
	}

	if o.KeyId != nil {
		v := *o.KeyId

		if err = d.Set("key_id", v); err != nil {
			return diag.Errorf("error reading key_id: %v", err)
		}
	}

	if o.KeyType != nil {
		v := *o.KeyType

		if err = d.Set("key_type", v); err != nil {
			return diag.Errorf("error reading key_type: %v", err)
		}
	}

	if o.Ntpserver != nil {
		if err = d.Set("ntpserver", flattenSystemNtpNtpserver(d, o.Ntpserver, "ntpserver", sort)); err != nil {
			return diag.Errorf("error reading ntpserver: %v", err)
		}
	}

	if o.Ntpsync != nil {
		v := *o.Ntpsync

		if err = d.Set("ntpsync", v); err != nil {
			return diag.Errorf("error reading ntpsync: %v", err)
		}
	}

	if o.ServerMode != nil {
		v := *o.ServerMode

		if err = d.Set("server_mode", v); err != nil {
			return diag.Errorf("error reading server_mode: %v", err)
		}
	}

	if o.SourceIp != nil {
		v := *o.SourceIp

		if err = d.Set("source_ip", v); err != nil {
			return diag.Errorf("error reading source_ip: %v", err)
		}
	}

	if o.SourceIp6 != nil {
		v := *o.SourceIp6

		if err = d.Set("source_ip6", v); err != nil {
			return diag.Errorf("error reading source_ip6: %v", err)
		}
	}

	if o.Syncinterval != nil {
		v := *o.Syncinterval

		if err = d.Set("syncinterval", v); err != nil {
			return diag.Errorf("error reading syncinterval: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	return nil
}

func expandSystemNtpInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemNtpInterface, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNtpInterface

	for i := range l {
		tmp := models.SystemNtpInterface{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.interface_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InterfaceName = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemNtpNtpserver(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemNtpNtpserver, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemNtpNtpserver

	for i := range l {
		tmp := models.SystemNtpNtpserver{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.authentication", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Authentication = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface_select_method", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InterfaceSelectMethod = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.key", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Key = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.key_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.KeyId = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ntpv3", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ntpv3 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.server", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Server = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemNtp(d *schema.ResourceData, sv string) (*models.SystemNtp, diag.Diagnostics) {
	obj := models.SystemNtp{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("authentication"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authentication", sv)
				diags = append(diags, e)
			}
			obj.Authentication = &v2
		}
	}
	if v, ok := d.GetOk("interface"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("interface", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemNtpInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Interface = t
		}
	} else if d.HasChange("interface") {
		old, new := d.GetChange("interface")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Interface = &[]models.SystemNtpInterface{}
		}
	}
	if v1, ok := d.GetOk("key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("key", sv)
				diags = append(diags, e)
			}
			obj.Key = &v2
		}
	}
	if v1, ok := d.GetOk("key_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("key_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.KeyId = &tmp
		}
	}
	if v1, ok := d.GetOk("key_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("key_type", sv)
				diags = append(diags, e)
			}
			obj.KeyType = &v2
		}
	}
	if v, ok := d.GetOk("ntpserver"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ntpserver", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemNtpNtpserver(d, v, "ntpserver", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Ntpserver = t
		}
	} else if d.HasChange("ntpserver") {
		old, new := d.GetChange("ntpserver")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Ntpserver = &[]models.SystemNtpNtpserver{}
		}
	}
	if v1, ok := d.GetOk("ntpsync"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ntpsync", sv)
				diags = append(diags, e)
			}
			obj.Ntpsync = &v2
		}
	}
	if v1, ok := d.GetOk("server_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_mode", sv)
				diags = append(diags, e)
			}
			obj.ServerMode = &v2
		}
	}
	if v1, ok := d.GetOk("source_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_ip", sv)
				diags = append(diags, e)
			}
			obj.SourceIp = &v2
		}
	}
	if v1, ok := d.GetOk("source_ip6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_ip6", sv)
				diags = append(diags, e)
			}
			obj.SourceIp6 = &v2
		}
	}
	if v1, ok := d.GetOk("syncinterval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("syncinterval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Syncinterval = &tmp
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
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemNtp(d *schema.ResourceData, sv string) (*models.SystemNtp, diag.Diagnostics) {
	obj := models.SystemNtp{}
	diags := diag.Diagnostics{}

	obj.Interface = &[]models.SystemNtpInterface{}
	obj.Ntpserver = &[]models.SystemNtpNtpserver{}

	return &obj, diags
}
