// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure user authentication setting.

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

func resourceUserSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Configure user authentication setting.",

		CreateContext: resourceUserSettingCreate,
		ReadContext:   resourceUserSettingRead,
		UpdateContext: resourceUserSettingUpdate,
		DeleteContext: resourceUserSettingDelete,

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
			"auth_blackout_time": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3600),

				Description: "Time in seconds an IP address is denied access after failing to authenticate five times within one minute.",
				Optional:    true,
				Computed:    true,
			},
			"auth_ca_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "HTTPS CA certificate for policy authentication.",
				Optional:    true,
				Computed:    true,
			},
			"auth_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "HTTPS server certificate for policy authentication.",
				Optional:    true,
				Computed:    true,
			},
			"auth_http_basic": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of HTTP basic authentication for identity-based firewall policies.",
				Optional:    true,
				Computed:    true,
			},
			"auth_invalid_max": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),

				Description: "Maximum number of failed authentication attempts before the user is blocked.",
				Optional:    true,
				Computed:    true,
			},
			"auth_lockout_duration": {
				Type: schema.TypeInt,

				Description: "Lockout period in seconds after too many login failures.",
				Optional:    true,
				Computed:    true,
			},
			"auth_lockout_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),

				Description: "Maximum number of failed login attempts before login lockout is triggered.",
				Optional:    true,
				Computed:    true,
			},
			"auth_on_demand": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"always", "implicitly"}, false),

				Description: "Always/implicitly trigger firewall authentication on demand.",
				Optional:    true,
				Computed:    true,
			},
			"auth_portal_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),

				Description: "Time in minutes before captive portal user have to re-authenticate (1 - 30 min, default 3 min).",
				Optional:    true,
				Computed:    true,
			},
			"auth_ports": {
				Type:        schema.TypeList,
				Description: "Set up non-standard ports for authentication with HTTP, HTTPS, FTP, and TELNET.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Non-standard port for firewall user authentication.",
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"http", "https", "ftp", "telnet"}, false),

							Description: "Service type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"auth_secure_http": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable redirecting HTTP user authentication to more secure HTTPS.",
				Optional:    true,
				Computed:    true,
			},
			"auth_src_mac": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable source MAC for user identity.",
				Optional:    true,
				Computed:    true,
			},
			"auth_ssl_allow_renegotiation": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Allow/forbid SSL re-negotiation for HTTPS authentication.",
				Optional:    true,
				Computed:    true,
			},
			"auth_ssl_max_proto_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"sslv3", "tlsv1", "tlsv1-1", "tlsv1-2", "tlsv1-3"}, false),

				Description: "Maximum supported protocol version for SSL/TLS connections (default is no limit).",
				Optional:    true,
				Computed:    true,
			},
			"auth_ssl_min_proto_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"default", "SSLv3", "TLSv1", "TLSv1-1", "TLSv1-2"}, false),

				Description: "Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).",
				Optional:    true,
				Computed:    true,
			},
			"auth_ssl_sigalgs": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"no-rsa-pss", "all"}, false),

				Description: "Set signature algorithms related to HTTPS authentication (affects TLS version <= 1.2 only, default is to enable all).",
				Optional:    true,
				Computed:    true,
			},
			"auth_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1440),

				Description: "Time in minutes before the firewall user authentication timeout requires the user to re-authenticate.",
				Optional:    true,
				Computed:    true,
			},
			"auth_timeout_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"idle-timeout", "hard-timeout", "new-session"}, false),

				Description: "Control if authenticated users have to login again after a hard timeout, after an idle timeout, or after a session timeout.",
				Optional:    true,
				Computed:    true,
			},
			"auth_type": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Supported firewall policy authentication protocols/methods.",
				Optional:         true,
				Computed:         true,
			},
			"per_policy_disclaimer": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable per policy disclaimer.",
				Optional:    true,
				Computed:    true,
			},
			"radius_ses_timeout_act": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"hard-timeout", "ignore-timeout"}, false),

				Description: "Set the RADIUS session timeout to a hard timeout or to ignore RADIUS server session timeouts.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceUserSettingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectUserSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateUserSetting(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserSetting")
	}

	return resourceUserSettingRead(ctx, d, meta)
}

func resourceUserSettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectUserSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateUserSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating UserSetting resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserSetting")
	}

	return resourceUserSettingRead(ctx, d, meta)
}

func resourceUserSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectUserSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateUserSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting UserSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserSetting resource: %v", err)
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

	diags := refreshObjectUserSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenUserSettingAuthPorts(v *[]models.UserSettingAuthPorts, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Port; tmp != nil {
				v["port"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectUserSetting(d *schema.ResourceData, o *models.UserSetting, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AuthBlackoutTime != nil {
		v := *o.AuthBlackoutTime

		if err = d.Set("auth_blackout_time", v); err != nil {
			return diag.Errorf("error reading auth_blackout_time: %v", err)
		}
	}

	if o.AuthCaCert != nil {
		v := *o.AuthCaCert

		if err = d.Set("auth_ca_cert", v); err != nil {
			return diag.Errorf("error reading auth_ca_cert: %v", err)
		}
	}

	if o.AuthCert != nil {
		v := *o.AuthCert

		if err = d.Set("auth_cert", v); err != nil {
			return diag.Errorf("error reading auth_cert: %v", err)
		}
	}

	if o.AuthHttpBasic != nil {
		v := *o.AuthHttpBasic

		if err = d.Set("auth_http_basic", v); err != nil {
			return diag.Errorf("error reading auth_http_basic: %v", err)
		}
	}

	if o.AuthInvalidMax != nil {
		v := *o.AuthInvalidMax

		if err = d.Set("auth_invalid_max", v); err != nil {
			return diag.Errorf("error reading auth_invalid_max: %v", err)
		}
	}

	if o.AuthLockoutDuration != nil {
		v := *o.AuthLockoutDuration

		if err = d.Set("auth_lockout_duration", v); err != nil {
			return diag.Errorf("error reading auth_lockout_duration: %v", err)
		}
	}

	if o.AuthLockoutThreshold != nil {
		v := *o.AuthLockoutThreshold

		if err = d.Set("auth_lockout_threshold", v); err != nil {
			return diag.Errorf("error reading auth_lockout_threshold: %v", err)
		}
	}

	if o.AuthOnDemand != nil {
		v := *o.AuthOnDemand

		if err = d.Set("auth_on_demand", v); err != nil {
			return diag.Errorf("error reading auth_on_demand: %v", err)
		}
	}

	if o.AuthPortalTimeout != nil {
		v := *o.AuthPortalTimeout

		if err = d.Set("auth_portal_timeout", v); err != nil {
			return diag.Errorf("error reading auth_portal_timeout: %v", err)
		}
	}

	if o.AuthPorts != nil {
		if err = d.Set("auth_ports", flattenUserSettingAuthPorts(o.AuthPorts, sort)); err != nil {
			return diag.Errorf("error reading auth_ports: %v", err)
		}
	}

	if o.AuthSecureHttp != nil {
		v := *o.AuthSecureHttp

		if err = d.Set("auth_secure_http", v); err != nil {
			return diag.Errorf("error reading auth_secure_http: %v", err)
		}
	}

	if o.AuthSrcMac != nil {
		v := *o.AuthSrcMac

		if err = d.Set("auth_src_mac", v); err != nil {
			return diag.Errorf("error reading auth_src_mac: %v", err)
		}
	}

	if o.AuthSslAllowRenegotiation != nil {
		v := *o.AuthSslAllowRenegotiation

		if err = d.Set("auth_ssl_allow_renegotiation", v); err != nil {
			return diag.Errorf("error reading auth_ssl_allow_renegotiation: %v", err)
		}
	}

	if o.AuthSslMaxProtoVersion != nil {
		v := *o.AuthSslMaxProtoVersion

		if err = d.Set("auth_ssl_max_proto_version", v); err != nil {
			return diag.Errorf("error reading auth_ssl_max_proto_version: %v", err)
		}
	}

	if o.AuthSslMinProtoVersion != nil {
		v := *o.AuthSslMinProtoVersion

		if err = d.Set("auth_ssl_min_proto_version", v); err != nil {
			return diag.Errorf("error reading auth_ssl_min_proto_version: %v", err)
		}
	}

	if o.AuthSslSigalgs != nil {
		v := *o.AuthSslSigalgs

		if err = d.Set("auth_ssl_sigalgs", v); err != nil {
			return diag.Errorf("error reading auth_ssl_sigalgs: %v", err)
		}
	}

	if o.AuthTimeout != nil {
		v := *o.AuthTimeout

		if err = d.Set("auth_timeout", v); err != nil {
			return diag.Errorf("error reading auth_timeout: %v", err)
		}
	}

	if o.AuthTimeoutType != nil {
		v := *o.AuthTimeoutType

		if err = d.Set("auth_timeout_type", v); err != nil {
			return diag.Errorf("error reading auth_timeout_type: %v", err)
		}
	}

	if o.AuthType != nil {
		v := *o.AuthType

		if err = d.Set("auth_type", v); err != nil {
			return diag.Errorf("error reading auth_type: %v", err)
		}
	}

	if o.PerPolicyDisclaimer != nil {
		v := *o.PerPolicyDisclaimer

		if err = d.Set("per_policy_disclaimer", v); err != nil {
			return diag.Errorf("error reading per_policy_disclaimer: %v", err)
		}
	}

	if o.RadiusSesTimeoutAct != nil {
		v := *o.RadiusSesTimeoutAct

		if err = d.Set("radius_ses_timeout_act", v); err != nil {
			return diag.Errorf("error reading radius_ses_timeout_act: %v", err)
		}
	}

	return nil
}

func expandUserSettingAuthPorts(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.UserSettingAuthPorts, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.UserSettingAuthPorts

	for i := range l {
		tmp := models.UserSettingAuthPorts{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Port = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectUserSetting(d *schema.ResourceData, sv string) (*models.UserSetting, diag.Diagnostics) {
	obj := models.UserSetting{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("auth_blackout_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_blackout_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AuthBlackoutTime = &tmp
		}
	}
	if v1, ok := d.GetOk("auth_ca_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_ca_cert", sv)
				diags = append(diags, e)
			}
			obj.AuthCaCert = &v2
		}
	}
	if v1, ok := d.GetOk("auth_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_cert", sv)
				diags = append(diags, e)
			}
			obj.AuthCert = &v2
		}
	}
	if v1, ok := d.GetOk("auth_http_basic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_http_basic", sv)
				diags = append(diags, e)
			}
			obj.AuthHttpBasic = &v2
		}
	}
	if v1, ok := d.GetOk("auth_invalid_max"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_invalid_max", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AuthInvalidMax = &tmp
		}
	}
	if v1, ok := d.GetOk("auth_lockout_duration"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_lockout_duration", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AuthLockoutDuration = &tmp
		}
	}
	if v1, ok := d.GetOk("auth_lockout_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_lockout_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AuthLockoutThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("auth_on_demand"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_on_demand", sv)
				diags = append(diags, e)
			}
			obj.AuthOnDemand = &v2
		}
	}
	if v1, ok := d.GetOk("auth_portal_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_portal_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AuthPortalTimeout = &tmp
		}
	}
	if v, ok := d.GetOk("auth_ports"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("auth_ports", sv)
			diags = append(diags, e)
		}
		t, err := expandUserSettingAuthPorts(d, v, "auth_ports", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.AuthPorts = t
		}
	} else if d.HasChange("auth_ports") {
		old, new := d.GetChange("auth_ports")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.AuthPorts = &[]models.UserSettingAuthPorts{}
		}
	}
	if v1, ok := d.GetOk("auth_secure_http"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_secure_http", sv)
				diags = append(diags, e)
			}
			obj.AuthSecureHttp = &v2
		}
	}
	if v1, ok := d.GetOk("auth_src_mac"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_src_mac", sv)
				diags = append(diags, e)
			}
			obj.AuthSrcMac = &v2
		}
	}
	if v1, ok := d.GetOk("auth_ssl_allow_renegotiation"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_ssl_allow_renegotiation", sv)
				diags = append(diags, e)
			}
			obj.AuthSslAllowRenegotiation = &v2
		}
	}
	if v1, ok := d.GetOk("auth_ssl_max_proto_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("auth_ssl_max_proto_version", sv)
				diags = append(diags, e)
			}
			obj.AuthSslMaxProtoVersion = &v2
		}
	}
	if v1, ok := d.GetOk("auth_ssl_min_proto_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_ssl_min_proto_version", sv)
				diags = append(diags, e)
			}
			obj.AuthSslMinProtoVersion = &v2
		}
	}
	if v1, ok := d.GetOk("auth_ssl_sigalgs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("auth_ssl_sigalgs", sv)
				diags = append(diags, e)
			}
			obj.AuthSslSigalgs = &v2
		}
	}
	if v1, ok := d.GetOk("auth_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AuthTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("auth_timeout_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_timeout_type", sv)
				diags = append(diags, e)
			}
			obj.AuthTimeoutType = &v2
		}
	}
	if v1, ok := d.GetOk("auth_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_type", sv)
				diags = append(diags, e)
			}
			obj.AuthType = &v2
		}
	}
	if v1, ok := d.GetOk("per_policy_disclaimer"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("per_policy_disclaimer", sv)
				diags = append(diags, e)
			}
			obj.PerPolicyDisclaimer = &v2
		}
	}
	if v1, ok := d.GetOk("radius_ses_timeout_act"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("radius_ses_timeout_act", sv)
				diags = append(diags, e)
			}
			obj.RadiusSesTimeoutAct = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectUserSetting(d *schema.ResourceData, sv string) (*models.UserSetting, diag.Diagnostics) {
	obj := models.UserSetting{}
	diags := diag.Diagnostics{}

	obj.AuthPorts = &[]models.UserSettingAuthPorts{}

	return &obj, diags
}
