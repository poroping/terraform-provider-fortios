// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure local users.

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

func resourceUserLocal() *schema.Resource {
	return &schema.Resource{
		Description: "Configure local users.",

		CreateContext: resourceUserLocalCreate,
		ReadContext:   resourceUserLocalRead,
		UpdateContext: resourceUserLocalUpdate,
		DeleteContext: resourceUserLocalDelete,

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
			"auth_concurrent_override": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable overriding the policy-auth-concurrent under config system global.",
				Optional:    true,
				Computed:    true,
			},
			"auth_concurrent_value": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),

				Description: "Maximum number of concurrent logins permitted from the same user.",
				Optional:    true,
				Computed:    true,
			},
			"authtimeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1440),

				Description: "Time in minutes before the authentication timeout for a user is reached.",
				Optional:    true,
				Computed:    true,
			},
			"email_to": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Two-factor recipient's email address.",
				Optional:    true,
				Computed:    true,
			},
			"fortitoken": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 16),

				Description: "Two-factor recipient's FortiToken serial number.",
				Optional:    true,
				Computed:    true,
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "User ID.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"ldap_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of LDAP server with which the user must authenticate.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "User name.",
				ForceNew:    true,
				Required:    true,
			},
			"passwd": {
				Type: schema.TypeString,

				Description: "User's password.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"passwd_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Password policy to apply to this user, as defined in config user password-policy.",
				Optional:    true,
				Computed:    true,
			},
			"passwd_time": {
				Type: schema.TypeString,

				Description: "Time of the last password update.",
				Optional:    true,
				Computed:    true,
			},
			"ppk_identity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "IKEv2 Postquantum Preshared Key Identity.",
				Optional:    true,
				Computed:    true,
			},
			"ppk_secret": {
				Type: schema.TypeString,

				Description: "IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"radius_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of RADIUS server with which the user must authenticate.",
				Optional:    true,
				Computed:    true,
			},
			"sms_custom_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Two-factor recipient's SMS server.",
				Optional:    true,
				Computed:    true,
			},
			"sms_phone": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Two-factor recipient's mobile phone number.",
				Optional:    true,
				Computed:    true,
			},
			"sms_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"fortiguard", "custom"}, false),

				Description: "Send SMS through FortiGuard or other external server.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable allowing the local user to authenticate with the FortiGate unit.",
				Optional:    true,
				Computed:    true,
			},
			"tacacs_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of TACACS+ server with which the user must authenticate.",
				Optional:    true,
				Computed:    true,
			},
			"two_factor": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "fortitoken", "fortitoken-cloud", "email", "sms"}, false),

				Description: "Enable/disable two-factor authentication.",
				Optional:    true,
				Computed:    true,
			},
			"two_factor_authentication": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"fortitoken", "email", "sms"}, false),

				Description: "Authentication method by FortiToken Cloud.",
				Optional:    true,
				Computed:    true,
			},
			"two_factor_notification": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"email", "sms"}, false),

				Description: "Notification method for user activation by FortiToken Cloud.",
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"password", "radius", "tacacs+", "ldap"}, false),

				Description: "Authentication method.",
				Optional:    true,
				Computed:    true,
			},
			"username_case_insensitivity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent).",
				Optional:    true,
				Computed:    true,
			},
			"username_case_sensitivity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent).",
				Optional:    true,
				Computed:    true,
			},
			"username_sensitivity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable case and accent sensitivity when performing username matching (accents are stripped and case is ignored when disabled).",
				Optional:    true,
				Computed:    true,
			},
			"workstation": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of the remote user workstation, if you want to limit the user to authenticate only from a particular workstation.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceUserLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating UserLocal resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectUserLocal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateUserLocal(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserLocal")
	}

	return resourceUserLocalRead(ctx, d, meta)
}

func resourceUserLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectUserLocal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateUserLocal(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating UserLocal resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserLocal")
	}

	return resourceUserLocalRead(ctx, d, meta)
}

func resourceUserLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteUserLocal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting UserLocal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserLocal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserLocal resource: %v", err)
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

	diags := refreshObjectUserLocal(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectUserLocal(d *schema.ResourceData, o *models.UserLocal, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AuthConcurrentOverride != nil {
		v := *o.AuthConcurrentOverride

		if err = d.Set("auth_concurrent_override", v); err != nil {
			return diag.Errorf("error reading auth_concurrent_override: %v", err)
		}
	}

	if o.AuthConcurrentValue != nil {
		v := *o.AuthConcurrentValue

		if err = d.Set("auth_concurrent_value", v); err != nil {
			return diag.Errorf("error reading auth_concurrent_value: %v", err)
		}
	}

	if o.Authtimeout != nil {
		v := *o.Authtimeout

		if err = d.Set("authtimeout", v); err != nil {
			return diag.Errorf("error reading authtimeout: %v", err)
		}
	}

	if o.EmailTo != nil {
		v := *o.EmailTo

		if err = d.Set("email_to", v); err != nil {
			return diag.Errorf("error reading email_to: %v", err)
		}
	}

	if o.Fortitoken != nil {
		v := *o.Fortitoken

		if err = d.Set("fortitoken", v); err != nil {
			return diag.Errorf("error reading fortitoken: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.LdapServer != nil {
		v := *o.LdapServer

		if err = d.Set("ldap_server", v); err != nil {
			return diag.Errorf("error reading ldap_server: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Passwd != nil {
		v := *o.Passwd

		if v == "ENC XXXX" {
		} else if err = d.Set("passwd", v); err != nil {
			return diag.Errorf("error reading passwd: %v", err)
		}
	}

	if o.PasswdPolicy != nil {
		v := *o.PasswdPolicy

		if err = d.Set("passwd_policy", v); err != nil {
			return diag.Errorf("error reading passwd_policy: %v", err)
		}
	}

	if o.PasswdTime != nil {
		v := *o.PasswdTime

		if err = d.Set("passwd_time", v); err != nil {
			return diag.Errorf("error reading passwd_time: %v", err)
		}
	}

	if o.PpkIdentity != nil {
		v := *o.PpkIdentity

		if err = d.Set("ppk_identity", v); err != nil {
			return diag.Errorf("error reading ppk_identity: %v", err)
		}
	}

	if o.PpkSecret != nil {
		v := *o.PpkSecret

		if v == "ENC XXXX" {
		} else if err = d.Set("ppk_secret", v); err != nil {
			return diag.Errorf("error reading ppk_secret: %v", err)
		}
	}

	if o.RadiusServer != nil {
		v := *o.RadiusServer

		if err = d.Set("radius_server", v); err != nil {
			return diag.Errorf("error reading radius_server: %v", err)
		}
	}

	if o.SmsCustomServer != nil {
		v := *o.SmsCustomServer

		if err = d.Set("sms_custom_server", v); err != nil {
			return diag.Errorf("error reading sms_custom_server: %v", err)
		}
	}

	if o.SmsPhone != nil {
		v := *o.SmsPhone

		if err = d.Set("sms_phone", v); err != nil {
			return diag.Errorf("error reading sms_phone: %v", err)
		}
	}

	if o.SmsServer != nil {
		v := *o.SmsServer

		if err = d.Set("sms_server", v); err != nil {
			return diag.Errorf("error reading sms_server: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.TacacsServer != nil {
		v := *o.TacacsServer

		if err = d.Set("tacacs_server", v); err != nil {
			return diag.Errorf("error reading tacacs_server: %v", err)
		}
	}

	if o.TwoFactor != nil {
		v := *o.TwoFactor

		if err = d.Set("two_factor", v); err != nil {
			return diag.Errorf("error reading two_factor: %v", err)
		}
	}

	if o.TwoFactorAuthentication != nil {
		v := *o.TwoFactorAuthentication

		if err = d.Set("two_factor_authentication", v); err != nil {
			return diag.Errorf("error reading two_factor_authentication: %v", err)
		}
	}

	if o.TwoFactorNotification != nil {
		v := *o.TwoFactorNotification

		if err = d.Set("two_factor_notification", v); err != nil {
			return diag.Errorf("error reading two_factor_notification: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	if o.UsernameCaseInsensitivity != nil {
		v := *o.UsernameCaseInsensitivity

		if err = d.Set("username_case_insensitivity", v); err != nil {
			return diag.Errorf("error reading username_case_insensitivity: %v", err)
		}
	}

	if o.UsernameCaseSensitivity != nil {
		v := *o.UsernameCaseSensitivity

		if err = d.Set("username_case_sensitivity", v); err != nil {
			return diag.Errorf("error reading username_case_sensitivity: %v", err)
		}
	}

	if o.UsernameSensitivity != nil {
		v := *o.UsernameSensitivity

		if err = d.Set("username_sensitivity", v); err != nil {
			return diag.Errorf("error reading username_sensitivity: %v", err)
		}
	}

	if o.Workstation != nil {
		v := *o.Workstation

		if err = d.Set("workstation", v); err != nil {
			return diag.Errorf("error reading workstation: %v", err)
		}
	}

	return nil
}

func getObjectUserLocal(d *schema.ResourceData, sv string) (*models.UserLocal, diag.Diagnostics) {
	obj := models.UserLocal{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("auth_concurrent_override"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_concurrent_override", sv)
				diags = append(diags, e)
			}
			obj.AuthConcurrentOverride = &v2
		}
	}
	if v1, ok := d.GetOk("auth_concurrent_value"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_concurrent_value", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AuthConcurrentValue = &tmp
		}
	}
	if v1, ok := d.GetOk("authtimeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authtimeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Authtimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("email_to"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("email_to", sv)
				diags = append(diags, e)
			}
			obj.EmailTo = &v2
		}
	}
	if v1, ok := d.GetOk("fortitoken"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fortitoken", sv)
				diags = append(diags, e)
			}
			obj.Fortitoken = &v2
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
	if v1, ok := d.GetOk("ldap_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ldap_server", sv)
				diags = append(diags, e)
			}
			obj.LdapServer = &v2
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
	if v1, ok := d.GetOk("passwd"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("passwd", sv)
				diags = append(diags, e)
			}
			obj.Passwd = &v2
		}
	}
	if v1, ok := d.GetOk("passwd_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("passwd_policy", sv)
				diags = append(diags, e)
			}
			obj.PasswdPolicy = &v2
		}
	}
	if v1, ok := d.GetOk("passwd_time"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("passwd_time", sv)
				diags = append(diags, e)
			}
			obj.PasswdTime = &v2
		}
	}
	if v1, ok := d.GetOk("ppk_identity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ppk_identity", sv)
				diags = append(diags, e)
			}
			obj.PpkIdentity = &v2
		}
	}
	if v1, ok := d.GetOk("ppk_secret"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ppk_secret", sv)
				diags = append(diags, e)
			}
			obj.PpkSecret = &v2
		}
	}
	if v1, ok := d.GetOk("radius_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("radius_server", sv)
				diags = append(diags, e)
			}
			obj.RadiusServer = &v2
		}
	}
	if v1, ok := d.GetOk("sms_custom_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sms_custom_server", sv)
				diags = append(diags, e)
			}
			obj.SmsCustomServer = &v2
		}
	}
	if v1, ok := d.GetOk("sms_phone"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sms_phone", sv)
				diags = append(diags, e)
			}
			obj.SmsPhone = &v2
		}
	}
	if v1, ok := d.GetOk("sms_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sms_server", sv)
				diags = append(diags, e)
			}
			obj.SmsServer = &v2
		}
	}
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	if v1, ok := d.GetOk("tacacs_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tacacs_server", sv)
				diags = append(diags, e)
			}
			obj.TacacsServer = &v2
		}
	}
	if v1, ok := d.GetOk("two_factor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("two_factor", sv)
				diags = append(diags, e)
			}
			obj.TwoFactor = &v2
		}
	}
	if v1, ok := d.GetOk("two_factor_authentication"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("two_factor_authentication", sv)
				diags = append(diags, e)
			}
			obj.TwoFactorAuthentication = &v2
		}
	}
	if v1, ok := d.GetOk("two_factor_notification"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("two_factor_notification", sv)
				diags = append(diags, e)
			}
			obj.TwoFactorNotification = &v2
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
	if v1, ok := d.GetOk("username_case_insensitivity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "v6.4.2") {
				e := utils.AttributeVersionWarning("username_case_insensitivity", sv)
				diags = append(diags, e)
			}
			obj.UsernameCaseInsensitivity = &v2
		}
	}
	if v1, ok := d.GetOk("username_case_sensitivity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("username_case_sensitivity", sv)
				diags = append(diags, e)
			}
			obj.UsernameCaseSensitivity = &v2
		}
	}
	if v1, ok := d.GetOk("username_sensitivity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.7", "v7.0.0") {
				e := utils.AttributeVersionWarning("username_sensitivity", sv)
				diags = append(diags, e)
			}
			obj.UsernameSensitivity = &v2
		}
	}
	if v1, ok := d.GetOk("workstation"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("workstation", sv)
				diags = append(diags, e)
			}
			obj.Workstation = &v2
		}
	}
	return &obj, diags
}
