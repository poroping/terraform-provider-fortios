// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure authentication setting.

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

func resourceAuthenticationSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Configure authentication setting.",

		CreateContext: resourceAuthenticationSettingCreate,
		ReadContext:   resourceAuthenticationSettingRead,
		UpdateContext: resourceAuthenticationSettingUpdate,
		DeleteContext: resourceAuthenticationSettingDelete,

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
			"active_auth_scheme": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Active authentication method (scheme name).",
				Optional:    true,
				Computed:    true,
			},
			"auth_https": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable redirecting HTTP user authentication to HTTPS.",
				Optional:    true,
				Computed:    true,
			},
			"captive_portal": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Captive portal host name.",
				Optional:    true,
				Computed:    true,
			},
			"captive_portal_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Captive portal IP address.",
				Optional:    true,
				Computed:    true,
			},
			"captive_portal_ip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Captive portal IPv6 address.",
				Optional:         true,
				Computed:         true,
			},
			"captive_portal_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Captive portal port number (1 - 65535, default = 7830).",
				Optional:    true,
				Computed:    true,
			},
			"captive_portal_ssl_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Captive portal SSL port number (1 - 65535, default = 7831).",
				Optional:    true,
				Computed:    true,
			},
			"captive_portal_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"fqdn", "ip"}, false),

				Description: "Captive portal type.",
				Optional:    true,
				Computed:    true,
			},
			"captive_portal6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "IPv6 captive portal host name.",
				Optional:    true,
				Computed:    true,
			},
			"cert_auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable redirecting certificate authentication to HTTPS portal.",
				Optional:    true,
				Computed:    true,
			},
			"cert_captive_portal": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Certificate captive portal host name.",
				Optional:    true,
				Computed:    true,
			},
			"cert_captive_portal_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Certificate captive portal IP address.",
				Optional:    true,
				Computed:    true,
			},
			"cert_captive_portal_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Certificate captive portal port number (1 - 65535, default = 7832).",
				Optional:    true,
				Computed:    true,
			},
			"cookie_max_age": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 10080),

				Description: "Persistent web portal cookie maximum age in minutes (30 - 10080 (1 week), default = 480 (8 hours)).",
				Optional:    true,
				Computed:    true,
			},
			"cookie_refresh_div": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 4),

				Description: "Refresh rate divider of persistent web portal cookie (default = 2). Refresh value = cookie-max-age/cookie-refresh-div.",
				Optional:    true,
				Computed:    true,
			},
			"dev_range": {
				Type:        schema.TypeList,
				Description: "Address range for the IP based device query.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ip_auth_cookie": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable persistent cookie on IP based web portal authentication (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"persistent_cookie": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable persistent cookie on web portal authentication (default = enable).",
				Optional:    true,
				Computed:    true,
			},
			"sso_auth_scheme": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Single-Sign-On authentication method (scheme name).",
				Optional:    true,
				Computed:    true,
			},
			"update_time": {
				Type: schema.TypeString,

				Description: "Time of the last update.",
				Optional:    true,
				Computed:    true,
			},
			"user_cert_ca": {
				Type:        schema.TypeList,
				Description: "CA certificate used for client certificate verification.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "CA certificate list.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceAuthenticationSettingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectAuthenticationSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateAuthenticationSetting(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("AuthenticationSetting")
	}

	return resourceAuthenticationSettingRead(ctx, d, meta)
}

func resourceAuthenticationSettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectAuthenticationSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateAuthenticationSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating AuthenticationSetting resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("AuthenticationSetting")
	}

	return resourceAuthenticationSettingRead(ctx, d, meta)
}

func resourceAuthenticationSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectAuthenticationSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateAuthenticationSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting AuthenticationSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAuthenticationSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadAuthenticationSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading AuthenticationSetting resource: %v", err)
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

	diags := refreshObjectAuthenticationSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenAuthenticationSettingDevRange(d *schema.ResourceData, v *[]models.AuthenticationSettingDevRange, prefix string, sort bool) interface{} {
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

func flattenAuthenticationSettingUserCertCa(d *schema.ResourceData, v *[]models.AuthenticationSettingUserCertCa, prefix string, sort bool) interface{} {
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

func refreshObjectAuthenticationSetting(d *schema.ResourceData, o *models.AuthenticationSetting, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.ActiveAuthScheme != nil {
		v := *o.ActiveAuthScheme

		if err = d.Set("active_auth_scheme", v); err != nil {
			return diag.Errorf("error reading active_auth_scheme: %v", err)
		}
	}

	if o.AuthHttps != nil {
		v := *o.AuthHttps

		if err = d.Set("auth_https", v); err != nil {
			return diag.Errorf("error reading auth_https: %v", err)
		}
	}

	if o.CaptivePortal != nil {
		v := *o.CaptivePortal

		if err = d.Set("captive_portal", v); err != nil {
			return diag.Errorf("error reading captive_portal: %v", err)
		}
	}

	if o.CaptivePortalIp != nil {
		v := *o.CaptivePortalIp

		if err = d.Set("captive_portal_ip", v); err != nil {
			return diag.Errorf("error reading captive_portal_ip: %v", err)
		}
	}

	if o.CaptivePortalIp6 != nil {
		v := *o.CaptivePortalIp6

		if err = d.Set("captive_portal_ip6", v); err != nil {
			return diag.Errorf("error reading captive_portal_ip6: %v", err)
		}
	}

	if o.CaptivePortalPort != nil {
		v := *o.CaptivePortalPort

		if err = d.Set("captive_portal_port", v); err != nil {
			return diag.Errorf("error reading captive_portal_port: %v", err)
		}
	}

	if o.CaptivePortalSslPort != nil {
		v := *o.CaptivePortalSslPort

		if err = d.Set("captive_portal_ssl_port", v); err != nil {
			return diag.Errorf("error reading captive_portal_ssl_port: %v", err)
		}
	}

	if o.CaptivePortalType != nil {
		v := *o.CaptivePortalType

		if err = d.Set("captive_portal_type", v); err != nil {
			return diag.Errorf("error reading captive_portal_type: %v", err)
		}
	}

	if o.CaptivePortal6 != nil {
		v := *o.CaptivePortal6

		if err = d.Set("captive_portal6", v); err != nil {
			return diag.Errorf("error reading captive_portal6: %v", err)
		}
	}

	if o.CertAuth != nil {
		v := *o.CertAuth

		if err = d.Set("cert_auth", v); err != nil {
			return diag.Errorf("error reading cert_auth: %v", err)
		}
	}

	if o.CertCaptivePortal != nil {
		v := *o.CertCaptivePortal

		if err = d.Set("cert_captive_portal", v); err != nil {
			return diag.Errorf("error reading cert_captive_portal: %v", err)
		}
	}

	if o.CertCaptivePortalIp != nil {
		v := *o.CertCaptivePortalIp

		if err = d.Set("cert_captive_portal_ip", v); err != nil {
			return diag.Errorf("error reading cert_captive_portal_ip: %v", err)
		}
	}

	if o.CertCaptivePortalPort != nil {
		v := *o.CertCaptivePortalPort

		if err = d.Set("cert_captive_portal_port", v); err != nil {
			return diag.Errorf("error reading cert_captive_portal_port: %v", err)
		}
	}

	if o.CookieMaxAge != nil {
		v := *o.CookieMaxAge

		if err = d.Set("cookie_max_age", v); err != nil {
			return diag.Errorf("error reading cookie_max_age: %v", err)
		}
	}

	if o.CookieRefreshDiv != nil {
		v := *o.CookieRefreshDiv

		if err = d.Set("cookie_refresh_div", v); err != nil {
			return diag.Errorf("error reading cookie_refresh_div: %v", err)
		}
	}

	if o.DevRange != nil {
		if err = d.Set("dev_range", flattenAuthenticationSettingDevRange(d, o.DevRange, "dev_range", sort)); err != nil {
			return diag.Errorf("error reading dev_range: %v", err)
		}
	}

	if o.IpAuthCookie != nil {
		v := *o.IpAuthCookie

		if err = d.Set("ip_auth_cookie", v); err != nil {
			return diag.Errorf("error reading ip_auth_cookie: %v", err)
		}
	}

	if o.PersistentCookie != nil {
		v := *o.PersistentCookie

		if err = d.Set("persistent_cookie", v); err != nil {
			return diag.Errorf("error reading persistent_cookie: %v", err)
		}
	}

	if o.SsoAuthScheme != nil {
		v := *o.SsoAuthScheme

		if err = d.Set("sso_auth_scheme", v); err != nil {
			return diag.Errorf("error reading sso_auth_scheme: %v", err)
		}
	}

	if o.UpdateTime != nil {
		v := *o.UpdateTime

		if err = d.Set("update_time", v); err != nil {
			return diag.Errorf("error reading update_time: %v", err)
		}
	}

	if o.UserCertCa != nil {
		if err = d.Set("user_cert_ca", flattenAuthenticationSettingUserCertCa(d, o.UserCertCa, "user_cert_ca", sort)); err != nil {
			return diag.Errorf("error reading user_cert_ca: %v", err)
		}
	}

	return nil
}

func expandAuthenticationSettingDevRange(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.AuthenticationSettingDevRange, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.AuthenticationSettingDevRange

	for i := range l {
		tmp := models.AuthenticationSettingDevRange{}
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

func expandAuthenticationSettingUserCertCa(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.AuthenticationSettingUserCertCa, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.AuthenticationSettingUserCertCa

	for i := range l {
		tmp := models.AuthenticationSettingUserCertCa{}
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

func getObjectAuthenticationSetting(d *schema.ResourceData, sv string) (*models.AuthenticationSetting, diag.Diagnostics) {
	obj := models.AuthenticationSetting{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("active_auth_scheme"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("active_auth_scheme", sv)
				diags = append(diags, e)
			}
			obj.ActiveAuthScheme = &v2
		}
	}
	if v1, ok := d.GetOk("auth_https"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_https", sv)
				diags = append(diags, e)
			}
			obj.AuthHttps = &v2
		}
	}
	if v1, ok := d.GetOk("captive_portal"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("captive_portal", sv)
				diags = append(diags, e)
			}
			obj.CaptivePortal = &v2
		}
	}
	if v1, ok := d.GetOk("captive_portal_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("captive_portal_ip", sv)
				diags = append(diags, e)
			}
			obj.CaptivePortalIp = &v2
		}
	}
	if v1, ok := d.GetOk("captive_portal_ip6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("captive_portal_ip6", sv)
				diags = append(diags, e)
			}
			obj.CaptivePortalIp6 = &v2
		}
	}
	if v1, ok := d.GetOk("captive_portal_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("captive_portal_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CaptivePortalPort = &tmp
		}
	}
	if v1, ok := d.GetOk("captive_portal_ssl_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("captive_portal_ssl_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CaptivePortalSslPort = &tmp
		}
	}
	if v1, ok := d.GetOk("captive_portal_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("captive_portal_type", sv)
				diags = append(diags, e)
			}
			obj.CaptivePortalType = &v2
		}
	}
	if v1, ok := d.GetOk("captive_portal6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("captive_portal6", sv)
				diags = append(diags, e)
			}
			obj.CaptivePortal6 = &v2
		}
	}
	if v1, ok := d.GetOk("cert_auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("cert_auth", sv)
				diags = append(diags, e)
			}
			obj.CertAuth = &v2
		}
	}
	if v1, ok := d.GetOk("cert_captive_portal"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("cert_captive_portal", sv)
				diags = append(diags, e)
			}
			obj.CertCaptivePortal = &v2
		}
	}
	if v1, ok := d.GetOk("cert_captive_portal_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("cert_captive_portal_ip", sv)
				diags = append(diags, e)
			}
			obj.CertCaptivePortalIp = &v2
		}
	}
	if v1, ok := d.GetOk("cert_captive_portal_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("cert_captive_portal_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CertCaptivePortalPort = &tmp
		}
	}
	if v1, ok := d.GetOk("cookie_max_age"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("cookie_max_age", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CookieMaxAge = &tmp
		}
	}
	if v1, ok := d.GetOk("cookie_refresh_div"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("cookie_refresh_div", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CookieRefreshDiv = &tmp
		}
	}
	if v, ok := d.GetOk("dev_range"); ok {
		if !utils.CheckVer(sv, "v7.0.0", "") {
			e := utils.AttributeVersionWarning("dev_range", sv)
			diags = append(diags, e)
		}
		t, err := expandAuthenticationSettingDevRange(d, v, "dev_range", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DevRange = t
		}
	} else if d.HasChange("dev_range") {
		old, new := d.GetChange("dev_range")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DevRange = &[]models.AuthenticationSettingDevRange{}
		}
	}
	if v1, ok := d.GetOk("ip_auth_cookie"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("ip_auth_cookie", sv)
				diags = append(diags, e)
			}
			obj.IpAuthCookie = &v2
		}
	}
	if v1, ok := d.GetOk("persistent_cookie"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("persistent_cookie", sv)
				diags = append(diags, e)
			}
			obj.PersistentCookie = &v2
		}
	}
	if v1, ok := d.GetOk("sso_auth_scheme"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sso_auth_scheme", sv)
				diags = append(diags, e)
			}
			obj.SsoAuthScheme = &v2
		}
	}
	if v1, ok := d.GetOk("update_time"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("update_time", sv)
				diags = append(diags, e)
			}
			obj.UpdateTime = &v2
		}
	}
	if v, ok := d.GetOk("user_cert_ca"); ok {
		if !utils.CheckVer(sv, "v7.0.0", "") {
			e := utils.AttributeVersionWarning("user_cert_ca", sv)
			diags = append(diags, e)
		}
		t, err := expandAuthenticationSettingUserCertCa(d, v, "user_cert_ca", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.UserCertCa = t
		}
	} else if d.HasChange("user_cert_ca") {
		old, new := d.GetChange("user_cert_ca")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.UserCertCa = &[]models.AuthenticationSettingUserCertCa{}
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectAuthenticationSetting(d *schema.ResourceData, sv string) (*models.AuthenticationSetting, diag.Diagnostics) {
	obj := models.AuthenticationSetting{}
	diags := diag.Diagnostics{}

	obj.DevRange = &[]models.AuthenticationSettingDevRange{}
	obj.UserCertCa = &[]models.AuthenticationSettingUserCertCa{}

	return &obj, diags
}
